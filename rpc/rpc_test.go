package rpc

import (
	"context"
	"testing"

	"go-delegator/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestEthClient(t *testing.T) {
	NetType = Testnet
	r := GetInstance()
	client := r.GetEthClient()
	if client == nil {
		t.Errorf("Failed to GetEthClient")
	}

	_, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		t.Fatalf("%s", err)
	}
}

func TestEthClientRaw(t *testing.T) {
	//client, err := ethclient.Dial("https://rinkeby.infura.io")
	client, err := ethclient.Dial(TestnetUrls[0])
	if err != nil {
		t.Fatalf("%s", err)
	}
	_, err = client.SuggestGasPrice(context.Background())
	if err != nil {
		t.Fatalf("%s", err)
	}
}

func BenchmarkEthClient(b *testing.B) {
	client, err := ethclient.Dial(TestnetUrls[0])
	if err != nil {
		b.Fatalf("%s", err)
	}

	address := common.HexToAddress("0x00F912f1F41203DaE29b37fc18db8Dbd3cA9833F")
	errCnt := 0
	for i := 0; i < b.N; i++ {
		_, err := client.BalanceAt(context.Background(), address, nil)
		if err != nil {
			errCnt++
		}
	}
	b.Logf("errCnt %d", errCnt)
}

func BenchmarkHttpClient(b *testing.B) {
	NetType = Testnet
	r := GetInstance()
	req := json.RPCRequest{
		Jsonrpc: "2.0",
		ID:      1,
		Method:  "eth_getBalance",
	}
	req.Params = append(req.Params, "0x00F912f1F41203DaE29b37fc18db8Dbd3cA9833F")
	req.Params = append(req.Params, "latest")

	errCnt := 0
	for i := 0; i < b.N; i++ {
		_, err := r.DoRPC(req)
		if err != nil {
			errCnt++
		}
	}
	b.Logf("errCnt %d", errCnt)
}

func TestRefreshUrlList(t *testing.T) {
	NetType = Testnet
	r := GetInstance()
	initLen := len(TestnetUrls)
	target := TestnetUrls[0]
	for i := 0; i < 30; i++ {
		r.refreshURLList(target)
	}
	if (initLen - 1) != availLen[Testnet] {
		t.Errorf("refreshUrlList is abnormal")
	}
}

func TestCall(t *testing.T) {
	NetType = Testnet
	r := GetInstance()
	if _, err := r.Call("0x11", "0x123"); err != nil {
		t.Errorf("Failed to RPC Call")
	}
}

func TestGasPrice(t *testing.T) {
	NetType = Testnet
	r := GetInstance()
	if ret := r.GetGasPrice(); ret == 0 {
		t.Errorf("Failed to get gas price")
	} else {
		t.Logf("%d", ret)
	}
}

func TestRpc(t *testing.T) {
	testMsg := "{\"jsonrpc\":\"2.0\",\"method\":\"web3_clientVersion\",\"params\":[\"a\",1],\"id\":100}"

	NetType = Testnet
	r := GetInstance()
	// Test with string param
	if _, err := r.DoRPC(testMsg); err != nil {
		t.Errorf("Failed to RPC with string: %s", err)
	}

	// Test with RpcRequest param
	testRPCRequest := json.GetRPCRequestFromJSON(testMsg)
	if _, err := r.DoRPC(testRPCRequest); err != nil {
		t.Errorf("Failed to RPC with RpcRequest: %s", err)
	}
}

func BenchmarkRpc(b *testing.B) {
	testMsg := "{\"jsonrpc\":\"2.0\",\"method\":\"web3_clientVersion\",\"params\":[\"a\",1],\"id\":100}"

	NetType = Testnet
	r := GetInstance()
	for i := 0; i < b.N; i++ {
		r.DoRPC(testMsg)
	}
}
