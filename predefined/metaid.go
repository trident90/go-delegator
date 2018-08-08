package predefined

import (
	"bytes"
	"math/big"

	proxyCommon "bitbucket.org/coinplugin/proxy/common"
	"bitbucket.org/coinplugin/proxy/crypto"
	"bitbucket.org/coinplugin/proxy/crypto/merkletree"
	"bitbucket.org/coinplugin/proxy/ipfs"
	"bitbucket.org/coinplugin/proxy/json"
	"bitbucket.org/coinplugin/proxy/log"
	"bitbucket.org/coinplugin/proxy/predefined/sc/identitymanager"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
)

var (
	zero = big.NewInt(0)
)

func makeErrorResponse(err Error) *json.RPCError {
	return &json.RPCError{
		Code:    err.ErrorCode(),
		Message: err.Error(),
	}
}

func getAddress(req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	requestTmpID := proxyCommon.RandomUint64()
	log.Debugd(requestTmpID, " Call getAddress Function")
	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	//1. Check paramter format
	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	reqParam := tmpParams.(metaIDGetAddressParams)

	log.Debugfd(requestTmpID, "parameter[MetaId] : %v", reqParam.MetaID)
	log.Debugd(requestTmpID, "PASS - 01. Check Parameter")

	address := &common.Address{}
	address, err := identitymanager.CallOwnerOf(reqParam.MetaID)
	if err != nil {
		log.Errorfd(requestTmpID, "CallOwnerOf Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	if address != nil {
		resp.Result = address.String()
	}
	return
}

func verifySignature(requestTmpID uint64, msg string, sig string, address *common.Address) (*common.Address, Error) {
	signedAddress, err := crypto.EcRecover(msg, sig)
	if err != nil {
		log.Errorfd(requestTmpID, "EcRecover Error: %v", err)
		return nil, &internalError{"Failed to EcRecover"}
	}
	log.Debugfd(requestTmpID, "PASS - Check Signature signAddress : %v", signedAddress.String())
	//비교할 address가 없을 경우, signedAddress만 리턴
	if address != nil {
		if !bytes.Equal(signedAddress.Bytes(), address.Bytes()) {
			log.Debugfd(requestTmpID, "Error  :Sign Error %v / %v ", address.String(), signedAddress.String())
			return &signedAddress, &invalidSignatureError{"Failed to verify signature"}
		}
	}
	log.Debugd(requestTmpID, "PASS - signedAddress is valid")
	return &signedAddress, nil
}

func registerMetaID(req json.RPCRequest) (resp json.RPCResponse, errRet error) {

	requestTmpID := proxyCommon.RandomUint64()
	log.Debugd(requestTmpID, " Call registerMetalID Function")

	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	//1. Check paramter format
	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	reqParam := tmpParams.(metaIDRegisterParams)
	log.Debugfd(requestTmpID, "parameter[Address] : %v", reqParam.Address.String())
	log.Debugfd(requestTmpID, "parameter[MetaId] : %v", reqParam.MetaID)
	log.Debugfd(requestTmpID, "parameter[Sig] : %v", reqParam.Signature)
	log.Debugfd(requestTmpID, "parameter[UserHashs] : %v", reqParam.UserHashs)

	log.Debugd(requestTmpID, "PASS - 01. Check Parameter")
	//2. check Signature
	//  if  address == ecrecover()  then Pass else Error
	// signedAddress, err := crypto.EcRecover(reqParam.MetaID.String(), reqParam.Signature.String())
	// if err != nil {
	// 	log.Errorfd(requestTmpID, "EcRecover Error: %v", err)
	// 	errObj := &internalError{"Failed to EcRecover"}
	// 	resp.Error = makeErrorResponse(errObj)
	// 	return
	// }
	// log.Debugfd(requestTmpID, "PASS - 02-1. Check Signature signAddress : %v", signedAddress.String())

	// if !bytes.Equal(signedAddress.Bytes(), reqParam.Address.Bytes()) {
	// 	log.Debugfd(requestTmpID, "Error  :Sign Error %v / %v ", reqParam.Address.String(), signedAddress.String())
	// 	//TODO:  ecRecover 에러와 address다른 부분 에러 분리
	// 	errObj := &invalidSignatureError{"Failed to verify signature"}
	// 	resp.Error = makeErrorResponse(errObj)
	// 	return
	// }

	_, errObj = verifySignature(requestTmpID, reqParam.MetaID.String(), reqParam.Signature.String(), &reqParam.Address)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(requestTmpID, "PASS - 02-2. signedAddress is valid")

	//3. Check Meta ID  by make merkle tree root.
	//   if metaId == root  then Pass  else  Error

	var usrHashs []merkletree.Content

	for _, userHash := range reqParam.UserHashs {
		usrHashs = append(usrHashs, merkletree.HashContent{H: userHash})
	}

	//Create a new Merkle Tree from the list of Content
	tree, _ := merkletree.NewTree(usrHashs)

	//Get the Merkle Root of the tree
	mr := tree.MerkleRoot()
	rootStr := hexutil.Encode(mr)

	log.Debugfd(requestTmpID, "PASS 03-1. Make Merkle Root: %v ", rootStr)

	if !bytes.Equal(mr, reqParam.MetaID) {
		log.Debugfd(requestTmpID, "Error - Invalid MetaID (MetaID,ROOT): %v,%v", reqParam.MetaID, rootStr)
		errObj := &invalidMetaIDError{"Failed to verify metaID"}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	log.Debugfd(requestTmpID, "PASS 03-2. MetaID valid : %v ", reqParam.MetaID)

	//4. Reqest Get UserAddress for Meta ID
	//  if GetAddress == nil  then Pass  else Error
	address := &common.Address{}
	address, err := identitymanager.CallOwnerOf(reqParam.MetaID)
	if err != nil {
		log.Errorfd(requestTmpID, "CallOwnerOf Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(requestTmpID, "PASS 04-1. Get Address OwnerOf MetaID")

	if address != nil {

		log.Debugf("Error-Already registered MetaID: %v", address.String())
		errObj := &alreadyExistsAddressError{"already exists address for metaID"} //"Cannot register Meta ID"
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(requestTmpID, "PASS 04-2. Get Address OwnerOf MetaID : nil")

	//5.get MetaID Count for Address
	count, err := identitymanager.CallBalanceOf(reqParam.Address)
	if err != nil {
		log.Errorfd(requestTmpID, "CallBalanceOf Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	if zero.Cmp(count) != 0 {
		log.Errord(requestTmpID, "already used address for registering metaID")
		errObj := &alreadyUsedAddressError{"already used address"}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	//6. Send Transaction for Calling SC Function [createNewMetaID]
	//identitymanager.CallCreateMetaID()

	trx, err := identitymanager.CallCreateMetaID(reqParam.MetaID, reqParam.Signature, reqParam.Address)
	if err != nil {
		log.Errorfd(requestTmpID, "CallCreateMetaID Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugfd(requestTmpID, "PASS 05. Call CreateMetaID : %v", trx.Hash().String())

	//  return txid
	resp.Result = trx.Hash().String()
	return
}

func updateMetaID(req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	requestTmpID := proxyCommon.RandomUint64()
	log.Debugd(requestTmpID, "Call updateMetaID Function")

	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	//1. Check paramter format
	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	reqParam := tmpParams.(metaIDUpdateParams)
	log.Debugfd(requestTmpID, "parameter[Address] : %v", reqParam.Address.String())
	log.Debugfd(requestTmpID, "parameter[NewMetaId] : %v", reqParam.NewMetaID)
	log.Debugfd(requestTmpID, "parameter[OldMetaId] : %v", reqParam.OldMetaID)
	log.Debugfd(requestTmpID, "parameter[Sig] : %v", reqParam.Signature)
	log.Debugfd(requestTmpID, "parameter[UserHashs] : %v", reqParam.UserHashs)
	log.Debugd(requestTmpID, "PASS - 01.Check Parameter")
	//2. check Signature
	//  if  address == ecrecover()  then Pass else Error
	// signedAddress, err := crypto.EcRecover(reqParam.NewMetaID.String(), reqParam.Signature.String())
	// if err != nil {
	// 	log.Errorfd(requestTmpID, "EcRecover Error : %v", err)
	// 	errObj := &internalError{"Failed to EcRecover"}
	// 	resp.Error = makeErrorResponse(errObj)
	// 	return
	// }
	// log.Debugfd(requestTmpID, "PASS - 02-1.EcRecover(Signd Address): %v", signedAddress.String())

	// if !bytes.Equal(signedAddress.Bytes(), reqParam.Address.Bytes()) {
	// 	log.Debugfd(requestTmpID, "Error  :Sign Error %v / %v ", reqParam.Address.String(), signedAddress.String())

	// 	errObj := &invalidSignatureError{"Failed to verify signature"}
	// 	resp.Error = makeErrorResponse(errObj)
	// 	return
	// }

	_, errObj = verifySignature(requestTmpID, reqParam.NewMetaID.String(), reqParam.Signature.String(), &reqParam.Address)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(requestTmpID, "PASS - 02-2. signd Address is valid")

	//3. Check New Meta ID  by make merkle tree root.
	//   if newMetaID == root  then Pass  else  Error

	var usrHashs []merkletree.Content

	for _, userHash := range reqParam.UserHashs {

		usrHashs = append(usrHashs, merkletree.HashContent{H: userHash})

	}

	tree, _ := merkletree.NewTree(usrHashs)

	//Get the Merkle Root of the tree
	mr := tree.MerkleRoot()

	root := hexutil.Bytes(mr)

	log.Debugfd(requestTmpID, "PASS - 03-1. Make Merkle Root: %v", root)

	if !bytes.Equal(root, reqParam.NewMetaID) {

		log.Debugfd(requestTmpID, "Error - Invalid MetaID (MetaID, ROOT) : %v, %x ", reqParam.NewMetaID, root)
		errObj := &invalidMetaIDError{"Failed to verify metaID"}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	log.Debugfd(requestTmpID, "PASS - 03-2. MetaID valid : %v", reqParam.NewMetaID)

	//4. Reqest Get UserAddress for old Meta ID
	//  if GetAddress ==  reqParam.Address   then pass
	//  else Error
	findAddress := &common.Address{}
	findAddress, err := identitymanager.CallOwnerOf(reqParam.OldMetaID)
	if err != nil {
		log.Errorfd(requestTmpID, "CallOwnerOf Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(requestTmpID, "PASS - 04-1. Get Address OwnerOf Old MetaID")

	if findAddress == nil {
		log.Debugd(requestTmpID, "not found address for old MetaID")
		errObj := &notExistsAddressError{"not found Address for old metaID"} //"Cannot Update Meta ID"
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugfd(requestTmpID, "PASS - 04-2. Get Address OwnerOf Old MetaID : %v", findAddress.String())

	if !bytes.Equal(findAddress.Bytes(), reqParam.Address.Bytes()) {
		log.Debugd(requestTmpID, "Error - Address is not valid")
		errObj := &invalidAddressError{"Cannot find valid user Address"}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(requestTmpID, "PASS - 04-3. Address Valid ")

	//5. Reqest Get UserAddress for New Meta ID
	//  if GetAddress ==  nil   then pass
	//  else Error
	//findNewAddress := &common.Address{}
	findNewAddress, err := identitymanager.CallOwnerOf(reqParam.NewMetaID)
	if err != nil {
		log.Errorfd(requestTmpID, "CallOwnerOf Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(requestTmpID, "PASS - 05-1. Get Address OwnerOf New MetaID")

	if findNewAddress != nil {
		log.Debugfd(requestTmpID, "already exists address for new MetaID(id/addr): %v, %v", reqParam.NewMetaID.String(), findNewAddress.String())
		errObj := &alreadyExistsAddressError{"already exists address for new MetaID"} //"Cannot Update Meta ID"
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugfd(requestTmpID, "PASS - 05-2. Get Address OwnerOf Old MetaID : %v", findAddress.String())

	//5. Send Transaction for Calling SC Function[ updateMetaID ]
	//  return txid
	trx, err := identitymanager.CallUpdateMetaID(reqParam.OldMetaID, reqParam.NewMetaID, reqParam.Signature, reqParam.Address)
	if err != nil {
		log.Errorfd(requestTmpID, "CallUpdateMetaID Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		// resp.Error = &json.RPCError{
		// 	Code:    errObj.ErrorCode(),
		// 	Message: errObj.Error(),
		// }
		return
	}
	log.Debugfd(requestTmpID, "PASS - 06. Call SC UpdateMetaID : %v", trx.Hash().String())
	resp.Result = trx.Hash().String()

	return
}

func backupUserData(req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	requestTmpID := proxyCommon.RandomUint64()
	log.Debugd(requestTmpID, "Call backupUserData Function")
	//var reqParam interface{}
	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	//1. Check paramter format
	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	reqParam := tmpParams.(metaIDBackupParams)
	log.Debugfd(requestTmpID, "parameter[Address] : %v", reqParam.Address.String())
	log.Debugfd(requestTmpID, "parameter[MetaId] : %v", reqParam.MetaID)
	log.Debugfd(requestTmpID, "parameter[EncData] : %v", reqParam.EncData)
	log.Debugfd(requestTmpID, "parameter[Sig] : %v", reqParam.Signature)
	log.Debugd(requestTmpID, "PASS - 01. Check Parameter")

	//2. check Signature
	//  if  address == ecrecover()  then Pass else Error
	// signedAddress, err := crypto.EcRecover(reqParam.MetaID.String(), reqParam.Signature.String())
	// if err != nil {
	// 	log.Errorfd(requestTmpID, "EcRecover Error : %v", err)
	// 	errObj := &internalError{"Failed to EcRecover"}
	// 	resp.Error = makeErrorResponse(errObj)
	// 	return
	// }

	// if !bytes.Equal(signedAddress.Bytes(), reqParam.Address.Bytes()) {
	// 	log.Debugfd(requestTmpID, "Error  :Sign Error %v / %v ", reqParam.Address.String(), signedAddress.String())
	// 	errObj := &invalidSignatureError{"Failed to verify signature"}
	// 	resp.Error = makeErrorResponse(errObj)
	// 	return
	// }
	_, errObj = verifySignature(requestTmpID, reqParam.MetaID.String(), reqParam.Signature.String(), &reqParam.Address)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(requestTmpID, "PASS - 02. Check Signature")
	//3. address == OwnerOf(metaID)
	//findAddress := &common.Address{}
	findAddress, err := identitymanager.CallOwnerOf(reqParam.MetaID)
	if err != nil {
		log.Errorfd(requestTmpID, "CallOwnerOf Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(requestTmpID, "PASS - 03-1. Get Address OwnerOf MetaID")

	if findAddress == nil {
		log.Debugd(requestTmpID, "not found address for MetaID")
		errObj := &notExistsAddressError{"not found address for metaID"}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugfd(requestTmpID, "PASS - 03-2. Get Address OwnerOfMetaID : %v", findAddress.String())

	if !bytes.Equal(findAddress.Bytes(), reqParam.Address.Bytes()) {
		log.Debugd(requestTmpID, "Error - Address is not valid")
		errObj := &invalidAddressError{"Cannot find valid user Address"}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(requestTmpID, "PASS - 03-3. Address Valid ")

	//4. Save file to IPFS
	ins := ipfs.GetInstance()
	fileHash, err := ins.Add(reqParam.EncData.String())
	if err != nil {
		log.Errorfd(requestTmpID, "ipfs.Add Error : %v", err)
		errObj := &saveBackupFileError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(requestTmpID, "PASS - 04-1. ipfs add ", fileHash)
	err = ins.Pin(fileHash)
	if err != nil {
		log.Errorfd(requestTmpID, "ipfs.Pin Error : %v", err)
		errObj := &saveBackupFileError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(requestTmpID, "PASS - 04-2. ipfs Pin ", fileHash)
	err = ins.PinByCluster(fileHash)
	//TODO: 당분간 Cluster 에러 PASS
	// if err != nil {
	// 	log.Errorf("ipfs.PinByCluster Error : %v", err)
	// 	errObj := &saveBackupFileError{err.Error()}
	// 	resp.Error = makeErrorResponse(errObj)
	// 	// resp.Error = &json.RPCError{
	// 	// 	Code:    errObj.ErrorCode(),
	// 	// 	Message: errObj.Error(),
	// 	// }
	// 	return
	// }
	log.Debugd(requestTmpID, "PASS - 04-3. ipfs PinByCluster ", fileHash)
	//  return fileID
	resp.Result = fileHash
	return
}

func getUserData(req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	requestTmpID := proxyCommon.RandomUint64()
	log.Debugd(requestTmpID, "Call getUserData Function")
	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	//1. Check paramter format
	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	reqParam := tmpParams.(metaIDGetUserDataParams)
	//log.Debugfd(requestTmpID, "parameter[NewAddress] : %v", reqParam.NewAddress.String())
	log.Debugfd(requestTmpID, "parameter[FileID] : %v", reqParam.FileID)
	//log.Debugfd(requestTmpID, "parameter[Sig] : %v", reqParam.Signature)
	log.Debugd(requestTmpID, "PASS - 01. Check Parameter")

	//TODO: 테스트로 인해 address/signature 필드 제거  추후 결정 필요
	// //2. check Signature
	// //  if  newAddress == ecrecover()  then Pass else Error

	// fHex := hexutil.Bytes(reqParam.FileID)
	// signedAddress, err := crypto.EcRecover(fHex.String(), reqParam.Signature.String())
	// if err != nil {
	// 	log.Errorf("EcRecover Error : %v", err)
	// 	errObj := &invalidSignatureError{"Failed to verify signature"}
	// 	resp.Error = &json.RPCError{
	// 		Code:    errObj.ErrorCode(),
	// 		Message: errObj.Error(),
	// 	}
	// 	return
	// }
	// log.Debugfd(requestTmpID, "PASS - 02-1. Check Signature signedAddress : %v ", signedAddress.String())

	// if !bytes.Equal(signedAddress.Bytes(), reqParam.NewAddress.Bytes()) {
	// 	log.Errorf("Error  :Sign Error %v / %v ", reqParam.NewAddress.String(), signedAddress.String())
	// 	errObj := &invalidSignatureError{"Failed to verify signature"}
	// 	resp.Error = &json.RPCError{
	// 		Code:    errObj.ErrorCode(),
	// 		Message: errObj.Error(),
	// 	}
	// 	return
	// }
	// log.Debugd(requestTmpID, "PASS - 02-2. signed Address Valid")

	//3. GET User Data from IPFS
	//return data
	ins := ipfs.GetInstance()
	ret := ins.Cat(reqParam.FileID)
	if ret == "" {
		log.Debugd(requestTmpID, "Error : Cannot find file ")
		errObj := &notFoundFileError{"not found backup file"}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(requestTmpID, "PASS - 03. ipfs.Cat")
	// //  return fileID
	resp.Result = ret
	return
}

func restoreUserData(req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	requestTmpID := proxyCommon.RandomUint64()
	log.Debugd(requestTmpID, "Call restoreUserData Function")
	//var reqParam interface{}
	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	//1. Check paramter format
	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	reqParam := tmpParams.(metaIDRestoreParams)
	log.Debugfd(requestTmpID, "parameter[NewAddress] : %v", reqParam.NewAddress.String())
	log.Debugfd(requestTmpID, "parameter[OldAddress] : %v", reqParam.OldAddress.String())
	log.Debugfd(requestTmpID, "parameter[NewMetaID] : %v", reqParam.NewMetaID)
	log.Debugfd(requestTmpID, "parameter[OldMetaID] : %v", reqParam.OldMetaID)
	log.Debugfd(requestTmpID, "parameter[userHash] : %v", reqParam.UserHashs)
	log.Debugfd(requestTmpID, "parameter[Sig] : %v", reqParam.Signature)
	log.Debugd(requestTmpID, "PASS - 01. Check Parameter")

	//2. check Signature
	// 	   signdata = newMetaID
	//  if  address == ecrecover()  then Pass else Error
	// signedAddress, err := crypto.EcRecover(reqParam.NewMetaID.String(), reqParam.Signature.String())
	// if err != nil {
	// 	log.Errorfd(requestTmpID, "EcRecover Error : %v", err)
	// 	errObj := &internalError{"Failed to EcRecover"}
	// 	resp.Error = makeErrorResponse(errObj)
	// 	return
	// }
	// log.Debugfd(requestTmpID, "PASS - 02-1. Check Signature signedAddress : %v ", signedAddress.String())

	// if !bytes.Equal(signedAddress.Bytes(), reqParam.NewAddress.Bytes()) {
	// 	log.Debugfd(requestTmpID, "Error  :Sign Error %v / %v ", reqParam.NewAddress.String(), signedAddress.String())
	// 	errObj := &invalidSignatureError{"Failed to verify signature"}
	// 	resp.Error = makeErrorResponse(errObj)
	// 	return
	// }
	_, errObj = verifySignature(requestTmpID, reqParam.NewMetaID.String(), reqParam.Signature.String(), &reqParam.NewAddress)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(requestTmpID, "PASS - 02-2. signed Address Valid")
	// 3. Verification new MetaID  by make merkle tree root.
	//   if newMetaID == root  then Pass  else  Error

	var usrHashs []merkletree.Content

	for _, userHash := range reqParam.UserHashs {
		usrHashs = append(usrHashs, merkletree.HashContent{H: userHash})
	}

	nTree, _ := merkletree.NewTree(usrHashs)

	//Get the Merkle Root of the tree
	nMr := nTree.MerkleRoot()

	nRoot := hexutil.Bytes(nMr)
	log.Debugfd(requestTmpID, "PASS - 03-1. Make Merkle Root For new MetaID: %v", nRoot)

	if !bytes.Equal(nRoot, reqParam.NewMetaID) {
		log.Debugfd(requestTmpID, "new MetaID invalid : %v, %x ", reqParam.NewMetaID, nRoot)
		errObj := &invalidMetaIDError{"Failed to verify New MetaID"}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	log.Debugfd(requestTmpID, "PASS - 03-2. New MetaID valid : %v ", reqParam.NewMetaID)
	// 4. Verification oldMetaID (merkleTree) by making merkle tree root with oldAddress

	oldHash := ethCrypto.Keccak256(reqParam.OldAddress.Bytes())
	//change newAddress to oldAddress
	usrHashs[addressHashIdx] = merkletree.HashContent{H: oldHash}

	oTree, _ := merkletree.NewTree(usrHashs)

	//Get the Merkle Root of the tree
	oMr := oTree.MerkleRoot()

	oRoot := hexutil.Bytes(oMr)

	log.Debugfd(requestTmpID, "PASS - 04-1. Make Merkle Root For old MetaID: %v", oRoot)

	if !bytes.Equal(oRoot, reqParam.OldMetaID) {
		log.Debugfd(requestTmpID, "Error - old MetaID invalid : %v, %x ", reqParam.OldMetaID, oRoot)
		errObj := &invalidOldMetaIDError{"Failed to verify Old MetaID"}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugfd(requestTmpID, "PASS - 04-2. Old MetaID valid : %v ", reqParam.OldMetaID)
	// 5. Get user Address for Old Meta ID
	//    getAddres == nil   then Error
	//    getAddress == oldAddress  then pass
	//findOldAddress := &common.Address{}
	findOldAddress, err := identitymanager.CallOwnerOf(reqParam.OldMetaID)
	if err != nil {
		log.Errorfd(requestTmpID, "CallOwnerOf Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(requestTmpID, "PASS - 05-1. Get Address OwnerOf Old MetaID")
	if findOldAddress == nil {
		log.Debugfd(requestTmpID, "Not Found Address for Old MetaID : %v", reqParam.OldMetaID)
		errObj := &notExistsAddressError{"not found address for old metaID"}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugfd(requestTmpID, "PASS - 05-2. Get Address OwnerOf Old MetaID : %v", findOldAddress.String())

	if !bytes.Equal(findOldAddress.Bytes(), reqParam.OldAddress.Bytes()) {
		log.Debugd(requestTmpID, "Old Address is not valid")
		errObj := &invalidAddressError{"Cannot find valid user Address"}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(requestTmpID, "PASS - 05-3. Old Address is valid")
	// 6. Get user Address for New Meta ID
	//    getAddres == nil   then Pass else error
	//findNewAddress := &common.Address{}
	findNewAddress, err := identitymanager.CallOwnerOf(reqParam.NewMetaID)
	if err != nil {
		log.Errorfd(requestTmpID, "CallOwnerOf Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	log.Debugd(requestTmpID, "PASS - 06-1. Get Address OwnerOf New MetaID")
	if findNewAddress != nil {
		log.Debugfd(requestTmpID, "address for new MetaID is already Existed : %v", findNewAddress.String())
		errObj := &alreadyExistsAddressError{"already exists address for new metaID"} //Cannot restore Meta ID
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(requestTmpID, "PASS - 06-2. Not Found Address for New MetaID")
	// 7. Call restoreMetaID  function
	trx, err := identitymanager.CallRestoreMetaID(reqParam.OldMetaID, reqParam.NewMetaID, reqParam.OldAddress, reqParam.NewAddress, reqParam.Signature)
	if err != nil {
		log.Errorfd(requestTmpID, "CallRestoreMetaID Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugfd(requestTmpID, "PASS - 07. Call SC restoreMetaID : %v", trx.Hash().String())
	resp.Result = trx.Hash().String()
	return
}

func revokeMetaID(req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	requestTmpID := proxyCommon.RandomUint64()
	log.Debugd(requestTmpID, "Call revokeMetaID Function")

	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	//1. Check paramter format
	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	reqParam := tmpParams.(metaIDRevokeParams)
	log.Debugfd(requestTmpID, "parameter[MetaID] : %v", reqParam.MetaID)
	log.Debugfd(requestTmpID, "parameter[Timestamp] : %v", reqParam.Timestamp)
	log.Debugfd(requestTmpID, "parameter[Signature] : %v", reqParam.Signature)
	log.Debugd(requestTmpID, "PASS - 01. Check Parameter")

	//	2. Get Address from  Signature
	//  signdata = metaID|timestamp
	mSize := len(reqParam.MetaID)
	tSize := len(reqParam.MetaID) + len(reqParam.Timestamp)
	var signData hexutil.Bytes
	signData = make(hexutil.Bytes, tSize, tSize)
	copy(signData, reqParam.MetaID)
	copy(signData[mSize:tSize], reqParam.Timestamp)
	log.Debugfd(requestTmpID, "sign data : %v", signData.String())

	signedAddress, errObj := verifySignature(requestTmpID, signData.String(), reqParam.Signature.String(), nil)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugfd(requestTmpID, "PASS - 02. Check Signature signedAddress : %v ", signedAddress.String())
	// 3. Get Address from Ownerof function(Smart Contract)
	//findAddress := &common.Address{}
	findAddress, err := identitymanager.CallOwnerOf(reqParam.MetaID)
	if err != nil {
		log.Errorfd(requestTmpID, "CallOwnerOf Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	log.Debugd(requestTmpID, "PASS - 03-1. Get Address OwnerOf MetaID")

	if findAddress == nil {
		log.Debugfd(requestTmpID, "Error - not found Address for MetaID (%v) ", reqParam.MetaID)
		errObj := &notExistsAddressError{"not found Address for metaID"}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugfd(requestTmpID, "PASS - 03-2. Get Address OwnerOf MetaID : %v", findAddress.String())

	if !bytes.Equal(findAddress.Bytes(), signedAddress.Bytes()) {
		log.Debugfd(requestTmpID, "signed Address Not Valid ")
		errObj := &invalidAddressError{"Cannot find valid user Address"}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(requestTmpID, "PASS - 03-3. signed Address Valid")

	// 4. Call deleteMetaID function
	trx, err := identitymanager.CallDeleteMetaID(reqParam.MetaID, reqParam.Timestamp, reqParam.Signature)
	if err != nil {
		log.Errorfd(requestTmpID, "CallDeleteMetaID Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugfd(requestTmpID, "PASS - 04. Call SC DeleteMetaID : %v", trx.Hash().String())
	resp.Result = trx.Hash().String()
	return
}
