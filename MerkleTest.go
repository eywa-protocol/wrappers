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
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_auditPath\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"blockMerkleProveTest\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"bridgeFrom\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"reqId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sel\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_auditPath\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"merkleProveTest\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610cae806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806358066ece1461003b57806368a246c814610064575b600080fd5b61004e6100493660046109fd565b610087565b60405161005b9190610b0e565b60405180910390f35b6100776100723660046109fd565b61009a565b60405161005b9493929190610b21565b606061009383836100c8565b9392505050565b60008060606000806100ac87876100c8565b90506100b781610200565b929a91995097509095509350505050565b6060600060606100d8858361024c565b9250905060006100e78261035a565b9050600060218488516100fa9190610b73565b6101049190610b8a565b905060008060005b838110156101b05761011e8a886103d2565b9750915061012c8a8861048d565b9750925060ff821661014957610142838661051b565b945061019e565b8160ff166001141561015f57610142858461051b565b60405162461bcd60e51b815260206004820152600f60248201526e1b595c9adb19541c9bdd9948195bd9608a1b60448201526064015b60405180910390fd5b806101a881610bac565b91505061010c565b508784146101f35760405162461bcd60e51b815260206004820152601060248201526f1b595c9adb19541c9bdd99481c9bdbdd60821b6044820152606401610195565b5092979650505050505050565b600080606081806102118682610599565b9095509050610220868261048d565b909450905061022f868261024c565b909350905061023e8682610599565b509496939550919392915050565b606060008061025b8585610631565b865190955090915061026d8286610bc7565b1115801561028357506102808185610bc7565b84105b6102db5760405162461bcd60e51b8152602060048201526024808201527f4e65787456617242797465732c206f66667365742065786365656473206d6178604482015263696d756d60e01b6064820152608401610195565b6060811580156102f657604051915060208201604052610340565b6040519150601f8316801560200281840101848101888315602002848c0101015b8183101561032f578051835260209283019201610317565b5050848452601f01601f1916604052505b508061034c8387610bc7565b9350935050505b9250929050565b60006002600083604051602001610372929190610bdf565b60408051601f198184030181529082905261038c91610c0e565b602060405180830381855afa1580156103a9573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906103cc9190610c2a565b92915050565b60008083518360016103e49190610bc7565b111580156103fb57506103f8836001610bc7565b83105b6104515760405162461bcd60e51b815260206004820152602160248201527f4e65787455696e74382c204f66667365742065786365656473206d6178696d756044820152606d60f81b6064820152608401610195565b6000604051846020870101518060001a82535060018101604052601f810351915050808460016104819190610bc7565b92509250509250929050565b600080835183602061049f9190610bc7565b111580156104b657506104b3836020610bc7565b83105b6105025760405162461bcd60e51b815260206004820181905260248201527f4e657874486173682c206f66667365742065786365656473206d6178696d756d6044820152606401610195565b6000602084018501519050808460206104819190610bc7565b604051600160f81b6020820152602181018390526041810182905260009060029060610160408051601f198184030181529082905261055991610c0e565b602060405180830381855afa158015610576573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906100939190610c2a565b60008083518360146105ab9190610bc7565b111580156105c257506105bf836014610bc7565b83105b61061a5760405162461bcd60e51b815260206004820152602360248201527f4e657874416464726573732c206f66667365742065786365656473206d6178696044820152626d756d60e81b6064820152608401610195565b83830160200151606081901c610481856014610bc7565b600080600061064085856103d2565b94509050600060fd60ff831614156106cf5761065c868661078c565b955061ffff16905060fd8110801590610677575061ffff8111155b6106c35760405162461bcd60e51b815260206004820152601f60248201527f4e65787455696e7431362c2076616c7565206f7574736964652072616e6765006044820152606401610195565b92508391506103539050565b8160ff1660fe1415610720576106e58686610845565b955063ffffffff16905061ffff81118015610704575063ffffffff8111155b6106c35760405162461bcd60e51b815260040161019590610c43565b8160ff1660ff1415610767576107368686610916565b955067ffffffffffffffff16905063ffffffff81116106c35760405162461bcd60e51b815260040161019590610c43565b5060ff811660fd81106106c35760405162461bcd60e51b815260040161019590610c43565b600080835183600261079e9190610bc7565b111580156107b557506107b2836002610bc7565b83105b61080c5760405162461bcd60e51b815260206004820152602260248201527f4e65787455696e7431362c206f66667365742065786365656473206d6178696d604482015261756d60f01b6064820152608401610195565b6000604051846020870101518060011a82538060001a60018301535060028101604052601e810351915050808460026104819190610bc7565b60008083518360046108579190610bc7565b1115801561086e575061086b836004610bc7565b83105b6108c55760405162461bcd60e51b815260206004820152602260248201527f4e65787455696e7433322c206f66667365742065786365656473206d6178696d604482015261756d60f01b6064820152608401610195565b600060405160046000600182038760208a0101515b838310156108fa5780821a838601536001830192506001820391506108da565b505050016040819052601f190151905080610481856004610bc7565b60008083518360086109289190610bc7565b1115801561093f575061093c836008610bc7565b83105b6109965760405162461bcd60e51b815260206004820152602260248201527f4e65787455696e7436342c206f66667365742065786365656473206d6178696d604482015261756d60f01b6064820152608401610195565b600060405160086000600182038760208a0101515b838310156109cb5780821a838601536001830192506001820391506109ab565b505050016040819052601f190151905080610481856008610bc7565b634e487b7160e01b600052604160045260246000fd5b60008060408385031215610a1057600080fd5b823567ffffffffffffffff80821115610a2857600080fd5b818501915085601f830112610a3c57600080fd5b813581811115610a4e57610a4e6109e7565b604051601f8201601f19908116603f01168101908382118183101715610a7657610a766109e7565b81604052828152886020848701011115610a8f57600080fd5b826020860160208301376000602093820184015298969091013596505050505050565b60005b83811015610acd578181015183820152602001610ab5565b83811115610adc576000848401525b50505050565b60008151808452610afa816020860160208601610ab2565b601f01601f19169290920160200192915050565b6020815260006100936020830184610ae2565b600060018060a01b03808716835285602084015260806040840152610b496080840186610ae2565b915080841660608401525095945050505050565b634e487b7160e01b600052601160045260246000fd5b600082821015610b8557610b85610b5d565b500390565b600082610ba757634e487b7160e01b600052601260045260246000fd5b500490565b6000600019821415610bc057610bc0610b5d565b5060010190565b60008219821115610bda57610bda610b5d565b500190565b60ff60f81b8360f81b16815260008251610c00816001850160208701610ab2565b919091016001019392505050565b60008251610c20818460208701610ab2565b9190910192915050565b600060208284031215610c3c57600080fd5b5051919050565b6020808252818101527f4e65787456617255696e742c2076616c7565206f7574736964652072616e676560408201526060019056fea26469706673582212204d05f5bfdf9cf96cba2d41094ef9dd9a106c9df14fc7eeacefd9f1ce11dff75164736f6c634300080a0033",
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
// Solidity: function blockMerkleProveTest(bytes _auditPath, bytes32 _root) pure returns(address bridgeFrom, bytes32 reqId, bytes sel, address receiveSide)
func (_MerkleTest *MerkleTestCaller) BlockMerkleProveTest(opts *bind.CallOpts, _auditPath []byte, _root [32]byte) (struct {
	BridgeFrom  common.Address
	ReqId       [32]byte
	Sel         []byte
	ReceiveSide common.Address
}, error) {
	var out []interface{}
	err := _MerkleTest.contract.Call(opts, &out, "blockMerkleProveTest", _auditPath, _root)

	outstruct := new(struct {
		BridgeFrom  common.Address
		ReqId       [32]byte
		Sel         []byte
		ReceiveSide common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BridgeFrom = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ReqId = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Sel = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.ReceiveSide = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// BlockMerkleProveTest is a free data retrieval call binding the contract method 0x68a246c8.
//
// Solidity: function blockMerkleProveTest(bytes _auditPath, bytes32 _root) pure returns(address bridgeFrom, bytes32 reqId, bytes sel, address receiveSide)
func (_MerkleTest *MerkleTestSession) BlockMerkleProveTest(_auditPath []byte, _root [32]byte) (struct {
	BridgeFrom  common.Address
	ReqId       [32]byte
	Sel         []byte
	ReceiveSide common.Address
}, error) {
	return _MerkleTest.Contract.BlockMerkleProveTest(&_MerkleTest.CallOpts, _auditPath, _root)
}

// BlockMerkleProveTest is a free data retrieval call binding the contract method 0x68a246c8.
//
// Solidity: function blockMerkleProveTest(bytes _auditPath, bytes32 _root) pure returns(address bridgeFrom, bytes32 reqId, bytes sel, address receiveSide)
func (_MerkleTest *MerkleTestCallerSession) BlockMerkleProveTest(_auditPath []byte, _root [32]byte) (struct {
	BridgeFrom  common.Address
	ReqId       [32]byte
	Sel         []byte
	ReceiveSide common.Address
}, error) {
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
