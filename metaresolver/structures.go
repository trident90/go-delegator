package metaresolver

import (
	encodingJson "encoding/json"
	"fmt"
	"math/big"
	"reflect"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"go-delegator/log"
	"gopkg.in/validator.v2"
)

const (
	// 	addressSize         = 20
	// 	hashSize            = 32
	// 	signatureSize       = 65
	// 	fileHashSize        = 46
	// 	timestampSize       = 14
	// 	minimumUserHashSize = 3
	// 	minimumUserDataSize = 10
	addressHashIdx = 2
)

var headerBytes = hexutil.MustDecode("0x1900")

type identityCreateParams struct {
	RecoveryAddress   common.Address   `json:"recovery_address"`
	AssociatedAddress common.Address   `json:"associated_address"`
	Providers         []common.Address `json:"providers"`
	Resolvers         []common.Address `json:"resolvers" validate:"min=1"` //itemlen=32 : type 이 []hexutil.Bytes 일때, 각 항목의 길이 체크
	V                 hexutil.Bytes    `json:"v" validate:"len=1"`
	R                 hexutil.Bytes    `json:"r" validate:"len=32"`
	S                 hexutil.Bytes    `json:"s" validate:"len=32"`
	Timestamp         *big.Int         `json:"timestamp"`
}

func (p *identityCreateParams) Keccak256(identityRegistryAddr *common.Address) (hash []byte, data []byte) {
	irAddrBytes := identityRegistryAddr.Bytes()
	providersBytes := addressArrayToBytes(p.Providers)
	resolversBytes := addressArrayToBytes(p.Resolvers)
	timestampBytes := bigIntToByte32(p.Timestamp)
	data = append(data, headerBytes...)
	data = append(data, irAddrBytes...)
	data = append(data, []byte("I authorize the creation of an Identity on my behalf.")...)
	data = append(data, p.RecoveryAddress.Bytes()...)
	data = append(data, p.AssociatedAddress.Bytes()...)
	data = append(data, providersBytes...)
	data = append(data, resolversBytes...)
	data = append(data, timestampBytes[:]...)
	hash = ethCrypto.Keccak256(data)
	return hash, data
}

type addAssociatedAddressDelegatedParams struct {
	ApprovingAddress common.Address  `json:"approving_address"`
	AddressToAdd     common.Address  `json:"address_to_add"`
	V                []hexutil.Bytes `json:"v" validate:"len=2,itemlen=1"`  //  itemlen=1
	R                []hexutil.Bytes `json:"r" validate:"len=2,itemlen=32"` //itemlen=32
	S                []hexutil.Bytes `json:"s" validate:"len=2,itemlen=32"` //itemlen=32
	Timestamp        [2]*big.Int     `json:"timestamp" validate:"len=2"`    //itemlen=32
}

func (p *addAssociatedAddressDelegatedParams) Keccak256(identityRegistryAddr *common.Address, ein *big.Int) (hash0 []byte, data0 []byte, hash1 []byte, data1 []byte) {
	irAddrBytes := identityRegistryAddr.Bytes()
	einBytes := bigIntToByte32(ein)

	timestamp0Bytes := bigIntToByte32(p.Timestamp[0])
	timestamp1Bytes := bigIntToByte32(p.Timestamp[1])
	var data []byte
	data = append(data, headerBytes...)
	data = append(data, irAddrBytes...)

	data0 = append(data, []byte("I authorize adding this address to my Identity.")...)
	data0 = append(data0, einBytes[:]...)
	data0 = append(data0, p.AddressToAdd.Bytes()...)
	data0 = append(data0, timestamp0Bytes[:]...)

	data1 = append(data, []byte("I authorize being added to this Identity.")...)
	data1 = append(data1, einBytes[:]...)
	data1 = append(data1, p.AddressToAdd.Bytes()...)
	data1 = append(data1, timestamp1Bytes[:]...)

	hash0 = ethCrypto.Keccak256(data0)
	hash1 = ethCrypto.Keccak256(data1)
	return hash0, data0, hash1, data1
}

type removeAssociatedAddressDelegatedParams struct {
	AddressToRemove common.Address `json:"address_to_remove"`
	V               hexutil.Bytes  `json:"v" validate:"len=1"`
	R               hexutil.Bytes  `json:"r" validate:"len=32"`
	S               hexutil.Bytes  `json:"s" validate:"len=32"`
	Timestamp       *big.Int       `json:"timestamp"`
}

func (p *removeAssociatedAddressDelegatedParams) Keccak256(identityRegistryAddr *common.Address, ein *big.Int) (hash []byte, data []byte) {
	irAddrBytes := identityRegistryAddr.Bytes()
	einBytes := bigIntToByte32(ein)
	timestampBytes := bigIntToByte32(p.Timestamp)

	data = append(data, headerBytes...)
	data = append(data, irAddrBytes...)
	data = append(data, []byte("I authorize removing this address from my Identity.")...)
	data = append(data, einBytes[:]...)
	data = append(data, p.AddressToRemove.Bytes()...)
	data = append(data, timestampBytes[:]...)
	hash = ethCrypto.Keccak256(data)
	return hash, data
}

type addKeyDelegatedParams struct {
	ResolverAddress   common.Address `json:"resolver_address"`
	AssociatedAddress common.Address `json:"associated_address"`
	Key               common.Address `json:"key"`
	Symbol            string         `json:"symbol"`
	V                 hexutil.Bytes  `json:"v" validate:"len=1"`
	R                 hexutil.Bytes  `json:"r" validate:"len=32"`
	S                 hexutil.Bytes  `json:"s" validate:"len=32"`
	Timestamp         *big.Int       `json:"timestamp"`
}

func (p *addKeyDelegatedParams) Keccak256() (hash []byte, data []byte) {
	timestampBytes := bigIntToByte32(p.Timestamp)
	data = append(data, headerBytes...)
	data = append(data, p.ResolverAddress.Bytes()...)
	data = append(data, []byte("I authorize the addition of a service key on my behalf.")...)
	data = append(data, p.Key.Bytes()...)
	data = append(data, []byte(p.Symbol)...)
	data = append(data, timestampBytes[:]...)
	hash = ethCrypto.Keccak256(data)
	return hash, data
}

type removeKeyDelegatedParams struct {
	ResolverAddress   common.Address `json:"resolver_address"`
	AssociatedAddress common.Address `json:"associated_address"`
	Key               common.Address `json:"key"`
	V                 hexutil.Bytes  `json:"v" validate:"len=1"`
	R                 hexutil.Bytes  `json:"r" validate:"len=32"`
	S                 hexutil.Bytes  `json:"s" validate:"len=32"`
	Timestamp         *big.Int       `json:"timestamp"`
}

func (p *removeKeyDelegatedParams) Keccak256() (hash []byte, data []byte) {
	timestampBytes := bigIntToByte32(p.Timestamp)
	data = append(data, headerBytes...)
	data = append(data, p.ResolverAddress.Bytes()...)
	data = append(data, []byte("I authorize the removal of a service key on my behalf.")...)
	data = append(data, p.Key.Bytes()...)
	data = append(data, timestampBytes[:]...)
	hash = ethCrypto.Keccak256(data)
	return hash, data
}

type removeKeysDelegatedParams struct {
	ResolverAddress   common.Address `json:"resolver_address"`
	AssociatedAddress common.Address `json:"associated_address"`
	V                 hexutil.Bytes  `json:"v" validate:"len=1"`
	R                 hexutil.Bytes  `json:"r" validate:"len=32"`
	S                 hexutil.Bytes  `json:"s" validate:"len=32"`
	Timestamp         *big.Int       `json:"timestamp"`
}

func (p *removeKeysDelegatedParams) Keccak256() (hash []byte, data []byte) {
	timestampBytes := bigIntToByte32(p.Timestamp)
	data = append(data, headerBytes...)
	data = append(data, p.ResolverAddress.Bytes()...)
	data = append(data, []byte("I authorize the removal of all service keys on my behalf.")...)
	data = append(data, timestampBytes[:]...)
	hash = ethCrypto.Keccak256(data)
	return hash, data
}

type addPublicKeyDelegatedParams struct {
	ResolverAddress   common.Address `json:"resolver_address"`
	AssociatedAddress common.Address `json:"associated_address"`
	PublicKey         hexutil.Bytes  `json:"public_key"`
	V                 hexutil.Bytes  `json:"v" validate:"len=1"`
	R                 hexutil.Bytes  `json:"r" validate:"len=32"`
	S                 hexutil.Bytes  `json:"s" validate:"len=32"`
	Timestamp         *big.Int       `json:"timestamp"`
}

func (p *addPublicKeyDelegatedParams) Keccak256() (hash []byte, data []byte) {
	timestampBytes := bigIntToByte32(p.Timestamp)
	data = append(data, headerBytes...)
	data = append(data, p.ResolverAddress.Bytes()...)
	data = append(data, []byte("I authorize the addition of a public key on my behalf.")...)
	data = append(data, p.AssociatedAddress.Bytes()...)
	data = append(data, p.PublicKey...)
	data = append(data, timestampBytes[:]...)
	hash = ethCrypto.Keccak256(data)
	return hash, data
}

type removePublicKeyDelegatedParams struct {
	ResolverAddress   common.Address `json:"resolver_address"`
	AssociatedAddress common.Address `json:"associated_address"`
	V                 hexutil.Bytes  `json:"v" validate:"len=1"`
	R                 hexutil.Bytes  `json:"r" validate:"len=32"`
	S                 hexutil.Bytes  `json:"s" validate:"len=32"`
	Timestamp         *big.Int       `json:"timestamp"`
}

func (p *removePublicKeyDelegatedParams) Keccak256() (hash []byte, data []byte) {
	timestampBytes := bigIntToByte32(p.Timestamp)
	data = append(data, headerBytes...)
	data = append(data, p.ResolverAddress.Bytes()...)
	data = append(data, []byte("I authorize the removal of a public key on my behalf.")...)
	data = append(data, p.AssociatedAddress.Bytes()...)
	data = append(data, timestampBytes[:]...)
	hash = ethCrypto.Keccak256(data)
	return hash, data
}

func init() {
	validator.SetValidationFunc("itemlen", checkBytesLength)
}
func fillParam(p interface{}, obj interface{}) Error {
	//	fmt.Println("fillParam-obj : ", obj)
	jsonBytes, err1 := encodingJson.Marshal(obj)
	//	fmt.Println("fillParam-jsonBytes : ", string(jsonBytes))
	err1 = encodingJson.Unmarshal(jsonBytes, p)

	if err1 != nil {
		log.Errorf("fillParam err : %v", err1)
		return &invalidMessageError{err1.Error()}
	}

	if errs := validator.Validate(p); errs != nil {
		// values not valid, deal with errors here

		return &invalidParamsError{errs.Error()}
	}
	// err := p.validate()
	// if err != nil {
	// 	return err
	// }
	return nil

}

func getParameter(method string, params []interface{}) (interface{}, Error) {
	if len(params) != 1 {
		return nil, &invalidParamsError{"Invalid params."}
	}
	obj := params[0]
	//log.Debug("param", obj)
	switch method {

	case "create_identity":
		var reqParam identityCreateParams
		err := fillParam(&reqParam, obj)
		if err != nil {
			return nil, err
		}
		return reqParam, nil

	case "add_associated_address_delegated":
		var reqParam addAssociatedAddressDelegatedParams
		err := fillParam(&reqParam, obj)
		if err != nil {
			return nil, err
		}
		return reqParam, nil
	case "remove_associated_address_delegated":
		var reqParam removeAssociatedAddressDelegatedParams
		err := fillParam(&reqParam, obj)
		if err != nil {
			return nil, err
		}
		return reqParam, nil
	case "add_key_delegated":
		var reqParam addKeyDelegatedParams
		err := fillParam(&reqParam, obj)
		if err != nil {
			return nil, err
		}
		return reqParam, nil

	case "remove_key_delegated":
		var reqParam removeKeyDelegatedParams
		err := fillParam(&reqParam, obj)
		if err != nil {
			return nil, err
		}
		return reqParam, nil

	case "remove_keys_delegated":
		var reqParam removeKeysDelegatedParams
		err := fillParam(&reqParam, obj)
		if err != nil {
			return nil, err
		}
		return reqParam, nil

	case "add_public_key_delegated":
		var reqParam addPublicKeyDelegatedParams
		err := fillParam(&reqParam, obj)
		if err != nil {
			return nil, err
		}
		return reqParam, nil

	case "remove_public_key_delegated":
		var reqParam removePublicKeyDelegatedParams
		err := fillParam(&reqParam, obj)
		if err != nil {
			return nil, err
		}
		return reqParam, nil

	default:
		return nil, &methodNotFoundError{method}

	}

}

func checkBytesLength(v interface{}, param string) error {
	//fmt.Printf("parameter[v] : %v\n", v)
	//fmt.Printf("parameter[param] : %v\n", param)
	st := reflect.ValueOf(v)
	//fmt.Printf("parameter[st] : %v\n", st)
	//fmt.Printf("parameter[st.Kind] : %v\n", st.Kind())
	if st.Kind() != reflect.Slice {
		return fmt.Errorf("validation error: invalid parameter")
	}
	p, err := asInt(param)
	//fmt.Printf("parameter[p] : %v\n", p)
	if err != nil {
		return validator.ErrBadParameter
	}

	var typeBytesList []hexutil.Bytes
	if st.Type() == reflect.TypeOf(typeBytesList) {
		//	fmt.Printf("Type is %v\n", st.Type())

		list := v.([]hexutil.Bytes)

		//	fmt.Printf("parameter[list] : %v\n", list)
		for _, hash := range list {
			//		fmt.Printf("parameter[hash] : %v\n", hash)
			if int64(len(hash)) != p {
				return fmt.Errorf("validation error: invalid element")
			}
		}
	} else {
		//	fmt.Printf("Unknown Type: : %v", st.Type())
		return fmt.Errorf("validation error :Unknown Type (%v)", st.Type())
	}
	return nil
}

// asInt returns the parameter as a int64
// or panics if it can't convert
func asInt(param string) (int64, error) {
	i, err := strconv.ParseInt(param, 0, 32)
	if err != nil {
		return 0, validator.ErrBadParameter
	}
	return i, nil
}
