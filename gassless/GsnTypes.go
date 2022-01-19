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

// GsnTypesABI is the input ABI used to generate the binding from.
const GsnTypesABI = "[]"

// GsnTypes is an auto generated Go binding around an Ethereum contract.
type GsnTypes struct {
	GsnTypesCaller     // Read-only binding to the contract
	GsnTypesTransactor // Write-only binding to the contract
	GsnTypesFilterer   // Log filterer for contract events
}

// GsnTypesCaller is an auto generated read-only Go binding around an Ethereum contract.
type GsnTypesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GsnTypesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GsnTypesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GsnTypesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GsnTypesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GsnTypesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GsnTypesSession struct {
	Contract     *GsnTypes         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GsnTypesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GsnTypesCallerSession struct {
	Contract *GsnTypesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// GsnTypesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GsnTypesTransactorSession struct {
	Contract     *GsnTypesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GsnTypesRaw is an auto generated low-level Go binding around an Ethereum contract.
type GsnTypesRaw struct {
	Contract *GsnTypes // Generic contract binding to access the raw methods on
}

// GsnTypesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GsnTypesCallerRaw struct {
	Contract *GsnTypesCaller // Generic read-only contract binding to access the raw methods on
}

// GsnTypesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GsnTypesTransactorRaw struct {
	Contract *GsnTypesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGsnTypes creates a new instance of GsnTypes, bound to a specific deployed contract.
func NewGsnTypes(address common.Address, backend bind.ContractBackend) (*GsnTypes, error) {
	contract, err := bindGsnTypes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GsnTypes{GsnTypesCaller: GsnTypesCaller{contract: contract}, GsnTypesTransactor: GsnTypesTransactor{contract: contract}, GsnTypesFilterer: GsnTypesFilterer{contract: contract}}, nil
}

// NewGsnTypesCaller creates a new read-only instance of GsnTypes, bound to a specific deployed contract.
func NewGsnTypesCaller(address common.Address, caller bind.ContractCaller) (*GsnTypesCaller, error) {
	contract, err := bindGsnTypes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GsnTypesCaller{contract: contract}, nil
}

// NewGsnTypesTransactor creates a new write-only instance of GsnTypes, bound to a specific deployed contract.
func NewGsnTypesTransactor(address common.Address, transactor bind.ContractTransactor) (*GsnTypesTransactor, error) {
	contract, err := bindGsnTypes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GsnTypesTransactor{contract: contract}, nil
}

// NewGsnTypesFilterer creates a new log filterer instance of GsnTypes, bound to a specific deployed contract.
func NewGsnTypesFilterer(address common.Address, filterer bind.ContractFilterer) (*GsnTypesFilterer, error) {
	contract, err := bindGsnTypes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GsnTypesFilterer{contract: contract}, nil
}

// bindGsnTypes binds a generic wrapper to an already deployed contract.
func bindGsnTypes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GsnTypesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GsnTypes *GsnTypesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GsnTypes.Contract.GsnTypesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GsnTypes *GsnTypesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GsnTypes.Contract.GsnTypesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GsnTypes *GsnTypesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GsnTypes.Contract.GsnTypesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GsnTypes *GsnTypesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GsnTypes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GsnTypes *GsnTypesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GsnTypes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GsnTypes *GsnTypesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GsnTypes.Contract.contract.Transact(opts, method, params...)
}
