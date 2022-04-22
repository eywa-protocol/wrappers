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

// RouterMetaData contains all meta data concerning the Router contract.

var RouterMetaData = &bind.MetaData{

	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"portal\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"synthesis\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"curveProxy\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userFrom\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionPrice\",\"type\":\"uint256\"}],\"name\":\"CrosschainPaymentEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_SYNTHESIZE_REQUEST_SIGNATURE_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_SYNTH_TRANSFER_REQUEST_SIGNATURE_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_UNSYNTHESIZE_REQUEST_SIGNATURE_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_trustedWorker\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIPortal.TransitData\",\"name\":\"transitData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structIPortal.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structIPortal.PermitData[]\",\"name\":\"permitData\",\"type\":\"tuple[]\"}],\"name\":\"batchSynthesizeRequestWithDataTransit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"emergencyUnburnRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"emergencyUnsyntesizeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remove\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinMintAmount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinDy\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"chain2address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"initialBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialChainID\",\"type\":\"uint256\"}],\"internalType\":\"structICurveProxy.MetaExchangeParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structICurveProxy.PermitData[]\",\"name\":\"permit\",\"type\":\"tuple[]\"},{\"internalType\":\"address[3]\",\"name\":\"token\",\"type\":\"address[3]\"},{\"internalType\":\"uint256[3]\",\"name\":\"amount\",\"type\":\"uint256[3]\"}],\"name\":\"metaExchangeRequestVia3pool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addAtCrosschainPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinMintAmountC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lpIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addAtHubPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinMintAmountH\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialChainID\",\"type\":\"uint256\"}],\"internalType\":\"structICurveProxy.MetaMintEUSD\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structICurveProxy.PermitData[]\",\"name\":\"permit\",\"type\":\"tuple[]\"},{\"internalType\":\"address[3]\",\"name\":\"token\",\"type\":\"address[3]\"},{\"internalType\":\"uint256[3]\",\"name\":\"amount\",\"type\":\"uint256[3]\"}],\"name\":\"mintEusdRequestVia3pool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"removeAtCrosschainPool\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinAmountC\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"removeAtHubPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountH\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"y\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinAmountH\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"internalType\":\"structICurveProxy.MetaRedeemEUSD\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structICurveProxy.PermitData\",\"name\":\"permit\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"redeemEusdRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"}],\"name\":\"removeTrustedWorker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"}],\"name\":\"setTrustedWorker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"rtoken\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structISynthesis.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"}],\"name\":\"synthTransferRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tokenReal\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenSynth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structISynthesis.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"executionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structRouter.DelegatedCallReceipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"synthTransferRequestPayNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structIPortal.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"executionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structRouter.DelegatedCallReceipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"synthesizeRequestPayNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structIPortal.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structIPortal.PermitData\",\"name\":\"permitData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"executionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structRouter.DelegatedCallReceipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"synthesizeRequestWithPermitPayNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structIPortal.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"}],\"name\":\"tokenSynthesizeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"pubkeys\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes1\",\"name\":\"txStateBump\",\"type\":\"bytes1\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"tokenSynthesizeRequestToSolana\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structIPortal.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structIPortal.PermitData\",\"name\":\"permitData\",\"type\":\"tuple\"}],\"name\":\"tokenSynthesizeRequestWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stoken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structISynthesis.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"}],\"name\":\"unsynthesizeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenSynth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structISynthesis.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"executionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structRouter.DelegatedCallReceipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"unsynthesizeRequestPayNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stoken\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"pubkeys\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"unsynthesizeRequestToSolana\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101206040523480156200001257600080fd5b506040516200305c3803806200305c8339810160408190526200003591620001a0565b60408051808201825260048152634559574160e01b60208083019182528351808501855260018152603160f81b908201529151902060c08181527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc660e08190524660a081815286517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f8188018190528189019690965260608101939093526080808401929092523083820152865180840390910181529190920194859052805193019290922090915261010052600080546001600160a01b031916339081178255918291907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350600280546001600160a01b039485166001600160a01b031991821617909155600380549385169382169390931790925560018054919093169116179055620001ea565b80516001600160a01b03811681146200019b57600080fd5b919050565b600080600060608486031215620001b657600080fd5b620001c18462000183565b9250620001d16020850162000183565b9150620001e16040850162000183565b90509250925092565b60805160a05160c05160e05161010051612e2d6200022f600039600061154c0152600061158e0152600061156d015260006114fc015260006115250152612e2d6000f3fe60806040526004361061019c5760003560e01c80637fe04792116100ec578063d3c522841161008a578063d9ab976411610064578063d9ab976414610433578063e840f9d314610453578063f2fde38b14610493578063f8900f78146104b357600080fd5b8063d3c52284146103eb578063d550f9e114610400578063d63f79b21461042057600080fd5b8063c3345883116100c6578063c334588314610385578063c51e547c146103a5578063c773abfd146103b8578063c89efd6b146103cb57600080fd5b80637fe047921461031d5780638da5cb5b1461033d5780639500125b1461036557600080fd5b80634d176e841161015957806371056f5e1161013357806371056f5e146102b3578063715018a6146102c857806372b90c02146102dd5780637ecebe00146102fd57600080fd5b80634d176e84146102605780634ec658e21461028057806369ab5d461461029357600080fd5b80630d0e5b6d146101a15780632acbaacf146101c35780632e88ce08146101eb5780633644e5151461020b578063439616681461022057806344c82d2314610240575b600080fd5b3480156101ad57600080fd5b506101c16101bc366004611de2565b6104d3565b005b3480156101cf57600080fd5b506101d86105ca565b6040519081526020015b60405180910390f35b3480156101f757600080fd5b506101c1610206366004611e49565b6105f2565b34801561021757600080fd5b506101d8610682565b34801561022c57600080fd5b506101c161023b366004611ee3565b610691565b34801561024c57600080fd5b506101c161025b366004611f76565b610708565b34801561026c57600080fd5b506101c161027b36600461205c565b61075e565b6101c161028e366004612143565b6109f6565b34801561029f57600080fd5b506101c16102ae3660046121ca565b610b2c565b3480156102bf57600080fd5b506101d8610bf4565b3480156102d457600080fd5b506101c1610c03565b3480156102e957600080fd5b506101c16102f8366004612247565b610c80565b34801561030957600080fd5b506101d8610318366004612247565b610ccb565b34801561032957600080fd5b506101c1610338366004612247565b610ceb565b34801561034957600080fd5b506000546040516001600160a01b0390911681526020016101e2565b34801561037157600080fd5b506101c1610380366004611ee3565b610d39565b34801561039157600080fd5b506101c16103a0366004612264565b610d75565b6101c16103b33660046122ad565b610ea1565b6101c16103c6366004612306565b610f83565b3480156103d757600080fd5b506101c16103e6366004612367565b611065565b3480156103f757600080fd5b506101d86110ef565b34801561040c57600080fd5b506101c161041b366004612367565b6110fe565b6101c161042e3660046122ad565b6111ba565b34801561043f57600080fd5b506101c161044e3660046122ad565b611232565b34801561045f57600080fd5b5061048361046e366004612247565b60046020526000908152604090205460ff1681565b60405190151581526020016101e2565b34801561049f57600080fd5b506101c16104ae366004612247565b611323565b3480156104bf57600080fd5b506101c16104ce36600461239f565b61140d565b6104df85333085611498565b60035460405163095ea7b360e01b81526001600160a01b039182166004820152602481018490529086169063095ea7b3906044016020604051808303816000875af1158015610532573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105569190612425565b5060035460405163966e0fe960e01b81526001600160a01b039091169063966e0fe99061059190889033908990899089908990600401612478565b600060405180830381600087803b1580156105ab57600080fd5b505af11580156105bf573d6000803e3d6000fd5b505050505050505050565b6040516020016105d9906124ba565b6040516020818303038152906040528051906020012081565b60025461060c90879033906001600160a01b031688611498565b600254604051632b17790f60e21b81526001600160a01b039091169063ac5de43c90610648908990899033908a908a908a908a90600401612520565b600060405180830381600087803b15801561066257600080fd5b505af1158015610676573d6000803e3d6000fd5b50505050505050505050565b600061068c6114f8565b905090565b600254604051630872c2cd60e31b81526001600160a01b03909116906343961668906106cd908a908a908a908a908a908a908a90600401612574565b600060405180830381600087803b1580156106e757600080fd5b505af11580156106fb573d6000803e3d6000fd5b5050505050505050505050565b60015461072690859033906001600160a01b031660808a0135611498565b600154604051635fcc4a3f60e01b81526001600160a01b0390911690635fcc4a3f90610648908990899088908890889060040161260c565b60005b8881101561097a57600088888381811061077d5761077d6126ea565b9050602002013511156109685782828281811061079c5761079c6126ea565b6107b292602060a0909202019081019150612700565b60ff1615610912578989828181106107cc576107cc6126ea565b90506020020160208101906107e19190612247565b6001600160a01b031663d505accf3330868686818110610803576108036126ea565b905060a00201608001602081019061081b919061271b565b61083d578b8b86818110610831576108316126ea565b90506020020135610841565b6000195b878787818110610853576108536126ea565b905060a002016060013588888881811061086f5761086f6126ea565b61088592602060a0909202019081019150612700565b898989818110610897576108976126ea565b905060a00201602001358a8a8a8181106108b3576108b36126ea565b905060a00201604001356040518863ffffffff1660e01b81526004016108df9796959493929190612738565b600060405180830381600087803b1580156108f957600080fd5b505af115801561090d573d6000803e3d6000fd5b505050505b6109688a8a83818110610927576109276126ea565b905060200201602081019061093c9190612247565b60025433906001600160a01b03168b8b8681811061095c5761095c6126ea565b90506020020135611498565b8061097281612779565b915050610761565b50600254604051632a3a615b60e01b81526001600160a01b0390911690632a3a615b906109b9908c908c908c908c9033908d908c908e9060040161286f565b600060405180830381600087803b1580156109d357600080fd5b505af11580156109e7573d6000803e3d6000fd5b50505050505050505050505050565b6000610a2b8460400135604051602001610a0f90612911565b60405160208183030381529060405280519060200120846115b2565b90506001600160a01b03871663d505accf3330610a4e60a088016080890161271b565b610a585789610a5c565b6000195b6060880135610a6e60208a018a612700565b89602001358a604001356040518863ffffffff1660e01b8152600401610a9a9796959493929190612738565b600060405180830381600087803b158015610ab457600080fd5b505af1158015610ac8573d6000803e3d6000fd5b50505050610ada826000013582611796565b600254610af490889033906001600160a01b031689611498565b600254604051637aa8486d60e11b81526001600160a01b039091169063f55090da906106cd908a908a9033908b908b90600401612975565b60005b6003811015610bbb576000828260038110610b4c57610b4c6126ea565b60200201351115610ba957610ba9838260038110610b6c57610b6c6126ea565b602002016020810190610b7f9190612247565b60025433906001600160a01b0316858560038110610b9f57610b9f6126ea565b6020020135611498565b80610bb381612779565b915050610b2f565b506001546040516301c2bc7d60e51b81526001600160a01b03909116906338578fa0906105919088908890889088908890600401612a37565b6040516020016105d990612911565b6000546001600160a01b03163314610c365760405162461bcd60e51b8152600401610c2d90612afa565b60405180910390fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b03163314610caa5760405162461bcd60e51b8152600401610c2d90612afa565b6001600160a01b03166000908152600460205260409020805460ff19169055565b6001600160a01b0381166000908152600560205260408120545b92915050565b6000546001600160a01b03163314610d155760405162461bcd60e51b8152600401610c2d90612afa565b6001600160a01b03166000908152600460205260409020805460ff19166001179055565b600354604051639500125b60e01b81526001600160a01b0390911690639500125b906106cd908a908a908a908a908a908a908a90600401612574565b60035460405163bca7382360e01b8152600481018690526000916001600160a01b03169063bca7382390602401602060405180830381865afa158015610dbf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610de39190612b2f565b9050610df181333087611498565b60035460405163095ea7b360e01b81526001600160a01b039182166004820152602481018690529082169063095ea7b3906044016020604051808303816000875af1158015610e44573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e689190612425565b5060035460405163b04640b760e01b81526001600160a01b039091169063b04640b7906105919088908890339089908990600401612b4c565b6000610eba8360400135604051602001610a0f906124ba565b9050610ec7823582611796565b610ed386333088611498565b60035460405163095ea7b360e01b81526001600160a01b039182166004820152602481018790529087169063095ea7b3906044016020604051808303816000875af1158015610f26573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f4a9190612425565b5060035460405163451c763560e01b81526001600160a01b039091169063451c763590610648908990899033908a908a90600401612975565b6000610f9c8360400135604051602001610a0f90612b7e565b9050610fa9823582611796565b610fb586333088611498565b60035460405163095ea7b360e01b81526001600160a01b039182166004820152602481018790529087169063095ea7b3906044016020604051808303816000875af1158015611008573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061102c9190612425565b5060035460405163b04640b760e01b81526001600160a01b039091169063b04640b7906106cd908a90899033908a908a90600401612b4c565b60025461107f90859033906001600160a01b031686611498565b600254604051637aa8486d60e11b81526001600160a01b039091169063f55090da906110b79087908790339088908890600401612975565b600060405180830381600087803b1580156110d157600080fd5b505af11580156110e5573d6000803e3d6000fd5b5050505050505050565b6040516020016105d990612b7e565b61110a84333086611498565b60035460405163095ea7b360e01b81526001600160a01b039182166004820152602481018590529085169063095ea7b3906044016020604051808303816000875af115801561115d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111819190612425565b5060035460405163451c763560e01b81526001600160a01b039091169063451c7635906110b79087908790339088908890600401612975565b60006111d38360400135604051602001610a0f90612911565b90506111e0823582611796565b6002546111fa90879033906001600160a01b031688611498565b600254604051637aa8486d60e11b81526001600160a01b039091169063f55090da90610648908990899033908a908a90600401612975565b6001600160a01b03851663d505accf333061125360a086016080870161271b565b61125d5787611261565b6000195b60608601356112736020880188612700565b876020013588604001356040518863ffffffff1660e01b815260040161129f9796959493929190612738565b600060405180830381600087803b1580156112b957600080fd5b505af11580156112cd573d6000803e3d6000fd5b50506002546112eb925087915033906001600160a01b031687611498565b600254604051637aa8486d60e11b81526001600160a01b039091169063f55090da906105919088908890339089908990600401612975565b6000546001600160a01b0316331461134d5760405162461bcd60e51b8152600401610c2d90612afa565b6001600160a01b0381166113b25760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610c2d565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b60005b600381101561145f57600082826003811061142d5761142d6126ea565b6020020135111561144d5761144d838260038110610b6c57610b6c6126ea565b8061145781612779565b915050611410565b50600154604051634638ef3f60e11b81526001600160a01b0390911690638c71de7e906105919088908890889088908890600401612be5565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b1790526114f29085906118b6565b50505050565b60007f000000000000000000000000000000000000000000000000000000000000000046141561154757507f000000000000000000000000000000000000000000000000000000000000000090565b61068c7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000061198d565b6000806115be336119d7565b604080517fb2cd5b2d9baed0ca8e8805aae5275d69b6a747b2fdb5c114eb96a962fdc1a6996020808301919091523360601b6bffffffffffffffffffffffff19168284015260548201899052863560748301526094820188905260b482018490528681013560d4808401919091528351808403909101815260f4909201909252805191012090915060006116a7611654836119ff565b6040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b90506116cc816116bd6060880160408901612700565b87606001358860800135611a4d565b6001600160a01b03811660009081526004602052604090205490945060ff166117455760405162461bcd60e51b815260206004820152602560248201527f526f757465723a20696e76616c6964207369676e61747572652066726f6d207760448201526437b935b2b960d91b6064820152608401610c2d565b846020013542111561178c5760405162461bcd60e51b815260206004820152601060248201526f526f757465723a20646561646c696e6560801b6044820152606401610c2d565b5050509392505050565b813410156117d75760405162461bcd60e51b815260206004820152600e60248201526d125b9d985b1a5908185b5bdd5b9d60921b6044820152606401610c2d565b6000816001600160a01b03163460405160006040518083038185875af1925050503d8060008114611824576040519150601f19603f3d011682016040523d82523d6000602084013e611829565b606091505b50509050806118715760405162461bcd60e51b81526020600482015260146024820152732330b4b632b2103a379039b2b7321022ba3432b960611b6044820152606401610c2d565b6040518381526001600160a01b0383169033907f141cdd7ab02a0ea7399fa91a81781d764708704497d78344275c700a6caa0be69060200160405180910390a3505050565b600061190b826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316611bf69092919063ffffffff16565b80519091501561198857808060200190518101906119299190612425565b6119885760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610c2d565b505050565b6040805160208101859052908101839052606081018290524660808201523060a082015260009060c0016040516020818303038152906040528051906020012090505b9392505050565b6001600160a01b03811660009081526005602052604090208054600181018255905b50919050565b6000610ce5611a0c6114f8565b8360405161190160f01b6020820152602281018390526042810182905260009060620160405160208183030381529060405280519060200120905092915050565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0821115611aca5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610c2d565b8360ff16601b1480611adf57508360ff16601c145b611b365760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610c2d565b6040805160008082526020820180845288905260ff871692820192909252606081018590526080810184905260019060a0016020604051602081039080840390855afa158015611b8a573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116611bed5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610c2d565b95945050505050565b6060611c058484600085611c0d565b949350505050565b606082471015611c6e5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610c2d565b843b611cbc5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610c2d565b600080866001600160a01b03168587604051611cd89190612da8565b60006040518083038185875af1925050503d8060008114611d15576040519150601f19603f3d011682016040523d82523d6000602084013e611d1a565b606091505b5091509150611d2a828286611d35565b979650505050505050565b60608315611d445750816119d0565b825115611d545782518084602001fd5b8160405162461bcd60e51b8152600401610c2d9190612dc4565b6001600160a01b0381168114611d8357600080fd5b50565b8035611d9181611d6e565b919050565b60008083601f840112611da857600080fd5b50813567ffffffffffffffff811115611dc057600080fd5b6020830191508360208260051b8501011115611ddb57600080fd5b9250929050565b600080600080600060808688031215611dfa57600080fd5b8535611e0581611d6e565b9450602086013567ffffffffffffffff811115611e2157600080fd5b611e2d88828901611d96565b9699909850959660408101359660609091013595509350505050565b60008060008060008060a08789031215611e6257600080fd5b8635611e6d81611d6e565b955060208701359450604087013567ffffffffffffffff811115611e9057600080fd5b611e9c89828a01611d96565b90955093505060608701356001600160f81b031981168114611ebd57600080fd5b80925050608087013590509295509295509295565b803560ff81168114611d9157600080fd5b600080600080600080600060e0888a031215611efe57600080fd5b873596506020880135611f1081611d6e565b95506040880135611f2081611d6e565b945060608801359350611f3560808901611ed2565b925060a0880135915060c0880135905092959891949750929550565b600061010082840312156119f957600080fd5b600060a082840312156119f957600080fd5b6000806000806000806102208789031215611f9057600080fd5b611f9a8888611f51565b9550611faa886101008901611f64565b94506101a0870135611fbb81611d6e565b93506101c0870135611fcc81611d6e565b92506101e0870135611fdd81611d6e565b8092505061020087013590509295509295509295565b6000604082840312156119f957600080fd5b6000606082840312156119f957600080fd5b60008083601f84011261202957600080fd5b50813567ffffffffffffffff81111561204157600080fd5b60208301915083602060a083028501011115611ddb57600080fd5b60008060008060008060008060006101008a8c03121561207b57600080fd5b893567ffffffffffffffff8082111561209357600080fd5b61209f8d838e01611d96565b909b50995060208c01359150808211156120b857600080fd5b6120c48d838e01611d96565b90995097508791506120d860408d01611d86565b965060608c01359150808211156120ee57600080fd5b6120fa8d838e01611ff3565b95506121098d60808e01612005565b945060e08c013591508082111561211f57600080fd5b5061212c8c828d01612017565b915080935050809150509295985092959850929598565b600080600080600080610200878903121561215d57600080fd5b863561216881611d6e565b955060208701359450604087013561217f81611d6e565b935061218e8860608901612005565b925061219d8860c08901611f64565b91506121ad886101608901611f64565b90509295509295509295565b8060608101831015610ce557600080fd5b60008060008060006101e086880312156121e357600080fd5b6121ed8787611f51565b945061010086013567ffffffffffffffff81111561220a57600080fd5b61221688828901612017565b909550935061222b90508761012088016121b9565b915061223b8761018088016121b9565b90509295509295909350565b60006020828403121561225957600080fd5b81356119d081611d6e565b60008060008060c0858703121561227a57600080fd5b8435935060208501359250604085013561229381611d6e565b91506122a28660608701612005565b905092959194509250565b600080600080600061016086880312156122c657600080fd5b85356122d181611d6e565b94506020860135935060408601356122e881611d6e565b92506122f78760608801612005565b915061223b8760c08801611f64565b600080600080600080610180878903121561232057600080fd5b86359550602087013561233281611d6e565b945060408701359350606087013561234981611d6e565b92506123588860808901612005565b91506121ad8860e08901611f64565b60008060008060c0858703121561237d57600080fd5b843561238881611d6e565b935060208501359250604085013561229381611d6e565b60008060008060008587036102e08112156123b957600080fd5b610200808212156123c957600080fd5b879650860135905067ffffffffffffffff8111156123e657600080fd5b6123f288828901612017565b909550935061240790508761022088016121b9565b915061223b8761028088016121b9565b8015158114611d8357600080fd5b60006020828403121561243757600080fd5b81516119d081612417565b81835260006001600160fb1b0383111561245b57600080fd5b8260051b8083602087013760009401602001938452509192915050565b6001600160a01b0387811682528616602082015260a0604082018190526000906124a59083018688612442565b60608301949094525060800152949350505050565b7f756e73796e74686573697a655265717565737428616464726573732c75696e7481527f3235362c616464726573732c616464726573732c5b616464726573732c61646460208201526d726573732c75696e743235365d2960901b6040820152604e0190565b6001600160a01b038881168252602082018890528616604082015260c0606082018190526000906125549083018688612442565b6001600160f81b03199490941660808301525060a0015295945050505050565b9687526001600160a01b039586166020880152939094166040860152606085019190915260ff16608084015260a083019190915260c082015260e00190565b8035600f81900b8114611d9157600080fd5b60ff6125d082611ed2565b16825260208101356020830152604081013560408301526060810135606083015260808101356125ff81612417565b8015156080840152505050565b6102008101863561261c81611d6e565b6001600160a01b03168252612633602088016125b3565b600f0b602083015260408701356040830152606087013561265381611d6e565b6001600160a01b031660608301526080878101359083015261267760a088016125b3565b61268660a0840182600f0b9052565b5060c087013560c083015261269d60e08801611d86565b6001600160a01b031660e08301526126b96101008301876125c5565b6001600160a01b0385166101a08301526001600160a01b0384166101c0830152826101e08301529695505050505050565b634e487b7160e01b600052603260045260246000fd5b60006020828403121561271257600080fd5b6119d082611ed2565b60006020828403121561272d57600080fd5b81356119d081612417565b6001600160a01b0397881681529590961660208601526040850193909352606084019190915260ff16608083015260a082015260c081019190915260e00190565b600060001982141561279b57634e487b7160e01b600052601160045260246000fd5b5060010190565b80356127ad81611d6e565b6001600160a01b0390811683526020820135906127c982611d6e565b166020830152604090810135910152565b600081356001600160e01b031981168082146127f557600080fd5b845250602082013536839003601e1901811261281057600080fd5b8201803567ffffffffffffffff81111561282957600080fd5b80360384131561283857600080fd5b60406020860152806040860152806020830160608701376000606082870101526060601f19601f8301168601019250505092915050565b6101008082528101889052600061012082018a825b8b8110156128b557813561289781611d6e565b6001600160a01b031683526020928301929190910190600101612884565b505082810360208401526128ca81898b612442565b6001600160a01b0388811660408601528716606085015290506128f060808401866127a2565b82810360e084015261290281856127da565b9b9a5050505050505050505050565b7f73796e74686573697a655265717565737428616464726573732c75696e74323581527f362c616464726573732c616464726573732c5b616464726573732c616464726560208201526b73732c75696e743235365d2960a01b6040820152604c0190565b6001600160a01b0386811682526020820186905284811660408301528316606082015260e081016129a960808301846127a2565b9695505050505050565b81835260208301925060008160005b848110156129e7576129d486836125c5565b60a09586019591909101906001016129c2565b5093949350505050565b8060005b60038110156114f2578135612a0981611d6e565b6001600160a01b0316845260209384019391909101906001016129f5565b6060818337506000606082015250565b60006101e08735612a4781611d6e565b6001600160a01b039081168452602089810135908501526040808a013590850152606089013590612a7782611d6e565b16606084015260808881013590840152612a9360a08901611d86565b6001600160a01b031660a0840152612aad60c08901611d86565b6001600160a01b03811660c08501525060e088013560e084015280610100840152612adb81840187896129b3565b915050612aec6101208301856129f1565b6129a9610180830184612a27565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b600060208284031215612b4157600080fd5b81516119d081611d6e565b858152602081018590526001600160a01b0384811660408301528316606082015260e081016129a960808301846127a2565b7f73796e74685472616e736665725265717565737428627974657333322c75696e81527f743235362c616464726573732c616464726573732c5b616464726573732c616460208201526e64726573732c75696e743235365d2960881b6040820152604f0190565b60006102e0612c0483612bf78a611d86565b6001600160a01b03169052565b612c1060208901611d86565b6001600160a01b03166020840152612c2a60408901611d86565b6001600160a01b0316604084015260608881013590840152612c4e608089016125b3565b612c5d6080850182600f0b9052565b50612c6a60a089016125b3565b612c7960a0850182600f0b9052565b5060c088013560c0840152612c9060e089016125b3565b612c9f60e0850182600f0b9052565b506101008881013590840152610120612cb9818a01611d86565b6001600160a01b031690840152610140612cd4898201611d86565b6001600160a01b031690840152610160612cef898201611d86565b6001600160a01b031690840152610180612d0a898201611d86565b6001600160a01b0316908401526101a088810135908401526101c0612d30818a01611d86565b6001600160a01b0316908401526101e088810135908401526102008301819052612d5d81840187896129b3565b915050612d6e6102208301856129f1565b6129a9610280830184612a27565b60005b83811015612d97578181015183820152602001612d7f565b838111156114f25750506000910152565b60008251612dba818460208701612d7c565b9190910192915050565b6020815260008251806020840152612de3816040850160208701612d7c565b601f01601f1916919091016040019291505056fea26469706673582212203f4b07b83b718df8ff3e8913aaebe41514d970471e8eaaad2fd3d447c98b82bf64736f6c634300080a0033",
}

// RouterABI is the input ABI used to generate the binding from.
// Deprecated: Use RouterMetaData.ABI instead.
var RouterABI = RouterMetaData.ABI

// RouterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RouterMetaData.Bin instead.
var RouterBin = RouterMetaData.Bin

// DeployRouter deploys a new Ethereum contract, binding an instance of Router to it.
func DeployRouter(auth *bind.TransactOpts, backend bind.ContractBackend, portal common.Address, synthesis common.Address, curveProxy common.Address) (common.Address, *types.Transaction, *Router, error) {
	parsed, err := RouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RouterBin), backend, portal, synthesis, curveProxy)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Router{RouterCaller: RouterCaller{contract: contract}, RouterTransactor: RouterTransactor{contract: contract}, RouterFilterer: RouterFilterer{contract: contract}}, nil
}

// Router is an auto generated Go binding around an Ethereum contract.
type Router struct {
	RouterCaller     // Read-only binding to the contract
	RouterTransactor // Write-only binding to the contract
	RouterFilterer   // Log filterer for contract events
}

// RouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type RouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
	gsn      *GsnCallOpts
}

func (_Router *RouterTransactor) SetGSNOptions(opts *GsnCallOpts) {
	_Router.gsn = opts
}

// RouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RouterSession struct {
	Contract     *Router           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RouterCallerSession struct {
	Contract *RouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RouterTransactorSession struct {
	Contract     *RouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type RouterRaw struct {
	Contract *Router // Generic contract binding to access the raw methods on
}

// RouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RouterCallerRaw struct {
	Contract *RouterCaller // Generic read-only contract binding to access the raw methods on
}

// RouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RouterTransactorRaw struct {
	Contract *RouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRouter creates a new instance of Router, bound to a specific deployed contract.
func NewRouter(address common.Address, backend bind.ContractBackend) (*Router, error) {
	contract, err := bindRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Router{RouterCaller: RouterCaller{contract: contract}, RouterTransactor: RouterTransactor{contract: contract}, RouterFilterer: RouterFilterer{contract: contract}}, nil
}

// NewRouterCaller creates a new read-only instance of Router, bound to a specific deployed contract.
func NewRouterCaller(address common.Address, caller bind.ContractCaller) (*RouterCaller, error) {
	contract, err := bindRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RouterCaller{contract: contract}, nil
}

// NewRouterTransactor creates a new write-only instance of Router, bound to a specific deployed contract.
func NewRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*RouterTransactor, error) {
	contract, err := bindRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RouterTransactor{contract: contract}, nil
}

// NewRouterFilterer creates a new log filterer instance of Router, bound to a specific deployed contract.
func NewRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*RouterFilterer, error) {
	contract, err := bindRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RouterFilterer{contract: contract}, nil
}

// bindRouter binds a generic wrapper to an already deployed contract.
func bindRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.RouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Router *RouterCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Router *RouterSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Router.Contract.DOMAINSEPARATOR(&_Router.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Router *RouterCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Router.Contract.DOMAINSEPARATOR(&_Router.CallOpts)
}

// SYNTHESIZEREQUESTSIGNATUREHASH is a free data retrieval call binding the contract method 0x71056f5e.
//
// Solidity: function _SYNTHESIZE_REQUEST_SIGNATURE_HASH() view returns(bytes32)
func (_Router *RouterCaller) SYNTHESIZEREQUESTSIGNATUREHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "_SYNTHESIZE_REQUEST_SIGNATURE_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SYNTHESIZEREQUESTSIGNATUREHASH is a free data retrieval call binding the contract method 0x71056f5e.
//
// Solidity: function _SYNTHESIZE_REQUEST_SIGNATURE_HASH() view returns(bytes32)
func (_Router *RouterSession) SYNTHESIZEREQUESTSIGNATUREHASH() ([32]byte, error) {
	return _Router.Contract.SYNTHESIZEREQUESTSIGNATUREHASH(&_Router.CallOpts)
}

// SYNTHESIZEREQUESTSIGNATUREHASH is a free data retrieval call binding the contract method 0x71056f5e.
//
// Solidity: function _SYNTHESIZE_REQUEST_SIGNATURE_HASH() view returns(bytes32)
func (_Router *RouterCallerSession) SYNTHESIZEREQUESTSIGNATUREHASH() ([32]byte, error) {
	return _Router.Contract.SYNTHESIZEREQUESTSIGNATUREHASH(&_Router.CallOpts)
}

// SYNTHTRANSFERREQUESTSIGNATUREHASH is a free data retrieval call binding the contract method 0xd3c52284.
//
// Solidity: function _SYNTH_TRANSFER_REQUEST_SIGNATURE_HASH() view returns(bytes32)
func (_Router *RouterCaller) SYNTHTRANSFERREQUESTSIGNATUREHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "_SYNTH_TRANSFER_REQUEST_SIGNATURE_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SYNTHTRANSFERREQUESTSIGNATUREHASH is a free data retrieval call binding the contract method 0xd3c52284.
//
// Solidity: function _SYNTH_TRANSFER_REQUEST_SIGNATURE_HASH() view returns(bytes32)
func (_Router *RouterSession) SYNTHTRANSFERREQUESTSIGNATUREHASH() ([32]byte, error) {
	return _Router.Contract.SYNTHTRANSFERREQUESTSIGNATUREHASH(&_Router.CallOpts)
}

// SYNTHTRANSFERREQUESTSIGNATUREHASH is a free data retrieval call binding the contract method 0xd3c52284.
//
// Solidity: function _SYNTH_TRANSFER_REQUEST_SIGNATURE_HASH() view returns(bytes32)
func (_Router *RouterCallerSession) SYNTHTRANSFERREQUESTSIGNATUREHASH() ([32]byte, error) {
	return _Router.Contract.SYNTHTRANSFERREQUESTSIGNATUREHASH(&_Router.CallOpts)
}

// UNSYNTHESIZEREQUESTSIGNATUREHASH is a free data retrieval call binding the contract method 0x2acbaacf.
//
// Solidity: function _UNSYNTHESIZE_REQUEST_SIGNATURE_HASH() view returns(bytes32)
func (_Router *RouterCaller) UNSYNTHESIZEREQUESTSIGNATUREHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "_UNSYNTHESIZE_REQUEST_SIGNATURE_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UNSYNTHESIZEREQUESTSIGNATUREHASH is a free data retrieval call binding the contract method 0x2acbaacf.
//
// Solidity: function _UNSYNTHESIZE_REQUEST_SIGNATURE_HASH() view returns(bytes32)
func (_Router *RouterSession) UNSYNTHESIZEREQUESTSIGNATUREHASH() ([32]byte, error) {
	return _Router.Contract.UNSYNTHESIZEREQUESTSIGNATUREHASH(&_Router.CallOpts)
}

// UNSYNTHESIZEREQUESTSIGNATUREHASH is a free data retrieval call binding the contract method 0x2acbaacf.
//
// Solidity: function _UNSYNTHESIZE_REQUEST_SIGNATURE_HASH() view returns(bytes32)
func (_Router *RouterCallerSession) UNSYNTHESIZEREQUESTSIGNATUREHASH() ([32]byte, error) {
	return _Router.Contract.UNSYNTHESIZEREQUESTSIGNATUREHASH(&_Router.CallOpts)
}

// TrustedWorker is a free data retrieval call binding the contract method 0xe840f9d3.
//
// Solidity: function _trustedWorker(address ) view returns(bool)
func (_Router *RouterCaller) TrustedWorker(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "_trustedWorker", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TrustedWorker is a free data retrieval call binding the contract method 0xe840f9d3.
//
// Solidity: function _trustedWorker(address ) view returns(bool)
func (_Router *RouterSession) TrustedWorker(arg0 common.Address) (bool, error) {
	return _Router.Contract.TrustedWorker(&_Router.CallOpts, arg0)
}

// TrustedWorker is a free data retrieval call binding the contract method 0xe840f9d3.
//
// Solidity: function _trustedWorker(address ) view returns(bool)
func (_Router *RouterCallerSession) TrustedWorker(arg0 common.Address) (bool, error) {
	return _Router.Contract.TrustedWorker(&_Router.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Router *RouterCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Router *RouterSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Router.Contract.Nonces(&_Router.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Router *RouterCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Router.Contract.Nonces(&_Router.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Router *RouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Router *RouterSession) Owner() (common.Address, error) {
	return _Router.Contract.Owner(&_Router.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Router *RouterCallerSession) Owner() (common.Address, error) {
	return _Router.Contract.Owner(&_Router.CallOpts)
}

// BatchSynthesizeRequestWithDataTransit is a paid mutator transaction binding the contract method 0x4d176e84.
//
// Solidity: function batchSynthesizeRequestWithDataTransit(address[] token, uint256[] amount, address to, (bytes4,bytes) transitData, (address,address,uint256) synthParams, (uint8,bytes32,bytes32,uint256,bool)[] permitData) returns()
func (_Router *RouterTransactor) BatchSynthesizeRequestWithDataTransit(opts *bind.TransactOpts, token []common.Address, amount []*big.Int, to common.Address, transitData IPortalTransitData, synthParams IPortalSynthParams, permitData []IPortalPermitData) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "batchSynthesizeRequestWithDataTransit", token, amount, to, transitData, synthParams, permitData)
}
func (_Router *RouterTransactor) BatchSynthesizeRequestWithDataTransitOverGsn(opts *bind.TransactOpts, token []common.Address, amount []*big.Int, to common.Address, transitData IPortalTransitData, synthParams IPortalSynthParams, permitData []IPortalPermitData) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "batchSynthesizeRequestWithDataTransit", token, amount, to, transitData, synthParams, permitData)
}

// BatchSynthesizeRequestWithDataTransit is a paid mutator transaction binding the contract method 0x4d176e84.
//
// Solidity: function batchSynthesizeRequestWithDataTransit(address[] token, uint256[] amount, address to, (bytes4,bytes) transitData, (address,address,uint256) synthParams, (uint8,bytes32,bytes32,uint256,bool)[] permitData) returns()
func (_Router *RouterSession) BatchSynthesizeRequestWithDataTransit(token []common.Address, amount []*big.Int, to common.Address, transitData IPortalTransitData, synthParams IPortalSynthParams, permitData []IPortalPermitData) (*types.Transaction, error) {
	return _Router.Contract.BatchSynthesizeRequestWithDataTransit(&_Router.TransactOpts, token, amount, to, transitData, synthParams, permitData)
}
func (_Router *RouterSession) BatchSynthesizeRequestWithDataTransitOverGsn(token []common.Address, amount []*big.Int, to common.Address, transitData IPortalTransitData, synthParams IPortalSynthParams, permitData []IPortalPermitData) (common.Hash, error) {
	return _Router.Contract.BatchSynthesizeRequestWithDataTransitOverGsn(&_Router.TransactOpts, token, amount, to, transitData, synthParams, permitData)
}

// BatchSynthesizeRequestWithDataTransit is a paid mutator transaction binding the contract method 0x4d176e84.
//
// Solidity: function batchSynthesizeRequestWithDataTransit(address[] token, uint256[] amount, address to, (bytes4,bytes) transitData, (address,address,uint256) synthParams, (uint8,bytes32,bytes32,uint256,bool)[] permitData) returns()
func (_Router *RouterTransactorSession) BatchSynthesizeRequestWithDataTransit(token []common.Address, amount []*big.Int, to common.Address, transitData IPortalTransitData, synthParams IPortalSynthParams, permitData []IPortalPermitData) (*types.Transaction, error) {
	return _Router.Contract.BatchSynthesizeRequestWithDataTransit(&_Router.TransactOpts, token, amount, to, transitData, synthParams, permitData)
}
func (_Router *RouterTransactorSession) BatchSynthesizeRequestWithDataTransitOverGsn(token []common.Address, amount []*big.Int, to common.Address, transitData IPortalTransitData, synthParams IPortalSynthParams, permitData []IPortalPermitData) (common.Hash, error) {
	return _Router.Contract.BatchSynthesizeRequestWithDataTransitOverGsn(&_Router.TransactOpts, token, amount, to, transitData, synthParams, permitData)
}

// EmergencyUnburnRequest is a paid mutator transaction binding the contract method 0x43961668.
//
// Solidity: function emergencyUnburnRequest(bytes32 txID, address receiveSide, address oppositeBridge, uint256 chainId, uint8 v, bytes32 r, bytes32 s) returns()
func (_Router *RouterTransactor) EmergencyUnburnRequest(opts *bind.TransactOpts, txID [32]byte, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "emergencyUnburnRequest", txID, receiveSide, oppositeBridge, chainId, v, r, s)
}
func (_Router *RouterTransactor) EmergencyUnburnRequestOverGsn(opts *bind.TransactOpts, txID [32]byte, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, v uint8, r [32]byte, s [32]byte) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "emergencyUnburnRequest", txID, receiveSide, oppositeBridge, chainId, v, r, s)
}

// EmergencyUnburnRequest is a paid mutator transaction binding the contract method 0x43961668.
//
// Solidity: function emergencyUnburnRequest(bytes32 txID, address receiveSide, address oppositeBridge, uint256 chainId, uint8 v, bytes32 r, bytes32 s) returns()
func (_Router *RouterSession) EmergencyUnburnRequest(txID [32]byte, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Router.Contract.EmergencyUnburnRequest(&_Router.TransactOpts, txID, receiveSide, oppositeBridge, chainId, v, r, s)
}
func (_Router *RouterSession) EmergencyUnburnRequestOverGsn(txID [32]byte, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, v uint8, r [32]byte, s [32]byte) (common.Hash, error) {
	return _Router.Contract.EmergencyUnburnRequestOverGsn(&_Router.TransactOpts, txID, receiveSide, oppositeBridge, chainId, v, r, s)
}

// EmergencyUnburnRequest is a paid mutator transaction binding the contract method 0x43961668.
//
// Solidity: function emergencyUnburnRequest(bytes32 txID, address receiveSide, address oppositeBridge, uint256 chainId, uint8 v, bytes32 r, bytes32 s) returns()
func (_Router *RouterTransactorSession) EmergencyUnburnRequest(txID [32]byte, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Router.Contract.EmergencyUnburnRequest(&_Router.TransactOpts, txID, receiveSide, oppositeBridge, chainId, v, r, s)
}
func (_Router *RouterTransactorSession) EmergencyUnburnRequestOverGsn(txID [32]byte, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, v uint8, r [32]byte, s [32]byte) (common.Hash, error) {
	return _Router.Contract.EmergencyUnburnRequestOverGsn(&_Router.TransactOpts, txID, receiveSide, oppositeBridge, chainId, v, r, s)
}

// EmergencyUnsyntesizeRequest is a paid mutator transaction binding the contract method 0x9500125b.
//
// Solidity: function emergencyUnsyntesizeRequest(bytes32 txID, address receiveSide, address oppositeBridge, uint256 chainId, uint8 v, bytes32 r, bytes32 s) returns()
func (_Router *RouterTransactor) EmergencyUnsyntesizeRequest(opts *bind.TransactOpts, txID [32]byte, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "emergencyUnsyntesizeRequest", txID, receiveSide, oppositeBridge, chainId, v, r, s)
}
func (_Router *RouterTransactor) EmergencyUnsyntesizeRequestOverGsn(opts *bind.TransactOpts, txID [32]byte, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, v uint8, r [32]byte, s [32]byte) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "emergencyUnsyntesizeRequest", txID, receiveSide, oppositeBridge, chainId, v, r, s)
}

// EmergencyUnsyntesizeRequest is a paid mutator transaction binding the contract method 0x9500125b.
//
// Solidity: function emergencyUnsyntesizeRequest(bytes32 txID, address receiveSide, address oppositeBridge, uint256 chainId, uint8 v, bytes32 r, bytes32 s) returns()
func (_Router *RouterSession) EmergencyUnsyntesizeRequest(txID [32]byte, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Router.Contract.EmergencyUnsyntesizeRequest(&_Router.TransactOpts, txID, receiveSide, oppositeBridge, chainId, v, r, s)
}
func (_Router *RouterSession) EmergencyUnsyntesizeRequestOverGsn(txID [32]byte, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, v uint8, r [32]byte, s [32]byte) (common.Hash, error) {
	return _Router.Contract.EmergencyUnsyntesizeRequestOverGsn(&_Router.TransactOpts, txID, receiveSide, oppositeBridge, chainId, v, r, s)
}

// EmergencyUnsyntesizeRequest is a paid mutator transaction binding the contract method 0x9500125b.
//
// Solidity: function emergencyUnsyntesizeRequest(bytes32 txID, address receiveSide, address oppositeBridge, uint256 chainId, uint8 v, bytes32 r, bytes32 s) returns()
func (_Router *RouterTransactorSession) EmergencyUnsyntesizeRequest(txID [32]byte, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Router.Contract.EmergencyUnsyntesizeRequest(&_Router.TransactOpts, txID, receiveSide, oppositeBridge, chainId, v, r, s)
}
func (_Router *RouterTransactorSession) EmergencyUnsyntesizeRequestOverGsn(txID [32]byte, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, v uint8, r [32]byte, s [32]byte) (common.Hash, error) {
	return _Router.Contract.EmergencyUnsyntesizeRequestOverGsn(&_Router.TransactOpts, txID, receiveSide, oppositeBridge, chainId, v, r, s)
}

// MetaExchangeRequestVia3pool is a paid mutator transaction binding the contract method 0xf8900f78.
//
// Solidity: function metaExchangeRequestVia3pool((address,address,address,uint256,int128,int128,uint256,int128,uint256,address,address,address,address,uint256,address,uint256) params, (uint8,bytes32,bytes32,uint256,bool)[] permit, address[3] token, uint256[3] amount) returns()
func (_Router *RouterTransactor) MetaExchangeRequestVia3pool(opts *bind.TransactOpts, params ICurveProxyMetaExchangeParams, permit []ICurveProxyPermitData, token [3]common.Address, amount [3]*big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "metaExchangeRequestVia3pool", params, permit, token, amount)
}
func (_Router *RouterTransactor) MetaExchangeRequestVia3poolOverGsn(opts *bind.TransactOpts, params ICurveProxyMetaExchangeParams, permit []ICurveProxyPermitData, token [3]common.Address, amount [3]*big.Int) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "metaExchangeRequestVia3pool", params, permit, token, amount)
}

// MetaExchangeRequestVia3pool is a paid mutator transaction binding the contract method 0xf8900f78.
//
// Solidity: function metaExchangeRequestVia3pool((address,address,address,uint256,int128,int128,uint256,int128,uint256,address,address,address,address,uint256,address,uint256) params, (uint8,bytes32,bytes32,uint256,bool)[] permit, address[3] token, uint256[3] amount) returns()
func (_Router *RouterSession) MetaExchangeRequestVia3pool(params ICurveProxyMetaExchangeParams, permit []ICurveProxyPermitData, token [3]common.Address, amount [3]*big.Int) (*types.Transaction, error) {
	return _Router.Contract.MetaExchangeRequestVia3pool(&_Router.TransactOpts, params, permit, token, amount)
}
func (_Router *RouterSession) MetaExchangeRequestVia3poolOverGsn(params ICurveProxyMetaExchangeParams, permit []ICurveProxyPermitData, token [3]common.Address, amount [3]*big.Int) (common.Hash, error) {
	return _Router.Contract.MetaExchangeRequestVia3poolOverGsn(&_Router.TransactOpts, params, permit, token, amount)
}

// MetaExchangeRequestVia3pool is a paid mutator transaction binding the contract method 0xf8900f78.
//
// Solidity: function metaExchangeRequestVia3pool((address,address,address,uint256,int128,int128,uint256,int128,uint256,address,address,address,address,uint256,address,uint256) params, (uint8,bytes32,bytes32,uint256,bool)[] permit, address[3] token, uint256[3] amount) returns()
func (_Router *RouterTransactorSession) MetaExchangeRequestVia3pool(params ICurveProxyMetaExchangeParams, permit []ICurveProxyPermitData, token [3]common.Address, amount [3]*big.Int) (*types.Transaction, error) {
	return _Router.Contract.MetaExchangeRequestVia3pool(&_Router.TransactOpts, params, permit, token, amount)
}
func (_Router *RouterTransactorSession) MetaExchangeRequestVia3poolOverGsn(params ICurveProxyMetaExchangeParams, permit []ICurveProxyPermitData, token [3]common.Address, amount [3]*big.Int) (common.Hash, error) {
	return _Router.Contract.MetaExchangeRequestVia3poolOverGsn(&_Router.TransactOpts, params, permit, token, amount)
}

// MintEusdRequestVia3pool is a paid mutator transaction binding the contract method 0x69ab5d46.
//
// Solidity: function mintEusdRequestVia3pool((address,uint256,uint256,address,uint256,address,address,uint256) params, (uint8,bytes32,bytes32,uint256,bool)[] permit, address[3] token, uint256[3] amount) returns()
func (_Router *RouterTransactor) MintEusdRequestVia3pool(opts *bind.TransactOpts, params ICurveProxyMetaMintEUSD, permit []ICurveProxyPermitData, token [3]common.Address, amount [3]*big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "mintEusdRequestVia3pool", params, permit, token, amount)
}
func (_Router *RouterTransactor) MintEusdRequestVia3poolOverGsn(opts *bind.TransactOpts, params ICurveProxyMetaMintEUSD, permit []ICurveProxyPermitData, token [3]common.Address, amount [3]*big.Int) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "mintEusdRequestVia3pool", params, permit, token, amount)
}

// MintEusdRequestVia3pool is a paid mutator transaction binding the contract method 0x69ab5d46.
//
// Solidity: function mintEusdRequestVia3pool((address,uint256,uint256,address,uint256,address,address,uint256) params, (uint8,bytes32,bytes32,uint256,bool)[] permit, address[3] token, uint256[3] amount) returns()
func (_Router *RouterSession) MintEusdRequestVia3pool(params ICurveProxyMetaMintEUSD, permit []ICurveProxyPermitData, token [3]common.Address, amount [3]*big.Int) (*types.Transaction, error) {
	return _Router.Contract.MintEusdRequestVia3pool(&_Router.TransactOpts, params, permit, token, amount)
}
func (_Router *RouterSession) MintEusdRequestVia3poolOverGsn(params ICurveProxyMetaMintEUSD, permit []ICurveProxyPermitData, token [3]common.Address, amount [3]*big.Int) (common.Hash, error) {
	return _Router.Contract.MintEusdRequestVia3poolOverGsn(&_Router.TransactOpts, params, permit, token, amount)
}

// MintEusdRequestVia3pool is a paid mutator transaction binding the contract method 0x69ab5d46.
//
// Solidity: function mintEusdRequestVia3pool((address,uint256,uint256,address,uint256,address,address,uint256) params, (uint8,bytes32,bytes32,uint256,bool)[] permit, address[3] token, uint256[3] amount) returns()
func (_Router *RouterTransactorSession) MintEusdRequestVia3pool(params ICurveProxyMetaMintEUSD, permit []ICurveProxyPermitData, token [3]common.Address, amount [3]*big.Int) (*types.Transaction, error) {
	return _Router.Contract.MintEusdRequestVia3pool(&_Router.TransactOpts, params, permit, token, amount)
}
func (_Router *RouterTransactorSession) MintEusdRequestVia3poolOverGsn(params ICurveProxyMetaMintEUSD, permit []ICurveProxyPermitData, token [3]common.Address, amount [3]*big.Int) (common.Hash, error) {
	return _Router.Contract.MintEusdRequestVia3poolOverGsn(&_Router.TransactOpts, params, permit, token, amount)
}

// RedeemEusdRequest is a paid mutator transaction binding the contract method 0x44c82d23.
//
// Solidity: function redeemEusdRequest((address,int128,uint256,address,uint256,int128,uint256,address) params, (uint8,bytes32,bytes32,uint256,bool) permit, address payToken, address receiveSide, address oppositeBridge, uint256 chainId) returns()
func (_Router *RouterTransactor) RedeemEusdRequest(opts *bind.TransactOpts, params ICurveProxyMetaRedeemEUSD, permit ICurveProxyPermitData, payToken common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "redeemEusdRequest", params, permit, payToken, receiveSide, oppositeBridge, chainId)
}
func (_Router *RouterTransactor) RedeemEusdRequestOverGsn(opts *bind.TransactOpts, params ICurveProxyMetaRedeemEUSD, permit ICurveProxyPermitData, payToken common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "redeemEusdRequest", params, permit, payToken, receiveSide, oppositeBridge, chainId)
}

// RedeemEusdRequest is a paid mutator transaction binding the contract method 0x44c82d23.
//
// Solidity: function redeemEusdRequest((address,int128,uint256,address,uint256,int128,uint256,address) params, (uint8,bytes32,bytes32,uint256,bool) permit, address payToken, address receiveSide, address oppositeBridge, uint256 chainId) returns()
func (_Router *RouterSession) RedeemEusdRequest(params ICurveProxyMetaRedeemEUSD, permit ICurveProxyPermitData, payToken common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _Router.Contract.RedeemEusdRequest(&_Router.TransactOpts, params, permit, payToken, receiveSide, oppositeBridge, chainId)
}
func (_Router *RouterSession) RedeemEusdRequestOverGsn(params ICurveProxyMetaRedeemEUSD, permit ICurveProxyPermitData, payToken common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (common.Hash, error) {
	return _Router.Contract.RedeemEusdRequestOverGsn(&_Router.TransactOpts, params, permit, payToken, receiveSide, oppositeBridge, chainId)
}

// RedeemEusdRequest is a paid mutator transaction binding the contract method 0x44c82d23.
//
// Solidity: function redeemEusdRequest((address,int128,uint256,address,uint256,int128,uint256,address) params, (uint8,bytes32,bytes32,uint256,bool) permit, address payToken, address receiveSide, address oppositeBridge, uint256 chainId) returns()
func (_Router *RouterTransactorSession) RedeemEusdRequest(params ICurveProxyMetaRedeemEUSD, permit ICurveProxyPermitData, payToken common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _Router.Contract.RedeemEusdRequest(&_Router.TransactOpts, params, permit, payToken, receiveSide, oppositeBridge, chainId)
}
func (_Router *RouterTransactorSession) RedeemEusdRequestOverGsn(params ICurveProxyMetaRedeemEUSD, permit ICurveProxyPermitData, payToken common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (common.Hash, error) {
	return _Router.Contract.RedeemEusdRequestOverGsn(&_Router.TransactOpts, params, permit, payToken, receiveSide, oppositeBridge, chainId)
}

// RemoveTrustedWorker is a paid mutator transaction binding the contract method 0x72b90c02.
//
// Solidity: function removeTrustedWorker(address worker) returns()
func (_Router *RouterTransactor) RemoveTrustedWorker(opts *bind.TransactOpts, worker common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "removeTrustedWorker", worker)
}
func (_Router *RouterTransactor) RemoveTrustedWorkerOverGsn(opts *bind.TransactOpts, worker common.Address) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "removeTrustedWorker", worker)
}

// RemoveTrustedWorker is a paid mutator transaction binding the contract method 0x72b90c02.
//
// Solidity: function removeTrustedWorker(address worker) returns()
func (_Router *RouterSession) RemoveTrustedWorker(worker common.Address) (*types.Transaction, error) {
	return _Router.Contract.RemoveTrustedWorker(&_Router.TransactOpts, worker)
}
func (_Router *RouterSession) RemoveTrustedWorkerOverGsn(worker common.Address) (common.Hash, error) {
	return _Router.Contract.RemoveTrustedWorkerOverGsn(&_Router.TransactOpts, worker)
}

// RemoveTrustedWorker is a paid mutator transaction binding the contract method 0x72b90c02.
//
// Solidity: function removeTrustedWorker(address worker) returns()
func (_Router *RouterTransactorSession) RemoveTrustedWorker(worker common.Address) (*types.Transaction, error) {
	return _Router.Contract.RemoveTrustedWorker(&_Router.TransactOpts, worker)
}
func (_Router *RouterTransactorSession) RemoveTrustedWorkerOverGsn(worker common.Address) (common.Hash, error) {
	return _Router.Contract.RemoveTrustedWorkerOverGsn(&_Router.TransactOpts, worker)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Router *RouterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "renounceOwnership")
}
func (_Router *RouterTransactor) RenounceOwnershipOverGsn(opts *bind.TransactOpts) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Router *RouterSession) RenounceOwnership() (*types.Transaction, error) {
	return _Router.Contract.RenounceOwnership(&_Router.TransactOpts)
}
func (_Router *RouterSession) RenounceOwnershipOverGsn() (common.Hash, error) {
	return _Router.Contract.RenounceOwnershipOverGsn(&_Router.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Router *RouterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Router.Contract.RenounceOwnership(&_Router.TransactOpts)
}
func (_Router *RouterTransactorSession) RenounceOwnershipOverGsn() (common.Hash, error) {
	return _Router.Contract.RenounceOwnershipOverGsn(&_Router.TransactOpts)
}

// SetTrustedWorker is a paid mutator transaction binding the contract method 0x7fe04792.
//
// Solidity: function setTrustedWorker(address worker) returns()
func (_Router *RouterTransactor) SetTrustedWorker(opts *bind.TransactOpts, worker common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "setTrustedWorker", worker)
}
func (_Router *RouterTransactor) SetTrustedWorkerOverGsn(opts *bind.TransactOpts, worker common.Address) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "setTrustedWorker", worker)
}

// SetTrustedWorker is a paid mutator transaction binding the contract method 0x7fe04792.
//
// Solidity: function setTrustedWorker(address worker) returns()
func (_Router *RouterSession) SetTrustedWorker(worker common.Address) (*types.Transaction, error) {
	return _Router.Contract.SetTrustedWorker(&_Router.TransactOpts, worker)
}
func (_Router *RouterSession) SetTrustedWorkerOverGsn(worker common.Address) (common.Hash, error) {
	return _Router.Contract.SetTrustedWorkerOverGsn(&_Router.TransactOpts, worker)
}

// SetTrustedWorker is a paid mutator transaction binding the contract method 0x7fe04792.
//
// Solidity: function setTrustedWorker(address worker) returns()
func (_Router *RouterTransactorSession) SetTrustedWorker(worker common.Address) (*types.Transaction, error) {
	return _Router.Contract.SetTrustedWorker(&_Router.TransactOpts, worker)
}
func (_Router *RouterTransactorSession) SetTrustedWorkerOverGsn(worker common.Address) (common.Hash, error) {
	return _Router.Contract.SetTrustedWorkerOverGsn(&_Router.TransactOpts, worker)
}

// SynthTransferRequest is a paid mutator transaction binding the contract method 0xc3345883.
//
// Solidity: function synthTransferRequest(bytes32 rtoken, uint256 amount, address to, (address,address,uint256) synthParams) returns()
func (_Router *RouterTransactor) SynthTransferRequest(opts *bind.TransactOpts, rtoken [32]byte, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "synthTransferRequest", rtoken, amount, to, synthParams)
}
func (_Router *RouterTransactor) SynthTransferRequestOverGsn(opts *bind.TransactOpts, rtoken [32]byte, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "synthTransferRequest", rtoken, amount, to, synthParams)
}

// SynthTransferRequest is a paid mutator transaction binding the contract method 0xc3345883.
//
// Solidity: function synthTransferRequest(bytes32 rtoken, uint256 amount, address to, (address,address,uint256) synthParams) returns()
func (_Router *RouterSession) SynthTransferRequest(rtoken [32]byte, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams) (*types.Transaction, error) {
	return _Router.Contract.SynthTransferRequest(&_Router.TransactOpts, rtoken, amount, to, synthParams)
}
func (_Router *RouterSession) SynthTransferRequestOverGsn(rtoken [32]byte, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams) (common.Hash, error) {
	return _Router.Contract.SynthTransferRequestOverGsn(&_Router.TransactOpts, rtoken, amount, to, synthParams)
}

// SynthTransferRequest is a paid mutator transaction binding the contract method 0xc3345883.
//
// Solidity: function synthTransferRequest(bytes32 rtoken, uint256 amount, address to, (address,address,uint256) synthParams) returns()
func (_Router *RouterTransactorSession) SynthTransferRequest(rtoken [32]byte, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams) (*types.Transaction, error) {
	return _Router.Contract.SynthTransferRequest(&_Router.TransactOpts, rtoken, amount, to, synthParams)
}
func (_Router *RouterTransactorSession) SynthTransferRequestOverGsn(rtoken [32]byte, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams) (common.Hash, error) {
	return _Router.Contract.SynthTransferRequestOverGsn(&_Router.TransactOpts, rtoken, amount, to, synthParams)
}

// SynthTransferRequestPayNative is a paid mutator transaction binding the contract method 0xc773abfd.
//
// Solidity: function synthTransferRequestPayNative(bytes32 tokenReal, address tokenSynth, uint256 amount, address to, (address,address,uint256) synthParams, (uint256,uint256,uint8,bytes32,bytes32) receipt) payable returns()
func (_Router *RouterTransactor) SynthTransferRequestPayNative(opts *bind.TransactOpts, tokenReal [32]byte, tokenSynth common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "synthTransferRequestPayNative", tokenReal, tokenSynth, amount, to, synthParams, receipt)
}
func (_Router *RouterTransactor) SynthTransferRequestPayNativeOverGsn(opts *bind.TransactOpts, tokenReal [32]byte, tokenSynth common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "synthTransferRequestPayNative", tokenReal, tokenSynth, amount, to, synthParams, receipt)
}

// SynthTransferRequestPayNative is a paid mutator transaction binding the contract method 0xc773abfd.
//
// Solidity: function synthTransferRequestPayNative(bytes32 tokenReal, address tokenSynth, uint256 amount, address to, (address,address,uint256) synthParams, (uint256,uint256,uint8,bytes32,bytes32) receipt) payable returns()
func (_Router *RouterSession) SynthTransferRequestPayNative(tokenReal [32]byte, tokenSynth common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.SynthTransferRequestPayNative(&_Router.TransactOpts, tokenReal, tokenSynth, amount, to, synthParams, receipt)
}
func (_Router *RouterSession) SynthTransferRequestPayNativeOverGsn(tokenReal [32]byte, tokenSynth common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.SynthTransferRequestPayNativeOverGsn(&_Router.TransactOpts, tokenReal, tokenSynth, amount, to, synthParams, receipt)
}

// SynthTransferRequestPayNative is a paid mutator transaction binding the contract method 0xc773abfd.
//
// Solidity: function synthTransferRequestPayNative(bytes32 tokenReal, address tokenSynth, uint256 amount, address to, (address,address,uint256) synthParams, (uint256,uint256,uint8,bytes32,bytes32) receipt) payable returns()
func (_Router *RouterTransactorSession) SynthTransferRequestPayNative(tokenReal [32]byte, tokenSynth common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.SynthTransferRequestPayNative(&_Router.TransactOpts, tokenReal, tokenSynth, amount, to, synthParams, receipt)
}
func (_Router *RouterTransactorSession) SynthTransferRequestPayNativeOverGsn(tokenReal [32]byte, tokenSynth common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.SynthTransferRequestPayNativeOverGsn(&_Router.TransactOpts, tokenReal, tokenSynth, amount, to, synthParams, receipt)
}

// SynthesizeRequestPayNative is a paid mutator transaction binding the contract method 0xd63f79b2.
//
// Solidity: function synthesizeRequestPayNative(address token, uint256 amount, address to, (address,address,uint256) synthParams, (uint256,uint256,uint8,bytes32,bytes32) receipt) payable returns()
func (_Router *RouterTransactor) SynthesizeRequestPayNative(opts *bind.TransactOpts, token common.Address, amount *big.Int, to common.Address, synthParams IPortalSynthParams, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "synthesizeRequestPayNative", token, amount, to, synthParams, receipt)
}
func (_Router *RouterTransactor) SynthesizeRequestPayNativeOverGsn(opts *bind.TransactOpts, token common.Address, amount *big.Int, to common.Address, synthParams IPortalSynthParams, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "synthesizeRequestPayNative", token, amount, to, synthParams, receipt)
}

// SynthesizeRequestPayNative is a paid mutator transaction binding the contract method 0xd63f79b2.
//
// Solidity: function synthesizeRequestPayNative(address token, uint256 amount, address to, (address,address,uint256) synthParams, (uint256,uint256,uint8,bytes32,bytes32) receipt) payable returns()
func (_Router *RouterSession) SynthesizeRequestPayNative(token common.Address, amount *big.Int, to common.Address, synthParams IPortalSynthParams, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.SynthesizeRequestPayNative(&_Router.TransactOpts, token, amount, to, synthParams, receipt)
}
func (_Router *RouterSession) SynthesizeRequestPayNativeOverGsn(token common.Address, amount *big.Int, to common.Address, synthParams IPortalSynthParams, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.SynthesizeRequestPayNativeOverGsn(&_Router.TransactOpts, token, amount, to, synthParams, receipt)
}

// SynthesizeRequestPayNative is a paid mutator transaction binding the contract method 0xd63f79b2.
//
// Solidity: function synthesizeRequestPayNative(address token, uint256 amount, address to, (address,address,uint256) synthParams, (uint256,uint256,uint8,bytes32,bytes32) receipt) payable returns()
func (_Router *RouterTransactorSession) SynthesizeRequestPayNative(token common.Address, amount *big.Int, to common.Address, synthParams IPortalSynthParams, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.SynthesizeRequestPayNative(&_Router.TransactOpts, token, amount, to, synthParams, receipt)
}
func (_Router *RouterTransactorSession) SynthesizeRequestPayNativeOverGsn(token common.Address, amount *big.Int, to common.Address, synthParams IPortalSynthParams, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.SynthesizeRequestPayNativeOverGsn(&_Router.TransactOpts, token, amount, to, synthParams, receipt)
}

// SynthesizeRequestWithPermitPayNative is a paid mutator transaction binding the contract method 0x4ec658e2.
//
// Solidity: function synthesizeRequestWithPermitPayNative(address token, uint256 amount, address to, (address,address,uint256) synthParams, (uint8,bytes32,bytes32,uint256,bool) permitData, (uint256,uint256,uint8,bytes32,bytes32) receipt) payable returns()
func (_Router *RouterTransactor) SynthesizeRequestWithPermitPayNative(opts *bind.TransactOpts, token common.Address, amount *big.Int, to common.Address, synthParams IPortalSynthParams, permitData IPortalPermitData, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "synthesizeRequestWithPermitPayNative", token, amount, to, synthParams, permitData, receipt)
}
func (_Router *RouterTransactor) SynthesizeRequestWithPermitPayNativeOverGsn(opts *bind.TransactOpts, token common.Address, amount *big.Int, to common.Address, synthParams IPortalSynthParams, permitData IPortalPermitData, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "synthesizeRequestWithPermitPayNative", token, amount, to, synthParams, permitData, receipt)
}

// SynthesizeRequestWithPermitPayNative is a paid mutator transaction binding the contract method 0x4ec658e2.
//
// Solidity: function synthesizeRequestWithPermitPayNative(address token, uint256 amount, address to, (address,address,uint256) synthParams, (uint8,bytes32,bytes32,uint256,bool) permitData, (uint256,uint256,uint8,bytes32,bytes32) receipt) payable returns()
func (_Router *RouterSession) SynthesizeRequestWithPermitPayNative(token common.Address, amount *big.Int, to common.Address, synthParams IPortalSynthParams, permitData IPortalPermitData, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.SynthesizeRequestWithPermitPayNative(&_Router.TransactOpts, token, amount, to, synthParams, permitData, receipt)
}
func (_Router *RouterSession) SynthesizeRequestWithPermitPayNativeOverGsn(token common.Address, amount *big.Int, to common.Address, synthParams IPortalSynthParams, permitData IPortalPermitData, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.SynthesizeRequestWithPermitPayNativeOverGsn(&_Router.TransactOpts, token, amount, to, synthParams, permitData, receipt)
}

// SynthesizeRequestWithPermitPayNative is a paid mutator transaction binding the contract method 0x4ec658e2.
//
// Solidity: function synthesizeRequestWithPermitPayNative(address token, uint256 amount, address to, (address,address,uint256) synthParams, (uint8,bytes32,bytes32,uint256,bool) permitData, (uint256,uint256,uint8,bytes32,bytes32) receipt) payable returns()
func (_Router *RouterTransactorSession) SynthesizeRequestWithPermitPayNative(token common.Address, amount *big.Int, to common.Address, synthParams IPortalSynthParams, permitData IPortalPermitData, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.SynthesizeRequestWithPermitPayNative(&_Router.TransactOpts, token, amount, to, synthParams, permitData, receipt)
}
func (_Router *RouterTransactorSession) SynthesizeRequestWithPermitPayNativeOverGsn(token common.Address, amount *big.Int, to common.Address, synthParams IPortalSynthParams, permitData IPortalPermitData, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.SynthesizeRequestWithPermitPayNativeOverGsn(&_Router.TransactOpts, token, amount, to, synthParams, permitData, receipt)
}

// TokenSynthesizeRequest is a paid mutator transaction binding the contract method 0xc89efd6b.
//
// Solidity: function tokenSynthesizeRequest(address token, uint256 amount, address to, (address,address,uint256) synthParams) returns()
func (_Router *RouterTransactor) TokenSynthesizeRequest(opts *bind.TransactOpts, token common.Address, amount *big.Int, to common.Address, synthParams IPortalSynthParams) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "tokenSynthesizeRequest", token, amount, to, synthParams)
}
func (_Router *RouterTransactor) TokenSynthesizeRequestOverGsn(opts *bind.TransactOpts, token common.Address, amount *big.Int, to common.Address, synthParams IPortalSynthParams) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "tokenSynthesizeRequest", token, amount, to, synthParams)
}

// TokenSynthesizeRequest is a paid mutator transaction binding the contract method 0xc89efd6b.
//
// Solidity: function tokenSynthesizeRequest(address token, uint256 amount, address to, (address,address,uint256) synthParams) returns()
func (_Router *RouterSession) TokenSynthesizeRequest(token common.Address, amount *big.Int, to common.Address, synthParams IPortalSynthParams) (*types.Transaction, error) {
	return _Router.Contract.TokenSynthesizeRequest(&_Router.TransactOpts, token, amount, to, synthParams)
}
func (_Router *RouterSession) TokenSynthesizeRequestOverGsn(token common.Address, amount *big.Int, to common.Address, synthParams IPortalSynthParams) (common.Hash, error) {
	return _Router.Contract.TokenSynthesizeRequestOverGsn(&_Router.TransactOpts, token, amount, to, synthParams)
}

// TokenSynthesizeRequest is a paid mutator transaction binding the contract method 0xc89efd6b.
//
// Solidity: function tokenSynthesizeRequest(address token, uint256 amount, address to, (address,address,uint256) synthParams) returns()
func (_Router *RouterTransactorSession) TokenSynthesizeRequest(token common.Address, amount *big.Int, to common.Address, synthParams IPortalSynthParams) (*types.Transaction, error) {
	return _Router.Contract.TokenSynthesizeRequest(&_Router.TransactOpts, token, amount, to, synthParams)
}
func (_Router *RouterTransactorSession) TokenSynthesizeRequestOverGsn(token common.Address, amount *big.Int, to common.Address, synthParams IPortalSynthParams) (common.Hash, error) {
	return _Router.Contract.TokenSynthesizeRequestOverGsn(&_Router.TransactOpts, token, amount, to, synthParams)
}

// TokenSynthesizeRequestToSolana is a paid mutator transaction binding the contract method 0x2e88ce08.
//
// Solidity: function tokenSynthesizeRequestToSolana(address token, uint256 amount, bytes32[] pubkeys, bytes1 txStateBump, uint256 chainId) returns()
func (_Router *RouterTransactor) TokenSynthesizeRequestToSolana(opts *bind.TransactOpts, token common.Address, amount *big.Int, pubkeys [][32]byte, txStateBump [1]byte, chainId *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "tokenSynthesizeRequestToSolana", token, amount, pubkeys, txStateBump, chainId)
}
func (_Router *RouterTransactor) TokenSynthesizeRequestToSolanaOverGsn(opts *bind.TransactOpts, token common.Address, amount *big.Int, pubkeys [][32]byte, txStateBump [1]byte, chainId *big.Int) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "tokenSynthesizeRequestToSolana", token, amount, pubkeys, txStateBump, chainId)
}

// TokenSynthesizeRequestToSolana is a paid mutator transaction binding the contract method 0x2e88ce08.
//
// Solidity: function tokenSynthesizeRequestToSolana(address token, uint256 amount, bytes32[] pubkeys, bytes1 txStateBump, uint256 chainId) returns()
func (_Router *RouterSession) TokenSynthesizeRequestToSolana(token common.Address, amount *big.Int, pubkeys [][32]byte, txStateBump [1]byte, chainId *big.Int) (*types.Transaction, error) {
	return _Router.Contract.TokenSynthesizeRequestToSolana(&_Router.TransactOpts, token, amount, pubkeys, txStateBump, chainId)
}
func (_Router *RouterSession) TokenSynthesizeRequestToSolanaOverGsn(token common.Address, amount *big.Int, pubkeys [][32]byte, txStateBump [1]byte, chainId *big.Int) (common.Hash, error) {
	return _Router.Contract.TokenSynthesizeRequestToSolanaOverGsn(&_Router.TransactOpts, token, amount, pubkeys, txStateBump, chainId)
}

// TokenSynthesizeRequestToSolana is a paid mutator transaction binding the contract method 0x2e88ce08.
//
// Solidity: function tokenSynthesizeRequestToSolana(address token, uint256 amount, bytes32[] pubkeys, bytes1 txStateBump, uint256 chainId) returns()
func (_Router *RouterTransactorSession) TokenSynthesizeRequestToSolana(token common.Address, amount *big.Int, pubkeys [][32]byte, txStateBump [1]byte, chainId *big.Int) (*types.Transaction, error) {
	return _Router.Contract.TokenSynthesizeRequestToSolana(&_Router.TransactOpts, token, amount, pubkeys, txStateBump, chainId)
}
func (_Router *RouterTransactorSession) TokenSynthesizeRequestToSolanaOverGsn(token common.Address, amount *big.Int, pubkeys [][32]byte, txStateBump [1]byte, chainId *big.Int) (common.Hash, error) {
	return _Router.Contract.TokenSynthesizeRequestToSolanaOverGsn(&_Router.TransactOpts, token, amount, pubkeys, txStateBump, chainId)
}

// TokenSynthesizeRequestWithPermit is a paid mutator transaction binding the contract method 0xd9ab9764.
//
// Solidity: function tokenSynthesizeRequestWithPermit(address token, uint256 amount, address to, (address,address,uint256) synthParams, (uint8,bytes32,bytes32,uint256,bool) permitData) returns()
func (_Router *RouterTransactor) TokenSynthesizeRequestWithPermit(opts *bind.TransactOpts, token common.Address, amount *big.Int, to common.Address, synthParams IPortalSynthParams, permitData IPortalPermitData) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "tokenSynthesizeRequestWithPermit", token, amount, to, synthParams, permitData)
}
func (_Router *RouterTransactor) TokenSynthesizeRequestWithPermitOverGsn(opts *bind.TransactOpts, token common.Address, amount *big.Int, to common.Address, synthParams IPortalSynthParams, permitData IPortalPermitData) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "tokenSynthesizeRequestWithPermit", token, amount, to, synthParams, permitData)
}

// TokenSynthesizeRequestWithPermit is a paid mutator transaction binding the contract method 0xd9ab9764.
//
// Solidity: function tokenSynthesizeRequestWithPermit(address token, uint256 amount, address to, (address,address,uint256) synthParams, (uint8,bytes32,bytes32,uint256,bool) permitData) returns()
func (_Router *RouterSession) TokenSynthesizeRequestWithPermit(token common.Address, amount *big.Int, to common.Address, synthParams IPortalSynthParams, permitData IPortalPermitData) (*types.Transaction, error) {
	return _Router.Contract.TokenSynthesizeRequestWithPermit(&_Router.TransactOpts, token, amount, to, synthParams, permitData)
}
func (_Router *RouterSession) TokenSynthesizeRequestWithPermitOverGsn(token common.Address, amount *big.Int, to common.Address, synthParams IPortalSynthParams, permitData IPortalPermitData) (common.Hash, error) {
	return _Router.Contract.TokenSynthesizeRequestWithPermitOverGsn(&_Router.TransactOpts, token, amount, to, synthParams, permitData)
}

// TokenSynthesizeRequestWithPermit is a paid mutator transaction binding the contract method 0xd9ab9764.
//
// Solidity: function tokenSynthesizeRequestWithPermit(address token, uint256 amount, address to, (address,address,uint256) synthParams, (uint8,bytes32,bytes32,uint256,bool) permitData) returns()
func (_Router *RouterTransactorSession) TokenSynthesizeRequestWithPermit(token common.Address, amount *big.Int, to common.Address, synthParams IPortalSynthParams, permitData IPortalPermitData) (*types.Transaction, error) {
	return _Router.Contract.TokenSynthesizeRequestWithPermit(&_Router.TransactOpts, token, amount, to, synthParams, permitData)
}
func (_Router *RouterTransactorSession) TokenSynthesizeRequestWithPermitOverGsn(token common.Address, amount *big.Int, to common.Address, synthParams IPortalSynthParams, permitData IPortalPermitData) (common.Hash, error) {
	return _Router.Contract.TokenSynthesizeRequestWithPermitOverGsn(&_Router.TransactOpts, token, amount, to, synthParams, permitData)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Router *RouterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "transferOwnership", newOwner)
}
func (_Router *RouterTransactor) TransferOwnershipOverGsn(opts *bind.TransactOpts, newOwner common.Address) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Router *RouterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Router.Contract.TransferOwnership(&_Router.TransactOpts, newOwner)
}
func (_Router *RouterSession) TransferOwnershipOverGsn(newOwner common.Address) (common.Hash, error) {
	return _Router.Contract.TransferOwnershipOverGsn(&_Router.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Router *RouterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Router.Contract.TransferOwnership(&_Router.TransactOpts, newOwner)
}
func (_Router *RouterTransactorSession) TransferOwnershipOverGsn(newOwner common.Address) (common.Hash, error) {
	return _Router.Contract.TransferOwnershipOverGsn(&_Router.TransactOpts, newOwner)
}

// UnsynthesizeRequest is a paid mutator transaction binding the contract method 0xd550f9e1.
//
// Solidity: function unsynthesizeRequest(address stoken, uint256 amount, address to, (address,address,uint256) synthParams) returns()
func (_Router *RouterTransactor) UnsynthesizeRequest(opts *bind.TransactOpts, stoken common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "unsynthesizeRequest", stoken, amount, to, synthParams)
}
func (_Router *RouterTransactor) UnsynthesizeRequestOverGsn(opts *bind.TransactOpts, stoken common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "unsynthesizeRequest", stoken, amount, to, synthParams)
}

// UnsynthesizeRequest is a paid mutator transaction binding the contract method 0xd550f9e1.
//
// Solidity: function unsynthesizeRequest(address stoken, uint256 amount, address to, (address,address,uint256) synthParams) returns()
func (_Router *RouterSession) UnsynthesizeRequest(stoken common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams) (*types.Transaction, error) {
	return _Router.Contract.UnsynthesizeRequest(&_Router.TransactOpts, stoken, amount, to, synthParams)
}
func (_Router *RouterSession) UnsynthesizeRequestOverGsn(stoken common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams) (common.Hash, error) {
	return _Router.Contract.UnsynthesizeRequestOverGsn(&_Router.TransactOpts, stoken, amount, to, synthParams)
}

// UnsynthesizeRequest is a paid mutator transaction binding the contract method 0xd550f9e1.
//
// Solidity: function unsynthesizeRequest(address stoken, uint256 amount, address to, (address,address,uint256) synthParams) returns()
func (_Router *RouterTransactorSession) UnsynthesizeRequest(stoken common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams) (*types.Transaction, error) {
	return _Router.Contract.UnsynthesizeRequest(&_Router.TransactOpts, stoken, amount, to, synthParams)
}
func (_Router *RouterTransactorSession) UnsynthesizeRequestOverGsn(stoken common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams) (common.Hash, error) {
	return _Router.Contract.UnsynthesizeRequestOverGsn(&_Router.TransactOpts, stoken, amount, to, synthParams)
}

// UnsynthesizeRequestPayNative is a paid mutator transaction binding the contract method 0xc51e547c.
//
// Solidity: function unsynthesizeRequestPayNative(address tokenSynth, uint256 amount, address to, (address,address,uint256) synthParams, (uint256,uint256,uint8,bytes32,bytes32) receipt) payable returns()
func (_Router *RouterTransactor) UnsynthesizeRequestPayNative(opts *bind.TransactOpts, tokenSynth common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "unsynthesizeRequestPayNative", tokenSynth, amount, to, synthParams, receipt)
}
func (_Router *RouterTransactor) UnsynthesizeRequestPayNativeOverGsn(opts *bind.TransactOpts, tokenSynth common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "unsynthesizeRequestPayNative", tokenSynth, amount, to, synthParams, receipt)
}

// UnsynthesizeRequestPayNative is a paid mutator transaction binding the contract method 0xc51e547c.
//
// Solidity: function unsynthesizeRequestPayNative(address tokenSynth, uint256 amount, address to, (address,address,uint256) synthParams, (uint256,uint256,uint8,bytes32,bytes32) receipt) payable returns()
func (_Router *RouterSession) UnsynthesizeRequestPayNative(tokenSynth common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.UnsynthesizeRequestPayNative(&_Router.TransactOpts, tokenSynth, amount, to, synthParams, receipt)
}
func (_Router *RouterSession) UnsynthesizeRequestPayNativeOverGsn(tokenSynth common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.UnsynthesizeRequestPayNativeOverGsn(&_Router.TransactOpts, tokenSynth, amount, to, synthParams, receipt)
}

// UnsynthesizeRequestPayNative is a paid mutator transaction binding the contract method 0xc51e547c.
//
// Solidity: function unsynthesizeRequestPayNative(address tokenSynth, uint256 amount, address to, (address,address,uint256) synthParams, (uint256,uint256,uint8,bytes32,bytes32) receipt) payable returns()
func (_Router *RouterTransactorSession) UnsynthesizeRequestPayNative(tokenSynth common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.UnsynthesizeRequestPayNative(&_Router.TransactOpts, tokenSynth, amount, to, synthParams, receipt)
}
func (_Router *RouterTransactorSession) UnsynthesizeRequestPayNativeOverGsn(tokenSynth common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.UnsynthesizeRequestPayNativeOverGsn(&_Router.TransactOpts, tokenSynth, amount, to, synthParams, receipt)
}

// UnsynthesizeRequestToSolana is a paid mutator transaction binding the contract method 0x0d0e5b6d.
//
// Solidity: function unsynthesizeRequestToSolana(address stoken, bytes32[] pubkeys, uint256 amount, uint256 chainId) returns()
func (_Router *RouterTransactor) UnsynthesizeRequestToSolana(opts *bind.TransactOpts, stoken common.Address, pubkeys [][32]byte, amount *big.Int, chainId *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "unsynthesizeRequestToSolana", stoken, pubkeys, amount, chainId)
}
func (_Router *RouterTransactor) UnsynthesizeRequestToSolanaOverGsn(opts *bind.TransactOpts, stoken common.Address, pubkeys [][32]byte, amount *big.Int, chainId *big.Int) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "unsynthesizeRequestToSolana", stoken, pubkeys, amount, chainId)
}

// UnsynthesizeRequestToSolana is a paid mutator transaction binding the contract method 0x0d0e5b6d.
//
// Solidity: function unsynthesizeRequestToSolana(address stoken, bytes32[] pubkeys, uint256 amount, uint256 chainId) returns()
func (_Router *RouterSession) UnsynthesizeRequestToSolana(stoken common.Address, pubkeys [][32]byte, amount *big.Int, chainId *big.Int) (*types.Transaction, error) {
	return _Router.Contract.UnsynthesizeRequestToSolana(&_Router.TransactOpts, stoken, pubkeys, amount, chainId)
}
func (_Router *RouterSession) UnsynthesizeRequestToSolanaOverGsn(stoken common.Address, pubkeys [][32]byte, amount *big.Int, chainId *big.Int) (common.Hash, error) {
	return _Router.Contract.UnsynthesizeRequestToSolanaOverGsn(&_Router.TransactOpts, stoken, pubkeys, amount, chainId)
}

// UnsynthesizeRequestToSolana is a paid mutator transaction binding the contract method 0x0d0e5b6d.
//
// Solidity: function unsynthesizeRequestToSolana(address stoken, bytes32[] pubkeys, uint256 amount, uint256 chainId) returns()
func (_Router *RouterTransactorSession) UnsynthesizeRequestToSolana(stoken common.Address, pubkeys [][32]byte, amount *big.Int, chainId *big.Int) (*types.Transaction, error) {
	return _Router.Contract.UnsynthesizeRequestToSolana(&_Router.TransactOpts, stoken, pubkeys, amount, chainId)
}
func (_Router *RouterTransactorSession) UnsynthesizeRequestToSolanaOverGsn(stoken common.Address, pubkeys [][32]byte, amount *big.Int, chainId *big.Int) (common.Hash, error) {
	return _Router.Contract.UnsynthesizeRequestToSolanaOverGsn(&_Router.TransactOpts, stoken, pubkeys, amount, chainId)
}

// RouterCrosschainPaymentEventIterator is returned from FilterCrosschainPaymentEvent and is used to iterate over the raw logs and unpacked data for CrosschainPaymentEvent events raised by the Router contract.
type RouterCrosschainPaymentEventIterator struct {
	Event *RouterCrosschainPaymentEvent // Event containing the contract specifics and raw log

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
func (it *RouterCrosschainPaymentEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterCrosschainPaymentEvent)
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
		it.Event = new(RouterCrosschainPaymentEvent)
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
func (it *RouterCrosschainPaymentEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterCrosschainPaymentEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterCrosschainPaymentEvent represents a CrosschainPaymentEvent event raised by the Router contract.
type RouterCrosschainPaymentEvent struct {
	UserFrom       common.Address
	Worker         common.Address
	ExecutionPrice *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCrosschainPaymentEvent is a free log retrieval operation binding the contract event 0x141cdd7ab02a0ea7399fa91a81781d764708704497d78344275c700a6caa0be6.
//
// Solidity: event CrosschainPaymentEvent(address indexed userFrom, address indexed worker, uint256 executionPrice)
func (_Router *RouterFilterer) FilterCrosschainPaymentEvent(opts *bind.FilterOpts, userFrom []common.Address, worker []common.Address) (*RouterCrosschainPaymentEventIterator, error) {

	var userFromRule []interface{}
	for _, userFromItem := range userFrom {
		userFromRule = append(userFromRule, userFromItem)
	}
	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "CrosschainPaymentEvent", userFromRule, workerRule)
	if err != nil {
		return nil, err
	}
	return &RouterCrosschainPaymentEventIterator{contract: _Router.contract, event: "CrosschainPaymentEvent", logs: logs, sub: sub}, nil
}

// WatchCrosschainPaymentEvent is a free log subscription operation binding the contract event 0x141cdd7ab02a0ea7399fa91a81781d764708704497d78344275c700a6caa0be6.
//
// Solidity: event CrosschainPaymentEvent(address indexed userFrom, address indexed worker, uint256 executionPrice)
func (_Router *RouterFilterer) WatchCrosschainPaymentEvent(opts *bind.WatchOpts, sink chan<- *RouterCrosschainPaymentEvent, userFrom []common.Address, worker []common.Address) (event.Subscription, error) {

	var userFromRule []interface{}
	for _, userFromItem := range userFrom {
		userFromRule = append(userFromRule, userFromItem)
	}
	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "CrosschainPaymentEvent", userFromRule, workerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterCrosschainPaymentEvent)
				if err := _Router.contract.UnpackLog(event, "CrosschainPaymentEvent", log); err != nil {
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

// ParseCrosschainPaymentEvent is a log parse operation binding the contract event 0x141cdd7ab02a0ea7399fa91a81781d764708704497d78344275c700a6caa0be6.
//
// Solidity: event CrosschainPaymentEvent(address indexed userFrom, address indexed worker, uint256 executionPrice)
func (_Router *RouterFilterer) ParseCrosschainPaymentEvent(log types.Log) (*RouterCrosschainPaymentEvent, error) {
	event := new(RouterCrosschainPaymentEvent)
	if err := _Router.contract.UnpackLog(event, "CrosschainPaymentEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Router contract.
type RouterOwnershipTransferredIterator struct {
	Event *RouterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RouterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterOwnershipTransferred)
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
		it.Event = new(RouterOwnershipTransferred)
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
func (it *RouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterOwnershipTransferred represents a OwnershipTransferred event raised by the Router contract.
type RouterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Router *RouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RouterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RouterOwnershipTransferredIterator{contract: _Router.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Router *RouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RouterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterOwnershipTransferred)
				if err := _Router.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Router *RouterFilterer) ParseOwnershipTransferred(log types.Log) (*RouterOwnershipTransferred, error) {
	event := new(RouterOwnershipTransferred)
	if err := _Router.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
