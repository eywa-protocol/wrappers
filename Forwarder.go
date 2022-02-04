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

// ForwarderMetaData contains all meta data concerning the Forwarder contract.
var ForwarderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"typeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"typeStr\",\"type\":\"string\"}],\"name\":\"RequestTypeRegistered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GENERIC_PARAMS\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"requestTypeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"suffixData\",\"type\":\"bytes\"}],\"name\":\"_getEncoded\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"requestTypeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"suffixData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"ret\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"execute2\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"executeAssemblyForwarderRequest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"reqAbiEncode\",\"type\":\"bytes\"}],\"name\":\"getAbiEncodeRequest\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"typeName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"typeSuffix\",\"type\":\"string\"}],\"name\":\"registerRequestType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"typeHashes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"requestTypeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"suffixData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060006040518060800160405280604a81526020016200163f604a9139604051602001620000409190620000fb565b60408051601f1981840301815291905290506200005d8162000064565b5062000174565b8051602080830191909120600081815291829052604091829020805460ff19166001179055905181907f64d6bce64323458c44643c51fe45113efc882082f7b7fd5f09f0d69d2eedb20290620000bc9085906200013f565b60405180910390a25050565b60005b83811015620000e5578181015183820152602001620000cb565b83811115620000f5576000848401525b50505050565b6e08cdee4eec2e4c8a4cae2eacae6e85608b1b8152600082516200012781600f850160208701620000c8565b602960f81b600f939091019283015250601001919050565b602081526000825180602084015262000160816040850160208701620000c8565b601f01601f19169190910160400192915050565b6114bb80620001846000396000f3fe6080604052600436106100955760003560e01c80639d5eaf64116100595780639d5eaf64146101a0578063c45c2674146101c0578063cdf5735a146101e0578063d9210be5146101f3578063f26625451461021357600080fd5b8063066a310c146100a157806321fe98df146100cc5780632d0335ab1461010c5780636a66a1e214610150578063992d82c31461017e57600080fd5b3661009c57005b600080fd5b3480156100ad57600080fd5b506100b6610226565b6040516100c39190610def565b60405180910390f35b3480156100d857600080fd5b506100fc6100e7366004610e02565b60006020819052908152604090205460ff1681565b60405190151581526020016100c3565b34801561011857600080fd5b50610142610127366004610e37565b6001600160a01b031660009081526001602052604090205490565b6040519081526020016100c3565b34801561015c57600080fd5b5061017061016b366004610f9a565b610242565b6040516100c3929190610fd7565b34801561018a57600080fd5b5061019e61019936600461103b565b61032c565b005b3480156101ac57600080fd5b506100b66101bb3660046110e3565b6103b4565b3480156101cc57600080fd5b506100b66101db366004611147565b610542565b6101706101ee3660046111b4565b6105e8565b3480156101ff57600080fd5b5061019e61020e366004611256565b61072e565b610170610221366004610f9a565b610827565b6040518060800160405280604a815260200161143c604a913981565b600060608260800151600161025791906112d8565b83516001600160a01b031660009081526001602090815260409182902092909255908401518185015160a0860151805193519293919290919060c060288237600080838387895af16102ad573d6000803e3d6000fd5b50603f87606001516102bf91906112f0565b5a116102cd576102cd611312565b604080516001808252818301909252600091602082018180368337019050509050600560f81b8160008151811061030657610306611328565b60200101906001600160f81b031916908160001a90535060019890975095505050505050565b6103358761092a565b6103ab87878787878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8b01819004810282018101909252898152925089915088908190840183828082843760009201919091525061098a92505050565b50505050505050565b60606000836000015184602001518560400151866060015187608001516040516020016104119594939291906001600160a01b03958616815293909416602084015260408301919091526060820152608081019190915260a00190565b60408051601f1981840301815291905284519091506001600160a01b031661046b5760405162461bcd60e51b81526020600482015260086024820152677265712e66726f6d60c01b60448201526064015b60405180910390fd5b60208401516001600160a01b03166104ae5760405162461bcd60e51b815260206004820152600660248201526507265712e74360d41b6044820152606401610462565b60808401516104eb5760405162461bcd60e51b81526020600482015260096024820152687265712e6e6f6e636560b81b6044820152606401610462565b828051906020012081805190602001201461053b5760405162461bcd60e51b815260206004820152601060248201526f0d0c2e6d0cae640dcdee840dac2e8c6d60831b6044820152606401610462565b9392505050565b606082846000015185602001518660400151876060015188608001518960a00151805190602001206040516020016105b0969594939291906001600160a01b03968716815294909516602085015260408401929092526060830152608082015260a081019190915260c00190565b60408051601f19818403018152908290526105d09291859060200161133e565b60405160208183030381529060405290509392505050565b600060606105f58861092a565b6106388888888888888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061098a92505050565b61064188610aa5565b87602001516001600160a01b03165a6040808b015160a08c01518c519251919261066d9260200161137b565b60408051601f1981840301815290829052610687916113b2565b600060405180830381858888f193505050503d80600081146106c5576040519150601f19603f3d011682016040523d82523d6000602084013e6106ca565b606091505b50909250905081610723573d80156106e857604051816000823e8181fd5b5060405162461bcd60e51b815260206004820152600f60248201526e756e6b6e6f776e206661696c75726560881b6044820152606401610462565b965096945050505050565b60005b838110156107d257600085858381811061074d5761074d611328565b909101356001600160f81b031916915050600560fb1b81148015906107805750602960f81b6001600160f81b0319821614155b6107bf5760405162461bcd60e51b815260206004820152601060248201526f696e76616c696420747970656e616d6560801b6044820152606401610462565b50806107ca816113ce565b915050610731565b50600084846040518060800160405280604a815260200161143c604a913985856040516020016108069594939291906113e9565b604051602081830303815290604052905061082081610ad2565b5050505050565b600060608260800151600161083c91906112d8565b6001600085600001516001600160a01b03166001600160a01b031681526020019081526020016000208190555060008084602001516001600160a01b0316856060015186604001518760a00151886000015160405160200161089f92919061137b565b60408051601f19818403018152908290526108b9916113b2565b600060405180830381858888f193505050503d80600081146108f7576040519150601f19603f3d011682016040523d82523d6000602084013e6108fc565b606091505b5091509150603f856060015161091291906112f0565b5a1161092057610920611312565b9094909350915050565b608081015181516001600160a01b0316600090815260016020526040902054146109875760405162461bcd60e51b815260206004820152600e60248201526d0dcdedcc6ca40dad2e6dac2e8c6d60931b6044820152606401610462565b50565b60008381526020819052604090205460ff166109e85760405162461bcd60e51b815260206004820152601860248201527f696e76616c6964207265717565737420747970656861736800000000000000006044820152606401610462565b6000846109f6878686610542565b8051602091820120604051610a2293920161190160f01b81526002810192909252602282015260420190565b60408051601f19818403018152919052805160209091012086519091506001600160a01b0316610a528284610b34565b6001600160a01b031614610a9d5760405162461bcd60e51b81526020600482015260126024820152710e6d2cedcc2e8eae4ca40dad2e6dac2e8c6d60731b6044820152606401610462565b505050505050565b80516001600160a01b03166000908152600160205260408120805491610aca836113ce565b919050555050565b8051602080830191909120600081815291829052604091829020805460ff19166001179055905181907f64d6bce64323458c44643c51fe45113efc882082f7b7fd5f09f0d69d2eedb20290610b28908590610def565b60405180910390a25050565b600080600080845160411415610b5e5750505060208201516040830151606084015160001a610bd4565b845160401415610b8c5750505060408201516020830151906001600160ff1b0381169060ff1c601b01610bd4565b60405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610462565b610be086828585610bea565b9695505050505050565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0821115610c675760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610462565b8360ff16601b1480610c7c57508360ff16601c145b610cd35760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610462565b6040805160008082526020820180845288905260ff871692820192909252606081018590526080810184905260019060a0016020604051602081039080840390855afa158015610d27573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116610d8a5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610462565b95945050505050565b60005b83811015610dae578181015183820152602001610d96565b83811115610dbd576000848401525b50505050565b60008151808452610ddb816020860160208601610d93565b601f01601f19169290920160200192915050565b60208152600061053b6020830184610dc3565b600060208284031215610e1457600080fd5b5035919050565b80356001600160a01b0381168114610e3257600080fd5b919050565b600060208284031215610e4957600080fd5b61053b82610e1b565b634e487b7160e01b600052604160045260246000fd5b600082601f830112610e7957600080fd5b813567ffffffffffffffff80821115610e9457610e94610e52565b604051601f8301601f19908116603f01168101908282118183101715610ebc57610ebc610e52565b81604052838152866020858801011115610ed557600080fd5b836020870160208301376000602085830101528094505050505092915050565b600060c08284031215610f0757600080fd5b60405160c0810167ffffffffffffffff8282108183111715610f2b57610f2b610e52565b81604052829350610f3b85610e1b565b8352610f4960208601610e1b565b602084015260408501356040840152606085013560608401526080850135608084015260a0850135915080821115610f8057600080fd5b50610f8d85828601610e68565b60a0830152505092915050565b600060208284031215610fac57600080fd5b813567ffffffffffffffff811115610fc357600080fd5b610fcf84828501610ef5565b949350505050565b8215158152604060208201526000610fcf6040830184610dc3565b60008083601f84011261100457600080fd5b50813567ffffffffffffffff81111561101c57600080fd5b60208301915083602082850101111561103457600080fd5b9250929050565b600080600080600080600060a0888a03121561105657600080fd5b873567ffffffffffffffff8082111561106e57600080fd5b61107a8b838c01610ef5565b985060208a0135975060408a0135965060608a013591508082111561109e57600080fd5b6110aa8b838c01610ff2565b909650945060808a01359150808211156110c357600080fd5b506110d08a828b01610ff2565b989b979a50959850939692959293505050565b600080604083850312156110f657600080fd5b823567ffffffffffffffff8082111561110e57600080fd5b61111a86838701610ef5565b9350602085013591508082111561113057600080fd5b5061113d85828601610e68565b9150509250929050565b60008060006060848603121561115c57600080fd5b833567ffffffffffffffff8082111561117457600080fd5b61118087838801610ef5565b945060208601359350604086013591508082111561119d57600080fd5b506111aa86828701610e68565b9150509250925092565b60008060008060008060a087890312156111cd57600080fd5b863567ffffffffffffffff808211156111e557600080fd5b6111f18a838b01610ef5565b97506020890135965060408901359550606089013591508082111561121557600080fd5b6112218a838b01610e68565b9450608089013591508082111561123757600080fd5b5061124489828a01610ff2565b979a9699509497509295939492505050565b6000806000806040858703121561126c57600080fd5b843567ffffffffffffffff8082111561128457600080fd5b61129088838901610ff2565b909650945060208701359150808211156112a957600080fd5b506112b687828801610ff2565b95989497509550505050565b634e487b7160e01b600052601160045260246000fd5b600082198211156112eb576112eb6112c2565b500190565b60008261130d57634e487b7160e01b600052601260045260246000fd5b500490565b634e487b7160e01b600052600160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b83815260008351611356816020850160208801610d93565b8083019050835161136e816020840160208801610d93565b0160200195945050505050565b6000835161138d818460208801610d93565b60609390931b6bffffffffffffffffffffffff19169190920190815260140192915050565b600082516113c4818460208701610d93565b9190910192915050565b60006000198214156113e2576113e26112c2565b5060010190565b84868237600085820160008152600560fb1b81528551611410816001840160208a01610d93565b600b60fa1b60019290910191820152838560028301376000930160020192835250909594505050505056fe616464726573732066726f6d2c6164647265737320746f2c75696e743235362076616c75652c75696e74323536206761732c75696e74323536206e6f6e63652c62797465732064617461a2646970667358221220e1238ea43860767b3a54e623395b86685ef0f2d5a3c676aeb81f3abb09a5669a64736f6c634300080a0033616464726573732066726f6d2c6164647265737320746f2c75696e743235362076616c75652c75696e74323536206761732c75696e74323536206e6f6e63652c62797465732064617461",
}

// ForwarderABI is the input ABI used to generate the binding from.
// Deprecated: Use ForwarderMetaData.ABI instead.
var ForwarderABI = ForwarderMetaData.ABI

// ForwarderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ForwarderMetaData.Bin instead.
var ForwarderBin = ForwarderMetaData.Bin

// DeployForwarder deploys a new Ethereum contract, binding an instance of Forwarder to it.
func DeployForwarder(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Forwarder, error) {
	parsed, err := ForwarderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ForwarderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Forwarder{ForwarderCaller: ForwarderCaller{contract: contract}, ForwarderTransactor: ForwarderTransactor{contract: contract}, ForwarderFilterer: ForwarderFilterer{contract: contract}}, nil
}

// Forwarder is an auto generated Go binding around an Ethereum contract.
type Forwarder struct {
	ForwarderCaller     // Read-only binding to the contract
	ForwarderTransactor // Write-only binding to the contract
	ForwarderFilterer   // Log filterer for contract events
}

// ForwarderCaller is an auto generated read-only Go binding around an Ethereum contract.
type ForwarderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForwarderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ForwarderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
	gsn      *GsnCallOpts
}

func (_Forwarder *ForwarderTransactor) SetGSNOptions(opts *GsnCallOpts) {
	_Forwarder.gsn = opts
}

// ForwarderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ForwarderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForwarderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ForwarderSession struct {
	Contract     *Forwarder        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ForwarderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ForwarderCallerSession struct {
	Contract *ForwarderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ForwarderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ForwarderTransactorSession struct {
	Contract     *ForwarderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ForwarderRaw is an auto generated low-level Go binding around an Ethereum contract.
type ForwarderRaw struct {
	Contract *Forwarder // Generic contract binding to access the raw methods on
}

// ForwarderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ForwarderCallerRaw struct {
	Contract *ForwarderCaller // Generic read-only contract binding to access the raw methods on
}

// ForwarderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ForwarderTransactorRaw struct {
	Contract *ForwarderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewForwarder creates a new instance of Forwarder, bound to a specific deployed contract.
func NewForwarder(address common.Address, backend bind.ContractBackend) (*Forwarder, error) {
	contract, err := bindForwarder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Forwarder{ForwarderCaller: ForwarderCaller{contract: contract}, ForwarderTransactor: ForwarderTransactor{contract: contract}, ForwarderFilterer: ForwarderFilterer{contract: contract}}, nil
}

// NewForwarderCaller creates a new read-only instance of Forwarder, bound to a specific deployed contract.
func NewForwarderCaller(address common.Address, caller bind.ContractCaller) (*ForwarderCaller, error) {
	contract, err := bindForwarder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ForwarderCaller{contract: contract}, nil
}

// NewForwarderTransactor creates a new write-only instance of Forwarder, bound to a specific deployed contract.
func NewForwarderTransactor(address common.Address, transactor bind.ContractTransactor) (*ForwarderTransactor, error) {
	contract, err := bindForwarder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ForwarderTransactor{contract: contract}, nil
}

// NewForwarderFilterer creates a new log filterer instance of Forwarder, bound to a specific deployed contract.
func NewForwarderFilterer(address common.Address, filterer bind.ContractFilterer) (*ForwarderFilterer, error) {
	contract, err := bindForwarder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ForwarderFilterer{contract: contract}, nil
}

// bindForwarder binds a generic wrapper to an already deployed contract.
func bindForwarder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ForwarderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Forwarder *ForwarderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Forwarder.Contract.ForwarderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Forwarder *ForwarderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Forwarder.Contract.ForwarderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Forwarder *ForwarderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Forwarder.Contract.ForwarderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Forwarder *ForwarderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Forwarder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Forwarder *ForwarderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Forwarder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Forwarder *ForwarderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Forwarder.Contract.contract.Transact(opts, method, params...)
}

// GENERICPARAMS is a free data retrieval call binding the contract method 0x066a310c.
//
// Solidity: function GENERIC_PARAMS() view returns(string)
func (_Forwarder *ForwarderCaller) GENERICPARAMS(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "GENERIC_PARAMS")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GENERICPARAMS is a free data retrieval call binding the contract method 0x066a310c.
//
// Solidity: function GENERIC_PARAMS() view returns(string)
func (_Forwarder *ForwarderSession) GENERICPARAMS() (string, error) {
	return _Forwarder.Contract.GENERICPARAMS(&_Forwarder.CallOpts)
}

// GENERICPARAMS is a free data retrieval call binding the contract method 0x066a310c.
//
// Solidity: function GENERIC_PARAMS() view returns(string)
func (_Forwarder *ForwarderCallerSession) GENERICPARAMS() (string, error) {
	return _Forwarder.Contract.GENERICPARAMS(&_Forwarder.CallOpts)
}

// GetEncoded is a free data retrieval call binding the contract method 0xc45c2674.
//
// Solidity: function _getEncoded((address,address,uint256,uint256,uint256,bytes) req, bytes32 requestTypeHash, bytes suffixData) pure returns(bytes)
func (_Forwarder *ForwarderCaller) GetEncoded(opts *bind.CallOpts, req IForwarderForwardRequest, requestTypeHash [32]byte, suffixData []byte) ([]byte, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "_getEncoded", req, requestTypeHash, suffixData)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetEncoded is a free data retrieval call binding the contract method 0xc45c2674.
//
// Solidity: function _getEncoded((address,address,uint256,uint256,uint256,bytes) req, bytes32 requestTypeHash, bytes suffixData) pure returns(bytes)
func (_Forwarder *ForwarderSession) GetEncoded(req IForwarderForwardRequest, requestTypeHash [32]byte, suffixData []byte) ([]byte, error) {
	return _Forwarder.Contract.GetEncoded(&_Forwarder.CallOpts, req, requestTypeHash, suffixData)
}

// GetEncoded is a free data retrieval call binding the contract method 0xc45c2674.
//
// Solidity: function _getEncoded((address,address,uint256,uint256,uint256,bytes) req, bytes32 requestTypeHash, bytes suffixData) pure returns(bytes)
func (_Forwarder *ForwarderCallerSession) GetEncoded(req IForwarderForwardRequest, requestTypeHash [32]byte, suffixData []byte) ([]byte, error) {
	return _Forwarder.Contract.GetEncoded(&_Forwarder.CallOpts, req, requestTypeHash, suffixData)
}

// GetAbiEncodeRequest is a free data retrieval call binding the contract method 0x9d5eaf64.
//
// Solidity: function getAbiEncodeRequest((address,address,uint256,uint256,uint256,bytes) req, bytes reqAbiEncode) pure returns(bytes)
func (_Forwarder *ForwarderCaller) GetAbiEncodeRequest(opts *bind.CallOpts, req IForwarderForwardRequest, reqAbiEncode []byte) ([]byte, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "getAbiEncodeRequest", req, reqAbiEncode)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetAbiEncodeRequest is a free data retrieval call binding the contract method 0x9d5eaf64.
//
// Solidity: function getAbiEncodeRequest((address,address,uint256,uint256,uint256,bytes) req, bytes reqAbiEncode) pure returns(bytes)
func (_Forwarder *ForwarderSession) GetAbiEncodeRequest(req IForwarderForwardRequest, reqAbiEncode []byte) ([]byte, error) {
	return _Forwarder.Contract.GetAbiEncodeRequest(&_Forwarder.CallOpts, req, reqAbiEncode)
}

// GetAbiEncodeRequest is a free data retrieval call binding the contract method 0x9d5eaf64.
//
// Solidity: function getAbiEncodeRequest((address,address,uint256,uint256,uint256,bytes) req, bytes reqAbiEncode) pure returns(bytes)
func (_Forwarder *ForwarderCallerSession) GetAbiEncodeRequest(req IForwarderForwardRequest, reqAbiEncode []byte) ([]byte, error) {
	return _Forwarder.Contract.GetAbiEncodeRequest(&_Forwarder.CallOpts, req, reqAbiEncode)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address from) view returns(uint256)
func (_Forwarder *ForwarderCaller) GetNonce(opts *bind.CallOpts, from common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "getNonce", from)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address from) view returns(uint256)
func (_Forwarder *ForwarderSession) GetNonce(from common.Address) (*big.Int, error) {
	return _Forwarder.Contract.GetNonce(&_Forwarder.CallOpts, from)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address from) view returns(uint256)
func (_Forwarder *ForwarderCallerSession) GetNonce(from common.Address) (*big.Int, error) {
	return _Forwarder.Contract.GetNonce(&_Forwarder.CallOpts, from)
}

// TypeHashes is a free data retrieval call binding the contract method 0x21fe98df.
//
// Solidity: function typeHashes(bytes32 ) view returns(bool)
func (_Forwarder *ForwarderCaller) TypeHashes(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "typeHashes", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TypeHashes is a free data retrieval call binding the contract method 0x21fe98df.
//
// Solidity: function typeHashes(bytes32 ) view returns(bool)
func (_Forwarder *ForwarderSession) TypeHashes(arg0 [32]byte) (bool, error) {
	return _Forwarder.Contract.TypeHashes(&_Forwarder.CallOpts, arg0)
}

// TypeHashes is a free data retrieval call binding the contract method 0x21fe98df.
//
// Solidity: function typeHashes(bytes32 ) view returns(bool)
func (_Forwarder *ForwarderCallerSession) TypeHashes(arg0 [32]byte) (bool, error) {
	return _Forwarder.Contract.TypeHashes(&_Forwarder.CallOpts, arg0)
}

// Verify is a free data retrieval call binding the contract method 0x992d82c3.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,bytes) req, bytes32 domainSeparator, bytes32 requestTypeHash, bytes suffixData, bytes sig) view returns()
func (_Forwarder *ForwarderCaller) Verify(opts *bind.CallOpts, req IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) error {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "verify", req, domainSeparator, requestTypeHash, suffixData, sig)

	if err != nil {
		return err
	}

	return err

}

// Verify is a free data retrieval call binding the contract method 0x992d82c3.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,bytes) req, bytes32 domainSeparator, bytes32 requestTypeHash, bytes suffixData, bytes sig) view returns()
func (_Forwarder *ForwarderSession) Verify(req IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) error {
	return _Forwarder.Contract.Verify(&_Forwarder.CallOpts, req, domainSeparator, requestTypeHash, suffixData, sig)
}

// Verify is a free data retrieval call binding the contract method 0x992d82c3.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,bytes) req, bytes32 domainSeparator, bytes32 requestTypeHash, bytes suffixData, bytes sig) view returns()
func (_Forwarder *ForwarderCallerSession) Verify(req IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) error {
	return _Forwarder.Contract.Verify(&_Forwarder.CallOpts, req, domainSeparator, requestTypeHash, suffixData, sig)
}

// Execute is a paid mutator transaction binding the contract method 0xcdf5735a.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes) req, bytes32 domainSeparator, bytes32 requestTypeHash, bytes suffixData, bytes sig) payable returns(bool success, bytes ret)
func (_Forwarder *ForwarderTransactor) Execute(opts *bind.TransactOpts, req IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) (common.Hash, error) {
	fmt.Printf("DBG: Wrapper run method = %s\n", "Execute")
	if UseGsnFlag && _Forwarder.gsn != nil {
		fmt.Printf("DBG: GsnWrap: Run gsn call\n")
		return GsnExecutor(_Forwarder.gsn, ForwarderMetaData.ABI, "execute", req, domainSeparator, requestTypeHash, suffixData, sig)
	}

	fmt.Printf("DBG: GsnWrap: Direct call\n")
	tx, err := _Forwarder.contract.Transact(opts, "execute", req, domainSeparator, requestTypeHash, suffixData, sig)
	if tx != nil {
		return tx.Hash(), err
	}
	return common.Hash{}, err
}

// Execute is a paid mutator transaction binding the contract method 0xcdf5735a.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes) req, bytes32 domainSeparator, bytes32 requestTypeHash, bytes suffixData, bytes sig) payable returns(bool success, bytes ret)
func (_Forwarder *ForwarderSession) Execute(req IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) (common.Hash, error) {
	return _Forwarder.Contract.Execute(&_Forwarder.TransactOpts, req, domainSeparator, requestTypeHash, suffixData, sig)
}

// Execute is a paid mutator transaction binding the contract method 0xcdf5735a.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes) req, bytes32 domainSeparator, bytes32 requestTypeHash, bytes suffixData, bytes sig) payable returns(bool success, bytes ret)
func (_Forwarder *ForwarderTransactorSession) Execute(req IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) (common.Hash, error) {
	return _Forwarder.Contract.Execute(&_Forwarder.TransactOpts, req, domainSeparator, requestTypeHash, suffixData, sig)
}

// Execute2 is a paid mutator transaction binding the contract method 0xf2662545.
//
// Solidity: function execute2((address,address,uint256,uint256,uint256,bytes) req) payable returns(bool, bytes)
func (_Forwarder *ForwarderTransactor) Execute2(opts *bind.TransactOpts, req IForwarderForwardRequest) (common.Hash, error) {
	fmt.Printf("DBG: Wrapper run method = %s\n", "Execute2")
	if UseGsnFlag && _Forwarder.gsn != nil {
		fmt.Printf("DBG: GsnWrap: Run gsn call\n")
		return GsnExecutor(_Forwarder.gsn, ForwarderMetaData.ABI, "execute2", req)
	}

	fmt.Printf("DBG: GsnWrap: Direct call\n")
	tx, err := _Forwarder.contract.Transact(opts, "execute2", req)
	if tx != nil {
		return tx.Hash(), err
	}
	return common.Hash{}, err
}

// Execute2 is a paid mutator transaction binding the contract method 0xf2662545.
//
// Solidity: function execute2((address,address,uint256,uint256,uint256,bytes) req) payable returns(bool, bytes)
func (_Forwarder *ForwarderSession) Execute2(req IForwarderForwardRequest) (common.Hash, error) {
	return _Forwarder.Contract.Execute2(&_Forwarder.TransactOpts, req)
}

// Execute2 is a paid mutator transaction binding the contract method 0xf2662545.
//
// Solidity: function execute2((address,address,uint256,uint256,uint256,bytes) req) payable returns(bool, bytes)
func (_Forwarder *ForwarderTransactorSession) Execute2(req IForwarderForwardRequest) (common.Hash, error) {
	return _Forwarder.Contract.Execute2(&_Forwarder.TransactOpts, req)
}

// ExecuteAssemblyForwarderRequest is a paid mutator transaction binding the contract method 0x6a66a1e2.
//
// Solidity: function executeAssemblyForwarderRequest((address,address,uint256,uint256,uint256,bytes) req) returns(bool, bytes)
func (_Forwarder *ForwarderTransactor) ExecuteAssemblyForwarderRequest(opts *bind.TransactOpts, req IForwarderForwardRequest) (common.Hash, error) {
	fmt.Printf("DBG: Wrapper run method = %s\n", "ExecuteAssemblyForwarderRequest")
	if UseGsnFlag && _Forwarder.gsn != nil {
		fmt.Printf("DBG: GsnWrap: Run gsn call\n")
		return GsnExecutor(_Forwarder.gsn, ForwarderMetaData.ABI, "executeAssemblyForwarderRequest", req)
	}

	fmt.Printf("DBG: GsnWrap: Direct call\n")
	tx, err := _Forwarder.contract.Transact(opts, "executeAssemblyForwarderRequest", req)
	if tx != nil {
		return tx.Hash(), err
	}
	return common.Hash{}, err
}

// ExecuteAssemblyForwarderRequest is a paid mutator transaction binding the contract method 0x6a66a1e2.
//
// Solidity: function executeAssemblyForwarderRequest((address,address,uint256,uint256,uint256,bytes) req) returns(bool, bytes)
func (_Forwarder *ForwarderSession) ExecuteAssemblyForwarderRequest(req IForwarderForwardRequest) (common.Hash, error) {
	return _Forwarder.Contract.ExecuteAssemblyForwarderRequest(&_Forwarder.TransactOpts, req)
}

// ExecuteAssemblyForwarderRequest is a paid mutator transaction binding the contract method 0x6a66a1e2.
//
// Solidity: function executeAssemblyForwarderRequest((address,address,uint256,uint256,uint256,bytes) req) returns(bool, bytes)
func (_Forwarder *ForwarderTransactorSession) ExecuteAssemblyForwarderRequest(req IForwarderForwardRequest) (common.Hash, error) {
	return _Forwarder.Contract.ExecuteAssemblyForwarderRequest(&_Forwarder.TransactOpts, req)
}

// RegisterRequestType is a paid mutator transaction binding the contract method 0xd9210be5.
//
// Solidity: function registerRequestType(string typeName, string typeSuffix) returns()
func (_Forwarder *ForwarderTransactor) RegisterRequestType(opts *bind.TransactOpts, typeName string, typeSuffix string) (common.Hash, error) {
	fmt.Printf("DBG: Wrapper run method = %s\n", "RegisterRequestType")
	if UseGsnFlag && _Forwarder.gsn != nil {
		fmt.Printf("DBG: GsnWrap: Run gsn call\n")
		return GsnExecutor(_Forwarder.gsn, ForwarderMetaData.ABI, "registerRequestType", typeName, typeSuffix)
	}

	fmt.Printf("DBG: GsnWrap: Direct call\n")
	tx, err := _Forwarder.contract.Transact(opts, "registerRequestType", typeName, typeSuffix)
	if tx != nil {
		return tx.Hash(), err
	}
	return common.Hash{}, err
}

// RegisterRequestType is a paid mutator transaction binding the contract method 0xd9210be5.
//
// Solidity: function registerRequestType(string typeName, string typeSuffix) returns()
func (_Forwarder *ForwarderSession) RegisterRequestType(typeName string, typeSuffix string) (common.Hash, error) {
	return _Forwarder.Contract.RegisterRequestType(&_Forwarder.TransactOpts, typeName, typeSuffix)
}

// RegisterRequestType is a paid mutator transaction binding the contract method 0xd9210be5.
//
// Solidity: function registerRequestType(string typeName, string typeSuffix) returns()
func (_Forwarder *ForwarderTransactorSession) RegisterRequestType(typeName string, typeSuffix string) (common.Hash, error) {
	return _Forwarder.Contract.RegisterRequestType(&_Forwarder.TransactOpts, typeName, typeSuffix)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Forwarder *ForwarderTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Forwarder.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Forwarder *ForwarderSession) Receive() (*types.Transaction, error) {
	return _Forwarder.Contract.Receive(&_Forwarder.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Forwarder *ForwarderTransactorSession) Receive() (*types.Transaction, error) {
	return _Forwarder.Contract.Receive(&_Forwarder.TransactOpts)
}

// ForwarderRequestTypeRegisteredIterator is returned from FilterRequestTypeRegistered and is used to iterate over the raw logs and unpacked data for RequestTypeRegistered events raised by the Forwarder contract.
type ForwarderRequestTypeRegisteredIterator struct {
	Event *ForwarderRequestTypeRegistered // Event containing the contract specifics and raw log

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
func (it *ForwarderRequestTypeRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForwarderRequestTypeRegistered)
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
		it.Event = new(ForwarderRequestTypeRegistered)
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
func (it *ForwarderRequestTypeRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForwarderRequestTypeRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForwarderRequestTypeRegistered represents a RequestTypeRegistered event raised by the Forwarder contract.
type ForwarderRequestTypeRegistered struct {
	TypeHash [32]byte
	TypeStr  string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRequestTypeRegistered is a free log retrieval operation binding the contract event 0x64d6bce64323458c44643c51fe45113efc882082f7b7fd5f09f0d69d2eedb202.
//
// Solidity: event RequestTypeRegistered(bytes32 indexed typeHash, string typeStr)
func (_Forwarder *ForwarderFilterer) FilterRequestTypeRegistered(opts *bind.FilterOpts, typeHash [][32]byte) (*ForwarderRequestTypeRegisteredIterator, error) {

	var typeHashRule []interface{}
	for _, typeHashItem := range typeHash {
		typeHashRule = append(typeHashRule, typeHashItem)
	}

	logs, sub, err := _Forwarder.contract.FilterLogs(opts, "RequestTypeRegistered", typeHashRule)
	if err != nil {
		return nil, err
	}
	return &ForwarderRequestTypeRegisteredIterator{contract: _Forwarder.contract, event: "RequestTypeRegistered", logs: logs, sub: sub}, nil
}

// WatchRequestTypeRegistered is a free log subscription operation binding the contract event 0x64d6bce64323458c44643c51fe45113efc882082f7b7fd5f09f0d69d2eedb202.
//
// Solidity: event RequestTypeRegistered(bytes32 indexed typeHash, string typeStr)
func (_Forwarder *ForwarderFilterer) WatchRequestTypeRegistered(opts *bind.WatchOpts, sink chan<- *ForwarderRequestTypeRegistered, typeHash [][32]byte) (event.Subscription, error) {

	var typeHashRule []interface{}
	for _, typeHashItem := range typeHash {
		typeHashRule = append(typeHashRule, typeHashItem)
	}

	logs, sub, err := _Forwarder.contract.WatchLogs(opts, "RequestTypeRegistered", typeHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForwarderRequestTypeRegistered)
				if err := _Forwarder.contract.UnpackLog(event, "RequestTypeRegistered", log); err != nil {
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

// ParseRequestTypeRegistered is a log parse operation binding the contract event 0x64d6bce64323458c44643c51fe45113efc882082f7b7fd5f09f0d69d2eedb202.
//
// Solidity: event RequestTypeRegistered(bytes32 indexed typeHash, string typeStr)
func (_Forwarder *ForwarderFilterer) ParseRequestTypeRegistered(log types.Log) (*ForwarderRequestTypeRegistered, error) {
	event := new(ForwarderRequestTypeRegistered)
	if err := _Forwarder.contract.UnpackLog(event, "RequestTypeRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
