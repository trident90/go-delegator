package nameservice

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/big"
	"sync"

	"bitbucket.org/coinplugin/proxy/crypto"
	"bitbucket.org/coinplugin/proxy/rpc"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var nsAddress = common.HexToAddress("0xb2cdd687dc0805f003a7020765886db5136b0db7")

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
			log.Print(err)
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
		log.Print(err)
		return nil, err
	}

	var contractName [32]byte
	copy(contractName[:], "MetaID")

	// result, err := instance.GetContractAddress(&bind.CallOpts{}, contractName)
	result, err := session.GetContractAddress(contractName)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	log.Printf("Meta ID Address: %x \n", result)
	return &result, nil
}

// GetIMContractAddress  retrun IdentityManager Contract Address
func GetIMContractAddress() (*common.Address, error) {
	session, err := getSession()
	if err != nil {
		log.Print(err)
	}
	var contractName [32]byte
	copy(contractName[:], "MetadiumIdentityManager")

	result, err := session.GetContractAddress(contractName)
	if err != nil {
		log.Print(err)
	}

	fmt.Printf("MetadiumIdentity Address: %x \n", result)
	return &result, nil
}
