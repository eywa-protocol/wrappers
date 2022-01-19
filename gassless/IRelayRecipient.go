// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gassless

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

// IRelayRecipientABI is the input ABI used to generate the binding from.
const IRelayRecipientABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionRecipient\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IRelayRecipient is an auto generated Go binding around an Ethereum contract.
type IRelayRecipient struct {
	IRelayRecipientCaller     // Read-only binding to the contract
	IRelayRecipientTransactor // Write-only binding to the contract
	IRelayRecipientFilterer   // Log filterer for contract events
}

// IRelayRecipientCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRelayRecipientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRelayRecipientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRelayRecipientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRelayRecipientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRelayRecipientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRelayRecipientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRelayRecipientSession struct {
	Contract     *IRelayRecipient  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRelayRecipientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRelayRecipientCallerSession struct {
	Contract *IRelayRecipientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IRelayRecipientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRelayRecipientTransactorSession struct {
	Contract     *IRelayRecipientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IRelayRecipientRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRelayRecipientRaw struct {
	Contract *IRelayRecipient // Generic contract binding to access the raw methods on
}

// IRelayRecipientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRelayRecipientCallerRaw struct {
	Contract *IRelayRecipientCaller // Generic read-only contract binding to access the raw methods on
}

// IRelayRecipientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRelayRecipientTransactorRaw struct {
	Contract *IRelayRecipientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRelayRecipient creates a new instance of IRelayRecipient, bound to a specific deployed contract.
func NewIRelayRecipient(address common.Address, backend bind.ContractBackend) (*IRelayRecipient, error) {
	contract, err := bindIRelayRecipient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRelayRecipient{IRelayRecipientCaller: IRelayRecipientCaller{contract: contract}, IRelayRecipientTransactor: IRelayRecipientTransactor{contract: contract}, IRelayRecipientFilterer: IRelayRecipientFilterer{contract: contract}}, nil
}

// NewIRelayRecipientCaller creates a new read-only instance of IRelayRecipient, bound to a specific deployed contract.
func NewIRelayRecipientCaller(address common.Address, caller bind.ContractCaller) (*IRelayRecipientCaller, error) {
	contract, err := bindIRelayRecipient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRelayRecipientCaller{contract: contract}, nil
}

// NewIRelayRecipientTransactor creates a new write-only instance of IRelayRecipient, bound to a specific deployed contract.
func NewIRelayRecipientTransactor(address common.Address, transactor bind.ContractTransactor) (*IRelayRecipientTransactor, error) {
	contract, err := bindIRelayRecipient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRelayRecipientTransactor{contract: contract}, nil
}

// NewIRelayRecipientFilterer creates a new log filterer instance of IRelayRecipient, bound to a specific deployed contract.
func NewIRelayRecipientFilterer(address common.Address, filterer bind.ContractFilterer) (*IRelayRecipientFilterer, error) {
	contract, err := bindIRelayRecipient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRelayRecipientFilterer{contract: contract}, nil
}

// bindIRelayRecipient binds a generic wrapper to an already deployed contract.
func bindIRelayRecipient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IRelayRecipientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRelayRecipient *IRelayRecipientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRelayRecipient.Contract.IRelayRecipientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRelayRecipient *IRelayRecipientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRelayRecipient.Contract.IRelayRecipientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRelayRecipient *IRelayRecipientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRelayRecipient.Contract.IRelayRecipientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRelayRecipient *IRelayRecipientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRelayRecipient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRelayRecipient *IRelayRecipientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRelayRecipient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRelayRecipient *IRelayRecipientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRelayRecipient.Contract.contract.Transact(opts, method, params...)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_IRelayRecipient *IRelayRecipientCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _IRelayRecipient.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_IRelayRecipient *IRelayRecipientSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _IRelayRecipient.Contract.IsTrustedForwarder(&_IRelayRecipient.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_IRelayRecipient *IRelayRecipientCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _IRelayRecipient.Contract.IsTrustedForwarder(&_IRelayRecipient.CallOpts, forwarder)
}

// VersionRecipient is a free data retrieval call binding the contract method 0x486ff0cd.
//
// Solidity: function versionRecipient() view returns(string)
func (_IRelayRecipient *IRelayRecipientCaller) VersionRecipient(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IRelayRecipient.contract.Call(opts, &out, "versionRecipient")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VersionRecipient is a free data retrieval call binding the contract method 0x486ff0cd.
//
// Solidity: function versionRecipient() view returns(string)
func (_IRelayRecipient *IRelayRecipientSession) VersionRecipient() (string, error) {
	return _IRelayRecipient.Contract.VersionRecipient(&_IRelayRecipient.CallOpts)
}

// VersionRecipient is a free data retrieval call binding the contract method 0x486ff0cd.
//
// Solidity: function versionRecipient() view returns(string)
func (_IRelayRecipient *IRelayRecipientCallerSession) VersionRecipient() (string, error) {
	return _IRelayRecipient.Contract.VersionRecipient(&_IRelayRecipient.CallOpts)
}
