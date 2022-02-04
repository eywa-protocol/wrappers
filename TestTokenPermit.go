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

// TestTokenPermitMetaData contains all meta data concerning the TestTokenPermit contract.
var TestTokenPermitMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnWithAllowanceDecrease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintWithAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101406040527f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9610120523480156200003757600080fd5b5060405162001576380380620015768339810160408190526200005a91620002cb565b604051806040016040528060048152602001634559574160e01b81525080604051806040016040528060018152602001603160f81b81525084848160039080519060200190620000ac92919062000158565b508051620000c290600490602084019062000158565b5050825160209384012082519284019290922060c083815260e08290524660a0818152604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f818a0181905281830198909852606081019590955260808086019390935230858301528051808603909201825293909201909252805194019390932090925261010052506200037292505050565b828054620001669062000335565b90600052602060002090601f0160209004810192826200018a5760008555620001d5565b82601f10620001a557805160ff1916838001178555620001d5565b82800160010185558215620001d5579182015b82811115620001d5578251825591602001919060010190620001b8565b50620001e3929150620001e7565b5090565b5b80821115620001e35760008155600101620001e8565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200022657600080fd5b81516001600160401b0380821115620002435762000243620001fe565b604051601f8301601f19908116603f011681019082821181831017156200026e576200026e620001fe565b816040528381526020925086838588010111156200028b57600080fd5b600091505b83821015620002af578582018301518183018401529082019062000290565b83821115620002c15760008385830101525b9695505050505050565b60008060408385031215620002df57600080fd5b82516001600160401b0380821115620002f757600080fd5b620003058683870162000214565b935060208501519150808211156200031c57600080fd5b506200032b8582860162000214565b9150509250929050565b600181811c908216806200034a57607f821691505b602082108114156200036c57634e487b7160e01b600052602260045260246000fd5b50919050565b60805160a05160c05160e05161010051610120516111b4620003c260003960006105a801526000610a3601526000610a8501526000610a60015260006109e401526000610a0d01526111b46000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c80637ecebe00116100a2578063a457c2d711610071578063a457c2d71461022a578063a9059cbb1461023d578063a918adf514610250578063d505accf14610263578063dd62ed3e1461027657600080fd5b80637ecebe00146101e957806395d89b41146101fc5780639be4e7b2146102045780639dc29fac1461021757600080fd5b8063313ce567116100e9578063313ce567146101815780633644e51514610190578063395093511461019857806340c10f19146101ab57806370a08231146101c057600080fd5b806306fdde031461011b578063095ea7b31461013957806318160ddd1461015c57806323b872dd1461016e575b600080fd5b610123610289565b6040516101309190610f20565b60405180910390f35b61014c610147366004610f91565b61031b565b6040519015158152602001610130565b6002545b604051908152602001610130565b61014c61017c366004610fbb565b610331565b60405160128152602001610130565b6101606103e7565b61014c6101a6366004610f91565b6103f6565b6101be6101b9366004610f91565b61042d565b005b6101606101ce366004610ff7565b6001600160a01b031660009081526020819052604090205490565b6101606101f7366004610ff7565b61043b565b61012361045b565b6101be610212366004610fbb565b61046a565b6101be610225366004610f91565b610493565b61014c610238366004610f91565b61049d565b61014c61024b366004610f91565b6104fa565b6101be61025e366004610fbb565b610507565b6101be610271366004611019565b610554565b61016061028436600461108c565b6106b8565b606060038054610298906110bf565b80601f01602080910402602001604051908101604052809291908181526020018280546102c4906110bf565b80156103115780601f106102e657610100808354040283529160200191610311565b820191906000526020600020905b8154815290600101906020018083116102f457829003601f168201915b5050505050905090565b60006103283384846106e3565b50600192915050565b600061033e848484610808565b6001600160a01b0384166000908152600160209081526040808320338452909152902054828110156103c85760405162461bcd60e51b815260206004820152602860248201527f45524332303a207472616e7366657220616d6f756e74206578636565647320616044820152676c6c6f77616e636560c01b60648201526084015b60405180910390fd5b6103dc85336103d7868561110a565b6106e3565b506001949350505050565b60006103f16109e0565b905090565b3360008181526001602090815260408083206001600160a01b038716845290915281205490916103289185906103d7908690611121565b6104378282610ad3565b5050565b6001600160a01b0381166000908152600560205260408120545b92915050565b606060048054610298906110bf565b6104748382610ad3565b61048e83838361048487876106b8565b6103d79190611121565b505050565b6104378282610bb2565b3360009081526001602090815260408083206001600160a01b0386168452909152812054828110156104e15760405162461bcd60e51b81526004016103bf90611139565b6104f033856103d7868561110a565b5060019392505050565b6000610328338484610808565b600061051384846106b8565b9050818110156105355760405162461bcd60e51b81526004016103bf90611139565b61054484846103d7858561110a565b61054e8483610bb2565b50505050565b834211156105a45760405162461bcd60e51b815260206004820152601d60248201527f45524332305065726d69743a206578706972656420646561646c696e6500000060448201526064016103bf565b60007f00000000000000000000000000000000000000000000000000000000000000008888886105d38c610d01565b6040805160208101969096526001600160a01b0394851690860152929091166060840152608083015260a082015260c0810186905260e001604051602081830303815290604052805190602001209050600061062e82610d29565b9050600061063e82878787610d77565b9050896001600160a01b0316816001600160a01b0316146106a15760405162461bcd60e51b815260206004820152601e60248201527f45524332305065726d69743a20696e76616c6964207369676e6174757265000060448201526064016103bf565b6106ac8a8a8a6106e3565b50505050505050505050565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6001600160a01b0383166107455760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b60648201526084016103bf565b6001600160a01b0382166107a65760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b60648201526084016103bf565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b6001600160a01b03831661086c5760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084016103bf565b6001600160a01b0382166108ce5760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b60648201526084016103bf565b6001600160a01b038316600090815260208190526040902054818110156109465760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b60648201526084016103bf565b610950828261110a565b6001600160a01b038086166000908152602081905260408082209390935590851681529081208054849290610986908490611121565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516109d291815260200190565b60405180910390a350505050565b60007f0000000000000000000000000000000000000000000000000000000000000000461415610a2f57507f000000000000000000000000000000000000000000000000000000000000000090565b50604080517f00000000000000000000000000000000000000000000000000000000000000006020808301919091527f0000000000000000000000000000000000000000000000000000000000000000828401527f000000000000000000000000000000000000000000000000000000000000000060608301524660808301523060a0808401919091528351808403909101815260c0909201909252805191012090565b6001600160a01b038216610b295760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016103bf565b8060026000828254610b3b9190611121565b90915550506001600160a01b03821660009081526020819052604081208054839290610b68908490611121565b90915550506040518181526001600160a01b038316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b6001600160a01b038216610c125760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b60648201526084016103bf565b6001600160a01b03821660009081526020819052604090205481811015610c865760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b60648201526084016103bf565b610c90828261110a565b6001600160a01b03841660009081526020819052604081209190915560028054849290610cbe90849061110a565b90915550506040518281526000906001600160a01b038516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020016107fb565b6001600160a01b03811660009081526005602052604090208054600181018255905b50919050565b6000610455610d366109e0565b8360405161190160f01b6020820152602281018390526042810182905260009060620160405160208183030381529060405280519060200120905092915050565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0821115610df45760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b60648201526084016103bf565b8360ff16601b1480610e0957508360ff16601c145b610e605760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b60648201526084016103bf565b6040805160008082526020820180845288905260ff871692820192909252606081018590526080810184905260019060a0016020604051602081039080840390855afa158015610eb4573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116610f175760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016103bf565b95945050505050565b600060208083528351808285015260005b81811015610f4d57858101830151858201604001528201610f31565b81811115610f5f576000604083870101525b50601f01601f1916929092016040019392505050565b80356001600160a01b0381168114610f8c57600080fd5b919050565b60008060408385031215610fa457600080fd5b610fad83610f75565b946020939093013593505050565b600080600060608486031215610fd057600080fd5b610fd984610f75565b9250610fe760208501610f75565b9150604084013590509250925092565b60006020828403121561100957600080fd5b61101282610f75565b9392505050565b600080600080600080600060e0888a03121561103457600080fd5b61103d88610f75565b965061104b60208901610f75565b95506040880135945060608801359350608088013560ff8116811461106f57600080fd5b9699959850939692959460a0840135945060c09093013592915050565b6000806040838503121561109f57600080fd5b6110a883610f75565b91506110b660208401610f75565b90509250929050565b600181811c908216806110d357607f821691505b60208210811415610d2357634e487b7160e01b600052602260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60008282101561111c5761111c6110f4565b500390565b60008219821115611134576111346110f4565b500190565b60208082526025908201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604082015264207a65726f60d81b60608201526080019056fea26469706673582212207d497252c46e285688f56bf43255ef58289a661ce8a227c4f5e51e546edf119d64736f6c634300080a0033",
}

// TestTokenPermitABI is the input ABI used to generate the binding from.
// Deprecated: Use TestTokenPermitMetaData.ABI instead.
var TestTokenPermitABI = TestTokenPermitMetaData.ABI

// TestTokenPermitBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestTokenPermitMetaData.Bin instead.
var TestTokenPermitBin = TestTokenPermitMetaData.Bin

// DeployTestTokenPermit deploys a new Ethereum contract, binding an instance of TestTokenPermit to it.
func DeployTestTokenPermit(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string) (common.Address, *types.Transaction, *TestTokenPermit, error) {
	parsed, err := TestTokenPermitMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestTokenPermitBin), backend, name_, symbol_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestTokenPermit{TestTokenPermitCaller: TestTokenPermitCaller{contract: contract}, TestTokenPermitTransactor: TestTokenPermitTransactor{contract: contract}, TestTokenPermitFilterer: TestTokenPermitFilterer{contract: contract}}, nil
}

// TestTokenPermit is an auto generated Go binding around an Ethereum contract.
type TestTokenPermit struct {
	TestTokenPermitCaller     // Read-only binding to the contract
	TestTokenPermitTransactor // Write-only binding to the contract
	TestTokenPermitFilterer   // Log filterer for contract events
}

// TestTokenPermitCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestTokenPermitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestTokenPermitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestTokenPermitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
	gsn      *GsnCallOpts
}

func (_TestTokenPermit *TestTokenPermitTransactor) SetGSNOptions(opts *GsnCallOpts) {
	_TestTokenPermit.gsn = opts
}

// TestTokenPermitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestTokenPermitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestTokenPermitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestTokenPermitSession struct {
	Contract     *TestTokenPermit  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestTokenPermitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestTokenPermitCallerSession struct {
	Contract *TestTokenPermitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// TestTokenPermitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestTokenPermitTransactorSession struct {
	Contract     *TestTokenPermitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// TestTokenPermitRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestTokenPermitRaw struct {
	Contract *TestTokenPermit // Generic contract binding to access the raw methods on
}

// TestTokenPermitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestTokenPermitCallerRaw struct {
	Contract *TestTokenPermitCaller // Generic read-only contract binding to access the raw methods on
}

// TestTokenPermitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestTokenPermitTransactorRaw struct {
	Contract *TestTokenPermitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestTokenPermit creates a new instance of TestTokenPermit, bound to a specific deployed contract.
func NewTestTokenPermit(address common.Address, backend bind.ContractBackend) (*TestTokenPermit, error) {
	contract, err := bindTestTokenPermit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestTokenPermit{TestTokenPermitCaller: TestTokenPermitCaller{contract: contract}, TestTokenPermitTransactor: TestTokenPermitTransactor{contract: contract}, TestTokenPermitFilterer: TestTokenPermitFilterer{contract: contract}}, nil
}

// NewTestTokenPermitCaller creates a new read-only instance of TestTokenPermit, bound to a specific deployed contract.
func NewTestTokenPermitCaller(address common.Address, caller bind.ContractCaller) (*TestTokenPermitCaller, error) {
	contract, err := bindTestTokenPermit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestTokenPermitCaller{contract: contract}, nil
}

// NewTestTokenPermitTransactor creates a new write-only instance of TestTokenPermit, bound to a specific deployed contract.
func NewTestTokenPermitTransactor(address common.Address, transactor bind.ContractTransactor) (*TestTokenPermitTransactor, error) {
	contract, err := bindTestTokenPermit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestTokenPermitTransactor{contract: contract}, nil
}

// NewTestTokenPermitFilterer creates a new log filterer instance of TestTokenPermit, bound to a specific deployed contract.
func NewTestTokenPermitFilterer(address common.Address, filterer bind.ContractFilterer) (*TestTokenPermitFilterer, error) {
	contract, err := bindTestTokenPermit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestTokenPermitFilterer{contract: contract}, nil
}

// bindTestTokenPermit binds a generic wrapper to an already deployed contract.
func bindTestTokenPermit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestTokenPermitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestTokenPermit *TestTokenPermitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestTokenPermit.Contract.TestTokenPermitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestTokenPermit *TestTokenPermitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.TestTokenPermitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestTokenPermit *TestTokenPermitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.TestTokenPermitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestTokenPermit *TestTokenPermitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestTokenPermit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestTokenPermit *TestTokenPermitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestTokenPermit *TestTokenPermitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_TestTokenPermit *TestTokenPermitCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TestTokenPermit.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_TestTokenPermit *TestTokenPermitSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _TestTokenPermit.Contract.DOMAINSEPARATOR(&_TestTokenPermit.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_TestTokenPermit *TestTokenPermitCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _TestTokenPermit.Contract.DOMAINSEPARATOR(&_TestTokenPermit.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TestTokenPermit *TestTokenPermitCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestTokenPermit.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TestTokenPermit *TestTokenPermitSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TestTokenPermit.Contract.Allowance(&_TestTokenPermit.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TestTokenPermit *TestTokenPermitCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TestTokenPermit.Contract.Allowance(&_TestTokenPermit.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TestTokenPermit *TestTokenPermitCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestTokenPermit.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TestTokenPermit *TestTokenPermitSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TestTokenPermit.Contract.BalanceOf(&_TestTokenPermit.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TestTokenPermit *TestTokenPermitCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TestTokenPermit.Contract.BalanceOf(&_TestTokenPermit.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TestTokenPermit *TestTokenPermitCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TestTokenPermit.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TestTokenPermit *TestTokenPermitSession) Decimals() (uint8, error) {
	return _TestTokenPermit.Contract.Decimals(&_TestTokenPermit.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TestTokenPermit *TestTokenPermitCallerSession) Decimals() (uint8, error) {
	return _TestTokenPermit.Contract.Decimals(&_TestTokenPermit.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestTokenPermit *TestTokenPermitCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestTokenPermit.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestTokenPermit *TestTokenPermitSession) Name() (string, error) {
	return _TestTokenPermit.Contract.Name(&_TestTokenPermit.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestTokenPermit *TestTokenPermitCallerSession) Name() (string, error) {
	return _TestTokenPermit.Contract.Name(&_TestTokenPermit.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_TestTokenPermit *TestTokenPermitCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestTokenPermit.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_TestTokenPermit *TestTokenPermitSession) Nonces(owner common.Address) (*big.Int, error) {
	return _TestTokenPermit.Contract.Nonces(&_TestTokenPermit.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_TestTokenPermit *TestTokenPermitCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _TestTokenPermit.Contract.Nonces(&_TestTokenPermit.CallOpts, owner)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestTokenPermit *TestTokenPermitCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestTokenPermit.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestTokenPermit *TestTokenPermitSession) Symbol() (string, error) {
	return _TestTokenPermit.Contract.Symbol(&_TestTokenPermit.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestTokenPermit *TestTokenPermitCallerSession) Symbol() (string, error) {
	return _TestTokenPermit.Contract.Symbol(&_TestTokenPermit.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestTokenPermit *TestTokenPermitCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestTokenPermit.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestTokenPermit *TestTokenPermitSession) TotalSupply() (*big.Int, error) {
	return _TestTokenPermit.Contract.TotalSupply(&_TestTokenPermit.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestTokenPermit *TestTokenPermitCallerSession) TotalSupply() (*big.Int, error) {
	return _TestTokenPermit.Contract.TotalSupply(&_TestTokenPermit.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TestTokenPermit *TestTokenPermitTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.contract.Transact(opts, "approve", spender, amount)
}
func (_TestTokenPermit *TestTokenPermitTransactor) ApproveOverGsn(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (common.Hash, error) {
	return GsnExecutor(_TestTokenPermit.gsn, TestTokenPermitMetaData.ABI, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TestTokenPermit *TestTokenPermitSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.Approve(&_TestTokenPermit.TransactOpts, spender, amount)
}
func (_TestTokenPermit *TestTokenPermitSession) ApproveOverGsn(spender common.Address, amount *big.Int) (common.Hash, error) {
	return _TestTokenPermit.Contract.ApproveOverGsn(&_TestTokenPermit.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TestTokenPermit *TestTokenPermitTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.Approve(&_TestTokenPermit.TransactOpts, spender, amount)
}
func (_TestTokenPermit *TestTokenPermitTransactorSession) ApproveOverGsn(spender common.Address, amount *big.Int) (common.Hash, error) {
	return _TestTokenPermit.Contract.ApproveOverGsn(&_TestTokenPermit.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_TestTokenPermit *TestTokenPermitTransactor) Burn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.contract.Transact(opts, "burn", account, amount)
}
func (_TestTokenPermit *TestTokenPermitTransactor) BurnOverGsn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (common.Hash, error) {
	return GsnExecutor(_TestTokenPermit.gsn, TestTokenPermitMetaData.ABI, "burn", account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_TestTokenPermit *TestTokenPermitSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.Burn(&_TestTokenPermit.TransactOpts, account, amount)
}
func (_TestTokenPermit *TestTokenPermitSession) BurnOverGsn(account common.Address, amount *big.Int) (common.Hash, error) {
	return _TestTokenPermit.Contract.BurnOverGsn(&_TestTokenPermit.TransactOpts, account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_TestTokenPermit *TestTokenPermitTransactorSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.Burn(&_TestTokenPermit.TransactOpts, account, amount)
}
func (_TestTokenPermit *TestTokenPermitTransactorSession) BurnOverGsn(account common.Address, amount *big.Int) (common.Hash, error) {
	return _TestTokenPermit.Contract.BurnOverGsn(&_TestTokenPermit.TransactOpts, account, amount)
}

// BurnWithAllowanceDecrease is a paid mutator transaction binding the contract method 0xa918adf5.
//
// Solidity: function burnWithAllowanceDecrease(address account, address spender, uint256 amount) returns()
func (_TestTokenPermit *TestTokenPermitTransactor) BurnWithAllowanceDecrease(opts *bind.TransactOpts, account common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.contract.Transact(opts, "burnWithAllowanceDecrease", account, spender, amount)
}
func (_TestTokenPermit *TestTokenPermitTransactor) BurnWithAllowanceDecreaseOverGsn(opts *bind.TransactOpts, account common.Address, spender common.Address, amount *big.Int) (common.Hash, error) {
	return GsnExecutor(_TestTokenPermit.gsn, TestTokenPermitMetaData.ABI, "burnWithAllowanceDecrease", account, spender, amount)
}

// BurnWithAllowanceDecrease is a paid mutator transaction binding the contract method 0xa918adf5.
//
// Solidity: function burnWithAllowanceDecrease(address account, address spender, uint256 amount) returns()
func (_TestTokenPermit *TestTokenPermitSession) BurnWithAllowanceDecrease(account common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.BurnWithAllowanceDecrease(&_TestTokenPermit.TransactOpts, account, spender, amount)
}
func (_TestTokenPermit *TestTokenPermitSession) BurnWithAllowanceDecreaseOverGsn(account common.Address, spender common.Address, amount *big.Int) (common.Hash, error) {
	return _TestTokenPermit.Contract.BurnWithAllowanceDecreaseOverGsn(&_TestTokenPermit.TransactOpts, account, spender, amount)
}

// BurnWithAllowanceDecrease is a paid mutator transaction binding the contract method 0xa918adf5.
//
// Solidity: function burnWithAllowanceDecrease(address account, address spender, uint256 amount) returns()
func (_TestTokenPermit *TestTokenPermitTransactorSession) BurnWithAllowanceDecrease(account common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.BurnWithAllowanceDecrease(&_TestTokenPermit.TransactOpts, account, spender, amount)
}
func (_TestTokenPermit *TestTokenPermitTransactorSession) BurnWithAllowanceDecreaseOverGsn(account common.Address, spender common.Address, amount *big.Int) (common.Hash, error) {
	return _TestTokenPermit.Contract.BurnWithAllowanceDecreaseOverGsn(&_TestTokenPermit.TransactOpts, account, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TestTokenPermit *TestTokenPermitTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}
func (_TestTokenPermit *TestTokenPermitTransactor) DecreaseAllowanceOverGsn(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (common.Hash, error) {
	return GsnExecutor(_TestTokenPermit.gsn, TestTokenPermitMetaData.ABI, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TestTokenPermit *TestTokenPermitSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.DecreaseAllowance(&_TestTokenPermit.TransactOpts, spender, subtractedValue)
}
func (_TestTokenPermit *TestTokenPermitSession) DecreaseAllowanceOverGsn(spender common.Address, subtractedValue *big.Int) (common.Hash, error) {
	return _TestTokenPermit.Contract.DecreaseAllowanceOverGsn(&_TestTokenPermit.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TestTokenPermit *TestTokenPermitTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.DecreaseAllowance(&_TestTokenPermit.TransactOpts, spender, subtractedValue)
}
func (_TestTokenPermit *TestTokenPermitTransactorSession) DecreaseAllowanceOverGsn(spender common.Address, subtractedValue *big.Int) (common.Hash, error) {
	return _TestTokenPermit.Contract.DecreaseAllowanceOverGsn(&_TestTokenPermit.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TestTokenPermit *TestTokenPermitTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}
func (_TestTokenPermit *TestTokenPermitTransactor) IncreaseAllowanceOverGsn(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (common.Hash, error) {
	return GsnExecutor(_TestTokenPermit.gsn, TestTokenPermitMetaData.ABI, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TestTokenPermit *TestTokenPermitSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.IncreaseAllowance(&_TestTokenPermit.TransactOpts, spender, addedValue)
}
func (_TestTokenPermit *TestTokenPermitSession) IncreaseAllowanceOverGsn(spender common.Address, addedValue *big.Int) (common.Hash, error) {
	return _TestTokenPermit.Contract.IncreaseAllowanceOverGsn(&_TestTokenPermit.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TestTokenPermit *TestTokenPermitTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.IncreaseAllowance(&_TestTokenPermit.TransactOpts, spender, addedValue)
}
func (_TestTokenPermit *TestTokenPermitTransactorSession) IncreaseAllowanceOverGsn(spender common.Address, addedValue *big.Int) (common.Hash, error) {
	return _TestTokenPermit.Contract.IncreaseAllowanceOverGsn(&_TestTokenPermit.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_TestTokenPermit *TestTokenPermitTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.contract.Transact(opts, "mint", account, amount)
}
func (_TestTokenPermit *TestTokenPermitTransactor) MintOverGsn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (common.Hash, error) {
	return GsnExecutor(_TestTokenPermit.gsn, TestTokenPermitMetaData.ABI, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_TestTokenPermit *TestTokenPermitSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.Mint(&_TestTokenPermit.TransactOpts, account, amount)
}
func (_TestTokenPermit *TestTokenPermitSession) MintOverGsn(account common.Address, amount *big.Int) (common.Hash, error) {
	return _TestTokenPermit.Contract.MintOverGsn(&_TestTokenPermit.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_TestTokenPermit *TestTokenPermitTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.Mint(&_TestTokenPermit.TransactOpts, account, amount)
}
func (_TestTokenPermit *TestTokenPermitTransactorSession) MintOverGsn(account common.Address, amount *big.Int) (common.Hash, error) {
	return _TestTokenPermit.Contract.MintOverGsn(&_TestTokenPermit.TransactOpts, account, amount)
}

// MintWithAllowance is a paid mutator transaction binding the contract method 0x9be4e7b2.
//
// Solidity: function mintWithAllowance(address account, address spender, uint256 amount) returns()
func (_TestTokenPermit *TestTokenPermitTransactor) MintWithAllowance(opts *bind.TransactOpts, account common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.contract.Transact(opts, "mintWithAllowance", account, spender, amount)
}
func (_TestTokenPermit *TestTokenPermitTransactor) MintWithAllowanceOverGsn(opts *bind.TransactOpts, account common.Address, spender common.Address, amount *big.Int) (common.Hash, error) {
	return GsnExecutor(_TestTokenPermit.gsn, TestTokenPermitMetaData.ABI, "mintWithAllowance", account, spender, amount)
}

// MintWithAllowance is a paid mutator transaction binding the contract method 0x9be4e7b2.
//
// Solidity: function mintWithAllowance(address account, address spender, uint256 amount) returns()
func (_TestTokenPermit *TestTokenPermitSession) MintWithAllowance(account common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.MintWithAllowance(&_TestTokenPermit.TransactOpts, account, spender, amount)
}
func (_TestTokenPermit *TestTokenPermitSession) MintWithAllowanceOverGsn(account common.Address, spender common.Address, amount *big.Int) (common.Hash, error) {
	return _TestTokenPermit.Contract.MintWithAllowanceOverGsn(&_TestTokenPermit.TransactOpts, account, spender, amount)
}

// MintWithAllowance is a paid mutator transaction binding the contract method 0x9be4e7b2.
//
// Solidity: function mintWithAllowance(address account, address spender, uint256 amount) returns()
func (_TestTokenPermit *TestTokenPermitTransactorSession) MintWithAllowance(account common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.MintWithAllowance(&_TestTokenPermit.TransactOpts, account, spender, amount)
}
func (_TestTokenPermit *TestTokenPermitTransactorSession) MintWithAllowanceOverGsn(account common.Address, spender common.Address, amount *big.Int) (common.Hash, error) {
	return _TestTokenPermit.Contract.MintWithAllowanceOverGsn(&_TestTokenPermit.TransactOpts, account, spender, amount)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_TestTokenPermit *TestTokenPermitTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _TestTokenPermit.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}
func (_TestTokenPermit *TestTokenPermitTransactor) PermitOverGsn(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (common.Hash, error) {
	return GsnExecutor(_TestTokenPermit.gsn, TestTokenPermitMetaData.ABI, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_TestTokenPermit *TestTokenPermitSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.Permit(&_TestTokenPermit.TransactOpts, owner, spender, value, deadline, v, r, s)
}
func (_TestTokenPermit *TestTokenPermitSession) PermitOverGsn(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (common.Hash, error) {
	return _TestTokenPermit.Contract.PermitOverGsn(&_TestTokenPermit.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_TestTokenPermit *TestTokenPermitTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.Permit(&_TestTokenPermit.TransactOpts, owner, spender, value, deadline, v, r, s)
}
func (_TestTokenPermit *TestTokenPermitTransactorSession) PermitOverGsn(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (common.Hash, error) {
	return _TestTokenPermit.Contract.PermitOverGsn(&_TestTokenPermit.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_TestTokenPermit *TestTokenPermitTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.contract.Transact(opts, "transfer", recipient, amount)
}
func (_TestTokenPermit *TestTokenPermitTransactor) TransferOverGsn(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (common.Hash, error) {
	return GsnExecutor(_TestTokenPermit.gsn, TestTokenPermitMetaData.ABI, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_TestTokenPermit *TestTokenPermitSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.Transfer(&_TestTokenPermit.TransactOpts, recipient, amount)
}
func (_TestTokenPermit *TestTokenPermitSession) TransferOverGsn(recipient common.Address, amount *big.Int) (common.Hash, error) {
	return _TestTokenPermit.Contract.TransferOverGsn(&_TestTokenPermit.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_TestTokenPermit *TestTokenPermitTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.Transfer(&_TestTokenPermit.TransactOpts, recipient, amount)
}
func (_TestTokenPermit *TestTokenPermitTransactorSession) TransferOverGsn(recipient common.Address, amount *big.Int) (common.Hash, error) {
	return _TestTokenPermit.Contract.TransferOverGsn(&_TestTokenPermit.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_TestTokenPermit *TestTokenPermitTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}
func (_TestTokenPermit *TestTokenPermitTransactor) TransferFromOverGsn(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (common.Hash, error) {
	return GsnExecutor(_TestTokenPermit.gsn, TestTokenPermitMetaData.ABI, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_TestTokenPermit *TestTokenPermitSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.TransferFrom(&_TestTokenPermit.TransactOpts, sender, recipient, amount)
}
func (_TestTokenPermit *TestTokenPermitSession) TransferFromOverGsn(sender common.Address, recipient common.Address, amount *big.Int) (common.Hash, error) {
	return _TestTokenPermit.Contract.TransferFromOverGsn(&_TestTokenPermit.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_TestTokenPermit *TestTokenPermitTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.TransferFrom(&_TestTokenPermit.TransactOpts, sender, recipient, amount)
}
func (_TestTokenPermit *TestTokenPermitTransactorSession) TransferFromOverGsn(sender common.Address, recipient common.Address, amount *big.Int) (common.Hash, error) {
	return _TestTokenPermit.Contract.TransferFromOverGsn(&_TestTokenPermit.TransactOpts, sender, recipient, amount)
}

// TestTokenPermitApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TestTokenPermit contract.
type TestTokenPermitApprovalIterator struct {
	Event *TestTokenPermitApproval // Event containing the contract specifics and raw log

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
func (it *TestTokenPermitApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestTokenPermitApproval)
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
		it.Event = new(TestTokenPermitApproval)
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
func (it *TestTokenPermitApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestTokenPermitApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestTokenPermitApproval represents a Approval event raised by the TestTokenPermit contract.
type TestTokenPermitApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestTokenPermit *TestTokenPermitFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TestTokenPermitApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TestTokenPermit.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TestTokenPermitApprovalIterator{contract: _TestTokenPermit.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestTokenPermit *TestTokenPermitFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TestTokenPermitApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TestTokenPermit.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestTokenPermitApproval)
				if err := _TestTokenPermit.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestTokenPermit *TestTokenPermitFilterer) ParseApproval(log types.Log) (*TestTokenPermitApproval, error) {
	event := new(TestTokenPermitApproval)
	if err := _TestTokenPermit.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestTokenPermitTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TestTokenPermit contract.
type TestTokenPermitTransferIterator struct {
	Event *TestTokenPermitTransfer // Event containing the contract specifics and raw log

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
func (it *TestTokenPermitTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestTokenPermitTransfer)
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
		it.Event = new(TestTokenPermitTransfer)
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
func (it *TestTokenPermitTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestTokenPermitTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestTokenPermitTransfer represents a Transfer event raised by the TestTokenPermit contract.
type TestTokenPermitTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestTokenPermit *TestTokenPermitFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TestTokenPermitTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestTokenPermit.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TestTokenPermitTransferIterator{contract: _TestTokenPermit.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestTokenPermit *TestTokenPermitFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TestTokenPermitTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestTokenPermit.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestTokenPermitTransfer)
				if err := _TestTokenPermit.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestTokenPermit *TestTokenPermitFilterer) ParseTransfer(log types.Log) (*TestTokenPermitTransfer, error) {
	event := new(TestTokenPermitTransfer)
	if err := _TestTokenPermit.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
