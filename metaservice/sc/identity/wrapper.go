package identity

import (
	"bytes"
	"fmt"
	"math/big"
	"sync"

	"github.com/metadium/go-delegator/log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/metadium/go-delegator/crypto"
	"github.com/metadium/go-delegator/rpc"
	//	"github.com/ethereum/go-ethereum/crypto"
)

var (
	once   sync.Once
	zero   = big.NewInt(0)
	glimit = uint64(4000000)

	recoveryFuncHash = hexutil.MustDecode("0x1d381240")
	emptyKeyVal      = new([32]byte)
)

func GetInstance(address common.Address) (*Identity, error) {
	_rpc := rpc.GetInstance()
	client := _rpc.GetEthClient()

	instance, err := NewIdentity(address, client)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if instance == nil {
		err := fmt.Errorf("Cannot get MetaID Instance")
		return nil, err
	}
	return instance, nil
}

//CallDelegatedExecute DelegatedExecute function call
func CallDelegatedExecute(instance *Identity, mgtAddress common.Address, to common.Address, value *big.Int, data hexutil.Bytes, metaNonce *big.Int, signature hexutil.Bytes) (*types.Transaction, error) {
	var trx *types.Transaction
	var err error

	// instance, err := getInstance(mgtAddress)
	// //session, err := getSession()

	// if err != nil {
	// 	log.Error(err)
	// 	return nil, err
	// }
	if instance == nil {
		err = fmt.Errorf("Error - Identity nil")
		return nil, err
	}

	tx := func(nonce uint64) error {
		_rpc := rpc.GetInstance()
		auth := crypto.GetTransactionOpts()
		auth.Nonce = big.NewInt(int64(nonce))
		auth.GasPrice = big.NewInt(int64(_rpc.GetGasPrice()))
		auth.GasLimit = glimit
		log.Debug("gas price : ", auth.GasPrice)
		trx, err = instance.DelegatedExecute(auth, to, value, data, metaNonce, signature)
		if err != nil {
			return err
		}
		return nil
	}
	c := crypto.GetInstance()
	res := c.ApplyNonce(tx)

	if !res {
		if err == nil {
			err = fmt.Errorf("call function Error - CallDelegatedExecute")
		}
		return nil, err
	}

	return trx, nil
}

//CallDelegatedApprove DelegatedApprove function call
func CallDelegatedApprove(instance *Identity, mgtAddress common.Address, id *big.Int, approve bool, metaNonce *big.Int, signature hexutil.Bytes) (*types.Transaction, error) {
	var trx *types.Transaction
	var err error

	// instance, err := getInstance(mgtAddress)
	// //session, err := getSession()

	// if err != nil {
	// 	log.Error(err)
	// 	return nil, err
	// }
	if instance == nil {
		err = fmt.Errorf("Error - Identity nil")
		return nil, err
	}

	tx := func(nonce uint64) error {
		_rpc := rpc.GetInstance()
		auth := crypto.GetTransactionOpts()
		auth.Nonce = big.NewInt(int64(nonce))
		auth.GasPrice = big.NewInt(int64(_rpc.GetGasPrice()))
		auth.GasLimit = glimit
		log.Debug("gas price : ", auth.GasPrice)
		trx, err = instance.DelegatedApprove(auth, id, approve, metaNonce, signature)

		if err != nil {
			return err
		}
		return nil
	}
	c := crypto.GetInstance()
	res := c.ApplyNonce(tx)

	if !res {
		if err == nil {
			err = fmt.Errorf("call function Error - CallDelegatedApprove")
		}
		return nil, err
	}

	return trx, nil
}

func CallGetTransactionCount(instance *Identity) (*big.Int, error) {
	var err error
	if instance == nil {
		err = fmt.Errorf("Error - Identity nil")
		return nil, err
	}

	result, err := instance.GetTransactionCount(&bind.CallOpts{})
	if err != nil {
		if err.Error() == "abi: unmarshalling empty output" {
			return common.Big0, nil
		}
		log.Error(err)
		return nil, err
	}

	log.Infof("Transction Count(%x): %x ", instance.Address, result)
	return result, nil

}

func CheckDelegateExecutePermission(instance *Identity, from common.Address, to common.Address, data hexutil.Bytes) error {
	var err error

	//1. convert addr to key
	keyBytes := CallAddrToKey(from)
	// keyBytes, err := CallAddrToKey(instance, from)
	// if err != nil {
	// 	return err
	// }
	// if keyBytes == nil {
	// 	return fmt.Errorf("fail to check Permission")
	// }
	//2. get Key
	key, err := CallGetKey(instance, keyBytes)
	if err != nil {
		return err
	}
	if key == nil {
		return fmt.Errorf("Not found key")
	}
	//3. check purpose
	if bytes.Equal(instance.Address.Bytes(), to.Bytes()) { //For Management or Recovery
		funcHash, err := CallGetFunctionSignature(instance, data)
		if err != nil {
			return err
		}
		for _, p := range key.Purposes {
			//management
			if common.Big1.Cmp(p) == 0 {
				log.Debug("vaild permission for SELF - management Key")
				return nil
			} else if big.NewInt(7).Cmp(p) == 0 && bytes.Equal(recoveryFuncHash, funcHash[:]) {
				log.Debug("vaild permission for SELF - recovery Key")
				return nil
			}
		}
	} else { //For Action
		for _, p := range key.Purposes {
			//management or action key
			if common.Big1.Cmp(p) == 0 {
				log.Debug("vaild permission - management Key")
				return nil
			} else if common.Big2.Cmp(p) == 0 {
				log.Debug("vaild permission - action Key")
				return nil
			}
		}
	}
	return fmt.Errorf("Not permission")
}

func CheckDelegateApprovePermission(instance *Identity, from common.Address) error {
	var err error

	//1. convert addr to key
	keyBytes := CallAddrToKey(from)

	//2. get Key
	key, err := CallGetKey(instance, keyBytes)
	if err != nil {
		return err
	}
	if key == nil {
		return fmt.Errorf("Not found key")
	}
	//3. check purpose
	//For Action
	for _, p := range key.Purposes {
		//management or action key
		if common.Big1.Cmp(p) == 0 {
			log.Debug("vaild permission - management Key")
			return nil
		} else if common.Big2.Cmp(p) == 0 {
			log.Debug("vaild permission - action Key")
			return nil
		}
	}

	return fmt.Errorf("Not permission")
}

func CallGetFunctionSignature(instance *Identity, data hexutil.Bytes) (*[4]byte, error) {
	var err error
	if instance == nil {
		err = fmt.Errorf("Error - Identity nil")
		return nil, err
	}

	result, err := instance.GetFunctionSignature(&bind.CallOpts{}, data)
	if err != nil {
		//데이터가 없을때 나는 에러 처리
		if err.Error() == "abi: unmarshalling empty output" {
			return nil, nil
		}
		log.Error(err)
		return nil, err
	}
	log.Debugf("functionSignature: %x ", result)
	return &result, nil
}

func CallGetKey(instance *Identity, key [32]byte) (*metaIDKey, error) {
	var err error
	if instance == nil {
		err = fmt.Errorf("Error - Identity nil")
		return nil, err
	}
	result, err := instance.GetKey(&bind.CallOpts{}, key)
	if err != nil {
		//데이터가 없을때 나는 에러 처리
		if err.Error() == "abi: unmarshalling empty output" {
			return nil, nil
		}
		log.Error(err)
		return nil, err
	}
	log.Debugf("key: %x ", result)
	return &result, nil
}

func CallGetKeysByPurpose(instance *Identity, purpose *big.Int) ([][32]byte, error) {
	var err error
	if instance == nil {
		err = fmt.Errorf("Error - Identity nil")
		return nil, err
	}
	result, err := instance.GetKeysByPurpose(&bind.CallOpts{}, purpose)
	if err != nil {
		//데이터가 없을때 나는 에러 처리
		if err.Error() == "abi: unmarshalling empty output" {
			return nil, nil
		}
		log.Error(err)
		return nil, err
	}
	log.Debugf("keys: %x ", result)
	return result, nil
}

/*
func CallAddrToKey(instance *Identity, address common.Address) ([32]byte, error) {

	var err error
	if instance == nil {
		err = fmt.Errorf("Error - Identity nil")
		return *emptyKeyVal, err
	}
	result, err := instance.AddrToKey(&bind.CallOpts{}, address)
	if err != nil {
		log.Error(err)
		return *emptyKeyVal, err
	}
	log.Debugf("address to key: %x ", result)
	return result, nil
}
*/
func CallAddrToKey(address common.Address) [32]byte {
	var result [32]byte
	copy(result[12:], address.Bytes())
	return result
}
