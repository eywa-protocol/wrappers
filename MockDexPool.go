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

// MockDexPoolMetaData contains all meta data concerning the MockDexPool contract.

var MockDexPoolMetaData = &bind.MetaData{

	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"RequestReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"reqId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"RequestReceivedV2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"reqId\",\"type\":\"bytes32\"}],\"name\":\"RequestSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"testData_\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"secondPartPool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oppBridge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"TestEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SOLANA_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOLANA_ED25519_SIG\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOLANA_INSTRUCTIONS\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOLANA_RENT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOLANA_SYSTEM_PROGRAM\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOLANA_TOKEN_PROGRAM\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"secondPartPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"calcRequestId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clearStats\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"doubleRequestError\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"doubleRequestIds\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"doubles\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_testData\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_reqId\",\"type\":\"bytes32\"}],\"name\":\"receiveRequestTest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"testData_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"secondPartPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"sendRequestTestV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"testData_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"secondPartPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"sendRequestTestV2Unsafe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"testData_\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"secondPartPool_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppBridge_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId_\",\"type\":\"uint256\"}],\"name\":\"sendTest2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"testStubPID_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"solBridgePID_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataAcc_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bridgePDASigner_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"testData_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"sendTestRequestToSolana\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"programId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"pubkey\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isSigner\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isWritable\",\"type\":\"bool\"}],\"internalType\":\"structSolanaSerialize.SolanaAccountMeta[]\",\"name\":\"accounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structSolanaSerialize.SolanaStandaloneInstruction\",\"name\":\"ix\",\"type\":\"tuple\"}],\"name\":\"serializeSolanaStandaloneInstruction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"sigHash\",\"outputs\":[{\"internalType\":\"bytes8\",\"name\":\"\",\"type\":\"bytes8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"thisChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c0604052601060809081526f676c6f62616c3a7365745f76616c756560801b60a0526200002d90620000a5565b600080546001600160401b03191660c09290921c91909117815560018190556005553480156200005c57600080fd5b5060405162001809380380620018098339810160408190526200007f9162000102565b600280546001600160a01b0319166001600160a01b03929092169190911790556200018c565b6000600282604051620000b9919062000134565b602060405180830381855afa158015620000d7573d6000803e3d6000fd5b5050506040513d601f19601f82011682018060405250810190620000fc919062000172565b92915050565b6000602082840312156200011557600080fd5b81516001600160a01b03811681146200012d57600080fd5b9392505050565b6000825160005b818110156200015757602081860181015185830152016200013b565b8181111562000167576000828501525b509190910192915050565b6000602082840312156200018557600080fd5b5051919050565b61166d806200019c6000396000f3fe608060405234801561001057600080fd5b506004361061014d5760003560e01c8063842ce88d116100c3578063e78cea921161007c578063e78cea92146102f4578063e9096b571461031f578063edff2c8214610327578063f54ea76b1461034e578063f9ee520e14610361578063faad85c81461037457600080fd5b8063842ce88d146102735780638aea61dc1461029a5780639d866985146102a3578063ac9be189146102c3578063d5854d70146102d8578063e3d29c5d146102eb57600080fd5b80632ab8c8b0116101155780632ab8c8b0146101bd5780632b1106e3146101dd57806352dd22ee14610204578063670490181461020c578063671c288414610238578063724fed5a1461026057600080fd5b8063016cbd5114610152578063054aabee1461016e578063067bd07a146101955780630a82d745146101a057806319d5c38e146101a8575b600080fd5b61015b60015481565b6040519081526020015b60405180910390f35b61015b7f037d46d67c93fbbe12f9428f838d40ff0570744927f48a64fcca70448000000081565b61015b631de44e3d81565b60045461015b565b6101bb6101b6366004610f4e565b610387565b005b6101d06101cb3660046110a3565b6104f8565b604051610165919061123f565b61015b7f06a7d517192c5c51218cc94c3d4af17f58daee089ba1fd44e3dbd98a0000000081565b61015b600081565b61021f61021a366004611252565b6106e9565b6040516001600160c01b03199091168152602001610165565b61024b6102463660046112a3565b610741565b60408051928352602083019190915201610165565b61015b61026e3660046112df565b6107e3565b61015b7f06a7d517187bd16635dad40455fdc2c0c124c68f215675a5dbbacb5f0800000081565b61015b60055481565b61015b6102b13660046112df565b60036020526000908152604090205481565b6102cb610804565b60405161016591906112f8565b6101bb6102e636600461133c565b61085c565b61015b60065481565b600254610307906001600160a01b031681565b6040516001600160a01b039091168152602001610165565b6101bb61089f565b61015b7f06ddf6e1d765a193d9cbe146ceeb79ac1cb485ed5f5b37913a8cf5857eff00a981565b6101bb61035c3660046113a2565b6108b2565b6101bb61036f3660046113e5565b610b74565b6101bb61038236600461140b565b610d66565b6001600160a01b0385166103d05760405162461bcd60e51b815260206004820152600b60248201526a424144204144445245535360a81b60448201526064015b60405180910390fd5b6000604051806060016040528060238152602001611615602391398051602090910120604051602481018990526044810185905260640160408051601f198184030181529181526020820180516001600160e01b03166001600160e01b031990941693909317909252600254915163329ef45f60e01b81529092506001600160a01b039091169063329ef45f906104779084908a908a908a908a9033908b9060040161142d565b6020604051808303816000875af1158015610496573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104ba919061147d565b506040518381527f1131472297a800fee664d1d89cfa8f7676ff07189ecc53f80bbb5f4969099db8906020015b60405180910390a150505050505050565b60208082015151604051606092600883811c62ff00ff1663ff00ff009490911b9390931692909217601081811c91901b179160009161054c9184910160e09190911b6001600160e01b031916815260040190565b6040516020818303038152906040529050606060005b85602001515181101561064a57856020015181815181106105855761058561149a565b602002602001015160000151866020015182815181106105a7576105a761149a565b602002602001015160200151876020015183815181106105c9576105c961149a565b6020026020010151604001516040516020016106029392919092835290151560f890811b6020840152901515901b602182015260220190565b604051602081830303815290604052915082826040516020016106269291906114b0565b60405160208183030381529060405292508080610642906114df565b915050610562565b50845160405161065e918491602001611508565b60408051601f198184030181528282529087015151600881811c62ff00ff1663ff00ff009290911b9190911617601081811c91901b17945092506106a8908390859060200161152a565b60405160208183030381529060405291508185604001516040516020016106d09291906114b0565b60408051601f1981840301815291905295945050505050565b60006002826040516106fb919061155c565b602060405180830381855afa158015610718573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061073b9190611578565b92915050565b600254604051632d0335ab60e01b8152336004820152600091829182916001600160a01b031690632d0335ab90602401602060405180830381865afa15801561078e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107b29190611578565b6006549091506000906107d6906001600160a01b038089169188918b163387610ea9565b9791965090945050505050565b600481815481106107f357600080fd5b600091825260209091200154905081565b6060600480548060200260200160405190810160405280929190818152602001828054801561085257602002820191906000526020600020905b81548152602001906001019080831161083e575b5050505050905090565b7f6732d52ba34d766a4a84b7fe8ccca6e4a4b89895ca1c38c9cc5b88ea59d97dd4848484846040516108919493929190611591565b60405180910390a150505050565b6108ab60046000610ef8565b6000600555565b631de44e3d81146108f95760405162461bcd60e51b81526020600482015260116024820152701a5b98dbdc9c9958dd0818da185a5b9259607a1b60448201526064016103c7565b600254604051632d0335ab60e01b81523360048201526000916001600160a01b031690632d0335ab90602401602060405180830381865afa158015610942573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109669190611578565b90506000610986888460065489336001600160a01b031660001b87610ea9565b6040805160028082526060820190925291925060009190816020015b60408051606081018252600080825260208083018290529282015282526000199092019101816109a257905050905060405180606001604052808881526020016000151581526020016001151581525081600081518110610a0557610a0561149a565b602002602001018190525060405180606001604052808781526020016001151581526020016001151581525081600181518110610a4457610a4461149a565b602090810291909101810191909152600254604080516060810182528c815280840185905260005482516001600160a01b03909416946331b433f994610ac694840192610aaf9260c09190911b918d91016001600160c01b0319929092168252600882015260280190565b6040516020818303038152906040528152506104f8565b8b8b631de44e3d87338a6040518863ffffffff1660e01b8152600401610af297969594939291906115c7565b6020604051808303816000875af1158015610b11573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b35919061147d565b506040518281527f1131472297a800fee664d1d89cfa8f7676ff07189ecc53f80bbb5f4969099db89060200160405180910390a1505050505050505050565b6001600160a01b038316610bb85760405162461bcd60e51b815260206004820152600b60248201526a424144204144445245535360a81b60448201526064016103c7565b600254604051632d0335ab60e01b81523360048201526000916001600160a01b031690632d0335ab90602401602060405180830381865afa158015610c01573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c259190611578565b600654909150600090610c49906001600160a01b0380871691869189163387610ea9565b90506000604051806060016040528060238152602001611615602391398051602090910120604051602481018990526044810184905260640160408051601f198184030181529181526020820180516001600160e01b03166001600160e01b031990941693909317909252600254915163329ef45f60e01b81529092506001600160a01b039091169063329ef45f90610cf29084908a908a908a90899033908c9060040161142d565b6020604051808303816000875af1158015610d11573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d35919061147d565b506040518281527f1131472297a800fee664d1d89cfa8f7676ff07189ecc53f80bbb5f4969099db8906020016104e7565b6002546001600160a01b03163314610db65760405162461bcd60e51b81526020600482015260136024820152724f4e4c59204345525441494e2042524944474560681b60448201526064016103c7565b60008181526003602052604090205415610e0057600480546001810182556000919091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b018190555b6000818152600360205260408120805491610e1a836114df565b909155505060058054906000610e2f836114df565b909155505060018290556040518281527f3e656638c321b7f315f9cd0ebbfbbb108d6d8fbad72ae54b16a2c1ac2dc1de0e9060200160405180910390a160408051828152602081018490527f254f734e74519a6ad860fd5a8c9fe8e49517cac9b3344a71c6e6d81e771d1509910160405180910390a15050565b60408051602080820194909452808201929092526060820195909552608081019390935260a083019190915260c0808301949094528251808303909401845260e0909101909152815191012090565b5080546000825590600052602060002090810190610f169190610f19565b50565b5b80821115610f2e5760008155600101610f1a565b5090565b80356001600160a01b0381168114610f4957600080fd5b919050565b60008060008060008060c08789031215610f6757600080fd5b86359550610f7760208801610f32565b9450610f8560408801610f32565b9350606087013592506080870135915060a087013590509295509295509295565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff81118282101715610fdf57610fdf610fa6565b60405290565b604051601f8201601f1916810167ffffffffffffffff8111828210171561100e5761100e610fa6565b604052919050565b8015158114610f1657600080fd5b600067ffffffffffffffff83111561103e5761103e610fa6565b611051601f8401601f1916602001610fe5565b905082815283838301111561106557600080fd5b828260208301376000602084830101529392505050565b600082601f83011261108d57600080fd5b61109c83833560208501611024565b9392505050565b600060208083850312156110b657600080fd5b823567ffffffffffffffff808211156110ce57600080fd5b818501915060608083880312156110e457600080fd5b6110ec610fbc565b83358152848401358381111561110157600080fd5b8401601f8101891361111257600080fd5b80358481111561112457611124610fa6565b611132878260051b01610fe5565b8181529084028201870190878101908b83111561114e57600080fd5b928801925b828410156111aa5785848d03121561116b5760008081fd5b611173610fbc565b843581528985013561118481611016565b818b015260408581013561119781611016565b9082015282529285019290880190611153565b80898601525050505060408401359450828511156111c757600080fd5b6111d38886860161107c565b6040820152979650505050505050565b60005b838110156111fe5781810151838201526020016111e6565b8381111561120d576000848401525b50505050565b6000815180845261122b8160208601602086016111e3565b601f01601f19169290920160200192915050565b60208152600061109c6020830184611213565b60006020828403121561126457600080fd5b813567ffffffffffffffff81111561127b57600080fd5b8201601f8101841361128c57600080fd5b61129b84823560208401611024565b949350505050565b6000806000606084860312156112b857600080fd5b6112c184610f32565b92506112cf60208501610f32565b9150604084013590509250925092565b6000602082840312156112f157600080fd5b5035919050565b6020808252825182820181905260009190848201906040850190845b8181101561133057835183529284019291840191600101611314565b50909695505050505050565b6000806000806080858703121561135257600080fd5b843567ffffffffffffffff81111561136957600080fd5b6113758782880161107c565b94505061138460208601610f32565b925061139260408601610f32565b9396929550929360600135925050565b60008060008060008060c087890312156113bb57600080fd5b505084359660208601359650604086013595606081013595506080810135945060a0013592509050565b600080600080608085870312156113fb57600080fd5b8435935061138460208601610f32565b6000806040838503121561141e57600080fd5b50508035926020909101359150565b60e08152600061144060e083018a611213565b6001600160a01b0398891660208401529688166040830152506060810194909452608084019290925290931660a082015260c00191909152919050565b60006020828403121561148f57600080fd5b815161109c81611016565b634e487b7160e01b600052603260045260246000fd5b600083516114c28184602088016111e3565b8351908301906114d68183602088016111e3565b01949350505050565b600060001982141561150157634e487b7160e01b600052601160045260246000fd5b5060010190565b6000835161151a8184602088016111e3565b9190910191825250602001919050565b6000835161153c8184602088016111e3565b60e09390931b6001600160e01b0319169190920190815260040192915050565b6000825161156e8184602087016111e3565b9190910192915050565b60006020828403121561158a57600080fd5b5051919050565b6080815260006115a46080830187611213565b6001600160a01b0395861660208401529390941660408201526060015292915050565b60e0815260006115da60e083018a611213565b6020830198909852506040810195909552606085019390935260808401919091526001600160a01b031660a083015260c09091015291905056fe7265636569766552657175657374546573742875696e743235362c6279746573333229a26469706673582212206b4698feaeca2556ee943e9db71be9e42591c9919e925230ddff9a1cac9e2bdb64736f6c634300080a0033",
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

// SOLANAED25519SIG is a free data retrieval call binding the contract method 0x054aabee.
//
// Solidity: function SOLANA_ED25519_SIG() view returns(bytes32)
func (_MockDexPool *MockDexPoolCaller) SOLANAED25519SIG(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MockDexPool.contract.Call(opts, &out, "SOLANA_ED25519_SIG")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SOLANAED25519SIG is a free data retrieval call binding the contract method 0x054aabee.
//
// Solidity: function SOLANA_ED25519_SIG() view returns(bytes32)
func (_MockDexPool *MockDexPoolSession) SOLANAED25519SIG() ([32]byte, error) {
	return _MockDexPool.Contract.SOLANAED25519SIG(&_MockDexPool.CallOpts)
}

// SOLANAED25519SIG is a free data retrieval call binding the contract method 0x054aabee.
//
// Solidity: function SOLANA_ED25519_SIG() view returns(bytes32)
func (_MockDexPool *MockDexPoolCallerSession) SOLANAED25519SIG() ([32]byte, error) {
	return _MockDexPool.Contract.SOLANAED25519SIG(&_MockDexPool.CallOpts)
}

// SOLANAINSTRUCTIONS is a free data retrieval call binding the contract method 0x842ce88d.
//
// Solidity: function SOLANA_INSTRUCTIONS() view returns(bytes32)
func (_MockDexPool *MockDexPoolCaller) SOLANAINSTRUCTIONS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MockDexPool.contract.Call(opts, &out, "SOLANA_INSTRUCTIONS")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SOLANAINSTRUCTIONS is a free data retrieval call binding the contract method 0x842ce88d.
//
// Solidity: function SOLANA_INSTRUCTIONS() view returns(bytes32)
func (_MockDexPool *MockDexPoolSession) SOLANAINSTRUCTIONS() ([32]byte, error) {
	return _MockDexPool.Contract.SOLANAINSTRUCTIONS(&_MockDexPool.CallOpts)
}

// SOLANAINSTRUCTIONS is a free data retrieval call binding the contract method 0x842ce88d.
//
// Solidity: function SOLANA_INSTRUCTIONS() view returns(bytes32)
func (_MockDexPool *MockDexPoolCallerSession) SOLANAINSTRUCTIONS() ([32]byte, error) {
	return _MockDexPool.Contract.SOLANAINSTRUCTIONS(&_MockDexPool.CallOpts)
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

// CalcRequestId is a free data retrieval call binding the contract method 0x671c2884.
//
// Solidity: function calcRequestId(address secondPartPool, address oppBridge, uint256 chainId) view returns(bytes32, uint256)
func (_MockDexPool *MockDexPoolCaller) CalcRequestId(opts *bind.CallOpts, secondPartPool common.Address, oppBridge common.Address, chainId *big.Int) ([32]byte, *big.Int, error) {
	var out []interface{}
	err := _MockDexPool.contract.Call(opts, &out, "calcRequestId", secondPartPool, oppBridge, chainId)

	if err != nil {
		return *new([32]byte), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// CalcRequestId is a free data retrieval call binding the contract method 0x671c2884.
//
// Solidity: function calcRequestId(address secondPartPool, address oppBridge, uint256 chainId) view returns(bytes32, uint256)
func (_MockDexPool *MockDexPoolSession) CalcRequestId(secondPartPool common.Address, oppBridge common.Address, chainId *big.Int) ([32]byte, *big.Int, error) {
	return _MockDexPool.Contract.CalcRequestId(&_MockDexPool.CallOpts, secondPartPool, oppBridge, chainId)
}

// CalcRequestId is a free data retrieval call binding the contract method 0x671c2884.
//
// Solidity: function calcRequestId(address secondPartPool, address oppBridge, uint256 chainId) view returns(bytes32, uint256)
func (_MockDexPool *MockDexPoolCallerSession) CalcRequestId(secondPartPool common.Address, oppBridge common.Address, chainId *big.Int) ([32]byte, *big.Int, error) {
	return _MockDexPool.Contract.CalcRequestId(&_MockDexPool.CallOpts, secondPartPool, oppBridge, chainId)
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

// DoubleRequestIds is a free data retrieval call binding the contract method 0x724fed5a.
//
// Solidity: function doubleRequestIds(uint256 ) view returns(bytes32)
func (_MockDexPool *MockDexPoolCaller) DoubleRequestIds(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MockDexPool.contract.Call(opts, &out, "doubleRequestIds", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DoubleRequestIds is a free data retrieval call binding the contract method 0x724fed5a.
//
// Solidity: function doubleRequestIds(uint256 ) view returns(bytes32)
func (_MockDexPool *MockDexPoolSession) DoubleRequestIds(arg0 *big.Int) ([32]byte, error) {
	return _MockDexPool.Contract.DoubleRequestIds(&_MockDexPool.CallOpts, arg0)
}

// DoubleRequestIds is a free data retrieval call binding the contract method 0x724fed5a.
//
// Solidity: function doubleRequestIds(uint256 ) view returns(bytes32)
func (_MockDexPool *MockDexPoolCallerSession) DoubleRequestIds(arg0 *big.Int) ([32]byte, error) {
	return _MockDexPool.Contract.DoubleRequestIds(&_MockDexPool.CallOpts, arg0)
}

// Doubles is a free data retrieval call binding the contract method 0xac9be189.
//
// Solidity: function doubles() view returns(bytes32[])
func (_MockDexPool *MockDexPoolCaller) Doubles(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _MockDexPool.contract.Call(opts, &out, "doubles")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// Doubles is a free data retrieval call binding the contract method 0xac9be189.
//
// Solidity: function doubles() view returns(bytes32[])
func (_MockDexPool *MockDexPoolSession) Doubles() ([][32]byte, error) {
	return _MockDexPool.Contract.Doubles(&_MockDexPool.CallOpts)
}

// Doubles is a free data retrieval call binding the contract method 0xac9be189.
//
// Solidity: function doubles() view returns(bytes32[])
func (_MockDexPool *MockDexPoolCallerSession) Doubles() ([][32]byte, error) {
	return _MockDexPool.Contract.Doubles(&_MockDexPool.CallOpts)
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

// ThisChainId is a free data retrieval call binding the contract method 0xe3d29c5d.
//
// Solidity: function thisChainId() view returns(uint256)
func (_MockDexPool *MockDexPoolCaller) ThisChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockDexPool.contract.Call(opts, &out, "thisChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ThisChainId is a free data retrieval call binding the contract method 0xe3d29c5d.
//
// Solidity: function thisChainId() view returns(uint256)
func (_MockDexPool *MockDexPoolSession) ThisChainId() (*big.Int, error) {
	return _MockDexPool.Contract.ThisChainId(&_MockDexPool.CallOpts)
}

// ThisChainId is a free data retrieval call binding the contract method 0xe3d29c5d.
//
// Solidity: function thisChainId() view returns(uint256)
func (_MockDexPool *MockDexPoolCallerSession) ThisChainId() (*big.Int, error) {
	return _MockDexPool.Contract.ThisChainId(&_MockDexPool.CallOpts)
}

// TotalRequests is a free data retrieval call binding the contract method 0x8aea61dc.
//
// Solidity: function totalRequests() view returns(uint256)
func (_MockDexPool *MockDexPoolCaller) TotalRequests(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockDexPool.contract.Call(opts, &out, "totalRequests")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRequests is a free data retrieval call binding the contract method 0x8aea61dc.
//
// Solidity: function totalRequests() view returns(uint256)
func (_MockDexPool *MockDexPoolSession) TotalRequests() (*big.Int, error) {
	return _MockDexPool.Contract.TotalRequests(&_MockDexPool.CallOpts)
}

// TotalRequests is a free data retrieval call binding the contract method 0x8aea61dc.
//
// Solidity: function totalRequests() view returns(uint256)
func (_MockDexPool *MockDexPoolCallerSession) TotalRequests() (*big.Int, error) {
	return _MockDexPool.Contract.TotalRequests(&_MockDexPool.CallOpts)
}

// ClearStats is a paid mutator transaction binding the contract method 0xe9096b57.
//
// Solidity: function clearStats() returns()
func (_MockDexPool *MockDexPoolTransactor) ClearStats(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockDexPool.contract.Transact(opts, "clearStats")
}
func (_MockDexPool *MockDexPoolTransactor) ClearStatsOverGsn(opts *bind.TransactOpts) (common.Hash, error) {
	return GsnExecutor(_MockDexPool.gsn, MockDexPoolMetaData.ABI, "clearStats")
}

// ClearStats is a paid mutator transaction binding the contract method 0xe9096b57.
//
// Solidity: function clearStats() returns()
func (_MockDexPool *MockDexPoolSession) ClearStats() (*types.Transaction, error) {
	return _MockDexPool.Contract.ClearStats(&_MockDexPool.TransactOpts)
}
func (_MockDexPool *MockDexPoolSession) ClearStatsOverGsn() (common.Hash, error) {
	return _MockDexPool.Contract.ClearStatsOverGsn(&_MockDexPool.TransactOpts)
}

// ClearStats is a paid mutator transaction binding the contract method 0xe9096b57.
//
// Solidity: function clearStats() returns()
func (_MockDexPool *MockDexPoolTransactorSession) ClearStats() (*types.Transaction, error) {
	return _MockDexPool.Contract.ClearStats(&_MockDexPool.TransactOpts)
}
func (_MockDexPool *MockDexPoolTransactorSession) ClearStatsOverGsn() (common.Hash, error) {
	return _MockDexPool.Contract.ClearStatsOverGsn(&_MockDexPool.TransactOpts)
}

// ReceiveRequestTest is a paid mutator transaction binding the contract method 0xfaad85c8.
//
// Solidity: function receiveRequestTest(uint256 _testData, bytes32 _reqId) returns()
func (_MockDexPool *MockDexPoolTransactor) ReceiveRequestTest(opts *bind.TransactOpts, _testData *big.Int, _reqId [32]byte) (*types.Transaction, error) {
	return _MockDexPool.contract.Transact(opts, "receiveRequestTest", _testData, _reqId)
}
func (_MockDexPool *MockDexPoolTransactor) ReceiveRequestTestOverGsn(opts *bind.TransactOpts, _testData *big.Int, _reqId [32]byte) (common.Hash, error) {
	return GsnExecutor(_MockDexPool.gsn, MockDexPoolMetaData.ABI, "receiveRequestTest", _testData, _reqId)
}

// ReceiveRequestTest is a paid mutator transaction binding the contract method 0xfaad85c8.
//
// Solidity: function receiveRequestTest(uint256 _testData, bytes32 _reqId) returns()
func (_MockDexPool *MockDexPoolSession) ReceiveRequestTest(_testData *big.Int, _reqId [32]byte) (*types.Transaction, error) {
	return _MockDexPool.Contract.ReceiveRequestTest(&_MockDexPool.TransactOpts, _testData, _reqId)
}
func (_MockDexPool *MockDexPoolSession) ReceiveRequestTestOverGsn(_testData *big.Int, _reqId [32]byte) (common.Hash, error) {
	return _MockDexPool.Contract.ReceiveRequestTestOverGsn(&_MockDexPool.TransactOpts, _testData, _reqId)
}

// ReceiveRequestTest is a paid mutator transaction binding the contract method 0xfaad85c8.
//
// Solidity: function receiveRequestTest(uint256 _testData, bytes32 _reqId) returns()
func (_MockDexPool *MockDexPoolTransactorSession) ReceiveRequestTest(_testData *big.Int, _reqId [32]byte) (*types.Transaction, error) {
	return _MockDexPool.Contract.ReceiveRequestTest(&_MockDexPool.TransactOpts, _testData, _reqId)
}
func (_MockDexPool *MockDexPoolTransactorSession) ReceiveRequestTestOverGsn(_testData *big.Int, _reqId [32]byte) (common.Hash, error) {
	return _MockDexPool.Contract.ReceiveRequestTestOverGsn(&_MockDexPool.TransactOpts, _testData, _reqId)
}

// SendRequestTestV2 is a paid mutator transaction binding the contract method 0xf9ee520e.
//
// Solidity: function sendRequestTestV2(uint256 testData_, address secondPartPool, address oppBridge, uint256 chainId) returns()
func (_MockDexPool *MockDexPoolTransactor) SendRequestTestV2(opts *bind.TransactOpts, testData_ *big.Int, secondPartPool common.Address, oppBridge common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _MockDexPool.contract.Transact(opts, "sendRequestTestV2", testData_, secondPartPool, oppBridge, chainId)
}
func (_MockDexPool *MockDexPoolTransactor) SendRequestTestV2OverGsn(opts *bind.TransactOpts, testData_ *big.Int, secondPartPool common.Address, oppBridge common.Address, chainId *big.Int) (common.Hash, error) {
	return GsnExecutor(_MockDexPool.gsn, MockDexPoolMetaData.ABI, "sendRequestTestV2", testData_, secondPartPool, oppBridge, chainId)
}

// SendRequestTestV2 is a paid mutator transaction binding the contract method 0xf9ee520e.
//
// Solidity: function sendRequestTestV2(uint256 testData_, address secondPartPool, address oppBridge, uint256 chainId) returns()
func (_MockDexPool *MockDexPoolSession) SendRequestTestV2(testData_ *big.Int, secondPartPool common.Address, oppBridge common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _MockDexPool.Contract.SendRequestTestV2(&_MockDexPool.TransactOpts, testData_, secondPartPool, oppBridge, chainId)
}
func (_MockDexPool *MockDexPoolSession) SendRequestTestV2OverGsn(testData_ *big.Int, secondPartPool common.Address, oppBridge common.Address, chainId *big.Int) (common.Hash, error) {
	return _MockDexPool.Contract.SendRequestTestV2OverGsn(&_MockDexPool.TransactOpts, testData_, secondPartPool, oppBridge, chainId)
}

// SendRequestTestV2 is a paid mutator transaction binding the contract method 0xf9ee520e.
//
// Solidity: function sendRequestTestV2(uint256 testData_, address secondPartPool, address oppBridge, uint256 chainId) returns()
func (_MockDexPool *MockDexPoolTransactorSession) SendRequestTestV2(testData_ *big.Int, secondPartPool common.Address, oppBridge common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _MockDexPool.Contract.SendRequestTestV2(&_MockDexPool.TransactOpts, testData_, secondPartPool, oppBridge, chainId)
}
func (_MockDexPool *MockDexPoolTransactorSession) SendRequestTestV2OverGsn(testData_ *big.Int, secondPartPool common.Address, oppBridge common.Address, chainId *big.Int) (common.Hash, error) {
	return _MockDexPool.Contract.SendRequestTestV2OverGsn(&_MockDexPool.TransactOpts, testData_, secondPartPool, oppBridge, chainId)
}

// SendRequestTestV2Unsafe is a paid mutator transaction binding the contract method 0x19d5c38e.
//
// Solidity: function sendRequestTestV2Unsafe(uint256 testData_, address secondPartPool, address oppBridge, uint256 chainId, bytes32 requestId, uint256 nonce) returns()
func (_MockDexPool *MockDexPoolTransactor) SendRequestTestV2Unsafe(opts *bind.TransactOpts, testData_ *big.Int, secondPartPool common.Address, oppBridge common.Address, chainId *big.Int, requestId [32]byte, nonce *big.Int) (*types.Transaction, error) {
	return _MockDexPool.contract.Transact(opts, "sendRequestTestV2Unsafe", testData_, secondPartPool, oppBridge, chainId, requestId, nonce)
}
func (_MockDexPool *MockDexPoolTransactor) SendRequestTestV2UnsafeOverGsn(opts *bind.TransactOpts, testData_ *big.Int, secondPartPool common.Address, oppBridge common.Address, chainId *big.Int, requestId [32]byte, nonce *big.Int) (common.Hash, error) {
	return GsnExecutor(_MockDexPool.gsn, MockDexPoolMetaData.ABI, "sendRequestTestV2Unsafe", testData_, secondPartPool, oppBridge, chainId, requestId, nonce)
}

// SendRequestTestV2Unsafe is a paid mutator transaction binding the contract method 0x19d5c38e.
//
// Solidity: function sendRequestTestV2Unsafe(uint256 testData_, address secondPartPool, address oppBridge, uint256 chainId, bytes32 requestId, uint256 nonce) returns()
func (_MockDexPool *MockDexPoolSession) SendRequestTestV2Unsafe(testData_ *big.Int, secondPartPool common.Address, oppBridge common.Address, chainId *big.Int, requestId [32]byte, nonce *big.Int) (*types.Transaction, error) {
	return _MockDexPool.Contract.SendRequestTestV2Unsafe(&_MockDexPool.TransactOpts, testData_, secondPartPool, oppBridge, chainId, requestId, nonce)
}
func (_MockDexPool *MockDexPoolSession) SendRequestTestV2UnsafeOverGsn(testData_ *big.Int, secondPartPool common.Address, oppBridge common.Address, chainId *big.Int, requestId [32]byte, nonce *big.Int) (common.Hash, error) {
	return _MockDexPool.Contract.SendRequestTestV2UnsafeOverGsn(&_MockDexPool.TransactOpts, testData_, secondPartPool, oppBridge, chainId, requestId, nonce)
}

// SendRequestTestV2Unsafe is a paid mutator transaction binding the contract method 0x19d5c38e.
//
// Solidity: function sendRequestTestV2Unsafe(uint256 testData_, address secondPartPool, address oppBridge, uint256 chainId, bytes32 requestId, uint256 nonce) returns()
func (_MockDexPool *MockDexPoolTransactorSession) SendRequestTestV2Unsafe(testData_ *big.Int, secondPartPool common.Address, oppBridge common.Address, chainId *big.Int, requestId [32]byte, nonce *big.Int) (*types.Transaction, error) {
	return _MockDexPool.Contract.SendRequestTestV2Unsafe(&_MockDexPool.TransactOpts, testData_, secondPartPool, oppBridge, chainId, requestId, nonce)
}
func (_MockDexPool *MockDexPoolTransactorSession) SendRequestTestV2UnsafeOverGsn(testData_ *big.Int, secondPartPool common.Address, oppBridge common.Address, chainId *big.Int, requestId [32]byte, nonce *big.Int) (common.Hash, error) {
	return _MockDexPool.Contract.SendRequestTestV2UnsafeOverGsn(&_MockDexPool.TransactOpts, testData_, secondPartPool, oppBridge, chainId, requestId, nonce)
}

// SendTest2 is a paid mutator transaction binding the contract method 0xd5854d70.
//
// Solidity: function sendTest2(bytes testData_, address secondPartPool_, address oppBridge_, uint256 chainId_) returns()
func (_MockDexPool *MockDexPoolTransactor) SendTest2(opts *bind.TransactOpts, testData_ []byte, secondPartPool_ common.Address, oppBridge_ common.Address, chainId_ *big.Int) (*types.Transaction, error) {
	return _MockDexPool.contract.Transact(opts, "sendTest2", testData_, secondPartPool_, oppBridge_, chainId_)
}
func (_MockDexPool *MockDexPoolTransactor) SendTest2OverGsn(opts *bind.TransactOpts, testData_ []byte, secondPartPool_ common.Address, oppBridge_ common.Address, chainId_ *big.Int) (common.Hash, error) {
	return GsnExecutor(_MockDexPool.gsn, MockDexPoolMetaData.ABI, "sendTest2", testData_, secondPartPool_, oppBridge_, chainId_)
}

// SendTest2 is a paid mutator transaction binding the contract method 0xd5854d70.
//
// Solidity: function sendTest2(bytes testData_, address secondPartPool_, address oppBridge_, uint256 chainId_) returns()
func (_MockDexPool *MockDexPoolSession) SendTest2(testData_ []byte, secondPartPool_ common.Address, oppBridge_ common.Address, chainId_ *big.Int) (*types.Transaction, error) {
	return _MockDexPool.Contract.SendTest2(&_MockDexPool.TransactOpts, testData_, secondPartPool_, oppBridge_, chainId_)
}
func (_MockDexPool *MockDexPoolSession) SendTest2OverGsn(testData_ []byte, secondPartPool_ common.Address, oppBridge_ common.Address, chainId_ *big.Int) (common.Hash, error) {
	return _MockDexPool.Contract.SendTest2OverGsn(&_MockDexPool.TransactOpts, testData_, secondPartPool_, oppBridge_, chainId_)
}

// SendTest2 is a paid mutator transaction binding the contract method 0xd5854d70.
//
// Solidity: function sendTest2(bytes testData_, address secondPartPool_, address oppBridge_, uint256 chainId_) returns()
func (_MockDexPool *MockDexPoolTransactorSession) SendTest2(testData_ []byte, secondPartPool_ common.Address, oppBridge_ common.Address, chainId_ *big.Int) (*types.Transaction, error) {
	return _MockDexPool.Contract.SendTest2(&_MockDexPool.TransactOpts, testData_, secondPartPool_, oppBridge_, chainId_)
}
func (_MockDexPool *MockDexPoolTransactorSession) SendTest2OverGsn(testData_ []byte, secondPartPool_ common.Address, oppBridge_ common.Address, chainId_ *big.Int) (common.Hash, error) {
	return _MockDexPool.Contract.SendTest2OverGsn(&_MockDexPool.TransactOpts, testData_, secondPartPool_, oppBridge_, chainId_)
}

// SendTestRequestToSolana is a paid mutator transaction binding the contract method 0xf54ea76b.
//
// Solidity: function sendTestRequestToSolana(bytes32 testStubPID_, bytes32 solBridgePID_, bytes32 dataAcc_, bytes32 bridgePDASigner_, uint256 testData_, uint256 chainId) returns()
func (_MockDexPool *MockDexPoolTransactor) SendTestRequestToSolana(opts *bind.TransactOpts, testStubPID_ [32]byte, solBridgePID_ [32]byte, dataAcc_ [32]byte, bridgePDASigner_ [32]byte, testData_ *big.Int, chainId *big.Int) (*types.Transaction, error) {
	return _MockDexPool.contract.Transact(opts, "sendTestRequestToSolana", testStubPID_, solBridgePID_, dataAcc_, bridgePDASigner_, testData_, chainId)
}
func (_MockDexPool *MockDexPoolTransactor) SendTestRequestToSolanaOverGsn(opts *bind.TransactOpts, testStubPID_ [32]byte, solBridgePID_ [32]byte, dataAcc_ [32]byte, bridgePDASigner_ [32]byte, testData_ *big.Int, chainId *big.Int) (common.Hash, error) {
	return GsnExecutor(_MockDexPool.gsn, MockDexPoolMetaData.ABI, "sendTestRequestToSolana", testStubPID_, solBridgePID_, dataAcc_, bridgePDASigner_, testData_, chainId)
}

// SendTestRequestToSolana is a paid mutator transaction binding the contract method 0xf54ea76b.
//
// Solidity: function sendTestRequestToSolana(bytes32 testStubPID_, bytes32 solBridgePID_, bytes32 dataAcc_, bytes32 bridgePDASigner_, uint256 testData_, uint256 chainId) returns()
func (_MockDexPool *MockDexPoolSession) SendTestRequestToSolana(testStubPID_ [32]byte, solBridgePID_ [32]byte, dataAcc_ [32]byte, bridgePDASigner_ [32]byte, testData_ *big.Int, chainId *big.Int) (*types.Transaction, error) {
	return _MockDexPool.Contract.SendTestRequestToSolana(&_MockDexPool.TransactOpts, testStubPID_, solBridgePID_, dataAcc_, bridgePDASigner_, testData_, chainId)
}
func (_MockDexPool *MockDexPoolSession) SendTestRequestToSolanaOverGsn(testStubPID_ [32]byte, solBridgePID_ [32]byte, dataAcc_ [32]byte, bridgePDASigner_ [32]byte, testData_ *big.Int, chainId *big.Int) (common.Hash, error) {
	return _MockDexPool.Contract.SendTestRequestToSolanaOverGsn(&_MockDexPool.TransactOpts, testStubPID_, solBridgePID_, dataAcc_, bridgePDASigner_, testData_, chainId)
}

// SendTestRequestToSolana is a paid mutator transaction binding the contract method 0xf54ea76b.
//
// Solidity: function sendTestRequestToSolana(bytes32 testStubPID_, bytes32 solBridgePID_, bytes32 dataAcc_, bytes32 bridgePDASigner_, uint256 testData_, uint256 chainId) returns()
func (_MockDexPool *MockDexPoolTransactorSession) SendTestRequestToSolana(testStubPID_ [32]byte, solBridgePID_ [32]byte, dataAcc_ [32]byte, bridgePDASigner_ [32]byte, testData_ *big.Int, chainId *big.Int) (*types.Transaction, error) {
	return _MockDexPool.Contract.SendTestRequestToSolana(&_MockDexPool.TransactOpts, testStubPID_, solBridgePID_, dataAcc_, bridgePDASigner_, testData_, chainId)
}
func (_MockDexPool *MockDexPoolTransactorSession) SendTestRequestToSolanaOverGsn(testStubPID_ [32]byte, solBridgePID_ [32]byte, dataAcc_ [32]byte, bridgePDASigner_ [32]byte, testData_ *big.Int, chainId *big.Int) (common.Hash, error) {
	return _MockDexPool.Contract.SendTestRequestToSolanaOverGsn(&_MockDexPool.TransactOpts, testStubPID_, solBridgePID_, dataAcc_, bridgePDASigner_, testData_, chainId)
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

// MockDexPoolTestEventIterator is returned from FilterTestEvent and is used to iterate over the raw logs and unpacked data for TestEvent events raised by the MockDexPool contract.
type MockDexPoolTestEventIterator struct {
	Event *MockDexPoolTestEvent // Event containing the contract specifics and raw log

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
func (it *MockDexPoolTestEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockDexPoolTestEvent)
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
		it.Event = new(MockDexPoolTestEvent)
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
func (it *MockDexPoolTestEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockDexPoolTestEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockDexPoolTestEvent represents a TestEvent event raised by the MockDexPool contract.
type MockDexPoolTestEvent struct {
	TestData       []byte
	SecondPartPool common.Address
	OppBridge      common.Address
	ChainId        *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTestEvent is a free log retrieval operation binding the contract event 0x6732d52ba34d766a4a84b7fe8ccca6e4a4b89895ca1c38c9cc5b88ea59d97dd4.
//
// Solidity: event TestEvent(bytes testData_, address secondPartPool, address oppBridge, uint256 chainId)
func (_MockDexPool *MockDexPoolFilterer) FilterTestEvent(opts *bind.FilterOpts) (*MockDexPoolTestEventIterator, error) {

	logs, sub, err := _MockDexPool.contract.FilterLogs(opts, "TestEvent")
	if err != nil {
		return nil, err
	}
	return &MockDexPoolTestEventIterator{contract: _MockDexPool.contract, event: "TestEvent", logs: logs, sub: sub}, nil
}

// WatchTestEvent is a free log subscription operation binding the contract event 0x6732d52ba34d766a4a84b7fe8ccca6e4a4b89895ca1c38c9cc5b88ea59d97dd4.
//
// Solidity: event TestEvent(bytes testData_, address secondPartPool, address oppBridge, uint256 chainId)
func (_MockDexPool *MockDexPoolFilterer) WatchTestEvent(opts *bind.WatchOpts, sink chan<- *MockDexPoolTestEvent) (event.Subscription, error) {

	logs, sub, err := _MockDexPool.contract.WatchLogs(opts, "TestEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockDexPoolTestEvent)
				if err := _MockDexPool.contract.UnpackLog(event, "TestEvent", log); err != nil {
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

// ParseTestEvent is a log parse operation binding the contract event 0x6732d52ba34d766a4a84b7fe8ccca6e4a4b89895ca1c38c9cc5b88ea59d97dd4.
//
// Solidity: event TestEvent(bytes testData_, address secondPartPool, address oppBridge, uint256 chainId)
func (_MockDexPool *MockDexPoolFilterer) ParseTestEvent(log types.Log) (*MockDexPoolTestEvent, error) {
	event := new(MockDexPoolTestEvent)
	if err := _MockDexPool.contract.UnpackLog(event, "TestEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
