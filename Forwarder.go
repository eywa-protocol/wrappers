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

// IForwarderForwardRequest is an auto generated low-level Go binding around an user-defined struct.
type IForwarderForwardRequest struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Gas   *big.Int
	Nonce *big.Int
	Data  []byte
}

// ForwarderABI is the input ABI used to generate the binding from.
const ForwarderABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"typeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"typeStr\",\"type\":\"string\"}],\"name\":\"RequestTypeRegistered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GENERIC_PARAMS\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"requestTypeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"suffixData\",\"type\":\"bytes\"}],\"name\":\"_getEncoded\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"requestTypeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"suffixData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"ret\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"execute2\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"executeAssemblyForwarderRequest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"reqAbiEncode\",\"type\":\"bytes\"}],\"name\":\"getAbiEncodeRequest\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"typeName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"typeSuffix\",\"type\":\"string\"}],\"name\":\"registerRequestType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"typeHashes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"requestTypeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"suffixData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// ForwarderBin is the compiled bytecode used for deploying new contracts.
var ForwarderBin = "0x60806040523480156200001157600080fd5b5060006040518060800160405280604a8152602001620016b7604a9139604051602001620000409190620000c8565b60408051601f1981840301815291905290506200005d8162000064565b5062000174565b8051602080830191909120600081815291829052604091829020805460ff19166001179055905181907f64d6bce64323458c44643c51fe45113efc882082f7b7fd5f09f0d69d2eedb20290620000bc9085906200010c565b60405180910390a25050565b60006e08cdee4eec2e4c8a4cae2eacae6e85608b1b82528251620000f481600f85016020870162000141565b602960f81b600f939091019283015250601001919050565b60006020825282518060208401526200012d81604085016020870162000141565b601f01601f19169190910160400192915050565b60005b838110156200015e57818101518382015260200162000144565b838111156200016e576000848401525b50505050565b61153380620001846000396000f3fe6080604052600436106100955760003560e01c80639d5eaf64116100595780639d5eaf6414610176578063c45c267414610196578063cdf5735a146101b6578063d9210be5146101c9578063f2662545146101e95761009c565b8063066a310c146100a157806321fe98df146100cc5780632d0335ab146100f95780636a66a1e214610126578063992d82c3146101545761009c565b3661009c57005b600080fd5b3480156100ad57600080fd5b506100b66101fc565b6040516100c39190611189565b60405180910390f35b3480156100d857600080fd5b506100ec6100e7366004610cf2565b610218565b6040516100c39190611145565b34801561010557600080fd5b50610119610114366004610cd8565b61022d565b6040516100c391906113fb565b34801561013257600080fd5b50610146610141366004610d73565b61024c565b6040516100c3929190611150565b34801561016057600080fd5b5061017461016f366004610dae565b610352565b005b34801561018257600080fd5b506100b6610191366004610f5a565b6103da565b3480156101a257600080fd5b506100b66101b1366004610ef0565b6104cd565b6101466101c4366004610e52565b610542565b3480156101d557600080fd5b506101746101e4366004610d0a565b610666565b6101466101f7366004610d73565b61074a565b6040518060800160405280604a81526020016114b4604a913981565b60006020819052908152604090205460ff1681565b6001600160a01b0381166000908152600160205260409020545b919050565b60006060826080015160016102619190611404565b83516001600160a01b031660009081526001602090815260409182902092909255908401518185015160a0860151805193519293919290919060c060288237600080838387895af16102b7573d6000803e3d6000fd5b50603f87606001516102c9919061141c565b5a116102e557634e487b7160e01b600052600160045260246000fd5b604080516001808252818301909252600091602082018180368337019050509050600560f81b8160008151811061032c57634e487b7160e01b600052603260045260246000fd5b60200101906001600160f81b031916908160001a90535060019650945050505050915091565b61035b8761085b565b6103d187878787878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8b01819004810282018101909252898152925089915088908190840183828082843760009201919091525061089a92505050565b50505050505050565b606060008360000151846020015185604001518660600151876080015160405160200161040b9594939291906110de565b60408051601f1981840301815291905284519091506001600160a01b031661044e5760405162461bcd60e51b8152600401610445906112c9565b60405180910390fd5b60208401516001600160a01b03166104785760405162461bcd60e51b8152600401610445906111d3565b60808401516104995760405162461bcd60e51b8152600401610445906112a6565b82805190602001208180519060200120146104c65760405162461bcd60e51b81526004016104459061127c565b9392505050565b606082846000015185602001518660400151876060015188608001518960a001518051906020012060405160200161050a9695949392919061110f565b60408051601f198184030181529082905261052a92918590602001610fe7565b60405160208183030381529060405290509392505050565b6000606061054f8861085b565b6105928888888888888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061089a92505050565b61059b8861094a565b87602001516001600160a01b03165a6040808b015160a08c01518c51925191926105c792602001611040565b60408051601f19818403018152908290526105e191611024565b600060405180830381858888f193505050503d806000811461061f576040519150601f19603f3d011682016040523d82523d6000602084013e610624565b606091505b5090925090508161065b573d801561064257604051816000823e8181fd5b5060405162461bcd60e51b8152600401610445906113d2565b965096945050505050565b60005b838110156106f557600085858381811061069357634e487b7160e01b600052603260045260246000fd5b909101356001600160f81b031916915050600560fb1b81148015906106c65750602960f81b6001600160f81b0319821614155b6106e25760405162461bcd60e51b81526004016104459061122a565b50806106ed8161146c565b915050610669565b50600084846040518060800160405280604a81526020016114b4604a91398585604051602001610729959493929190611077565b604051602081830303815290604052905061074381610977565b5050505050565b600060608260800151600161075f9190611404565b6001600085600001516001600160a01b03166001600160a01b031681526020019081526020016000208190555060008084602001516001600160a01b0316856060015186604001518760a0015188600001516040516020016107c2929190611040565b60408051601f19818403018152908290526107dc91611024565b600060405180830381858888f193505050503d806000811461081a576040519150601f19603f3d011682016040523d82523d6000602084013e61081f565b606091505b5091509150603f8560600151610835919061141c565b5a1161085157634e487b7160e01b600052600160045260246000fd5b9092509050915091565b608081015181516001600160a01b0316600090815260016020526040902054146108975760405162461bcd60e51b815260040161044590611254565b50565b60008381526020819052604090205460ff166108c85760405162461bcd60e51b81526004016104459061136f565b6000846108d68786866104cd565b80516020918201206040516108ec9392016110c3565b60408051601f19818403018152919052805160209091012086519091506001600160a01b031661091c82846109d9565b6001600160a01b0316146109425760405162461bcd60e51b8152600401610445906113a6565b505050505050565b80516001600160a01b0316600090815260016020526040812080549161096f8361146c565b919050555050565b8051602080830191909120600081815291829052604091829020805460ff19166001179055905181907f64d6bce64323458c44643c51fe45113efc882082f7b7fd5f09f0d69d2eedb202906109cd908590611189565b60405180910390a25050565b600080600080845160411415610a035750505060208201516040830151606084015160001a610a49565b845160401415610a315750505060408201516020830151906001600160ff1b0381169060ff1c601b01610a49565b60405162461bcd60e51b8152600401610445906111f3565b610a5586828585610a5f565b9695505050505050565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0821115610aa15760405162461bcd60e51b8152600401610445906112eb565b8360ff16601b1480610ab657508360ff16601c145b610ad25760405162461bcd60e51b81526004016104459061132d565b600060018686868660405160008152602001604052604051610af7949392919061116b565b6020604051602081039080840390855afa158015610b19573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116610b4c5760405162461bcd60e51b81526004016104459061119c565b95945050505050565b80356001600160a01b038116811461024757600080fd5b60008083601f840112610b7d578081fd5b50813567ffffffffffffffff811115610b94578182fd5b602083019150836020828501011115610bac57600080fd5b9250929050565b600082601f830112610bc3578081fd5b813567ffffffffffffffff80821115610bde57610bde61149d565b604051601f8301601f191681016020018281118282101715610c0257610c0261149d565b604052828152848301602001861015610c19578384fd5b82602086016020830137918201602001929092529392505050565b600060c08284031215610c45578081fd5b60405160c0810167ffffffffffffffff8282108183111715610c6957610c6961149d565b81604052829350610c7985610b55565b8352610c8760208601610b55565b602084015260408501356040840152606085013560608401526080850135608084015260a0850135915080821115610cbe57600080fd5b50610ccb85828601610bb3565b60a0830152505092915050565b600060208284031215610ce9578081fd5b6104c682610b55565b600060208284031215610d03578081fd5b5035919050565b60008060008060408587031215610d1f578283fd5b843567ffffffffffffffff80821115610d36578485fd5b610d4288838901610b6c565b90965094506020870135915080821115610d5a578384fd5b50610d6787828801610b6c565b95989497509550505050565b600060208284031215610d84578081fd5b813567ffffffffffffffff811115610d9a578182fd5b610da684828501610c34565b949350505050565b600080600080600080600060a0888a031215610dc8578283fd5b873567ffffffffffffffff80821115610ddf578485fd5b610deb8b838c01610c34565b985060208a0135975060408a0135965060608a0135915080821115610e0e578485fd5b610e1a8b838c01610b6c565b909650945060808a0135915080821115610e32578384fd5b50610e3f8a828b01610b6c565b989b979a50959850939692959293505050565b60008060008060008060a08789031215610e6a578182fd5b863567ffffffffffffffff80821115610e81578384fd5b610e8d8a838b01610c34565b975060208901359650604089013595506060890135915080821115610eb0578384fd5b610ebc8a838b01610bb3565b94506080890135915080821115610ed1578384fd5b50610ede89828a01610b6c565b979a9699509497509295939492505050565b600080600060608486031215610f04578283fd5b833567ffffffffffffffff80821115610f1b578485fd5b610f2787838801610c34565b9450602086013593506040860135915080821115610f43578283fd5b50610f5086828701610bb3565b9150509250925092565b60008060408385031215610f6c578182fd5b823567ffffffffffffffff80821115610f83578384fd5b610f8f86838701610c34565b93506020850135915080821115610fa4578283fd5b50610fb185828601610bb3565b9150509250929050565b60008151808452610fd381602086016020860161143c565b601f01601f19169290920160200192915050565b60008482528351610fff81602085016020880161143c565b8083019050835161101781602084016020880161143c565b0160200195945050505050565b6000825161103681846020870161143c565b9190910192915050565b6000835161105281846020880161143c565b60609390931b6bffffffffffffffffffffffff19169190920190815260140192915050565b600085878337600560fb1b828701908152855161109b816001840160208a0161143c565b600b60fa1b600192909101918201528385600283013790920160020191825250949350505050565b61190160f01b81526002810192909252602282015260420190565b6001600160a01b03958616815293909416602084015260408301919091526060820152608081019190915260a00190565b6001600160a01b03968716815294909516602085015260408401929092526060830152608082015260a081019190915260c00190565b901515815260200190565b6000831515825260406020830152610da66040830184610fbb565b93845260ff9290921660208401526040830152606082015260800190565b6000602082526104c66020830184610fbb565b60208082526018908201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604082015260600190565b60208082526006908201526507265712e74360d41b604082015260600190565b6020808252601f908201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604082015260600190565b60208082526010908201526f696e76616c696420747970656e616d6560801b604082015260600190565b6020808252600e908201526d0dcdedcc6ca40dad2e6dac2e8c6d60931b604082015260600190565b60208082526010908201526f0d0c2e6d0cae640dcdee840dac2e8c6d60831b604082015260600190565b6020808252600990820152687265712e6e6f6e636560b81b604082015260600190565b6020808252600890820152677265712e66726f6d60c01b604082015260600190565b60208082526022908201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604082015261756560f01b606082015260800190565b60208082526022908201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604082015261756560f01b606082015260800190565b60208082526018908201527f696e76616c696420726571756573742074797065686173680000000000000000604082015260600190565b6020808252601290820152710e6d2cedcc2e8eae4ca40dad2e6dac2e8c6d60731b604082015260600190565b6020808252600f908201526e756e6b6e6f776e206661696c75726560881b604082015260600190565b90815260200190565b6000821982111561141757611417611487565b500190565b60008261143757634e487b7160e01b81526012600452602481fd5b500490565b60005b8381101561145757818101518382015260200161143f565b83811115611466576000848401525b50505050565b600060001982141561148057611480611487565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fdfe616464726573732066726f6d2c6164647265737320746f2c75696e743235362076616c75652c75696e74323536206761732c75696e74323536206e6f6e63652c62797465732064617461a2646970667358221220e7c9287db504aacfc6faef395807f310d0b2f4c39226ad3d104e4105f8032a9c64736f6c63430008000033616464726573732066726f6d2c6164647265737320746f2c75696e743235362076616c75652c75696e74323536206761732c75696e74323536206e6f6e63652c62797465732064617461"

// DeployForwarder deploys a new Ethereum contract, binding an instance of Forwarder to it.
func DeployForwarder(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Forwarder, error) {
	parsed, err := abi.JSON(strings.NewReader(ForwarderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ForwarderBin), backend)
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
// Solidity: function getAbiEncodeRequest((address,address,uint256,uint256,uint256,bytes) req, bytes reqAbiEncode) view returns(bytes)
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
// Solidity: function getAbiEncodeRequest((address,address,uint256,uint256,uint256,bytes) req, bytes reqAbiEncode) view returns(bytes)
func (_Forwarder *ForwarderSession) GetAbiEncodeRequest(req IForwarderForwardRequest, reqAbiEncode []byte) ([]byte, error) {
	return _Forwarder.Contract.GetAbiEncodeRequest(&_Forwarder.CallOpts, req, reqAbiEncode)
}

// GetAbiEncodeRequest is a free data retrieval call binding the contract method 0x9d5eaf64.
//
// Solidity: function getAbiEncodeRequest((address,address,uint256,uint256,uint256,bytes) req, bytes reqAbiEncode) view returns(bytes)
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
func (_Forwarder *ForwarderTransactor) Execute(opts *bind.TransactOpts, req IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "execute", req, domainSeparator, requestTypeHash, suffixData, sig)
}

// Execute is a paid mutator transaction binding the contract method 0xcdf5735a.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes) req, bytes32 domainSeparator, bytes32 requestTypeHash, bytes suffixData, bytes sig) payable returns(bool success, bytes ret)
func (_Forwarder *ForwarderSession) Execute(req IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) (*types.Transaction, error) {
	return _Forwarder.Contract.Execute(&_Forwarder.TransactOpts, req, domainSeparator, requestTypeHash, suffixData, sig)
}

// Execute is a paid mutator transaction binding the contract method 0xcdf5735a.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes) req, bytes32 domainSeparator, bytes32 requestTypeHash, bytes suffixData, bytes sig) payable returns(bool success, bytes ret)
func (_Forwarder *ForwarderTransactorSession) Execute(req IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) (*types.Transaction, error) {
	return _Forwarder.Contract.Execute(&_Forwarder.TransactOpts, req, domainSeparator, requestTypeHash, suffixData, sig)
}

// Execute2 is a paid mutator transaction binding the contract method 0xf2662545.
//
// Solidity: function execute2((address,address,uint256,uint256,uint256,bytes) req) payable returns(bool, bytes)
func (_Forwarder *ForwarderTransactor) Execute2(opts *bind.TransactOpts, req IForwarderForwardRequest) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "execute2", req)
}

// Execute2 is a paid mutator transaction binding the contract method 0xf2662545.
//
// Solidity: function execute2((address,address,uint256,uint256,uint256,bytes) req) payable returns(bool, bytes)
func (_Forwarder *ForwarderSession) Execute2(req IForwarderForwardRequest) (*types.Transaction, error) {
	return _Forwarder.Contract.Execute2(&_Forwarder.TransactOpts, req)
}

// Execute2 is a paid mutator transaction binding the contract method 0xf2662545.
//
// Solidity: function execute2((address,address,uint256,uint256,uint256,bytes) req) payable returns(bool, bytes)
func (_Forwarder *ForwarderTransactorSession) Execute2(req IForwarderForwardRequest) (*types.Transaction, error) {
	return _Forwarder.Contract.Execute2(&_Forwarder.TransactOpts, req)
}

// ExecuteAssemblyForwarderRequest is a paid mutator transaction binding the contract method 0x6a66a1e2.
//
// Solidity: function executeAssemblyForwarderRequest((address,address,uint256,uint256,uint256,bytes) req) returns(bool, bytes)
func (_Forwarder *ForwarderTransactor) ExecuteAssemblyForwarderRequest(opts *bind.TransactOpts, req IForwarderForwardRequest) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "executeAssemblyForwarderRequest", req)
}

// ExecuteAssemblyForwarderRequest is a paid mutator transaction binding the contract method 0x6a66a1e2.
//
// Solidity: function executeAssemblyForwarderRequest((address,address,uint256,uint256,uint256,bytes) req) returns(bool, bytes)
func (_Forwarder *ForwarderSession) ExecuteAssemblyForwarderRequest(req IForwarderForwardRequest) (*types.Transaction, error) {
	return _Forwarder.Contract.ExecuteAssemblyForwarderRequest(&_Forwarder.TransactOpts, req)
}

// ExecuteAssemblyForwarderRequest is a paid mutator transaction binding the contract method 0x6a66a1e2.
//
// Solidity: function executeAssemblyForwarderRequest((address,address,uint256,uint256,uint256,bytes) req) returns(bool, bytes)
func (_Forwarder *ForwarderTransactorSession) ExecuteAssemblyForwarderRequest(req IForwarderForwardRequest) (*types.Transaction, error) {
	return _Forwarder.Contract.ExecuteAssemblyForwarderRequest(&_Forwarder.TransactOpts, req)
}

// RegisterRequestType is a paid mutator transaction binding the contract method 0xd9210be5.
//
// Solidity: function registerRequestType(string typeName, string typeSuffix) returns()
func (_Forwarder *ForwarderTransactor) RegisterRequestType(opts *bind.TransactOpts, typeName string, typeSuffix string) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "registerRequestType", typeName, typeSuffix)
}

// RegisterRequestType is a paid mutator transaction binding the contract method 0xd9210be5.
//
// Solidity: function registerRequestType(string typeName, string typeSuffix) returns()
func (_Forwarder *ForwarderSession) RegisterRequestType(typeName string, typeSuffix string) (*types.Transaction, error) {
	return _Forwarder.Contract.RegisterRequestType(&_Forwarder.TransactOpts, typeName, typeSuffix)
}

// RegisterRequestType is a paid mutator transaction binding the contract method 0xd9210be5.
//
// Solidity: function registerRequestType(string typeName, string typeSuffix) returns()
func (_Forwarder *ForwarderTransactorSession) RegisterRequestType(typeName string, typeSuffix string) (*types.Transaction, error) {
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
