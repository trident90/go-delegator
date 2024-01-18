package metaresolver

import (
	"fmt"

	"go-delegator/json"
	"go-delegator/log"
	"go-delegator/metaresolver/sc/servicekeyresolver"
)

func addKeyDelegated(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {

	log.Debugd(reqID, " Call addKeyDelegated Function")

	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc
	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	reqParam := tmpParams.(addKeyDelegatedParams)
	log.Debugfd(reqID, "parameter[ResolverAddress] : %x", reqParam.ResolverAddress)
	log.Debugfd(reqID, "parameter[AssociatedAddress] : %x", reqParam.AssociatedAddress)
	log.Debugfd(reqID, "parameter[Key] : %x", reqParam.Key)
	log.Debugfd(reqID, "parameter[Symbol] : %v", reqParam.Symbol)
	log.Debugfd(reqID, "parameter[V] : %v", reqParam.V)
	log.Debugfd(reqID, "parameter[R] : %v", reqParam.R)
	log.Debugfd(reqID, "parameter[S] : %v", reqParam.S)
	log.Debugfd(reqID, "parameter[Timestamp] : %v", reqParam.Timestamp)
	log.Debugd(reqID, "PASS - 01. Check Parameter")

	//2.  is in servicekeyresolver.GetAddressList
	isValidResovlerAddress := servicekeyresolver.ContainsInAddresses(reqParam.ResolverAddress)
	log.Debugfd(reqID, "Reslover Address[%x] is Valid : %v", reqParam.ResolverAddress, isValidResovlerAddress)
	if !isValidResovlerAddress {
		err := fmt.Errorf("Is not valid resolver address")
		errObj = &invalidAddressError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	//3. get instance ServiceKeyResolver
	instance, err := servicekeyresolver.GetInstance(reqParam.ResolverAddress)
	if err != nil {
		errObj := &notExistsAddressError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	if instance == nil {
		err = fmt.Errorf("Cannot get ServiceKeyResolver Instance")
		errObj := &notExistsAddressError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(reqID, "PASS - 02. Get ServiceKeyResolver")

	// 4. CallAddKeyDelegated
	var rBytes, sBytes [32]byte
	copy(rBytes[:], reqParam.R)
	copy(sBytes[:], reqParam.S)
	trx, err := servicekeyresolver.CallAddKeyDelegated(reqID, instance, reqParam.AssociatedAddress, reqParam.Key, reqParam.Symbol, reqParam.V[0], rBytes, sBytes, reqParam.Timestamp)
	if err != nil {
		log.Errorfd(reqID, "CallAddKeyDelegated Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugfd(reqID, "PASS 05. Call addKeyDelegated : %v", trx.Hash().String())

	//  return txid
	resp.Result = trx.Hash().String()
	return
}

func removeKeyDelegated(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	log.Debugd(reqID, " Call removeKeyDelegated Function")

	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc
	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	reqParam := tmpParams.(removeKeyDelegatedParams)
	log.Debugfd(reqID, "parameter[ResolverAddress] : %x", reqParam.ResolverAddress)
	log.Debugfd(reqID, "parameter[AssociatedAddress] : %x", reqParam.AssociatedAddress)
	log.Debugfd(reqID, "parameter[Key] : %x", reqParam.Key)
	log.Debugfd(reqID, "parameter[V] : %v", reqParam.V)
	log.Debugfd(reqID, "parameter[R] : %v", reqParam.R)
	log.Debugfd(reqID, "parameter[S] : %v", reqParam.S)
	log.Debugfd(reqID, "parameter[Timestamp] : %v", reqParam.Timestamp)
	log.Debugd(reqID, "PASS - 01. Check Parameter")

	//2.  is in servicekeyresolver.GetAddressList
	isValidResovlerAddress := servicekeyresolver.ContainsInAddresses(reqParam.ResolverAddress)
	log.Debugfd(reqID, "Reslover Address[%x] is Valid : %v", reqParam.ResolverAddress, isValidResovlerAddress)
	if !isValidResovlerAddress {
		err := fmt.Errorf("Is not valid resolver address")
		errObj = &invalidAddressError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	//3. get instance ServiceKeyResolver
	instance, err := servicekeyresolver.GetInstance(reqParam.ResolverAddress)
	if err != nil {
		errObj := &notExistsAddressError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	if instance == nil {
		err = fmt.Errorf("Cannot get ServiceKeyResolver Instance")
		errObj := &notExistsAddressError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(reqID, "PASS - 02. Get ServiceKeyResolver")

	// 4. CallDelegatedApprove
	var rBytes, sBytes [32]byte
	copy(rBytes[:], reqParam.R)
	copy(sBytes[:], reqParam.S)
	trx, err := servicekeyresolver.CallRemoveKeyDelegated(reqID, instance, reqParam.AssociatedAddress, reqParam.Key, reqParam.V[0], rBytes, sBytes, reqParam.Timestamp)
	if err != nil {
		log.Errorfd(reqID, "CallRemoveKeyDelegated Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugfd(reqID, "PASS 05. Call RemoveKeyDelegated : %v", trx.Hash().String())

	//return txid
	resp.Result = trx.Hash().String()

	return
}

func removeKeysDelegated(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	log.Debugd(reqID, " Call removeKeysDelegated Function")

	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc
	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	reqParam := tmpParams.(removeKeysDelegatedParams)
	log.Debugfd(reqID, "parameter[ResolverAddress] : %x", reqParam.ResolverAddress)
	log.Debugfd(reqID, "parameter[AssociatedAddress] : %x", reqParam.AssociatedAddress)
	log.Debugfd(reqID, "parameter[V] : %v", reqParam.V)
	log.Debugfd(reqID, "parameter[R] : %v", reqParam.R)
	log.Debugfd(reqID, "parameter[S] : %v", reqParam.S)
	log.Debugfd(reqID, "parameter[Timestamp] : %v", reqParam.Timestamp)
	log.Debugd(reqID, "PASS - 01. Check Parameter")

	//2.  is in servicekeyresolver.GetAddressList
	isValidResovlerAddress := servicekeyresolver.ContainsInAddresses(reqParam.ResolverAddress)
	log.Debugfd(reqID, "Reslover Address[%x] is Valid : %v", reqParam.ResolverAddress, isValidResovlerAddress)
	if !isValidResovlerAddress {
		err := fmt.Errorf("Is not valid resolver address")
		errObj = &invalidAddressError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	//3. get instance ServiceKeyResolver
	instance, err := servicekeyresolver.GetInstance(reqParam.ResolverAddress)
	if err != nil {
		errObj := &notExistsAddressError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	if instance == nil {
		err = fmt.Errorf("Cannot get ServiceKeyResolver Instance")
		errObj := &notExistsAddressError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(reqID, "PASS - 02. Get ServiceKeyResolver")

	// 4. CallDelegatedApprove
	var rBytes, sBytes [32]byte
	copy(rBytes[:], reqParam.R)
	copy(sBytes[:], reqParam.S)
	trx, err := servicekeyresolver.CallRemoveKeysDelegated(reqID, instance, reqParam.AssociatedAddress, reqParam.V[0], rBytes, sBytes, reqParam.Timestamp)
	if err != nil {
		log.Errorfd(reqID, "CallRemoveKeysDelegated Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugfd(reqID, "PASS 05. Call RemoveKeysDelegated : %v", trx.Hash().String())

	//return txid
	resp.Result = trx.Hash().String()

	return
}
