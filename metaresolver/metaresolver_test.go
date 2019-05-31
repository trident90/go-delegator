package metaresolver

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/metadium/go-delegator/crypto"
	"github.com/metadium/go-delegator/json"
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

	resolverAddr := servicekeyresolver.GetAddress()
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

	resolverAddr := servicekeyresolver.GetAddress()
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
func TestRemoveKeysDelegated(t *testing.T) {
	defaultSetting()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "remove_keys_delegated",
		ID:      1,
	}

	time := big.NewInt(1558404392)

	resolverAddr := servicekeyresolver.GetAddress()
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
