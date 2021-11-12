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

// MockDexPoolABI is the input ABI used to generate the binding from.
const MockDexPoolABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"reqId\",\"type\":\"bytes32\"}],\"name\":\"RequestSended\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_testData\",\"type\":\"uint256\"}],\"name\":\"receiveRequestTest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"testData_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"secondPartPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"sendRequestTestV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MockDexPoolBin is the compiled bytecode used for deploying new contracts.
var MockDexPoolBin = "0x60806040526000805534801561001457600080fd5b506040516104d13803806104d183398101604081905261003391610058565b600180546001600160a01b0319166001600160a01b0392909216919091179055610086565b600060208284031215610069578081fd5b81516001600160a01b038116811461007f578182fd5b9392505050565b61043c806100956000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063016cbd5114610051578063e78cea921461006f578063f16f710314610084578063f9ee520e14610099575b600080fd5b6100596100ac565b6040516100669190610335565b60405180910390f35b6100776100b2565b6040516100669190610321565b6100976100923660046102b9565b6100c1565b005b6100976100a73660046102d1565b6100f9565b60005481565b6001546001600160a01b031681565b6001546001600160a01b031633146100f45760405162461bcd60e51b81526004016100eb906103b4565b60405180910390fd5b600055565b6001600160a01b03831661011f5760405162461bcd60e51b81526004016100eb906103e1565b604080518082018252601b81527f7265636569766552657175657374546573742875696e74323536290000000000602090910152516000907ff16f710300ade8663e80c8937af0bceb0204f46c88990d0957b94e12b18138de90610187908790602401610335565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909416939093179092526001549151633675e4e160e11b81529092506000916001600160a01b031690636cebc9c2906101f290859089908990899060040161033e565b602060405180830381600087803b15801561020c57600080fd5b505af1158015610220573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061024491906102a1565b90507fa67b89382ffb8a1a6987e2a4c7a66d2383ed74b29abbf2a8b9d2adac76b0d2c8816040516102759190610335565b60405180910390a1505050505050565b80356001600160a01b038116811461029c57600080fd5b919050565b6000602082840312156102b2578081fd5b5051919050565b6000602082840312156102ca578081fd5b5035919050565b600080600080608085870312156102e6578283fd5b843593506102f660208601610285565b925061030460408601610285565b9396929550929360600135925050565b6001600160a01b03169052565b6001600160a01b0391909116815260200190565b90815260200190565b6000608082528551806080840152815b8181101561036b57602081890181015160a086840101520161034e565b8181111561037c578260a083860101525b50601f01601f1916820160a00190506103986020830186610314565b6103a56040830185610314565b82606083015295945050505050565b6020808252601390820152724f4e4c59204345525441494e2042524944474560681b604082015260600190565b6020808252600b908201526a424144204144445245535360a81b60408201526060019056fea2646970667358221220ab15b4c36a0a3e2c44055dd12988dc565930d14052d4337fd8f3c2b6a04a591964736f6c63430008000033"

// DeployMockDexPool deploys a new Ethereum contract, binding an instance of MockDexPool to it.
func DeployMockDexPool(auth *bind.TransactOpts, backend bind.ContractBackend, _bridge common.Address) (common.Address, *types.Transaction, *MockDexPool, error) {
	parsed, err := abi.JSON(strings.NewReader(MockDexPoolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MockDexPoolBin), backend, _bridge)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockDexPool{MockDexPoolCaller: MockDexPoolCaller{contract: contract}, MockDexPoolTransactor: MockDexPoolTransactor{contract: contract}, MockDexPoolFilterer: MockDexPoolFilterer{contract: contract}}, nil
}

// MockDexPool is an auto generated Go binding around an Ethereum contract.
type MockDexPool struct {
	MockDexPoolCaller     // Read-only binding to the contract
	MockDexPoolTransactor // Write-only binding to the contract
	MockDexPoolFilterer   // Log filterer for contract events
}

// MockDexPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockDexPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockDexPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockDexPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockDexPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockDexPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockDexPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockDexPoolSession struct {
	Contract     *MockDexPool      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockDexPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockDexPoolCallerSession struct {
	Contract *MockDexPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MockDexPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockDexPoolTransactorSession struct {
	Contract     *MockDexPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MockDexPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockDexPoolRaw struct {
	Contract *MockDexPool // Generic contract binding to access the raw methods on
}

// MockDexPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockDexPoolCallerRaw struct {
	Contract *MockDexPoolCaller // Generic read-only contract binding to access the raw methods on
}

// MockDexPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockDexPoolTransactorRaw struct {
	Contract *MockDexPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockDexPool creates a new instance of MockDexPool, bound to a specific deployed contract.
func NewMockDexPool(address common.Address, backend bind.ContractBackend) (*MockDexPool, error) {
	contract, err := bindMockDexPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockDexPool{MockDexPoolCaller: MockDexPoolCaller{contract: contract}, MockDexPoolTransactor: MockDexPoolTransactor{contract: contract}, MockDexPoolFilterer: MockDexPoolFilterer{contract: contract}}, nil
}

// NewMockDexPoolCaller creates a new read-only instance of MockDexPool, bound to a specific deployed contract.
func NewMockDexPoolCaller(address common.Address, caller bind.ContractCaller) (*MockDexPoolCaller, error) {
	contract, err := bindMockDexPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockDexPoolCaller{contract: contract}, nil
}

// NewMockDexPoolTransactor creates a new write-only instance of MockDexPool, bound to a specific deployed contract.
func NewMockDexPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*MockDexPoolTransactor, error) {
	contract, err := bindMockDexPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockDexPoolTransactor{contract: contract}, nil
}

// NewMockDexPoolFilterer creates a new log filterer instance of MockDexPool, bound to a specific deployed contract.
func NewMockDexPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*MockDexPoolFilterer, error) {
	contract, err := bindMockDexPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockDexPoolFilterer{contract: contract}, nil
}

// bindMockDexPool binds a generic wrapper to an already deployed contract.
func bindMockDexPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MockDexPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockDexPool *MockDexPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockDexPool.Contract.MockDexPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockDexPool *MockDexPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockDexPool.Contract.MockDexPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockDexPool *MockDexPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockDexPool.Contract.MockDexPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockDexPool *MockDexPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockDexPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockDexPool *MockDexPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockDexPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockDexPool *MockDexPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockDexPool.Contract.contract.Transact(opts, method, params...)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_MockDexPool *MockDexPoolCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockDexPool.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_MockDexPool *MockDexPoolSession) Bridge() (common.Address, error) {
	return _MockDexPool.Contract.Bridge(&_MockDexPool.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_MockDexPool *MockDexPoolCallerSession) Bridge() (common.Address, error) {
	return _MockDexPool.Contract.Bridge(&_MockDexPool.CallOpts)
}

// TestData is a free data retrieval call binding the contract method 0x016cbd51.
//
// Solidity: function testData() view returns(uint256)
func (_MockDexPool *MockDexPoolCaller) TestData(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockDexPool.contract.Call(opts, &out, "testData")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TestData is a free data retrieval call binding the contract method 0x016cbd51.
//
// Solidity: function testData() view returns(uint256)
func (_MockDexPool *MockDexPoolSession) TestData() (*big.Int, error) {
	return _MockDexPool.Contract.TestData(&_MockDexPool.CallOpts)
}

// TestData is a free data retrieval call binding the contract method 0x016cbd51.
//
// Solidity: function testData() view returns(uint256)
func (_MockDexPool *MockDexPoolCallerSession) TestData() (*big.Int, error) {
	return _MockDexPool.Contract.TestData(&_MockDexPool.CallOpts)
}

// ReceiveRequestTest is a paid mutator transaction binding the contract method 0xf16f7103.
//
// Solidity: function receiveRequestTest(uint256 _testData) returns()
func (_MockDexPool *MockDexPoolTransactor) ReceiveRequestTest(opts *bind.TransactOpts, _testData *big.Int) (*types.Transaction, error) {
	return _MockDexPool.contract.Transact(opts, "receiveRequestTest", _testData)
}

// ReceiveRequestTest is a paid mutator transaction binding the contract method 0xf16f7103.
//
// Solidity: function receiveRequestTest(uint256 _testData) returns()
func (_MockDexPool *MockDexPoolSession) ReceiveRequestTest(_testData *big.Int) (*types.Transaction, error) {
	return _MockDexPool.Contract.ReceiveRequestTest(&_MockDexPool.TransactOpts, _testData)
}

// ReceiveRequestTest is a paid mutator transaction binding the contract method 0xf16f7103.
//
// Solidity: function receiveRequestTest(uint256 _testData) returns()
func (_MockDexPool *MockDexPoolTransactorSession) ReceiveRequestTest(_testData *big.Int) (*types.Transaction, error) {
	return _MockDexPool.Contract.ReceiveRequestTest(&_MockDexPool.TransactOpts, _testData)
}

// SendRequestTestV2 is a paid mutator transaction binding the contract method 0xf9ee520e.
//
// Solidity: function sendRequestTestV2(uint256 testData_, address secondPartPool, address oppBridge, uint256 chainId) returns()
func (_MockDexPool *MockDexPoolTransactor) SendRequestTestV2(opts *bind.TransactOpts, testData_ *big.Int, secondPartPool common.Address, oppBridge common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _MockDexPool.contract.Transact(opts, "sendRequestTestV2", testData_, secondPartPool, oppBridge, chainId)
}

// SendRequestTestV2 is a paid mutator transaction binding the contract method 0xf9ee520e.
//
// Solidity: function sendRequestTestV2(uint256 testData_, address secondPartPool, address oppBridge, uint256 chainId) returns()
func (_MockDexPool *MockDexPoolSession) SendRequestTestV2(testData_ *big.Int, secondPartPool common.Address, oppBridge common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _MockDexPool.Contract.SendRequestTestV2(&_MockDexPool.TransactOpts, testData_, secondPartPool, oppBridge, chainId)
}

// SendRequestTestV2 is a paid mutator transaction binding the contract method 0xf9ee520e.
//
// Solidity: function sendRequestTestV2(uint256 testData_, address secondPartPool, address oppBridge, uint256 chainId) returns()
func (_MockDexPool *MockDexPoolTransactorSession) SendRequestTestV2(testData_ *big.Int, secondPartPool common.Address, oppBridge common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _MockDexPool.Contract.SendRequestTestV2(&_MockDexPool.TransactOpts, testData_, secondPartPool, oppBridge, chainId)
}

// MockDexPoolRequestSendedIterator is returned from FilterRequestSended and is used to iterate over the raw logs and unpacked data for RequestSended events raised by the MockDexPool contract.
type MockDexPoolRequestSendedIterator struct {
	Event *MockDexPoolRequestSended // Event containing the contract specifics and raw log

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
func (it *MockDexPoolRequestSendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockDexPoolRequestSended)
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
		it.Event = new(MockDexPoolRequestSended)
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
func (it *MockDexPoolRequestSendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockDexPoolRequestSendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockDexPoolRequestSended represents a RequestSended event raised by the MockDexPool contract.
type MockDexPoolRequestSended struct {
	ReqId [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRequestSended is a free log retrieval operation binding the contract event 0xa67b89382ffb8a1a6987e2a4c7a66d2383ed74b29abbf2a8b9d2adac76b0d2c8.
//
// Solidity: event RequestSended(bytes32 reqId)
func (_MockDexPool *MockDexPoolFilterer) FilterRequestSended(opts *bind.FilterOpts) (*MockDexPoolRequestSendedIterator, error) {

	logs, sub, err := _MockDexPool.contract.FilterLogs(opts, "RequestSended")
	if err != nil {
		return nil, err
	}
	return &MockDexPoolRequestSendedIterator{contract: _MockDexPool.contract, event: "RequestSended", logs: logs, sub: sub}, nil
}

// WatchRequestSended is a free log subscription operation binding the contract event 0xa67b89382ffb8a1a6987e2a4c7a66d2383ed74b29abbf2a8b9d2adac76b0d2c8.
//
// Solidity: event RequestSended(bytes32 reqId)
func (_MockDexPool *MockDexPoolFilterer) WatchRequestSended(opts *bind.WatchOpts, sink chan<- *MockDexPoolRequestSended) (event.Subscription, error) {

	logs, sub, err := _MockDexPool.contract.WatchLogs(opts, "RequestSended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockDexPoolRequestSended)
				if err := _MockDexPool.contract.UnpackLog(event, "RequestSended", log); err != nil {
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

// ParseRequestSended is a log parse operation binding the contract event 0xa67b89382ffb8a1a6987e2a4c7a66d2383ed74b29abbf2a8b9d2adac76b0d2c8.
//
// Solidity: event RequestSended(bytes32 reqId)
func (_MockDexPool *MockDexPoolFilterer) ParseRequestSended(log types.Log) (*MockDexPoolRequestSended, error) {
	event := new(MockDexPoolRequestSended)
	if err := _MockDexPool.contract.UnpackLog(event, "RequestSended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
