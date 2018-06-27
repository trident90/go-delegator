package nameservice

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var nsAddress = common.HexToAddress("0x24211f4606e1e069df4a1f34ebd47f963a888ce2")

/*
func getSession(){
	service, err := getNSService()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	session := &NameserviceSession{
		Contract: service,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasLimit: big.NewInt(3141592),
		},
	}
}
*/
func getMetaClient() (*ethclient.Client, error) {
	//	if metaClient == nil {
	client, err := ethclient.Dial("http://13.125.247.228:8545")
	if err != nil {
		return nil, err
	}
	//		metaClient = client
	//	}
	return client, nil
}

// var nsService *Nameservice
func getNSService() (*Nameservice, error) {
	// if nsService == nil {
	client, err := getMetaClient()
	if err != nil {
		return nil, err
	}
	instance, err := NewNameservice(nsAddress, client)
	if err != nil {
		return nil, err
	}
	//	} else {
	return instance, nil
	//	}

}
func GetIDContractAddress() (*common.Address, error) {

	instance, err := getNSService()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var contractName [32]byte
	copy(contractName[:], "MetaID")

	result, err := instance.GetContractAddress(&bind.CallOpts{}, contractName)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Printf("Meta ID Address: %x \n", result)
	return &result, nil
}

func GetIMContractAddress() (*common.Address, error) {
	service, err := getNSService()
	if err != nil {
		log.Fatal(err)
	}
	var contractName [32]byte
	copy(contractName[:], "MetadiumIdentityManager")

	result, err := service.GetContractAddress(&bind.CallOpts{}, contractName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("MetadiumIdentity Address: %x \n", result)
	return &result, nil
}
