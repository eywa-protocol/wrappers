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
	From       common.Address
	To         common.Address
	Value      *big.Int
	Gas        *big.Int
	Nonce      *big.Int
	Data       []byte
	ValidUntil *big.Int
}

// ForwarderABI is the input ABI used to generate the binding from.
const ForwarderABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"domainValue\",\"type\":\"bytes\"}],\"name\":\"DomainRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"typeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"typeStr\",\"type\":\"string\"}],\"name\":\"RequestTypeRegistered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GENERIC_PARAMS\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validUntil\",\"type\":\"uint256\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"requestTypeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"suffixData\",\"type\":\"bytes\"}],\"name\":\"_getEncoded\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validUntil\",\"type\":\"uint256\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"requestTypeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"suffixData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"ret\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validUntil\",\"type\":\"uint256\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"execute2\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validUntil\",\"type\":\"uint256\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"executeAssemblyForwarderRequest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validUntil\",\"type\":\"uint256\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"reqAbiEncode\",\"type\":\"bytes\"}],\"name\":\"getAbiEncodeRequest\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"registerDomainSeparator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"typeName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"typeSuffix\",\"type\":\"string\"}],\"name\":\"registerRequestType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"typeHashes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validUntil\",\"type\":\"uint256\"}],\"internalType\":\"structIForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"requestTypeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"suffixData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// ForwarderBin is the compiled bytecode used for deploying new contracts.
var ForwarderBin = "0x60806040523480156200001157600080fd5b5060006040518060800160405280604a81526020016200176a604a9139604051602001620000409190620000fb565b60408051601f1981840301815291905290506200005d8162000064565b5062000174565b8051602080830191909120600081815291829052604091829020805460ff19166001179055905181907f64d6bce64323458c44643c51fe45113efc882082f7b7fd5f09f0d69d2eedb20290620000bc9085906200013f565b60405180910390a25050565b60005b83811015620000e5578181015183820152602001620000cb565b83811115620000f5576000848401525b50505050565b6e08cdee4eec2e4c8a4cae2eacae6e85608b1b8152600082516200012781600f850160208701620000c8565b602960f81b600f939091019283015250601001919050565b602081526000825180602084015262000160816040850160208701620000c8565b601f01601f19169190910160400192915050565b6115e680620001846000396000f3fe6080604052600436106100a05760003560e01c80639c7b4592116100645780639c7b4592146101a9578063a3cc9a4d146101cb578063ad9f99c7146101de578063d9210be5146101fe578063e024dc7f1461021e578063e2b62f2d1461023157600080fd5b8063066a310c146100ac57806321fe98df146100d75780632d0335ab146101175780637296665a1461015b57806390bf16071461018957600080fd5b366100a757005b600080fd5b3480156100b857600080fd5b506100c1610251565b6040516100ce9190610ec5565b60405180910390f35b3480156100e357600080fd5b506101076100f2366004610ed8565b60006020819052908152604090205460ff1681565b60405190151581526020016100ce565b34801561012357600080fd5b5061014d610132366004610f0d565b6001600160a01b031660009081526001602052604090205490565b6040519081526020016100ce565b34801561016757600080fd5b5061017b610176366004611086565b61026d565b6040516100ce9291906110c3565b34801561019557600080fd5b506100c16101a43660046110de565b610357565b3480156101b557600080fd5b506101c96101c436600461118b565b6104e5565b005b61017b6101d9366004611086565b610590565b3480156101ea57600080fd5b506101c96101f93660046111f7565b610693565b34801561020a57600080fd5b506101c961021936600461118b565b61071b565b61017b61022c36600461129f565b610814565b34801561023d57600080fd5b506100c161024c366004611341565b61095a565b6040518060800160405280604a8152602001611567604a913981565b600060608260800151600161028291906113c4565b83516001600160a01b031660009081526001602090815260409182902092909255908401518185015160a0860151805193519293919290919060c060288237600080838387895af16102d8573d6000803e3d6000fd5b50603f87606001516102ea91906113dc565b5a116102f8576102f86113fe565b604080516001808252818301909252600091602082018180368337019050509050600560f81b8160008151811061033157610331611414565b60200101906001600160f81b031916908160001a90535060019890975095505050505050565b60606000836000015184602001518560400151866060015187608001516040516020016103b49594939291906001600160a01b03958616815293909416602084015260408301919091526060820152608081019190915260a00190565b60408051601f1981840301815291905284519091506001600160a01b031661040e5760405162461bcd60e51b81526020600482015260086024820152677265712e66726f6d60c01b60448201526064015b60405180910390fd5b60208401516001600160a01b03166104515760405162461bcd60e51b815260206004820152600660248201526507265712e74360d41b6044820152606401610405565b608084015161048e5760405162461bcd60e51b81526020600482015260096024820152687265712e6e6f6e636560b81b6044820152606401610405565b82805190602001208180519060200120146104de5760405162461bcd60e51b815260206004820152601060248201526f0d0c2e6d0cae640dcdee840dac2e8c6d60831b6044820152606401610405565b9392505050565b83836040516104f592919061142a565b60405180910390207f4bc68689cbe89a4a6333a3ab0a70093874da3e5bfb71e93102027f3f073687d8838360405161052e92919061143a565b60405180910390a260405162461bcd60e51b815260206004820152602960248201527f7265676973746572446f6d61696e536570617261746f722829206e6f7420696d6044820152681c1b195b595b9d195960ba1b6064820152608401610405565b60006060826080015160016105a591906113c4565b6001600085600001516001600160a01b03166001600160a01b031681526020019081526020016000208190555060008084602001516001600160a01b0316856060015186604001518760a001518860000151604051602001610608929190611469565b60408051601f1981840301815290829052610622916114a0565b600060405180830381858888f193505050503d8060008114610660576040519150601f19603f3d011682016040523d82523d6000602084013e610665565b606091505b5091509150603f856060015161067b91906113dc565b5a11610689576106896113fe565b9094909350915050565b61069c87610a00565b61071287878787878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8b018190048102820181019092528981529250899150889081908401838280828437600092019190915250610a6092505050565b50505050505050565b60005b838110156107bf57600085858381811061073a5761073a611414565b909101356001600160f81b031916915050600560fb1b811480159061076d5750602960f81b6001600160f81b0319821614155b6107ac5760405162461bcd60e51b815260206004820152601060248201526f696e76616c696420747970656e616d6560801b6044820152606401610405565b50806107b7816114bc565b91505061071e565b50600084846040518060800160405280604a8152602001611567604a913985856040516020016107f39594939291906114d7565b604051602081830303815290604052905061080d81610b7b565b5050505050565b6000606061082188610a00565b6108648888888888888080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610a6092505050565b61086d88610bdd565b87602001516001600160a01b03165a6040808b015160a08c01518c519251919261089992602001611469565b60408051601f19818403018152908290526108b3916114a0565b600060405180830381858888f193505050503d80600081146108f1576040519150601f19603f3d011682016040523d82523d6000602084013e6108f6565b606091505b5090925090508161094f573d801561091457604051816000823e8181fd5b5060405162461bcd60e51b815260206004820152600f60248201526e756e6b6e6f776e206661696c75726560881b6044820152606401610405565b965096945050505050565b606082846000015185602001518660400151876060015188608001518960a00151805190602001206040516020016109c8969594939291906001600160a01b03968716815294909516602085015260408401929092526060830152608082015260a081019190915260c00190565b60408051601f19818403018152908290526109e892918590602001611529565b60405160208183030381529060405290509392505050565b608081015181516001600160a01b031660009081526001602052604090205414610a5d5760405162461bcd60e51b815260206004820152600e60248201526d0dcdedcc6ca40dad2e6dac2e8c6d60931b6044820152606401610405565b50565b60008381526020819052604090205460ff16610abe5760405162461bcd60e51b815260206004820152601860248201527f696e76616c6964207265717565737420747970656861736800000000000000006044820152606401610405565b600084610acc87868661095a565b8051602091820120604051610af893920161190160f01b81526002810192909252602282015260420190565b60408051601f19818403018152919052805160209091012086519091506001600160a01b0316610b288284610c0a565b6001600160a01b031614610b735760405162461bcd60e51b81526020600482015260126024820152710e6d2cedcc2e8eae4ca40dad2e6dac2e8c6d60731b6044820152606401610405565b505050505050565b8051602080830191909120600081815291829052604091829020805460ff19166001179055905181907f64d6bce64323458c44643c51fe45113efc882082f7b7fd5f09f0d69d2eedb20290610bd1908590610ec5565b60405180910390a25050565b80516001600160a01b03166000908152600160205260408120805491610c02836114bc565b919050555050565b600080600080845160411415610c345750505060208201516040830151606084015160001a610caa565b845160401415610c625750505060408201516020830151906001600160ff1b0381169060ff1c601b01610caa565b60405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610405565b610cb686828585610cc0565b9695505050505050565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0821115610d3d5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610405565b8360ff16601b1480610d5257508360ff16601c145b610da95760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610405565b6040805160008082526020820180845288905260ff871692820192909252606081018590526080810184905260019060a0016020604051602081039080840390855afa158015610dfd573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116610e605760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610405565b95945050505050565b60005b83811015610e84578181015183820152602001610e6c565b83811115610e93576000848401525b50505050565b60008151808452610eb1816020860160208601610e69565b601f01601f19169290920160200192915050565b6020815260006104de6020830184610e99565b600060208284031215610eea57600080fd5b5035919050565b80356001600160a01b0381168114610f0857600080fd5b919050565b600060208284031215610f1f57600080fd5b6104de82610ef1565b634e487b7160e01b600052604160045260246000fd5b60405160e0810167ffffffffffffffff81118282101715610f6157610f61610f28565b60405290565b600082601f830112610f7857600080fd5b813567ffffffffffffffff80821115610f9357610f93610f28565b604051601f8301601f19908116603f01168101908282118183101715610fbb57610fbb610f28565b81604052838152866020858801011115610fd457600080fd5b836020870160208301376000602085830101528094505050505092915050565b600060e0828403121561100657600080fd5b61100e610f3e565b905061101982610ef1565b815261102760208301610ef1565b602082015260408201356040820152606082013560608201526080820135608082015260a082013567ffffffffffffffff81111561106457600080fd5b61107084828501610f67565b60a08301525060c082013560c082015292915050565b60006020828403121561109857600080fd5b813567ffffffffffffffff8111156110af57600080fd5b6110bb84828501610ff4565b949350505050565b82151581526040602082015260006110bb6040830184610e99565b600080604083850312156110f157600080fd5b823567ffffffffffffffff8082111561110957600080fd5b61111586838701610ff4565b9350602085013591508082111561112b57600080fd5b5061113885828601610f67565b9150509250929050565b60008083601f84011261115457600080fd5b50813567ffffffffffffffff81111561116c57600080fd5b60208301915083602082850101111561118457600080fd5b9250929050565b600080600080604085870312156111a157600080fd5b843567ffffffffffffffff808211156111b957600080fd5b6111c588838901611142565b909650945060208701359150808211156111de57600080fd5b506111eb87828801611142565b95989497509550505050565b600080600080600080600060a0888a03121561121257600080fd5b873567ffffffffffffffff8082111561122a57600080fd5b6112368b838c01610ff4565b985060208a0135975060408a0135965060608a013591508082111561125a57600080fd5b6112668b838c01611142565b909650945060808a013591508082111561127f57600080fd5b5061128c8a828b01611142565b989b979a50959850939692959293505050565b60008060008060008060a087890312156112b857600080fd5b863567ffffffffffffffff808211156112d057600080fd5b6112dc8a838b01610ff4565b97506020890135965060408901359550606089013591508082111561130057600080fd5b61130c8a838b01610f67565b9450608089013591508082111561132257600080fd5b5061132f89828a01611142565b979a9699509497509295939492505050565b60008060006060848603121561135657600080fd5b833567ffffffffffffffff8082111561136e57600080fd5b61137a87838801610ff4565b945060208601359350604086013591508082111561139757600080fd5b506113a486828701610f67565b9150509250925092565b634e487b7160e01b600052601160045260246000fd5b600082198211156113d7576113d76113ae565b500190565b6000826113f957634e487b7160e01b600052601260045260246000fd5b500490565b634e487b7160e01b600052600160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b8183823760009101908152919050565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b6000835161147b818460208801610e69565b60609390931b6bffffffffffffffffffffffff19169190920190815260140192915050565b600082516114b2818460208701610e69565b9190910192915050565b60006000198214156114d0576114d06113ae565b5060010190565b84868237600085820160008152600560fb1b815285516114fe816001840160208a01610e69565b600b60fa1b600192909101918201528385600283013760009301600201928352509095945050505050565b83815260008351611541816020850160208801610e69565b80830190508351611559816020840160208801610e69565b016020019594505050505056fe616464726573732066726f6d2c6164647265737320746f2c75696e743235362076616c75652c75696e74323536206761732c75696e74323536206e6f6e63652c62797465732064617461a26469706673582212203c3f305649ce100885663e30cc6111a3509d9933b66d2e05c643bf38d61460f064736f6c634300080a0033616464726573732066726f6d2c6164647265737320746f2c75696e743235362076616c75652c75696e74323536206761732c75696e74323536206e6f6e63652c62797465732064617461"

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

// GetEncoded is a free data retrieval call binding the contract method 0xe2b62f2d.
//
// Solidity: function _getEncoded((address,address,uint256,uint256,uint256,bytes,uint256) req, bytes32 requestTypeHash, bytes suffixData) pure returns(bytes)
func (_Forwarder *ForwarderCaller) GetEncoded(opts *bind.CallOpts, req IForwarderForwardRequest, requestTypeHash [32]byte, suffixData []byte) ([]byte, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "_getEncoded", req, requestTypeHash, suffixData)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetEncoded is a free data retrieval call binding the contract method 0xe2b62f2d.
//
// Solidity: function _getEncoded((address,address,uint256,uint256,uint256,bytes,uint256) req, bytes32 requestTypeHash, bytes suffixData) pure returns(bytes)
func (_Forwarder *ForwarderSession) GetEncoded(req IForwarderForwardRequest, requestTypeHash [32]byte, suffixData []byte) ([]byte, error) {
	return _Forwarder.Contract.GetEncoded(&_Forwarder.CallOpts, req, requestTypeHash, suffixData)
}

// GetEncoded is a free data retrieval call binding the contract method 0xe2b62f2d.
//
// Solidity: function _getEncoded((address,address,uint256,uint256,uint256,bytes,uint256) req, bytes32 requestTypeHash, bytes suffixData) pure returns(bytes)
func (_Forwarder *ForwarderCallerSession) GetEncoded(req IForwarderForwardRequest, requestTypeHash [32]byte, suffixData []byte) ([]byte, error) {
	return _Forwarder.Contract.GetEncoded(&_Forwarder.CallOpts, req, requestTypeHash, suffixData)
}

// GetAbiEncodeRequest is a free data retrieval call binding the contract method 0x90bf1607.
//
// Solidity: function getAbiEncodeRequest((address,address,uint256,uint256,uint256,bytes,uint256) req, bytes reqAbiEncode) pure returns(bytes)
func (_Forwarder *ForwarderCaller) GetAbiEncodeRequest(opts *bind.CallOpts, req IForwarderForwardRequest, reqAbiEncode []byte) ([]byte, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "getAbiEncodeRequest", req, reqAbiEncode)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetAbiEncodeRequest is a free data retrieval call binding the contract method 0x90bf1607.
//
// Solidity: function getAbiEncodeRequest((address,address,uint256,uint256,uint256,bytes,uint256) req, bytes reqAbiEncode) pure returns(bytes)
func (_Forwarder *ForwarderSession) GetAbiEncodeRequest(req IForwarderForwardRequest, reqAbiEncode []byte) ([]byte, error) {
	return _Forwarder.Contract.GetAbiEncodeRequest(&_Forwarder.CallOpts, req, reqAbiEncode)
}

// GetAbiEncodeRequest is a free data retrieval call binding the contract method 0x90bf1607.
//
// Solidity: function getAbiEncodeRequest((address,address,uint256,uint256,uint256,bytes,uint256) req, bytes reqAbiEncode) pure returns(bytes)
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

// Verify is a free data retrieval call binding the contract method 0xad9f99c7.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,bytes,uint256) req, bytes32 domainSeparator, bytes32 requestTypeHash, bytes suffixData, bytes sig) view returns()
func (_Forwarder *ForwarderCaller) Verify(opts *bind.CallOpts, req IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) error {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "verify", req, domainSeparator, requestTypeHash, suffixData, sig)

	if err != nil {
		return err
	}

	return err

}

// Verify is a free data retrieval call binding the contract method 0xad9f99c7.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,bytes,uint256) req, bytes32 domainSeparator, bytes32 requestTypeHash, bytes suffixData, bytes sig) view returns()
func (_Forwarder *ForwarderSession) Verify(req IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) error {
	return _Forwarder.Contract.Verify(&_Forwarder.CallOpts, req, domainSeparator, requestTypeHash, suffixData, sig)
}

// Verify is a free data retrieval call binding the contract method 0xad9f99c7.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,bytes,uint256) req, bytes32 domainSeparator, bytes32 requestTypeHash, bytes suffixData, bytes sig) view returns()
func (_Forwarder *ForwarderCallerSession) Verify(req IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) error {
	return _Forwarder.Contract.Verify(&_Forwarder.CallOpts, req, domainSeparator, requestTypeHash, suffixData, sig)
}

// Execute is a paid mutator transaction binding the contract method 0xe024dc7f.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes,uint256) req, bytes32 domainSeparator, bytes32 requestTypeHash, bytes suffixData, bytes sig) payable returns(bool success, bytes ret)
func (_Forwarder *ForwarderTransactor) Execute(opts *bind.TransactOpts, req IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "execute", req, domainSeparator, requestTypeHash, suffixData, sig)
}

// Execute is a paid mutator transaction binding the contract method 0xe024dc7f.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes,uint256) req, bytes32 domainSeparator, bytes32 requestTypeHash, bytes suffixData, bytes sig) payable returns(bool success, bytes ret)
func (_Forwarder *ForwarderSession) Execute(req IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) (*types.Transaction, error) {
	return _Forwarder.Contract.Execute(&_Forwarder.TransactOpts, req, domainSeparator, requestTypeHash, suffixData, sig)
}

// Execute is a paid mutator transaction binding the contract method 0xe024dc7f.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes,uint256) req, bytes32 domainSeparator, bytes32 requestTypeHash, bytes suffixData, bytes sig) payable returns(bool success, bytes ret)
func (_Forwarder *ForwarderTransactorSession) Execute(req IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) (*types.Transaction, error) {
	return _Forwarder.Contract.Execute(&_Forwarder.TransactOpts, req, domainSeparator, requestTypeHash, suffixData, sig)
}

// Execute2 is a paid mutator transaction binding the contract method 0xa3cc9a4d.
//
// Solidity: function execute2((address,address,uint256,uint256,uint256,bytes,uint256) req) payable returns(bool, bytes)
func (_Forwarder *ForwarderTransactor) Execute2(opts *bind.TransactOpts, req IForwarderForwardRequest) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "execute2", req)
}

// Execute2 is a paid mutator transaction binding the contract method 0xa3cc9a4d.
//
// Solidity: function execute2((address,address,uint256,uint256,uint256,bytes,uint256) req) payable returns(bool, bytes)
func (_Forwarder *ForwarderSession) Execute2(req IForwarderForwardRequest) (*types.Transaction, error) {
	return _Forwarder.Contract.Execute2(&_Forwarder.TransactOpts, req)
}

// Execute2 is a paid mutator transaction binding the contract method 0xa3cc9a4d.
//
// Solidity: function execute2((address,address,uint256,uint256,uint256,bytes,uint256) req) payable returns(bool, bytes)
func (_Forwarder *ForwarderTransactorSession) Execute2(req IForwarderForwardRequest) (*types.Transaction, error) {
	return _Forwarder.Contract.Execute2(&_Forwarder.TransactOpts, req)
}

// ExecuteAssemblyForwarderRequest is a paid mutator transaction binding the contract method 0x7296665a.
//
// Solidity: function executeAssemblyForwarderRequest((address,address,uint256,uint256,uint256,bytes,uint256) req) returns(bool, bytes)
func (_Forwarder *ForwarderTransactor) ExecuteAssemblyForwarderRequest(opts *bind.TransactOpts, req IForwarderForwardRequest) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "executeAssemblyForwarderRequest", req)
}

// ExecuteAssemblyForwarderRequest is a paid mutator transaction binding the contract method 0x7296665a.
//
// Solidity: function executeAssemblyForwarderRequest((address,address,uint256,uint256,uint256,bytes,uint256) req) returns(bool, bytes)
func (_Forwarder *ForwarderSession) ExecuteAssemblyForwarderRequest(req IForwarderForwardRequest) (*types.Transaction, error) {
	return _Forwarder.Contract.ExecuteAssemblyForwarderRequest(&_Forwarder.TransactOpts, req)
}

// ExecuteAssemblyForwarderRequest is a paid mutator transaction binding the contract method 0x7296665a.
//
// Solidity: function executeAssemblyForwarderRequest((address,address,uint256,uint256,uint256,bytes,uint256) req) returns(bool, bytes)
func (_Forwarder *ForwarderTransactorSession) ExecuteAssemblyForwarderRequest(req IForwarderForwardRequest) (*types.Transaction, error) {
	return _Forwarder.Contract.ExecuteAssemblyForwarderRequest(&_Forwarder.TransactOpts, req)
}

// RegisterDomainSeparator is a paid mutator transaction binding the contract method 0x9c7b4592.
//
// Solidity: function registerDomainSeparator(string name, string version) returns()
func (_Forwarder *ForwarderTransactor) RegisterDomainSeparator(opts *bind.TransactOpts, name string, version string) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "registerDomainSeparator", name, version)
}

// RegisterDomainSeparator is a paid mutator transaction binding the contract method 0x9c7b4592.
//
// Solidity: function registerDomainSeparator(string name, string version) returns()
func (_Forwarder *ForwarderSession) RegisterDomainSeparator(name string, version string) (*types.Transaction, error) {
	return _Forwarder.Contract.RegisterDomainSeparator(&_Forwarder.TransactOpts, name, version)
}

// RegisterDomainSeparator is a paid mutator transaction binding the contract method 0x9c7b4592.
//
// Solidity: function registerDomainSeparator(string name, string version) returns()
func (_Forwarder *ForwarderTransactorSession) RegisterDomainSeparator(name string, version string) (*types.Transaction, error) {
	return _Forwarder.Contract.RegisterDomainSeparator(&_Forwarder.TransactOpts, name, version)
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

// ForwarderDomainRegisteredIterator is returned from FilterDomainRegistered and is used to iterate over the raw logs and unpacked data for DomainRegistered events raised by the Forwarder contract.
type ForwarderDomainRegisteredIterator struct {
	Event *ForwarderDomainRegistered // Event containing the contract specifics and raw log

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
func (it *ForwarderDomainRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForwarderDomainRegistered)
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
		it.Event = new(ForwarderDomainRegistered)
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
func (it *ForwarderDomainRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForwarderDomainRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForwarderDomainRegistered represents a DomainRegistered event raised by the Forwarder contract.
type ForwarderDomainRegistered struct {
	DomainSeparator [32]byte
	DomainValue     []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDomainRegistered is a free log retrieval operation binding the contract event 0x4bc68689cbe89a4a6333a3ab0a70093874da3e5bfb71e93102027f3f073687d8.
//
// Solidity: event DomainRegistered(bytes32 indexed domainSeparator, bytes domainValue)
func (_Forwarder *ForwarderFilterer) FilterDomainRegistered(opts *bind.FilterOpts, domainSeparator [][32]byte) (*ForwarderDomainRegisteredIterator, error) {

	var domainSeparatorRule []interface{}
	for _, domainSeparatorItem := range domainSeparator {
		domainSeparatorRule = append(domainSeparatorRule, domainSeparatorItem)
	}

	logs, sub, err := _Forwarder.contract.FilterLogs(opts, "DomainRegistered", domainSeparatorRule)
	if err != nil {
		return nil, err
	}
	return &ForwarderDomainRegisteredIterator{contract: _Forwarder.contract, event: "DomainRegistered", logs: logs, sub: sub}, nil
}

// WatchDomainRegistered is a free log subscription operation binding the contract event 0x4bc68689cbe89a4a6333a3ab0a70093874da3e5bfb71e93102027f3f073687d8.
//
// Solidity: event DomainRegistered(bytes32 indexed domainSeparator, bytes domainValue)
func (_Forwarder *ForwarderFilterer) WatchDomainRegistered(opts *bind.WatchOpts, sink chan<- *ForwarderDomainRegistered, domainSeparator [][32]byte) (event.Subscription, error) {

	var domainSeparatorRule []interface{}
	for _, domainSeparatorItem := range domainSeparator {
		domainSeparatorRule = append(domainSeparatorRule, domainSeparatorItem)
	}

	logs, sub, err := _Forwarder.contract.WatchLogs(opts, "DomainRegistered", domainSeparatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForwarderDomainRegistered)
				if err := _Forwarder.contract.UnpackLog(event, "DomainRegistered", log); err != nil {
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

// ParseDomainRegistered is a log parse operation binding the contract event 0x4bc68689cbe89a4a6333a3ab0a70093874da3e5bfb71e93102027f3f073687d8.
//
// Solidity: event DomainRegistered(bytes32 indexed domainSeparator, bytes domainValue)
func (_Forwarder *ForwarderFilterer) ParseDomainRegistered(log types.Log) (*ForwarderDomainRegistered, error) {
	event := new(ForwarderDomainRegistered)
	if err := _Forwarder.contract.UnpackLog(event, "DomainRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
