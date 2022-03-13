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

// BlockTestMetaData contains all meta data concerning the BlockTest contract.
var BlockTestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"prevBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"epochBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"transactionsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceHeigh\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"blockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"selector\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"}],\"name\":\"oracleRequestTest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bridge\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"selector\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"oppositeBridge\",\"type\":\"bytes32\"}],\"name\":\"oracleRequestTestSolana\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610472806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80635e7eea6914610046578063732a56a41461006b578063b6785f3c1461007e575b600080fd5b610059610054366004610226565b610091565b60405190815260200160405180910390f35b61005961007936600461029e565b610112565b61005961008c366004610314565b6101c0565b6000600286868686866040516020016100ae959493929190610383565b60408051601f19818403018152908290526100c8916103a5565b602060405180830381855afa1580156100e5573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061010891906103e0565b9695505050505050565b6040516001600160c01b031960c088811b8216602084015260288301889052604883018790526068830186905284811b8216608884015283901b16609082015260009060029060980160408051601f1981840301815290829052610175916103a5565b602060405180830381855afa158015610192573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906101b591906103e0565b979650505050505050565b6000600286868686866040516020016100ae9594939291906103f9565b60008083601f8401126101ef57600080fd5b50813567ffffffffffffffff81111561020757600080fd5b60208301915083602082850101111561021f57600080fd5b9250929050565b60008060008060006080868803121561023e57600080fd5b8535945060208601359350604086013567ffffffffffffffff81111561026357600080fd5b61026f888289016101dd565b96999598509660600135949350505050565b803567ffffffffffffffff8116811461029957600080fd5b919050565b60008060008060008060c087890312156102b757600080fd5b6102c087610281565b95506020870135945060408701359350606087013592506102e360808801610281565b91506102f160a08801610281565b90509295509295509295565b80356001600160a01b038116811461029957600080fd5b60008060008060006080868803121561032c57600080fd5b610335866102fd565b945060208601359350604086013567ffffffffffffffff81111561035857600080fd5b610364888289016101dd565b90945092506103779050606087016102fd565b90509295509295909350565b8581528460208201528284604083013760409201918201526060019392505050565b6000825160005b818110156103c657602081860181015185830152016103ac565b818111156103d5576000828501525b509190910192915050565b6000602082840312156103f257600080fd5b5051919050565b60006bffffffffffffffffffffffff19808860601b1683528660148401528486603485013760609390931b9092169201603481019290925250604801939250505056fea2646970667358221220244187d2acfa0e29568777b3ae3c98dd63dc00008230a07509375a8c0f1a9d1664736f6c634300080a0033",
}

// BlockTestABI is the input ABI used to generate the binding from.
// Deprecated: Use BlockTestMetaData.ABI instead.
var BlockTestABI = BlockTestMetaData.ABI

// BlockTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BlockTestMetaData.Bin instead.
var BlockTestBin = BlockTestMetaData.Bin

// DeployBlockTest deploys a new Ethereum contract, binding an instance of BlockTest to it.
func DeployBlockTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BlockTest, error) {
	parsed, err := BlockTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BlockTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BlockTest{BlockTestCaller: BlockTestCaller{contract: contract}, BlockTestTransactor: BlockTestTransactor{contract: contract}, BlockTestFilterer: BlockTestFilterer{contract: contract}}, nil
}

// BlockTest is an auto generated Go binding around an Ethereum contract.
type BlockTest struct {
	BlockTestCaller     // Read-only binding to the contract
	BlockTestTransactor // Write-only binding to the contract
	BlockTestFilterer   // Log filterer for contract events
}

// BlockTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlockTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
	gsn      *GsnCallOpts
}

func (_BlockTest *BlockTestTransactor) SetGSNOptions(opts *GsnCallOpts) {
	_BlockTest.gsn = opts
}

// BlockTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockTestSession struct {
	Contract     *BlockTest        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlockTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockTestCallerSession struct {
	Contract *BlockTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BlockTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockTestTransactorSession struct {
	Contract     *BlockTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BlockTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlockTestRaw struct {
	Contract *BlockTest // Generic contract binding to access the raw methods on
}

// BlockTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockTestCallerRaw struct {
	Contract *BlockTestCaller // Generic read-only contract binding to access the raw methods on
}

// BlockTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockTestTransactorRaw struct {
	Contract *BlockTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockTest creates a new instance of BlockTest, bound to a specific deployed contract.
func NewBlockTest(address common.Address, backend bind.ContractBackend) (*BlockTest, error) {
	contract, err := bindBlockTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlockTest{BlockTestCaller: BlockTestCaller{contract: contract}, BlockTestTransactor: BlockTestTransactor{contract: contract}, BlockTestFilterer: BlockTestFilterer{contract: contract}}, nil
}

// NewBlockTestCaller creates a new read-only instance of BlockTest, bound to a specific deployed contract.
func NewBlockTestCaller(address common.Address, caller bind.ContractCaller) (*BlockTestCaller, error) {
	contract, err := bindBlockTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockTestCaller{contract: contract}, nil
}

// NewBlockTestTransactor creates a new write-only instance of BlockTest, bound to a specific deployed contract.
func NewBlockTestTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockTestTransactor, error) {
	contract, err := bindBlockTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockTestTransactor{contract: contract}, nil
}

// NewBlockTestFilterer creates a new log filterer instance of BlockTest, bound to a specific deployed contract.
func NewBlockTestFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockTestFilterer, error) {
	contract, err := bindBlockTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockTestFilterer{contract: contract}, nil
}

// bindBlockTest binds a generic wrapper to an already deployed contract.
func bindBlockTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlockTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockTest *BlockTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockTest.Contract.BlockTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockTest *BlockTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockTest.Contract.BlockTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockTest *BlockTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockTest.Contract.BlockTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockTest *BlockTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockTest *BlockTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockTest *BlockTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockTest.Contract.contract.Transact(opts, method, params...)
}

// BlockHash is a free data retrieval call binding the contract method 0x732a56a4.
//
// Solidity: function blockHash(uint64 chainId, bytes32 prevBlockHash, bytes32 epochBlockHash, bytes32 transactionsRoot, uint64 sourceHeigh, uint64 height) pure returns(bytes32)
func (_BlockTest *BlockTestCaller) BlockHash(opts *bind.CallOpts, chainId uint64, prevBlockHash [32]byte, epochBlockHash [32]byte, transactionsRoot [32]byte, sourceHeigh uint64, height uint64) ([32]byte, error) {
	var out []interface{}
	err := _BlockTest.contract.Call(opts, &out, "blockHash", chainId, prevBlockHash, epochBlockHash, transactionsRoot, sourceHeigh, height)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BlockHash is a free data retrieval call binding the contract method 0x732a56a4.
//
// Solidity: function blockHash(uint64 chainId, bytes32 prevBlockHash, bytes32 epochBlockHash, bytes32 transactionsRoot, uint64 sourceHeigh, uint64 height) pure returns(bytes32)
func (_BlockTest *BlockTestSession) BlockHash(chainId uint64, prevBlockHash [32]byte, epochBlockHash [32]byte, transactionsRoot [32]byte, sourceHeigh uint64, height uint64) ([32]byte, error) {
	return _BlockTest.Contract.BlockHash(&_BlockTest.CallOpts, chainId, prevBlockHash, epochBlockHash, transactionsRoot, sourceHeigh, height)
}

// BlockHash is a free data retrieval call binding the contract method 0x732a56a4.
//
// Solidity: function blockHash(uint64 chainId, bytes32 prevBlockHash, bytes32 epochBlockHash, bytes32 transactionsRoot, uint64 sourceHeigh, uint64 height) pure returns(bytes32)
func (_BlockTest *BlockTestCallerSession) BlockHash(chainId uint64, prevBlockHash [32]byte, epochBlockHash [32]byte, transactionsRoot [32]byte, sourceHeigh uint64, height uint64) ([32]byte, error) {
	return _BlockTest.Contract.BlockHash(&_BlockTest.CallOpts, chainId, prevBlockHash, epochBlockHash, transactionsRoot, sourceHeigh, height)
}

// OracleRequestTest is a free data retrieval call binding the contract method 0xb6785f3c.
//
// Solidity: function oracleRequestTest(address bridge, bytes32 requestId, bytes selector, address receiveSide) pure returns(bytes32)
func (_BlockTest *BlockTestCaller) OracleRequestTest(opts *bind.CallOpts, bridge common.Address, requestId [32]byte, selector []byte, receiveSide common.Address) ([32]byte, error) {
	var out []interface{}
	err := _BlockTest.contract.Call(opts, &out, "oracleRequestTest", bridge, requestId, selector, receiveSide)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OracleRequestTest is a free data retrieval call binding the contract method 0xb6785f3c.
//
// Solidity: function oracleRequestTest(address bridge, bytes32 requestId, bytes selector, address receiveSide) pure returns(bytes32)
func (_BlockTest *BlockTestSession) OracleRequestTest(bridge common.Address, requestId [32]byte, selector []byte, receiveSide common.Address) ([32]byte, error) {
	return _BlockTest.Contract.OracleRequestTest(&_BlockTest.CallOpts, bridge, requestId, selector, receiveSide)
}

// OracleRequestTest is a free data retrieval call binding the contract method 0xb6785f3c.
//
// Solidity: function oracleRequestTest(address bridge, bytes32 requestId, bytes selector, address receiveSide) pure returns(bytes32)
func (_BlockTest *BlockTestCallerSession) OracleRequestTest(bridge common.Address, requestId [32]byte, selector []byte, receiveSide common.Address) ([32]byte, error) {
	return _BlockTest.Contract.OracleRequestTest(&_BlockTest.CallOpts, bridge, requestId, selector, receiveSide)
}

// OracleRequestTestSolana is a free data retrieval call binding the contract method 0x5e7eea69.
//
// Solidity: function oracleRequestTestSolana(bytes32 bridge, bytes32 requestId, bytes selector, bytes32 oppositeBridge) pure returns(bytes32)
func (_BlockTest *BlockTestCaller) OracleRequestTestSolana(opts *bind.CallOpts, bridge [32]byte, requestId [32]byte, selector []byte, oppositeBridge [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BlockTest.contract.Call(opts, &out, "oracleRequestTestSolana", bridge, requestId, selector, oppositeBridge)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OracleRequestTestSolana is a free data retrieval call binding the contract method 0x5e7eea69.
//
// Solidity: function oracleRequestTestSolana(bytes32 bridge, bytes32 requestId, bytes selector, bytes32 oppositeBridge) pure returns(bytes32)
func (_BlockTest *BlockTestSession) OracleRequestTestSolana(bridge [32]byte, requestId [32]byte, selector []byte, oppositeBridge [32]byte) ([32]byte, error) {
	return _BlockTest.Contract.OracleRequestTestSolana(&_BlockTest.CallOpts, bridge, requestId, selector, oppositeBridge)
}

// OracleRequestTestSolana is a free data retrieval call binding the contract method 0x5e7eea69.
//
// Solidity: function oracleRequestTestSolana(bytes32 bridge, bytes32 requestId, bytes selector, bytes32 oppositeBridge) pure returns(bytes32)
func (_BlockTest *BlockTestCallerSession) OracleRequestTestSolana(bridge [32]byte, requestId [32]byte, selector []byte, oppositeBridge [32]byte) ([32]byte, error) {
	return _BlockTest.Contract.OracleRequestTestSolana(&_BlockTest.CallOpts, bridge, requestId, selector, oppositeBridge)
}
