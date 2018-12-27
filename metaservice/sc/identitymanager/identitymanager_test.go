package identitymanager

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/metadium/go-delegator/crypto"
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

func TestCallCreateMetaID(t *testing.T) {
	defaultSetting()

	address := common.HexToAddress("0x961c20596e7ec441723fbb168461f4b51371d8aa") //e052cb04e4fe4d3ca69d247b4eff2aff35613b0e
	trx, err := CallCreateMetaID(address)
	if err != nil {
		t.Error("Error CallCreateMetaID", err)
	}

	scAddress := ethCrypto.CreateAddress(*trx.To(), trx.Nonce())
	fmt.Printf("Meta ID Address: %x \n", scAddress)
	fmt.Printf("trxid : %v", trx.Hash().String())
}

func TestCallGetDeployedMetaIds(t *testing.T) {
	defaultSetting()
	addrs, err := CallGetDeployedMetaIds()

	if err != nil {
		t.Error("Error - ", err)
	}
	fmt.Println(addrs)
	for _, p := range addrs {
		//management
		fmt.Printf("%x \n", p)
	}

}

func TestCreateAddress(t *testing.T) {
	address := common.HexToAddress("0x1318ca3CD01CFa04C7D129Aa5f9cF7145736C0be")
	scAddress := ethCrypto.CreateAddress(address, 1)
	fmt.Printf("Meta ID Address: %x \n", scAddress)
}
