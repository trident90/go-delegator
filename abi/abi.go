// Package abi implements smart contract call helper
package abi

import (
	"fmt"
	"math/big"
	"strings"

	"go-delegator/crypto"
	"go-delegator/json"
	"go-delegator/rpc"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

var (
	zero = big.NewInt(0)
)

// Pack makes packed data with inputs on ABI
func Pack(abi abi.ABI, name string, args ...interface{}) (string, error) {
	data, err := abi.Pack(name, args...)
	if err != nil {
		return "", err
	}
	return hexutil.Encode(data), nil
}

// Unpack fills output into given ABI
func Unpack(abi abi.ABI, v interface{}, name string, output string) error {
	// TODO: debug
	//var data []byte
	var err error
	/*
		if output[:2] == "0x" {
			data, err = hex.DecodeString(output[2:])
		} else {
			data, err = hex.DecodeString(output)
		}
	*/

	if err != nil {
		return err
	}
	//return abi.Unpack(v, name, data)
	return err
}

// Call gets contract value with contract address and name
func Call(abi abi.ABI, to, name string, inputs []interface{}) (resp json.RPCResponse, err error) {
	data, err := Pack(abi, name, inputs...)
	if err != nil {
		return
	}

	r := rpc.GetInstance()
	respStr, err := r.Call(to, data)
	if err != nil {
		return
	}

	resp = json.GetRPCResponseFromJSON(respStr)
	return
}

// SendTransaction calls smart contract with ABI using eth_sendTransaction
func SendTransaction(abi abi.ABI, to, name string, inputs []interface{}, gas int) (resp json.RPCResponse, err error) {
	var data string
	if data, err = Pack(abi, name, inputs...); err != nil {
		return
	}

	c := crypto.GetInstance()
	r := rpc.GetInstance()
	respStr, err := r.SendTransaction(c.GetAddress(), to, data, gas)
	if err != nil {
		return
	}

	resp = json.GetRPCResponseFromJSON(respStr)
	return
}

// SendTransactionWithSign calls smart contract with ABI using eth_sendRawTransaction
func SendTransactionWithSign(abi abi.ABI, to, name string, inputs []interface{}, gasLimit, gasPrice uint64) (resp json.RPCResponse, err error) {
	var data []byte
	if data, err = abi.Pack(name, inputs...); err != nil {
		return
	}

	c := crypto.GetInstance()
	r := rpc.GetInstance()

	// Make TX function to get nonce
	tx := func(nonce uint64) (err error) {
		tx := types.NewTransaction(nonce, common.HexToAddress(to), zero, uint64(gasLimit), big.NewInt(int64(gasPrice)), data)
		if tx, err = c.SignTx(tx); err != nil {
			return
		}

		var rlpTx []byte
		if rlpTx, err = rlp.EncodeToBytes(tx); err != nil {
			return
		}

		var respStr string
		if respStr, err = r.SendRawTransaction(rlpTx); err != nil {
			return
		}

		if resp = json.GetRPCResponseFromJSON(respStr); resp.Error == nil {
			return fmt.Errorf("%s", resp.Error.Message)
		}
		return
	}

	c.ApplyNonce(tx)
	return
}

// GetAbiFromJSON returns ABI object from JSON string
func GetAbiFromJSON(raw string) (abi.ABI, error) {
	return abi.JSON(strings.NewReader(raw))
}

// getAbiFromAddress is NOT YET SUPPORTED
// TODO: use eth.compile.solidity?
func getAbiFromAddress(addr string) (abi abi.ABI) {
	r := rpc.GetInstance()
	respStr, err := r.GetCode(addr)
	if err != nil {
		return
	}

	json.GetRPCResponseFromJSON(respStr)
	return
}
