// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registry

import (
	"strings"

	ethereum "github.com/metadium/go-metadium"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// RegistryABI is the input ABI used to generate the binding from.
const RegistryABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"permissions\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"contracts\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"setter\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"SetContractDomain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_contract\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"granted\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"SetPermission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"bytes32\"},{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setContractDomain\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"getContractAddress\",\"outputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contract\",\"type\":\"bytes32\"},{\"name\":\"_granted\",\"type\":\"address\"},{\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"setPermission\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_contract\",\"type\":\"bytes32\"},{\"name\":\"_granted\",\"type\":\"address\"}],\"name\":\"getPermission\",\"outputs\":[{\"name\":\"found\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RegistryBin is the compiled bytecode used for deploying new contracts.
const RegistryBin = `608060405260008054600160a060020a0319163317905561059a806100256000396000f30060806040526004361061008d5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166304af66ad81146100925780630d2020dd146100ca5780633ec50c6c146100fe578063599e4c701461012257806360d6c7cf1461014b5780638da5cb5b1461016f578063ec56a37314610184578063f2fde38b1461019c575b600080fd5b34801561009e57600080fd5b506100b6600435600160a060020a03602435166101bf565b604080519115158252519081900360200190f35b3480156100d657600080fd5b506100e26004356102c6565b60408051600160a060020a039092168252519081900360200190f35b34801561010a57600080fd5b506100b6600435600160a060020a0360243516610367565b34801561012e57600080fd5b506100b6600435600160a060020a03602435166044351515610387565b34801561015757600080fd5b506100b6600435600160a060020a0360243516610485565b34801561017b57600080fd5b506100e26104b0565b34801561019057600080fd5b506100e26004356104bf565b3480156101a857600080fd5b506101bd600160a060020a03600435166104da565b005b60008054600160a060020a031633146101d757600080fd5b600160a060020a038216151561024e57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f616464726573732073686f756c64206265206e6f6e2d7a65726f000000000000604482015290519081900360640190fd5b600083815260016020908152604091829020805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03861690811790915582513381529251909286927f37724a4a9968ac9654e6ee52f3d0c93e5ef8863e057254ee2e36e8ad3e8429db92918290030190a350600192915050565b600081815260016020526040812054600160a060020a0316151561034b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f616464726573732073686f756c64206265206e6f6e2d7a65726f000000000000604482015290519081900360640190fd5b50600090815260016020526040902054600160a060020a031690565b600260209081526000928352604080842090915290825290205460ff1681565b60008054600160a060020a0316331461039f57600080fd5b600160a060020a038316151561041657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f616464726573732073686f756c64206265206e6f6e2d7a65726f000000000000604482015290519081900360640190fd5b6000848152600260209081526040808320600160a060020a03871680855290835292819020805460ff19168615159081179091558151908152905187927fe9f5231bbfb4b32867755b94562215cff6c8998489de8ba20926f8d0980e7818928290030190a35060019392505050565b6000918252600260209081526040808420600160a060020a0393909316845291905290205460ff1690565b600054600160a060020a031681565b600160205260009081526040902054600160a060020a031681565b600054600160a060020a031633146104f157600080fd5b600160a060020a038116151561050657600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820fdff7bf855d2278cc858686fbe566ea3437cb39f9567787f1a3866731c2c1ecb0029`

// DeployRegistry deploys a new Ethereum contract, binding an instance of Registry to it.
func DeployRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Registry, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// Contracts is a free data retrieval call binding the contract method 0xec56a373.
//
// Solidity: function contracts( bytes32) constant returns(address)
func (_Registry *RegistryCaller) Contracts(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "contracts", arg0)
	return *ret0, err
}

// Contracts is a free data retrieval call binding the contract method 0xec56a373.
//
// Solidity: function contracts( bytes32) constant returns(address)
func (_Registry *RegistrySession) Contracts(arg0 [32]byte) (common.Address, error) {
	return _Registry.Contract.Contracts(&_Registry.CallOpts, arg0)
}

// Contracts is a free data retrieval call binding the contract method 0xec56a373.
//
// Solidity: function contracts( bytes32) constant returns(address)
func (_Registry *RegistryCallerSession) Contracts(arg0 [32]byte) (common.Address, error) {
	return _Registry.Contract.Contracts(&_Registry.CallOpts, arg0)
}

// GetContractAddress is a free data retrieval call binding the contract method 0x0d2020dd.
//
// Solidity: function getContractAddress(_name bytes32) constant returns(addr address)
func (_Registry *RegistryCaller) GetContractAddress(opts *bind.CallOpts, _name [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "getContractAddress", _name)
	return *ret0, err
}

// GetContractAddress is a free data retrieval call binding the contract method 0x0d2020dd.
//
// Solidity: function getContractAddress(_name bytes32) constant returns(addr address)
func (_Registry *RegistrySession) GetContractAddress(_name [32]byte) (common.Address, error) {
	return _Registry.Contract.GetContractAddress(&_Registry.CallOpts, _name)
}

// GetContractAddress is a free data retrieval call binding the contract method 0x0d2020dd.
//
// Solidity: function getContractAddress(_name bytes32) constant returns(addr address)
func (_Registry *RegistryCallerSession) GetContractAddress(_name [32]byte) (common.Address, error) {
	return _Registry.Contract.GetContractAddress(&_Registry.CallOpts, _name)
}

// GetPermission is a free data retrieval call binding the contract method 0x60d6c7cf.
//
// Solidity: function getPermission(_contract bytes32, _granted address) constant returns(found bool)
func (_Registry *RegistryCaller) GetPermission(opts *bind.CallOpts, _contract [32]byte, _granted common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "getPermission", _contract, _granted)
	return *ret0, err
}

// GetPermission is a free data retrieval call binding the contract method 0x60d6c7cf.
//
// Solidity: function getPermission(_contract bytes32, _granted address) constant returns(found bool)
func (_Registry *RegistrySession) GetPermission(_contract [32]byte, _granted common.Address) (bool, error) {
	return _Registry.Contract.GetPermission(&_Registry.CallOpts, _contract, _granted)
}

// GetPermission is a free data retrieval call binding the contract method 0x60d6c7cf.
//
// Solidity: function getPermission(_contract bytes32, _granted address) constant returns(found bool)
func (_Registry *RegistryCallerSession) GetPermission(_contract [32]byte, _granted common.Address) (bool, error) {
	return _Registry.Contract.GetPermission(&_Registry.CallOpts, _contract, _granted)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Registry *RegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Registry *RegistrySession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Registry *RegistryCallerSession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// Permissions is a free data retrieval call binding the contract method 0x3ec50c6c.
//
// Solidity: function permissions( bytes32,  address) constant returns(bool)
func (_Registry *RegistryCaller) Permissions(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "permissions", arg0, arg1)
	return *ret0, err
}

// Permissions is a free data retrieval call binding the contract method 0x3ec50c6c.
//
// Solidity: function permissions( bytes32,  address) constant returns(bool)
func (_Registry *RegistrySession) Permissions(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _Registry.Contract.Permissions(&_Registry.CallOpts, arg0, arg1)
}

// Permissions is a free data retrieval call binding the contract method 0x3ec50c6c.
//
// Solidity: function permissions( bytes32,  address) constant returns(bool)
func (_Registry *RegistryCallerSession) Permissions(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _Registry.Contract.Permissions(&_Registry.CallOpts, arg0, arg1)
}

// SetContractDomain is a paid mutator transaction binding the contract method 0x04af66ad.
//
// Solidity: function setContractDomain(_name bytes32, _addr address) returns(success bool)
func (_Registry *RegistryTransactor) SetContractDomain(opts *bind.TransactOpts, _name [32]byte, _addr common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setContractDomain", _name, _addr)
}

// SetContractDomain is a paid mutator transaction binding the contract method 0x04af66ad.
//
// Solidity: function setContractDomain(_name bytes32, _addr address) returns(success bool)
func (_Registry *RegistrySession) SetContractDomain(_name [32]byte, _addr common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetContractDomain(&_Registry.TransactOpts, _name, _addr)
}

// SetContractDomain is a paid mutator transaction binding the contract method 0x04af66ad.
//
// Solidity: function setContractDomain(_name bytes32, _addr address) returns(success bool)
func (_Registry *RegistryTransactorSession) SetContractDomain(_name [32]byte, _addr common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetContractDomain(&_Registry.TransactOpts, _name, _addr)
}

// SetPermission is a paid mutator transaction binding the contract method 0x599e4c70.
//
// Solidity: function setPermission(_contract bytes32, _granted address, _status bool) returns(success bool)
func (_Registry *RegistryTransactor) SetPermission(opts *bind.TransactOpts, _contract [32]byte, _granted common.Address, _status bool) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setPermission", _contract, _granted, _status)
}

// SetPermission is a paid mutator transaction binding the contract method 0x599e4c70.
//
// Solidity: function setPermission(_contract bytes32, _granted address, _status bool) returns(success bool)
func (_Registry *RegistrySession) SetPermission(_contract [32]byte, _granted common.Address, _status bool) (*types.Transaction, error) {
	return _Registry.Contract.SetPermission(&_Registry.TransactOpts, _contract, _granted, _status)
}

// SetPermission is a paid mutator transaction binding the contract method 0x599e4c70.
//
// Solidity: function setPermission(_contract bytes32, _granted address, _status bool) returns(success bool)
func (_Registry *RegistryTransactorSession) SetPermission(_contract [32]byte, _granted common.Address, _status bool) (*types.Transaction, error) {
	return _Registry.Contract.SetPermission(&_Registry.TransactOpts, _contract, _granted, _status)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Registry *RegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Registry *RegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwnership(&_Registry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Registry *RegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwnership(&_Registry.TransactOpts, newOwner)
}

// RegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Registry contract.
type RegistryOwnershipTransferredIterator struct {
	Event *RegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryOwnershipTransferred)
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
		it.Event = new(RegistryOwnershipTransferred)
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
func (it *RegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryOwnershipTransferred represents a OwnershipTransferred event raised by the Registry contract.
type RegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Registry *RegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RegistryOwnershipTransferredIterator{contract: _Registry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Registry *RegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryOwnershipTransferred)
				if err := _Registry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// RegistrySetContractDomainIterator is returned from FilterSetContractDomain and is used to iterate over the raw logs and unpacked data for SetContractDomain events raised by the Registry contract.
type RegistrySetContractDomainIterator struct {
	Event *RegistrySetContractDomain // Event containing the contract specifics and raw log

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
func (it *RegistrySetContractDomainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrySetContractDomain)
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
		it.Event = new(RegistrySetContractDomain)
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
func (it *RegistrySetContractDomainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrySetContractDomainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrySetContractDomain represents a SetContractDomain event raised by the Registry contract.
type RegistrySetContractDomain struct {
	Setter common.Address
	Name   [32]byte
	Addr   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetContractDomain is a free log retrieval operation binding the contract event 0x37724a4a9968ac9654e6ee52f3d0c93e5ef8863e057254ee2e36e8ad3e8429db.
//
// Solidity: e SetContractDomain(setter address, name indexed bytes32, addr indexed address)
func (_Registry *RegistryFilterer) FilterSetContractDomain(opts *bind.FilterOpts, name [][32]byte, addr []common.Address) (*RegistrySetContractDomainIterator, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "SetContractDomain", nameRule, addrRule)
	if err != nil {
		return nil, err
	}
	return &RegistrySetContractDomainIterator{contract: _Registry.contract, event: "SetContractDomain", logs: logs, sub: sub}, nil
}

// WatchSetContractDomain is a free log subscription operation binding the contract event 0x37724a4a9968ac9654e6ee52f3d0c93e5ef8863e057254ee2e36e8ad3e8429db.
//
// Solidity: e SetContractDomain(setter address, name indexed bytes32, addr indexed address)
func (_Registry *RegistryFilterer) WatchSetContractDomain(opts *bind.WatchOpts, sink chan<- *RegistrySetContractDomain, name [][32]byte, addr []common.Address) (event.Subscription, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "SetContractDomain", nameRule, addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrySetContractDomain)
				if err := _Registry.contract.UnpackLog(event, "SetContractDomain", log); err != nil {
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

// RegistrySetPermissionIterator is returned from FilterSetPermission and is used to iterate over the raw logs and unpacked data for SetPermission events raised by the Registry contract.
type RegistrySetPermissionIterator struct {
	Event *RegistrySetPermission // Event containing the contract specifics and raw log

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
func (it *RegistrySetPermissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrySetPermission)
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
		it.Event = new(RegistrySetPermission)
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
func (it *RegistrySetPermissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrySetPermissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrySetPermission represents a SetPermission event raised by the Registry contract.
type RegistrySetPermission struct {
	Contract [32]byte
	Granted  common.Address
	Status   bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetPermission is a free log retrieval operation binding the contract event 0xe9f5231bbfb4b32867755b94562215cff6c8998489de8ba20926f8d0980e7818.
//
// Solidity: e SetPermission(_contract indexed bytes32, granted indexed address, status bool)
func (_Registry *RegistryFilterer) FilterSetPermission(opts *bind.FilterOpts, _contract [][32]byte, granted []common.Address) (*RegistrySetPermissionIterator, error) {

	var _contractRule []interface{}
	for _, _contractItem := range _contract {
		_contractRule = append(_contractRule, _contractItem)
	}
	var grantedRule []interface{}
	for _, grantedItem := range granted {
		grantedRule = append(grantedRule, grantedItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "SetPermission", _contractRule, grantedRule)
	if err != nil {
		return nil, err
	}
	return &RegistrySetPermissionIterator{contract: _Registry.contract, event: "SetPermission", logs: logs, sub: sub}, nil
}

// WatchSetPermission is a free log subscription operation binding the contract event 0xe9f5231bbfb4b32867755b94562215cff6c8998489de8ba20926f8d0980e7818.
//
// Solidity: e SetPermission(_contract indexed bytes32, granted indexed address, status bool)
func (_Registry *RegistryFilterer) WatchSetPermission(opts *bind.WatchOpts, sink chan<- *RegistrySetPermission, _contract [][32]byte, granted []common.Address) (event.Subscription, error) {

	var _contractRule []interface{}
	for _, _contractItem := range _contract {
		_contractRule = append(_contractRule, _contractItem)
	}
	var grantedRule []interface{}
	for _, grantedItem := range granted {
		grantedRule = append(grantedRule, grantedItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "SetPermission", _contractRule, grantedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrySetPermission)
				if err := _Registry.contract.UnpackLog(event, "SetPermission", log); err != nil {
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
