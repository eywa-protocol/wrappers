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

// RelayHubValidatorABI is the input ABI used to generate the binding from.
const RelayHubValidatorABI = "[]"

// RelayHubValidatorBin is the compiled bytecode used for deploying new contracts.
var RelayHubValidatorBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122020417900793e34b5ffb7664293bd2f5390050370e58d981e80aa647175859bff64736f6c634300080a0033"

// DeployRelayHubValidator deploys a new Ethereum contract, binding an instance of RelayHubValidator to it.
func DeployRelayHubValidator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RelayHubValidator, error) {
	parsed, err := abi.JSON(strings.NewReader(RelayHubValidatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RelayHubValidatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RelayHubValidator{RelayHubValidatorCaller: RelayHubValidatorCaller{contract: contract}, RelayHubValidatorTransactor: RelayHubValidatorTransactor{contract: contract}, RelayHubValidatorFilterer: RelayHubValidatorFilterer{contract: contract}}, nil
}

// RelayHubValidator is an auto generated Go binding around an Ethereum contract.
type RelayHubValidator struct {
	RelayHubValidatorCaller     // Read-only binding to the contract
	RelayHubValidatorTransactor // Write-only binding to the contract
	RelayHubValidatorFilterer   // Log filterer for contract events
}

// RelayHubValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type RelayHubValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayHubValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RelayHubValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayHubValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RelayHubValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayHubValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RelayHubValidatorSession struct {
	Contract     *RelayHubValidator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RelayHubValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RelayHubValidatorCallerSession struct {
	Contract *RelayHubValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// RelayHubValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RelayHubValidatorTransactorSession struct {
	Contract     *RelayHubValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// RelayHubValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type RelayHubValidatorRaw struct {
	Contract *RelayHubValidator // Generic contract binding to access the raw methods on
}

// RelayHubValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RelayHubValidatorCallerRaw struct {
	Contract *RelayHubValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// RelayHubValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RelayHubValidatorTransactorRaw struct {
	Contract *RelayHubValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRelayHubValidator creates a new instance of RelayHubValidator, bound to a specific deployed contract.
func NewRelayHubValidator(address common.Address, backend bind.ContractBackend) (*RelayHubValidator, error) {
	contract, err := bindRelayHubValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RelayHubValidator{RelayHubValidatorCaller: RelayHubValidatorCaller{contract: contract}, RelayHubValidatorTransactor: RelayHubValidatorTransactor{contract: contract}, RelayHubValidatorFilterer: RelayHubValidatorFilterer{contract: contract}}, nil
}

// NewRelayHubValidatorCaller creates a new read-only instance of RelayHubValidator, bound to a specific deployed contract.
func NewRelayHubValidatorCaller(address common.Address, caller bind.ContractCaller) (*RelayHubValidatorCaller, error) {
	contract, err := bindRelayHubValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RelayHubValidatorCaller{contract: contract}, nil
}

// NewRelayHubValidatorTransactor creates a new write-only instance of RelayHubValidator, bound to a specific deployed contract.
func NewRelayHubValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*RelayHubValidatorTransactor, error) {
	contract, err := bindRelayHubValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RelayHubValidatorTransactor{contract: contract}, nil
}

// NewRelayHubValidatorFilterer creates a new log filterer instance of RelayHubValidator, bound to a specific deployed contract.
func NewRelayHubValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*RelayHubValidatorFilterer, error) {
	contract, err := bindRelayHubValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RelayHubValidatorFilterer{contract: contract}, nil
}

// bindRelayHubValidator binds a generic wrapper to an already deployed contract.
func bindRelayHubValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RelayHubValidatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RelayHubValidator *RelayHubValidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RelayHubValidator.Contract.RelayHubValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RelayHubValidator *RelayHubValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayHubValidator.Contract.RelayHubValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RelayHubValidator *RelayHubValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RelayHubValidator.Contract.RelayHubValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RelayHubValidator *RelayHubValidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RelayHubValidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RelayHubValidator *RelayHubValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayHubValidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RelayHubValidator *RelayHubValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RelayHubValidator.Contract.contract.Transact(opts, method, params...)
}
