// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nameservice

import (
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// NameserviceABI is the input ABI used to generate the binding from.
const NameserviceABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"permissions\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"contracts\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"bytes32\"},{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setContractDomain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"getContractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contract\",\"type\":\"bytes32\"},{\"name\":\"_granted\",\"type\":\"address\"},{\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"setPermission\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_contract\",\"type\":\"bytes32\"},{\"name\":\"_granted\",\"type\":\"address\"}],\"name\":\"getPermission\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// NameserviceBin is the compiled bytecode used for deploying new contracts.
const NameserviceBin = `6080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506108e3806100536000396000f30060806040526004361061008e576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806304af66ad146100935780630d2020dd146100e45780633ec50c6c14610155578063599e4c70146101be57806360d6c7cf146102335780638da5cb5b1461029c578063ec56a373146102f3578063f2fde38b14610364575b600080fd5b34801561009f57600080fd5b506100e26004803603810190808035600019169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506103a7565b005b3480156100f057600080fd5b50610113600480360381019080803560001916906020019092919050505061049c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561016157600080fd5b506101a46004803603810190808035600019169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610557565b604051808215151515815260200191505060405180910390f35b3480156101ca57600080fd5b506102196004803603810190808035600019169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803515159060200190929190505050610586565b604051808215151515815260200191505060405180910390f35b34801561023f57600080fd5b506102826004803603810190808035600019169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061069a565b604051808215151515815260200191505060405180910390f35b3480156102a857600080fd5b506102b161070a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156102ff57600080fd5b50610322600480360381019080803560001916906020019092919050505061072f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561037057600080fd5b506103a5600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610762565b005b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561040257600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561043e57600080fd5b8060016000846000191660001916815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b60008073ffffffffffffffffffffffffffffffffffffffff1660016000846000191660001916815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415151561051457600080fd5b60016000836000191660001916815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b60026020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156105e357600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415151561061f57600080fd5b8160026000866000191660001916815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600190509392505050565b600060026000846000191660001916815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60016020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156107bd57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156107f957600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a723058209197617b7cd7db0615c420312cd98633903150642ed008cb70299b5d64f8a85d0029`

// DeployNameservice deploys a new Ethereum contract, binding an instance of Nameservice to it.
func DeployNameservice(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Nameservice, error) {
	parsed, err := abi.JSON(strings.NewReader(NameserviceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NameserviceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Nameservice{NameserviceCaller: NameserviceCaller{contract: contract}, NameserviceTransactor: NameserviceTransactor{contract: contract}, NameserviceFilterer: NameserviceFilterer{contract: contract}}, nil
}

// Nameservice is an auto generated Go binding around an Ethereum contract.
type Nameservice struct {
	NameserviceCaller     // Read-only binding to the contract
	NameserviceTransactor // Write-only binding to the contract
	NameserviceFilterer   // Log filterer for contract events
}

// NameserviceCaller is an auto generated read-only Go binding around an Ethereum contract.
type NameserviceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NameserviceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NameserviceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NameserviceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NameserviceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NameserviceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NameserviceSession struct {
	Contract     *Nameservice      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NameserviceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NameserviceCallerSession struct {
	Contract *NameserviceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// NameserviceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NameserviceTransactorSession struct {
	Contract     *NameserviceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NameserviceRaw is an auto generated low-level Go binding around an Ethereum contract.
type NameserviceRaw struct {
	Contract *Nameservice // Generic contract binding to access the raw methods on
}

// NameserviceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NameserviceCallerRaw struct {
	Contract *NameserviceCaller // Generic read-only contract binding to access the raw methods on
}

// NameserviceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NameserviceTransactorRaw struct {
	Contract *NameserviceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNameservice creates a new instance of Nameservice, bound to a specific deployed contract.
func NewNameservice(address common.Address, backend bind.ContractBackend) (*Nameservice, error) {
	contract, err := bindNameservice(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nameservice{NameserviceCaller: NameserviceCaller{contract: contract}, NameserviceTransactor: NameserviceTransactor{contract: contract}, NameserviceFilterer: NameserviceFilterer{contract: contract}}, nil
}

// NewNameserviceCaller creates a new read-only instance of Nameservice, bound to a specific deployed contract.
func NewNameserviceCaller(address common.Address, caller bind.ContractCaller) (*NameserviceCaller, error) {
	contract, err := bindNameservice(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NameserviceCaller{contract: contract}, nil
}

// NewNameserviceTransactor creates a new write-only instance of Nameservice, bound to a specific deployed contract.
func NewNameserviceTransactor(address common.Address, transactor bind.ContractTransactor) (*NameserviceTransactor, error) {
	contract, err := bindNameservice(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NameserviceTransactor{contract: contract}, nil
}

// NewNameserviceFilterer creates a new log filterer instance of Nameservice, bound to a specific deployed contract.
func NewNameserviceFilterer(address common.Address, filterer bind.ContractFilterer) (*NameserviceFilterer, error) {
	contract, err := bindNameservice(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NameserviceFilterer{contract: contract}, nil
}

// bindNameservice binds a generic wrapper to an already deployed contract.
func bindNameservice(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NameserviceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nameservice *NameserviceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Nameservice.Contract.NameserviceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nameservice *NameserviceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nameservice.Contract.NameserviceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nameservice *NameserviceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nameservice.Contract.NameserviceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nameservice *NameserviceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Nameservice.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nameservice *NameserviceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nameservice.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nameservice *NameserviceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nameservice.Contract.contract.Transact(opts, method, params...)
}

// Contracts is a free data retrieval call binding the contract method 0xec56a373.
//
// Solidity: function contracts( bytes32) constant returns(address)
func (_Nameservice *NameserviceCaller) Contracts(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Nameservice.contract.Call(opts, out, "contracts", arg0)
	return *ret0, err
}

// Contracts is a free data retrieval call binding the contract method 0xec56a373.
//
// Solidity: function contracts( bytes32) constant returns(address)
func (_Nameservice *NameserviceSession) Contracts(arg0 [32]byte) (common.Address, error) {
	return _Nameservice.Contract.Contracts(&_Nameservice.CallOpts, arg0)
}

// Contracts is a free data retrieval call binding the contract method 0xec56a373.
//
// Solidity: function contracts( bytes32) constant returns(address)
func (_Nameservice *NameserviceCallerSession) Contracts(arg0 [32]byte) (common.Address, error) {
	return _Nameservice.Contract.Contracts(&_Nameservice.CallOpts, arg0)
}

// GetContractAddress is a free data retrieval call binding the contract method 0x0d2020dd.
//
// Solidity: function getContractAddress(_name bytes32) constant returns(address)
func (_Nameservice *NameserviceCaller) GetContractAddress(opts *bind.CallOpts, _name [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Nameservice.contract.Call(opts, out, "getContractAddress", _name)
	return *ret0, err
}

// GetContractAddress is a free data retrieval call binding the contract method 0x0d2020dd.
//
// Solidity: function getContractAddress(_name bytes32) constant returns(address)
func (_Nameservice *NameserviceSession) GetContractAddress(_name [32]byte) (common.Address, error) {
	return _Nameservice.Contract.GetContractAddress(&_Nameservice.CallOpts, _name)
}

// GetContractAddress is a free data retrieval call binding the contract method 0x0d2020dd.
//
// Solidity: function getContractAddress(_name bytes32) constant returns(address)
func (_Nameservice *NameserviceCallerSession) GetContractAddress(_name [32]byte) (common.Address, error) {
	return _Nameservice.Contract.GetContractAddress(&_Nameservice.CallOpts, _name)
}

// GetPermission is a free data retrieval call binding the contract method 0x60d6c7cf.
//
// Solidity: function getPermission(_contract bytes32, _granted address) constant returns(bool)
func (_Nameservice *NameserviceCaller) GetPermission(opts *bind.CallOpts, _contract [32]byte, _granted common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Nameservice.contract.Call(opts, out, "getPermission", _contract, _granted)
	return *ret0, err
}

// GetPermission is a free data retrieval call binding the contract method 0x60d6c7cf.
//
// Solidity: function getPermission(_contract bytes32, _granted address) constant returns(bool)
func (_Nameservice *NameserviceSession) GetPermission(_contract [32]byte, _granted common.Address) (bool, error) {
	return _Nameservice.Contract.GetPermission(&_Nameservice.CallOpts, _contract, _granted)
}

// GetPermission is a free data retrieval call binding the contract method 0x60d6c7cf.
//
// Solidity: function getPermission(_contract bytes32, _granted address) constant returns(bool)
func (_Nameservice *NameserviceCallerSession) GetPermission(_contract [32]byte, _granted common.Address) (bool, error) {
	return _Nameservice.Contract.GetPermission(&_Nameservice.CallOpts, _contract, _granted)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Nameservice *NameserviceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Nameservice.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Nameservice *NameserviceSession) Owner() (common.Address, error) {
	return _Nameservice.Contract.Owner(&_Nameservice.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Nameservice *NameserviceCallerSession) Owner() (common.Address, error) {
	return _Nameservice.Contract.Owner(&_Nameservice.CallOpts)
}

// Permissions is a free data retrieval call binding the contract method 0x3ec50c6c.
//
// Solidity: function permissions( bytes32,  address) constant returns(bool)
func (_Nameservice *NameserviceCaller) Permissions(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Nameservice.contract.Call(opts, out, "permissions", arg0, arg1)
	return *ret0, err
}

// Permissions is a free data retrieval call binding the contract method 0x3ec50c6c.
//
// Solidity: function permissions( bytes32,  address) constant returns(bool)
func (_Nameservice *NameserviceSession) Permissions(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _Nameservice.Contract.Permissions(&_Nameservice.CallOpts, arg0, arg1)
}

// Permissions is a free data retrieval call binding the contract method 0x3ec50c6c.
//
// Solidity: function permissions( bytes32,  address) constant returns(bool)
func (_Nameservice *NameserviceCallerSession) Permissions(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _Nameservice.Contract.Permissions(&_Nameservice.CallOpts, arg0, arg1)
}

// SetContractDomain is a paid mutator transaction binding the contract method 0x04af66ad.
//
// Solidity: function setContractDomain(_name bytes32, _addr address) returns()
func (_Nameservice *NameserviceTransactor) SetContractDomain(opts *bind.TransactOpts, _name [32]byte, _addr common.Address) (*types.Transaction, error) {
	return _Nameservice.contract.Transact(opts, "setContractDomain", _name, _addr)
}

// SetContractDomain is a paid mutator transaction binding the contract method 0x04af66ad.
//
// Solidity: function setContractDomain(_name bytes32, _addr address) returns()
func (_Nameservice *NameserviceSession) SetContractDomain(_name [32]byte, _addr common.Address) (*types.Transaction, error) {
	return _Nameservice.Contract.SetContractDomain(&_Nameservice.TransactOpts, _name, _addr)
}

// SetContractDomain is a paid mutator transaction binding the contract method 0x04af66ad.
//
// Solidity: function setContractDomain(_name bytes32, _addr address) returns()
func (_Nameservice *NameserviceTransactorSession) SetContractDomain(_name [32]byte, _addr common.Address) (*types.Transaction, error) {
	return _Nameservice.Contract.SetContractDomain(&_Nameservice.TransactOpts, _name, _addr)
}

// SetPermission is a paid mutator transaction binding the contract method 0x599e4c70.
//
// Solidity: function setPermission(_contract bytes32, _granted address, _status bool) returns(bool)
func (_Nameservice *NameserviceTransactor) SetPermission(opts *bind.TransactOpts, _contract [32]byte, _granted common.Address, _status bool) (*types.Transaction, error) {
	return _Nameservice.contract.Transact(opts, "setPermission", _contract, _granted, _status)
}

// SetPermission is a paid mutator transaction binding the contract method 0x599e4c70.
//
// Solidity: function setPermission(_contract bytes32, _granted address, _status bool) returns(bool)
func (_Nameservice *NameserviceSession) SetPermission(_contract [32]byte, _granted common.Address, _status bool) (*types.Transaction, error) {
	return _Nameservice.Contract.SetPermission(&_Nameservice.TransactOpts, _contract, _granted, _status)
}

// SetPermission is a paid mutator transaction binding the contract method 0x599e4c70.
//
// Solidity: function setPermission(_contract bytes32, _granted address, _status bool) returns(bool)
func (_Nameservice *NameserviceTransactorSession) SetPermission(_contract [32]byte, _granted common.Address, _status bool) (*types.Transaction, error) {
	return _Nameservice.Contract.SetPermission(&_Nameservice.TransactOpts, _contract, _granted, _status)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Nameservice *NameserviceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Nameservice.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Nameservice *NameserviceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Nameservice.Contract.TransferOwnership(&_Nameservice.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Nameservice *NameserviceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Nameservice.Contract.TransferOwnership(&_Nameservice.TransactOpts, newOwner)
}

// NameserviceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Nameservice contract.
type NameserviceOwnershipTransferredIterator struct {
	Event *NameserviceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NameserviceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameserviceOwnershipTransferred)
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
		it.Event = new(NameserviceOwnershipTransferred)
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
func (it *NameserviceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameserviceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameserviceOwnershipTransferred represents a OwnershipTransferred event raised by the Nameservice contract.
type NameserviceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Nameservice *NameserviceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NameserviceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Nameservice.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NameserviceOwnershipTransferredIterator{contract: _Nameservice.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Nameservice *NameserviceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NameserviceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Nameservice.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameserviceOwnershipTransferred)
				if err := _Nameservice.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
