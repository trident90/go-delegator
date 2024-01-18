package metaservice

import (
	encodingJson "encoding/json"
	"fmt"
	"math/big"
	"reflect"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

type metaIDCreateParams struct {
	Address   common.Address `json:"address" validate:"len=20"`
	Signature hexutil.Bytes  `json:"signature" validate:"len=65"` // Sign(Address)
	Recaptcha string         `json:"recaptcha"`
}

type metaIDExecuteParams struct {
	MetaID    common.Address `json:"meta_id" validate:"len=20"`
	From      common.Address `json:"from" validate:"len=20"`
	To        common.Address `json:"to" validate:"len=20"`
	Value     *big.Int       `json:"value"`
	Data      hexutil.Bytes  `json:"data"`
	Nonce     *big.Int       `json:"nonce"`
	Signature hexutil.Bytes  `json:"signature" validate:"len=65"`
}

type metaIDApproveParams struct {
	MetaID common.Address `json:"meta_id" validate:"len=20"`
	From   common.Address `json:"from" validate:"len=20"`
	//	To        common.Address `json:"to" validate:"len=20"`
	Id        hexutil.Bytes `json:"id" validate:"len=32"`
	Approve   bool          `json:"approve"`
	Nonce     *big.Int      `json:"nonce"`
	Signature hexutil.Bytes `json:"signature" validate:"len=65"`
}

type metaIDBackupParams struct {
	Address   common.Address `json:"address" validate:"len=20"`
	MetaID    hexutil.Bytes  `json:"meta_id" validate:"len=20"`
	EncData   hexutil.Bytes  `json:"enc_data" validate:"min=66"`
	Signature hexutil.Bytes  `json:"signature" validate:"len=65"` // Sign(MetaId)
}

type metaIDGetUserDataParams struct {
	//TODO: 테스트로 인해 address/signature 필드 제거  추후 결정 필요
	//	NewAddress common.Address `json:"address_new"`
	FileID string `json:"file_id" validate:"len=46"`
	//	Signature  hexutil.Bytes  `json:"signature" validate:"len=65"` // Sign(FileId)
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

func getParameter(method string, params []interface{}) (interface{}, Error) {
	if len(params) != 1 {
		return nil, &invalidParamsError{"Invalid params."}
	}
	obj := params[0]
	log.Debug("param", obj)
	switch method {

	case "create_meta_id":
		var reqParam metaIDCreateParams
		err := fillParam(&reqParam, obj)
		if err != nil {
			return nil, err
		}
		return reqParam, nil

	case "delegated_execute":
		var reqParam metaIDExecuteParams
		err := fillParam(&reqParam, obj)
		if err != nil {
			return nil, err
		}
		return reqParam, nil

	case "delegated_approve":
		var reqParam metaIDApproveParams
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

	case "get_identity_manager_address", "get_topic_registry_address", "get_attestation_agency_registry_address", "get_achievement_manager_address", "get_achievement_address":
		return nil, nil

	default:
		return nil, &methodNotFoundError{method}

	}

}

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

// asInt returns the parameter as a int64
// or panics if it can't convert
func asInt(param string) (int64, error) {
	i, err := strconv.ParseInt(param, 0, 32)
	if err != nil {
		return 0, validator.ErrBadParameter
	}
	return i, nil
}
