package identitymanager

import (
	"fmt"
	"log"
	"math/big"
	"sync"

	"bitbucket.org/coinplugin/proxy/crypto"
	"bitbucket.org/coinplugin/proxy/predefined/sc/nameservice"
	"bitbucket.org/coinplugin/proxy/rpc"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	//	"github.com/ethereum/go-ethereum/crypto"
)

var once sync.Once

// var session *IdentitymanagerSession
var instance *Identitymanager

// func getSession() (*IdentitymanagerSession, error) {
// 	once.Do(func() {
// 		service, err := getIMService()
// 		//auth := bind.NewKeyedTransactor(getPrivateKey())
// 		auth := crypto.GetTransactionOpts()
// 		auth.Value = big.NewInt(0)
// 		auth.GasLimit = uint64(300000)
// 		if err != nil {
// 			log.Fatal(err)
// 			return
// 		}
// 		session = &IdentitymanagerSession{
// 			Contract: service,
// 			CallOpts: bind.CallOpts{
// 				Pending: true,
// 			},
// 			TransactOpts: *auth,
// 		}
// 	})
// 	if session == nil {
// 		err := fmt.Errorf("Cannot make session for IdentityManager")
// 		return nil, err
// 	}
// 	return session, nil
// }
// func getIMService() (*Identitymanager, error) {

// 	_rpc := rpc.GetInstance()
// 	client := _rpc.GetEthClient()

// 	imAddress, err := nameservice.GetIMContractAddress()
// 	service, err := NewIdentitymanager(*imAddress, client)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return service, nil

// }

func getService() (*Identitymanager, error) {
	once.Do(func() {
		_rpc := rpc.GetInstance()
		client := _rpc.GetEthClient()
		//var err error
		imAddress, err := nameservice.GetIMContractAddress()
		if err != nil {
			log.Fatal(err)
			return
		}
		instance, err = NewIdentitymanager(*imAddress, client)
		if err != nil {
			log.Fatal(err)
			return
		}
	})

	if instance == nil {
		err := fmt.Errorf("Cannot make service for IdentityManager")
		return nil, err
	}
	return instance, nil
	//	}

}

//CallOwnerOf ownerOf function call
func CallOwnerOf(metaID hexutil.Bytes) (*common.Address, error) {

	service, err := getService()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	//tmpBigInt := hexutil.MustDecodeBig(metaID.String())
	tmpBigInt := new(big.Int)
	tmpBigInt.SetBytes(metaID)
	result, err := service.OwnerOf(&bind.CallOpts{}, tmpBigInt)
	if err != nil {
		if err.Error() == "abi: unmarshalling empty output" {
			return nil, nil
		}
		log.Fatal(err)
		return nil, err
	}
	log.Printf("Onwer Address: %x \n", result)
	return &result, nil
}

//CallCreateMetaID createMetaID function call
func CallCreateMetaID(metaID hexutil.Bytes, sig hexutil.Bytes, userAddress common.Address) (*types.Transaction, error) {
	var trx *types.Transaction
	var err error

	var metaPack metaPackage
	metaPack = &metaPackageV1{
		Version:           1,
		UserSenderAddress: userAddress,
	}
	pack := metaPack.Serialize()
	log.Printf("Pack : %x", pack)
	service, err := getService()
	//session, err := getSession()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	tx := func(nonce uint64) error {
		_rpc := rpc.GetInstance()
		auth := crypto.GetTransactionOpts()
		auth.Nonce = big.NewInt(int64(nonce))
		auth.GasPrice = big.NewInt(int64(_rpc.GetGasPrice()))

		var id [32]byte
		copy(id[:], metaID)

		trx, err = service.CreateMetaID(auth, id, sig, pack)
		if err != nil {
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

	// var id [32]byte
	// copy(id[:], metaID)
	// result, err := service.CreateMetaID(auth, id, sig, pack)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return nil, err
	// }
	return trx, nil
}

//CallUpdateMetaID  updateMetaID function call of IentityManager Smart Contract
func CallUpdateMetaID(oldMetaID hexutil.Bytes, newMetaID hexutil.Bytes, sig hexutil.Bytes, userAddress common.Address) (*types.Transaction, error) {
	var trx *types.Transaction
	var err error

	var metaPack metaPackage
	metaPack = &metaPackageV1{
		Version:           1,
		UserSenderAddress: userAddress,
	}
	pack := metaPack.Serialize()
	log.Printf("Pack : %x", pack)
	service, err := getService()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// auth := crypto.GetTransactionOpts()
	// auth.Nonce = big.NewInt(int64(nonce))
	// var oldID, newID [32]byte
	// copy(oldID[:], oldMetaID)
	// copy(newID[:], newMetaID)
	// result, err := service.UpdateMetaID(auth, oldID, newID, sig, pack)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return nil, err
	// }

	tx := func(nonce uint64) error {
		_rpc := rpc.GetInstance()
		auth := crypto.GetTransactionOpts()
		auth.Nonce = big.NewInt(int64(nonce))
		auth.GasPrice = big.NewInt(int64(_rpc.GetGasPrice()))

		var oldID, newID [32]byte
		copy(oldID[:], oldMetaID)
		copy(newID[:], newMetaID)

		trx, err = service.UpdateMetaID(auth, oldID, newID, sig, pack)
		if err != nil {
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

//CallRestoreMetaID   change new meta id,address  from old meta id, address
//TODO: Contract 구현이 안되어 있음.
func CallRestoreMetaID(oldMetaID hexutil.Bytes, newMetaID hexutil.Bytes, oldAddress common.Address, newAddress common.Address, sig hexutil.Bytes) (*types.Transaction, error) {
	// var trx *types.Transaction
	var err error

	// var metaPack metaPackage
	// metaPack = &metaPackageV1{
	// 	Version:           1,
	// 	UserSenderAddress: Address,
	// }
	// pack := metaPack.Serialize()
	// log.Printf("Pack : %x", pack)
	// service, err := getService()

	// if err != nil {
	// 	log.Fatal(err)
	// 	return nil, err
	// }

	// tx := func(nonce uint64) error {
	// 	_rpc := rpc.GetInstance()
	// 	auth := crypto.GetTransactionOpts()
	// 	auth.Nonce = big.NewInt(int64(nonce))
	// 	auth.GasPrice = big.NewInt(int64(_rpc.GetGasPrice()))

	// 	var oldID, newID [32]byte
	// 	copy(oldID[:], oldMetaID)
	// 	copy(newID[:], newMetaID)

	// 	trx, err = service.UpdateMetaID(auth, oldID, newID, sig, pack)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	return nil
	// }
	// c := crypto.GetInstance()
	// res := c.ApplyNonce(tx)

	// if !res {
	// 	if err == nil {
	// 		err = fmt.Errorf("call function Error - CreateMetaID")
	// 	}
	// 	return nil, err
	// }
	// return trx, nil
	err = fmt.Errorf("Not implemented function.")
	return nil, err
}

// CallDeleteMetaID is delete MetaID
func CallDeleteMetaID(metaID hexutil.Bytes, timestamp hexutil.Bytes, sig hexutil.Bytes) (*types.Transaction, error) {

	service, err := getService()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	/*
		auth := crypto.GetTransactionOpts()
		auth.Nonce = big.NewInt(int64(nonce))

		var id [32]byte
		copy(id[:], metaID)

		result, err := service.DeleteMetaID(auth, id, timestamp, sig)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		return result, nil
	*/
	var trx *types.Transaction

	tx := func(nonce uint64) error {
		_rpc := rpc.GetInstance()
		auth := crypto.GetTransactionOpts()
		auth.Nonce = big.NewInt(int64(nonce))
		auth.GasPrice = big.NewInt(int64(_rpc.GetGasPrice()))
		var id [32]byte
		copy(id[:], metaID)

		trx, err = service.DeleteMetaID(auth, id, timestamp, sig)
		if err != nil {
			return err
		}
		return nil
	}
	c := crypto.GetInstance()
	res := c.ApplyNonce(tx)

	if !res {
		if err == nil {
			err = fmt.Errorf("call function Error -DeleteMetaID")
		}
		return nil, err
	}

	// var id [32]byte
	// copy(id[:], metaID)
	// result, err := service.CreateMetaID(auth, id, sig, pack)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return nil, err
	// }
	return trx, nil

}

type metaPackage interface {
	Serialize() []byte
}
type metaPackageV1 struct {
	Version           byte
	UserSenderAddress common.Address
	AttestationMask   [32]byte
	Status            [32]byte
}

func (t *metaPackageV1) Serialize() []byte {
	var bytes [1 + 20 + 32 + 32]byte
	bytes[0] = t.Version
	copy(bytes[1:21], t.UserSenderAddress[:])
	copy(bytes[21:53], t.AttestationMask[:])
	copy(bytes[53:85], t.Status[:])

	return bytes[:]
}

//CallEcverify Ecverify function call
func CallEcverify(msg hexutil.Bytes, sig hexutil.Bytes, address common.Address) (bool, error) {

	service, err := getService()

	if err != nil {
		log.Fatal(err)
		return false, err
	}
	var msgBytes [32]byte
	copy(msgBytes[:], msg)

	result, err1 := service.Ecverify(&bind.CallOpts{}, msgBytes, sig, address)

	if err1 != nil {
		log.Fatal(err1)
		return false, err1
	}
	log.Printf("Ecverfiy: %v \n", result)
	return result, nil
}
