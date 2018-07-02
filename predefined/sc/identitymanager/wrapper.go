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
var session *IdentitymanagerSession

func getSession() (*IdentitymanagerSession, error) {
	once.Do(func() {
		service, err := getIMService()
		//auth := bind.NewKeyedTransactor(getPrivateKey())
		auth := crypto.GetTransactionOpts()
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

// var nsService *Nameservice
func getIMService() (*Identitymanager, error) {

	_rpc := rpc.GetInstance()
	client := _rpc.GetEthClient()

	imAddress, err := nameservice.GetIMContractAddress()
	instance, err := NewIdentitymanager(*imAddress, client)
	if err != nil {
		return nil, err
	}
	//	} else {
	return instance, nil
	//	}

}

//CallOwnerOf ownerOf function call
func CallOwnerOf(metaID hexutil.Bytes) (*common.Address, error) {

	service, err := getIMService()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	tmpBigInt := hexutil.MustDecodeBig(metaID.String())

	result, err := service.OwnerOf(&bind.CallOpts{}, tmpBigInt)
	if err != nil {
		if err.Error() == "abi: unmarshalling empty output" {
			return nil, nil
		}
		log.Fatal(err)
		return nil, err
	}
	fmt.Printf("Onwer Address: %x \n", result)
	return &result, nil
}

//CallCreateMetaID createMetaID function call
func CallCreateMetaID(metaID hexutil.Bytes, sig hexutil.Bytes, userAddress common.Address) (*types.Transaction, error) {
	var metaPack metaPackage
	metaPack = &metaPackageV1{
		Version:           1,
		UserSenderAddress: userAddress,
	}
	pack := metaPack.Serialize()
	fmt.Printf("Pack : %x", pack)
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

//CallUpdateMetaID  updateMetaID function call of IentityManager Smart Contract
func CallUpdateMetaID(oldMetaID hexutil.Bytes, newMetaID hexutil.Bytes, sig hexutil.Bytes, userAddress common.Address) (*types.Transaction, error) {

	var metaPack metaPackage
	metaPack = &metaPackageV1{
		Version:           1,
		UserSenderAddress: userAddress,
	}
	pack := metaPack.Serialize()
	fmt.Printf("Pack : %x", pack)
	session, err := getSession()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var oldID, newID [32]byte
	copy(oldID[:], oldMetaID)
	copy(newID[:], newMetaID)
	result, err := session.UpdateMetaID(oldID, newID, sig, pack)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil
}

// CallDeleteMetaID is delete MetaID
// TODO: timestamp 파라메터 추가 필요 (Smart contract 수정/반영후, 조치 예정 )
func CallDeleteMetaID(metaID hexutil.Bytes, sig hexutil.Bytes, userAddress common.Address) (*types.Transaction, error) {

	var metaPack metaPackage
	metaPack = &metaPackageV1{
		Version:           1,
		UserSenderAddress: userAddress,
	}
	pack := metaPack.Serialize()
	fmt.Printf("Pack : %x", pack)
	session, err := getSession()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var id [32]byte
	copy(id[:], metaID)

	result, err := session.DeleteMetaID(id, sig, pack)
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

//CallEcverify Ecverify function call
func CallEcverify(msg hexutil.Bytes, sig hexutil.Bytes, address common.Address) (bool, error) {

	service, err := getIMService()

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
	fmt.Printf("Ecverfiy: %v \n", result)
	return result, nil
}
