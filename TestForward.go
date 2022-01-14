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

// TestForwardABI is the input ABI used to generate the binding from.
const TestForwardABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_forwarder\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"FooCalled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_val\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_str\",\"type\":\"string\"}],\"name\":\"foo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"str\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedForwarder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"val\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionRecipient\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TestForwardBin is the compiled bytecode used for deploying new contracts.
var TestForwardBin = "0x60018055600280546001600160a01b031916905560c0604052600c60808190526b48656c6c6f20776f726c642160a01b60a090815261004191600491906100cd565b5034801561004e57600080fd5b5060405161075938038061075983398101604081905261006d91610166565b6001600160a01b03811661009c5760405162461bcd60e51b815260040161009390610194565b60405180910390fd5b6100a5816100ab565b506101f5565b600080546001600160a01b0319166001600160a01b0392909216919091179055565b8280546100d9906101ba565b90600052602060002090601f0160209004810192826100fb5760008555610141565b82601f1061011457805160ff1916838001178555610141565b82800160010185558215610141579182015b82811115610141578251825591602001919060010190610126565b5061014d929150610151565b5090565b5b8082111561014d5760008155600101610152565b600060208284031215610177578081fd5b81516001600160a01b038116811461018d578182fd5b9392505050565b6020808252600c908201526b5a45524f204144445245535360a01b604082015260600190565b6002810460018216806101ce57607f821691505b602082108114156101ef57634e487b7160e01b600052602260045260246000fd5b50919050565b610555806102046000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c806367e404ce1161005b57806367e404ce146100d55780637da0a877146100ea578063c15bae84146100f2578063c5d1c995146100fa5761007d565b80633c6bb43614610082578063486ff0cd146100a0578063572b6c05146100b5575b600080fd5b61008a61010f565b60405161009791906104c5565b60405180910390f35b6100a8610115565b604051610097919061044e565b6100c86100c3366004610354565b6101a3565b6040516100979190610443565b6100dd6101b7565b604051610097919061042f565b6100dd6101c6565b6100a86101d6565b61010d610108366004610382565b6101e3565b005b60015481565b60048054610122906104ce565b80601f016020809104026020016040519081016040528092919081815260200182805461014e906104ce565b801561019b5780601f106101705761010080835404028352916020019161019b565b820191906000526020600020905b81548152906001019060200180831161017e57829003601f168201915b505050505081565b6000546001600160a01b0390811691161490565b6002546001600160a01b031681565b6000546001600160a01b03165b90565b60038054610122906104ce565b816102095760405162461bcd60e51b8152600401610200906104a1565b60405180910390fd5b600182905580516102219060039060208401906102bb565b5061022a61028b565b600280546001600160a01b0319166001600160a01b03928316179081905560015460405191909216917f21d588436ad968dec202ff08608dc590023f4b084bfecdc43ac3ca2fd9922a3a9161027f91906104c5565b60405180910390a25050565b6000601436108015906102a257506102a2336101a3565b156102b6575060131936013560601c6101d3565b503390565b8280546102c7906104ce565b90600052602060002090601f0160209004810192826102e9576000855561032f565b82601f1061030257805160ff191683800117855561032f565b8280016001018555821561032f579182015b8281111561032f578251825591602001919060010190610314565b5061033b92915061033f565b5090565b5b8082111561033b5760008155600101610340565b600060208284031215610365578081fd5b81356001600160a01b038116811461037b578182fd5b9392505050565b60008060408385031215610394578081fd5b8235915060208084013567ffffffffffffffff808211156103b3578384fd5b818601915086601f8301126103c6578384fd5b8135818111156103d8576103d8610509565b604051601f8201601f19168101850183811182821017156103fb576103fb610509565b6040528181528382018501891015610411578586fd5b81858501868301378585838301015280955050505050509250929050565b6001600160a01b0391909116815260200190565b901515815260200190565b6000602080835283518082850152825b8181101561047a5785810183015185820160400152820161045e565b8181111561048b5783604083870101525b50601f01601f1916929092016040019392505050565b6020808252600a90820152695a45524f2056414c554560b01b604082015260600190565b90815260200190565b6002810460018216806104e257607f821691505b6020821081141561050357634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fdfea2646970667358221220ba8171a88ae46aa3981215c31c5bbe35d19c76c67f09eed76f5235b5c7a2410f64736f6c63430008000033"

// DeployTestForward deploys a new Ethereum contract, binding an instance of TestForward to it.
func DeployTestForward(auth *bind.TransactOpts, backend bind.ContractBackend, _forwarder common.Address) (common.Address, *types.Transaction, *TestForward, error) {
	parsed, err := abi.JSON(strings.NewReader(TestForwardABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TestForwardBin), backend, _forwarder)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestForward{TestForwardCaller: TestForwardCaller{contract: contract}, TestForwardTransactor: TestForwardTransactor{contract: contract}, TestForwardFilterer: TestForwardFilterer{contract: contract}}, nil
}

// TestForward is an auto generated Go binding around an Ethereum contract.
type TestForward struct {
	TestForwardCaller     // Read-only binding to the contract
	TestForwardTransactor // Write-only binding to the contract
	TestForwardFilterer   // Log filterer for contract events
}

// TestForwardCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestForwardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestForwardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestForwardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestForwardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestForwardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestForwardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestForwardSession struct {
	Contract     *TestForward      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestForwardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestForwardCallerSession struct {
	Contract *TestForwardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TestForwardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestForwardTransactorSession struct {
	Contract     *TestForwardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TestForwardRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestForwardRaw struct {
	Contract *TestForward // Generic contract binding to access the raw methods on
}

// TestForwardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestForwardCallerRaw struct {
	Contract *TestForwardCaller // Generic read-only contract binding to access the raw methods on
}

// TestForwardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestForwardTransactorRaw struct {
	Contract *TestForwardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestForward creates a new instance of TestForward, bound to a specific deployed contract.
func NewTestForward(address common.Address, backend bind.ContractBackend) (*TestForward, error) {
	contract, err := bindTestForward(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestForward{TestForwardCaller: TestForwardCaller{contract: contract}, TestForwardTransactor: TestForwardTransactor{contract: contract}, TestForwardFilterer: TestForwardFilterer{contract: contract}}, nil
}

// NewTestForwardCaller creates a new read-only instance of TestForward, bound to a specific deployed contract.
func NewTestForwardCaller(address common.Address, caller bind.ContractCaller) (*TestForwardCaller, error) {
	contract, err := bindTestForward(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestForwardCaller{contract: contract}, nil
}

// NewTestForwardTransactor creates a new write-only instance of TestForward, bound to a specific deployed contract.
func NewTestForwardTransactor(address common.Address, transactor bind.ContractTransactor) (*TestForwardTransactor, error) {
	contract, err := bindTestForward(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestForwardTransactor{contract: contract}, nil
}

// NewTestForwardFilterer creates a new log filterer instance of TestForward, bound to a specific deployed contract.
func NewTestForwardFilterer(address common.Address, filterer bind.ContractFilterer) (*TestForwardFilterer, error) {
	contract, err := bindTestForward(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestForwardFilterer{contract: contract}, nil
}

// bindTestForward binds a generic wrapper to an already deployed contract.
func bindTestForward(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestForwardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestForward *TestForwardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestForward.Contract.TestForwardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestForward *TestForwardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestForward.Contract.TestForwardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestForward *TestForwardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestForward.Contract.TestForwardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestForward *TestForwardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestForward.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestForward *TestForwardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestForward.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestForward *TestForwardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestForward.Contract.contract.Transact(opts, method, params...)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_TestForward *TestForwardCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _TestForward.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_TestForward *TestForwardSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _TestForward.Contract.IsTrustedForwarder(&_TestForward.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_TestForward *TestForwardCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _TestForward.Contract.IsTrustedForwarder(&_TestForward.CallOpts, forwarder)
}

// Sender is a free data retrieval call binding the contract method 0x67e404ce.
//
// Solidity: function sender() view returns(address)
func (_TestForward *TestForwardCaller) Sender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestForward.contract.Call(opts, &out, "sender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sender is a free data retrieval call binding the contract method 0x67e404ce.
//
// Solidity: function sender() view returns(address)
func (_TestForward *TestForwardSession) Sender() (common.Address, error) {
	return _TestForward.Contract.Sender(&_TestForward.CallOpts)
}

// Sender is a free data retrieval call binding the contract method 0x67e404ce.
//
// Solidity: function sender() view returns(address)
func (_TestForward *TestForwardCallerSession) Sender() (common.Address, error) {
	return _TestForward.Contract.Sender(&_TestForward.CallOpts)
}

// Str is a free data retrieval call binding the contract method 0xc15bae84.
//
// Solidity: function str() view returns(string)
func (_TestForward *TestForwardCaller) Str(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestForward.contract.Call(opts, &out, "str")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Str is a free data retrieval call binding the contract method 0xc15bae84.
//
// Solidity: function str() view returns(string)
func (_TestForward *TestForwardSession) Str() (string, error) {
	return _TestForward.Contract.Str(&_TestForward.CallOpts)
}

// Str is a free data retrieval call binding the contract method 0xc15bae84.
//
// Solidity: function str() view returns(string)
func (_TestForward *TestForwardCallerSession) Str() (string, error) {
	return _TestForward.Contract.Str(&_TestForward.CallOpts)
}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_TestForward *TestForwardCaller) TrustedForwarder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestForward.contract.Call(opts, &out, "trustedForwarder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_TestForward *TestForwardSession) TrustedForwarder() (common.Address, error) {
	return _TestForward.Contract.TrustedForwarder(&_TestForward.CallOpts)
}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_TestForward *TestForwardCallerSession) TrustedForwarder() (common.Address, error) {
	return _TestForward.Contract.TrustedForwarder(&_TestForward.CallOpts)
}

// Val is a free data retrieval call binding the contract method 0x3c6bb436.
//
// Solidity: function val() view returns(uint256)
func (_TestForward *TestForwardCaller) Val(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestForward.contract.Call(opts, &out, "val")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Val is a free data retrieval call binding the contract method 0x3c6bb436.
//
// Solidity: function val() view returns(uint256)
func (_TestForward *TestForwardSession) Val() (*big.Int, error) {
	return _TestForward.Contract.Val(&_TestForward.CallOpts)
}

// Val is a free data retrieval call binding the contract method 0x3c6bb436.
//
// Solidity: function val() view returns(uint256)
func (_TestForward *TestForwardCallerSession) Val() (*big.Int, error) {
	return _TestForward.Contract.Val(&_TestForward.CallOpts)
}

// VersionRecipient is a free data retrieval call binding the contract method 0x486ff0cd.
//
// Solidity: function versionRecipient() view returns(string)
func (_TestForward *TestForwardCaller) VersionRecipient(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestForward.contract.Call(opts, &out, "versionRecipient")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VersionRecipient is a free data retrieval call binding the contract method 0x486ff0cd.
//
// Solidity: function versionRecipient() view returns(string)
func (_TestForward *TestForwardSession) VersionRecipient() (string, error) {
	return _TestForward.Contract.VersionRecipient(&_TestForward.CallOpts)
}

// VersionRecipient is a free data retrieval call binding the contract method 0x486ff0cd.
//
// Solidity: function versionRecipient() view returns(string)
func (_TestForward *TestForwardCallerSession) VersionRecipient() (string, error) {
	return _TestForward.Contract.VersionRecipient(&_TestForward.CallOpts)
}

// Foo is a paid mutator transaction binding the contract method 0xc5d1c995.
//
// Solidity: function foo(uint256 _val, string _str) returns()
func (_TestForward *TestForwardTransactor) Foo(opts *bind.TransactOpts, _val *big.Int, _str string) (*types.Transaction, error) {
	return _TestForward.contract.Transact(opts, "foo", _val, _str)
}

// Foo is a paid mutator transaction binding the contract method 0xc5d1c995.
//
// Solidity: function foo(uint256 _val, string _str) returns()
func (_TestForward *TestForwardSession) Foo(_val *big.Int, _str string) (*types.Transaction, error) {
	return _TestForward.Contract.Foo(&_TestForward.TransactOpts, _val, _str)
}

// Foo is a paid mutator transaction binding the contract method 0xc5d1c995.
//
// Solidity: function foo(uint256 _val, string _str) returns()
func (_TestForward *TestForwardTransactorSession) Foo(_val *big.Int, _str string) (*types.Transaction, error) {
	return _TestForward.Contract.Foo(&_TestForward.TransactOpts, _val, _str)
}

// TestForwardFooCalledIterator is returned from FilterFooCalled and is used to iterate over the raw logs and unpacked data for FooCalled events raised by the TestForward contract.
type TestForwardFooCalledIterator struct {
	Event *TestForwardFooCalled // Event containing the contract specifics and raw log

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
func (it *TestForwardFooCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestForwardFooCalled)
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
		it.Event = new(TestForwardFooCalled)
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
func (it *TestForwardFooCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestForwardFooCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestForwardFooCalled represents a FooCalled event raised by the TestForward contract.
type TestForwardFooCalled struct {
	Caller common.Address
	Val    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFooCalled is a free log retrieval operation binding the contract event 0x21d588436ad968dec202ff08608dc590023f4b084bfecdc43ac3ca2fd9922a3a.
//
// Solidity: event FooCalled(address indexed caller, uint256 val)
func (_TestForward *TestForwardFilterer) FilterFooCalled(opts *bind.FilterOpts, caller []common.Address) (*TestForwardFooCalledIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _TestForward.contract.FilterLogs(opts, "FooCalled", callerRule)
	if err != nil {
		return nil, err
	}
	return &TestForwardFooCalledIterator{contract: _TestForward.contract, event: "FooCalled", logs: logs, sub: sub}, nil
}

// WatchFooCalled is a free log subscription operation binding the contract event 0x21d588436ad968dec202ff08608dc590023f4b084bfecdc43ac3ca2fd9922a3a.
//
// Solidity: event FooCalled(address indexed caller, uint256 val)
func (_TestForward *TestForwardFilterer) WatchFooCalled(opts *bind.WatchOpts, sink chan<- *TestForwardFooCalled, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _TestForward.contract.WatchLogs(opts, "FooCalled", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestForwardFooCalled)
				if err := _TestForward.contract.UnpackLog(event, "FooCalled", log); err != nil {
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

// ParseFooCalled is a log parse operation binding the contract event 0x21d588436ad968dec202ff08608dc590023f4b084bfecdc43ac3ca2fd9922a3a.
//
// Solidity: event FooCalled(address indexed caller, uint256 val)
func (_TestForward *TestForwardFilterer) ParseFooCalled(log types.Log) (*TestForwardFooCalled, error) {
	event := new(TestForwardFooCalled)
	if err := _TestForward.contract.UnpackLog(event, "FooCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
