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

// MerkleTestMetaData contains all meta data concerning the MerkleTest contract.

var MerkleTestMetaData = &bind.MetaData{

	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_auditPath\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"blockMerkleProveTest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"reqId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bridgeFrom\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"sel\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_auditPath\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"merkleProveTest\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610cac806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806358066ece1461003b57806368a246c814610064575b600080fd5b61004e610049366004610a00565b610087565b60405161005b9190610b11565b60405180910390f35b610077610072366004610a00565b61009a565b60405161005b9493929190610b24565b606061009383836100c9565b9392505050565b6000806000606060006100ad87876100c9565b90506100b881610201565b929a91995097509095509350505050565b6060600060606100d9858361024f565b9250905060006100e88261035d565b9050600060218488516100fb9190610b71565b6101059190610b88565b905060008060005b838110156101b15761011f8a886103d5565b9750915061012d8a88610490565b9750925060ff821661014a57610143838661051e565b945061019f565b8160ff166001141561016057610143858461051e565b60405162461bcd60e51b815260206004820152600f60248201526e1b595c9adb19541c9bdd9948195bd9608a1b60448201526064015b60405180910390fd5b806101a981610baa565b91505061010d565b508784146101f45760405162461bcd60e51b815260206004820152601060248201526f1b595c9adb19541c9bdd99481c9bdbdd60821b6044820152606401610196565b5092979650505050505050565b6000806000606060006102148682610490565b90955090506102238682610490565b9094509050610232868261059c565b9093509050610241868261024f565b509496939550919392915050565b606060008061025e8585610634565b86519095509091506102708286610bc5565b1115801561028657506102838185610bc5565b84105b6102de5760405162461bcd60e51b8152602060048201526024808201527f4e65787456617242797465732c206f66667365742065786365656473206d6178604482015263696d756d60e01b6064820152608401610196565b6060811580156102f957604051915060208201604052610343565b6040519150601f8316801560200281840101848101888315602002848c0101015b8183101561033257805183526020928301920161031a565b5050848452601f01601f1916604052505b508061034f8387610bc5565b9350935050505b9250929050565b60006002600083604051602001610375929190610bdd565b60408051601f198184030181529082905261038f91610c0c565b602060405180830381855afa1580156103ac573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906103cf9190610c28565b92915050565b60008083518360016103e79190610bc5565b111580156103fe57506103fb836001610bc5565b83105b6104545760405162461bcd60e51b815260206004820152602160248201527f4e65787455696e74382c204f66667365742065786365656473206d6178696d756044820152606d60f81b6064820152608401610196565b6000604051846020870101518060001a82535060018101604052601f810351915050808460016104849190610bc5565b92509250509250929050565b60008083518360206104a29190610bc5565b111580156104b957506104b6836020610bc5565b83105b6105055760405162461bcd60e51b815260206004820181905260248201527f4e657874486173682c206f66667365742065786365656473206d6178696d756d6044820152606401610196565b6000602084018501519050808460206104849190610bc5565b604051600160f81b6020820152602181018390526041810182905260009060029060610160408051601f198184030181529082905261055c91610c0c565b602060405180830381855afa158015610579573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906100939190610c28565b60008083518360146105ae9190610bc5565b111580156105c557506105c2836014610bc5565b83105b61061d5760405162461bcd60e51b815260206004820152602360248201527f4e657874416464726573732c206f66667365742065786365656473206d6178696044820152626d756d60e81b6064820152608401610196565b83830160200151606081901c610484856014610bc5565b600080600061064385856103d5565b94509050600060fd60ff831614156106d25761065f868661078f565b955061ffff16905060fd811080159061067a575061ffff8111155b6106c65760405162461bcd60e51b815260206004820152601f60248201527f4e65787455696e7431362c2076616c7565206f7574736964652072616e6765006044820152606401610196565b92508391506103569050565b8160ff1660fe1415610723576106e88686610848565b955063ffffffff16905061ffff81118015610707575063ffffffff8111155b6106c65760405162461bcd60e51b815260040161019690610c41565b8160ff1660ff141561076a576107398686610919565b955067ffffffffffffffff16905063ffffffff81116106c65760405162461bcd60e51b815260040161019690610c41565b5060ff811660fd81106106c65760405162461bcd60e51b815260040161019690610c41565b60008083518360026107a19190610bc5565b111580156107b857506107b5836002610bc5565b83105b61080f5760405162461bcd60e51b815260206004820152602260248201527f4e65787455696e7431362c206f66667365742065786365656473206d6178696d604482015261756d60f01b6064820152608401610196565b6000604051846020870101518060011a82538060001a60018301535060028101604052601e810351915050808460026104849190610bc5565b600080835183600461085a9190610bc5565b11158015610871575061086e836004610bc5565b83105b6108c85760405162461bcd60e51b815260206004820152602260248201527f4e65787455696e7433322c206f66667365742065786365656473206d6178696d604482015261756d60f01b6064820152608401610196565b600060405160046000600182038760208a0101515b838310156108fd5780821a838601536001830192506001820391506108dd565b505050016040819052601f190151905080610484856004610bc5565b600080835183600861092b9190610bc5565b11158015610942575061093f836008610bc5565b83105b6109995760405162461bcd60e51b815260206004820152602260248201527f4e65787455696e7436342c206f66667365742065786365656473206d6178696d604482015261756d60f01b6064820152608401610196565b600060405160086000600182038760208a0101515b838310156109ce5780821a838601536001830192506001820391506109ae565b505050016040819052601f190151905080610484856008610bc5565b634e487b7160e01b600052604160045260246000fd5b60008060408385031215610a1357600080fd5b823567ffffffffffffffff80821115610a2b57600080fd5b818501915085601f830112610a3f57600080fd5b813581811115610a5157610a516109ea565b604051601f8201601f19908116603f01168101908382118183101715610a7957610a796109ea565b81604052828152886020848701011115610a9257600080fd5b826020860160208301376000602093820184015298969091013596505050505050565b60005b83811015610ad0578181015183820152602001610ab8565b83811115610adf576000848401525b50505050565b60008151808452610afd816020860160208601610ab5565b601f01601f19169290920160200192915050565b6020815260006100936020830184610ae5565b84815283602082015260018060a01b0383166040820152608060608201526000610b516080830184610ae5565b9695505050505050565b634e487b7160e01b600052601160045260246000fd5b600082821015610b8357610b83610b5b565b500390565b600082610ba557634e487b7160e01b600052601260045260246000fd5b500490565b6000600019821415610bbe57610bbe610b5b565b5060010190565b60008219821115610bd857610bd8610b5b565b500190565b60ff60f81b8360f81b16815260008251610bfe816001850160208701610ab5565b919091016001019392505050565b60008251610c1e818460208701610ab5565b9190910192915050565b600060208284031215610c3a57600080fd5b5051919050565b6020808252818101527f4e65787456617255696e742c2076616c7565206f7574736964652072616e676560408201526060019056fea26469706673582212203ddff05258ad504efa62a1a7bcfad7fa31737d43a5a44b441c4f87284b398a5f64736f6c634300080a0033",
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
// Solidity: function blockMerkleProveTest(bytes _auditPath, bytes32 _root) pure returns(bytes32 reqId, bytes32 bridgeFrom, address receiveSide, bytes sel)
func (_MerkleTest *MerkleTestCaller) BlockMerkleProveTest(opts *bind.CallOpts, _auditPath []byte, _root [32]byte) (struct {
	ReqId       [32]byte
	BridgeFrom  [32]byte
	ReceiveSide common.Address
	Sel         []byte
}, error) {
	var out []interface{}
	err := _MerkleTest.contract.Call(opts, &out, "blockMerkleProveTest", _auditPath, _root)

	outstruct := new(struct {
		ReqId       [32]byte
		BridgeFrom  [32]byte
		ReceiveSide common.Address
		Sel         []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReqId = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.BridgeFrom = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.ReceiveSide = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Sel = *abi.ConvertType(out[3], new([]byte)).(*[]byte)

	return *outstruct, err

}

// BlockMerkleProveTest is a free data retrieval call binding the contract method 0x68a246c8.
//
// Solidity: function blockMerkleProveTest(bytes _auditPath, bytes32 _root) pure returns(bytes32 reqId, bytes32 bridgeFrom, address receiveSide, bytes sel)
func (_MerkleTest *MerkleTestSession) BlockMerkleProveTest(_auditPath []byte, _root [32]byte) (struct {
	ReqId       [32]byte
	BridgeFrom  [32]byte
	ReceiveSide common.Address
	Sel         []byte
}, error) {
	return _MerkleTest.Contract.BlockMerkleProveTest(&_MerkleTest.CallOpts, _auditPath, _root)
}

// BlockMerkleProveTest is a free data retrieval call binding the contract method 0x68a246c8.
//
// Solidity: function blockMerkleProveTest(bytes _auditPath, bytes32 _root) pure returns(bytes32 reqId, bytes32 bridgeFrom, address receiveSide, bytes sel)
func (_MerkleTest *MerkleTestCallerSession) BlockMerkleProveTest(_auditPath []byte, _root [32]byte) (struct {
	ReqId       [32]byte
	BridgeFrom  [32]byte
	ReceiveSide common.Address
	Sel         []byte
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
