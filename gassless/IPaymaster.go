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

// GsnTypesRelayData is an auto generated low-level Go binding around an user-defined struct.
type GsnTypesRelayData struct {
	GasPrice      *big.Int
	PctRelayFee   *big.Int
	BaseRelayFee  *big.Int
	RelayWorker   common.Address
	Paymaster     common.Address
	Forwarder     common.Address
	PaymasterData []byte
	ClientId      *big.Int
}

// GsnTypesRelayRequest is an auto generated low-level Go binding around an user-defined struct.
type GsnTypesRelayRequest struct {
	Request   IForwarderForwardRequest
	RelayData GsnTypesRelayData
}

// IForwarderForwardRequest is an auto generated low-level Go binding around an user-defined struct.
type IForwarderForwardRequest struct {
	From       common.Address
	To         common.Address
	Value      *big.Int
	Gas        *big.Int
	Nonce      *big.Int
	Data       []byte
	ValidUntil *big.Int
}

// IPaymasterGasAndDataLimits is an auto generated low-level Go binding around an user-defined struct.
type IPaymasterGasAndDataLimits struct {
	AcceptanceBudget        *big.Int
	PreRelayedCallGasLimit  *big.Int
	PostRelayedCallGasLimit *big.Int
	CalldataSizeLimit       *big.Int
}

// IPaymasterABI is the input ABI used to generate the binding from.
const IPaymasterABI = "[{\"inputs\":[],\"name\":\"getGasAndDataLimits\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"acceptanceBudget\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preRelayedCallGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"postRelayedCallGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldataSizeLimit\",\"type\":\"uint256\"}],\"internalType\":\"structIPaymaster.GasAndDataLimits\",\"name\":\"limits\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getHubAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRelayHubDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"gasUseWithoutPost\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pctRelayFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRelayFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayWorker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"paymasterData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"clientId\",\"type\":\"uint256\"}],\"internalType\":\"structGsnTypes.RelayData\",\"name\":\"relayData\",\"type\":\"tuple\"}],\"name\":\"postRelayedCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validUntil\",\"type\":\"uint256\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"request\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pctRelayFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRelayFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayWorker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"paymasterData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"clientId\",\"type\":\"uint256\"}],\"internalType\":\"structGsnTypes.RelayData\",\"name\":\"relayData\",\"type\":\"tuple\"}],\"internalType\":\"structGsnTypes.RelayRequest\",\"name\":\"relayRequest\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approvalData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"maxPossibleGas\",\"type\":\"uint256\"}],\"name\":\"preRelayedCall\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"rejectOnRecipientRevert\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedForwarder\",\"outputs\":[{\"internalType\":\"contractIForwarder\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionPaymaster\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IPaymaster is an auto generated Go binding around an Ethereum contract.
type IPaymaster struct {
	IPaymasterCaller     // Read-only binding to the contract
	IPaymasterTransactor // Write-only binding to the contract
	IPaymasterFilterer   // Log filterer for contract events
}

// IPaymasterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPaymasterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPaymasterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPaymasterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPaymasterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPaymasterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPaymasterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPaymasterSession struct {
	Contract     *IPaymaster       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPaymasterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPaymasterCallerSession struct {
	Contract *IPaymasterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IPaymasterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPaymasterTransactorSession struct {
	Contract     *IPaymasterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IPaymasterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPaymasterRaw struct {
	Contract *IPaymaster // Generic contract binding to access the raw methods on
}

// IPaymasterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPaymasterCallerRaw struct {
	Contract *IPaymasterCaller // Generic read-only contract binding to access the raw methods on
}

// IPaymasterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPaymasterTransactorRaw struct {
	Contract *IPaymasterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPaymaster creates a new instance of IPaymaster, bound to a specific deployed contract.
func NewIPaymaster(address common.Address, backend bind.ContractBackend) (*IPaymaster, error) {
	contract, err := bindIPaymaster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPaymaster{IPaymasterCaller: IPaymasterCaller{contract: contract}, IPaymasterTransactor: IPaymasterTransactor{contract: contract}, IPaymasterFilterer: IPaymasterFilterer{contract: contract}}, nil
}

// NewIPaymasterCaller creates a new read-only instance of IPaymaster, bound to a specific deployed contract.
func NewIPaymasterCaller(address common.Address, caller bind.ContractCaller) (*IPaymasterCaller, error) {
	contract, err := bindIPaymaster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPaymasterCaller{contract: contract}, nil
}

// NewIPaymasterTransactor creates a new write-only instance of IPaymaster, bound to a specific deployed contract.
func NewIPaymasterTransactor(address common.Address, transactor bind.ContractTransactor) (*IPaymasterTransactor, error) {
	contract, err := bindIPaymaster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPaymasterTransactor{contract: contract}, nil
}

// NewIPaymasterFilterer creates a new log filterer instance of IPaymaster, bound to a specific deployed contract.
func NewIPaymasterFilterer(address common.Address, filterer bind.ContractFilterer) (*IPaymasterFilterer, error) {
	contract, err := bindIPaymaster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPaymasterFilterer{contract: contract}, nil
}

// bindIPaymaster binds a generic wrapper to an already deployed contract.
func bindIPaymaster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPaymasterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPaymaster *IPaymasterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPaymaster.Contract.IPaymasterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPaymaster *IPaymasterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPaymaster.Contract.IPaymasterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPaymaster *IPaymasterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPaymaster.Contract.IPaymasterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPaymaster *IPaymasterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPaymaster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPaymaster *IPaymasterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPaymaster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPaymaster *IPaymasterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPaymaster.Contract.contract.Transact(opts, method, params...)
}

// GetGasAndDataLimits is a free data retrieval call binding the contract method 0xb039a88f.
//
// Solidity: function getGasAndDataLimits() view returns((uint256,uint256,uint256,uint256) limits)
func (_IPaymaster *IPaymasterCaller) GetGasAndDataLimits(opts *bind.CallOpts) (IPaymasterGasAndDataLimits, error) {
	var out []interface{}
	err := _IPaymaster.contract.Call(opts, &out, "getGasAndDataLimits")

	if err != nil {
		return *new(IPaymasterGasAndDataLimits), err
	}

	out0 := *abi.ConvertType(out[0], new(IPaymasterGasAndDataLimits)).(*IPaymasterGasAndDataLimits)

	return out0, err

}

// GetGasAndDataLimits is a free data retrieval call binding the contract method 0xb039a88f.
//
// Solidity: function getGasAndDataLimits() view returns((uint256,uint256,uint256,uint256) limits)
func (_IPaymaster *IPaymasterSession) GetGasAndDataLimits() (IPaymasterGasAndDataLimits, error) {
	return _IPaymaster.Contract.GetGasAndDataLimits(&_IPaymaster.CallOpts)
}

// GetGasAndDataLimits is a free data retrieval call binding the contract method 0xb039a88f.
//
// Solidity: function getGasAndDataLimits() view returns((uint256,uint256,uint256,uint256) limits)
func (_IPaymaster *IPaymasterCallerSession) GetGasAndDataLimits() (IPaymasterGasAndDataLimits, error) {
	return _IPaymaster.Contract.GetGasAndDataLimits(&_IPaymaster.CallOpts)
}

// GetHubAddr is a free data retrieval call binding the contract method 0x74e861d6.
//
// Solidity: function getHubAddr() view returns(address)
func (_IPaymaster *IPaymasterCaller) GetHubAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPaymaster.contract.Call(opts, &out, "getHubAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetHubAddr is a free data retrieval call binding the contract method 0x74e861d6.
//
// Solidity: function getHubAddr() view returns(address)
func (_IPaymaster *IPaymasterSession) GetHubAddr() (common.Address, error) {
	return _IPaymaster.Contract.GetHubAddr(&_IPaymaster.CallOpts)
}

// GetHubAddr is a free data retrieval call binding the contract method 0x74e861d6.
//
// Solidity: function getHubAddr() view returns(address)
func (_IPaymaster *IPaymasterCallerSession) GetHubAddr() (common.Address, error) {
	return _IPaymaster.Contract.GetHubAddr(&_IPaymaster.CallOpts)
}

// GetRelayHubDeposit is a free data retrieval call binding the contract method 0x2afe31c1.
//
// Solidity: function getRelayHubDeposit() view returns(uint256)
func (_IPaymaster *IPaymasterCaller) GetRelayHubDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPaymaster.contract.Call(opts, &out, "getRelayHubDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRelayHubDeposit is a free data retrieval call binding the contract method 0x2afe31c1.
//
// Solidity: function getRelayHubDeposit() view returns(uint256)
func (_IPaymaster *IPaymasterSession) GetRelayHubDeposit() (*big.Int, error) {
	return _IPaymaster.Contract.GetRelayHubDeposit(&_IPaymaster.CallOpts)
}

// GetRelayHubDeposit is a free data retrieval call binding the contract method 0x2afe31c1.
//
// Solidity: function getRelayHubDeposit() view returns(uint256)
func (_IPaymaster *IPaymasterCallerSession) GetRelayHubDeposit() (*big.Int, error) {
	return _IPaymaster.Contract.GetRelayHubDeposit(&_IPaymaster.CallOpts)
}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_IPaymaster *IPaymasterCaller) TrustedForwarder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPaymaster.contract.Call(opts, &out, "trustedForwarder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_IPaymaster *IPaymasterSession) TrustedForwarder() (common.Address, error) {
	return _IPaymaster.Contract.TrustedForwarder(&_IPaymaster.CallOpts)
}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_IPaymaster *IPaymasterCallerSession) TrustedForwarder() (common.Address, error) {
	return _IPaymaster.Contract.TrustedForwarder(&_IPaymaster.CallOpts)
}

// VersionPaymaster is a free data retrieval call binding the contract method 0x921276ea.
//
// Solidity: function versionPaymaster() view returns(string)
func (_IPaymaster *IPaymasterCaller) VersionPaymaster(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IPaymaster.contract.Call(opts, &out, "versionPaymaster")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VersionPaymaster is a free data retrieval call binding the contract method 0x921276ea.
//
// Solidity: function versionPaymaster() view returns(string)
func (_IPaymaster *IPaymasterSession) VersionPaymaster() (string, error) {
	return _IPaymaster.Contract.VersionPaymaster(&_IPaymaster.CallOpts)
}

// VersionPaymaster is a free data retrieval call binding the contract method 0x921276ea.
//
// Solidity: function versionPaymaster() view returns(string)
func (_IPaymaster *IPaymasterCallerSession) VersionPaymaster() (string, error) {
	return _IPaymaster.Contract.VersionPaymaster(&_IPaymaster.CallOpts)
}

// PostRelayedCall is a paid mutator transaction binding the contract method 0x76fa01c3.
//
// Solidity: function postRelayedCall(bytes context, bool success, uint256 gasUseWithoutPost, (uint256,uint256,uint256,address,address,address,bytes,uint256) relayData) returns()
func (_IPaymaster *IPaymasterTransactor) PostRelayedCall(opts *bind.TransactOpts, context []byte, success bool, gasUseWithoutPost *big.Int, relayData GsnTypesRelayData) (*types.Transaction, error) {
	return _IPaymaster.contract.Transact(opts, "postRelayedCall", context, success, gasUseWithoutPost, relayData)
}

// PostRelayedCall is a paid mutator transaction binding the contract method 0x76fa01c3.
//
// Solidity: function postRelayedCall(bytes context, bool success, uint256 gasUseWithoutPost, (uint256,uint256,uint256,address,address,address,bytes,uint256) relayData) returns()
func (_IPaymaster *IPaymasterSession) PostRelayedCall(context []byte, success bool, gasUseWithoutPost *big.Int, relayData GsnTypesRelayData) (*types.Transaction, error) {
	return _IPaymaster.Contract.PostRelayedCall(&_IPaymaster.TransactOpts, context, success, gasUseWithoutPost, relayData)
}

// PostRelayedCall is a paid mutator transaction binding the contract method 0x76fa01c3.
//
// Solidity: function postRelayedCall(bytes context, bool success, uint256 gasUseWithoutPost, (uint256,uint256,uint256,address,address,address,bytes,uint256) relayData) returns()
func (_IPaymaster *IPaymasterTransactorSession) PostRelayedCall(context []byte, success bool, gasUseWithoutPost *big.Int, relayData GsnTypesRelayData) (*types.Transaction, error) {
	return _IPaymaster.Contract.PostRelayedCall(&_IPaymaster.TransactOpts, context, success, gasUseWithoutPost, relayData)
}

// PreRelayedCall is a paid mutator transaction binding the contract method 0x00be5dd4.
//
// Solidity: function preRelayedCall(((address,address,uint256,uint256,uint256,bytes,uint256),(uint256,uint256,uint256,address,address,address,bytes,uint256)) relayRequest, bytes signature, bytes approvalData, uint256 maxPossibleGas) returns(bytes context, bool rejectOnRecipientRevert)
func (_IPaymaster *IPaymasterTransactor) PreRelayedCall(opts *bind.TransactOpts, relayRequest GsnTypesRelayRequest, signature []byte, approvalData []byte, maxPossibleGas *big.Int) (*types.Transaction, error) {
	return _IPaymaster.contract.Transact(opts, "preRelayedCall", relayRequest, signature, approvalData, maxPossibleGas)
}

// PreRelayedCall is a paid mutator transaction binding the contract method 0x00be5dd4.
//
// Solidity: function preRelayedCall(((address,address,uint256,uint256,uint256,bytes,uint256),(uint256,uint256,uint256,address,address,address,bytes,uint256)) relayRequest, bytes signature, bytes approvalData, uint256 maxPossibleGas) returns(bytes context, bool rejectOnRecipientRevert)
func (_IPaymaster *IPaymasterSession) PreRelayedCall(relayRequest GsnTypesRelayRequest, signature []byte, approvalData []byte, maxPossibleGas *big.Int) (*types.Transaction, error) {
	return _IPaymaster.Contract.PreRelayedCall(&_IPaymaster.TransactOpts, relayRequest, signature, approvalData, maxPossibleGas)
}

// PreRelayedCall is a paid mutator transaction binding the contract method 0x00be5dd4.
//
// Solidity: function preRelayedCall(((address,address,uint256,uint256,uint256,bytes,uint256),(uint256,uint256,uint256,address,address,address,bytes,uint256)) relayRequest, bytes signature, bytes approvalData, uint256 maxPossibleGas) returns(bytes context, bool rejectOnRecipientRevert)
func (_IPaymaster *IPaymasterTransactorSession) PreRelayedCall(relayRequest GsnTypesRelayRequest, signature []byte, approvalData []byte, maxPossibleGas *big.Int) (*types.Transaction, error) {
	return _IPaymaster.Contract.PreRelayedCall(&_IPaymaster.TransactOpts, relayRequest, signature, approvalData, maxPossibleGas)
}
