// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wrappers

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BridgeABI is the input ABI used to generate the binding from.
const BridgeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"listNode\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"requestType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"selector\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainid\",\"type\":\"uint256\"}],\"name\":\"OracleRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"reqId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridgeFrom\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"senderSide\",\"type\":\"address\"}],\"name\":\"ReceiveRequest\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_listNode\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"addContractBind\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reqId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"b\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridgeFrom\",\"type\":\"address\"}],\"name\":\"receiveRequestV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_selector\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"transmitRequestV2\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedForwarder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionRecipient\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BridgeBin is the compiled bytecode used for deploying new contracts.
var BridgeBin = "0x60c06040526005608081905264322e322e3360d81b60a09081526200002691908162000089565b503480156200003457600080fd5b5060405162000d3238038062000d3283398101604081905262000057916200014c565b600080546001600160a01b039384166001600160a01b03199182161790915560048054929093169116179055620001c0565b828054620000979062000183565b90600052602060002090601f016020900481019282620000bb576000855562000106565b82601f10620000d657805160ff191683800117855562000106565b8280016001018555821562000106579182015b8281111562000106578251825591602001919060010190620000e9565b506200011492915062000118565b5090565b5b8082111562000114576000815560010162000119565b80516001600160a01b03811681146200014757600080fd5b919050565b600080604083850312156200015f578182fd5b6200016a836200012f565b91506200017a602084016200012f565b90509250929050565b6002810460018216806200019857607f821691505b60208210811415620001ba57634e487b7160e01b600052602260045260246000fd5b50919050565b610b6280620001d06000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80637da0a8771161005b5780637da0a877146100e0578063aa764d7c146100f5578063b37d24e8146100fd578063dc8958ba146101125761007d565b8063486ff0cd14610082578063572b6c05146100a05780636cebc9c2146100c0575b600080fd5b61008a610125565b604051610097919061090a565b60405180910390f35b6100b36100ae3660046106f9565b6101b3565b60405161009791906108cf565b6100d36100ce3660046107e2565b6101ca565b60405161009791906108da565b6100e86102ef565b60405161009791906108bb565b6100e86102fe565b61011061010b36600461077c565b61030d565b005b61011061012036600461071a565b6104b4565b6005805461013290610adb565b80601f016020809104026020016040519081016040528092919081815260200182805461015e90610adb565b80156101ab5780601f10610180576101008083540402835291602001916101ab565b820191906000526020600020905b81548152906001019060200180831161018e57829003601f168201915b505050505081565b6004546001600160a01b038281169116145b919050565b60008383816001600160a01b0316600260006101e46105d2565b6001600160a01b03908116825260208083019390935260409182016000908120868316825290935291205416146102365760405162461bcd60e51b815260040161022d9061097e565b60405180910390fd5b60006102448887878a610603565b6001600160a01b038088166000908152600160208181526040808420948d168452939052919020549192506102799190610a87565b6001600160a01b038088166000908152600160209081526040808320938c16835292905281902091909155517f5c3fb349179e8d8347e69078d6b9912b9724461c39bf776a56925267afc55aff906102dc90309084908c908c908c908c90610a1e565b60405180910390a1979650505050505050565b6004546001600160a01b031681565b6000546001600160a01b031681565b6001600160a01b038083166000908152600260209081526040808320858516845282528083205460018352818420941680845293825280832054905161035792889146910161088e565b604051602081830303815290604052805190602001209050600080856001600160a01b03168760405161038a9190610872565b6000604051808303816000865af19150503d80600081146103c7576040519150601f19603f3d011682016040523d82523d6000602084013e6103cc565b606091505b50915091508180156103f65750805115806103f65750808060200190518101906103f6919061075c565b6104125760405162461bcd60e51b815260040161022d906109fe565b6001600160a01b03808616600090815260016020818152604080842094891684529390529190205461044391610a87565b6001600160a01b038087166000908152600160209081526040808320938916835292905281902091909155517f87565ac85e588baf96cb01bd0fb239c4ab035e0d56eb65f9d1c3d68b80ccfd29906104a2908a908990899089906108e3565b60405180910390a15050505050505050565b6001600160a01b0381166104da5760405162461bcd60e51b815260040161022d906109d5565b6001600160a01b0383166105005760405162461bcd60e51b815260040161022d906109aa565b6001600160a01b03811660009081526003602052604090205460ff16156105395760405162461bcd60e51b815260040161022d9061091d565b6001600160a01b0383811660009081526002602090815260408083208685168452909152902054161561057e5760405162461bcd60e51b815260040161022d90610947565b6001600160a01b03928316600090815260026020908152604080832094861683529381528382208054959093166001600160a01b03199095168517909255928352600390529020805460ff19166001179055565b6000601436108015906105e957506105e9336101b3565b156105fd575060131936013560601c610600565b50335b90565b6001600160a01b03808416600090815260016020908152604080832093851683529281528282205492519192839261063f92899188910161088e565b60408051808303601f1901815291905280516020909101209695505050505050565b80356001600160a01b03811681146101c557600080fd5b600082601f830112610688578081fd5b813567ffffffffffffffff808211156106a3576106a3610b16565b604051601f8301601f1916810160200182811182821017156106c7576106c7610b16565b6040528281528483016020018610156106de578384fd5b82602086016020830137918201602001929092529392505050565b60006020828403121561070a578081fd5b61071382610661565b9392505050565b60008060006060848603121561072e578182fd5b61073784610661565b925061074560208501610661565b915061075360408501610661565b90509250925092565b60006020828403121561076d578081fd5b81518015158114610713578182fd5b60008060008060808587031215610791578081fd5b84359350602085013567ffffffffffffffff8111156107ae578182fd5b6107ba87828801610678565b9350506107c960408601610661565b91506107d760608601610661565b905092959194509250565b600080600080608085870312156107f7578384fd5b843567ffffffffffffffff81111561080d578485fd5b61081987828801610678565b94505061082860208601610661565b925061083660408601610661565b9396929550929360600135925050565b6000815180845261085e816020860160208601610aab565b601f01601f19169290920160200192915050565b60008251610884818460208701610aab565b9190910192915050565b600084825283516108a6816020850160208801610aab565b60209201918201929092526040019392505050565b6001600160a01b0391909116815260200190565b901515815260200190565b90815260200190565b9384526001600160a01b039283166020850152908216604084015216606082015260800190565b6000602082526107136020830184610846565b60208082526010908201526f1513c8105314915051164811561254d560821b604082015260600190565b60208082526017908201527f55504441544520444f4553204e4f5420414c4c4f574544000000000000000000604082015260600190565b602080825260129082015271155395149554d511510810d3d395149050d560721b604082015260600190565b6020808252601190820152704e554c4c20414444524553532046524f4d60781b604082015260600190565b6020808252600f908201526e4e554c4c204144445245535320544f60881b604082015260600190565b60208082526006908201526511905253115160d21b604082015260600190565b600060e08252600a60e0830152691cd95d14995c5d595cdd60b21b61010083015261012060018060a01b03808a166020850152886040850152816060850152610a6982850189610846565b96811660808501529490941660a08301525060c00152509392505050565b60008219821115610aa657634e487b7160e01b81526011600452602481fd5b500190565b60005b83811015610ac6578181015183820152602001610aae565b83811115610ad5576000848401525b50505050565b600281046001821680610aef57607f821691505b60208210811415610b1057634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fdfea264697066735822122062497ef7869fe8bcb6482b81520244f1b92480f2032890f34f7bafbd00a5ec7064736f6c63430008000033"

// DeployBridge deploys a new Ethereum contract, binding an instance of Bridge to it.
func DeployBridge(auth *bind.TransactOpts, backend bind.ContractBackend, listNode common.Address, forwarder common.Address) (common.Address, *types.Transaction, *Bridge, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BridgeBin), backend, listNode, forwarder)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// ListNode is a free data retrieval call binding the contract method 0xaa764d7c.
//
// Solidity: function _listNode() view returns(address)
func (_Bridge *BridgeCaller) ListNode(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_listNode")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ListNode is a free data retrieval call binding the contract method 0xaa764d7c.
//
// Solidity: function _listNode() view returns(address)
func (_Bridge *BridgeSession) ListNode() (common.Address, error) {
	return _Bridge.Contract.ListNode(&_Bridge.CallOpts)
}

// ListNode is a free data retrieval call binding the contract method 0xaa764d7c.
//
// Solidity: function _listNode() view returns(address)
func (_Bridge *BridgeCallerSession) ListNode() (common.Address, error) {
	return _Bridge.Contract.ListNode(&_Bridge.CallOpts)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Bridge *BridgeCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Bridge *BridgeSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Bridge.Contract.IsTrustedForwarder(&_Bridge.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Bridge *BridgeCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Bridge.Contract.IsTrustedForwarder(&_Bridge.CallOpts, forwarder)
}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_Bridge *BridgeCaller) TrustedForwarder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "trustedForwarder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_Bridge *BridgeSession) TrustedForwarder() (common.Address, error) {
	return _Bridge.Contract.TrustedForwarder(&_Bridge.CallOpts)
}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_Bridge *BridgeCallerSession) TrustedForwarder() (common.Address, error) {
	return _Bridge.Contract.TrustedForwarder(&_Bridge.CallOpts)
}

// VersionRecipient is a free data retrieval call binding the contract method 0x486ff0cd.
//
// Solidity: function versionRecipient() view returns(string)
func (_Bridge *BridgeCaller) VersionRecipient(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "versionRecipient")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VersionRecipient is a free data retrieval call binding the contract method 0x486ff0cd.
//
// Solidity: function versionRecipient() view returns(string)
func (_Bridge *BridgeSession) VersionRecipient() (string, error) {
	return _Bridge.Contract.VersionRecipient(&_Bridge.CallOpts)
}

// VersionRecipient is a free data retrieval call binding the contract method 0x486ff0cd.
//
// Solidity: function versionRecipient() view returns(string)
func (_Bridge *BridgeCallerSession) VersionRecipient() (string, error) {
	return _Bridge.Contract.VersionRecipient(&_Bridge.CallOpts)
}

// AddContractBind is a paid mutator transaction binding the contract method 0xdc8958ba.
//
// Solidity: function addContractBind(address from, address oppositeBridge, address to) returns()
func (_Bridge *BridgeTransactor) AddContractBind(opts *bind.TransactOpts, from common.Address, oppositeBridge common.Address, to common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "addContractBind", from, oppositeBridge, to)
}

// AddContractBind is a paid mutator transaction binding the contract method 0xdc8958ba.
//
// Solidity: function addContractBind(address from, address oppositeBridge, address to) returns()
func (_Bridge *BridgeSession) AddContractBind(from common.Address, oppositeBridge common.Address, to common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AddContractBind(&_Bridge.TransactOpts, from, oppositeBridge, to)
}

// AddContractBind is a paid mutator transaction binding the contract method 0xdc8958ba.
//
// Solidity: function addContractBind(address from, address oppositeBridge, address to) returns()
func (_Bridge *BridgeTransactorSession) AddContractBind(from common.Address, oppositeBridge common.Address, to common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AddContractBind(&_Bridge.TransactOpts, from, oppositeBridge, to)
}

// ReceiveRequestV2 is a paid mutator transaction binding the contract method 0xb37d24e8.
//
// Solidity: function receiveRequestV2(bytes32 reqId, bytes b, address receiveSide, address bridgeFrom) returns()
func (_Bridge *BridgeTransactor) ReceiveRequestV2(opts *bind.TransactOpts, reqId [32]byte, b []byte, receiveSide common.Address, bridgeFrom common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "receiveRequestV2", reqId, b, receiveSide, bridgeFrom)
}

// ReceiveRequestV2 is a paid mutator transaction binding the contract method 0xb37d24e8.
//
// Solidity: function receiveRequestV2(bytes32 reqId, bytes b, address receiveSide, address bridgeFrom) returns()
func (_Bridge *BridgeSession) ReceiveRequestV2(reqId [32]byte, b []byte, receiveSide common.Address, bridgeFrom common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.ReceiveRequestV2(&_Bridge.TransactOpts, reqId, b, receiveSide, bridgeFrom)
}

// ReceiveRequestV2 is a paid mutator transaction binding the contract method 0xb37d24e8.
//
// Solidity: function receiveRequestV2(bytes32 reqId, bytes b, address receiveSide, address bridgeFrom) returns()
func (_Bridge *BridgeTransactorSession) ReceiveRequestV2(reqId [32]byte, b []byte, receiveSide common.Address, bridgeFrom common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.ReceiveRequestV2(&_Bridge.TransactOpts, reqId, b, receiveSide, bridgeFrom)
}

// TransmitRequestV2 is a paid mutator transaction binding the contract method 0x6cebc9c2.
//
// Solidity: function transmitRequestV2(bytes _selector, address receiveSide, address oppositeBridge, uint256 chainId) returns(bytes32)
func (_Bridge *BridgeTransactor) TransmitRequestV2(opts *bind.TransactOpts, _selector []byte, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "transmitRequestV2", _selector, receiveSide, oppositeBridge, chainId)
}

// TransmitRequestV2 is a paid mutator transaction binding the contract method 0x6cebc9c2.
//
// Solidity: function transmitRequestV2(bytes _selector, address receiveSide, address oppositeBridge, uint256 chainId) returns(bytes32)
func (_Bridge *BridgeSession) TransmitRequestV2(_selector []byte, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.TransmitRequestV2(&_Bridge.TransactOpts, _selector, receiveSide, oppositeBridge, chainId)
}

// TransmitRequestV2 is a paid mutator transaction binding the contract method 0x6cebc9c2.
//
// Solidity: function transmitRequestV2(bytes _selector, address receiveSide, address oppositeBridge, uint256 chainId) returns(bytes32)
func (_Bridge *BridgeTransactorSession) TransmitRequestV2(_selector []byte, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.TransmitRequestV2(&_Bridge.TransactOpts, _selector, receiveSide, oppositeBridge, chainId)
}

// BridgeOracleRequestIterator is returned from FilterOracleRequest and is used to iterate over the raw logs and unpacked data for OracleRequest events raised by the Bridge contract.
type BridgeOracleRequestIterator struct {
	Event *BridgeOracleRequest // Event containing the contract specifics and raw log

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
func (it *BridgeOracleRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeOracleRequest)
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
		it.Event = new(BridgeOracleRequest)
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
func (it *BridgeOracleRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeOracleRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeOracleRequest represents a OracleRequest event raised by the Bridge contract.
type BridgeOracleRequest struct {
	RequestType    string
	Bridge         common.Address
	RequestId      [32]byte
	Selector       []byte
	ReceiveSide    common.Address
	OppositeBridge common.Address
	Chainid        *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOracleRequest is a free log retrieval operation binding the contract event 0x5c3fb349179e8d8347e69078d6b9912b9724461c39bf776a56925267afc55aff.
//
// Solidity: event OracleRequest(string requestType, address bridge, bytes32 requestId, bytes selector, address receiveSide, address oppositeBridge, uint256 chainid)
func (_Bridge *BridgeFilterer) FilterOracleRequest(opts *bind.FilterOpts) (*BridgeOracleRequestIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "OracleRequest")
	if err != nil {
		return nil, err
	}
	return &BridgeOracleRequestIterator{contract: _Bridge.contract, event: "OracleRequest", logs: logs, sub: sub}, nil
}

// WatchOracleRequest is a free log subscription operation binding the contract event 0x5c3fb349179e8d8347e69078d6b9912b9724461c39bf776a56925267afc55aff.
//
// Solidity: event OracleRequest(string requestType, address bridge, bytes32 requestId, bytes selector, address receiveSide, address oppositeBridge, uint256 chainid)
func (_Bridge *BridgeFilterer) WatchOracleRequest(opts *bind.WatchOpts, sink chan<- *BridgeOracleRequest) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "OracleRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeOracleRequest)
				if err := _Bridge.contract.UnpackLog(event, "OracleRequest", log); err != nil {
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
// Solidity: event OracleRequest(string requestType, address bridge, bytes32 requestId, bytes selector, address receiveSide, address oppositeBridge, uint256 chainid)
func (_Bridge *BridgeFilterer) ParseOracleRequest(log types.Log) (*BridgeOracleRequest, error) {
	event := new(BridgeOracleRequest)
	if err := _Bridge.contract.UnpackLog(event, "OracleRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeReceiveRequestIterator is returned from FilterReceiveRequest and is used to iterate over the raw logs and unpacked data for ReceiveRequest events raised by the Bridge contract.
type BridgeReceiveRequestIterator struct {
	Event *BridgeReceiveRequest // Event containing the contract specifics and raw log

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
func (it *BridgeReceiveRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeReceiveRequest)
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
		it.Event = new(BridgeReceiveRequest)
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
func (it *BridgeReceiveRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeReceiveRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeReceiveRequest represents a ReceiveRequest event raised by the Bridge contract.
type BridgeReceiveRequest struct {
	ReqId       [32]byte
	ReceiveSide common.Address
	BridgeFrom  common.Address
	SenderSide  common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReceiveRequest is a free log retrieval operation binding the contract event 0x87565ac85e588baf96cb01bd0fb239c4ab035e0d56eb65f9d1c3d68b80ccfd29.
//
// Solidity: event ReceiveRequest(bytes32 reqId, address receiveSide, address bridgeFrom, address senderSide)
func (_Bridge *BridgeFilterer) FilterReceiveRequest(opts *bind.FilterOpts) (*BridgeReceiveRequestIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ReceiveRequest")
	if err != nil {
		return nil, err
	}
	return &BridgeReceiveRequestIterator{contract: _Bridge.contract, event: "ReceiveRequest", logs: logs, sub: sub}, nil
}

// WatchReceiveRequest is a free log subscription operation binding the contract event 0x87565ac85e588baf96cb01bd0fb239c4ab035e0d56eb65f9d1c3d68b80ccfd29.
//
// Solidity: event ReceiveRequest(bytes32 reqId, address receiveSide, address bridgeFrom, address senderSide)
func (_Bridge *BridgeFilterer) WatchReceiveRequest(opts *bind.WatchOpts, sink chan<- *BridgeReceiveRequest) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ReceiveRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeReceiveRequest)
				if err := _Bridge.contract.UnpackLog(event, "ReceiveRequest", log); err != nil {
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

// ParseReceiveRequest is a log parse operation binding the contract event 0x87565ac85e588baf96cb01bd0fb239c4ab035e0d56eb65f9d1c3d68b80ccfd29.
//
// Solidity: event ReceiveRequest(bytes32 reqId, address receiveSide, address bridgeFrom, address senderSide)
func (_Bridge *BridgeFilterer) ParseReceiveRequest(log types.Log) (*BridgeReceiveRequest, error) {
	event := new(BridgeReceiveRequest)
	if err := _Bridge.contract.UnpackLog(event, "ReceiveRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
