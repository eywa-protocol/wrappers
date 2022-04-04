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

// ISynthesisMetaData contains all meta data concerning the ISynthesis contract.
var ISynthesisMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stoken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"burnSyntheticToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stoken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"pubkeys\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"burnSyntheticTokenToSolana\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"emergencyUnsyntesizeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"pubkeys\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes1\",\"name\":\"bumpSynthesizeRequest\",\"type\":\"bytes1\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"emergencyUnsyntesizeRequestToSolana\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_rtoken\",\"type\":\"bytes32\"}],\"name\":\"getRepresentation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tokenReal\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structISynthesis.SynthParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"synthTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ISynthesisABI is the input ABI used to generate the binding from.
// Deprecated: Use ISynthesisMetaData.ABI instead.
var ISynthesisABI = ISynthesisMetaData.ABI

// ISynthesis is an auto generated Go binding around an Ethereum contract.
type ISynthesis struct {
	ISynthesisCaller     // Read-only binding to the contract
	ISynthesisTransactor // Write-only binding to the contract
	ISynthesisFilterer   // Log filterer for contract events
}

// ISynthesisCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISynthesisCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynthesisTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISynthesisTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
	gsn      *GsnCallOpts
}

func (_ISynthesis *ISynthesisTransactor) SetGSNOptions(opts *GsnCallOpts) {
	_ISynthesis.gsn = opts
}

// ISynthesisFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISynthesisFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynthesisSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISynthesisSession struct {
	Contract     *ISynthesis       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISynthesisCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISynthesisCallerSession struct {
	Contract *ISynthesisCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ISynthesisTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISynthesisTransactorSession struct {
	Contract     *ISynthesisTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ISynthesisRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISynthesisRaw struct {
	Contract *ISynthesis // Generic contract binding to access the raw methods on
}

// ISynthesisCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISynthesisCallerRaw struct {
	Contract *ISynthesisCaller // Generic read-only contract binding to access the raw methods on
}

// ISynthesisTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISynthesisTransactorRaw struct {
	Contract *ISynthesisTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISynthesis creates a new instance of ISynthesis, bound to a specific deployed contract.
func NewISynthesis(address common.Address, backend bind.ContractBackend) (*ISynthesis, error) {
	contract, err := bindISynthesis(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISynthesis{ISynthesisCaller: ISynthesisCaller{contract: contract}, ISynthesisTransactor: ISynthesisTransactor{contract: contract}, ISynthesisFilterer: ISynthesisFilterer{contract: contract}}, nil
}

// NewISynthesisCaller creates a new read-only instance of ISynthesis, bound to a specific deployed contract.
func NewISynthesisCaller(address common.Address, caller bind.ContractCaller) (*ISynthesisCaller, error) {
	contract, err := bindISynthesis(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISynthesisCaller{contract: contract}, nil
}

// NewISynthesisTransactor creates a new write-only instance of ISynthesis, bound to a specific deployed contract.
func NewISynthesisTransactor(address common.Address, transactor bind.ContractTransactor) (*ISynthesisTransactor, error) {
	contract, err := bindISynthesis(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISynthesisTransactor{contract: contract}, nil
}

// NewISynthesisFilterer creates a new log filterer instance of ISynthesis, bound to a specific deployed contract.
func NewISynthesisFilterer(address common.Address, filterer bind.ContractFilterer) (*ISynthesisFilterer, error) {
	contract, err := bindISynthesis(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISynthesisFilterer{contract: contract}, nil
}

// bindISynthesis binds a generic wrapper to an already deployed contract.
func bindISynthesis(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISynthesisABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISynthesis *ISynthesisRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISynthesis.Contract.ISynthesisCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISynthesis *ISynthesisRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynthesis.Contract.ISynthesisTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISynthesis *ISynthesisRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISynthesis.Contract.ISynthesisTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISynthesis *ISynthesisCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISynthesis.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISynthesis *ISynthesisTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynthesis.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISynthesis *ISynthesisTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISynthesis.Contract.contract.Transact(opts, method, params...)
}

// GetRepresentation is a free data retrieval call binding the contract method 0xbca73823.
//
// Solidity: function getRepresentation(bytes32 _rtoken) view returns(address)
func (_ISynthesis *ISynthesisCaller) GetRepresentation(opts *bind.CallOpts, _rtoken [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ISynthesis.contract.Call(opts, &out, "getRepresentation", _rtoken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRepresentation is a free data retrieval call binding the contract method 0xbca73823.
//
// Solidity: function getRepresentation(bytes32 _rtoken) view returns(address)
func (_ISynthesis *ISynthesisSession) GetRepresentation(_rtoken [32]byte) (common.Address, error) {
	return _ISynthesis.Contract.GetRepresentation(&_ISynthesis.CallOpts, _rtoken)
}

// GetRepresentation is a free data retrieval call binding the contract method 0xbca73823.
//
// Solidity: function getRepresentation(bytes32 _rtoken) view returns(address)
func (_ISynthesis *ISynthesisCallerSession) GetRepresentation(_rtoken [32]byte) (common.Address, error) {
	return _ISynthesis.Contract.GetRepresentation(&_ISynthesis.CallOpts, _rtoken)
}

// BurnSyntheticToken is a paid mutator transaction binding the contract method 0xbeee7378.
//
// Solidity: function burnSyntheticToken(address stoken, uint256 amount, address from, address to, address receiveSide, address oppositeBridge, uint256 chainId) returns()
func (_ISynthesis *ISynthesisTransactor) BurnSyntheticToken(opts *bind.TransactOpts, stoken common.Address, amount *big.Int, from common.Address, to common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _ISynthesis.contract.Transact(opts, "burnSyntheticToken", stoken, amount, from, to, receiveSide, oppositeBridge, chainId)
}
func (_ISynthesis *ISynthesisTransactor) BurnSyntheticTokenOverGsn(opts *bind.TransactOpts, stoken common.Address, amount *big.Int, from common.Address, to common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (common.Hash, error) {
	return GsnExecutor(_ISynthesis.gsn, ISynthesisMetaData.ABI, "burnSyntheticToken", stoken, amount, from, to, receiveSide, oppositeBridge, chainId)
}

// BurnSyntheticToken is a paid mutator transaction binding the contract method 0xbeee7378.
//
// Solidity: function burnSyntheticToken(address stoken, uint256 amount, address from, address to, address receiveSide, address oppositeBridge, uint256 chainId) returns()
func (_ISynthesis *ISynthesisSession) BurnSyntheticToken(stoken common.Address, amount *big.Int, from common.Address, to common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _ISynthesis.Contract.BurnSyntheticToken(&_ISynthesis.TransactOpts, stoken, amount, from, to, receiveSide, oppositeBridge, chainId)
}
func (_ISynthesis *ISynthesisSession) BurnSyntheticTokenOverGsn(stoken common.Address, amount *big.Int, from common.Address, to common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (common.Hash, error) {
	return _ISynthesis.Contract.BurnSyntheticTokenOverGsn(&_ISynthesis.TransactOpts, stoken, amount, from, to, receiveSide, oppositeBridge, chainId)
}

// BurnSyntheticToken is a paid mutator transaction binding the contract method 0xbeee7378.
//
// Solidity: function burnSyntheticToken(address stoken, uint256 amount, address from, address to, address receiveSide, address oppositeBridge, uint256 chainId) returns()
func (_ISynthesis *ISynthesisTransactorSession) BurnSyntheticToken(stoken common.Address, amount *big.Int, from common.Address, to common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _ISynthesis.Contract.BurnSyntheticToken(&_ISynthesis.TransactOpts, stoken, amount, from, to, receiveSide, oppositeBridge, chainId)
}
func (_ISynthesis *ISynthesisTransactorSession) BurnSyntheticTokenOverGsn(stoken common.Address, amount *big.Int, from common.Address, to common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (common.Hash, error) {
	return _ISynthesis.Contract.BurnSyntheticTokenOverGsn(&_ISynthesis.TransactOpts, stoken, amount, from, to, receiveSide, oppositeBridge, chainId)
}

// BurnSyntheticTokenToSolana is a paid mutator transaction binding the contract method 0x966e0fe9.
//
// Solidity: function burnSyntheticTokenToSolana(address stoken, address from, bytes32[] pubkeys, uint256 amount, uint256 chainId) returns()
func (_ISynthesis *ISynthesisTransactor) BurnSyntheticTokenToSolana(opts *bind.TransactOpts, stoken common.Address, from common.Address, pubkeys [][32]byte, amount *big.Int, chainId *big.Int) (*types.Transaction, error) {
	return _ISynthesis.contract.Transact(opts, "burnSyntheticTokenToSolana", stoken, from, pubkeys, amount, chainId)
}
func (_ISynthesis *ISynthesisTransactor) BurnSyntheticTokenToSolanaOverGsn(opts *bind.TransactOpts, stoken common.Address, from common.Address, pubkeys [][32]byte, amount *big.Int, chainId *big.Int) (common.Hash, error) {
	return GsnExecutor(_ISynthesis.gsn, ISynthesisMetaData.ABI, "burnSyntheticTokenToSolana", stoken, from, pubkeys, amount, chainId)
}

// BurnSyntheticTokenToSolana is a paid mutator transaction binding the contract method 0x966e0fe9.
//
// Solidity: function burnSyntheticTokenToSolana(address stoken, address from, bytes32[] pubkeys, uint256 amount, uint256 chainId) returns()
func (_ISynthesis *ISynthesisSession) BurnSyntheticTokenToSolana(stoken common.Address, from common.Address, pubkeys [][32]byte, amount *big.Int, chainId *big.Int) (*types.Transaction, error) {
	return _ISynthesis.Contract.BurnSyntheticTokenToSolana(&_ISynthesis.TransactOpts, stoken, from, pubkeys, amount, chainId)
}
func (_ISynthesis *ISynthesisSession) BurnSyntheticTokenToSolanaOverGsn(stoken common.Address, from common.Address, pubkeys [][32]byte, amount *big.Int, chainId *big.Int) (common.Hash, error) {
	return _ISynthesis.Contract.BurnSyntheticTokenToSolanaOverGsn(&_ISynthesis.TransactOpts, stoken, from, pubkeys, amount, chainId)
}

// BurnSyntheticTokenToSolana is a paid mutator transaction binding the contract method 0x966e0fe9.
//
// Solidity: function burnSyntheticTokenToSolana(address stoken, address from, bytes32[] pubkeys, uint256 amount, uint256 chainId) returns()
func (_ISynthesis *ISynthesisTransactorSession) BurnSyntheticTokenToSolana(stoken common.Address, from common.Address, pubkeys [][32]byte, amount *big.Int, chainId *big.Int) (*types.Transaction, error) {
	return _ISynthesis.Contract.BurnSyntheticTokenToSolana(&_ISynthesis.TransactOpts, stoken, from, pubkeys, amount, chainId)
}
func (_ISynthesis *ISynthesisTransactorSession) BurnSyntheticTokenToSolanaOverGsn(stoken common.Address, from common.Address, pubkeys [][32]byte, amount *big.Int, chainId *big.Int) (common.Hash, error) {
	return _ISynthesis.Contract.BurnSyntheticTokenToSolanaOverGsn(&_ISynthesis.TransactOpts, stoken, from, pubkeys, amount, chainId)
}

// EmergencyUnsyntesizeRequest is a paid mutator transaction binding the contract method 0x9500125b.
//
// Solidity: function emergencyUnsyntesizeRequest(bytes32 txID, address receiveSide, address oppositeBridge, uint256 chainId, uint8 v, bytes32 r, bytes32 s) returns()
func (_ISynthesis *ISynthesisTransactor) EmergencyUnsyntesizeRequest(opts *bind.TransactOpts, txID [32]byte, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ISynthesis.contract.Transact(opts, "emergencyUnsyntesizeRequest", txID, receiveSide, oppositeBridge, chainId, v, r, s)
}
func (_ISynthesis *ISynthesisTransactor) EmergencyUnsyntesizeRequestOverGsn(opts *bind.TransactOpts, txID [32]byte, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, v uint8, r [32]byte, s [32]byte) (common.Hash, error) {
	return GsnExecutor(_ISynthesis.gsn, ISynthesisMetaData.ABI, "emergencyUnsyntesizeRequest", txID, receiveSide, oppositeBridge, chainId, v, r, s)
}

// EmergencyUnsyntesizeRequest is a paid mutator transaction binding the contract method 0x9500125b.
//
// Solidity: function emergencyUnsyntesizeRequest(bytes32 txID, address receiveSide, address oppositeBridge, uint256 chainId, uint8 v, bytes32 r, bytes32 s) returns()
func (_ISynthesis *ISynthesisSession) EmergencyUnsyntesizeRequest(txID [32]byte, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ISynthesis.Contract.EmergencyUnsyntesizeRequest(&_ISynthesis.TransactOpts, txID, receiveSide, oppositeBridge, chainId, v, r, s)
}
func (_ISynthesis *ISynthesisSession) EmergencyUnsyntesizeRequestOverGsn(txID [32]byte, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, v uint8, r [32]byte, s [32]byte) (common.Hash, error) {
	return _ISynthesis.Contract.EmergencyUnsyntesizeRequestOverGsn(&_ISynthesis.TransactOpts, txID, receiveSide, oppositeBridge, chainId, v, r, s)
}

// EmergencyUnsyntesizeRequest is a paid mutator transaction binding the contract method 0x9500125b.
//
// Solidity: function emergencyUnsyntesizeRequest(bytes32 txID, address receiveSide, address oppositeBridge, uint256 chainId, uint8 v, bytes32 r, bytes32 s) returns()
func (_ISynthesis *ISynthesisTransactorSession) EmergencyUnsyntesizeRequest(txID [32]byte, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ISynthesis.Contract.EmergencyUnsyntesizeRequest(&_ISynthesis.TransactOpts, txID, receiveSide, oppositeBridge, chainId, v, r, s)
}
func (_ISynthesis *ISynthesisTransactorSession) EmergencyUnsyntesizeRequestOverGsn(txID [32]byte, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, v uint8, r [32]byte, s [32]byte) (common.Hash, error) {
	return _ISynthesis.Contract.EmergencyUnsyntesizeRequestOverGsn(&_ISynthesis.TransactOpts, txID, receiveSide, oppositeBridge, chainId, v, r, s)
}

// EmergencyUnsyntesizeRequestToSolana is a paid mutator transaction binding the contract method 0x78f986e7.
//
// Solidity: function emergencyUnsyntesizeRequestToSolana(address from, bytes32[] pubkeys, bytes1 bumpSynthesizeRequest, uint256 chainId) returns()
func (_ISynthesis *ISynthesisTransactor) EmergencyUnsyntesizeRequestToSolana(opts *bind.TransactOpts, from common.Address, pubkeys [][32]byte, bumpSynthesizeRequest [1]byte, chainId *big.Int) (*types.Transaction, error) {
	return _ISynthesis.contract.Transact(opts, "emergencyUnsyntesizeRequestToSolana", from, pubkeys, bumpSynthesizeRequest, chainId)
}
func (_ISynthesis *ISynthesisTransactor) EmergencyUnsyntesizeRequestToSolanaOverGsn(opts *bind.TransactOpts, from common.Address, pubkeys [][32]byte, bumpSynthesizeRequest [1]byte, chainId *big.Int) (common.Hash, error) {
	return GsnExecutor(_ISynthesis.gsn, ISynthesisMetaData.ABI, "emergencyUnsyntesizeRequestToSolana", from, pubkeys, bumpSynthesizeRequest, chainId)
}

// EmergencyUnsyntesizeRequestToSolana is a paid mutator transaction binding the contract method 0x78f986e7.
//
// Solidity: function emergencyUnsyntesizeRequestToSolana(address from, bytes32[] pubkeys, bytes1 bumpSynthesizeRequest, uint256 chainId) returns()
func (_ISynthesis *ISynthesisSession) EmergencyUnsyntesizeRequestToSolana(from common.Address, pubkeys [][32]byte, bumpSynthesizeRequest [1]byte, chainId *big.Int) (*types.Transaction, error) {
	return _ISynthesis.Contract.EmergencyUnsyntesizeRequestToSolana(&_ISynthesis.TransactOpts, from, pubkeys, bumpSynthesizeRequest, chainId)
}
func (_ISynthesis *ISynthesisSession) EmergencyUnsyntesizeRequestToSolanaOverGsn(from common.Address, pubkeys [][32]byte, bumpSynthesizeRequest [1]byte, chainId *big.Int) (common.Hash, error) {
	return _ISynthesis.Contract.EmergencyUnsyntesizeRequestToSolanaOverGsn(&_ISynthesis.TransactOpts, from, pubkeys, bumpSynthesizeRequest, chainId)
}

// EmergencyUnsyntesizeRequestToSolana is a paid mutator transaction binding the contract method 0x78f986e7.
//
// Solidity: function emergencyUnsyntesizeRequestToSolana(address from, bytes32[] pubkeys, bytes1 bumpSynthesizeRequest, uint256 chainId) returns()
func (_ISynthesis *ISynthesisTransactorSession) EmergencyUnsyntesizeRequestToSolana(from common.Address, pubkeys [][32]byte, bumpSynthesizeRequest [1]byte, chainId *big.Int) (*types.Transaction, error) {
	return _ISynthesis.Contract.EmergencyUnsyntesizeRequestToSolana(&_ISynthesis.TransactOpts, from, pubkeys, bumpSynthesizeRequest, chainId)
}
func (_ISynthesis *ISynthesisTransactorSession) EmergencyUnsyntesizeRequestToSolanaOverGsn(from common.Address, pubkeys [][32]byte, bumpSynthesizeRequest [1]byte, chainId *big.Int) (common.Hash, error) {
	return _ISynthesis.Contract.EmergencyUnsyntesizeRequestToSolanaOverGsn(&_ISynthesis.TransactOpts, from, pubkeys, bumpSynthesizeRequest, chainId)
}

// SynthTransfer is a paid mutator transaction binding the contract method 0xb04640b7.
//
// Solidity: function synthTransfer(bytes32 tokenReal, uint256 amount, address from, address to, (address,address,uint256) params) returns()
func (_ISynthesis *ISynthesisTransactor) SynthTransfer(opts *bind.TransactOpts, tokenReal [32]byte, amount *big.Int, from common.Address, to common.Address, params ISynthesisSynthParams) (*types.Transaction, error) {
	return _ISynthesis.contract.Transact(opts, "synthTransfer", tokenReal, amount, from, to, params)
}
func (_ISynthesis *ISynthesisTransactor) SynthTransferOverGsn(opts *bind.TransactOpts, tokenReal [32]byte, amount *big.Int, from common.Address, to common.Address, params ISynthesisSynthParams) (common.Hash, error) {
	return GsnExecutor(_ISynthesis.gsn, ISynthesisMetaData.ABI, "synthTransfer", tokenReal, amount, from, to, params)
}

// SynthTransfer is a paid mutator transaction binding the contract method 0xb04640b7.
//
// Solidity: function synthTransfer(bytes32 tokenReal, uint256 amount, address from, address to, (address,address,uint256) params) returns()
func (_ISynthesis *ISynthesisSession) SynthTransfer(tokenReal [32]byte, amount *big.Int, from common.Address, to common.Address, params ISynthesisSynthParams) (*types.Transaction, error) {
	return _ISynthesis.Contract.SynthTransfer(&_ISynthesis.TransactOpts, tokenReal, amount, from, to, params)
}
func (_ISynthesis *ISynthesisSession) SynthTransferOverGsn(tokenReal [32]byte, amount *big.Int, from common.Address, to common.Address, params ISynthesisSynthParams) (common.Hash, error) {
	return _ISynthesis.Contract.SynthTransferOverGsn(&_ISynthesis.TransactOpts, tokenReal, amount, from, to, params)
}

// SynthTransfer is a paid mutator transaction binding the contract method 0xb04640b7.
//
// Solidity: function synthTransfer(bytes32 tokenReal, uint256 amount, address from, address to, (address,address,uint256) params) returns()
func (_ISynthesis *ISynthesisTransactorSession) SynthTransfer(tokenReal [32]byte, amount *big.Int, from common.Address, to common.Address, params ISynthesisSynthParams) (*types.Transaction, error) {
	return _ISynthesis.Contract.SynthTransfer(&_ISynthesis.TransactOpts, tokenReal, amount, from, to, params)
}
func (_ISynthesis *ISynthesisTransactorSession) SynthTransferOverGsn(tokenReal [32]byte, amount *big.Int, from common.Address, to common.Address, params ISynthesisSynthParams) (common.Hash, error) {
	return _ISynthesis.Contract.SynthTransferOverGsn(&_ISynthesis.TransactOpts, tokenReal, amount, from, to, params)
}
