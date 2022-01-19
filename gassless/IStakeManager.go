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

// IStakeManagerStakeInfo is an auto generated low-level Go binding around an user-defined struct.
type IStakeManagerStakeInfo struct {
	Stake         *big.Int
	UnstakeDelay  *big.Int
	WithdrawBlock *big.Int
	Owner         common.Address
}

// IStakeManagerABI is the input ABI used to generate the binding from.
const IStakeManagerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayHub\",\"type\":\"address\"}],\"name\":\"HubAuthorized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayHub\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"removalBlock\",\"type\":\"uint256\"}],\"name\":\"HubUnauthorized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unstakeDelay\",\"type\":\"uint256\"}],\"name\":\"StakeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"StakePenalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawBlock\",\"type\":\"uint256\"}],\"name\":\"StakeUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayHub\",\"type\":\"address\"}],\"name\":\"authorizeHubByManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"relayHub\",\"type\":\"address\"}],\"name\":\"authorizeHubByOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayManager\",\"type\":\"address\"}],\"name\":\"getStakeInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawBlock\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structIStakeManager.StakeInfo\",\"name\":\"stakeInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"relayHub\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minUnstakeDelay\",\"type\":\"uint256\"}],\"name\":\"isRelayManagerStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxUnstakeDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayManager\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"penalizeRelayManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setRelayManagerOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"unstakeDelay\",\"type\":\"uint256\"}],\"name\":\"stakeForRelayManager\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayHub\",\"type\":\"address\"}],\"name\":\"unauthorizeHubByManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"relayHub\",\"type\":\"address\"}],\"name\":\"unauthorizeHubByOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayManager\",\"type\":\"address\"}],\"name\":\"unlockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionSM\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayManager\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IStakeManager is an auto generated Go binding around an Ethereum contract.
type IStakeManager struct {
	IStakeManagerCaller     // Read-only binding to the contract
	IStakeManagerTransactor // Write-only binding to the contract
	IStakeManagerFilterer   // Log filterer for contract events
}

// IStakeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStakeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStakeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStakeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStakeManagerSession struct {
	Contract     *IStakeManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IStakeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStakeManagerCallerSession struct {
	Contract *IStakeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IStakeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStakeManagerTransactorSession struct {
	Contract     *IStakeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IStakeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStakeManagerRaw struct {
	Contract *IStakeManager // Generic contract binding to access the raw methods on
}

// IStakeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStakeManagerCallerRaw struct {
	Contract *IStakeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IStakeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStakeManagerTransactorRaw struct {
	Contract *IStakeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStakeManager creates a new instance of IStakeManager, bound to a specific deployed contract.
func NewIStakeManager(address common.Address, backend bind.ContractBackend) (*IStakeManager, error) {
	contract, err := bindIStakeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStakeManager{IStakeManagerCaller: IStakeManagerCaller{contract: contract}, IStakeManagerTransactor: IStakeManagerTransactor{contract: contract}, IStakeManagerFilterer: IStakeManagerFilterer{contract: contract}}, nil
}

// NewIStakeManagerCaller creates a new read-only instance of IStakeManager, bound to a specific deployed contract.
func NewIStakeManagerCaller(address common.Address, caller bind.ContractCaller) (*IStakeManagerCaller, error) {
	contract, err := bindIStakeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStakeManagerCaller{contract: contract}, nil
}

// NewIStakeManagerTransactor creates a new write-only instance of IStakeManager, bound to a specific deployed contract.
func NewIStakeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IStakeManagerTransactor, error) {
	contract, err := bindIStakeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStakeManagerTransactor{contract: contract}, nil
}

// NewIStakeManagerFilterer creates a new log filterer instance of IStakeManager, bound to a specific deployed contract.
func NewIStakeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IStakeManagerFilterer, error) {
	contract, err := bindIStakeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStakeManagerFilterer{contract: contract}, nil
}

// bindIStakeManager binds a generic wrapper to an already deployed contract.
func bindIStakeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IStakeManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStakeManager *IStakeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStakeManager.Contract.IStakeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStakeManager *IStakeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakeManager.Contract.IStakeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStakeManager *IStakeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStakeManager.Contract.IStakeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStakeManager *IStakeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStakeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStakeManager *IStakeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStakeManager *IStakeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStakeManager.Contract.contract.Transact(opts, method, params...)
}

// GetStakeInfo is a free data retrieval call binding the contract method 0xc3453153.
//
// Solidity: function getStakeInfo(address relayManager) view returns((uint256,uint256,uint256,address) stakeInfo)
func (_IStakeManager *IStakeManagerCaller) GetStakeInfo(opts *bind.CallOpts, relayManager common.Address) (IStakeManagerStakeInfo, error) {
	var out []interface{}
	err := _IStakeManager.contract.Call(opts, &out, "getStakeInfo", relayManager)

	if err != nil {
		return *new(IStakeManagerStakeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IStakeManagerStakeInfo)).(*IStakeManagerStakeInfo)

	return out0, err

}

// GetStakeInfo is a free data retrieval call binding the contract method 0xc3453153.
//
// Solidity: function getStakeInfo(address relayManager) view returns((uint256,uint256,uint256,address) stakeInfo)
func (_IStakeManager *IStakeManagerSession) GetStakeInfo(relayManager common.Address) (IStakeManagerStakeInfo, error) {
	return _IStakeManager.Contract.GetStakeInfo(&_IStakeManager.CallOpts, relayManager)
}

// GetStakeInfo is a free data retrieval call binding the contract method 0xc3453153.
//
// Solidity: function getStakeInfo(address relayManager) view returns((uint256,uint256,uint256,address) stakeInfo)
func (_IStakeManager *IStakeManagerCallerSession) GetStakeInfo(relayManager common.Address) (IStakeManagerStakeInfo, error) {
	return _IStakeManager.Contract.GetStakeInfo(&_IStakeManager.CallOpts, relayManager)
}

// IsRelayManagerStaked is a free data retrieval call binding the contract method 0x6de8dd41.
//
// Solidity: function isRelayManagerStaked(address relayManager, address relayHub, uint256 minAmount, uint256 minUnstakeDelay) view returns(bool)
func (_IStakeManager *IStakeManagerCaller) IsRelayManagerStaked(opts *bind.CallOpts, relayManager common.Address, relayHub common.Address, minAmount *big.Int, minUnstakeDelay *big.Int) (bool, error) {
	var out []interface{}
	err := _IStakeManager.contract.Call(opts, &out, "isRelayManagerStaked", relayManager, relayHub, minAmount, minUnstakeDelay)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRelayManagerStaked is a free data retrieval call binding the contract method 0x6de8dd41.
//
// Solidity: function isRelayManagerStaked(address relayManager, address relayHub, uint256 minAmount, uint256 minUnstakeDelay) view returns(bool)
func (_IStakeManager *IStakeManagerSession) IsRelayManagerStaked(relayManager common.Address, relayHub common.Address, minAmount *big.Int, minUnstakeDelay *big.Int) (bool, error) {
	return _IStakeManager.Contract.IsRelayManagerStaked(&_IStakeManager.CallOpts, relayManager, relayHub, minAmount, minUnstakeDelay)
}

// IsRelayManagerStaked is a free data retrieval call binding the contract method 0x6de8dd41.
//
// Solidity: function isRelayManagerStaked(address relayManager, address relayHub, uint256 minAmount, uint256 minUnstakeDelay) view returns(bool)
func (_IStakeManager *IStakeManagerCallerSession) IsRelayManagerStaked(relayManager common.Address, relayHub common.Address, minAmount *big.Int, minUnstakeDelay *big.Int) (bool, error) {
	return _IStakeManager.Contract.IsRelayManagerStaked(&_IStakeManager.CallOpts, relayManager, relayHub, minAmount, minUnstakeDelay)
}

// MaxUnstakeDelay is a free data retrieval call binding the contract method 0x4e02131c.
//
// Solidity: function maxUnstakeDelay() view returns(uint256)
func (_IStakeManager *IStakeManagerCaller) MaxUnstakeDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IStakeManager.contract.Call(opts, &out, "maxUnstakeDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxUnstakeDelay is a free data retrieval call binding the contract method 0x4e02131c.
//
// Solidity: function maxUnstakeDelay() view returns(uint256)
func (_IStakeManager *IStakeManagerSession) MaxUnstakeDelay() (*big.Int, error) {
	return _IStakeManager.Contract.MaxUnstakeDelay(&_IStakeManager.CallOpts)
}

// MaxUnstakeDelay is a free data retrieval call binding the contract method 0x4e02131c.
//
// Solidity: function maxUnstakeDelay() view returns(uint256)
func (_IStakeManager *IStakeManagerCallerSession) MaxUnstakeDelay() (*big.Int, error) {
	return _IStakeManager.Contract.MaxUnstakeDelay(&_IStakeManager.CallOpts)
}

// VersionSM is a free data retrieval call binding the contract method 0x47116c6e.
//
// Solidity: function versionSM() view returns(string)
func (_IStakeManager *IStakeManagerCaller) VersionSM(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IStakeManager.contract.Call(opts, &out, "versionSM")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VersionSM is a free data retrieval call binding the contract method 0x47116c6e.
//
// Solidity: function versionSM() view returns(string)
func (_IStakeManager *IStakeManagerSession) VersionSM() (string, error) {
	return _IStakeManager.Contract.VersionSM(&_IStakeManager.CallOpts)
}

// VersionSM is a free data retrieval call binding the contract method 0x47116c6e.
//
// Solidity: function versionSM() view returns(string)
func (_IStakeManager *IStakeManagerCallerSession) VersionSM() (string, error) {
	return _IStakeManager.Contract.VersionSM(&_IStakeManager.CallOpts)
}

// AuthorizeHubByManager is a paid mutator transaction binding the contract method 0xd48a9d43.
//
// Solidity: function authorizeHubByManager(address relayHub) returns()
func (_IStakeManager *IStakeManagerTransactor) AuthorizeHubByManager(opts *bind.TransactOpts, relayHub common.Address) (*types.Transaction, error) {
	return _IStakeManager.contract.Transact(opts, "authorizeHubByManager", relayHub)
}

// AuthorizeHubByManager is a paid mutator transaction binding the contract method 0xd48a9d43.
//
// Solidity: function authorizeHubByManager(address relayHub) returns()
func (_IStakeManager *IStakeManagerSession) AuthorizeHubByManager(relayHub common.Address) (*types.Transaction, error) {
	return _IStakeManager.Contract.AuthorizeHubByManager(&_IStakeManager.TransactOpts, relayHub)
}

// AuthorizeHubByManager is a paid mutator transaction binding the contract method 0xd48a9d43.
//
// Solidity: function authorizeHubByManager(address relayHub) returns()
func (_IStakeManager *IStakeManagerTransactorSession) AuthorizeHubByManager(relayHub common.Address) (*types.Transaction, error) {
	return _IStakeManager.Contract.AuthorizeHubByManager(&_IStakeManager.TransactOpts, relayHub)
}

// AuthorizeHubByOwner is a paid mutator transaction binding the contract method 0x7835d296.
//
// Solidity: function authorizeHubByOwner(address relayManager, address relayHub) returns()
func (_IStakeManager *IStakeManagerTransactor) AuthorizeHubByOwner(opts *bind.TransactOpts, relayManager common.Address, relayHub common.Address) (*types.Transaction, error) {
	return _IStakeManager.contract.Transact(opts, "authorizeHubByOwner", relayManager, relayHub)
}

// AuthorizeHubByOwner is a paid mutator transaction binding the contract method 0x7835d296.
//
// Solidity: function authorizeHubByOwner(address relayManager, address relayHub) returns()
func (_IStakeManager *IStakeManagerSession) AuthorizeHubByOwner(relayManager common.Address, relayHub common.Address) (*types.Transaction, error) {
	return _IStakeManager.Contract.AuthorizeHubByOwner(&_IStakeManager.TransactOpts, relayManager, relayHub)
}

// AuthorizeHubByOwner is a paid mutator transaction binding the contract method 0x7835d296.
//
// Solidity: function authorizeHubByOwner(address relayManager, address relayHub) returns()
func (_IStakeManager *IStakeManagerTransactorSession) AuthorizeHubByOwner(relayManager common.Address, relayHub common.Address) (*types.Transaction, error) {
	return _IStakeManager.Contract.AuthorizeHubByOwner(&_IStakeManager.TransactOpts, relayManager, relayHub)
}

// PenalizeRelayManager is a paid mutator transaction binding the contract method 0x09a08af4.
//
// Solidity: function penalizeRelayManager(address relayManager, address beneficiary, uint256 amount) returns()
func (_IStakeManager *IStakeManagerTransactor) PenalizeRelayManager(opts *bind.TransactOpts, relayManager common.Address, beneficiary common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStakeManager.contract.Transact(opts, "penalizeRelayManager", relayManager, beneficiary, amount)
}

// PenalizeRelayManager is a paid mutator transaction binding the contract method 0x09a08af4.
//
// Solidity: function penalizeRelayManager(address relayManager, address beneficiary, uint256 amount) returns()
func (_IStakeManager *IStakeManagerSession) PenalizeRelayManager(relayManager common.Address, beneficiary common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStakeManager.Contract.PenalizeRelayManager(&_IStakeManager.TransactOpts, relayManager, beneficiary, amount)
}

// PenalizeRelayManager is a paid mutator transaction binding the contract method 0x09a08af4.
//
// Solidity: function penalizeRelayManager(address relayManager, address beneficiary, uint256 amount) returns()
func (_IStakeManager *IStakeManagerTransactorSession) PenalizeRelayManager(relayManager common.Address, beneficiary common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStakeManager.Contract.PenalizeRelayManager(&_IStakeManager.TransactOpts, relayManager, beneficiary, amount)
}

// SetRelayManagerOwner is a paid mutator transaction binding the contract method 0xfece3dd4.
//
// Solidity: function setRelayManagerOwner(address owner) returns()
func (_IStakeManager *IStakeManagerTransactor) SetRelayManagerOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _IStakeManager.contract.Transact(opts, "setRelayManagerOwner", owner)
}

// SetRelayManagerOwner is a paid mutator transaction binding the contract method 0xfece3dd4.
//
// Solidity: function setRelayManagerOwner(address owner) returns()
func (_IStakeManager *IStakeManagerSession) SetRelayManagerOwner(owner common.Address) (*types.Transaction, error) {
	return _IStakeManager.Contract.SetRelayManagerOwner(&_IStakeManager.TransactOpts, owner)
}

// SetRelayManagerOwner is a paid mutator transaction binding the contract method 0xfece3dd4.
//
// Solidity: function setRelayManagerOwner(address owner) returns()
func (_IStakeManager *IStakeManagerTransactorSession) SetRelayManagerOwner(owner common.Address) (*types.Transaction, error) {
	return _IStakeManager.Contract.SetRelayManagerOwner(&_IStakeManager.TransactOpts, owner)
}

// StakeForRelayManager is a paid mutator transaction binding the contract method 0xf32102db.
//
// Solidity: function stakeForRelayManager(address relayManager, uint256 unstakeDelay) payable returns()
func (_IStakeManager *IStakeManagerTransactor) StakeForRelayManager(opts *bind.TransactOpts, relayManager common.Address, unstakeDelay *big.Int) (*types.Transaction, error) {
	return _IStakeManager.contract.Transact(opts, "stakeForRelayManager", relayManager, unstakeDelay)
}

// StakeForRelayManager is a paid mutator transaction binding the contract method 0xf32102db.
//
// Solidity: function stakeForRelayManager(address relayManager, uint256 unstakeDelay) payable returns()
func (_IStakeManager *IStakeManagerSession) StakeForRelayManager(relayManager common.Address, unstakeDelay *big.Int) (*types.Transaction, error) {
	return _IStakeManager.Contract.StakeForRelayManager(&_IStakeManager.TransactOpts, relayManager, unstakeDelay)
}

// StakeForRelayManager is a paid mutator transaction binding the contract method 0xf32102db.
//
// Solidity: function stakeForRelayManager(address relayManager, uint256 unstakeDelay) payable returns()
func (_IStakeManager *IStakeManagerTransactorSession) StakeForRelayManager(relayManager common.Address, unstakeDelay *big.Int) (*types.Transaction, error) {
	return _IStakeManager.Contract.StakeForRelayManager(&_IStakeManager.TransactOpts, relayManager, unstakeDelay)
}

// UnauthorizeHubByManager is a paid mutator transaction binding the contract method 0xf9bce311.
//
// Solidity: function unauthorizeHubByManager(address relayHub) returns()
func (_IStakeManager *IStakeManagerTransactor) UnauthorizeHubByManager(opts *bind.TransactOpts, relayHub common.Address) (*types.Transaction, error) {
	return _IStakeManager.contract.Transact(opts, "unauthorizeHubByManager", relayHub)
}

// UnauthorizeHubByManager is a paid mutator transaction binding the contract method 0xf9bce311.
//
// Solidity: function unauthorizeHubByManager(address relayHub) returns()
func (_IStakeManager *IStakeManagerSession) UnauthorizeHubByManager(relayHub common.Address) (*types.Transaction, error) {
	return _IStakeManager.Contract.UnauthorizeHubByManager(&_IStakeManager.TransactOpts, relayHub)
}

// UnauthorizeHubByManager is a paid mutator transaction binding the contract method 0xf9bce311.
//
// Solidity: function unauthorizeHubByManager(address relayHub) returns()
func (_IStakeManager *IStakeManagerTransactorSession) UnauthorizeHubByManager(relayHub common.Address) (*types.Transaction, error) {
	return _IStakeManager.Contract.UnauthorizeHubByManager(&_IStakeManager.TransactOpts, relayHub)
}

// UnauthorizeHubByOwner is a paid mutator transaction binding the contract method 0xf48f8ac7.
//
// Solidity: function unauthorizeHubByOwner(address relayManager, address relayHub) returns()
func (_IStakeManager *IStakeManagerTransactor) UnauthorizeHubByOwner(opts *bind.TransactOpts, relayManager common.Address, relayHub common.Address) (*types.Transaction, error) {
	return _IStakeManager.contract.Transact(opts, "unauthorizeHubByOwner", relayManager, relayHub)
}

// UnauthorizeHubByOwner is a paid mutator transaction binding the contract method 0xf48f8ac7.
//
// Solidity: function unauthorizeHubByOwner(address relayManager, address relayHub) returns()
func (_IStakeManager *IStakeManagerSession) UnauthorizeHubByOwner(relayManager common.Address, relayHub common.Address) (*types.Transaction, error) {
	return _IStakeManager.Contract.UnauthorizeHubByOwner(&_IStakeManager.TransactOpts, relayManager, relayHub)
}

// UnauthorizeHubByOwner is a paid mutator transaction binding the contract method 0xf48f8ac7.
//
// Solidity: function unauthorizeHubByOwner(address relayManager, address relayHub) returns()
func (_IStakeManager *IStakeManagerTransactorSession) UnauthorizeHubByOwner(relayManager common.Address, relayHub common.Address) (*types.Transaction, error) {
	return _IStakeManager.Contract.UnauthorizeHubByOwner(&_IStakeManager.TransactOpts, relayManager, relayHub)
}

// UnlockStake is a paid mutator transaction binding the contract method 0x4a1ce599.
//
// Solidity: function unlockStake(address relayManager) returns()
func (_IStakeManager *IStakeManagerTransactor) UnlockStake(opts *bind.TransactOpts, relayManager common.Address) (*types.Transaction, error) {
	return _IStakeManager.contract.Transact(opts, "unlockStake", relayManager)
}

// UnlockStake is a paid mutator transaction binding the contract method 0x4a1ce599.
//
// Solidity: function unlockStake(address relayManager) returns()
func (_IStakeManager *IStakeManagerSession) UnlockStake(relayManager common.Address) (*types.Transaction, error) {
	return _IStakeManager.Contract.UnlockStake(&_IStakeManager.TransactOpts, relayManager)
}

// UnlockStake is a paid mutator transaction binding the contract method 0x4a1ce599.
//
// Solidity: function unlockStake(address relayManager) returns()
func (_IStakeManager *IStakeManagerTransactorSession) UnlockStake(relayManager common.Address) (*types.Transaction, error) {
	return _IStakeManager.Contract.UnlockStake(&_IStakeManager.TransactOpts, relayManager)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address relayManager) returns()
func (_IStakeManager *IStakeManagerTransactor) WithdrawStake(opts *bind.TransactOpts, relayManager common.Address) (*types.Transaction, error) {
	return _IStakeManager.contract.Transact(opts, "withdrawStake", relayManager)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address relayManager) returns()
func (_IStakeManager *IStakeManagerSession) WithdrawStake(relayManager common.Address) (*types.Transaction, error) {
	return _IStakeManager.Contract.WithdrawStake(&_IStakeManager.TransactOpts, relayManager)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address relayManager) returns()
func (_IStakeManager *IStakeManagerTransactorSession) WithdrawStake(relayManager common.Address) (*types.Transaction, error) {
	return _IStakeManager.Contract.WithdrawStake(&_IStakeManager.TransactOpts, relayManager)
}

// IStakeManagerHubAuthorizedIterator is returned from FilterHubAuthorized and is used to iterate over the raw logs and unpacked data for HubAuthorized events raised by the IStakeManager contract.
type IStakeManagerHubAuthorizedIterator struct {
	Event *IStakeManagerHubAuthorized // Event containing the contract specifics and raw log

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
func (it *IStakeManagerHubAuthorizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakeManagerHubAuthorized)
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
		it.Event = new(IStakeManagerHubAuthorized)
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
func (it *IStakeManagerHubAuthorizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakeManagerHubAuthorizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakeManagerHubAuthorized represents a HubAuthorized event raised by the IStakeManager contract.
type IStakeManagerHubAuthorized struct {
	RelayManager common.Address
	RelayHub     common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterHubAuthorized is a free log retrieval operation binding the contract event 0xe292c4f6e9f34c975f4958cd5650a8111352feae914a67b06407957105421021.
//
// Solidity: event HubAuthorized(address indexed relayManager, address indexed relayHub)
func (_IStakeManager *IStakeManagerFilterer) FilterHubAuthorized(opts *bind.FilterOpts, relayManager []common.Address, relayHub []common.Address) (*IStakeManagerHubAuthorizedIterator, error) {

	var relayManagerRule []interface{}
	for _, relayManagerItem := range relayManager {
		relayManagerRule = append(relayManagerRule, relayManagerItem)
	}
	var relayHubRule []interface{}
	for _, relayHubItem := range relayHub {
		relayHubRule = append(relayHubRule, relayHubItem)
	}

	logs, sub, err := _IStakeManager.contract.FilterLogs(opts, "HubAuthorized", relayManagerRule, relayHubRule)
	if err != nil {
		return nil, err
	}
	return &IStakeManagerHubAuthorizedIterator{contract: _IStakeManager.contract, event: "HubAuthorized", logs: logs, sub: sub}, nil
}

// WatchHubAuthorized is a free log subscription operation binding the contract event 0xe292c4f6e9f34c975f4958cd5650a8111352feae914a67b06407957105421021.
//
// Solidity: event HubAuthorized(address indexed relayManager, address indexed relayHub)
func (_IStakeManager *IStakeManagerFilterer) WatchHubAuthorized(opts *bind.WatchOpts, sink chan<- *IStakeManagerHubAuthorized, relayManager []common.Address, relayHub []common.Address) (event.Subscription, error) {

	var relayManagerRule []interface{}
	for _, relayManagerItem := range relayManager {
		relayManagerRule = append(relayManagerRule, relayManagerItem)
	}
	var relayHubRule []interface{}
	for _, relayHubItem := range relayHub {
		relayHubRule = append(relayHubRule, relayHubItem)
	}

	logs, sub, err := _IStakeManager.contract.WatchLogs(opts, "HubAuthorized", relayManagerRule, relayHubRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakeManagerHubAuthorized)
				if err := _IStakeManager.contract.UnpackLog(event, "HubAuthorized", log); err != nil {
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

// ParseHubAuthorized is a log parse operation binding the contract event 0xe292c4f6e9f34c975f4958cd5650a8111352feae914a67b06407957105421021.
//
// Solidity: event HubAuthorized(address indexed relayManager, address indexed relayHub)
func (_IStakeManager *IStakeManagerFilterer) ParseHubAuthorized(log types.Log) (*IStakeManagerHubAuthorized, error) {
	event := new(IStakeManagerHubAuthorized)
	if err := _IStakeManager.contract.UnpackLog(event, "HubAuthorized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakeManagerHubUnauthorizedIterator is returned from FilterHubUnauthorized and is used to iterate over the raw logs and unpacked data for HubUnauthorized events raised by the IStakeManager contract.
type IStakeManagerHubUnauthorizedIterator struct {
	Event *IStakeManagerHubUnauthorized // Event containing the contract specifics and raw log

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
func (it *IStakeManagerHubUnauthorizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakeManagerHubUnauthorized)
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
		it.Event = new(IStakeManagerHubUnauthorized)
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
func (it *IStakeManagerHubUnauthorizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakeManagerHubUnauthorizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakeManagerHubUnauthorized represents a HubUnauthorized event raised by the IStakeManager contract.
type IStakeManagerHubUnauthorized struct {
	RelayManager common.Address
	RelayHub     common.Address
	RemovalBlock *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterHubUnauthorized is a free log retrieval operation binding the contract event 0x8d941c9b73ba7e59671a59eed85054004624684182b0e4bdb56c35937bac65a6.
//
// Solidity: event HubUnauthorized(address indexed relayManager, address indexed relayHub, uint256 removalBlock)
func (_IStakeManager *IStakeManagerFilterer) FilterHubUnauthorized(opts *bind.FilterOpts, relayManager []common.Address, relayHub []common.Address) (*IStakeManagerHubUnauthorizedIterator, error) {

	var relayManagerRule []interface{}
	for _, relayManagerItem := range relayManager {
		relayManagerRule = append(relayManagerRule, relayManagerItem)
	}
	var relayHubRule []interface{}
	for _, relayHubItem := range relayHub {
		relayHubRule = append(relayHubRule, relayHubItem)
	}

	logs, sub, err := _IStakeManager.contract.FilterLogs(opts, "HubUnauthorized", relayManagerRule, relayHubRule)
	if err != nil {
		return nil, err
	}
	return &IStakeManagerHubUnauthorizedIterator{contract: _IStakeManager.contract, event: "HubUnauthorized", logs: logs, sub: sub}, nil
}

// WatchHubUnauthorized is a free log subscription operation binding the contract event 0x8d941c9b73ba7e59671a59eed85054004624684182b0e4bdb56c35937bac65a6.
//
// Solidity: event HubUnauthorized(address indexed relayManager, address indexed relayHub, uint256 removalBlock)
func (_IStakeManager *IStakeManagerFilterer) WatchHubUnauthorized(opts *bind.WatchOpts, sink chan<- *IStakeManagerHubUnauthorized, relayManager []common.Address, relayHub []common.Address) (event.Subscription, error) {

	var relayManagerRule []interface{}
	for _, relayManagerItem := range relayManager {
		relayManagerRule = append(relayManagerRule, relayManagerItem)
	}
	var relayHubRule []interface{}
	for _, relayHubItem := range relayHub {
		relayHubRule = append(relayHubRule, relayHubItem)
	}

	logs, sub, err := _IStakeManager.contract.WatchLogs(opts, "HubUnauthorized", relayManagerRule, relayHubRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakeManagerHubUnauthorized)
				if err := _IStakeManager.contract.UnpackLog(event, "HubUnauthorized", log); err != nil {
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

// ParseHubUnauthorized is a log parse operation binding the contract event 0x8d941c9b73ba7e59671a59eed85054004624684182b0e4bdb56c35937bac65a6.
//
// Solidity: event HubUnauthorized(address indexed relayManager, address indexed relayHub, uint256 removalBlock)
func (_IStakeManager *IStakeManagerFilterer) ParseHubUnauthorized(log types.Log) (*IStakeManagerHubUnauthorized, error) {
	event := new(IStakeManagerHubUnauthorized)
	if err := _IStakeManager.contract.UnpackLog(event, "HubUnauthorized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakeManagerOwnerSetIterator is returned from FilterOwnerSet and is used to iterate over the raw logs and unpacked data for OwnerSet events raised by the IStakeManager contract.
type IStakeManagerOwnerSetIterator struct {
	Event *IStakeManagerOwnerSet // Event containing the contract specifics and raw log

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
func (it *IStakeManagerOwnerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakeManagerOwnerSet)
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
		it.Event = new(IStakeManagerOwnerSet)
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
func (it *IStakeManagerOwnerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakeManagerOwnerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakeManagerOwnerSet represents a OwnerSet event raised by the IStakeManager contract.
type IStakeManagerOwnerSet struct {
	RelayManager common.Address
	Owner        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnerSet is a free log retrieval operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed relayManager, address indexed owner)
func (_IStakeManager *IStakeManagerFilterer) FilterOwnerSet(opts *bind.FilterOpts, relayManager []common.Address, owner []common.Address) (*IStakeManagerOwnerSetIterator, error) {

	var relayManagerRule []interface{}
	for _, relayManagerItem := range relayManager {
		relayManagerRule = append(relayManagerRule, relayManagerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _IStakeManager.contract.FilterLogs(opts, "OwnerSet", relayManagerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &IStakeManagerOwnerSetIterator{contract: _IStakeManager.contract, event: "OwnerSet", logs: logs, sub: sub}, nil
}

// WatchOwnerSet is a free log subscription operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed relayManager, address indexed owner)
func (_IStakeManager *IStakeManagerFilterer) WatchOwnerSet(opts *bind.WatchOpts, sink chan<- *IStakeManagerOwnerSet, relayManager []common.Address, owner []common.Address) (event.Subscription, error) {

	var relayManagerRule []interface{}
	for _, relayManagerItem := range relayManager {
		relayManagerRule = append(relayManagerRule, relayManagerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _IStakeManager.contract.WatchLogs(opts, "OwnerSet", relayManagerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakeManagerOwnerSet)
				if err := _IStakeManager.contract.UnpackLog(event, "OwnerSet", log); err != nil {
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

// ParseOwnerSet is a log parse operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed relayManager, address indexed owner)
func (_IStakeManager *IStakeManagerFilterer) ParseOwnerSet(log types.Log) (*IStakeManagerOwnerSet, error) {
	event := new(IStakeManagerOwnerSet)
	if err := _IStakeManager.contract.UnpackLog(event, "OwnerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakeManagerStakeAddedIterator is returned from FilterStakeAdded and is used to iterate over the raw logs and unpacked data for StakeAdded events raised by the IStakeManager contract.
type IStakeManagerStakeAddedIterator struct {
	Event *IStakeManagerStakeAdded // Event containing the contract specifics and raw log

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
func (it *IStakeManagerStakeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakeManagerStakeAdded)
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
		it.Event = new(IStakeManagerStakeAdded)
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
func (it *IStakeManagerStakeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakeManagerStakeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakeManagerStakeAdded represents a StakeAdded event raised by the IStakeManager contract.
type IStakeManagerStakeAdded struct {
	RelayManager common.Address
	Owner        common.Address
	Stake        *big.Int
	UnstakeDelay *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStakeAdded is a free log retrieval operation binding the contract event 0xef7c8dfef14cbefdf829b8f066b068b677992411137321d64b3ed4538c2b3637.
//
// Solidity: event StakeAdded(address indexed relayManager, address indexed owner, uint256 stake, uint256 unstakeDelay)
func (_IStakeManager *IStakeManagerFilterer) FilterStakeAdded(opts *bind.FilterOpts, relayManager []common.Address, owner []common.Address) (*IStakeManagerStakeAddedIterator, error) {

	var relayManagerRule []interface{}
	for _, relayManagerItem := range relayManager {
		relayManagerRule = append(relayManagerRule, relayManagerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _IStakeManager.contract.FilterLogs(opts, "StakeAdded", relayManagerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &IStakeManagerStakeAddedIterator{contract: _IStakeManager.contract, event: "StakeAdded", logs: logs, sub: sub}, nil
}

// WatchStakeAdded is a free log subscription operation binding the contract event 0xef7c8dfef14cbefdf829b8f066b068b677992411137321d64b3ed4538c2b3637.
//
// Solidity: event StakeAdded(address indexed relayManager, address indexed owner, uint256 stake, uint256 unstakeDelay)
func (_IStakeManager *IStakeManagerFilterer) WatchStakeAdded(opts *bind.WatchOpts, sink chan<- *IStakeManagerStakeAdded, relayManager []common.Address, owner []common.Address) (event.Subscription, error) {

	var relayManagerRule []interface{}
	for _, relayManagerItem := range relayManager {
		relayManagerRule = append(relayManagerRule, relayManagerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _IStakeManager.contract.WatchLogs(opts, "StakeAdded", relayManagerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakeManagerStakeAdded)
				if err := _IStakeManager.contract.UnpackLog(event, "StakeAdded", log); err != nil {
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

// ParseStakeAdded is a log parse operation binding the contract event 0xef7c8dfef14cbefdf829b8f066b068b677992411137321d64b3ed4538c2b3637.
//
// Solidity: event StakeAdded(address indexed relayManager, address indexed owner, uint256 stake, uint256 unstakeDelay)
func (_IStakeManager *IStakeManagerFilterer) ParseStakeAdded(log types.Log) (*IStakeManagerStakeAdded, error) {
	event := new(IStakeManagerStakeAdded)
	if err := _IStakeManager.contract.UnpackLog(event, "StakeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakeManagerStakePenalizedIterator is returned from FilterStakePenalized and is used to iterate over the raw logs and unpacked data for StakePenalized events raised by the IStakeManager contract.
type IStakeManagerStakePenalizedIterator struct {
	Event *IStakeManagerStakePenalized // Event containing the contract specifics and raw log

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
func (it *IStakeManagerStakePenalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakeManagerStakePenalized)
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
		it.Event = new(IStakeManagerStakePenalized)
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
func (it *IStakeManagerStakePenalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakeManagerStakePenalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakeManagerStakePenalized represents a StakePenalized event raised by the IStakeManager contract.
type IStakeManagerStakePenalized struct {
	RelayManager common.Address
	Beneficiary  common.Address
	Reward       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStakePenalized is a free log retrieval operation binding the contract event 0x2f2ba0bf4c9bedc2210a4da5b5811c2a4fd28e62c51bb90c3ea6fdce00808eb0.
//
// Solidity: event StakePenalized(address indexed relayManager, address indexed beneficiary, uint256 reward)
func (_IStakeManager *IStakeManagerFilterer) FilterStakePenalized(opts *bind.FilterOpts, relayManager []common.Address, beneficiary []common.Address) (*IStakeManagerStakePenalizedIterator, error) {

	var relayManagerRule []interface{}
	for _, relayManagerItem := range relayManager {
		relayManagerRule = append(relayManagerRule, relayManagerItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _IStakeManager.contract.FilterLogs(opts, "StakePenalized", relayManagerRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &IStakeManagerStakePenalizedIterator{contract: _IStakeManager.contract, event: "StakePenalized", logs: logs, sub: sub}, nil
}

// WatchStakePenalized is a free log subscription operation binding the contract event 0x2f2ba0bf4c9bedc2210a4da5b5811c2a4fd28e62c51bb90c3ea6fdce00808eb0.
//
// Solidity: event StakePenalized(address indexed relayManager, address indexed beneficiary, uint256 reward)
func (_IStakeManager *IStakeManagerFilterer) WatchStakePenalized(opts *bind.WatchOpts, sink chan<- *IStakeManagerStakePenalized, relayManager []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var relayManagerRule []interface{}
	for _, relayManagerItem := range relayManager {
		relayManagerRule = append(relayManagerRule, relayManagerItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _IStakeManager.contract.WatchLogs(opts, "StakePenalized", relayManagerRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakeManagerStakePenalized)
				if err := _IStakeManager.contract.UnpackLog(event, "StakePenalized", log); err != nil {
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

// ParseStakePenalized is a log parse operation binding the contract event 0x2f2ba0bf4c9bedc2210a4da5b5811c2a4fd28e62c51bb90c3ea6fdce00808eb0.
//
// Solidity: event StakePenalized(address indexed relayManager, address indexed beneficiary, uint256 reward)
func (_IStakeManager *IStakeManagerFilterer) ParseStakePenalized(log types.Log) (*IStakeManagerStakePenalized, error) {
	event := new(IStakeManagerStakePenalized)
	if err := _IStakeManager.contract.UnpackLog(event, "StakePenalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakeManagerStakeUnlockedIterator is returned from FilterStakeUnlocked and is used to iterate over the raw logs and unpacked data for StakeUnlocked events raised by the IStakeManager contract.
type IStakeManagerStakeUnlockedIterator struct {
	Event *IStakeManagerStakeUnlocked // Event containing the contract specifics and raw log

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
func (it *IStakeManagerStakeUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakeManagerStakeUnlocked)
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
		it.Event = new(IStakeManagerStakeUnlocked)
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
func (it *IStakeManagerStakeUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakeManagerStakeUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakeManagerStakeUnlocked represents a StakeUnlocked event raised by the IStakeManager contract.
type IStakeManagerStakeUnlocked struct {
	RelayManager  common.Address
	Owner         common.Address
	WithdrawBlock *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterStakeUnlocked is a free log retrieval operation binding the contract event 0x9ffc6168de1eb7f1d16200f614753cd7edce5a2186aab1c612199dd7316cd7c4.
//
// Solidity: event StakeUnlocked(address indexed relayManager, address indexed owner, uint256 withdrawBlock)
func (_IStakeManager *IStakeManagerFilterer) FilterStakeUnlocked(opts *bind.FilterOpts, relayManager []common.Address, owner []common.Address) (*IStakeManagerStakeUnlockedIterator, error) {

	var relayManagerRule []interface{}
	for _, relayManagerItem := range relayManager {
		relayManagerRule = append(relayManagerRule, relayManagerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _IStakeManager.contract.FilterLogs(opts, "StakeUnlocked", relayManagerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &IStakeManagerStakeUnlockedIterator{contract: _IStakeManager.contract, event: "StakeUnlocked", logs: logs, sub: sub}, nil
}

// WatchStakeUnlocked is a free log subscription operation binding the contract event 0x9ffc6168de1eb7f1d16200f614753cd7edce5a2186aab1c612199dd7316cd7c4.
//
// Solidity: event StakeUnlocked(address indexed relayManager, address indexed owner, uint256 withdrawBlock)
func (_IStakeManager *IStakeManagerFilterer) WatchStakeUnlocked(opts *bind.WatchOpts, sink chan<- *IStakeManagerStakeUnlocked, relayManager []common.Address, owner []common.Address) (event.Subscription, error) {

	var relayManagerRule []interface{}
	for _, relayManagerItem := range relayManager {
		relayManagerRule = append(relayManagerRule, relayManagerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _IStakeManager.contract.WatchLogs(opts, "StakeUnlocked", relayManagerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakeManagerStakeUnlocked)
				if err := _IStakeManager.contract.UnpackLog(event, "StakeUnlocked", log); err != nil {
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

// ParseStakeUnlocked is a log parse operation binding the contract event 0x9ffc6168de1eb7f1d16200f614753cd7edce5a2186aab1c612199dd7316cd7c4.
//
// Solidity: event StakeUnlocked(address indexed relayManager, address indexed owner, uint256 withdrawBlock)
func (_IStakeManager *IStakeManagerFilterer) ParseStakeUnlocked(log types.Log) (*IStakeManagerStakeUnlocked, error) {
	event := new(IStakeManagerStakeUnlocked)
	if err := _IStakeManager.contract.UnpackLog(event, "StakeUnlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakeManagerStakeWithdrawnIterator is returned from FilterStakeWithdrawn and is used to iterate over the raw logs and unpacked data for StakeWithdrawn events raised by the IStakeManager contract.
type IStakeManagerStakeWithdrawnIterator struct {
	Event *IStakeManagerStakeWithdrawn // Event containing the contract specifics and raw log

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
func (it *IStakeManagerStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakeManagerStakeWithdrawn)
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
		it.Event = new(IStakeManagerStakeWithdrawn)
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
func (it *IStakeManagerStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakeManagerStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakeManagerStakeWithdrawn represents a StakeWithdrawn event raised by the IStakeManager contract.
type IStakeManagerStakeWithdrawn struct {
	RelayManager common.Address
	Owner        common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawn is a free log retrieval operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed relayManager, address indexed owner, uint256 amount)
func (_IStakeManager *IStakeManagerFilterer) FilterStakeWithdrawn(opts *bind.FilterOpts, relayManager []common.Address, owner []common.Address) (*IStakeManagerStakeWithdrawnIterator, error) {

	var relayManagerRule []interface{}
	for _, relayManagerItem := range relayManager {
		relayManagerRule = append(relayManagerRule, relayManagerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _IStakeManager.contract.FilterLogs(opts, "StakeWithdrawn", relayManagerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &IStakeManagerStakeWithdrawnIterator{contract: _IStakeManager.contract, event: "StakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawn is a free log subscription operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed relayManager, address indexed owner, uint256 amount)
func (_IStakeManager *IStakeManagerFilterer) WatchStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *IStakeManagerStakeWithdrawn, relayManager []common.Address, owner []common.Address) (event.Subscription, error) {

	var relayManagerRule []interface{}
	for _, relayManagerItem := range relayManager {
		relayManagerRule = append(relayManagerRule, relayManagerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _IStakeManager.contract.WatchLogs(opts, "StakeWithdrawn", relayManagerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakeManagerStakeWithdrawn)
				if err := _IStakeManager.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
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

// ParseStakeWithdrawn is a log parse operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed relayManager, address indexed owner, uint256 amount)
func (_IStakeManager *IStakeManagerFilterer) ParseStakeWithdrawn(log types.Log) (*IStakeManagerStakeWithdrawn, error) {
	event := new(IStakeManagerStakeWithdrawn)
	if err := _IStakeManager.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
