package identityregistry

import (
	"fmt"
	"math/big"
	"sync"

	"go-delegator/config"
	"go-delegator/log"

	"go-delegator/crypto"
	"go-delegator/rpc"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	once   sync.Once
	zero   = big.NewInt(0)
	glimit = uint64(2000000)
)

var instance *Identityregistry

func getService() (*Identityregistry, error) {
	once.Do(func() {
		_rpc := rpc.GetInstance()
		client := _rpc.GetEthClient()
		var err error

		instance, err = NewIdentityregistry(common.HexToAddress(config.Config.IdentityRegistryAddress), client)
		if err != nil {
			log.Error(err)
			return
		}
	})

	if instance == nil {
		err := fmt.Errorf("Cannot make service for IdentityRegistry")
		return nil, err
	}
	return instance, nil
}

// CallCreateIdentity CreateIdentity function call
func CallCreateIdentity(reqID uint64, recoveryAddress common.Address, associatedAddress common.Address, providers []common.Address, resolvers []common.Address, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {

	var trx *types.Transaction
	var err error
	service, err := getService()
	//session, err := getSession()

	if err != nil {
		log.Error(err)
		return nil, err
	}

	tx := func(nonce uint64) error {
		_rpc := rpc.GetInstance()
		auth := crypto.GetTransactionOpts()
		auth.Nonce = big.NewInt(int64(nonce))
		auth.GasPrice = big.NewInt(int64(_rpc.GetGasPrice()))
		auth.GasLimit = glimit
		trx, err = service.CreateIdentityDelegated(auth, recoveryAddress, associatedAddress, providers, resolvers, v, r, s, timestamp)
		if err != nil {
			log.Error(err)
			return err
		}
		log.Debugfd(reqID, "trxid : %v\n", trx.Hash().String())

		return nil
	}
	c := crypto.GetInstance()
	res := c.ApplyNonce(tx)

	if !res {
		if err == nil {
			err = fmt.Errorf("call function Error - CreateIdentity")
		}
		return nil, err
	}

	return trx, nil
}

// CallAddAssociatedAddressDelegated  AddAssociatedAddressDelegated function call
func CallAddAssociatedAddressDelegated(reqID uint64, approvingAddress common.Address, addressToAdd common.Address, v [2]uint8, r [2][32]byte, s [2][32]byte, timestamp [2]*big.Int) (*types.Transaction, error) {

	var trx *types.Transaction
	var err error
	service, err := getService()
	//session, err := getSession()

	if err != nil {
		log.Error(err)
		return nil, err
	}

	tx := func(nonce uint64) error {

		_rpc := rpc.GetInstance()
		auth := crypto.GetTransactionOpts()
		auth.Nonce = big.NewInt(int64(nonce))
		auth.GasPrice = big.NewInt(int64(_rpc.GetGasPrice()))
		auth.GasLimit = glimit
		trx, err = service.AddAssociatedAddressDelegated(auth, approvingAddress, addressToAdd, v, r, s, timestamp)
		if err != nil {
			log.Error(err)
			return err
		}
		log.Debugfd(reqID, "trxid : %v", trx.Hash().String())

		return nil
	}
	c := crypto.GetInstance()
	res := c.ApplyNonce(tx)
	if !res {
		if err == nil {
			err = fmt.Errorf("call function Error - AddAssociatedAddressDelegated")
		}
		return nil, err
	}

	return trx, nil
}

// CallRemoveAssociatedAddressDelegated  RemoveAssociatedAddressDelegated function call
func CallRemoveAssociatedAddressDelegated(reqID uint64, addressToRemove common.Address, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {

	var trx *types.Transaction
	var err error
	service, err := getService()
	//session, err := getSession()

	if err != nil {
		log.Error(err)
		return nil, err
	}

	tx := func(nonce uint64) error {

		_rpc := rpc.GetInstance()
		auth := crypto.GetTransactionOpts()
		auth.Nonce = big.NewInt(int64(nonce))
		auth.GasPrice = big.NewInt(int64(_rpc.GetGasPrice()))
		auth.GasLimit = glimit
		trx, err = service.RemoveAssociatedAddressDelegated(auth, addressToRemove, v, r, s, timestamp)
		if err != nil {
			log.Error(err)
			return err
		}
		log.Debugfd(reqID, "trxid : %v", trx.Hash().String())

		return nil
	}
	c := crypto.GetInstance()
	res := c.ApplyNonce(tx)

	if !res {
		if err == nil {
			err = fmt.Errorf("call function Error - RemoveAssociatedAddressDelegated")
		}
		return nil, err
	}

	return trx, nil
}

// CallAddResolversFor  AddResolversFor function call
func CallAddResolversFor(reqID uint64, ein *big.Int, resolvers []common.Address) (*types.Transaction, error) {

	var trx *types.Transaction
	var err error
	service, err := getService()
	//session, err := getSession()

	if err != nil {
		log.Error(err)
		return nil, err
	}

	tx := func(nonce uint64) error {

		_rpc := rpc.GetInstance()
		auth := crypto.GetTransactionOpts()
		auth.Nonce = big.NewInt(int64(nonce))
		auth.GasPrice = big.NewInt(int64(_rpc.GetGasPrice()))
		auth.GasLimit = glimit
		trx, err = service.AddResolversFor(auth, ein, resolvers)
		if err != nil {
			log.Error(err)
			return err
		}
		log.Debugfd(reqID, "trxid : %v", trx.Hash().String())

		return nil
	}
	c := crypto.GetInstance()
	res := c.ApplyNonce(tx)
	if !res {
		if err == nil {
			err = fmt.Errorf("call function Error - AddAssociatedAddressDelegated")
		}
		return nil, err
	}

	return trx, nil
}

// CallGetEIN  get ein for associated address
func CallGetEIN(reqID uint64, associatedAddress common.Address) (*big.Int, error) {
	var err error

	service, err := getService()

	if err != nil {
		log.Error(err)
		return nil, err
	}

	result, err := service.GetEIN(&bind.CallOpts{}, associatedAddress)
	if err != nil {
		// if err.Error() == "abi: unmarshalling empty output" {
		// 	return common.Big0, nil
		// }
		log.Error(err)
		return nil, err
	}

	log.Debugfd(reqID, "EIN for (%x): %x ", associatedAddress, result)
	return result, nil

}

// CallIsProviderFor Checks whether the passed provider is set for the passed EIN.
func CallIsProviderFor(reqID uint64, ein *big.Int, provider common.Address) (bool, error) {
	var err error

	service, err := getService()

	if err != nil {
		log.Error(err)
		return false, err
	}

	result, err := service.IsProviderFor(&bind.CallOpts{}, ein, provider)
	if err != nil {
		// if err.Error() == "abi: unmarshalling empty output" {
		// 	return common.Big0, nil
		// }
		log.Error(err)
		return false, err
	}

	log.Debugfd(reqID, "IsProviderFor for (%v): %v ", ein, result)
	return result, nil

}

// CallIsResolverFor Checks whether the passed resolver is set for the passed EIN.
func CallIsResolverFor(reqID uint64, ein *big.Int, provider common.Address) (bool, error) {
	var err error

	service, err := getService()

	if err != nil {
		log.Error(err)
		return false, err
	}

	result, err := service.IsResolverFor(&bind.CallOpts{}, ein, provider)
	if err != nil {
		// if err.Error() == "abi: unmarshalling empty output" {
		// 	return common.Big0, nil
		// }
		log.Error(err)
		return false, err
	}

	log.Debugfd(reqID, "IsResolverFor for (%v): %v ", ein, result)
	return result, nil

}

// GetAddress Get IdentityRegistry Contract Address deployed by metadium
func GetAddress() *common.Address {
	irAddr := common.HexToAddress(config.Config.IdentityRegistryAddress)
	return &irAddr
}
