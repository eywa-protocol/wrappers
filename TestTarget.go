// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wrappers

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

// TestTargetABI is the input ABI used to generate the binding from.
const TestTargetABI = "[{\"inputs\":[],\"name\":\"bytesData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setTestAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_inputBytes\",\"type\":\"bytes\"}],\"name\":\"setTestBytes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_inputString\",\"type\":\"string\"}],\"name\":\"setTestString\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_testUint\",\"type\":\"uint256\"}],\"name\":\"setTestUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stringData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tesAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TestTargetBin is the compiled bytecode used for deploying new contracts.
var TestTargetBin = "0x608060405234801561001057600080fd5b50610590806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063656687001161005b57806365668700146100f3578063b3a0eb7d14610108578063e3850c3d14610110578063eea237681461011857610088565b806309b4699d1461008d578063315e2f1b146100b65780634aec35ed146100cb578063584f0228146100e0575b600080fd5b6100a061009b366004610476565b61012b565b6040516100ad9190610500565b60405180910390f35b6100c96100c4366004610409565b610133565b005b6100d3610144565b6040516100ad91906104d9565b6100d36100ee36600461033a565b610153565b6100fb610175565b6040516100ad91906104ed565b6100fb610203565b6100a0610210565b6100c9610126366004610368565b610216565b600081905590565b61013f6003838361022d565b505050565b6001546001600160a01b031681565b600180546001600160a01b0319166001600160a01b0392831617908190551690565b6003805461018290610509565b80601f01602080910402602001604051908101604052809291908181526020018280546101ae90610509565b80156101fb5780601f106101d0576101008083540402835291602001916101fb565b820191906000526020600020905b8154815290600101906020018083116101de57829003601f168201915b505050505081565b6002805461018290610509565b60005481565b80516102299060029060208401906102b1565b5050565b82805461023990610509565b90600052602060002090601f01602090048101928261025b57600085556102a1565b82601f106102745782800160ff198235161785556102a1565b828001600101855582156102a1579182015b828111156102a1578235825591602001919060010190610286565b506102ad929150610325565b5090565b8280546102bd90610509565b90600052602060002090601f0160209004810192826102df57600085556102a1565b82601f106102f857805160ff19168380011785556102a1565b828001600101855582156102a1579182015b828111156102a157825182559160200191906001019061030a565b5b808211156102ad5760008155600101610326565b60006020828403121561034b578081fd5b81356001600160a01b0381168114610361578182fd5b9392505050565b6000602080838503121561037a578182fd5b823567ffffffffffffffff80821115610391578384fd5b818501915085601f8301126103a4578384fd5b8135818111156103b6576103b6610544565b604051601f8201601f19168101850183811182821017156103d9576103d9610544565b60405281815283820185018810156103ef578586fd5b818585018683013790810190930193909352509392505050565b6000806020838503121561041b578081fd5b823567ffffffffffffffff80821115610432578283fd5b818501915085601f830112610445578283fd5b813581811115610453578384fd5b866020828501011115610464578384fd5b60209290920196919550909350505050565b600060208284031215610487578081fd5b5035919050565b60008151808452815b818110156104b357602081850181015186830182015201610497565b818111156104c45782602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b0391909116815260200190565b600060208252610361602083018461048e565b90815260200190565b60028104600182168061051d57607f821691505b6020821081141561053e57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fdfea26469706673582212200c427e7ff00507470b643a8df3c02f2090cbd335fe6b473856b815059220f0b064736f6c63430008000033"

// DeployTestTarget deploys a new Ethereum contract, binding an instance of TestTarget to it.
func DeployTestTarget(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestTarget, error) {
	parsed, err := abi.JSON(strings.NewReader(TestTargetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TestTargetBin), backend)
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
func (_TestTarget *TestTargetTransactor) SetTestAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _TestTarget.contract.Transact(opts, "setTestAddress", _address)
}

// SetTestAddress is a paid mutator transaction binding the contract method 0x584f0228.
//
// Solidity: function setTestAddress(address _address) returns(address)
func (_TestTarget *TestTargetSession) SetTestAddress(_address common.Address) (*types.Transaction, error) {
	return _TestTarget.Contract.SetTestAddress(&_TestTarget.TransactOpts, _address)
}

// SetTestAddress is a paid mutator transaction binding the contract method 0x584f0228.
//
// Solidity: function setTestAddress(address _address) returns(address)
func (_TestTarget *TestTargetTransactorSession) SetTestAddress(_address common.Address) (*types.Transaction, error) {
	return _TestTarget.Contract.SetTestAddress(&_TestTarget.TransactOpts, _address)
}

// SetTestBytes is a paid mutator transaction binding the contract method 0xeea23768.
//
// Solidity: function setTestBytes(bytes _inputBytes) returns()
func (_TestTarget *TestTargetTransactor) SetTestBytes(opts *bind.TransactOpts, _inputBytes []byte) (*types.Transaction, error) {
	return _TestTarget.contract.Transact(opts, "setTestBytes", _inputBytes)
}

// SetTestBytes is a paid mutator transaction binding the contract method 0xeea23768.
//
// Solidity: function setTestBytes(bytes _inputBytes) returns()
func (_TestTarget *TestTargetSession) SetTestBytes(_inputBytes []byte) (*types.Transaction, error) {
	return _TestTarget.Contract.SetTestBytes(&_TestTarget.TransactOpts, _inputBytes)
}

// SetTestBytes is a paid mutator transaction binding the contract method 0xeea23768.
//
// Solidity: function setTestBytes(bytes _inputBytes) returns()
func (_TestTarget *TestTargetTransactorSession) SetTestBytes(_inputBytes []byte) (*types.Transaction, error) {
	return _TestTarget.Contract.SetTestBytes(&_TestTarget.TransactOpts, _inputBytes)
}

// SetTestString is a paid mutator transaction binding the contract method 0x315e2f1b.
//
// Solidity: function setTestString(string _inputString) returns()
func (_TestTarget *TestTargetTransactor) SetTestString(opts *bind.TransactOpts, _inputString string) (*types.Transaction, error) {
	return _TestTarget.contract.Transact(opts, "setTestString", _inputString)
}

// SetTestString is a paid mutator transaction binding the contract method 0x315e2f1b.
//
// Solidity: function setTestString(string _inputString) returns()
func (_TestTarget *TestTargetSession) SetTestString(_inputString string) (*types.Transaction, error) {
	return _TestTarget.Contract.SetTestString(&_TestTarget.TransactOpts, _inputString)
}

// SetTestString is a paid mutator transaction binding the contract method 0x315e2f1b.
//
// Solidity: function setTestString(string _inputString) returns()
func (_TestTarget *TestTargetTransactorSession) SetTestString(_inputString string) (*types.Transaction, error) {
	return _TestTarget.Contract.SetTestString(&_TestTarget.TransactOpts, _inputString)
}

// SetTestUint is a paid mutator transaction binding the contract method 0x09b4699d.
//
// Solidity: function setTestUint(uint256 _testUint) returns(uint256)
func (_TestTarget *TestTargetTransactor) SetTestUint(opts *bind.TransactOpts, _testUint *big.Int) (*types.Transaction, error) {
	return _TestTarget.contract.Transact(opts, "setTestUint", _testUint)
}

// SetTestUint is a paid mutator transaction binding the contract method 0x09b4699d.
//
// Solidity: function setTestUint(uint256 _testUint) returns(uint256)
func (_TestTarget *TestTargetSession) SetTestUint(_testUint *big.Int) (*types.Transaction, error) {
	return _TestTarget.Contract.SetTestUint(&_TestTarget.TransactOpts, _testUint)
}

// SetTestUint is a paid mutator transaction binding the contract method 0x09b4699d.
//
// Solidity: function setTestUint(uint256 _testUint) returns(uint256)
func (_TestTarget *TestTargetTransactorSession) SetTestUint(_testUint *big.Int) (*types.Transaction, error) {
	return _TestTarget.Contract.SetTestUint(&_TestTarget.TransactOpts, _testUint)
}
