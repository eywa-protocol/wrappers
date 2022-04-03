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

// ErrorsMetaData contains all meta data concerning the Errors contract.
var ErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DATA_INCONSISTENCY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEPOSIT_IS_LOCKED\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EMISSION_ANNUAL_RATE_IS_TOO_HIGH\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EMISSION_ANNUAL_RATE_IS_TOO_LOW\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_IS_TOO_HIGH\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_IS_TOO_LOW\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INSUFFICIENT_DEPOSIT\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NOT_DEPOSIT_OWNER\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SAME_VALUE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ZERO_PROFIT\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x61036e61003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100b35760003560e01c8063529186761161007b57806352918676146101c4578063538ba4f9146101f15780638e8af4f91461021c578063a57c37fc1461024f578063cba6fe231461028b578063da461352146102b557600080fd5b806312aad774146100b857806313a5e9db146100fe5780631646dec31461012757806325ce8936146101635780634b7767d414610194575b600080fd5b6100e8604051806040016040528060118152602001702727aa2fa222a827a9a4aa2fa7aba722a960791b81525081565b6040516100f591906102e3565b60405180910390f35b6100e86040518060400160405280600a81526020016953414d455f56414c554560b01b81525081565b6100e86040518060400160405280602081526020017f454d495353494f4e5f414e4e55414c5f524154455f49535f544f4f5f4849474881525081565b6100e860405180604001604052806012815260200171444154415f494e434f4e53495354454e435960701b81525081565b6100e86040518060400160405280601181526020017011115413d4d25517d254d7d313d0d2d151607a1b81525081565b6100e86040518060400160405280600e81526020016d4645455f49535f544f4f5f4c4f5760901b81525081565b6100e86040518060400160405280600c81526020016b5a45524f5f4144445245535360a01b81525081565b6100e860405180604001604052806014815260200173125394d551919250d251539517d1115413d4d25560621b81525081565b6100e86040518060400160405280601f81526020017f454d495353494f4e5f414e4e55414c5f524154455f49535f544f4f5f4c4f570081525081565b6100e86040518060400160405280600b81526020016a16915493d7d41493d1925560aa1b81525081565b6100e86040518060400160405280600f81526020016e08c8a8abe92a6bea89e9ebe90928e9608b1b81525081565b600060208083528351808285015260005b81811015610310578581018301518582016040015282016102f4565b81811115610322576000604083870101525b50601f01601f191692909201604001939250505056fea2646970667358221220f9ac670baea2c12018ea541d92cb226e923014188409d9ae18eb659d57a8e64b64736f6c634300080a0033",
}

// ErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use ErrorsMetaData.ABI instead.
var ErrorsABI = ErrorsMetaData.ABI

// ErrorsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ErrorsMetaData.Bin instead.
var ErrorsBin = ErrorsMetaData.Bin

// DeployErrors deploys a new Ethereum contract, binding an instance of Errors to it.
func DeployErrors(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Errors, error) {
	parsed, err := ErrorsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ErrorsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Errors{ErrorsCaller: ErrorsCaller{contract: contract}, ErrorsTransactor: ErrorsTransactor{contract: contract}, ErrorsFilterer: ErrorsFilterer{contract: contract}}, nil
}

// Errors is an auto generated Go binding around an Ethereum contract.
type Errors struct {
	ErrorsCaller     // Read-only binding to the contract
	ErrorsTransactor // Write-only binding to the contract
	ErrorsFilterer   // Log filterer for contract events
}

// ErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
	gsn      *GsnCallOpts
}

func (_Errors *ErrorsTransactor) SetGSNOptions(opts *GsnCallOpts) {
	_Errors.gsn = opts
}

// ErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ErrorsSession struct {
	Contract     *Errors           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ErrorsCallerSession struct {
	Contract *ErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ErrorsTransactorSession struct {
	Contract     *ErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ErrorsRaw struct {
	Contract *Errors // Generic contract binding to access the raw methods on
}

// ErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ErrorsCallerRaw struct {
	Contract *ErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// ErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ErrorsTransactorRaw struct {
	Contract *ErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErrors creates a new instance of Errors, bound to a specific deployed contract.
func NewErrors(address common.Address, backend bind.ContractBackend) (*Errors, error) {
	contract, err := bindErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Errors{ErrorsCaller: ErrorsCaller{contract: contract}, ErrorsTransactor: ErrorsTransactor{contract: contract}, ErrorsFilterer: ErrorsFilterer{contract: contract}}, nil
}

// NewErrorsCaller creates a new read-only instance of Errors, bound to a specific deployed contract.
func NewErrorsCaller(address common.Address, caller bind.ContractCaller) (*ErrorsCaller, error) {
	contract, err := bindErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ErrorsCaller{contract: contract}, nil
}

// NewErrorsTransactor creates a new write-only instance of Errors, bound to a specific deployed contract.
func NewErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*ErrorsTransactor, error) {
	contract, err := bindErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ErrorsTransactor{contract: contract}, nil
}

// NewErrorsFilterer creates a new log filterer instance of Errors, bound to a specific deployed contract.
func NewErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*ErrorsFilterer, error) {
	contract, err := bindErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ErrorsFilterer{contract: contract}, nil
}

// bindErrors binds a generic wrapper to an already deployed contract.
func bindErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ErrorsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Errors *ErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Errors.Contract.ErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Errors *ErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Errors.Contract.ErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Errors *ErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Errors.Contract.ErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Errors *ErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Errors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Errors *ErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Errors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Errors *ErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Errors.Contract.contract.Transact(opts, method, params...)
}

// DATAINCONSISTENCY is a free data retrieval call binding the contract method 0x25ce8936.
//
// Solidity: function DATA_INCONSISTENCY() view returns(string)
func (_Errors *ErrorsCaller) DATAINCONSISTENCY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "DATA_INCONSISTENCY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DATAINCONSISTENCY is a free data retrieval call binding the contract method 0x25ce8936.
//
// Solidity: function DATA_INCONSISTENCY() view returns(string)
func (_Errors *ErrorsSession) DATAINCONSISTENCY() (string, error) {
	return _Errors.Contract.DATAINCONSISTENCY(&_Errors.CallOpts)
}

// DATAINCONSISTENCY is a free data retrieval call binding the contract method 0x25ce8936.
//
// Solidity: function DATA_INCONSISTENCY() view returns(string)
func (_Errors *ErrorsCallerSession) DATAINCONSISTENCY() (string, error) {
	return _Errors.Contract.DATAINCONSISTENCY(&_Errors.CallOpts)
}

// DEPOSITISLOCKED is a free data retrieval call binding the contract method 0x4b7767d4.
//
// Solidity: function DEPOSIT_IS_LOCKED() view returns(string)
func (_Errors *ErrorsCaller) DEPOSITISLOCKED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "DEPOSIT_IS_LOCKED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DEPOSITISLOCKED is a free data retrieval call binding the contract method 0x4b7767d4.
//
// Solidity: function DEPOSIT_IS_LOCKED() view returns(string)
func (_Errors *ErrorsSession) DEPOSITISLOCKED() (string, error) {
	return _Errors.Contract.DEPOSITISLOCKED(&_Errors.CallOpts)
}

// DEPOSITISLOCKED is a free data retrieval call binding the contract method 0x4b7767d4.
//
// Solidity: function DEPOSIT_IS_LOCKED() view returns(string)
func (_Errors *ErrorsCallerSession) DEPOSITISLOCKED() (string, error) {
	return _Errors.Contract.DEPOSITISLOCKED(&_Errors.CallOpts)
}

// EMISSIONANNUALRATEISTOOHIGH is a free data retrieval call binding the contract method 0x1646dec3.
//
// Solidity: function EMISSION_ANNUAL_RATE_IS_TOO_HIGH() view returns(string)
func (_Errors *ErrorsCaller) EMISSIONANNUALRATEISTOOHIGH(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "EMISSION_ANNUAL_RATE_IS_TOO_HIGH")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EMISSIONANNUALRATEISTOOHIGH is a free data retrieval call binding the contract method 0x1646dec3.
//
// Solidity: function EMISSION_ANNUAL_RATE_IS_TOO_HIGH() view returns(string)
func (_Errors *ErrorsSession) EMISSIONANNUALRATEISTOOHIGH() (string, error) {
	return _Errors.Contract.EMISSIONANNUALRATEISTOOHIGH(&_Errors.CallOpts)
}

// EMISSIONANNUALRATEISTOOHIGH is a free data retrieval call binding the contract method 0x1646dec3.
//
// Solidity: function EMISSION_ANNUAL_RATE_IS_TOO_HIGH() view returns(string)
func (_Errors *ErrorsCallerSession) EMISSIONANNUALRATEISTOOHIGH() (string, error) {
	return _Errors.Contract.EMISSIONANNUALRATEISTOOHIGH(&_Errors.CallOpts)
}

// EMISSIONANNUALRATEISTOOLOW is a free data retrieval call binding the contract method 0xa57c37fc.
//
// Solidity: function EMISSION_ANNUAL_RATE_IS_TOO_LOW() view returns(string)
func (_Errors *ErrorsCaller) EMISSIONANNUALRATEISTOOLOW(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "EMISSION_ANNUAL_RATE_IS_TOO_LOW")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EMISSIONANNUALRATEISTOOLOW is a free data retrieval call binding the contract method 0xa57c37fc.
//
// Solidity: function EMISSION_ANNUAL_RATE_IS_TOO_LOW() view returns(string)
func (_Errors *ErrorsSession) EMISSIONANNUALRATEISTOOLOW() (string, error) {
	return _Errors.Contract.EMISSIONANNUALRATEISTOOLOW(&_Errors.CallOpts)
}

// EMISSIONANNUALRATEISTOOLOW is a free data retrieval call binding the contract method 0xa57c37fc.
//
// Solidity: function EMISSION_ANNUAL_RATE_IS_TOO_LOW() view returns(string)
func (_Errors *ErrorsCallerSession) EMISSIONANNUALRATEISTOOLOW() (string, error) {
	return _Errors.Contract.EMISSIONANNUALRATEISTOOLOW(&_Errors.CallOpts)
}

// FEEISTOOHIGH is a free data retrieval call binding the contract method 0xda461352.
//
// Solidity: function FEE_IS_TOO_HIGH() view returns(string)
func (_Errors *ErrorsCaller) FEEISTOOHIGH(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "FEE_IS_TOO_HIGH")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// FEEISTOOHIGH is a free data retrieval call binding the contract method 0xda461352.
//
// Solidity: function FEE_IS_TOO_HIGH() view returns(string)
func (_Errors *ErrorsSession) FEEISTOOHIGH() (string, error) {
	return _Errors.Contract.FEEISTOOHIGH(&_Errors.CallOpts)
}

// FEEISTOOHIGH is a free data retrieval call binding the contract method 0xda461352.
//
// Solidity: function FEE_IS_TOO_HIGH() view returns(string)
func (_Errors *ErrorsCallerSession) FEEISTOOHIGH() (string, error) {
	return _Errors.Contract.FEEISTOOHIGH(&_Errors.CallOpts)
}

// FEEISTOOLOW is a free data retrieval call binding the contract method 0x52918676.
//
// Solidity: function FEE_IS_TOO_LOW() view returns(string)
func (_Errors *ErrorsCaller) FEEISTOOLOW(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "FEE_IS_TOO_LOW")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// FEEISTOOLOW is a free data retrieval call binding the contract method 0x52918676.
//
// Solidity: function FEE_IS_TOO_LOW() view returns(string)
func (_Errors *ErrorsSession) FEEISTOOLOW() (string, error) {
	return _Errors.Contract.FEEISTOOLOW(&_Errors.CallOpts)
}

// FEEISTOOLOW is a free data retrieval call binding the contract method 0x52918676.
//
// Solidity: function FEE_IS_TOO_LOW() view returns(string)
func (_Errors *ErrorsCallerSession) FEEISTOOLOW() (string, error) {
	return _Errors.Contract.FEEISTOOLOW(&_Errors.CallOpts)
}

// INSUFFICIENTDEPOSIT is a free data retrieval call binding the contract method 0x8e8af4f9.
//
// Solidity: function INSUFFICIENT_DEPOSIT() view returns(string)
func (_Errors *ErrorsCaller) INSUFFICIENTDEPOSIT(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "INSUFFICIENT_DEPOSIT")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// INSUFFICIENTDEPOSIT is a free data retrieval call binding the contract method 0x8e8af4f9.
//
// Solidity: function INSUFFICIENT_DEPOSIT() view returns(string)
func (_Errors *ErrorsSession) INSUFFICIENTDEPOSIT() (string, error) {
	return _Errors.Contract.INSUFFICIENTDEPOSIT(&_Errors.CallOpts)
}

// INSUFFICIENTDEPOSIT is a free data retrieval call binding the contract method 0x8e8af4f9.
//
// Solidity: function INSUFFICIENT_DEPOSIT() view returns(string)
func (_Errors *ErrorsCallerSession) INSUFFICIENTDEPOSIT() (string, error) {
	return _Errors.Contract.INSUFFICIENTDEPOSIT(&_Errors.CallOpts)
}

// NOTDEPOSITOWNER is a free data retrieval call binding the contract method 0x12aad774.
//
// Solidity: function NOT_DEPOSIT_OWNER() view returns(string)
func (_Errors *ErrorsCaller) NOTDEPOSITOWNER(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "NOT_DEPOSIT_OWNER")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NOTDEPOSITOWNER is a free data retrieval call binding the contract method 0x12aad774.
//
// Solidity: function NOT_DEPOSIT_OWNER() view returns(string)
func (_Errors *ErrorsSession) NOTDEPOSITOWNER() (string, error) {
	return _Errors.Contract.NOTDEPOSITOWNER(&_Errors.CallOpts)
}

// NOTDEPOSITOWNER is a free data retrieval call binding the contract method 0x12aad774.
//
// Solidity: function NOT_DEPOSIT_OWNER() view returns(string)
func (_Errors *ErrorsCallerSession) NOTDEPOSITOWNER() (string, error) {
	return _Errors.Contract.NOTDEPOSITOWNER(&_Errors.CallOpts)
}

// SAMEVALUE is a free data retrieval call binding the contract method 0x13a5e9db.
//
// Solidity: function SAME_VALUE() view returns(string)
func (_Errors *ErrorsCaller) SAMEVALUE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "SAME_VALUE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SAMEVALUE is a free data retrieval call binding the contract method 0x13a5e9db.
//
// Solidity: function SAME_VALUE() view returns(string)
func (_Errors *ErrorsSession) SAMEVALUE() (string, error) {
	return _Errors.Contract.SAMEVALUE(&_Errors.CallOpts)
}

// SAMEVALUE is a free data retrieval call binding the contract method 0x13a5e9db.
//
// Solidity: function SAME_VALUE() view returns(string)
func (_Errors *ErrorsCallerSession) SAMEVALUE() (string, error) {
	return _Errors.Contract.SAMEVALUE(&_Errors.CallOpts)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() view returns(string)
func (_Errors *ErrorsCaller) ZEROADDRESS(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "ZERO_ADDRESS")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() view returns(string)
func (_Errors *ErrorsSession) ZEROADDRESS() (string, error) {
	return _Errors.Contract.ZEROADDRESS(&_Errors.CallOpts)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() view returns(string)
func (_Errors *ErrorsCallerSession) ZEROADDRESS() (string, error) {
	return _Errors.Contract.ZEROADDRESS(&_Errors.CallOpts)
}

// ZEROPROFIT is a free data retrieval call binding the contract method 0xcba6fe23.
//
// Solidity: function ZERO_PROFIT() view returns(string)
func (_Errors *ErrorsCaller) ZEROPROFIT(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "ZERO_PROFIT")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ZEROPROFIT is a free data retrieval call binding the contract method 0xcba6fe23.
//
// Solidity: function ZERO_PROFIT() view returns(string)
func (_Errors *ErrorsSession) ZEROPROFIT() (string, error) {
	return _Errors.Contract.ZEROPROFIT(&_Errors.CallOpts)
}

// ZEROPROFIT is a free data retrieval call binding the contract method 0xcba6fe23.
//
// Solidity: function ZERO_PROFIT() view returns(string)
func (_Errors *ErrorsCallerSession) ZEROPROFIT() (string, error) {
	return _Errors.Contract.ZEROPROFIT(&_Errors.CallOpts)
}
