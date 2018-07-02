package nameservice

import (
	"fmt"
	"log"
	"math/big"
	"sync"

	"bitbucket.org/coinplugin/proxy/crypto"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var nsAddress = common.HexToAddress("0xe06dbbcb9af642017012d5021d586536673b955f")

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
			log.Fatal(err)
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

func getMetaClient() (*ethclient.Client, error) {
	client, err := ethclient.Dial("http://13.125.247.228:8545")
	if err != nil {
		return nil, err
	}
	return client, nil
}

// var nsService *Nameservice
func getNSService() (*Nameservice, error) {

	client, err := getMetaClient()
	if err != nil {
		return nil, err
	}
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
		log.Fatal(err)
		return nil, err
	}

	var contractName [32]byte
	copy(contractName[:], "MetaID")

	// result, err := instance.GetContractAddress(&bind.CallOpts{}, contractName)
	result, err := session.GetContractAddress(contractName)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Printf("Meta ID Address: %x \n", result)
	return &result, nil
}

// GetIMContractAddress  retrun IdentityManager Contract Address
func GetIMContractAddress() (*common.Address, error) {
	session, err := getSession()
	if err != nil {
		log.Fatal(err)
	}
	var contractName [32]byte
	copy(contractName[:], "MetadiumIdentityManager")

	result, err := session.GetContractAddress(contractName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("MetadiumIdentity Address: %x \n", result)
	return &result, nil
}
