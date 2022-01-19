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

// IRelayHubRelayHubConfig is an auto generated low-level Go binding around an user-defined struct.
type IRelayHubRelayHubConfig struct {
	MaxWorkerCount               *big.Int
	GasReserve                   *big.Int
	PostOverhead                 *big.Int
	GasOverhead                  *big.Int
	MaximumRecipientDeposit      *big.Int
	MinimumUnstakeDelay          *big.Int
	MinimumStake                 *big.Int
	DataGasCostPerByte           *big.Int
	ExternalCallDataCostOverhead *big.Int
}

// IRelayHubABI is the input ABI used to generate the binding from.
const IRelayHubABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromBlock\",\"type\":\"uint256\"}],\"name\":\"HubDeprecated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxWorkerCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"postOverhead\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasOverhead\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximumRecipientDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumUnstakeDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dataGasCostPerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"externalCallDataCostOverhead\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIRelayHub.RelayHubConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"RelayHubConfigured\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseRelayFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pctRelayFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"relayUrl\",\"type\":\"string\"}],\"name\":\"RelayServerRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"newRelayWorkers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"workersCount\",\"type\":\"uint256\"}],\"name\":\"RelayWorkersAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayWorker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"innerGasUsed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"TransactionRejectedByPaymaster\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayWorker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"enumIRelayHub.RelayCallStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"charge\",\"type\":\"uint256\"}],\"name\":\"TransactionRelayed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumIRelayHub.RelayCallStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnValue\",\"type\":\"bytes\"}],\"name\":\"TransactionResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newRelayWorkers\",\"type\":\"address[]\"}],\"name\":\"addRelayWorkers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pctRelayFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRelayFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayWorker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"paymasterData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"clientId\",\"type\":\"uint256\"}],\"internalType\":\"structGsnTypes.RelayData\",\"name\":\"relayData\",\"type\":\"tuple\"}],\"name\":\"calculateCharge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"calldataGasCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"depositFor\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromBlock\",\"type\":\"uint256\"}],\"name\":\"deprecateHub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deprecationBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfiguration\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxWorkerCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"postOverhead\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasOverhead\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximumRecipientDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumUnstakeDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dataGasCostPerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"externalCallDataCostOverhead\",\"type\":\"uint256\"}],\"internalType\":\"structIRelayHub.RelayHubConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isDeprecated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayManager\",\"type\":\"address\"}],\"name\":\"isRelayManagerStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayWorker\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"penalize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"penalizer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"baseRelayFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pctRelayFee\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"registerRelayServer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxAcceptanceBudget\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validUntil\",\"type\":\"uint256\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"request\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pctRelayFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRelayFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayWorker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"paymasterData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"clientId\",\"type\":\"uint256\"}],\"internalType\":\"structGsnTypes.RelayData\",\"name\":\"relayData\",\"type\":\"tuple\"}],\"internalType\":\"structGsnTypes.RelayRequest\",\"name\":\"relayRequest\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approvalData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"externalGasLimit\",\"type\":\"uint256\"}],\"name\":\"relayCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"paymasterAccepted\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnValue\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxWorkerCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"postOverhead\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasOverhead\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximumRecipientDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumUnstakeDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dataGasCostPerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"externalCallDataCostOverhead\",\"type\":\"uint256\"}],\"internalType\":\"structIRelayHub.RelayHubConfig\",\"name\":\"_config\",\"type\":\"tuple\"}],\"name\":\"setConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeManager\",\"outputs\":[{\"internalType\":\"contractIStakeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionHub\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"dest\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"workerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"}],\"name\":\"workerToManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IRelayHub is an auto generated Go binding around an Ethereum contract.
type IRelayHub struct {
	IRelayHubCaller     // Read-only binding to the contract
	IRelayHubTransactor // Write-only binding to the contract
	IRelayHubFilterer   // Log filterer for contract events
}

// IRelayHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRelayHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRelayHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRelayHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRelayHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRelayHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRelayHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRelayHubSession struct {
	Contract     *IRelayHub        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRelayHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRelayHubCallerSession struct {
	Contract *IRelayHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IRelayHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRelayHubTransactorSession struct {
	Contract     *IRelayHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IRelayHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRelayHubRaw struct {
	Contract *IRelayHub // Generic contract binding to access the raw methods on
}

// IRelayHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRelayHubCallerRaw struct {
	Contract *IRelayHubCaller // Generic read-only contract binding to access the raw methods on
}

// IRelayHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRelayHubTransactorRaw struct {
	Contract *IRelayHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRelayHub creates a new instance of IRelayHub, bound to a specific deployed contract.
func NewIRelayHub(address common.Address, backend bind.ContractBackend) (*IRelayHub, error) {
	contract, err := bindIRelayHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRelayHub{IRelayHubCaller: IRelayHubCaller{contract: contract}, IRelayHubTransactor: IRelayHubTransactor{contract: contract}, IRelayHubFilterer: IRelayHubFilterer{contract: contract}}, nil
}

// NewIRelayHubCaller creates a new read-only instance of IRelayHub, bound to a specific deployed contract.
func NewIRelayHubCaller(address common.Address, caller bind.ContractCaller) (*IRelayHubCaller, error) {
	contract, err := bindIRelayHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRelayHubCaller{contract: contract}, nil
}

// NewIRelayHubTransactor creates a new write-only instance of IRelayHub, bound to a specific deployed contract.
func NewIRelayHubTransactor(address common.Address, transactor bind.ContractTransactor) (*IRelayHubTransactor, error) {
	contract, err := bindIRelayHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRelayHubTransactor{contract: contract}, nil
}

// NewIRelayHubFilterer creates a new log filterer instance of IRelayHub, bound to a specific deployed contract.
func NewIRelayHubFilterer(address common.Address, filterer bind.ContractFilterer) (*IRelayHubFilterer, error) {
	contract, err := bindIRelayHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRelayHubFilterer{contract: contract}, nil
}

// bindIRelayHub binds a generic wrapper to an already deployed contract.
func bindIRelayHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IRelayHubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRelayHub *IRelayHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRelayHub.Contract.IRelayHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRelayHub *IRelayHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRelayHub.Contract.IRelayHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRelayHub *IRelayHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRelayHub.Contract.IRelayHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRelayHub *IRelayHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRelayHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRelayHub *IRelayHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRelayHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRelayHub *IRelayHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRelayHub.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address target) view returns(uint256)
func (_IRelayHub *IRelayHubCaller) BalanceOf(opts *bind.CallOpts, target common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IRelayHub.contract.Call(opts, &out, "balanceOf", target)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address target) view returns(uint256)
func (_IRelayHub *IRelayHubSession) BalanceOf(target common.Address) (*big.Int, error) {
	return _IRelayHub.Contract.BalanceOf(&_IRelayHub.CallOpts, target)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address target) view returns(uint256)
func (_IRelayHub *IRelayHubCallerSession) BalanceOf(target common.Address) (*big.Int, error) {
	return _IRelayHub.Contract.BalanceOf(&_IRelayHub.CallOpts, target)
}

// CalculateCharge is a free data retrieval call binding the contract method 0x8e53548b.
//
// Solidity: function calculateCharge(uint256 gasUsed, (uint256,uint256,uint256,address,address,address,bytes,uint256) relayData) view returns(uint256)
func (_IRelayHub *IRelayHubCaller) CalculateCharge(opts *bind.CallOpts, gasUsed *big.Int, relayData GsnTypesRelayData) (*big.Int, error) {
	var out []interface{}
	err := _IRelayHub.contract.Call(opts, &out, "calculateCharge", gasUsed, relayData)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateCharge is a free data retrieval call binding the contract method 0x8e53548b.
//
// Solidity: function calculateCharge(uint256 gasUsed, (uint256,uint256,uint256,address,address,address,bytes,uint256) relayData) view returns(uint256)
func (_IRelayHub *IRelayHubSession) CalculateCharge(gasUsed *big.Int, relayData GsnTypesRelayData) (*big.Int, error) {
	return _IRelayHub.Contract.CalculateCharge(&_IRelayHub.CallOpts, gasUsed, relayData)
}

// CalculateCharge is a free data retrieval call binding the contract method 0x8e53548b.
//
// Solidity: function calculateCharge(uint256 gasUsed, (uint256,uint256,uint256,address,address,address,bytes,uint256) relayData) view returns(uint256)
func (_IRelayHub *IRelayHubCallerSession) CalculateCharge(gasUsed *big.Int, relayData GsnTypesRelayData) (*big.Int, error) {
	return _IRelayHub.Contract.CalculateCharge(&_IRelayHub.CallOpts, gasUsed, relayData)
}

// CalldataGasCost is a free data retrieval call binding the contract method 0x26595b9d.
//
// Solidity: function calldataGasCost(uint256 length) view returns(uint256)
func (_IRelayHub *IRelayHubCaller) CalldataGasCost(opts *bind.CallOpts, length *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IRelayHub.contract.Call(opts, &out, "calldataGasCost", length)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalldataGasCost is a free data retrieval call binding the contract method 0x26595b9d.
//
// Solidity: function calldataGasCost(uint256 length) view returns(uint256)
func (_IRelayHub *IRelayHubSession) CalldataGasCost(length *big.Int) (*big.Int, error) {
	return _IRelayHub.Contract.CalldataGasCost(&_IRelayHub.CallOpts, length)
}

// CalldataGasCost is a free data retrieval call binding the contract method 0x26595b9d.
//
// Solidity: function calldataGasCost(uint256 length) view returns(uint256)
func (_IRelayHub *IRelayHubCallerSession) CalldataGasCost(length *big.Int) (*big.Int, error) {
	return _IRelayHub.Contract.CalldataGasCost(&_IRelayHub.CallOpts, length)
}

// DeprecationBlock is a free data retrieval call binding the contract method 0xd6a71c0d.
//
// Solidity: function deprecationBlock() view returns(uint256)
func (_IRelayHub *IRelayHubCaller) DeprecationBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IRelayHub.contract.Call(opts, &out, "deprecationBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeprecationBlock is a free data retrieval call binding the contract method 0xd6a71c0d.
//
// Solidity: function deprecationBlock() view returns(uint256)
func (_IRelayHub *IRelayHubSession) DeprecationBlock() (*big.Int, error) {
	return _IRelayHub.Contract.DeprecationBlock(&_IRelayHub.CallOpts)
}

// DeprecationBlock is a free data retrieval call binding the contract method 0xd6a71c0d.
//
// Solidity: function deprecationBlock() view returns(uint256)
func (_IRelayHub *IRelayHubCallerSession) DeprecationBlock() (*big.Int, error) {
	return _IRelayHub.Contract.DeprecationBlock(&_IRelayHub.CallOpts)
}

// GetConfiguration is a free data retrieval call binding the contract method 0x6bd50cef.
//
// Solidity: function getConfiguration() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) config)
func (_IRelayHub *IRelayHubCaller) GetConfiguration(opts *bind.CallOpts) (IRelayHubRelayHubConfig, error) {
	var out []interface{}
	err := _IRelayHub.contract.Call(opts, &out, "getConfiguration")

	if err != nil {
		return *new(IRelayHubRelayHubConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IRelayHubRelayHubConfig)).(*IRelayHubRelayHubConfig)

	return out0, err

}

// GetConfiguration is a free data retrieval call binding the contract method 0x6bd50cef.
//
// Solidity: function getConfiguration() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) config)
func (_IRelayHub *IRelayHubSession) GetConfiguration() (IRelayHubRelayHubConfig, error) {
	return _IRelayHub.Contract.GetConfiguration(&_IRelayHub.CallOpts)
}

// GetConfiguration is a free data retrieval call binding the contract method 0x6bd50cef.
//
// Solidity: function getConfiguration() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) config)
func (_IRelayHub *IRelayHubCallerSession) GetConfiguration() (IRelayHubRelayHubConfig, error) {
	return _IRelayHub.Contract.GetConfiguration(&_IRelayHub.CallOpts)
}

// IsDeprecated is a free data retrieval call binding the contract method 0xc7178230.
//
// Solidity: function isDeprecated() view returns(bool)
func (_IRelayHub *IRelayHubCaller) IsDeprecated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IRelayHub.contract.Call(opts, &out, "isDeprecated")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDeprecated is a free data retrieval call binding the contract method 0xc7178230.
//
// Solidity: function isDeprecated() view returns(bool)
func (_IRelayHub *IRelayHubSession) IsDeprecated() (bool, error) {
	return _IRelayHub.Contract.IsDeprecated(&_IRelayHub.CallOpts)
}

// IsDeprecated is a free data retrieval call binding the contract method 0xc7178230.
//
// Solidity: function isDeprecated() view returns(bool)
func (_IRelayHub *IRelayHubCallerSession) IsDeprecated() (bool, error) {
	return _IRelayHub.Contract.IsDeprecated(&_IRelayHub.CallOpts)
}

// IsRelayManagerStaked is a free data retrieval call binding the contract method 0x2ad311b5.
//
// Solidity: function isRelayManagerStaked(address relayManager) view returns(bool)
func (_IRelayHub *IRelayHubCaller) IsRelayManagerStaked(opts *bind.CallOpts, relayManager common.Address) (bool, error) {
	var out []interface{}
	err := _IRelayHub.contract.Call(opts, &out, "isRelayManagerStaked", relayManager)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRelayManagerStaked is a free data retrieval call binding the contract method 0x2ad311b5.
//
// Solidity: function isRelayManagerStaked(address relayManager) view returns(bool)
func (_IRelayHub *IRelayHubSession) IsRelayManagerStaked(relayManager common.Address) (bool, error) {
	return _IRelayHub.Contract.IsRelayManagerStaked(&_IRelayHub.CallOpts, relayManager)
}

// IsRelayManagerStaked is a free data retrieval call binding the contract method 0x2ad311b5.
//
// Solidity: function isRelayManagerStaked(address relayManager) view returns(bool)
func (_IRelayHub *IRelayHubCallerSession) IsRelayManagerStaked(relayManager common.Address) (bool, error) {
	return _IRelayHub.Contract.IsRelayManagerStaked(&_IRelayHub.CallOpts, relayManager)
}

// Penalizer is a free data retrieval call binding the contract method 0xc4775a68.
//
// Solidity: function penalizer() view returns(address)
func (_IRelayHub *IRelayHubCaller) Penalizer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IRelayHub.contract.Call(opts, &out, "penalizer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Penalizer is a free data retrieval call binding the contract method 0xc4775a68.
//
// Solidity: function penalizer() view returns(address)
func (_IRelayHub *IRelayHubSession) Penalizer() (common.Address, error) {
	return _IRelayHub.Contract.Penalizer(&_IRelayHub.CallOpts)
}

// Penalizer is a free data retrieval call binding the contract method 0xc4775a68.
//
// Solidity: function penalizer() view returns(address)
func (_IRelayHub *IRelayHubCallerSession) Penalizer() (common.Address, error) {
	return _IRelayHub.Contract.Penalizer(&_IRelayHub.CallOpts)
}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_IRelayHub *IRelayHubCaller) StakeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IRelayHub.contract.Call(opts, &out, "stakeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_IRelayHub *IRelayHubSession) StakeManager() (common.Address, error) {
	return _IRelayHub.Contract.StakeManager(&_IRelayHub.CallOpts)
}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_IRelayHub *IRelayHubCallerSession) StakeManager() (common.Address, error) {
	return _IRelayHub.Contract.StakeManager(&_IRelayHub.CallOpts)
}

// VersionHub is a free data retrieval call binding the contract method 0xd904c732.
//
// Solidity: function versionHub() view returns(string)
func (_IRelayHub *IRelayHubCaller) VersionHub(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IRelayHub.contract.Call(opts, &out, "versionHub")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VersionHub is a free data retrieval call binding the contract method 0xd904c732.
//
// Solidity: function versionHub() view returns(string)
func (_IRelayHub *IRelayHubSession) VersionHub() (string, error) {
	return _IRelayHub.Contract.VersionHub(&_IRelayHub.CallOpts)
}

// VersionHub is a free data retrieval call binding the contract method 0xd904c732.
//
// Solidity: function versionHub() view returns(string)
func (_IRelayHub *IRelayHubCallerSession) VersionHub() (string, error) {
	return _IRelayHub.Contract.VersionHub(&_IRelayHub.CallOpts)
}

// WorkerCount is a free data retrieval call binding the contract method 0x194ac307.
//
// Solidity: function workerCount(address manager) view returns(uint256)
func (_IRelayHub *IRelayHubCaller) WorkerCount(opts *bind.CallOpts, manager common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IRelayHub.contract.Call(opts, &out, "workerCount", manager)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WorkerCount is a free data retrieval call binding the contract method 0x194ac307.
//
// Solidity: function workerCount(address manager) view returns(uint256)
func (_IRelayHub *IRelayHubSession) WorkerCount(manager common.Address) (*big.Int, error) {
	return _IRelayHub.Contract.WorkerCount(&_IRelayHub.CallOpts, manager)
}

// WorkerCount is a free data retrieval call binding the contract method 0x194ac307.
//
// Solidity: function workerCount(address manager) view returns(uint256)
func (_IRelayHub *IRelayHubCallerSession) WorkerCount(manager common.Address) (*big.Int, error) {
	return _IRelayHub.Contract.WorkerCount(&_IRelayHub.CallOpts, manager)
}

// WorkerToManager is a free data retrieval call binding the contract method 0xca998f56.
//
// Solidity: function workerToManager(address worker) view returns(address)
func (_IRelayHub *IRelayHubCaller) WorkerToManager(opts *bind.CallOpts, worker common.Address) (common.Address, error) {
	var out []interface{}
	err := _IRelayHub.contract.Call(opts, &out, "workerToManager", worker)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WorkerToManager is a free data retrieval call binding the contract method 0xca998f56.
//
// Solidity: function workerToManager(address worker) view returns(address)
func (_IRelayHub *IRelayHubSession) WorkerToManager(worker common.Address) (common.Address, error) {
	return _IRelayHub.Contract.WorkerToManager(&_IRelayHub.CallOpts, worker)
}

// WorkerToManager is a free data retrieval call binding the contract method 0xca998f56.
//
// Solidity: function workerToManager(address worker) view returns(address)
func (_IRelayHub *IRelayHubCallerSession) WorkerToManager(worker common.Address) (common.Address, error) {
	return _IRelayHub.Contract.WorkerToManager(&_IRelayHub.CallOpts, worker)
}

// AddRelayWorkers is a paid mutator transaction binding the contract method 0xc2da0786.
//
// Solidity: function addRelayWorkers(address[] newRelayWorkers) returns()
func (_IRelayHub *IRelayHubTransactor) AddRelayWorkers(opts *bind.TransactOpts, newRelayWorkers []common.Address) (*types.Transaction, error) {
	return _IRelayHub.contract.Transact(opts, "addRelayWorkers", newRelayWorkers)
}

// AddRelayWorkers is a paid mutator transaction binding the contract method 0xc2da0786.
//
// Solidity: function addRelayWorkers(address[] newRelayWorkers) returns()
func (_IRelayHub *IRelayHubSession) AddRelayWorkers(newRelayWorkers []common.Address) (*types.Transaction, error) {
	return _IRelayHub.Contract.AddRelayWorkers(&_IRelayHub.TransactOpts, newRelayWorkers)
}

// AddRelayWorkers is a paid mutator transaction binding the contract method 0xc2da0786.
//
// Solidity: function addRelayWorkers(address[] newRelayWorkers) returns()
func (_IRelayHub *IRelayHubTransactorSession) AddRelayWorkers(newRelayWorkers []common.Address) (*types.Transaction, error) {
	return _IRelayHub.Contract.AddRelayWorkers(&_IRelayHub.TransactOpts, newRelayWorkers)
}

// DepositFor is a paid mutator transaction binding the contract method 0xaa67c919.
//
// Solidity: function depositFor(address target) payable returns()
func (_IRelayHub *IRelayHubTransactor) DepositFor(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _IRelayHub.contract.Transact(opts, "depositFor", target)
}

// DepositFor is a paid mutator transaction binding the contract method 0xaa67c919.
//
// Solidity: function depositFor(address target) payable returns()
func (_IRelayHub *IRelayHubSession) DepositFor(target common.Address) (*types.Transaction, error) {
	return _IRelayHub.Contract.DepositFor(&_IRelayHub.TransactOpts, target)
}

// DepositFor is a paid mutator transaction binding the contract method 0xaa67c919.
//
// Solidity: function depositFor(address target) payable returns()
func (_IRelayHub *IRelayHubTransactorSession) DepositFor(target common.Address) (*types.Transaction, error) {
	return _IRelayHub.Contract.DepositFor(&_IRelayHub.TransactOpts, target)
}

// DeprecateHub is a paid mutator transaction binding the contract method 0xaf595dfc.
//
// Solidity: function deprecateHub(uint256 fromBlock) returns()
func (_IRelayHub *IRelayHubTransactor) DeprecateHub(opts *bind.TransactOpts, fromBlock *big.Int) (*types.Transaction, error) {
	return _IRelayHub.contract.Transact(opts, "deprecateHub", fromBlock)
}

// DeprecateHub is a paid mutator transaction binding the contract method 0xaf595dfc.
//
// Solidity: function deprecateHub(uint256 fromBlock) returns()
func (_IRelayHub *IRelayHubSession) DeprecateHub(fromBlock *big.Int) (*types.Transaction, error) {
	return _IRelayHub.Contract.DeprecateHub(&_IRelayHub.TransactOpts, fromBlock)
}

// DeprecateHub is a paid mutator transaction binding the contract method 0xaf595dfc.
//
// Solidity: function deprecateHub(uint256 fromBlock) returns()
func (_IRelayHub *IRelayHubTransactorSession) DeprecateHub(fromBlock *big.Int) (*types.Transaction, error) {
	return _IRelayHub.Contract.DeprecateHub(&_IRelayHub.TransactOpts, fromBlock)
}

// Penalize is a paid mutator transaction binding the contract method 0xebcd31ac.
//
// Solidity: function penalize(address relayWorker, address beneficiary) returns()
func (_IRelayHub *IRelayHubTransactor) Penalize(opts *bind.TransactOpts, relayWorker common.Address, beneficiary common.Address) (*types.Transaction, error) {
	return _IRelayHub.contract.Transact(opts, "penalize", relayWorker, beneficiary)
}

// Penalize is a paid mutator transaction binding the contract method 0xebcd31ac.
//
// Solidity: function penalize(address relayWorker, address beneficiary) returns()
func (_IRelayHub *IRelayHubSession) Penalize(relayWorker common.Address, beneficiary common.Address) (*types.Transaction, error) {
	return _IRelayHub.Contract.Penalize(&_IRelayHub.TransactOpts, relayWorker, beneficiary)
}

// Penalize is a paid mutator transaction binding the contract method 0xebcd31ac.
//
// Solidity: function penalize(address relayWorker, address beneficiary) returns()
func (_IRelayHub *IRelayHubTransactorSession) Penalize(relayWorker common.Address, beneficiary common.Address) (*types.Transaction, error) {
	return _IRelayHub.Contract.Penalize(&_IRelayHub.TransactOpts, relayWorker, beneficiary)
}

// RegisterRelayServer is a paid mutator transaction binding the contract method 0x83b71871.
//
// Solidity: function registerRelayServer(uint256 baseRelayFee, uint256 pctRelayFee, string url) returns()
func (_IRelayHub *IRelayHubTransactor) RegisterRelayServer(opts *bind.TransactOpts, baseRelayFee *big.Int, pctRelayFee *big.Int, url string) (*types.Transaction, error) {
	return _IRelayHub.contract.Transact(opts, "registerRelayServer", baseRelayFee, pctRelayFee, url)
}

// RegisterRelayServer is a paid mutator transaction binding the contract method 0x83b71871.
//
// Solidity: function registerRelayServer(uint256 baseRelayFee, uint256 pctRelayFee, string url) returns()
func (_IRelayHub *IRelayHubSession) RegisterRelayServer(baseRelayFee *big.Int, pctRelayFee *big.Int, url string) (*types.Transaction, error) {
	return _IRelayHub.Contract.RegisterRelayServer(&_IRelayHub.TransactOpts, baseRelayFee, pctRelayFee, url)
}

// RegisterRelayServer is a paid mutator transaction binding the contract method 0x83b71871.
//
// Solidity: function registerRelayServer(uint256 baseRelayFee, uint256 pctRelayFee, string url) returns()
func (_IRelayHub *IRelayHubTransactorSession) RegisterRelayServer(baseRelayFee *big.Int, pctRelayFee *big.Int, url string) (*types.Transaction, error) {
	return _IRelayHub.Contract.RegisterRelayServer(&_IRelayHub.TransactOpts, baseRelayFee, pctRelayFee, url)
}

// RelayCall is a paid mutator transaction binding the contract method 0x10c45431.
//
// Solidity: function relayCall(uint256 maxAcceptanceBudget, ((address,address,uint256,uint256,uint256,bytes,uint256),(uint256,uint256,uint256,address,address,address,bytes,uint256)) relayRequest, bytes signature, bytes approvalData, uint256 externalGasLimit) returns(bool paymasterAccepted, bytes returnValue)
func (_IRelayHub *IRelayHubTransactor) RelayCall(opts *bind.TransactOpts, maxAcceptanceBudget *big.Int, relayRequest GsnTypesRelayRequest, signature []byte, approvalData []byte, externalGasLimit *big.Int) (*types.Transaction, error) {
	return _IRelayHub.contract.Transact(opts, "relayCall", maxAcceptanceBudget, relayRequest, signature, approvalData, externalGasLimit)
}

// RelayCall is a paid mutator transaction binding the contract method 0x10c45431.
//
// Solidity: function relayCall(uint256 maxAcceptanceBudget, ((address,address,uint256,uint256,uint256,bytes,uint256),(uint256,uint256,uint256,address,address,address,bytes,uint256)) relayRequest, bytes signature, bytes approvalData, uint256 externalGasLimit) returns(bool paymasterAccepted, bytes returnValue)
func (_IRelayHub *IRelayHubSession) RelayCall(maxAcceptanceBudget *big.Int, relayRequest GsnTypesRelayRequest, signature []byte, approvalData []byte, externalGasLimit *big.Int) (*types.Transaction, error) {
	return _IRelayHub.Contract.RelayCall(&_IRelayHub.TransactOpts, maxAcceptanceBudget, relayRequest, signature, approvalData, externalGasLimit)
}

// RelayCall is a paid mutator transaction binding the contract method 0x10c45431.
//
// Solidity: function relayCall(uint256 maxAcceptanceBudget, ((address,address,uint256,uint256,uint256,bytes,uint256),(uint256,uint256,uint256,address,address,address,bytes,uint256)) relayRequest, bytes signature, bytes approvalData, uint256 externalGasLimit) returns(bool paymasterAccepted, bytes returnValue)
func (_IRelayHub *IRelayHubTransactorSession) RelayCall(maxAcceptanceBudget *big.Int, relayRequest GsnTypesRelayRequest, signature []byte, approvalData []byte, externalGasLimit *big.Int) (*types.Transaction, error) {
	return _IRelayHub.Contract.RelayCall(&_IRelayHub.TransactOpts, maxAcceptanceBudget, relayRequest, signature, approvalData, externalGasLimit)
}

// SetConfiguration is a paid mutator transaction binding the contract method 0xc651bce8.
//
// Solidity: function setConfiguration((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) _config) returns()
func (_IRelayHub *IRelayHubTransactor) SetConfiguration(opts *bind.TransactOpts, _config IRelayHubRelayHubConfig) (*types.Transaction, error) {
	return _IRelayHub.contract.Transact(opts, "setConfiguration", _config)
}

// SetConfiguration is a paid mutator transaction binding the contract method 0xc651bce8.
//
// Solidity: function setConfiguration((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) _config) returns()
func (_IRelayHub *IRelayHubSession) SetConfiguration(_config IRelayHubRelayHubConfig) (*types.Transaction, error) {
	return _IRelayHub.Contract.SetConfiguration(&_IRelayHub.TransactOpts, _config)
}

// SetConfiguration is a paid mutator transaction binding the contract method 0xc651bce8.
//
// Solidity: function setConfiguration((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) _config) returns()
func (_IRelayHub *IRelayHubTransactorSession) SetConfiguration(_config IRelayHubRelayHubConfig) (*types.Transaction, error) {
	return _IRelayHub.Contract.SetConfiguration(&_IRelayHub.TransactOpts, _config)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address dest) returns()
func (_IRelayHub *IRelayHubTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, dest common.Address) (*types.Transaction, error) {
	return _IRelayHub.contract.Transact(opts, "withdraw", amount, dest)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address dest) returns()
func (_IRelayHub *IRelayHubSession) Withdraw(amount *big.Int, dest common.Address) (*types.Transaction, error) {
	return _IRelayHub.Contract.Withdraw(&_IRelayHub.TransactOpts, amount, dest)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address dest) returns()
func (_IRelayHub *IRelayHubTransactorSession) Withdraw(amount *big.Int, dest common.Address) (*types.Transaction, error) {
	return _IRelayHub.Contract.Withdraw(&_IRelayHub.TransactOpts, amount, dest)
}

// IRelayHubDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the IRelayHub contract.
type IRelayHubDepositedIterator struct {
	Event *IRelayHubDeposited // Event containing the contract specifics and raw log

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
func (it *IRelayHubDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRelayHubDeposited)
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
		it.Event = new(IRelayHubDeposited)
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
func (it *IRelayHubDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRelayHubDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRelayHubDeposited represents a Deposited event raised by the IRelayHub contract.
type IRelayHubDeposited struct {
	Paymaster common.Address
	From      common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address indexed paymaster, address indexed from, uint256 amount)
func (_IRelayHub *IRelayHubFilterer) FilterDeposited(opts *bind.FilterOpts, paymaster []common.Address, from []common.Address) (*IRelayHubDepositedIterator, error) {

	var paymasterRule []interface{}
	for _, paymasterItem := range paymaster {
		paymasterRule = append(paymasterRule, paymasterItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IRelayHub.contract.FilterLogs(opts, "Deposited", paymasterRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &IRelayHubDepositedIterator{contract: _IRelayHub.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address indexed paymaster, address indexed from, uint256 amount)
func (_IRelayHub *IRelayHubFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *IRelayHubDeposited, paymaster []common.Address, from []common.Address) (event.Subscription, error) {

	var paymasterRule []interface{}
	for _, paymasterItem := range paymaster {
		paymasterRule = append(paymasterRule, paymasterItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IRelayHub.contract.WatchLogs(opts, "Deposited", paymasterRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRelayHubDeposited)
				if err := _IRelayHub.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address indexed paymaster, address indexed from, uint256 amount)
func (_IRelayHub *IRelayHubFilterer) ParseDeposited(log types.Log) (*IRelayHubDeposited, error) {
	event := new(IRelayHubDeposited)
	if err := _IRelayHub.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRelayHubHubDeprecatedIterator is returned from FilterHubDeprecated and is used to iterate over the raw logs and unpacked data for HubDeprecated events raised by the IRelayHub contract.
type IRelayHubHubDeprecatedIterator struct {
	Event *IRelayHubHubDeprecated // Event containing the contract specifics and raw log

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
func (it *IRelayHubHubDeprecatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRelayHubHubDeprecated)
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
		it.Event = new(IRelayHubHubDeprecated)
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
func (it *IRelayHubHubDeprecatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRelayHubHubDeprecatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRelayHubHubDeprecated represents a HubDeprecated event raised by the IRelayHub contract.
type IRelayHubHubDeprecated struct {
	FromBlock *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterHubDeprecated is a free log retrieval operation binding the contract event 0x1c0aa0c666483fbf0cf795d9d646ea3552d1e3008162ba9ab1d6d6dfd8c6ec6b.
//
// Solidity: event HubDeprecated(uint256 fromBlock)
func (_IRelayHub *IRelayHubFilterer) FilterHubDeprecated(opts *bind.FilterOpts) (*IRelayHubHubDeprecatedIterator, error) {

	logs, sub, err := _IRelayHub.contract.FilterLogs(opts, "HubDeprecated")
	if err != nil {
		return nil, err
	}
	return &IRelayHubHubDeprecatedIterator{contract: _IRelayHub.contract, event: "HubDeprecated", logs: logs, sub: sub}, nil
}

// WatchHubDeprecated is a free log subscription operation binding the contract event 0x1c0aa0c666483fbf0cf795d9d646ea3552d1e3008162ba9ab1d6d6dfd8c6ec6b.
//
// Solidity: event HubDeprecated(uint256 fromBlock)
func (_IRelayHub *IRelayHubFilterer) WatchHubDeprecated(opts *bind.WatchOpts, sink chan<- *IRelayHubHubDeprecated) (event.Subscription, error) {

	logs, sub, err := _IRelayHub.contract.WatchLogs(opts, "HubDeprecated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRelayHubHubDeprecated)
				if err := _IRelayHub.contract.UnpackLog(event, "HubDeprecated", log); err != nil {
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

// ParseHubDeprecated is a log parse operation binding the contract event 0x1c0aa0c666483fbf0cf795d9d646ea3552d1e3008162ba9ab1d6d6dfd8c6ec6b.
//
// Solidity: event HubDeprecated(uint256 fromBlock)
func (_IRelayHub *IRelayHubFilterer) ParseHubDeprecated(log types.Log) (*IRelayHubHubDeprecated, error) {
	event := new(IRelayHubHubDeprecated)
	if err := _IRelayHub.contract.UnpackLog(event, "HubDeprecated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRelayHubRelayHubConfiguredIterator is returned from FilterRelayHubConfigured and is used to iterate over the raw logs and unpacked data for RelayHubConfigured events raised by the IRelayHub contract.
type IRelayHubRelayHubConfiguredIterator struct {
	Event *IRelayHubRelayHubConfigured // Event containing the contract specifics and raw log

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
func (it *IRelayHubRelayHubConfiguredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRelayHubRelayHubConfigured)
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
		it.Event = new(IRelayHubRelayHubConfigured)
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
func (it *IRelayHubRelayHubConfiguredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRelayHubRelayHubConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRelayHubRelayHubConfigured represents a RelayHubConfigured event raised by the IRelayHub contract.
type IRelayHubRelayHubConfigured struct {
	Config IRelayHubRelayHubConfig
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRelayHubConfigured is a free log retrieval operation binding the contract event 0x918ee002eb112844e457f37ea6b320c5572bc73957ebb0423ffcfb03d7b939d7.
//
// Solidity: event RelayHubConfigured((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) config)
func (_IRelayHub *IRelayHubFilterer) FilterRelayHubConfigured(opts *bind.FilterOpts) (*IRelayHubRelayHubConfiguredIterator, error) {

	logs, sub, err := _IRelayHub.contract.FilterLogs(opts, "RelayHubConfigured")
	if err != nil {
		return nil, err
	}
	return &IRelayHubRelayHubConfiguredIterator{contract: _IRelayHub.contract, event: "RelayHubConfigured", logs: logs, sub: sub}, nil
}

// WatchRelayHubConfigured is a free log subscription operation binding the contract event 0x918ee002eb112844e457f37ea6b320c5572bc73957ebb0423ffcfb03d7b939d7.
//
// Solidity: event RelayHubConfigured((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) config)
func (_IRelayHub *IRelayHubFilterer) WatchRelayHubConfigured(opts *bind.WatchOpts, sink chan<- *IRelayHubRelayHubConfigured) (event.Subscription, error) {

	logs, sub, err := _IRelayHub.contract.WatchLogs(opts, "RelayHubConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRelayHubRelayHubConfigured)
				if err := _IRelayHub.contract.UnpackLog(event, "RelayHubConfigured", log); err != nil {
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

// ParseRelayHubConfigured is a log parse operation binding the contract event 0x918ee002eb112844e457f37ea6b320c5572bc73957ebb0423ffcfb03d7b939d7.
//
// Solidity: event RelayHubConfigured((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) config)
func (_IRelayHub *IRelayHubFilterer) ParseRelayHubConfigured(log types.Log) (*IRelayHubRelayHubConfigured, error) {
	event := new(IRelayHubRelayHubConfigured)
	if err := _IRelayHub.contract.UnpackLog(event, "RelayHubConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRelayHubRelayServerRegisteredIterator is returned from FilterRelayServerRegistered and is used to iterate over the raw logs and unpacked data for RelayServerRegistered events raised by the IRelayHub contract.
type IRelayHubRelayServerRegisteredIterator struct {
	Event *IRelayHubRelayServerRegistered // Event containing the contract specifics and raw log

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
func (it *IRelayHubRelayServerRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRelayHubRelayServerRegistered)
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
		it.Event = new(IRelayHubRelayServerRegistered)
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
func (it *IRelayHubRelayServerRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRelayHubRelayServerRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRelayHubRelayServerRegistered represents a RelayServerRegistered event raised by the IRelayHub contract.
type IRelayHubRelayServerRegistered struct {
	RelayManager common.Address
	BaseRelayFee *big.Int
	PctRelayFee  *big.Int
	RelayUrl     string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRelayServerRegistered is a free log retrieval operation binding the contract event 0x77f2d8afec4b9d82ffa0dea525320620292bd1067f575964994d5c4501479aed.
//
// Solidity: event RelayServerRegistered(address indexed relayManager, uint256 baseRelayFee, uint256 pctRelayFee, string relayUrl)
func (_IRelayHub *IRelayHubFilterer) FilterRelayServerRegistered(opts *bind.FilterOpts, relayManager []common.Address) (*IRelayHubRelayServerRegisteredIterator, error) {

	var relayManagerRule []interface{}
	for _, relayManagerItem := range relayManager {
		relayManagerRule = append(relayManagerRule, relayManagerItem)
	}

	logs, sub, err := _IRelayHub.contract.FilterLogs(opts, "RelayServerRegistered", relayManagerRule)
	if err != nil {
		return nil, err
	}
	return &IRelayHubRelayServerRegisteredIterator{contract: _IRelayHub.contract, event: "RelayServerRegistered", logs: logs, sub: sub}, nil
}

// WatchRelayServerRegistered is a free log subscription operation binding the contract event 0x77f2d8afec4b9d82ffa0dea525320620292bd1067f575964994d5c4501479aed.
//
// Solidity: event RelayServerRegistered(address indexed relayManager, uint256 baseRelayFee, uint256 pctRelayFee, string relayUrl)
func (_IRelayHub *IRelayHubFilterer) WatchRelayServerRegistered(opts *bind.WatchOpts, sink chan<- *IRelayHubRelayServerRegistered, relayManager []common.Address) (event.Subscription, error) {

	var relayManagerRule []interface{}
	for _, relayManagerItem := range relayManager {
		relayManagerRule = append(relayManagerRule, relayManagerItem)
	}

	logs, sub, err := _IRelayHub.contract.WatchLogs(opts, "RelayServerRegistered", relayManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRelayHubRelayServerRegistered)
				if err := _IRelayHub.contract.UnpackLog(event, "RelayServerRegistered", log); err != nil {
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

// ParseRelayServerRegistered is a log parse operation binding the contract event 0x77f2d8afec4b9d82ffa0dea525320620292bd1067f575964994d5c4501479aed.
//
// Solidity: event RelayServerRegistered(address indexed relayManager, uint256 baseRelayFee, uint256 pctRelayFee, string relayUrl)
func (_IRelayHub *IRelayHubFilterer) ParseRelayServerRegistered(log types.Log) (*IRelayHubRelayServerRegistered, error) {
	event := new(IRelayHubRelayServerRegistered)
	if err := _IRelayHub.contract.UnpackLog(event, "RelayServerRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRelayHubRelayWorkersAddedIterator is returned from FilterRelayWorkersAdded and is used to iterate over the raw logs and unpacked data for RelayWorkersAdded events raised by the IRelayHub contract.
type IRelayHubRelayWorkersAddedIterator struct {
	Event *IRelayHubRelayWorkersAdded // Event containing the contract specifics and raw log

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
func (it *IRelayHubRelayWorkersAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRelayHubRelayWorkersAdded)
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
		it.Event = new(IRelayHubRelayWorkersAdded)
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
func (it *IRelayHubRelayWorkersAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRelayHubRelayWorkersAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRelayHubRelayWorkersAdded represents a RelayWorkersAdded event raised by the IRelayHub contract.
type IRelayHubRelayWorkersAdded struct {
	RelayManager    common.Address
	NewRelayWorkers []common.Address
	WorkersCount    *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRelayWorkersAdded is a free log retrieval operation binding the contract event 0xebf4a9bffb39f7c5dbf3f65540183b9381ae226ac3d0a45b4cad484713bd4a28.
//
// Solidity: event RelayWorkersAdded(address indexed relayManager, address[] newRelayWorkers, uint256 workersCount)
func (_IRelayHub *IRelayHubFilterer) FilterRelayWorkersAdded(opts *bind.FilterOpts, relayManager []common.Address) (*IRelayHubRelayWorkersAddedIterator, error) {

	var relayManagerRule []interface{}
	for _, relayManagerItem := range relayManager {
		relayManagerRule = append(relayManagerRule, relayManagerItem)
	}

	logs, sub, err := _IRelayHub.contract.FilterLogs(opts, "RelayWorkersAdded", relayManagerRule)
	if err != nil {
		return nil, err
	}
	return &IRelayHubRelayWorkersAddedIterator{contract: _IRelayHub.contract, event: "RelayWorkersAdded", logs: logs, sub: sub}, nil
}

// WatchRelayWorkersAdded is a free log subscription operation binding the contract event 0xebf4a9bffb39f7c5dbf3f65540183b9381ae226ac3d0a45b4cad484713bd4a28.
//
// Solidity: event RelayWorkersAdded(address indexed relayManager, address[] newRelayWorkers, uint256 workersCount)
func (_IRelayHub *IRelayHubFilterer) WatchRelayWorkersAdded(opts *bind.WatchOpts, sink chan<- *IRelayHubRelayWorkersAdded, relayManager []common.Address) (event.Subscription, error) {

	var relayManagerRule []interface{}
	for _, relayManagerItem := range relayManager {
		relayManagerRule = append(relayManagerRule, relayManagerItem)
	}

	logs, sub, err := _IRelayHub.contract.WatchLogs(opts, "RelayWorkersAdded", relayManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRelayHubRelayWorkersAdded)
				if err := _IRelayHub.contract.UnpackLog(event, "RelayWorkersAdded", log); err != nil {
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

// ParseRelayWorkersAdded is a log parse operation binding the contract event 0xebf4a9bffb39f7c5dbf3f65540183b9381ae226ac3d0a45b4cad484713bd4a28.
//
// Solidity: event RelayWorkersAdded(address indexed relayManager, address[] newRelayWorkers, uint256 workersCount)
func (_IRelayHub *IRelayHubFilterer) ParseRelayWorkersAdded(log types.Log) (*IRelayHubRelayWorkersAdded, error) {
	event := new(IRelayHubRelayWorkersAdded)
	if err := _IRelayHub.contract.UnpackLog(event, "RelayWorkersAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRelayHubTransactionRejectedByPaymasterIterator is returned from FilterTransactionRejectedByPaymaster and is used to iterate over the raw logs and unpacked data for TransactionRejectedByPaymaster events raised by the IRelayHub contract.
type IRelayHubTransactionRejectedByPaymasterIterator struct {
	Event *IRelayHubTransactionRejectedByPaymaster // Event containing the contract specifics and raw log

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
func (it *IRelayHubTransactionRejectedByPaymasterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRelayHubTransactionRejectedByPaymaster)
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
		it.Event = new(IRelayHubTransactionRejectedByPaymaster)
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
func (it *IRelayHubTransactionRejectedByPaymasterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRelayHubTransactionRejectedByPaymasterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRelayHubTransactionRejectedByPaymaster represents a TransactionRejectedByPaymaster event raised by the IRelayHub contract.
type IRelayHubTransactionRejectedByPaymaster struct {
	RelayManager common.Address
	Paymaster    common.Address
	From         common.Address
	To           common.Address
	RelayWorker  common.Address
	Selector     [4]byte
	InnerGasUsed *big.Int
	Reason       []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTransactionRejectedByPaymaster is a free log retrieval operation binding the contract event 0xddb88484d11f800b80fe63aa67488ec56ee001d85896d528912c5d850cbcd06a.
//
// Solidity: event TransactionRejectedByPaymaster(address indexed relayManager, address indexed paymaster, address indexed from, address to, address relayWorker, bytes4 selector, uint256 innerGasUsed, bytes reason)
func (_IRelayHub *IRelayHubFilterer) FilterTransactionRejectedByPaymaster(opts *bind.FilterOpts, relayManager []common.Address, paymaster []common.Address, from []common.Address) (*IRelayHubTransactionRejectedByPaymasterIterator, error) {

	var relayManagerRule []interface{}
	for _, relayManagerItem := range relayManager {
		relayManagerRule = append(relayManagerRule, relayManagerItem)
	}
	var paymasterRule []interface{}
	for _, paymasterItem := range paymaster {
		paymasterRule = append(paymasterRule, paymasterItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IRelayHub.contract.FilterLogs(opts, "TransactionRejectedByPaymaster", relayManagerRule, paymasterRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &IRelayHubTransactionRejectedByPaymasterIterator{contract: _IRelayHub.contract, event: "TransactionRejectedByPaymaster", logs: logs, sub: sub}, nil
}

// WatchTransactionRejectedByPaymaster is a free log subscription operation binding the contract event 0xddb88484d11f800b80fe63aa67488ec56ee001d85896d528912c5d850cbcd06a.
//
// Solidity: event TransactionRejectedByPaymaster(address indexed relayManager, address indexed paymaster, address indexed from, address to, address relayWorker, bytes4 selector, uint256 innerGasUsed, bytes reason)
func (_IRelayHub *IRelayHubFilterer) WatchTransactionRejectedByPaymaster(opts *bind.WatchOpts, sink chan<- *IRelayHubTransactionRejectedByPaymaster, relayManager []common.Address, paymaster []common.Address, from []common.Address) (event.Subscription, error) {

	var relayManagerRule []interface{}
	for _, relayManagerItem := range relayManager {
		relayManagerRule = append(relayManagerRule, relayManagerItem)
	}
	var paymasterRule []interface{}
	for _, paymasterItem := range paymaster {
		paymasterRule = append(paymasterRule, paymasterItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IRelayHub.contract.WatchLogs(opts, "TransactionRejectedByPaymaster", relayManagerRule, paymasterRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRelayHubTransactionRejectedByPaymaster)
				if err := _IRelayHub.contract.UnpackLog(event, "TransactionRejectedByPaymaster", log); err != nil {
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

// ParseTransactionRejectedByPaymaster is a log parse operation binding the contract event 0xddb88484d11f800b80fe63aa67488ec56ee001d85896d528912c5d850cbcd06a.
//
// Solidity: event TransactionRejectedByPaymaster(address indexed relayManager, address indexed paymaster, address indexed from, address to, address relayWorker, bytes4 selector, uint256 innerGasUsed, bytes reason)
func (_IRelayHub *IRelayHubFilterer) ParseTransactionRejectedByPaymaster(log types.Log) (*IRelayHubTransactionRejectedByPaymaster, error) {
	event := new(IRelayHubTransactionRejectedByPaymaster)
	if err := _IRelayHub.contract.UnpackLog(event, "TransactionRejectedByPaymaster", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRelayHubTransactionRelayedIterator is returned from FilterTransactionRelayed and is used to iterate over the raw logs and unpacked data for TransactionRelayed events raised by the IRelayHub contract.
type IRelayHubTransactionRelayedIterator struct {
	Event *IRelayHubTransactionRelayed // Event containing the contract specifics and raw log

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
func (it *IRelayHubTransactionRelayedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRelayHubTransactionRelayed)
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
		it.Event = new(IRelayHubTransactionRelayed)
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
func (it *IRelayHubTransactionRelayedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRelayHubTransactionRelayedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRelayHubTransactionRelayed represents a TransactionRelayed event raised by the IRelayHub contract.
type IRelayHubTransactionRelayed struct {
	RelayManager common.Address
	RelayWorker  common.Address
	From         common.Address
	To           common.Address
	Paymaster    common.Address
	Selector     [4]byte
	Status       uint8
	Charge       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTransactionRelayed is a free log retrieval operation binding the contract event 0xc9aa709786a3d5fe2cc947abc1ba8cbb0f6decb57aa74b84eb7f558125fee454.
//
// Solidity: event TransactionRelayed(address indexed relayManager, address indexed relayWorker, address indexed from, address to, address paymaster, bytes4 selector, uint8 status, uint256 charge)
func (_IRelayHub *IRelayHubFilterer) FilterTransactionRelayed(opts *bind.FilterOpts, relayManager []common.Address, relayWorker []common.Address, from []common.Address) (*IRelayHubTransactionRelayedIterator, error) {

	var relayManagerRule []interface{}
	for _, relayManagerItem := range relayManager {
		relayManagerRule = append(relayManagerRule, relayManagerItem)
	}
	var relayWorkerRule []interface{}
	for _, relayWorkerItem := range relayWorker {
		relayWorkerRule = append(relayWorkerRule, relayWorkerItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IRelayHub.contract.FilterLogs(opts, "TransactionRelayed", relayManagerRule, relayWorkerRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &IRelayHubTransactionRelayedIterator{contract: _IRelayHub.contract, event: "TransactionRelayed", logs: logs, sub: sub}, nil
}

// WatchTransactionRelayed is a free log subscription operation binding the contract event 0xc9aa709786a3d5fe2cc947abc1ba8cbb0f6decb57aa74b84eb7f558125fee454.
//
// Solidity: event TransactionRelayed(address indexed relayManager, address indexed relayWorker, address indexed from, address to, address paymaster, bytes4 selector, uint8 status, uint256 charge)
func (_IRelayHub *IRelayHubFilterer) WatchTransactionRelayed(opts *bind.WatchOpts, sink chan<- *IRelayHubTransactionRelayed, relayManager []common.Address, relayWorker []common.Address, from []common.Address) (event.Subscription, error) {

	var relayManagerRule []interface{}
	for _, relayManagerItem := range relayManager {
		relayManagerRule = append(relayManagerRule, relayManagerItem)
	}
	var relayWorkerRule []interface{}
	for _, relayWorkerItem := range relayWorker {
		relayWorkerRule = append(relayWorkerRule, relayWorkerItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IRelayHub.contract.WatchLogs(opts, "TransactionRelayed", relayManagerRule, relayWorkerRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRelayHubTransactionRelayed)
				if err := _IRelayHub.contract.UnpackLog(event, "TransactionRelayed", log); err != nil {
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

// ParseTransactionRelayed is a log parse operation binding the contract event 0xc9aa709786a3d5fe2cc947abc1ba8cbb0f6decb57aa74b84eb7f558125fee454.
//
// Solidity: event TransactionRelayed(address indexed relayManager, address indexed relayWorker, address indexed from, address to, address paymaster, bytes4 selector, uint8 status, uint256 charge)
func (_IRelayHub *IRelayHubFilterer) ParseTransactionRelayed(log types.Log) (*IRelayHubTransactionRelayed, error) {
	event := new(IRelayHubTransactionRelayed)
	if err := _IRelayHub.contract.UnpackLog(event, "TransactionRelayed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRelayHubTransactionResultIterator is returned from FilterTransactionResult and is used to iterate over the raw logs and unpacked data for TransactionResult events raised by the IRelayHub contract.
type IRelayHubTransactionResultIterator struct {
	Event *IRelayHubTransactionResult // Event containing the contract specifics and raw log

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
func (it *IRelayHubTransactionResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRelayHubTransactionResult)
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
		it.Event = new(IRelayHubTransactionResult)
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
func (it *IRelayHubTransactionResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRelayHubTransactionResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRelayHubTransactionResult represents a TransactionResult event raised by the IRelayHub contract.
type IRelayHubTransactionResult struct {
	Status      uint8
	ReturnValue []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransactionResult is a free log retrieval operation binding the contract event 0xa1478a4242848419db824250a0dddc645dca0d6a9b12ab1fd79b00145a0ba98e.
//
// Solidity: event TransactionResult(uint8 status, bytes returnValue)
func (_IRelayHub *IRelayHubFilterer) FilterTransactionResult(opts *bind.FilterOpts) (*IRelayHubTransactionResultIterator, error) {

	logs, sub, err := _IRelayHub.contract.FilterLogs(opts, "TransactionResult")
	if err != nil {
		return nil, err
	}
	return &IRelayHubTransactionResultIterator{contract: _IRelayHub.contract, event: "TransactionResult", logs: logs, sub: sub}, nil
}

// WatchTransactionResult is a free log subscription operation binding the contract event 0xa1478a4242848419db824250a0dddc645dca0d6a9b12ab1fd79b00145a0ba98e.
//
// Solidity: event TransactionResult(uint8 status, bytes returnValue)
func (_IRelayHub *IRelayHubFilterer) WatchTransactionResult(opts *bind.WatchOpts, sink chan<- *IRelayHubTransactionResult) (event.Subscription, error) {

	logs, sub, err := _IRelayHub.contract.WatchLogs(opts, "TransactionResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRelayHubTransactionResult)
				if err := _IRelayHub.contract.UnpackLog(event, "TransactionResult", log); err != nil {
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

// ParseTransactionResult is a log parse operation binding the contract event 0xa1478a4242848419db824250a0dddc645dca0d6a9b12ab1fd79b00145a0ba98e.
//
// Solidity: event TransactionResult(uint8 status, bytes returnValue)
func (_IRelayHub *IRelayHubFilterer) ParseTransactionResult(log types.Log) (*IRelayHubTransactionResult, error) {
	event := new(IRelayHubTransactionResult)
	if err := _IRelayHub.contract.UnpackLog(event, "TransactionResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRelayHubWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the IRelayHub contract.
type IRelayHubWithdrawnIterator struct {
	Event *IRelayHubWithdrawn // Event containing the contract specifics and raw log

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
func (it *IRelayHubWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRelayHubWithdrawn)
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
		it.Event = new(IRelayHubWithdrawn)
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
func (it *IRelayHubWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRelayHubWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRelayHubWithdrawn represents a Withdrawn event raised by the IRelayHub contract.
type IRelayHubWithdrawn struct {
	Account common.Address
	Dest    common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed account, address indexed dest, uint256 amount)
func (_IRelayHub *IRelayHubFilterer) FilterWithdrawn(opts *bind.FilterOpts, account []common.Address, dest []common.Address) (*IRelayHubWithdrawnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var destRule []interface{}
	for _, destItem := range dest {
		destRule = append(destRule, destItem)
	}

	logs, sub, err := _IRelayHub.contract.FilterLogs(opts, "Withdrawn", accountRule, destRule)
	if err != nil {
		return nil, err
	}
	return &IRelayHubWithdrawnIterator{contract: _IRelayHub.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed account, address indexed dest, uint256 amount)
func (_IRelayHub *IRelayHubFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *IRelayHubWithdrawn, account []common.Address, dest []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var destRule []interface{}
	for _, destItem := range dest {
		destRule = append(destRule, destItem)
	}

	logs, sub, err := _IRelayHub.contract.WatchLogs(opts, "Withdrawn", accountRule, destRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRelayHubWithdrawn)
				if err := _IRelayHub.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed account, address indexed dest, uint256 amount)
func (_IRelayHub *IRelayHubFilterer) ParseWithdrawn(log types.Log) (*IRelayHubWithdrawn, error) {
	event := new(IRelayHubWithdrawn)
	if err := _IRelayHub.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
