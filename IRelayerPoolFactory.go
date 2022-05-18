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

// IRelayerPoolFactoryMetaData contains all meta data concerning the IRelayerPoolFactory contract.

var IRelayerPoolFactoryMetaData = &bind.MetaData{

	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_relayerFeeNumerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_emissionAnnualRateNumerator\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"contractRelayerPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IRelayerPoolFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use IRelayerPoolFactoryMetaData.ABI instead.
var IRelayerPoolFactoryABI = IRelayerPoolFactoryMetaData.ABI

// IRelayerPoolFactory is an auto generated Go binding around an Ethereum contract.
type IRelayerPoolFactory struct {
	IRelayerPoolFactoryCaller     // Read-only binding to the contract
	IRelayerPoolFactoryTransactor // Write-only binding to the contract
	IRelayerPoolFactoryFilterer   // Log filterer for contract events
}

// IRelayerPoolFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRelayerPoolFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRelayerPoolFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRelayerPoolFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
	gsn      *GsnCallOpts
}

func (_IRelayerPoolFactory *IRelayerPoolFactoryTransactor) SetGSNOptions(opts *GsnCallOpts) {
	_IRelayerPoolFactory.gsn = opts
}

// IRelayerPoolFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRelayerPoolFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRelayerPoolFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRelayerPoolFactorySession struct {
	Contract     *IRelayerPoolFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IRelayerPoolFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRelayerPoolFactoryCallerSession struct {
	Contract *IRelayerPoolFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IRelayerPoolFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRelayerPoolFactoryTransactorSession struct {
	Contract     *IRelayerPoolFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IRelayerPoolFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRelayerPoolFactoryRaw struct {
	Contract *IRelayerPoolFactory // Generic contract binding to access the raw methods on
}

// IRelayerPoolFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRelayerPoolFactoryCallerRaw struct {
	Contract *IRelayerPoolFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IRelayerPoolFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRelayerPoolFactoryTransactorRaw struct {
	Contract *IRelayerPoolFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRelayerPoolFactory creates a new instance of IRelayerPoolFactory, bound to a specific deployed contract.
func NewIRelayerPoolFactory(address common.Address, backend bind.ContractBackend) (*IRelayerPoolFactory, error) {
	contract, err := bindIRelayerPoolFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRelayerPoolFactory{IRelayerPoolFactoryCaller: IRelayerPoolFactoryCaller{contract: contract}, IRelayerPoolFactoryTransactor: IRelayerPoolFactoryTransactor{contract: contract}, IRelayerPoolFactoryFilterer: IRelayerPoolFactoryFilterer{contract: contract}}, nil
}

// NewIRelayerPoolFactoryCaller creates a new read-only instance of IRelayerPoolFactory, bound to a specific deployed contract.
func NewIRelayerPoolFactoryCaller(address common.Address, caller bind.ContractCaller) (*IRelayerPoolFactoryCaller, error) {
	contract, err := bindIRelayerPoolFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRelayerPoolFactoryCaller{contract: contract}, nil
}

// NewIRelayerPoolFactoryTransactor creates a new write-only instance of IRelayerPoolFactory, bound to a specific deployed contract.
func NewIRelayerPoolFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IRelayerPoolFactoryTransactor, error) {
	contract, err := bindIRelayerPoolFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRelayerPoolFactoryTransactor{contract: contract}, nil
}

// NewIRelayerPoolFactoryFilterer creates a new log filterer instance of IRelayerPoolFactory, bound to a specific deployed contract.
func NewIRelayerPoolFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IRelayerPoolFactoryFilterer, error) {
	contract, err := bindIRelayerPoolFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRelayerPoolFactoryFilterer{contract: contract}, nil
}

// bindIRelayerPoolFactory binds a generic wrapper to an already deployed contract.
func bindIRelayerPoolFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IRelayerPoolFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRelayerPoolFactory *IRelayerPoolFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRelayerPoolFactory.Contract.IRelayerPoolFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRelayerPoolFactory *IRelayerPoolFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRelayerPoolFactory.Contract.IRelayerPoolFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRelayerPoolFactory *IRelayerPoolFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRelayerPoolFactory.Contract.IRelayerPoolFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRelayerPoolFactory *IRelayerPoolFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRelayerPoolFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRelayerPoolFactory *IRelayerPoolFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRelayerPoolFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRelayerPoolFactory *IRelayerPoolFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRelayerPoolFactory.Contract.contract.Transact(opts, method, params...)
}

// Create is a paid mutator transaction binding the contract method 0xcfcae1a5.
//
// Solidity: function create(address _owner, address _rewardToken, address _depositToken, uint256 _relayerFeeNumerator, uint256 _emissionAnnualRateNumerator, address _vault) returns(address)
func (_IRelayerPoolFactory *IRelayerPoolFactoryTransactor) Create(opts *bind.TransactOpts, _owner common.Address, _rewardToken common.Address, _depositToken common.Address, _relayerFeeNumerator *big.Int, _emissionAnnualRateNumerator *big.Int, _vault common.Address) (*types.Transaction, error) {
	return _IRelayerPoolFactory.contract.Transact(opts, "create", _owner, _rewardToken, _depositToken, _relayerFeeNumerator, _emissionAnnualRateNumerator, _vault)
}
func (_IRelayerPoolFactory *IRelayerPoolFactoryTransactor) CreateOverGsn(opts *bind.TransactOpts, _owner common.Address, _rewardToken common.Address, _depositToken common.Address, _relayerFeeNumerator *big.Int, _emissionAnnualRateNumerator *big.Int, _vault common.Address) (common.Hash, error) {
	return GsnExecutor(_IRelayerPoolFactory.gsn, IRelayerPoolFactoryMetaData.ABI, "create", _owner, _rewardToken, _depositToken, _relayerFeeNumerator, _emissionAnnualRateNumerator, _vault)
}

// Create is a paid mutator transaction binding the contract method 0xcfcae1a5.
//
// Solidity: function create(address _owner, address _rewardToken, address _depositToken, uint256 _relayerFeeNumerator, uint256 _emissionAnnualRateNumerator, address _vault) returns(address)
func (_IRelayerPoolFactory *IRelayerPoolFactorySession) Create(_owner common.Address, _rewardToken common.Address, _depositToken common.Address, _relayerFeeNumerator *big.Int, _emissionAnnualRateNumerator *big.Int, _vault common.Address) (*types.Transaction, error) {
	return _IRelayerPoolFactory.Contract.Create(&_IRelayerPoolFactory.TransactOpts, _owner, _rewardToken, _depositToken, _relayerFeeNumerator, _emissionAnnualRateNumerator, _vault)
}
func (_IRelayerPoolFactory *IRelayerPoolFactorySession) CreateOverGsn(_owner common.Address, _rewardToken common.Address, _depositToken common.Address, _relayerFeeNumerator *big.Int, _emissionAnnualRateNumerator *big.Int, _vault common.Address) (common.Hash, error) {
	return _IRelayerPoolFactory.Contract.CreateOverGsn(&_IRelayerPoolFactory.TransactOpts, _owner, _rewardToken, _depositToken, _relayerFeeNumerator, _emissionAnnualRateNumerator, _vault)
}

// Create is a paid mutator transaction binding the contract method 0xcfcae1a5.
//
// Solidity: function create(address _owner, address _rewardToken, address _depositToken, uint256 _relayerFeeNumerator, uint256 _emissionAnnualRateNumerator, address _vault) returns(address)
func (_IRelayerPoolFactory *IRelayerPoolFactoryTransactorSession) Create(_owner common.Address, _rewardToken common.Address, _depositToken common.Address, _relayerFeeNumerator *big.Int, _emissionAnnualRateNumerator *big.Int, _vault common.Address) (*types.Transaction, error) {
	return _IRelayerPoolFactory.Contract.Create(&_IRelayerPoolFactory.TransactOpts, _owner, _rewardToken, _depositToken, _relayerFeeNumerator, _emissionAnnualRateNumerator, _vault)
}
func (_IRelayerPoolFactory *IRelayerPoolFactoryTransactorSession) CreateOverGsn(_owner common.Address, _rewardToken common.Address, _depositToken common.Address, _relayerFeeNumerator *big.Int, _emissionAnnualRateNumerator *big.Int, _vault common.Address) (common.Hash, error) {
	return _IRelayerPoolFactory.Contract.CreateOverGsn(&_IRelayerPoolFactory.TransactOpts, _owner, _rewardToken, _depositToken, _relayerFeeNumerator, _emissionAnnualRateNumerator, _vault)
}
