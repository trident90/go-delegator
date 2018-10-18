package metaservice

import (
	"bitbucket.org/coinplugin/proxy/json"
	"bitbucket.org/coinplugin/proxy/log"
	"bitbucket.org/coinplugin/proxy/metaservice/sc/registry"
)

func getRegistryAddress(reqId uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	log.Debugd(reqId, "Call getRegistryAddress Function")
	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	address, err := registry.GetRegistryContractAddress()
	if err != nil {
		log.Errorfd(reqId, "getRegistryAddress Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	if address != nil {
		resp.Result = address.String()
	}
	return
}
func getIdentityManagerAddress(reqId uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	log.Debugd(reqId, "Call getIdentityManagerAddress Function")
	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	address, err := registry.GetIMContractAddress()
	if err != nil {
		log.Errorfd(reqId, "getIdentityManagerAddress Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	if address != nil {
		resp.Result = address.String()
	}
	return
}

func getTopicRegistryAddress(reqId uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	log.Debugd(reqId, "Call getTopicRegistryAddress Function")
	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	address, err := registry.GetTRContractAddress()
	if err != nil {
		log.Errorfd(reqId, "getTopicRegistryAddress Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	if address != nil {
		resp.Result = address.String()
	}
	return
}

func getAttestationAgencyRegistryAddress(reqId uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	log.Debugd(reqId, "Call getAttestationAgencyRegistryAddress Function")
	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	address, err := registry.GetAARContractAddress()
	if err != nil {
		log.Errorfd(reqId, "getAttestationAgencyRegistryAddress Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	if address != nil {
		resp.Result = address.String()
	}
	return
}

func getAchievementManagerAddress(reqId uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	log.Debugd(reqId, "Call getAchievementManagerAddress Function")
	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	address, err := registry.GetAMContractAddress()
	if err != nil {
		log.Errorfd(reqId, "getAchievementManagerAddress Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	if address != nil {
		resp.Result = address.String()
	}
	return
}

func getAchievementAddress(reqId uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	log.Debugd(reqId, "Call getAchievementAddress Function")
	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	address, err := registry.GetAcContractAddress()
	if err != nil {
		log.Errorfd(reqId, "getAchievementAddress Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	if address != nil {
		resp.Result = address.String()
	}
	return
}

func getAllSystemAddress(reqId uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	log.Debugd(reqId, "Call getAllSystemAddress Function")
	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	addresses, err := registry.GetAllContractAddress()
	if err != nil {
		log.Errorfd(reqId, "getAllSystemAddress Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	if addresses != nil {
		resp.Result = addresses
	}
	return
}
