// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wrappers

import (
	"errors"
	"fmt"
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

// ICurveProxyMetaData contains all meta data concerning the ICurveProxy contract.
var ICurveProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addAtCrosschainPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinMintAmountC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lpIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addAtHubPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinMintAmountH\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialChainID\",\"type\":\"uint256\"}],\"internalType\":\"structICurveProxy.MetaMintEUSD\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structICurveProxy.PermitData[]\",\"name\":\"permit\",\"type\":\"tuple[]\"},{\"internalType\":\"address[3]\",\"name\":\"token\",\"type\":\"address[3]\"},{\"internalType\":\"uint256[3]\",\"name\":\"amount\",\"type\":\"uint256[3]\"}],\"name\":\"addLiquidity3PoolMintEUSD\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remove\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinMintAmount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinDy\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"chain2address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"initialBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialChainID\",\"type\":\"uint256\"}],\"internalType\":\"structICurveProxy.MetaExchangeParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structICurveProxy.PermitData[]\",\"name\":\"permit\",\"type\":\"tuple[]\"},{\"internalType\":\"address[3]\",\"name\":\"token\",\"type\":\"address[3]\"},{\"internalType\":\"uint256[3]\",\"name\":\"amount\",\"type\":\"uint256[3]\"}],\"name\":\"metaExchange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"removeAtCrosschainPool\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinAmountC\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"removeAtHubPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountH\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"y\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinAmountH\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"internalType\":\"structICurveProxy.MetaRedeemEUSD\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structICurveProxy.PermitData\",\"name\":\"permit\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"redeemEUSD\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ICurveProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use ICurveProxyMetaData.ABI instead.
var ICurveProxyABI = ICurveProxyMetaData.ABI

// ICurveProxy is an auto generated Go binding around an Ethereum contract.
type ICurveProxy struct {
	ICurveProxyCaller     // Read-only binding to the contract
	ICurveProxyTransactor // Write-only binding to the contract
	ICurveProxyFilterer   // Log filterer for contract events
}

// ICurveProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICurveProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICurveProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICurveProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
	gsn      *GsnCallOpts
}

func (_ICurveProxy *ICurveProxyTransactor) SetGSNOptions(opts *GsnCallOpts) {
	_ICurveProxy.gsn = opts
}

// ICurveProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICurveProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICurveProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICurveProxySession struct {
	Contract     *ICurveProxy      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICurveProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICurveProxyCallerSession struct {
	Contract *ICurveProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ICurveProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICurveProxyTransactorSession struct {
	Contract     *ICurveProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ICurveProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICurveProxyRaw struct {
	Contract *ICurveProxy // Generic contract binding to access the raw methods on
}

// ICurveProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICurveProxyCallerRaw struct {
	Contract *ICurveProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ICurveProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICurveProxyTransactorRaw struct {
	Contract *ICurveProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICurveProxy creates a new instance of ICurveProxy, bound to a specific deployed contract.
func NewICurveProxy(address common.Address, backend bind.ContractBackend) (*ICurveProxy, error) {
	contract, err := bindICurveProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICurveProxy{ICurveProxyCaller: ICurveProxyCaller{contract: contract}, ICurveProxyTransactor: ICurveProxyTransactor{contract: contract}, ICurveProxyFilterer: ICurveProxyFilterer{contract: contract}}, nil
}

// NewICurveProxyCaller creates a new read-only instance of ICurveProxy, bound to a specific deployed contract.
func NewICurveProxyCaller(address common.Address, caller bind.ContractCaller) (*ICurveProxyCaller, error) {
	contract, err := bindICurveProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICurveProxyCaller{contract: contract}, nil
}

// NewICurveProxyTransactor creates a new write-only instance of ICurveProxy, bound to a specific deployed contract.
func NewICurveProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ICurveProxyTransactor, error) {
	contract, err := bindICurveProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICurveProxyTransactor{contract: contract}, nil
}

// NewICurveProxyFilterer creates a new log filterer instance of ICurveProxy, bound to a specific deployed contract.
func NewICurveProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ICurveProxyFilterer, error) {
	contract, err := bindICurveProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICurveProxyFilterer{contract: contract}, nil
}

// bindICurveProxy binds a generic wrapper to an already deployed contract.
func bindICurveProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ICurveProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICurveProxy *ICurveProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICurveProxy.Contract.ICurveProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICurveProxy *ICurveProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICurveProxy.Contract.ICurveProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICurveProxy *ICurveProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICurveProxy.Contract.ICurveProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICurveProxy *ICurveProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICurveProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICurveProxy *ICurveProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICurveProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICurveProxy *ICurveProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICurveProxy.Contract.contract.Transact(opts, method, params...)
}

// AddLiquidity3PoolMintEUSD is a paid mutator transaction binding the contract method 0x38578fa0.
//
// Solidity: function addLiquidity3PoolMintEUSD((address,uint256,uint256,address,uint256,address,address,uint256) params, (uint8,bytes32,bytes32,uint256,bool)[] permit, address[3] token, uint256[3] amount) returns()
func (_ICurveProxy *ICurveProxyTransactor) AddLiquidity3PoolMintEUSD(opts *bind.TransactOpts, params ICurveProxyMetaMintEUSD, permit []ICurveProxyPermitData, token [3]common.Address, amount [3]*big.Int) (*types.Transaction, error) {
	return _ICurveProxy.contract.Transact(opts, "addLiquidity3PoolMintEUSD", params, permit, token, amount)
}
func (_ICurveProxy *ICurveProxyTransactor) AddLiquidity3PoolMintEUSDOverGsn(opts *bind.TransactOpts, params ICurveProxyMetaMintEUSD, permit []ICurveProxyPermitData, token [3]common.Address, amount [3]*big.Int) (common.Hash, error) {
	return GsnExecutor(_ICurveProxy.gsn, ICurveProxyMetaData.ABI, "addLiquidity3PoolMintEUSD", params, permit, token, amount)
}

// AddLiquidity3PoolMintEUSD is a paid mutator transaction binding the contract method 0x38578fa0.
//
// Solidity: function addLiquidity3PoolMintEUSD((address,uint256,uint256,address,uint256,address,address,uint256) params, (uint8,bytes32,bytes32,uint256,bool)[] permit, address[3] token, uint256[3] amount) returns()
func (_ICurveProxy *ICurveProxySession) AddLiquidity3PoolMintEUSD(params ICurveProxyMetaMintEUSD, permit []ICurveProxyPermitData, token [3]common.Address, amount [3]*big.Int) (*types.Transaction, error) {
	return _ICurveProxy.Contract.AddLiquidity3PoolMintEUSD(&_ICurveProxy.TransactOpts, params, permit, token, amount)
}
func (_ICurveProxy *ICurveProxySession) AddLiquidity3PoolMintEUSDOverGsn(params ICurveProxyMetaMintEUSD, permit []ICurveProxyPermitData, token [3]common.Address, amount [3]*big.Int) (common.Hash, error) {
	return _ICurveProxy.Contract.AddLiquidity3PoolMintEUSDOverGsn(&_ICurveProxy.TransactOpts, params, permit, token, amount)
}

// AddLiquidity3PoolMintEUSD is a paid mutator transaction binding the contract method 0x38578fa0.
//
// Solidity: function addLiquidity3PoolMintEUSD((address,uint256,uint256,address,uint256,address,address,uint256) params, (uint8,bytes32,bytes32,uint256,bool)[] permit, address[3] token, uint256[3] amount) returns()
func (_ICurveProxy *ICurveProxyTransactorSession) AddLiquidity3PoolMintEUSD(params ICurveProxyMetaMintEUSD, permit []ICurveProxyPermitData, token [3]common.Address, amount [3]*big.Int) (*types.Transaction, error) {
	return _ICurveProxy.Contract.AddLiquidity3PoolMintEUSD(&_ICurveProxy.TransactOpts, params, permit, token, amount)
}
func (_ICurveProxy *ICurveProxyTransactorSession) AddLiquidity3PoolMintEUSDOverGsn(params ICurveProxyMetaMintEUSD, permit []ICurveProxyPermitData, token [3]common.Address, amount [3]*big.Int) (common.Hash, error) {
	return _ICurveProxy.Contract.AddLiquidity3PoolMintEUSDOverGsn(&_ICurveProxy.TransactOpts, params, permit, token, amount)
}

// MetaExchange is a paid mutator transaction binding the contract method 0x8c71de7e.
//
// Solidity: function metaExchange((address,address,address,uint256,int128,int128,uint256,int128,uint256,address,address,address,address,uint256,address,uint256) params, (uint8,bytes32,bytes32,uint256,bool)[] permit, address[3] token, uint256[3] amount) returns()
func (_ICurveProxy *ICurveProxyTransactor) MetaExchange(opts *bind.TransactOpts, params ICurveProxyMetaExchangeParams, permit []ICurveProxyPermitData, token [3]common.Address, amount [3]*big.Int) (*types.Transaction, error) {
	return _ICurveProxy.contract.Transact(opts, "metaExchange", params, permit, token, amount)
}
func (_ICurveProxy *ICurveProxyTransactor) MetaExchangeOverGsn(opts *bind.TransactOpts, params ICurveProxyMetaExchangeParams, permit []ICurveProxyPermitData, token [3]common.Address, amount [3]*big.Int) (common.Hash, error) {
	return GsnExecutor(_ICurveProxy.gsn, ICurveProxyMetaData.ABI, "metaExchange", params, permit, token, amount)
}

// MetaExchange is a paid mutator transaction binding the contract method 0x8c71de7e.
//
// Solidity: function metaExchange((address,address,address,uint256,int128,int128,uint256,int128,uint256,address,address,address,address,uint256,address,uint256) params, (uint8,bytes32,bytes32,uint256,bool)[] permit, address[3] token, uint256[3] amount) returns()
func (_ICurveProxy *ICurveProxySession) MetaExchange(params ICurveProxyMetaExchangeParams, permit []ICurveProxyPermitData, token [3]common.Address, amount [3]*big.Int) (*types.Transaction, error) {
	return _ICurveProxy.Contract.MetaExchange(&_ICurveProxy.TransactOpts, params, permit, token, amount)
}
func (_ICurveProxy *ICurveProxySession) MetaExchangeOverGsn(params ICurveProxyMetaExchangeParams, permit []ICurveProxyPermitData, token [3]common.Address, amount [3]*big.Int) (common.Hash, error) {
	return _ICurveProxy.Contract.MetaExchangeOverGsn(&_ICurveProxy.TransactOpts, params, permit, token, amount)
}

// MetaExchange is a paid mutator transaction binding the contract method 0x8c71de7e.
//
// Solidity: function metaExchange((address,address,address,uint256,int128,int128,uint256,int128,uint256,address,address,address,address,uint256,address,uint256) params, (uint8,bytes32,bytes32,uint256,bool)[] permit, address[3] token, uint256[3] amount) returns()
func (_ICurveProxy *ICurveProxyTransactorSession) MetaExchange(params ICurveProxyMetaExchangeParams, permit []ICurveProxyPermitData, token [3]common.Address, amount [3]*big.Int) (*types.Transaction, error) {
	return _ICurveProxy.Contract.MetaExchange(&_ICurveProxy.TransactOpts, params, permit, token, amount)
}
func (_ICurveProxy *ICurveProxyTransactorSession) MetaExchangeOverGsn(params ICurveProxyMetaExchangeParams, permit []ICurveProxyPermitData, token [3]common.Address, amount [3]*big.Int) (common.Hash, error) {
	return _ICurveProxy.Contract.MetaExchangeOverGsn(&_ICurveProxy.TransactOpts, params, permit, token, amount)
}

// RedeemEUSD is a paid mutator transaction binding the contract method 0x5fcc4a3f.
//
// Solidity: function redeemEUSD((address,int128,uint256,address,uint256,int128,uint256,address) params, (uint8,bytes32,bytes32,uint256,bool) permit, address receiveSide, address oppositeBridge, uint256 chainId) returns()
func (_ICurveProxy *ICurveProxyTransactor) RedeemEUSD(opts *bind.TransactOpts, params ICurveProxyMetaRedeemEUSD, permit ICurveProxyPermitData, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _ICurveProxy.contract.Transact(opts, "redeemEUSD", params, permit, receiveSide, oppositeBridge, chainId)
}
func (_ICurveProxy *ICurveProxyTransactor) RedeemEUSDOverGsn(opts *bind.TransactOpts, params ICurveProxyMetaRedeemEUSD, permit ICurveProxyPermitData, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (common.Hash, error) {
	return GsnExecutor(_ICurveProxy.gsn, ICurveProxyMetaData.ABI, "redeemEUSD", params, permit, receiveSide, oppositeBridge, chainId)
}

// RedeemEUSD is a paid mutator transaction binding the contract method 0x5fcc4a3f.
//
// Solidity: function redeemEUSD((address,int128,uint256,address,uint256,int128,uint256,address) params, (uint8,bytes32,bytes32,uint256,bool) permit, address receiveSide, address oppositeBridge, uint256 chainId) returns()
func (_ICurveProxy *ICurveProxySession) RedeemEUSD(params ICurveProxyMetaRedeemEUSD, permit ICurveProxyPermitData, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _ICurveProxy.Contract.RedeemEUSD(&_ICurveProxy.TransactOpts, params, permit, receiveSide, oppositeBridge, chainId)
}
func (_ICurveProxy *ICurveProxySession) RedeemEUSDOverGsn(params ICurveProxyMetaRedeemEUSD, permit ICurveProxyPermitData, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (common.Hash, error) {
	return _ICurveProxy.Contract.RedeemEUSDOverGsn(&_ICurveProxy.TransactOpts, params, permit, receiveSide, oppositeBridge, chainId)
}

// RedeemEUSD is a paid mutator transaction binding the contract method 0x5fcc4a3f.
//
// Solidity: function redeemEUSD((address,int128,uint256,address,uint256,int128,uint256,address) params, (uint8,bytes32,bytes32,uint256,bool) permit, address receiveSide, address oppositeBridge, uint256 chainId) returns()
func (_ICurveProxy *ICurveProxyTransactorSession) RedeemEUSD(params ICurveProxyMetaRedeemEUSD, permit ICurveProxyPermitData, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _ICurveProxy.Contract.RedeemEUSD(&_ICurveProxy.TransactOpts, params, permit, receiveSide, oppositeBridge, chainId)
}
func (_ICurveProxy *ICurveProxyTransactorSession) RedeemEUSDOverGsn(params ICurveProxyMetaRedeemEUSD, permit ICurveProxyPermitData, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (common.Hash, error) {
	return _ICurveProxy.Contract.RedeemEUSDOverGsn(&_ICurveProxy.TransactOpts, params, permit, receiveSide, oppositeBridge, chainId)
}
