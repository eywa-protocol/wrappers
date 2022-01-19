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

// MinLibBytesABI is the input ABI used to generate the binding from.
const MinLibBytesABI = "[]"

// MinLibBytesBin is the compiled bytecode used for deploying new contracts.
var MinLibBytesBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212205b9a3621936e2bbc27afb3586645a6edbc46cf910868dd8fbd1d367d623c1b9164736f6c634300080a0033"

// DeployMinLibBytes deploys a new Ethereum contract, binding an instance of MinLibBytes to it.
func DeployMinLibBytes(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MinLibBytes, error) {
	parsed, err := abi.JSON(strings.NewReader(MinLibBytesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MinLibBytesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MinLibBytes{MinLibBytesCaller: MinLibBytesCaller{contract: contract}, MinLibBytesTransactor: MinLibBytesTransactor{contract: contract}, MinLibBytesFilterer: MinLibBytesFilterer{contract: contract}}, nil
}

// MinLibBytes is an auto generated Go binding around an Ethereum contract.
type MinLibBytes struct {
	MinLibBytesCaller     // Read-only binding to the contract
	MinLibBytesTransactor // Write-only binding to the contract
	MinLibBytesFilterer   // Log filterer for contract events
}

// MinLibBytesCaller is an auto generated read-only Go binding around an Ethereum contract.
type MinLibBytesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinLibBytesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MinLibBytesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinLibBytesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MinLibBytesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinLibBytesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MinLibBytesSession struct {
	Contract     *MinLibBytes      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MinLibBytesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MinLibBytesCallerSession struct {
	Contract *MinLibBytesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MinLibBytesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MinLibBytesTransactorSession struct {
	Contract     *MinLibBytesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MinLibBytesRaw is an auto generated low-level Go binding around an Ethereum contract.
type MinLibBytesRaw struct {
	Contract *MinLibBytes // Generic contract binding to access the raw methods on
}

// MinLibBytesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MinLibBytesCallerRaw struct {
	Contract *MinLibBytesCaller // Generic read-only contract binding to access the raw methods on
}

// MinLibBytesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MinLibBytesTransactorRaw struct {
	Contract *MinLibBytesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMinLibBytes creates a new instance of MinLibBytes, bound to a specific deployed contract.
func NewMinLibBytes(address common.Address, backend bind.ContractBackend) (*MinLibBytes, error) {
	contract, err := bindMinLibBytes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MinLibBytes{MinLibBytesCaller: MinLibBytesCaller{contract: contract}, MinLibBytesTransactor: MinLibBytesTransactor{contract: contract}, MinLibBytesFilterer: MinLibBytesFilterer{contract: contract}}, nil
}

// NewMinLibBytesCaller creates a new read-only instance of MinLibBytes, bound to a specific deployed contract.
func NewMinLibBytesCaller(address common.Address, caller bind.ContractCaller) (*MinLibBytesCaller, error) {
	contract, err := bindMinLibBytes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MinLibBytesCaller{contract: contract}, nil
}

// NewMinLibBytesTransactor creates a new write-only instance of MinLibBytes, bound to a specific deployed contract.
func NewMinLibBytesTransactor(address common.Address, transactor bind.ContractTransactor) (*MinLibBytesTransactor, error) {
	contract, err := bindMinLibBytes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MinLibBytesTransactor{contract: contract}, nil
}

// NewMinLibBytesFilterer creates a new log filterer instance of MinLibBytes, bound to a specific deployed contract.
func NewMinLibBytesFilterer(address common.Address, filterer bind.ContractFilterer) (*MinLibBytesFilterer, error) {
	contract, err := bindMinLibBytes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MinLibBytesFilterer{contract: contract}, nil
}

// bindMinLibBytes binds a generic wrapper to an already deployed contract.
func bindMinLibBytes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MinLibBytesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinLibBytes *MinLibBytesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MinLibBytes.Contract.MinLibBytesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinLibBytes *MinLibBytesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinLibBytes.Contract.MinLibBytesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinLibBytes *MinLibBytesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinLibBytes.Contract.MinLibBytesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinLibBytes *MinLibBytesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MinLibBytes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinLibBytes *MinLibBytesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinLibBytes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinLibBytes *MinLibBytesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinLibBytes.Contract.contract.Transact(opts, method, params...)
}
