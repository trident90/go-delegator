package predefined

import (
	encodingJson "encoding/json"
	"fmt"
	"strings"

	"bitbucket.org/coinplugin/geth/common"
	"bitbucket.org/coinplugin/geth/common/hexutil"
	"bitbucket.org/coinplugin/proxy/crypto"
	"bitbucket.org/coinplugin/proxy/json"
)

const (
	addressSize         = 20
	hashSize            = 32
	signatureSize       = 65
	fileHashSize        = 46
	timestampSize       = 52
	minimumUserHashSize = 3
)

type metaIDRegisterParams struct {
	Address   common.Address  `json:"address"`
	MetaID    hexutil.Bytes   `json:"meta_id"`
	UserHashs []hexutil.Bytes `json:"user_hash_list"`
	Signature hexutil.Bytes   `json:"signature"` // Sign(MetaId)
}

type metaIDUpdateParams struct {
	Address   common.Address  `json:"address"`
	OldMetaID hexutil.Bytes   `json:"meta_id_old"`
	NewMetaID hexutil.Bytes   `json:"meta_id_new"`
	UserHashs []hexutil.Bytes `json:"user_hash_list"`
	Signature hexutil.Bytes   `json:"signature"` // Sign(NewMetaId)
}

type metaIDBackupParams struct {
	Address   common.Address `json:"address"`
	MetaID    hexutil.Bytes  `json:"meta_id"`
	encData   hexutil.Bytes  `json:"enc_data"`
	Signature hexutil.Bytes  `json:"signature"` // Sign(MetaId)
}

type metaIDGetUserDataParams struct {
	NewAddress common.Address `json:"address_new"`
	FileID     string         `json:"file_id"`
	Signature  hexutil.Bytes  `json:"signature"` // Sign(FileId)
}

type metaIDRestoreParams struct {
	NewAddress common.Address  `json:"address_new"`
	OldAddress common.Address  `json:"address_old"`
	NewMetaID  hexutil.Bytes   `json:"meta_id_new"`
	OldMetaID  hexutil.Bytes   `json:"meta_id_old"`
	UserHashs  []hexutil.Bytes `json:"user_hash_list"`
	Signature  hexutil.Bytes   `json:"signature"` //sign(NewMetaId)
}

type metaIDRevokeParams struct {
	MetaID    hexutil.Bytes `json:"meta_id"`
	Timestamp hexutil.Bytes `json:"timestamp"`
	Signature hexutil.Bytes `json:"signature"` //sign(MetaId|timestamp)
}

type metaParam interface {
	validate() Error
}

//validate is function of checking if Parameter is valid or invalid. if invalid, return invalidParamsError
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
	} else {
		for _, hash := range p.UserHashs {
			if len(hash) != hashSize {
				return &invalidParamsError{fmt.Sprintf("Invalid user_hash_list parameter element  %v", len(p.MetaID))}
			}
		}
	}
	return nil
}

func (p *metaIDUpdateParams) validate() (err Error) {
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
	} else {
		for _, hash := range p.UserHashs {
			if len(hash) != hashSize {
				return &invalidParamsError{"Invalid user_hash_list parameter element."}
			}
		}
	}
	return
}

func (p *metaIDBackupParams) validate() (err Error) {

	if len(p.Address) != addressSize {
		return &invalidParamsError{"Invalid address parameter."}
	}
	if len(p.MetaID) != hashSize {
		return &invalidParamsError{"Invalid meta_id parameter."}
	}
	if len(p.encData) > minimumUserHashSize {
		return &invalidParamsError{"Invalid enc_data parameter."}
	}
	if len(p.Signature) != signatureSize {
		return &invalidParamsError{"Invalid signature parameter."}
	}
	return
}
func (p *metaIDGetUserDataParams) validate() (err Error) {
	if len(p.NewAddress) != addressSize {
		return &invalidParamsError{"Invalid address_new parameter."}
	}

	if len(p.FileID) < fileHashSize {
		return &invalidParamsError{"Invalid meta_id parameter."}
	}

	if len(p.Signature) != signatureSize {
		return &invalidParamsError{"Invalid signature parameter."}
	}
	return
}

func (p *metaIDRestoreParams) validate() (err Error) {
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
	} else {
		for _, hash := range p.UserHashs {
			if len(hash) != hashSize {
				return &invalidParamsError{"Invalid user_hash_list parameter element."}
			}
		}
	}
	return
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

func fillParam(p metaParam, obj interface{}) Error {
	jsonBytes, err1 := encodingJson.Marshal(obj)

	err1 = encodingJson.Unmarshal(jsonBytes, p)

	if err1 != nil {
		fmt.Println("err : ", err1)
		return &invalidMessageError{err1.Error()}
	}

	err := p.validate()
	if err != nil {
		return err
	}
	return nil

}

// func (p *metaIDUpdateParams) fillParam(obj interface{}) Error {
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

func getParameter(method string, obj interface{}) (interface{}, Error) {

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

	default:
		return nil, &methodNotFoundError{method}

	}

}

func register_meta_id(req json.RpcRequest) (resp json.RpcResponse, err error) {
	//var reqParam interface{}
	resp.Id = req.Id
	resp.Jsonrpc = req.Jsonrpc

	//1. Check paramter format
	tmpParams, errObj := getParameter(req.Method, req.Params[0])
	if errObj != nil {

		resp.Error.Code = errObj.ErrorCode()
		resp.Error.Message = errObj.Error()
		err = fmt.Errorf(errObj.Error())
		return
	}
	reqParam := tmpParams.(metaIDRegisterParams)
	fmt.Println("parameter[Address] : ", reqParam.Address)
	fmt.Println("parameter[MetaId] : ", reqParam.MetaID)
	fmt.Println("parameter[Sig] : ", reqParam.Signature)
	fmt.Println("parameter[UserHashs] : ", reqParam.UserHashs)

	//2. check Signature
	//  if  address == ecrecover()  then Pass else Error
	signedAddress, err := crypto.EcRecover(reqParam.MetaID.String(), reqParam.Signature.String())
	if err != nil {
		fmt.Println("Error :", err)
		errObj := &invalidSignatureError{"Failed to verify signature"}
		resp.Error.Code = errObj.ErrorCode()
		resp.Error.Message = errObj.Error()
		err = fmt.Errorf(errObj.Error())
		return
	}

	fmt.Println("Address :", signedAddress)

	if signedAddress != strings.ToLower(reqParam.Address.Hex()) {
		fmt.Printf("Error  :Sign Error %v / %v \n", reqParam.Address.Hex(), signedAddress)
		errObj := &invalidSignatureError{"Failed to verify signature"}
		resp.Error.Code = errObj.ErrorCode()
		resp.Error.Message = errObj.Error()
		err = fmt.Errorf(errObj.Error())
		return
	}

	//3. Check Meta ID  by make merkle tree root.
	//   if metaId == root  then Pass  else  Error

	//4. Reqest Get UserAddress for Meta ID
	//  if GetAddress == nil  then Pass  else Error

	//5. Send Transaction for Calling SC Function
	//  return txid
	return
}

func update_meta_id(req json.RpcRequest) (resp json.RpcResponse, err error) {
	return
}

func backup_user_data(req json.RpcRequest) (resp json.RpcResponse, err error) {
	return
}

func get_user_data(req json.RpcRequest) (resp json.RpcResponse, err error) {
	return
}

func restore_user_data(req json.RpcRequest) (resp json.RpcResponse, err error) {
	return
}

func revoke_meta_id(req json.RpcRequest) (resp json.RpcResponse, err error) {
	return
}
