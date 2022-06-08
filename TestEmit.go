// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wrappers

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
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

// TestEmitMetaData contains all meta data concerning the TestEmit contract.

var TestEmitMetaData = &bind.MetaData{

	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"requestType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"selector\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"OracleRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"requestType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"bridge\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"selector\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"oppositeBridge\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"OracleRequestSolana\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"reqId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"bridgeFrom\",\"type\":\"bytes32\"}],\"name\":\"ReceiveRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"what\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"when\",\"type\":\"uint256\"}],\"name\":\"TestEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"from\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"oppositeBridge\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"}],\"name\":\"addContractBind\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"selector\",\"type\":\"bytes\"}],\"name\":\"testOracleRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bridgeFrom\",\"type\":\"bytes32\"}],\"name\":\"testReceiveRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"what\",\"type\":\"string\"}],\"name\":\"testTestEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061058e806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80632b4602dc1461005c5780632d0335ab14610071578063bf564af6146100ac578063e27ad9a5146100bf578063fc28534d146100d2575b600080fd5b61006f61006a366004610293565b6100e5565b005b61009a61007f3660046102b5565b6001600160a01b031660009081526020819052604090205490565b60405190815260200160405180910390f35b61006f6100ba3660046102e5565b61012c565b61006f6100cd366004610361565b610174565b61006f6100e03660046103a3565b610243565b6040805183815230602082018190529181018390527ffb4571f3f72d75f82b13143fd35e1935981b5bd063d4bc973acc83934d8777f29060600160405180910390a1505050565b6040513090819085907fa8bfc8b51168a5f110dd8ec0f0f752dff26fae23d0e2cd422653e78f92d94586906101669087908790429061045e565b60405180910390a350505050565b806101c65760405162461bcd60e51b815260206004820152601c60248201527f4272696467653a20696e76616c69642027746f2720616464726573730000000060448201526064015b60405180910390fd5b826102135760405162461bcd60e51b815260206004820152601e60248201527f4272696467653a20696e76616c6964202766726f6d272061646472657373000060448201526064016101bd565b6000928352600160208181526040808620948652938152838520928552919091529120805460ff19169091179055565b604051309063cafebabe907f5c3fb349179e8d8347e69078d6b9912b9724461c39bf776a56925267afc55aff9061028590849087908790839081908890610497565b60405180910390a150505050565b600080604083850312156102a657600080fd5b50508035926020909101359150565b6000602082840312156102c757600080fd5b81356001600160a01b03811681146102de57600080fd5b9392505050565b6000806000604084860312156102fa57600080fd5b83359250602084013567ffffffffffffffff8082111561031957600080fd5b818601915086601f83011261032d57600080fd5b81358181111561033c57600080fd5b87602082850101111561034e57600080fd5b6020830194508093505050509250925092565b60008060006060848603121561037657600080fd5b505081359360208301359350604090920135919050565b634e487b7160e01b600052604160045260246000fd5b600080604083850312156103b657600080fd5b82359150602083013567ffffffffffffffff808211156103d557600080fd5b818501915085601f8301126103e957600080fd5b8135818111156103fb576103fb61038d565b604051601f8201601f19908116603f011681019083821181831017156104235761042361038d565b8160405282815288602084870101111561043c57600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b604081528260408201528284606083013760006060848301015260006060601f19601f8601168301019050826020830152949350505050565b60e08152600a60e0820152691cd95d14995c5d595cdd60b21b6101008201526000610120602060018060a01b038a168185015288604085015281606085015287518083860152600092505b80831015610501578883018201518584016101400152918101916104e2565b8083111561051457600061014082870101525b601f01601f1916840161014001925061053b91505060808301866001600160a01b03169052565b6001600160a01b039390931660a082015260c0015294935050505056fea26469706673582212203e659f8aa3ab60e357cdd6b24dfbc9f42024d08ce982d4ab1de10d05055f72da64736f6c634300080a0033",
}

// TestEmitABI is the input ABI used to generate the binding from.
// Deprecated: Use TestEmitMetaData.ABI instead.
var TestEmitABI = TestEmitMetaData.ABI

// TestEmitBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestEmitMetaData.Bin instead.
var TestEmitBin = TestEmitMetaData.Bin

// DeployTestEmit deploys a new Ethereum contract, binding an instance of TestEmit to it.
func DeployTestEmit(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestEmit, error) {
	parsed, err := TestEmitMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestEmitBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestEmit{TestEmitCaller: TestEmitCaller{contract: contract}, TestEmitTransactor: TestEmitTransactor{contract: contract}, TestEmitFilterer: TestEmitFilterer{contract: contract}}, nil
}

// TestEmit is an auto generated Go binding around an Ethereum contract.
type TestEmit struct {
	TestEmitCaller     // Read-only binding to the contract
	TestEmitTransactor // Write-only binding to the contract
	TestEmitFilterer   // Log filterer for contract events
}

// TestEmitCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestEmitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestEmitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestEmitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
	gsn      *GsnCallOpts
}

func (_TestEmit *TestEmitTransactor) SetGSNOptions(opts *GsnCallOpts) {
	_TestEmit.gsn = opts
}

// TestEmitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestEmitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestEmitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestEmitSession struct {
	Contract     *TestEmit         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestEmitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestEmitCallerSession struct {
	Contract *TestEmitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TestEmitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestEmitTransactorSession struct {
	Contract     *TestEmitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TestEmitRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestEmitRaw struct {
	Contract *TestEmit // Generic contract binding to access the raw methods on
}

// TestEmitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestEmitCallerRaw struct {
	Contract *TestEmitCaller // Generic read-only contract binding to access the raw methods on
}

// TestEmitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestEmitTransactorRaw struct {
	Contract *TestEmitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestEmit creates a new instance of TestEmit, bound to a specific deployed contract.
func NewTestEmit(address common.Address, backend bind.ContractBackend) (*TestEmit, error) {
	contract, err := bindTestEmit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestEmit{TestEmitCaller: TestEmitCaller{contract: contract}, TestEmitTransactor: TestEmitTransactor{contract: contract}, TestEmitFilterer: TestEmitFilterer{contract: contract}}, nil
}

// NewTestEmitCaller creates a new read-only instance of TestEmit, bound to a specific deployed contract.
func NewTestEmitCaller(address common.Address, caller bind.ContractCaller) (*TestEmitCaller, error) {
	contract, err := bindTestEmit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestEmitCaller{contract: contract}, nil
}

// NewTestEmitTransactor creates a new write-only instance of TestEmit, bound to a specific deployed contract.
func NewTestEmitTransactor(address common.Address, transactor bind.ContractTransactor) (*TestEmitTransactor, error) {
	contract, err := bindTestEmit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestEmitTransactor{contract: contract}, nil
}

// NewTestEmitFilterer creates a new log filterer instance of TestEmit, bound to a specific deployed contract.
func NewTestEmitFilterer(address common.Address, filterer bind.ContractFilterer) (*TestEmitFilterer, error) {
	contract, err := bindTestEmit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestEmitFilterer{contract: contract}, nil
}

// bindTestEmit binds a generic wrapper to an already deployed contract.
func bindTestEmit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestEmitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestEmit *TestEmitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestEmit.Contract.TestEmitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestEmit *TestEmitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestEmit.Contract.TestEmitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestEmit *TestEmitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestEmit.Contract.TestEmitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestEmit *TestEmitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestEmit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestEmit *TestEmitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestEmit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestEmit *TestEmitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestEmit.Contract.contract.Transact(opts, method, params...)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address from) view returns(uint256)
func (_TestEmit *TestEmitCaller) GetNonce(opts *bind.CallOpts, from common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestEmit.contract.Call(opts, &out, "getNonce", from)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address from) view returns(uint256)
func (_TestEmit *TestEmitSession) GetNonce(from common.Address) (*big.Int, error) {
	return _TestEmit.Contract.GetNonce(&_TestEmit.CallOpts, from)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address from) view returns(uint256)
func (_TestEmit *TestEmitCallerSession) GetNonce(from common.Address) (*big.Int, error) {
	return _TestEmit.Contract.GetNonce(&_TestEmit.CallOpts, from)
}

// AddContractBind is a paid mutator transaction binding the contract method 0xe27ad9a5.
//
// Solidity: function addContractBind(bytes32 from, bytes32 oppositeBridge, bytes32 to) returns()
func (_TestEmit *TestEmitTransactor) AddContractBind(opts *bind.TransactOpts, from [32]byte, oppositeBridge [32]byte, to [32]byte) (*types.Transaction, error) {
	return _TestEmit.contract.Transact(opts, "addContractBind", from, oppositeBridge, to)
}
func (_TestEmit *TestEmitTransactor) AddContractBindOverGsn(opts *bind.TransactOpts, from [32]byte, oppositeBridge [32]byte, to [32]byte) (common.Hash, error) {
	return GsnExecutor(_TestEmit.gsn, TestEmitMetaData.ABI, "addContractBind", from, oppositeBridge, to)
}

// AddContractBind is a paid mutator transaction binding the contract method 0xe27ad9a5.
//
// Solidity: function addContractBind(bytes32 from, bytes32 oppositeBridge, bytes32 to) returns()
func (_TestEmit *TestEmitSession) AddContractBind(from [32]byte, oppositeBridge [32]byte, to [32]byte) (*types.Transaction, error) {
	return _TestEmit.Contract.AddContractBind(&_TestEmit.TransactOpts, from, oppositeBridge, to)
}
func (_TestEmit *TestEmitSession) AddContractBindOverGsn(from [32]byte, oppositeBridge [32]byte, to [32]byte) (common.Hash, error) {
	return _TestEmit.Contract.AddContractBindOverGsn(&_TestEmit.TransactOpts, from, oppositeBridge, to)
}

// AddContractBind is a paid mutator transaction binding the contract method 0xe27ad9a5.
//
// Solidity: function addContractBind(bytes32 from, bytes32 oppositeBridge, bytes32 to) returns()
func (_TestEmit *TestEmitTransactorSession) AddContractBind(from [32]byte, oppositeBridge [32]byte, to [32]byte) (*types.Transaction, error) {
	return _TestEmit.Contract.AddContractBind(&_TestEmit.TransactOpts, from, oppositeBridge, to)
}
func (_TestEmit *TestEmitTransactorSession) AddContractBindOverGsn(from [32]byte, oppositeBridge [32]byte, to [32]byte) (common.Hash, error) {
	return _TestEmit.Contract.AddContractBindOverGsn(&_TestEmit.TransactOpts, from, oppositeBridge, to)
}

// TestOracleRequest is a paid mutator transaction binding the contract method 0xfc28534d.
//
// Solidity: function testOracleRequest(bytes32 requestId, bytes selector) returns()
func (_TestEmit *TestEmitTransactor) TestOracleRequest(opts *bind.TransactOpts, requestId [32]byte, selector []byte) (*types.Transaction, error) {
	return _TestEmit.contract.Transact(opts, "testOracleRequest", requestId, selector)
}
func (_TestEmit *TestEmitTransactor) TestOracleRequestOverGsn(opts *bind.TransactOpts, requestId [32]byte, selector []byte) (common.Hash, error) {
	return GsnExecutor(_TestEmit.gsn, TestEmitMetaData.ABI, "testOracleRequest", requestId, selector)
}

// TestOracleRequest is a paid mutator transaction binding the contract method 0xfc28534d.
//
// Solidity: function testOracleRequest(bytes32 requestId, bytes selector) returns()
func (_TestEmit *TestEmitSession) TestOracleRequest(requestId [32]byte, selector []byte) (*types.Transaction, error) {
	return _TestEmit.Contract.TestOracleRequest(&_TestEmit.TransactOpts, requestId, selector)
}
func (_TestEmit *TestEmitSession) TestOracleRequestOverGsn(requestId [32]byte, selector []byte) (common.Hash, error) {
	return _TestEmit.Contract.TestOracleRequestOverGsn(&_TestEmit.TransactOpts, requestId, selector)
}

// TestOracleRequest is a paid mutator transaction binding the contract method 0xfc28534d.
//
// Solidity: function testOracleRequest(bytes32 requestId, bytes selector) returns()
func (_TestEmit *TestEmitTransactorSession) TestOracleRequest(requestId [32]byte, selector []byte) (*types.Transaction, error) {
	return _TestEmit.Contract.TestOracleRequest(&_TestEmit.TransactOpts, requestId, selector)
}
func (_TestEmit *TestEmitTransactorSession) TestOracleRequestOverGsn(requestId [32]byte, selector []byte) (common.Hash, error) {
	return _TestEmit.Contract.TestOracleRequestOverGsn(&_TestEmit.TransactOpts, requestId, selector)
}

// TestReceiveRequest is a paid mutator transaction binding the contract method 0x2b4602dc.
//
// Solidity: function testReceiveRequest(bytes32 requestId, bytes32 bridgeFrom) returns()
func (_TestEmit *TestEmitTransactor) TestReceiveRequest(opts *bind.TransactOpts, requestId [32]byte, bridgeFrom [32]byte) (*types.Transaction, error) {
	return _TestEmit.contract.Transact(opts, "testReceiveRequest", requestId, bridgeFrom)
}
func (_TestEmit *TestEmitTransactor) TestReceiveRequestOverGsn(opts *bind.TransactOpts, requestId [32]byte, bridgeFrom [32]byte) (common.Hash, error) {
	return GsnExecutor(_TestEmit.gsn, TestEmitMetaData.ABI, "testReceiveRequest", requestId, bridgeFrom)
}

// TestReceiveRequest is a paid mutator transaction binding the contract method 0x2b4602dc.
//
// Solidity: function testReceiveRequest(bytes32 requestId, bytes32 bridgeFrom) returns()
func (_TestEmit *TestEmitSession) TestReceiveRequest(requestId [32]byte, bridgeFrom [32]byte) (*types.Transaction, error) {
	return _TestEmit.Contract.TestReceiveRequest(&_TestEmit.TransactOpts, requestId, bridgeFrom)
}
func (_TestEmit *TestEmitSession) TestReceiveRequestOverGsn(requestId [32]byte, bridgeFrom [32]byte) (common.Hash, error) {
	return _TestEmit.Contract.TestReceiveRequestOverGsn(&_TestEmit.TransactOpts, requestId, bridgeFrom)
}

// TestReceiveRequest is a paid mutator transaction binding the contract method 0x2b4602dc.
//
// Solidity: function testReceiveRequest(bytes32 requestId, bytes32 bridgeFrom) returns()
func (_TestEmit *TestEmitTransactorSession) TestReceiveRequest(requestId [32]byte, bridgeFrom [32]byte) (*types.Transaction, error) {
	return _TestEmit.Contract.TestReceiveRequest(&_TestEmit.TransactOpts, requestId, bridgeFrom)
}
func (_TestEmit *TestEmitTransactorSession) TestReceiveRequestOverGsn(requestId [32]byte, bridgeFrom [32]byte) (common.Hash, error) {
	return _TestEmit.Contract.TestReceiveRequestOverGsn(&_TestEmit.TransactOpts, requestId, bridgeFrom)
}

// TestTestEvent is a paid mutator transaction binding the contract method 0xbf564af6.
//
// Solidity: function testTestEvent(bytes32 id, string what) returns()
func (_TestEmit *TestEmitTransactor) TestTestEvent(opts *bind.TransactOpts, id [32]byte, what string) (*types.Transaction, error) {
	return _TestEmit.contract.Transact(opts, "testTestEvent", id, what)
}
func (_TestEmit *TestEmitTransactor) TestTestEventOverGsn(opts *bind.TransactOpts, id [32]byte, what string) (common.Hash, error) {
	return GsnExecutor(_TestEmit.gsn, TestEmitMetaData.ABI, "testTestEvent", id, what)
}

// TestTestEvent is a paid mutator transaction binding the contract method 0xbf564af6.
//
// Solidity: function testTestEvent(bytes32 id, string what) returns()
func (_TestEmit *TestEmitSession) TestTestEvent(id [32]byte, what string) (*types.Transaction, error) {
	return _TestEmit.Contract.TestTestEvent(&_TestEmit.TransactOpts, id, what)
}
func (_TestEmit *TestEmitSession) TestTestEventOverGsn(id [32]byte, what string) (common.Hash, error) {
	return _TestEmit.Contract.TestTestEventOverGsn(&_TestEmit.TransactOpts, id, what)
}

// TestTestEvent is a paid mutator transaction binding the contract method 0xbf564af6.
//
// Solidity: function testTestEvent(bytes32 id, string what) returns()
func (_TestEmit *TestEmitTransactorSession) TestTestEvent(id [32]byte, what string) (*types.Transaction, error) {
	return _TestEmit.Contract.TestTestEvent(&_TestEmit.TransactOpts, id, what)
}
func (_TestEmit *TestEmitTransactorSession) TestTestEventOverGsn(id [32]byte, what string) (common.Hash, error) {
	return _TestEmit.Contract.TestTestEventOverGsn(&_TestEmit.TransactOpts, id, what)
}

// TestEmitOracleRequestIterator is returned from FilterOracleRequest and is used to iterate over the raw logs and unpacked data for OracleRequest events raised by the TestEmit contract.
type TestEmitOracleRequestIterator struct {
	Event *TestEmitOracleRequest // Event containing the contract specifics and raw log

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
func (it *TestEmitOracleRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestEmitOracleRequest)
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
		it.Event = new(TestEmitOracleRequest)
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
func (it *TestEmitOracleRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestEmitOracleRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestEmitOracleRequest represents a OracleRequest event raised by the TestEmit contract.
type TestEmitOracleRequest struct {
	RequestType    string
	Bridge         common.Address
	RequestId      [32]byte
	Selector       []byte
	ReceiveSide    common.Address
	OppositeBridge common.Address
	ChainId        *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOracleRequest is a free log retrieval operation binding the contract event 0x5c3fb349179e8d8347e69078d6b9912b9724461c39bf776a56925267afc55aff.
//
// Solidity: event OracleRequest(string requestType, address bridge, bytes32 requestId, bytes selector, address receiveSide, address oppositeBridge, uint256 chainId)
func (_TestEmit *TestEmitFilterer) FilterOracleRequest(opts *bind.FilterOpts) (*TestEmitOracleRequestIterator, error) {

	logs, sub, err := _TestEmit.contract.FilterLogs(opts, "OracleRequest")
	if err != nil {
		return nil, err
	}
	return &TestEmitOracleRequestIterator{contract: _TestEmit.contract, event: "OracleRequest", logs: logs, sub: sub}, nil
}

// WatchOracleRequest is a free log subscription operation binding the contract event 0x5c3fb349179e8d8347e69078d6b9912b9724461c39bf776a56925267afc55aff.
//
// Solidity: event OracleRequest(string requestType, address bridge, bytes32 requestId, bytes selector, address receiveSide, address oppositeBridge, uint256 chainId)
func (_TestEmit *TestEmitFilterer) WatchOracleRequest(opts *bind.WatchOpts, sink chan<- *TestEmitOracleRequest) (event.Subscription, error) {

	logs, sub, err := _TestEmit.contract.WatchLogs(opts, "OracleRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestEmitOracleRequest)
				if err := _TestEmit.contract.UnpackLog(event, "OracleRequest", log); err != nil {
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

// ParseOracleRequest is a log parse operation binding the contract event 0x5c3fb349179e8d8347e69078d6b9912b9724461c39bf776a56925267afc55aff.
//
// Solidity: event OracleRequest(string requestType, address bridge, bytes32 requestId, bytes selector, address receiveSide, address oppositeBridge, uint256 chainId)
func (_TestEmit *TestEmitFilterer) ParseOracleRequest(log types.Log) (*TestEmitOracleRequest, error) {
	event := new(TestEmitOracleRequest)
	if err := _TestEmit.contract.UnpackLog(event, "OracleRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestEmitOracleRequestSolanaIterator is returned from FilterOracleRequestSolana and is used to iterate over the raw logs and unpacked data for OracleRequestSolana events raised by the TestEmit contract.
type TestEmitOracleRequestSolanaIterator struct {
	Event *TestEmitOracleRequestSolana // Event containing the contract specifics and raw log

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
func (it *TestEmitOracleRequestSolanaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestEmitOracleRequestSolana)
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
		it.Event = new(TestEmitOracleRequestSolana)
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
func (it *TestEmitOracleRequestSolanaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestEmitOracleRequestSolanaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestEmitOracleRequestSolana represents a OracleRequestSolana event raised by the TestEmit contract.
type TestEmitOracleRequestSolana struct {
	RequestType    string
	Bridge         [32]byte
	RequestId      [32]byte
	Selector       []byte
	OppositeBridge [32]byte
	ChainId        *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOracleRequestSolana is a free log retrieval operation binding the contract event 0x2b0dad29f6c03635bda7007bdf022fea8d9c71e34bd592b687af9284f41aea13.
//
// Solidity: event OracleRequestSolana(string requestType, bytes32 bridge, bytes32 requestId, bytes selector, bytes32 oppositeBridge, uint256 chainId)
func (_TestEmit *TestEmitFilterer) FilterOracleRequestSolana(opts *bind.FilterOpts) (*TestEmitOracleRequestSolanaIterator, error) {

	logs, sub, err := _TestEmit.contract.FilterLogs(opts, "OracleRequestSolana")
	if err != nil {
		return nil, err
	}
	return &TestEmitOracleRequestSolanaIterator{contract: _TestEmit.contract, event: "OracleRequestSolana", logs: logs, sub: sub}, nil
}

// WatchOracleRequestSolana is a free log subscription operation binding the contract event 0x2b0dad29f6c03635bda7007bdf022fea8d9c71e34bd592b687af9284f41aea13.
//
// Solidity: event OracleRequestSolana(string requestType, bytes32 bridge, bytes32 requestId, bytes selector, bytes32 oppositeBridge, uint256 chainId)
func (_TestEmit *TestEmitFilterer) WatchOracleRequestSolana(opts *bind.WatchOpts, sink chan<- *TestEmitOracleRequestSolana) (event.Subscription, error) {

	logs, sub, err := _TestEmit.contract.WatchLogs(opts, "OracleRequestSolana")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestEmitOracleRequestSolana)
				if err := _TestEmit.contract.UnpackLog(event, "OracleRequestSolana", log); err != nil {
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

// ParseOracleRequestSolana is a log parse operation binding the contract event 0x2b0dad29f6c03635bda7007bdf022fea8d9c71e34bd592b687af9284f41aea13.
//
// Solidity: event OracleRequestSolana(string requestType, bytes32 bridge, bytes32 requestId, bytes selector, bytes32 oppositeBridge, uint256 chainId)
func (_TestEmit *TestEmitFilterer) ParseOracleRequestSolana(log types.Log) (*TestEmitOracleRequestSolana, error) {
	event := new(TestEmitOracleRequestSolana)
	if err := _TestEmit.contract.UnpackLog(event, "OracleRequestSolana", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestEmitReceiveRequestIterator is returned from FilterReceiveRequest and is used to iterate over the raw logs and unpacked data for ReceiveRequest events raised by the TestEmit contract.
type TestEmitReceiveRequestIterator struct {
	Event *TestEmitReceiveRequest // Event containing the contract specifics and raw log

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
func (it *TestEmitReceiveRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestEmitReceiveRequest)
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
		it.Event = new(TestEmitReceiveRequest)
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
func (it *TestEmitReceiveRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestEmitReceiveRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestEmitReceiveRequest represents a ReceiveRequest event raised by the TestEmit contract.
type TestEmitReceiveRequest struct {
	ReqId       [32]byte
	ReceiveSide common.Address
	BridgeFrom  [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReceiveRequest is a free log retrieval operation binding the contract event 0xfb4571f3f72d75f82b13143fd35e1935981b5bd063d4bc973acc83934d8777f2.
//
// Solidity: event ReceiveRequest(bytes32 reqId, address receiveSide, bytes32 bridgeFrom)
func (_TestEmit *TestEmitFilterer) FilterReceiveRequest(opts *bind.FilterOpts) (*TestEmitReceiveRequestIterator, error) {

	logs, sub, err := _TestEmit.contract.FilterLogs(opts, "ReceiveRequest")
	if err != nil {
		return nil, err
	}
	return &TestEmitReceiveRequestIterator{contract: _TestEmit.contract, event: "ReceiveRequest", logs: logs, sub: sub}, nil
}

// WatchReceiveRequest is a free log subscription operation binding the contract event 0xfb4571f3f72d75f82b13143fd35e1935981b5bd063d4bc973acc83934d8777f2.
//
// Solidity: event ReceiveRequest(bytes32 reqId, address receiveSide, bytes32 bridgeFrom)
func (_TestEmit *TestEmitFilterer) WatchReceiveRequest(opts *bind.WatchOpts, sink chan<- *TestEmitReceiveRequest) (event.Subscription, error) {

	logs, sub, err := _TestEmit.contract.WatchLogs(opts, "ReceiveRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestEmitReceiveRequest)
				if err := _TestEmit.contract.UnpackLog(event, "ReceiveRequest", log); err != nil {
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

// ParseReceiveRequest is a log parse operation binding the contract event 0xfb4571f3f72d75f82b13143fd35e1935981b5bd063d4bc973acc83934d8777f2.
//
// Solidity: event ReceiveRequest(bytes32 reqId, address receiveSide, bytes32 bridgeFrom)
func (_TestEmit *TestEmitFilterer) ParseReceiveRequest(log types.Log) (*TestEmitReceiveRequest, error) {
	event := new(TestEmitReceiveRequest)
	if err := _TestEmit.contract.UnpackLog(event, "ReceiveRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestEmitTestEventIterator is returned from FilterTestEvent and is used to iterate over the raw logs and unpacked data for TestEvent events raised by the TestEmit contract.
type TestEmitTestEventIterator struct {
	Event *TestEmitTestEvent // Event containing the contract specifics and raw log

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
func (it *TestEmitTestEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestEmitTestEvent)
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
		it.Event = new(TestEmitTestEvent)
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
func (it *TestEmitTestEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestEmitTestEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestEmitTestEvent represents a TestEvent event raised by the TestEmit contract.
type TestEmitTestEvent struct {
	Id   [32]byte
	Who  common.Address
	What string
	When *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTestEvent is a free log retrieval operation binding the contract event 0xa8bfc8b51168a5f110dd8ec0f0f752dff26fae23d0e2cd422653e78f92d94586.
//
// Solidity: event TestEvent(bytes32 indexed id, address indexed who, string what, uint256 when)
func (_TestEmit *TestEmitFilterer) FilterTestEvent(opts *bind.FilterOpts, id [][32]byte, who []common.Address) (*TestEmitTestEventIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _TestEmit.contract.FilterLogs(opts, "TestEvent", idRule, whoRule)
	if err != nil {
		return nil, err
	}
	return &TestEmitTestEventIterator{contract: _TestEmit.contract, event: "TestEvent", logs: logs, sub: sub}, nil
}

// WatchTestEvent is a free log subscription operation binding the contract event 0xa8bfc8b51168a5f110dd8ec0f0f752dff26fae23d0e2cd422653e78f92d94586.
//
// Solidity: event TestEvent(bytes32 indexed id, address indexed who, string what, uint256 when)
func (_TestEmit *TestEmitFilterer) WatchTestEvent(opts *bind.WatchOpts, sink chan<- *TestEmitTestEvent, id [][32]byte, who []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _TestEmit.contract.WatchLogs(opts, "TestEvent", idRule, whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestEmitTestEvent)
				if err := _TestEmit.contract.UnpackLog(event, "TestEvent", log); err != nil {
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

// ParseTestEvent is a log parse operation binding the contract event 0xa8bfc8b51168a5f110dd8ec0f0f752dff26fae23d0e2cd422653e78f92d94586.
//
// Solidity: event TestEvent(bytes32 indexed id, address indexed who, string what, uint256 when)
func (_TestEmit *TestEmitFilterer) ParseTestEvent(log types.Log) (*TestEmitTestEvent, error) {
	event := new(TestEmitTestEvent)
	if err := _TestEmit.contract.UnpackLog(event, "TestEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
