package metaresolver

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/metadium/go-delegator/crypto"
	"github.com/metadium/go-delegator/json"
	"github.com/metadium/go-delegator/metaresolver/sc/identityregistry"
	"github.com/metadium/go-delegator/metaresolver/sc/servicekeyresolver"
	log "github.com/sirupsen/logrus"
)

func defaultSetting() {

	path := "/tmp/testKey"

	file, err := os.Open("/tmp/testKeyPass")
	if err != nil {
		log.Panicf("fail reading file : %s", err)
	}
	defer file.Close()
	r := bufio.NewReader(file)
	data, _, _ := r.ReadLine()
	passphrase := string(data)

	go func() { crypto.PathChan <- path }()
	go func() { crypto.PassphraseChan <- passphrase }()
	crypto.GetInstance()

}
func TestGetProviderAddresses(t *testing.T) {
	defaultSetting()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "get_provider_addresses",
		ID:      1,
	}

	//param := map[string]interface{}{}
	//req.Params = append(req.Params, param)
	fmt.Println("request : ", req.String())

	resp, err := Forward(req)
	if resp.String() == "" || err != nil {
		fmt.Println("Error : ", err)
		t.Errorf("Failed to start main")
	}
	fmt.Println("Response : ", resp.String())
}

func TestGetIdentityRegistryAddress(t *testing.T) {
	defaultSetting()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "get_identity_registry_address",
		ID:      1,
	}

	//param := map[string]interface{}{}
	//req.Params = append(req.Params, param)
	fmt.Println("request : ", req.String())

	resp, err := Forward(req)
	if resp.String() == "" || err != nil {
		fmt.Println("Error : ", err)
		t.Errorf("Failed to start main")
	}
	fmt.Println("Response : ", resp.String())
}
func TestGetServiceKeyAddress(t *testing.T) {
	defaultSetting()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "get_service_key_address",
		ID:      1,
	}

	//param := map[string]interface{}{}
	//req.Params = append(req.Params, param)
	fmt.Println("request : ", req.String())

	resp, err := Forward(req)
	if resp.String() == "" || err != nil {
		fmt.Println("Error : ", err)
		t.Errorf("Failed to start main")
	}
	fmt.Println("Response : ", resp.String())
}
func TestGetServiceKeyAllAddresses(t *testing.T) {
	defaultSetting()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "get_service_key_all_addresses",
		ID:      1,
	}

	//param := map[string]interface{}{}
	//req.Params = append(req.Params, param)
	fmt.Println("request : ", req.String())

	resp, err := Forward(req)
	if resp.String() == "" || err != nil {
		fmt.Println("Error : ", err)
		t.Errorf("Failed to start main")
	}
	fmt.Println("Response : ", resp.String())
}
func TestGetResolverAddresses(t *testing.T) {
	defaultSetting()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "get_resolver_addresses",
		ID:      1,
	}

	//param := map[string]interface{}{}
	//req.Params = append(req.Params, param)
	fmt.Println("request : ", req.String())

	resp, err := Forward(req)
	if resp.String() == "" || err != nil {
		fmt.Println("Error : ", err)
		t.Errorf("Failed to start main")
	}
	fmt.Println("Response : ", resp.String())
}

func TestGetAllServiceAddresses(t *testing.T) {
	defaultSetting()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "get_all_service_addresses",
		ID:      1,
	}

	//param := map[string]interface{}{}
	//req.Params = append(req.Params, param)
	fmt.Println("request : ", req.String())

	resp, err := Forward(req)
	if resp.String() == "" || err != nil {
		fmt.Println("Error : ", err)
		t.Errorf("Failed to start main")
	}
	fmt.Println("Response : ", resp.String())
}

/*
func TestCreateIdentity(t *testing.T) {
	defaultSetting()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "create_identity",
		ID:      1,
	}
	time := big.NewInt(1557822880)

	// timebytes, _ = time.MarshalJSON()
	// fmt.Println("request timebytes: ", timebytes)
	// text, _ = time.MarshalText()
	// fmt.Println("request timebytes: ", text)
	timeHex := hexutil.EncodeBig(time)
	fmt.Println("request timeHex: ", timeHex)
	param := map[string]interface{}{

		"recovery_address":   "0x961c20596e7EC441723FBb168461f4B51371D8aA",
		"associated_address": "0x961c20596e7EC441723FBb168461f4B51371D8aA",
		"providers":          []string{"0x084f8293F1b047D3A217025B24cd7b5aCe8fC657"},
		"resolvers":          []string{"0xF4F9790205ee559A379C519E04042b20560EefaD"},
		"v":                  "0x1c",
		"r":                  "0xa61f685c83edb7e81aa2b9fed45ab2afb036255f61212d5d31aff4ce3145de7c",
		"s":                  "0x2002552a52228d44ab557bbf98614511e5fb8fc95e9a5fc2af5c6241084ae310",
		"timestamp":          time,
	}

	req.Params = append(req.Params, param)
	fmt.Println("request : ", req.String())

	resp, err := Forward(req)
	if resp.String() == "" || err != nil {
		fmt.Println("Error : ", err)
		t.Errorf("Failed to start main")
	}
	fmt.Println("Response : ", resp.String())
}
func TestAddAssociatedAddressDelegated(t *testing.T) {
	defaultSetting()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "add_associated_address_delegated",
		ID:      1,
	}

	//ein := big.NewInt(4)
	approveTime := big.NewInt(1558334743)
	addTime := big.NewInt(1558334744)
	reqParam := addAssociatedAddressDelegatedParams{
		ApprovingAddress: common.HexToAddress("0x961c20596e7EC441723FBb168461f4B51371D8aA"),
		AddressToAdd:     common.HexToAddress("0xd3Cb37aE6a81EbF1b5C3D9422636b0dB48767B72"),
		V:                []hexutil.Bytes{hexutil.MustDecode("0x1c"), hexutil.MustDecode("0x1b")},
		R:                []hexutil.Bytes{hexutil.MustDecode("0x963f81a9faf9c990d13c5ce13c67d889295f7b555fcef4f13da2b0155cb56e36"), hexutil.MustDecode("0xa4614ef85d4a607a2db75b3cc72d311c778d41f0cf48d7423b4faca8e876b654")},
		S:                []hexutil.Bytes{hexutil.MustDecode("0x4e7c210e0762db64eea4abaa433a02444886a33a12132621aa21eebb151d1191"), hexutil.MustDecode("0x735a0b735f74091f5c05cc2625773b1f53705e73de94571c4c9030e2012b646f")},
		Timestamp:        [2]*big.Int{approveTime, addTime},
	}

	req.Params = append(req.Params, reqParam)
	fmt.Println("request : ", req.String())

	resp, err := Forward(req)
	if resp.String() == "" || err != nil {
		fmt.Println("Error : ", err)
		t.Errorf("Failed to start main")
	}
	fmt.Println("Response : ", resp.String())
}

func TestRemoveAssociatedAddressDelegated(t *testing.T) {
	defaultSetting()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "remove_associated_address_delegated",
		ID:      1,
	}

	time := big.NewInt(1558404392)

	reqParam := removeAssociatedAddressDelegatedParams{
		AddressToRemove: common.HexToAddress("0xd3Cb37aE6a81EbF1b5C3D9422636b0dB48767B72"),
		V:               hexutil.MustDecode("0x1b"),
		R:               hexutil.MustDecode("0x3fad6b80c2c66d950e52d797c2cf0dc21d04d0e5ff0feaa987b154a99eab0d6b"),
		S:               hexutil.MustDecode("0x2558a44b87bec5d7065f785805c329508ce38013f21fd048dd2350480d6b8578"),
		Timestamp:       time,
	}

	req.Params = append(req.Params, reqParam)
	fmt.Println("request : ", req.String())

	resp, err := Forward(req)
	if resp.String() == "" || err != nil {
		fmt.Println("Error : ", err)
		t.Errorf("Failed to start main")
	}
	fmt.Println("Response : ", resp.String())
}

func TestAddKeyDelegated(t *testing.T) {
	defaultSetting()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "add_key_delegated",
		ID:      1,
	}

	time := big.NewInt(1558404392)

	resolverAddr, _ := servicekeyresolver.GetAddress()
	reqParam := addKeyDelegatedParams{
		ResolverAddress:   *resolverAddr,
		AssociatedAddress: common.HexToAddress("0x961c20596e7EC441723FBb168461f4B51371D8aA"),
		Key:               common.HexToAddress("0xd3Cb37aE6a81EbF1b5C3D9422636b0dB48767B72"),
		Symbol:            "test",
		V:                 hexutil.MustDecode("0x1b"),
		R:                 hexutil.MustDecode("0x120b79114f7b5a73c9e2728e2e12ce51b6fb0858174a90eae812366b8dd52d4c"),
		S:                 hexutil.MustDecode("0x4d655153c224d80b4a27c873c54611d23e6d95364387cc018e9fe00c3a5af605"),
		Timestamp:         time,
	}

	req.Params = append(req.Params, reqParam)
	fmt.Println("request : ", req.String())

	resp, err := Forward(req)
	if resp.String() == "" || err != nil {
		fmt.Println("Error : ", err)
		t.Errorf("Failed to start main")
	}
	fmt.Println("Response : ", resp.String())
}


func TestRemoveKeyDelegated(t *testing.T) {
	defaultSetting()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "remove_key_delegated",
		ID:      1,
	}

	time := big.NewInt(1558404392)

	resolverAddr, _ := servicekeyresolver.GetAddress()
	reqParam := removeKeyDelegatedParams{
		ResolverAddress:   *resolverAddr,
		AssociatedAddress: common.HexToAddress("0x961c20596e7EC441723FBb168461f4B51371D8aA"),
		Key:               common.HexToAddress("0xd3Cb37aE6a81EbF1b5C3D9422636b0dB48767B72"),
		V:                 hexutil.MustDecode("0x1b"),
		R:                 hexutil.MustDecode("0x1ea57bbf64a19f9763a69d5cb719ed47c294fcde62e8f1f99ceb9349e1ade502"),
		S:                 hexutil.MustDecode("0x6a3257a1b33b5c0eb8abb28ea736b945646e1da6aaaf58cc852b645ba210d395"),
		Timestamp:         time,
	}

	req.Params = append(req.Params, reqParam)
	fmt.Println("request : ", req.String())

	resp, err := Forward(req)
	if resp.String() == "" || err != nil {
		fmt.Println("Error : ", err)
		t.Errorf("Failed to start main")
	}
	fmt.Println("Response : ", resp.String())
}
*/

/*
func TestRemoveKeysDelegated(t *testing.T) {
	defaultSetting()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "remove_keys_delegated",
		ID:      1,
	}

	time := big.NewInt(1558404392)

	resolverAddr, _ := servicekeyresolver.GetAddress()
	reqParam := removeKeyDelegatedParams{
		ResolverAddress:   *resolverAddr,
		AssociatedAddress: common.HexToAddress("0x961c20596e7EC441723FBb168461f4B51371D8aA"),
		V:                 hexutil.MustDecode("0x1c"),
		R:                 hexutil.MustDecode("0x56a0b769505922060ccae317aa473519661cdc95b22c5d6a0557bb74c1102b57"),
		S:                 hexutil.MustDecode("0x29750cdc324005a0b834844c33fb533c148888ab8490058d93524498ebd51b9b"),
		Timestamp:         time,
	}

	req.Params = append(req.Params, reqParam)
	fmt.Println("request : ", req.String())

	resp, err := Forward(req)
	if resp.String() == "" || err != nil {
		fmt.Println("Error : ", err)
		t.Errorf("Failed to start main")
	}
	fmt.Println("Response : ", resp.String())
}
*/

func TestCreateIdentity_test(t *testing.T) {
	defaultSetting()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "create_identity_test",
		ID:      1,
	}

	time := big.NewInt(1557822880)
	reqParam := identityCreateParams{
		RecoveryAddress:   common.HexToAddress("0x961c20596e7EC441723FBb168461f4B51371D8aA"),
		AssociatedAddress: common.HexToAddress("0x961c20596e7EC441723FBb168461f4B51371D8aA"),
		Providers:         []common.Address{common.HexToAddress("0x084f8293F1b047D3A217025B24cd7b5aCe8fC657")},
		Resolvers:         []common.Address{common.HexToAddress("0xF4F9790205ee559A379C519E04042b20560EefaD")},
		V:                 hexutil.MustDecode("0x1c"),
		R:                 hexutil.MustDecode("0xa61f685c83edb7e81aa2b9fed45ab2afb036255f61212d5d31aff4ce3145de7c"),
		S:                 hexutil.MustDecode("0x2002552a52228d44ab557bbf98614511e5fb8fc95e9a5fc2af5c6241084ae310"),
		Timestamp:         time,
	}

	identityRegistryAddr, _ := identityregistry.GetAddress()
	irAddrBytes := identityRegistryAddr.Bytes()
	providersBytes := addressArrayToBytes(reqParam.Providers)
	resolversBytes := addressArrayToBytes(reqParam.Resolvers)
	timestampBytes := bigIntToByte32(reqParam.Timestamp)

	var signData []byte
	signData = append(signData, headerBytes...)
	signData = append(signData, irAddrBytes...)
	signData = append(signData, []byte("I authorize the creation of an Identity on my behalf.")...)
	signData = append(signData, reqParam.RecoveryAddress.Bytes()...)
	signData = append(signData, reqParam.AssociatedAddress.Bytes()...)
	signData = append(signData, providersBytes...)
	signData = append(signData, resolversBytes...)
	signData = append(signData, timestampBytes[:]...)
	fmt.Printf("SignData : %x \n", signData)
	keccakBytes := ethCrypto.Keccak256(signData)
	fmt.Printf("keccakBytes : %x \n", keccakBytes)

	prikey, err := ethCrypto.HexToECDSA("01b149603ca8f537bbb4e45d22e77df9054e50d826bb5f0a34e9ce460432b596")
	if err != nil {
		fmt.Printf("Failed to get key %s", err)
	}
	var tsig hexutil.Bytes
	tsig, err = signBytes(signData, prikey)
	//tsig, err := ethCrypto.Sign(signHash(signData), prikey)
	if err != nil {
		fmt.Printf("Failed to sign %s", err)
	}
	fmt.Printf("signature : %v %v\n", tsig, len(tsig))
	rsig := tsig[:32]
	ssig := tsig[32:64]
	v := tsig[64]

	fmt.Printf("V,R,S : %v,%v,%v \n", v, rsig, ssig)
	reqParam.V[0] = v + 27
	reqParam.R = rsig
	reqParam.S = ssig

	var pubkeyBytes, shash hexutil.Bytes
	shash = signHash(keccakBytes)
	fmt.Printf("shash : %v \n", shash)
	pubkeyBytes, _ = ethCrypto.Ecrecover(shash, tsig)
	fmt.Printf("pubkeyBytes : %v \n", pubkeyBytes)
	pubKey, _ := ethCrypto.UnmarshalPubkey(pubkeyBytes)
	//var addrBytes common.Address
	addrbytes := ethCrypto.PubkeyToAddress(*pubKey)
	fmt.Printf("address : %x \n", addrbytes)

	req.Params = append(req.Params, reqParam)
	fmt.Println("request : ", req.String())

	resp, err := Forward(req)
	if resp.String() == "" || err != nil {
		fmt.Println("Error : ", err)
		t.Errorf("Failed to start main")
	}
	fmt.Println("Response : ", resp.String())

}
func TestAddAssociatedAddressDelegatedTest(t *testing.T) {
	defaultSetting()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "add_associated_address_delegated_test",
		ID:      1,
	}

	ein := big.NewInt(4)
	approveTime := big.NewInt(1558334743)
	addTime := big.NewInt(1558334744)
	reqParam := addAssociatedAddressDelegatedParams{
		ApprovingAddress: common.HexToAddress("0x961c20596e7EC441723FBb168461f4B51371D8aA"),
		AddressToAdd:     common.HexToAddress("0xd3Cb37aE6a81EbF1b5C3D9422636b0dB48767B72"),
		// V:                [2]hexutil.Bytes{hexutil.MustDecode("0x1c"), hexutil.MustDecode("0x1c")},
		// R:                [2]hexutil.Bytes{hexutil.MustDecode("0xa61f685c83edb7e81aa2b9fed45ab2afb036255f61212d5d31aff4ce3145de7c"), hexutil.MustDecode("0xa61f685c83edb7e81aa2b9fed45ab2afb036255f61212d5d31aff4ce3145de7c")},
		// S:                [2]hexutil.Bytes{hexutil.MustDecode("0x2002552a52228d44ab557bbf98614511e5fb8fc95e9a5fc2af5c6241084ae310"), hexutil.MustDecode("0x2002552a52228d44ab557bbf98614511e5fb8fc95e9a5fc2af5c6241084ae310")},
		Timestamp: [2]*big.Int{approveTime, addTime},
	}

	identityRegistryAddr, _ := identityregistry.GetAddress()
	irAddrBytes := identityRegistryAddr.Bytes()

	approveTimestampBytes := bigIntToByte32(reqParam.Timestamp[0])
	addTimestampBytes := bigIntToByte32(reqParam.Timestamp[1])
	einBytes := bigIntToByte32(ein)

	approvePrikey, err := ethCrypto.HexToECDSA("01b149603ca8f537bbb4e45d22e77df9054e50d826bb5f0a34e9ce460432b596")
	if err != nil {
		fmt.Printf("Failed to get key %s", err)
	}
	addPrikey, err := ethCrypto.HexToECDSA("7c57396b8f48edb490c0efbcdc81b5198c54216546233a4272decd7ebbbda4bb")
	if err != nil {
		fmt.Printf("Failed to get key %s", err)
	}

	//sign  for approve
	var approveData []byte
	approveData = append(approveData, headerBytes...)
	approveData = append(approveData, irAddrBytes...)
	approveData = append(approveData, []byte("I authorize adding this address to my Identity.")...)
	approveData = append(approveData, einBytes[:]...)
	approveData = append(approveData, reqParam.AddressToAdd.Bytes()...)
	approveData = append(approveData, approveTimestampBytes[:]...)
	fmt.Printf("SignData for Approve : %x \n", approveData)
	// approveHash := ethCrypto.Keccak256(approveData)
	// fmt.Printf("keccakBytes : %x \n", approveHash)

	var approveSig hexutil.Bytes
	approveSig, err = signBytes(approveData, approvePrikey)
	//tsig, err := ethCrypto.Sign(signHash(signData), prikey)
	if err != nil {
		fmt.Printf("Failed to sign %s", err)
	}
	fmt.Printf("signature : %v %v\n", approveSig, len(approveSig))
	approveR := approveSig[:32]
	approveS := approveSig[32:64]
	approveV := approveSig[64]

	fmt.Printf("V,R,S : %v,%v,%v \n", approveV, approveR, approveS)
	var vB1 hexutil.Bytes
	vB1 = append(vB1, approveV+27)
	//vB1 = append(vB1, approveV+27)

	// reqParam.V[0] = append(reqParam.V[0], approveV+27)
	// reqParam.V[0] = append(reqParam.V[0], approveV+27)
	//reqParam.V[0] = approveV + 27
	// reqParam.R[0] = approveR
	// reqParam.S[0] = approveS

	//sign  for add
	var addData []byte
	addData = append(addData, headerBytes...)
	addData = append(addData, irAddrBytes...)
	addData = append(addData, []byte("I authorize being added to this Identity.")...)
	addData = append(addData, einBytes[:]...)
	addData = append(addData, reqParam.AddressToAdd.Bytes()...)
	addData = append(addData, addTimestampBytes[:]...)
	fmt.Printf("SignData for Add : %x \n", addData)

	var addSig hexutil.Bytes
	addSig, err = signBytes(addData, addPrikey)

	if err != nil {
		fmt.Printf("Failed to sign %s", err)
	}
	fmt.Printf("signature : %v %v\n", addSig, len(addSig))
	addR := addSig[:32]
	addS := addSig[32:64]
	addV := addSig[64]

	fmt.Printf("V,R,S : %v,%v,%v \n", addV, addR, addS)
	var vB2 hexutil.Bytes
	vB2 = append(vB2, addV+27)
	//vB2 = append(vB2, approveV+27)
	// reqParam.V[1] = append(reqParam.V[1], addV+27)
	// reqParam.V[1] = append(reqParam.V[1], addV+27)
	//reqParam.V[1] = addV + 27
	// reqParam.R[1] = addR
	// reqParam.S[1] = addS

	reqParam.V = []hexutil.Bytes{vB1, vB2}
	reqParam.R = []hexutil.Bytes{approveR, addR}
	reqParam.S = []hexutil.Bytes{approveS, addS}
	req.Params = append(req.Params, reqParam)
	fmt.Println("request : ", req.String())

	//jsonBytes, err1 := hexutil..Marshal(obj)

	resp, err := Forward(req)
	if resp.String() == "" || err != nil {
		fmt.Println("Error : ", err)
		t.Errorf("Failed to start main")
	}
	fmt.Println("Response : ", resp.String())
}

func TestRemoveAssociatedAddressDelegatedTest(t *testing.T) {
	defaultSetting()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "remove_associated_address_delegated_test",
		ID:      1,
	}

	ein := big.NewInt(4)
	time := big.NewInt(1558404392)

	reqParam := removeAssociatedAddressDelegatedParams{
		AddressToRemove: common.HexToAddress("0xd3Cb37aE6a81EbF1b5C3D9422636b0dB48767B72"),
		Timestamp:       time,
	}

	identityRegistryAddr, _ := identityregistry.GetAddress()
	irAddrBytes := identityRegistryAddr.Bytes()

	timestampBytes := bigIntToByte32(reqParam.Timestamp)

	einBytes := bigIntToByte32(ein)

	prikey, err := ethCrypto.HexToECDSA("7c57396b8f48edb490c0efbcdc81b5198c54216546233a4272decd7ebbbda4bb")
	if err != nil {
		fmt.Printf("Failed to get key %s", err)
	}

	//sign  for approve
	var removeData []byte
	removeData = append(removeData, headerBytes...)
	removeData = append(removeData, irAddrBytes...)
	removeData = append(removeData, []byte("I authorize removing this address from my Identity.")...)
	removeData = append(removeData, einBytes[:]...)
	removeData = append(removeData, reqParam.AddressToRemove.Bytes()...)
	removeData = append(removeData, timestampBytes[:]...)
	fmt.Printf("SignData for Approve : %x \n", removeData)
	// approveHash := ethCrypto.Keccak256(approveData)
	// fmt.Printf("keccakBytes : %x \n", approveHash)

	var removeSig hexutil.Bytes
	removeSig, err = signBytes(removeData, prikey)
	//tsig, err := ethCrypto.Sign(signHash(signData), prikey)
	if err != nil {
		fmt.Printf("Failed to sign %s", err)
	}
	fmt.Printf("signature : %v %v\n", removeSig, len(removeSig))
	removeR := removeSig[:32]
	removeS := removeSig[32:64]
	removeV := removeSig[64]

	fmt.Printf("V,R,S : %v,%v,%v \n", removeV, removeR, removeS)
	var vB1 hexutil.Bytes
	vB1 = append(vB1, removeV+27)

	reqParam.V = vB1
	reqParam.R = removeR
	reqParam.S = removeS
	req.Params = append(req.Params, reqParam)
	fmt.Println("request : ", req.String())

	//jsonBytes, err1 := hexutil..Marshal(obj)

	resp, err := Forward(req)
	if resp.String() == "" || err != nil {
		fmt.Println("Error : ", err)
		t.Errorf("Failed to start main")
	}
	fmt.Println("Response : ", resp.String())
}

func TestAddKeyDelegatedTest(t *testing.T) {
	defaultSetting()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "add_key_delegated_test",
		ID:      1,
	}

	time := big.NewInt(1558404392)

	resolverAddr, _ := servicekeyresolver.GetAddress()
	reqParam := addKeyDelegatedParams{
		ResolverAddress:   *resolverAddr,
		AssociatedAddress: common.HexToAddress("0x961c20596e7EC441723FBb168461f4B51371D8aA"),
		Key:               common.HexToAddress("0xd3Cb37aE6a81EbF1b5C3D9422636b0dB48767B72"),
		Symbol:            "test",
		Timestamp:         time,
	}

	resolverAddrBytes := resolverAddr.Bytes()

	timestampBytes := bigIntToByte32(reqParam.Timestamp)

	//einBytes := bigIntToByte32(ein)

	prikey, err := ethCrypto.HexToECDSA("01b149603ca8f537bbb4e45d22e77df9054e50d826bb5f0a34e9ce460432b596")
	if err != nil {
		fmt.Printf("Failed to get key %s", err)
	}

	//sign  for approve
	var data []byte
	data = append(data, headerBytes...)
	data = append(data, resolverAddrBytes...)
	data = append(data, []byte("I authorize the addition of a service key on my behalf.")...)
	data = append(data, reqParam.Key.Bytes()...)
	data = append(data, []byte(reqParam.Symbol)...)
	data = append(data, timestampBytes[:]...)
	fmt.Printf("SignData : '%x' \n", data)
	hash := ethCrypto.Keccak256(data)
	fmt.Printf("hash : '%x' \n", hash)
	var sig hexutil.Bytes
	//sig, err = signBytes(data, prikey)
	sig, err = ethCrypto.Sign(hash, prikey)
	if err != nil {
		fmt.Printf("Failed to sign %s", err)
	}
	fmt.Printf("signature : %v %v\n", sig, len(sig))
	r := sig[:32]
	s := sig[32:64]
	v := sig[64]

	fmt.Printf("V,R,S : %v,%v,%v \n", v, r, s)
	var vB hexutil.Bytes
	vB = append(vB, v+27)

	reqParam.V = vB
	reqParam.R = r
	reqParam.S = s
	req.Params = append(req.Params, reqParam)
	fmt.Println("request : ", req.String())

	//jsonBytes, err1 := hexutil..Marshal(obj)

	resp, err := Forward(req)
	if resp.String() == "" || err != nil {
		fmt.Println("Error : ", err)
		t.Errorf("Failed to start main")
	}
	fmt.Println("Response : ", resp.String())
}

func TestRemoveKeyDelegatedTest(t *testing.T) {
	defaultSetting()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "remove_key_delegated_test",
		ID:      1,
	}

	time := big.NewInt(1558404392)

	resolverAddr, _ := servicekeyresolver.GetAddress()
	reqParam := removeKeyDelegatedParams{
		ResolverAddress:   *resolverAddr,
		AssociatedAddress: common.HexToAddress("0x961c20596e7EC441723FBb168461f4B51371D8aA"),
		Key:               common.HexToAddress("0xd3Cb37aE6a81EbF1b5C3D9422636b0dB48767B72"),
		Timestamp:         time,
	}

	resolverAddrBytes := resolverAddr.Bytes()

	timestampBytes := bigIntToByte32(reqParam.Timestamp)

	//einBytes := bigIntToByte32(ein)

	prikey, err := ethCrypto.HexToECDSA("01b149603ca8f537bbb4e45d22e77df9054e50d826bb5f0a34e9ce460432b596")
	if err != nil {
		fmt.Printf("Failed to get key %s", err)
	}

	//sign  for approve
	var data []byte
	data = append(data, headerBytes...)
	data = append(data, resolverAddrBytes...)
	data = append(data, []byte("I authorize the removal of a service key on my behalf.")...)
	data = append(data, reqParam.Key.Bytes()...)
	data = append(data, timestampBytes[:]...)
	fmt.Printf("SignData : '%x' \n", data)
	hash := ethCrypto.Keccak256(data)
	fmt.Printf("hash : '%x' \n", hash)
	var sig hexutil.Bytes
	//sig, err = signBytes(data, prikey)
	sig, err = ethCrypto.Sign(hash, prikey)
	if err != nil {
		fmt.Printf("Failed to sign %s", err)
	}
	fmt.Printf("signature : %v %v\n", sig, len(sig))
	r := sig[:32]
	s := sig[32:64]
	v := sig[64]

	fmt.Printf("V,R,S : %v,%v,%v \n", v, r, s)
	var vB hexutil.Bytes
	vB = append(vB, v+27)

	reqParam.V = vB
	reqParam.R = r
	reqParam.S = s
	req.Params = append(req.Params, reqParam)
	fmt.Println("request : ", req.String())

	//jsonBytes, err1 := hexutil..Marshal(obj)

	resp, err := Forward(req)
	if resp.String() == "" || err != nil {
		fmt.Println("Error : ", err)
		t.Errorf("Failed to start main")
	}
	fmt.Println("Response : ", resp.String())
}

func TestRemoveKeysDelegatedTest(t *testing.T) {
	defaultSetting()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "remove_keys_delegated_test",
		ID:      1,
	}

	time := big.NewInt(1558404392)

	resolverAddr, _ := servicekeyresolver.GetAddress()
	reqParam := removeKeysDelegatedParams{
		ResolverAddress:   *resolverAddr,
		AssociatedAddress: common.HexToAddress("0x961c20596e7EC441723FBb168461f4B51371D8aA"),
		Timestamp:         time,
	}

	resolverAddrBytes := resolverAddr.Bytes()

	timestampBytes := bigIntToByte32(reqParam.Timestamp)

	//einBytes := bigIntToByte32(ein)

	prikey, err := ethCrypto.HexToECDSA("01b149603ca8f537bbb4e45d22e77df9054e50d826bb5f0a34e9ce460432b596")
	if err != nil {
		fmt.Printf("Failed to get key %s", err)
	}

	//sign  for approve
	var data []byte
	data = append(data, headerBytes...)
	data = append(data, resolverAddrBytes...)
	data = append(data, []byte("I authorize the removal of all service keys on my behalf.")...)

	data = append(data, timestampBytes[:]...)
	fmt.Printf("SignData : '%x' \n", data)
	hash := ethCrypto.Keccak256(data)
	fmt.Printf("hash : '%x' \n", hash)
	var sig hexutil.Bytes
	//sig, err = signBytes(data, prikey)
	sig, err = ethCrypto.Sign(hash, prikey)
	if err != nil {
		fmt.Printf("Failed to sign %s", err)
	}
	fmt.Printf("signature : %v %v\n", sig, len(sig))
	r := sig[:32]
	s := sig[32:64]
	v := sig[64]

	fmt.Printf("V,R,S : %v,%v,%v \n", v, r, s)
	var vB hexutil.Bytes
	vB = append(vB, v+27)

	reqParam.V = vB
	reqParam.R = r
	reqParam.S = s
	req.Params = append(req.Params, reqParam)
	fmt.Println("request : ", req.String())

	//jsonBytes, err1 := hexutil..Marshal(obj)

	resp, err := Forward(req)
	if resp.String() == "" || err != nil {
		fmt.Println("Error : ", err)
		t.Errorf("Failed to start main")
	}
	fmt.Println("Response : ", resp.String())
}

/*
func TestCreateIdentity_parameterCheck(t *testing.T) {
	defaultSetting()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "create_identity_test",
		ID:      1,
	}

	time := big.NewInt(1558443167)
	reqParam := identityCreateParams{
		RecoveryAddress:   common.HexToAddress("0xdE627A778479645De3203Ff50aD22057116Af526"),
		AssociatedAddress: common.HexToAddress("0xdE627A778479645De3203Ff50aD22057116Af526"),
		Providers:         []common.Address{common.HexToAddress("0x084f8293F1b047D3A217025B24cd7b5aCe8fC657")},
		Resolvers:         []common.Address{common.HexToAddress("0xF4F9790205ee559A379C519E04042b20560EefaD")},
		V:                 hexutil.MustDecode("0x1c"),
		R:                 hexutil.MustDecode("0x4f76ea16cc92e0549478b0023cade0d78f87c1872de8b3a415c7f36ceacfcde8"),
		S:                 hexutil.MustDecode("0x26d3c47666df65ed3df096f357a3669946772c61a61dc28d24336751078ed306"),
		Timestamp:         time,
	}
	//{"associated_address":"0xdE627A778479645De3203Ff50aD22057116Af526","providers":["0x084f8293f1b047d3a217025b24cd7b5ace8fc657"],"r":"0x4f76ea16cc92e0549478b0023cade0d78f87c1872de8b3a415c7f36ceacfcde8","recovery_address":"0xdE627A778479645De3203Ff50aD22057116Af526","resolvers":["0xf4f9790205ee559a379c519e04042b20560eefad"],"s":"0x26d3c47666df65ed3df096f357a3669946772c61a61dc28d24336751078ed306","timestamp":1558443167,"v":"0x1C"}
	//{"associated_address":"0xD318C906B02E0A268710F6fF23F822093Fc7FCbe","providers":["0x084f8293f1b047d3a217025b24cd7b5ace8fc657"],"r":"0xe19d091e1409e8893adebd9d77ddd5e37193b1946e1bd7c8e77284e635aa6339","recovery_address":"0xD318C906B02E0A268710F6fF23F822093Fc7FCbe","resolvers":["0xf4f9790205ee559a379c519e04042b20560eefad"],"s":"0x662d03f43112ca42eaa780083433b091fbd6d2d88fff76105254dc99b2006239","timestamp":1558442428,"v":"0x1B"}

	identityRegistryAddr, _ := identityregistry.GetAddress()
	irAddrBytes := identityRegistryAddr.Bytes()
	providersBytes := addressArrayToBytes(reqParam.Providers)
	resolversBytes := addressArrayToBytes(reqParam.Resolvers)
	timestampBytes := bigIntToByte32(reqParam.Timestamp)

	var signData []byte
	signData = append(signData, headerBytes...)
	signData = append(signData, irAddrBytes...)
	signData = append(signData, []byte("I authorize the creation of an Identity on my behalf.")...)
	signData = append(signData, reqParam.RecoveryAddress.Bytes()...)
	signData = append(signData, reqParam.AssociatedAddress.Bytes()...)
	signData = append(signData, providersBytes...)
	signData = append(signData, resolversBytes...)
	signData = append(signData, timestampBytes[:]...)
	fmt.Printf("SignData : %x \n", signData)
	keccakBytes := ethCrypto.Keccak256(signData)
	fmt.Printf("keccakBytes : %x \n", keccakBytes)

	req.Params = append(req.Params, reqParam)
	fmt.Println("request : ", req.String())

	resp, err := Forward(req)
	if resp.String() == "" || err != nil {
		fmt.Println("Error : ", err)
		t.Errorf("Failed to start main")
	}
	fmt.Println("Response : ", resp.String())

}
*/
