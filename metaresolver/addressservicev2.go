package metaresolver

import (
	"fmt"

	"go-delegator/config"
	"go-delegator/crypto"
	"go-delegator/json"
	"go-delegator/log"
	"go-delegator/metaresolver/sc/identityregistry"
	"go-delegator/metaresolver/sc/publickeyresolver"
	"go-delegator/metaresolver/sc/servicekeyresolver"

	"github.com/ethereum/go-ethereum/common"
)

func getIdentityRegistryAddress(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	log.Debugd(reqID, "Call getIdentityRegistryAddress Function")
	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	address := identityregistry.GetAddress()
	if address != nil {
		resp.Result = address
	} else {
		err := fmt.Errorf("address is nil")
		log.Errorfd(reqID, "getIdentityRegistryAddress Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
	}
	return
}
func getServiceKeyAddress(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	log.Debugd(reqID, "Call getServiceKeyAddress Function")
	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	address := servicekeyresolver.GetAddress()
	if address != nil {
		resp.Result = address
	} else {
		err := fmt.Errorf("address is nil")
		log.Errorfd(reqID, "getServiceKeyAddress Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
	}
	return
}
func getServiceKeyAllAddresses(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	log.Debugd(reqID, "Call getServiceKeyAllAddresses Function")
	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	addresses := servicekeyresolver.GetAddressList()
	resp.Result = addresses

	if addresses != nil {
		resp.Result = addresses
	} else {
		err := fmt.Errorf("address is nil")
		log.Errorfd(reqID, "getServiceKeyAllAddresses Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
	}

	return
}

func getPublicKeyAddress(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	log.Debugd(reqID, "Call getPublicKeyAddress Function")
	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	address := publickeyresolver.GetAddress()
	if address != nil {
		resp.Result = address
	} else {
		err := fmt.Errorf("address is nil")
		log.Errorfd(reqID, "getPublicKeyAddress Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
	}
	return
}
func getPublicKeyAllAddresses(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	log.Debugd(reqID, "Call getPublicKeyAllAddresses Function")
	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	addresses := publickeyresolver.GetAddressList()
	resp.Result = addresses

	if addresses != nil {
		resp.Result = addresses
	} else {
		err := fmt.Errorf("address is nil")
		log.Errorfd(reqID, "getPublicKeyAllAddresses Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
	}

	return
}

func getResolverAddresses(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	log.Debugd(reqID, "Call getResolverAddresses Function")
	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc
	resp.Result = makeResolverAddresses()
	return
}
func getProviderAddresses(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	log.Debugd(reqID, "Call getProviderAddresses Function")
	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	resp.Result = config.Config.ProviderAddresses
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

	irAddr := identityregistry.GetAddress()
	skrAddr := servicekeyresolver.GetAddress()
	skrAddrList := servicekeyresolver.GetAddressList()
	resolverAddrs := makeResolverAddresses()
	pkrAddr := publickeyresolver.GetAddress()
	pkrAddrList := publickeyresolver.GetAddressList()

	result := map[string]interface{}{
		"identity_registry": irAddr,
		"providers":         config.Config.ProviderAddresses,
		"resolvers":         resolverAddrs,
		"service_key":       skrAddr,
		"service_key_all":   skrAddrList,
		"public_key":        pkrAddr,
		"public_key_all":    pkrAddrList,
	}
	return result, nil
}
func getDelegatorAddress() (*common.Address, error) {
	delegatorAddr := common.HexToAddress(crypto.GetInstance().GetAddress())
	return &delegatorAddr, nil
}
func makeResolverAddresses() []common.Address {
	addrs := []common.Address{*servicekeyresolver.GetAddress(), *publickeyresolver.GetAddress()}
	return addrs
}
