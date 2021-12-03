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

// MerkleTestABI is the input ABI used to generate the binding from.
const MerkleTestABI = "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_auditPath\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"merkleProve\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// MerkleTestBin is the compiled bytecode used for deploying new contracts.
var MerkleTestBin = "0x608060405234801561001057600080fd5b50610bb4806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063efd3b91014610030575b600080fd5b61004361003e3660046106e6565b610059565b60405161005091906107fa565b60405180910390f35b6060600060606100698583610150565b92509050600061007882610222565b90506000602184885161008b9190610af0565b6100959190610ad0565b905060008060005b83811015610123576100af8a8861029a565b975091506100bd8a8861031b565b9750925060ff82166100da576100d38386610379565b9450610111565b8160ff16600114156100f0576100d38584610379565b60405162461bcd60e51b8152600401610108906108f6565b60405180910390fd5b8061011b81610b37565b91505061009d565b508784146101435760405162461bcd60e51b8152600401610108906108a5565b5092979650505050505050565b606060008061015f85856103f7565b86519095509091506101718286610ab8565b1115801561018757506101848185610ab8565b84105b6101a35760405162461bcd60e51b815260040161010890610a74565b6060811580156101be57604051915060208201604052610208565b6040519150601f8316801560200281840101848101888315602002848c0101015b818310156101f75780518352602092830192016101df565b5050848452601f01601f1916604052505b50806102148387610ab8565b9350935050505b9250929050565b6000600260008360405160200161023a9291906107cb565b60408051601f1981840301815290829052610254916107af565b602060405180830381855afa158015610271573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061029491906106ce565b92915050565b60008083518360016102ac9190610ab8565b111580156102c357506102c0836001610ab8565b83105b6102df5760405162461bcd60e51b815260040161010890610864565b6000604051846020870101518060001a82535060018101604052601f8103519150508084600161030f9190610ab8565b92509250509250929050565b600080835183602061032d9190610ab8565b111580156103445750610341836020610ab8565b83105b6103605760405162461bcd60e51b8152600401610108906109bb565b60006020840185015190508084602061030f9190610ab8565b60006002600160f81b84846040516020016103969392919061078d565b60408051601f19818403018152908290526103b0916107af565b602060405180830381855afa1580156103cd573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906103f091906106ce565b9392505050565b6000806000610406858561029a565b94509050600060fd60ff83161415610465576104228686610522565b955061ffff16905060fd811080159061043d575061ffff8111155b6104595760405162461bcd60e51b81526004016101089061082d565b925083915061021b9050565b8160ff1660fe14156104b65761047b86866105a0565b955063ffffffff16905061ffff8111801561049a575063ffffffff8111155b6104595760405162461bcd60e51b815260040161010890610944565b8160ff1660ff14156104fd576104cc8686610637565b955067ffffffffffffffff16905063ffffffff81116104595760405162461bcd60e51b815260040161010890610944565b5060ff811660fd81106104595760405162461bcd60e51b815260040161010890610944565b60008083518360026105349190610ab8565b1115801561054b5750610548836002610ab8565b83105b6105675760405162461bcd60e51b815260040161010890610979565b6000604051846020870101518060011a82538060001a60018301535060028101604052601e8103519150508084600261030f9190610ab8565b60008083518360046105b29190610ab8565b111580156105c957506105c6836004610ab8565b83105b6105e55760405162461bcd60e51b815260040161010890610a32565b600060405160046000600182038760208a0101515b8383101561061a5780821a838601536001830192506001820391506105fa565b50505081810160405260200390035190508061030f856004610ab8565b60008083518360086106499190610ab8565b11158015610660575061065d836008610ab8565b83105b61067c5760405162461bcd60e51b8152600401610108906109f0565b600060405160086000600182038760208a0101515b838310156106b15780821a83860153600183019250600182039150610691565b50505081810160405260200390035190508061030f856008610ab8565b6000602082840312156106df578081fd5b5051919050565b600080604083850312156106f8578081fd5b823567ffffffffffffffff8082111561070f578283fd5b818501915085601f830112610722578283fd5b813560208282111561073657610736610b68565b604051601f8301601f191681018201848111828210171561075957610759610b68565b604052828152848301820189101561076f578586fd5b82828601838301379182018101949094529694909201359450505050565b6001600160f81b03199390931683526001830191909152602182015260410190565b600082516107c1818460208701610b07565b9190910192915050565b600060ff60f81b8460f81b16825282516107ec816001850160208701610b07565b919091016001019392505050565b6000602082528251806020840152610819816040850160208701610b07565b601f01601f19169190910160400192915050565b6020808252601f908201527f4e65787455696e7431362c2076616c7565206f7574736964652072616e676500604082015260600190565b60208082526021908201527f4e65787455696e74382c204f66667365742065786365656473206d6178696d756040820152606d60f81b606082015260800190565b60208082526031908201527f6d65726b6c6550726f76652c2065787065637420726f6f74206973206e6f7420604082015270195c5d585b081858dd1d585b081c9bdbdd607a1b606082015260800190565b6020808252602e908201527f6d65726b6c6550726f76652c204e6578744279746520666f7220706f7369746960408201526d1bdb881a5b999bc819985a5b195960921b606082015260800190565b6020808252818101527f4e65787456617255696e742c2076616c7565206f7574736964652072616e6765604082015260600190565b60208082526022908201527f4e65787455696e7431362c206f66667365742065786365656473206d6178696d604082015261756d60f01b606082015260800190565b6020808252818101527f4e657874486173682c206f66667365742065786365656473206d6178696d756d604082015260600190565b60208082526022908201527f4e65787455696e7436342c206f66667365742065786365656473206d6178696d604082015261756d60f01b606082015260800190565b60208082526022908201527f4e65787455696e7433322c206f66667365742065786365656473206d6178696d604082015261756d60f01b606082015260800190565b60208082526024908201527f4e65787456617242797465732c206f66667365742065786365656473206d6178604082015263696d756d60e01b606082015260800190565b60008219821115610acb57610acb610b52565b500190565b600082610aeb57634e487b7160e01b81526012600452602481fd5b500490565b600082821015610b0257610b02610b52565b500390565b60005b83811015610b22578181015183820152602001610b0a565b83811115610b31576000848401525b50505050565b6000600019821415610b4b57610b4b610b52565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fdfea2646970667358221220755ff1e5531a450a750f5fcf05c3f5879783c4138ab5a84fdcd873cca038230d64736f6c63430008000033"

// DeployMerkleTest deploys a new Ethereum contract, binding an instance of MerkleTest to it.
func DeployMerkleTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MerkleTest, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleTestABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MerkleTestBin), backend)
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
