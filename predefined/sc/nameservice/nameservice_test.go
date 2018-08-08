package nameservice

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"bitbucket.org/coinplugin/proxy/crypto"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func defaultSetting() {

	path := "/tmp/testKey"

	file, err := os.Open("/tmp/testKeyPass")
	if err != nil {
		fmt.Printf("fail reading file : %s \n", err)
		panic("Fail to read key file")
	}
	defer file.Close()
	r := bufio.NewReader(file)
	data, _, err := r.ReadLine()
	passphrase := string(data)

	go func() { crypto.PathChan <- path }()
	go func() { crypto.PassphraseChan <- passphrase }()
	crypto.GetInstance()

}
func TestGetMetaIDContractAddress(t *testing.T) {
	client, err := ethclient.Dial("http://13.125.247.228:8545")
	if err != nil {
		fmt.Println(err)
	}

	address := common.HexToAddress("0x24211f4606e1e069df4a1f34ebd47f963a888ce2")
	instance, err := NewNameservice(address, client)

	if err != nil {
		fmt.Println(err)
	}
	var contractName [32]byte
	copy(contractName[:], "MetaID")

	result, err := instance.GetContractAddress(&bind.CallOpts{}, contractName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Meta ID Address: %x \n", result)
	fmt.Println("test END")
}

func TestGetMetadiumIdentityManagerContractAddress(t *testing.T) {
	client, err := ethclient.Dial("http://13.125.247.228:8545")
	if err != nil {
		fmt.Println(err)
	}

	address := common.HexToAddress("0x24211f4606e1e069df4a1f34ebd47f963a888ce2")
	instance, err := NewNameservice(address, client)

	if err != nil {
		fmt.Println(err)
	}
	var contractName [32]byte
	copy(contractName[:], "MetadiumIdentityManager")

	result, err := instance.GetContractAddress(&bind.CallOpts{}, contractName)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("MetadiumIdentity Address: %x \n", result)

}

func TestGetIDContractAddress(t *testing.T) {

	defaultSetting()
	address, err := GetIDContractAddress()
	if err != nil {
		fmt.Println(err)
		t.Error("error ", err)
	}
	fmt.Printf("ID Address: %x", *address)
}

func TestGetIMContractAddress(t *testing.T) {
	defaultSetting()
	address, err := GetIMContractAddress()
	if err != nil {
		fmt.Println(err)
		t.Error("error ", err)
	}
	fmt.Printf("IM Address: %x", *address)
}
