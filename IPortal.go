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

// IPortalMetaData contains all meta data concerning the IPortal contract.
var IPortalMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"emergencyUnburnRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"pubkeys\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"emergencyUnburnRequestToSolana\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structIPortal.SynthParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"synthesize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structIPortal.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIPortal.TransitData\",\"name\":\"transitData\",\"type\":\"tuple\"}],\"name\":\"synthesizeBatchWithDataTransit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"pubkeys\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes1\",\"name\":\"txStateBump\",\"type\":\"bytes1\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"synthesizeToSolana\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structIPortal.PermitData\",\"name\":\"permitData\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"synthesizeWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IPortalABI is the input ABI used to generate the binding from.
// Deprecated: Use IPortalMetaData.ABI instead.
var IPortalABI = IPortalMetaData.ABI

// IPortal is an auto generated Go binding around an Ethereum contract.
type IPortal struct {
	IPortalCaller     // Read-only binding to the contract
	IPortalTransactor // Write-only binding to the contract
	IPortalFilterer   // Log filterer for contract events
}

// IPortalCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPortalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPortalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPortalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
	gsn      *GsnCallOpts
}

func (_IPortal *IPortalTransactor) SetGSNOptions(opts *GsnCallOpts) {
	_IPortal.gsn = opts
}

// IPortalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPortalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPortalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPortalSession struct {
	Contract     *IPortal          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPortalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPortalCallerSession struct {
	Contract *IPortalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IPortalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPortalTransactorSession struct {
	Contract     *IPortalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IPortalRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPortalRaw struct {
	Contract *IPortal // Generic contract binding to access the raw methods on
}

// IPortalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPortalCallerRaw struct {
	Contract *IPortalCaller // Generic read-only contract binding to access the raw methods on
}

// IPortalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPortalTransactorRaw struct {
	Contract *IPortalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPortal creates a new instance of IPortal, bound to a specific deployed contract.
func NewIPortal(address common.Address, backend bind.ContractBackend) (*IPortal, error) {
	contract, err := bindIPortal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPortal{IPortalCaller: IPortalCaller{contract: contract}, IPortalTransactor: IPortalTransactor{contract: contract}, IPortalFilterer: IPortalFilterer{contract: contract}}, nil
}

// NewIPortalCaller creates a new read-only instance of IPortal, bound to a specific deployed contract.
func NewIPortalCaller(address common.Address, caller bind.ContractCaller) (*IPortalCaller, error) {
	contract, err := bindIPortal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPortalCaller{contract: contract}, nil
}

// NewIPortalTransactor creates a new write-only instance of IPortal, bound to a specific deployed contract.
func NewIPortalTransactor(address common.Address, transactor bind.ContractTransactor) (*IPortalTransactor, error) {
	contract, err := bindIPortal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPortalTransactor{contract: contract}, nil
}

// NewIPortalFilterer creates a new log filterer instance of IPortal, bound to a specific deployed contract.
func NewIPortalFilterer(address common.Address, filterer bind.ContractFilterer) (*IPortalFilterer, error) {
	contract, err := bindIPortal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPortalFilterer{contract: contract}, nil
}

// bindIPortal binds a generic wrapper to an already deployed contract.
func bindIPortal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPortalABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPortal *IPortalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPortal.Contract.IPortalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPortal *IPortalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPortal.Contract.IPortalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPortal *IPortalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPortal.Contract.IPortalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPortal *IPortalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPortal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPortal *IPortalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPortal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPortal *IPortalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPortal.Contract.contract.Transact(opts, method, params...)
}

// EmergencyUnburnRequest is a paid mutator transaction binding the contract method 0x43961668.
//
// Solidity: function emergencyUnburnRequest(bytes32 txID, address receiveSide, address oppositeBridge, uint256 chainId, uint8 v, bytes32 r, bytes32 s) returns()
func (_IPortal *IPortalTransactor) EmergencyUnburnRequest(opts *bind.TransactOpts, txID [32]byte, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPortal.contract.Transact(opts, "emergencyUnburnRequest", txID, receiveSide, oppositeBridge, chainId, v, r, s)
}
func (_IPortal *IPortalTransactor) EmergencyUnburnRequestOverGsn(opts *bind.TransactOpts, txID [32]byte, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, v uint8, r [32]byte, s [32]byte) (common.Hash, error) {
	return GsnExecutor(_IPortal.gsn, IPortalMetaData.ABI, "emergencyUnburnRequest", txID, receiveSide, oppositeBridge, chainId, v, r, s)
}

// EmergencyUnburnRequest is a paid mutator transaction binding the contract method 0x43961668.
//
// Solidity: function emergencyUnburnRequest(bytes32 txID, address receiveSide, address oppositeBridge, uint256 chainId, uint8 v, bytes32 r, bytes32 s) returns()
func (_IPortal *IPortalSession) EmergencyUnburnRequest(txID [32]byte, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPortal.Contract.EmergencyUnburnRequest(&_IPortal.TransactOpts, txID, receiveSide, oppositeBridge, chainId, v, r, s)
}
func (_IPortal *IPortalSession) EmergencyUnburnRequestOverGsn(txID [32]byte, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, v uint8, r [32]byte, s [32]byte) (common.Hash, error) {
	return _IPortal.Contract.EmergencyUnburnRequestOverGsn(&_IPortal.TransactOpts, txID, receiveSide, oppositeBridge, chainId, v, r, s)
}

// EmergencyUnburnRequest is a paid mutator transaction binding the contract method 0x43961668.
//
// Solidity: function emergencyUnburnRequest(bytes32 txID, address receiveSide, address oppositeBridge, uint256 chainId, uint8 v, bytes32 r, bytes32 s) returns()
func (_IPortal *IPortalTransactorSession) EmergencyUnburnRequest(txID [32]byte, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IPortal.Contract.EmergencyUnburnRequest(&_IPortal.TransactOpts, txID, receiveSide, oppositeBridge, chainId, v, r, s)
}
func (_IPortal *IPortalTransactorSession) EmergencyUnburnRequestOverGsn(txID [32]byte, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, v uint8, r [32]byte, s [32]byte) (common.Hash, error) {
	return _IPortal.Contract.EmergencyUnburnRequestOverGsn(&_IPortal.TransactOpts, txID, receiveSide, oppositeBridge, chainId, v, r, s)
}

// EmergencyUnburnRequestToSolana is a paid mutator transaction binding the contract method 0xa9c4475f.
//
// Solidity: function emergencyUnburnRequestToSolana(bytes32 txID, address from, bytes32[] pubkeys, uint256 chainId) returns()
func (_IPortal *IPortalTransactor) EmergencyUnburnRequestToSolana(opts *bind.TransactOpts, txID [32]byte, from common.Address, pubkeys [][32]byte, chainId *big.Int) (*types.Transaction, error) {
	return _IPortal.contract.Transact(opts, "emergencyUnburnRequestToSolana", txID, from, pubkeys, chainId)
}
func (_IPortal *IPortalTransactor) EmergencyUnburnRequestToSolanaOverGsn(opts *bind.TransactOpts, txID [32]byte, from common.Address, pubkeys [][32]byte, chainId *big.Int) (common.Hash, error) {
	return GsnExecutor(_IPortal.gsn, IPortalMetaData.ABI, "emergencyUnburnRequestToSolana", txID, from, pubkeys, chainId)
}

// EmergencyUnburnRequestToSolana is a paid mutator transaction binding the contract method 0xa9c4475f.
//
// Solidity: function emergencyUnburnRequestToSolana(bytes32 txID, address from, bytes32[] pubkeys, uint256 chainId) returns()
func (_IPortal *IPortalSession) EmergencyUnburnRequestToSolana(txID [32]byte, from common.Address, pubkeys [][32]byte, chainId *big.Int) (*types.Transaction, error) {
	return _IPortal.Contract.EmergencyUnburnRequestToSolana(&_IPortal.TransactOpts, txID, from, pubkeys, chainId)
}
func (_IPortal *IPortalSession) EmergencyUnburnRequestToSolanaOverGsn(txID [32]byte, from common.Address, pubkeys [][32]byte, chainId *big.Int) (common.Hash, error) {
	return _IPortal.Contract.EmergencyUnburnRequestToSolanaOverGsn(&_IPortal.TransactOpts, txID, from, pubkeys, chainId)
}

// EmergencyUnburnRequestToSolana is a paid mutator transaction binding the contract method 0xa9c4475f.
//
// Solidity: function emergencyUnburnRequestToSolana(bytes32 txID, address from, bytes32[] pubkeys, uint256 chainId) returns()
func (_IPortal *IPortalTransactorSession) EmergencyUnburnRequestToSolana(txID [32]byte, from common.Address, pubkeys [][32]byte, chainId *big.Int) (*types.Transaction, error) {
	return _IPortal.Contract.EmergencyUnburnRequestToSolana(&_IPortal.TransactOpts, txID, from, pubkeys, chainId)
}
func (_IPortal *IPortalTransactorSession) EmergencyUnburnRequestToSolanaOverGsn(txID [32]byte, from common.Address, pubkeys [][32]byte, chainId *big.Int) (common.Hash, error) {
	return _IPortal.Contract.EmergencyUnburnRequestToSolanaOverGsn(&_IPortal.TransactOpts, txID, from, pubkeys, chainId)
}

// Synthesize is a paid mutator transaction binding the contract method 0xf55090da.
//
// Solidity: function synthesize(address token, uint256 amount, address from, address to, (address,address,uint256) params) returns()
func (_IPortal *IPortalTransactor) Synthesize(opts *bind.TransactOpts, token common.Address, amount *big.Int, from common.Address, to common.Address, params IPortalSynthParams) (*types.Transaction, error) {
	return _IPortal.contract.Transact(opts, "synthesize", token, amount, from, to, params)
}
func (_IPortal *IPortalTransactor) SynthesizeOverGsn(opts *bind.TransactOpts, token common.Address, amount *big.Int, from common.Address, to common.Address, params IPortalSynthParams) (common.Hash, error) {
	return GsnExecutor(_IPortal.gsn, IPortalMetaData.ABI, "synthesize", token, amount, from, to, params)
}

// Synthesize is a paid mutator transaction binding the contract method 0xf55090da.
//
// Solidity: function synthesize(address token, uint256 amount, address from, address to, (address,address,uint256) params) returns()
func (_IPortal *IPortalSession) Synthesize(token common.Address, amount *big.Int, from common.Address, to common.Address, params IPortalSynthParams) (*types.Transaction, error) {
	return _IPortal.Contract.Synthesize(&_IPortal.TransactOpts, token, amount, from, to, params)
}
func (_IPortal *IPortalSession) SynthesizeOverGsn(token common.Address, amount *big.Int, from common.Address, to common.Address, params IPortalSynthParams) (common.Hash, error) {
	return _IPortal.Contract.SynthesizeOverGsn(&_IPortal.TransactOpts, token, amount, from, to, params)
}

// Synthesize is a paid mutator transaction binding the contract method 0xf55090da.
//
// Solidity: function synthesize(address token, uint256 amount, address from, address to, (address,address,uint256) params) returns()
func (_IPortal *IPortalTransactorSession) Synthesize(token common.Address, amount *big.Int, from common.Address, to common.Address, params IPortalSynthParams) (*types.Transaction, error) {
	return _IPortal.Contract.Synthesize(&_IPortal.TransactOpts, token, amount, from, to, params)
}
func (_IPortal *IPortalTransactorSession) SynthesizeOverGsn(token common.Address, amount *big.Int, from common.Address, to common.Address, params IPortalSynthParams) (common.Hash, error) {
	return _IPortal.Contract.SynthesizeOverGsn(&_IPortal.TransactOpts, token, amount, from, to, params)
}

// SynthesizeBatchWithDataTransit is a paid mutator transaction binding the contract method 0x2a3a615b.
//
// Solidity: function synthesizeBatchWithDataTransit(address[] token, uint256[] amount, address from, address to, (address,address,uint256) synthParams, (bytes4,bytes) transitData) returns()
func (_IPortal *IPortalTransactor) SynthesizeBatchWithDataTransit(opts *bind.TransactOpts, token []common.Address, amount []*big.Int, from common.Address, to common.Address, synthParams IPortalSynthParams, transitData IPortalTransitData) (*types.Transaction, error) {
	return _IPortal.contract.Transact(opts, "synthesizeBatchWithDataTransit", token, amount, from, to, synthParams, transitData)
}
func (_IPortal *IPortalTransactor) SynthesizeBatchWithDataTransitOverGsn(opts *bind.TransactOpts, token []common.Address, amount []*big.Int, from common.Address, to common.Address, synthParams IPortalSynthParams, transitData IPortalTransitData) (common.Hash, error) {
	return GsnExecutor(_IPortal.gsn, IPortalMetaData.ABI, "synthesizeBatchWithDataTransit", token, amount, from, to, synthParams, transitData)
}

// SynthesizeBatchWithDataTransit is a paid mutator transaction binding the contract method 0x2a3a615b.
//
// Solidity: function synthesizeBatchWithDataTransit(address[] token, uint256[] amount, address from, address to, (address,address,uint256) synthParams, (bytes4,bytes) transitData) returns()
func (_IPortal *IPortalSession) SynthesizeBatchWithDataTransit(token []common.Address, amount []*big.Int, from common.Address, to common.Address, synthParams IPortalSynthParams, transitData IPortalTransitData) (*types.Transaction, error) {
	return _IPortal.Contract.SynthesizeBatchWithDataTransit(&_IPortal.TransactOpts, token, amount, from, to, synthParams, transitData)
}
func (_IPortal *IPortalSession) SynthesizeBatchWithDataTransitOverGsn(token []common.Address, amount []*big.Int, from common.Address, to common.Address, synthParams IPortalSynthParams, transitData IPortalTransitData) (common.Hash, error) {
	return _IPortal.Contract.SynthesizeBatchWithDataTransitOverGsn(&_IPortal.TransactOpts, token, amount, from, to, synthParams, transitData)
}

// SynthesizeBatchWithDataTransit is a paid mutator transaction binding the contract method 0x2a3a615b.
//
// Solidity: function synthesizeBatchWithDataTransit(address[] token, uint256[] amount, address from, address to, (address,address,uint256) synthParams, (bytes4,bytes) transitData) returns()
func (_IPortal *IPortalTransactorSession) SynthesizeBatchWithDataTransit(token []common.Address, amount []*big.Int, from common.Address, to common.Address, synthParams IPortalSynthParams, transitData IPortalTransitData) (*types.Transaction, error) {
	return _IPortal.Contract.SynthesizeBatchWithDataTransit(&_IPortal.TransactOpts, token, amount, from, to, synthParams, transitData)
}
func (_IPortal *IPortalTransactorSession) SynthesizeBatchWithDataTransitOverGsn(token []common.Address, amount []*big.Int, from common.Address, to common.Address, synthParams IPortalSynthParams, transitData IPortalTransitData) (common.Hash, error) {
	return _IPortal.Contract.SynthesizeBatchWithDataTransitOverGsn(&_IPortal.TransactOpts, token, amount, from, to, synthParams, transitData)
}

// SynthesizeToSolana is a paid mutator transaction binding the contract method 0xac5de43c.
//
// Solidity: function synthesizeToSolana(address token, uint256 amount, address from, bytes32[] pubkeys, bytes1 txStateBump, uint256 chainId) returns()
func (_IPortal *IPortalTransactor) SynthesizeToSolana(opts *bind.TransactOpts, token common.Address, amount *big.Int, from common.Address, pubkeys [][32]byte, txStateBump [1]byte, chainId *big.Int) (*types.Transaction, error) {
	return _IPortal.contract.Transact(opts, "synthesizeToSolana", token, amount, from, pubkeys, txStateBump, chainId)
}
func (_IPortal *IPortalTransactor) SynthesizeToSolanaOverGsn(opts *bind.TransactOpts, token common.Address, amount *big.Int, from common.Address, pubkeys [][32]byte, txStateBump [1]byte, chainId *big.Int) (common.Hash, error) {
	return GsnExecutor(_IPortal.gsn, IPortalMetaData.ABI, "synthesizeToSolana", token, amount, from, pubkeys, txStateBump, chainId)
}

// SynthesizeToSolana is a paid mutator transaction binding the contract method 0xac5de43c.
//
// Solidity: function synthesizeToSolana(address token, uint256 amount, address from, bytes32[] pubkeys, bytes1 txStateBump, uint256 chainId) returns()
func (_IPortal *IPortalSession) SynthesizeToSolana(token common.Address, amount *big.Int, from common.Address, pubkeys [][32]byte, txStateBump [1]byte, chainId *big.Int) (*types.Transaction, error) {
	return _IPortal.Contract.SynthesizeToSolana(&_IPortal.TransactOpts, token, amount, from, pubkeys, txStateBump, chainId)
}
func (_IPortal *IPortalSession) SynthesizeToSolanaOverGsn(token common.Address, amount *big.Int, from common.Address, pubkeys [][32]byte, txStateBump [1]byte, chainId *big.Int) (common.Hash, error) {
	return _IPortal.Contract.SynthesizeToSolanaOverGsn(&_IPortal.TransactOpts, token, amount, from, pubkeys, txStateBump, chainId)
}

// SynthesizeToSolana is a paid mutator transaction binding the contract method 0xac5de43c.
//
// Solidity: function synthesizeToSolana(address token, uint256 amount, address from, bytes32[] pubkeys, bytes1 txStateBump, uint256 chainId) returns()
func (_IPortal *IPortalTransactorSession) SynthesizeToSolana(token common.Address, amount *big.Int, from common.Address, pubkeys [][32]byte, txStateBump [1]byte, chainId *big.Int) (*types.Transaction, error) {
	return _IPortal.Contract.SynthesizeToSolana(&_IPortal.TransactOpts, token, amount, from, pubkeys, txStateBump, chainId)
}
func (_IPortal *IPortalTransactorSession) SynthesizeToSolanaOverGsn(token common.Address, amount *big.Int, from common.Address, pubkeys [][32]byte, txStateBump [1]byte, chainId *big.Int) (common.Hash, error) {
	return _IPortal.Contract.SynthesizeToSolanaOverGsn(&_IPortal.TransactOpts, token, amount, from, pubkeys, txStateBump, chainId)
}

// SynthesizeWithPermit is a paid mutator transaction binding the contract method 0x5be21d97.
//
// Solidity: function synthesizeWithPermit((uint8,bytes32,bytes32,uint256,bool) permitData, address token, uint256 amount, address from, address to, address receiveSide, address oppositeBridge, uint256 chainId) returns()
func (_IPortal *IPortalTransactor) SynthesizeWithPermit(opts *bind.TransactOpts, permitData IPortalPermitData, token common.Address, amount *big.Int, from common.Address, to common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _IPortal.contract.Transact(opts, "synthesizeWithPermit", permitData, token, amount, from, to, receiveSide, oppositeBridge, chainId)
}
func (_IPortal *IPortalTransactor) SynthesizeWithPermitOverGsn(opts *bind.TransactOpts, permitData IPortalPermitData, token common.Address, amount *big.Int, from common.Address, to common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (common.Hash, error) {
	return GsnExecutor(_IPortal.gsn, IPortalMetaData.ABI, "synthesizeWithPermit", permitData, token, amount, from, to, receiveSide, oppositeBridge, chainId)
}

// SynthesizeWithPermit is a paid mutator transaction binding the contract method 0x5be21d97.
//
// Solidity: function synthesizeWithPermit((uint8,bytes32,bytes32,uint256,bool) permitData, address token, uint256 amount, address from, address to, address receiveSide, address oppositeBridge, uint256 chainId) returns()
func (_IPortal *IPortalSession) SynthesizeWithPermit(permitData IPortalPermitData, token common.Address, amount *big.Int, from common.Address, to common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _IPortal.Contract.SynthesizeWithPermit(&_IPortal.TransactOpts, permitData, token, amount, from, to, receiveSide, oppositeBridge, chainId)
}
func (_IPortal *IPortalSession) SynthesizeWithPermitOverGsn(permitData IPortalPermitData, token common.Address, amount *big.Int, from common.Address, to common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (common.Hash, error) {
	return _IPortal.Contract.SynthesizeWithPermitOverGsn(&_IPortal.TransactOpts, permitData, token, amount, from, to, receiveSide, oppositeBridge, chainId)
}

// SynthesizeWithPermit is a paid mutator transaction binding the contract method 0x5be21d97.
//
// Solidity: function synthesizeWithPermit((uint8,bytes32,bytes32,uint256,bool) permitData, address token, uint256 amount, address from, address to, address receiveSide, address oppositeBridge, uint256 chainId) returns()
func (_IPortal *IPortalTransactorSession) SynthesizeWithPermit(permitData IPortalPermitData, token common.Address, amount *big.Int, from common.Address, to common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _IPortal.Contract.SynthesizeWithPermit(&_IPortal.TransactOpts, permitData, token, amount, from, to, receiveSide, oppositeBridge, chainId)
}
func (_IPortal *IPortalTransactorSession) SynthesizeWithPermitOverGsn(permitData IPortalPermitData, token common.Address, amount *big.Int, from common.Address, to common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (common.Hash, error) {
	return _IPortal.Contract.SynthesizeWithPermitOverGsn(&_IPortal.TransactOpts, permitData, token, amount, from, to, receiveSide, oppositeBridge, chainId)
}
