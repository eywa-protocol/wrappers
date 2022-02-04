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
)

// MockDexPoolMetaData contains all meta data concerning the MockDexPool contract.
var MockDexPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"RequestReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"reqId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"RequestReceivedV2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"reqId\",\"type\":\"bytes32\"}],\"name\":\"RequestSent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SOLANA_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOLANA_RENT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOLANA_SYSTEM_PROGRAM\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOLANA_TOKEN_PROGRAM\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"doubleRequestError\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_testData\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_reqId\",\"type\":\"bytes32\"}],\"name\":\"receiveRequestTest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"testData_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"secondPartPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"sendRequestTestV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"testStubPID_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"solBridgePID_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataAcc_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bridgePDASigner_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"testData_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"sendTestRequestToSolana\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"programId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"pubkey\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isSigner\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isWritable\",\"type\":\"bool\"}],\"internalType\":\"structSolanaSerialize.SolanaAccountMeta[]\",\"name\":\"accounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structSolanaSerialize.SolanaStandaloneInstruction\",\"name\":\"ix\",\"type\":\"tuple\"}],\"name\":\"serializeSolanaStandaloneInstruction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"sigHash\",\"outputs\":[{\"internalType\":\"bytes8\",\"name\":\"\",\"type\":\"bytes8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c0604052601060809081526f676c6f62616c3a7365745f76616c756560801b60a0526200002d90620000a5565b600080546001600160401b03191660c09290921c91909117815560018190556004553480156200005c57600080fd5b50604051620012a6380380620012a68339810160408190526200007f9162000102565b600280546001600160a01b0319166001600160a01b03929092169190911790556200018c565b6000600282604051620000b9919062000134565b602060405180830381855afa158015620000d7573d6000803e3d6000fd5b5050506040513d601f19601f82011682018060405250810190620000fc919062000172565b92915050565b6000602082840312156200011557600080fd5b81516001600160a01b03811681146200012d57600080fd5b9392505050565b6000825160005b818110156200015757602081860181015185830152016200013b565b8181111562000167576000828501525b509190910192915050565b6000602082840312156200018557600080fd5b5051919050565b61110a806200019c6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c8063670490181161008c578063edff2c8211610066578063edff2c82146101ca578063f54ea76b146101f1578063f9ee520e14610206578063faad85c81461021957600080fd5b806367049018146101535780639d8669851461017f578063e78cea921461019f57600080fd5b8063016cbd51146100d4578063067bd07a146100f05780630a82d745146100fb5780632ab8c8b0146101045780632b1106e31461012457806352dd22ee1461014b575b600080fd5b6100dd60015481565b6040519081526020015b60405180910390f35b6100dd631de44e3d81565b6100dd60045481565b610117610112366004610c22565b61022c565b6040516100e79190610dbe565b6100dd7f06a7d517192c5c51218cc94c3d4af17f58daee089ba1fd44e3dbd98a0000000081565b6100dd600081565b610166610161366004610dd1565b61041d565b6040516001600160c01b031990911681526020016100e7565b6100dd61018d366004610e22565b60036020526000908152604090205481565b6002546101b2906001600160a01b031681565b6040516001600160a01b0390911681526020016100e7565b6100dd7f06ddf6e1d765a193d9cbe146ceeb79ac1cb485ed5f5b37913a8cf5857eff00a981565b6102046101ff366004610e3b565b610475565b005b610204610214366004610e9a565b6107aa565b610204610227366004610ede565b610a14565b60208082015151604051606092600883811c62ff00ff1663ff00ff009490911b9390931692909217601081811c91901b17916000916102809184910160e09190911b6001600160e01b031916815260040190565b6040516020818303038152906040529050606060005b85602001515181101561037e57856020015181815181106102b9576102b9610f00565b602002602001015160000151866020015182815181106102db576102db610f00565b602002602001015160200151876020015183815181106102fd576102fd610f00565b6020026020010151604001516040516020016103369392919092835290151560f890811b6020840152901515901b602182015260220190565b6040516020818303038152906040529150828260405160200161035a929190610f16565b6040516020818303038152906040529250808061037690610f45565b915050610296565b508451604051610392918491602001610f6e565b60408051601f198184030181528282529087015151600881811c62ff00ff1663ff00ff009290911b9190911617601081811c91901b17945092506103dc9083908590602001610f90565b6040516020818303038152906040529150818560400151604051602001610404929190610f16565b60408051601f1981840301815291905295945050505050565b600060028260405161042f9190610fc2565b602060405180830381855afa15801561044c573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061046f9190610fde565b92915050565b631de44e3d81146104c15760405162461bcd60e51b81526020600482015260116024820152701a5b98dbdc9c9958dd0818da185a5b9251607a1b60448201526064015b60405180910390fd5b600254604051632d0335ab60e01b81523360048201526000916001600160a01b031690632d0335ab90602401602060405180830381865afa15801561050a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061052e9190610fde565b60025460405163379a90cb60e11b8152600481018a90526024810185905260448101889052336064820152608481018390529192506000916001600160a01b0390911690636f3521969060a401602060405180830381865afa158015610598573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105bc9190610fde565b6040805160028082526060820190925291925060009190816020015b60408051606081018252600080825260208083018290529282015282526000199092019101816105d85790505090506040518060600160405280888152602001600015158152602001600115158152508160008151811061063b5761063b610f00565b60200260200101819052506040518060600160405280878152602001600115158152602001600115158152508160018151811061067a5761067a610f00565b602090810291909101810191909152600254604080516060810182528c815280840185905260005482516001600160a01b03909416946331b433f9946106fc948401926106e59260c09190911b918d91016001600160c01b0319929092168252600882015260280190565b60405160208183030381529060405281525061022c565b8b8b631de44e3d87338a6040518863ffffffff1660e01b81526004016107289796959493929190610ff7565b6020604051808303816000875af1158015610747573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061076b9190611044565b506040518281527f1131472297a800fee664d1d89cfa8f7676ff07189ecc53f80bbb5f4969099db89060200160405180910390a1505050505050505050565b6001600160a01b0383166107ee5760405162461bcd60e51b815260206004820152600b60248201526a424144204144445245535360a81b60448201526064016104b8565b600254604051632d0335ab60e01b81523360048201526000916001600160a01b031690632d0335ab90602401602060405180830381865afa158015610837573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061085b9190610fde565b60025460405163379a90cb60e11b81526001600160a01b0386811660048301526024820186905287811660448301523360648301526084820184905292935060009290911690636f3521969060a401602060405180830381865afa1580156108c7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108eb9190610fde565b905060006040518060600160405280602381526020016110b2602391398051602090910120604051602481018990526044810184905260640160408051601f198184030181529181526020820180516001600160e01b03166001600160e01b031990941693909317909252600254915163329ef45f60e01b81529092506001600160a01b039091169063329ef45f906109949084908a908a908a90899033908c90600401611061565b6020604051808303816000875af11580156109b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109d79190611044565b506040518281527f1131472297a800fee664d1d89cfa8f7676ff07189ecc53f80bbb5f4969099db89060200160405180910390a150505050505050565b6002546001600160a01b03163314610a645760405162461bcd60e51b81526020600482015260136024820152724f4e4c59204345525441494e2042524944474560681b60448201526064016104b8565b60008181526003602052604090205415610a8e5760048054906000610a8883610f45565b91905055505b6000818152600360205260408120805491610aa883610f45565b909155505060018290556040518281527f3e656638c321b7f315f9cd0ebbfbbb108d6d8fbad72ae54b16a2c1ac2dc1de0e9060200160405180910390a160408051828152602081018490527f254f734e74519a6ad860fd5a8c9fe8e49517cac9b3344a71c6e6d81e771d1509910160405180910390a15050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff81118282101715610b5b57610b5b610b22565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715610b8a57610b8a610b22565b604052919050565b8015158114610ba057600080fd5b50565b600067ffffffffffffffff831115610bbd57610bbd610b22565b610bd0601f8401601f1916602001610b61565b9050828152838383011115610be457600080fd5b828260208301376000602084830101529392505050565b600082601f830112610c0c57600080fd5b610c1b83833560208501610ba3565b9392505050565b60006020808385031215610c3557600080fd5b823567ffffffffffffffff80821115610c4d57600080fd5b81850191506060808388031215610c6357600080fd5b610c6b610b38565b833581528484013583811115610c8057600080fd5b8401601f81018913610c9157600080fd5b803584811115610ca357610ca3610b22565b610cb1878260051b01610b61565b8181529084028201870190878101908b831115610ccd57600080fd5b928801925b82841015610d295785848d031215610cea5760008081fd5b610cf2610b38565b8435815289850135610d0381610b92565b818b0152604085810135610d1681610b92565b9082015282529285019290880190610cd2565b8089860152505050506040840135945082851115610d4657600080fd5b610d5288868601610bfb565b6040820152979650505050505050565b60005b83811015610d7d578181015183820152602001610d65565b83811115610d8c576000848401525b50505050565b60008151808452610daa816020860160208601610d62565b601f01601f19169290920160200192915050565b602081526000610c1b6020830184610d92565b600060208284031215610de357600080fd5b813567ffffffffffffffff811115610dfa57600080fd5b8201601f81018413610e0b57600080fd5b610e1a84823560208401610ba3565b949350505050565b600060208284031215610e3457600080fd5b5035919050565b60008060008060008060c08789031215610e5457600080fd5b505084359660208601359650604086013595606081013595506080810135945060a0013592509050565b80356001600160a01b0381168114610e9557600080fd5b919050565b60008060008060808587031215610eb057600080fd5b84359350610ec060208601610e7e565b9250610ece60408601610e7e565b9396929550929360600135925050565b60008060408385031215610ef157600080fd5b50508035926020909101359150565b634e487b7160e01b600052603260045260246000fd5b60008351610f28818460208801610d62565b835190830190610f3c818360208801610d62565b01949350505050565b6000600019821415610f6757634e487b7160e01b600052601160045260246000fd5b5060010190565b60008351610f80818460208801610d62565b9190910191825250602001919050565b60008351610fa2818460208801610d62565b60e09390931b6001600160e01b0319169190920190815260040192915050565b60008251610fd4818460208701610d62565b9190910192915050565b600060208284031215610ff057600080fd5b5051919050565b60e08152600061100a60e083018a610d92565b6020830198909852506040810195909552606085019390935260808401919091526001600160a01b031660a083015260c090910152919050565b60006020828403121561105657600080fd5b8151610c1b81610b92565b60e08152600061107460e083018a610d92565b6001600160a01b0398891660208401529688166040830152506060810194909452608084019290925290931660a082015260c0019190915291905056fe7265636569766552657175657374546573742875696e743235362c6279746573333229a2646970667358221220175a1b41023c044a0c488209907f316acfd202d18d233cbaa5a74a47f1c2fd9264736f6c634300080a0033",
}

// MockDexPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use MockDexPoolMetaData.ABI instead.
var MockDexPoolABI = MockDexPoolMetaData.ABI

// MockDexPoolBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockDexPoolMetaData.Bin instead.
var MockDexPoolBin = MockDexPoolMetaData.Bin

// DeployMockDexPool deploys a new Ethereum contract, binding an instance of MockDexPool to it.
func DeployMockDexPool(auth *bind.TransactOpts, backend bind.ContractBackend, _bridge common.Address) (common.Address, *types.Transaction, *MockDexPool, error) {
	parsed, err := MockDexPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockDexPoolBin), backend, _bridge)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockDexPool{MockDexPoolCaller: MockDexPoolCaller{contract: contract}, MockDexPoolTransactor: MockDexPoolTransactor{contract: contract}, MockDexPoolFilterer: MockDexPoolFilterer{contract: contract}}, nil
}

// MockDexPool is an auto generated Go binding around an Ethereum contract.
type MockDexPool struct {
	MockDexPoolCaller     // Read-only binding to the contract
	MockDexPoolTransactor // Write-only binding to the contract
	MockDexPoolFilterer   // Log filterer for contract events
}

// MockDexPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockDexPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockDexPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockDexPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
	gsn      *GsnCallOpts
}

func (_MockDexPool *MockDexPoolTransactor) SetGSNOptions(opts *GsnCallOpts) {
	_MockDexPool.gsn = opts
}

// MockDexPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockDexPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockDexPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockDexPoolSession struct {
	Contract     *MockDexPool      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockDexPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockDexPoolCallerSession struct {
	Contract *MockDexPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MockDexPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockDexPoolTransactorSession struct {
	Contract     *MockDexPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MockDexPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockDexPoolRaw struct {
	Contract *MockDexPool // Generic contract binding to access the raw methods on
}

// MockDexPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockDexPoolCallerRaw struct {
	Contract *MockDexPoolCaller // Generic read-only contract binding to access the raw methods on
}

// MockDexPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockDexPoolTransactorRaw struct {
	Contract *MockDexPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockDexPool creates a new instance of MockDexPool, bound to a specific deployed contract.
func NewMockDexPool(address common.Address, backend bind.ContractBackend) (*MockDexPool, error) {
	contract, err := bindMockDexPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockDexPool{MockDexPoolCaller: MockDexPoolCaller{contract: contract}, MockDexPoolTransactor: MockDexPoolTransactor{contract: contract}, MockDexPoolFilterer: MockDexPoolFilterer{contract: contract}}, nil
}

// NewMockDexPoolCaller creates a new read-only instance of MockDexPool, bound to a specific deployed contract.
func NewMockDexPoolCaller(address common.Address, caller bind.ContractCaller) (*MockDexPoolCaller, error) {
	contract, err := bindMockDexPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockDexPoolCaller{contract: contract}, nil
}

// NewMockDexPoolTransactor creates a new write-only instance of MockDexPool, bound to a specific deployed contract.
func NewMockDexPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*MockDexPoolTransactor, error) {
	contract, err := bindMockDexPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockDexPoolTransactor{contract: contract}, nil
}

// NewMockDexPoolFilterer creates a new log filterer instance of MockDexPool, bound to a specific deployed contract.
func NewMockDexPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*MockDexPoolFilterer, error) {
	contract, err := bindMockDexPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockDexPoolFilterer{contract: contract}, nil
}

// bindMockDexPool binds a generic wrapper to an already deployed contract.
func bindMockDexPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MockDexPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockDexPool *MockDexPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockDexPool.Contract.MockDexPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockDexPool *MockDexPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockDexPool.Contract.MockDexPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockDexPool *MockDexPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockDexPool.Contract.MockDexPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockDexPool *MockDexPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockDexPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockDexPool *MockDexPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockDexPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockDexPool *MockDexPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockDexPool.Contract.contract.Transact(opts, method, params...)
}

// SOLANACHAINID is a free data retrieval call binding the contract method 0x067bd07a.
//
// Solidity: function SOLANA_CHAIN_ID() view returns(uint256)
func (_MockDexPool *MockDexPoolCaller) SOLANACHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockDexPool.contract.Call(opts, &out, "SOLANA_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SOLANACHAINID is a free data retrieval call binding the contract method 0x067bd07a.
//
// Solidity: function SOLANA_CHAIN_ID() view returns(uint256)
func (_MockDexPool *MockDexPoolSession) SOLANACHAINID() (*big.Int, error) {
	return _MockDexPool.Contract.SOLANACHAINID(&_MockDexPool.CallOpts)
}

// SOLANACHAINID is a free data retrieval call binding the contract method 0x067bd07a.
//
// Solidity: function SOLANA_CHAIN_ID() view returns(uint256)
func (_MockDexPool *MockDexPoolCallerSession) SOLANACHAINID() (*big.Int, error) {
	return _MockDexPool.Contract.SOLANACHAINID(&_MockDexPool.CallOpts)
}

// SOLANARENT is a free data retrieval call binding the contract method 0x2b1106e3.
//
// Solidity: function SOLANA_RENT() view returns(bytes32)
func (_MockDexPool *MockDexPoolCaller) SOLANARENT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MockDexPool.contract.Call(opts, &out, "SOLANA_RENT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SOLANARENT is a free data retrieval call binding the contract method 0x2b1106e3.
//
// Solidity: function SOLANA_RENT() view returns(bytes32)
func (_MockDexPool *MockDexPoolSession) SOLANARENT() ([32]byte, error) {
	return _MockDexPool.Contract.SOLANARENT(&_MockDexPool.CallOpts)
}

// SOLANARENT is a free data retrieval call binding the contract method 0x2b1106e3.
//
// Solidity: function SOLANA_RENT() view returns(bytes32)
func (_MockDexPool *MockDexPoolCallerSession) SOLANARENT() ([32]byte, error) {
	return _MockDexPool.Contract.SOLANARENT(&_MockDexPool.CallOpts)
}

// SOLANASYSTEMPROGRAM is a free data retrieval call binding the contract method 0x52dd22ee.
//
// Solidity: function SOLANA_SYSTEM_PROGRAM() view returns(bytes32)
func (_MockDexPool *MockDexPoolCaller) SOLANASYSTEMPROGRAM(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MockDexPool.contract.Call(opts, &out, "SOLANA_SYSTEM_PROGRAM")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SOLANASYSTEMPROGRAM is a free data retrieval call binding the contract method 0x52dd22ee.
//
// Solidity: function SOLANA_SYSTEM_PROGRAM() view returns(bytes32)
func (_MockDexPool *MockDexPoolSession) SOLANASYSTEMPROGRAM() ([32]byte, error) {
	return _MockDexPool.Contract.SOLANASYSTEMPROGRAM(&_MockDexPool.CallOpts)
}

// SOLANASYSTEMPROGRAM is a free data retrieval call binding the contract method 0x52dd22ee.
//
// Solidity: function SOLANA_SYSTEM_PROGRAM() view returns(bytes32)
func (_MockDexPool *MockDexPoolCallerSession) SOLANASYSTEMPROGRAM() ([32]byte, error) {
	return _MockDexPool.Contract.SOLANASYSTEMPROGRAM(&_MockDexPool.CallOpts)
}

// SOLANATOKENPROGRAM is a free data retrieval call binding the contract method 0xedff2c82.
//
// Solidity: function SOLANA_TOKEN_PROGRAM() view returns(bytes32)
func (_MockDexPool *MockDexPoolCaller) SOLANATOKENPROGRAM(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MockDexPool.contract.Call(opts, &out, "SOLANA_TOKEN_PROGRAM")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SOLANATOKENPROGRAM is a free data retrieval call binding the contract method 0xedff2c82.
//
// Solidity: function SOLANA_TOKEN_PROGRAM() view returns(bytes32)
func (_MockDexPool *MockDexPoolSession) SOLANATOKENPROGRAM() ([32]byte, error) {
	return _MockDexPool.Contract.SOLANATOKENPROGRAM(&_MockDexPool.CallOpts)
}

// SOLANATOKENPROGRAM is a free data retrieval call binding the contract method 0xedff2c82.
//
// Solidity: function SOLANA_TOKEN_PROGRAM() view returns(bytes32)
func (_MockDexPool *MockDexPoolCallerSession) SOLANATOKENPROGRAM() ([32]byte, error) {
	return _MockDexPool.Contract.SOLANATOKENPROGRAM(&_MockDexPool.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_MockDexPool *MockDexPoolCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockDexPool.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_MockDexPool *MockDexPoolSession) Bridge() (common.Address, error) {
	return _MockDexPool.Contract.Bridge(&_MockDexPool.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_MockDexPool *MockDexPoolCallerSession) Bridge() (common.Address, error) {
	return _MockDexPool.Contract.Bridge(&_MockDexPool.CallOpts)
}

// DoubleRequestError is a free data retrieval call binding the contract method 0x0a82d745.
//
// Solidity: function doubleRequestError() view returns(uint256)
func (_MockDexPool *MockDexPoolCaller) DoubleRequestError(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockDexPool.contract.Call(opts, &out, "doubleRequestError")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DoubleRequestError is a free data retrieval call binding the contract method 0x0a82d745.
//
// Solidity: function doubleRequestError() view returns(uint256)
func (_MockDexPool *MockDexPoolSession) DoubleRequestError() (*big.Int, error) {
	return _MockDexPool.Contract.DoubleRequestError(&_MockDexPool.CallOpts)
}

// DoubleRequestError is a free data retrieval call binding the contract method 0x0a82d745.
//
// Solidity: function doubleRequestError() view returns(uint256)
func (_MockDexPool *MockDexPoolCallerSession) DoubleRequestError() (*big.Int, error) {
	return _MockDexPool.Contract.DoubleRequestError(&_MockDexPool.CallOpts)
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) view returns(uint256)
func (_MockDexPool *MockDexPoolCaller) Requests(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _MockDexPool.contract.Call(opts, &out, "requests", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) view returns(uint256)
func (_MockDexPool *MockDexPoolSession) Requests(arg0 [32]byte) (*big.Int, error) {
	return _MockDexPool.Contract.Requests(&_MockDexPool.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) view returns(uint256)
func (_MockDexPool *MockDexPoolCallerSession) Requests(arg0 [32]byte) (*big.Int, error) {
	return _MockDexPool.Contract.Requests(&_MockDexPool.CallOpts, arg0)
}

// SerializeSolanaStandaloneInstruction is a free data retrieval call binding the contract method 0x2ab8c8b0.
//
// Solidity: function serializeSolanaStandaloneInstruction((bytes32,(bytes32,bool,bool)[],bytes) ix) pure returns(bytes)
func (_MockDexPool *MockDexPoolCaller) SerializeSolanaStandaloneInstruction(opts *bind.CallOpts, ix SolanaSerializeSolanaStandaloneInstruction) ([]byte, error) {
	var out []interface{}
	err := _MockDexPool.contract.Call(opts, &out, "serializeSolanaStandaloneInstruction", ix)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// SerializeSolanaStandaloneInstruction is a free data retrieval call binding the contract method 0x2ab8c8b0.
//
// Solidity: function serializeSolanaStandaloneInstruction((bytes32,(bytes32,bool,bool)[],bytes) ix) pure returns(bytes)
func (_MockDexPool *MockDexPoolSession) SerializeSolanaStandaloneInstruction(ix SolanaSerializeSolanaStandaloneInstruction) ([]byte, error) {
	return _MockDexPool.Contract.SerializeSolanaStandaloneInstruction(&_MockDexPool.CallOpts, ix)
}

// SerializeSolanaStandaloneInstruction is a free data retrieval call binding the contract method 0x2ab8c8b0.
//
// Solidity: function serializeSolanaStandaloneInstruction((bytes32,(bytes32,bool,bool)[],bytes) ix) pure returns(bytes)
func (_MockDexPool *MockDexPoolCallerSession) SerializeSolanaStandaloneInstruction(ix SolanaSerializeSolanaStandaloneInstruction) ([]byte, error) {
	return _MockDexPool.Contract.SerializeSolanaStandaloneInstruction(&_MockDexPool.CallOpts, ix)
}

// SigHash is a free data retrieval call binding the contract method 0x67049018.
//
// Solidity: function sigHash(string _data) pure returns(bytes8)
func (_MockDexPool *MockDexPoolCaller) SigHash(opts *bind.CallOpts, _data string) ([8]byte, error) {
	var out []interface{}
	err := _MockDexPool.contract.Call(opts, &out, "sigHash", _data)

	if err != nil {
		return *new([8]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([8]byte)).(*[8]byte)

	return out0, err

}

// SigHash is a free data retrieval call binding the contract method 0x67049018.
//
// Solidity: function sigHash(string _data) pure returns(bytes8)
func (_MockDexPool *MockDexPoolSession) SigHash(_data string) ([8]byte, error) {
	return _MockDexPool.Contract.SigHash(&_MockDexPool.CallOpts, _data)
}

// SigHash is a free data retrieval call binding the contract method 0x67049018.
//
// Solidity: function sigHash(string _data) pure returns(bytes8)
func (_MockDexPool *MockDexPoolCallerSession) SigHash(_data string) ([8]byte, error) {
	return _MockDexPool.Contract.SigHash(&_MockDexPool.CallOpts, _data)
}

// TestData is a free data retrieval call binding the contract method 0x016cbd51.
//
// Solidity: function testData() view returns(uint256)
func (_MockDexPool *MockDexPoolCaller) TestData(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockDexPool.contract.Call(opts, &out, "testData")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TestData is a free data retrieval call binding the contract method 0x016cbd51.
//
// Solidity: function testData() view returns(uint256)
func (_MockDexPool *MockDexPoolSession) TestData() (*big.Int, error) {
	return _MockDexPool.Contract.TestData(&_MockDexPool.CallOpts)
}

// TestData is a free data retrieval call binding the contract method 0x016cbd51.
//
// Solidity: function testData() view returns(uint256)
func (_MockDexPool *MockDexPoolCallerSession) TestData() (*big.Int, error) {
	return _MockDexPool.Contract.TestData(&_MockDexPool.CallOpts)
}

// ReceiveRequestTest is a paid mutator transaction binding the contract method 0xfaad85c8.
//
// Solidity: function receiveRequestTest(uint256 _testData, bytes32 _reqId) returns()
func (_MockDexPool *MockDexPoolTransactor) ReceiveRequestTest(opts *bind.TransactOpts, _testData *big.Int, _reqId [32]byte) (common.Hash, error) {
	fmt.Printf("DBG: Wrapper run method = %s\n", "ReceiveRequestTest")
	return GsnWrap(
		_MockDexPool.gsn,
		func() (common.Hash, error) {
			tx, errIn := _MockDexPool.contract.Transact(opts, "receiveRequestTest", _testData, _reqId)
			if tx != nil {
				return tx.Hash(), errIn
			}
			return common.Hash{}, errIn
		},
		func() (common.Hash, error) {
			return GsnExecutor(_MockDexPool.gsn, MockDexPoolMetaData.ABI, "receiveRequestTest", _testData, _reqId)
		},
	)
}

// ReceiveRequestTest is a paid mutator transaction binding the contract method 0xfaad85c8.
//
// Solidity: function receiveRequestTest(uint256 _testData, bytes32 _reqId) returns()
func (_MockDexPool *MockDexPoolSession) ReceiveRequestTest(_testData *big.Int, _reqId [32]byte) (common.Hash, error) {
	return _MockDexPool.Contract.ReceiveRequestTest(&_MockDexPool.TransactOpts, _testData, _reqId)
}

// ReceiveRequestTest is a paid mutator transaction binding the contract method 0xfaad85c8.
//
// Solidity: function receiveRequestTest(uint256 _testData, bytes32 _reqId) returns()
func (_MockDexPool *MockDexPoolTransactorSession) ReceiveRequestTest(_testData *big.Int, _reqId [32]byte) (common.Hash, error) {
	return _MockDexPool.Contract.ReceiveRequestTest(&_MockDexPool.TransactOpts, _testData, _reqId)
}

// SendRequestTestV2 is a paid mutator transaction binding the contract method 0xf9ee520e.
//
// Solidity: function sendRequestTestV2(uint256 testData_, address secondPartPool, address oppBridge, uint256 chainId) returns()
func (_MockDexPool *MockDexPoolTransactor) SendRequestTestV2(opts *bind.TransactOpts, testData_ *big.Int, secondPartPool common.Address, oppBridge common.Address, chainId *big.Int) (common.Hash, error) {
	fmt.Printf("DBG: Wrapper run method = %s\n", "SendRequestTestV2")
	return GsnWrap(
		_MockDexPool.gsn,
		func() (common.Hash, error) {
			tx, errIn := _MockDexPool.contract.Transact(opts, "sendRequestTestV2", testData_, secondPartPool, oppBridge, chainId)
			if tx != nil {
				return tx.Hash(), errIn
			}
			return common.Hash{}, errIn
		},
		func() (common.Hash, error) {
			return GsnExecutor(_MockDexPool.gsn, MockDexPoolMetaData.ABI, "sendRequestTestV2", testData_, secondPartPool, oppBridge, chainId)
		},
	)
}

// SendRequestTestV2 is a paid mutator transaction binding the contract method 0xf9ee520e.
//
// Solidity: function sendRequestTestV2(uint256 testData_, address secondPartPool, address oppBridge, uint256 chainId) returns()
func (_MockDexPool *MockDexPoolSession) SendRequestTestV2(testData_ *big.Int, secondPartPool common.Address, oppBridge common.Address, chainId *big.Int) (common.Hash, error) {
	return _MockDexPool.Contract.SendRequestTestV2(&_MockDexPool.TransactOpts, testData_, secondPartPool, oppBridge, chainId)
}

// SendRequestTestV2 is a paid mutator transaction binding the contract method 0xf9ee520e.
//
// Solidity: function sendRequestTestV2(uint256 testData_, address secondPartPool, address oppBridge, uint256 chainId) returns()
func (_MockDexPool *MockDexPoolTransactorSession) SendRequestTestV2(testData_ *big.Int, secondPartPool common.Address, oppBridge common.Address, chainId *big.Int) (common.Hash, error) {
	return _MockDexPool.Contract.SendRequestTestV2(&_MockDexPool.TransactOpts, testData_, secondPartPool, oppBridge, chainId)
}

// SendTestRequestToSolana is a paid mutator transaction binding the contract method 0xf54ea76b.
//
// Solidity: function sendTestRequestToSolana(bytes32 testStubPID_, bytes32 solBridgePID_, bytes32 dataAcc_, bytes32 bridgePDASigner_, uint256 testData_, uint256 chainId) returns()
func (_MockDexPool *MockDexPoolTransactor) SendTestRequestToSolana(opts *bind.TransactOpts, testStubPID_ [32]byte, solBridgePID_ [32]byte, dataAcc_ [32]byte, bridgePDASigner_ [32]byte, testData_ *big.Int, chainId *big.Int) (common.Hash, error) {
	fmt.Printf("DBG: Wrapper run method = %s\n", "SendTestRequestToSolana")
	return GsnWrap(
		_MockDexPool.gsn,
		func() (common.Hash, error) {
			tx, errIn := _MockDexPool.contract.Transact(opts, "sendTestRequestToSolana", testStubPID_, solBridgePID_, dataAcc_, bridgePDASigner_, testData_, chainId)
			if tx != nil {
				return tx.Hash(), errIn
			}
			return common.Hash{}, errIn
		},
		func() (common.Hash, error) {
			return GsnExecutor(_MockDexPool.gsn, MockDexPoolMetaData.ABI, "sendTestRequestToSolana", testStubPID_, solBridgePID_, dataAcc_, bridgePDASigner_, testData_, chainId)
		},
	)
}

// SendTestRequestToSolana is a paid mutator transaction binding the contract method 0xf54ea76b.
//
// Solidity: function sendTestRequestToSolana(bytes32 testStubPID_, bytes32 solBridgePID_, bytes32 dataAcc_, bytes32 bridgePDASigner_, uint256 testData_, uint256 chainId) returns()
func (_MockDexPool *MockDexPoolSession) SendTestRequestToSolana(testStubPID_ [32]byte, solBridgePID_ [32]byte, dataAcc_ [32]byte, bridgePDASigner_ [32]byte, testData_ *big.Int, chainId *big.Int) (common.Hash, error) {
	return _MockDexPool.Contract.SendTestRequestToSolana(&_MockDexPool.TransactOpts, testStubPID_, solBridgePID_, dataAcc_, bridgePDASigner_, testData_, chainId)
}

// SendTestRequestToSolana is a paid mutator transaction binding the contract method 0xf54ea76b.
//
// Solidity: function sendTestRequestToSolana(bytes32 testStubPID_, bytes32 solBridgePID_, bytes32 dataAcc_, bytes32 bridgePDASigner_, uint256 testData_, uint256 chainId) returns()
func (_MockDexPool *MockDexPoolTransactorSession) SendTestRequestToSolana(testStubPID_ [32]byte, solBridgePID_ [32]byte, dataAcc_ [32]byte, bridgePDASigner_ [32]byte, testData_ *big.Int, chainId *big.Int) (common.Hash, error) {
	return _MockDexPool.Contract.SendTestRequestToSolana(&_MockDexPool.TransactOpts, testStubPID_, solBridgePID_, dataAcc_, bridgePDASigner_, testData_, chainId)
}

// MockDexPoolRequestReceivedIterator is returned from FilterRequestReceived and is used to iterate over the raw logs and unpacked data for RequestReceived events raised by the MockDexPool contract.
type MockDexPoolRequestReceivedIterator struct {
	Event *MockDexPoolRequestReceived // Event containing the contract specifics and raw log

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
func (it *MockDexPoolRequestReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockDexPoolRequestReceived)
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
		it.Event = new(MockDexPoolRequestReceived)
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
func (it *MockDexPoolRequestReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockDexPoolRequestReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockDexPoolRequestReceived represents a RequestReceived event raised by the MockDexPool contract.
type MockDexPoolRequestReceived struct {
	Data *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRequestReceived is a free log retrieval operation binding the contract event 0x3e656638c321b7f315f9cd0ebbfbbb108d6d8fbad72ae54b16a2c1ac2dc1de0e.
//
// Solidity: event RequestReceived(uint256 data)
func (_MockDexPool *MockDexPoolFilterer) FilterRequestReceived(opts *bind.FilterOpts) (*MockDexPoolRequestReceivedIterator, error) {

	logs, sub, err := _MockDexPool.contract.FilterLogs(opts, "RequestReceived")
	if err != nil {
		return nil, err
	}
	return &MockDexPoolRequestReceivedIterator{contract: _MockDexPool.contract, event: "RequestReceived", logs: logs, sub: sub}, nil
}

// WatchRequestReceived is a free log subscription operation binding the contract event 0x3e656638c321b7f315f9cd0ebbfbbb108d6d8fbad72ae54b16a2c1ac2dc1de0e.
//
// Solidity: event RequestReceived(uint256 data)
func (_MockDexPool *MockDexPoolFilterer) WatchRequestReceived(opts *bind.WatchOpts, sink chan<- *MockDexPoolRequestReceived) (event.Subscription, error) {

	logs, sub, err := _MockDexPool.contract.WatchLogs(opts, "RequestReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockDexPoolRequestReceived)
				if err := _MockDexPool.contract.UnpackLog(event, "RequestReceived", log); err != nil {
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

// ParseRequestReceived is a log parse operation binding the contract event 0x3e656638c321b7f315f9cd0ebbfbbb108d6d8fbad72ae54b16a2c1ac2dc1de0e.
//
// Solidity: event RequestReceived(uint256 data)
func (_MockDexPool *MockDexPoolFilterer) ParseRequestReceived(log types.Log) (*MockDexPoolRequestReceived, error) {
	event := new(MockDexPoolRequestReceived)
	if err := _MockDexPool.contract.UnpackLog(event, "RequestReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockDexPoolRequestReceivedV2Iterator is returned from FilterRequestReceivedV2 and is used to iterate over the raw logs and unpacked data for RequestReceivedV2 events raised by the MockDexPool contract.
type MockDexPoolRequestReceivedV2Iterator struct {
	Event *MockDexPoolRequestReceivedV2 // Event containing the contract specifics and raw log

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
func (it *MockDexPoolRequestReceivedV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockDexPoolRequestReceivedV2)
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
		it.Event = new(MockDexPoolRequestReceivedV2)
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
func (it *MockDexPoolRequestReceivedV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockDexPoolRequestReceivedV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockDexPoolRequestReceivedV2 represents a RequestReceivedV2 event raised by the MockDexPool contract.
type MockDexPoolRequestReceivedV2 struct {
	ReqId [32]byte
	Data  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRequestReceivedV2 is a free log retrieval operation binding the contract event 0x254f734e74519a6ad860fd5a8c9fe8e49517cac9b3344a71c6e6d81e771d1509.
//
// Solidity: event RequestReceivedV2(bytes32 reqId, uint256 data)
func (_MockDexPool *MockDexPoolFilterer) FilterRequestReceivedV2(opts *bind.FilterOpts) (*MockDexPoolRequestReceivedV2Iterator, error) {

	logs, sub, err := _MockDexPool.contract.FilterLogs(opts, "RequestReceivedV2")
	if err != nil {
		return nil, err
	}
	return &MockDexPoolRequestReceivedV2Iterator{contract: _MockDexPool.contract, event: "RequestReceivedV2", logs: logs, sub: sub}, nil
}

// WatchRequestReceivedV2 is a free log subscription operation binding the contract event 0x254f734e74519a6ad860fd5a8c9fe8e49517cac9b3344a71c6e6d81e771d1509.
//
// Solidity: event RequestReceivedV2(bytes32 reqId, uint256 data)
func (_MockDexPool *MockDexPoolFilterer) WatchRequestReceivedV2(opts *bind.WatchOpts, sink chan<- *MockDexPoolRequestReceivedV2) (event.Subscription, error) {

	logs, sub, err := _MockDexPool.contract.WatchLogs(opts, "RequestReceivedV2")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockDexPoolRequestReceivedV2)
				if err := _MockDexPool.contract.UnpackLog(event, "RequestReceivedV2", log); err != nil {
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

// ParseRequestReceivedV2 is a log parse operation binding the contract event 0x254f734e74519a6ad860fd5a8c9fe8e49517cac9b3344a71c6e6d81e771d1509.
//
// Solidity: event RequestReceivedV2(bytes32 reqId, uint256 data)
func (_MockDexPool *MockDexPoolFilterer) ParseRequestReceivedV2(log types.Log) (*MockDexPoolRequestReceivedV2, error) {
	event := new(MockDexPoolRequestReceivedV2)
	if err := _MockDexPool.contract.UnpackLog(event, "RequestReceivedV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockDexPoolRequestSentIterator is returned from FilterRequestSent and is used to iterate over the raw logs and unpacked data for RequestSent events raised by the MockDexPool contract.
type MockDexPoolRequestSentIterator struct {
	Event *MockDexPoolRequestSent // Event containing the contract specifics and raw log

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
func (it *MockDexPoolRequestSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockDexPoolRequestSent)
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
		it.Event = new(MockDexPoolRequestSent)
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
func (it *MockDexPoolRequestSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockDexPoolRequestSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockDexPoolRequestSent represents a RequestSent event raised by the MockDexPool contract.
type MockDexPoolRequestSent struct {
	ReqId [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRequestSent is a free log retrieval operation binding the contract event 0x1131472297a800fee664d1d89cfa8f7676ff07189ecc53f80bbb5f4969099db8.
//
// Solidity: event RequestSent(bytes32 reqId)
func (_MockDexPool *MockDexPoolFilterer) FilterRequestSent(opts *bind.FilterOpts) (*MockDexPoolRequestSentIterator, error) {

	logs, sub, err := _MockDexPool.contract.FilterLogs(opts, "RequestSent")
	if err != nil {
		return nil, err
	}
	return &MockDexPoolRequestSentIterator{contract: _MockDexPool.contract, event: "RequestSent", logs: logs, sub: sub}, nil
}

// WatchRequestSent is a free log subscription operation binding the contract event 0x1131472297a800fee664d1d89cfa8f7676ff07189ecc53f80bbb5f4969099db8.
//
// Solidity: event RequestSent(bytes32 reqId)
func (_MockDexPool *MockDexPoolFilterer) WatchRequestSent(opts *bind.WatchOpts, sink chan<- *MockDexPoolRequestSent) (event.Subscription, error) {

	logs, sub, err := _MockDexPool.contract.WatchLogs(opts, "RequestSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockDexPoolRequestSent)
				if err := _MockDexPool.contract.UnpackLog(event, "RequestSent", log); err != nil {
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

// ParseRequestSent is a log parse operation binding the contract event 0x1131472297a800fee664d1d89cfa8f7676ff07189ecc53f80bbb5f4969099db8.
//
// Solidity: event RequestSent(bytes32 reqId)
func (_MockDexPool *MockDexPoolFilterer) ParseRequestSent(log types.Log) (*MockDexPoolRequestSent, error) {
	event := new(MockDexPoolRequestSent)
	if err := _MockDexPool.contract.UnpackLog(event, "RequestSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
