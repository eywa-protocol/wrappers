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

// RouterMetaData contains all meta data concerning the Router contract.
var RouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"portal\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"synthesis\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"curveProxy\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userFrom\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionPrice\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"}],\"name\":\"CrosschainPaymentEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_DELEGATED_SYNTHESIZE_REQUEST_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_DELEGATED_SYNTH_TRANSFER_REQUEST_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_DELEGATED_UNSYNTHESIZE_REQUEST_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_trustedWorker\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"transitData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structIPortal.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structIPortal.PermitData[]\",\"name\":\"permitData\",\"type\":\"tuple[]\"}],\"name\":\"batchSynthesizeRequestWithDataTransit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tokenReal\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenSynth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structISynthesis.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"executionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8[2]\",\"name\":\"v\",\"type\":\"uint8[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"r\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"s\",\"type\":\"bytes32[2]\"}],\"internalType\":\"structRouter.DelegatedCallReceipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"delegatedSynthTransferRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structIPortal.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"executionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8[2]\",\"name\":\"v\",\"type\":\"uint8[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"r\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"s\",\"type\":\"bytes32[2]\"}],\"internalType\":\"structRouter.DelegatedCallReceipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"delegatedTokenSynthesizeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structIPortal.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structIPortal.PermitData\",\"name\":\"permitData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"executionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8[2]\",\"name\":\"v\",\"type\":\"uint8[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"r\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"s\",\"type\":\"bytes32[2]\"}],\"internalType\":\"structRouter.DelegatedCallReceipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"delegatedTokenSynthesizeRequestWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stoken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"executionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8[2]\",\"name\":\"v\",\"type\":\"uint8[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"r\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"s\",\"type\":\"bytes32[2]\"}],\"internalType\":\"structRouter.DelegatedCallReceipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"delegatedUnsynthesizeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"emergencyUnburnRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"emergencyUnsyntesizeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remove\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinMintAmount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinDy\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"chain2address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"initialBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialChainID\",\"type\":\"uint256\"}],\"internalType\":\"structICurveProxy.MetaExchangeParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structICurveProxy.PermitData[]\",\"name\":\"permit\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address[3]\",\"name\":\"token\",\"type\":\"address[3]\"},{\"internalType\":\"uint256[3]\",\"name\":\"amount\",\"type\":\"uint256[3]\"}],\"name\":\"metaExchangeRequestVia3pool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addAtCrosschainPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinMintAmountC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lpIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addAtHubPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinMintAmountH\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialChainID\",\"type\":\"uint256\"}],\"internalType\":\"structICurveProxy.MetaMintEUSD\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structICurveProxy.PermitData[]\",\"name\":\"permit\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address[3]\",\"name\":\"token\",\"type\":\"address[3]\"},{\"internalType\":\"uint256[3]\",\"name\":\"amount\",\"type\":\"uint256[3]\"}],\"name\":\"mintEusdRequestVia3pool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"removeAtCrosschainPool\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinAmountC\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"removeAtHubPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountH\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"y\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinAmountH\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"internalType\":\"structICurveProxy.MetaRedeemEUSD\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structICurveProxy.PermitData\",\"name\":\"permit\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"redeemEusdRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"}],\"name\":\"removeTrustedWorker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"}],\"name\":\"setTrustedWorker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tokenReal\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structISynthesis.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"}],\"name\":\"synthTransferRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structIPortal.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"}],\"name\":\"tokenSynthesizeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"pubkeys\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes1\",\"name\":\"txStateBump\",\"type\":\"bytes1\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"tokenSynthesizeRequestToSolana\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structIPortal.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structIPortal.PermitData\",\"name\":\"permitData\",\"type\":\"tuple\"}],\"name\":\"tokenSynthesizeRequestWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stoken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"unsynthesizeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stoken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"pubkeys\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"unsynthesizeRequestToSolana\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162003295380380620032958339810160408190526200003591620001a0565b60408051808201825260048152634559574160e01b60208083019182528351808501855260018152603160f81b908201529151902060c08181527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc660e08190524660a081815286517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f8188018190528189019690965260608101939093526080808401929092523083820152865180840390910181529190920194859052805193019290922090915261010052600080546001600160a01b031916339081178255918291907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350600280546001600160a01b039485166001600160a01b031991821617909155600380549385169382169390931790925560018054919093169116179055620001ea565b80516001600160a01b03811681146200019b57600080fd5b919050565b600080600060608486031215620001b657600080fd5b620001c18462000183565b9250620001d16020850162000183565b9150620001e16040850162000183565b90509250925092565b60805160a05160c05160e051610100516130666200022f60003960006112440152600061128601526000611265015260006111f40152600061121d01526130666000f3fe608060405234801561001057600080fd5b50600436106101a95760003560e01c806372b90c02116100f9578063a74aeff711610097578063d7c7d41611610071578063d7c7d41614610349578063db1c62e51461035c578063e840f9d31461036f578063f2fde38b146103a257600080fd5b8063a74aeff71461031b578063c067de1714610323578063c7026a471461033657600080fd5b80637fe04792116100d35780637fe04792146102c75780638da5cb5b146102da578063905e4be0146102f55780639500125b1461030857600080fd5b806372b90c021461028e57806379da42b9146102a15780637ecebe00146102b457600080fd5b80633d5ac6cb11610166578063604477a611610140578063604477a61461024d578063605ae5441461026057806361bf99ac14610273578063715018a61461028657600080fd5b80633d5ac6cb1461021457806343961668146102275780634b7568fb1461023a57600080fd5b80630b414a0d146101ae57806320eada55146101c3578063259ab8a4146101de57806326363d22146101f15780633644e515146101f95780633cce304514610201575b600080fd5b6101c16101bc366004611c1b565b6103b5565b005b6101cb610446565b6040519081526020015b60405180910390f35b6101c16101ec366004611cbf565b61046e565b6101cb610596565b6101cb6105a5565b6101c161020f366004611d70565b6105b4565b6101c1610222366004611df9565b6106ac565b6101c1610235366004611e8c565b610773565b6101c1610248366004611efa565b6107af565b6101c161025b366004611feb565b610837565b6101c161026e3660046120fa565b61094f565b6101c1610281366004612188565b610a17565b6101c1610b43565b6101c161029c3660046121d8565b610bc0565b6101c16102af3660046121f5565b610c0b565b6101cb6102c23660046121d8565b610c96565b6101c16102d53660046121d8565b610cb6565b6000546040516001600160a01b0390911681526020016101d5565b6101c161030336600461227e565b610d04565b6101c1610316366004611e8c565b610db4565b6101cb610df0565b6101c1610331366004612312565b610dff565b6101c161034436600461237c565b610f17565b6101c16103573660046123d5565b610f90565b6101c161036a366004612459565b611050565b61039261037d3660046121d8565b60046020526000908152604090205460ff1681565b60405190151581526020016101d5565b6101c16103b03660046121d8565b6110a6565b6001546103d390869086906001600160a01b031660808b0135611190565b600154604051635fcc4a3f60e01b81526001600160a01b0390911690635fcc4a3f9061040b908a908a9088908890889060040161255d565b600060405180830381600087803b15801561042557600080fd5b505af1158015610439573d6000803e3d6000fd5b5050505050505050505050565b6040516020016104559061263b565b6040516020818303038152906040528051906020012081565b6001600160a01b03851663d505accf843061048f60a08601608087016126e5565b610499578761049d565b6000195b60608601356104af6020880188612702565b876020013588604001356040518863ffffffff1660e01b81526004016104db979695949392919061271d565b600060405180830381600087803b1580156104f557600080fd5b505af1158015610509573d6000803e3d6000fd5b5050600254610527925087915085906001600160a01b031687611190565b600254604051630529fc9f60e31b81526001600160a01b039091169063294fe4f89061055d9088908890889088906004016127ad565b600060405180830381600087803b15801561057757600080fd5b505af115801561058b573d6000803e3d6000fd5b505050505050505050565b604051602001610455906127d9565b60006105af6111f0565b905090565b6105c086863085611190565b60035460405163095ea7b360e01b81526001600160a01b039182166004820152602481018490529087169063095ea7b3906044016020604051808303816000875af1158015610613573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106379190612888565b5060035460405163966e0fe960e01b81526001600160a01b039091169063966e0fe990610672908990899089908990899089906004016128db565b600060405180830381600087803b15801561068c57600080fd5b505af11580156106a0573d6000803e3d6000fd5b50505050505050505050565b6106de868585886040516020016106c29061291d565b60405160208183030381529060405280519060200120866112aa565b6003546000906106fe908890889088906001600160a01b031686356115c7565b60035460405163b04640b760e01b81529192506001600160a01b03169063b04640b790610737908b9085908a908a908a906004016129f8565b600060405180830381600087803b15801561075157600080fd5b505af1158015610765573d6000803e3d6000fd5b505050505050505050505050565b600254604051630872c2cd60e31b81526001600160a01b039091169063439616689061040b908a908a908a908a908a908a908a90600401612a5a565b6002546107c990859084906001600160a01b031686611190565b600254604051630529fc9f60e31b81526001600160a01b039091169063294fe4f8906107ff9087908790879087906004016127ad565b600060405180830381600087803b15801561081957600080fd5b505af115801561082d573d6000803e3d6000fd5b5050505050505050565b60005b8a8110156108cb5760008a8a8381811061085657610856612a99565b9050602002013511156108b9576108b98c8c8381811061087857610878612a99565b905060200201602081019061088d91906121d8565b6002548a906001600160a01b03168d8d868181106108ad576108ad612a99565b90506020020135611190565b806108c381612ac5565b91505061083a565b50600254604051630b505d3760e41b81526001600160a01b039091169063b505d37090610910908e908e908e908e908e908b908f908f908f908e908e90600401612b47565b600060405180830381600087803b15801561092a57600080fd5b505af115801561093e573d6000803e3d6000fd5b505050505050505050505050505050565b60005b60038110156109de57600082826003811061096f5761096f612a99565b602002013511156109cc576109cc83826003811061098f5761098f612a99565b6020020160208101906109a291906121d8565b60025486906001600160a01b03168585600381106109c2576109c2612a99565b6020020135611190565b806109d681612ac5565b915050610952565b506001546040516301c2bc7d60e51b81526001600160a01b03909116906338578fa0906106729089908990899088908890600401612c52565b60035460405163bca7382360e01b8152600481018790526000916001600160a01b03169063bca7382390602401602060405180830381865afa158015610a61573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a859190612d1f565b9050610a9381853088611190565b60035460405163095ea7b360e01b81526001600160a01b039182166004820152602481018790529082169063095ea7b3906044016020604051808303816000875af1158015610ae6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b0a9190612888565b5060035460405163b04640b760e01b81526001600160a01b039091169063b04640b79061067290899089908990899089906004016129f8565b6000546001600160a01b03163314610b765760405162461bcd60e51b8152600401610b6d90612d3c565b60405180910390fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b03163314610bea5760405162461bcd60e51b8152600401610b6d90612d3c565b6001600160a01b03166000908152600460205260409020805460ff19169055565b60005b6003811015610c5d576000828260038110610c2b57610c2b612a99565b60200201351115610c4b57610c4b83826003811061098f5761098f612a99565b80610c5581612ac5565b915050610c0e565b50600154604051634638ef3f60e11b81526001600160a01b0390911690638c71de7e906106729089908990899088908890600401612d71565b6001600160a01b0381166000908152600560205260408120545b92915050565b6000546001600160a01b03163314610ce05760405162461bcd60e51b8152600401610b6d90612d3c565b6001600160a01b03166000908152600460205260409020805460ff19166001179055565b610d1a8887878a6040516020016106c29061263b565b600354600090610d3a908a908a908a906001600160a01b031686356115c7565b6003546040516317ddce6f60e31b81529192506001600160a01b03169063beee737890610d77908c9085908c908c908c908c908c90600401612f08565b600060405180830381600087803b158015610d9157600080fd5b505af1158015610da5573d6000803e3d6000fd5b50505050505050505050505050565b600354604051639500125b60e01b81526001600160a01b0390911690639500125b9061040b908a908a908a908a908a908a908a90600401612a5a565b6040516020016104559061291d565b610e218685610e1160208701876121d8565b886040516020016106c2906127d9565b6001600160a01b03861663d505accf8530610e4260a08701608088016126e5565b610e4c5788610e50565b6000195b6060870135610e626020890189612702565b886020013589604001356040518863ffffffff1660e01b8152600401610e8e979695949392919061271d565b600060405180830381600087803b158015610ea857600080fd5b505af1158015610ebc573d6000803e3d6000fd5b505060025460009250610ee091508890889088906001600160a01b031686356116aa565b600254604051630529fc9f60e31b81529192506001600160a01b03169063294fe4f89061040b908a9085908a908a906004016127ad565b610f398584610f2960208601866121d8565b876040516020016106c2906127d9565b600254600090610f59908790879087906001600160a01b031686356116aa565b600254604051630529fc9f60e31b81529192506001600160a01b03169063294fe4f8906106729089908590899089906004016127ad565b610f9c87863089611190565b60035460405163095ea7b360e01b81526001600160a01b039182166004820152602481018890529088169063095ea7b3906044016020604051808303816000875af1158015610fef573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110139190612888565b506003546040516317ddce6f60e31b81526001600160a01b039091169063beee73789061040b908a908a908a908a908a908a908a90600401612f08565b60025461106a90889087906001600160a01b031689611190565b600254604051632b17790f60e21b81526001600160a01b039091169063ac5de43c9061040b908a908a908a908a908a908a908a90600401612f4a565b6000546001600160a01b031633146110d05760405162461bcd60e51b8152600401610b6d90612d3c565b6001600160a01b0381166111355760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610b6d565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b1790526111ea908590611716565b50505050565b60007f000000000000000000000000000000000000000000000000000000000000000046141561123f57507f000000000000000000000000000000000000000000000000000000000000000090565b6105af7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000006117ed565b60006113627f72933cec627dd074a3496d54ab05f9976222e324124cfc405e828350b54cda91888888888735896112e085611837565b6040805160208181019a909a526bffffffffffffffffffffffff196060998a1b81169282019290925296881b811660548801529490961b9093166068850152607c840191909152609c83015260bc82015260dc8101919091529084013560fc82015261011c015b6040516020818303038152906040528051906020012061185f565b905060006114427f0e668f7b944092ca963fc5538a4911e923a0140f636afd096a882dfef9838ab28989898988358a61139a85611837565b60208c01356113af60608e0160408f01612702565b60408051602081019b909b526060998a1b6bffffffffffffffffffffffff19908116918c019190915297891b881660548b01529590971b9095166068880152607c870192909252609c86015260bc85015260dc84019190915260fc83019190915260f81b6001600160f81b03191661011c820152608085013561011d82015260c085013561013d82015261015d01611347565b905060006114708361145a6060870160408801612702565b608087013560c0880160005b60200201356118ad565b90506000611498836114886080880160608901612702565b60a088013560c089016001611466565b6001600160a01b03831660009081526004602052604090205490915060ff166115115760405162461bcd60e51b815260206004820152602560248201527f526f757465723a20696e76616c6964207369676e61747572652066726f6d207760448201526437b935b2b960d91b6064820152608401610b6d565b886001600160a01b0316816001600160a01b0316146115805760405162461bcd60e51b815260206004820152602560248201527f526f757465723a20696e76616c6964207369676e61747572652066726f6d207360448201526432b73232b960d91b6064820152608401610b6d565b84602001354211156106a05760405162461bcd60e51b815260206004820152601060248201526f526f757465723a20646561646c696e6560801b6044820152606401610b6d565b60006115d586853385611190565b6115df8286612f9e565b60405163095ea7b360e01b81526001600160a01b038581166004830152602482018390529192509087169063095ea7b3906044016020604051808303816000875af1158015611632573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116569190612888565b50604080516001600160a01b038881168252602082018590523392908716917f8adfc333732362694bf9044b570e1ae5d25a5b143b8aa3c267aee92adae48bbe91015b60405180910390a395945050505050565b60006116b886853385611190565b6116c28286612f9e565b90506116d086858584611190565b604080516001600160a01b038881168252602082018590523392908716917f8adfc333732362694bf9044b570e1ae5d25a5b143b8aa3c267aee92adae48bbe9101611699565b600061176b826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316611a569092919063ffffffff16565b8051909150156117e857808060200190518101906117899190612888565b6117e85760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610b6d565b505050565b6040805160208101859052908101839052606081018290524660808201523060a082015260009060c0016040516020818303038152906040528051906020012090505b9392505050565b6001600160a01b03811660009081526005602052604090208054600181018255905b50919050565b6000610cb061186c6111f0565b8360405161190160f01b6020820152602281018390526042810182905260009060620160405160208183030381529060405280519060200120905092915050565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a082111561192a5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610b6d565b8360ff16601b148061193f57508360ff16601c145b6119965760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610b6d565b6040805160008082526020820180845288905260ff871692820192909252606081018590526080810184905260019060a0016020604051602081039080840390855afa1580156119ea573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116611a4d5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610b6d565b95945050505050565b6060611a658484600085611a6d565b949350505050565b606082471015611ace5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610b6d565b843b611b1c5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610b6d565b600080866001600160a01b03168587604051611b389190612fe1565b60006040518083038185875af1925050503d8060008114611b75576040519150601f19603f3d011682016040523d82523d6000602084013e611b7a565b606091505b5091509150611b8a828286611b95565b979650505050505050565b60608315611ba4575081611830565b825115611bb45782518084602001fd5b8160405162461bcd60e51b8152600401610b6d9190612ffd565b6000610100828403121561185957600080fd5b600060a0828403121561185957600080fd5b6001600160a01b0381168114611c0857600080fd5b50565b8035611c1681611bf3565b919050565b6000806000806000806000610240888a031215611c3757600080fd5b611c418989611bce565b9650611c51896101008a01611be1565b95506101a0880135611c6281611bf3565b94506101c0880135611c7381611bf3565b93506101e0880135611c8481611bf3565b9250610200880135611c9581611bf3565b80925050610220880135905092959891949750929550565b60006080828403121561185957600080fd5b60008060008060006101808688031215611cd857600080fd5b8535611ce381611bf3565b9450602086013593506040860135611cfa81611bf3565b9250611d098760608801611cad565b9150611d188760e08801611be1565b90509295509295909350565b60008083601f840112611d3657600080fd5b50813567ffffffffffffffff811115611d4e57600080fd5b6020830191508360208260051b8501011115611d6957600080fd5b9250929050565b60008060008060008060a08789031215611d8957600080fd5b8635611d9481611bf3565b95506020870135611da481611bf3565b9450604087013567ffffffffffffffff811115611dc057600080fd5b611dcc89828a01611d24565b979a9699509760608101359660809091013595509350505050565b60006060828403121561185957600080fd5b6000806000806000806000610200888a031215611e1557600080fd5b873596506020880135611e2781611bf3565b9550604088013594506060880135611e3e81611bf3565b93506080880135611e4e81611bf3565b9250611e5d8960a08a01611de7565b9150611e6d896101008a01611bce565b905092959891949750929550565b803560ff81168114611c1657600080fd5b600080600080600080600060e0888a031215611ea757600080fd5b873596506020880135611eb981611bf3565b95506040880135611ec981611bf3565b945060608801359350611ede60808901611e7b565b925060a0880135915060c0880135905092959891949750929550565b60008060008060e08587031215611f1057600080fd5b8435611f1b81611bf3565b9350602085013592506040850135611f3281611bf3565b9150611f418660608701611cad565b905092959194509250565b80356001600160e01b031981168114611c1657600080fd5b60008083601f840112611f7657600080fd5b50813567ffffffffffffffff811115611f8e57600080fd5b602083019150836020828501011115611d6957600080fd5b60008083601f840112611fb857600080fd5b50813567ffffffffffffffff811115611fd057600080fd5b60208301915083602060a083028501011115611d6957600080fd5b60008060008060008060008060008060006101408c8e03121561200d57600080fd5b67ffffffffffffffff808d35111561202457600080fd5b6120318e8e358f01611d24565b909c509a5060208d013581101561204757600080fd5b6120578e60208f01358f01611d24565b909a50985061206860408e01611c0b565b975061207660608e01611f4c565b96508060808e0135111561208957600080fd5b6120998e60808f01358f01611f64565b90965094506120ab8e60a08f01611cad565b9350806101208e013511156120bf57600080fd5b506120d18d6101208e01358e01611fa6565b81935080925050509295989b509295989b9093969950565b8060608101831015610cb057600080fd5b600080600080600080610200878903121561211457600080fd5b61211e8888611bce565b955061010087013567ffffffffffffffff81111561213b57600080fd5b61214789828a01611fa6565b90965094505061012087013561215c81611bf3565b925061216c8861014089016120e9565b915061217c886101a089016120e9565b90509295509295509295565b600080600080600060e086880312156121a057600080fd5b853594506020860135935060408601356121b981611bf3565b925060608601356121c981611bf3565b9150611d188760808801611de7565b6000602082840312156121ea57600080fd5b813561183081611bf3565b60008060008060008086880361030081121561221057600080fd5b6102008082121561222057600080fd5b889750870135905067ffffffffffffffff81111561223d57600080fd5b61224989828a01611fa6565b90965094505061022087013561225e81611bf3565b925061226e8861024089016120e9565b915061217c886102a089016120e9565b6000806000806000806000806101e0898b03121561229b57600080fd5b88356122a681611bf3565b97506020890135965060408901356122bd81611bf3565b955060608901356122cd81611bf3565b945060808901356122dd81611bf3565b935060a08901356122ed81611bf3565b925060c089013591506123038a60e08b01611bce565b90509295985092959890939650565b600080600080600080610280878903121561232c57600080fd5b863561233781611bf3565b955060208701359450604087013561234e81611bf3565b935061235d8860608901611cad565b925061236c8860e08901611be1565b915061217c886101808901611bce565b60008060008060006101e0868803121561239557600080fd5b85356123a081611bf3565b94506020860135935060408601356123b781611bf3565b92506123c68760608801611cad565b9150611d188760e08801611bce565b600080600080600080600060e0888a0312156123f057600080fd5b87356123fb81611bf3565b965060208801359550604088013561241281611bf3565b9450606088013561242281611bf3565b9350608088013561243281611bf3565b925060a088013561244281611bf3565b8092505060c0880135905092959891949750929550565b600080600080600080600060c0888a03121561247457600080fd5b873561247f81611bf3565b965060208801359550604088013561249681611bf3565b9450606088013567ffffffffffffffff8111156124b257600080fd5b6124be8a828b01611d24565b90955093505060808801356001600160f81b0319811681146124df57600080fd5b8092505060a0880135905092959891949750929550565b8035600f81900b8114611c1657600080fd5b8015158114611c0857600080fd5b60ff61252182611e7b565b168252602081013560208301526040810135604083015260608101356060830152608081013561255081612508565b8015156080840152505050565b6102008101863561256d81611bf3565b6001600160a01b03168252612584602088016124f6565b600f0b60208301526040870135604083015260608701356125a481611bf3565b6001600160a01b03166060830152608087810135908301526125c860a088016124f6565b6125d760a0840182600f0b9052565b5060c087013560c08301526125ee60e08801611c0b565b6001600160a01b031660e083015261260a610100830187612516565b6001600160a01b0385166101a08301526001600160a01b0384166101c0830152826101e08301529695505050505050565b7f64656c656761746564556e73796e74686573697a65526571756573742861646481527f726573732c75696e743235362c616464726573732c616464726573732c61646460208201527f726573732c616464726573732c75696e743235362c5b75696e743235362c756960408201527f6e743235362c75696e74385b325d2c627974657333325b325d2c627974657333606082015265325b325d5d2960d01b608082015260860190565b6000602082840312156126f757600080fd5b813561183081612508565b60006020828403121561271457600080fd5b61183082611e7b565b6001600160a01b0397881681529590961660208601526040850193909352606084019190915260ff16608083015260a082015260c081019190915260e00190565b803561276981611bf3565b6001600160a01b03908116835260208201359061278582611bf3565b908116602084015260408201359061279c82611bf3565b166040830152606090810135910152565b6001600160a01b038581168252602082018590528316604082015260e08101611a4d606083018461275e565b7f64656c656761746564546f6b656e53796e74686573697a65526571756573742881527f616464726573732c75696e743235362c616464726573732c5b6164647265737360208201527f2c616464726573732c616464726573732c75696e743235365d2c5b75696e743260408201527f35362c75696e743235362c75696e74385b325d2c627974657333325b325d2c6260608201526a7974657333325b325d5d2960a81b6080820152608b0190565b60006020828403121561289a57600080fd5b815161183081612508565b81835260006001600160fb1b038311156128be57600080fd5b8260051b8083602087013760009401602001938452509192915050565b6001600160a01b0387811682528616602082015260a06040820181905260009061290890830186886128a5565b60608301949094525060800152949350505050565b7f64656c656761746564546f6b656e53796e74686573697a65526571756573742881527f64656c65676174656453796e74685472616e736665725265717565737428627960208201527f74657333322c616464726573732c75696e743235362c616464726573732c616460408201527f64726573732c5b616464726573732c616464726573732c75696e743235365d2c60608201527f5b75696e743235362c75696e743235362c75696e74385b325d2c627974657333608082015270325b325d2c627974657333325b325d5d2960781b60a082015260b10190565b858152602081018590526001600160a01b038481166040830152838116606083015260e08201908335612a2a81611bf3565b811660808401526020840135612a3f81611bf3565b1660a08301526040929092013560c090910152949350505050565b9687526001600160a01b039586166020880152939094166040860152606085019190915260ff16608084015260a083019190915260c082015260e00190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415612ad957612ad9612aaf565b5060010190565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b81835260208301925060008160005b84811015612b3d57612b2a8683612516565b60a0958601959190910190600101612b18565b5093949350505050565b61014080825281018b9052600061016082018d825b8e811015612b8d578135612b6f81611bf3565b6001600160a01b031683526020928301929190910190600101612b5c565b50508281036020840152612ba2818c8e6128a5565b6001600160a01b038b1660408501529050612bc0606084018a61275e565b6001600160e01b0319881660e0840152828103610100840152612be4818789612ae0565b9050828103610120840152612bfa818587612b09565b9e9d5050505050505050505050505050565b8060005b60038110156111ea578135612c2481611bf3565b6001600160a01b031684526020938401939190910190600101612c10565b6060818337506000606082015250565b60006101e08735612c6281611bf3565b6001600160a01b039081168452602089810135908501526040808a013590850152606089013590612c9282611bf3565b16606084015260808881013590840152612cae60a08901611c0b565b6001600160a01b031660a0840152612cc860c08901611c0b565b6001600160a01b03811660c08501525060e088013560e084015280610100840152612cf68184018789612b09565b915050612d07610120830185612c0c565b612d15610180830184612c42565b9695505050505050565b600060208284031215612d3157600080fd5b815161183081611bf3565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b60006102e0612d9083612d838a611c0b565b6001600160a01b03169052565b612d9c60208901611c0b565b6001600160a01b03166020840152612db660408901611c0b565b6001600160a01b0316604084015260608881013590840152612dda608089016124f6565b612de96080850182600f0b9052565b50612df660a089016124f6565b612e0560a0850182600f0b9052565b5060c088013560c0840152612e1c60e089016124f6565b612e2b60e0850182600f0b9052565b506101008881013590840152610120612e45818a01611c0b565b6001600160a01b031690840152610140612e60898201611c0b565b6001600160a01b031690840152610160612e7b898201611c0b565b6001600160a01b031690840152610180612e96898201611c0b565b6001600160a01b0316908401526101a088810135908401526101c0612ebc818a01611c0b565b6001600160a01b0316908401526101e088810135908401526102008301819052612ee98184018789612b09565b915050612efa610220830185612c0c565b612d15610280830184612c42565b6001600160a01b0397881681526020810196909652938616604086015291851660608501528416608084015290921660a082015260c081019190915260e00190565b6001600160a01b038881168252602082018890528616604082015260c060608201819052600090612f7e90830186886128a5565b6001600160f81b03199490941660808301525060a0015295945050505050565b600082821015612fb057612fb0612aaf565b500390565b60005b83811015612fd0578181015183820152602001612fb8565b838111156111ea5750506000910152565b60008251612ff3818460208701612fb5565b9190910192915050565b602081526000825180602084015261301c816040850160208701612fb5565b601f01601f1916919091016040019291505056fea264697066735822122007442c20a67e525a63372307cbc88dadbf58d36a29099208da135b1d46b4417a64736f6c634300080a0033",
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

// DELEGATEDSYNTHESIZEREQUESTTYPEHASH is a free data retrieval call binding the contract method 0x26363d22.
//
// Solidity: function _DELEGATED_SYNTHESIZE_REQUEST_TYPEHASH() view returns(bytes32)
func (_Router *RouterCaller) DELEGATEDSYNTHESIZEREQUESTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "_DELEGATED_SYNTHESIZE_REQUEST_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELEGATEDSYNTHESIZEREQUESTTYPEHASH is a free data retrieval call binding the contract method 0x26363d22.
//
// Solidity: function _DELEGATED_SYNTHESIZE_REQUEST_TYPEHASH() view returns(bytes32)
func (_Router *RouterSession) DELEGATEDSYNTHESIZEREQUESTTYPEHASH() ([32]byte, error) {
	return _Router.Contract.DELEGATEDSYNTHESIZEREQUESTTYPEHASH(&_Router.CallOpts)
}

// DELEGATEDSYNTHESIZEREQUESTTYPEHASH is a free data retrieval call binding the contract method 0x26363d22.
//
// Solidity: function _DELEGATED_SYNTHESIZE_REQUEST_TYPEHASH() view returns(bytes32)
func (_Router *RouterCallerSession) DELEGATEDSYNTHESIZEREQUESTTYPEHASH() ([32]byte, error) {
	return _Router.Contract.DELEGATEDSYNTHESIZEREQUESTTYPEHASH(&_Router.CallOpts)
}

// DELEGATEDSYNTHTRANSFERREQUESTTYPEHASH is a free data retrieval call binding the contract method 0xa74aeff7.
//
// Solidity: function _DELEGATED_SYNTH_TRANSFER_REQUEST_TYPEHASH() view returns(bytes32)
func (_Router *RouterCaller) DELEGATEDSYNTHTRANSFERREQUESTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "_DELEGATED_SYNTH_TRANSFER_REQUEST_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELEGATEDSYNTHTRANSFERREQUESTTYPEHASH is a free data retrieval call binding the contract method 0xa74aeff7.
//
// Solidity: function _DELEGATED_SYNTH_TRANSFER_REQUEST_TYPEHASH() view returns(bytes32)
func (_Router *RouterSession) DELEGATEDSYNTHTRANSFERREQUESTTYPEHASH() ([32]byte, error) {
	return _Router.Contract.DELEGATEDSYNTHTRANSFERREQUESTTYPEHASH(&_Router.CallOpts)
}

// DELEGATEDSYNTHTRANSFERREQUESTTYPEHASH is a free data retrieval call binding the contract method 0xa74aeff7.
//
// Solidity: function _DELEGATED_SYNTH_TRANSFER_REQUEST_TYPEHASH() view returns(bytes32)
func (_Router *RouterCallerSession) DELEGATEDSYNTHTRANSFERREQUESTTYPEHASH() ([32]byte, error) {
	return _Router.Contract.DELEGATEDSYNTHTRANSFERREQUESTTYPEHASH(&_Router.CallOpts)
}

// DELEGATEDUNSYNTHESIZEREQUESTTYPEHASH is a free data retrieval call binding the contract method 0x20eada55.
//
// Solidity: function _DELEGATED_UNSYNTHESIZE_REQUEST_TYPEHASH() view returns(bytes32)
func (_Router *RouterCaller) DELEGATEDUNSYNTHESIZEREQUESTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "_DELEGATED_UNSYNTHESIZE_REQUEST_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELEGATEDUNSYNTHESIZEREQUESTTYPEHASH is a free data retrieval call binding the contract method 0x20eada55.
//
// Solidity: function _DELEGATED_UNSYNTHESIZE_REQUEST_TYPEHASH() view returns(bytes32)
func (_Router *RouterSession) DELEGATEDUNSYNTHESIZEREQUESTTYPEHASH() ([32]byte, error) {
	return _Router.Contract.DELEGATEDUNSYNTHESIZEREQUESTTYPEHASH(&_Router.CallOpts)
}

// DELEGATEDUNSYNTHESIZEREQUESTTYPEHASH is a free data retrieval call binding the contract method 0x20eada55.
//
// Solidity: function _DELEGATED_UNSYNTHESIZE_REQUEST_TYPEHASH() view returns(bytes32)
func (_Router *RouterCallerSession) DELEGATEDUNSYNTHESIZEREQUESTTYPEHASH() ([32]byte, error) {
	return _Router.Contract.DELEGATEDUNSYNTHESIZEREQUESTTYPEHASH(&_Router.CallOpts)
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

// BatchSynthesizeRequestWithDataTransit is a paid mutator transaction binding the contract method 0x604477a6.
//
// Solidity: function batchSynthesizeRequestWithDataTransit(address[] token, uint256[] amount, address from, bytes4 selector, bytes transitData, (address,address,address,uint256) synthParams, (uint8,bytes32,bytes32,uint256,bool)[] permitData) returns()
func (_Router *RouterTransactor) BatchSynthesizeRequestWithDataTransit(opts *bind.TransactOpts, token []common.Address, amount []*big.Int, from common.Address, selector [4]byte, transitData []byte, synthParams IPortalSynthParams, permitData []IPortalPermitData) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "batchSynthesizeRequestWithDataTransit", token, amount, from, selector, transitData, synthParams, permitData)
}
func (_Router *RouterTransactor) BatchSynthesizeRequestWithDataTransitOverGsn(opts *bind.TransactOpts, token []common.Address, amount []*big.Int, from common.Address, selector [4]byte, transitData []byte, synthParams IPortalSynthParams, permitData []IPortalPermitData) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "batchSynthesizeRequestWithDataTransit", token, amount, from, selector, transitData, synthParams, permitData)
}

// BatchSynthesizeRequestWithDataTransit is a paid mutator transaction binding the contract method 0x604477a6.
//
// Solidity: function batchSynthesizeRequestWithDataTransit(address[] token, uint256[] amount, address from, bytes4 selector, bytes transitData, (address,address,address,uint256) synthParams, (uint8,bytes32,bytes32,uint256,bool)[] permitData) returns()
func (_Router *RouterSession) BatchSynthesizeRequestWithDataTransit(token []common.Address, amount []*big.Int, from common.Address, selector [4]byte, transitData []byte, synthParams IPortalSynthParams, permitData []IPortalPermitData) (*types.Transaction, error) {
	return _Router.Contract.BatchSynthesizeRequestWithDataTransit(&_Router.TransactOpts, token, amount, from, selector, transitData, synthParams, permitData)
}
func (_Router *RouterSession) BatchSynthesizeRequestWithDataTransitOverGsn(token []common.Address, amount []*big.Int, from common.Address, selector [4]byte, transitData []byte, synthParams IPortalSynthParams, permitData []IPortalPermitData) (common.Hash, error) {
	return _Router.Contract.BatchSynthesizeRequestWithDataTransitOverGsn(&_Router.TransactOpts, token, amount, from, selector, transitData, synthParams, permitData)
}

// BatchSynthesizeRequestWithDataTransit is a paid mutator transaction binding the contract method 0x604477a6.
//
// Solidity: function batchSynthesizeRequestWithDataTransit(address[] token, uint256[] amount, address from, bytes4 selector, bytes transitData, (address,address,address,uint256) synthParams, (uint8,bytes32,bytes32,uint256,bool)[] permitData) returns()
func (_Router *RouterTransactorSession) BatchSynthesizeRequestWithDataTransit(token []common.Address, amount []*big.Int, from common.Address, selector [4]byte, transitData []byte, synthParams IPortalSynthParams, permitData []IPortalPermitData) (*types.Transaction, error) {
	return _Router.Contract.BatchSynthesizeRequestWithDataTransit(&_Router.TransactOpts, token, amount, from, selector, transitData, synthParams, permitData)
}
func (_Router *RouterTransactorSession) BatchSynthesizeRequestWithDataTransitOverGsn(token []common.Address, amount []*big.Int, from common.Address, selector [4]byte, transitData []byte, synthParams IPortalSynthParams, permitData []IPortalPermitData) (common.Hash, error) {
	return _Router.Contract.BatchSynthesizeRequestWithDataTransitOverGsn(&_Router.TransactOpts, token, amount, from, selector, transitData, synthParams, permitData)
}

// DelegatedSynthTransferRequest is a paid mutator transaction binding the contract method 0x3d5ac6cb.
//
// Solidity: function delegatedSynthTransferRequest(bytes32 tokenReal, address tokenSynth, uint256 amount, address from, address to, (address,address,uint256) synthParams, (uint256,uint256,uint8[2],bytes32[2],bytes32[2]) receipt) returns()
func (_Router *RouterTransactor) DelegatedSynthTransferRequest(opts *bind.TransactOpts, tokenReal [32]byte, tokenSynth common.Address, amount *big.Int, from common.Address, to common.Address, synthParams ISynthesisSynthParams, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "delegatedSynthTransferRequest", tokenReal, tokenSynth, amount, from, to, synthParams, receipt)
}
func (_Router *RouterTransactor) DelegatedSynthTransferRequestOverGsn(opts *bind.TransactOpts, tokenReal [32]byte, tokenSynth common.Address, amount *big.Int, from common.Address, to common.Address, synthParams ISynthesisSynthParams, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "delegatedSynthTransferRequest", tokenReal, tokenSynth, amount, from, to, synthParams, receipt)
}

// DelegatedSynthTransferRequest is a paid mutator transaction binding the contract method 0x3d5ac6cb.
//
// Solidity: function delegatedSynthTransferRequest(bytes32 tokenReal, address tokenSynth, uint256 amount, address from, address to, (address,address,uint256) synthParams, (uint256,uint256,uint8[2],bytes32[2],bytes32[2]) receipt) returns()
func (_Router *RouterSession) DelegatedSynthTransferRequest(tokenReal [32]byte, tokenSynth common.Address, amount *big.Int, from common.Address, to common.Address, synthParams ISynthesisSynthParams, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.DelegatedSynthTransferRequest(&_Router.TransactOpts, tokenReal, tokenSynth, amount, from, to, synthParams, receipt)
}
func (_Router *RouterSession) DelegatedSynthTransferRequestOverGsn(tokenReal [32]byte, tokenSynth common.Address, amount *big.Int, from common.Address, to common.Address, synthParams ISynthesisSynthParams, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.DelegatedSynthTransferRequestOverGsn(&_Router.TransactOpts, tokenReal, tokenSynth, amount, from, to, synthParams, receipt)
}

// DelegatedSynthTransferRequest is a paid mutator transaction binding the contract method 0x3d5ac6cb.
//
// Solidity: function delegatedSynthTransferRequest(bytes32 tokenReal, address tokenSynth, uint256 amount, address from, address to, (address,address,uint256) synthParams, (uint256,uint256,uint8[2],bytes32[2],bytes32[2]) receipt) returns()
func (_Router *RouterTransactorSession) DelegatedSynthTransferRequest(tokenReal [32]byte, tokenSynth common.Address, amount *big.Int, from common.Address, to common.Address, synthParams ISynthesisSynthParams, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.DelegatedSynthTransferRequest(&_Router.TransactOpts, tokenReal, tokenSynth, amount, from, to, synthParams, receipt)
}
func (_Router *RouterTransactorSession) DelegatedSynthTransferRequestOverGsn(tokenReal [32]byte, tokenSynth common.Address, amount *big.Int, from common.Address, to common.Address, synthParams ISynthesisSynthParams, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.DelegatedSynthTransferRequestOverGsn(&_Router.TransactOpts, tokenReal, tokenSynth, amount, from, to, synthParams, receipt)
}

// DelegatedTokenSynthesizeRequest is a paid mutator transaction binding the contract method 0xc7026a47.
//
// Solidity: function delegatedTokenSynthesizeRequest(address token, uint256 amount, address from, (address,address,address,uint256) synthParams, (uint256,uint256,uint8[2],bytes32[2],bytes32[2]) receipt) returns()
func (_Router *RouterTransactor) DelegatedTokenSynthesizeRequest(opts *bind.TransactOpts, token common.Address, amount *big.Int, from common.Address, synthParams IPortalSynthParams, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "delegatedTokenSynthesizeRequest", token, amount, from, synthParams, receipt)
}
func (_Router *RouterTransactor) DelegatedTokenSynthesizeRequestOverGsn(opts *bind.TransactOpts, token common.Address, amount *big.Int, from common.Address, synthParams IPortalSynthParams, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "delegatedTokenSynthesizeRequest", token, amount, from, synthParams, receipt)
}

// DelegatedTokenSynthesizeRequest is a paid mutator transaction binding the contract method 0xc7026a47.
//
// Solidity: function delegatedTokenSynthesizeRequest(address token, uint256 amount, address from, (address,address,address,uint256) synthParams, (uint256,uint256,uint8[2],bytes32[2],bytes32[2]) receipt) returns()
func (_Router *RouterSession) DelegatedTokenSynthesizeRequest(token common.Address, amount *big.Int, from common.Address, synthParams IPortalSynthParams, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.DelegatedTokenSynthesizeRequest(&_Router.TransactOpts, token, amount, from, synthParams, receipt)
}
func (_Router *RouterSession) DelegatedTokenSynthesizeRequestOverGsn(token common.Address, amount *big.Int, from common.Address, synthParams IPortalSynthParams, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.DelegatedTokenSynthesizeRequestOverGsn(&_Router.TransactOpts, token, amount, from, synthParams, receipt)
}

// DelegatedTokenSynthesizeRequest is a paid mutator transaction binding the contract method 0xc7026a47.
//
// Solidity: function delegatedTokenSynthesizeRequest(address token, uint256 amount, address from, (address,address,address,uint256) synthParams, (uint256,uint256,uint8[2],bytes32[2],bytes32[2]) receipt) returns()
func (_Router *RouterTransactorSession) DelegatedTokenSynthesizeRequest(token common.Address, amount *big.Int, from common.Address, synthParams IPortalSynthParams, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.DelegatedTokenSynthesizeRequest(&_Router.TransactOpts, token, amount, from, synthParams, receipt)
}
func (_Router *RouterTransactorSession) DelegatedTokenSynthesizeRequestOverGsn(token common.Address, amount *big.Int, from common.Address, synthParams IPortalSynthParams, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.DelegatedTokenSynthesizeRequestOverGsn(&_Router.TransactOpts, token, amount, from, synthParams, receipt)
}

// DelegatedTokenSynthesizeRequestWithPermit is a paid mutator transaction binding the contract method 0xc067de17.
//
// Solidity: function delegatedTokenSynthesizeRequestWithPermit(address token, uint256 amount, address from, (address,address,address,uint256) synthParams, (uint8,bytes32,bytes32,uint256,bool) permitData, (uint256,uint256,uint8[2],bytes32[2],bytes32[2]) receipt) returns()
func (_Router *RouterTransactor) DelegatedTokenSynthesizeRequestWithPermit(opts *bind.TransactOpts, token common.Address, amount *big.Int, from common.Address, synthParams IPortalSynthParams, permitData IPortalPermitData, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "delegatedTokenSynthesizeRequestWithPermit", token, amount, from, synthParams, permitData, receipt)
}
func (_Router *RouterTransactor) DelegatedTokenSynthesizeRequestWithPermitOverGsn(opts *bind.TransactOpts, token common.Address, amount *big.Int, from common.Address, synthParams IPortalSynthParams, permitData IPortalPermitData, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "delegatedTokenSynthesizeRequestWithPermit", token, amount, from, synthParams, permitData, receipt)
}

// DelegatedTokenSynthesizeRequestWithPermit is a paid mutator transaction binding the contract method 0xc067de17.
//
// Solidity: function delegatedTokenSynthesizeRequestWithPermit(address token, uint256 amount, address from, (address,address,address,uint256) synthParams, (uint8,bytes32,bytes32,uint256,bool) permitData, (uint256,uint256,uint8[2],bytes32[2],bytes32[2]) receipt) returns()
func (_Router *RouterSession) DelegatedTokenSynthesizeRequestWithPermit(token common.Address, amount *big.Int, from common.Address, synthParams IPortalSynthParams, permitData IPortalPermitData, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.DelegatedTokenSynthesizeRequestWithPermit(&_Router.TransactOpts, token, amount, from, synthParams, permitData, receipt)
}
func (_Router *RouterSession) DelegatedTokenSynthesizeRequestWithPermitOverGsn(token common.Address, amount *big.Int, from common.Address, synthParams IPortalSynthParams, permitData IPortalPermitData, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.DelegatedTokenSynthesizeRequestWithPermitOverGsn(&_Router.TransactOpts, token, amount, from, synthParams, permitData, receipt)
}

// DelegatedTokenSynthesizeRequestWithPermit is a paid mutator transaction binding the contract method 0xc067de17.
//
// Solidity: function delegatedTokenSynthesizeRequestWithPermit(address token, uint256 amount, address from, (address,address,address,uint256) synthParams, (uint8,bytes32,bytes32,uint256,bool) permitData, (uint256,uint256,uint8[2],bytes32[2],bytes32[2]) receipt) returns()
func (_Router *RouterTransactorSession) DelegatedTokenSynthesizeRequestWithPermit(token common.Address, amount *big.Int, from common.Address, synthParams IPortalSynthParams, permitData IPortalPermitData, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.DelegatedTokenSynthesizeRequestWithPermit(&_Router.TransactOpts, token, amount, from, synthParams, permitData, receipt)
}
func (_Router *RouterTransactorSession) DelegatedTokenSynthesizeRequestWithPermitOverGsn(token common.Address, amount *big.Int, from common.Address, synthParams IPortalSynthParams, permitData IPortalPermitData, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.DelegatedTokenSynthesizeRequestWithPermitOverGsn(&_Router.TransactOpts, token, amount, from, synthParams, permitData, receipt)
}

// DelegatedUnsynthesizeRequest is a paid mutator transaction binding the contract method 0x905e4be0.
//
// Solidity: function delegatedUnsynthesizeRequest(address stoken, uint256 amount, address from, address to, address receiveSide, address oppositeBridge, uint256 chainId, (uint256,uint256,uint8[2],bytes32[2],bytes32[2]) receipt) returns()
func (_Router *RouterTransactor) DelegatedUnsynthesizeRequest(opts *bind.TransactOpts, stoken common.Address, amount *big.Int, from common.Address, to common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "delegatedUnsynthesizeRequest", stoken, amount, from, to, receiveSide, oppositeBridge, chainId, receipt)
}
func (_Router *RouterTransactor) DelegatedUnsynthesizeRequestOverGsn(opts *bind.TransactOpts, stoken common.Address, amount *big.Int, from common.Address, to common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "delegatedUnsynthesizeRequest", stoken, amount, from, to, receiveSide, oppositeBridge, chainId, receipt)
}

// DelegatedUnsynthesizeRequest is a paid mutator transaction binding the contract method 0x905e4be0.
//
// Solidity: function delegatedUnsynthesizeRequest(address stoken, uint256 amount, address from, address to, address receiveSide, address oppositeBridge, uint256 chainId, (uint256,uint256,uint8[2],bytes32[2],bytes32[2]) receipt) returns()
func (_Router *RouterSession) DelegatedUnsynthesizeRequest(stoken common.Address, amount *big.Int, from common.Address, to common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.DelegatedUnsynthesizeRequest(&_Router.TransactOpts, stoken, amount, from, to, receiveSide, oppositeBridge, chainId, receipt)
}
func (_Router *RouterSession) DelegatedUnsynthesizeRequestOverGsn(stoken common.Address, amount *big.Int, from common.Address, to common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.DelegatedUnsynthesizeRequestOverGsn(&_Router.TransactOpts, stoken, amount, from, to, receiveSide, oppositeBridge, chainId, receipt)
}

// DelegatedUnsynthesizeRequest is a paid mutator transaction binding the contract method 0x905e4be0.
//
// Solidity: function delegatedUnsynthesizeRequest(address stoken, uint256 amount, address from, address to, address receiveSide, address oppositeBridge, uint256 chainId, (uint256,uint256,uint8[2],bytes32[2],bytes32[2]) receipt) returns()
func (_Router *RouterTransactorSession) DelegatedUnsynthesizeRequest(stoken common.Address, amount *big.Int, from common.Address, to common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.DelegatedUnsynthesizeRequest(&_Router.TransactOpts, stoken, amount, from, to, receiveSide, oppositeBridge, chainId, receipt)
}
func (_Router *RouterTransactorSession) DelegatedUnsynthesizeRequestOverGsn(stoken common.Address, amount *big.Int, from common.Address, to common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.DelegatedUnsynthesizeRequestOverGsn(&_Router.TransactOpts, stoken, amount, from, to, receiveSide, oppositeBridge, chainId, receipt)
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

// MetaExchangeRequestVia3pool is a paid mutator transaction binding the contract method 0x79da42b9.
//
// Solidity: function metaExchangeRequestVia3pool((address,address,address,uint256,int128,int128,uint256,int128,uint256,address,address,address,address,uint256,address,uint256) params, (uint8,bytes32,bytes32,uint256,bool)[] permit, address from, address[3] token, uint256[3] amount) returns()
func (_Router *RouterTransactor) MetaExchangeRequestVia3pool(opts *bind.TransactOpts, params ICurveProxyMetaExchangeParams, permit []ICurveProxyPermitData, from common.Address, token [3]common.Address, amount [3]*big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "metaExchangeRequestVia3pool", params, permit, from, token, amount)
}
func (_Router *RouterTransactor) MetaExchangeRequestVia3poolOverGsn(opts *bind.TransactOpts, params ICurveProxyMetaExchangeParams, permit []ICurveProxyPermitData, from common.Address, token [3]common.Address, amount [3]*big.Int) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "metaExchangeRequestVia3pool", params, permit, from, token, amount)
}

// MetaExchangeRequestVia3pool is a paid mutator transaction binding the contract method 0x79da42b9.
//
// Solidity: function metaExchangeRequestVia3pool((address,address,address,uint256,int128,int128,uint256,int128,uint256,address,address,address,address,uint256,address,uint256) params, (uint8,bytes32,bytes32,uint256,bool)[] permit, address from, address[3] token, uint256[3] amount) returns()
func (_Router *RouterSession) MetaExchangeRequestVia3pool(params ICurveProxyMetaExchangeParams, permit []ICurveProxyPermitData, from common.Address, token [3]common.Address, amount [3]*big.Int) (*types.Transaction, error) {
	return _Router.Contract.MetaExchangeRequestVia3pool(&_Router.TransactOpts, params, permit, from, token, amount)
}
func (_Router *RouterSession) MetaExchangeRequestVia3poolOverGsn(params ICurveProxyMetaExchangeParams, permit []ICurveProxyPermitData, from common.Address, token [3]common.Address, amount [3]*big.Int) (common.Hash, error) {
	return _Router.Contract.MetaExchangeRequestVia3poolOverGsn(&_Router.TransactOpts, params, permit, from, token, amount)
}

// MetaExchangeRequestVia3pool is a paid mutator transaction binding the contract method 0x79da42b9.
//
// Solidity: function metaExchangeRequestVia3pool((address,address,address,uint256,int128,int128,uint256,int128,uint256,address,address,address,address,uint256,address,uint256) params, (uint8,bytes32,bytes32,uint256,bool)[] permit, address from, address[3] token, uint256[3] amount) returns()
func (_Router *RouterTransactorSession) MetaExchangeRequestVia3pool(params ICurveProxyMetaExchangeParams, permit []ICurveProxyPermitData, from common.Address, token [3]common.Address, amount [3]*big.Int) (*types.Transaction, error) {
	return _Router.Contract.MetaExchangeRequestVia3pool(&_Router.TransactOpts, params, permit, from, token, amount)
}
func (_Router *RouterTransactorSession) MetaExchangeRequestVia3poolOverGsn(params ICurveProxyMetaExchangeParams, permit []ICurveProxyPermitData, from common.Address, token [3]common.Address, amount [3]*big.Int) (common.Hash, error) {
	return _Router.Contract.MetaExchangeRequestVia3poolOverGsn(&_Router.TransactOpts, params, permit, from, token, amount)
}

// MintEusdRequestVia3pool is a paid mutator transaction binding the contract method 0x605ae544.
//
// Solidity: function mintEusdRequestVia3pool((address,uint256,uint256,address,uint256,address,address,uint256) params, (uint8,bytes32,bytes32,uint256,bool)[] permit, address from, address[3] token, uint256[3] amount) returns()
func (_Router *RouterTransactor) MintEusdRequestVia3pool(opts *bind.TransactOpts, params ICurveProxyMetaMintEUSD, permit []ICurveProxyPermitData, from common.Address, token [3]common.Address, amount [3]*big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "mintEusdRequestVia3pool", params, permit, from, token, amount)
}
func (_Router *RouterTransactor) MintEusdRequestVia3poolOverGsn(opts *bind.TransactOpts, params ICurveProxyMetaMintEUSD, permit []ICurveProxyPermitData, from common.Address, token [3]common.Address, amount [3]*big.Int) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "mintEusdRequestVia3pool", params, permit, from, token, amount)
}

// MintEusdRequestVia3pool is a paid mutator transaction binding the contract method 0x605ae544.
//
// Solidity: function mintEusdRequestVia3pool((address,uint256,uint256,address,uint256,address,address,uint256) params, (uint8,bytes32,bytes32,uint256,bool)[] permit, address from, address[3] token, uint256[3] amount) returns()
func (_Router *RouterSession) MintEusdRequestVia3pool(params ICurveProxyMetaMintEUSD, permit []ICurveProxyPermitData, from common.Address, token [3]common.Address, amount [3]*big.Int) (*types.Transaction, error) {
	return _Router.Contract.MintEusdRequestVia3pool(&_Router.TransactOpts, params, permit, from, token, amount)
}
func (_Router *RouterSession) MintEusdRequestVia3poolOverGsn(params ICurveProxyMetaMintEUSD, permit []ICurveProxyPermitData, from common.Address, token [3]common.Address, amount [3]*big.Int) (common.Hash, error) {
	return _Router.Contract.MintEusdRequestVia3poolOverGsn(&_Router.TransactOpts, params, permit, from, token, amount)
}

// MintEusdRequestVia3pool is a paid mutator transaction binding the contract method 0x605ae544.
//
// Solidity: function mintEusdRequestVia3pool((address,uint256,uint256,address,uint256,address,address,uint256) params, (uint8,bytes32,bytes32,uint256,bool)[] permit, address from, address[3] token, uint256[3] amount) returns()
func (_Router *RouterTransactorSession) MintEusdRequestVia3pool(params ICurveProxyMetaMintEUSD, permit []ICurveProxyPermitData, from common.Address, token [3]common.Address, amount [3]*big.Int) (*types.Transaction, error) {
	return _Router.Contract.MintEusdRequestVia3pool(&_Router.TransactOpts, params, permit, from, token, amount)
}
func (_Router *RouterTransactorSession) MintEusdRequestVia3poolOverGsn(params ICurveProxyMetaMintEUSD, permit []ICurveProxyPermitData, from common.Address, token [3]common.Address, amount [3]*big.Int) (common.Hash, error) {
	return _Router.Contract.MintEusdRequestVia3poolOverGsn(&_Router.TransactOpts, params, permit, from, token, amount)
}

// RedeemEusdRequest is a paid mutator transaction binding the contract method 0x0b414a0d.
//
// Solidity: function redeemEusdRequest((address,int128,uint256,address,uint256,int128,uint256,address) params, (uint8,bytes32,bytes32,uint256,bool) permit, address payToken, address from, address receiveSide, address oppositeBridge, uint256 chainId) returns()
func (_Router *RouterTransactor) RedeemEusdRequest(opts *bind.TransactOpts, params ICurveProxyMetaRedeemEUSD, permit ICurveProxyPermitData, payToken common.Address, from common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "redeemEusdRequest", params, permit, payToken, from, receiveSide, oppositeBridge, chainId)
}
func (_Router *RouterTransactor) RedeemEusdRequestOverGsn(opts *bind.TransactOpts, params ICurveProxyMetaRedeemEUSD, permit ICurveProxyPermitData, payToken common.Address, from common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "redeemEusdRequest", params, permit, payToken, from, receiveSide, oppositeBridge, chainId)
}

// RedeemEusdRequest is a paid mutator transaction binding the contract method 0x0b414a0d.
//
// Solidity: function redeemEusdRequest((address,int128,uint256,address,uint256,int128,uint256,address) params, (uint8,bytes32,bytes32,uint256,bool) permit, address payToken, address from, address receiveSide, address oppositeBridge, uint256 chainId) returns()
func (_Router *RouterSession) RedeemEusdRequest(params ICurveProxyMetaRedeemEUSD, permit ICurveProxyPermitData, payToken common.Address, from common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _Router.Contract.RedeemEusdRequest(&_Router.TransactOpts, params, permit, payToken, from, receiveSide, oppositeBridge, chainId)
}
func (_Router *RouterSession) RedeemEusdRequestOverGsn(params ICurveProxyMetaRedeemEUSD, permit ICurveProxyPermitData, payToken common.Address, from common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (common.Hash, error) {
	return _Router.Contract.RedeemEusdRequestOverGsn(&_Router.TransactOpts, params, permit, payToken, from, receiveSide, oppositeBridge, chainId)
}

// RedeemEusdRequest is a paid mutator transaction binding the contract method 0x0b414a0d.
//
// Solidity: function redeemEusdRequest((address,int128,uint256,address,uint256,int128,uint256,address) params, (uint8,bytes32,bytes32,uint256,bool) permit, address payToken, address from, address receiveSide, address oppositeBridge, uint256 chainId) returns()
func (_Router *RouterTransactorSession) RedeemEusdRequest(params ICurveProxyMetaRedeemEUSD, permit ICurveProxyPermitData, payToken common.Address, from common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _Router.Contract.RedeemEusdRequest(&_Router.TransactOpts, params, permit, payToken, from, receiveSide, oppositeBridge, chainId)
}
func (_Router *RouterTransactorSession) RedeemEusdRequestOverGsn(params ICurveProxyMetaRedeemEUSD, permit ICurveProxyPermitData, payToken common.Address, from common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (common.Hash, error) {
	return _Router.Contract.RedeemEusdRequestOverGsn(&_Router.TransactOpts, params, permit, payToken, from, receiveSide, oppositeBridge, chainId)
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

// SynthTransferRequest is a paid mutator transaction binding the contract method 0x61bf99ac.
//
// Solidity: function synthTransferRequest(bytes32 tokenReal, uint256 amount, address from, address to, (address,address,uint256) synthParams) returns()
func (_Router *RouterTransactor) SynthTransferRequest(opts *bind.TransactOpts, tokenReal [32]byte, amount *big.Int, from common.Address, to common.Address, synthParams ISynthesisSynthParams) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "synthTransferRequest", tokenReal, amount, from, to, synthParams)
}
func (_Router *RouterTransactor) SynthTransferRequestOverGsn(opts *bind.TransactOpts, tokenReal [32]byte, amount *big.Int, from common.Address, to common.Address, synthParams ISynthesisSynthParams) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "synthTransferRequest", tokenReal, amount, from, to, synthParams)
}

// SynthTransferRequest is a paid mutator transaction binding the contract method 0x61bf99ac.
//
// Solidity: function synthTransferRequest(bytes32 tokenReal, uint256 amount, address from, address to, (address,address,uint256) synthParams) returns()
func (_Router *RouterSession) SynthTransferRequest(tokenReal [32]byte, amount *big.Int, from common.Address, to common.Address, synthParams ISynthesisSynthParams) (*types.Transaction, error) {
	return _Router.Contract.SynthTransferRequest(&_Router.TransactOpts, tokenReal, amount, from, to, synthParams)
}
func (_Router *RouterSession) SynthTransferRequestOverGsn(tokenReal [32]byte, amount *big.Int, from common.Address, to common.Address, synthParams ISynthesisSynthParams) (common.Hash, error) {
	return _Router.Contract.SynthTransferRequestOverGsn(&_Router.TransactOpts, tokenReal, amount, from, to, synthParams)
}

// SynthTransferRequest is a paid mutator transaction binding the contract method 0x61bf99ac.
//
// Solidity: function synthTransferRequest(bytes32 tokenReal, uint256 amount, address from, address to, (address,address,uint256) synthParams) returns()
func (_Router *RouterTransactorSession) SynthTransferRequest(tokenReal [32]byte, amount *big.Int, from common.Address, to common.Address, synthParams ISynthesisSynthParams) (*types.Transaction, error) {
	return _Router.Contract.SynthTransferRequest(&_Router.TransactOpts, tokenReal, amount, from, to, synthParams)
}
func (_Router *RouterTransactorSession) SynthTransferRequestOverGsn(tokenReal [32]byte, amount *big.Int, from common.Address, to common.Address, synthParams ISynthesisSynthParams) (common.Hash, error) {
	return _Router.Contract.SynthTransferRequestOverGsn(&_Router.TransactOpts, tokenReal, amount, from, to, synthParams)
}

// TokenSynthesizeRequest is a paid mutator transaction binding the contract method 0x4b7568fb.
//
// Solidity: function tokenSynthesizeRequest(address token, uint256 amount, address from, (address,address,address,uint256) synthParams) returns()
func (_Router *RouterTransactor) TokenSynthesizeRequest(opts *bind.TransactOpts, token common.Address, amount *big.Int, from common.Address, synthParams IPortalSynthParams) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "tokenSynthesizeRequest", token, amount, from, synthParams)
}
func (_Router *RouterTransactor) TokenSynthesizeRequestOverGsn(opts *bind.TransactOpts, token common.Address, amount *big.Int, from common.Address, synthParams IPortalSynthParams) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "tokenSynthesizeRequest", token, amount, from, synthParams)
}

// TokenSynthesizeRequest is a paid mutator transaction binding the contract method 0x4b7568fb.
//
// Solidity: function tokenSynthesizeRequest(address token, uint256 amount, address from, (address,address,address,uint256) synthParams) returns()
func (_Router *RouterSession) TokenSynthesizeRequest(token common.Address, amount *big.Int, from common.Address, synthParams IPortalSynthParams) (*types.Transaction, error) {
	return _Router.Contract.TokenSynthesizeRequest(&_Router.TransactOpts, token, amount, from, synthParams)
}
func (_Router *RouterSession) TokenSynthesizeRequestOverGsn(token common.Address, amount *big.Int, from common.Address, synthParams IPortalSynthParams) (common.Hash, error) {
	return _Router.Contract.TokenSynthesizeRequestOverGsn(&_Router.TransactOpts, token, amount, from, synthParams)
}

// TokenSynthesizeRequest is a paid mutator transaction binding the contract method 0x4b7568fb.
//
// Solidity: function tokenSynthesizeRequest(address token, uint256 amount, address from, (address,address,address,uint256) synthParams) returns()
func (_Router *RouterTransactorSession) TokenSynthesizeRequest(token common.Address, amount *big.Int, from common.Address, synthParams IPortalSynthParams) (*types.Transaction, error) {
	return _Router.Contract.TokenSynthesizeRequest(&_Router.TransactOpts, token, amount, from, synthParams)
}
func (_Router *RouterTransactorSession) TokenSynthesizeRequestOverGsn(token common.Address, amount *big.Int, from common.Address, synthParams IPortalSynthParams) (common.Hash, error) {
	return _Router.Contract.TokenSynthesizeRequestOverGsn(&_Router.TransactOpts, token, amount, from, synthParams)
}

// TokenSynthesizeRequestToSolana is a paid mutator transaction binding the contract method 0xdb1c62e5.
//
// Solidity: function tokenSynthesizeRequestToSolana(address token, uint256 amount, address from, bytes32[] pubkeys, bytes1 txStateBump, uint256 chainId) returns()
func (_Router *RouterTransactor) TokenSynthesizeRequestToSolana(opts *bind.TransactOpts, token common.Address, amount *big.Int, from common.Address, pubkeys [][32]byte, txStateBump [1]byte, chainId *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "tokenSynthesizeRequestToSolana", token, amount, from, pubkeys, txStateBump, chainId)
}
func (_Router *RouterTransactor) TokenSynthesizeRequestToSolanaOverGsn(opts *bind.TransactOpts, token common.Address, amount *big.Int, from common.Address, pubkeys [][32]byte, txStateBump [1]byte, chainId *big.Int) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "tokenSynthesizeRequestToSolana", token, amount, from, pubkeys, txStateBump, chainId)
}

// TokenSynthesizeRequestToSolana is a paid mutator transaction binding the contract method 0xdb1c62e5.
//
// Solidity: function tokenSynthesizeRequestToSolana(address token, uint256 amount, address from, bytes32[] pubkeys, bytes1 txStateBump, uint256 chainId) returns()
func (_Router *RouterSession) TokenSynthesizeRequestToSolana(token common.Address, amount *big.Int, from common.Address, pubkeys [][32]byte, txStateBump [1]byte, chainId *big.Int) (*types.Transaction, error) {
	return _Router.Contract.TokenSynthesizeRequestToSolana(&_Router.TransactOpts, token, amount, from, pubkeys, txStateBump, chainId)
}
func (_Router *RouterSession) TokenSynthesizeRequestToSolanaOverGsn(token common.Address, amount *big.Int, from common.Address, pubkeys [][32]byte, txStateBump [1]byte, chainId *big.Int) (common.Hash, error) {
	return _Router.Contract.TokenSynthesizeRequestToSolanaOverGsn(&_Router.TransactOpts, token, amount, from, pubkeys, txStateBump, chainId)
}

// TokenSynthesizeRequestToSolana is a paid mutator transaction binding the contract method 0xdb1c62e5.
//
// Solidity: function tokenSynthesizeRequestToSolana(address token, uint256 amount, address from, bytes32[] pubkeys, bytes1 txStateBump, uint256 chainId) returns()
func (_Router *RouterTransactorSession) TokenSynthesizeRequestToSolana(token common.Address, amount *big.Int, from common.Address, pubkeys [][32]byte, txStateBump [1]byte, chainId *big.Int) (*types.Transaction, error) {
	return _Router.Contract.TokenSynthesizeRequestToSolana(&_Router.TransactOpts, token, amount, from, pubkeys, txStateBump, chainId)
}
func (_Router *RouterTransactorSession) TokenSynthesizeRequestToSolanaOverGsn(token common.Address, amount *big.Int, from common.Address, pubkeys [][32]byte, txStateBump [1]byte, chainId *big.Int) (common.Hash, error) {
	return _Router.Contract.TokenSynthesizeRequestToSolanaOverGsn(&_Router.TransactOpts, token, amount, from, pubkeys, txStateBump, chainId)
}

// TokenSynthesizeRequestWithPermit is a paid mutator transaction binding the contract method 0x259ab8a4.
//
// Solidity: function tokenSynthesizeRequestWithPermit(address token, uint256 amount, address from, (address,address,address,uint256) synthParams, (uint8,bytes32,bytes32,uint256,bool) permitData) returns()
func (_Router *RouterTransactor) TokenSynthesizeRequestWithPermit(opts *bind.TransactOpts, token common.Address, amount *big.Int, from common.Address, synthParams IPortalSynthParams, permitData IPortalPermitData) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "tokenSynthesizeRequestWithPermit", token, amount, from, synthParams, permitData)
}
func (_Router *RouterTransactor) TokenSynthesizeRequestWithPermitOverGsn(opts *bind.TransactOpts, token common.Address, amount *big.Int, from common.Address, synthParams IPortalSynthParams, permitData IPortalPermitData) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "tokenSynthesizeRequestWithPermit", token, amount, from, synthParams, permitData)
}

// TokenSynthesizeRequestWithPermit is a paid mutator transaction binding the contract method 0x259ab8a4.
//
// Solidity: function tokenSynthesizeRequestWithPermit(address token, uint256 amount, address from, (address,address,address,uint256) synthParams, (uint8,bytes32,bytes32,uint256,bool) permitData) returns()
func (_Router *RouterSession) TokenSynthesizeRequestWithPermit(token common.Address, amount *big.Int, from common.Address, synthParams IPortalSynthParams, permitData IPortalPermitData) (*types.Transaction, error) {
	return _Router.Contract.TokenSynthesizeRequestWithPermit(&_Router.TransactOpts, token, amount, from, synthParams, permitData)
}
func (_Router *RouterSession) TokenSynthesizeRequestWithPermitOverGsn(token common.Address, amount *big.Int, from common.Address, synthParams IPortalSynthParams, permitData IPortalPermitData) (common.Hash, error) {
	return _Router.Contract.TokenSynthesizeRequestWithPermitOverGsn(&_Router.TransactOpts, token, amount, from, synthParams, permitData)
}

// TokenSynthesizeRequestWithPermit is a paid mutator transaction binding the contract method 0x259ab8a4.
//
// Solidity: function tokenSynthesizeRequestWithPermit(address token, uint256 amount, address from, (address,address,address,uint256) synthParams, (uint8,bytes32,bytes32,uint256,bool) permitData) returns()
func (_Router *RouterTransactorSession) TokenSynthesizeRequestWithPermit(token common.Address, amount *big.Int, from common.Address, synthParams IPortalSynthParams, permitData IPortalPermitData) (*types.Transaction, error) {
	return _Router.Contract.TokenSynthesizeRequestWithPermit(&_Router.TransactOpts, token, amount, from, synthParams, permitData)
}
func (_Router *RouterTransactorSession) TokenSynthesizeRequestWithPermitOverGsn(token common.Address, amount *big.Int, from common.Address, synthParams IPortalSynthParams, permitData IPortalPermitData) (common.Hash, error) {
	return _Router.Contract.TokenSynthesizeRequestWithPermitOverGsn(&_Router.TransactOpts, token, amount, from, synthParams, permitData)
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

// UnsynthesizeRequest is a paid mutator transaction binding the contract method 0xd7c7d416.
//
// Solidity: function unsynthesizeRequest(address stoken, uint256 amount, address from, address to, address receiveSide, address oppositeBridge, uint256 chainId) returns()
func (_Router *RouterTransactor) UnsynthesizeRequest(opts *bind.TransactOpts, stoken common.Address, amount *big.Int, from common.Address, to common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "unsynthesizeRequest", stoken, amount, from, to, receiveSide, oppositeBridge, chainId)
}
func (_Router *RouterTransactor) UnsynthesizeRequestOverGsn(opts *bind.TransactOpts, stoken common.Address, amount *big.Int, from common.Address, to common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "unsynthesizeRequest", stoken, amount, from, to, receiveSide, oppositeBridge, chainId)
}

// UnsynthesizeRequest is a paid mutator transaction binding the contract method 0xd7c7d416.
//
// Solidity: function unsynthesizeRequest(address stoken, uint256 amount, address from, address to, address receiveSide, address oppositeBridge, uint256 chainId) returns()
func (_Router *RouterSession) UnsynthesizeRequest(stoken common.Address, amount *big.Int, from common.Address, to common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _Router.Contract.UnsynthesizeRequest(&_Router.TransactOpts, stoken, amount, from, to, receiveSide, oppositeBridge, chainId)
}
func (_Router *RouterSession) UnsynthesizeRequestOverGsn(stoken common.Address, amount *big.Int, from common.Address, to common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (common.Hash, error) {
	return _Router.Contract.UnsynthesizeRequestOverGsn(&_Router.TransactOpts, stoken, amount, from, to, receiveSide, oppositeBridge, chainId)
}

// UnsynthesizeRequest is a paid mutator transaction binding the contract method 0xd7c7d416.
//
// Solidity: function unsynthesizeRequest(address stoken, uint256 amount, address from, address to, address receiveSide, address oppositeBridge, uint256 chainId) returns()
func (_Router *RouterTransactorSession) UnsynthesizeRequest(stoken common.Address, amount *big.Int, from common.Address, to common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _Router.Contract.UnsynthesizeRequest(&_Router.TransactOpts, stoken, amount, from, to, receiveSide, oppositeBridge, chainId)
}
func (_Router *RouterTransactorSession) UnsynthesizeRequestOverGsn(stoken common.Address, amount *big.Int, from common.Address, to common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int) (common.Hash, error) {
	return _Router.Contract.UnsynthesizeRequestOverGsn(&_Router.TransactOpts, stoken, amount, from, to, receiveSide, oppositeBridge, chainId)
}

// UnsynthesizeRequestToSolana is a paid mutator transaction binding the contract method 0x3cce3045.
//
// Solidity: function unsynthesizeRequestToSolana(address stoken, address from, bytes32[] pubkeys, uint256 amount, uint256 chainId) returns()
func (_Router *RouterTransactor) UnsynthesizeRequestToSolana(opts *bind.TransactOpts, stoken common.Address, from common.Address, pubkeys [][32]byte, amount *big.Int, chainId *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "unsynthesizeRequestToSolana", stoken, from, pubkeys, amount, chainId)
}
func (_Router *RouterTransactor) UnsynthesizeRequestToSolanaOverGsn(opts *bind.TransactOpts, stoken common.Address, from common.Address, pubkeys [][32]byte, amount *big.Int, chainId *big.Int) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "unsynthesizeRequestToSolana", stoken, from, pubkeys, amount, chainId)
}

// UnsynthesizeRequestToSolana is a paid mutator transaction binding the contract method 0x3cce3045.
//
// Solidity: function unsynthesizeRequestToSolana(address stoken, address from, bytes32[] pubkeys, uint256 amount, uint256 chainId) returns()
func (_Router *RouterSession) UnsynthesizeRequestToSolana(stoken common.Address, from common.Address, pubkeys [][32]byte, amount *big.Int, chainId *big.Int) (*types.Transaction, error) {
	return _Router.Contract.UnsynthesizeRequestToSolana(&_Router.TransactOpts, stoken, from, pubkeys, amount, chainId)
}
func (_Router *RouterSession) UnsynthesizeRequestToSolanaOverGsn(stoken common.Address, from common.Address, pubkeys [][32]byte, amount *big.Int, chainId *big.Int) (common.Hash, error) {
	return _Router.Contract.UnsynthesizeRequestToSolanaOverGsn(&_Router.TransactOpts, stoken, from, pubkeys, amount, chainId)
}

// UnsynthesizeRequestToSolana is a paid mutator transaction binding the contract method 0x3cce3045.
//
// Solidity: function unsynthesizeRequestToSolana(address stoken, address from, bytes32[] pubkeys, uint256 amount, uint256 chainId) returns()
func (_Router *RouterTransactorSession) UnsynthesizeRequestToSolana(stoken common.Address, from common.Address, pubkeys [][32]byte, amount *big.Int, chainId *big.Int) (*types.Transaction, error) {
	return _Router.Contract.UnsynthesizeRequestToSolana(&_Router.TransactOpts, stoken, from, pubkeys, amount, chainId)
}
func (_Router *RouterTransactorSession) UnsynthesizeRequestToSolanaOverGsn(stoken common.Address, from common.Address, pubkeys [][32]byte, amount *big.Int, chainId *big.Int) (common.Hash, error) {
	return _Router.Contract.UnsynthesizeRequestToSolanaOverGsn(&_Router.TransactOpts, stoken, from, pubkeys, amount, chainId)
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
	PayToken       common.Address
	ExecutionPrice *big.Int
	Worker         common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCrosschainPaymentEvent is a free log retrieval operation binding the contract event 0x8adfc333732362694bf9044b570e1ae5d25a5b143b8aa3c267aee92adae48bbe.
//
// Solidity: event CrosschainPaymentEvent(address indexed userFrom, address payToken, uint256 executionPrice, address indexed worker)
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

// WatchCrosschainPaymentEvent is a free log subscription operation binding the contract event 0x8adfc333732362694bf9044b570e1ae5d25a5b143b8aa3c267aee92adae48bbe.
//
// Solidity: event CrosschainPaymentEvent(address indexed userFrom, address payToken, uint256 executionPrice, address indexed worker)
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

// ParseCrosschainPaymentEvent is a log parse operation binding the contract event 0x8adfc333732362694bf9044b570e1ae5d25a5b143b8aa3c267aee92adae48bbe.
//
// Solidity: event CrosschainPaymentEvent(address indexed userFrom, address payToken, uint256 executionPrice, address indexed worker)
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
