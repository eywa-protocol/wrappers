// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wrappers

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = fmt.Errorf("")
)

// IERC20PermitMetaData contains all meta data concerning the IERC20Permit contract.

var IERC20PermitMetaData = &bind.MetaData{

	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC20PermitABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20PermitMetaData.ABI instead.
var IERC20PermitABI = IERC20PermitMetaData.ABI

// IERC20Permit is an auto generated Go binding around an Ethereum contract.
type IERC20Permit struct {
	IERC20PermitCaller     // Read-only binding to the contract
	IERC20PermitTransactor // Write-only binding to the contract
	IERC20PermitFilterer   // Log filterer for contract events
}

// IERC20PermitCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20PermitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20PermitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20PermitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
	gsn      *GsnCallOpts
}

func (_IERC20Permit *IERC20PermitTransactor) SetGSNOptions(opts *GsnCallOpts) {
	_IERC20Permit.gsn = opts
}

// IERC20PermitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20PermitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20PermitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20PermitSession struct {
	Contract     *IERC20Permit     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20PermitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20PermitCallerSession struct {
	Contract *IERC20PermitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IERC20PermitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20PermitTransactorSession struct {
	Contract     *IERC20PermitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IERC20PermitRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20PermitRaw struct {
	Contract *IERC20Permit // Generic contract binding to access the raw methods on
}

// IERC20PermitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20PermitCallerRaw struct {
	Contract *IERC20PermitCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20PermitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20PermitTransactorRaw struct {
	Contract *IERC20PermitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20Permit creates a new instance of IERC20Permit, bound to a specific deployed contract.
func NewIERC20Permit(address common.Address, backend bind.ContractBackend) (*IERC20Permit, error) {
	contract, err := bindIERC20Permit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20Permit{IERC20PermitCaller: IERC20PermitCaller{contract: contract}, IERC20PermitTransactor: IERC20PermitTransactor{contract: contract}, IERC20PermitFilterer: IERC20PermitFilterer{contract: contract}}, nil
}

// NewIERC20PermitCaller creates a new read-only instance of IERC20Permit, bound to a specific deployed contract.
func NewIERC20PermitCaller(address common.Address, caller bind.ContractCaller) (*IERC20PermitCaller, error) {
	contract, err := bindIERC20Permit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20PermitCaller{contract: contract}, nil
}

// NewIERC20PermitTransactor creates a new write-only instance of IERC20Permit, bound to a specific deployed contract.
func NewIERC20PermitTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20PermitTransactor, error) {
	contract, err := bindIERC20Permit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20PermitTransactor{contract: contract}, nil
}

// NewIERC20PermitFilterer creates a new log filterer instance of IERC20Permit, bound to a specific deployed contract.
func NewIERC20PermitFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20PermitFilterer, error) {
	contract, err := bindIERC20Permit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20PermitFilterer{contract: contract}, nil
}

// bindIERC20Permit binds a generic wrapper to an already deployed contract.
func bindIERC20Permit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20PermitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Permit *IERC20PermitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Permit.Contract.IERC20PermitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Permit *IERC20PermitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Permit.Contract.IERC20PermitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Permit *IERC20PermitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Permit.Contract.IERC20PermitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Permit *IERC20PermitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Permit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Permit *IERC20PermitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Permit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Permit *IERC20PermitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Permit.Contract.contract.Transact(opts, method, params...)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC20Permit *IERC20PermitTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC20Permit.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}
func (_IERC20Permit *IERC20PermitTransactor) PermitOverGsn(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (common.Hash, error) {
	return GsnExecutor(_IERC20Permit.gsn, IERC20PermitMetaData.ABI, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC20Permit *IERC20PermitSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC20Permit.Contract.Permit(&_IERC20Permit.TransactOpts, owner, spender, value, deadline, v, r, s)
}
func (_IERC20Permit *IERC20PermitSession) PermitOverGsn(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (common.Hash, error) {
	return _IERC20Permit.Contract.PermitOverGsn(&_IERC20Permit.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC20Permit *IERC20PermitTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC20Permit.Contract.Permit(&_IERC20Permit.TransactOpts, owner, spender, value, deadline, v, r, s)
}
func (_IERC20Permit *IERC20PermitTransactorSession) PermitOverGsn(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (common.Hash, error) {
	return _IERC20Permit.Contract.PermitOverGsn(&_IERC20Permit.TransactOpts, owner, spender, value, deadline, v, r, s)
}
