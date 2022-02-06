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
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_auditPath\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"merkleProve\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610b66806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063efd3b91014610030575b600080fd5b61004361003e3660046108fd565b610059565b60405161005091906109e2565b60405180910390f35b60606000606061006985836101dd565b925090506000610078826102eb565b90506000602184885161008b9190610a2b565b6100959190610a42565b905060008060005b83811015610166576100af8a88610363565b975091506100bd8a8861041e565b9750925060ff82166100da576100d383866104ac565b9450610154565b8160ff16600114156100f0576100d385846104ac565b60405162461bcd60e51b815260206004820152602e60248201527f6d65726b6c6550726f76652c204e6578744279746520666f7220706f7369746960448201526d1bdb881a5b999bc819985a5b195960921b60648201526084015b60405180910390fd5b8061015e81610a64565b91505061009d565b508784146101d05760405162461bcd60e51b815260206004820152603160248201527f6d65726b6c6550726f76652c2065787065637420726f6f74206973206e6f7420604482015270195c5d585b081858dd1d585b081c9bdbdd607a1b606482015260840161014b565b5092979650505050505050565b60606000806101ec8585610531565b86519095509091506101fe8286610a7f565b1115801561021457506102118185610a7f565b84105b61026c5760405162461bcd60e51b8152602060048201526024808201527f4e65787456617242797465732c206f66667365742065786365656473206d6178604482015263696d756d60e01b606482015260840161014b565b606081158015610287576040519150602082016040526102d1565b6040519150601f8316801560200281840101848101888315602002848c0101015b818310156102c05780518352602092830192016102a8565b5050848452601f01601f1916604052505b50806102dd8387610a7f565b9350935050505b9250929050565b60006002600083604051602001610303929190610a97565b60408051601f198184030181529082905261031d91610ac6565b602060405180830381855afa15801561033a573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061035d9190610ae2565b92915050565b60008083518360016103759190610a7f565b1115801561038c5750610389836001610a7f565b83105b6103e25760405162461bcd60e51b815260206004820152602160248201527f4e65787455696e74382c204f66667365742065786365656473206d6178696d756044820152606d60f81b606482015260840161014b565b6000604051846020870101518060001a82535060018101604052601f810351915050808460016104129190610a7f565b92509250509250929050565b60008083518360206104309190610a7f565b111580156104475750610444836020610a7f565b83105b6104935760405162461bcd60e51b815260206004820181905260248201527f4e657874486173682c206f66667365742065786365656473206d6178696d756d604482015260640161014b565b6000602084018501519050808460206104129190610a7f565b604051600160f81b6020820152602181018390526041810182905260009060029060610160408051601f19818403018152908290526104ea91610ac6565b602060405180830381855afa158015610507573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061052a9190610ae2565b9392505050565b60008060006105408585610363565b94509050600060fd60ff831614156105cf5761055c868661068c565b955061ffff16905060fd8110801590610577575061ffff8111155b6105c35760405162461bcd60e51b815260206004820152601f60248201527f4e65787455696e7431362c2076616c7565206f7574736964652072616e676500604482015260640161014b565b92508391506102e49050565b8160ff1660fe1415610620576105e58686610745565b955063ffffffff16905061ffff81118015610604575063ffffffff8111155b6105c35760405162461bcd60e51b815260040161014b90610afb565b8160ff1660ff1415610667576106368686610816565b955067ffffffffffffffff16905063ffffffff81116105c35760405162461bcd60e51b815260040161014b90610afb565b5060ff811660fd81106105c35760405162461bcd60e51b815260040161014b90610afb565b600080835183600261069e9190610a7f565b111580156106b557506106b2836002610a7f565b83105b61070c5760405162461bcd60e51b815260206004820152602260248201527f4e65787455696e7431362c206f66667365742065786365656473206d6178696d604482015261756d60f01b606482015260840161014b565b6000604051846020870101518060011a82538060001a60018301535060028101604052601e810351915050808460026104129190610a7f565b60008083518360046107579190610a7f565b1115801561076e575061076b836004610a7f565b83105b6107c55760405162461bcd60e51b815260206004820152602260248201527f4e65787455696e7433322c206f66667365742065786365656473206d6178696d604482015261756d60f01b606482015260840161014b565b600060405160046000600182038760208a0101515b838310156107fa5780821a838601536001830192506001820391506107da565b505050016040819052601f190151905080610412856004610a7f565b60008083518360086108289190610a7f565b1115801561083f575061083c836008610a7f565b83105b6108965760405162461bcd60e51b815260206004820152602260248201527f4e65787455696e7436342c206f66667365742065786365656473206d6178696d604482015261756d60f01b606482015260840161014b565b600060405160086000600182038760208a0101515b838310156108cb5780821a838601536001830192506001820391506108ab565b505050016040819052601f190151905080610412856008610a7f565b634e487b7160e01b600052604160045260246000fd5b6000806040838503121561091057600080fd5b823567ffffffffffffffff8082111561092857600080fd5b818501915085601f83011261093c57600080fd5b81358181111561094e5761094e6108e7565b604051601f8201601f19908116603f01168101908382118183101715610976576109766108e7565b8160405282815288602084870101111561098f57600080fd5b826020860160208301376000602093820184015298969091013596505050505050565b60005b838110156109cd5781810151838201526020016109b5565b838111156109dc576000848401525b50505050565b6020815260008251806020840152610a018160408501602087016109b2565b601f01601f19169190910160400192915050565b634e487b7160e01b600052601160045260246000fd5b600082821015610a3d57610a3d610a15565b500390565b600082610a5f57634e487b7160e01b600052601260045260246000fd5b500490565b6000600019821415610a7857610a78610a15565b5060010190565b60008219821115610a9257610a92610a15565b500190565b60ff60f81b8360f81b16815260008251610ab88160018501602087016109b2565b919091016001019392505050565b60008251610ad88184602087016109b2565b9190910192915050565b600060208284031215610af457600080fd5b5051919050565b6020808252818101527f4e65787456617255696e742c2076616c7565206f7574736964652072616e676560408201526060019056fea2646970667358221220d0175395a438fa0b39f9487f50cc0264961a9ccbca1f826defa8d520aea7d58264736f6c634300080a0033",
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

// MerkleProve is a free data retrieval call binding the contract method 0xefd3b910.
//
// Solidity: function merkleProve(bytes _auditPath, bytes32 _root) pure returns(bytes)
func (_MerkleTest *MerkleTestCaller) MerkleProve(opts *bind.CallOpts, _auditPath []byte, _root [32]byte) ([]byte, error) {
	var out []interface{}
	err := _MerkleTest.contract.Call(opts, &out, "merkleProve", _auditPath, _root)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// MerkleProve is a free data retrieval call binding the contract method 0xefd3b910.
//
// Solidity: function merkleProve(bytes _auditPath, bytes32 _root) pure returns(bytes)
func (_MerkleTest *MerkleTestSession) MerkleProve(_auditPath []byte, _root [32]byte) ([]byte, error) {
	return _MerkleTest.Contract.MerkleProve(&_MerkleTest.CallOpts, _auditPath, _root)
}

// MerkleProve is a free data retrieval call binding the contract method 0xefd3b910.
//
// Solidity: function merkleProve(bytes _auditPath, bytes32 _root) pure returns(bytes)
func (_MerkleTest *MerkleTestCallerSession) MerkleProve(_auditPath []byte, _root [32]byte) ([]byte, error) {
	return _MerkleTest.Contract.MerkleProve(&_MerkleTest.CallOpts, _auditPath, _root)
}
