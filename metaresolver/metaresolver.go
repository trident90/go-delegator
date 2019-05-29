// Package metaservice manages metaservice functions for RPC request
package metaresolver

import (
	"fmt"
	"math/big"
	"runtime/debug"

	proxyCommon "github.com/metadium/go-delegator/common"
	"github.com/metadium/go-delegator/json"
	"github.com/metadium/go-delegator/log"
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

// Forward delivers RPCRequest to predefined function and returns that
func Forward(req json.RPCRequest) (resp json.RPCResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("Internal Error(Panic) : %s", debug.Stack())
			resp.Error = &json.RPCError{
				Code:    -32603,
				Message: "Internal Error",
			}
		}
	}()
	for k, v := range predefinedPaths {
		if k == req.Method {
			requestID := proxyCommon.RandomUint64()
			return v.(func(uint64, json.RPCRequest) (json.RPCResponse, error))(requestID, req)
		}
	}
	err = fmt.Errorf("predefined NOT FOUND")

	return resp, err
}

// Contains check if given path is in predefined or not
func Contains(path string) bool {
	for k := range predefinedPaths {
		if k == path {
			return true
		}
	}
	return false
}

var predefinedPaths = map[string]interface{}{
	"create_identity":                     createIdentity,
	"add_associated_address_delegated":    addAssociatedAddressDelegated,
	"remove_associated_address_delegated": removeAssociatedAddressDelegated,

	"add_key_delegated":     addKeyDelegated,
	"remove_key_delegated":  removeKeyDelegated,
	"remove_keys_delegated": removeKeysDelegated,

	"get_provider_addresses":        getProviderAddresses,
	"get_identity_registry_address": getIdentityRegistryAddress,
	"get_service_key_address":       getServiceKeyAddress,
	"get_service_key_all_addresses": getServiceKeyAllAddresses,
	"get_resolver_addresses":        getResolverAddresses,
	"get_all_service_addresses":     getAllServiceAddresses,

	"create_identity_test":                     createIdentityTest,
	"add_associated_address_delegated_test":    addAssociatedAddressDelegatedTest,
	"remove_associated_address_delegated_test": removeAssociatedAddressDelegatedTest,
	"add_key_delegated_test":                   addKeyDelegatedTest,
	"remove_key_delegated_test":                removeKeyDelegatedTest,
	"remove_keys_delegated_test":               removeKeysDelegatedTest,
}
