package publickeyresolver

import (
	"bytes"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"go-delegator/log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go-delegator/crypto"
	"go-delegator/rpc"
)

var (
	once   sync.Once
	zero   = big.NewInt(0)
	glimit = uint64(2000000)
)

//GetInstance get Servicekeyresolver Instance
func GetInstance(address common.Address) (*Publickeyresolver, error) {
	_rpc := rpc.GetInstance()
	client := _rpc.GetEthClient()

	instance, err := NewPublickeyresolver(address, client)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if instance == nil {
		err := fmt.Errorf("Cannot get Publickeyresolver Instance")
		return nil, err
	}
	return instance, nil
}

//CallAddPublicKeyDelegated addKeyDelegated function call
func CallAddPublicKeyDelegated(reqID uint64, instance *Publickeyresolver, associatedAddress common.Address, publickey hexutil.Bytes, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	var trx *types.Transaction
	var err error

	if instance == nil {
		err = fmt.Errorf("Error - Publickeyresolver nil")
		return nil, err
	}

	tx := func(nonce uint64) error {
		_rpc := rpc.GetInstance()
		auth := crypto.GetTransactionOpts()
		auth.Nonce = big.NewInt(int64(nonce))
		auth.GasPrice = big.NewInt(int64(_rpc.GetGasPrice()))
		auth.GasLimit = glimit
		log.Debug("gas price : ", auth.GasPrice)
		trx, err = instance.AddPublicKeyDelegated(auth, associatedAddress, publickey, v, r, s, timestamp)
		if err != nil {
			return err
		}
		log.Debugfd(reqID, "trxid : %v", trx.Hash().String())
		return nil
	}
	c := crypto.GetInstance()
	res := c.ApplyNonce(tx)

	if !res {
		if err == nil {
			err = fmt.Errorf("call function Error - CallAddPublicKeyDelegated")
		}
		return nil, err
	}
	return trx, nil

}

//CallRemovePublicKeyDelegated RemoveKeyDelegated function call
func CallRemovePublicKeyDelegated(reqID uint64, instance *Publickeyresolver, associatedAddress common.Address, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	var trx *types.Transaction
	var err error

	if instance == nil {
		err = fmt.Errorf("Error - PublicKeyResolver nil")
		return nil, err
	}

	tx := func(nonce uint64) error {
		_rpc := rpc.GetInstance()
		auth := crypto.GetTransactionOpts()
		auth.Nonce = big.NewInt(int64(nonce))
		auth.GasPrice = big.NewInt(int64(_rpc.GetGasPrice()))
		auth.GasLimit = glimit
		log.Debug("gas price : ", auth.GasPrice)
		trx, err = instance.RemovePublicKeyDelegated(auth, associatedAddress, v, r, s, timestamp)
		if err != nil {
			return err
		}
		log.Debugfd(reqID, "trxid : %v", trx.Hash().String())
		return nil
	}
	c := crypto.GetInstance()
	res := c.ApplyNonce(tx)

	if !res {
		if err == nil {
			err = fmt.Errorf("call function Error - CallRemovePublicKeyDelegated")
		}
		return nil, err
	}

	return trx, nil

}

//GetAddress return current ServiceKeyResolver contract address
func GetAddress() *common.Address {
	return &pkrAddress
}

//GetAddressList Return all old and current ServiceKeyResolver contract address list
func GetAddressList() []common.Address {
	return allPkrAddress
}

func ContainsInAddresses(_address common.Address) bool {

	for _, a := range allPkrAddress {
		if bytes.Equal(a.Bytes(), _address.Bytes()) {
			return true
		}
	}
	return false
}
