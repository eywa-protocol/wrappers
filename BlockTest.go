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

// BlockTestMetaData contains all meta data concerning the BlockTest contract.

var BlockTestMetaData = &bind.MetaData{

	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"prevBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"epochBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"transactionsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceHeigh\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"blockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"blockHeaderRawDataTest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"allBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blockTxHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"epochRequestTxRawDataTest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"txNewEpochNum\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"txNewKey\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"txNewEpochParticipantsNum\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"oracleRequestTxRawDataTest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"reqId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bridgeFrom\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"sel\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"solanaRequestTxRawDataTest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"reqId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bridgeFrom\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"oppositeBridge\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sel\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610f87806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80634c08fdf31461005c578063594c73f814610088578063732a56a4146100ac578063845d481e146100cd578063c9698895146100f5575b600080fd5b61006f61006a366004610ca6565b610119565b60405161007f9493929190610d74565b60405180910390f35b61009b610096366004610ca6565b6101dd565b60405161007f959493929190610dad565b6100bf6100ba366004610dfd565b6102a6565b60405190815260200161007f565b6100e06100db366004610ca6565b610354565b6040805192835260208301919091520161007f565b610108610103366004610ca6565b6103db565b60405161007f959493929190610e5c565b6000806060600060028686604051602001610135929190610e87565b60408051601f198184030181529082905261014f91610e97565b602060405180830381855afa15801561016c573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061018f9190610eb3565b93506101d086868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061049392505050565b9598919750955092505050565b6000806000806060600287876040516020016101fa929190610e87565b60408051601f198184030181529082905261021491610e97565b602060405180830381855afa158015610231573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906102549190610eb3565b945061029587878080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506104ce92505050565b979a92995090975095945092505050565b6040516001600160c01b031960c088811b8216602084015260288301889052604883018790526068830186905284811b8216608884015283901b16609082015260009060029060980160408051601f198184030181529082905261030991610e97565b602060405180830381855afa158015610326573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906103499190610eb3565b979650505050505050565b6000806002848460405160200161036c929190610e87565b60408051601f198184030181529082905261038691610e97565b602060405180830381855afa1580156103a3573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906103c69190610eb3565b91506103d2848461051c565b90509250929050565b6000806000806060600287876040516020016103f8929190610e87565b60408051601f198184030181529082905261041291610e97565b602060405180830381855afa15801561042f573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906104529190610eb3565b945061029587878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061057092505050565b6000606081806104a385826105a1565b90945090506104b28582610684565b90925090506104c18582610733565b5093959394509092915050565b6000806000606060006104e18682610840565b90955090506104f08682610840565b90945090506104ff86826108ce565b909350905061050e8682610733565b509496939550919392915050565b600061056961052f606860488587610ecc565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061096692505050565b9392505050565b6000806000606060006105838682610840565b90955090506105928682610840565b90945090506104ff8682610840565b60008083518360046105b39190610ef6565b111580156105ca57506105c7836004610ef6565b83105b6106265760405162461bcd60e51b815260206004820152602260248201527f4e65787455696e7433322c206f66667365742065786365656473206d6178696d604482015261756d60f01b60648201526084015b60405180910390fd5b600060405160046000600182038760208a0101515b8383101561065b5780821a8386015360018301925060018203915061063b565b505050016040819052601f190151905080610677856004610ef6565b92509250505b9250929050565b60008083518360016106969190610ef6565b111580156106ad57506106aa836001610ef6565b83105b6107035760405162461bcd60e51b815260206004820152602160248201527f4e65787455696e74382c204f66667365742065786365656473206d6178696d756044820152606d60f81b606482015260840161061d565b6000604051846020870101518060001a82535060018101604052601f810351915050808460016106779190610ef6565b606060008061074285856109c1565b86519095509091506107548286610ef6565b1115801561076a57506107678185610ef6565b84105b6107c25760405162461bcd60e51b8152602060048201526024808201527f4e65787456617242797465732c206f66667365742065786365656473206d6178604482015263696d756d60e01b606482015260840161061d565b6060811580156107dd57604051915060208201604052610827565b6040519150601f8316801560200281840101848101888315602002848c0101015b818310156108165780518352602092830192016107fe565b5050848452601f01601f1916604052505b50806108338387610ef6565b9350935050509250929050565b60008083518360206108529190610ef6565b111580156108695750610866836020610ef6565b83105b6108b55760405162461bcd60e51b815260206004820181905260248201527f4e657874486173682c206f66667365742065786365656473206d6178696d756d604482015260640161061d565b6000602084018501519050808460206106779190610ef6565b60008083518360146108e09190610ef6565b111580156108f757506108f4836014610ef6565b83105b61094f5760405162461bcd60e51b815260206004820152602360248201527f4e657874416464726573732c206f66667365742065786365656473206d6178696044820152626d756d60e81b606482015260840161061d565b83830160200151606081901c610677856014610ef6565b600081516020146109b95760405162461bcd60e51b815260206004820152601760248201527f6279746573206c656e677468206973206e6f742033322e000000000000000000604482015260640161061d565b506020015190565b60008060006109d08585610684565b94509050600060fd60ff83161415610a5f576109ec8686610b1c565b955061ffff16905060fd8110801590610a07575061ffff8111155b610a535760405162461bcd60e51b815260206004820152601f60248201527f4e65787455696e7431362c2076616c7565206f7574736964652072616e676500604482015260640161061d565b925083915061067d9050565b8160ff1660fe1415610ab057610a7586866105a1565b955063ffffffff16905061ffff81118015610a94575063ffffffff8111155b610a535760405162461bcd60e51b815260040161061d90610f1c565b8160ff1660ff1415610af757610ac68686610bd5565b955067ffffffffffffffff16905063ffffffff8111610a535760405162461bcd60e51b815260040161061d90610f1c565b5060ff811660fd8110610a535760405162461bcd60e51b815260040161061d90610f1c565b6000808351836002610b2e9190610ef6565b11158015610b455750610b42836002610ef6565b83105b610b9c5760405162461bcd60e51b815260206004820152602260248201527f4e65787455696e7431362c206f66667365742065786365656473206d6178696d604482015261756d60f01b606482015260840161061d565b6000604051846020870101518060011a82538060001a60018301535060028101604052601e810351915050808460026106779190610ef6565b6000808351836008610be79190610ef6565b11158015610bfe5750610bfb836008610ef6565b83105b610c555760405162461bcd60e51b815260206004820152602260248201527f4e65787455696e7436342c206f66667365742065786365656473206d6178696d604482015261756d60f01b606482015260840161061d565b600060405160086000600182038760208a0101515b83831015610c8a5780821a83860153600183019250600182039150610c6a565b505050016040819052601f190151905080610677856008610ef6565b60008060208385031215610cb957600080fd5b823567ffffffffffffffff80821115610cd157600080fd5b818501915085601f830112610ce557600080fd5b813581811115610cf457600080fd5b866020828501011115610d0657600080fd5b60209290920196919550909350505050565b60005b83811015610d33578181015183820152602001610d1b565b83811115610d42576000848401525b50505050565b60008151808452610d60816020860160208601610d18565b601f01601f19169290920160200192915050565b84815263ffffffff84166020820152608060408201526000610d996080830185610d48565b905060ff8316606083015295945050505050565b85815284602082015283604082015260018060a01b038316606082015260a06080820152600061034960a0830184610d48565b803567ffffffffffffffff81168114610df857600080fd5b919050565b60008060008060008060c08789031215610e1657600080fd5b610e1f87610de0565b9550602087013594506040870135935060608701359250610e4260808801610de0565b9150610e5060a08801610de0565b90509295509295509295565b85815284602082015283604082015282606082015260a06080820152600061034960a0830184610d48565b8183823760009101908152919050565b60008251610ea9818460208701610d18565b9190910192915050565b600060208284031215610ec557600080fd5b5051919050565b60008085851115610edc57600080fd5b83861115610ee957600080fd5b5050820193919092039150565b60008219821115610f1757634e487b7160e01b600052601160045260246000fd5b500190565b6020808252818101527f4e65787456617255696e742c2076616c7565206f7574736964652072616e676560408201526060019056fea2646970667358221220a606e262078f7d5d3c02762b81fde80a4c13d1591aca572470f653fe6af59b8d64736f6c634300080a0033",
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

// EpochRequestTxRawDataTest is a free data retrieval call binding the contract method 0x4c08fdf3.
//
// Solidity: function epochRequestTxRawDataTest(bytes _data) pure returns(bytes32 txHash, uint32 txNewEpochNum, bytes txNewKey, uint8 txNewEpochParticipantsNum)
func (_BlockTest *BlockTestCaller) EpochRequestTxRawDataTest(opts *bind.CallOpts, _data []byte) (struct {
	TxHash                    [32]byte
	TxNewEpochNum             uint32
	TxNewKey                  []byte
	TxNewEpochParticipantsNum uint8
}, error) {
	var out []interface{}
	err := _BlockTest.contract.Call(opts, &out, "epochRequestTxRawDataTest", _data)

	outstruct := new(struct {
		TxHash                    [32]byte
		TxNewEpochNum             uint32
		TxNewKey                  []byte
		TxNewEpochParticipantsNum uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TxHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.TxNewEpochNum = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.TxNewKey = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.TxNewEpochParticipantsNum = *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return *outstruct, err

}

// EpochRequestTxRawDataTest is a free data retrieval call binding the contract method 0x4c08fdf3.
//
// Solidity: function epochRequestTxRawDataTest(bytes _data) pure returns(bytes32 txHash, uint32 txNewEpochNum, bytes txNewKey, uint8 txNewEpochParticipantsNum)
func (_BlockTest *BlockTestSession) EpochRequestTxRawDataTest(_data []byte) (struct {
	TxHash                    [32]byte
	TxNewEpochNum             uint32
	TxNewKey                  []byte
	TxNewEpochParticipantsNum uint8
}, error) {
	return _BlockTest.Contract.EpochRequestTxRawDataTest(&_BlockTest.CallOpts, _data)
}

// EpochRequestTxRawDataTest is a free data retrieval call binding the contract method 0x4c08fdf3.
//
// Solidity: function epochRequestTxRawDataTest(bytes _data) pure returns(bytes32 txHash, uint32 txNewEpochNum, bytes txNewKey, uint8 txNewEpochParticipantsNum)
func (_BlockTest *BlockTestCallerSession) EpochRequestTxRawDataTest(_data []byte) (struct {
	TxHash                    [32]byte
	TxNewEpochNum             uint32
	TxNewKey                  []byte
	TxNewEpochParticipantsNum uint8
}, error) {
	return _BlockTest.Contract.EpochRequestTxRawDataTest(&_BlockTest.CallOpts, _data)
}

// OracleRequestTxRawDataTest is a free data retrieval call binding the contract method 0x594c73f8.
//
// Solidity: function oracleRequestTxRawDataTest(bytes _data) pure returns(bytes32 txHash, bytes32 reqId, bytes32 bridgeFrom, address receiveSide, bytes sel)
func (_BlockTest *BlockTestCaller) OracleRequestTxRawDataTest(opts *bind.CallOpts, _data []byte) (struct {
	TxHash      [32]byte
	ReqId       [32]byte
	BridgeFrom  [32]byte
	ReceiveSide common.Address
	Sel         []byte
}, error) {
	var out []interface{}
	err := _BlockTest.contract.Call(opts, &out, "oracleRequestTxRawDataTest", _data)

	outstruct := new(struct {
		TxHash      [32]byte
		ReqId       [32]byte
		BridgeFrom  [32]byte
		ReceiveSide common.Address
		Sel         []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TxHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.ReqId = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.BridgeFrom = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.ReceiveSide = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Sel = *abi.ConvertType(out[4], new([]byte)).(*[]byte)

	return *outstruct, err

}

// OracleRequestTxRawDataTest is a free data retrieval call binding the contract method 0x594c73f8.
//
// Solidity: function oracleRequestTxRawDataTest(bytes _data) pure returns(bytes32 txHash, bytes32 reqId, bytes32 bridgeFrom, address receiveSide, bytes sel)
func (_BlockTest *BlockTestSession) OracleRequestTxRawDataTest(_data []byte) (struct {
	TxHash      [32]byte
	ReqId       [32]byte
	BridgeFrom  [32]byte
	ReceiveSide common.Address
	Sel         []byte
}, error) {
	return _BlockTest.Contract.OracleRequestTxRawDataTest(&_BlockTest.CallOpts, _data)
}

// OracleRequestTxRawDataTest is a free data retrieval call binding the contract method 0x594c73f8.
//
// Solidity: function oracleRequestTxRawDataTest(bytes _data) pure returns(bytes32 txHash, bytes32 reqId, bytes32 bridgeFrom, address receiveSide, bytes sel)
func (_BlockTest *BlockTestCallerSession) OracleRequestTxRawDataTest(_data []byte) (struct {
	TxHash      [32]byte
	ReqId       [32]byte
	BridgeFrom  [32]byte
	ReceiveSide common.Address
	Sel         []byte
}, error) {
	return _BlockTest.Contract.OracleRequestTxRawDataTest(&_BlockTest.CallOpts, _data)
}

// SolanaRequestTxRawDataTest is a free data retrieval call binding the contract method 0xc9698895.
//
// Solidity: function solanaRequestTxRawDataTest(bytes _data) pure returns(bytes32 txHash, bytes32 reqId, bytes32 bridgeFrom, bytes32 oppositeBridge, bytes sel)
func (_BlockTest *BlockTestCaller) SolanaRequestTxRawDataTest(opts *bind.CallOpts, _data []byte) (struct {
	TxHash         [32]byte
	ReqId          [32]byte
	BridgeFrom     [32]byte
	OppositeBridge [32]byte
	Sel            []byte
}, error) {
	var out []interface{}
	err := _BlockTest.contract.Call(opts, &out, "solanaRequestTxRawDataTest", _data)

	outstruct := new(struct {
		TxHash         [32]byte
		ReqId          [32]byte
		BridgeFrom     [32]byte
		OppositeBridge [32]byte
		Sel            []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TxHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.ReqId = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.BridgeFrom = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.OppositeBridge = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.Sel = *abi.ConvertType(out[4], new([]byte)).(*[]byte)

	return *outstruct, err

}

// SolanaRequestTxRawDataTest is a free data retrieval call binding the contract method 0xc9698895.
//
// Solidity: function solanaRequestTxRawDataTest(bytes _data) pure returns(bytes32 txHash, bytes32 reqId, bytes32 bridgeFrom, bytes32 oppositeBridge, bytes sel)
func (_BlockTest *BlockTestSession) SolanaRequestTxRawDataTest(_data []byte) (struct {
	TxHash         [32]byte
	ReqId          [32]byte
	BridgeFrom     [32]byte
	OppositeBridge [32]byte
	Sel            []byte
}, error) {
	return _BlockTest.Contract.SolanaRequestTxRawDataTest(&_BlockTest.CallOpts, _data)
}

// SolanaRequestTxRawDataTest is a free data retrieval call binding the contract method 0xc9698895.
//
// Solidity: function solanaRequestTxRawDataTest(bytes _data) pure returns(bytes32 txHash, bytes32 reqId, bytes32 bridgeFrom, bytes32 oppositeBridge, bytes sel)
func (_BlockTest *BlockTestCallerSession) SolanaRequestTxRawDataTest(_data []byte) (struct {
	TxHash         [32]byte
	ReqId          [32]byte
	BridgeFrom     [32]byte
	OppositeBridge [32]byte
	Sel            []byte
}, error) {
	return _BlockTest.Contract.SolanaRequestTxRawDataTest(&_BlockTest.CallOpts, _data)
}
