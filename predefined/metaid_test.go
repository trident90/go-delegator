package predefined

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"

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
		"meta_id": "0xdb6dd8f5917a3c2f84a280f365ac137549e62d647b6cfba05a0f2c5e8e60e972",
		"user_hash_list": []string{
			"0xcbb6887db9a92ac59fa4c101d88f80b69fead4eea556ea486f276a05c3a99144",
			"0x9e50e9545037b5ff006c9e421f63a6401554dfe20d2fb16d0ce04fa02c69b8f6",
			"0x288a85888972f2c365b88075940528d5a36451d9cf3dce6251e980e17e19846a",
			"0xed10665502644999df542f6ce39506cd1c84d8aab9886df731032beb38313cbf"},
		"signature": "0x542e3d9af6758e80e4f2500395898a9cd9f5e3bd91b3053ad60d4cb0147b608c7de2ff608ae7528bc59cc6a7ba006020678d2ee5b5f88a97204734953cd2cccb1b",
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
		"address":     "0x961c20596e7EC441723FBb168461f4B51371D8aA",
		"meta_id_old": "0xdb6dd8f5917a3c2f84a280f365ac137549e62d647b6cfba05a0f2c5e8e60e972",
		"meta_id_new": "0xc30255d8455a90fbb0ce11405b501624be9287370a41779312944c8739a9f79d",
		"user_hash_list": []string{
			"0xcbb6887db9a92ac59fa4c101d88f80b69fead4eea556ea486f276a05c3a99144",
			"0x9e50e9545037b5ff006c9e421f63a6401554dfe20d2fb16d0ce04fa02c69b8f6",
			"0x288a85888972f2c365b88075940528d5a36451d9cf3dce6251e980e17e19846a",
			"0xe3114197cb884bbaf2fd098ef55437c832e69bb0b5aee4b535eb5db7b0cdedc5"},
		"signature": "0x0f86138d24ecd75efd982e8c68555c97c70d5abfa67ab99edeca462b9b546d706888437eecc91f1ea5a0e815606188b6e91353bd867c9ad08cf2f62781aefe391c",
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
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		Method:  "restore_user_data",
		ID:      1,
	}

	param := map[string]interface{}{
		"address_old": "0x961c20596e7EC441723FBb168461f4B51371D8aA",
		"address_new": "0xd3cb37ae6a81ebf1b5c3d9422636b0db48767b72",
		"meta_id_old": "0xbbf9e4bcbee9add053ed959fb3d8bc808e6bb7d0a16bc1137f19030d5b4499ff",
		"meta_id_new": "0x9face836403006cd9f013c62e6e4dc73a7ad37e1fa362a595e1b7a808d486796",

		"user_hash_list": []string{
			"0xa21be2c1a1951cafd81a22bffbd1d48294d45525b129d2002c9424c1d5dfb8ea",
			"0x9e50e9545037b5ff006c9e421f63a6401554dfe20d2fb16d0ce04fa02c69b8f6",
			"0x288a85888972f2c365b88075940528d5a36451d9cf3dce6251e980e17e19846a",
			"0xed10665502644999df542f6ce39506cd1c84d8aab9886df731032beb38313cbf"},
		"signature": "0x0d6f04ea002960f37a496b50dabbcc79e2a602136b374a9000586c553b621eed45cf325f0a86c610b5ab9750f2530dec308c9dee9ec92279f85c9a552cc255a51b",
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
