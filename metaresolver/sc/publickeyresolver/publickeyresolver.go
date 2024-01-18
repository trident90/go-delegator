// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package publickeyresolver

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

// PublickeyresolverABI is the input ABI used to generate the binding from.
const PublickeyresolverABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"signatureTimeout\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"isSigned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"identityRegistryAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"ein\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"publicKey\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"delegated\",\"type\":\"bool\"}],\"name\":\"PublicKeyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"ein\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"delegated\",\"type\":\"bool\"}],\"name\":\"PublicKeyRemoved\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"calculateAddress\",\"outputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"associatedAddress\",\"type\":\"address\"},{\"name\":\"publicKey\",\"type\":\"bytes\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"addPublicKeyDelegated\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"addPublicKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"associatedAddress\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"removePublicKeyDelegated\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removePublicKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getPublicKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PublickeyresolverBin is the compiled bytecode used for deploying new contracts.
const PublickeyresolverBin = `60806040526201518060025534801561001757600080fd5b506040516020806114ae8339810180604052602081101561003757600080fd5b505160008054600160a060020a03909216600160a060020a0319909216919091179055611445806100696000396000f3fe6080604052600436106100985763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663173c5249811461009d5780633c31b7ea146100ed5780634e21a5171461018f5780635437b67c1461020c578063857cdbb8146102335780638677ebe8146102db578063a061bc741461033d578063a3f4df7e14610352578063e8a4c04e14610367575b600080fd5b3480156100a957600080fd5b506100eb600480360360a08110156100c057600080fd5b50600160a060020a038135169060ff6020820135169060408101359060608101359060800135610436565b005b3480156100f957600080fd5b506100eb600480360360c081101561011057600080fd5b600160a060020a03823516919081019060408101602082013564010000000081111561013b57600080fd5b82018360208201111561014d57600080fd5b8035906020019184600183028401116401000000008311171561016f57600080fd5b919350915060ff813516906020810135906040810135906060013561074c565b34801561019b57600080fd5b506100eb600480360360208110156101b257600080fd5b8101906020810181356401000000008111156101cd57600080fd5b8201836020820111156101df57600080fd5b8035906020019184600183028401116401000000008311171561020157600080fd5b509092509050610abc565b34801561021857600080fd5b50610221610b76565b60408051918252519081900360200190f35b34801561023f57600080fd5b506102666004803603602081101561025657600080fd5b5035600160a060020a0316610b7c565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102a0578181015183820152602001610288565b50505050905090810190601f1680156102cd5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156102e757600080fd5b50610329600480360360a08110156102fe57600080fd5b50600160a060020a038135169060208101359060ff6040820135169060608101359060800135610c26565b604080519115158252519081900360200190f35b34801561034957600080fd5b506100eb610c52565b34801561035e57600080fd5b50610266610cd7565b34801561037357600080fd5b5061041a6004803603602081101561038a57600080fd5b8101906020810181356401000000008111156103a557600080fd5b8201836020820111156103b757600080fd5b803590602001918460018302840111640100000000831117156103d957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610d0e945050505050565b60408051600160a060020a039092168252519081900360200190f35b8080421015801561044a5750600254810142105b15156104a0576040805160e560020a62461bcd02815260206004820152601760248201527f54696d657374616d70206973206e6f742076616c69642e000000000000000000604482015290519081900360640190fd5b600080546040805160e060020a6305c62c2f028152600160a060020a038a81166004830152915191909216916305c62c2f916024808301926020929190829003018186803b1580156104f157600080fd5b505afa158015610505573d6000803e3d6000fd5b505050506040513d602081101561051b57600080fd5b5051600054604080517f53a9698a000000000000000000000000000000000000000000000000000000008152600481018490523360248201529051929350600160a060020a03909116916353a9698a91604480820192602092909190829003018186803b15801561058b57600080fd5b505afa15801561059f573d6000803e3d6000fd5b505050506040513d60208110156105b557600080fd5b5051151561060d576040805160e560020a62461bcd02815260206004820152601f60248201527f4f6e6c792070726f76696465722063616e2062652064656c6567617465642e00604482015290519081900360640190fd5b604080517f1900000000000000000000000000000000000000000000000000000000000000602080830191909152600060218301526c0100000000000000000000000030810260228401527f4920617574686f72697a65207468652072656d6f76616c206f6620612070756260368401527f6c6963206b6579206f6e206d7920626568616c662e00000000000000000000006056840152600160a060020a038b1602606b830152607f80830187905283518084039091018152609f90920190925280519101206106e1908890888888610c26565b1515610737576040805160e560020a62461bcd02815260206004820152601260248201527f5065726d697373696f6e2064656e6965642e0000000000000000000000000000604482015290519081900360640190fd5b61074381886001610d1e565b50505050505050565b808042101580156107605750600254810142105b15156107b6576040805160e560020a62461bcd02815260206004820152601760248201527f54696d657374616d70206973206e6f742076616c69642e000000000000000000604482015290519081900360640190fd5b600080546040805160e060020a6305c62c2f028152600160a060020a038c81166004830152915191909216916305c62c2f916024808301926020929190829003018186803b15801561080757600080fd5b505afa15801561081b573d6000803e3d6000fd5b505050506040513d602081101561083157600080fd5b5051600054604080517f53a9698a000000000000000000000000000000000000000000000000000000008152600481018490523360248201529051929350600160a060020a03909116916353a9698a91604480820192602092909190829003018186803b1580156108a157600080fd5b505afa1580156108b5573d6000803e3d6000fd5b505050506040513d60208110156108cb57600080fd5b50511515610923576040805160e560020a62461bcd02815260206004820152601f60248201527f4f6e6c792070726f76696465722063616e2062652064656c6567617465642e00604482015290519081900360640190fd5b6040517f190000000000000000000000000000000000000000000000000000000000000060208201818152600060218401819052306c0100000000000000000000000081810260228701527f4920617574686f72697a6520746865206164646974696f6e206f66206120707560368701527f626c6963206b6579206f6e206d7920626568616c662e000000000000000000006056870152600160a060020a038f1602606c860152610a17948e949385918f918f918c91906080018484808284378083019250505082815260200197505050505050505060405160208183030381529060405280519060200120888888610c26565b1515610a6d576040805160e560020a62461bcd02815260206004820152601260248201527f5065726d697373696f6e2064656e6965642e0000000000000000000000000000604482015290519081900360640190fd5b610ab1818a8a8a8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525060019250610e99915050565b505050505050505050565b6000546040805160e060020a6305c62c2f0281523360048201529051610b7292600160a060020a0316916305c62c2f916024808301926020929190829003018186803b158015610b0b57600080fd5b505afa158015610b1f573d6000803e3d6000fd5b505050506040513d6020811015610b3557600080fd5b5051604080516020601f86018190048102820181019092528481523391869086908190840183828082843760009201829052509250610e99915050565b5050565b60025481565b600160a060020a03811660009081526001602081815260409283902080548451600294821615610100026000190190911693909304601f81018390048302840183019094528383526060939091830182828015610c1a5780601f10610bef57610100808354040283529160200191610c1a565b820191906000526020600020905b815481529060010190602001808311610bfd57829003601f168201915b50505050509050919050565b6000610c3586868686866111f4565b80610c485750610c48868686868661127c565b9695505050505050565b6000546040805160e060020a6305c62c2f0281523360048201529051610cd592600160a060020a0316916305c62c2f916024808301926020929190829003018186803b158015610ca157600080fd5b505afa158015610cb5573d6000803e3d6000fd5b505050506040513d6020811015610ccb57600080fd5b5051336000610d1e565b565b60408051808201909152601181527f5075626c69634b65795265736f6c766572000000000000000000000000000000602082015281565b8051602090910120600081905290565b600054604080517fd4b1cdcc0000000000000000000000000000000000000000000000000000000081526004810186905230602482015290518592600160a060020a03169163d4b1cdcc916044808301926020929190829003018186803b158015610d8857600080fd5b505afa158015610d9c573d6000803e3d6000fd5b505050506040513d6020811015610db257600080fd5b50511515610e30576040805160e560020a62461bcd02815260206004820152603560248201527f5468652063616c6c696e67206964656e7469747920646f6573206e6f7420686160448201527f76652074686973207265736f6c766572207365742e0000000000000000000000606482015290519081900360840190fd5b600160a060020a0383166000908152600160205260408120610e5191611337565b60408051831515815290518591600160a060020a038616917f627d6da185122bce6a368e7e5c48234968d76011fc1277894aa208786d8ece1e9181900360200190a350505050565b600054604080517fd4b1cdcc0000000000000000000000000000000000000000000000000000000081526004810187905230602482015290518692600160a060020a03169163d4b1cdcc916044808301926020929190829003018186803b158015610f0357600080fd5b505afa158015610f17573d6000803e3d6000fd5b505050506040513d6020811015610f2d57600080fd5b50511515610fab576040805160e560020a62461bcd02815260206004820152603560248201527f5468652063616c6c696e67206964656e7469747920646f6573206e6f7420686160448201527f76652074686973207265736f6c766572207365742e0000000000000000000000606482015290519081900360840190fd5b838381600160a060020a0316610fc082610d0e565b600160a060020a03161461106a576040805160e560020a62461bcd02815260206004820152604260248201527f5468652061646472657373206973206e6f74207468652073616d65206173207460448201527f68617420636f6e7665727465642066726f6d20746865207075626c6963206b6560648201527f792e000000000000000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b600160a060020a0386166000908152600160208190526040909120546002918116156101000260001901160415611111576040805160e560020a62461bcd02815260206004820152602160248201527f4b65792077617320616c726561647920616464656420627920736f6d656f6e6560448201527f2e00000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a0386166000908152600160209081526040909120865161113a9288019061137e565b508686600160a060020a03167fc604da33cee952ded8ce0db2b489092b367cc77da1b4965a9b76e48455b4797c8787604051808060200183151515158152602001828103825284818151815260200191508051906020019080838360005b838110156111b0578181015183820152602001611198565b50505050905090810190601f1680156111dd5780820380516001836020036101000a031916815260200191505b50935050505060405180910390a350505050505050565b600085600160a060020a031660018686868660405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa15801561125d573d6000803e3d6000fd5b50505060206040510351600160a060020a031614905095945050505050565b604080518082018252601c8082527f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020808401918252935160009461132c938b9386938c9301918291908083835b602083106112ea5780518252601f1990920191602091820191016112cb565b51815160209384036101000a600019018019909216911617905292019384525060408051808503815293820190528251920191909120915088905087876111f4565b979650505050505050565b50805460018160011615610100020316600290046000825580601f1061135d575061137b565b601f01602090049060005260206000209081019061137b91906113fc565b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106113bf57805160ff19168380011785556113ec565b828001600101855582156113ec579182015b828111156113ec5782518255916020019190600101906113d1565b506113f89291506113fc565b5090565b61141691905b808211156113f85760008155600101611402565b9056fea165627a7a723058209fa1a0f6f83f1c7808aa3b225c1ef27e96bbce10419fbd7f22a75ab1b54b6c800029`

// DeployPublickeyresolver deploys a new Ethereum contract, binding an instance of Publickeyresolver to it.
func DeployPublickeyresolver(auth *bind.TransactOpts, backend bind.ContractBackend, identityRegistryAddress common.Address) (common.Address, *types.Transaction, *Publickeyresolver, error) {
	parsed, err := abi.JSON(strings.NewReader(PublickeyresolverABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PublickeyresolverBin), backend, identityRegistryAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Publickeyresolver{PublickeyresolverCaller: PublickeyresolverCaller{contract: contract}, PublickeyresolverTransactor: PublickeyresolverTransactor{contract: contract}, PublickeyresolverFilterer: PublickeyresolverFilterer{contract: contract}}, nil
}

// Publickeyresolver is an auto generated Go binding around an Ethereum contract.
type Publickeyresolver struct {
	PublickeyresolverCaller     // Read-only binding to the contract
	PublickeyresolverTransactor // Write-only binding to the contract
	PublickeyresolverFilterer   // Log filterer for contract events
}

// PublickeyresolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type PublickeyresolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublickeyresolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PublickeyresolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublickeyresolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PublickeyresolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublickeyresolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PublickeyresolverSession struct {
	Contract     *Publickeyresolver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PublickeyresolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PublickeyresolverCallerSession struct {
	Contract *PublickeyresolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// PublickeyresolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PublickeyresolverTransactorSession struct {
	Contract     *PublickeyresolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PublickeyresolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type PublickeyresolverRaw struct {
	Contract *Publickeyresolver // Generic contract binding to access the raw methods on
}

// PublickeyresolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PublickeyresolverCallerRaw struct {
	Contract *PublickeyresolverCaller // Generic read-only contract binding to access the raw methods on
}

// PublickeyresolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PublickeyresolverTransactorRaw struct {
	Contract *PublickeyresolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPublickeyresolver creates a new instance of Publickeyresolver, bound to a specific deployed contract.
func NewPublickeyresolver(address common.Address, backend bind.ContractBackend) (*Publickeyresolver, error) {
	contract, err := bindPublickeyresolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Publickeyresolver{PublickeyresolverCaller: PublickeyresolverCaller{contract: contract}, PublickeyresolverTransactor: PublickeyresolverTransactor{contract: contract}, PublickeyresolverFilterer: PublickeyresolverFilterer{contract: contract}}, nil
}

// NewPublickeyresolverCaller creates a new read-only instance of Publickeyresolver, bound to a specific deployed contract.
func NewPublickeyresolverCaller(address common.Address, caller bind.ContractCaller) (*PublickeyresolverCaller, error) {
	contract, err := bindPublickeyresolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PublickeyresolverCaller{contract: contract}, nil
}

// NewPublickeyresolverTransactor creates a new write-only instance of Publickeyresolver, bound to a specific deployed contract.
func NewPublickeyresolverTransactor(address common.Address, transactor bind.ContractTransactor) (*PublickeyresolverTransactor, error) {
	contract, err := bindPublickeyresolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PublickeyresolverTransactor{contract: contract}, nil
}

// NewPublickeyresolverFilterer creates a new log filterer instance of Publickeyresolver, bound to a specific deployed contract.
func NewPublickeyresolverFilterer(address common.Address, filterer bind.ContractFilterer) (*PublickeyresolverFilterer, error) {
	contract, err := bindPublickeyresolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PublickeyresolverFilterer{contract: contract}, nil
}

// bindPublickeyresolver binds a generic wrapper to an already deployed contract.
func bindPublickeyresolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PublickeyresolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Publickeyresolver *PublickeyresolverRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	var results []interface{}
	err := _Publickeyresolver.Contract.PublickeyresolverCaller.contract.Call(opts, &results, method, params...)
	result = results[0]
	return err
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Publickeyresolver *PublickeyresolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Publickeyresolver.Contract.PublickeyresolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Publickeyresolver *PublickeyresolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Publickeyresolver.Contract.PublickeyresolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Publickeyresolver *PublickeyresolverCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	var results []interface{}
	err := _Publickeyresolver.Contract.contract.Call(opts, &results, method, params...)
	result = results[0]
	return err
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Publickeyresolver *PublickeyresolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Publickeyresolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Publickeyresolver *PublickeyresolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Publickeyresolver.Contract.contract.Transact(opts, method, params...)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() constant returns(string)
func (_Publickeyresolver *PublickeyresolverCaller) NAME(opts *bind.CallOpts) (string, error) {
	var results []interface{}
	err := _Publickeyresolver.contract.Call(opts, &results, "NAME")
	return results[0].(string), err
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() constant returns(string)
func (_Publickeyresolver *PublickeyresolverSession) NAME() (string, error) {
	return _Publickeyresolver.Contract.NAME(&_Publickeyresolver.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() constant returns(string)
func (_Publickeyresolver *PublickeyresolverCallerSession) NAME() (string, error) {
	return _Publickeyresolver.Contract.NAME(&_Publickeyresolver.CallOpts)
}

// CalculateAddress is a free data retrieval call binding the contract method 0xe8a4c04e.
//
// Solidity: function calculateAddress(bytes publicKey) constant returns(address addr)
func (_Publickeyresolver *PublickeyresolverCaller) CalculateAddress(opts *bind.CallOpts, publicKey []byte) (common.Address, error) {
	var results []interface{}
	err := _Publickeyresolver.contract.Call(opts, &results, "calculateAddress", publicKey)
	return results[0].(common.Address), err
}

// CalculateAddress is a free data retrieval call binding the contract method 0xe8a4c04e.
//
// Solidity: function calculateAddress(bytes publicKey) constant returns(address addr)
func (_Publickeyresolver *PublickeyresolverSession) CalculateAddress(publicKey []byte) (common.Address, error) {
	return _Publickeyresolver.Contract.CalculateAddress(&_Publickeyresolver.CallOpts, publicKey)
}

// CalculateAddress is a free data retrieval call binding the contract method 0xe8a4c04e.
//
// Solidity: function calculateAddress(bytes publicKey) constant returns(address addr)
func (_Publickeyresolver *PublickeyresolverCallerSession) CalculateAddress(publicKey []byte) (common.Address, error) {
	return _Publickeyresolver.Contract.CalculateAddress(&_Publickeyresolver.CallOpts, publicKey)
}

// GetPublicKey is a free data retrieval call binding the contract method 0x857cdbb8.
//
// Solidity: function getPublicKey(address addr) constant returns(bytes)
func (_Publickeyresolver *PublickeyresolverCaller) GetPublicKey(opts *bind.CallOpts, addr common.Address) ([]byte, error) {
	var results []interface{}
	err := _Publickeyresolver.contract.Call(opts, &results, "getPublicKey", addr)
	return results[0].([]byte), err
}

// GetPublicKey is a free data retrieval call binding the contract method 0x857cdbb8.
//
// Solidity: function getPublicKey(address addr) constant returns(bytes)
func (_Publickeyresolver *PublickeyresolverSession) GetPublicKey(addr common.Address) ([]byte, error) {
	return _Publickeyresolver.Contract.GetPublicKey(&_Publickeyresolver.CallOpts, addr)
}

// GetPublicKey is a free data retrieval call binding the contract method 0x857cdbb8.
//
// Solidity: function getPublicKey(address addr) constant returns(bytes)
func (_Publickeyresolver *PublickeyresolverCallerSession) GetPublicKey(addr common.Address) ([]byte, error) {
	return _Publickeyresolver.Contract.GetPublicKey(&_Publickeyresolver.CallOpts, addr)
}

// IsSigned is a free data retrieval call binding the contract method 0x8677ebe8.
//
// Solidity: function isSigned(address _address, bytes32 messageHash, uint8 v, bytes32 r, bytes32 s) constant returns(bool)
func (_Publickeyresolver *PublickeyresolverCaller) IsSigned(opts *bind.CallOpts, _address common.Address, messageHash [32]byte, v uint8, r [32]byte, s [32]byte) (bool, error) {
	var results []interface{}
	err := _Publickeyresolver.contract.Call(opts, &results, "isSigned", _address, messageHash, v, r, s)
	return results[0].(bool), err
}

// IsSigned is a free data retrieval call binding the contract method 0x8677ebe8.
//
// Solidity: function isSigned(address _address, bytes32 messageHash, uint8 v, bytes32 r, bytes32 s) constant returns(bool)
func (_Publickeyresolver *PublickeyresolverSession) IsSigned(_address common.Address, messageHash [32]byte, v uint8, r [32]byte, s [32]byte) (bool, error) {
	return _Publickeyresolver.Contract.IsSigned(&_Publickeyresolver.CallOpts, _address, messageHash, v, r, s)
}

// IsSigned is a free data retrieval call binding the contract method 0x8677ebe8.
//
// Solidity: function isSigned(address _address, bytes32 messageHash, uint8 v, bytes32 r, bytes32 s) constant returns(bool)
func (_Publickeyresolver *PublickeyresolverCallerSession) IsSigned(_address common.Address, messageHash [32]byte, v uint8, r [32]byte, s [32]byte) (bool, error) {
	return _Publickeyresolver.Contract.IsSigned(&_Publickeyresolver.CallOpts, _address, messageHash, v, r, s)
}

// SignatureTimeout is a free data retrieval call binding the contract method 0x5437b67c.
//
// Solidity: function signatureTimeout() constant returns(uint256)
func (_Publickeyresolver *PublickeyresolverCaller) SignatureTimeout(opts *bind.CallOpts) (*big.Int, error) {
	var results []interface{}
	err := _Publickeyresolver.contract.Call(opts, &results, "signatureTimeout")
	return results[0].(*big.Int), err
}

// SignatureTimeout is a free data retrieval call binding the contract method 0x5437b67c.
//
// Solidity: function signatureTimeout() constant returns(uint256)
func (_Publickeyresolver *PublickeyresolverSession) SignatureTimeout() (*big.Int, error) {
	return _Publickeyresolver.Contract.SignatureTimeout(&_Publickeyresolver.CallOpts)
}

// SignatureTimeout is a free data retrieval call binding the contract method 0x5437b67c.
//
// Solidity: function signatureTimeout() constant returns(uint256)
func (_Publickeyresolver *PublickeyresolverCallerSession) SignatureTimeout() (*big.Int, error) {
	return _Publickeyresolver.Contract.SignatureTimeout(&_Publickeyresolver.CallOpts)
}

// AddPublicKey is a paid mutator transaction binding the contract method 0x4e21a517.
//
// Solidity: function addPublicKey(bytes publicKey) returns()
func (_Publickeyresolver *PublickeyresolverTransactor) AddPublicKey(opts *bind.TransactOpts, publicKey []byte) (*types.Transaction, error) {
	return _Publickeyresolver.contract.Transact(opts, "addPublicKey", publicKey)
}

// AddPublicKey is a paid mutator transaction binding the contract method 0x4e21a517.
//
// Solidity: function addPublicKey(bytes publicKey) returns()
func (_Publickeyresolver *PublickeyresolverSession) AddPublicKey(publicKey []byte) (*types.Transaction, error) {
	return _Publickeyresolver.Contract.AddPublicKey(&_Publickeyresolver.TransactOpts, publicKey)
}

// AddPublicKey is a paid mutator transaction binding the contract method 0x4e21a517.
//
// Solidity: function addPublicKey(bytes publicKey) returns()
func (_Publickeyresolver *PublickeyresolverTransactorSession) AddPublicKey(publicKey []byte) (*types.Transaction, error) {
	return _Publickeyresolver.Contract.AddPublicKey(&_Publickeyresolver.TransactOpts, publicKey)
}

// AddPublicKeyDelegated is a paid mutator transaction binding the contract method 0x3c31b7ea.
//
// Solidity: function addPublicKeyDelegated(address associatedAddress, bytes publicKey, uint8 v, bytes32 r, bytes32 s, uint256 timestamp) returns()
func (_Publickeyresolver *PublickeyresolverTransactor) AddPublicKeyDelegated(opts *bind.TransactOpts, associatedAddress common.Address, publicKey []byte, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	return _Publickeyresolver.contract.Transact(opts, "addPublicKeyDelegated", associatedAddress, publicKey, v, r, s, timestamp)
}

// AddPublicKeyDelegated is a paid mutator transaction binding the contract method 0x3c31b7ea.
//
// Solidity: function addPublicKeyDelegated(address associatedAddress, bytes publicKey, uint8 v, bytes32 r, bytes32 s, uint256 timestamp) returns()
func (_Publickeyresolver *PublickeyresolverSession) AddPublicKeyDelegated(associatedAddress common.Address, publicKey []byte, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	return _Publickeyresolver.Contract.AddPublicKeyDelegated(&_Publickeyresolver.TransactOpts, associatedAddress, publicKey, v, r, s, timestamp)
}

// AddPublicKeyDelegated is a paid mutator transaction binding the contract method 0x3c31b7ea.
//
// Solidity: function addPublicKeyDelegated(address associatedAddress, bytes publicKey, uint8 v, bytes32 r, bytes32 s, uint256 timestamp) returns()
func (_Publickeyresolver *PublickeyresolverTransactorSession) AddPublicKeyDelegated(associatedAddress common.Address, publicKey []byte, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	return _Publickeyresolver.Contract.AddPublicKeyDelegated(&_Publickeyresolver.TransactOpts, associatedAddress, publicKey, v, r, s, timestamp)
}

// RemovePublicKey is a paid mutator transaction binding the contract method 0xa061bc74.
//
// Solidity: function removePublicKey() returns()
func (_Publickeyresolver *PublickeyresolverTransactor) RemovePublicKey(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Publickeyresolver.contract.Transact(opts, "removePublicKey")
}

// RemovePublicKey is a paid mutator transaction binding the contract method 0xa061bc74.
//
// Solidity: function removePublicKey() returns()
func (_Publickeyresolver *PublickeyresolverSession) RemovePublicKey() (*types.Transaction, error) {
	return _Publickeyresolver.Contract.RemovePublicKey(&_Publickeyresolver.TransactOpts)
}

// RemovePublicKey is a paid mutator transaction binding the contract method 0xa061bc74.
//
// Solidity: function removePublicKey() returns()
func (_Publickeyresolver *PublickeyresolverTransactorSession) RemovePublicKey() (*types.Transaction, error) {
	return _Publickeyresolver.Contract.RemovePublicKey(&_Publickeyresolver.TransactOpts)
}

// RemovePublicKeyDelegated is a paid mutator transaction binding the contract method 0x173c5249.
//
// Solidity: function removePublicKeyDelegated(address associatedAddress, uint8 v, bytes32 r, bytes32 s, uint256 timestamp) returns()
func (_Publickeyresolver *PublickeyresolverTransactor) RemovePublicKeyDelegated(opts *bind.TransactOpts, associatedAddress common.Address, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	return _Publickeyresolver.contract.Transact(opts, "removePublicKeyDelegated", associatedAddress, v, r, s, timestamp)
}

// RemovePublicKeyDelegated is a paid mutator transaction binding the contract method 0x173c5249.
//
// Solidity: function removePublicKeyDelegated(address associatedAddress, uint8 v, bytes32 r, bytes32 s, uint256 timestamp) returns()
func (_Publickeyresolver *PublickeyresolverSession) RemovePublicKeyDelegated(associatedAddress common.Address, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	return _Publickeyresolver.Contract.RemovePublicKeyDelegated(&_Publickeyresolver.TransactOpts, associatedAddress, v, r, s, timestamp)
}

// RemovePublicKeyDelegated is a paid mutator transaction binding the contract method 0x173c5249.
//
// Solidity: function removePublicKeyDelegated(address associatedAddress, uint8 v, bytes32 r, bytes32 s, uint256 timestamp) returns()
func (_Publickeyresolver *PublickeyresolverTransactorSession) RemovePublicKeyDelegated(associatedAddress common.Address, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	return _Publickeyresolver.Contract.RemovePublicKeyDelegated(&_Publickeyresolver.TransactOpts, associatedAddress, v, r, s, timestamp)
}

// PublickeyresolverPublicKeyAddedIterator is returned from FilterPublicKeyAdded and is used to iterate over the raw logs and unpacked data for PublicKeyAdded events raised by the Publickeyresolver contract.
type PublickeyresolverPublicKeyAddedIterator struct {
	Event *PublickeyresolverPublicKeyAdded // Event containing the contract specifics and raw log

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
func (it *PublickeyresolverPublicKeyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublickeyresolverPublicKeyAdded)
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
		it.Event = new(PublickeyresolverPublicKeyAdded)
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
func (it *PublickeyresolverPublicKeyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublickeyresolverPublicKeyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublickeyresolverPublicKeyAdded represents a PublicKeyAdded event raised by the Publickeyresolver contract.
type PublickeyresolverPublicKeyAdded struct {
	Addr      common.Address
	Ein       *big.Int
	PublicKey []byte
	Delegated bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPublicKeyAdded is a free log retrieval operation binding the contract event 0xc604da33cee952ded8ce0db2b489092b367cc77da1b4965a9b76e48455b4797c.
//
// Solidity: event PublicKeyAdded(address indexed addr, uint256 indexed ein, bytes publicKey, bool delegated)
func (_Publickeyresolver *PublickeyresolverFilterer) FilterPublicKeyAdded(opts *bind.FilterOpts, addr []common.Address, ein []*big.Int) (*PublickeyresolverPublicKeyAddedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var einRule []interface{}
	for _, einItem := range ein {
		einRule = append(einRule, einItem)
	}

	logs, sub, err := _Publickeyresolver.contract.FilterLogs(opts, "PublicKeyAdded", addrRule, einRule)
	if err != nil {
		return nil, err
	}
	return &PublickeyresolverPublicKeyAddedIterator{contract: _Publickeyresolver.contract, event: "PublicKeyAdded", logs: logs, sub: sub}, nil
}

// WatchPublicKeyAdded is a free log subscription operation binding the contract event 0xc604da33cee952ded8ce0db2b489092b367cc77da1b4965a9b76e48455b4797c.
//
// Solidity: event PublicKeyAdded(address indexed addr, uint256 indexed ein, bytes publicKey, bool delegated)
func (_Publickeyresolver *PublickeyresolverFilterer) WatchPublicKeyAdded(opts *bind.WatchOpts, sink chan<- *PublickeyresolverPublicKeyAdded, addr []common.Address, ein []*big.Int) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var einRule []interface{}
	for _, einItem := range ein {
		einRule = append(einRule, einItem)
	}

	logs, sub, err := _Publickeyresolver.contract.WatchLogs(opts, "PublicKeyAdded", addrRule, einRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublickeyresolverPublicKeyAdded)
				if err := _Publickeyresolver.contract.UnpackLog(event, "PublicKeyAdded", log); err != nil {
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

// PublickeyresolverPublicKeyRemovedIterator is returned from FilterPublicKeyRemoved and is used to iterate over the raw logs and unpacked data for PublicKeyRemoved events raised by the Publickeyresolver contract.
type PublickeyresolverPublicKeyRemovedIterator struct {
	Event *PublickeyresolverPublicKeyRemoved // Event containing the contract specifics and raw log

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
func (it *PublickeyresolverPublicKeyRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublickeyresolverPublicKeyRemoved)
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
		it.Event = new(PublickeyresolverPublicKeyRemoved)
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
func (it *PublickeyresolverPublicKeyRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublickeyresolverPublicKeyRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublickeyresolverPublicKeyRemoved represents a PublicKeyRemoved event raised by the Publickeyresolver contract.
type PublickeyresolverPublicKeyRemoved struct {
	Addr      common.Address
	Ein       *big.Int
	Delegated bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPublicKeyRemoved is a free log retrieval operation binding the contract event 0x627d6da185122bce6a368e7e5c48234968d76011fc1277894aa208786d8ece1e.
//
// Solidity: event PublicKeyRemoved(address indexed addr, uint256 indexed ein, bool delegated)
func (_Publickeyresolver *PublickeyresolverFilterer) FilterPublicKeyRemoved(opts *bind.FilterOpts, addr []common.Address, ein []*big.Int) (*PublickeyresolverPublicKeyRemovedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var einRule []interface{}
	for _, einItem := range ein {
		einRule = append(einRule, einItem)
	}

	logs, sub, err := _Publickeyresolver.contract.FilterLogs(opts, "PublicKeyRemoved", addrRule, einRule)
	if err != nil {
		return nil, err
	}
	return &PublickeyresolverPublicKeyRemovedIterator{contract: _Publickeyresolver.contract, event: "PublicKeyRemoved", logs: logs, sub: sub}, nil
}

// WatchPublicKeyRemoved is a free log subscription operation binding the contract event 0x627d6da185122bce6a368e7e5c48234968d76011fc1277894aa208786d8ece1e.
//
// Solidity: event PublicKeyRemoved(address indexed addr, uint256 indexed ein, bool delegated)
func (_Publickeyresolver *PublickeyresolverFilterer) WatchPublicKeyRemoved(opts *bind.WatchOpts, sink chan<- *PublickeyresolverPublicKeyRemoved, addr []common.Address, ein []*big.Int) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var einRule []interface{}
	for _, einItem := range ein {
		einRule = append(einRule, einItem)
	}

	logs, sub, err := _Publickeyresolver.contract.WatchLogs(opts, "PublicKeyRemoved", addrRule, einRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublickeyresolverPublicKeyRemoved)
				if err := _Publickeyresolver.contract.UnpackLog(event, "PublicKeyRemoved", log); err != nil {
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
