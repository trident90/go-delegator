package metaservice

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"go-delegator/crypto"
	"go-delegator/ipfs"
	"go-delegator/json"
	"go-delegator/log"
	"go-delegator/metaservice/sc/identity"
	"go-delegator/metaservice/sc/identitymanager"
)

func createMetaID(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {

	log.Debugd(reqID, " Call createMetaID Function")

	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	reqParam := tmpParams.(metaIDCreateParams)
	// 1. check recaptcha
	if useReCaptcha {
		err := checkRecaptcha(reqParam.Recaptcha)
		if err != nil {
			errObj := &verifyRecaptchaError{err.Error()}
			resp.Error = makeErrorResponse(errObj)
		}
		log.Debugd(reqID, "Verified reCaptcha")
	} else {
		log.Debugd(reqID, "SKIP reCaptcha")
	}
	// 2. Verify signature
	_, errObj = verifySignature(reqID, reqParam.Address.String(), reqParam.Signature.String(), &reqParam.Address)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}

	// 3. CallCreateMetaID
	trx, err := identitymanager.CallCreateMetaID(reqParam.Address)
	if err != nil {
		log.Errorfd(reqID, "CallCreateMetaID Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugfd(reqID, "PASS - Call CreateMetaID : %v", trx.Hash().String())

	//  return txid
	resp.Result = trx.Hash().String()
	return
}

func delegatedExecute(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {

	log.Debugd(reqID, " Call delegatedExecute Function")

	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc
	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	reqParam := tmpParams.(metaIDExecuteParams)
	log.Debugfd(reqID, "parameter[MetaID] : %x", reqParam.MetaID)
	log.Debugfd(reqID, "parameter[from] : %x", reqParam.From)
	log.Debugfd(reqID, "parameter[to] : %x", reqParam.To)
	log.Debugfd(reqID, "parameter[data] : %v", reqParam.Data)
	log.Debugfd(reqID, "parameter[nonce] : %v", reqParam.Nonce)
	log.Debugfd(reqID, "parameter[value] : %v", reqParam.Value)
	log.Debugfd(reqID, "parameter[Signature] : %v", reqParam.Signature)
	log.Debugd(reqID, "PASS - 01. Check Parameter")

	//1. get instance MetaID
	instance, err := identity.GetInstance(reqParam.MetaID)
	if err != nil {
		errObj := &notExistsAddressError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	if instance == nil {
		err = fmt.Errorf("Cannot get MetaID Instance")
		errObj := &notExistsAddressError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(reqID, "PASS - 02. Get Identity")

	//2. validate signature
	//   - get sign address

	var executeSigData hexutil.Bytes
	valueBytes := intToByte32(reqParam.Value)
	nonceBytes := intToByte32(reqParam.Nonce)
	executeSigData = concatBytes(reqParam.To.Bytes(), valueBytes[:], reqParam.Data, nonceBytes[:])

	log.Debugd(reqID, "executeSigData : ", executeSigData.String())

	signAddr, errObj := verifySignature(reqID, hexutil.Encode(ethCrypto.Keccak256(executeSigData)), reqParam.Signature.String(), &reqParam.From)
	// if err != nil {
	// 	errObj := &internalError{err.Error()}
	// 	resp.Error = makeErrorResponse(errObj)
	// 	return
	// }
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(reqID, "PASS - 03. Verify Sign : ", signAddr.String())

	//3. Check permission
	err = identity.CheckDelegateExecutePermission(instance, reqParam.From, reqParam.To, reqParam.Data)
	if err != nil {
		errObj := &invalidPermissionError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(reqID, "PASS - 04. Check Permission ")

	// 4. CallDelegatedExecute
	trx, err := identity.CallDelegatedExecute(instance, reqParam.From, reqParam.To, reqParam.Value, reqParam.Data, reqParam.Nonce, reqParam.Signature)
	if err != nil {
		log.Errorfd(reqID, "CallDelegatedExecute Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugfd(reqID, "PASS 05. Call DelegatedExecute : %v", trx.Hash().String())

	//  return txid
	resp.Result = trx.Hash().String()
	return
}

func delegatedApprove(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	log.Debugd(reqID, " Call delegatedApprove Function")

	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc
	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	reqParam := tmpParams.(metaIDApproveParams)
	log.Debugfd(reqID, "parameter[MetaID] : %x", reqParam.MetaID)
	log.Debugfd(reqID, "parameter[from] : %x", reqParam.From)

	log.Debugfd(reqID, "parameter[id] : %v", reqParam.Id)
	log.Debugfd(reqID, "parameter[approve] : %v", reqParam.Approve)
	log.Debugfd(reqID, "parameter[nonce] : %v", reqParam.Nonce)
	log.Debugfd(reqID, "parameter[Signature] : %v", reqParam.Signature)
	log.Debugd(reqID, "PASS - 01. Check Parameter")

	//1. get instance MetaID
	instance, err := identity.GetInstance(reqParam.MetaID)
	if err != nil {
		errObj := &notExistsAddressError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	if instance == nil {
		err = fmt.Errorf("Cannot get MetaID Instance")
		errObj := &notExistsAddressError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(reqID, "PASS - 02. Get Identity")
	//2. validate signature
	//   - get sign address

	var executeSigData hexutil.Bytes
	// idBytes := intToByte32(reqParam.Id)
	idBigInt := new(big.Int)
	idBigInt.SetBytes(reqParam.Id)

	nonceBytes := intToByte32(reqParam.Nonce)
	_approve := hexutil.MustDecode("0x00")
	if reqParam.Approve {
		_approve = hexutil.MustDecode("0x01")
	}
	executeSigData = concatBytes(reqParam.Id, _approve[:], nonceBytes[:])

	log.Debugd(reqID, "executeSigData : ", executeSigData.String())

	signAddr, errObj := verifySignature(reqID, hexutil.Encode(ethCrypto.Keccak256(executeSigData)), reqParam.Signature.String(), &reqParam.From)
	// if err != nil {
	// 	errObj := &internalError{err.Error()}
	// 	resp.Error = makeErrorResponse(errObj)
	// 	return
	// }
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(reqID, "PASS - 03. Verify Sign : ", signAddr.String())

	//3. Check permission
	err = identity.CheckDelegateApprovePermission(instance, reqParam.From)
	if err != nil {
		errObj := &invalidPermissionError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(reqID, "PASS - 04. Check Permission ")

	// 4. CallDelegatedApprove
	trx, err := identity.CallDelegatedApprove(instance, reqParam.From, idBigInt, reqParam.Approve, reqParam.Nonce, reqParam.Signature)
	if err != nil {
		log.Errorfd(reqID, "CallDelegatedApprove Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugfd(reqID, "PASS 05. Call DelegatedApprove : %v", trx.Hash().String())

	//return txid
	resp.Result = trx.Hash().String()

	return
}

func backupUserData(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {

	log.Debugd(reqID, "Call backupUserData Function")
	//var reqParam interface{}
	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	//1. Check parameter format
	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	reqParam := tmpParams.(metaIDBackupParams)
	log.Debugfd(reqID, "parameter[Address] : %v", reqParam.Address.String())
	log.Debugfd(reqID, "parameter[MetaId] : %v", reqParam.MetaID)
	log.Debugfd(reqID, "parameter[EncData] : %v", reqParam.EncData)
	log.Debugfd(reqID, "parameter[Sig] : %v", reqParam.Signature)
	log.Debugd(reqID, "PASS - 01. Check Parameter")

	//2. check Signature
	_, errObj = verifySignature(reqID, reqParam.MetaID.String(), reqParam.Signature.String(), &reqParam.Address)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(reqID, "PASS - 02. Check Signature")
	//TODO: address가 metaID contract에 managementKey로 등록 되어 있는지 체크 여부 필요할지.
	/*
		//3. address == managementKey check permission
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
	*/
	//4. Save file to IPFS
	ins := ipfs.GetInstance()
	fileHash, err := ins.Add(reqParam.EncData.String())
	if err != nil {
		log.Errorfd(reqID, "ipfs.Add Error : %v", err)
		errObj := &saveBackupFileError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(reqID, "PASS - 04-1. ipfs add ", fileHash)
	err = ins.Pin(fileHash)
	if err != nil {
		log.Errorfd(reqID, "ipfs.Pin Error : %v", err)
		errObj := &saveBackupFileError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(reqID, "PASS - 04-2. ipfs Pin ", fileHash)
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
	log.Debugd(reqID, "PASS - 04-3. ipfs PinByCluster ", fileHash)
	//  return fileID
	resp.Result = fileHash

	return
}

func getUserData(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	//requestTmpID := proxyCommon.RandomUint64()
	log.Debugd(reqID, "Call getUserData Function")
	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	//1. Check parameter format
	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	reqParam := tmpParams.(metaIDGetUserDataParams)
	//log.Debugfd(requestTmpID, "parameter[NewAddress] : %v", reqParam.NewAddress.String())
	log.Debugfd(reqID, "parameter[FileID] : %v", reqParam.FileID)
	//log.Debugfd(requestTmpID, "parameter[Sig] : %v", reqParam.Signature)
	log.Debugd(reqID, "PASS - 01. Check Parameter")

	//2. GET User Data from IPFS
	//return data
	ins := ipfs.GetInstance()
	ret := ins.Cat(reqParam.FileID)
	if ret == "" {
		log.Debugd(reqID, "Error : Cannot find file ")
		errObj := &notFoundFileError{"not found backup file"}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(reqID, "PASS - 02. ipfs.Cat")
	// //  return fileID
	resp.Result = ret

	return
}

func verifySignature(reqID uint64, msg string, sig string, address *common.Address) (*common.Address, Error) {

	signedAddress, err := crypto.EcRecover(msg, sig)
	if err != nil {
		log.Errorfd(reqID, "EcRecover Error: %v", err)
		return nil, &internalError{"Failed to EcRecover"}
	}
	log.Debugfd(reqID, "PASS - Check Signature signAddress : %v", signedAddress.String())
	//비교할 address가 없을 경우, signedAddress만 리턴
	if address != nil {
		if !bytes.Equal(signedAddress.Bytes(), address.Bytes()) {
			log.Debugfd(reqID, "Error  :Sign Error %v / %v ", address.String(), signedAddress.String())
			return &signedAddress, &invalidSignatureError{"Failed to verify signature"}
		}
		log.Debugd(reqID, "PASS - signedAddress is valid")
	} else {
		log.Debugd(reqID, "PASS - return signedAddress")
	}

	return &signedAddress, nil
}
