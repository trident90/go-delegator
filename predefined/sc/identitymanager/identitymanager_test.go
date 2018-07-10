package identitymanager

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"

	"bitbucket.org/coinplugin/proxy/crypto"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func defaultSetting() {
	path := "/tmp/testKey"

	file, err := os.Open("/tmp/testKeyPass")
	if err != nil {
		log.Panicf("fail reading file : %s", err)
	}
	defer file.Close()
	r := bufio.NewReader(file)
	data, _, err := r.ReadLine()
	passphrase := string(data)

	go func() { crypto.PathChan <- path }()
	go func() { crypto.PassphraseChan <- passphrase }()
	crypto.GetInstance()

}

/*
func TestGetOwnerAddress2(t *testing.T) {
	client, err := ethclient.Dial("http://13.125.247.228:8545")
	if err != nil {
		log.Print(err)
	}
	address := common.HexToAddress("0xc0b39803ae89ffa15b06a9b2784a1504b48eeb30")
	//address := common.HexToAddress("0x71f446d16475d80322e880121dc7996512172031")
	instance, err := NewIdentitymanager(address, client)

	if err != nil {
		log.Print(err)
	}

	var contractName [32]byte
	copy(contractName[:], "MetaID")
	//metaID := hexutil.MustDecodeBig("0x1b442640e0333cb03054940e3cda07da982d2b57af68c3df8d0557b47a77d0bc")
	metaID := hexutil.MustDecodeBig("0xdb6dd8f5917a3c2f84a280f365ac137549e62d647b6cfba05a0f2c5e8e60e972")
	result, err := instance.OwnerOf(&bind.CallOpts{}, metaID)
	if err != nil {
		if err.Error() == "abi: unmarshalling empty output" {
			fmt.Println("not registered ID")
		} else {
			log.Print(err)
		}
	}
	fmt.Printf("Onwer Address: %x \n", result)
}
*/

func TestGetOwnerAddress(t *testing.T) {
	var address *common.Address
	var err error
	defaultSetting()

	// metaID := hexutil.MustDecode("0x1b442640e0333cb03054940e3cda07da982d2b57af68c3df8d0557b47a77d0bc")		//084f8293f1b047d3a217025b24cd7b5ace8fc657
	//metaID := hexutil.MustDecode("0xdb6dd8f5917a3c2f84a280f365ac137549e62d647b6cfba05a0f2c5e8e60e972") //created 961c20596e7ec441723fbb168461f4b51371d8aa
	//metaID := hexutil.MustDecode("0xc30255d8455a90fbb0ce11405b501624be9287370a41779312944c8739a9f79d") //Updated 961c20596e7ec441723fbb168461f4b51371d8aa
	//metaID := hexutil.MustDecode("0x628defba1494674ff937a7b88c453e25ae00b89dc925819007fa6f13a505a74c")
	metaID := hexutil.MustDecode("0x400c2c9ae5f66133fb6a48d476756b9048a345a0666676e416cdaddf93ae09c5")
	address, err = CallOwnerOf(metaID)
	if err != nil {
		log.Print(err)
	}

	if address != nil {

		fmt.Println("Address : ", address.String()) //, address.Hex(), hex.EncodeToString(address.Bytes())
	} else {
		fmt.Println("Address not Exists ")
	}
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

func TestCallCreateMetaID(t *testing.T) {
	defaultSetting()

	address := common.HexToAddress("0x961c20596e7ec441723fbb168461f4b51371d8aa")
	metaID := hexutil.MustDecode("0xdb6dd8f5917a3c2f84a280f365ac137549e62d647b6cfba05a0f2c5e8e60e972")
	sig := hexutil.MustDecode("0x542e3d9af6758e80e4f2500395898a9cd9f5e3bd91b3053ad60d4cb0147b608c7de2ff608ae7528bc59cc6a7ba006020678d2ee5b5f88a97204734953cd2cccb1b")
	//sig := hexutil.MustDecode("0x94b7ca242f431df15ba6c1019f3e4e52b8e1c177615b254140456d7814d202751224329073bf22bd4779d98dbaa02c65c9e1ac867932f2828197e012cffd45ef1c")
	//txId := hexutil.MustDecode("0xaea34e618fe1defc7472d0b2a0cdffe3407d5cadee59f1d050d523e20d71ca49")
	//fmt.Printf("txId : %v", txId)

	trx, err := CallCreateMetaID(metaID, sig, address)
	if err != nil {
		t.Error("Error CallCreateMetaID", err)
	}

	fmt.Printf("trxid : %v", trx.Hash().String())
}

func TestCallUpdateMetaID(t *testing.T) {
	defaultSetting()

	address := common.HexToAddress("0x961c20596e7ec441723fbb168461f4b51371d8aa")
	oldMetaID := hexutil.MustDecode("0xdb6dd8f5917a3c2f84a280f365ac137549e62d647b6cfba05a0f2c5e8e60e972")
	newMetaID := hexutil.MustDecode("0xc30255d8455a90fbb0ce11405b501624be9287370a41779312944c8739a9f79d")
	sig := hexutil.MustDecode("0x0f86138d24ecd75efd982e8c68555c97c70d5abfa67ab99edeca462b9b546d706888437eecc91f1ea5a0e815606188b6e91353bd867c9ad08cf2f62781aefe391c")

	trx, err := CallUpdateMetaID(oldMetaID, newMetaID, sig, address)
	if err != nil {
		t.Error("Error CallCreateMetaID", err)
	}

	fmt.Printf("trxid : %v", trx.Hash().String())
}

func TestCallDeleteMetaID(t *testing.T) {
	defaultSetting()
	metaID := hexutil.MustDecode("0xdb6dd8f5917a3c2f84a280f365ac137549e62d647b6cfba05a0f2c5e8e60e972")
	timestamp := hexutil.Bytes("20180703142012")
	sig := hexutil.MustDecode("0x3f593b366bf85cc98c877cc0c7e406017ced62b2913494457597bf78a65e8fd26023727bd2e9f7ac5dadf394744476a69fd43f5cbcc30c5b5987d6404ba4a4321b")
	trx, err := CallDeleteMetaID(metaID, timestamp, sig)
	if err != nil {
		t.Error("Error CallCreateMetaID", err)
	}
	//metaID|timestamp
	//0xdb6dd8f5917a3c2f84a280f365ac137549e62d647b6cfba05a0f2c5e8e60e9723230313830373033313432303132
	fmt.Printf("trxid : %v", trx.Hash().String())
}

/*
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
	fmt.Println("LOG", receipt.Logs)
	//fmt.Println("MarshalJson :", receipt.)


}
*/

func TestEcverify(t *testing.T) {
	defaultSetting()

	address := common.HexToAddress("0x961c20596e7ec441723fbb168461f4b51371d8aa")
	metaID := hexutil.MustDecode("0xdb6dd8f5917a3c2f84a280f365ac137549e62d647b6cfba05a0f2c5e8e60e972")
	//sig := hexutil.MustDecode("0x542e3d9af6758e80e4f2500395898a9cd9f5e3bd91b3053ad60d4cb0147b608c7de2ff608ae7528bc59cc6a7ba006020678d2ee5b5f88a97204734953cd2cccb1b")
	sig := hexutil.MustDecode("0x94b7ca242f431df15ba6c1019f3e4e52b8e1c177615b254140456d7814d202751224329073bf22bd4779d98dbaa02c65c9e1ac867932f2828197e012cffd45ef1c")
	// res, err := CallEcverify(metaID, sig, address)
	// if err != nil {
	// 	log.Print(err)
	// }
	// fmt.Println(res)

	service, err := getService()

	if err != nil {
		log.Print(err)
	}
	var msgBytes [32]byte
	copy(msgBytes[:], metaID)

	result, err1 := service.Ecverify(&bind.CallOpts{}, msgBytes, sig, address)

	if err1 != nil {
		log.Print(err1)

	}
	fmt.Printf("Ecverfiy: %v \n", result)

}

// func TestPubKey(t *testing.T) {

// 	ct := crypto.GetInstance()
// 	fmt.Println(ct.Address)
// }
