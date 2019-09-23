package servicekeyresolver

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/metadium/go-delegator/log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/metadium/go-delegator/crypto"
	"github.com/metadium/go-delegator/rpc"
)

var (
	once   sync.Once
	zero   = big.NewInt(0)
	glimit = uint64(2000000)
)

//GetInstance get Servicekeyresolver Instance
func GetInstance(address common.Address) (*Servicekeyresolver, error) {
	_rpc := rpc.GetInstance()
	client := _rpc.GetEthClient()

	instance, err := NewServicekeyresolver(address, client)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if instance == nil {
		err := fmt.Errorf("Cannot get Servicekeyresolver Instance")
		return nil, err
	}
	return instance, nil
}

//CallAddKeyDelegated addKeyDelegated function call
func CallAddKeyDelegated(reqID uint64, instance *Servicekeyresolver, associatedAddress common.Address, key common.Address, symbol string, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	var trx *types.Transaction
	var err error

	if instance == nil {
		err = fmt.Errorf("Error - ServiceKeyResolver nil")
		return nil, err
	}

	tx := func(nonce uint64) error {
		_rpc := rpc.GetInstance()
		auth := crypto.GetTransactionOpts()
		auth.Nonce = big.NewInt(int64(nonce))
		auth.GasPrice = big.NewInt(int64(_rpc.GetGasPrice()))
		auth.GasLimit = glimit
		log.Debug("gas price : ", auth.GasPrice)
		trx, err = instance.AddKeyDelegated(auth, associatedAddress, key, symbol, v, r, s, timestamp)
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
			err = fmt.Errorf("call function Error - CallAddKeyDelegated")
		}
		return nil, err
	}
	return trx, nil

}

//CallRemoveKeyDelegated RemoveKeyDelegated function call
func CallRemoveKeyDelegated(reqID uint64, instance *Servicekeyresolver, associatedAddress common.Address, key common.Address, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	var trx *types.Transaction
	var err error

	if instance == nil {
		err = fmt.Errorf("Error - ServiceKeyResolver nil")
		return nil, err
	}

	tx := func(nonce uint64) error {
		_rpc := rpc.GetInstance()
		auth := crypto.GetTransactionOpts()
		auth.Nonce = big.NewInt(int64(nonce))
		auth.GasPrice = big.NewInt(int64(_rpc.GetGasPrice()))
		auth.GasLimit = glimit
		log.Debug("gas price : ", auth.GasPrice)
		trx, err = instance.RemoveKeyDelegated(auth, associatedAddress, key, v, r, s, timestamp)
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
			err = fmt.Errorf("call function Error - CallRemoveKeyDelegated")
		}
		return nil, err
	}

	return trx, nil

}

//CallRemoveKeysDelegated RemoveKeysDelegated function call
func CallRemoveKeysDelegated(reqID uint64, instance *Servicekeyresolver, associatedAddress common.Address, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	var trx *types.Transaction
	var err error

	if instance == nil {
		err = fmt.Errorf("Error - ServiceKeyResolver nil")
		return nil, err
	}

	tx := func(nonce uint64) error {
		_rpc := rpc.GetInstance()
		auth := crypto.GetTransactionOpts()
		auth.Nonce = big.NewInt(int64(nonce))
		auth.GasPrice = big.NewInt(int64(_rpc.GetGasPrice()))
		auth.GasLimit = glimit
		log.Debug("gas price : ", auth.GasPrice)

		trx, err = instance.RemoveKeysDelegated(auth, associatedAddress, v, r, s, timestamp)
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
			err = fmt.Errorf("call function Error - CallRemoveKeyDelegated")
		}
		return nil, err
	}

	return trx, nil

}

//GetAddress return current ServiceKeyResolver contract address
func GetAddress() *common.Address {
	return &skrAddress
}

//GetAddressList Return all old and current ServiceKeyResolver contract address list
func GetAddressList() []common.Address {
	return allSkrAddress
}
