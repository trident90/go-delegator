package metaresolver

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/metadium/go-delegator/crypto"
	"github.com/metadium/go-delegator/json"
	"github.com/metadium/go-delegator/log"
	"github.com/metadium/go-delegator/metaresolver/sc/identityregistry"
	"github.com/metadium/go-delegator/metaresolver/sc/servicekeyresolver"
)

func getIdentityRegistryAddress(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	log.Debugd(reqID, "Call getIdentityRegistryAddress Function")
	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	address, err := identityregistry.GetAddress()
	if err != nil {
		log.Errorfd(reqID, "getIdentityRegistryAddress Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	if address != nil {
		resp.Result = address.String()
	}
	return
}
func getServiceKeyAddress(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	log.Debugd(reqID, "Call getServiceKeyAddress Function")
	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	address, err := servicekeyresolver.GetAddress()
	if err != nil {
		log.Errorfd(reqID, "getServiceKeyAddress Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	if address != nil {
		resp.Result = address.String()
	}
	return
}
func getServiceKeyAllAddresses(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	log.Debugd(reqID, "Call getServiceKeyAllAddresses Function")
	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	addresses, err := servicekeyresolver.GetAddressList()
	if err != nil {
		log.Errorfd(reqID, "getServiceKeyAllAddresses Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	if addresses != nil {
		resp.Result = addresses
	}

	return
}
func getResolverAddresses(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	log.Debugd(reqID, "Call getResolverAddresses Function")
	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc
	resp.Result = resolverAddresses
	return
}
func getProviderAddresses(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	log.Debugd(reqID, "Call getProviderAddresses Function")
	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	resp.Result = providerAddresses
	return
}

func getAllServiceAddresses(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	log.Debugd(reqID, "Call getAllServiceAddresses Function")
	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	addresses, err := makeAllServiceAddressMap()
	if err != nil {
		log.Errorfd(reqID, "getAllServiceAddresses Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	if addresses != nil {
		resp.Result = addresses
	}
	return
}

// getAllServiceAddress  get All Contract Addresses for metaresovler
func makeAllServiceAddressMap() (map[string]interface{}, error) {

	irAddr, err := identityregistry.GetAddress()
	if err != nil {
		log.Error("getAllServiceAddress() ", err)
		return nil, err
	}

	skrAddr, err := servicekeyresolver.GetAddress()
	if err != nil {
		log.Error("getAllServiceAddress() ", err)
		return nil, err
	}

	skrAddrList, err := servicekeyresolver.GetAddressList()
	if err != nil {
		log.Error("getAllServiceAddress() ", err)
		return nil, err
	}

	result := map[string]interface{}{
		"identity_registry": irAddr,
		"providers":         providerAddresses,
		"resolvers":         resolverAddresses,
		"service_key":       skrAddr,
		"service_key_all":   skrAddrList,
	}
	return result, nil
}
func getDelegatorAddress() (*common.Address, error) {
	delegatorAddr := common.HexToAddress(crypto.GetInstance().GetAddress())
	return &delegatorAddr, nil
}
