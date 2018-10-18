package registry

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"testing"

	"bitbucket.org/coinplugin/proxy/crypto"
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
func TestGetIMContractAddress(t *testing.T) {
	defaultSetting()
	result, err := GetIMContractAddress()
	if err != nil {
		t.Fatal(err)
	}
	address := common.HexToAddress("0x1318ca3cd01cfa04c7d129aa5f9cf7145736c0be")
	if !bytes.Equal(result.Bytes(), address.Bytes()) {
		t.Fatal("Not equal address", address.String())
	}
	fmt.Printf("Meta Identity Manager Address: %x \n", result)
	fmt.Println("test END")
}

func TestGetTRContractAddress(t *testing.T) {
	result, err := GetTRContractAddress()
	if err != nil {
		t.Fatal(err)
	}
	address := common.HexToAddress("0xb0143348314974f24e38bf9ca98db297028c64d9")
	if !bytes.Equal(result.Bytes(), address.Bytes()) {
		t.Fatal("Not equal address", address.String())
	}
	fmt.Printf("Topic Registry Address: %x \n", result)
	fmt.Println("test END")
}

func TestGetAARContractAddress(t *testing.T) {
	result, err := GetAARContractAddress()
	if err != nil {
		t.Fatal(err)
	}
	address := common.HexToAddress("0x0")
	if !bytes.Equal(result.Bytes(), address.Bytes()) {
		t.Fatal("Not equal address", address.String())
	}
	fmt.Printf("Attestation Agency Registry Address: %x \n", result)
	fmt.Println("test END")
}

func TestGetAMContractAddress(t *testing.T) {
	result, err := GetAMContractAddress()
	if err != nil {
		t.Fatal(err)
	}
	address := common.HexToAddress("0xa940417275057ee1732739a1dee5b4988fd878c0")
	if !bytes.Equal(result.Bytes(), address.Bytes()) {
		t.Fatal("Not equal address", address.String())
	}
	fmt.Printf("Achievement Manager Address: %x \n", result)
	fmt.Println("test END")
}

func TestGetAcContractAddress(t *testing.T) {
	result, err := GetAcContractAddress()
	if err != nil {
		t.Fatal(err)
	}
	address := common.HexToAddress("0x91f2778fb288482bad5d9cd7cdbed53a0577b8d0")
	if !bytes.Equal(result.Bytes(), address.Bytes()) {
		t.Fatal("Not equal address", address.String())
	}
	fmt.Printf("Achievement Address: %x \n", result)
	fmt.Println("test END")
}
