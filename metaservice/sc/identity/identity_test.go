package identity

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/metadium/go-delegator/crypto"
	"github.com/ethereum/go-ethereum/common"
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
func TestCallGetTransactionCount(t *testing.T) {
	defaultSetting()

	//mgtAddress := common.HexToAddress("0x961c20596e7ec441723fbb168461f4b51371d8aa") // SC: 0x1840c5049536874fcFfC62ed9E0fC6cE7b773c49
	scAddress := common.HexToAddress("0x1840c5049536874fcFfC62ed9E0fC6cE7b773c49")

	identity, err := GetInstance(scAddress)

	if err != nil {
		t.Error("Error get Identity", err)
	}

	nonce, err := CallGetTransactionCount(identity)

	if err != nil {
		t.Error("Error CallCreateMetaID", err)
	}
	fmt.Printf("Transaction count : %x \n", nonce)
}

func TestCallGetKey(t *testing.T) {
	defaultSetting()
	//mgtAddress := common.HexToAddress("0x0000000000000000000000000000000000000000")
	keyAddress := common.HexToAddress("0xd3Cb37aE6a81EbF1b5C3D9422636b0dB48767B72") // SC: 0x1840c5049536874fcFfC62ed9E0fC6cE7b773c49
	scAddress := common.HexToAddress("0xe052cb04e4fe4d3ca69d247b4eff2aff35613b0e")

	instance, err := GetInstance(scAddress)

	if err != nil {
		t.Error("Error get Identity", err)
	}
	keyBytes := CallAddrToKey(keyAddress)
	// if err != nil {
	// 	t.Error("Error addrToKey", err)
	// }
	if keyBytes == *emptyKeyVal {
		t.Error("Key empty")
	}
	fmt.Printf("mgt Key  : %x \n", keyBytes)
	metaKey, err := CallGetKey(instance, keyBytes)
	if err != nil {
		t.Error("Error getKey", err)
	}
	fmt.Printf("Get Key : %x \n", metaKey)
}

func TestCallGetKeysByPurpose(t *testing.T) {
	defaultSetting()

	//mgtAddress := common.HexToAddress("0x961c20596e7ec441723fbb168461f4b51371d8aa") // SC: 0x1840c5049536874fcFfC62ed9E0fC6cE7b773c49
	scAddress := common.HexToAddress("0xe052cb04e4fe4d3ca69d247b4eff2aff35613b0e")

	instance, err := GetInstance(scAddress)

	keys, err := CallGetKeysByPurpose(instance, common.Big2)
	if err != nil {
		t.Error("Error getKey", err)
	}
	fmt.Printf("GetKeysByPurpose : %x \n", keys)
	for _, p := range keys {
		//management
		fmt.Printf("%x \n", p)
	}

}

/*
func TestCallDelegatedExecute(t *testing.T) {
	defaultSetting()

	mgtAddress := common.HexToAddress("0x961c20596e7ec441723fbb168461f4b51371d8aa")
	scAddress := common.HexToAddress("0x1840c5049536874fcFfC62ed9E0fC6cE7b773c49")

	identity, err := GetInstance(scAddress)
	if err != nil {
		t.Error("Error CallCreateMetaID", err)
	}
	value := big.NewInt(0)
	data := hexutil.MustDecode("0x00")
	nonce := big.NewInt(0)
	sig := hexutil.MustDecode("0x00")


	trx, err := CallDelegatedExecute(identity, mgtAddress, scAddress, value, data, executeID, sig)
	if err != nil {
		t.Error("Error CallCreateMetaID", err)
	}
	cryptoIns := crypto.GetInstance()
	proxyAddress := common.HexToAddress(cryptoIns.GetAddress())

	scAddress := ethCrypto.CreateAddress(proxyAddress, trx.Nonce())
	fmt.Printf("Meta ID Address: %x \n", scAddress)
	fmt.Printf("trxid : %v", trx.Hash().String())
}
*/
