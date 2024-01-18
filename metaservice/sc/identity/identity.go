// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package identity

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IdentityABI is the input ABI used to generate the binding from.
const IdentityABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERC165ID\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"REGISTRY_TOPIC\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MANAGEMENT_KEY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getKey\",\"outputs\":[{\"name\":\"purposes\",\"type\":\"uint256[]\"},{\"name\":\"keyType\",\"type\":\"uint256\"},{\"name\":\"key\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"issuer\",\"type\":\"address\"},{\"name\":\"topic\",\"type\":\"uint256\"}],\"name\":\"getClaimId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumClaims\",\"outputs\":[{\"name\":\"num\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"subject\",\"type\":\"address\"},{\"name\":\"topic\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"claimToSign\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"},{\"name\":\"_purpose\",\"type\":\"uint256\"},{\"name\":\"_keyType\",\"type\":\"uint256\"}],\"name\":\"addKey\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_claimId\",\"type\":\"bytes32\"}],\"name\":\"isClaimExist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_SCHEME\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_topic\",\"type\":\"uint256\"}],\"name\":\"getClaimIdsByType\",\"outputs\":[{\"name\":\"claimIds\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RSA_TYPE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTransactionCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"actionThreshold\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"toSign\",\"type\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"getSignatureAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ECDSA_TYPE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"},{\"name\":\"_nonce\",\"type\":\"uint256\"},{\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"delegatedExecute\",\"outputs\":[{\"name\":\"executionId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_claimId\",\"type\":\"bytes32\"}],\"name\":\"removeClaim\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"},{\"name\":\"_purpose\",\"type\":\"uint256\"}],\"name\":\"removeKey\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_claimId\",\"type\":\"bytes32\"}],\"name\":\"refreshClaim\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"execution\",\"outputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"needsApprove\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addrToKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"METAID_TOPIC\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_func\",\"type\":\"bytes4\"},{\"name\":\"_executable\",\"type\":\"bool\"}],\"name\":\"setFunc\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_approve\",\"type\":\"bool\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ACTION_KEY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_func\",\"type\":\"bytes4\"}],\"name\":\"keyCanExecute\",\"outputs\":[{\"name\":\"executable\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERC725ID\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ECDSA_SCHEME\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_purpose\",\"type\":\"uint256\"}],\"name\":\"getKeysByPurpose\",\"outputs\":[{\"name\":\"keys\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ENCRYPTION_KEY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RESIDENCE_TOPIC\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"managementThreshold\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PROFILE_TOPIC\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CUSTOM_KEY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_topic\",\"type\":\"uint256\"},{\"name\":\"_scheme\",\"type\":\"uint256\"},{\"name\":\"issuer\",\"type\":\"address\"},{\"name\":\"_signature\",\"type\":\"bytes\"},{\"name\":\"_data\",\"type\":\"bytes\"},{\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"addClaim\",\"outputs\":[{\"name\":\"claimRequestId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"LABEL_TOPIC\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"name\":\"executionId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RESTORE_KEY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERC735ID\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"getFunctionSignature\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CLAIM_SIGNER_KEY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_claimId\",\"type\":\"bytes32\"}],\"name\":\"getClaim\",\"outputs\":[{\"name\":\"topic\",\"type\":\"uint256\"},{\"name\":\"scheme\",\"type\":\"uint256\"},{\"name\":\"issuer\",\"type\":\"address\"},{\"name\":\"signature\",\"type\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"uri\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numKeys\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"changeManagementThreshold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNonce\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"},{\"name\":\"purpose\",\"type\":\"uint256\"}],\"name\":\"keyHasPurpose\",\"outputs\":[{\"name\":\"exists\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ASSIST_KEY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_topic\",\"type\":\"uint256\"},{\"name\":\"_scheme\",\"type\":\"uint256\"},{\"name\":\"issuer\",\"type\":\"address\"},{\"name\":\"_signature\",\"type\":\"bytes\"},{\"name\":\"_data\",\"type\":\"bytes\"},{\"name\":\"_uri\",\"type\":\"string\"},{\"name\":\"_idSignature\",\"type\":\"bytes\"}],\"name\":\"addClaimByProxy\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_approve\",\"type\":\"bool\"},{\"name\":\"_nonce\",\"type\":\"uint256\"},{\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"delegatedApprove\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"changeActionThreshold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DELEGATE_KEY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RSA_SCHEME\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"destroyAndSend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"self\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"getExecutionId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numClaims\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_managementKey\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"claimRequestId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"topic\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"scheme\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"signature\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"ClaimRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"claimId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"topic\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"scheme\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"signature\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"ClaimAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"claimId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"topic\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"scheme\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"signature\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"ClaimRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"claimId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"topic\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"scheme\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"signature\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"ClaimChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"purpose\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"keyType\",\"type\":\"uint256\"}],\"name\":\"KeyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"purpose\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"keyType\",\"type\":\"uint256\"}],\"name\":\"KeyRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"executionId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ExecutionRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"executionId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"executionId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"Approved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"executionId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ExecutionFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"LogPause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"LogUnpause\",\"type\":\"event\"}]"

// IdentityBin is the compiled bytecode used for deploying new contracts.
const IdentityBin = `60806040526001600055600180556000600560006101000a81548160ff02191690831515021790555060016007553480156200003a57600080fd5b5060405160208062006fe58339810180604052810190808051906020019092919050505060006001600660006200007f6200030e640100000000026401000000009004565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600660006200010162000339640100000000026401000000009004565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600660006200018362000442640100000000026401000000009004565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060006101000a81548160ff0219169083151502179055506200020082620004dc640100000000026401000000009004565b90506200021f81600180620004ff640100000000026401000000009004565b6200023d8160026001620004ff640100000000026401000000009004565b6200025b8160036001620004ff640100000000026401000000009004565b6001600081905550600180819055506001600660006200028962000442640100000000026401000000009004565b620002a262000339640100000000026401000000009004565b187bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060006101000a81548160ff02191690831515021790555050506200063a565b60006301ffc9a77c010000000000000000000000000000000000000000000000000000000002905090565b60006353d413c57c01000000000000000000000000000000000000000000000000000000000263747442d37c01000000000000000000000000000000000000000000000000000000000263b61d27f67c010000000000000000000000000000000000000000000000000000000002631d3812407c010000000000000000000000000000000000000000000000000000000002639010f7267c01000000000000000000000000000000000000000000000000000000000263d202158d7c0100000000000000000000000000000000000000000000000000000000026312aaac707c010000000000000000000000000000000000000000000000000000000002181818181818905090565b6000634eee424a7c01000000000000000000000000000000000000000000000000000000000263b1a34e0d7c01000000000000000000000000000000000000000000000000000000000263262b54f57c01000000000000000000000000000000000000000000000000000000000263c9100bcb7c010000000000000000000000000000000000000000000000000000000002181818905090565b60008173ffffffffffffffffffffffffffffffffffffffff166001029050919050565b6200052783838360026200055f6401000000000262005de51790939291906401000000009004565b808284600019167f480000bb1edad8ca1470381cc334b1917fbd51c6531f3a623ea8e0ec7e38a6e960405160405180910390a4505050565b6000846000016000856000191660001916815260200190815260200160002090508060000183908060018154018082558091505090600182039060005260206000200160009091929091909150555060006001028160020154600019161415620005da57838160020181600019169055508181600101819055505b84600101600084815260200190815260200160002084908060018154018082558091505090600182039060005260206000200160009091929091909150906000191690555084600201600081548092919060010191905055505050505050565b61699b806200064a6000396000f3006080604052600436106102e0576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806301ffc9a7146102e257806302e7491e146103465780630440b43a146103af578063058b316c146103da57806312aaac7014610405578063190db862146104a1578063199029861461050a5780631d203be8146105355780631d381240146105e4578063237434f014610641578063251de3e91461068a578063262b54f5146106b55780632d32d442146107375780632e7700f01461076257806338f4edd41461078d5780633b8a12c8146107b85780633f4ba83a1461086f57806349991ec8146108865780634da34c2c146108b15780634eee424a146109a857806353d413c5146109f15780635c975abb14610a445780635d7bc3fc14610a735780635dccebba14610abc57806363f14c3c14610ba35780636e4c431114610c02578063710ca55014610c79578063724a4b3b14610ca4578063747442d314610d4257806375e5598c14610d93578063765b304214610dbe5780637d96fa5814610e5057806382d0944614610eb95780638456cb5914610ee45780639010f72614610efb5780639e140cc814610f7d578063a550f0c714610fa8578063aa0a514214610fd3578063ae62838614610ffe578063affed0e014611029578063b132734e14611054578063b1a34e0d1461107f578063b1e9f64c146111bc578063b61d27f6146111e7578063b9133d631461128e578063bf2f20ad146112b9578063c32b351814611322578063c6702187146113dd578063c9100bcb14611408578063c9d24ecc146115cb578063ccfe5868146115f6578063d087d28814611623578063d202158d1461164e578063dbfa74b7146116a1578063e0610ba3146116cc578063e804c01b14611853578063e99896b8146118f4578063ead09fab14611921578063f22d08a61461194c578063f5074f4114611977578063f7444137146119ba578063fc0fc84914611a8b575b005b3480156102ee57600080fd5b5061032c60048036038101908080357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19169060200190929190505050611ab6565b604051808215151515815260200191505060405180910390f35b34801561035257600080fd5b5061035b611b1e565b60405180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b3480156103bb57600080fd5b506103c4611b49565b6040518082815260200191505060405180910390f35b3480156103e657600080fd5b506103ef611b4e565b6040518082815260200191505060405180910390f35b34801561041157600080fd5b506104346004803603810190808035600019169060200190929190505050611b53565b60405180806020018481526020018360001916600019168152602001828103825285818151815260200191508051906020019060200280838360005b8381101561048b578082015181840152602081019050610470565b5050505090500194505050505060405180910390f35b3480156104ad57600080fd5b506104ec600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611c1d565b60405180826000191660001916815260200191505060405180910390f35b34801561051657600080fd5b5061051f611cec565b6040518082815260200191505060405180910390f35b34801561054157600080fd5b506105c6600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611cf6565b60405180826000191660001916815260200191505060405180910390f35b3480156105f057600080fd5b5061062760048036038101908080356000191690602001909291908035906020019092919080359060200190929190505050611e1b565b604051808215151515815260200191505060405180910390f35b34801561064d57600080fd5b506106706004803603810190808035600019169060200190929190505050611e87565b604051808215151515815260200191505060405180910390f35b34801561069657600080fd5b5061069f611efe565b6040518082815260200191505060405180910390f35b3480156106c157600080fd5b506106e060048036038101908080359060200190929190505050611f03565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015610723578082015181840152602081019050610708565b505050509050019250505060405180910390f35b34801561074357600080fd5b5061074c611f72565b6040518082815260200191505060405180910390f35b34801561076e57600080fd5b50610777611f77565b6040518082815260200191505060405180910390f35b34801561079957600080fd5b506107a2611f81565b6040518082815260200191505060405180910390f35b3480156107c457600080fd5b5061082d6004803603810190808035600019169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611f87565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561087b57600080fd5b506108846120bb565b005b34801561089257600080fd5b5061089b612132565b6040518082815260200191505060405180910390f35b3480156108bd57600080fd5b50610992600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929080359060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050612137565b6040518082815260200191505060405180910390f35b3480156109b457600080fd5b506109d76004803603810190808035600019169060200190929190505050612314565b604051808215151515815260200191505060405180910390f35b3480156109fd57600080fd5b50610a2a600480360381019080803560001916906020019092919080359060200190929190505050612aef565b604051808215151515815260200191505060405180910390f35b348015610a5057600080fd5b50610a59612bcb565b604051808215151515815260200191505060405180910390f35b348015610a7f57600080fd5b50610aa26004803603810190808035600019169060200190929190505050612bde565b604051808215151515815260200191505060405180910390f35b348015610ac857600080fd5b50610ae7600480360381019080803590602001909291905050506130c8565b604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200184815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b83811015610b65578082015181840152602081019050610b4a565b50505050905090810190601f168015610b925780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b348015610baf57600080fd5b50610be4600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506131b0565b60405180826000191660001916815260200191505060405180910390f35b348015610c0e57600080fd5b50610c3760048036038101908080359060200190929190803590602001909291905050506131d3565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610c8557600080fd5b50610c8e613220565b6040518082815260200191505060405180910390f35b348015610cb057600080fd5b50610d286004803603810190808035600019169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19169060200190929190803515159060200190929190505050613225565b604051808215151515815260200191505060405180910390f35b348015610d4e57600080fd5b50610d796004803603810190808035906020019092919080351515906020019092919050505061329c565b604051808215151515815260200191505060405180910390f35b348015610d9f57600080fd5b50610da86132cd565b6040518082815260200191505060405180910390f35b348015610dca57600080fd5b50610e366004803603810190808035600019169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690602001909291905050506132d2565b604051808215151515815260200191505060405180910390f35b348015610e5c57600080fd5b50610e65613398565b60405180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b348015610ec557600080fd5b50610ece6134a1565b6040518082815260200191505060405180910390f35b348015610ef057600080fd5b50610ef96134a6565b005b348015610f0757600080fd5b50610f266004803603810190808035906020019092919050505061351e565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015610f69578082015181840152602081019050610f4e565b505050509050019250505060405180910390f35b348015610f8957600080fd5b50610f92613590565b6040518082815260200191505060405180910390f35b348015610fb457600080fd5b50610fbd613595565b6040518082815260200191505060405180910390f35b348015610fdf57600080fd5b50610fe861359a565b6040518082815260200191505060405180910390f35b34801561100a57600080fd5b506110136135a0565b6040518082815260200191505060405180910390f35b34801561103557600080fd5b5061103e6135a5565b6040518082815260200191505060405180910390f35b34801561106057600080fd5b506110696135ab565b6040518082815260200191505060405180910390f35b34801561108b57600080fd5b506111a66004803603810190808035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506135b0565b6040518082815260200191505060405180910390f35b3480156111c857600080fd5b506111d1613b48565b6040518082815260200191505060405180910390f35b3480156111f357600080fd5b50611278600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050613b4d565b6040518082815260200191505060405180910390f35b34801561129a57600080fd5b506112a3613b80565b6040518082815260200191505060405180910390f35b3480156112c557600080fd5b506112ce613b85565b60405180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b34801561132e57600080fd5b50611389600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050613c1f565b60405180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b3480156113e957600080fd5b506113f2613d15565b6040518082815260200191505060405180910390f35b34801561141457600080fd5b506114376004803603810190808035600019169060200190929190505050613d1a565b604051808781526020018681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001806020018060200180602001848103845287818151815260200191508051906020019080838360005b838110156114bd5780820151818401526020810190506114a2565b50505050905090810190601f1680156114ea5780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019080838360005b83811015611523578082015181840152602081019050611508565b50505050905090810190601f1680156115505780820380516001836020036101000a031916815260200191505b50848103825285818151815260200191508051906020019080838360005b8381101561158957808201518184015260208101905061156e565b50505050905090810190601f1680156115b65780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390f35b3480156115d757600080fd5b506115e0614019565b6040518082815260200191505060405180910390f35b34801561160257600080fd5b5061162160048036038101908080359060200190929190505050614025565b005b34801561162f57600080fd5b5061163861408c565b6040518082815260200191505060405180910390f35b34801561165a57600080fd5b50611687600480360381019080803560001916906020019092919080359060200190929190505050614096565b604051808215151515815260200191505060405180910390f35b3480156116ad57600080fd5b506116b66140b6565b6040518082815260200191505060405180910390f35b3480156116d857600080fd5b506118396004803603810190808035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506140bb565b604051808215151515815260200191505060405180910390f35b34801561185f57600080fd5b506118da6004803603810190808035906020019092919080351515906020019092919080359060200190929190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506145a8565b604051808215151515815260200191505060405180910390f35b34801561190057600080fd5b5061191f60048036038101908080359060200190929190505050614719565b005b34801561192d57600080fd5b50611936614780565b6040518082815260200191505060405180910390f35b34801561195857600080fd5b50611961614785565b6040518082815260200191505060405180910390f35b34801561198357600080fd5b506119b8600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061478a565b005b3480156119c657600080fd5b50611a75600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001909291905050506147f2565b6040518082815260200191505060405180910390f35b348015611a9757600080fd5b50611aa0614968565b6040518082815260200191505060405180910390f35b600060066000837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060009054906101000a900460ff169050919050565b60006301ffc9a77c010000000000000000000000000000000000000000000000000000000002905090565b600381565b600181565b6060600080611b6061651e565b6002600001600086600019166000191681526020019081526020016000206060604051908101604052908160008201805480602002602001604051908101604052809291908181526020018280548015611bd957602002820191906000526020600020905b815481526020019060010190808311611bc5575b505050505081526020016001820154815260200160028201546000191660001916815250509050806000015193508060200151925080604001519150509193909250565b60008282604051602001808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c01000000000000000000000000028152601401828152602001925050506040516020818303038152906040526040518082805190602001908083835b602083101515611cb75780518252602082019150602081019050602083039250611c92565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020905092915050565b6000600c54905090565b6000838383604051602001808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c0100000000000000000000000002815260140183815260200182805190602001908083835b602083101515611d7a5780518252602082019150602081019050602083039250611d55565b6001836020036101000a03801982511681845116808217855250505050505090500193505050506040516020818303038152906040526040518082805190602001908083835b602083101515611de55780518252602082019150602081019050602083039250611dc0565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902090509392505050565b6000611e2561496e565b1515611e3057600080fd5b600560009054906101000a900460ff16151515611e4c57600080fd5b611e62848460026149e39092919063ffffffff16565b15611e705760009050611e80565b611e7b848484614b00565b600190505b9392505050565b60008073ffffffffffffffffffffffffffffffffffffffff16600a6000846000191660001916815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614159050919050565b600381565b6060600b6000838152602001908152602001600020805480602002602001604051908101604052809291908181526020018280548015611f6657602002820191906000526020600020905b81546000191681526020019060010190808311611f4e575b50505050509050919050565b600281565b6000600754905090565b60015481565b60006120b3826040805190810160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a333200000000815250856040516020018083805190602001908083835b602083101515611ffc5780518252602082019150602081019050602083039250611fd7565b6001836020036101000a0380198251168184511680821785525050505050509050018260001916600019168152602001925050506040516020818303038152906040526040518082805190602001908083835b602083101515612074578051825260208201915060208101905060208303925061204f565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060001916614b5090919063ffffffff16565b905092915050565b6120c361496e565b15156120ce57600080fd5b600560009054906101000a900460ff1615156120e957600080fd5b6000600560006101000a81548160ff0219169083151502179055507f730c1faaa977b67dacf1e2451ef54556e04a07d577785ff79f6d31f73502efc960405160405180910390a1565b600181565b600080600560009054906101000a900460ff1615151561215657600080fd5b600754841415156121cf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f6e6f6e6365206d69736d6174636800000000000000000000000000000000000081525060200191505060405180910390fd5b6122fa87878787604051602001808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c0100000000000000000000000002815260140184815260200183805190602001908083835b6020831015156122555780518252602082019150602081019050602083039250612230565b6001836020036101000a0380198251168184511680821785525050505050509050018281526020019450505050506040516020818303038152906040526040518082805190602001908083835b6020831015156122c757805182526020820191506020810190506020830392506122a2565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902084611f87565b905061230881888888614c48565b91505095945050505050565b600061231e616543565b600080600560009054906101000a900460ff1615151561233d57600080fd5b846000600a6000836000191660001916815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff16141515156123a757600080fd5b6123af61496e565b156123b957612501565b8073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156123f257612500565b6124226123fd613398565b8273ffffffffffffffffffffffffffffffffffffffff16614ea190919063ffffffff16565b156124fa578073ffffffffffffffffffffffffffffffffffffffff1663d202158d61244c336131b0565b60026040518363ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180836000191660001916815260200182815260200192505050602060405180830381600087803b1580156124af57600080fd5b505af11580156124c3573d6000803e3d6000fd5b505050506040513d60208110156124d957600080fd5b810190808051906020019092919050505015156124f557600080fd5b6124ff565b600080fd5b5b5b600a6000886000191660001916815260200190815260200160002060c0604051908101604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561262b5780601f106126005761010080835404028352916020019161262b565b820191906000526020600020905b81548152906001019060200180831161260e57829003601f168201915b50505050508152602001600482018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156126cd5780601f106126a2576101008083540402835291602001916126cd565b820191906000526020600020905b8154815290600101906020018083116126b057829003601f168201915b50505050508152602001600582018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561276f5780601f106127445761010080835404028352916020019161276f565b820191906000526020600020905b81548152906001019060200180831161275257829003601f168201915b5050505050815250509450600073ffffffffffffffffffffffffffffffffffffffff16856040015173ffffffffffffffffffffffffffffffffffffffff16141515156127ba57600080fd5b600a6000886000191660001916815260200190815260200160002060008082016000905560018201600090556002820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905560038201600061281d9190616590565b60048201600061282d9190616590565b60058201600061283d91906165d8565b5050600b6000866000015181526020019081526020016000209350600092505b838054905083101561291b578660001916848481548110151561287c57fe5b906000526020600020015460001916141561290e578360018580549050038154811015156128a657fe5b906000526020600020015484848154811015156128bf57fe5b9060005260206000200181600019169055508360018580549050038154811015156128e657fe5b9060005260206000200160009055838054809190600190036129089190616620565b5061291b565b828060010193505061285d565b600c6000815480929190600190039190505550846040015173ffffffffffffffffffffffffffffffffffffffff16856000015188600019167f3cf57863a89432c61c4a27073c6ee39e8a764bff5a05aebfbcdcdc80b2e6130a886020015189606001518a608001518b60a0015160405180858152602001806020018060200180602001848103845287818151815260200191508051906020019080838360005b838110156129d65780820151818401526020810190506129bb565b50505050905090810190601f168015612a035780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019080838360005b83811015612a3c578082015181840152602081019050612a21565b50505050905090810190601f168015612a695780820380516001836020036101000a031916815260200191505b50848103825285818151815260200191508051906020019080838360005b83811015612aa2578082015181840152602081019050612a87565b50505050905090810190601f168015612acf5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a4600195505050505050919050565b600080612afa61496e565b1515612b0557600080fd5b600560009054906101000a900460ff16151515612b2157600080fd5b6001831415612b4f57612b3f60016002614f8890919063ffffffff16565b600054101515612b4e57600080fd5b5b612b65848460026149e39092919063ffffffff16565b1515612b745760009150612bc4565b612b8a84846002614fab9092919063ffffffff16565b9050808385600019167f585a4aef50f8267a92b32412b331b20f7f8b96f2245b253b9cc50dcc621d339760405160405180910390a4600191505b5092915050565b600560009054906101000a900460ff1681565b6000612be8616543565b600560009054906101000a900460ff16151515612c0457600080fd5b826000600a6000836000191660001916815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff1614151515612c6e57600080fd5b612c7661496e565b15612c8057612dc8565b8073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415612cb957612dc7565b612ce9612cc4613398565b8273ffffffffffffffffffffffffffffffffffffffff16614ea190919063ffffffff16565b15612dc1578073ffffffffffffffffffffffffffffffffffffffff1663d202158d612d13336131b0565b60026040518363ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180836000191660001916815260200182815260200192505050602060405180830381600087803b158015612d7657600080fd5b505af1158015612d8a573d6000803e3d6000fd5b505050506040513d6020811015612da057600080fd5b81019080805190602001909291905050501515612dbc57600080fd5b612dc6565b600080fd5b5b5b600a6000866000191660001916815260200190815260200160002060c0604051908101604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015612ef25780601f10612ec757610100808354040283529160200191612ef2565b820191906000526020600020905b815481529060010190602001808311612ed557829003601f168201915b50505050508152602001600482018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015612f945780601f10612f6957610100808354040283529160200191612f94565b820191906000526020600020905b815481529060010190602001808311612f7757829003601f168201915b50505050508152602001600582018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156130365780601f1061300b57610100808354040283529160200191613036565b820191906000526020600020905b81548152906001019060200180831161301957829003601f168201915b5050505050815250509250600073ffffffffffffffffffffffffffffffffffffffff16836040015173ffffffffffffffffffffffffffffffffffffffff161415151561308157600080fd5b6130a2836000015184602001518560400151866060015187608001516151f0565b15156130bb576130b185612314565b50600093506130c0565b600193505b505050919050565b60086020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001015490806002018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156131a05780601f10613175576101008083540402835291602001916131a0565b820191906000526020600020905b81548152906001019060200180831161318357829003601f168201915b5050505050908060030154905084565b60008173ffffffffffffffffffffffffffffffffffffffff166001029050919050565b6009602052816000526040600020818154811015156131ee57fe5b906000526020600020016000915091509054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600181565b600061322f61496e565b151561323a57600080fd5b600560009054906101000a900460ff1615151561325657600080fd5b61326a8560026153c090919063ffffffff16565b151561327557600080fd5b61328f85858585600261548290949392919063ffffffff16565b5060019050949350505050565b6000600560009054906101000a900460ff161515156132ba57600080fd5b6132c5338484615559565b905092915050565b600281565b600060026000016000856000191660001916815260200190815260200160002060030160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060009054906101000a900460ff1690509392505050565b60006353d413c57c01000000000000000000000000000000000000000000000000000000000263747442d37c01000000000000000000000000000000000000000000000000000000000263b61d27f67c010000000000000000000000000000000000000000000000000000000002631d3812407c010000000000000000000000000000000000000000000000000000000002639010f7267c01000000000000000000000000000000000000000000000000000000000263d202158d7c0100000000000000000000000000000000000000000000000000000000026312aaac707c010000000000000000000000000000000000000000000000000000000002181818181818905090565b600181565b6134ae61496e565b15156134b957600080fd5b600560009054906101000a900460ff161515156134d557600080fd5b6001600560006101000a81548160ff0219169083151502179055507f4b314b34e912fda7f95e7d23e9c8c95f82f0aff1984e4ce592a0b005f905562460405160405180910390a1565b60606002600101600083815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561358457602002820191906000526020600020905b8154600019168152602001906001019080831161356c575b50505050509050919050565b600481565b600281565b60005481565b600481565b60075481565b600881565b600080600080600560009054906101000a900460ff161515156135d257600080fd5b6135df8a8a8a8a8a6151f0565b15156135ea57600080fd5b6135f261496e565b925082151561389a573073ffffffffffffffffffffffffffffffffffffffff1663b61d27f630600080366040518563ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001848152602001806020018281038252848482818152602001925080828437820191505095505050505050602060405180830381600087803b1580156136c157600080fd5b505af11580156136d5573d6000803e3d6000fd5b505050506040513d60208110156136eb57600080fd5b810190808051906020019092919050505093508773ffffffffffffffffffffffffffffffffffffffff168a857fe6b6db97dedfb44cbced367387e40b3f92dbee1b1be87f9367f7b673532b0d698c8b8b8b60405180858152602001806020018060200180602001848103845287818151815260200191508051906020019080838360005b8381101561378a57808201518184015260208101905061376f565b50505050905090810190601f1680156137b75780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019080838360005b838110156137f05780820151818401526020810190506137d5565b50505050905090810190601f16801561381d5780820380516001836020036101000a031916815260200191505b50848103825285818151815260200191508051906020019080838360005b8381101561385657808201518184015260208101905061383b565b50505050905090810190601f1680156138835780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a4613b3b565b6138a4888b611c1d565b9150600073ffffffffffffffffffffffffffffffffffffffff16600a6000846000191660001916815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561392d57613928828b8b8b8b8b8b615ade565b613b3a565b600a6000836000191660001916815260200190815260200160002090508881600101819055508681600301908051906020019061396b92919061664c565b508581600401908051906020019061398492919061664c565b508481600501908051906020019061399d9291906166cc565b508773ffffffffffffffffffffffffffffffffffffffff168a83600019167f3bab293fc00db832d7619a9299914251b8747c036867ec056cbd506f60135b138c8b8b8b60405180858152602001806020018060200180602001848103845287818151815260200191508051906020019080838360005b83811015613a2e578082015181840152602081019050613a13565b50505050905090810190601f168015613a5b5780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019080838360005b83811015613a94578082015181840152602081019050613a79565b50505050905090810190601f168015613ac15780820380516001836020036101000a031916815260200191505b50848103825285818151815260200191508051906020019080838360005b83811015613afa578082015181840152602081019050613adf565b50505050905090810190601f168015613b275780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a45b5b5050509695505050505050565b600581565b6000600560009054906101000a900460ff16151515613b6b57600080fd5b613b7733858585614c48565b90509392505050565b600781565b6000634eee424a7c01000000000000000000000000000000000000000000000000000000000263b1a34e0d7c01000000000000000000000000000000000000000000000000000000000263262b54f57c01000000000000000000000000000000000000000000000000000000000263c9100bcb7c010000000000000000000000000000000000000000000000000000000002181818905090565b60008060008090505b6004811015613d0b576008810260ff7f0100000000000000000000000000000000000000000000000000000000000000028583815181101515613c6757fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f010000000000000000000000000000000000000000000000000000000000000002167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19169060020a9004821791508080600101915050613c28565b8192505050919050565b600381565b60008060006060806060613d2c616543565b600a6000896000191660001916815260200190815260200160002060c0604051908101604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015613e565780601f10613e2b57610100808354040283529160200191613e56565b820191906000526020600020905b815481529060010190602001808311613e3957829003601f168201915b50505050508152602001600482018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015613ef85780601f10613ecd57610100808354040283529160200191613ef8565b820191906000526020600020905b815481529060010190602001808311613edb57829003601f168201915b50505050508152602001600582018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015613f9a5780601f10613f6f57610100808354040283529160200191613f9a565b820191906000526020600020905b815481529060010190602001808311613f7d57829003601f168201915b5050505050815250509050600073ffffffffffffffffffffffffffffffffffffffff16816040015173ffffffffffffffffffffffffffffffffffffffff1614151515613fe557600080fd5b80600001519650806020015195508060400151945080606001519350806080015192508060a0015191505091939550919395565b60006002800154905090565b6000600560009054906101000a900460ff1615151561404357600080fd5b61404b61496e565b151561405657600080fd5b60008211151561406557600080fd5b61406f600161351e565b51905080821115151561408157600080fd5b816000819055505050565b6000600754905090565b60006140ae838360026149e39092919063ffffffff16565b905092915050565b600581565b600080600080600560009054906101000a900460ff161515156140dd57600080fd5b6140ea8b8b8b8b8b6151f0565b15156140f557600080fd5b6142ca8b8b8b8b8b8b604051602001808781526020018681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c0100000000000000000000000002815260140184805190602001908083835b602083101515614183578051825260208201915060208101905060208303925061415e565b6001836020036101000a03801982511681845116808217855250505050505090500183805190602001908083835b6020831015156141d657805182526020820191506020810190506020830392506141b1565b6001836020036101000a03801982511681845116808217855250505050505090500182805190602001908083835b6020831015156142295780518252602082019150602081019050602083039250614204565b6001836020036101000a03801982511681845116808217855250505050505090500196505050505050506040516020818303038152906040526040518082805190602001908083835b6020831015156142975780518252602082019150602081019050602083039250614272565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902086611f87565b92506142eb6142d8846131b0565b600160026149e39092919063ffffffff16565b15156142f657600080fd5b614300898c611c1d565b9150600073ffffffffffffffffffffffffffffffffffffffff16600a6000846000191660001916815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561438957614384828c8c8c8c8c8c615ade565b614596565b600a600083600019166000191681526020019081526020016000209050898160010181905550878160030190805190602001906143c792919061664c565b50868160040190805190602001906143e092919061664c565b50858160050190805190602001906143f99291906166cc565b508873ffffffffffffffffffffffffffffffffffffffff168b83600019167f3bab293fc00db832d7619a9299914251b8747c036867ec056cbd506f60135b138d8c8c8c60405180858152602001806020018060200180602001848103845287818151815260200191508051906020019080838360005b8381101561448a57808201518184015260208101905061446f565b50505050905090810190601f1680156144b75780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019080838360005b838110156144f05780820151818401526020810190506144d5565b50505050905090810190601f16801561451d5780820380516001836020036101000a031916815260200191505b50848103825285818151815260200191508051906020019080838360005b8381101561455657808201518184015260208101905061453b565b50505050905090810190601f1680156145835780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a45b60019350505050979650505050505050565b600080600560009054906101000a900460ff161515156145c757600080fd5b60075484141515614640576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f6e6f6e6365206d69736d6174636800000000000000000000000000000000000081525060200191505060405180910390fd5b6147018686866040516020018084815260200183151515157f010000000000000000000000000000000000000000000000000000000000000002815260010182815260200193505050506040516020818303038152906040526040518082805190602001908083835b6020831015156146ce57805182526020820191506020810190506020830392506146a9565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902084611f87565b905061470e818787615559565b915050949350505050565b6000600560009054906101000a900460ff1615151561473757600080fd5b61473f61496e565b151561474a57600080fd5b60008211151561475957600080fd5b614763600261351e565b51905080821115151561477557600080fd5b816001819055505050565b600681565b600281565b61479261496e565b151561479d57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156147d957600080fd5b8073ffffffffffffffffffffffffffffffffffffffff16ff5b60008585858585604051602001808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c0100000000000000000000000002815260140184815260200183805190602001908083835b6020831015156148b95780518252602082019150602081019050602083039250614894565b6001836020036101000a038019825116818451168082178552505050505050905001828152602001955050505050506040516020818303038152906040526040518082805190602001908083835b60208310151561492c5780518252602082019150602081019050602083039250614907565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060019004905095945050505050565b600c5481565b60003073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156149ad57600190506149e0565b60016000541415156149be57600080fd5b6149dd6149ca336131b0565b600160026149e39092919063ffffffff16565b90505b90565b60006149ed61651e565b600085600001600086600019166000191681526020019081526020016000206060604051908101604052908160008201805480602002602001604051908101604052809291908181526020018280548015614a6757602002820191906000526020600020905b815481526020019060010190808311614a53575b50505050508152602001600182015481526020016002820154600019166000191681525050915060006001028260400151600019161415614aab5760009250614af7565b600090505b816000015151811015614af65783826000015182815181101515614ad057fe5b906020019060200201511415614ae95760019250614af7565b8080600101915050614ab0565b5b50509392505050565b614b188383836002615de5909392919063ffffffff16565b808284600019167f480000bb1edad8ca1470381cc334b1917fbd51c6531f3a623ea8e0ec7e38a6e960405160405180910390a4505050565b60008060008060418551141515614b6a5760009350614c3f565b6020850151925060408501519150606085015160001a9050601b8160ff161015614b9557601b810190505b601b8160ff1614158015614bad5750601c8160ff1614155b15614bbb5760009350614c3f565b600186828585604051600081526020016040526040518085600019166000191681526020018460ff1660ff1681526020018360001916600019168152602001826000191660001916815260200194505050505060206040516020810390808403906000865af1158015614c32573d6000803e3d6000fd5b5050506020604051035193505b50505092915050565b600080614c5361674c565b614c5e878786615ebf565b9150614c6f308787876007546147f2565b9250848673ffffffffffffffffffffffffffffffffffffffff16847f8afcfabcb00e47a53a8fc3e9f23ff47ee1926194bb1350dd007c50b412a6cee8876040518080602001828103825283818151815260200191508051906020019080838360005b83811015614cec578082015181840152602081019050614cd1565b50505050905090810190601f168015614d195780820380516001836020036101000a031916815260200191505b509250505060405180910390a46007600081548092919060010191905055506080604051908101604052808773ffffffffffffffffffffffffffffffffffffffff1681526020018681526020018581526020018381525090506000821415614d8d57614d87838260006161df565b50614e94565b806008600085815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101556040820151816002019080519060200190614e0e92919061678b565b5060608201518160030155905050600960008481526020019081526020016000208790806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505b8292505050949350505050565b6000806000614ed3856301ffc9a77c0100000000000000000000000000000000000000000000000000000000026164cd565b80925081935050506000821480614eea5750600081145b15614ef85760009250614f80565b614f258563ffffffff7c0100000000000000000000000000000000000000000000000000000000026164cd565b80925081935050506000821480614f3d575060008114155b15614f4b5760009250614f80565b614f5585856164cd565b8092508193505050600182148015614f6d5750600181145b15614f7b5760019250614f80565b600092505b505092915050565b600082600101600083815260200190815260200160002080549050905092915050565b600080600080866000016000876000191660001916815260200190815260200160002060010154935086600001600087600019166000191681526020019081526020016000206000019250600091505b82805490508210156150c25784838381548110151561501657fe5b906000526020600020015414156150b55782600184805490500381548110151561503c57fe5b9060005260206000200154838381548110151561505557fe5b906000526020600020018190555082600184805490500381548110151561507857fe5b90600052602060002001600090558280548091906001900361509a919061680b565b508660020160008154809291906001900391905055506150c2565b8180600101925050614ffb565b600083805490501415615110578660000160008760001916600019168152602001908152602001600020600080820160006150fd9190616837565b6001820160009055600282016000905550505b8660010160008681526020019081526020016000209050600091505b80805490508210156151e6578560001916818381548110151561514b57fe5b90600052602060002001546000191614156151d95780600182805490500381548110151561517557fe5b9060005260206000200154818381548110151561518e57fe5b9060005260206000200181600019169055508060018280549050038154811015156151b557fe5b9060005260206000200160009055808054809190600190036151d79190616620565b505b818060010192505061512c565b5050509392505050565b60008060018614156153b15761521061520a308986611cf6565b85611f87565b90508073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16141561524f57600191506153b6565b3073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614156152a9576152a261528f826131b0565b600360026149e39092919063ffffffff16565b91506153b6565b6152d96152b4613398565b8673ffffffffffffffffffffffffffffffffffffffff16614ea190919063ffffffff16565b156153a8578473ffffffffffffffffffffffffffffffffffffffff1663d202158d615303836131b0565b60036040518363ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180836000191660001916815260200182815260200192505050602060405180830381600087803b15801561536657600080fd5b505af115801561537a573d6000803e3d6000fd5b505050506040513d602081101561539057600080fd5b810190808051906020019092919050505091506153b6565b600091506153b6565b600091505b5095945050505050565b60006153ca61651e565b8360000160008460001916600019168152602001908152602001600020606060405190810160405290816000820180548060200260200160405190810160405280929190818152602001828054801561544257602002820191906000526020600020905b81548152602001906001019080831161542e575b5050505050815260200160018201548152602001600282015460001916600019168152505090506000600102816040015160001916141591505092915050565b60008086600001600087600019166000191681526020019081526020016000209050828160030160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000867bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060006101000a81548160ff0219169083151502179055508291505095945050505050565b6000806000806000861415151561556f57600080fd5b60086000878152602001908152602001600020925060008360000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515156155ce57600080fd5b615699878460000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16856002018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561568f5780601f106156645761010080835404028352916020019161568f565b820191906000526020600020905b81548152906001019060200180831161567257829003601f168201915b5050505050615ebf565b50857fb3932da477fe5d6c8ff2eafef050c0f3a1af18fc07121001482600f36f3715d886604051808215151515815260200191505060405180910390a260096000878152602001908152602001600020915084151561588e57600090505b8180549050811015615885578673ffffffffffffffffffffffffffffffffffffffff16828281548110151561572857fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156158785781600183805490500381548110151561578457fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682828154811015156157bd57fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600183805490500381548110151561581a57fe5b9060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690558180548091906001900361585b9190616858565b506001836003016000828254019250508190555060019350615ad4565b80806001019150506156f7565b60009350615ad4565b600090505b818054905081101561591f578673ffffffffffffffffffffffffffffffffffffffff1682828154811015156158c457fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415151561591257600080fd5b8080600101915050615893565b818790806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505060018360030160008282540392505081905550600083600301541415615acf57615ac88684608060405190810160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201548152602001600282018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015615aae5780601f10615a8357610100808354040283529160200191615aae565b820191906000526020600020905b815481529060010190602001808311615a9157829003601f168201915b5050505050815260200160038201548152505060016161df565b9350615ad4565b600193505b5050509392505050565b60c0604051908101604052808781526020018681526020018573ffffffffffffffffffffffffffffffffffffffff16815260200184815260200183815260200182815250600a60008960001916600019168152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816003019080519060200190615bb492919061678b565b506080820151816004019080519060200190615bd192919061678b565b5060a0820151816005019080519060200190615bee929190616884565b50905050600b60008781526020019081526020016000208790806001815401808255809150509060018203906000526020600020016000909192909190915090600019169055506001600c5401600c819055508373ffffffffffffffffffffffffffffffffffffffff168688600019167f46149b18aa084502c3f12bc75e19eda8bda8d102b82cce8474677a6d0d5f43c58887878760405180858152602001806020018060200180602001848103845287818151815260200191508051906020019080838360005b83811015615cd1578082015181840152602081019050615cb6565b50505050905090810190601f168015615cfe5780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019080838360005b83811015615d37578082015181840152602081019050615d1c565b50505050905090810190601f168015615d645780820380516001836020036101000a031916815260200191505b50848103825285818151815260200191508051906020019080838360005b83811015615d9d578082015181840152602081019050615d82565b50505050905090810190601f168015615dca5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a450505050505050565b6000846000016000856000191660001916815260200190815260200160002090508060000183908060018154018082558091505090600182039060005260206000200160009091929091909150555060006001028160020154600019161415615e5f57838160020181600019169055508181600101819055505b84600101600084815260200190815260200160002084908060018154018082558091505090600182039060005260206000200160009091929091909150906000191690555084600201600081548092919060010191905055505050505050565b60003073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415616009573073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415615f33576000549050616004565b615f52615f3f856131b0565b600160026149e39092919063ffffffff16565b80615fec5750615f77615f64856131b0565b600760026149e39092919063ffffffff16565b8015615feb5750631d3812407c0100000000000000000000000000000000000000000000000000000000027bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916615fca83613c1f565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b5b15615ffe576001600054039050616003565b600080fd5b5b6161d8565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415151561604557600080fd5b3073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614156160835760015490506161d7565b6160a161608f856131b0565b6002806149e39092919063ffffffff16565b806160c757506160c66160b3856131b0565b600660026149e39092919063ffffffff16565b5b806161c057506160ec6160d9856131b0565b600860026149e39092919063ffffffff16565b80156161bf575060026000016000616103866131b0565b6000191660001916815260200190815260200160002060030160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600061616484613c1f565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060009054906101000a900460ff165b5b156161d157600180540390506161d6565b600080fd5b5b5b9392505050565b6000806000846000015173ffffffffffffffffffffffffffffffffffffffff161415151561620c57600080fd5b836000015173ffffffffffffffffffffffffffffffffffffffff168460200151856040015160405180828051906020019080838360005b8381101561625e578082015181840152602081019050616243565b50505050905090810190601f16801561628b5780820380516001836020036101000a031916815260200191505b5091505060006040518083038185875af1925050509050801515616373578360200151846000015173ffffffffffffffffffffffffffffffffffffffff16867fe10c49d9f7c71da23262367013434763cfdb2332267641728d25cd712c5c6a6887604001516040518080602001828103825283818151815260200191508051906020019080838360005b83811015616330578082015181840152602081019050616315565b50505050905090810190601f16801561635d5780820380516001836020036101000a031916815260200191505b509250505060405180910390a4600091506164c5565b8360200151846000015173ffffffffffffffffffffffffffffffffffffffff16867f1f920dbda597d7bf95035464170fa58d0a4b57f13a1c315ace6793b9f63688b887604001516040518080602001828103825283818151815260200191508051906020019080838360005b838110156163fa5780820151818401526020810190506163df565b50505050905090810190601f1680156164275780820380516001836020036101000a031916815260200191505b509250505060405180910390a482151561644457600191506164c5565b60086000868152602001908152602001600020600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905560018201600090556002820160006164979190616590565b600382016000905550506009600086815260200190815260200160002060006164c09190616904565b600191505b509392505050565b60008060006301ffc9a77c010000000000000000000000000000000000000000000000000000000002905060405181815284600482015260208160208389617530fa93508051925050509250929050565b6060604051908101604052806060815260200160008152602001600080191681525090565b60c0604051908101604052806000815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016060815260200160608152602001606081525090565b50805460018160011615610100020316600290046000825580601f106165b657506165d5565b601f0160209004906000526020600020908101906165d49190616925565b5b50565b50805460018160011615610100020316600290046000825580601f106165fe575061661d565b601f01602090049060005260206000209081019061661c9190616925565b5b50565b81548183558181111561664757818360005260206000209182019101616646919061694a565b5b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061668d57805160ff19168380011785556166bb565b828001600101855582156166bb579182015b828111156166ba57825182559160200191906001019061669f565b5b5090506166c89190616925565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061670d57805160ff191683800117855561673b565b8280016001018555821561673b579182015b8281111561673a57825182559160200191906001019061671f565b5b5090506167489190616925565b5090565b608060405190810160405280600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160608152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106167cc57805160ff19168380011785556167fa565b828001600101855582156167fa579182015b828111156167f95782518255916020019190600101906167de565b5b5090506168079190616925565b5090565b815481835581811115616832578183600052602060002091820191016168319190616925565b5b505050565b50805460008255906000526020600020908101906168559190616925565b50565b81548183558181111561687f5781836000526020600020918201910161687e9190616925565b5b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106168c557805160ff19168380011785556168f3565b828001600101855582156168f3579182015b828111156168f25782518255916020019190600101906168d7565b5b5090506169009190616925565b5090565b50805460008255906000526020600020908101906169229190616925565b50565b61694791905b8082111561694357600081600090555060010161692b565b5090565b90565b61696c91905b80821115616968576000816000905550600101616950565b5090565b905600a165627a7a723058200d6c15bc4ca99efd2ba9b51757557595d474bf81cc0f4c68694cb5a99d64c8720029`

// DeployIdentity deploys a new Ethereum contract, binding an instance of Identity to it.
func DeployIdentity(auth *bind.TransactOpts, backend bind.ContractBackend, _managementKey common.Address) (common.Address, *types.Transaction, *Identity, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IdentityBin), backend, _managementKey)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Identity{IdentityCaller: IdentityCaller{contract: contract}, IdentityTransactor: IdentityTransactor{contract: contract}, IdentityFilterer: IdentityFilterer{contract: contract}}, nil
}

// Identity is an auto generated Go binding around an Ethereum contract.
type Identity struct {
	Address            *common.Address
	IdentityCaller     // Read-only binding to the contract
	IdentityTransactor // Write-only binding to the contract
	IdentityFilterer   // Log filterer for contract events
}

// IdentityCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentitySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentitySession struct {
	Contract     *Identity         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdentityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityCallerSession struct {
	Contract *IdentityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IdentityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityTransactorSession struct {
	Contract     *IdentityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IdentityRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityRaw struct {
	Contract *Identity // Generic contract binding to access the raw methods on
}

// IdentityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityCallerRaw struct {
	Contract *IdentityCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityTransactorRaw struct {
	Contract *IdentityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentity creates a new instance of Identity, bound to a specific deployed contract.
func NewIdentity(address common.Address, backend bind.ContractBackend) (*Identity, error) {
	contract, err := bindIdentity(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Identity{Address: &address, IdentityCaller: IdentityCaller{contract: contract}, IdentityTransactor: IdentityTransactor{contract: contract}, IdentityFilterer: IdentityFilterer{contract: contract}}, nil
}

// NewIdentityCaller creates a new read-only instance of Identity, bound to a specific deployed contract.
func NewIdentityCaller(address common.Address, caller bind.ContractCaller) (*IdentityCaller, error) {
	contract, err := bindIdentity(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityCaller{contract: contract}, nil
}

// NewIdentityTransactor creates a new write-only instance of Identity, bound to a specific deployed contract.
func NewIdentityTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityTransactor, error) {
	contract, err := bindIdentity(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityTransactor{contract: contract}, nil
}

// NewIdentityFilterer creates a new log filterer instance of Identity, bound to a specific deployed contract.
func NewIdentityFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityFilterer, error) {
	contract, err := bindIdentity(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityFilterer{contract: contract}, nil
}

// bindIdentity binds a generic wrapper to an already deployed contract.
func bindIdentity(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Identity *IdentityRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	var results []interface{}
	err := _Identity.Contract.IdentityCaller.contract.Call(opts, &results, method, params...)
	result = results[0]
	return err
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Identity *IdentityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identity.Contract.IdentityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Identity *IdentityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Identity.Contract.IdentityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Identity *IdentityCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	var results []interface{}
	err := _Identity.Contract.contract.Call(opts, &results, method, params...)
	result = results[0]
	return err
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Identity *IdentityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identity.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Identity *IdentityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Identity.Contract.contract.Transact(opts, method, params...)
}

// ACTIONKEY is a free data retrieval call binding the contract method 0x75e5598c.
//
// Solidity: function ACTION_KEY() constant returns(uint256)
func (_Identity *IdentityCaller) ACTIONKEY(opts *bind.CallOpts) (*big.Int, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "ACTION_KEY")
	return results[0].(*big.Int), err
}

// ACTIONKEY is a free data retrieval call binding the contract method 0x75e5598c.
//
// Solidity: function ACTION_KEY() constant returns(uint256)
func (_Identity *IdentitySession) ACTIONKEY() (*big.Int, error) {
	return _Identity.Contract.ACTIONKEY(&_Identity.CallOpts)
}

// ACTIONKEY is a free data retrieval call binding the contract method 0x75e5598c.
//
// Solidity: function ACTION_KEY() constant returns(uint256)
func (_Identity *IdentityCallerSession) ACTIONKEY() (*big.Int, error) {
	return _Identity.Contract.ACTIONKEY(&_Identity.CallOpts)
}

// ASSISTKEY is a free data retrieval call binding the contract method 0xdbfa74b7.
//
// Solidity: function ASSIST_KEY() constant returns(uint256)
func (_Identity *IdentityCaller) ASSISTKEY(opts *bind.CallOpts) (*big.Int, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "ASSIST_KEY")
	return results[0].(*big.Int), err
}

// ASSISTKEY is a free data retrieval call binding the contract method 0xdbfa74b7.
//
// Solidity: function ASSIST_KEY() constant returns(uint256)
func (_Identity *IdentitySession) ASSISTKEY() (*big.Int, error) {
	return _Identity.Contract.ASSISTKEY(&_Identity.CallOpts)
}

// ASSISTKEY is a free data retrieval call binding the contract method 0xdbfa74b7.
//
// Solidity: function ASSIST_KEY() constant returns(uint256)
func (_Identity *IdentityCallerSession) ASSISTKEY() (*big.Int, error) {
	return _Identity.Contract.ASSISTKEY(&_Identity.CallOpts)
}

// CLAIMSIGNERKEY is a free data retrieval call binding the contract method 0xc6702187.
//
// Solidity: function CLAIM_SIGNER_KEY() constant returns(uint256)
func (_Identity *IdentityCaller) CLAIMSIGNERKEY(opts *bind.CallOpts) (*big.Int, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "CLAIM_SIGNER_KEY")
	return results[0].(*big.Int), err
}

// CLAIMSIGNERKEY is a free data retrieval call binding the contract method 0xc6702187.
//
// Solidity: function CLAIM_SIGNER_KEY() constant returns(uint256)
func (_Identity *IdentitySession) CLAIMSIGNERKEY() (*big.Int, error) {
	return _Identity.Contract.CLAIMSIGNERKEY(&_Identity.CallOpts)
}

// CLAIMSIGNERKEY is a free data retrieval call binding the contract method 0xc6702187.
//
// Solidity: function CLAIM_SIGNER_KEY() constant returns(uint256)
func (_Identity *IdentityCallerSession) CLAIMSIGNERKEY() (*big.Int, error) {
	return _Identity.Contract.CLAIMSIGNERKEY(&_Identity.CallOpts)
}

// CONTRACTSCHEME is a free data retrieval call binding the contract method 0x251de3e9.
//
// Solidity: function CONTRACT_SCHEME() constant returns(uint256)
func (_Identity *IdentityCaller) CONTRACTSCHEME(opts *bind.CallOpts) (*big.Int, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "CLAIM_SIGNER_KEY")
	return results[0].(*big.Int), err
}

// CONTRACTSCHEME is a free data retrieval call binding the contract method 0x251de3e9.
//
// Solidity: function CONTRACT_SCHEME() constant returns(uint256)
func (_Identity *IdentitySession) CONTRACTSCHEME() (*big.Int, error) {
	return _Identity.Contract.CONTRACTSCHEME(&_Identity.CallOpts)
}

// CONTRACTSCHEME is a free data retrieval call binding the contract method 0x251de3e9.
//
// Solidity: function CONTRACT_SCHEME() constant returns(uint256)
func (_Identity *IdentityCallerSession) CONTRACTSCHEME() (*big.Int, error) {
	return _Identity.Contract.CONTRACTSCHEME(&_Identity.CallOpts)
}

// CUSTOMKEY is a free data retrieval call binding the contract method 0xb132734e.
//
// Solidity: function CUSTOM_KEY() constant returns(uint256)
func (_Identity *IdentityCaller) CUSTOMKEY(opts *bind.CallOpts) (*big.Int, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "CUSTOM_KEY")
	return results[0].(*big.Int), err
}

// CUSTOMKEY is a free data retrieval call binding the contract method 0xb132734e.
//
// Solidity: function CUSTOM_KEY() constant returns(uint256)
func (_Identity *IdentitySession) CUSTOMKEY() (*big.Int, error) {
	return _Identity.Contract.CUSTOMKEY(&_Identity.CallOpts)
}

// CUSTOMKEY is a free data retrieval call binding the contract method 0xb132734e.
//
// Solidity: function CUSTOM_KEY() constant returns(uint256)
func (_Identity *IdentityCallerSession) CUSTOMKEY() (*big.Int, error) {
	return _Identity.Contract.CUSTOMKEY(&_Identity.CallOpts)
}

// DELEGATEKEY is a free data retrieval call binding the contract method 0xead09fab.
//
// Solidity: function DELEGATE_KEY() constant returns(uint256)
func (_Identity *IdentityCaller) DELEGATEKEY(opts *bind.CallOpts) (*big.Int, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "DELEGATE_KEY")
	return results[0].(*big.Int), err
}

// DELEGATEKEY is a free data retrieval call binding the contract method 0xead09fab.
//
// Solidity: function DELEGATE_KEY() constant returns(uint256)
func (_Identity *IdentitySession) DELEGATEKEY() (*big.Int, error) {
	return _Identity.Contract.DELEGATEKEY(&_Identity.CallOpts)
}

// DELEGATEKEY is a free data retrieval call binding the contract method 0xead09fab.
//
// Solidity: function DELEGATE_KEY() constant returns(uint256)
func (_Identity *IdentityCallerSession) DELEGATEKEY() (*big.Int, error) {
	return _Identity.Contract.DELEGATEKEY(&_Identity.CallOpts)
}

// ECDSASCHEME is a free data retrieval call binding the contract method 0x82d09446.
//
// Solidity: function ECDSA_SCHEME() constant returns(uint256)
func (_Identity *IdentityCaller) ECDSASCHEME(opts *bind.CallOpts) (*big.Int, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "ECDSA_SCHEME")
	return results[0].(*big.Int), err
}

// ECDSASCHEME is a free data retrieval call binding the contract method 0x82d09446.
//
// Solidity: function ECDSA_SCHEME() constant returns(uint256)
func (_Identity *IdentitySession) ECDSASCHEME() (*big.Int, error) {
	return _Identity.Contract.ECDSASCHEME(&_Identity.CallOpts)
}

// ECDSASCHEME is a free data retrieval call binding the contract method 0x82d09446.
//
// Solidity: function ECDSA_SCHEME() constant returns(uint256)
func (_Identity *IdentityCallerSession) ECDSASCHEME() (*big.Int, error) {
	return _Identity.Contract.ECDSASCHEME(&_Identity.CallOpts)
}

// ECDSATYPE is a free data retrieval call binding the contract method 0x49991ec8.
//
// Solidity: function ECDSA_TYPE() constant returns(uint256)
func (_Identity *IdentityCaller) ECDSATYPE(opts *bind.CallOpts) (*big.Int, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "ECDSA_TYPE")
	return results[0].(*big.Int), err
}

// ECDSATYPE is a free data retrieval call binding the contract method 0x49991ec8.
//
// Solidity: function ECDSA_TYPE() constant returns(uint256)
func (_Identity *IdentitySession) ECDSATYPE() (*big.Int, error) {
	return _Identity.Contract.ECDSATYPE(&_Identity.CallOpts)
}

// ECDSATYPE is a free data retrieval call binding the contract method 0x49991ec8.
//
// Solidity: function ECDSA_TYPE() constant returns(uint256)
func (_Identity *IdentityCallerSession) ECDSATYPE() (*big.Int, error) {
	return _Identity.Contract.ECDSATYPE(&_Identity.CallOpts)
}

// ENCRYPTIONKEY is a free data retrieval call binding the contract method 0x9e140cc8.
//
// Solidity: function ENCRYPTION_KEY() constant returns(uint256)
func (_Identity *IdentityCaller) ENCRYPTIONKEY(opts *bind.CallOpts) (*big.Int, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "ENCRYPTION_KEY")
	return results[0].(*big.Int), err
}

// ENCRYPTIONKEY is a free data retrieval call binding the contract method 0x9e140cc8.
//
// Solidity: function ENCRYPTION_KEY() constant returns(uint256)
func (_Identity *IdentitySession) ENCRYPTIONKEY() (*big.Int, error) {
	return _Identity.Contract.ENCRYPTIONKEY(&_Identity.CallOpts)
}

// ENCRYPTIONKEY is a free data retrieval call binding the contract method 0x9e140cc8.
//
// Solidity: function ENCRYPTION_KEY() constant returns(uint256)
func (_Identity *IdentityCallerSession) ENCRYPTIONKEY() (*big.Int, error) {
	return _Identity.Contract.ENCRYPTIONKEY(&_Identity.CallOpts)
}

// ERC165ID is a free data retrieval call binding the contract method 0x02e7491e.
//
// Solidity: function ERC165ID() constant returns(bytes4)
func (_Identity *IdentityCaller) ERC165ID(opts *bind.CallOpts) ([4]byte, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "ERC165ID")
	return results[0].([4]byte), err
}

// ERC165ID is a free data retrieval call binding the contract method 0x02e7491e.
//
// Solidity: function ERC165ID() constant returns(bytes4)
func (_Identity *IdentitySession) ERC165ID() ([4]byte, error) {
	return _Identity.Contract.ERC165ID(&_Identity.CallOpts)
}

// ERC165ID is a free data retrieval call binding the contract method 0x02e7491e.
//
// Solidity: function ERC165ID() constant returns(bytes4)
func (_Identity *IdentityCallerSession) ERC165ID() ([4]byte, error) {
	return _Identity.Contract.ERC165ID(&_Identity.CallOpts)
}

// ERC725ID is a free data retrieval call binding the contract method 0x7d96fa58.
//
// Solidity: function ERC725ID() constant returns(bytes4)
func (_Identity *IdentityCaller) ERC725ID(opts *bind.CallOpts) ([4]byte, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "ERC725ID")
	return results[0].([4]byte), err
}

// ERC725ID is a free data retrieval call binding the contract method 0x7d96fa58.
//
// Solidity: function ERC725ID() constant returns(bytes4)
func (_Identity *IdentitySession) ERC725ID() ([4]byte, error) {
	return _Identity.Contract.ERC725ID(&_Identity.CallOpts)
}

// ERC725ID is a free data retrieval call binding the contract method 0x7d96fa58.
//
// Solidity: function ERC725ID() constant returns(bytes4)
func (_Identity *IdentityCallerSession) ERC725ID() ([4]byte, error) {
	return _Identity.Contract.ERC725ID(&_Identity.CallOpts)
}

// ERC735ID is a free data retrieval call binding the contract method 0xbf2f20ad.
//
// Solidity: function ERC735ID() constant returns(bytes4)
func (_Identity *IdentityCaller) ERC735ID(opts *bind.CallOpts) ([4]byte, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "ERC735ID")
	return results[0].([4]byte), err
}

// ERC735ID is a free data retrieval call binding the contract method 0xbf2f20ad.
//
// Solidity: function ERC735ID() constant returns(bytes4)
func (_Identity *IdentitySession) ERC735ID() ([4]byte, error) {
	return _Identity.Contract.ERC735ID(&_Identity.CallOpts)
}

// ERC735ID is a free data retrieval call binding the contract method 0xbf2f20ad.
//
// Solidity: function ERC735ID() constant returns(bytes4)
func (_Identity *IdentityCallerSession) ERC735ID() ([4]byte, error) {
	return _Identity.Contract.ERC735ID(&_Identity.CallOpts)
}

// LABELTOPIC is a free data retrieval call binding the contract method 0xb1e9f64c.
//
// Solidity: function LABEL_TOPIC() constant returns(uint256)
func (_Identity *IdentityCaller) LABELTOPIC(opts *bind.CallOpts) (*big.Int, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "LABEL_TOPIC")
	return results[0].(*big.Int), err
}

// LABELTOPIC is a free data retrieval call binding the contract method 0xb1e9f64c.
//
// Solidity: function LABEL_TOPIC() constant returns(uint256)
func (_Identity *IdentitySession) LABELTOPIC() (*big.Int, error) {
	return _Identity.Contract.LABELTOPIC(&_Identity.CallOpts)
}

// LABELTOPIC is a free data retrieval call binding the contract method 0xb1e9f64c.
//
// Solidity: function LABEL_TOPIC() constant returns(uint256)
func (_Identity *IdentityCallerSession) LABELTOPIC() (*big.Int, error) {
	return _Identity.Contract.LABELTOPIC(&_Identity.CallOpts)
}

// MANAGEMENTKEY is a free data retrieval call binding the contract method 0x058b316c.
//
// Solidity: function MANAGEMENT_KEY() constant returns(uint256)
func (_Identity *IdentityCaller) MANAGEMENTKEY(opts *bind.CallOpts) (*big.Int, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "MANAGEMENT_KEY")
	return results[0].(*big.Int), err
}

// MANAGEMENTKEY is a free data retrieval call binding the contract method 0x058b316c.
//
// Solidity: function MANAGEMENT_KEY() constant returns(uint256)
func (_Identity *IdentitySession) MANAGEMENTKEY() (*big.Int, error) {
	return _Identity.Contract.MANAGEMENTKEY(&_Identity.CallOpts)
}

// MANAGEMENTKEY is a free data retrieval call binding the contract method 0x058b316c.
//
// Solidity: function MANAGEMENT_KEY() constant returns(uint256)
func (_Identity *IdentityCallerSession) MANAGEMENTKEY() (*big.Int, error) {
	return _Identity.Contract.MANAGEMENTKEY(&_Identity.CallOpts)
}

// METAIDTOPIC is a free data retrieval call binding the contract method 0x710ca550.
//
// Solidity: function METAID_TOPIC() constant returns(uint256)
func (_Identity *IdentityCaller) METAIDTOPIC(opts *bind.CallOpts) (*big.Int, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "METAID_TOPIC")
	return results[0].(*big.Int), err
}

// METAIDTOPIC is a free data retrieval call binding the contract method 0x710ca550.
//
// Solidity: function METAID_TOPIC() constant returns(uint256)
func (_Identity *IdentitySession) METAIDTOPIC() (*big.Int, error) {
	return _Identity.Contract.METAIDTOPIC(&_Identity.CallOpts)
}

// METAIDTOPIC is a free data retrieval call binding the contract method 0x710ca550.
//
// Solidity: function METAID_TOPIC() constant returns(uint256)
func (_Identity *IdentityCallerSession) METAIDTOPIC() (*big.Int, error) {
	return _Identity.Contract.METAIDTOPIC(&_Identity.CallOpts)
}

// PROFILETOPIC is a free data retrieval call binding the contract method 0xae628386.
//
// Solidity: function PROFILE_TOPIC() constant returns(uint256)
func (_Identity *IdentityCaller) PROFILETOPIC(opts *bind.CallOpts) (*big.Int, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "PROFILE_TOPIC")
	return results[0].(*big.Int), err
}

// PROFILETOPIC is a free data retrieval call binding the contract method 0xae628386.
//
// Solidity: function PROFILE_TOPIC() constant returns(uint256)
func (_Identity *IdentitySession) PROFILETOPIC() (*big.Int, error) {
	return _Identity.Contract.PROFILETOPIC(&_Identity.CallOpts)
}

// PROFILETOPIC is a free data retrieval call binding the contract method 0xae628386.
//
// Solidity: function PROFILE_TOPIC() constant returns(uint256)
func (_Identity *IdentityCallerSession) PROFILETOPIC() (*big.Int, error) {
	return _Identity.Contract.PROFILETOPIC(&_Identity.CallOpts)
}

// REGISTRYTOPIC is a free data retrieval call binding the contract method 0x0440b43a.
//
// Solidity: function REGISTRY_TOPIC() constant returns(uint256)
func (_Identity *IdentityCaller) REGISTRYTOPIC(opts *bind.CallOpts) (*big.Int, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "REGISTRY_TOPIC")
	return results[0].(*big.Int), err
}

// REGISTRYTOPIC is a free data retrieval call binding the contract method 0x0440b43a.
//
// Solidity: function REGISTRY_TOPIC() constant returns(uint256)
func (_Identity *IdentitySession) REGISTRYTOPIC() (*big.Int, error) {
	return _Identity.Contract.REGISTRYTOPIC(&_Identity.CallOpts)
}

// REGISTRYTOPIC is a free data retrieval call binding the contract method 0x0440b43a.
//
// Solidity: function REGISTRY_TOPIC() constant returns(uint256)
func (_Identity *IdentityCallerSession) REGISTRYTOPIC() (*big.Int, error) {
	return _Identity.Contract.REGISTRYTOPIC(&_Identity.CallOpts)
}

// RESIDENCETOPIC is a free data retrieval call binding the contract method 0xa550f0c7.
//
// Solidity: function RESIDENCE_TOPIC() constant returns(uint256)
func (_Identity *IdentityCaller) RESIDENCETOPIC(opts *bind.CallOpts) (*big.Int, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "RESIDENCE_TOPIC")
	return results[0].(*big.Int), err
}

// RESIDENCETOPIC is a free data retrieval call binding the contract method 0xa550f0c7.
//
// Solidity: function RESIDENCE_TOPIC() constant returns(uint256)
func (_Identity *IdentitySession) RESIDENCETOPIC() (*big.Int, error) {
	return _Identity.Contract.RESIDENCETOPIC(&_Identity.CallOpts)
}

// RESIDENCETOPIC is a free data retrieval call binding the contract method 0xa550f0c7.
//
// Solidity: function RESIDENCE_TOPIC() constant returns(uint256)
func (_Identity *IdentityCallerSession) RESIDENCETOPIC() (*big.Int, error) {
	return _Identity.Contract.RESIDENCETOPIC(&_Identity.CallOpts)
}

// RESTOREKEY is a free data retrieval call binding the contract method 0xb9133d63.
//
// Solidity: function RESTORE_KEY() constant returns(uint256)
func (_Identity *IdentityCaller) RESTOREKEY(opts *bind.CallOpts) (*big.Int, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "RESTORE_KEY")
	return results[0].(*big.Int), err
}

// RESTOREKEY is a free data retrieval call binding the contract method 0xb9133d63.
//
// Solidity: function RESTORE_KEY() constant returns(uint256)
func (_Identity *IdentitySession) RESTOREKEY() (*big.Int, error) {
	return _Identity.Contract.RESTOREKEY(&_Identity.CallOpts)
}

// RESTOREKEY is a free data retrieval call binding the contract method 0xb9133d63.
//
// Solidity: function RESTORE_KEY() constant returns(uint256)
func (_Identity *IdentityCallerSession) RESTOREKEY() (*big.Int, error) {
	return _Identity.Contract.RESTOREKEY(&_Identity.CallOpts)
}

// RSASCHEME is a free data retrieval call binding the contract method 0xf22d08a6.
//
// Solidity: function RSA_SCHEME() constant returns(uint256)
func (_Identity *IdentityCaller) RSASCHEME(opts *bind.CallOpts) (*big.Int, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "RSA_SCHEME")
	return results[0].(*big.Int), err
}

// RSASCHEME is a free data retrieval call binding the contract method 0xf22d08a6.
//
// Solidity: function RSA_SCHEME() constant returns(uint256)
func (_Identity *IdentitySession) RSASCHEME() (*big.Int, error) {
	return _Identity.Contract.RSASCHEME(&_Identity.CallOpts)
}

// RSASCHEME is a free data retrieval call binding the contract method 0xf22d08a6.
//
// Solidity: function RSA_SCHEME() constant returns(uint256)
func (_Identity *IdentityCallerSession) RSASCHEME() (*big.Int, error) {
	return _Identity.Contract.RSASCHEME(&_Identity.CallOpts)
}

// RSATYPE is a free data retrieval call binding the contract method 0x2d32d442.
//
// Solidity: function RSA_TYPE() constant returns(uint256)
func (_Identity *IdentityCaller) RSATYPE(opts *bind.CallOpts) (*big.Int, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "RSA_TYPE")
	return results[0].(*big.Int), err
}

// RSATYPE is a free data retrieval call binding the contract method 0x2d32d442.
//
// Solidity: function RSA_TYPE() constant returns(uint256)
func (_Identity *IdentitySession) RSATYPE() (*big.Int, error) {
	return _Identity.Contract.RSATYPE(&_Identity.CallOpts)
}

// RSATYPE is a free data retrieval call binding the contract method 0x2d32d442.
//
// Solidity: function RSA_TYPE() constant returns(uint256)
func (_Identity *IdentityCallerSession) RSATYPE() (*big.Int, error) {
	return _Identity.Contract.RSATYPE(&_Identity.CallOpts)
}

// ActionThreshold is a free data retrieval call binding the contract method 0x38f4edd4.
//
// Solidity: function actionThreshold() constant returns(uint256)
func (_Identity *IdentityCaller) ActionThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "actionThreshold")
	return results[0].(*big.Int), err
}

// ActionThreshold is a free data retrieval call binding the contract method 0x38f4edd4.
//
// Solidity: function actionThreshold() constant returns(uint256)
func (_Identity *IdentitySession) ActionThreshold() (*big.Int, error) {
	return _Identity.Contract.ActionThreshold(&_Identity.CallOpts)
}

// ActionThreshold is a free data retrieval call binding the contract method 0x38f4edd4.
//
// Solidity: function actionThreshold() constant returns(uint256)
func (_Identity *IdentityCallerSession) ActionThreshold() (*big.Int, error) {
	return _Identity.Contract.ActionThreshold(&_Identity.CallOpts)
}

// AddrToKey is a free data retrieval call binding the contract method 0x63f14c3c.
//
// Solidity: function addrToKey(address addr) constant returns(bytes32)
func (_Identity *IdentityCaller) AddrToKey(opts *bind.CallOpts, addr common.Address) ([32]byte, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "addrToKey", addr)
	return results[0].([32]byte), err
}

// AddrToKey is a free data retrieval call binding the contract method 0x63f14c3c.
//
// Solidity: function addrToKey(address addr) constant returns(bytes32)
func (_Identity *IdentitySession) AddrToKey(addr common.Address) ([32]byte, error) {
	return _Identity.Contract.AddrToKey(&_Identity.CallOpts, addr)
}

// AddrToKey is a free data retrieval call binding the contract method 0x63f14c3c.
//
// Solidity: function addrToKey(address addr) constant returns(bytes32)
func (_Identity *IdentityCallerSession) AddrToKey(addr common.Address) ([32]byte, error) {
	return _Identity.Contract.AddrToKey(&_Identity.CallOpts, addr)
}

// Approved is a free data retrieval call binding the contract method 0x6e4c4311.
//
// Solidity: function approved(uint256 , uint256 ) constant returns(address)
func (_Identity *IdentityCaller) Approved(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "approved", arg0, arg1)
	return results[0].(common.Address), err
}

// Approved is a free data retrieval call binding the contract method 0x6e4c4311.
//
// Solidity: function approved(uint256 , uint256 ) constant returns(address)
func (_Identity *IdentitySession) Approved(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Identity.Contract.Approved(&_Identity.CallOpts, arg0, arg1)
}

// Approved is a free data retrieval call binding the contract method 0x6e4c4311.
//
// Solidity: function approved(uint256 , uint256 ) constant returns(address)
func (_Identity *IdentityCallerSession) Approved(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Identity.Contract.Approved(&_Identity.CallOpts, arg0, arg1)
}

// ClaimToSign is a free data retrieval call binding the contract method 0x1d203be8.
//
// Solidity: function claimToSign(address subject, uint256 topic, bytes data) constant returns(bytes32)
func (_Identity *IdentityCaller) ClaimToSign(opts *bind.CallOpts, subject common.Address, topic *big.Int, data []byte) ([32]byte, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "claimToSign", subject, topic, data)
	return results[0].([32]byte), err
}

// ClaimToSign is a free data retrieval call binding the contract method 0x1d203be8.
//
// Solidity: function claimToSign(address subject, uint256 topic, bytes data) constant returns(bytes32)
func (_Identity *IdentitySession) ClaimToSign(subject common.Address, topic *big.Int, data []byte) ([32]byte, error) {
	return _Identity.Contract.ClaimToSign(&_Identity.CallOpts, subject, topic, data)
}

// ClaimToSign is a free data retrieval call binding the contract method 0x1d203be8.
//
// Solidity: function claimToSign(address subject, uint256 topic, bytes data) constant returns(bytes32)
func (_Identity *IdentityCallerSession) ClaimToSign(subject common.Address, topic *big.Int, data []byte) ([32]byte, error) {
	return _Identity.Contract.ClaimToSign(&_Identity.CallOpts, subject, topic, data)
}

type Method5dccebbaData struct {
	To           common.Address
	Value        *big.Int
	Data         []byte
	NeedsApprove *big.Int
}
// Execution is a free data retrieval call binding the contract method 0x5dccebba.
//
// Solidity: function execution(uint256 ) constant returns(address to, uint256 value, bytes data, uint256 needsApprove)
func (_Identity *IdentityCaller) Execution(opts *bind.CallOpts, arg0 *big.Int) (Method5dccebbaData, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "execution", arg0)
	return results[0].(Method5dccebbaData), err
}

// Execution is a free data retrieval call binding the contract method 0x5dccebba.
//
// Solidity: function execution(uint256 ) constant returns(address to, uint256 value, bytes data, uint256 needsApprove)
func (_Identity *IdentitySession) Execution(arg0 *big.Int) (Method5dccebbaData, error) {
	return _Identity.Contract.Execution(&_Identity.CallOpts, arg0)
}

// Execution is a free data retrieval call binding the contract method 0x5dccebba.
//
// Solidity: function execution(uint256 ) constant returns(address to, uint256 value, bytes data, uint256 needsApprove)
func (_Identity *IdentityCallerSession) Execution(arg0 *big.Int) (Method5dccebbaData, error) {
	return _Identity.Contract.Execution(&_Identity.CallOpts, arg0)
}

type Methodc9100bcbData struct {
	Topic     *big.Int
	Scheme    *big.Int
	Issuer    common.Address
	Signature []byte
	Data      []byte
	Uri       string
}
// GetClaim is a free data retrieval call binding the contract method 0xc9100bcb.
//
// Solidity: function getClaim(bytes32 _claimId) constant returns(uint256 topic, uint256 scheme, address issuer, bytes signature, bytes data, string uri)
func (_Identity *IdentityCaller) GetClaim(opts *bind.CallOpts, _claimId [32]byte) (Methodc9100bcbData, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "getClaim", _claimId)
	return results[0].(Methodc9100bcbData), err
}

// GetClaim is a free data retrieval call binding the contract method 0xc9100bcb.
//
// Solidity: function getClaim(bytes32 _claimId) constant returns(uint256 topic, uint256 scheme, address issuer, bytes signature, bytes data, string uri)
func (_Identity *IdentitySession) GetClaim(_claimId [32]byte) (Methodc9100bcbData, error) {
	return _Identity.Contract.GetClaim(&_Identity.CallOpts, _claimId)
}

// GetClaim is a free data retrieval call binding the contract method 0xc9100bcb.
//
// Solidity: function getClaim(bytes32 _claimId) constant returns(uint256 topic, uint256 scheme, address issuer, bytes signature, bytes data, string uri)
func (_Identity *IdentityCallerSession) GetClaim(_claimId [32]byte) (Methodc9100bcbData, error) {
	return _Identity.Contract.GetClaim(&_Identity.CallOpts, _claimId)
}

// GetClaimId is a free data retrieval call binding the contract method 0x190db862.
//
// Solidity: function getClaimId(address issuer, uint256 topic) constant returns(bytes32)
func (_Identity *IdentityCaller) GetClaimId(opts *bind.CallOpts, issuer common.Address, topic *big.Int) ([32]byte, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "getClaimId", issuer, topic)
	return results[0].([32]byte), err
}

// GetClaimId is a free data retrieval call binding the contract method 0x190db862.
//
// Solidity: function getClaimId(address issuer, uint256 topic) constant returns(bytes32)
func (_Identity *IdentitySession) GetClaimId(issuer common.Address, topic *big.Int) ([32]byte, error) {
	return _Identity.Contract.GetClaimId(&_Identity.CallOpts, issuer, topic)
}

// GetClaimId is a free data retrieval call binding the contract method 0x190db862.
//
// Solidity: function getClaimId(address issuer, uint256 topic) constant returns(bytes32)
func (_Identity *IdentityCallerSession) GetClaimId(issuer common.Address, topic *big.Int) ([32]byte, error) {
	return _Identity.Contract.GetClaimId(&_Identity.CallOpts, issuer, topic)
}

// GetClaimIdsByType is a free data retrieval call binding the contract method 0x262b54f5.
//
// Solidity: function getClaimIdsByType(uint256 _topic) constant returns(bytes32[] claimIds)
func (_Identity *IdentityCaller) GetClaimIdsByType(opts *bind.CallOpts, _topic *big.Int) ([][32]byte, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "getClaimIdsByType", _topic)
	return results[0].([][32]byte), err
}

// GetClaimIdsByType is a free data retrieval call binding the contract method 0x262b54f5.
//
// Solidity: function getClaimIdsByType(uint256 _topic) constant returns(bytes32[] claimIds)
func (_Identity *IdentitySession) GetClaimIdsByType(_topic *big.Int) ([][32]byte, error) {
	return _Identity.Contract.GetClaimIdsByType(&_Identity.CallOpts, _topic)
}

// GetClaimIdsByType is a free data retrieval call binding the contract method 0x262b54f5.
//
// Solidity: function getClaimIdsByType(uint256 _topic) constant returns(bytes32[] claimIds)
func (_Identity *IdentityCallerSession) GetClaimIdsByType(_topic *big.Int) ([][32]byte, error) {
	return _Identity.Contract.GetClaimIdsByType(&_Identity.CallOpts, _topic)
}

// GetExecutionId is a free data retrieval call binding the contract method 0xf7444137.
//
// Solidity: function getExecutionId(address self, address _to, uint256 _value, bytes _data, uint256 _nonce) constant returns(uint256)
func (_Identity *IdentityCaller) GetExecutionId(opts *bind.CallOpts, self common.Address, _to common.Address, _value *big.Int, _data []byte, _nonce *big.Int) (*big.Int, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "getExecutionId", self, _to, _value, _data, _nonce)
	return results[0].(*big.Int), err
}

// GetExecutionId is a free data retrieval call binding the contract method 0xf7444137.
//
// Solidity: function getExecutionId(address self, address _to, uint256 _value, bytes _data, uint256 _nonce) constant returns(uint256)
func (_Identity *IdentitySession) GetExecutionId(self common.Address, _to common.Address, _value *big.Int, _data []byte, _nonce *big.Int) (*big.Int, error) {
	return _Identity.Contract.GetExecutionId(&_Identity.CallOpts, self, _to, _value, _data, _nonce)
}

// GetExecutionId is a free data retrieval call binding the contract method 0xf7444137.
//
// Solidity: function getExecutionId(address self, address _to, uint256 _value, bytes _data, uint256 _nonce) constant returns(uint256)
func (_Identity *IdentityCallerSession) GetExecutionId(self common.Address, _to common.Address, _value *big.Int, _data []byte, _nonce *big.Int) (*big.Int, error) {
	return _Identity.Contract.GetExecutionId(&_Identity.CallOpts, self, _to, _value, _data, _nonce)
}

// GetFunctionSignature is a free data retrieval call binding the contract method 0xc32b3518.
//
// Solidity: function getFunctionSignature(bytes b) constant returns(bytes4)
func (_Identity *IdentityCaller) GetFunctionSignature(opts *bind.CallOpts, b []byte) ([4]byte, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "getFunctionSignature", b)
	return results[0].([4]byte), err
}

// GetFunctionSignature is a free data retrieval call binding the contract method 0xc32b3518.
//
// Solidity: function getFunctionSignature(bytes b) constant returns(bytes4)
func (_Identity *IdentitySession) GetFunctionSignature(b []byte) ([4]byte, error) {
	return _Identity.Contract.GetFunctionSignature(&_Identity.CallOpts, b)
}

// GetFunctionSignature is a free data retrieval call binding the contract method 0xc32b3518.
//
// Solidity: function getFunctionSignature(bytes b) constant returns(bytes4)
func (_Identity *IdentityCallerSession) GetFunctionSignature(b []byte) ([4]byte, error) {
	return _Identity.Contract.GetFunctionSignature(&_Identity.CallOpts, b)
}

// GetKey is a free data retrieval call binding the contract method 0x12aaac70.
//
// Solidity: function getKey(bytes32 _key) constant returns(uint256[] purposes, uint256 keyType, bytes32 key)
func (_Identity *IdentityCaller) GetKey(opts *bind.CallOpts, _key [32]byte) (metaIDKey, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "getKey", _key)
	return results[0].(metaIDKey), err
}

// GetKey is a free data retrieval call binding the contract method 0x12aaac70.
//
// Solidity: function getKey(bytes32 _key) constant returns(uint256[] purposes, uint256 keyType, bytes32 key)
func (_Identity *IdentitySession) GetKey(_key [32]byte) (metaIDKey, error) {
	return _Identity.Contract.GetKey(&_Identity.CallOpts, _key)
}

// GetKey is a free data retrieval call binding the contract method 0x12aaac70.
//
// Solidity: function getKey(bytes32 _key) constant returns(uint256[] purposes, uint256 keyType, bytes32 key)
func (_Identity *IdentityCallerSession) GetKey(_key [32]byte) (metaIDKey, error) {
	return _Identity.Contract.GetKey(&_Identity.CallOpts, _key)
}

// GetKeysByPurpose is a free data retrieval call binding the contract method 0x9010f726.
//
// Solidity: function getKeysByPurpose(uint256 _purpose) constant returns(bytes32[] keys)
func (_Identity *IdentityCaller) GetKeysByPurpose(opts *bind.CallOpts, _purpose *big.Int) ([][32]byte, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "getKeysByPurpose", _purpose)
	return results[0].([][32]byte), err
}

// GetKeysByPurpose is a free data retrieval call binding the contract method 0x9010f726.
//
// Solidity: function getKeysByPurpose(uint256 _purpose) constant returns(bytes32[] keys)
func (_Identity *IdentitySession) GetKeysByPurpose(_purpose *big.Int) ([][32]byte, error) {
	return _Identity.Contract.GetKeysByPurpose(&_Identity.CallOpts, _purpose)
}

// GetKeysByPurpose is a free data retrieval call binding the contract method 0x9010f726.
//
// Solidity: function getKeysByPurpose(uint256 _purpose) constant returns(bytes32[] keys)
func (_Identity *IdentityCallerSession) GetKeysByPurpose(_purpose *big.Int) ([][32]byte, error) {
	return _Identity.Contract.GetKeysByPurpose(&_Identity.CallOpts, _purpose)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() constant returns(uint256)
func (_Identity *IdentityCaller) GetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "getNonce")
	return results[0].(*big.Int), err
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() constant returns(uint256)
func (_Identity *IdentitySession) GetNonce() (*big.Int, error) {
	return _Identity.Contract.GetNonce(&_Identity.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() constant returns(uint256)
func (_Identity *IdentityCallerSession) GetNonce() (*big.Int, error) {
	return _Identity.Contract.GetNonce(&_Identity.CallOpts)
}

// GetNumClaims is a free data retrieval call binding the contract method 0x19902986.
//
// Solidity: function getNumClaims() constant returns(uint256 num)
func (_Identity *IdentityCaller) GetNumClaims(opts *bind.CallOpts) (*big.Int, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "getNumClaims")
	return results[0].(*big.Int), err
}

// GetNumClaims is a free data retrieval call binding the contract method 0x19902986.
//
// Solidity: function getNumClaims() constant returns(uint256 num)
func (_Identity *IdentitySession) GetNumClaims() (*big.Int, error) {
	return _Identity.Contract.GetNumClaims(&_Identity.CallOpts)
}

// GetNumClaims is a free data retrieval call binding the contract method 0x19902986.
//
// Solidity: function getNumClaims() constant returns(uint256 num)
func (_Identity *IdentityCallerSession) GetNumClaims() (*big.Int, error) {
	return _Identity.Contract.GetNumClaims(&_Identity.CallOpts)
}

// GetSignatureAddress is a free data retrieval call binding the contract method 0x3b8a12c8.
//
// Solidity: function getSignatureAddress(bytes32 toSign, bytes signature) constant returns(address)
func (_Identity *IdentityCaller) GetSignatureAddress(opts *bind.CallOpts, toSign [32]byte, signature []byte) (common.Address, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "getSignatureAddress", toSign, signature)
	return results[0].(common.Address), err
}

// GetSignatureAddress is a free data retrieval call binding the contract method 0x3b8a12c8.
//
// Solidity: function getSignatureAddress(bytes32 toSign, bytes signature) constant returns(address)
func (_Identity *IdentitySession) GetSignatureAddress(toSign [32]byte, signature []byte) (common.Address, error) {
	return _Identity.Contract.GetSignatureAddress(&_Identity.CallOpts, toSign, signature)
}

// GetSignatureAddress is a free data retrieval call binding the contract method 0x3b8a12c8.
//
// Solidity: function getSignatureAddress(bytes32 toSign, bytes signature) constant returns(address)
func (_Identity *IdentityCallerSession) GetSignatureAddress(toSign [32]byte, signature []byte) (common.Address, error) {
	return _Identity.Contract.GetSignatureAddress(&_Identity.CallOpts, toSign, signature)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x2e7700f0.
//
// Solidity: function getTransactionCount() constant returns(uint256)
func (_Identity *IdentityCaller) GetTransactionCount(opts *bind.CallOpts) (*big.Int, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "getTransactionCount")
	return results[0].(*big.Int), err
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x2e7700f0.
//
// Solidity: function getTransactionCount() constant returns(uint256)
func (_Identity *IdentitySession) GetTransactionCount() (*big.Int, error) {
	return _Identity.Contract.GetTransactionCount(&_Identity.CallOpts)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x2e7700f0.
//
// Solidity: function getTransactionCount() constant returns(uint256)
func (_Identity *IdentityCallerSession) GetTransactionCount() (*big.Int, error) {
	return _Identity.Contract.GetTransactionCount(&_Identity.CallOpts)
}

// IsClaimExist is a free data retrieval call binding the contract method 0x237434f0.
//
// Solidity: function isClaimExist(bytes32 _claimId) constant returns(bool)
func (_Identity *IdentityCaller) IsClaimExist(opts *bind.CallOpts, _claimId [32]byte) (bool, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "isClaimExist", _claimId)
	return results[0].(bool), err
}

// IsClaimExist is a free data retrieval call binding the contract method 0x237434f0.
//
// Solidity: function isClaimExist(bytes32 _claimId) constant returns(bool)
func (_Identity *IdentitySession) IsClaimExist(_claimId [32]byte) (bool, error) {
	return _Identity.Contract.IsClaimExist(&_Identity.CallOpts, _claimId)
}

// IsClaimExist is a free data retrieval call binding the contract method 0x237434f0.
//
// Solidity: function isClaimExist(bytes32 _claimId) constant returns(bool)
func (_Identity *IdentityCallerSession) IsClaimExist(_claimId [32]byte) (bool, error) {
	return _Identity.Contract.IsClaimExist(&_Identity.CallOpts, _claimId)
}

// KeyCanExecute is a free data retrieval call binding the contract method 0x765b3042.
//
// Solidity: function keyCanExecute(bytes32 _key, address _to, bytes4 _func) constant returns(bool executable)
func (_Identity *IdentityCaller) KeyCanExecute(opts *bind.CallOpts, _key [32]byte, _to common.Address, _func [4]byte) (bool, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "keyCanExecute", _key, _to, _func)
	return results[0].(bool), err
}

// KeyCanExecute is a free data retrieval call binding the contract method 0x765b3042.
//
// Solidity: function keyCanExecute(bytes32 _key, address _to, bytes4 _func) constant returns(bool executable)
func (_Identity *IdentitySession) KeyCanExecute(_key [32]byte, _to common.Address, _func [4]byte) (bool, error) {
	return _Identity.Contract.KeyCanExecute(&_Identity.CallOpts, _key, _to, _func)
}

// KeyCanExecute is a free data retrieval call binding the contract method 0x765b3042.
//
// Solidity: function keyCanExecute(bytes32 _key, address _to, bytes4 _func) constant returns(bool executable)
func (_Identity *IdentityCallerSession) KeyCanExecute(_key [32]byte, _to common.Address, _func [4]byte) (bool, error) {
	return _Identity.Contract.KeyCanExecute(&_Identity.CallOpts, _key, _to, _func)
}

// KeyHasPurpose is a free data retrieval call binding the contract method 0xd202158d.
//
// Solidity: function keyHasPurpose(bytes32 _key, uint256 purpose) constant returns(bool exists)
func (_Identity *IdentityCaller) KeyHasPurpose(opts *bind.CallOpts, _key [32]byte, purpose *big.Int) (bool, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "keyHasPurpose", _key, purpose)
	return results[0].(bool), err
}

// KeyHasPurpose is a free data retrieval call binding the contract method 0xd202158d.
//
// Solidity: function keyHasPurpose(bytes32 _key, uint256 purpose) constant returns(bool exists)
func (_Identity *IdentitySession) KeyHasPurpose(_key [32]byte, purpose *big.Int) (bool, error) {
	return _Identity.Contract.KeyHasPurpose(&_Identity.CallOpts, _key, purpose)
}

// KeyHasPurpose is a free data retrieval call binding the contract method 0xd202158d.
//
// Solidity: function keyHasPurpose(bytes32 _key, uint256 purpose) constant returns(bool exists)
func (_Identity *IdentityCallerSession) KeyHasPurpose(_key [32]byte, purpose *big.Int) (bool, error) {
	return _Identity.Contract.KeyHasPurpose(&_Identity.CallOpts, _key, purpose)
}

// ManagementThreshold is a free data retrieval call binding the contract method 0xaa0a5142.
//
// Solidity: function managementThreshold() constant returns(uint256)
func (_Identity *IdentityCaller) ManagementThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "managementThreshold")
	return results[0].(*big.Int), err
}

// ManagementThreshold is a free data retrieval call binding the contract method 0xaa0a5142.
//
// Solidity: function managementThreshold() constant returns(uint256)
func (_Identity *IdentitySession) ManagementThreshold() (*big.Int, error) {
	return _Identity.Contract.ManagementThreshold(&_Identity.CallOpts)
}

// ManagementThreshold is a free data retrieval call binding the contract method 0xaa0a5142.
//
// Solidity: function managementThreshold() constant returns(uint256)
func (_Identity *IdentityCallerSession) ManagementThreshold() (*big.Int, error) {
	return _Identity.Contract.ManagementThreshold(&_Identity.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() constant returns(uint256)
func (_Identity *IdentityCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "nonce")
	return results[0].(*big.Int), err
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() constant returns(uint256)
func (_Identity *IdentitySession) Nonce() (*big.Int, error) {
	return _Identity.Contract.Nonce(&_Identity.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() constant returns(uint256)
func (_Identity *IdentityCallerSession) Nonce() (*big.Int, error) {
	return _Identity.Contract.Nonce(&_Identity.CallOpts)
}

// NumClaims is a free data retrieval call binding the contract method 0xfc0fc849.
//
// Solidity: function numClaims() constant returns(uint256)
func (_Identity *IdentityCaller) NumClaims(opts *bind.CallOpts) (*big.Int, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "numClaims")
	return results[0].(*big.Int), err
}

// NumClaims is a free data retrieval call binding the contract method 0xfc0fc849.
//
// Solidity: function numClaims() constant returns(uint256)
func (_Identity *IdentitySession) NumClaims() (*big.Int, error) {
	return _Identity.Contract.NumClaims(&_Identity.CallOpts)
}

// NumClaims is a free data retrieval call binding the contract method 0xfc0fc849.
//
// Solidity: function numClaims() constant returns(uint256)
func (_Identity *IdentityCallerSession) NumClaims() (*big.Int, error) {
	return _Identity.Contract.NumClaims(&_Identity.CallOpts)
}

// NumKeys is a free data retrieval call binding the contract method 0xc9d24ecc.
//
// Solidity: function numKeys() constant returns(uint256)
func (_Identity *IdentityCaller) NumKeys(opts *bind.CallOpts) (*big.Int, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "numKeys")
	return results[0].(*big.Int), err
}

// NumKeys is a free data retrieval call binding the contract method 0xc9d24ecc.
//
// Solidity: function numKeys() constant returns(uint256)
func (_Identity *IdentitySession) NumKeys() (*big.Int, error) {
	return _Identity.Contract.NumKeys(&_Identity.CallOpts)
}

// NumKeys is a free data retrieval call binding the contract method 0xc9d24ecc.
//
// Solidity: function numKeys() constant returns(uint256)
func (_Identity *IdentityCallerSession) NumKeys() (*big.Int, error) {
	return _Identity.Contract.NumKeys(&_Identity.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Identity *IdentityCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "paused")
	return results[0].(bool), err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Identity *IdentitySession) Paused() (bool, error) {
	return _Identity.Contract.Paused(&_Identity.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Identity *IdentityCallerSession) Paused() (bool, error) {
	return _Identity.Contract.Paused(&_Identity.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) constant returns(bool)
func (_Identity *IdentityCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var results []interface{}
	err := _Identity.contract.Call(opts, &results, "supportsInterface", interfaceID)
	return results[0].(bool), err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) constant returns(bool)
func (_Identity *IdentitySession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _Identity.Contract.SupportsInterface(&_Identity.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) constant returns(bool)
func (_Identity *IdentityCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _Identity.Contract.SupportsInterface(&_Identity.CallOpts, interfaceID)
}

// AddClaim is a paid mutator transaction binding the contract method 0xb1a34e0d.
//
// Solidity: function addClaim(uint256 _topic, uint256 _scheme, address issuer, bytes _signature, bytes _data, string _uri) returns(uint256 claimRequestId)
func (_Identity *IdentityTransactor) AddClaim(opts *bind.TransactOpts, _topic *big.Int, _scheme *big.Int, issuer common.Address, _signature []byte, _data []byte, _uri string) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "addClaim", _topic, _scheme, issuer, _signature, _data, _uri)
}

// AddClaim is a paid mutator transaction binding the contract method 0xb1a34e0d.
//
// Solidity: function addClaim(uint256 _topic, uint256 _scheme, address issuer, bytes _signature, bytes _data, string _uri) returns(uint256 claimRequestId)
func (_Identity *IdentitySession) AddClaim(_topic *big.Int, _scheme *big.Int, issuer common.Address, _signature []byte, _data []byte, _uri string) (*types.Transaction, error) {
	return _Identity.Contract.AddClaim(&_Identity.TransactOpts, _topic, _scheme, issuer, _signature, _data, _uri)
}

// AddClaim is a paid mutator transaction binding the contract method 0xb1a34e0d.
//
// Solidity: function addClaim(uint256 _topic, uint256 _scheme, address issuer, bytes _signature, bytes _data, string _uri) returns(uint256 claimRequestId)
func (_Identity *IdentityTransactorSession) AddClaim(_topic *big.Int, _scheme *big.Int, issuer common.Address, _signature []byte, _data []byte, _uri string) (*types.Transaction, error) {
	return _Identity.Contract.AddClaim(&_Identity.TransactOpts, _topic, _scheme, issuer, _signature, _data, _uri)
}

// AddClaimByProxy is a paid mutator transaction binding the contract method 0xe0610ba3.
//
// Solidity: function addClaimByProxy(uint256 _topic, uint256 _scheme, address issuer, bytes _signature, bytes _data, string _uri, bytes _idSignature) returns(bool success)
func (_Identity *IdentityTransactor) AddClaimByProxy(opts *bind.TransactOpts, _topic *big.Int, _scheme *big.Int, issuer common.Address, _signature []byte, _data []byte, _uri string, _idSignature []byte) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "addClaimByProxy", _topic, _scheme, issuer, _signature, _data, _uri, _idSignature)
}

// AddClaimByProxy is a paid mutator transaction binding the contract method 0xe0610ba3.
//
// Solidity: function addClaimByProxy(uint256 _topic, uint256 _scheme, address issuer, bytes _signature, bytes _data, string _uri, bytes _idSignature) returns(bool success)
func (_Identity *IdentitySession) AddClaimByProxy(_topic *big.Int, _scheme *big.Int, issuer common.Address, _signature []byte, _data []byte, _uri string, _idSignature []byte) (*types.Transaction, error) {
	return _Identity.Contract.AddClaimByProxy(&_Identity.TransactOpts, _topic, _scheme, issuer, _signature, _data, _uri, _idSignature)
}

// AddClaimByProxy is a paid mutator transaction binding the contract method 0xe0610ba3.
//
// Solidity: function addClaimByProxy(uint256 _topic, uint256 _scheme, address issuer, bytes _signature, bytes _data, string _uri, bytes _idSignature) returns(bool success)
func (_Identity *IdentityTransactorSession) AddClaimByProxy(_topic *big.Int, _scheme *big.Int, issuer common.Address, _signature []byte, _data []byte, _uri string, _idSignature []byte) (*types.Transaction, error) {
	return _Identity.Contract.AddClaimByProxy(&_Identity.TransactOpts, _topic, _scheme, issuer, _signature, _data, _uri, _idSignature)
}

// AddKey is a paid mutator transaction binding the contract method 0x1d381240.
//
// Solidity: function addKey(bytes32 _key, uint256 _purpose, uint256 _keyType) returns(bool success)
func (_Identity *IdentityTransactor) AddKey(opts *bind.TransactOpts, _key [32]byte, _purpose *big.Int, _keyType *big.Int) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "addKey", _key, _purpose, _keyType)
}

// AddKey is a paid mutator transaction binding the contract method 0x1d381240.
//
// Solidity: function addKey(bytes32 _key, uint256 _purpose, uint256 _keyType) returns(bool success)
func (_Identity *IdentitySession) AddKey(_key [32]byte, _purpose *big.Int, _keyType *big.Int) (*types.Transaction, error) {
	return _Identity.Contract.AddKey(&_Identity.TransactOpts, _key, _purpose, _keyType)
}

// AddKey is a paid mutator transaction binding the contract method 0x1d381240.
//
// Solidity: function addKey(bytes32 _key, uint256 _purpose, uint256 _keyType) returns(bool success)
func (_Identity *IdentityTransactorSession) AddKey(_key [32]byte, _purpose *big.Int, _keyType *big.Int) (*types.Transaction, error) {
	return _Identity.Contract.AddKey(&_Identity.TransactOpts, _key, _purpose, _keyType)
}

// Approve is a paid mutator transaction binding the contract method 0x747442d3.
//
// Solidity: function approve(uint256 _id, bool _approve) returns(bool success)
func (_Identity *IdentityTransactor) Approve(opts *bind.TransactOpts, _id *big.Int, _approve bool) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "approve", _id, _approve)
}

// Approve is a paid mutator transaction binding the contract method 0x747442d3.
//
// Solidity: function approve(uint256 _id, bool _approve) returns(bool success)
func (_Identity *IdentitySession) Approve(_id *big.Int, _approve bool) (*types.Transaction, error) {
	return _Identity.Contract.Approve(&_Identity.TransactOpts, _id, _approve)
}

// Approve is a paid mutator transaction binding the contract method 0x747442d3.
//
// Solidity: function approve(uint256 _id, bool _approve) returns(bool success)
func (_Identity *IdentityTransactorSession) Approve(_id *big.Int, _approve bool) (*types.Transaction, error) {
	return _Identity.Contract.Approve(&_Identity.TransactOpts, _id, _approve)
}

// ChangeActionThreshold is a paid mutator transaction binding the contract method 0xe99896b8.
//
// Solidity: function changeActionThreshold(uint256 threshold) returns()
func (_Identity *IdentityTransactor) ChangeActionThreshold(opts *bind.TransactOpts, threshold *big.Int) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "changeActionThreshold", threshold)
}

// ChangeActionThreshold is a paid mutator transaction binding the contract method 0xe99896b8.
//
// Solidity: function changeActionThreshold(uint256 threshold) returns()
func (_Identity *IdentitySession) ChangeActionThreshold(threshold *big.Int) (*types.Transaction, error) {
	return _Identity.Contract.ChangeActionThreshold(&_Identity.TransactOpts, threshold)
}

// ChangeActionThreshold is a paid mutator transaction binding the contract method 0xe99896b8.
//
// Solidity: function changeActionThreshold(uint256 threshold) returns()
func (_Identity *IdentityTransactorSession) ChangeActionThreshold(threshold *big.Int) (*types.Transaction, error) {
	return _Identity.Contract.ChangeActionThreshold(&_Identity.TransactOpts, threshold)
}

// ChangeManagementThreshold is a paid mutator transaction binding the contract method 0xccfe5868.
//
// Solidity: function changeManagementThreshold(uint256 threshold) returns()
func (_Identity *IdentityTransactor) ChangeManagementThreshold(opts *bind.TransactOpts, threshold *big.Int) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "changeManagementThreshold", threshold)
}

// ChangeManagementThreshold is a paid mutator transaction binding the contract method 0xccfe5868.
//
// Solidity: function changeManagementThreshold(uint256 threshold) returns()
func (_Identity *IdentitySession) ChangeManagementThreshold(threshold *big.Int) (*types.Transaction, error) {
	return _Identity.Contract.ChangeManagementThreshold(&_Identity.TransactOpts, threshold)
}

// ChangeManagementThreshold is a paid mutator transaction binding the contract method 0xccfe5868.
//
// Solidity: function changeManagementThreshold(uint256 threshold) returns()
func (_Identity *IdentityTransactorSession) ChangeManagementThreshold(threshold *big.Int) (*types.Transaction, error) {
	return _Identity.Contract.ChangeManagementThreshold(&_Identity.TransactOpts, threshold)
}

// DelegatedApprove is a paid mutator transaction binding the contract method 0xe804c01b.
//
// Solidity: function delegatedApprove(uint256 _id, bool _approve, uint256 _nonce, bytes _sig) returns(bool success)
func (_Identity *IdentityTransactor) DelegatedApprove(opts *bind.TransactOpts, _id *big.Int, _approve bool, _nonce *big.Int, _sig []byte) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "delegatedApprove", _id, _approve, _nonce, _sig)
}

// DelegatedApprove is a paid mutator transaction binding the contract method 0xe804c01b.
//
// Solidity: function delegatedApprove(uint256 _id, bool _approve, uint256 _nonce, bytes _sig) returns(bool success)
func (_Identity *IdentitySession) DelegatedApprove(_id *big.Int, _approve bool, _nonce *big.Int, _sig []byte) (*types.Transaction, error) {
	return _Identity.Contract.DelegatedApprove(&_Identity.TransactOpts, _id, _approve, _nonce, _sig)
}

// DelegatedApprove is a paid mutator transaction binding the contract method 0xe804c01b.
//
// Solidity: function delegatedApprove(uint256 _id, bool _approve, uint256 _nonce, bytes _sig) returns(bool success)
func (_Identity *IdentityTransactorSession) DelegatedApprove(_id *big.Int, _approve bool, _nonce *big.Int, _sig []byte) (*types.Transaction, error) {
	return _Identity.Contract.DelegatedApprove(&_Identity.TransactOpts, _id, _approve, _nonce, _sig)
}

// DelegatedExecute is a paid mutator transaction binding the contract method 0x4da34c2c.
//
// Solidity: function delegatedExecute(address _to, uint256 _value, bytes _data, uint256 _nonce, bytes _sig) returns(uint256 executionId)
func (_Identity *IdentityTransactor) DelegatedExecute(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _data []byte, _nonce *big.Int, _sig []byte) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "delegatedExecute", _to, _value, _data, _nonce, _sig)
}

// DelegatedExecute is a paid mutator transaction binding the contract method 0x4da34c2c.
//
// Solidity: function delegatedExecute(address _to, uint256 _value, bytes _data, uint256 _nonce, bytes _sig) returns(uint256 executionId)
func (_Identity *IdentitySession) DelegatedExecute(_to common.Address, _value *big.Int, _data []byte, _nonce *big.Int, _sig []byte) (*types.Transaction, error) {
	return _Identity.Contract.DelegatedExecute(&_Identity.TransactOpts, _to, _value, _data, _nonce, _sig)
}

// DelegatedExecute is a paid mutator transaction binding the contract method 0x4da34c2c.
//
// Solidity: function delegatedExecute(address _to, uint256 _value, bytes _data, uint256 _nonce, bytes _sig) returns(uint256 executionId)
func (_Identity *IdentityTransactorSession) DelegatedExecute(_to common.Address, _value *big.Int, _data []byte, _nonce *big.Int, _sig []byte) (*types.Transaction, error) {
	return _Identity.Contract.DelegatedExecute(&_Identity.TransactOpts, _to, _value, _data, _nonce, _sig)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(address _recipient) returns()
func (_Identity *IdentityTransactor) DestroyAndSend(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "destroyAndSend", _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(address _recipient) returns()
func (_Identity *IdentitySession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _Identity.Contract.DestroyAndSend(&_Identity.TransactOpts, _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(address _recipient) returns()
func (_Identity *IdentityTransactorSession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _Identity.Contract.DestroyAndSend(&_Identity.TransactOpts, _recipient)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address _to, uint256 _value, bytes _data) returns(uint256 executionId)
func (_Identity *IdentityTransactor) Execute(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "execute", _to, _value, _data)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address _to, uint256 _value, bytes _data) returns(uint256 executionId)
func (_Identity *IdentitySession) Execute(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Identity.Contract.Execute(&_Identity.TransactOpts, _to, _value, _data)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address _to, uint256 _value, bytes _data) returns(uint256 executionId)
func (_Identity *IdentityTransactorSession) Execute(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Identity.Contract.Execute(&_Identity.TransactOpts, _to, _value, _data)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Identity *IdentityTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Identity *IdentitySession) Pause() (*types.Transaction, error) {
	return _Identity.Contract.Pause(&_Identity.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Identity *IdentityTransactorSession) Pause() (*types.Transaction, error) {
	return _Identity.Contract.Pause(&_Identity.TransactOpts)
}

// RefreshClaim is a paid mutator transaction binding the contract method 0x5d7bc3fc.
//
// Solidity: function refreshClaim(bytes32 _claimId) returns(bool)
func (_Identity *IdentityTransactor) RefreshClaim(opts *bind.TransactOpts, _claimId [32]byte) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "refreshClaim", _claimId)
}

// RefreshClaim is a paid mutator transaction binding the contract method 0x5d7bc3fc.
//
// Solidity: function refreshClaim(bytes32 _claimId) returns(bool)
func (_Identity *IdentitySession) RefreshClaim(_claimId [32]byte) (*types.Transaction, error) {
	return _Identity.Contract.RefreshClaim(&_Identity.TransactOpts, _claimId)
}

// RefreshClaim is a paid mutator transaction binding the contract method 0x5d7bc3fc.
//
// Solidity: function refreshClaim(bytes32 _claimId) returns(bool)
func (_Identity *IdentityTransactorSession) RefreshClaim(_claimId [32]byte) (*types.Transaction, error) {
	return _Identity.Contract.RefreshClaim(&_Identity.TransactOpts, _claimId)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0x4eee424a.
//
// Solidity: function removeClaim(bytes32 _claimId) returns(bool success)
func (_Identity *IdentityTransactor) RemoveClaim(opts *bind.TransactOpts, _claimId [32]byte) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "removeClaim", _claimId)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0x4eee424a.
//
// Solidity: function removeClaim(bytes32 _claimId) returns(bool success)
func (_Identity *IdentitySession) RemoveClaim(_claimId [32]byte) (*types.Transaction, error) {
	return _Identity.Contract.RemoveClaim(&_Identity.TransactOpts, _claimId)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0x4eee424a.
//
// Solidity: function removeClaim(bytes32 _claimId) returns(bool success)
func (_Identity *IdentityTransactorSession) RemoveClaim(_claimId [32]byte) (*types.Transaction, error) {
	return _Identity.Contract.RemoveClaim(&_Identity.TransactOpts, _claimId)
}

// RemoveKey is a paid mutator transaction binding the contract method 0x53d413c5.
//
// Solidity: function removeKey(bytes32 _key, uint256 _purpose) returns(bool success)
func (_Identity *IdentityTransactor) RemoveKey(opts *bind.TransactOpts, _key [32]byte, _purpose *big.Int) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "removeKey", _key, _purpose)
}

// RemoveKey is a paid mutator transaction binding the contract method 0x53d413c5.
//
// Solidity: function removeKey(bytes32 _key, uint256 _purpose) returns(bool success)
func (_Identity *IdentitySession) RemoveKey(_key [32]byte, _purpose *big.Int) (*types.Transaction, error) {
	return _Identity.Contract.RemoveKey(&_Identity.TransactOpts, _key, _purpose)
}

// RemoveKey is a paid mutator transaction binding the contract method 0x53d413c5.
//
// Solidity: function removeKey(bytes32 _key, uint256 _purpose) returns(bool success)
func (_Identity *IdentityTransactorSession) RemoveKey(_key [32]byte, _purpose *big.Int) (*types.Transaction, error) {
	return _Identity.Contract.RemoveKey(&_Identity.TransactOpts, _key, _purpose)
}

// SetFunc is a paid mutator transaction binding the contract method 0x724a4b3b.
//
// Solidity: function setFunc(bytes32 _key, address _to, bytes4 _func, bool _executable) returns(bool success)
func (_Identity *IdentityTransactor) SetFunc(opts *bind.TransactOpts, _key [32]byte, _to common.Address, _func [4]byte, _executable bool) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "setFunc", _key, _to, _func, _executable)
}

// SetFunc is a paid mutator transaction binding the contract method 0x724a4b3b.
//
// Solidity: function setFunc(bytes32 _key, address _to, bytes4 _func, bool _executable) returns(bool success)
func (_Identity *IdentitySession) SetFunc(_key [32]byte, _to common.Address, _func [4]byte, _executable bool) (*types.Transaction, error) {
	return _Identity.Contract.SetFunc(&_Identity.TransactOpts, _key, _to, _func, _executable)
}

// SetFunc is a paid mutator transaction binding the contract method 0x724a4b3b.
//
// Solidity: function setFunc(bytes32 _key, address _to, bytes4 _func, bool _executable) returns(bool success)
func (_Identity *IdentityTransactorSession) SetFunc(_key [32]byte, _to common.Address, _func [4]byte, _executable bool) (*types.Transaction, error) {
	return _Identity.Contract.SetFunc(&_Identity.TransactOpts, _key, _to, _func, _executable)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Identity *IdentityTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Identity *IdentitySession) Unpause() (*types.Transaction, error) {
	return _Identity.Contract.Unpause(&_Identity.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Identity *IdentityTransactorSession) Unpause() (*types.Transaction, error) {
	return _Identity.Contract.Unpause(&_Identity.TransactOpts)
}

// IdentityApprovedIterator is returned from FilterApproved and is used to iterate over the raw logs and unpacked data for Approved events raised by the Identity contract.
type IdentityApprovedIterator struct {
	Event *IdentityApproved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IdentityApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityApproved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IdentityApproved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IdentityApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityApproved represents a Approved event raised by the Identity contract.
type IdentityApproved struct {
	ExecutionId *big.Int
	Approved    bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterApproved is a free log retrieval operation binding the contract event 0xb3932da477fe5d6c8ff2eafef050c0f3a1af18fc07121001482600f36f3715d8.
//
// Solidity: event Approved(uint256 indexed executionId, bool approved)
func (_Identity *IdentityFilterer) FilterApproved(opts *bind.FilterOpts, executionId []*big.Int) (*IdentityApprovedIterator, error) {

	var executionIdRule []interface{}
	for _, executionIdItem := range executionId {
		executionIdRule = append(executionIdRule, executionIdItem)
	}

	logs, sub, err := _Identity.contract.FilterLogs(opts, "Approved", executionIdRule)
	if err != nil {
		return nil, err
	}
	return &IdentityApprovedIterator{contract: _Identity.contract, event: "Approved", logs: logs, sub: sub}, nil
}

// WatchApproved is a free log subscription operation binding the contract event 0xb3932da477fe5d6c8ff2eafef050c0f3a1af18fc07121001482600f36f3715d8.
//
// Solidity: event Approved(uint256 indexed executionId, bool approved)
func (_Identity *IdentityFilterer) WatchApproved(opts *bind.WatchOpts, sink chan<- *IdentityApproved, executionId []*big.Int) (event.Subscription, error) {

	var executionIdRule []interface{}
	for _, executionIdItem := range executionId {
		executionIdRule = append(executionIdRule, executionIdItem)
	}

	logs, sub, err := _Identity.contract.WatchLogs(opts, "Approved", executionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityApproved)
				if err := _Identity.contract.UnpackLog(event, "Approved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// IdentityClaimAddedIterator is returned from FilterClaimAdded and is used to iterate over the raw logs and unpacked data for ClaimAdded events raised by the Identity contract.
type IdentityClaimAddedIterator struct {
	Event *IdentityClaimAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IdentityClaimAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityClaimAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IdentityClaimAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IdentityClaimAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityClaimAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityClaimAdded represents a ClaimAdded event raised by the Identity contract.
type IdentityClaimAdded struct {
	ClaimId   [32]byte
	Topic     *big.Int
	Scheme    *big.Int
	Issuer    common.Address
	Signature []byte
	Data      []byte
	Uri       string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimAdded is a free log retrieval operation binding the contract event 0x46149b18aa084502c3f12bc75e19eda8bda8d102b82cce8474677a6d0d5f43c5.
//
// Solidity: event ClaimAdded(bytes32 indexed claimId, uint256 indexed topic, uint256 scheme, address indexed issuer, bytes signature, bytes data, string uri)
func (_Identity *IdentityFilterer) FilterClaimAdded(opts *bind.FilterOpts, claimId [][32]byte, topic []*big.Int, issuer []common.Address) (*IdentityClaimAddedIterator, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _Identity.contract.FilterLogs(opts, "ClaimAdded", claimIdRule, topicRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &IdentityClaimAddedIterator{contract: _Identity.contract, event: "ClaimAdded", logs: logs, sub: sub}, nil
}

// WatchClaimAdded is a free log subscription operation binding the contract event 0x46149b18aa084502c3f12bc75e19eda8bda8d102b82cce8474677a6d0d5f43c5.
//
// Solidity: event ClaimAdded(bytes32 indexed claimId, uint256 indexed topic, uint256 scheme, address indexed issuer, bytes signature, bytes data, string uri)
func (_Identity *IdentityFilterer) WatchClaimAdded(opts *bind.WatchOpts, sink chan<- *IdentityClaimAdded, claimId [][32]byte, topic []*big.Int, issuer []common.Address) (event.Subscription, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _Identity.contract.WatchLogs(opts, "ClaimAdded", claimIdRule, topicRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityClaimAdded)
				if err := _Identity.contract.UnpackLog(event, "ClaimAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// IdentityClaimChangedIterator is returned from FilterClaimChanged and is used to iterate over the raw logs and unpacked data for ClaimChanged events raised by the Identity contract.
type IdentityClaimChangedIterator struct {
	Event *IdentityClaimChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IdentityClaimChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityClaimChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IdentityClaimChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IdentityClaimChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityClaimChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityClaimChanged represents a ClaimChanged event raised by the Identity contract.
type IdentityClaimChanged struct {
	ClaimId   [32]byte
	Topic     *big.Int
	Scheme    *big.Int
	Issuer    common.Address
	Signature []byte
	Data      []byte
	Uri       string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimChanged is a free log retrieval operation binding the contract event 0x3bab293fc00db832d7619a9299914251b8747c036867ec056cbd506f60135b13.
//
// Solidity: event ClaimChanged(bytes32 indexed claimId, uint256 indexed topic, uint256 scheme, address indexed issuer, bytes signature, bytes data, string uri)
func (_Identity *IdentityFilterer) FilterClaimChanged(opts *bind.FilterOpts, claimId [][32]byte, topic []*big.Int, issuer []common.Address) (*IdentityClaimChangedIterator, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _Identity.contract.FilterLogs(opts, "ClaimChanged", claimIdRule, topicRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &IdentityClaimChangedIterator{contract: _Identity.contract, event: "ClaimChanged", logs: logs, sub: sub}, nil
}

// WatchClaimChanged is a free log subscription operation binding the contract event 0x3bab293fc00db832d7619a9299914251b8747c036867ec056cbd506f60135b13.
//
// Solidity: event ClaimChanged(bytes32 indexed claimId, uint256 indexed topic, uint256 scheme, address indexed issuer, bytes signature, bytes data, string uri)
func (_Identity *IdentityFilterer) WatchClaimChanged(opts *bind.WatchOpts, sink chan<- *IdentityClaimChanged, claimId [][32]byte, topic []*big.Int, issuer []common.Address) (event.Subscription, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _Identity.contract.WatchLogs(opts, "ClaimChanged", claimIdRule, topicRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityClaimChanged)
				if err := _Identity.contract.UnpackLog(event, "ClaimChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// IdentityClaimRemovedIterator is returned from FilterClaimRemoved and is used to iterate over the raw logs and unpacked data for ClaimRemoved events raised by the Identity contract.
type IdentityClaimRemovedIterator struct {
	Event *IdentityClaimRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IdentityClaimRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityClaimRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IdentityClaimRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IdentityClaimRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityClaimRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityClaimRemoved represents a ClaimRemoved event raised by the Identity contract.
type IdentityClaimRemoved struct {
	ClaimId   [32]byte
	Topic     *big.Int
	Scheme    *big.Int
	Issuer    common.Address
	Signature []byte
	Data      []byte
	Uri       string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimRemoved is a free log retrieval operation binding the contract event 0x3cf57863a89432c61c4a27073c6ee39e8a764bff5a05aebfbcdcdc80b2e6130a.
//
// Solidity: event ClaimRemoved(bytes32 indexed claimId, uint256 indexed topic, uint256 scheme, address indexed issuer, bytes signature, bytes data, string uri)
func (_Identity *IdentityFilterer) FilterClaimRemoved(opts *bind.FilterOpts, claimId [][32]byte, topic []*big.Int, issuer []common.Address) (*IdentityClaimRemovedIterator, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _Identity.contract.FilterLogs(opts, "ClaimRemoved", claimIdRule, topicRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &IdentityClaimRemovedIterator{contract: _Identity.contract, event: "ClaimRemoved", logs: logs, sub: sub}, nil
}

// WatchClaimRemoved is a free log subscription operation binding the contract event 0x3cf57863a89432c61c4a27073c6ee39e8a764bff5a05aebfbcdcdc80b2e6130a.
//
// Solidity: event ClaimRemoved(bytes32 indexed claimId, uint256 indexed topic, uint256 scheme, address indexed issuer, bytes signature, bytes data, string uri)
func (_Identity *IdentityFilterer) WatchClaimRemoved(opts *bind.WatchOpts, sink chan<- *IdentityClaimRemoved, claimId [][32]byte, topic []*big.Int, issuer []common.Address) (event.Subscription, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _Identity.contract.WatchLogs(opts, "ClaimRemoved", claimIdRule, topicRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityClaimRemoved)
				if err := _Identity.contract.UnpackLog(event, "ClaimRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// IdentityClaimRequestedIterator is returned from FilterClaimRequested and is used to iterate over the raw logs and unpacked data for ClaimRequested events raised by the Identity contract.
type IdentityClaimRequestedIterator struct {
	Event *IdentityClaimRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IdentityClaimRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityClaimRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IdentityClaimRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IdentityClaimRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityClaimRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityClaimRequested represents a ClaimRequested event raised by the Identity contract.
type IdentityClaimRequested struct {
	ClaimRequestId *big.Int
	Topic          *big.Int
	Scheme         *big.Int
	Issuer         common.Address
	Signature      []byte
	Data           []byte
	Uri            string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterClaimRequested is a free log retrieval operation binding the contract event 0xe6b6db97dedfb44cbced367387e40b3f92dbee1b1be87f9367f7b673532b0d69.
//
// Solidity: event ClaimRequested(uint256 indexed claimRequestId, uint256 indexed topic, uint256 scheme, address indexed issuer, bytes signature, bytes data, string uri)
func (_Identity *IdentityFilterer) FilterClaimRequested(opts *bind.FilterOpts, claimRequestId []*big.Int, topic []*big.Int, issuer []common.Address) (*IdentityClaimRequestedIterator, error) {

	var claimRequestIdRule []interface{}
	for _, claimRequestIdItem := range claimRequestId {
		claimRequestIdRule = append(claimRequestIdRule, claimRequestIdItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _Identity.contract.FilterLogs(opts, "ClaimRequested", claimRequestIdRule, topicRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &IdentityClaimRequestedIterator{contract: _Identity.contract, event: "ClaimRequested", logs: logs, sub: sub}, nil
}

// WatchClaimRequested is a free log subscription operation binding the contract event 0xe6b6db97dedfb44cbced367387e40b3f92dbee1b1be87f9367f7b673532b0d69.
//
// Solidity: event ClaimRequested(uint256 indexed claimRequestId, uint256 indexed topic, uint256 scheme, address indexed issuer, bytes signature, bytes data, string uri)
func (_Identity *IdentityFilterer) WatchClaimRequested(opts *bind.WatchOpts, sink chan<- *IdentityClaimRequested, claimRequestId []*big.Int, topic []*big.Int, issuer []common.Address) (event.Subscription, error) {

	var claimRequestIdRule []interface{}
	for _, claimRequestIdItem := range claimRequestId {
		claimRequestIdRule = append(claimRequestIdRule, claimRequestIdItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _Identity.contract.WatchLogs(opts, "ClaimRequested", claimRequestIdRule, topicRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityClaimRequested)
				if err := _Identity.contract.UnpackLog(event, "ClaimRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// IdentityExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the Identity contract.
type IdentityExecutedIterator struct {
	Event *IdentityExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IdentityExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IdentityExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IdentityExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityExecuted represents a Executed event raised by the Identity contract.
type IdentityExecuted struct {
	ExecutionId *big.Int
	To          common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0x1f920dbda597d7bf95035464170fa58d0a4b57f13a1c315ace6793b9f63688b8.
//
// Solidity: event Executed(uint256 indexed executionId, address indexed to, uint256 indexed value, bytes data)
func (_Identity *IdentityFilterer) FilterExecuted(opts *bind.FilterOpts, executionId []*big.Int, to []common.Address, value []*big.Int) (*IdentityExecutedIterator, error) {

	var executionIdRule []interface{}
	for _, executionIdItem := range executionId {
		executionIdRule = append(executionIdRule, executionIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Identity.contract.FilterLogs(opts, "Executed", executionIdRule, toRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &IdentityExecutedIterator{contract: _Identity.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0x1f920dbda597d7bf95035464170fa58d0a4b57f13a1c315ace6793b9f63688b8.
//
// Solidity: event Executed(uint256 indexed executionId, address indexed to, uint256 indexed value, bytes data)
func (_Identity *IdentityFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *IdentityExecuted, executionId []*big.Int, to []common.Address, value []*big.Int) (event.Subscription, error) {

	var executionIdRule []interface{}
	for _, executionIdItem := range executionId {
		executionIdRule = append(executionIdRule, executionIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Identity.contract.WatchLogs(opts, "Executed", executionIdRule, toRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityExecuted)
				if err := _Identity.contract.UnpackLog(event, "Executed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// IdentityExecutionFailedIterator is returned from FilterExecutionFailed and is used to iterate over the raw logs and unpacked data for ExecutionFailed events raised by the Identity contract.
type IdentityExecutionFailedIterator struct {
	Event *IdentityExecutionFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IdentityExecutionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityExecutionFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IdentityExecutionFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IdentityExecutionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityExecutionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityExecutionFailed represents a ExecutionFailed event raised by the Identity contract.
type IdentityExecutionFailed struct {
	ExecutionId *big.Int
	To          common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecutionFailed is a free log retrieval operation binding the contract event 0xe10c49d9f7c71da23262367013434763cfdb2332267641728d25cd712c5c6a68.
//
// Solidity: event ExecutionFailed(uint256 indexed executionId, address indexed to, uint256 indexed value, bytes data)
func (_Identity *IdentityFilterer) FilterExecutionFailed(opts *bind.FilterOpts, executionId []*big.Int, to []common.Address, value []*big.Int) (*IdentityExecutionFailedIterator, error) {

	var executionIdRule []interface{}
	for _, executionIdItem := range executionId {
		executionIdRule = append(executionIdRule, executionIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Identity.contract.FilterLogs(opts, "ExecutionFailed", executionIdRule, toRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &IdentityExecutionFailedIterator{contract: _Identity.contract, event: "ExecutionFailed", logs: logs, sub: sub}, nil
}

// WatchExecutionFailed is a free log subscription operation binding the contract event 0xe10c49d9f7c71da23262367013434763cfdb2332267641728d25cd712c5c6a68.
//
// Solidity: event ExecutionFailed(uint256 indexed executionId, address indexed to, uint256 indexed value, bytes data)
func (_Identity *IdentityFilterer) WatchExecutionFailed(opts *bind.WatchOpts, sink chan<- *IdentityExecutionFailed, executionId []*big.Int, to []common.Address, value []*big.Int) (event.Subscription, error) {

	var executionIdRule []interface{}
	for _, executionIdItem := range executionId {
		executionIdRule = append(executionIdRule, executionIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Identity.contract.WatchLogs(opts, "ExecutionFailed", executionIdRule, toRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityExecutionFailed)
				if err := _Identity.contract.UnpackLog(event, "ExecutionFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// IdentityExecutionRequestedIterator is returned from FilterExecutionRequested and is used to iterate over the raw logs and unpacked data for ExecutionRequested events raised by the Identity contract.
type IdentityExecutionRequestedIterator struct {
	Event *IdentityExecutionRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IdentityExecutionRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityExecutionRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IdentityExecutionRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IdentityExecutionRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityExecutionRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityExecutionRequested represents a ExecutionRequested event raised by the Identity contract.
type IdentityExecutionRequested struct {
	ExecutionId *big.Int
	To          common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecutionRequested is a free log retrieval operation binding the contract event 0x8afcfabcb00e47a53a8fc3e9f23ff47ee1926194bb1350dd007c50b412a6cee8.
//
// Solidity: event ExecutionRequested(uint256 indexed executionId, address indexed to, uint256 indexed value, bytes data)
func (_Identity *IdentityFilterer) FilterExecutionRequested(opts *bind.FilterOpts, executionId []*big.Int, to []common.Address, value []*big.Int) (*IdentityExecutionRequestedIterator, error) {

	var executionIdRule []interface{}
	for _, executionIdItem := range executionId {
		executionIdRule = append(executionIdRule, executionIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Identity.contract.FilterLogs(opts, "ExecutionRequested", executionIdRule, toRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &IdentityExecutionRequestedIterator{contract: _Identity.contract, event: "ExecutionRequested", logs: logs, sub: sub}, nil
}

// WatchExecutionRequested is a free log subscription operation binding the contract event 0x8afcfabcb00e47a53a8fc3e9f23ff47ee1926194bb1350dd007c50b412a6cee8.
//
// Solidity: event ExecutionRequested(uint256 indexed executionId, address indexed to, uint256 indexed value, bytes data)
func (_Identity *IdentityFilterer) WatchExecutionRequested(opts *bind.WatchOpts, sink chan<- *IdentityExecutionRequested, executionId []*big.Int, to []common.Address, value []*big.Int) (event.Subscription, error) {

	var executionIdRule []interface{}
	for _, executionIdItem := range executionId {
		executionIdRule = append(executionIdRule, executionIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Identity.contract.WatchLogs(opts, "ExecutionRequested", executionIdRule, toRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityExecutionRequested)
				if err := _Identity.contract.UnpackLog(event, "ExecutionRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// IdentityKeyAddedIterator is returned from FilterKeyAdded and is used to iterate over the raw logs and unpacked data for KeyAdded events raised by the Identity contract.
type IdentityKeyAddedIterator struct {
	Event *IdentityKeyAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IdentityKeyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityKeyAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IdentityKeyAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IdentityKeyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityKeyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityKeyAdded represents a KeyAdded event raised by the Identity contract.
type IdentityKeyAdded struct {
	Key     [32]byte
	Purpose *big.Int
	KeyType *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterKeyAdded is a free log retrieval operation binding the contract event 0x480000bb1edad8ca1470381cc334b1917fbd51c6531f3a623ea8e0ec7e38a6e9.
//
// Solidity: event KeyAdded(bytes32 indexed key, uint256 indexed purpose, uint256 indexed keyType)
func (_Identity *IdentityFilterer) FilterKeyAdded(opts *bind.FilterOpts, key [][32]byte, purpose []*big.Int, keyType []*big.Int) (*IdentityKeyAddedIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}
	var purposeRule []interface{}
	for _, purposeItem := range purpose {
		purposeRule = append(purposeRule, purposeItem)
	}
	var keyTypeRule []interface{}
	for _, keyTypeItem := range keyType {
		keyTypeRule = append(keyTypeRule, keyTypeItem)
	}

	logs, sub, err := _Identity.contract.FilterLogs(opts, "KeyAdded", keyRule, purposeRule, keyTypeRule)
	if err != nil {
		return nil, err
	}
	return &IdentityKeyAddedIterator{contract: _Identity.contract, event: "KeyAdded", logs: logs, sub: sub}, nil
}

// WatchKeyAdded is a free log subscription operation binding the contract event 0x480000bb1edad8ca1470381cc334b1917fbd51c6531f3a623ea8e0ec7e38a6e9.
//
// Solidity: event KeyAdded(bytes32 indexed key, uint256 indexed purpose, uint256 indexed keyType)
func (_Identity *IdentityFilterer) WatchKeyAdded(opts *bind.WatchOpts, sink chan<- *IdentityKeyAdded, key [][32]byte, purpose []*big.Int, keyType []*big.Int) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}
	var purposeRule []interface{}
	for _, purposeItem := range purpose {
		purposeRule = append(purposeRule, purposeItem)
	}
	var keyTypeRule []interface{}
	for _, keyTypeItem := range keyType {
		keyTypeRule = append(keyTypeRule, keyTypeItem)
	}

	logs, sub, err := _Identity.contract.WatchLogs(opts, "KeyAdded", keyRule, purposeRule, keyTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityKeyAdded)
				if err := _Identity.contract.UnpackLog(event, "KeyAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// IdentityKeyRemovedIterator is returned from FilterKeyRemoved and is used to iterate over the raw logs and unpacked data for KeyRemoved events raised by the Identity contract.
type IdentityKeyRemovedIterator struct {
	Event *IdentityKeyRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IdentityKeyRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityKeyRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IdentityKeyRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IdentityKeyRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityKeyRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityKeyRemoved represents a KeyRemoved event raised by the Identity contract.
type IdentityKeyRemoved struct {
	Key     [32]byte
	Purpose *big.Int
	KeyType *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterKeyRemoved is a free log retrieval operation binding the contract event 0x585a4aef50f8267a92b32412b331b20f7f8b96f2245b253b9cc50dcc621d3397.
//
// Solidity: event KeyRemoved(bytes32 indexed key, uint256 indexed purpose, uint256 indexed keyType)
func (_Identity *IdentityFilterer) FilterKeyRemoved(opts *bind.FilterOpts, key [][32]byte, purpose []*big.Int, keyType []*big.Int) (*IdentityKeyRemovedIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}
	var purposeRule []interface{}
	for _, purposeItem := range purpose {
		purposeRule = append(purposeRule, purposeItem)
	}
	var keyTypeRule []interface{}
	for _, keyTypeItem := range keyType {
		keyTypeRule = append(keyTypeRule, keyTypeItem)
	}

	logs, sub, err := _Identity.contract.FilterLogs(opts, "KeyRemoved", keyRule, purposeRule, keyTypeRule)
	if err != nil {
		return nil, err
	}
	return &IdentityKeyRemovedIterator{contract: _Identity.contract, event: "KeyRemoved", logs: logs, sub: sub}, nil
}

// WatchKeyRemoved is a free log subscription operation binding the contract event 0x585a4aef50f8267a92b32412b331b20f7f8b96f2245b253b9cc50dcc621d3397.
//
// Solidity: event KeyRemoved(bytes32 indexed key, uint256 indexed purpose, uint256 indexed keyType)
func (_Identity *IdentityFilterer) WatchKeyRemoved(opts *bind.WatchOpts, sink chan<- *IdentityKeyRemoved, key [][32]byte, purpose []*big.Int, keyType []*big.Int) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}
	var purposeRule []interface{}
	for _, purposeItem := range purpose {
		purposeRule = append(purposeRule, purposeItem)
	}
	var keyTypeRule []interface{}
	for _, keyTypeItem := range keyType {
		keyTypeRule = append(keyTypeRule, keyTypeItem)
	}

	logs, sub, err := _Identity.contract.WatchLogs(opts, "KeyRemoved", keyRule, purposeRule, keyTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityKeyRemoved)
				if err := _Identity.contract.UnpackLog(event, "KeyRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// IdentityLogPauseIterator is returned from FilterLogPause and is used to iterate over the raw logs and unpacked data for LogPause events raised by the Identity contract.
type IdentityLogPauseIterator struct {
	Event *IdentityLogPause // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IdentityLogPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityLogPause)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IdentityLogPause)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IdentityLogPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityLogPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityLogPause represents a LogPause event raised by the Identity contract.
type IdentityLogPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogPause is a free log retrieval operation binding the contract event 0x4b314b34e912fda7f95e7d23e9c8c95f82f0aff1984e4ce592a0b005f9055624.
//
// Solidity: event LogPause()
func (_Identity *IdentityFilterer) FilterLogPause(opts *bind.FilterOpts) (*IdentityLogPauseIterator, error) {

	logs, sub, err := _Identity.contract.FilterLogs(opts, "LogPause")
	if err != nil {
		return nil, err
	}
	return &IdentityLogPauseIterator{contract: _Identity.contract, event: "LogPause", logs: logs, sub: sub}, nil
}

// WatchLogPause is a free log subscription operation binding the contract event 0x4b314b34e912fda7f95e7d23e9c8c95f82f0aff1984e4ce592a0b005f9055624.
//
// Solidity: event LogPause()
func (_Identity *IdentityFilterer) WatchLogPause(opts *bind.WatchOpts, sink chan<- *IdentityLogPause) (event.Subscription, error) {

	logs, sub, err := _Identity.contract.WatchLogs(opts, "LogPause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityLogPause)
				if err := _Identity.contract.UnpackLog(event, "LogPause", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// IdentityLogUnpauseIterator is returned from FilterLogUnpause and is used to iterate over the raw logs and unpacked data for LogUnpause events raised by the Identity contract.
type IdentityLogUnpauseIterator struct {
	Event *IdentityLogUnpause // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IdentityLogUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityLogUnpause)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IdentityLogUnpause)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IdentityLogUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityLogUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityLogUnpause represents a LogUnpause event raised by the Identity contract.
type IdentityLogUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogUnpause is a free log retrieval operation binding the contract event 0x730c1faaa977b67dacf1e2451ef54556e04a07d577785ff79f6d31f73502efc9.
//
// Solidity: event LogUnpause()
func (_Identity *IdentityFilterer) FilterLogUnpause(opts *bind.FilterOpts) (*IdentityLogUnpauseIterator, error) {

	logs, sub, err := _Identity.contract.FilterLogs(opts, "LogUnpause")
	if err != nil {
		return nil, err
	}
	return &IdentityLogUnpauseIterator{contract: _Identity.contract, event: "LogUnpause", logs: logs, sub: sub}, nil
}

// WatchLogUnpause is a free log subscription operation binding the contract event 0x730c1faaa977b67dacf1e2451ef54556e04a07d577785ff79f6d31f73502efc9.
//
// Solidity: event LogUnpause()
func (_Identity *IdentityFilterer) WatchLogUnpause(opts *bind.WatchOpts, sink chan<- *IdentityLogUnpause) (event.Subscription, error) {

	logs, sub, err := _Identity.contract.WatchLogs(opts, "LogUnpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityLogUnpause)
				if err := _Identity.contract.UnpackLog(event, "LogUnpause", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

type metaIDKey struct {
	Purposes []*big.Int
	KeyType  *big.Int
	Key      [32]byte
}
