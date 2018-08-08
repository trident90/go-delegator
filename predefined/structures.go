package predefined

import (
	encodingJson "encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"bitbucket.org/coinplugin/proxy/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

type metaIDRegisterParams struct {
	Address   common.Address  `json:"address" validate:"len=20"`
	MetaID    hexutil.Bytes   `json:"meta_id" validate:"len=32"`
	UserHashs []hexutil.Bytes `json:"user_hash_list" validate:"min=3,itemlen=32"`
	Signature hexutil.Bytes   `json:"signature" validate:"len=65"` // Sign(MetaId)
}

type metaIDUpdateParams struct {
	Address   common.Address  `json:"address" validate:"len=20"`
	OldMetaID hexutil.Bytes   `json:"meta_id_old" validate:"len=32"`
	NewMetaID hexutil.Bytes   `json:"meta_id_new" validate:"len=32"`
	UserHashs []hexutil.Bytes `json:"user_hash_list" validate:"min=3,itemlen=32"`
	Signature hexutil.Bytes   `json:"signature" validate:"len=65"` // Sign(NewMetaId)
}

type metaIDBackupParams struct {
	Address   common.Address `json:"address" validate:"len=20"`
	MetaID    hexutil.Bytes  `json:"meta_id" validate:"len=32"`
	EncData   hexutil.Bytes  `json:"enc_data" validate:"min=66"`
	Signature hexutil.Bytes  `json:"signature" validate:"len=65"` // Sign(MetaId)
}

type metaIDGetUserDataParams struct {
	//TODO: 테스트로 인해 address/signature 필드 제거  추후 결정 필요
	//	NewAddress common.Address `json:"address_new"`
	FileID string `json:"file_id" validate:"len=46"`
	//	Signature  hexutil.Bytes  `json:"signature" validate:"len=65"` // Sign(FileId)
}

type metaIDRestoreParams struct {
	NewAddress common.Address  `json:"address_new" validate:"len=20"`
	OldAddress common.Address  `json:"address_old" validate:"len=20"`
	NewMetaID  hexutil.Bytes   `json:"meta_id_new" validate:"len=32"`
	OldMetaID  hexutil.Bytes   `json:"meta_id_old" validate:"len=32"`
	UserHashs  []hexutil.Bytes `json:"user_hash_list" validate:"min=3,itemlen=32"`
	Signature  hexutil.Bytes   `json:"signature" validate:"len=65"` //sign(NewMetaId)
}

type metaIDRevokeParams struct {
	MetaID    hexutil.Bytes `json:"meta_id" validate:"len=32"`
	Timestamp hexutil.Bytes `json:"timestamp" validate:"len=14"`
	Signature hexutil.Bytes `json:"signature" validate:"len=65"` //sign(MetaId|timestamp)
}
type metaIDGetAddressParams struct {
	MetaID hexutil.Bytes `json:"meta_id" validate:"len=32"`
}

func init() {
	validator.SetValidationFunc("itemlen", checkBytesLength)
}
func fillParam(p interface{}, obj interface{}) Error {
	jsonBytes, err1 := encodingJson.Marshal(obj)

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

// fillParam is function of converting map type  to struct type and check Valid parameter
// func (p *metaIDRegisterParams) fillParam(obj interface{}) Error {
// 	jsonBytes, err1 := encodingJson.Marshal(obj)

// 	err1 = encodingJson.Unmarshal(jsonBytes, p)

// 	if err1 != nil {
// 		fmt.Println("err : ", err1)
// 		return &invalidMessageError{err1.Error()}
// 	}

// 	err := p.validate()
// 	if err != nil {
// 		return err
// 	}
// 	return nil

// }

func getParameter(method string, params []interface{}) (interface{}, Error) {
	if len(params) != 1 {
		return nil, &invalidParamsError{"Invalid params."}
	}
	obj := params[0]
	switch method {
	case "register_meta_id":
		var reqParam metaIDRegisterParams
		err := fillParam(&reqParam, obj)
		if err != nil {
			return nil, err
		}

		return reqParam, nil

	case "update_meta_id":
		var reqParam metaIDUpdateParams
		err := fillParam(&reqParam, obj)
		if err != nil {
			return nil, err
		}
		return reqParam, nil

	case "backup_user_data":
		var reqParam metaIDBackupParams
		err := fillParam(&reqParam, obj)
		if err != nil {
			return nil, err
		}
		return reqParam, nil

	case "get_user_data":
		var reqParam metaIDGetUserDataParams
		err := fillParam(&reqParam, obj)
		if err != nil {
			return nil, err
		}
		return reqParam, nil

	case "restore_user_data":
		var reqParam metaIDRestoreParams
		err := fillParam(&reqParam, obj)
		if err != nil {
			return nil, err
		}
		return reqParam, nil

	case "revoke_meta_id":
		var reqParam metaIDRevokeParams
		err := fillParam(&reqParam, obj)
		if err != nil {
			return nil, err
		}
		return reqParam, nil

	case "get_address":
		var reqParam metaIDGetAddressParams
		err := fillParam(&reqParam, obj)
		if err != nil {
			return nil, err
		}
		return reqParam, nil
	default:
		return nil, &methodNotFoundError{method}

	}

}

/*
type metaParam interface {
	validate() Error
}

//validate is function of checking if Parameter is valid or invalid. if invalid, return invalidParamsError
//TODO: validate lib 적용 고려 (supported by 승곤)
func (p *metaIDRegisterParams) validate() Error {

	if len(p.Address) != addressSize {
		return &invalidParamsError{fmt.Sprintf("Invalid address parameter %v", len(p.Address))}
	}
	if len(p.MetaID) != hashSize {
		return &invalidParamsError{fmt.Sprintf("Invalid meta_id parameter %v", len(p.MetaID))}
	}
	if len(p.Signature) != signatureSize {
		return &invalidParamsError{fmt.Sprintf("Invalid signature parameter %v", len(p.Signature))}
	}
	if len(p.UserHashs) < minimumUserHashSize {
		return &invalidParamsError{fmt.Sprintf("Invalid user_hash_list parameter  %v", len(p.UserHashs))}
	}
	for _, hash := range p.UserHashs {
		if len(hash) != hashSize {
			return &invalidParamsError{fmt.Sprintf("Invalid user_hash_list parameter element  %v", len(p.MetaID))}
		}
	}

	return nil
}

func (p *metaIDUpdateParams) validate() Error {
	if len(p.Address) != addressSize {
		return &invalidParamsError{"Invalid address parameter."}
	}
	if len(p.OldMetaID) != hashSize {
		return &invalidParamsError{"Invalid meta_id_old parameter."}
	}
	if len(p.NewMetaID) != hashSize {
		return &invalidParamsError{"Invalid meta_id_new parameter."}
	}
	if len(p.Signature) != signatureSize {
		return &invalidParamsError{"Invalid signature parameter."}
	}
	if len(p.UserHashs) < minimumUserHashSize {
		return &invalidParamsError{"Invalid user_hash_list parameter."}
	}
	for _, hash := range p.UserHashs {
		if len(hash) != hashSize {
			return &invalidParamsError{"Invalid user_hash_list parameter element."}
		}
	}

	return nil
}

func (p *metaIDBackupParams) validate() Error {

	if len(p.Address) != addressSize {
		return &invalidParamsError{"Invalid address parameter."}
	}
	if len(p.MetaID) != hashSize {
		return &invalidParamsError{"Invalid meta_id parameter."}
	}
	if len(p.EncData) < minimumUserHashSize {
		return &invalidParamsError{"Invalid enc_data parameter."}
	}
	if len(p.Signature) != signatureSize {
		return &invalidParamsError{"Invalid signature parameter."}
	}
	return nil
}
func (p *metaIDGetUserDataParams) validate() Error {
	//TODO: 테스트로 인해 address/signature 필드 제거  추후 결정 필요
	// if len(p.NewAddress) != addressSize {
	// 	return &invalidParamsError{"Invalid address_new parameter."}
	// }

	if len(p.FileID) != fileHashSize {
		return &invalidParamsError{"Invalid meta_id parameter."}
	}

	// if len(p.Signature) != signatureSize {
	// 	return &invalidParamsError{"Invalid signature parameter."}
	// }
	return nil
}

func (p *metaIDRestoreParams) validate() Error {
	if len(p.NewAddress) != addressSize {
		return &invalidParamsError{"Invalid address_new parameter."}
	}
	if len(p.OldAddress) != addressSize {
		return &invalidParamsError{"Invalid address_old parameter."}
	}
	if len(p.NewMetaID) != hashSize {
		return &invalidParamsError{"Invalid meta_id_new parameter."}
	}
	if len(p.OldMetaID) != hashSize {
		return &invalidParamsError{"Invalid meta_id_old parameter."}
	}
	if len(p.Signature) != signatureSize {
		return &invalidParamsError{"Invalid signature parameter."}
	}

	if len(p.UserHashs) < minimumUserHashSize {
		return &invalidParamsError{"Invalid user_hash_list parameter."}
	}
	for _, hash := range p.UserHashs {
		if len(hash) != hashSize {
			return &invalidParamsError{"Invalid user_hash_list parameter element."}
		}
	}

	return nil
}

func (p *metaIDRevokeParams) validate() Error {

	if len(p.MetaID) != hashSize {
		return &invalidParamsError{"Invalid meta_id parameter."}
	}
	if len(p.Timestamp) != timestampSize {
		return &invalidParamsError{"Invalid timesamp parameter."}
	}
	if len(p.Signature) != signatureSize {
		return &invalidParamsError{"Invalid signature parameter."}
	}

	return nil
}
*/
func checkBytesLength(v interface{}, param string) error {

	st := reflect.ValueOf(v)
	if st.Kind() != reflect.Slice {
		return fmt.Errorf("user_hash_list: invalid parameter")
	}
	p, err := asInt(param)
	if err != nil {
		return validator.ErrBadParameter
	}

	list := v.([]hexutil.Bytes)

	for _, hash := range list {
		if int64(len(hash)) != p {
			return fmt.Errorf("user_hash_list: invalid element")
		}
	}
	return nil
}

// asInt retuns the parameter as a int64
// or panics if it can't convert
func asInt(param string) (int64, error) {
	i, err := strconv.ParseInt(param, 0, 32)
	if err != nil {
		return 0, validator.ErrBadParameter
	}
	return i, nil
}
