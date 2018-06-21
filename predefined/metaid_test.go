package predefined

import (
	"fmt"
	"testing"

	"bitbucket.org/coinplugin/proxy/json"
)

func TestMetaIdRegister(t *testing.T) {
	req := json.RpcRequest{
		Jsonrpc: "2.0",
		Method:  "register_meta_id",
		Id:      1,
	}
	param := map[string]interface{}{
		"address": "0x961c20596e7EC441723FBb168461f4B51371D8aA",
		"meta_id": "0xdb6dd8f5917a3c2f84a280f365ac137549e62d647b6cfba05a0f2c5e8e60e972",
		"user_hash_list": []string{"0xcbb6887db9a92ac59fa4c101d88f80b69fead4eea556ea486f276a05c3a99144",
			"0x9e50e9545037b5ff006c9e421f63a6401554dfe20d2fb16d0ce04fa02c69b8f6", "0x288a85888972f2c365b88075940528d5a36451d9cf3dce6251e980e17e19846a", "0xed10665502644999df542f6ce39506cd1c84d8aab9886df731032beb38313cbf"},
		"signature": "0x542e3d9af6758e80e4f2500395898a9cd9f5e3bd91b3053ad60d4cb0147b608c7de2ff608ae7528bc59cc6a7ba006020678d2ee5b5f88a97204734953cd2cccc1b",
	}
	req.Params = append(req.Params, param)
	fmt.Println("reqest : ", req.String())
	// resp, err := main.Handler(nil, events.APIGatewayProxyRequest{
	// 	Body: req.String(),
	// })
	resp, err := register_meta_id(req)
	if resp.String() == "" || err != nil {
		fmt.Println("Error : ", err)
		t.Errorf("Failed to start main")
	}
	fmt.Println("Response : ", resp.String())
}
