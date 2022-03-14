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

// MerkleTestMetaData contains all meta data concerning the MerkleTest contract.
var MerkleTestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_auditPath\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"blockMerkleProveTest\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_auditPath\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"merkleProveTest\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610c9b806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806358066ece1461003b57806368a246c814610064575b600080fd5b61004e6100493660046109ea565b610087565b60405161005b9190610afb565b60405180910390f35b6100776100723660046109ea565b61009a565b60405161005b9493929190610b0e565b60606100938383610105565b9392505050565b60008060606000806100ac8787610105565b905060008080606060006100c0868661023d565b955093506100ce86866102de565b955092506100dc868661036c565b955091506100ea868661023d565b50606094851c9d939c50919a505090911c9650945050505050565b606060006060610115858361036c565b92509050600061012482610479565b9050600060218488516101379190610b60565b6101419190610b77565b905060008060005b838110156101ed5761015b8a886104f1565b975091506101698a886102de565b9750925060ff82166101865761017f83866105a0565b94506101db565b8160ff166001141561019c5761017f85846105a0565b60405162461bcd60e51b815260206004820152600f60248201526e1b595c9adb19541c9bdd9948195bd9608a1b60448201526064015b60405180910390fd5b806101e581610b99565b915050610149565b508784146102305760405162461bcd60e51b815260206004820152601060248201526f1b595c9adb19541c9bdd99481c9bdbdd60821b60448201526064016101d2565b5092979650505050505050565b600080835183601461024f9190610bb4565b111580156102665750610263836014610bb4565b83105b6102be5760405162461bcd60e51b815260206004820152602360248201527f4e657874427974657332302c206f66667365742065786365656473206d6178696044820152626d756d60e81b60648201526084016101d2565b83830160200151806102d1856014610bb4565b92509250505b9250929050565b60008083518360206102f09190610bb4565b111580156103075750610304836020610bb4565b83105b6103535760405162461bcd60e51b815260206004820181905260248201527f4e657874486173682c206f66667365742065786365656473206d6178696d756d60448201526064016101d2565b6000602084018501519050808460206102d19190610bb4565b606060008061037b858561061e565b865190955090915061038d8286610bb4565b111580156103a357506103a08185610bb4565b84105b6103fb5760405162461bcd60e51b8152602060048201526024808201527f4e65787456617242797465732c206f66667365742065786365656473206d6178604482015263696d756d60e01b60648201526084016101d2565b60608115801561041657604051915060208201604052610460565b6040519150601f8316801560200281840101848101888315602002848c0101015b8183101561044f578051835260209283019201610437565b5050848452601f01601f1916604052505b508061046c8387610bb4565b9350935050509250929050565b60006002600083604051602001610491929190610bcc565b60408051601f19818403018152908290526104ab91610bfb565b602060405180830381855afa1580156104c8573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906104eb9190610c17565b92915050565b60008083518360016105039190610bb4565b1115801561051a5750610517836001610bb4565b83105b6105705760405162461bcd60e51b815260206004820152602160248201527f4e65787455696e74382c204f66667365742065786365656473206d6178696d756044820152606d60f81b60648201526084016101d2565b6000604051846020870101518060001a82535060018101604052601f810351915050808460016102d19190610bb4565b604051600160f81b6020820152602181018390526041810182905260009060029060610160408051601f19818403018152908290526105de91610bfb565b602060405180830381855afa1580156105fb573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906100939190610c17565b600080600061062d85856104f1565b94509050600060fd60ff831614156106bc576106498686610779565b955061ffff16905060fd8110801590610664575061ffff8111155b6106b05760405162461bcd60e51b815260206004820152601f60248201527f4e65787455696e7431362c2076616c7565206f7574736964652072616e67650060448201526064016101d2565b92508391506102d79050565b8160ff1660fe141561070d576106d28686610832565b955063ffffffff16905061ffff811180156106f1575063ffffffff8111155b6106b05760405162461bcd60e51b81526004016101d290610c30565b8160ff1660ff1415610754576107238686610903565b955067ffffffffffffffff16905063ffffffff81116106b05760405162461bcd60e51b81526004016101d290610c30565b5060ff811660fd81106106b05760405162461bcd60e51b81526004016101d290610c30565b600080835183600261078b9190610bb4565b111580156107a2575061079f836002610bb4565b83105b6107f95760405162461bcd60e51b815260206004820152602260248201527f4e65787455696e7431362c206f66667365742065786365656473206d6178696d604482015261756d60f01b60648201526084016101d2565b6000604051846020870101518060011a82538060001a60018301535060028101604052601e810351915050808460026102d19190610bb4565b60008083518360046108449190610bb4565b1115801561085b5750610858836004610bb4565b83105b6108b25760405162461bcd60e51b815260206004820152602260248201527f4e65787455696e7433322c206f66667365742065786365656473206d6178696d604482015261756d60f01b60648201526084016101d2565b600060405160046000600182038760208a0101515b838310156108e75780821a838601536001830192506001820391506108c7565b505050016040819052601f1901519050806102d1856004610bb4565b60008083518360086109159190610bb4565b1115801561092c5750610929836008610bb4565b83105b6109835760405162461bcd60e51b815260206004820152602260248201527f4e65787455696e7436342c206f66667365742065786365656473206d6178696d604482015261756d60f01b60648201526084016101d2565b600060405160086000600182038760208a0101515b838310156109b85780821a83860153600183019250600182039150610998565b505050016040819052601f1901519050806102d1856008610bb4565b634e487b7160e01b600052604160045260246000fd5b600080604083850312156109fd57600080fd5b823567ffffffffffffffff80821115610a1557600080fd5b818501915085601f830112610a2957600080fd5b813581811115610a3b57610a3b6109d4565b604051601f8201601f19908116603f01168101908382118183101715610a6357610a636109d4565b81604052828152886020848701011115610a7c57600080fd5b826020860160208301376000602093820184015298969091013596505050505050565b60005b83811015610aba578181015183820152602001610aa2565b83811115610ac9576000848401525b50505050565b60008151808452610ae7816020860160208601610a9f565b601f01601f19169290920160200192915050565b6020815260006100936020830184610acf565b600060018060a01b03808716835285602084015260806040840152610b366080840186610acf565b915080841660608401525095945050505050565b634e487b7160e01b600052601160045260246000fd5b600082821015610b7257610b72610b4a565b500390565b600082610b9457634e487b7160e01b600052601260045260246000fd5b500490565b6000600019821415610bad57610bad610b4a565b5060010190565b60008219821115610bc757610bc7610b4a565b500190565b60ff60f81b8360f81b16815260008251610bed816001850160208701610a9f565b919091016001019392505050565b60008251610c0d818460208701610a9f565b9190910192915050565b600060208284031215610c2957600080fd5b5051919050565b6020808252818101527f4e65787456617255696e742c2076616c7565206f7574736964652072616e676560408201526060019056fea26469706673582212202528395180b501cd6441a8ba32e4a55a90faf13d7cdec75193ac67a4a92377b464736f6c634300080a0033",
}

// MerkleTestABI is the input ABI used to generate the binding from.
// Deprecated: Use MerkleTestMetaData.ABI instead.
var MerkleTestABI = MerkleTestMetaData.ABI

// MerkleTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MerkleTestMetaData.Bin instead.
var MerkleTestBin = MerkleTestMetaData.Bin

// DeployMerkleTest deploys a new Ethereum contract, binding an instance of MerkleTest to it.
func DeployMerkleTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MerkleTest, error) {
	parsed, err := MerkleTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MerkleTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MerkleTest{MerkleTestCaller: MerkleTestCaller{contract: contract}, MerkleTestTransactor: MerkleTestTransactor{contract: contract}, MerkleTestFilterer: MerkleTestFilterer{contract: contract}}, nil
}

// MerkleTest is an auto generated Go binding around an Ethereum contract.
type MerkleTest struct {
	MerkleTestCaller     // Read-only binding to the contract
	MerkleTestTransactor // Write-only binding to the contract
	MerkleTestFilterer   // Log filterer for contract events
}

// MerkleTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
	gsn      *GsnCallOpts
}

func (_MerkleTest *MerkleTestTransactor) SetGSNOptions(opts *GsnCallOpts) {
	_MerkleTest.gsn = opts
}

// MerkleTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleTestSession struct {
	Contract     *MerkleTest       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleTestCallerSession struct {
	Contract *MerkleTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MerkleTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleTestTransactorSession struct {
	Contract     *MerkleTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MerkleTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleTestRaw struct {
	Contract *MerkleTest // Generic contract binding to access the raw methods on
}

// MerkleTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleTestCallerRaw struct {
	Contract *MerkleTestCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleTestTransactorRaw struct {
	Contract *MerkleTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleTest creates a new instance of MerkleTest, bound to a specific deployed contract.
func NewMerkleTest(address common.Address, backend bind.ContractBackend) (*MerkleTest, error) {
	contract, err := bindMerkleTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleTest{MerkleTestCaller: MerkleTestCaller{contract: contract}, MerkleTestTransactor: MerkleTestTransactor{contract: contract}, MerkleTestFilterer: MerkleTestFilterer{contract: contract}}, nil
}

// NewMerkleTestCaller creates a new read-only instance of MerkleTest, bound to a specific deployed contract.
func NewMerkleTestCaller(address common.Address, caller bind.ContractCaller) (*MerkleTestCaller, error) {
	contract, err := bindMerkleTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleTestCaller{contract: contract}, nil
}

// NewMerkleTestTransactor creates a new write-only instance of MerkleTest, bound to a specific deployed contract.
func NewMerkleTestTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleTestTransactor, error) {
	contract, err := bindMerkleTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleTestTransactor{contract: contract}, nil
}

// NewMerkleTestFilterer creates a new log filterer instance of MerkleTest, bound to a specific deployed contract.
func NewMerkleTestFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleTestFilterer, error) {
	contract, err := bindMerkleTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleTestFilterer{contract: contract}, nil
}

// bindMerkleTest binds a generic wrapper to an already deployed contract.
func bindMerkleTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleTest *MerkleTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleTest.Contract.MerkleTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleTest *MerkleTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleTest.Contract.MerkleTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleTest *MerkleTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleTest.Contract.MerkleTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleTest *MerkleTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleTest *MerkleTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleTest *MerkleTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleTest.Contract.contract.Transact(opts, method, params...)
}

// BlockMerkleProveTest is a free data retrieval call binding the contract method 0x68a246c8.
//
// Solidity: function blockMerkleProveTest(bytes _auditPath, bytes32 _root) pure returns(address, bytes32, bytes, address)
func (_MerkleTest *MerkleTestCaller) BlockMerkleProveTest(opts *bind.CallOpts, _auditPath []byte, _root [32]byte) (common.Address, [32]byte, []byte, common.Address, error) {
	var out []interface{}
	err := _MerkleTest.contract.Call(opts, &out, "blockMerkleProveTest", _auditPath, _root)

	if err != nil {
		return *new(common.Address), *new([32]byte), *new([]byte), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	out2 := *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	out3 := *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return out0, out1, out2, out3, err

}

// BlockMerkleProveTest is a free data retrieval call binding the contract method 0x68a246c8.
//
// Solidity: function blockMerkleProveTest(bytes _auditPath, bytes32 _root) pure returns(address, bytes32, bytes, address)
func (_MerkleTest *MerkleTestSession) BlockMerkleProveTest(_auditPath []byte, _root [32]byte) (common.Address, [32]byte, []byte, common.Address, error) {
	return _MerkleTest.Contract.BlockMerkleProveTest(&_MerkleTest.CallOpts, _auditPath, _root)
}

// BlockMerkleProveTest is a free data retrieval call binding the contract method 0x68a246c8.
//
// Solidity: function blockMerkleProveTest(bytes _auditPath, bytes32 _root) pure returns(address, bytes32, bytes, address)
func (_MerkleTest *MerkleTestCallerSession) BlockMerkleProveTest(_auditPath []byte, _root [32]byte) (common.Address, [32]byte, []byte, common.Address, error) {
	return _MerkleTest.Contract.BlockMerkleProveTest(&_MerkleTest.CallOpts, _auditPath, _root)
}

// MerkleProveTest is a free data retrieval call binding the contract method 0x58066ece.
//
// Solidity: function merkleProveTest(bytes _auditPath, bytes32 _root) pure returns(bytes)
func (_MerkleTest *MerkleTestCaller) MerkleProveTest(opts *bind.CallOpts, _auditPath []byte, _root [32]byte) ([]byte, error) {
	var out []interface{}
	err := _MerkleTest.contract.Call(opts, &out, "merkleProveTest", _auditPath, _root)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// MerkleProveTest is a free data retrieval call binding the contract method 0x58066ece.
//
// Solidity: function merkleProveTest(bytes _auditPath, bytes32 _root) pure returns(bytes)
func (_MerkleTest *MerkleTestSession) MerkleProveTest(_auditPath []byte, _root [32]byte) ([]byte, error) {
	return _MerkleTest.Contract.MerkleProveTest(&_MerkleTest.CallOpts, _auditPath, _root)
}

// MerkleProveTest is a free data retrieval call binding the contract method 0x58066ece.
//
// Solidity: function merkleProveTest(bytes _auditPath, bytes32 _root) pure returns(bytes)
func (_MerkleTest *MerkleTestCallerSession) MerkleProveTest(_auditPath []byte, _root [32]byte) ([]byte, error) {
	return _MerkleTest.Contract.MerkleProveTest(&_MerkleTest.CallOpts, _auditPath, _root)
}
