package nameservice

import (
	"fmt"
	"log"
	"os"
	"testing"

	"bitbucket.org/coinplugin/proxy/crypto"
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

func TestGetIDContractAddress(t *testing.T) {
	defaultSetting()
	address, err := GetIDContractAddress()
	if err != nil {
		log.Fatal(err)
		t.Error("error ", err)
	}
	fmt.Printf("ID Address: %x", *address)
}

func TestGetIMContractAddress(t *testing.T) {
	defaultSetting()
	address, err := GetIMContractAddress()
	if err != nil {
		log.Fatal(err)
		t.Error("error ", err)
	}
	fmt.Printf("IM Address: %x", *address)
}

func defaultSetting() {
	os.Setenv(crypto.Passphrase, "testtesttest")
	os.Setenv(crypto.Path, "/Users/ywshin/keyStore/UTC--2018-06-19T03-29-35.987669228Z--084f8293f1b047d3a217025b24cd7b5ace8fc657")
}
