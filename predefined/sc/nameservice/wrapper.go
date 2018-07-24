package nameservice

import (
	"fmt"
	"math/big"
	"sync"

	"bitbucket.org/coinplugin/proxy/log"

	"bitbucket.org/coinplugin/proxy/crypto"
	"bitbucket.org/coinplugin/proxy/rpc"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var once sync.Once
var session *NameserviceSession

func getSession() (*NameserviceSession, error) {
	once.Do(func() {
		service, err := getNSService()
		//auth := bind.NewKeyedTransactor(getPrivateKey())
		auth := crypto.GetTransactionOpts()
		auth.Value = big.NewInt(0)
		auth.GasLimit = uint64(300000)
		if err != nil {
			log.Error(err)
			return
		}
		session = &NameserviceSession{
			Contract: service,
			CallOpts: bind.CallOpts{
				Pending: true,
			},
			TransactOpts: *auth,
		}
	})
	if session == nil {
		err := fmt.Errorf("Cannot make session for NameService")
		return nil, err
	}
	return session, nil
}

func getNSService() (*Nameservice, error) {
	_rpc := rpc.GetInstance()
	client := _rpc.GetEthClient()

	instance, err := NewNameservice(nsAddress, client)
	if err != nil {
		return nil, err
	}
	return instance, nil
}

// GetIDContractAddress  retrun MetaID Contract Address
func GetIDContractAddress() (*common.Address, error) {

	session, err := getSession()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var contractName [32]byte
	copy(contractName[:], "MetaID")

	// result, err := instance.GetContractAddress(&bind.CallOpts{}, contractName)
	result, err := session.GetContractAddress(contractName)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Infof("Meta ID Address: %x ", result)
	return &result, nil
}

// GetIMContractAddress  retrun IdentityManager Contract Address
func GetIMContractAddress() (*common.Address, error) {
	session, err := getSession()
	if err != nil {
		log.Error(err)
	}
	var contractName [32]byte
	copy(contractName[:], "MetadiumIdentityManager")

	result, err := session.GetContractAddress(contractName)
	if err != nil {
		log.Error(err)
	}

	log.Infof("Metadium Identity Manager Address: %x ", result)
	return &result, nil
}
