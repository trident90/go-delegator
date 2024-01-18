package registry

import (
	"fmt"
	"sync"

	"go-delegator/log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"go-delegator/rpc"
)

var once sync.Once
var session *RegistrySession

func getSession() (*RegistrySession, error) {
	once.Do(func() {
		registry, err := getRegistry()

		// auth := crypto.GetTransactionOpts()
		// auth.Value = big.NewInt(0)
		// auth.GasLimit = uint64(300000)
		if err != nil {
			log.Error(err)
			return
		}
		session = &RegistrySession{
			Contract: registry,
			CallOpts: bind.CallOpts{
				Pending: false,
			},
			//			TransactOpts: *auth,
		}
	})
	if session == nil {
		err := fmt.Errorf("Cannot make session for Registry")
		return nil, err
	}
	return session, nil
}

func getRegistry() (*Registry, error) {
	_rpc := rpc.GetInstance()
	client := _rpc.GetEthClient()

	instance, err := NewRegistry(nsAddress, client)
	if err != nil {
		return nil, err
	}
	return instance, nil
}

//GetRegistryContractAddress get Registry Contract Address
func GetRegistryContractAddress() (*common.Address, error) {

	return &nsAddress, nil
}

func callGetContractAddress(name string) (*common.Address, error) {
	session, err := getSession()
	if err != nil {
		log.Error("callGetContractAddress() ", err)
		return nil, err
	}
	var contractName [32]byte
	copy(contractName[:], name)

	result, err := session.GetContractAddress(contractName)
	if err != nil {
		if err.Error() == "abi: unmarshalling empty output" {
			return nil, nil
		}
		log.Error("callGetContractAddress( ", name, " ): ", err)
		return nil, err
	}

	log.Infof("Get Contract Address( %s ): %x ", name, result)
	return &result, nil
}

// GetIMContractAddress  get IdentityManager Contract Address
func GetIMContractAddress() (*common.Address, error) {

	result, err := callGetContractAddress("IdentityManager")
	log.Infof("Identity Manager Address: %x ", result)
	return result, err
}

// GetTRContractAddress  get Topic Registry Contract Address
func GetTRContractAddress() (*common.Address, error) {

	result, err := callGetContractAddress("TopicRegistry")
	log.Infof("Topic Registry Address: %x ", result)
	return result, err
}

// GetAARContractAddress  get Attestation Agency Registry Contract Address
func GetAARContractAddress() (*common.Address, error) {
	result, err := callGetContractAddress("AttestationAgencyRegistry")
	log.Infof("Attestation Agency Registry Address: %x ", result)
	return result, err
}

// GetAMContractAddress  get Achievement Manager Contract Address
func GetAMContractAddress() (*common.Address, error) {
	result, err := callGetContractAddress("AchievementManager")
	log.Infof("Achievement Manager Address: %x ", result)
	return result, err
}

// GetAcContractAddress  get Achievement Contract Address
func GetAcContractAddress() (*common.Address, error) {
	result, err := callGetContractAddress("Achievement")
	log.Infof("Achievement Address: %x ", result)
	return result, err
}

// GetAllContractAddress  get All System Contract Addresses
func GetAllContractAddress() (map[string]*common.Address, error) {

	regAddr, err := GetRegistryContractAddress()
	if err != nil {
		log.Error("GetAllContractAddress() ", err)
		return nil, err
	}
	imAddr, err := GetIMContractAddress()
	if err != nil {
		log.Error("GetAllContractAddress() ", err)
		return nil, err
	}
	trAddr, err := GetTRContractAddress()
	if err != nil {
		log.Error("GetAllContractAddress() ", err)
		return nil, err
	}
	aarAddr, err := GetAARContractAddress()
	if err != nil {
		log.Error("GetAllContractAddress() ", err)
		return nil, err
	}
	amAddr, err := GetAMContractAddress()
	if err != nil {
		log.Error("GetAllContractAddress() ", err)
		return nil, err
	}
	acAddr, err := GetAcContractAddress()
	if err != nil {
		log.Error("GetAllContractAddress() ", err)
		return nil, err
	}

	result := map[string]*common.Address{
		"registry":                    regAddr,
		"identity_manager":            imAddr,
		"topic_registry":              trAddr,
		"attestation_agency_registry": aarAddr,
		"achievement_manager":         amAddr,
		"achievement":                 acAddr,
	}

	return result, nil
}
