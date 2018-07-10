package predefined

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"

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
func TestMetaIdRegister(t *testing.T) {
	defaultSetting()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "register_meta_id",
		ID:      1,
	}

	param := map[string]interface{}{
		"address": "0x961c20596e7EC441723FBb168461f4B51371D8aA",
		"meta_id": "0x400c2c9ae5f66133fb6a48d476756b9048a345a0666676e416cdaddf93ae09c5",
		"user_hash_list": []string{
			"0xa21be2c1a1951cafd81a22bffbd1d48294d45525b129d2002c9424c1d5dfb8ea",
			"0x4603ce8e282fe90eea9074d43991ba46afa0b3b23f3bc3148a055e6c51fc204a",
			"0xf26afd7d9e642849530645e15cd62af0e0cbbb5d2f41d1881c48b0da76a4d9a6",
			"0x9e50e9545037b5ff006c9e421f63a6401554dfe20d2fb16d0ce04fa02c69b8f6",
			"0x288a85888972f2c365b88075940528d5a36451d9cf3dce6251e980e17e19846a",
			"0xed10665502644999df542f6ce39506cd1c84d8aab9886df731032beb38313cbf",
		},
		"signature": "0xf0589f079f48101a17fcce0360ed2c614652d08c757bb7b31f7146508168acc33e59b723568b2977c4bc6849de99e5429ccf6cd27d9ab1d83c055a32267386bb1c",
	}
	req.Params = append(req.Params, param)
	fmt.Println("reqest : ", req.String())
	// resp, err := main.Handler(nil, events.APIGatewayProxyRequest{
	// 	Body: req.String(),
	// })
	resp, err := registerMetaID(req)
	if resp.String() == "" || err != nil {
		fmt.Println("Error : ", err)
		t.Errorf("Failed to start main")
	}
	fmt.Println("Response : ", resp.String())
}

func TestMetaIdUpdate(t *testing.T) {
	defaultSetting()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "update_meta_id",
		ID:      1,
	}

	/*
	   Address   common.Address  `json:"address"`
	   	OldMetaID hexutil.Bytes   `json:"meta_id_old"`
	   	NewMetaID hexutil.Bytes   `json:"meta_id_new"`
	   	UserHashs []hexutil.Bytes `json:"user_hash_list"`
	   	Signature hexutil.Bytes   `json:"signature"` // Sign(NewMetaId)

	*/

	param := map[string]interface{}{
		// "address":     "0x961c20596e7EC441723FBb168461f4B51371D8aA",
		// "meta_id_old": "0xdb6dd8f5917a3c2f84a280f365ac137549e62d647b6cfba05a0f2c5e8e60e972",
		// "meta_id_new": "0xc30255d8455a90fbb0ce11405b501624be9287370a41779312944c8739a9f79d",
		// "user_hash_list": []string{
		// 	"0xcbb6887db9a92ac59fa4c101d88f80b69fead4eea556ea486f276a05c3a99144",
		// 	"0x9e50e9545037b5ff006c9e421f63a6401554dfe20d2fb16d0ce04fa02c69b8f6",
		// 	"0x288a85888972f2c365b88075940528d5a36451d9cf3dce6251e980e17e19846a",
		// 	"0xe3114197cb884bbaf2fd098ef55437c832e69bb0b5aee4b535eb5db7b0cdedc5"},
		// "signature": "0x0f86138d24ecd75efd982e8c68555c97c70d5abfa67ab99edeca462b9b546d706888437eecc91f1ea5a0e815606188b6e91353bd867c9ad08cf2f62781aefe391c",

		"address":     "0x513e27db4538aa4f6ea01a2e255da95c24af95fb",
		"signature":   "0xfd295061731453421c83927127e1b0897c67af65a9c34d08b04889b3e79fbdf03af4135b17b1903ca951f26cf9c7d8d8c882ab463b38fd4efaf38eaef46738c61b",
		"meta_id_old": "0x4ee2a5ede4fad2a84f8fbd137bcbe7f7fa067a067a31a24761eef0bac2fa457c",
		"user_hash_list": []string{
			"0x801bf53cced7d71f6122ff82b59a26baab1cf22885642bbe9501977a5c166a80",
			"0xf267912b08591469b312faa9fcc645236dc6499ac72e1e270eed7ef3ca3bf5b5",
			"0x9ad58839944a807cdd6028d35ef295204ca8e8c36705e125d66fdaf2dc09b2ea",
			"0x841005121d892789ad599d0a94884006ae13dc1b1da9ec891de6239e5c353a8e",
			"0x8573f67506b9cd2207b6c0eb7721e31ec6f40475b7a11d94d3ab411301564121",
			"0xcf889c892a73b8fdc904acc9c3c120775ba61f85a3e5a798894b72fec354f2dd",
			"0x5803334362f5f567d40a2e67e4ef89358fe13187ac0319cc6883254f812e880c",
			"0x5803334362f5f567d40a2e67e4ef89358fe13187ac0319cc6883254f812e880c"},
		"meta_id_new": "0x2f3f47ef470af29b250ac795962f41f1fb8576944cc29e2f1ff8fccb1009ad92",
	}
	req.Params = append(req.Params, param)
	fmt.Println("reqest : ", req.String())

	resp, err := updateMetaID(req)
	if resp.String() == "" || err != nil {
		fmt.Println("Error : ", err)
		t.Errorf("Failed to start main")
	}
	fmt.Println("Response : ", resp.String())
}

func TestMetaIdRevoke(t *testing.T) {
	defaultSetting()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "revoke_meta_id",
		ID:      1,
	}

	param := map[string]interface{}{

		// 	"meta_id":   "0xdb6dd8f5917a3c2f84a280f365ac137549e62d647b6cfba05a0f2c5e8e60e972",
		// 	"timestamp": "0x3230313830373033313432303132",
		//	"signature": "0x3f593b366bf85cc98c877cc0c7e406017ced62b2913494457597bf78a65e8fd26023727bd2e9f7ac5dadf394744476a69fd43f5cbcc30c5b5987d6404ba4a4321b",

		"meta_id":   "0xc30255d8455a90fbb0ce11405b501624be9287370a41779312944c8739a9f79d",
		"timestamp": "0x3230313830373033313432303132",
		"signature": "0x2e0828fc90ec82ac850555dfda43fa900c55c76026fe0f769e055c17d11d324e5c473cb640e6e900a4f91451af83442dabde021fc7e46acf3b4944deaff577731c",
	}
	req.Params = append(req.Params, param)
	fmt.Println("reqest : ", req.String())

	resp, err := revokeMetaID(req)
	if resp.String() == "" || err != nil {
		fmt.Println("Error : ", err)
		t.Errorf("Failed to start main")
	}
	fmt.Println("Response : ", resp.String())
}
func TestBackupUserData(t *testing.T) {
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "backup_user_data",
		ID:      1,
	}

	param := map[string]interface{}{
		"address":   "0x961c20596e7EC441723FBb168461f4B51371D8aA",
		"meta_id":   "0xdb6dd8f5917a3c2f84a280f365ac137549e62d647b6cfba05a0f2c5e8e60e972",
		"enc_data":  "0xcbb6887db9a92ac59fa4c101d88f80b69fead4eea556ea486f276a05c3a991449e50e9545037b5ff006c9e421f63a6401554dfe20d2fb16d0ce04fa02c69b8f6288a85888972f2c365b88075940528d5a36451d9cf3dce6251e980e17e19846aed10665502644999df542f6ce39506cd1c84d8aab9886df731032beb38313cbf",
		"signature": "0x542e3d9af6758e80e4f2500395898a9cd9f5e3bd91b3053ad60d4cb0147b608c7de2ff608ae7528bc59cc6a7ba006020678d2ee5b5f88a97204734953cd2cccb1b",
	}
	req.Params = append(req.Params, param)
	fmt.Println("reqest : ", req.String())

	resp, err := backupUserData(req)
	if resp.String() == "" || err != nil {
		fmt.Println("Error : ", err)
		t.Errorf("Failed to start main")
	}
	fmt.Println("Response : ", resp.String()) //QmSVUyM2piAKC5Yc8YZUKh843aatQJWTDViozKscSk8WLd

}

func TestGetUserData(t *testing.T) {
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "get_user_data",
		ID:      1,
	}

	param := map[string]interface{}{
		"address_new": "0xd3cb37ae6a81ebf1b5c3d9422636b0db48767b72", // new Address
		"file_id":     "QmSVUyM2piAKC5Yc8YZUKh843aatQJWTDViozKscSk8WLd",
		"signature":   "0x8766f15e2e99342a9afac5c48dc88d8c3970d3b40ced7f268a220010ac47e348244096109988de30a315bb13559f5763158c5d8a48988c4b6126ffa323d71f251b",
	}
	req.Params = append(req.Params, param)
	fmt.Println("reqest : ", req.String())

	resp, err := getUserData(req)
	if resp.String() == "" || err != nil {
		fmt.Println("Error : ", err)
		t.Errorf("Failed to start main")
	}
	fmt.Println("Response : ", resp.String()) //QmSVUyM2piAKC5Yc8YZUKh843aatQJWTDViozKscSk8WLd

}
func TestRestoreUserData(t *testing.T) {
	defaultSetting()

	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "restore_user_data",
		ID:      1,
	}

	param := map[string]interface{}{
		"address_old": "0x961c20596e7EC441723FBb168461f4B51371D8aA",
		"address_new": "0xd3cb37ae6a81ebf1b5c3d9422636b0db48767b72",
		"meta_id_old": "0x400c2c9ae5f66133fb6a48d476756b9048a345a0666676e416cdaddf93ae09c5",

		"meta_id_new": "0x8f458c9b0b1aba820270bd06ba7019367b3044b0a92b40afb0ab852d723c61aa",

		"user_hash_list": []string{
			"0xa21be2c1a1951cafd81a22bffbd1d48294d45525b129d2002c9424c1d5dfb8ea",
			"0x4603ce8e282fe90eea9074d43991ba46afa0b3b23f3bc3148a055e6c51fc204a",
			"0xa21be2c1a1951cafd81a22bffbd1d48294d45525b129d2002c9424c1d5dfb8ea",
			"0x9e50e9545037b5ff006c9e421f63a6401554dfe20d2fb16d0ce04fa02c69b8f6",
			"0x288a85888972f2c365b88075940528d5a36451d9cf3dce6251e980e17e19846a",
			"0xed10665502644999df542f6ce39506cd1c84d8aab9886df731032beb38313cbf"},
		"signature": "0x8556ac67c81a2dfeed7cc55e270b0d88ebe2a1eb48e1181bfffc98dd75900548073f5485e0a6595896d2dace8ce2445f59f6150d10c7033e5b0fe0d9be419bb51c",
	}
	req.Params = append(req.Params, param)
	fmt.Println("reqest : ", req.String())

	resp, err := restoreUserData(req)
	if resp.String() == "" || err != nil {
		fmt.Println("Error : ", err)
		t.Errorf("Failed to start main")
	}
	fmt.Println("Response : ", resp.String()) //QmSVUyM2piAKC5Yc8YZUKh843aatQJWTDViozKscSk8WLd

}
func TestGetUserDataError(t *testing.T) {
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "get_user_data",
		ID:      1,
	}

	// param := map[string]interface{}{
	// 	"address_new": "0xd3cb37ae6a81ebf1b5c3d9422636b0db48767b72", // new Address
	// 	"file_id":     "QmSVUyM2piAKC5Yc8YZUKh843aatQJWTDViozKscSk8WLd",
	// 	"signature":   "0x8766f15e2e99342a9afac5c48dc88d8c3970d3b40ced7f268a220010ac47e348244096109988de30a315bb13559f5763158c5d8a48988c4b6126ffa323d71f251b",
	// }
	// req.Params = append(req.Params, param)
	fmt.Println("reqest : ", req.String())

	resp, err := getUserData(req)
	if resp.String() == "" || err != nil {
		fmt.Println("Error : ", err)
		t.Errorf("Failed to start main")
	}
	fmt.Println("Response : ", resp.String()) //QmSVUyM2piAKC5Yc8YZUKh843aatQJWTDViozKscSk8WLd

}
func TestConvert(t *testing.T) {

	metaId := "0x059e1e9bee92b3d2c425434ac3e7f076c10049b79bbfdf98fb0c5408217531f2"
	bytes := hexutil.MustDecode(metaId)
	//tmpBigInt := hexutil.MustDecodeBig(metaId)
	n := new(big.Int)
	n.SetBytes(bytes)
	fmt.Println(n)

}
