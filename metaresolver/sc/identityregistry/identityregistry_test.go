package identityregistry

import (
	"bufio"
	"fmt"
	"os"
	"testing"
	"time"

	abi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
	data, _, _ := r.ReadLine()
	passphrase := string(data)

	go func() { crypto.PathChan <- path }()
	go func() { crypto.PassphraseChan <- passphrase }()
	crypto.GetInstance()

}

/*
func TestCallCreateIdentity(t *testing.T) {
	defaultSetting()

	associatedAddress := common.HexToAddress("0x961c20596e7ec441723fbb168461f4b51371d8aa") //e052cb04e4fe4d3ca69d247b4eff2aff35613b0e
	recoveryAddress := common.HexToAddress("0x961c20596e7ec441723fbb168461f4b51371d8aa")


	var executeSigData hexutil.Bytes
	valueBytes := intToByte32(reqParam.Value)
	nonceBytes := intToByte32(reqParam.Nonce)

	executeSigData = concatBytes("0x19","0x00",irAddress.Bytes(),"I authorize the creation of an Identity on my behalf.",recoveryAddress.Bytes(),associatedAddress.Bytes(), providers.Bytes(), resolvers.Bytes(), timestamp)
	keccak256(
		abi.encodePacked(
			byte(0x19), byte(0), address(this),
			"I authorize the creation of an Identity on my behalf.",
			recoveryAddress, associatedAddress, providers, resolvers, timestamp
		)
	)

	ethCrypto.Keccak256(executeSigData))


	//recoveryAddress common.Address, associatedAddress common.Address, providers []common.Address, resolvers []common.Address, v uint8, r [32]byte, s [32]byte, timestamp *big.Int
	var providers [1]common.Address
	provider[0] := crypto.GetInstance().GetAddress();

	trx, err := CallCreateIdentity(recoveryAddress, associatedAddress, providers, resolvers, v, r, s, timestamp)
	if err != nil {
		t.Error("Error CallCreateMetaID", err)
	}

	//	scAddress := ethCrypto.CreateAddress(*trx.To(), trx.Nonce())
	//	fmt.Printf("Meta ID Address: %x \n", scAddress)
	fmt.Printf("trxid : %v", trx.Hash().String())
}
*/
func TestCallCreateIdentity(t *testing.T) {
	//defaultSetting()

	bytes2Ty, _ := abi.NewType("bytes2", nil)
	addrTy, _ := abi.NewType("address", nil)
	strTy, _ := abi.NewType("string", nil)
	addrArrayTy, _ := abi.NewType("address[]", nil)
	uintTy, _ := abi.NewType("uint256", nil)
	arguments := abi.Arguments{
		{
			Type: bytes2Ty, //"0x19"
		},
		{
			Type: addrTy, //identityRegistry Address
		},
		{
			Type: strTy, //Message
		},
		{
			Type: addrTy, //recoveryAddress
		},
		{
			Type: addrTy, //associatedAddress
		},
		{
			Type: addrArrayTy, //providers
		},
		{
			Type: addrArrayTy, //resolvers
		},
		{
			Type: uintTy, //timestamp
		},
	}

	addrArrayArguments := abi.Arguments{
		{
			Type: addrArrayTy, //resolvers
		},
		{
			Type: addrArrayTy, //resolvers
		},
		{
			Type: uintTy, //timestamp
		},
	}
	fmt.Printf("TEST1 \n")
	associatedAddress := common.HexToAddress("0x961c20596e7ec441723fbb168461f4b51371d8aa") //e052cb04e4fe4d3ca69d247b4eff2aff35613b0e
	recoveryAddress := common.HexToAddress("0x961c20596e7ec441723fbb168461f4b51371d8aa")
	var providers [1]common.Address
	var resolvers []common.Address
	providers[0] = common.HexToAddress("0x961c20596e7ec441723fbb168461f4b51371d8aa")
	fmt.Printf("TEST2-providers: %v \n", providers)
	var addrPacked hexutil.Bytes

	timestamp := time.Now().UTC().Unix()
	fmt.Printf("TEST3-timestamp: %v \n", timestamp)
	addrPacked, _ = addrArrayArguments.Pack(providers, resolvers, timestamp)

	fmt.Printf("TEST2.1-provider/resolver Pack: %v \n", addrPacked.String())
	var signBytes hexutil.Bytes
	signBytes, _ = arguments.Pack(
		"0x1900",
		irAddress,
		"I authorize the creation of an Identity on my behalf.",
		recoveryAddress,
		associatedAddress,
		providers,
		resolvers,
		timestamp,
	)
	fmt.Printf("signBytes : %v", signBytes)
	// var executeSigData hexutil.Bytes
	// prefix, _ := hexutil.Decode("0x1900")
	// msg := []byte("I authorize the creation of an Identity on my behalf.")
	//executeSigData = concatBytes(prefix, irAddress.Bytes(), msg, recoveryAddress.Bytes(), associatedAddress.Bytes(), providers, resolvers, timestamp)
	// ethCrypto.Keccak256(executeSigData))
}
func concatBytes(args ...[]byte) []byte {
	var ret []byte
	for _, a := range args {
		ret = append(ret, a...)
	}
	return ret
}
