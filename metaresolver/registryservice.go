package metaresolver

import (
	"go-delegator/json"
	"go-delegator/log"
	"go-delegator/metaresolver/sc/identityregistry"
)

func createIdentity(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {

	log.Debugd(reqID, " Call createIdentity Function")

	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	// var providers [1]common.Address
	// delegatorAddr, _ := getDelegatorAddress()
	// providers[0] = *delegatorAddr
	reqParam := tmpParams.(identityCreateParams)

	log.Debugfd(reqID, "parameter[RecoveryAddress] : %x \n", reqParam.RecoveryAddress)
	log.Debugfd(reqID, "parameter[AssociatedAddress] : %x \n", reqParam.AssociatedAddress)
	log.Debugfd(reqID, "parameter[Providers] : %x \n", reqParam.Providers)
	log.Debugfd(reqID, "parameter[Resolvers] : %v \n", reqParam.Resolvers)
	log.Debugfd(reqID, "parameter[V] : %v", reqParam.V)
	log.Debugfd(reqID, "parameter[R] : %v", reqParam.R)
	log.Debugfd(reqID, "parameter[S] : %v", reqParam.S)
	log.Debugfd(reqID, "parameter[Timestamp] : %v", reqParam.Timestamp)

	log.Debugd(reqID, "PASS - 01. Check Parameter \n")
	// // 2. verify signature
	// identityRegistryAddr, _ := identityregistry.GetAddress()
	// keccakBytes, _ := reqParam.Keccak256(identityRegistryAddr)
	// errObj = verifySignature(reqID, keccakBytes, reqParam.V[0], reqParam.R, reqParam.S, reqParam.AssociatedAddress)
	// if errObj != nil {
	// 	resp.Error = makeErrorResponse(errObj)
	// 	return
	// }

	// 3. CallCreateIdentity
	var rBytes, sBytes [32]byte
	copy(rBytes[:], reqParam.R)
	copy(sBytes[:], reqParam.S)
	trx, err := identityregistry.CallCreateIdentity(reqID, reqParam.RecoveryAddress, reqParam.AssociatedAddress, reqParam.Providers, reqParam.Resolvers, reqParam.V[0], rBytes, sBytes, reqParam.Timestamp)
	if err != nil {
		log.Errorfd(reqID, "CallCreateIdentity Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugfd(reqID, "PASS - Call CreateIdentity : %v", trx.Hash().String())

	//  return txid
	resp.Result = trx.Hash().String()

	return
}

func addAssociatedAddressDelegated(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	log.Debugd(reqID, " Call addAssociatedAddressDelegated Function")

	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	reqParam := tmpParams.(addAssociatedAddressDelegatedParams)

	log.Debugfd(reqID, "parameter[ApprovingAddress] : %x \n", reqParam.ApprovingAddress)
	log.Debugfd(reqID, "parameter[AddressToAdd] : %x \n", reqParam.AddressToAdd)
	log.Debugfd(reqID, "parameter[V] : %v", reqParam.V)
	log.Debugfd(reqID, "parameter[R] : %v", reqParam.R)
	log.Debugfd(reqID, "parameter[S] : %v", reqParam.S)
	log.Debugfd(reqID, "parameter[Timestamp] : %v", reqParam.Timestamp)

	log.Debugd(reqID, "PASS - 01. Check Parameter \n")
	var vBytes [2]byte
	vBytes[0] = reqParam.V[0][0]
	vBytes[1] = reqParam.V[1][0]
	var rBytes, sBytes [2][32]byte
	copy(rBytes[0][:], reqParam.R[0])
	copy(rBytes[1][:], reqParam.R[1])
	copy(sBytes[0][:], reqParam.S[0])
	copy(sBytes[1][:], reqParam.S[1])

	trx, err := identityregistry.CallAddAssociatedAddressDelegated(reqID, reqParam.ApprovingAddress, reqParam.AddressToAdd, vBytes, rBytes, sBytes, reqParam.Timestamp)
	if err != nil {
		log.Errorfd(reqID, "CallAddAssociatedAddressDelegated Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugfd(reqID, "PASS - Call AddAssociatedAddressDelegated : %v", trx.Hash().String())

	//  return txid
	resp.Result = trx.Hash().String()

	return
}

func removeAssociatedAddressDelegated(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	log.Debugd(reqID, " Call removeAssociatedAddressDelegated Function")

	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	reqParam := tmpParams.(removeAssociatedAddressDelegatedParams)

	log.Debugfd(reqID, "parameter[AddressToRemove] : %x \n", reqParam.AddressToRemove)
	log.Debugfd(reqID, "parameter[V] : %v", reqParam.V)
	log.Debugfd(reqID, "parameter[R] : %v", reqParam.R)
	log.Debugfd(reqID, "parameter[S] : %v", reqParam.S)
	log.Debugfd(reqID, "parameter[Timestamp] : %v", reqParam.Timestamp)

	log.Debugd(reqID, "PASS - 01. Check Parameter \n")

	var rBytes, sBytes [32]byte
	copy(rBytes[:], reqParam.R)
	copy(sBytes[:], reqParam.S)
	trx, err := identityregistry.CallRemoveAssociatedAddressDelegated(reqID, reqParam.AddressToRemove, reqParam.V[0], rBytes, sBytes, reqParam.Timestamp)
	if err != nil {
		log.Errorfd(reqID, "CallRemoveAssociatedAddressDelegated Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugfd(reqID, "PASS - Call RemoveAssociatedAddressDelegated : %v", trx.Hash().String())

	//  return txid
	resp.Result = trx.Hash().String()

	return
}
