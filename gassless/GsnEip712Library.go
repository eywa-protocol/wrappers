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

// GsnEip712LibraryABI is the input ABI used to generate the binding from.
const GsnEip712LibraryABI = "[{\"inputs\":[],\"name\":\"EIP712DOMAIN_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GENERIC_PARAMS\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYDATA_TYPE\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYDATA_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAY_REQUEST_NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAY_REQUEST_SUFFIX\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAY_REQUEST_TYPE\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAY_REQUEST_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// GsnEip712LibraryBin is the compiled bytecode used for deploying new contracts.
var GsnEip712LibraryBin = "0x61053d61003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100925760003560e01c8063abf0d3f411610065578063abf0d3f4146100f0578063c46cf83f14610106578063c49f91d31461010e578063cc0c62b21461013557600080fd5b8063066a310c146100975780636225e61b146100b5578063931cd38f146100bd578063987673f7146100c5575b600080fd5b61009f61013d565b6040516100ac9190610354565b60405180910390f35b61009f610159565b61009f6101f1565b61009f6040518060400160405280600c81526020016b14995b185e54995c5d595cdd60a21b81525081565b6100f861021a565b6040519081526020016100ac565b61009f6102b9565b6100f87f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f81565b6100f86102d5565b6040518060800160405280605d815260200161040a605d913981565b6040518060400160405280600c81526020016b14995b185e54995c5d595cdd60a21b8152506040518060800160405280605d815260200161040a605d91396040518060e0016040528060a1815260200161046760a191396040516020016101c0919061036e565b60408051601f19818403018152908290526101df9392916020016103aa565b60405160208183030381529060405281565b6040518060e0016040528060a1815260200161046760a191396040516020016101df919061036e565b6040518060400160405280600c81526020016b14995b185e54995c5d595cdd60a21b8152506040518060800160405280605d815260200161040a605d91396040518060e0016040528060a1815260200161046760a19139604051602001610281919061036e565b60408051601f19818403018152908290526102a09392916020016103aa565b6040516020818303038152906040528051906020012081565b6040518060e0016040528060a1815260200161046760a1913981565b6040518060e0016040528060a1815260200161046760a191398051906020012081565b60005b838110156103135781810151838201526020016102fb565b83811115610322576000848401525b50505050565b600081518084526103408160208601602086016102f8565b601f01601f19169290920160200192915050565b6020815260006103676020830184610328565b9392505050565b7352656c6179446174612072656c6179446174612960601b81526000825161039d8160148501602087016102f8565b9190910160140192915050565b600084516103bc8184602089016102f8565b600560fb1b90830190815284516103da8160018401602089016102f8565b600b60fa1b6001929091019182015283516103fc8160028401602088016102f8565b016002019594505050505056fe616464726573732066726f6d2c6164647265737320746f2c75696e743235362076616c75652c75696e74323536206761732c75696e74323536206e6f6e63652c627974657320646174612c75696e743235362076616c6964556e74696c52656c6179446174612875696e743235362067617350726963652c75696e743235362070637452656c61794665652c75696e74323536206261736552656c61794665652c616464726573732072656c6179576f726b65722c61646472657373207061796d61737465722c6164647265737320666f727761726465722c6279746573207061796d6173746572446174612c75696e7432353620636c69656e74496429a264697066735822122064f22092d681473d5fb53edfcbb89e709826789b31a541586635dfcc108bdde264736f6c634300080a0033"

// DeployGsnEip712Library deploys a new Ethereum contract, binding an instance of GsnEip712Library to it.
func DeployGsnEip712Library(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GsnEip712Library, error) {
	parsed, err := abi.JSON(strings.NewReader(GsnEip712LibraryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GsnEip712LibraryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GsnEip712Library{GsnEip712LibraryCaller: GsnEip712LibraryCaller{contract: contract}, GsnEip712LibraryTransactor: GsnEip712LibraryTransactor{contract: contract}, GsnEip712LibraryFilterer: GsnEip712LibraryFilterer{contract: contract}}, nil
}

// GsnEip712Library is an auto generated Go binding around an Ethereum contract.
type GsnEip712Library struct {
	GsnEip712LibraryCaller     // Read-only binding to the contract
	GsnEip712LibraryTransactor // Write-only binding to the contract
	GsnEip712LibraryFilterer   // Log filterer for contract events
}

// GsnEip712LibraryCaller is an auto generated read-only Go binding around an Ethereum contract.
type GsnEip712LibraryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GsnEip712LibraryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GsnEip712LibraryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GsnEip712LibraryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GsnEip712LibraryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GsnEip712LibrarySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GsnEip712LibrarySession struct {
	Contract     *GsnEip712Library // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GsnEip712LibraryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GsnEip712LibraryCallerSession struct {
	Contract *GsnEip712LibraryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// GsnEip712LibraryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GsnEip712LibraryTransactorSession struct {
	Contract     *GsnEip712LibraryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// GsnEip712LibraryRaw is an auto generated low-level Go binding around an Ethereum contract.
type GsnEip712LibraryRaw struct {
	Contract *GsnEip712Library // Generic contract binding to access the raw methods on
}

// GsnEip712LibraryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GsnEip712LibraryCallerRaw struct {
	Contract *GsnEip712LibraryCaller // Generic read-only contract binding to access the raw methods on
}

// GsnEip712LibraryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GsnEip712LibraryTransactorRaw struct {
	Contract *GsnEip712LibraryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGsnEip712Library creates a new instance of GsnEip712Library, bound to a specific deployed contract.
func NewGsnEip712Library(address common.Address, backend bind.ContractBackend) (*GsnEip712Library, error) {
	contract, err := bindGsnEip712Library(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GsnEip712Library{GsnEip712LibraryCaller: GsnEip712LibraryCaller{contract: contract}, GsnEip712LibraryTransactor: GsnEip712LibraryTransactor{contract: contract}, GsnEip712LibraryFilterer: GsnEip712LibraryFilterer{contract: contract}}, nil
}

// NewGsnEip712LibraryCaller creates a new read-only instance of GsnEip712Library, bound to a specific deployed contract.
func NewGsnEip712LibraryCaller(address common.Address, caller bind.ContractCaller) (*GsnEip712LibraryCaller, error) {
	contract, err := bindGsnEip712Library(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GsnEip712LibraryCaller{contract: contract}, nil
}

// NewGsnEip712LibraryTransactor creates a new write-only instance of GsnEip712Library, bound to a specific deployed contract.
func NewGsnEip712LibraryTransactor(address common.Address, transactor bind.ContractTransactor) (*GsnEip712LibraryTransactor, error) {
	contract, err := bindGsnEip712Library(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GsnEip712LibraryTransactor{contract: contract}, nil
}

// NewGsnEip712LibraryFilterer creates a new log filterer instance of GsnEip712Library, bound to a specific deployed contract.
func NewGsnEip712LibraryFilterer(address common.Address, filterer bind.ContractFilterer) (*GsnEip712LibraryFilterer, error) {
	contract, err := bindGsnEip712Library(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GsnEip712LibraryFilterer{contract: contract}, nil
}

// bindGsnEip712Library binds a generic wrapper to an already deployed contract.
func bindGsnEip712Library(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GsnEip712LibraryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GsnEip712Library *GsnEip712LibraryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GsnEip712Library.Contract.GsnEip712LibraryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GsnEip712Library *GsnEip712LibraryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GsnEip712Library.Contract.GsnEip712LibraryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GsnEip712Library *GsnEip712LibraryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GsnEip712Library.Contract.GsnEip712LibraryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GsnEip712Library *GsnEip712LibraryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GsnEip712Library.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GsnEip712Library *GsnEip712LibraryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GsnEip712Library.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GsnEip712Library *GsnEip712LibraryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GsnEip712Library.Contract.contract.Transact(opts, method, params...)
}

// EIP712DOMAINTYPEHASH is a free data retrieval call binding the contract method 0xc49f91d3.
//
// Solidity: function EIP712DOMAIN_TYPEHASH() view returns(bytes32)
func (_GsnEip712Library *GsnEip712LibraryCaller) EIP712DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GsnEip712Library.contract.Call(opts, &out, "EIP712DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EIP712DOMAINTYPEHASH is a free data retrieval call binding the contract method 0xc49f91d3.
//
// Solidity: function EIP712DOMAIN_TYPEHASH() view returns(bytes32)
func (_GsnEip712Library *GsnEip712LibrarySession) EIP712DOMAINTYPEHASH() ([32]byte, error) {
	return _GsnEip712Library.Contract.EIP712DOMAINTYPEHASH(&_GsnEip712Library.CallOpts)
}

// EIP712DOMAINTYPEHASH is a free data retrieval call binding the contract method 0xc49f91d3.
//
// Solidity: function EIP712DOMAIN_TYPEHASH() view returns(bytes32)
func (_GsnEip712Library *GsnEip712LibraryCallerSession) EIP712DOMAINTYPEHASH() ([32]byte, error) {
	return _GsnEip712Library.Contract.EIP712DOMAINTYPEHASH(&_GsnEip712Library.CallOpts)
}

// GENERICPARAMS is a free data retrieval call binding the contract method 0x066a310c.
//
// Solidity: function GENERIC_PARAMS() view returns(string)
func (_GsnEip712Library *GsnEip712LibraryCaller) GENERICPARAMS(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GsnEip712Library.contract.Call(opts, &out, "GENERIC_PARAMS")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GENERICPARAMS is a free data retrieval call binding the contract method 0x066a310c.
//
// Solidity: function GENERIC_PARAMS() view returns(string)
func (_GsnEip712Library *GsnEip712LibrarySession) GENERICPARAMS() (string, error) {
	return _GsnEip712Library.Contract.GENERICPARAMS(&_GsnEip712Library.CallOpts)
}

// GENERICPARAMS is a free data retrieval call binding the contract method 0x066a310c.
//
// Solidity: function GENERIC_PARAMS() view returns(string)
func (_GsnEip712Library *GsnEip712LibraryCallerSession) GENERICPARAMS() (string, error) {
	return _GsnEip712Library.Contract.GENERICPARAMS(&_GsnEip712Library.CallOpts)
}

// RELAYDATATYPE is a free data retrieval call binding the contract method 0xc46cf83f.
//
// Solidity: function RELAYDATA_TYPE() view returns(bytes)
func (_GsnEip712Library *GsnEip712LibraryCaller) RELAYDATATYPE(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _GsnEip712Library.contract.Call(opts, &out, "RELAYDATA_TYPE")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// RELAYDATATYPE is a free data retrieval call binding the contract method 0xc46cf83f.
//
// Solidity: function RELAYDATA_TYPE() view returns(bytes)
func (_GsnEip712Library *GsnEip712LibrarySession) RELAYDATATYPE() ([]byte, error) {
	return _GsnEip712Library.Contract.RELAYDATATYPE(&_GsnEip712Library.CallOpts)
}

// RELAYDATATYPE is a free data retrieval call binding the contract method 0xc46cf83f.
//
// Solidity: function RELAYDATA_TYPE() view returns(bytes)
func (_GsnEip712Library *GsnEip712LibraryCallerSession) RELAYDATATYPE() ([]byte, error) {
	return _GsnEip712Library.Contract.RELAYDATATYPE(&_GsnEip712Library.CallOpts)
}

// RELAYDATATYPEHASH is a free data retrieval call binding the contract method 0xcc0c62b2.
//
// Solidity: function RELAYDATA_TYPEHASH() view returns(bytes32)
func (_GsnEip712Library *GsnEip712LibraryCaller) RELAYDATATYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GsnEip712Library.contract.Call(opts, &out, "RELAYDATA_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RELAYDATATYPEHASH is a free data retrieval call binding the contract method 0xcc0c62b2.
//
// Solidity: function RELAYDATA_TYPEHASH() view returns(bytes32)
func (_GsnEip712Library *GsnEip712LibrarySession) RELAYDATATYPEHASH() ([32]byte, error) {
	return _GsnEip712Library.Contract.RELAYDATATYPEHASH(&_GsnEip712Library.CallOpts)
}

// RELAYDATATYPEHASH is a free data retrieval call binding the contract method 0xcc0c62b2.
//
// Solidity: function RELAYDATA_TYPEHASH() view returns(bytes32)
func (_GsnEip712Library *GsnEip712LibraryCallerSession) RELAYDATATYPEHASH() ([32]byte, error) {
	return _GsnEip712Library.Contract.RELAYDATATYPEHASH(&_GsnEip712Library.CallOpts)
}

// RELAYREQUESTNAME is a free data retrieval call binding the contract method 0x987673f7.
//
// Solidity: function RELAY_REQUEST_NAME() view returns(string)
func (_GsnEip712Library *GsnEip712LibraryCaller) RELAYREQUESTNAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GsnEip712Library.contract.Call(opts, &out, "RELAY_REQUEST_NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// RELAYREQUESTNAME is a free data retrieval call binding the contract method 0x987673f7.
//
// Solidity: function RELAY_REQUEST_NAME() view returns(string)
func (_GsnEip712Library *GsnEip712LibrarySession) RELAYREQUESTNAME() (string, error) {
	return _GsnEip712Library.Contract.RELAYREQUESTNAME(&_GsnEip712Library.CallOpts)
}

// RELAYREQUESTNAME is a free data retrieval call binding the contract method 0x987673f7.
//
// Solidity: function RELAY_REQUEST_NAME() view returns(string)
func (_GsnEip712Library *GsnEip712LibraryCallerSession) RELAYREQUESTNAME() (string, error) {
	return _GsnEip712Library.Contract.RELAYREQUESTNAME(&_GsnEip712Library.CallOpts)
}

// RELAYREQUESTSUFFIX is a free data retrieval call binding the contract method 0x931cd38f.
//
// Solidity: function RELAY_REQUEST_SUFFIX() view returns(string)
func (_GsnEip712Library *GsnEip712LibraryCaller) RELAYREQUESTSUFFIX(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GsnEip712Library.contract.Call(opts, &out, "RELAY_REQUEST_SUFFIX")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// RELAYREQUESTSUFFIX is a free data retrieval call binding the contract method 0x931cd38f.
//
// Solidity: function RELAY_REQUEST_SUFFIX() view returns(string)
func (_GsnEip712Library *GsnEip712LibrarySession) RELAYREQUESTSUFFIX() (string, error) {
	return _GsnEip712Library.Contract.RELAYREQUESTSUFFIX(&_GsnEip712Library.CallOpts)
}

// RELAYREQUESTSUFFIX is a free data retrieval call binding the contract method 0x931cd38f.
//
// Solidity: function RELAY_REQUEST_SUFFIX() view returns(string)
func (_GsnEip712Library *GsnEip712LibraryCallerSession) RELAYREQUESTSUFFIX() (string, error) {
	return _GsnEip712Library.Contract.RELAYREQUESTSUFFIX(&_GsnEip712Library.CallOpts)
}

// RELAYREQUESTTYPE is a free data retrieval call binding the contract method 0x6225e61b.
//
// Solidity: function RELAY_REQUEST_TYPE() view returns(bytes)
func (_GsnEip712Library *GsnEip712LibraryCaller) RELAYREQUESTTYPE(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _GsnEip712Library.contract.Call(opts, &out, "RELAY_REQUEST_TYPE")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// RELAYREQUESTTYPE is a free data retrieval call binding the contract method 0x6225e61b.
//
// Solidity: function RELAY_REQUEST_TYPE() view returns(bytes)
func (_GsnEip712Library *GsnEip712LibrarySession) RELAYREQUESTTYPE() ([]byte, error) {
	return _GsnEip712Library.Contract.RELAYREQUESTTYPE(&_GsnEip712Library.CallOpts)
}

// RELAYREQUESTTYPE is a free data retrieval call binding the contract method 0x6225e61b.
//
// Solidity: function RELAY_REQUEST_TYPE() view returns(bytes)
func (_GsnEip712Library *GsnEip712LibraryCallerSession) RELAYREQUESTTYPE() ([]byte, error) {
	return _GsnEip712Library.Contract.RELAYREQUESTTYPE(&_GsnEip712Library.CallOpts)
}

// RELAYREQUESTTYPEHASH is a free data retrieval call binding the contract method 0xabf0d3f4.
//
// Solidity: function RELAY_REQUEST_TYPEHASH() view returns(bytes32)
func (_GsnEip712Library *GsnEip712LibraryCaller) RELAYREQUESTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GsnEip712Library.contract.Call(opts, &out, "RELAY_REQUEST_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RELAYREQUESTTYPEHASH is a free data retrieval call binding the contract method 0xabf0d3f4.
//
// Solidity: function RELAY_REQUEST_TYPEHASH() view returns(bytes32)
func (_GsnEip712Library *GsnEip712LibrarySession) RELAYREQUESTTYPEHASH() ([32]byte, error) {
	return _GsnEip712Library.Contract.RELAYREQUESTTYPEHASH(&_GsnEip712Library.CallOpts)
}

// RELAYREQUESTTYPEHASH is a free data retrieval call binding the contract method 0xabf0d3f4.
//
// Solidity: function RELAY_REQUEST_TYPEHASH() view returns(bytes32)
func (_GsnEip712Library *GsnEip712LibraryCallerSession) RELAYREQUESTTYPEHASH() ([32]byte, error) {
	return _GsnEip712Library.Contract.RELAYREQUESTTYPEHASH(&_GsnEip712Library.CallOpts)
}
