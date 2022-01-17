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
var TestForwardBin = "0x60018055600280546001600160a01b031916905560c0604052600c60808190526b48656c6c6f20776f726c642160a01b60a090815261004191600491906100d7565b5034801561004e57600080fd5b5060405161070f38038061070f83398101604081905261006d91610170565b6001600160a01b0381166100b65760405162461bcd60e51b815260206004820152600c60248201526b5a45524f204144445245535360a01b604482015260640160405180910390fd5b600080546001600160a01b0319166001600160a01b038316179055506101db565b8280546100e3906101a0565b90600052602060002090601f016020900481019282610105576000855561014b565b82601f1061011e57805160ff191683800117855561014b565b8280016001018555821561014b579182015b8281111561014b578251825591602001919060010190610130565b5061015792915061015b565b5090565b5b80821115610157576000815560010161015c565b60006020828403121561018257600080fd5b81516001600160a01b038116811461019957600080fd5b9392505050565b600181811c908216806101b457607f821691505b602082108114156101d557634e487b7160e01b600052602260045260246000fd5b50919050565b610525806101ea6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c806367e404ce1161005b57806367e404ce146100e55780637da0a87714610110578063c15bae8414610121578063c5d1c9951461012957600080fd5b80633c6bb43614610082578063486ff0cd1461009e578063572b6c05146100b3575b600080fd5b61008b60015481565b6040519081526020015b60405180910390f35b6100a661013e565b604051610095919061035e565b6100d56100c13660046103b3565b6000546001600160a01b0391821691161490565b6040519015158152602001610095565b6002546100f8906001600160a01b031681565b6040516001600160a01b039091168152602001610095565b6000546001600160a01b03166100f8565b6100a66101cc565b61013c6101373660046103f9565b6101d9565b005b6004805461014b906104b4565b80601f0160208091040260200160405190810160405280929190818152602001828054610177906104b4565b80156101c45780601f10610199576101008083540402835291602001916101c4565b820191906000526020600020905b8154815290600101906020018083116101a757829003601f168201915b505050505081565b6003805461014b906104b4565b816102175760405162461bcd60e51b815260206004820152600a6024820152695a45524f2056414c554560b01b604482015260640160405180910390fd5b6001829055805161022f9060039060208401906102c5565b50610238610291565b600280546001600160a01b0319166001600160a01b039290921691821790556001546040519081527f21d588436ad968dec202ff08608dc590023f4b084bfecdc43ac3ca2fd9922a3a9060200160405180910390a25050565b6000601436108015906102ae57506000546001600160a01b031633145b156102c0575060131936013560601c90565b503390565b8280546102d1906104b4565b90600052602060002090601f0160209004810192826102f35760008555610339565b82601f1061030c57805160ff1916838001178555610339565b82800160010185558215610339579182015b8281111561033957825182559160200191906001019061031e565b50610345929150610349565b5090565b5b80821115610345576000815560010161034a565b600060208083528351808285015260005b8181101561038b5785810183015185820160400152820161036f565b8181111561039d576000604083870101525b50601f01601f1916929092016040019392505050565b6000602082840312156103c557600080fd5b81356001600160a01b03811681146103dc57600080fd5b9392505050565b634e487b7160e01b600052604160045260246000fd5b6000806040838503121561040c57600080fd5b82359150602083013567ffffffffffffffff8082111561042b57600080fd5b818501915085601f83011261043f57600080fd5b813581811115610451576104516103e3565b604051601f8201601f19908116603f01168101908382118183101715610479576104796103e3565b8160405282815288602084870101111561049257600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b600181811c908216806104c857607f821691505b602082108114156104e957634e487b7160e01b600052602260045260246000fd5b5091905056fea264697066735822122039d21a847eb52a5ef801d5b9d6a2a7d19144ee41f4e23873950fd06b145d8fc164736f6c634300080a0033"

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
