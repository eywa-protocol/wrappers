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

// RelayerPoolABI is the input ABI used to generate the binding from.
const RelayerPoolABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_relayerFeeNumerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_emissionAnnualRateNumerator\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockTill\",\"type\":\"uint256\"}],\"name\":\"DepositPut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rest\",\"type\":\"uint256\"}],\"name\":\"DepositWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"EmissionAnnualRateNumeratorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"harvestForPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardForPool\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardPerTokenNumeratorBefore\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardPerTokenNumerator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDeposit\",\"type\":\"uint256\"}],\"name\":\"HarvestPoolReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"RelayerFeeNumeratorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumRelayerPool.RelayerStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"RelayerStatusSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userDeposit\",\"type\":\"uint256\"}],\"name\":\"UserHarvestReward\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MIN_RELAYER_COLLATERAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_RELAYER_STAKING_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_STAKING_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lockTill\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emissionAnnualRateNumerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_depositId\",\"type\":\"uint256\"}],\"name\":\"getDeposit\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockTill\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvestMyReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvestPoolReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastHarvestRewardTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relayerFeeNumerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relayerStatus\",\"outputs\":[{\"internalType\":\"enumRelayerPool.RelayerStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerTokenNumerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setEmissionAnnualRateNumerator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setRelayerFeeNumerator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_depositId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RelayerPoolBin is the compiled bytecode used for deploying new contracts.
var RelayerPoolBin = "0x60806040523480156200001157600080fd5b5060405162001a5838038062001a5883398101604081905262000034916200038b565b600160005560408051808201909152600e81526d4645455f49535f544f4f5f4c4f5760901b602082015260648410156200008c5760405162461bcd60e51b8152600401620000839190620003f9565b60405180910390fd5b5060408051808201909152600f81526e08c8a8abe92a6bea89e9ebe90928e9608b1b6020820152612710841115620000d95760405162461bcd60e51b8152600401620000839190620003f9565b50600c8390556301e13380620000f2836127106200044f565b10156040518060400160405280601f81526020017f454d495353494f4e5f414e4e55414c5f524154455f49535f544f4f5f4c4f5700815250906200014b5760405162461bcd60e51b8152600401620000839190620003f9565b5063bbf81e008211156040518060400160405280602081526020017f454d495353494f4e5f414e4e55414c5f524154455f49535f544f4f5f4849474881525090620001ab5760405162461bcd60e51b8152600401620000839190620003f9565b50600d82905560408051808201909152600c81526b5a45524f5f4144445245535360a01b60208201526001600160a01b038616620001fe5760405162461bcd60e51b8152600401620000839190620003f9565b50601180546001600160a01b0319166001600160a01b038781169190911790915560408051808201909152600c81526b5a45524f5f4144445245535360a01b6020820152908516620002655760405162461bcd60e51b8152600401620000839190620003f9565b50601080546001600160a01b0319166001600160a01b038681169190911790915560408051808201909152600c81526b5a45524f5f4144445245535360a01b6020820152908716620002cc5760405162461bcd60e51b8152600401620000839190620003f9565b50600e80546001600160a01b038089166001600160a01b031992831617909255600f80549091163317905542600b5560408051808201909152600c81526b5a45524f5f4144445245535360a01b6020820152908216620003415760405162461bcd60e51b8152600401620000839190620003f9565b50601280546001600160a01b0319166001600160a01b0392909216919091179055506200047b9350505050565b80516001600160a01b03811681146200038657600080fd5b919050565b60008060008060008060c08789031215620003a4578182fd5b620003af876200036e565b9550620003bf602088016200036e565b9450620003cf604088016200036e565b93506060870151925060808701519150620003ed60a088016200036e565b90509295509295509295565b6000602080835283518082850152825b81811015620004275785810183015185820160400152820162000409565b81811115620004395783604083870101525b50601f01601f1916929092016040019392505050565b60008160001904831182151516156200047657634e487b7160e01b81526011600452602481fd5b500290565b6115cd806200048b6000396000f3fe608060405234801561001057600080fd5b50600436106101425760003560e01c8063979853a2116100b8578063d1239f371161007c578063d1239f371461023f578063e61a958214610247578063ec51364e1461024f578063f6153ccd14610262578063f7c618c11461026a578063fbfa77cf1461027257610142565b8063979853a2146101e75780639f9fb968146101ef578063b02c43d014610211578063b6b55f2514610224578063c89039c51461023757610142565b80635bf399d41161010a5780635bf399d41461019f5780637b103999146101a75780638363deb6146101bc5780638be9d5bf146101c45780638da5cb5b146101cc578063959ccfcb146101d457610142565b8063061887b3146101475780631b49a92614610165578063343d597c1461017a578063441a3e7014610182578063519b9a9e14610197575b600080fd5b61014f61027a565b60405161015c9190611485565b60405180910390f35b61016d610280565b60405161015c9190611293565b61014f610289565b6101956101903660046111e4565b610290565b005b61014f610689565b61014f61068f565b6101af61069b565b60405161015c9190611221565b6101956106aa565b6101956106b4565b6101af6106bc565b6101956101e23660046111cc565b6106cb565b61014f6107c9565b6102026101fd3660046111cc565b6107d0565b60405161015c93929190611272565b61020261021f3660046111cc565b610811565b6101956102323660046111cc565b61083c565b6101af610a45565b61014f610a54565b61014f610a5a565b61019561025d3660046111cc565b610a60565b61014f610b86565b6101af610b8c565b6101af610b9b565b600b5481565b60025460ff1681565b6212750081565b6000828152600460209081526040918290208054835180850190945260118452702727aa2fa222a827a9a4aa2fa7aba722a960791b9284019290925291906001600160a01b031633146102ff5760405162461bcd60e51b81526004016102f691906112bb565b60405180910390fd5b50806002015482111560405180604001604052806014815260200173125394d551919250d251539517d1115413d4d25560621b815250906103535760405162461bcd60e51b81526004016102f691906112bb565b5080600101544210156040518060400160405280601181526020017011115413d4d25517d254d7d313d0d2d151607a1b815250906103a45760405162461bcd60e51b81526004016102f691906112bb565b503360009081526005602052604090206103be9084610baa565b60405180604001604052806012815260200171444154415f494e434f4e53495354454e435960701b815250906104075760405162461bcd60e51b81526004016102f691906112bb565b50610410610bbf565b610418610d1f565b670de0b6b3a76400006009548361042f9190611504565b61043991906114e4565b3360009081526008602052604081208054909190610458908490611523565b909155505060028101548210156104c9578181600201600082825461047d9190611523565b90915550506002810154604051849133917f4959c2f9b36bf25a8c0c413e521adc24f9eeaacd05f94c0eee57ace9bf8b4634916104bc9187919061148e565b60405180910390a3610599565b600083815260046020908152604080832080546001600160a01b031916815560018101849055600201839055338352600590915290206105099084610df7565b60405180604001604052806012815260200171444154415f494e434f4e53495354454e435960701b815250906105525760405162461bcd60e51b81526004016102f691906112bb565b5082336001600160a01b03167f4959c2f9b36bf25a8c0c413e521adc24f9eeaacd05f94c0eee57ace9bf8b463484600060405161059092919061148e565b60405180910390a35b33600090815260066020526040812080548492906105b8908490611523565b9250508190555081600760008282546105d19190611523565b9091555050600e546001600160a01b031633141561066d57600e546001600160a01b031660009081526006602081905260409091205461061091611504565b60075411156106315760405162461bcd60e51b81526004016102f690611358565b600a54600e546001600160a01b0316600090815260066020526040902054101561066d5760405162461bcd60e51b81526004016102f690611428565b601054610684906001600160a01b03163384610e03565b505050565b600d5481565b670de0b6b3a764000081565b600f546001600160a01b031681565b6106b2610d1f565b565b6106b2610bbf565b600e546001600160a01b031681565b600e546001600160a01b031633146106f55760405162461bcd60e51b81526004016102f690611334565b60408051808201909152600e81526d4645455f49535f544f4f5f4c4f5760901b6020820152606482101561073c5760405162461bcd60e51b81526004016102f691906112bb565b5060408051808201909152600f81526e08c8a8abe92a6bea89e9ebe90928e9608b1b60208201526127108211156107865760405162461bcd60e51b81526004016102f691906112bb565b50600c81905560405133907f673d70d178eaf2ea132ef62a70f12374909ee3364342c07575e53a4625e51185906107be908490611485565b60405180910390a250565b6224ea0081565b600090815260046020908152604091829020825160608101845281546001600160a01b03168082526001830154938201849052600290920154930183905292565b6004602052600090815260409020805460018201546002909201546001600160a01b03909116919083565b610844610bbf565b61084c610d1f565b600380546000918261085d83611566565b90915550600e549091506000906001600160a01b031633141561088e576108876224ea00426114cc565b905061089e565b61089b62127500426114cc565b90505b60408051606081018252338082526020808301858152838501888152600088815260048452868120955186546001600160a01b0319166001600160a01b03909116178655915160018601555160029094019390935590825260069052908120805485929061090d9084906114cc565b92505081905550826007600082825461092691906114cc565b9091555050600954670de0b6b3a7640000906109429085611504565b61094c91906114e4565b336000908152600860205260408120805490919061096b9084906114cc565b9091555050600e546001600160a01b031633146109ca57600e546001600160a01b03166000908152600660208190526040909120546109a991611504565b60075411156109ca5760405162461bcd60e51b81526004016102f690611358565b3360009081526005602052604090206109e39083610e59565b506010546109fc906001600160a01b0316333086610e65565b81336001600160a01b03167f7bacd633fd842a77f4647e920fa8675614c879eb3267a371ef7208363c01cdee8584604051610a3892919061148e565b60405180910390a3505050565b6010546001600160a01b031681565b60095481565b600c5481565b600e546001600160a01b03163314610a8a5760405162461bcd60e51b81526004016102f690611334565b6301e13380610a9b82612710611504565b10156040518060400160405280601f81526020017f454d495353494f4e5f414e4e55414c5f524154455f49535f544f4f5f4c4f570081525090610af15760405162461bcd60e51b81526004016102f691906112bb565b5063bbf81e008111156040518060400160405280602081526020017f454d495353494f4e5f414e4e55414c5f524154455f49535f544f4f5f4849474881525090610b4e5760405162461bcd60e51b81526004016102f691906112bb565b50600d81905560405133907f4956113bd03fef454b6d6c1c9525fc977eb4f4fb1d01f45c85ea86ecaa46b887906107be908490611485565b60075481565b6011546001600160a01b031681565b6012546001600160a01b031681565b6000610bb68383610e8c565b90505b92915050565b6000600b5442610bcf9190611523565b905060006301e13380600d5483600754610be99190611504565b610bf39190611504565b610bfd91906114e4565b905080610c0b5750506106b2565b42600b55600c5460009061271090610c239084611504565b610c2d91906114e4565b90506000610c3b8284611523565b60095460075491925090610c57670de0b6b3a764000084611504565b610c6191906114e4565b60096000828254610c7291906114cc565b9091555050601254601154610c95916001600160a01b0391821691163087610e65565b8215610cbf57601254600e54601154610cbf926001600160a01b0391821692908216911686610e65565b600e546009546007546040516001600160a01b039093169233927fadd1c21c3f7b25aa60bceee6e16a6f94dfd90fbb5809c95e74ffc682c3cfc6a792610d10928b928b928b928b928b92909161149c565b60405180910390a35050505050565b3360009081526006602081815260408084205460088352908420549290915260095490929190670de0b6b3a764000090610d5a908590611504565b610d6491906114e4565b610d6e9190611523565b905080610d7c5750506106b2565b601154610d93906001600160a01b03163383610e03565b3360009081526008602052604081208054839290610db29084906114cc565b909155505060405133907fbeb525c2d0c03a594fed4cea83f02931246206387586fb3704201174b5ea8b7190610deb908490869061148e565b60405180910390a25050565b6000610bb68383610ea4565b6106848363a9059cbb60e01b8484604051602401610e22929190611259565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152610fbb565b6000610bb6838361104a565b610e86846323b872dd60e01b858585604051602401610e2293929190611235565b50505050565b60009081526001919091016020526040902054151590565b60008181526001830160205260408120548015610fb1576000610ec8600183611523565b8554909150600090610edc90600190611523565b90506000866000018281548110610f0357634e487b7160e01b600052603260045260246000fd5b9060005260206000200154905080876000018481548110610f3457634e487b7160e01b600052603260045260246000fd5b600091825260208083209091019290925582815260018901909152604090208490558654879080610f7557634e487b7160e01b600052603160045260246000fd5b60019003818190600052602060002001600090559055866001016000878152602001908152602001600020600090556001945050505050610bb9565b6000915050610bb9565b6000611010826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166110949092919063ffffffff16565b805190915015610684578080602001905181019061102e91906111ac565b6106845760405162461bcd60e51b81526004016102f6906113de565b60006110568383610e8c565b61108c57508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610bb9565b506000610bb9565b60606110a384846000856110ad565b90505b9392505050565b6060824710156110cf5760405162461bcd60e51b81526004016102f6906112ee565b6110d88561116d565b6110f45760405162461bcd60e51b81526004016102f6906113a7565b600080866001600160a01b031685876040516111109190611205565b60006040518083038185875af1925050503d806000811461114d576040519150601f19603f3d011682016040523d82523d6000602084013e611152565b606091505b5091509150611162828286611173565b979650505050505050565b3b151590565b606083156111825750816110a6565b8251156111925782518084602001fd5b8160405162461bcd60e51b81526004016102f691906112bb565b6000602082840312156111bd578081fd5b815180151581146110a6578182fd5b6000602082840312156111dd578081fd5b5035919050565b600080604083850312156111f6578081fd5b50508035926020909101359150565b6000825161121781846020870161153a565b9190910192915050565b6001600160a01b0391909116815260200190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b6001600160a01b03929092168252602082015260400190565b6001600160a01b039390931683526020830191909152604082015260600190565b60208101600483106112b557634e487b7160e01b600052602160045260246000fd5b91905290565b60006020825282518060208401526112da81604085016020870161153a565b601f01601f19169190910160400192915050565b60208082526026908201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6040820152651c8818d85b1b60d21b606082015260800190565b6020808252600a908201526937b7363c9037bbb732b960b11b604082015260600190565b6020808252602f908201527f736d616c6c206f776e6572207374616b6520286f776e65725374616b65722a3660408201526e203e3d20746f74616c5374616b652960881b606082015260800190565b6020808252601d908201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604082015260600190565b6020808252602a908201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6040820152691bdd081cdd58d8d9595960b21b606082015260800190565b60208082526037908201527f736d616c6c206f776e6572207374616b6520286f776e65725374616b65203e3d60408201527f205f6d696e5f6f776e65725f636f6c6c61746572616c29000000000000000000606082015260800190565b90815260200190565b918252602082015260400190565b968752602087019590955260408601939093526060850191909152608084015260a083015260c082015260e00190565b600082198211156114df576114df611581565b500190565b6000826114ff57634e487b7160e01b81526012600452602481fd5b500490565b600081600019048311821515161561151e5761151e611581565b500290565b60008282101561153557611535611581565b500390565b60005b8381101561155557818101518382015260200161153d565b83811115610e865750506000910152565b600060001982141561157a5761157a611581565b5060010190565b634e487b7160e01b600052601160045260246000fdfea264697066735822122090f7137b4f892c6551ca0f07e8a78007b3bac3588f4ef844823df36e066b4d8464736f6c63430008000033"

// DeployRelayerPool deploys a new Ethereum contract, binding an instance of RelayerPool to it.
func DeployRelayerPool(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _rewardToken common.Address, _depositToken common.Address, _relayerFeeNumerator *big.Int, _emissionAnnualRateNumerator *big.Int, _vault common.Address) (common.Address, *types.Transaction, *RelayerPool, error) {
	parsed, err := abi.JSON(strings.NewReader(RelayerPoolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RelayerPoolBin), backend, _owner, _rewardToken, _depositToken, _relayerFeeNumerator, _emissionAnnualRateNumerator, _vault)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RelayerPool{RelayerPoolCaller: RelayerPoolCaller{contract: contract}, RelayerPoolTransactor: RelayerPoolTransactor{contract: contract}, RelayerPoolFilterer: RelayerPoolFilterer{contract: contract}}, nil
}

// RelayerPool is an auto generated Go binding around an Ethereum contract.
type RelayerPool struct {
	RelayerPoolCaller     // Read-only binding to the contract
	RelayerPoolTransactor // Write-only binding to the contract
	RelayerPoolFilterer   // Log filterer for contract events
}

// RelayerPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type RelayerPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RelayerPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RelayerPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RelayerPoolSession struct {
	Contract     *RelayerPool      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RelayerPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RelayerPoolCallerSession struct {
	Contract *RelayerPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// RelayerPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RelayerPoolTransactorSession struct {
	Contract     *RelayerPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RelayerPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type RelayerPoolRaw struct {
	Contract *RelayerPool // Generic contract binding to access the raw methods on
}

// RelayerPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RelayerPoolCallerRaw struct {
	Contract *RelayerPoolCaller // Generic read-only contract binding to access the raw methods on
}

// RelayerPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RelayerPoolTransactorRaw struct {
	Contract *RelayerPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRelayerPool creates a new instance of RelayerPool, bound to a specific deployed contract.
func NewRelayerPool(address common.Address, backend bind.ContractBackend) (*RelayerPool, error) {
	contract, err := bindRelayerPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RelayerPool{RelayerPoolCaller: RelayerPoolCaller{contract: contract}, RelayerPoolTransactor: RelayerPoolTransactor{contract: contract}, RelayerPoolFilterer: RelayerPoolFilterer{contract: contract}}, nil
}

// NewRelayerPoolCaller creates a new read-only instance of RelayerPool, bound to a specific deployed contract.
func NewRelayerPoolCaller(address common.Address, caller bind.ContractCaller) (*RelayerPoolCaller, error) {
	contract, err := bindRelayerPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RelayerPoolCaller{contract: contract}, nil
}

// NewRelayerPoolTransactor creates a new write-only instance of RelayerPool, bound to a specific deployed contract.
func NewRelayerPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*RelayerPoolTransactor, error) {
	contract, err := bindRelayerPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RelayerPoolTransactor{contract: contract}, nil
}

// NewRelayerPoolFilterer creates a new log filterer instance of RelayerPool, bound to a specific deployed contract.
func NewRelayerPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*RelayerPoolFilterer, error) {
	contract, err := bindRelayerPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RelayerPoolFilterer{contract: contract}, nil
}

// bindRelayerPool binds a generic wrapper to an already deployed contract.
func bindRelayerPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RelayerPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RelayerPool *RelayerPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RelayerPool.Contract.RelayerPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RelayerPool *RelayerPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayerPool.Contract.RelayerPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RelayerPool *RelayerPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RelayerPool.Contract.RelayerPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RelayerPool *RelayerPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RelayerPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RelayerPool *RelayerPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayerPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RelayerPool *RelayerPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RelayerPool.Contract.contract.Transact(opts, method, params...)
}

// MINRELAYERCOLLATERAL is a free data retrieval call binding the contract method 0x5bf399d4.
//
// Solidity: function MIN_RELAYER_COLLATERAL() view returns(uint256)
func (_RelayerPool *RelayerPoolCaller) MINRELAYERCOLLATERAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RelayerPool.contract.Call(opts, &out, "MIN_RELAYER_COLLATERAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINRELAYERCOLLATERAL is a free data retrieval call binding the contract method 0x5bf399d4.
//
// Solidity: function MIN_RELAYER_COLLATERAL() view returns(uint256)
func (_RelayerPool *RelayerPoolSession) MINRELAYERCOLLATERAL() (*big.Int, error) {
	return _RelayerPool.Contract.MINRELAYERCOLLATERAL(&_RelayerPool.CallOpts)
}

// MINRELAYERCOLLATERAL is a free data retrieval call binding the contract method 0x5bf399d4.
//
// Solidity: function MIN_RELAYER_COLLATERAL() view returns(uint256)
func (_RelayerPool *RelayerPoolCallerSession) MINRELAYERCOLLATERAL() (*big.Int, error) {
	return _RelayerPool.Contract.MINRELAYERCOLLATERAL(&_RelayerPool.CallOpts)
}

// MINRELAYERSTAKINGTIME is a free data retrieval call binding the contract method 0x979853a2.
//
// Solidity: function MIN_RELAYER_STAKING_TIME() view returns(uint256)
func (_RelayerPool *RelayerPoolCaller) MINRELAYERSTAKINGTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RelayerPool.contract.Call(opts, &out, "MIN_RELAYER_STAKING_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINRELAYERSTAKINGTIME is a free data retrieval call binding the contract method 0x979853a2.
//
// Solidity: function MIN_RELAYER_STAKING_TIME() view returns(uint256)
func (_RelayerPool *RelayerPoolSession) MINRELAYERSTAKINGTIME() (*big.Int, error) {
	return _RelayerPool.Contract.MINRELAYERSTAKINGTIME(&_RelayerPool.CallOpts)
}

// MINRELAYERSTAKINGTIME is a free data retrieval call binding the contract method 0x979853a2.
//
// Solidity: function MIN_RELAYER_STAKING_TIME() view returns(uint256)
func (_RelayerPool *RelayerPoolCallerSession) MINRELAYERSTAKINGTIME() (*big.Int, error) {
	return _RelayerPool.Contract.MINRELAYERSTAKINGTIME(&_RelayerPool.CallOpts)
}

// MINSTAKINGTIME is a free data retrieval call binding the contract method 0x343d597c.
//
// Solidity: function MIN_STAKING_TIME() view returns(uint256)
func (_RelayerPool *RelayerPoolCaller) MINSTAKINGTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RelayerPool.contract.Call(opts, &out, "MIN_STAKING_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINSTAKINGTIME is a free data retrieval call binding the contract method 0x343d597c.
//
// Solidity: function MIN_STAKING_TIME() view returns(uint256)
func (_RelayerPool *RelayerPoolSession) MINSTAKINGTIME() (*big.Int, error) {
	return _RelayerPool.Contract.MINSTAKINGTIME(&_RelayerPool.CallOpts)
}

// MINSTAKINGTIME is a free data retrieval call binding the contract method 0x343d597c.
//
// Solidity: function MIN_STAKING_TIME() view returns(uint256)
func (_RelayerPool *RelayerPoolCallerSession) MINSTAKINGTIME() (*big.Int, error) {
	return _RelayerPool.Contract.MINSTAKINGTIME(&_RelayerPool.CallOpts)
}

// DepositToken is a free data retrieval call binding the contract method 0xc89039c5.
//
// Solidity: function depositToken() view returns(address)
func (_RelayerPool *RelayerPoolCaller) DepositToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RelayerPool.contract.Call(opts, &out, "depositToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DepositToken is a free data retrieval call binding the contract method 0xc89039c5.
//
// Solidity: function depositToken() view returns(address)
func (_RelayerPool *RelayerPoolSession) DepositToken() (common.Address, error) {
	return _RelayerPool.Contract.DepositToken(&_RelayerPool.CallOpts)
}

// DepositToken is a free data retrieval call binding the contract method 0xc89039c5.
//
// Solidity: function depositToken() view returns(address)
func (_RelayerPool *RelayerPoolCallerSession) DepositToken() (common.Address, error) {
	return _RelayerPool.Contract.DepositToken(&_RelayerPool.CallOpts)
}

// Deposits is a free data retrieval call binding the contract method 0xb02c43d0.
//
// Solidity: function deposits(uint256 ) view returns(address user, uint256 lockTill, uint256 amount)
func (_RelayerPool *RelayerPoolCaller) Deposits(opts *bind.CallOpts, arg0 *big.Int) (struct {
	User     common.Address
	LockTill *big.Int
	Amount   *big.Int
}, error) {
	var out []interface{}
	err := _RelayerPool.contract.Call(opts, &out, "deposits", arg0)

	outstruct := new(struct {
		User     common.Address
		LockTill *big.Int
		Amount   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.LockTill = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Deposits is a free data retrieval call binding the contract method 0xb02c43d0.
//
// Solidity: function deposits(uint256 ) view returns(address user, uint256 lockTill, uint256 amount)
func (_RelayerPool *RelayerPoolSession) Deposits(arg0 *big.Int) (struct {
	User     common.Address
	LockTill *big.Int
	Amount   *big.Int
}, error) {
	return _RelayerPool.Contract.Deposits(&_RelayerPool.CallOpts, arg0)
}

// Deposits is a free data retrieval call binding the contract method 0xb02c43d0.
//
// Solidity: function deposits(uint256 ) view returns(address user, uint256 lockTill, uint256 amount)
func (_RelayerPool *RelayerPoolCallerSession) Deposits(arg0 *big.Int) (struct {
	User     common.Address
	LockTill *big.Int
	Amount   *big.Int
}, error) {
	return _RelayerPool.Contract.Deposits(&_RelayerPool.CallOpts, arg0)
}

// EmissionAnnualRateNumerator is a free data retrieval call binding the contract method 0x519b9a9e.
//
// Solidity: function emissionAnnualRateNumerator() view returns(uint256)
func (_RelayerPool *RelayerPoolCaller) EmissionAnnualRateNumerator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RelayerPool.contract.Call(opts, &out, "emissionAnnualRateNumerator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EmissionAnnualRateNumerator is a free data retrieval call binding the contract method 0x519b9a9e.
//
// Solidity: function emissionAnnualRateNumerator() view returns(uint256)
func (_RelayerPool *RelayerPoolSession) EmissionAnnualRateNumerator() (*big.Int, error) {
	return _RelayerPool.Contract.EmissionAnnualRateNumerator(&_RelayerPool.CallOpts)
}

// EmissionAnnualRateNumerator is a free data retrieval call binding the contract method 0x519b9a9e.
//
// Solidity: function emissionAnnualRateNumerator() view returns(uint256)
func (_RelayerPool *RelayerPoolCallerSession) EmissionAnnualRateNumerator() (*big.Int, error) {
	return _RelayerPool.Contract.EmissionAnnualRateNumerator(&_RelayerPool.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0x9f9fb968.
//
// Solidity: function getDeposit(uint256 _depositId) view returns(address user, uint256 amount, uint256 lockTill)
func (_RelayerPool *RelayerPoolCaller) GetDeposit(opts *bind.CallOpts, _depositId *big.Int) (struct {
	User     common.Address
	Amount   *big.Int
	LockTill *big.Int
}, error) {
	var out []interface{}
	err := _RelayerPool.contract.Call(opts, &out, "getDeposit", _depositId)

	outstruct := new(struct {
		User     common.Address
		Amount   *big.Int
		LockTill *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LockTill = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetDeposit is a free data retrieval call binding the contract method 0x9f9fb968.
//
// Solidity: function getDeposit(uint256 _depositId) view returns(address user, uint256 amount, uint256 lockTill)
func (_RelayerPool *RelayerPoolSession) GetDeposit(_depositId *big.Int) (struct {
	User     common.Address
	Amount   *big.Int
	LockTill *big.Int
}, error) {
	return _RelayerPool.Contract.GetDeposit(&_RelayerPool.CallOpts, _depositId)
}

// GetDeposit is a free data retrieval call binding the contract method 0x9f9fb968.
//
// Solidity: function getDeposit(uint256 _depositId) view returns(address user, uint256 amount, uint256 lockTill)
func (_RelayerPool *RelayerPoolCallerSession) GetDeposit(_depositId *big.Int) (struct {
	User     common.Address
	Amount   *big.Int
	LockTill *big.Int
}, error) {
	return _RelayerPool.Contract.GetDeposit(&_RelayerPool.CallOpts, _depositId)
}

// LastHarvestRewardTimestamp is a free data retrieval call binding the contract method 0x061887b3.
//
// Solidity: function lastHarvestRewardTimestamp() view returns(uint256)
func (_RelayerPool *RelayerPoolCaller) LastHarvestRewardTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RelayerPool.contract.Call(opts, &out, "lastHarvestRewardTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastHarvestRewardTimestamp is a free data retrieval call binding the contract method 0x061887b3.
//
// Solidity: function lastHarvestRewardTimestamp() view returns(uint256)
func (_RelayerPool *RelayerPoolSession) LastHarvestRewardTimestamp() (*big.Int, error) {
	return _RelayerPool.Contract.LastHarvestRewardTimestamp(&_RelayerPool.CallOpts)
}

// LastHarvestRewardTimestamp is a free data retrieval call binding the contract method 0x061887b3.
//
// Solidity: function lastHarvestRewardTimestamp() view returns(uint256)
func (_RelayerPool *RelayerPoolCallerSession) LastHarvestRewardTimestamp() (*big.Int, error) {
	return _RelayerPool.Contract.LastHarvestRewardTimestamp(&_RelayerPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RelayerPool *RelayerPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RelayerPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RelayerPool *RelayerPoolSession) Owner() (common.Address, error) {
	return _RelayerPool.Contract.Owner(&_RelayerPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RelayerPool *RelayerPoolCallerSession) Owner() (common.Address, error) {
	return _RelayerPool.Contract.Owner(&_RelayerPool.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_RelayerPool *RelayerPoolCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RelayerPool.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_RelayerPool *RelayerPoolSession) Registry() (common.Address, error) {
	return _RelayerPool.Contract.Registry(&_RelayerPool.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_RelayerPool *RelayerPoolCallerSession) Registry() (common.Address, error) {
	return _RelayerPool.Contract.Registry(&_RelayerPool.CallOpts)
}

// RelayerFeeNumerator is a free data retrieval call binding the contract method 0xe61a9582.
//
// Solidity: function relayerFeeNumerator() view returns(uint256)
func (_RelayerPool *RelayerPoolCaller) RelayerFeeNumerator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RelayerPool.contract.Call(opts, &out, "relayerFeeNumerator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RelayerFeeNumerator is a free data retrieval call binding the contract method 0xe61a9582.
//
// Solidity: function relayerFeeNumerator() view returns(uint256)
func (_RelayerPool *RelayerPoolSession) RelayerFeeNumerator() (*big.Int, error) {
	return _RelayerPool.Contract.RelayerFeeNumerator(&_RelayerPool.CallOpts)
}

// RelayerFeeNumerator is a free data retrieval call binding the contract method 0xe61a9582.
//
// Solidity: function relayerFeeNumerator() view returns(uint256)
func (_RelayerPool *RelayerPoolCallerSession) RelayerFeeNumerator() (*big.Int, error) {
	return _RelayerPool.Contract.RelayerFeeNumerator(&_RelayerPool.CallOpts)
}

// RelayerStatus is a free data retrieval call binding the contract method 0x1b49a926.
//
// Solidity: function relayerStatus() view returns(uint8)
func (_RelayerPool *RelayerPoolCaller) RelayerStatus(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RelayerPool.contract.Call(opts, &out, "relayerStatus")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// RelayerStatus is a free data retrieval call binding the contract method 0x1b49a926.
//
// Solidity: function relayerStatus() view returns(uint8)
func (_RelayerPool *RelayerPoolSession) RelayerStatus() (uint8, error) {
	return _RelayerPool.Contract.RelayerStatus(&_RelayerPool.CallOpts)
}

// RelayerStatus is a free data retrieval call binding the contract method 0x1b49a926.
//
// Solidity: function relayerStatus() view returns(uint8)
func (_RelayerPool *RelayerPoolCallerSession) RelayerStatus() (uint8, error) {
	return _RelayerPool.Contract.RelayerStatus(&_RelayerPool.CallOpts)
}

// RewardPerTokenNumerator is a free data retrieval call binding the contract method 0xd1239f37.
//
// Solidity: function rewardPerTokenNumerator() view returns(uint256)
func (_RelayerPool *RelayerPoolCaller) RewardPerTokenNumerator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RelayerPool.contract.Call(opts, &out, "rewardPerTokenNumerator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerTokenNumerator is a free data retrieval call binding the contract method 0xd1239f37.
//
// Solidity: function rewardPerTokenNumerator() view returns(uint256)
func (_RelayerPool *RelayerPoolSession) RewardPerTokenNumerator() (*big.Int, error) {
	return _RelayerPool.Contract.RewardPerTokenNumerator(&_RelayerPool.CallOpts)
}

// RewardPerTokenNumerator is a free data retrieval call binding the contract method 0xd1239f37.
//
// Solidity: function rewardPerTokenNumerator() view returns(uint256)
func (_RelayerPool *RelayerPoolCallerSession) RewardPerTokenNumerator() (*big.Int, error) {
	return _RelayerPool.Contract.RewardPerTokenNumerator(&_RelayerPool.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_RelayerPool *RelayerPoolCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RelayerPool.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_RelayerPool *RelayerPoolSession) RewardToken() (common.Address, error) {
	return _RelayerPool.Contract.RewardToken(&_RelayerPool.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_RelayerPool *RelayerPoolCallerSession) RewardToken() (common.Address, error) {
	return _RelayerPool.Contract.RewardToken(&_RelayerPool.CallOpts)
}

// TotalDeposit is a free data retrieval call binding the contract method 0xf6153ccd.
//
// Solidity: function totalDeposit() view returns(uint256)
func (_RelayerPool *RelayerPoolCaller) TotalDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RelayerPool.contract.Call(opts, &out, "totalDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDeposit is a free data retrieval call binding the contract method 0xf6153ccd.
//
// Solidity: function totalDeposit() view returns(uint256)
func (_RelayerPool *RelayerPoolSession) TotalDeposit() (*big.Int, error) {
	return _RelayerPool.Contract.TotalDeposit(&_RelayerPool.CallOpts)
}

// TotalDeposit is a free data retrieval call binding the contract method 0xf6153ccd.
//
// Solidity: function totalDeposit() view returns(uint256)
func (_RelayerPool *RelayerPoolCallerSession) TotalDeposit() (*big.Int, error) {
	return _RelayerPool.Contract.TotalDeposit(&_RelayerPool.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_RelayerPool *RelayerPoolCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RelayerPool.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_RelayerPool *RelayerPoolSession) Vault() (common.Address, error) {
	return _RelayerPool.Contract.Vault(&_RelayerPool.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_RelayerPool *RelayerPoolCallerSession) Vault() (common.Address, error) {
	return _RelayerPool.Contract.Vault(&_RelayerPool.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_RelayerPool *RelayerPoolTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _RelayerPool.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_RelayerPool *RelayerPoolSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _RelayerPool.Contract.Deposit(&_RelayerPool.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_RelayerPool *RelayerPoolTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _RelayerPool.Contract.Deposit(&_RelayerPool.TransactOpts, _amount)
}

// HarvestMyReward is a paid mutator transaction binding the contract method 0x8363deb6.
//
// Solidity: function harvestMyReward() returns()
func (_RelayerPool *RelayerPoolTransactor) HarvestMyReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayerPool.contract.Transact(opts, "harvestMyReward")
}

// HarvestMyReward is a paid mutator transaction binding the contract method 0x8363deb6.
//
// Solidity: function harvestMyReward() returns()
func (_RelayerPool *RelayerPoolSession) HarvestMyReward() (*types.Transaction, error) {
	return _RelayerPool.Contract.HarvestMyReward(&_RelayerPool.TransactOpts)
}

// HarvestMyReward is a paid mutator transaction binding the contract method 0x8363deb6.
//
// Solidity: function harvestMyReward() returns()
func (_RelayerPool *RelayerPoolTransactorSession) HarvestMyReward() (*types.Transaction, error) {
	return _RelayerPool.Contract.HarvestMyReward(&_RelayerPool.TransactOpts)
}

// HarvestPoolReward is a paid mutator transaction binding the contract method 0x8be9d5bf.
//
// Solidity: function harvestPoolReward() returns()
func (_RelayerPool *RelayerPoolTransactor) HarvestPoolReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayerPool.contract.Transact(opts, "harvestPoolReward")
}

// HarvestPoolReward is a paid mutator transaction binding the contract method 0x8be9d5bf.
//
// Solidity: function harvestPoolReward() returns()
func (_RelayerPool *RelayerPoolSession) HarvestPoolReward() (*types.Transaction, error) {
	return _RelayerPool.Contract.HarvestPoolReward(&_RelayerPool.TransactOpts)
}

// HarvestPoolReward is a paid mutator transaction binding the contract method 0x8be9d5bf.
//
// Solidity: function harvestPoolReward() returns()
func (_RelayerPool *RelayerPoolTransactorSession) HarvestPoolReward() (*types.Transaction, error) {
	return _RelayerPool.Contract.HarvestPoolReward(&_RelayerPool.TransactOpts)
}

// SetEmissionAnnualRateNumerator is a paid mutator transaction binding the contract method 0xec51364e.
//
// Solidity: function setEmissionAnnualRateNumerator(uint256 _value) returns()
func (_RelayerPool *RelayerPoolTransactor) SetEmissionAnnualRateNumerator(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _RelayerPool.contract.Transact(opts, "setEmissionAnnualRateNumerator", _value)
}

// SetEmissionAnnualRateNumerator is a paid mutator transaction binding the contract method 0xec51364e.
//
// Solidity: function setEmissionAnnualRateNumerator(uint256 _value) returns()
func (_RelayerPool *RelayerPoolSession) SetEmissionAnnualRateNumerator(_value *big.Int) (*types.Transaction, error) {
	return _RelayerPool.Contract.SetEmissionAnnualRateNumerator(&_RelayerPool.TransactOpts, _value)
}

// SetEmissionAnnualRateNumerator is a paid mutator transaction binding the contract method 0xec51364e.
//
// Solidity: function setEmissionAnnualRateNumerator(uint256 _value) returns()
func (_RelayerPool *RelayerPoolTransactorSession) SetEmissionAnnualRateNumerator(_value *big.Int) (*types.Transaction, error) {
	return _RelayerPool.Contract.SetEmissionAnnualRateNumerator(&_RelayerPool.TransactOpts, _value)
}

// SetRelayerFeeNumerator is a paid mutator transaction binding the contract method 0x959ccfcb.
//
// Solidity: function setRelayerFeeNumerator(uint256 _value) returns()
func (_RelayerPool *RelayerPoolTransactor) SetRelayerFeeNumerator(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _RelayerPool.contract.Transact(opts, "setRelayerFeeNumerator", _value)
}

// SetRelayerFeeNumerator is a paid mutator transaction binding the contract method 0x959ccfcb.
//
// Solidity: function setRelayerFeeNumerator(uint256 _value) returns()
func (_RelayerPool *RelayerPoolSession) SetRelayerFeeNumerator(_value *big.Int) (*types.Transaction, error) {
	return _RelayerPool.Contract.SetRelayerFeeNumerator(&_RelayerPool.TransactOpts, _value)
}

// SetRelayerFeeNumerator is a paid mutator transaction binding the contract method 0x959ccfcb.
//
// Solidity: function setRelayerFeeNumerator(uint256 _value) returns()
func (_RelayerPool *RelayerPoolTransactorSession) SetRelayerFeeNumerator(_value *big.Int) (*types.Transaction, error) {
	return _RelayerPool.Contract.SetRelayerFeeNumerator(&_RelayerPool.TransactOpts, _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _depositId, uint256 _amount) returns()
func (_RelayerPool *RelayerPoolTransactor) Withdraw(opts *bind.TransactOpts, _depositId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _RelayerPool.contract.Transact(opts, "withdraw", _depositId, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _depositId, uint256 _amount) returns()
func (_RelayerPool *RelayerPoolSession) Withdraw(_depositId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _RelayerPool.Contract.Withdraw(&_RelayerPool.TransactOpts, _depositId, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _depositId, uint256 _amount) returns()
func (_RelayerPool *RelayerPoolTransactorSession) Withdraw(_depositId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _RelayerPool.Contract.Withdraw(&_RelayerPool.TransactOpts, _depositId, _amount)
}

// RelayerPoolDepositPutIterator is returned from FilterDepositPut and is used to iterate over the raw logs and unpacked data for DepositPut events raised by the RelayerPool contract.
type RelayerPoolDepositPutIterator struct {
	Event *RelayerPoolDepositPut // Event containing the contract specifics and raw log

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
func (it *RelayerPoolDepositPutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerPoolDepositPut)
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
		it.Event = new(RelayerPoolDepositPut)
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
func (it *RelayerPoolDepositPutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerPoolDepositPutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerPoolDepositPut represents a DepositPut event raised by the RelayerPool contract.
type RelayerPoolDepositPut struct {
	User     common.Address
	Id       *big.Int
	Amount   *big.Int
	LockTill *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDepositPut is a free log retrieval operation binding the contract event 0x7bacd633fd842a77f4647e920fa8675614c879eb3267a371ef7208363c01cdee.
//
// Solidity: event DepositPut(address indexed user, uint256 indexed id, uint256 amount, uint256 lockTill)
func (_RelayerPool *RelayerPoolFilterer) FilterDepositPut(opts *bind.FilterOpts, user []common.Address, id []*big.Int) (*RelayerPoolDepositPutIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RelayerPool.contract.FilterLogs(opts, "DepositPut", userRule, idRule)
	if err != nil {
		return nil, err
	}
	return &RelayerPoolDepositPutIterator{contract: _RelayerPool.contract, event: "DepositPut", logs: logs, sub: sub}, nil
}

// WatchDepositPut is a free log subscription operation binding the contract event 0x7bacd633fd842a77f4647e920fa8675614c879eb3267a371ef7208363c01cdee.
//
// Solidity: event DepositPut(address indexed user, uint256 indexed id, uint256 amount, uint256 lockTill)
func (_RelayerPool *RelayerPoolFilterer) WatchDepositPut(opts *bind.WatchOpts, sink chan<- *RelayerPoolDepositPut, user []common.Address, id []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RelayerPool.contract.WatchLogs(opts, "DepositPut", userRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerPoolDepositPut)
				if err := _RelayerPool.contract.UnpackLog(event, "DepositPut", log); err != nil {
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

// ParseDepositPut is a log parse operation binding the contract event 0x7bacd633fd842a77f4647e920fa8675614c879eb3267a371ef7208363c01cdee.
//
// Solidity: event DepositPut(address indexed user, uint256 indexed id, uint256 amount, uint256 lockTill)
func (_RelayerPool *RelayerPoolFilterer) ParseDepositPut(log types.Log) (*RelayerPoolDepositPut, error) {
	event := new(RelayerPoolDepositPut)
	if err := _RelayerPool.contract.UnpackLog(event, "DepositPut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelayerPoolDepositWithdrawnIterator is returned from FilterDepositWithdrawn and is used to iterate over the raw logs and unpacked data for DepositWithdrawn events raised by the RelayerPool contract.
type RelayerPoolDepositWithdrawnIterator struct {
	Event *RelayerPoolDepositWithdrawn // Event containing the contract specifics and raw log

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
func (it *RelayerPoolDepositWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerPoolDepositWithdrawn)
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
		it.Event = new(RelayerPoolDepositWithdrawn)
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
func (it *RelayerPoolDepositWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerPoolDepositWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerPoolDepositWithdrawn represents a DepositWithdrawn event raised by the RelayerPool contract.
type RelayerPoolDepositWithdrawn struct {
	User   common.Address
	Id     *big.Int
	Amount *big.Int
	Rest   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDepositWithdrawn is a free log retrieval operation binding the contract event 0x4959c2f9b36bf25a8c0c413e521adc24f9eeaacd05f94c0eee57ace9bf8b4634.
//
// Solidity: event DepositWithdrawn(address indexed user, uint256 indexed id, uint256 amount, uint256 rest)
func (_RelayerPool *RelayerPoolFilterer) FilterDepositWithdrawn(opts *bind.FilterOpts, user []common.Address, id []*big.Int) (*RelayerPoolDepositWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RelayerPool.contract.FilterLogs(opts, "DepositWithdrawn", userRule, idRule)
	if err != nil {
		return nil, err
	}
	return &RelayerPoolDepositWithdrawnIterator{contract: _RelayerPool.contract, event: "DepositWithdrawn", logs: logs, sub: sub}, nil
}

// WatchDepositWithdrawn is a free log subscription operation binding the contract event 0x4959c2f9b36bf25a8c0c413e521adc24f9eeaacd05f94c0eee57ace9bf8b4634.
//
// Solidity: event DepositWithdrawn(address indexed user, uint256 indexed id, uint256 amount, uint256 rest)
func (_RelayerPool *RelayerPoolFilterer) WatchDepositWithdrawn(opts *bind.WatchOpts, sink chan<- *RelayerPoolDepositWithdrawn, user []common.Address, id []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RelayerPool.contract.WatchLogs(opts, "DepositWithdrawn", userRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerPoolDepositWithdrawn)
				if err := _RelayerPool.contract.UnpackLog(event, "DepositWithdrawn", log); err != nil {
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

// ParseDepositWithdrawn is a log parse operation binding the contract event 0x4959c2f9b36bf25a8c0c413e521adc24f9eeaacd05f94c0eee57ace9bf8b4634.
//
// Solidity: event DepositWithdrawn(address indexed user, uint256 indexed id, uint256 amount, uint256 rest)
func (_RelayerPool *RelayerPoolFilterer) ParseDepositWithdrawn(log types.Log) (*RelayerPoolDepositWithdrawn, error) {
	event := new(RelayerPoolDepositWithdrawn)
	if err := _RelayerPool.contract.UnpackLog(event, "DepositWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelayerPoolEmissionAnnualRateNumeratorSetIterator is returned from FilterEmissionAnnualRateNumeratorSet and is used to iterate over the raw logs and unpacked data for EmissionAnnualRateNumeratorSet events raised by the RelayerPool contract.
type RelayerPoolEmissionAnnualRateNumeratorSetIterator struct {
	Event *RelayerPoolEmissionAnnualRateNumeratorSet // Event containing the contract specifics and raw log

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
func (it *RelayerPoolEmissionAnnualRateNumeratorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerPoolEmissionAnnualRateNumeratorSet)
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
		it.Event = new(RelayerPoolEmissionAnnualRateNumeratorSet)
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
func (it *RelayerPoolEmissionAnnualRateNumeratorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerPoolEmissionAnnualRateNumeratorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerPoolEmissionAnnualRateNumeratorSet represents a EmissionAnnualRateNumeratorSet event raised by the RelayerPool contract.
type RelayerPoolEmissionAnnualRateNumeratorSet struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmissionAnnualRateNumeratorSet is a free log retrieval operation binding the contract event 0x4956113bd03fef454b6d6c1c9525fc977eb4f4fb1d01f45c85ea86ecaa46b887.
//
// Solidity: event EmissionAnnualRateNumeratorSet(address indexed sender, uint256 value)
func (_RelayerPool *RelayerPoolFilterer) FilterEmissionAnnualRateNumeratorSet(opts *bind.FilterOpts, sender []common.Address) (*RelayerPoolEmissionAnnualRateNumeratorSetIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RelayerPool.contract.FilterLogs(opts, "EmissionAnnualRateNumeratorSet", senderRule)
	if err != nil {
		return nil, err
	}
	return &RelayerPoolEmissionAnnualRateNumeratorSetIterator{contract: _RelayerPool.contract, event: "EmissionAnnualRateNumeratorSet", logs: logs, sub: sub}, nil
}

// WatchEmissionAnnualRateNumeratorSet is a free log subscription operation binding the contract event 0x4956113bd03fef454b6d6c1c9525fc977eb4f4fb1d01f45c85ea86ecaa46b887.
//
// Solidity: event EmissionAnnualRateNumeratorSet(address indexed sender, uint256 value)
func (_RelayerPool *RelayerPoolFilterer) WatchEmissionAnnualRateNumeratorSet(opts *bind.WatchOpts, sink chan<- *RelayerPoolEmissionAnnualRateNumeratorSet, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RelayerPool.contract.WatchLogs(opts, "EmissionAnnualRateNumeratorSet", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerPoolEmissionAnnualRateNumeratorSet)
				if err := _RelayerPool.contract.UnpackLog(event, "EmissionAnnualRateNumeratorSet", log); err != nil {
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

// ParseEmissionAnnualRateNumeratorSet is a log parse operation binding the contract event 0x4956113bd03fef454b6d6c1c9525fc977eb4f4fb1d01f45c85ea86ecaa46b887.
//
// Solidity: event EmissionAnnualRateNumeratorSet(address indexed sender, uint256 value)
func (_RelayerPool *RelayerPoolFilterer) ParseEmissionAnnualRateNumeratorSet(log types.Log) (*RelayerPoolEmissionAnnualRateNumeratorSet, error) {
	event := new(RelayerPoolEmissionAnnualRateNumeratorSet)
	if err := _RelayerPool.contract.UnpackLog(event, "EmissionAnnualRateNumeratorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelayerPoolHarvestPoolRewardIterator is returned from FilterHarvestPoolReward and is used to iterate over the raw logs and unpacked data for HarvestPoolReward events raised by the RelayerPool contract.
type RelayerPoolHarvestPoolRewardIterator struct {
	Event *RelayerPoolHarvestPoolReward // Event containing the contract specifics and raw log

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
func (it *RelayerPoolHarvestPoolRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerPoolHarvestPoolReward)
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
		it.Event = new(RelayerPoolHarvestPoolReward)
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
func (it *RelayerPoolHarvestPoolRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerPoolHarvestPoolRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerPoolHarvestPoolReward represents a HarvestPoolReward event raised by the RelayerPool contract.
type RelayerPoolHarvestPoolReward struct {
	Sender                        common.Address
	HarvestForPeriod              *big.Int
	Profit                        *big.Int
	FeeReceiver                   common.Address
	Fee                           *big.Int
	RewardForPool                 *big.Int
	RewardPerTokenNumeratorBefore *big.Int
	RewardPerTokenNumerator       *big.Int
	TotalDeposit                  *big.Int
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterHarvestPoolReward is a free log retrieval operation binding the contract event 0xadd1c21c3f7b25aa60bceee6e16a6f94dfd90fbb5809c95e74ffc682c3cfc6a7.
//
// Solidity: event HarvestPoolReward(address indexed sender, uint256 harvestForPeriod, uint256 profit, address indexed feeReceiver, uint256 fee, uint256 rewardForPool, uint256 rewardPerTokenNumeratorBefore, uint256 rewardPerTokenNumerator, uint256 totalDeposit)
func (_RelayerPool *RelayerPoolFilterer) FilterHarvestPoolReward(opts *bind.FilterOpts, sender []common.Address, feeReceiver []common.Address) (*RelayerPoolHarvestPoolRewardIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var feeReceiverRule []interface{}
	for _, feeReceiverItem := range feeReceiver {
		feeReceiverRule = append(feeReceiverRule, feeReceiverItem)
	}

	logs, sub, err := _RelayerPool.contract.FilterLogs(opts, "HarvestPoolReward", senderRule, feeReceiverRule)
	if err != nil {
		return nil, err
	}
	return &RelayerPoolHarvestPoolRewardIterator{contract: _RelayerPool.contract, event: "HarvestPoolReward", logs: logs, sub: sub}, nil
}

// WatchHarvestPoolReward is a free log subscription operation binding the contract event 0xadd1c21c3f7b25aa60bceee6e16a6f94dfd90fbb5809c95e74ffc682c3cfc6a7.
//
// Solidity: event HarvestPoolReward(address indexed sender, uint256 harvestForPeriod, uint256 profit, address indexed feeReceiver, uint256 fee, uint256 rewardForPool, uint256 rewardPerTokenNumeratorBefore, uint256 rewardPerTokenNumerator, uint256 totalDeposit)
func (_RelayerPool *RelayerPoolFilterer) WatchHarvestPoolReward(opts *bind.WatchOpts, sink chan<- *RelayerPoolHarvestPoolReward, sender []common.Address, feeReceiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var feeReceiverRule []interface{}
	for _, feeReceiverItem := range feeReceiver {
		feeReceiverRule = append(feeReceiverRule, feeReceiverItem)
	}

	logs, sub, err := _RelayerPool.contract.WatchLogs(opts, "HarvestPoolReward", senderRule, feeReceiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerPoolHarvestPoolReward)
				if err := _RelayerPool.contract.UnpackLog(event, "HarvestPoolReward", log); err != nil {
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

// ParseHarvestPoolReward is a log parse operation binding the contract event 0xadd1c21c3f7b25aa60bceee6e16a6f94dfd90fbb5809c95e74ffc682c3cfc6a7.
//
// Solidity: event HarvestPoolReward(address indexed sender, uint256 harvestForPeriod, uint256 profit, address indexed feeReceiver, uint256 fee, uint256 rewardForPool, uint256 rewardPerTokenNumeratorBefore, uint256 rewardPerTokenNumerator, uint256 totalDeposit)
func (_RelayerPool *RelayerPoolFilterer) ParseHarvestPoolReward(log types.Log) (*RelayerPoolHarvestPoolReward, error) {
	event := new(RelayerPoolHarvestPoolReward)
	if err := _RelayerPool.contract.UnpackLog(event, "HarvestPoolReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelayerPoolRelayerFeeNumeratorSetIterator is returned from FilterRelayerFeeNumeratorSet and is used to iterate over the raw logs and unpacked data for RelayerFeeNumeratorSet events raised by the RelayerPool contract.
type RelayerPoolRelayerFeeNumeratorSetIterator struct {
	Event *RelayerPoolRelayerFeeNumeratorSet // Event containing the contract specifics and raw log

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
func (it *RelayerPoolRelayerFeeNumeratorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerPoolRelayerFeeNumeratorSet)
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
		it.Event = new(RelayerPoolRelayerFeeNumeratorSet)
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
func (it *RelayerPoolRelayerFeeNumeratorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerPoolRelayerFeeNumeratorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerPoolRelayerFeeNumeratorSet represents a RelayerFeeNumeratorSet event raised by the RelayerPool contract.
type RelayerPoolRelayerFeeNumeratorSet struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRelayerFeeNumeratorSet is a free log retrieval operation binding the contract event 0x673d70d178eaf2ea132ef62a70f12374909ee3364342c07575e53a4625e51185.
//
// Solidity: event RelayerFeeNumeratorSet(address indexed sender, uint256 value)
func (_RelayerPool *RelayerPoolFilterer) FilterRelayerFeeNumeratorSet(opts *bind.FilterOpts, sender []common.Address) (*RelayerPoolRelayerFeeNumeratorSetIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RelayerPool.contract.FilterLogs(opts, "RelayerFeeNumeratorSet", senderRule)
	if err != nil {
		return nil, err
	}
	return &RelayerPoolRelayerFeeNumeratorSetIterator{contract: _RelayerPool.contract, event: "RelayerFeeNumeratorSet", logs: logs, sub: sub}, nil
}

// WatchRelayerFeeNumeratorSet is a free log subscription operation binding the contract event 0x673d70d178eaf2ea132ef62a70f12374909ee3364342c07575e53a4625e51185.
//
// Solidity: event RelayerFeeNumeratorSet(address indexed sender, uint256 value)
func (_RelayerPool *RelayerPoolFilterer) WatchRelayerFeeNumeratorSet(opts *bind.WatchOpts, sink chan<- *RelayerPoolRelayerFeeNumeratorSet, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RelayerPool.contract.WatchLogs(opts, "RelayerFeeNumeratorSet", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerPoolRelayerFeeNumeratorSet)
				if err := _RelayerPool.contract.UnpackLog(event, "RelayerFeeNumeratorSet", log); err != nil {
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

// ParseRelayerFeeNumeratorSet is a log parse operation binding the contract event 0x673d70d178eaf2ea132ef62a70f12374909ee3364342c07575e53a4625e51185.
//
// Solidity: event RelayerFeeNumeratorSet(address indexed sender, uint256 value)
func (_RelayerPool *RelayerPoolFilterer) ParseRelayerFeeNumeratorSet(log types.Log) (*RelayerPoolRelayerFeeNumeratorSet, error) {
	event := new(RelayerPoolRelayerFeeNumeratorSet)
	if err := _RelayerPool.contract.UnpackLog(event, "RelayerFeeNumeratorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelayerPoolRelayerStatusSetIterator is returned from FilterRelayerStatusSet and is used to iterate over the raw logs and unpacked data for RelayerStatusSet events raised by the RelayerPool contract.
type RelayerPoolRelayerStatusSetIterator struct {
	Event *RelayerPoolRelayerStatusSet // Event containing the contract specifics and raw log

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
func (it *RelayerPoolRelayerStatusSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerPoolRelayerStatusSet)
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
		it.Event = new(RelayerPoolRelayerStatusSet)
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
func (it *RelayerPoolRelayerStatusSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerPoolRelayerStatusSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerPoolRelayerStatusSet represents a RelayerStatusSet event raised by the RelayerPool contract.
type RelayerPoolRelayerStatusSet struct {
	Sender common.Address
	Status uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRelayerStatusSet is a free log retrieval operation binding the contract event 0x22f527cc6aaeb89226980b7a4e1051e149dc4d11f1d1e957968c437998ccc3a8.
//
// Solidity: event RelayerStatusSet(address indexed sender, uint8 status)
func (_RelayerPool *RelayerPoolFilterer) FilterRelayerStatusSet(opts *bind.FilterOpts, sender []common.Address) (*RelayerPoolRelayerStatusSetIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RelayerPool.contract.FilterLogs(opts, "RelayerStatusSet", senderRule)
	if err != nil {
		return nil, err
	}
	return &RelayerPoolRelayerStatusSetIterator{contract: _RelayerPool.contract, event: "RelayerStatusSet", logs: logs, sub: sub}, nil
}

// WatchRelayerStatusSet is a free log subscription operation binding the contract event 0x22f527cc6aaeb89226980b7a4e1051e149dc4d11f1d1e957968c437998ccc3a8.
//
// Solidity: event RelayerStatusSet(address indexed sender, uint8 status)
func (_RelayerPool *RelayerPoolFilterer) WatchRelayerStatusSet(opts *bind.WatchOpts, sink chan<- *RelayerPoolRelayerStatusSet, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RelayerPool.contract.WatchLogs(opts, "RelayerStatusSet", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerPoolRelayerStatusSet)
				if err := _RelayerPool.contract.UnpackLog(event, "RelayerStatusSet", log); err != nil {
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

// ParseRelayerStatusSet is a log parse operation binding the contract event 0x22f527cc6aaeb89226980b7a4e1051e149dc4d11f1d1e957968c437998ccc3a8.
//
// Solidity: event RelayerStatusSet(address indexed sender, uint8 status)
func (_RelayerPool *RelayerPoolFilterer) ParseRelayerStatusSet(log types.Log) (*RelayerPoolRelayerStatusSet, error) {
	event := new(RelayerPoolRelayerStatusSet)
	if err := _RelayerPool.contract.UnpackLog(event, "RelayerStatusSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelayerPoolUserHarvestRewardIterator is returned from FilterUserHarvestReward and is used to iterate over the raw logs and unpacked data for UserHarvestReward events raised by the RelayerPool contract.
type RelayerPoolUserHarvestRewardIterator struct {
	Event *RelayerPoolUserHarvestReward // Event containing the contract specifics and raw log

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
func (it *RelayerPoolUserHarvestRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerPoolUserHarvestReward)
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
		it.Event = new(RelayerPoolUserHarvestReward)
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
func (it *RelayerPoolUserHarvestRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerPoolUserHarvestRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerPoolUserHarvestReward represents a UserHarvestReward event raised by the RelayerPool contract.
type RelayerPoolUserHarvestReward struct {
	User        common.Address
	UserReward  *big.Int
	UserDeposit *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUserHarvestReward is a free log retrieval operation binding the contract event 0xbeb525c2d0c03a594fed4cea83f02931246206387586fb3704201174b5ea8b71.
//
// Solidity: event UserHarvestReward(address indexed user, uint256 userReward, uint256 userDeposit)
func (_RelayerPool *RelayerPoolFilterer) FilterUserHarvestReward(opts *bind.FilterOpts, user []common.Address) (*RelayerPoolUserHarvestRewardIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _RelayerPool.contract.FilterLogs(opts, "UserHarvestReward", userRule)
	if err != nil {
		return nil, err
	}
	return &RelayerPoolUserHarvestRewardIterator{contract: _RelayerPool.contract, event: "UserHarvestReward", logs: logs, sub: sub}, nil
}

// WatchUserHarvestReward is a free log subscription operation binding the contract event 0xbeb525c2d0c03a594fed4cea83f02931246206387586fb3704201174b5ea8b71.
//
// Solidity: event UserHarvestReward(address indexed user, uint256 userReward, uint256 userDeposit)
func (_RelayerPool *RelayerPoolFilterer) WatchUserHarvestReward(opts *bind.WatchOpts, sink chan<- *RelayerPoolUserHarvestReward, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _RelayerPool.contract.WatchLogs(opts, "UserHarvestReward", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerPoolUserHarvestReward)
				if err := _RelayerPool.contract.UnpackLog(event, "UserHarvestReward", log); err != nil {
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

// ParseUserHarvestReward is a log parse operation binding the contract event 0xbeb525c2d0c03a594fed4cea83f02931246206387586fb3704201174b5ea8b71.
//
// Solidity: event UserHarvestReward(address indexed user, uint256 userReward, uint256 userDeposit)
func (_RelayerPool *RelayerPoolFilterer) ParseUserHarvestReward(log types.Log) (*RelayerPoolUserHarvestReward, error) {
	event := new(RelayerPoolUserHarvestReward)
	if err := _RelayerPool.contract.UnpackLog(event, "UserHarvestReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
