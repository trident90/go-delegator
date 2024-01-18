package metaresolver

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"go-delegator/crypto"
	"go-delegator/json"
	"go-delegator/log"
	"go-delegator/metaresolver/sc/identityregistry"
	"go-delegator/metaresolver/sc/publickeyresolver"
)

func addPublicKeyDelegated(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {

	log.Debugd(reqID, " Call addPublicKeyDelegated Function")

	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc
	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	reqParam := tmpParams.(addPublicKeyDelegatedParams)
	log.Debugfd(reqID, "parameter[ResolverAddress] : %x", reqParam.ResolverAddress)
	log.Debugfd(reqID, "parameter[AssociatedAddress] : %x", reqParam.AssociatedAddress)
	log.Debugfd(reqID, "parameter[PublicKey] : %x", reqParam.PublicKey)
	log.Debugfd(reqID, "parameter[V] : %v", reqParam.V)
	log.Debugfd(reqID, "parameter[R] : %v", reqParam.R)
	log.Debugfd(reqID, "parameter[S] : %v", reqParam.S)
	log.Debugfd(reqID, "parameter[Timestamp] : %v", reqParam.Timestamp)
	log.Debugd(reqID, "PASS - 01. Check Parameter")

	//2.  is in publickeyresolver.GetAddressList
	isValidResovlerAddress := publickeyresolver.ContainsInAddresses(reqParam.ResolverAddress)
	log.Debugfd(reqID, "Reslover Address[%x] is Valid : %v", reqParam.ResolverAddress, isValidResovlerAddress)
	if !isValidResovlerAddress {
		err := fmt.Errorf("Is not valid resolver address")
		errObj = &invalidAddressError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	//3. GET EIN
	ein, err := identityregistry.CallGetEIN(reqID, reqParam.AssociatedAddress)
	if err != nil {
		errObj = &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	log.Debugfd(reqID, "EIN is %v", ein)

	//4. Check IsProviderFor
	providerAddr := common.HexToAddress(crypto.GetInstance().GetAddress())

	isProvider, err := identityregistry.CallIsProviderFor(reqID, ein, providerAddr)
	if err != nil {
		errObj = &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	if !isProvider {
		err = fmt.Errorf("Is not provider Address for user")
		errObj = &invalidAddressError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	//5. get isResolverFor
	isResolver, err := identityregistry.CallIsResolverFor(reqID, ein, reqParam.ResolverAddress)
	if err != nil {
		errObj = &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	if !isResolver {
		log.Infofd(reqID, "publicKeyResolver address is not added in identity(%v) [%x]", ein, reqParam.AssociatedAddress)
		// err = fmt.Errorf("Is not resolver for user")
		// errObj = &internalError{err.Error()}
		// resp.Error = makeErrorResponse(errObj)
		// return

		// 5-1 Add PublicKeyResolver address to resolvers
		tx, err := identityregistry.CallAddResolversFor(reqID, ein, []common.Address{reqParam.ResolverAddress})
		if err != nil {
			errObj = &internalError{err.Error()}
			resp.Error = makeErrorResponse(errObj)
			return
		}
		log.Infofd(reqID, "Transaction for Adding Resolver for (%v) : %x", ein, tx)
	}

	//6. get instance PublicKeyResolver
	instance, err := publickeyresolver.GetInstance(reqParam.ResolverAddress)
	if err != nil {
		errObj := &notExistsAddressError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	if instance == nil {
		err = fmt.Errorf("Cannot get PublicKeyResolver Instance")
		errObj := &notExistsAddressError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(reqID, "PASS - 02. Get PublicKeyResolver")

	//7. CallAddPublicKeyDelegated
	var rBytes, sBytes [32]byte
	copy(rBytes[:], reqParam.R)
	copy(sBytes[:], reqParam.S)
	trx, err := publickeyresolver.CallAddPublicKeyDelegated(reqID, instance, reqParam.AssociatedAddress, reqParam.PublicKey, reqParam.V[0], rBytes, sBytes, reqParam.Timestamp)
	if err != nil {
		log.Errorfd(reqID, "CallAddPublicKeyDelegated Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugfd(reqID, "PASS 05. Call addPublicKeyDelegated : %v", trx.Hash().String())

	//  return txid
	resp.Result = trx.Hash().String()
	return
}

func removePublicKeyDelegated(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	log.Debugd(reqID, " Call removePublicKeyDelegated Function")

	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc
	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	reqParam := tmpParams.(removePublicKeyDelegatedParams)
	log.Debugfd(reqID, "parameter[ResolverAddress] : %x", reqParam.ResolverAddress)
	log.Debugfd(reqID, "parameter[AssociatedAddress] : %x", reqParam.AssociatedAddress)
	log.Debugfd(reqID, "parameter[V] : %v", reqParam.V)
	log.Debugfd(reqID, "parameter[R] : %v", reqParam.R)
	log.Debugfd(reqID, "parameter[S] : %v", reqParam.S)
	log.Debugfd(reqID, "parameter[Timestamp] : %v", reqParam.Timestamp)
	log.Debugd(reqID, "PASS - 01. Check Parameter")

	//2.  is in publickeyresolver.GetAddressList
	isValidResovlerAddress := publickeyresolver.ContainsInAddresses(reqParam.ResolverAddress)
	log.Debugfd(reqID, "Reslover Address[%x] is Valid : %v", reqParam.ResolverAddress, isValidResovlerAddress)
	if !isValidResovlerAddress {
		err := fmt.Errorf("Is not valid resolver address")
		errObj = &invalidAddressError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	//3. get instance PublicKeyResolver
	instance, err := publickeyresolver.GetInstance(reqParam.ResolverAddress)
	if err != nil {
		errObj := &notExistsAddressError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	if instance == nil {
		err = fmt.Errorf("Cannot get PublicKeyResolver Instance")
		errObj := &notExistsAddressError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(reqID, "PASS - 02. Get PublicKeyResolver")

	// 4. CallDelegatedApprove
	var rBytes, sBytes [32]byte
	copy(rBytes[:], reqParam.R)
	copy(sBytes[:], reqParam.S)
	trx, err := publickeyresolver.CallRemovePublicKeyDelegated(reqID, instance, reqParam.AssociatedAddress, reqParam.V[0], rBytes, sBytes, reqParam.Timestamp)
	if err != nil {
		log.Errorfd(reqID, "CallRemovePublicKeyDelegated Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugfd(reqID, "PASS 05. Call RemovePublicKeyDelegated : %v", trx.Hash().String())

	//return txid
	resp.Result = trx.Hash().String()

	return
}
