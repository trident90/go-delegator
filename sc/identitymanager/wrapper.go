package identitymanager

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"sync"

	"bitbucket.org/coinplugin/geth/crypto"
	"bitbucket.org/coinplugin/proxy/sc/nameservice"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getPrivateKey() *ecdsa.PrivateKey {
	privKey := "01b149603ca8f537bbb4e45d22e77df9054e50d826bb5f0a34e9ce460432b596"
	key, _ := crypto.HexToECDSA(privKey)
	return key

}

var once sync.Once
var session *IdentitymanagerSession

func getSession() (*IdentitymanagerSession, error) {
	once.Do(func() {
		service, err := getIMService()
		auth := bind.NewKeyedTransactor(getPrivateKey())
		auth.Value = big.NewInt(0)
		auth.GasLimit = uint64(300000)
		if err != nil {
			log.Fatal(err)
			return
		}
		session = &IdentitymanagerSession{
			Contract: service,
			CallOpts: bind.CallOpts{
				Pending: true,
			},
			TransactOpts: *auth,
		}
	})
	if session == nil {
		err := fmt.Errorf("Cannot make session for IdentityManager")
		return nil, err
	}
	return session, nil
}

func getMetaClient() (*ethclient.Client, error) {
	//	if metaClient == nil {
	client, err := ethclient.Dial("http://13.125.247.228:8545")
	if err != nil {
		return nil, err
	}
	//		metaClient = client
	//	}
	return client, nil
}

// var nsService *Nameservice
func getIMService() (*Identitymanager, error) {
	// if nsService == nil {
	client, err := getMetaClient()
	if err != nil {
		return nil, err
	}

	imAddress, err := nameservice.GetIMContractAddress()
	instance, err := NewIdentitymanager(*imAddress, client)
	if err != nil {
		return nil, err
	}
	//	} else {
	return instance, nil
	//	}

}
func GetOwnerAddress(metaID hexutil.Bytes) (*common.Address, error) {

	service, err := getIMService()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	tmpBigInt := hexutil.MustDecodeBig(metaID.String())

	result, err := service.OwnerOf(&bind.CallOpts{}, tmpBigInt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Onwer Address: %x \n", result)
	return &result, nil
}

// func (_Identitymanager *IdentitymanagerTransactor) CreateMetaID(opts *bind.TransactOpts, _metaID [32]byte, _sig []byte, _metaPackage []byte) (*types.Transaction, error) {
// 	return _Identitymanager.contract.Transact(opts, "createMetaID", _metaID, _sig, _metaPackage)
// }

func CallCreateMetaID(metaID hexutil.Bytes, sig hexutil.Bytes, userAddress common.Address) (*types.Transaction, error) {
	var metaPack metaPackage
	metaPack = &metaPackageV1{
		Version:           1,
		UserSenderAddress: userAddress,
	}
	pack := metaPack.Serialize()

	session, err := getSession()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var id [32]byte
	copy(id[:], metaID)
	result, err := session.CreateMetaID(id, sig, pack)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil
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
