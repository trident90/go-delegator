// Package metaservice manages metaservice functions for RPC request
package metaservice

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
	"create_meta_id":                          createMetaID,
	"delegated_execute":                       delegatedExecute,
	"delegated_approve":                       delegatedApprove,
	"backup_user_data":                        backupUserData,
	"get_user_data":                           getUserData,
	"get_registry_address":                    getRegistryAddress,
	"get_identity_manager_address":            getIdentityManagerAddress,
	"get_topic_registry_address":              getTopicRegistryAddress,
	"get_attestation_agency_registry_address": getAttestationAgencyRegistryAddress,
	"get_achievement_manager_address":         getAchievementManagerAddress,
	"get_achievement_address":                 getAchievementAddress,
	"get_all_system_address":                  getAllSystemAddress,
}
