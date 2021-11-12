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

// TestTokenPermitABI is the input ABI used to generate the binding from.
const TestTokenPermitABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnWithAllowanceDecrease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintWithAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TestTokenPermitBin is the compiled bytecode used for deploying new contracts.
var TestTokenPermitBin = "0x6101406040527f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9610120523480156200003757600080fd5b506040516200175f3803806200175f8339810160408190526200005a91620002b9565b604051806040016040528060048152602001634559574160e01b81525080604051806040016040528060018152602001603160f81b81525084848160039080519060200190620000ac92919062000168565b508051620000c290600490602084019062000168565b5050825160208085019190912083519184019190912060c082905260e08190524660a0529091507f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f620001178184846200012c565b60805261010052506200039f95505050505050565b600083838346306040516020016200014995949392919062000320565b6040516020818303038152906040528051906020012090509392505050565b82805462000176906200034c565b90600052602060002090601f0160209004810192826200019a5760008555620001e5565b82601f10620001b557805160ff1916838001178555620001e5565b82800160010185558215620001e5579182015b82811115620001e5578251825591602001919060010190620001c8565b50620001f3929150620001f7565b5090565b5b80821115620001f35760008155600101620001f8565b600082601f8301126200021f578081fd5b81516001600160401b03808211156200023c576200023c62000389565b6040516020601f8401601f191682018101838111838210171562000264576200026462000389565b60405283825285840181018710156200027b578485fd5b8492505b838310156200029e57858301810151828401820152918201916200027f565b83831115620002af57848185840101525b5095945050505050565b60008060408385031215620002cc578182fd5b82516001600160401b0380821115620002e3578384fd5b620002f1868387016200020e565b9350602085015191508082111562000307578283fd5b5062000316858286016200020e565b9150509250929050565b9485526020850193909352604084019190915260608301526001600160a01b0316608082015260a00190565b6002810460018216806200036157607f821691505b602082108114156200038357634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b60805160a05160c05160e0516101005161012051611370620003ef60003960006105c8015260006108e7015260006109290152600061090801526000610895015260006108be01526113706000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c80637ecebe00116100a2578063a457c2d711610071578063a457c2d71461021a578063a9059cbb1461022d578063a918adf514610240578063d505accf14610253578063dd62ed3e1461026657610116565b80637ecebe00146101d957806395d89b41146101ec5780639be4e7b2146101f45780639dc29fac1461020757610116565b8063313ce567116100e9578063313ce567146101815780633644e51514610196578063395093511461019e57806340c10f19146101b157806370a08231146101c657610116565b806306fdde031461011b578063095ea7b31461013957806318160ddd1461015957806323b872dd1461016e575b600080fd5b610123610279565b6040516101309190610e9b565b60405180910390f35b61014c610147366004610dc5565b61030c565b6040516101309190610e09565b610161610329565b6040516101309190610e14565b61014c61017c366004610d19565b61032f565b6101896103cf565b60405161013091906112b2565b6101616103d4565b61014c6101ac366004610dc5565b6103e3565b6101c46101bf366004610dc5565b610432565b005b6101616101d4366004610cc6565b610440565b6101616101e7366004610cc6565b61045f565b610123610486565b6101c4610202366004610d19565b610495565b6101c4610215366004610dc5565b6104be565b61014c610228366004610dc5565b6104c8565b61014c61023b366004610dc5565b610543565b6101c461024e366004610d19565b610557565b6101c4610261366004610d54565b6105a4565b610161610274366004610ce7565b610686565b606060038054610288906112ef565b80601f01602080910402602001604051908101604052809291908181526020018280546102b4906112ef565b80156103015780601f106102d657610100808354040283529160200191610301565b820191906000526020600020905b8154815290600101906020018083116102e457829003601f168201915b505050505090505b90565b60006103206103196106b1565b84846106b5565b50600192915050565b60025490565b600061033c848484610769565b6001600160a01b03841660009081526001602052604081208161035d6106b1565b6001600160a01b03166001600160a01b03168152602001908152602001600020549050828110156103a95760405162461bcd60e51b81526004016103a090611124565b60405180910390fd5b6103c4856103b56106b1565b6103bf86856112d8565b6106b5565b506001949350505050565b601290565b60006103de610891565b905090565b60006103206103f06106b1565b8484600160006103fe6106b1565b6001600160a01b03908116825260208083019390935260409182016000908120918b16815292529020546103bf91906112c0565b61043c8282610954565b5050565b6001600160a01b0381166000908152602081905260409020545b919050565b6001600160a01b038116600090815260056020526040812061048090610a14565b92915050565b606060048054610288906112ef565b61049f8382610954565b6104b98383836104af8787610686565b6103bf91906112c0565b505050565b61043c8282610a18565b600080600160006104d76106b1565b6001600160a01b03908116825260208083019390935260409182016000908120918816815292529020549050828110156105235760405162461bcd60e51b81526004016103a090611236565b61053961052e6106b1565b856103bf86856112d8565b5060019392505050565b60006103206105506106b1565b8484610769565b60006105638484610686565b9050818110156105855760405162461bcd60e51b81526004016103a090611236565b61059484846103bf85856112d8565b61059e8483610a18565b50505050565b834211156105c45760405162461bcd60e51b81526004016103a090610fec565b60007f00000000000000000000000000000000000000000000000000000000000000008888886105f38c610afe565b8960405160200161060996959493929190610e1d565b604051602081830303815290604052805190602001209050600061062c82610b30565b9050600061063c82878787610b43565b9050896001600160a01b0316816001600160a01b03161461066f5760405162461bcd60e51b81526004016103a0906110ed565b61067a8a8a8a6106b5565b50505050505050505050565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b3390565b6001600160a01b0383166106db5760405162461bcd60e51b81526004016103a0906111f2565b6001600160a01b0382166107015760405162461bcd60e51b81526004016103a090610faa565b6001600160a01b0380841660008181526001602090815260408083209487168084529490915290819020849055517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259061075c908590610e14565b60405180910390a3505050565b6001600160a01b03831661078f5760405162461bcd60e51b81526004016103a0906111ad565b6001600160a01b0382166107b55760405162461bcd60e51b81526004016103a090610f25565b6107c08383836104b9565b6001600160a01b038316600090815260208190526040902054818110156107f95760405162461bcd60e51b81526004016103a090611023565b61080382826112d8565b6001600160a01b0380861660009081526020819052604080822093909355908516815290812080548492906108399084906112c0565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516108839190610e14565b60405180910390a350505050565b60007f00000000000000000000000000000000000000000000000000000000000000004614156108e257507f0000000000000000000000000000000000000000000000000000000000000000610309565b61094d7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000610c39565b9050610309565b6001600160a01b03821661097a5760405162461bcd60e51b81526004016103a09061127b565b610986600083836104b9565b806002600082825461099891906112c0565b90915550506001600160a01b038216600090815260208190526040812080548392906109c59084906112c0565b90915550506040516001600160a01b038316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90610a08908590610e14565b60405180910390a35050565b5490565b6001600160a01b038216610a3e5760405162461bcd60e51b81526004016103a09061116c565b610a4a826000836104b9565b6001600160a01b03821660009081526020819052604090205481811015610a835760405162461bcd60e51b81526004016103a090610f68565b610a8d82826112d8565b6001600160a01b03841660009081526020819052604081209190915560028054849290610abb9084906112d8565b90915550506040516000906001600160a01b038516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9061075c908690610e14565b6001600160a01b0381166000908152600560205260408120610b1f81610a14565b9150610b2a81610c73565b50919050565b6000610480610b3d610891565b83610c7c565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0821115610b855760405162461bcd60e51b81526004016103a090611069565b8360ff16601b1480610b9a57508360ff16601c145b610bb65760405162461bcd60e51b81526004016103a0906110ab565b600060018686868660405160008152602001604052604051610bdb9493929190610e7d565b6020604051602081039080840390855afa158015610bfd573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116610c305760405162461bcd60e51b81526004016103a090610eee565b95945050505050565b60008383834630604051602001610c54959493929190610e51565b6040516020818303038152906040528051906020012090509392505050565b80546001019055565b60008282604051602001610c91929190610dee565b60405160208183030381529060405280519060200120905092915050565b80356001600160a01b038116811461045a57600080fd5b600060208284031215610cd7578081fd5b610ce082610caf565b9392505050565b60008060408385031215610cf9578081fd5b610d0283610caf565b9150610d1060208401610caf565b90509250929050565b600080600060608486031215610d2d578081fd5b610d3684610caf565b9250610d4460208501610caf565b9150604084013590509250925092565b600080600080600080600060e0888a031215610d6e578283fd5b610d7788610caf565b9650610d8560208901610caf565b95506040880135945060608801359350608088013560ff81168114610da8578384fd5b9699959850939692959460a0840135945060c09093013592915050565b60008060408385031215610dd7578182fd5b610de083610caf565b946020939093013593505050565b61190160f01b81526002810192909252602282015260420190565b901515815260200190565b90815260200190565b9586526001600160a01b0394851660208701529290931660408501526060840152608083019190915260a082015260c00190565b9485526020850193909352604084019190915260608301526001600160a01b0316608082015260a00190565b93845260ff9290921660208401526040830152606082015260800190565b6000602080835283518082850152825b81811015610ec757858101830151858201604001528201610eab565b81811115610ed85783604083870101525b50601f01601f1916929092016040019392505050565b60208082526018908201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604082015260600190565b60208082526023908201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260408201526265737360e81b606082015260800190565b60208082526022908201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604082015261636560f01b606082015260800190565b60208082526022908201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604082015261737360f01b606082015260800190565b6020808252601d908201527f45524332305065726d69743a206578706972656420646561646c696e65000000604082015260600190565b60208082526026908201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604082015265616c616e636560d01b606082015260800190565b60208082526022908201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604082015261756560f01b606082015260800190565b60208082526022908201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604082015261756560f01b606082015260800190565b6020808252601e908201527f45524332305065726d69743a20696e76616c6964207369676e61747572650000604082015260600190565b60208082526028908201527f45524332303a207472616e7366657220616d6f756e74206578636565647320616040820152676c6c6f77616e636560c01b606082015260800190565b60208082526021908201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736040820152607360f81b606082015260800190565b60208082526025908201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604082015264647265737360d81b606082015260800190565b60208082526024908201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646040820152637265737360e01b606082015260800190565b60208082526025908201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604082015264207a65726f60d81b606082015260800190565b6020808252601f908201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604082015260600190565b60ff91909116815260200190565b600082198211156112d3576112d3611324565b500190565b6000828210156112ea576112ea611324565b500390565b60028104600182168061130357607f821691505b60208210811415610b2a57634e487b7160e01b600052602260045260246000fd5b634e487b7160e01b600052601160045260246000fdfea264697066735822122030d7f921892734d5817def7c921f8d76290b106fcb113e443f3b5f6e9e066ee064736f6c63430008000033"

// DeployTestTokenPermit deploys a new Ethereum contract, binding an instance of TestTokenPermit to it.
func DeployTestTokenPermit(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string) (common.Address, *types.Transaction, *TestTokenPermit, error) {
	parsed, err := abi.JSON(strings.NewReader(TestTokenPermitABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TestTokenPermitBin), backend, name_, symbol_)
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

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TestTokenPermit *TestTokenPermitSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.Approve(&_TestTokenPermit.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TestTokenPermit *TestTokenPermitTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.Approve(&_TestTokenPermit.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_TestTokenPermit *TestTokenPermitTransactor) Burn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.contract.Transact(opts, "burn", account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_TestTokenPermit *TestTokenPermitSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.Burn(&_TestTokenPermit.TransactOpts, account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_TestTokenPermit *TestTokenPermitTransactorSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.Burn(&_TestTokenPermit.TransactOpts, account, amount)
}

// BurnWithAllowanceDecrease is a paid mutator transaction binding the contract method 0xa918adf5.
//
// Solidity: function burnWithAllowanceDecrease(address account, address spender, uint256 amount) returns()
func (_TestTokenPermit *TestTokenPermitTransactor) BurnWithAllowanceDecrease(opts *bind.TransactOpts, account common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.contract.Transact(opts, "burnWithAllowanceDecrease", account, spender, amount)
}

// BurnWithAllowanceDecrease is a paid mutator transaction binding the contract method 0xa918adf5.
//
// Solidity: function burnWithAllowanceDecrease(address account, address spender, uint256 amount) returns()
func (_TestTokenPermit *TestTokenPermitSession) BurnWithAllowanceDecrease(account common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.BurnWithAllowanceDecrease(&_TestTokenPermit.TransactOpts, account, spender, amount)
}

// BurnWithAllowanceDecrease is a paid mutator transaction binding the contract method 0xa918adf5.
//
// Solidity: function burnWithAllowanceDecrease(address account, address spender, uint256 amount) returns()
func (_TestTokenPermit *TestTokenPermitTransactorSession) BurnWithAllowanceDecrease(account common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.BurnWithAllowanceDecrease(&_TestTokenPermit.TransactOpts, account, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TestTokenPermit *TestTokenPermitTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TestTokenPermit *TestTokenPermitSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.DecreaseAllowance(&_TestTokenPermit.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TestTokenPermit *TestTokenPermitTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.DecreaseAllowance(&_TestTokenPermit.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TestTokenPermit *TestTokenPermitTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TestTokenPermit *TestTokenPermitSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.IncreaseAllowance(&_TestTokenPermit.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TestTokenPermit *TestTokenPermitTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.IncreaseAllowance(&_TestTokenPermit.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_TestTokenPermit *TestTokenPermitTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_TestTokenPermit *TestTokenPermitSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.Mint(&_TestTokenPermit.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_TestTokenPermit *TestTokenPermitTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.Mint(&_TestTokenPermit.TransactOpts, account, amount)
}

// MintWithAllowance is a paid mutator transaction binding the contract method 0x9be4e7b2.
//
// Solidity: function mintWithAllowance(address account, address spender, uint256 amount) returns()
func (_TestTokenPermit *TestTokenPermitTransactor) MintWithAllowance(opts *bind.TransactOpts, account common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.contract.Transact(opts, "mintWithAllowance", account, spender, amount)
}

// MintWithAllowance is a paid mutator transaction binding the contract method 0x9be4e7b2.
//
// Solidity: function mintWithAllowance(address account, address spender, uint256 amount) returns()
func (_TestTokenPermit *TestTokenPermitSession) MintWithAllowance(account common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.MintWithAllowance(&_TestTokenPermit.TransactOpts, account, spender, amount)
}

// MintWithAllowance is a paid mutator transaction binding the contract method 0x9be4e7b2.
//
// Solidity: function mintWithAllowance(address account, address spender, uint256 amount) returns()
func (_TestTokenPermit *TestTokenPermitTransactorSession) MintWithAllowance(account common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.MintWithAllowance(&_TestTokenPermit.TransactOpts, account, spender, amount)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_TestTokenPermit *TestTokenPermitTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _TestTokenPermit.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_TestTokenPermit *TestTokenPermitSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.Permit(&_TestTokenPermit.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_TestTokenPermit *TestTokenPermitTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.Permit(&_TestTokenPermit.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_TestTokenPermit *TestTokenPermitTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_TestTokenPermit *TestTokenPermitSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.Transfer(&_TestTokenPermit.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_TestTokenPermit *TestTokenPermitTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.Transfer(&_TestTokenPermit.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_TestTokenPermit *TestTokenPermitTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_TestTokenPermit *TestTokenPermitSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.TransferFrom(&_TestTokenPermit.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_TestTokenPermit *TestTokenPermitTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestTokenPermit.Contract.TransferFrom(&_TestTokenPermit.TransactOpts, sender, recipient, amount)
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
