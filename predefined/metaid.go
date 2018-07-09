package predefined

import (
	"bytes"
	encodingJson "encoding/json"
	"fmt"
	"log"

	"bitbucket.org/coinplugin/proxy/crypto"
	"bitbucket.org/coinplugin/proxy/ipfs"
	"bitbucket.org/coinplugin/proxy/json"
	"bitbucket.org/coinplugin/proxy/predefined/merkletree"
	"bitbucket.org/coinplugin/proxy/predefined/sc/identitymanager"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
)

const (
	addressSize         = 20
	hashSize            = 32
	signatureSize       = 65
	fileHashSize        = 46
	timestampSize       = 14
	minimumUserHashSize = 3
	minimumUserDataSize = 10
	addressHashIdx      = 0
)

type metaIDRegisterParams struct {
	Address   common.Address  `json:"address"`
	MetaID    hexutil.Bytes   `json:"meta_id"`
	UserHashs []hexutil.Bytes `json:"user_hash_list"`
	Signature hexutil.Bytes   `json:"signature"` // Sign(MetaId)
}

type metaIDUpdateParams struct {
	Address   common.Address  `json:"address"`
	OldMetaID hexutil.Bytes   `json:"meta_id_old"`
	NewMetaID hexutil.Bytes   `json:"meta_id_new"`
	UserHashs []hexutil.Bytes `json:"user_hash_list"`
	Signature hexutil.Bytes   `json:"signature"` // Sign(NewMetaId)
}

type metaIDBackupParams struct {
	Address   common.Address `json:"address"`
	MetaID    hexutil.Bytes  `json:"meta_id"`
	EncData   hexutil.Bytes  `json:"enc_data"`
	Signature hexutil.Bytes  `json:"signature"` // Sign(MetaId)
}

type metaIDGetUserDataParams struct {
	NewAddress common.Address `json:"address_new"`
	FileID     string         `json:"file_id"`
	Signature  hexutil.Bytes  `json:"signature"` // Sign(FileId)
}

type metaIDRestoreParams struct {
	NewAddress common.Address  `json:"address_new"`
	OldAddress common.Address  `json:"address_old"`
	NewMetaID  hexutil.Bytes   `json:"meta_id_new"`
	OldMetaID  hexutil.Bytes   `json:"meta_id_old"`
	UserHashs  []hexutil.Bytes `json:"user_hash_list"`
	Signature  hexutil.Bytes   `json:"signature"` //sign(NewMetaId)
}

type metaIDRevokeParams struct {
	MetaID    hexutil.Bytes `json:"meta_id"`
	Timestamp hexutil.Bytes `json:"timestamp"`
	Signature hexutil.Bytes `json:"signature"` //sign(MetaId|timestamp)
}

type metaParam interface {
	validate() Error
}

//validate is function of checking if Parameter is valid or invalid. if invalid, return invalidParamsError
func (p *metaIDRegisterParams) validate() Error {

	if len(p.Address) != addressSize {
		return &invalidParamsError{fmt.Sprintf("Invalid address parameter %v", len(p.Address))}
	}
	if len(p.MetaID) != hashSize {
		return &invalidParamsError{fmt.Sprintf("Invalid meta_id parameter %v", len(p.MetaID))}
	}
	if len(p.Signature) != signatureSize {
		return &invalidParamsError{fmt.Sprintf("Invalid signature parameter %v", len(p.Signature))}
	}
	if len(p.UserHashs) < minimumUserHashSize {
		return &invalidParamsError{fmt.Sprintf("Invalid user_hash_list parameter  %v", len(p.UserHashs))}
	}
	for _, hash := range p.UserHashs {
		if len(hash) != hashSize {
			return &invalidParamsError{fmt.Sprintf("Invalid user_hash_list parameter element  %v", len(p.MetaID))}
		}
	}

	return nil
}

func (p *metaIDUpdateParams) validate() Error {
	if len(p.Address) != addressSize {
		return &invalidParamsError{"Invalid address parameter."}
	}
	if len(p.OldMetaID) != hashSize {
		return &invalidParamsError{"Invalid meta_id_old parameter."}
	}
	if len(p.NewMetaID) != hashSize {
		return &invalidParamsError{"Invalid meta_id_new parameter."}
	}
	if len(p.Signature) != signatureSize {
		return &invalidParamsError{"Invalid signature parameter."}
	}
	if len(p.UserHashs) < minimumUserHashSize {
		return &invalidParamsError{"Invalid user_hash_list parameter."}
	}
	for _, hash := range p.UserHashs {
		if len(hash) != hashSize {
			return &invalidParamsError{"Invalid user_hash_list parameter element."}
		}
	}

	return nil
}

func (p *metaIDBackupParams) validate() Error {

	if len(p.Address) != addressSize {
		return &invalidParamsError{"Invalid address parameter."}
	}
	if len(p.MetaID) != hashSize {
		return &invalidParamsError{"Invalid meta_id parameter."}
	}
	if len(p.EncData) < minimumUserHashSize {
		return &invalidParamsError{"Invalid enc_data parameter."}
	}
	if len(p.Signature) != signatureSize {
		return &invalidParamsError{"Invalid signature parameter."}
	}
	return nil
}
func (p *metaIDGetUserDataParams) validate() Error {
	if len(p.NewAddress) != addressSize {
		return &invalidParamsError{"Invalid address_new parameter."}
	}

	if len(p.FileID) < fileHashSize {
		return &invalidParamsError{"Invalid meta_id parameter."}
	}

	if len(p.Signature) != signatureSize {
		return &invalidParamsError{"Invalid signature parameter."}
	}
	return nil
}

func (p *metaIDRestoreParams) validate() Error {
	if len(p.NewAddress) != addressSize {
		return &invalidParamsError{"Invalid address_new parameter."}
	}
	if len(p.OldAddress) != addressSize {
		return &invalidParamsError{"Invalid address_old parameter."}
	}
	if len(p.NewMetaID) != hashSize {
		return &invalidParamsError{"Invalid meta_id_new parameter."}
	}
	if len(p.OldMetaID) != hashSize {
		return &invalidParamsError{"Invalid meta_id_old parameter."}
	}
	if len(p.Signature) != signatureSize {
		return &invalidParamsError{"Invalid signature parameter."}
	}

	if len(p.UserHashs) < minimumUserHashSize {
		return &invalidParamsError{"Invalid user_hash_list parameter."}
	}
	for _, hash := range p.UserHashs {
		if len(hash) != hashSize {
			return &invalidParamsError{"Invalid user_hash_list parameter element."}
		}
	}

	return nil
}

func (p *metaIDRevokeParams) validate() Error {

	if len(p.MetaID) != hashSize {
		return &invalidParamsError{"Invalid meta_id parameter."}
	}
	if len(p.Timestamp) != timestampSize {
		return &invalidParamsError{"Invalid timesamp parameter."}
	}
	if len(p.Signature) != signatureSize {
		return &invalidParamsError{"Invalid signature parameter."}
	}

	return nil
}

// fillParam is function of converting map type  to struct type and check Valid parameter
// func (p *metaIDRegisterParams) fillParam(obj interface{}) Error {
// 	jsonBytes, err1 := encodingJson.Marshal(obj)

// 	err1 = encodingJson.Unmarshal(jsonBytes, p)

// 	if err1 != nil {
// 		fmt.Println("err : ", err1)
// 		return &invalidMessageError{err1.Error()}
// 	}

// 	err := p.validate()
// 	if err != nil {
// 		return err
// 	}
// 	return nil

// }

func fillParam(p metaParam, obj interface{}) Error {
	jsonBytes, err1 := encodingJson.Marshal(obj)

	err1 = encodingJson.Unmarshal(jsonBytes, p)

	if err1 != nil {
		log.Println("err : ", err1)
		return &invalidMessageError{err1.Error()}
	}

	err := p.validate()
	if err != nil {
		return err
	}
	return nil

}

// func (p *metaIDUpdateParams) fillParam(obj interface{}) Error {
// 	jsonBytes, err1 := encodingJson.Marshal(obj)

// 	err1 = encodingJson.Unmarshal(jsonBytes, p)

// 	if err1 != nil {
// 		fmt.Println("err : ", err1)
// 		return &invalidMessageError{err1.Error()}
// 	}

// 	err := p.validate()
// 	if err != nil {
// 		return err
// 	}
// 	return nil

// }

func getParameter(method string, params []interface{}) (interface{}, Error) {
	if len(params) != 1 {
		return nil, &invalidParamsError{"Invalid params."}
	}
	obj := params[0]
	switch method {
	case "register_meta_id":
		var reqParam metaIDRegisterParams
		err := fillParam(&reqParam, obj)
		if err != nil {
			return nil, err
		}

		return reqParam, nil

	case "update_meta_id":
		var reqParam metaIDUpdateParams
		err := fillParam(&reqParam, obj)
		if err != nil {
			return nil, err
		}
		return reqParam, nil

	case "backup_user_data":
		var reqParam metaIDBackupParams
		err := fillParam(&reqParam, obj)
		if err != nil {
			return nil, err
		}
		return reqParam, nil

	case "get_user_data":
		var reqParam metaIDGetUserDataParams
		err := fillParam(&reqParam, obj)
		if err != nil {
			return nil, err
		}
		return reqParam, nil

	case "restore_user_data":
		var reqParam metaIDRestoreParams
		err := fillParam(&reqParam, obj)
		if err != nil {
			return nil, err
		}
		return reqParam, nil

	case "revoke_meta_id":
		var reqParam metaIDRevokeParams
		err := fillParam(&reqParam, obj)
		if err != nil {
			return nil, err
		}
		return reqParam, nil

	default:
		return nil, &methodNotFoundError{method}

	}

}

func registerMetaID(req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	//var reqParam interface{}
	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	//1. Check paramter format
	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}
	reqParam := tmpParams.(metaIDRegisterParams)
	log.Println("parameter[Address] : ", reqParam.Address)
	log.Println("parameter[MetaId] : ", reqParam.MetaID)
	log.Println("parameter[Sig] : ", reqParam.Signature)
	log.Println("parameter[UserHashs] : ", reqParam.UserHashs)

	//2. check Signature
	//  if  address == ecrecover()  then Pass else Error
	signedAddress, err := crypto.EcRecover(reqParam.MetaID.String(), reqParam.Signature.String())
	if err != nil {
		log.Println("Error :", err)
		errObj := &invalidSignatureError{"Failed to verify signature"}
		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}

	if !bytes.Equal(signedAddress.Bytes(), reqParam.Address.Bytes()) {
		log.Printf("Error  :Sign Error %v / %v \n", reqParam.Address.Hex(), signedAddress)
		errObj := &invalidSignatureError{"Failed to verify signature"}
		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}

	//3. Check Meta ID  by make merkle tree root.
	//   if metaId == root  then Pass  else  Error

	var usrHashs []merkletree.Content

	//raws := []string{"abc", "123"}
	for _, userHash := range reqParam.UserHashs {

		usrHashs = append(usrHashs, HashContent{h: userHash})
	}

	//Create a new Merkle Tree from the list of Content
	tree, _ := merkletree.NewTree(usrHashs)

	//Get the Merkle Root of the tree
	mr := tree.MerkleRoot()
	root := hexutil.Bytes(mr)

	log.Printf("root hash: %v \n", root)

	if root.String() != reqParam.MetaID.String() {
		log.Printf("MetaID invalid : %v, %x \n", reqParam.MetaID, root)
		errObj := &invalidMetaIDError{"Failed to verify signature"}
		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}

	log.Printf("MetaID valid : %v \n", reqParam.MetaID)

	//4. Reqest Get UserAddress for Meta ID
	//  if GetAddress == nil  then Pass  else Error
	address := &common.Address{}
	address, err = identitymanager.CallOwnerOf(reqParam.MetaID)
	if err != nil {
		errObj := &internalError{err.Error()}

		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}
	if address != nil {

		errObj := &alreadyExistsError{"Cannot register Meta ID"}
		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}

	//5. Send Transaction for Calling SC Function [createNewMetaID]
	//identitymanager.CallCreateMetaID()
	//  return txid

	trx, err := identitymanager.CallCreateMetaID(reqParam.MetaID, reqParam.Signature, reqParam.Address)
	if err != nil {
		errObj := &internalError{err.Error()}

		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}

	// resp.Result = trx.Hash().String()

	/*
		var trx *types.Transaction
		tx := func(nonce uint64) error {
			var err error
			trx, err = identitymanager.CallCreateMetaID(reqParam.MetaID, reqParam.Signature, reqParam.Address, nonce)
			if err != nil {
				return err
			}
			return nil
		}
		c := crypto.GetInstance()
		res := c.ApplyNonce(tx)

		if !res {
			errObj := &internalError{"call function Error - CreateMetaID"}

			resp.Error = &json.RPCError{
				Code:    errObj.ErrorCode(),
				Message: errObj.Error(),
			}
		}
	*/

	resp.Result = trx.Hash().String()
	return
}

func updateMetaID(req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	//var reqParam interface{}
	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	//1. Check paramter format
	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}
	reqParam := tmpParams.(metaIDUpdateParams)
	log.Println("parameter[Address] : ", reqParam.Address)
	log.Println("parameter[NewMetaId] : ", reqParam.NewMetaID)
	log.Println("parameter[OldMetaId] : ", reqParam.OldMetaID)
	log.Println("parameter[Sig] : ", reqParam.Signature)
	log.Println("parameter[UserHashs] : ", reqParam.UserHashs)

	//2. check Signature
	//  if  address == ecrecover()  then Pass else Error
	signedAddress, err := crypto.EcRecover(reqParam.NewMetaID.String(), reqParam.Signature.String())
	if err != nil {
		log.Println("Error :", err)
		errObj := &invalidSignatureError{"Failed to verify signature"}
		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}

	if !bytes.Equal(signedAddress.Bytes(), reqParam.Address.Bytes()) {
		log.Printf("Error  :Sign Error %v / %v \n", reqParam.Address.Hex(), signedAddress)
		errObj := &invalidSignatureError{"Failed to verify signature"}
		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}

	//3. Check New Meta ID  by make merkle tree root.
	//   if newMetaID == root  then Pass  else  Error

	var usrHashs []merkletree.Content

	for _, userHash := range reqParam.UserHashs {

		usrHashs = append(usrHashs, HashContent{h: userHash})

	}

	tree, _ := merkletree.NewTree(usrHashs)

	//Get the Merkle Root of the tree
	mr := tree.MerkleRoot()

	root := hexutil.Bytes(mr)

	log.Printf("root hash: %v \n", root)

	if root.String() != reqParam.NewMetaID.String() {
		log.Printf("MetaID invalid : %v, %x \n", reqParam.NewMetaID, root)
		errObj := &invalidSignatureError{"Failed to verify signature"}
		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}

	log.Printf("MetaID valid : %v \n", reqParam.NewMetaID)

	//4. Reqest Get UserAddress for New Meta ID
	//  if GetAddress ==  reqParam.Address   then pass
	//  else Error
	findAddress := &common.Address{}
	findAddress, err = identitymanager.CallOwnerOf(reqParam.OldMetaID)
	if err != nil {
		errObj := &internalError{err.Error()}

		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}
	if findAddress == nil {

		errObj := &notExistsError{"Cannot Update Meta ID"}
		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}
	if !bytes.Equal(findAddress.Bytes(), reqParam.Address.Bytes()) {
		errObj := &invalidAddressError{"Cannot find valid user Address"}
		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}
	//5. Send Transaction for Calling SC Function[ updateMetaID ]
	//  return txid
	trx, err := identitymanager.CallUpdateMetaID(reqParam.OldMetaID, reqParam.NewMetaID, reqParam.Signature, reqParam.Address)
	if err != nil {
		errObj := &internalError{err.Error()}

		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}
	resp.Result = trx.Hash().String()
	return
}

func backupUserData(req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	//var reqParam interface{}
	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	//1. Check paramter format
	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}
	reqParam := tmpParams.(metaIDBackupParams)
	log.Println("parameter[Address] : ", reqParam.Address)
	log.Println("parameter[MetaId] : ", reqParam.MetaID)
	log.Println("parameter[EncData] : ", reqParam.EncData)
	log.Println("parameter[Sig] : ", reqParam.Signature)

	//2. check Signature
	//  if  address == ecrecover()  then Pass else Error
	signedAddress, err := crypto.EcRecover(reqParam.MetaID.String(), reqParam.Signature.String())
	if err != nil {
		log.Println("Error :", err)
		errObj := &invalidSignatureError{"Failed to verify signature"}
		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}

	if !bytes.Equal(signedAddress.Bytes(), reqParam.Address.Bytes()) {
		log.Printf("Error  :Sign Error %v / %v \n", reqParam.Address.Hex(), signedAddress)
		errObj := &invalidSignatureError{"Failed to verify signature"}
		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}

	//4. Save file to IPFS
	ins := ipfs.GetInstance()
	fileHash, err := ins.Add(reqParam.EncData.String())
	if err != nil {
		log.Println("Error :", err)
		errObj := &internalError{err.Error()}
		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}
	// //  return fileID
	resp.Result = fileHash
	return
}

func getUserData(req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	//var reqParam interface{}
	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	//1. Check paramter format
	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}
	reqParam := tmpParams.(metaIDGetUserDataParams)
	log.Println("parameter[NewAddress] : ", reqParam.NewAddress)
	log.Println("parameter[FileID] : ", reqParam.FileID)
	log.Println("parameter[Sig] : ", reqParam.Signature)

	//2. check Signature
	//  if  newAddress == ecrecover()  then Pass else Error
	fHex := hexutil.Bytes(reqParam.FileID)
	signedAddress, err := crypto.EcRecover(fHex.String(), reqParam.Signature.String())
	if err != nil {
		log.Println("Error :", err)
		errObj := &invalidSignatureError{"Failed to verify signature"}
		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}

	if !bytes.Equal(signedAddress.Bytes(), reqParam.NewAddress.Bytes()) {
		log.Printf("Error  :Sign Error %v / %v \n", reqParam.NewAddress.Hex(), signedAddress)
		errObj := &invalidSignatureError{"Failed to verify signature"}
		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}
	//3. GET User Data from IPFS
	//return data
	ins := ipfs.GetInstance()
	ret := ins.Cat(reqParam.FileID)
	if ret == "" {
		log.Println("Error : Cannot find file ")
		errObj := &internalError{"Cannot find file"}
		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}
	// //  return fileID
	resp.Result = ret
	return
}

func restoreUserData(req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	//var reqParam interface{}
	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	//1. Check paramter format
	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}
	reqParam := tmpParams.(metaIDRestoreParams)
	log.Println("parameter[NewAddress] : ", reqParam.NewAddress)
	log.Println("parameter[OldAddress] : ", reqParam.OldAddress)
	log.Println("parameter[NewMetaID] : ", reqParam.NewMetaID)
	log.Println("parameter[OldMetaID] : ", reqParam.OldMetaID)
	log.Println("parameter[userHash] : ", reqParam.UserHashs)
	log.Println("parameter[Sig] : ", reqParam.Signature)
	//	2. Get Address from  Signature
	// 	   signdata = newMetaID

	//2. check Signature
	//  if  address == ecrecover()  then Pass else Error
	signedAddress, err := crypto.EcRecover(reqParam.NewMetaID.String(), reqParam.Signature.String())
	if err != nil {
		log.Println("Error :", err)
		errObj := &invalidSignatureError{"Failed to verify signature"}
		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}

	if !bytes.Equal(signedAddress.Bytes(), reqParam.NewAddress.Bytes()) {
		log.Printf("Error  :Sign Error %v / %v \n", reqParam.NewAddress.Hex(), signedAddress)
		errObj := &invalidSignatureError{"Failed to verify signature"}
		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}

	// 3. Verification new MetaID  by make merkle tree root.
	//   if newMetaID == root  then Pass  else  Error

	var usrHashs []merkletree.Content

	for _, userHash := range reqParam.UserHashs {

		usrHashs = append(usrHashs, HashContent{h: userHash})

	}

	nTree, _ := merkletree.NewTree(usrHashs)

	//Get the Merkle Root of the tree
	nMr := nTree.MerkleRoot()

	nRoot := hexutil.Bytes(nMr)

	log.Printf("root hash: %v \n", nRoot)

	if nRoot.String() != reqParam.NewMetaID.String() {
		log.Printf("new MetaID invalid : %v, %x \n", reqParam.NewMetaID, nRoot)
		errObj := &invalidSignatureError{"Failed to verify New MetaID"}
		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}

	log.Printf("MetaID valid : %v \n", reqParam.NewMetaID)
	// 4. Verification oldMetaID (merkleTree) by making merkle tree root with oldAddress

	oldHash := ethCrypto.Keccak256(reqParam.OldAddress.Bytes())
	//change newAddress to oldAddress
	usrHashs[addressHashIdx] = HashContent{h: oldHash}

	oTree, _ := merkletree.NewTree(usrHashs)

	//Get the Merkle Root of the tree
	oMr := oTree.MerkleRoot()

	oRoot := hexutil.Bytes(oMr)

	log.Printf("root hash: %v \n", oRoot)

	if oRoot.String() != reqParam.OldMetaID.String() {
		log.Printf("old MetaID invalid : %v, %x \n", reqParam.OldMetaID, oRoot)
		errObj := &invalidSignatureError{"Failed to verify OldMetaID"}
		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}

	// 5. Get user Address for Old Meta ID
	//    getAddres == nil   then Error
	//    getAddress == oldAddress  then pass
	findOldAddress := &common.Address{}
	findOldAddress, err = identitymanager.CallOwnerOf(reqParam.OldMetaID)
	if err != nil {
		errObj := &internalError{err.Error()}

		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}
	if findOldAddress == nil {

		errObj := &notExistsError{"Canot get address for old Meta ID"}
		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}
	if !bytes.Equal(findOldAddress.Bytes(), reqParam.OldAddress.Bytes()) {
		errObj := &invalidAddressError{"Cannot find valid user Address"}
		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}
	// 6. Get user Address for New Meta ID
	//    getAddres == nil   then Pass else error
	findNewAddress := &common.Address{}
	findNewAddress, err = identitymanager.CallOwnerOf(reqParam.NewMetaID)
	if err != nil {
		errObj := &internalError{err.Error()}

		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}
	if findNewAddress != nil {

		errObj := &alreadyExistsError{"Cannot restore Meta ID"}
		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}
	// 7. Call restoreMetaID  function
	trx, err := identitymanager.CallRestoreMetaID(reqParam.OldMetaID, reqParam.NewMetaID, reqParam.OldAddress, reqParam.NewAddress, reqParam.Signature)
	if err != nil {
		errObj := &internalError{err.Error()}

		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}
	resp.Result = trx.Hash().String()
	return
}

func revokeMetaID(req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	//var reqParam interface{}
	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	//1. Check paramter format
	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}
	reqParam := tmpParams.(metaIDRevokeParams)
	log.Println("parameter[MetaID] : ", reqParam.MetaID)
	log.Println("parameter[Timestamp] : ", reqParam.Timestamp)
	log.Println("parameter[Signature] : ", reqParam.Signature)

	//	2. Get Address from  Signature
	//  signdata = metaID|timestamp
	mSize := len(reqParam.MetaID)
	tSize := len(reqParam.MetaID) + len(reqParam.Timestamp)
	var signData hexutil.Bytes
	signData = make(hexutil.Bytes, tSize, tSize)
	copy(signData, reqParam.MetaID)
	copy(signData[mSize:tSize], reqParam.Timestamp)

	signedAddress, err := crypto.EcRecover(signData.String(), reqParam.Signature.String())
	if err != nil {
		log.Println("Error :", err)
		errObj := &invalidSignatureError{"Failed to verify signature"}
		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}

	// 3. Get Address from Ownerof function(Smart Contract)
	findAddress := &common.Address{}
	findAddress, err = identitymanager.CallOwnerOf(reqParam.MetaID)
	if err != nil {
		errObj := &internalError{err.Error()}

		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}
	if findAddress == nil {

		errObj := &notExistsError{"Cannot revoke Meta ID"}
		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}

	if !bytes.Equal(findAddress.Bytes(), signedAddress.Bytes()) {
		errObj := &invalidAddressError{"Cannot find valid user Address"}
		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}
	// 4. Call deleteMetaID function
	trx, err := identitymanager.CallDeleteMetaID(reqParam.MetaID, reqParam.Timestamp, reqParam.Signature)
	if err != nil {
		errObj := &internalError{err.Error()}

		resp.Error = &json.RPCError{
			Code:    errObj.ErrorCode(),
			Message: errObj.Error(),
		}
		return
	}
	resp.Result = trx.Hash().String()
	return
}

// HashContent Content For metaID
type HashContent struct {
	h hexutil.Bytes
}

// CalculateHash function to calculate hash
func (t HashContent) CalculateHash() []byte {
	return t.h
}

// Equals function to compare other Content
func (t HashContent) Equals(other merkletree.Content) bool {
	return t.h.String() == other.(HashContent).h.String()
}
