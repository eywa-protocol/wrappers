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

// BasePaymasterABI is the input ABI used to generate the binding from.
const BasePaymasterABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CALLDATA_SIZE_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FORWARDER_HUB_OVERHEAD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAYMASTER_ACCEPTANCE_BUDGET\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POST_RELAYED_CALL_GAS_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRE_RELAYED_CALL_GAS_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validUntil\",\"type\":\"uint256\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"request\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pctRelayFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRelayFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayWorker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"paymasterData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"clientId\",\"type\":\"uint256\"}],\"internalType\":\"structGsnTypes.RelayData\",\"name\":\"relayData\",\"type\":\"tuple\"}],\"internalType\":\"structGsnTypes.RelayRequest\",\"name\":\"relayRequest\",\"type\":\"tuple\"}],\"name\":\"_verifyForwarder\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGasAndDataLimits\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"acceptanceBudget\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preRelayedCallGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"postRelayedCallGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldataSizeLimit\",\"type\":\"uint256\"}],\"internalType\":\"structIPaymaster.GasAndDataLimits\",\"name\":\"limits\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getHubAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRelayHubDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"gasUseWithoutPost\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pctRelayFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRelayFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayWorker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"paymasterData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"clientId\",\"type\":\"uint256\"}],\"internalType\":\"structGsnTypes.RelayData\",\"name\":\"relayData\",\"type\":\"tuple\"}],\"name\":\"postRelayedCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validUntil\",\"type\":\"uint256\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"request\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pctRelayFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRelayFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayWorker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"paymasterData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"clientId\",\"type\":\"uint256\"}],\"internalType\":\"structGsnTypes.RelayData\",\"name\":\"relayData\",\"type\":\"tuple\"}],\"internalType\":\"structGsnTypes.RelayRequest\",\"name\":\"relayRequest\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approvalData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"maxPossibleGas\",\"type\":\"uint256\"}],\"name\":\"preRelayedCall\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"rejectOnRecipientRevert\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRelayHub\",\"name\":\"hub\",\"type\":\"address\"}],\"name\":\"setRelayHub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIForwarder\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"setTrustedForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedForwarder\",\"outputs\":[{\"internalType\":\"contractIForwarder\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionPaymaster\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"withdrawRelayHubDepositTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// BasePaymaster is an auto generated Go binding around an Ethereum contract.
type BasePaymaster struct {
	BasePaymasterCaller     // Read-only binding to the contract
	BasePaymasterTransactor // Write-only binding to the contract
	BasePaymasterFilterer   // Log filterer for contract events
}

// BasePaymasterCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasePaymasterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasePaymasterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasePaymasterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasePaymasterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasePaymasterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasePaymasterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasePaymasterSession struct {
	Contract     *BasePaymaster    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasePaymasterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasePaymasterCallerSession struct {
	Contract *BasePaymasterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BasePaymasterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasePaymasterTransactorSession struct {
	Contract     *BasePaymasterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BasePaymasterRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasePaymasterRaw struct {
	Contract *BasePaymaster // Generic contract binding to access the raw methods on
}

// BasePaymasterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasePaymasterCallerRaw struct {
	Contract *BasePaymasterCaller // Generic read-only contract binding to access the raw methods on
}

// BasePaymasterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasePaymasterTransactorRaw struct {
	Contract *BasePaymasterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasePaymaster creates a new instance of BasePaymaster, bound to a specific deployed contract.
func NewBasePaymaster(address common.Address, backend bind.ContractBackend) (*BasePaymaster, error) {
	contract, err := bindBasePaymaster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasePaymaster{BasePaymasterCaller: BasePaymasterCaller{contract: contract}, BasePaymasterTransactor: BasePaymasterTransactor{contract: contract}, BasePaymasterFilterer: BasePaymasterFilterer{contract: contract}}, nil
}

// NewBasePaymasterCaller creates a new read-only instance of BasePaymaster, bound to a specific deployed contract.
func NewBasePaymasterCaller(address common.Address, caller bind.ContractCaller) (*BasePaymasterCaller, error) {
	contract, err := bindBasePaymaster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasePaymasterCaller{contract: contract}, nil
}

// NewBasePaymasterTransactor creates a new write-only instance of BasePaymaster, bound to a specific deployed contract.
func NewBasePaymasterTransactor(address common.Address, transactor bind.ContractTransactor) (*BasePaymasterTransactor, error) {
	contract, err := bindBasePaymaster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasePaymasterTransactor{contract: contract}, nil
}

// NewBasePaymasterFilterer creates a new log filterer instance of BasePaymaster, bound to a specific deployed contract.
func NewBasePaymasterFilterer(address common.Address, filterer bind.ContractFilterer) (*BasePaymasterFilterer, error) {
	contract, err := bindBasePaymaster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasePaymasterFilterer{contract: contract}, nil
}

// bindBasePaymaster binds a generic wrapper to an already deployed contract.
func bindBasePaymaster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasePaymasterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasePaymaster *BasePaymasterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasePaymaster.Contract.BasePaymasterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasePaymaster *BasePaymasterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasePaymaster.Contract.BasePaymasterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasePaymaster *BasePaymasterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasePaymaster.Contract.BasePaymasterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasePaymaster *BasePaymasterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasePaymaster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasePaymaster *BasePaymasterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasePaymaster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasePaymaster *BasePaymasterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasePaymaster.Contract.contract.Transact(opts, method, params...)
}

// CALLDATASIZELIMIT is a free data retrieval call binding the contract method 0x5c5e3db1.
//
// Solidity: function CALLDATA_SIZE_LIMIT() view returns(uint256)
func (_BasePaymaster *BasePaymasterCaller) CALLDATASIZELIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasePaymaster.contract.Call(opts, &out, "CALLDATA_SIZE_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CALLDATASIZELIMIT is a free data retrieval call binding the contract method 0x5c5e3db1.
//
// Solidity: function CALLDATA_SIZE_LIMIT() view returns(uint256)
func (_BasePaymaster *BasePaymasterSession) CALLDATASIZELIMIT() (*big.Int, error) {
	return _BasePaymaster.Contract.CALLDATASIZELIMIT(&_BasePaymaster.CallOpts)
}

// CALLDATASIZELIMIT is a free data retrieval call binding the contract method 0x5c5e3db1.
//
// Solidity: function CALLDATA_SIZE_LIMIT() view returns(uint256)
func (_BasePaymaster *BasePaymasterCallerSession) CALLDATASIZELIMIT() (*big.Int, error) {
	return _BasePaymaster.Contract.CALLDATASIZELIMIT(&_BasePaymaster.CallOpts)
}

// FORWARDERHUBOVERHEAD is a free data retrieval call binding the contract method 0xb90b41cf.
//
// Solidity: function FORWARDER_HUB_OVERHEAD() view returns(uint256)
func (_BasePaymaster *BasePaymasterCaller) FORWARDERHUBOVERHEAD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasePaymaster.contract.Call(opts, &out, "FORWARDER_HUB_OVERHEAD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FORWARDERHUBOVERHEAD is a free data retrieval call binding the contract method 0xb90b41cf.
//
// Solidity: function FORWARDER_HUB_OVERHEAD() view returns(uint256)
func (_BasePaymaster *BasePaymasterSession) FORWARDERHUBOVERHEAD() (*big.Int, error) {
	return _BasePaymaster.Contract.FORWARDERHUBOVERHEAD(&_BasePaymaster.CallOpts)
}

// FORWARDERHUBOVERHEAD is a free data retrieval call binding the contract method 0xb90b41cf.
//
// Solidity: function FORWARDER_HUB_OVERHEAD() view returns(uint256)
func (_BasePaymaster *BasePaymasterCallerSession) FORWARDERHUBOVERHEAD() (*big.Int, error) {
	return _BasePaymaster.Contract.FORWARDERHUBOVERHEAD(&_BasePaymaster.CallOpts)
}

// PAYMASTERACCEPTANCEBUDGET is a free data retrieval call binding the contract method 0xdf463a66.
//
// Solidity: function PAYMASTER_ACCEPTANCE_BUDGET() view returns(uint256)
func (_BasePaymaster *BasePaymasterCaller) PAYMASTERACCEPTANCEBUDGET(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasePaymaster.contract.Call(opts, &out, "PAYMASTER_ACCEPTANCE_BUDGET")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PAYMASTERACCEPTANCEBUDGET is a free data retrieval call binding the contract method 0xdf463a66.
//
// Solidity: function PAYMASTER_ACCEPTANCE_BUDGET() view returns(uint256)
func (_BasePaymaster *BasePaymasterSession) PAYMASTERACCEPTANCEBUDGET() (*big.Int, error) {
	return _BasePaymaster.Contract.PAYMASTERACCEPTANCEBUDGET(&_BasePaymaster.CallOpts)
}

// PAYMASTERACCEPTANCEBUDGET is a free data retrieval call binding the contract method 0xdf463a66.
//
// Solidity: function PAYMASTER_ACCEPTANCE_BUDGET() view returns(uint256)
func (_BasePaymaster *BasePaymasterCallerSession) PAYMASTERACCEPTANCEBUDGET() (*big.Int, error) {
	return _BasePaymaster.Contract.PAYMASTERACCEPTANCEBUDGET(&_BasePaymaster.CallOpts)
}

// POSTRELAYEDCALLGASLIMIT is a free data retrieval call binding the contract method 0xbbdaa3c9.
//
// Solidity: function POST_RELAYED_CALL_GAS_LIMIT() view returns(uint256)
func (_BasePaymaster *BasePaymasterCaller) POSTRELAYEDCALLGASLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasePaymaster.contract.Call(opts, &out, "POST_RELAYED_CALL_GAS_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// POSTRELAYEDCALLGASLIMIT is a free data retrieval call binding the contract method 0xbbdaa3c9.
//
// Solidity: function POST_RELAYED_CALL_GAS_LIMIT() view returns(uint256)
func (_BasePaymaster *BasePaymasterSession) POSTRELAYEDCALLGASLIMIT() (*big.Int, error) {
	return _BasePaymaster.Contract.POSTRELAYEDCALLGASLIMIT(&_BasePaymaster.CallOpts)
}

// POSTRELAYEDCALLGASLIMIT is a free data retrieval call binding the contract method 0xbbdaa3c9.
//
// Solidity: function POST_RELAYED_CALL_GAS_LIMIT() view returns(uint256)
func (_BasePaymaster *BasePaymasterCallerSession) POSTRELAYEDCALLGASLIMIT() (*big.Int, error) {
	return _BasePaymaster.Contract.POSTRELAYEDCALLGASLIMIT(&_BasePaymaster.CallOpts)
}

// PRERELAYEDCALLGASLIMIT is a free data retrieval call binding the contract method 0xf9c002f7.
//
// Solidity: function PRE_RELAYED_CALL_GAS_LIMIT() view returns(uint256)
func (_BasePaymaster *BasePaymasterCaller) PRERELAYEDCALLGASLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasePaymaster.contract.Call(opts, &out, "PRE_RELAYED_CALL_GAS_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRERELAYEDCALLGASLIMIT is a free data retrieval call binding the contract method 0xf9c002f7.
//
// Solidity: function PRE_RELAYED_CALL_GAS_LIMIT() view returns(uint256)
func (_BasePaymaster *BasePaymasterSession) PRERELAYEDCALLGASLIMIT() (*big.Int, error) {
	return _BasePaymaster.Contract.PRERELAYEDCALLGASLIMIT(&_BasePaymaster.CallOpts)
}

// PRERELAYEDCALLGASLIMIT is a free data retrieval call binding the contract method 0xf9c002f7.
//
// Solidity: function PRE_RELAYED_CALL_GAS_LIMIT() view returns(uint256)
func (_BasePaymaster *BasePaymasterCallerSession) PRERELAYEDCALLGASLIMIT() (*big.Int, error) {
	return _BasePaymaster.Contract.PRERELAYEDCALLGASLIMIT(&_BasePaymaster.CallOpts)
}

// VerifyForwarder is a free data retrieval call binding the contract method 0xa5dcd07b.
//
// Solidity: function _verifyForwarder(((address,address,uint256,uint256,uint256,bytes,uint256),(uint256,uint256,uint256,address,address,address,bytes,uint256)) relayRequest) view returns()
func (_BasePaymaster *BasePaymasterCaller) VerifyForwarder(opts *bind.CallOpts, relayRequest GsnTypesRelayRequest) error {
	var out []interface{}
	err := _BasePaymaster.contract.Call(opts, &out, "_verifyForwarder", relayRequest)

	if err != nil {
		return err
	}

	return err

}

// VerifyForwarder is a free data retrieval call binding the contract method 0xa5dcd07b.
//
// Solidity: function _verifyForwarder(((address,address,uint256,uint256,uint256,bytes,uint256),(uint256,uint256,uint256,address,address,address,bytes,uint256)) relayRequest) view returns()
func (_BasePaymaster *BasePaymasterSession) VerifyForwarder(relayRequest GsnTypesRelayRequest) error {
	return _BasePaymaster.Contract.VerifyForwarder(&_BasePaymaster.CallOpts, relayRequest)
}

// VerifyForwarder is a free data retrieval call binding the contract method 0xa5dcd07b.
//
// Solidity: function _verifyForwarder(((address,address,uint256,uint256,uint256,bytes,uint256),(uint256,uint256,uint256,address,address,address,bytes,uint256)) relayRequest) view returns()
func (_BasePaymaster *BasePaymasterCallerSession) VerifyForwarder(relayRequest GsnTypesRelayRequest) error {
	return _BasePaymaster.Contract.VerifyForwarder(&_BasePaymaster.CallOpts, relayRequest)
}

// GetGasAndDataLimits is a free data retrieval call binding the contract method 0xb039a88f.
//
// Solidity: function getGasAndDataLimits() view returns((uint256,uint256,uint256,uint256) limits)
func (_BasePaymaster *BasePaymasterCaller) GetGasAndDataLimits(opts *bind.CallOpts) (IPaymasterGasAndDataLimits, error) {
	var out []interface{}
	err := _BasePaymaster.contract.Call(opts, &out, "getGasAndDataLimits")

	if err != nil {
		return *new(IPaymasterGasAndDataLimits), err
	}

	out0 := *abi.ConvertType(out[0], new(IPaymasterGasAndDataLimits)).(*IPaymasterGasAndDataLimits)

	return out0, err

}

// GetGasAndDataLimits is a free data retrieval call binding the contract method 0xb039a88f.
//
// Solidity: function getGasAndDataLimits() view returns((uint256,uint256,uint256,uint256) limits)
func (_BasePaymaster *BasePaymasterSession) GetGasAndDataLimits() (IPaymasterGasAndDataLimits, error) {
	return _BasePaymaster.Contract.GetGasAndDataLimits(&_BasePaymaster.CallOpts)
}

// GetGasAndDataLimits is a free data retrieval call binding the contract method 0xb039a88f.
//
// Solidity: function getGasAndDataLimits() view returns((uint256,uint256,uint256,uint256) limits)
func (_BasePaymaster *BasePaymasterCallerSession) GetGasAndDataLimits() (IPaymasterGasAndDataLimits, error) {
	return _BasePaymaster.Contract.GetGasAndDataLimits(&_BasePaymaster.CallOpts)
}

// GetHubAddr is a free data retrieval call binding the contract method 0x74e861d6.
//
// Solidity: function getHubAddr() view returns(address)
func (_BasePaymaster *BasePaymasterCaller) GetHubAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasePaymaster.contract.Call(opts, &out, "getHubAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetHubAddr is a free data retrieval call binding the contract method 0x74e861d6.
//
// Solidity: function getHubAddr() view returns(address)
func (_BasePaymaster *BasePaymasterSession) GetHubAddr() (common.Address, error) {
	return _BasePaymaster.Contract.GetHubAddr(&_BasePaymaster.CallOpts)
}

// GetHubAddr is a free data retrieval call binding the contract method 0x74e861d6.
//
// Solidity: function getHubAddr() view returns(address)
func (_BasePaymaster *BasePaymasterCallerSession) GetHubAddr() (common.Address, error) {
	return _BasePaymaster.Contract.GetHubAddr(&_BasePaymaster.CallOpts)
}

// GetRelayHubDeposit is a free data retrieval call binding the contract method 0x2afe31c1.
//
// Solidity: function getRelayHubDeposit() view returns(uint256)
func (_BasePaymaster *BasePaymasterCaller) GetRelayHubDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasePaymaster.contract.Call(opts, &out, "getRelayHubDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRelayHubDeposit is a free data retrieval call binding the contract method 0x2afe31c1.
//
// Solidity: function getRelayHubDeposit() view returns(uint256)
func (_BasePaymaster *BasePaymasterSession) GetRelayHubDeposit() (*big.Int, error) {
	return _BasePaymaster.Contract.GetRelayHubDeposit(&_BasePaymaster.CallOpts)
}

// GetRelayHubDeposit is a free data retrieval call binding the contract method 0x2afe31c1.
//
// Solidity: function getRelayHubDeposit() view returns(uint256)
func (_BasePaymaster *BasePaymasterCallerSession) GetRelayHubDeposit() (*big.Int, error) {
	return _BasePaymaster.Contract.GetRelayHubDeposit(&_BasePaymaster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BasePaymaster *BasePaymasterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasePaymaster.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BasePaymaster *BasePaymasterSession) Owner() (common.Address, error) {
	return _BasePaymaster.Contract.Owner(&_BasePaymaster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BasePaymaster *BasePaymasterCallerSession) Owner() (common.Address, error) {
	return _BasePaymaster.Contract.Owner(&_BasePaymaster.CallOpts)
}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_BasePaymaster *BasePaymasterCaller) TrustedForwarder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasePaymaster.contract.Call(opts, &out, "trustedForwarder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_BasePaymaster *BasePaymasterSession) TrustedForwarder() (common.Address, error) {
	return _BasePaymaster.Contract.TrustedForwarder(&_BasePaymaster.CallOpts)
}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_BasePaymaster *BasePaymasterCallerSession) TrustedForwarder() (common.Address, error) {
	return _BasePaymaster.Contract.TrustedForwarder(&_BasePaymaster.CallOpts)
}

// VersionPaymaster is a free data retrieval call binding the contract method 0x921276ea.
//
// Solidity: function versionPaymaster() view returns(string)
func (_BasePaymaster *BasePaymasterCaller) VersionPaymaster(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BasePaymaster.contract.Call(opts, &out, "versionPaymaster")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VersionPaymaster is a free data retrieval call binding the contract method 0x921276ea.
//
// Solidity: function versionPaymaster() view returns(string)
func (_BasePaymaster *BasePaymasterSession) VersionPaymaster() (string, error) {
	return _BasePaymaster.Contract.VersionPaymaster(&_BasePaymaster.CallOpts)
}

// VersionPaymaster is a free data retrieval call binding the contract method 0x921276ea.
//
// Solidity: function versionPaymaster() view returns(string)
func (_BasePaymaster *BasePaymasterCallerSession) VersionPaymaster() (string, error) {
	return _BasePaymaster.Contract.VersionPaymaster(&_BasePaymaster.CallOpts)
}

// PostRelayedCall is a paid mutator transaction binding the contract method 0x76fa01c3.
//
// Solidity: function postRelayedCall(bytes context, bool success, uint256 gasUseWithoutPost, (uint256,uint256,uint256,address,address,address,bytes,uint256) relayData) returns()
func (_BasePaymaster *BasePaymasterTransactor) PostRelayedCall(opts *bind.TransactOpts, context []byte, success bool, gasUseWithoutPost *big.Int, relayData GsnTypesRelayData) (*types.Transaction, error) {
	return _BasePaymaster.contract.Transact(opts, "postRelayedCall", context, success, gasUseWithoutPost, relayData)
}

// PostRelayedCall is a paid mutator transaction binding the contract method 0x76fa01c3.
//
// Solidity: function postRelayedCall(bytes context, bool success, uint256 gasUseWithoutPost, (uint256,uint256,uint256,address,address,address,bytes,uint256) relayData) returns()
func (_BasePaymaster *BasePaymasterSession) PostRelayedCall(context []byte, success bool, gasUseWithoutPost *big.Int, relayData GsnTypesRelayData) (*types.Transaction, error) {
	return _BasePaymaster.Contract.PostRelayedCall(&_BasePaymaster.TransactOpts, context, success, gasUseWithoutPost, relayData)
}

// PostRelayedCall is a paid mutator transaction binding the contract method 0x76fa01c3.
//
// Solidity: function postRelayedCall(bytes context, bool success, uint256 gasUseWithoutPost, (uint256,uint256,uint256,address,address,address,bytes,uint256) relayData) returns()
func (_BasePaymaster *BasePaymasterTransactorSession) PostRelayedCall(context []byte, success bool, gasUseWithoutPost *big.Int, relayData GsnTypesRelayData) (*types.Transaction, error) {
	return _BasePaymaster.Contract.PostRelayedCall(&_BasePaymaster.TransactOpts, context, success, gasUseWithoutPost, relayData)
}

// PreRelayedCall is a paid mutator transaction binding the contract method 0x00be5dd4.
//
// Solidity: function preRelayedCall(((address,address,uint256,uint256,uint256,bytes,uint256),(uint256,uint256,uint256,address,address,address,bytes,uint256)) relayRequest, bytes signature, bytes approvalData, uint256 maxPossibleGas) returns(bytes context, bool rejectOnRecipientRevert)
func (_BasePaymaster *BasePaymasterTransactor) PreRelayedCall(opts *bind.TransactOpts, relayRequest GsnTypesRelayRequest, signature []byte, approvalData []byte, maxPossibleGas *big.Int) (*types.Transaction, error) {
	return _BasePaymaster.contract.Transact(opts, "preRelayedCall", relayRequest, signature, approvalData, maxPossibleGas)
}

// PreRelayedCall is a paid mutator transaction binding the contract method 0x00be5dd4.
//
// Solidity: function preRelayedCall(((address,address,uint256,uint256,uint256,bytes,uint256),(uint256,uint256,uint256,address,address,address,bytes,uint256)) relayRequest, bytes signature, bytes approvalData, uint256 maxPossibleGas) returns(bytes context, bool rejectOnRecipientRevert)
func (_BasePaymaster *BasePaymasterSession) PreRelayedCall(relayRequest GsnTypesRelayRequest, signature []byte, approvalData []byte, maxPossibleGas *big.Int) (*types.Transaction, error) {
	return _BasePaymaster.Contract.PreRelayedCall(&_BasePaymaster.TransactOpts, relayRequest, signature, approvalData, maxPossibleGas)
}

// PreRelayedCall is a paid mutator transaction binding the contract method 0x00be5dd4.
//
// Solidity: function preRelayedCall(((address,address,uint256,uint256,uint256,bytes,uint256),(uint256,uint256,uint256,address,address,address,bytes,uint256)) relayRequest, bytes signature, bytes approvalData, uint256 maxPossibleGas) returns(bytes context, bool rejectOnRecipientRevert)
func (_BasePaymaster *BasePaymasterTransactorSession) PreRelayedCall(relayRequest GsnTypesRelayRequest, signature []byte, approvalData []byte, maxPossibleGas *big.Int) (*types.Transaction, error) {
	return _BasePaymaster.Contract.PreRelayedCall(&_BasePaymaster.TransactOpts, relayRequest, signature, approvalData, maxPossibleGas)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BasePaymaster *BasePaymasterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasePaymaster.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BasePaymaster *BasePaymasterSession) RenounceOwnership() (*types.Transaction, error) {
	return _BasePaymaster.Contract.RenounceOwnership(&_BasePaymaster.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BasePaymaster *BasePaymasterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BasePaymaster.Contract.RenounceOwnership(&_BasePaymaster.TransactOpts)
}

// SetRelayHub is a paid mutator transaction binding the contract method 0x7bb05264.
//
// Solidity: function setRelayHub(address hub) returns()
func (_BasePaymaster *BasePaymasterTransactor) SetRelayHub(opts *bind.TransactOpts, hub common.Address) (*types.Transaction, error) {
	return _BasePaymaster.contract.Transact(opts, "setRelayHub", hub)
}

// SetRelayHub is a paid mutator transaction binding the contract method 0x7bb05264.
//
// Solidity: function setRelayHub(address hub) returns()
func (_BasePaymaster *BasePaymasterSession) SetRelayHub(hub common.Address) (*types.Transaction, error) {
	return _BasePaymaster.Contract.SetRelayHub(&_BasePaymaster.TransactOpts, hub)
}

// SetRelayHub is a paid mutator transaction binding the contract method 0x7bb05264.
//
// Solidity: function setRelayHub(address hub) returns()
func (_BasePaymaster *BasePaymasterTransactorSession) SetRelayHub(hub common.Address) (*types.Transaction, error) {
	return _BasePaymaster.Contract.SetRelayHub(&_BasePaymaster.TransactOpts, hub)
}

// SetTrustedForwarder is a paid mutator transaction binding the contract method 0xda742228.
//
// Solidity: function setTrustedForwarder(address forwarder) returns()
func (_BasePaymaster *BasePaymasterTransactor) SetTrustedForwarder(opts *bind.TransactOpts, forwarder common.Address) (*types.Transaction, error) {
	return _BasePaymaster.contract.Transact(opts, "setTrustedForwarder", forwarder)
}

// SetTrustedForwarder is a paid mutator transaction binding the contract method 0xda742228.
//
// Solidity: function setTrustedForwarder(address forwarder) returns()
func (_BasePaymaster *BasePaymasterSession) SetTrustedForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _BasePaymaster.Contract.SetTrustedForwarder(&_BasePaymaster.TransactOpts, forwarder)
}

// SetTrustedForwarder is a paid mutator transaction binding the contract method 0xda742228.
//
// Solidity: function setTrustedForwarder(address forwarder) returns()
func (_BasePaymaster *BasePaymasterTransactorSession) SetTrustedForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _BasePaymaster.Contract.SetTrustedForwarder(&_BasePaymaster.TransactOpts, forwarder)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BasePaymaster *BasePaymasterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BasePaymaster.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BasePaymaster *BasePaymasterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BasePaymaster.Contract.TransferOwnership(&_BasePaymaster.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BasePaymaster *BasePaymasterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BasePaymaster.Contract.TransferOwnership(&_BasePaymaster.TransactOpts, newOwner)
}

// WithdrawRelayHubDepositTo is a paid mutator transaction binding the contract method 0x2d14c4b7.
//
// Solidity: function withdrawRelayHubDepositTo(uint256 amount, address target) returns()
func (_BasePaymaster *BasePaymasterTransactor) WithdrawRelayHubDepositTo(opts *bind.TransactOpts, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _BasePaymaster.contract.Transact(opts, "withdrawRelayHubDepositTo", amount, target)
}

// WithdrawRelayHubDepositTo is a paid mutator transaction binding the contract method 0x2d14c4b7.
//
// Solidity: function withdrawRelayHubDepositTo(uint256 amount, address target) returns()
func (_BasePaymaster *BasePaymasterSession) WithdrawRelayHubDepositTo(amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _BasePaymaster.Contract.WithdrawRelayHubDepositTo(&_BasePaymaster.TransactOpts, amount, target)
}

// WithdrawRelayHubDepositTo is a paid mutator transaction binding the contract method 0x2d14c4b7.
//
// Solidity: function withdrawRelayHubDepositTo(uint256 amount, address target) returns()
func (_BasePaymaster *BasePaymasterTransactorSession) WithdrawRelayHubDepositTo(amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _BasePaymaster.Contract.WithdrawRelayHubDepositTo(&_BasePaymaster.TransactOpts, amount, target)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BasePaymaster *BasePaymasterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasePaymaster.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BasePaymaster *BasePaymasterSession) Receive() (*types.Transaction, error) {
	return _BasePaymaster.Contract.Receive(&_BasePaymaster.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BasePaymaster *BasePaymasterTransactorSession) Receive() (*types.Transaction, error) {
	return _BasePaymaster.Contract.Receive(&_BasePaymaster.TransactOpts)
}

// BasePaymasterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BasePaymaster contract.
type BasePaymasterOwnershipTransferredIterator struct {
	Event *BasePaymasterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BasePaymasterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasePaymasterOwnershipTransferred)
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
		it.Event = new(BasePaymasterOwnershipTransferred)
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
func (it *BasePaymasterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasePaymasterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasePaymasterOwnershipTransferred represents a OwnershipTransferred event raised by the BasePaymaster contract.
type BasePaymasterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BasePaymaster *BasePaymasterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BasePaymasterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BasePaymaster.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BasePaymasterOwnershipTransferredIterator{contract: _BasePaymaster.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BasePaymaster *BasePaymasterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BasePaymasterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BasePaymaster.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasePaymasterOwnershipTransferred)
				if err := _BasePaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BasePaymaster *BasePaymasterFilterer) ParseOwnershipTransferred(log types.Log) (*BasePaymasterOwnershipTransferred, error) {
	event := new(BasePaymasterOwnershipTransferred)
	if err := _BasePaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
