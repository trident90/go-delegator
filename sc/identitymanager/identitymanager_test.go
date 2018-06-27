package identitymanager

import (
	"context"
	"fmt"
	"log"
	"testing"

	"bitbucket.org/coinplugin/geth/common/hexutil"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestGetOwnerAddress(t *testing.T) {
	client, err := ethclient.Dial("http://13.125.247.228:8545")
	if err != nil {
		log.Fatal(err)
	}

	// privateKey, err := crypto.HexToECDSA("01b149603ca8f537bbb4e45d22e77df9054e50d826bb5f0a34e9ce460432b596")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//auth := bind.NewKeyedTransactor(privateKey)

	address := common.HexToAddress("0xc0b39803ae89ffa15b06a9b2784a1504b48eeb30")
	instance, err := NewIdentitymanager(address, client)

	if err != nil {
		log.Fatal(err)
	}
	var contractName [32]byte
	copy(contractName[:], "MetaID")
	metaID := hexutil.MustDecodeBig("0x1b442640e0333cb03054940e3cda07da982d2b57af68c3df8d0557b47a77d0bc")

	result, err := instance.OwnerOf(&bind.CallOpts{}, metaID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Onwer Address: %x \n", result)
}

func TestGetOwnerAddress2(t *testing.T) {
	var address *common.Address
	var err error

	metaID := hexutil.MustDecode("0x1b442640e0333cb03054940e3cda07da982d2b57af68c3df8d0557b47a77d0bc")

	address, err = GetOwnerAddress(metaID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Address : ", address.String())
}

func TestMakeMetaPackage(*testing.T) {
	version := byte(1)
	address := common.HexToAddress("0x961c20596e7ec441723fbb168461f4b51371d8aa")
	fmt.Printf("address : %v \n", address)
	mp := &metaPackageV1{
		Version:           version,
		UserSenderAddress: address,
	}
	bytes := mp.Serialize()
	fmt.Printf("bytes : %v", bytes)
}

/*
func TestCallCreateMetaID(t *testing.T) {

	address := common.HexToAddress("0x961c20596e7ec441723fbb168461f4b51371d8aa")
	metaID := hexutil.MustDecode("0xdb6dd8f5917a3c2f84a280f365ac137549e62d647b6cfba05a0f2c5e8e60e972")
	sig := hexutil.MustDecode("0x542e3d9af6758e80e4f2500395898a9cd9f5e3bd91b3053ad60d4cb0147b608c7de2ff608ae7528bc59cc6a7ba006020678d2ee5b5f88a97204734953cd2cdcb1b")
	txId := hexutil.MustDecode("0xaea34e618fe1defc7472d0b2a0cdffe3407d5cadee59f1d050d523e20d71ca49")
	fmt.Printf("txId : %v", txId)
	trx, err := CallCreateMetaID(metaID, sig, address)
	if err != nil {
		t.Error("Error CallCreateMetaID", err)
	}

	fmt.Printf("trxid : %v", trx.Hash().String())
}
*/

func TestGetResult(t *testing.T) {
	txHash := common.HexToHash("0x06d31f639409c8d4eb564fb790ae9d890d619784bfe17eaddf57f27bf9296e55")
	client, _ := getMetaClient()
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("txid ", tx.Hash().Hex())
	fmt.Println("isPending :", isPending)
	fmt.Println(tx.Data())

	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Status ", receipt.Status)
	fmt.Println("Status ", receipt.PostState)
	//fmt.Println("MarshalJson :", receipt.)

}
