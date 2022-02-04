// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wrappers

import (
	"errors"
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
)

// TestTargetMetaData contains all meta data concerning the TestTarget contract.
var TestTargetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"bytesData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setTestAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_inputBytes\",\"type\":\"bytes\"}],\"name\":\"setTestBytes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_inputString\",\"type\":\"string\"}],\"name\":\"setTestString\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_testUint\",\"type\":\"uint256\"}],\"name\":\"setTestUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stringData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tesAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610582806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063656687001161005b5780636566870014610126578063b3a0eb7d1461013b578063e3850c3d14610143578063eea237681461014c57600080fd5b806309b4699d1461008d578063315e2f1b146100b65780634aec35ed146100cb578063584f0228146100f6575b600080fd5b6100a361009b36600461032f565b600081905590565b6040519081526020015b60405180910390f35b6100c96100c4366004610348565b61015f565b005b6001546100de906001600160a01b031681565b6040516001600160a01b0390911681526020016100ad565b6100de6101043660046103ba565b600180546001600160a01b039092166001600160a01b03199092168217905590565b61012e610170565b6040516100ad9190610437565b61012e6101fe565b6100a360005481565b6100c961015a366004610460565b61020b565b61016b60038383610222565b505050565b6003805461017d90610511565b80601f01602080910402602001604051908101604052809291908181526020018280546101a990610511565b80156101f65780601f106101cb576101008083540402835291602001916101f6565b820191906000526020600020905b8154815290600101906020018083116101d957829003601f168201915b505050505081565b6002805461017d90610511565b805161021e9060029060208401906102a6565b5050565b82805461022e90610511565b90600052602060002090601f0160209004810192826102505760008555610296565b82601f106102695782800160ff19823516178555610296565b82800160010185558215610296579182015b8281111561029657823582559160200191906001019061027b565b506102a292915061031a565b5090565b8280546102b290610511565b90600052602060002090601f0160209004810192826102d45760008555610296565b82601f106102ed57805160ff1916838001178555610296565b82800160010185558215610296579182015b828111156102965782518255916020019190600101906102ff565b5b808211156102a2576000815560010161031b565b60006020828403121561034157600080fd5b5035919050565b6000806020838503121561035b57600080fd5b823567ffffffffffffffff8082111561037357600080fd5b818501915085601f83011261038757600080fd5b81358181111561039657600080fd5b8660208285010111156103a857600080fd5b60209290920196919550909350505050565b6000602082840312156103cc57600080fd5b81356001600160a01b03811681146103e357600080fd5b9392505050565b6000815180845260005b81811015610410576020818501810151868301820152016103f4565b81811115610422576000602083870101525b50601f01601f19169290920160200192915050565b6020815260006103e360208301846103ea565b634e487b7160e01b600052604160045260246000fd5b60006020828403121561047257600080fd5b813567ffffffffffffffff8082111561048a57600080fd5b818401915084601f83011261049e57600080fd5b8135818111156104b0576104b061044a565b604051601f8201601f19908116603f011681019083821181831017156104d8576104d861044a565b816040528281528760208487010111156104f157600080fd5b826020860160208301376000928101602001929092525095945050505050565b600181811c9082168061052557607f821691505b6020821081141561054657634e487b7160e01b600052602260045260246000fd5b5091905056fea26469706673582212200adf7bb27605f00f46890f08a1ce5aa3b3be0dc92e7b5f4dea854f37827ed0c964736f6c634300080a0033",
}

// TestTargetABI is the input ABI used to generate the binding from.
// Deprecated: Use TestTargetMetaData.ABI instead.
var TestTargetABI = TestTargetMetaData.ABI

// TestTargetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestTargetMetaData.Bin instead.
var TestTargetBin = TestTargetMetaData.Bin

// DeployTestTarget deploys a new Ethereum contract, binding an instance of TestTarget to it.
func DeployTestTarget(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestTarget, error) {
	parsed, err := TestTargetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestTargetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestTarget{TestTargetCaller: TestTargetCaller{contract: contract}, TestTargetTransactor: TestTargetTransactor{contract: contract}, TestTargetFilterer: TestTargetFilterer{contract: contract}}, nil
}

// TestTarget is an auto generated Go binding around an Ethereum contract.
type TestTarget struct {
	TestTargetCaller     // Read-only binding to the contract
	TestTargetTransactor // Write-only binding to the contract
	TestTargetFilterer   // Log filterer for contract events
}

// TestTargetCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestTargetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestTargetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestTargetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
	gsn      *GsnCallOpts
}

func (_TestTarget *TestTargetTransactor) SetGSNOptions(opts *GsnCallOpts) {
	_TestTarget.gsn = opts
}

// TestTargetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestTargetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestTargetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestTargetSession struct {
	Contract     *TestTarget       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestTargetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestTargetCallerSession struct {
	Contract *TestTargetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TestTargetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestTargetTransactorSession struct {
	Contract     *TestTargetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TestTargetRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestTargetRaw struct {
	Contract *TestTarget // Generic contract binding to access the raw methods on
}

// TestTargetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestTargetCallerRaw struct {
	Contract *TestTargetCaller // Generic read-only contract binding to access the raw methods on
}

// TestTargetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestTargetTransactorRaw struct {
	Contract *TestTargetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestTarget creates a new instance of TestTarget, bound to a specific deployed contract.
func NewTestTarget(address common.Address, backend bind.ContractBackend) (*TestTarget, error) {
	contract, err := bindTestTarget(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestTarget{TestTargetCaller: TestTargetCaller{contract: contract}, TestTargetTransactor: TestTargetTransactor{contract: contract}, TestTargetFilterer: TestTargetFilterer{contract: contract}}, nil
}

// NewTestTargetCaller creates a new read-only instance of TestTarget, bound to a specific deployed contract.
func NewTestTargetCaller(address common.Address, caller bind.ContractCaller) (*TestTargetCaller, error) {
	contract, err := bindTestTarget(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestTargetCaller{contract: contract}, nil
}

// NewTestTargetTransactor creates a new write-only instance of TestTarget, bound to a specific deployed contract.
func NewTestTargetTransactor(address common.Address, transactor bind.ContractTransactor) (*TestTargetTransactor, error) {
	contract, err := bindTestTarget(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestTargetTransactor{contract: contract}, nil
}

// NewTestTargetFilterer creates a new log filterer instance of TestTarget, bound to a specific deployed contract.
func NewTestTargetFilterer(address common.Address, filterer bind.ContractFilterer) (*TestTargetFilterer, error) {
	contract, err := bindTestTarget(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestTargetFilterer{contract: contract}, nil
}

// bindTestTarget binds a generic wrapper to an already deployed contract.
func bindTestTarget(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestTargetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestTarget *TestTargetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestTarget.Contract.TestTargetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestTarget *TestTargetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestTarget.Contract.TestTargetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestTarget *TestTargetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestTarget.Contract.TestTargetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestTarget *TestTargetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestTarget.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestTarget *TestTargetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestTarget.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestTarget *TestTargetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestTarget.Contract.contract.Transact(opts, method, params...)
}

// BytesData is a free data retrieval call binding the contract method 0xb3a0eb7d.
//
// Solidity: function bytesData() view returns(bytes)
func (_TestTarget *TestTargetCaller) BytesData(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _TestTarget.contract.Call(opts, &out, "bytesData")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// BytesData is a free data retrieval call binding the contract method 0xb3a0eb7d.
//
// Solidity: function bytesData() view returns(bytes)
func (_TestTarget *TestTargetSession) BytesData() ([]byte, error) {
	return _TestTarget.Contract.BytesData(&_TestTarget.CallOpts)
}

// BytesData is a free data retrieval call binding the contract method 0xb3a0eb7d.
//
// Solidity: function bytesData() view returns(bytes)
func (_TestTarget *TestTargetCallerSession) BytesData() ([]byte, error) {
	return _TestTarget.Contract.BytesData(&_TestTarget.CallOpts)
}

// StringData is a free data retrieval call binding the contract method 0x65668700.
//
// Solidity: function stringData() view returns(string)
func (_TestTarget *TestTargetCaller) StringData(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestTarget.contract.Call(opts, &out, "stringData")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// StringData is a free data retrieval call binding the contract method 0x65668700.
//
// Solidity: function stringData() view returns(string)
func (_TestTarget *TestTargetSession) StringData() (string, error) {
	return _TestTarget.Contract.StringData(&_TestTarget.CallOpts)
}

// StringData is a free data retrieval call binding the contract method 0x65668700.
//
// Solidity: function stringData() view returns(string)
func (_TestTarget *TestTargetCallerSession) StringData() (string, error) {
	return _TestTarget.Contract.StringData(&_TestTarget.CallOpts)
}

// TesAddress is a free data retrieval call binding the contract method 0x4aec35ed.
//
// Solidity: function tesAddress() view returns(address)
func (_TestTarget *TestTargetCaller) TesAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestTarget.contract.Call(opts, &out, "tesAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TesAddress is a free data retrieval call binding the contract method 0x4aec35ed.
//
// Solidity: function tesAddress() view returns(address)
func (_TestTarget *TestTargetSession) TesAddress() (common.Address, error) {
	return _TestTarget.Contract.TesAddress(&_TestTarget.CallOpts)
}

// TesAddress is a free data retrieval call binding the contract method 0x4aec35ed.
//
// Solidity: function tesAddress() view returns(address)
func (_TestTarget *TestTargetCallerSession) TesAddress() (common.Address, error) {
	return _TestTarget.Contract.TesAddress(&_TestTarget.CallOpts)
}

// TestUint is a free data retrieval call binding the contract method 0xe3850c3d.
//
// Solidity: function testUint() view returns(uint256)
func (_TestTarget *TestTargetCaller) TestUint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestTarget.contract.Call(opts, &out, "testUint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TestUint is a free data retrieval call binding the contract method 0xe3850c3d.
//
// Solidity: function testUint() view returns(uint256)
func (_TestTarget *TestTargetSession) TestUint() (*big.Int, error) {
	return _TestTarget.Contract.TestUint(&_TestTarget.CallOpts)
}

// TestUint is a free data retrieval call binding the contract method 0xe3850c3d.
//
// Solidity: function testUint() view returns(uint256)
func (_TestTarget *TestTargetCallerSession) TestUint() (*big.Int, error) {
	return _TestTarget.Contract.TestUint(&_TestTarget.CallOpts)
}

// SetTestAddress is a paid mutator transaction binding the contract method 0x584f0228.
//
// Solidity: function setTestAddress(address _address) returns(address)
func (_TestTarget *TestTargetTransactor) SetTestAddress(opts *bind.TransactOpts, _address common.Address) (common.Hash, error) {
	return GsnWrap(
		_TestTarget.gsn,
		func() (common.Hash, error) {
			tx, errIn := _TestTarget.contract.Transact(opts, "setTestAddress", _address)
			if tx != nil {
				return tx.Hash(), errIn
			}
			return common.Hash{}, errIn
		},
		func() (common.Hash, error) {
			return GsnExecutor(_TestTarget.gsn, TestTargetMetaData.ABI, "setTestAddress", _address)
		},
	)
}

// SetTestAddress is a paid mutator transaction binding the contract method 0x584f0228.
//
// Solidity: function setTestAddress(address _address) returns(address)
func (_TestTarget *TestTargetSession) SetTestAddress(_address common.Address) (common.Hash, error) {
	return _TestTarget.Contract.SetTestAddress(&_TestTarget.TransactOpts, _address)
}

// SetTestAddress is a paid mutator transaction binding the contract method 0x584f0228.
//
// Solidity: function setTestAddress(address _address) returns(address)
func (_TestTarget *TestTargetTransactorSession) SetTestAddress(_address common.Address) (common.Hash, error) {
	return _TestTarget.Contract.SetTestAddress(&_TestTarget.TransactOpts, _address)
}

// SetTestBytes is a paid mutator transaction binding the contract method 0xeea23768.
//
// Solidity: function setTestBytes(bytes _inputBytes) returns()
func (_TestTarget *TestTargetTransactor) SetTestBytes(opts *bind.TransactOpts, _inputBytes []byte) (common.Hash, error) {
	return GsnWrap(
		_TestTarget.gsn,
		func() (common.Hash, error) {
			tx, errIn := _TestTarget.contract.Transact(opts, "setTestBytes", _inputBytes)
			if tx != nil {
				return tx.Hash(), errIn
			}
			return common.Hash{}, errIn
		},
		func() (common.Hash, error) {
			return GsnExecutor(_TestTarget.gsn, TestTargetMetaData.ABI, "setTestBytes", _inputBytes)
		},
	)
}

// SetTestBytes is a paid mutator transaction binding the contract method 0xeea23768.
//
// Solidity: function setTestBytes(bytes _inputBytes) returns()
func (_TestTarget *TestTargetSession) SetTestBytes(_inputBytes []byte) (common.Hash, error) {
	return _TestTarget.Contract.SetTestBytes(&_TestTarget.TransactOpts, _inputBytes)
}

// SetTestBytes is a paid mutator transaction binding the contract method 0xeea23768.
//
// Solidity: function setTestBytes(bytes _inputBytes) returns()
func (_TestTarget *TestTargetTransactorSession) SetTestBytes(_inputBytes []byte) (common.Hash, error) {
	return _TestTarget.Contract.SetTestBytes(&_TestTarget.TransactOpts, _inputBytes)
}

// SetTestString is a paid mutator transaction binding the contract method 0x315e2f1b.
//
// Solidity: function setTestString(string _inputString) returns()
func (_TestTarget *TestTargetTransactor) SetTestString(opts *bind.TransactOpts, _inputString string) (common.Hash, error) {
	return GsnWrap(
		_TestTarget.gsn,
		func() (common.Hash, error) {
			tx, errIn := _TestTarget.contract.Transact(opts, "setTestString", _inputString)
			if tx != nil {
				return tx.Hash(), errIn
			}
			return common.Hash{}, errIn
		},
		func() (common.Hash, error) {
			return GsnExecutor(_TestTarget.gsn, TestTargetMetaData.ABI, "setTestString", _inputString)
		},
	)
}

// SetTestString is a paid mutator transaction binding the contract method 0x315e2f1b.
//
// Solidity: function setTestString(string _inputString) returns()
func (_TestTarget *TestTargetSession) SetTestString(_inputString string) (common.Hash, error) {
	return _TestTarget.Contract.SetTestString(&_TestTarget.TransactOpts, _inputString)
}

// SetTestString is a paid mutator transaction binding the contract method 0x315e2f1b.
//
// Solidity: function setTestString(string _inputString) returns()
func (_TestTarget *TestTargetTransactorSession) SetTestString(_inputString string) (common.Hash, error) {
	return _TestTarget.Contract.SetTestString(&_TestTarget.TransactOpts, _inputString)
}

// SetTestUint is a paid mutator transaction binding the contract method 0x09b4699d.
//
// Solidity: function setTestUint(uint256 _testUint) returns(uint256)
func (_TestTarget *TestTargetTransactor) SetTestUint(opts *bind.TransactOpts, _testUint *big.Int) (common.Hash, error) {
	return GsnWrap(
		_TestTarget.gsn,
		func() (common.Hash, error) {
			tx, errIn := _TestTarget.contract.Transact(opts, "setTestUint", _testUint)
			if tx != nil {
				return tx.Hash(), errIn
			}
			return common.Hash{}, errIn
		},
		func() (common.Hash, error) {
			return GsnExecutor(_TestTarget.gsn, TestTargetMetaData.ABI, "setTestUint", _testUint)
		},
	)
}

// SetTestUint is a paid mutator transaction binding the contract method 0x09b4699d.
//
// Solidity: function setTestUint(uint256 _testUint) returns(uint256)
func (_TestTarget *TestTargetSession) SetTestUint(_testUint *big.Int) (common.Hash, error) {
	return _TestTarget.Contract.SetTestUint(&_TestTarget.TransactOpts, _testUint)
}

// SetTestUint is a paid mutator transaction binding the contract method 0x09b4699d.
//
// Solidity: function setTestUint(uint256 _testUint) returns(uint256)
func (_TestTarget *TestTargetTransactorSession) SetTestUint(_testUint *big.Int) (common.Hash, error) {
	return _TestTarget.Contract.SetTestUint(&_TestTarget.TransactOpts, _testUint)
}
