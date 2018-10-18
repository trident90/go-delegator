package metaservice

import (
	"bufio"
	"crypto/ecdsa"
	"fmt"
	"os"
	"strings"
	"testing"

	"bitbucket.org/coinplugin/proxy/metaservice/sc/identity"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	log "github.com/sirupsen/logrus"

	"bitbucket.org/coinplugin/proxy/crypto"
	"bitbucket.org/coinplugin/proxy/json"
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

func TestGetRegistryAddress(t *testing.T) {
	defaultSetting()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "get_registry_address",
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
func TestGetIdentityManagerAddress(t *testing.T) {
	defaultSetting()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "get_identity_manager_address",
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
func TestGetTopicRegistryAddress(t *testing.T) {
	defaultSetting()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "get_topic_registry_address",
		ID:      1,
	}

	// param := map[string]interface{}{}
	// req.Params = append(req.Params, param)
	fmt.Println("request : ", req.String())

	resp, err := Forward(req)
	if resp.String() == "" || err != nil {
		fmt.Println("Error : ", err)
		t.Errorf("Failed to start main")
	}
	fmt.Println("Response : ", resp.String())
}
func TestGetAttestationAgencyRegistryAddress(t *testing.T) {
	defaultSetting()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "get_attestation_agency_registry_address",
		ID:      1,
	}

	// param := map[string]interface{}{}
	// req.Params = append(req.Params, param)
	fmt.Println("request : ", req.String())

	resp, err := Forward(req)
	if resp.String() == "" || err != nil {
		fmt.Println("Error : ", err)
		t.Errorf("Failed to start main")
	}
	fmt.Println("Response : ", resp.String())
}

func TestGetAchievementManagerAddress(t *testing.T) {
	defaultSetting()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "get_achievement_manager_address",
		ID:      1,
	}

	param := map[string]interface{}{}
	req.Params = append(req.Params, param)
	fmt.Println("request : ", req.String())

	resp, err := Forward(req)
	if resp.String() == "" || err != nil {
		fmt.Println("Error : ", err)
		t.Errorf("Failed to start main")
	}
	fmt.Println("Response : ", resp.String())
}

func TestGetAchievementAddress(t *testing.T) {
	defaultSetting()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "get_achievement_address",
		ID:      1,
	}

	// param := map[string]interface{}{}
	// req.Params = append(req.Params, param)
	fmt.Println("request : ", req.String())

	resp, err := Forward(req)
	if resp.String() == "" || err != nil {
		fmt.Println("Error : ", err)
		t.Errorf("Failed to start main")
	}
	fmt.Println("Response : ", resp.String())
}
func TestGetAllSystemAddress(t *testing.T) {
	defaultSetting()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "get_all_system_address",
		ID:      1,
	}

	// param := map[string]interface{}{}
	// req.Params = append(req.Params, param)
	fmt.Println("request : ", req.String())

	resp, err := Forward(req)
	if resp.String() == "" || err != nil {
		fmt.Println("Error : ", err)
		t.Errorf("Failed to start main")
	}
	fmt.Println("Response : ", resp.String())
}

func TestCreateMetaId(t *testing.T) {
	defaultSetting()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "create_meta_id",
		ID:      1,
	}

	param := map[string]interface{}{
		"address":   "0x961c20596e7EC441723FBb168461f4B51371D8aA",
		"signature": "0x2bdbb3388163719b81d0554ecbe81b6553c91d586f0238cd40dd65d3f2c4ea1e4aa2e9005ac46cd3548046f55c8ae89834010f50394c94476616cfc81b5cf2ab1c",
		"recaptcha": "",
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

func TestAddActionKey(t *testing.T) {
	defaultSetting()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "delegated_execute",
		ID:      1,
	}
	fromAddr := common.HexToAddress("0x961c20596e7EC441723FBb168461f4B51371D8aA")
	scAddr := common.HexToAddress("0xe052cb04e4fe4d3ca69d247b4eff2aff35613b0e")
	toAddr := common.HexToAddress("0xe052cb04e4fe4d3ca69d247b4eff2aff35613b0e")
	ins, err := identity.GetInstance(scAddr)
	if err != nil {
		t.Fatal(err)
	}
	nonce, err := identity.CallGetTransactionCount(ins)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("nonce :", nonce)
	identityAbi, err := abi.JSON(strings.NewReader(identity.IdentityABI))
	if err != nil {
		t.Fatal(err)
	}

	actionAddr := common.HexToAddress("0xd3Cb37aE6a81EbF1b5C3D9422636b0dB48767B72")
	key := identity.CallAddrToKey(actionAddr)
	var data hexutil.Bytes
	data, err = identityAbi.Pack("addKey", key, common.Big1, common.Big2)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("data : %x \n", data)

	signPrivkey, err := ethCrypto.HexToECDSA("01b149603ca8f537bbb4e45d22e77df9054e50d826bb5f0a34e9ce460432b596")
	if err != nil {
		t.Fatalf("Failed to get key %s", err)
	}

	//Make executeSig
	var executeSigData hexutil.Bytes
	valueBytes := intToByte32(common.Big0)
	nonceBytes := intToByte32(nonce)
	executeSigData = concatBytes(toAddr.Bytes(), valueBytes[:], data, nonceBytes[:])
	fmt.Printf("executeSigData : %v \n", executeSigData)
	var sig hexutil.Bytes
	sig, err = signBytes(executeSigData, signPrivkey)
	if err != nil {
		t.Fatalf("Failed to sign %s", err)
	}
	fmt.Printf("signature : %x \n", sig)

	param := map[string]interface{}{
		"meta_id":   toAddr,
		"from":      fromAddr,
		"to":        toAddr,
		"value":     common.Big0,
		"data":      data,
		"nonce":     nonce,
		"signature": sig.String(),
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

func TestAddSelfClaim(t *testing.T) {
	defaultSetting()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "delegated_execute",
		ID:      1,
	}
	fromAddr := common.HexToAddress("0x961c20596e7EC441723FBb168461f4B51371D8aA")
	scAddr := common.HexToAddress("0xe052cb04e4fe4d3ca69d247b4eff2aff35613b0e")
	toAddr := common.HexToAddress("0xe052cb04e4fe4d3ca69d247b4eff2aff35613b0e")
	signPrivkey, _ := ethCrypto.HexToECDSA("01b149603ca8f537bbb4e45d22e77df9054e50d826bb5f0a34e9ce460432b596")

	ins, err := identity.GetInstance(scAddr)
	if err != nil {
		t.Fatal(err)
	}

	nonce, err := identity.CallGetTransactionCount(ins)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("nonce :", nonce)

	//1. claim Data
	_topic := common.Big1
	_scheme := common.Big1
	_issuer := common.HexToAddress("0x961c20596e7EC441723FBb168461f4B51371D8aA")
	_data := hexutil.MustDecode("0x1b442640e0333cb03054940e3cda07da982d2b57af68c3df8d0557b47a77d0bc")
	_uri := "MetaPrint"

	//2. claim.signature 생성
	//2-1. Sign Data 생성
	// signdata : scAddr + _topic + data
	var _signData hexutil.Bytes

	//_topicBytes, err := claimSignArgs.Pack(_topic)
	_topicBytes := intToByte32(_topic)
	_signData = concatBytes(_signData, _topicBytes[:], _data)
	fmt.Println("_signData : ", _signData.String())

	//2-2 sign  data
	var _signature hexutil.Bytes
	_signature, err = signBytes(_signData, signPrivkey)
	fmt.Println("signature : ", _signature.String())

	//3.  data(DelegateExcute parameter) 생성
	identityAbi, err := abi.JSON(strings.NewReader(identity.IdentityABI))
	if err != nil {
		t.Fatal(err)
	}
	// "addClaim":  _topic, _scheme, issuer, _signature, _data, _uri
	var executeData hexutil.Bytes
	executeData, err = identityAbi.Pack("addClaim", _topic, _scheme, _issuer, _signature, _data, _uri)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("executeData : %v \n", executeData)

	//4. Make signature (DelegateExcute parameter) 생성
	var executeSigData hexutil.Bytes
	valueBytes := intToByte32(common.Big0)
	nonceBytes := intToByte32(nonce)
	executeSigData = concatBytes(toAddr.Bytes(), valueBytes[:], executeData, nonceBytes[:])

	fmt.Printf("executeSigData :  %v \n", executeSigData)
	var executeSig hexutil.Bytes
	executeSig, err = signBytes(executeSigData, signPrivkey)
	if err != nil {
		t.Fatalf("Failed to sign %s", err)
	}
	fmt.Printf("exec signature : %v \n", executeSig)

	//5. make Request Data
	param := map[string]interface{}{
		"meta_id":   toAddr,
		"from":      fromAddr,
		"to":        toAddr,
		"value":     common.Big0,
		"data":      executeData,
		"nonce":     nonce,
		"signature": executeSig,
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
func signBytes(bmsg []byte, privKey *ecdsa.PrivateKey) ([]byte, error) {
	bMsg := ethCrypto.Keccak256(bmsg)
	return ethCrypto.Sign(signHash(bMsg), privKey)
}
func signHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return ethCrypto.Keccak256([]byte(msg))
}
