package identitymanager

import (
	"fmt"
	"math/big"
	"sync"

	"go-delegator/log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go-delegator/crypto"
	"go-delegator/metaservice/sc/registry"
	"go-delegator/rpc"
	//	"github.com/ethereum/go-ethereum/crypto"
)

var (
	once   sync.Once
	zero   = big.NewInt(0)
	glimit = uint64(4000000)
)

var instance *Identitymanager

func getService() (*Identitymanager, error) {
	once.Do(func() {
		_rpc := rpc.GetInstance()
		client := _rpc.GetEthClient()
		//var err error
		imAddress, err := registry.GetIMContractAddress()
		if err != nil {
			log.Error(err)
			return
		}
		instance, err = NewIdentitymanager(*imAddress, client)
		if err != nil {
			log.Error(err)
			return
		}
	})

	if instance == nil {
		err := fmt.Errorf("Cannot make service for IdentityManager")
		return nil, err
	}
	return instance, nil
}

//CallCreateMetaID createMetaID function call
func CallCreateMetaID(mgtAddress common.Address) (*types.Transaction, error) {
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

		trx, err = service.CreateMetaId(auth, mgtAddress)
		if err != nil {
			log.Error(err)
			return err
		}
		return nil
	}
	c := crypto.GetInstance()
	res := c.ApplyNonce(tx)

	if !res {
		if err == nil {
			err = fmt.Errorf("call function Error - CreateMetaID")
		}
		return nil, err
	}

	return trx, nil
}

/* Not Used
func CallGetDeployedMetaIds() ([]common.Address, error) {
	service, err := getService()

	if err != nil {
		log.Error(err)
		return nil, err
	}
	addrs, err := service.GetDeployedMetaIds(&bind.CallOpts{})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Debug("Address List: ", addrs)
	return addrs, nil
}
*/
/* Not Used
func CreateAddress(b common.Address, nonce uint64) common.Address {
    data, _ := rlp.EncodeToBytes([]interface{}{b, nonce})
    return common.BytesToAddress(Keccak256(data)[12:])
}
*/
