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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"prevBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"epochBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"transactionsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceHeigh\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"blockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"blockHeaderRawDataTest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"allBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blockTxHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610410806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063732a56a41461003b578063845d481e14610061575b600080fd5b61004e61004936600461027b565b610089565b6040519081526020015b60405180910390f35b61007461006f3660046102da565b610137565b60408051928352602083019190915201610058565b6040516001600160c01b031960c088811b8216602084015260288301889052604883018790526068830186905284811b8216608884015283901b16609082015260009060029060980160408051601f19818403018152908290526100ec9161034c565b602060405180830381855afa158015610109573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061012c9190610387565b979650505050505050565b6000806002848460405160200161014f9291906103a0565b60408051601f19818403018152908290526101699161034c565b602060405180830381855afa158015610186573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906101a99190610387565b91506101f66101bc6068604886886103b0565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506101ff92505050565b90509250929050565b600081516020146102565760405162461bcd60e51b815260206004820152601760248201527f6279746573206c656e677468206973206e6f742033322e000000000000000000604482015260640160405180910390fd5b506020015190565b803567ffffffffffffffff8116811461027657600080fd5b919050565b60008060008060008060c0878903121561029457600080fd5b61029d8761025e565b95506020870135945060408701359350606087013592506102c06080880161025e565b91506102ce60a0880161025e565b90509295509295509295565b600080602083850312156102ed57600080fd5b823567ffffffffffffffff8082111561030557600080fd5b818501915085601f83011261031957600080fd5b81358181111561032857600080fd5b86602082850101111561033a57600080fd5b60209290920196919550909350505050565b6000825160005b8181101561036d5760208186018101518583015201610353565b8181111561037c576000828501525b509190910192915050565b60006020828403121561039957600080fd5b5051919050565b8183823760009101908152919050565b600080858511156103c057600080fd5b838611156103cd57600080fd5b505082019391909203915056fea26469706673582212201fe96e69cacbda1fb92db7799563aad720e0a54993a30f15f3d4907ecd9a769b64736f6c634300080a0033",
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

// BlockHeaderRawDataTest is a free data retrieval call binding the contract method 0x845d481e.
//
// Solidity: function blockHeaderRawDataTest(bytes _data) pure returns(bytes32 allBlockHash, bytes32 blockTxHash)
func (_BlockTest *BlockTestCaller) BlockHeaderRawDataTest(opts *bind.CallOpts, _data []byte) (struct {
	AllBlockHash [32]byte
	BlockTxHash  [32]byte
}, error) {
	var out []interface{}
	err := _BlockTest.contract.Call(opts, &out, "blockHeaderRawDataTest", _data)

	outstruct := new(struct {
		AllBlockHash [32]byte
		BlockTxHash  [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AllBlockHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.BlockTxHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// BlockHeaderRawDataTest is a free data retrieval call binding the contract method 0x845d481e.
//
// Solidity: function blockHeaderRawDataTest(bytes _data) pure returns(bytes32 allBlockHash, bytes32 blockTxHash)
func (_BlockTest *BlockTestSession) BlockHeaderRawDataTest(_data []byte) (struct {
	AllBlockHash [32]byte
	BlockTxHash  [32]byte
}, error) {
	return _BlockTest.Contract.BlockHeaderRawDataTest(&_BlockTest.CallOpts, _data)
}

// BlockHeaderRawDataTest is a free data retrieval call binding the contract method 0x845d481e.
//
// Solidity: function blockHeaderRawDataTest(bytes _data) pure returns(bytes32 allBlockHash, bytes32 blockTxHash)
func (_BlockTest *BlockTestCallerSession) BlockHeaderRawDataTest(_data []byte) (struct {
	AllBlockHash [32]byte
	BlockTxHash  [32]byte
}, error) {
	return _BlockTest.Contract.BlockHeaderRawDataTest(&_BlockTest.CallOpts, _data)
}
