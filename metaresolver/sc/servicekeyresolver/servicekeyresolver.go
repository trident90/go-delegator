// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package servicekeyresolver

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ServicekeyresolverABI is the input ABI used to generate the binding from.
const ServicekeyresolverABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"signatureTimeout\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"isSigned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"identityRegistryAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"key\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"ein\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"KeyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"key\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"ein\",\"type\":\"uint256\"}],\"name\":\"KeyRemoved\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"associatedAddress\",\"type\":\"address\"},{\"name\":\"key\",\"type\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"addKeyDelegated\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"addKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"associatedAddress\",\"type\":\"address\"},{\"name\":\"key\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"removeKeyDelegated\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"address\"}],\"name\":\"removeKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"associatedAddress\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"removeKeysDelegated\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeKeys\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"address\"},{\"name\":\"ein\",\"type\":\"uint256\"}],\"name\":\"isKeyFor\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"address\"}],\"name\":\"getSymbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ein\",\"type\":\"uint256\"}],\"name\":\"getKeys\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ServicekeyresolverBin is the compiled bytecode used for deploying new contracts.
const ServicekeyresolverBin = `60806040526201518060045534801561001757600080fd5b50604051602080611ba98339810180604052602081101561003757600080fd5b505160008054600160a060020a03909216600160a060020a0319909216919091179055611b40806100696000396000f3fe6080604052600436106100ae5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166335434b5d81146100b35780635437b67c1461010d578063600667261461013457806369e78499146101825780637f1ccc25146101b55780638677ebe814610202578063871fc606146102505780638a4ac03f146102dd5780638d357fa314610388578063a36fee1714610402578063c9b2e52214610417575b600080fd5b3480156100bf57600080fd5b5061010b600480360360c08110156100d657600080fd5b50600160a060020a03813581169160208101359091169060ff6040820135169060608101359060808101359060a001356104bf565b005b34801561011957600080fd5b506101226107d4565b60408051918252519081900360200190f35b34801561014057600080fd5b5061010b600480360360a081101561015757600080fd5b50600160a060020a038135169060ff60208201351690604081013590606081013590608001356107da565b34801561018e57600080fd5b5061010b600480360360208110156101a557600080fd5b5035600160a060020a0316610adc565b3480156101c157600080fd5b506101ee600480360360408110156101d857600080fd5b50600160a060020a038135169060200135610b60565b604080519115158252519081900360200190f35b34801561020e57600080fd5b506101ee600480360360a081101561022557600080fd5b50600160a060020a038135169060208101359060ff6040820135169060608101359060800135610c8e565b34801561025c57600080fd5b5061010b6004803603604081101561027357600080fd5b600160a060020a03823516919081019060408101602082013564010000000081111561029e57600080fd5b8201836020820111156102b057600080fd5b803590602001918460018302840111640100000000831117156102d257600080fd5b509092509050610cba565b3480156102e957600080fd5b5061010b600480360360e081101561030057600080fd5b600160a060020a03823581169260208101359091169181019060608101604082013564010000000081111561033457600080fd5b82018360208201111561034657600080fd5b8035906020019184600183028401116401000000008311171561036857600080fd5b919350915060ff8135169060208101359060408101359060600135610d75565b34801561039457600080fd5b506103b2600480360360208110156103ab57600080fd5b50356110e3565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156103ee5781810151838201526020016103d6565b505050509050019250505060405180910390f35b34801561040e57600080fd5b5061010b611265565b34801561042357600080fd5b5061044a6004803603602081101561043a57600080fd5b5035600160a060020a03166112e7565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561048457818101518382015260200161046c565b50505050905090810190601f1680156104b15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b808042101580156104d35750600454810142105b1515610529576040805160e560020a62461bcd02815260206004820152601760248201527f54696d657374616d70206973206e6f742076616c69642e000000000000000000604482015290519081900360640190fd5b600080546040805160e060020a6305c62c2f028152600160a060020a038b81166004830152915191909216916305c62c2f916024808301926020929190829003018186803b15801561057a57600080fd5b505afa15801561058e573d6000803e3d6000fd5b505050506040513d60208110156105a457600080fd5b5051600054604080517f53a9698a000000000000000000000000000000000000000000000000000000008152600481018490523360248201529051929350600160a060020a03909116916353a9698a91604480820192602092909190829003018186803b15801561061457600080fd5b505afa158015610628573d6000803e3d6000fd5b505050506040513d602081101561063e57600080fd5b50511515610696576040805160e560020a62461bcd02815260206004820152601f60248201527f4f6e6c792070726f76696465722063616e2062652064656c6567617465642e00604482015290519081900360640190fd5b604080517f1900000000000000000000000000000000000000000000000000000000000000602080830191909152600060218301526c0100000000000000000000000030810260228401527f4920617574686f72697a65207468652072656d6f76616c206f6620612073657260368401527f76696365206b6579206f6e206d7920626568616c662e000000000000000000006056840152600160a060020a038b1602606c83015260808083018790528351808403909101815260a0909201909252805191012061076a908990888888610c8e565b15156107c0576040805160e560020a62461bcd02815260206004820152601260248201527f5065726d697373696f6e2064656e6965642e0000000000000000000000000000604482015290519081900360640190fd5b6107ca8188611392565b5050505050505050565b60045481565b808042101580156107ee5750600454810142105b1515610844576040805160e560020a62461bcd02815260206004820152601760248201527f54696d657374616d70206973206e6f742076616c69642e000000000000000000604482015290519081900360640190fd5b600080546040805160e060020a6305c62c2f028152600160a060020a038a81166004830152915191909216916305c62c2f916024808301926020929190829003018186803b15801561089557600080fd5b505afa1580156108a9573d6000803e3d6000fd5b505050506040513d60208110156108bf57600080fd5b5051600054604080517f53a9698a000000000000000000000000000000000000000000000000000000008152600481018490523360248201529051929350600160a060020a03909116916353a9698a91604480820192602092909190829003018186803b15801561092f57600080fd5b505afa158015610943573d6000803e3d6000fd5b505050506040513d602081101561095957600080fd5b505115156109b1576040805160e560020a62461bcd02815260206004820152601f60248201527f4f6e6c792070726f76696465722063616e2062652064656c6567617465642e00604482015290519081900360640190fd5b604080517f1900000000000000000000000000000000000000000000000000000000000000602080830191909152600060218301526c01000000000000000000000000300260228301527f4920617574686f72697a65207468652072656d6f76616c206f6620616c6c207360368301527f657276696365206b657973206f6e206d7920626568616c662e000000000000006056830152606f80830187905283518084039091018152608f9092019092528051910120610a74908890888888610c8e565b1515610aca576040805160e560020a62461bcd02815260206004820152601260248201527f5065726d697373696f6e2064656e6965642e0000000000000000000000000000604482015290519081900360640190fd5b610ad3816114dc565b50505050505050565b6000546040805160e060020a6305c62c2f0281523360048201529051610b5d92600160a060020a0316916305c62c2f916024808301926020929190829003018186803b158015610b2b57600080fd5b505afa158015610b3f573d6000803e3d6000fd5b505050506040513d6020811015610b5557600080fd5b505182611392565b50565b60008054604080517f5b5aed3a0000000000000000000000000000000000000000000000000000000081526004810185905290518492600160a060020a031691635b5aed3a916024808301926020929190829003018186803b158015610bc557600080fd5b505afa158015610bd9573d6000803e3d6000fd5b505050506040513d6020811015610bef57600080fd5b50511515610c6d576040805160e560020a62461bcd02815260206004820152602760248201527f546865207265666572656e636564206964656e7469747920646f6573206e6f7460448201527f2065786973742e00000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b5050600160a060020a03919091166000908152600260205260409020541490565b6000610c9d8686868686611550565b80610cb05750610cb086868686866115d8565b9695505050505050565b6000546040805160e060020a6305c62c2f0281523360048201529051610d7092600160a060020a0316916305c62c2f916024808301926020929190829003018186803b158015610d0957600080fd5b505afa158015610d1d573d6000803e3d6000fd5b505050506040513d6020811015610d3357600080fd5b5051604080516020601f86018190048102820181019092528481528691869086908190840183828082843760009201919091525061169392505050565b505050565b80804210158015610d895750600454810142105b1515610ddf576040805160e560020a62461bcd02815260206004820152601760248201527f54696d657374616d70206973206e6f742076616c69642e000000000000000000604482015290519081900360640190fd5b600080546040805160e060020a6305c62c2f028152600160a060020a038d81166004830152915191909216916305c62c2f916024808301926020929190829003018186803b158015610e3057600080fd5b505afa158015610e44573d6000803e3d6000fd5b505050506040513d6020811015610e5a57600080fd5b5051600054604080517f53a9698a000000000000000000000000000000000000000000000000000000008152600481018490523360248201529051929350600160a060020a03909116916353a9698a91604480820192602092909190829003018186803b158015610eca57600080fd5b505afa158015610ede573d6000803e3d6000fd5b505050506040513d6020811015610ef457600080fd5b50511515610f4c576040805160e560020a62461bcd02815260206004820152601f60248201527f4f6e6c792070726f76696465722063616e2062652064656c6567617465642e00604482015290519081900360640190fd5b6040517f190000000000000000000000000000000000000000000000000000000000000060208201818152600060218401819052306c0100000000000000000000000081810260228701527f4920617574686f72697a6520746865206164646974696f6e206f66206120736560368701527f7276696365206b6579206f6e206d7920626568616c662e0000000000000000006056870152600160a060020a038f1602606d860152611040948f94938f918f918f918c91906081018484808284378083019250505082815260200197505050505050505060405160208183030381529060405280519060200120888888610c8e565b1515611096576040805160e560020a62461bcd02815260206004820152601260248201527f5065726d697373696f6e2064656e6965642e0000000000000000000000000000604482015290519081900360640190fd5b6110d7818a8a8a8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061169392505050565b50505050505050505050565b600054604080517f5b5aed3a0000000000000000000000000000000000000000000000000000000081526004810184905290516060928492600160a060020a0390911691635b5aed3a91602480820192602092909190829003018186803b15801561114d57600080fd5b505afa158015611161573d6000803e3d6000fd5b505050506040513d602081101561117757600080fd5b505115156111f5576040805160e560020a62461bcd02815260206004820152602760248201527f546865207265666572656e636564206964656e7469747920646f6573206e6f7460448201527f2065786973742e00000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600083815260016020908152604091829020805483518184028101840190945280845290929183919083018282801561125757602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311611239575b505050505092505050919050565b6000546040805160e060020a6305c62c2f02815233600482015290516112e592600160a060020a0316916305c62c2f916024808301926020929190829003018186803b1580156112b457600080fd5b505afa1580156112c8573d6000803e3d6000fd5b505050506040513d60208110156112de57600080fd5b50516114dc565b565b600160a060020a03811660009081526003602090815260409182902080548351601f60026000196101006001861615020190931692909204918201849004840281018401909452808452606093928301828280156113865780601f1061135b57610100808354040283529160200191611386565b820191906000526020600020905b81548152906001019060200180831161136957829003601f168201915b50505050509050919050565b600054604080517fd4b1cdcc0000000000000000000000000000000000000000000000000000000081526004810185905230602482015290518492600160a060020a03169163d4b1cdcc916044808301926020929190829003018186803b1580156113fc57600080fd5b505afa158015611410573d6000803e3d6000fd5b505050506040513d602081101561142657600080fd5b505115156114a4576040805160e560020a62461bcd02815260206004820152603560248201527f5468652063616c6c696e67206964656e7469747920646f6573206e6f7420686160448201527f76652074686973207265736f6c766572207365742e0000000000000000000000606482015290519081900360840190fd5b600160a060020a038216600090815260026020908152604080832083905585835260019091529020610d70908363ffffffff61182116565b6000818152600160205260408120905b6114f582611963565b81101561154057600060026000846000018481548110151561151357fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020556001016114ec565b5061154c816000611a5b565b5050565b600085600160a060020a031660018686868660405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa1580156115b9573d6000803e3d6000fd5b50505060206040510351600160a060020a031614905095945050505050565b604080518082018252601c8082527f19457468657265756d205369676e6564204d6573736167653a0a33320000000060208084019182529351600094611688938b9386938c9301918291908083835b602083106116465780518252601f199092019160209182019101611627565b51815160209384036101000a60001901801990921691161790529201938452506040805180850381529382019052825192019190912091508890508787611550565b979650505050505050565b600054604080517fd4b1cdcc0000000000000000000000000000000000000000000000000000000081526004810186905230602482015290518592600160a060020a03169163d4b1cdcc916044808301926020929190829003018186803b1580156116fd57600080fd5b505afa158015611711573d6000803e3d6000fd5b505050506040513d602081101561172757600080fd5b505115156117a5576040805160e560020a62461bcd02815260206004820152603560248201527f5468652063616c6c696e67206964656e7469747920646f6573206e6f7420686160448201527f76652074686973207265736f6c766572207365742e0000000000000000000000606482015290519081900360840190fd5b600160a060020a03831660009081526002602090815260408083208790556003825290912083516117d892850190611a79565b5060008481526001602052604090206117f7908463ffffffff61196716565b151561181b57600084815260016020526040902061181b908463ffffffff6119fc16565b50505050565b61182b8282611967565b1561154c5781600161183c82611963565b0381548110151561184957fe5b6000918252602080832090910154600160a060020a0384811684526001860190925260409092205484549190921691849160001990910190811061188957fe5b6000918252602080832091909101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03948516179055918316815260018401918290526040812054845490929190859060001985019081106118e657fe5b6000918252602080832090910154600160a060020a039081168452838201949094526040928301822094909455918416825260018501909252908120558154829080151561193057fe5b6000828152602090208101600019908101805473ffffffffffffffffffffffffffffffffffffffff191690550190555050565b5490565b600160a060020a0381166000908152600183016020526040812054811080156119ab5750600160a060020a0382166000908152600184016020526040902054835410155b80156119f55750600160a060020a038216600081815260018501602052604090205484548591600019019081106119de57fe5b600091825260209091200154600160a060020a0316145b9392505050565b611a068282611967565b151561154c578154600180820180855560008581526020808220909401805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03969096169586179055938452930190526040902055565b5080546000825590600052602060002090810190610b5d9190611af7565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611aba57805160ff1916838001178555611ae7565b82800160010185558215611ae7579182015b82811115611ae7578251825591602001919060010190611acc565b50611af3929150611af7565b5090565b611b1191905b80821115611af35760008155600101611afd565b9056fea165627a7a7230582027bf5119839dca99bce0f3cac3bb5e927e47093d3a0dca1d56f00909b11ca1900029`

// DeployServicekeyresolver deploys a new Ethereum contract, binding an instance of Servicekeyresolver to it.
func DeployServicekeyresolver(auth *bind.TransactOpts, backend bind.ContractBackend, identityRegistryAddress common.Address) (common.Address, *types.Transaction, *Servicekeyresolver, error) {
	parsed, err := abi.JSON(strings.NewReader(ServicekeyresolverABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ServicekeyresolverBin), backend, identityRegistryAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Servicekeyresolver{ServicekeyresolverCaller: ServicekeyresolverCaller{contract: contract}, ServicekeyresolverTransactor: ServicekeyresolverTransactor{contract: contract}, ServicekeyresolverFilterer: ServicekeyresolverFilterer{contract: contract}}, nil
}

// Servicekeyresolver is an auto generated Go binding around an Ethereum contract.
type Servicekeyresolver struct {
	ServicekeyresolverCaller     // Read-only binding to the contract
	ServicekeyresolverTransactor // Write-only binding to the contract
	ServicekeyresolverFilterer   // Log filterer for contract events
}

// ServicekeyresolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ServicekeyresolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServicekeyresolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ServicekeyresolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServicekeyresolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ServicekeyresolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServicekeyresolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ServicekeyresolverSession struct {
	Contract     *Servicekeyresolver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ServicekeyresolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ServicekeyresolverCallerSession struct {
	Contract *ServicekeyresolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ServicekeyresolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ServicekeyresolverTransactorSession struct {
	Contract     *ServicekeyresolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ServicekeyresolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ServicekeyresolverRaw struct {
	Contract *Servicekeyresolver // Generic contract binding to access the raw methods on
}

// ServicekeyresolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ServicekeyresolverCallerRaw struct {
	Contract *ServicekeyresolverCaller // Generic read-only contract binding to access the raw methods on
}

// ServicekeyresolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ServicekeyresolverTransactorRaw struct {
	Contract *ServicekeyresolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewServicekeyresolver creates a new instance of Servicekeyresolver, bound to a specific deployed contract.
func NewServicekeyresolver(address common.Address, backend bind.ContractBackend) (*Servicekeyresolver, error) {
	contract, err := bindServicekeyresolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Servicekeyresolver{ServicekeyresolverCaller: ServicekeyresolverCaller{contract: contract}, ServicekeyresolverTransactor: ServicekeyresolverTransactor{contract: contract}, ServicekeyresolverFilterer: ServicekeyresolverFilterer{contract: contract}}, nil
}

// NewServicekeyresolverCaller creates a new read-only instance of Servicekeyresolver, bound to a specific deployed contract.
func NewServicekeyresolverCaller(address common.Address, caller bind.ContractCaller) (*ServicekeyresolverCaller, error) {
	contract, err := bindServicekeyresolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ServicekeyresolverCaller{contract: contract}, nil
}

// NewServicekeyresolverTransactor creates a new write-only instance of Servicekeyresolver, bound to a specific deployed contract.
func NewServicekeyresolverTransactor(address common.Address, transactor bind.ContractTransactor) (*ServicekeyresolverTransactor, error) {
	contract, err := bindServicekeyresolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ServicekeyresolverTransactor{contract: contract}, nil
}

// NewServicekeyresolverFilterer creates a new log filterer instance of Servicekeyresolver, bound to a specific deployed contract.
func NewServicekeyresolverFilterer(address common.Address, filterer bind.ContractFilterer) (*ServicekeyresolverFilterer, error) {
	contract, err := bindServicekeyresolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ServicekeyresolverFilterer{contract: contract}, nil
}

// bindServicekeyresolver binds a generic wrapper to an already deployed contract.
func bindServicekeyresolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ServicekeyresolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Servicekeyresolver *ServicekeyresolverRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Servicekeyresolver.Contract.ServicekeyresolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Servicekeyresolver *ServicekeyresolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Servicekeyresolver.Contract.ServicekeyresolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Servicekeyresolver *ServicekeyresolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Servicekeyresolver.Contract.ServicekeyresolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Servicekeyresolver *ServicekeyresolverCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Servicekeyresolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Servicekeyresolver *ServicekeyresolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Servicekeyresolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Servicekeyresolver *ServicekeyresolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Servicekeyresolver.Contract.contract.Transact(opts, method, params...)
}

// GetKeys is a free data retrieval call binding the contract method 0x8d357fa3.
//
// Solidity: function getKeys(uint256 ein) constant returns(address[])
func (_Servicekeyresolver *ServicekeyresolverCaller) GetKeys(opts *bind.CallOpts, ein *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Servicekeyresolver.contract.Call(opts, out, "getKeys", ein)
	return *ret0, err
}

// GetKeys is a free data retrieval call binding the contract method 0x8d357fa3.
//
// Solidity: function getKeys(uint256 ein) constant returns(address[])
func (_Servicekeyresolver *ServicekeyresolverSession) GetKeys(ein *big.Int) ([]common.Address, error) {
	return _Servicekeyresolver.Contract.GetKeys(&_Servicekeyresolver.CallOpts, ein)
}

// GetKeys is a free data retrieval call binding the contract method 0x8d357fa3.
//
// Solidity: function getKeys(uint256 ein) constant returns(address[])
func (_Servicekeyresolver *ServicekeyresolverCallerSession) GetKeys(ein *big.Int) ([]common.Address, error) {
	return _Servicekeyresolver.Contract.GetKeys(&_Servicekeyresolver.CallOpts, ein)
}

// GetSymbol is a free data retrieval call binding the contract method 0xc9b2e522.
//
// Solidity: function getSymbol(address key) constant returns(string)
func (_Servicekeyresolver *ServicekeyresolverCaller) GetSymbol(opts *bind.CallOpts, key common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Servicekeyresolver.contract.Call(opts, out, "getSymbol", key)
	return *ret0, err
}

// GetSymbol is a free data retrieval call binding the contract method 0xc9b2e522.
//
// Solidity: function getSymbol(address key) constant returns(string)
func (_Servicekeyresolver *ServicekeyresolverSession) GetSymbol(key common.Address) (string, error) {
	return _Servicekeyresolver.Contract.GetSymbol(&_Servicekeyresolver.CallOpts, key)
}

// GetSymbol is a free data retrieval call binding the contract method 0xc9b2e522.
//
// Solidity: function getSymbol(address key) constant returns(string)
func (_Servicekeyresolver *ServicekeyresolverCallerSession) GetSymbol(key common.Address) (string, error) {
	return _Servicekeyresolver.Contract.GetSymbol(&_Servicekeyresolver.CallOpts, key)
}

// IsKeyFor is a free data retrieval call binding the contract method 0x7f1ccc25.
//
// Solidity: function isKeyFor(address key, uint256 ein) constant returns(bool)
func (_Servicekeyresolver *ServicekeyresolverCaller) IsKeyFor(opts *bind.CallOpts, key common.Address, ein *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Servicekeyresolver.contract.Call(opts, out, "isKeyFor", key, ein)
	return *ret0, err
}

// IsKeyFor is a free data retrieval call binding the contract method 0x7f1ccc25.
//
// Solidity: function isKeyFor(address key, uint256 ein) constant returns(bool)
func (_Servicekeyresolver *ServicekeyresolverSession) IsKeyFor(key common.Address, ein *big.Int) (bool, error) {
	return _Servicekeyresolver.Contract.IsKeyFor(&_Servicekeyresolver.CallOpts, key, ein)
}

// IsKeyFor is a free data retrieval call binding the contract method 0x7f1ccc25.
//
// Solidity: function isKeyFor(address key, uint256 ein) constant returns(bool)
func (_Servicekeyresolver *ServicekeyresolverCallerSession) IsKeyFor(key common.Address, ein *big.Int) (bool, error) {
	return _Servicekeyresolver.Contract.IsKeyFor(&_Servicekeyresolver.CallOpts, key, ein)
}

// IsSigned is a free data retrieval call binding the contract method 0x8677ebe8.
//
// Solidity: function isSigned(address _address, bytes32 messageHash, uint8 v, bytes32 r, bytes32 s) constant returns(bool)
func (_Servicekeyresolver *ServicekeyresolverCaller) IsSigned(opts *bind.CallOpts, _address common.Address, messageHash [32]byte, v uint8, r [32]byte, s [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Servicekeyresolver.contract.Call(opts, out, "isSigned", _address, messageHash, v, r, s)
	return *ret0, err
}

// IsSigned is a free data retrieval call binding the contract method 0x8677ebe8.
//
// Solidity: function isSigned(address _address, bytes32 messageHash, uint8 v, bytes32 r, bytes32 s) constant returns(bool)
func (_Servicekeyresolver *ServicekeyresolverSession) IsSigned(_address common.Address, messageHash [32]byte, v uint8, r [32]byte, s [32]byte) (bool, error) {
	return _Servicekeyresolver.Contract.IsSigned(&_Servicekeyresolver.CallOpts, _address, messageHash, v, r, s)
}

// IsSigned is a free data retrieval call binding the contract method 0x8677ebe8.
//
// Solidity: function isSigned(address _address, bytes32 messageHash, uint8 v, bytes32 r, bytes32 s) constant returns(bool)
func (_Servicekeyresolver *ServicekeyresolverCallerSession) IsSigned(_address common.Address, messageHash [32]byte, v uint8, r [32]byte, s [32]byte) (bool, error) {
	return _Servicekeyresolver.Contract.IsSigned(&_Servicekeyresolver.CallOpts, _address, messageHash, v, r, s)
}

// SignatureTimeout is a free data retrieval call binding the contract method 0x5437b67c.
//
// Solidity: function signatureTimeout() constant returns(uint256)
func (_Servicekeyresolver *ServicekeyresolverCaller) SignatureTimeout(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Servicekeyresolver.contract.Call(opts, out, "signatureTimeout")
	return *ret0, err
}

// SignatureTimeout is a free data retrieval call binding the contract method 0x5437b67c.
//
// Solidity: function signatureTimeout() constant returns(uint256)
func (_Servicekeyresolver *ServicekeyresolverSession) SignatureTimeout() (*big.Int, error) {
	return _Servicekeyresolver.Contract.SignatureTimeout(&_Servicekeyresolver.CallOpts)
}

// SignatureTimeout is a free data retrieval call binding the contract method 0x5437b67c.
//
// Solidity: function signatureTimeout() constant returns(uint256)
func (_Servicekeyresolver *ServicekeyresolverCallerSession) SignatureTimeout() (*big.Int, error) {
	return _Servicekeyresolver.Contract.SignatureTimeout(&_Servicekeyresolver.CallOpts)
}

// AddKey is a paid mutator transaction binding the contract method 0x871fc606.
//
// Solidity: function addKey(address key, string symbol) returns()
func (_Servicekeyresolver *ServicekeyresolverTransactor) AddKey(opts *bind.TransactOpts, key common.Address, symbol string) (*types.Transaction, error) {
	return _Servicekeyresolver.contract.Transact(opts, "addKey", key, symbol)
}

// AddKey is a paid mutator transaction binding the contract method 0x871fc606.
//
// Solidity: function addKey(address key, string symbol) returns()
func (_Servicekeyresolver *ServicekeyresolverSession) AddKey(key common.Address, symbol string) (*types.Transaction, error) {
	return _Servicekeyresolver.Contract.AddKey(&_Servicekeyresolver.TransactOpts, key, symbol)
}

// AddKey is a paid mutator transaction binding the contract method 0x871fc606.
//
// Solidity: function addKey(address key, string symbol) returns()
func (_Servicekeyresolver *ServicekeyresolverTransactorSession) AddKey(key common.Address, symbol string) (*types.Transaction, error) {
	return _Servicekeyresolver.Contract.AddKey(&_Servicekeyresolver.TransactOpts, key, symbol)
}

// AddKeyDelegated is a paid mutator transaction binding the contract method 0x8a4ac03f.
//
// Solidity: function addKeyDelegated(address associatedAddress, address key, string symbol, uint8 v, bytes32 r, bytes32 s, uint256 timestamp) returns()
func (_Servicekeyresolver *ServicekeyresolverTransactor) AddKeyDelegated(opts *bind.TransactOpts, associatedAddress common.Address, key common.Address, symbol string, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	return _Servicekeyresolver.contract.Transact(opts, "addKeyDelegated", associatedAddress, key, symbol, v, r, s, timestamp)
}

// AddKeyDelegated is a paid mutator transaction binding the contract method 0x8a4ac03f.
//
// Solidity: function addKeyDelegated(address associatedAddress, address key, string symbol, uint8 v, bytes32 r, bytes32 s, uint256 timestamp) returns()
func (_Servicekeyresolver *ServicekeyresolverSession) AddKeyDelegated(associatedAddress common.Address, key common.Address, symbol string, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	return _Servicekeyresolver.Contract.AddKeyDelegated(&_Servicekeyresolver.TransactOpts, associatedAddress, key, symbol, v, r, s, timestamp)
}

// AddKeyDelegated is a paid mutator transaction binding the contract method 0x8a4ac03f.
//
// Solidity: function addKeyDelegated(address associatedAddress, address key, string symbol, uint8 v, bytes32 r, bytes32 s, uint256 timestamp) returns()
func (_Servicekeyresolver *ServicekeyresolverTransactorSession) AddKeyDelegated(associatedAddress common.Address, key common.Address, symbol string, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	return _Servicekeyresolver.Contract.AddKeyDelegated(&_Servicekeyresolver.TransactOpts, associatedAddress, key, symbol, v, r, s, timestamp)
}

// RemoveKey is a paid mutator transaction binding the contract method 0x69e78499.
//
// Solidity: function removeKey(address key) returns()
func (_Servicekeyresolver *ServicekeyresolverTransactor) RemoveKey(opts *bind.TransactOpts, key common.Address) (*types.Transaction, error) {
	return _Servicekeyresolver.contract.Transact(opts, "removeKey", key)
}

// RemoveKey is a paid mutator transaction binding the contract method 0x69e78499.
//
// Solidity: function removeKey(address key) returns()
func (_Servicekeyresolver *ServicekeyresolverSession) RemoveKey(key common.Address) (*types.Transaction, error) {
	return _Servicekeyresolver.Contract.RemoveKey(&_Servicekeyresolver.TransactOpts, key)
}

// RemoveKey is a paid mutator transaction binding the contract method 0x69e78499.
//
// Solidity: function removeKey(address key) returns()
func (_Servicekeyresolver *ServicekeyresolverTransactorSession) RemoveKey(key common.Address) (*types.Transaction, error) {
	return _Servicekeyresolver.Contract.RemoveKey(&_Servicekeyresolver.TransactOpts, key)
}

// RemoveKeyDelegated is a paid mutator transaction binding the contract method 0x35434b5d.
//
// Solidity: function removeKeyDelegated(address associatedAddress, address key, uint8 v, bytes32 r, bytes32 s, uint256 timestamp) returns()
func (_Servicekeyresolver *ServicekeyresolverTransactor) RemoveKeyDelegated(opts *bind.TransactOpts, associatedAddress common.Address, key common.Address, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	return _Servicekeyresolver.contract.Transact(opts, "removeKeyDelegated", associatedAddress, key, v, r, s, timestamp)
}

// RemoveKeyDelegated is a paid mutator transaction binding the contract method 0x35434b5d.
//
// Solidity: function removeKeyDelegated(address associatedAddress, address key, uint8 v, bytes32 r, bytes32 s, uint256 timestamp) returns()
func (_Servicekeyresolver *ServicekeyresolverSession) RemoveKeyDelegated(associatedAddress common.Address, key common.Address, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	return _Servicekeyresolver.Contract.RemoveKeyDelegated(&_Servicekeyresolver.TransactOpts, associatedAddress, key, v, r, s, timestamp)
}

// RemoveKeyDelegated is a paid mutator transaction binding the contract method 0x35434b5d.
//
// Solidity: function removeKeyDelegated(address associatedAddress, address key, uint8 v, bytes32 r, bytes32 s, uint256 timestamp) returns()
func (_Servicekeyresolver *ServicekeyresolverTransactorSession) RemoveKeyDelegated(associatedAddress common.Address, key common.Address, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	return _Servicekeyresolver.Contract.RemoveKeyDelegated(&_Servicekeyresolver.TransactOpts, associatedAddress, key, v, r, s, timestamp)
}

// RemoveKeys is a paid mutator transaction binding the contract method 0xa36fee17.
//
// Solidity: function removeKeys() returns()
func (_Servicekeyresolver *ServicekeyresolverTransactor) RemoveKeys(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Servicekeyresolver.contract.Transact(opts, "removeKeys")
}

// RemoveKeys is a paid mutator transaction binding the contract method 0xa36fee17.
//
// Solidity: function removeKeys() returns()
func (_Servicekeyresolver *ServicekeyresolverSession) RemoveKeys() (*types.Transaction, error) {
	return _Servicekeyresolver.Contract.RemoveKeys(&_Servicekeyresolver.TransactOpts)
}

// RemoveKeys is a paid mutator transaction binding the contract method 0xa36fee17.
//
// Solidity: function removeKeys() returns()
func (_Servicekeyresolver *ServicekeyresolverTransactorSession) RemoveKeys() (*types.Transaction, error) {
	return _Servicekeyresolver.Contract.RemoveKeys(&_Servicekeyresolver.TransactOpts)
}

// RemoveKeysDelegated is a paid mutator transaction binding the contract method 0x60066726.
//
// Solidity: function removeKeysDelegated(address associatedAddress, uint8 v, bytes32 r, bytes32 s, uint256 timestamp) returns()
func (_Servicekeyresolver *ServicekeyresolverTransactor) RemoveKeysDelegated(opts *bind.TransactOpts, associatedAddress common.Address, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	return _Servicekeyresolver.contract.Transact(opts, "removeKeysDelegated", associatedAddress, v, r, s, timestamp)
}

// RemoveKeysDelegated is a paid mutator transaction binding the contract method 0x60066726.
//
// Solidity: function removeKeysDelegated(address associatedAddress, uint8 v, bytes32 r, bytes32 s, uint256 timestamp) returns()
func (_Servicekeyresolver *ServicekeyresolverSession) RemoveKeysDelegated(associatedAddress common.Address, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	return _Servicekeyresolver.Contract.RemoveKeysDelegated(&_Servicekeyresolver.TransactOpts, associatedAddress, v, r, s, timestamp)
}

// RemoveKeysDelegated is a paid mutator transaction binding the contract method 0x60066726.
//
// Solidity: function removeKeysDelegated(address associatedAddress, uint8 v, bytes32 r, bytes32 s, uint256 timestamp) returns()
func (_Servicekeyresolver *ServicekeyresolverTransactorSession) RemoveKeysDelegated(associatedAddress common.Address, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	return _Servicekeyresolver.Contract.RemoveKeysDelegated(&_Servicekeyresolver.TransactOpts, associatedAddress, v, r, s, timestamp)
}

// ServicekeyresolverKeyAddedIterator is returned from FilterKeyAdded and is used to iterate over the raw logs and unpacked data for KeyAdded events raised by the Servicekeyresolver contract.
type ServicekeyresolverKeyAddedIterator struct {
	Event *ServicekeyresolverKeyAdded // Event containing the contract specifics and raw log

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
func (it *ServicekeyresolverKeyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServicekeyresolverKeyAdded)
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
		it.Event = new(ServicekeyresolverKeyAdded)
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
func (it *ServicekeyresolverKeyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServicekeyresolverKeyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServicekeyresolverKeyAdded represents a KeyAdded event raised by the Servicekeyresolver contract.
type ServicekeyresolverKeyAdded struct {
	Key    common.Address
	Ein    *big.Int
	Symbol string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterKeyAdded is a free log retrieval operation binding the contract event 0x6c41fd67f4ec3e259af9b0e4b420fa3e24388bd588ca87e0c54522a001c9b686.
//
// Solidity: event KeyAdded(address indexed key, uint256 indexed ein, string symbol)
func (_Servicekeyresolver *ServicekeyresolverFilterer) FilterKeyAdded(opts *bind.FilterOpts, key []common.Address, ein []*big.Int) (*ServicekeyresolverKeyAddedIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}
	var einRule []interface{}
	for _, einItem := range ein {
		einRule = append(einRule, einItem)
	}

	logs, sub, err := _Servicekeyresolver.contract.FilterLogs(opts, "KeyAdded", keyRule, einRule)
	if err != nil {
		return nil, err
	}
	return &ServicekeyresolverKeyAddedIterator{contract: _Servicekeyresolver.contract, event: "KeyAdded", logs: logs, sub: sub}, nil
}

// WatchKeyAdded is a free log subscription operation binding the contract event 0x6c41fd67f4ec3e259af9b0e4b420fa3e24388bd588ca87e0c54522a001c9b686.
//
// Solidity: event KeyAdded(address indexed key, uint256 indexed ein, string symbol)
func (_Servicekeyresolver *ServicekeyresolverFilterer) WatchKeyAdded(opts *bind.WatchOpts, sink chan<- *ServicekeyresolverKeyAdded, key []common.Address, ein []*big.Int) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}
	var einRule []interface{}
	for _, einItem := range ein {
		einRule = append(einRule, einItem)
	}

	logs, sub, err := _Servicekeyresolver.contract.WatchLogs(opts, "KeyAdded", keyRule, einRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServicekeyresolverKeyAdded)
				if err := _Servicekeyresolver.contract.UnpackLog(event, "KeyAdded", log); err != nil {
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

// ServicekeyresolverKeyRemovedIterator is returned from FilterKeyRemoved and is used to iterate over the raw logs and unpacked data for KeyRemoved events raised by the Servicekeyresolver contract.
type ServicekeyresolverKeyRemovedIterator struct {
	Event *ServicekeyresolverKeyRemoved // Event containing the contract specifics and raw log

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
func (it *ServicekeyresolverKeyRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServicekeyresolverKeyRemoved)
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
		it.Event = new(ServicekeyresolverKeyRemoved)
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
func (it *ServicekeyresolverKeyRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServicekeyresolverKeyRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServicekeyresolverKeyRemoved represents a KeyRemoved event raised by the Servicekeyresolver contract.
type ServicekeyresolverKeyRemoved struct {
	Key common.Address
	Ein *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterKeyRemoved is a free log retrieval operation binding the contract event 0xe96ba5805e91ce4b5225d90ad1aac15c207472188f51f24025974341360f0f8a.
//
// Solidity: event KeyRemoved(address indexed key, uint256 indexed ein)
func (_Servicekeyresolver *ServicekeyresolverFilterer) FilterKeyRemoved(opts *bind.FilterOpts, key []common.Address, ein []*big.Int) (*ServicekeyresolverKeyRemovedIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}
	var einRule []interface{}
	for _, einItem := range ein {
		einRule = append(einRule, einItem)
	}

	logs, sub, err := _Servicekeyresolver.contract.FilterLogs(opts, "KeyRemoved", keyRule, einRule)
	if err != nil {
		return nil, err
	}
	return &ServicekeyresolverKeyRemovedIterator{contract: _Servicekeyresolver.contract, event: "KeyRemoved", logs: logs, sub: sub}, nil
}

// WatchKeyRemoved is a free log subscription operation binding the contract event 0xe96ba5805e91ce4b5225d90ad1aac15c207472188f51f24025974341360f0f8a.
//
// Solidity: event KeyRemoved(address indexed key, uint256 indexed ein)
func (_Servicekeyresolver *ServicekeyresolverFilterer) WatchKeyRemoved(opts *bind.WatchOpts, sink chan<- *ServicekeyresolverKeyRemoved, key []common.Address, ein []*big.Int) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}
	var einRule []interface{}
	for _, einItem := range ein {
		einRule = append(einRule, einItem)
	}

	logs, sub, err := _Servicekeyresolver.contract.WatchLogs(opts, "KeyRemoved", keyRule, einRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServicekeyresolverKeyRemoved)
				if err := _Servicekeyresolver.contract.UnpackLog(event, "KeyRemoved", log); err != nil {
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
