package nameservice

import (
	"fmt"
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestGetMetaIDContractAddress(t *testing.T) {
	client, err := ethclient.Dial("http://13.125.247.228:8545")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x24211f4606e1e069df4a1f34ebd47f963a888ce2")
	instance, err := NewNameservice(address, client)

	if err != nil {
		log.Fatal(err)
	}
	var contractName [32]byte
	copy(contractName[:], "MetaID")

	result, err := instance.GetContractAddress(&bind.CallOpts{}, contractName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Meta ID Address: %x \n", result)
	fmt.Println("test END")
}

func TestGetMetadiumIdentityManagerContractAddress(t *testing.T) {
	client, err := ethclient.Dial("http://13.125.247.228:8545")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x24211f4606e1e069df4a1f34ebd47f963a888ce2")
	instance, err := NewNameservice(address, client)

	if err != nil {
		log.Fatal(err)
	}
	var contractName [32]byte
	copy(contractName[:], "MetadiumIdentityManager")

	result, err := instance.GetContractAddress(&bind.CallOpts{}, contractName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("MetadiumIdentity Address: %x \n", result)

}

/*
func TestMultiAction(t *testing.T) {
	client, err := ethclient.Dial("http://13.125.247.228:8545")
	if err != nil {
		log.Fatal(err)
	}

	nsAddress := common.HexToAddress("0x24211f4606e1e069df4a1f34ebd47f963a888ce2")

	nsService, err := NewNameservice(nsAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	imAddress := common.HexToAddress("0xc0b39803ae89ffa15b06a9b2784a1504b48eeb30")
	imService, err := identitymanager.NewIdentitymanager(imAddress, client)

	if err != nil {
		log.Fatal(err)
	}
	go getMetaIDAddress(nsService)
	go getIMAddress(nsService)
	go getOwnerAddress(imService)
	go getMetaIDAddress(nsService)
	go getIMAddress(nsService)
	go getOwnerAddress(imService)
	go getMetaIDAddress(nsService)
	go getIMAddress(nsService)
	go getOwnerAddress(imService)
	go getMetaIDAddress(nsService)
	go getIMAddress(nsService)
	go getOwnerAddress(imService)
	go getMetaIDAddress(nsService)
	go getIMAddress(nsService)
	go getOwnerAddress(imService)
	go getMetaIDAddress(nsService)
	go getIMAddress(nsService)
	go getOwnerAddress(imService)

	time.Sleep(time.Second * 30)

}
func getMetaIDAddress(service *Nameservice) {
	var contractName [32]byte
	copy(contractName[:], "MetaID")

	result, err := service.GetContractAddress(&bind.CallOpts{}, contractName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Meta ID Address: %x \n", result)
}

func getIMAddress(service *Nameservice) {
	var contractName [32]byte
	copy(contractName[:], "MetadiumIdentityManager")

	result, err := service.GetContractAddress(&bind.CallOpts{}, contractName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("MetadiumIdentity Address: %x \n", result)
}

func getOwnerAddress(service *identitymanager.Identitymanager) {

	metaID := hexutil.MustDecodeBig("0x1b442640e0333cb03054940e3cda07da982d2b57af68c3df8d0557b47a77d0bc")

	result, err := service.OwnerOf(&bind.CallOpts{}, metaID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Onwer Address: %x \n", result)
}
*/

func TestGetIDContractAddress(t *testing.T) {
	address, err := GetIDContractAddress()
	if err != nil {
		log.Fatal(err)
		t.Error("error ", err)
	}
	fmt.Printf("ID Address: %x", *address)
}

func TestGetIMContractAddress(t *testing.T) {
	address, err := GetIMContractAddress()
	if err != nil {
		log.Fatal(err)
		t.Error("error ", err)
	}
	fmt.Printf("IM Address: %x", *address)
}
