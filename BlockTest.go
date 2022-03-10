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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"prevBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"epochBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"transactionsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceHeigh\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"blockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"selector\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"}],\"name\":\"oracleRequestTest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"selector\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"oppositeBridge\",\"type\":\"bytes32\"}],\"name\":\"oracleRequestTestSolana\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061048b806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063732a56a41461004657806392ee7a5f1461006b578063b6785f3c1461007e575b600080fd5b6100596100543660046101fa565b610091565b60405190815260200160405180910390f35b6100596100793660046102b9565b61013f565b61005961008c36600461031b565b6101c0565b6040516001600160c01b031960c088811b8216602084015260288301889052604883018790526068830186905284811b8216608884015283901b16609082015260009060029060980160408051601f19818403018152908290526100f49161038a565b602060405180830381855afa158015610111573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061013491906103c5565b979650505050505050565b60006002868686868660405160200161015c9594939291906103de565b60408051601f19818403018152908290526101769161038a565b602060405180830381855afa158015610193573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906101b691906103c5565b9695505050505050565b60006002868686868660405160200161015c959493929190610412565b803567ffffffffffffffff811681146101f557600080fd5b919050565b60008060008060008060c0878903121561021357600080fd5b61021c876101dd565b955060208701359450604087013593506060870135925061023f608088016101dd565b915061024d60a088016101dd565b90509295509295509295565b80356001600160a01b03811681146101f557600080fd5b60008083601f84011261028257600080fd5b50813567ffffffffffffffff81111561029a57600080fd5b6020830191508360208285010111156102b257600080fd5b9250929050565b6000806000806000608086880312156102d157600080fd5b6102da86610259565b945060208601359350604086013567ffffffffffffffff8111156102fd57600080fd5b61030988828901610270565b96999598509660600135949350505050565b60008060008060006080868803121561033357600080fd5b61033c86610259565b945060208601359350604086013567ffffffffffffffff81111561035f57600080fd5b61036b88828901610270565b909450925061037e905060608701610259565b90509295509295909350565b6000825160005b818110156103ab5760208186018101518583015201610391565b818111156103ba576000828501525b509190910192915050565b6000602082840312156103d757600080fd5b5051919050565b6bffffffffffffffffffffffff198660601b1681528460148201528284603483013760349201918201526054019392505050565b60006bffffffffffffffffffffffff19808860601b1683528660148401528486603485013760609390931b9092169201603481019290925250604801939250505056fea2646970667358221220d6fe083ab7aea32d7c9d3e505285066db02ceea4f853bfb9a16ffba9d76727c564736f6c634300080a0033",
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

// OracleRequestTestSolana is a free data retrieval call binding the contract method 0x92ee7a5f.
//
// Solidity: function oracleRequestTestSolana(address bridge, bytes32 requestId, bytes selector, bytes32 oppositeBridge) pure returns(bytes32)
func (_BlockTest *BlockTestCaller) OracleRequestTestSolana(opts *bind.CallOpts, bridge common.Address, requestId [32]byte, selector []byte, oppositeBridge [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BlockTest.contract.Call(opts, &out, "oracleRequestTestSolana", bridge, requestId, selector, oppositeBridge)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OracleRequestTestSolana is a free data retrieval call binding the contract method 0x92ee7a5f.
//
// Solidity: function oracleRequestTestSolana(address bridge, bytes32 requestId, bytes selector, bytes32 oppositeBridge) pure returns(bytes32)
func (_BlockTest *BlockTestSession) OracleRequestTestSolana(bridge common.Address, requestId [32]byte, selector []byte, oppositeBridge [32]byte) ([32]byte, error) {
	return _BlockTest.Contract.OracleRequestTestSolana(&_BlockTest.CallOpts, bridge, requestId, selector, oppositeBridge)
}

// OracleRequestTestSolana is a free data retrieval call binding the contract method 0x92ee7a5f.
//
// Solidity: function oracleRequestTestSolana(address bridge, bytes32 requestId, bytes selector, bytes32 oppositeBridge) pure returns(bytes32)
func (_BlockTest *BlockTestCallerSession) OracleRequestTestSolana(bridge common.Address, requestId [32]byte, selector []byte, oppositeBridge [32]byte) ([32]byte, error) {
	return _BlockTest.Contract.OracleRequestTestSolana(&_BlockTest.CallOpts, bridge, requestId, selector, oppositeBridge)
}
