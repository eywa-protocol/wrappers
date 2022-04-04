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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"portal\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"synthesis\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"curveProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"localTreasury\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userFrom\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionPrice\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"}],\"name\":\"CrosschainPaymentEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_trustedWorker\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"transitData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structIPortal.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structIPortal.PermitData[]\",\"name\":\"permitData\",\"type\":\"tuple[]\"}],\"name\":\"batchSynthesizeRequestWithDataTransit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"transitData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structIPortal.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structIPortal.PermitData[]\",\"name\":\"permitData\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"executionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8[2]\",\"name\":\"v\",\"type\":\"uint8[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"r\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"s\",\"type\":\"bytes32[2]\"}],\"internalType\":\"structRouter.DelegatedCallReceipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"delegatedBatchSynthesizeRequestWithDataTransit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remove\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinMintAmount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinDy\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"chain2address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"initialBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialChainID\",\"type\":\"uint256\"}],\"internalType\":\"structICurveProxy.MetaExchangeParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structICurveProxy.PermitData[]\",\"name\":\"permit\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address[3]\",\"name\":\"token\",\"type\":\"address[3]\"},{\"internalType\":\"uint256[3]\",\"name\":\"amount\",\"type\":\"uint256[3]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"executionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8[2]\",\"name\":\"v\",\"type\":\"uint8[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"r\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"s\",\"type\":\"bytes32[2]\"}],\"internalType\":\"structRouter.DelegatedCallReceipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"delegatedMetaExchangeRequestVia3pool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addAtCrosschainPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinMintAmountC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lpIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addAtHubPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinMintAmountH\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialChainID\",\"type\":\"uint256\"}],\"internalType\":\"structICurveProxy.MetaMintEUSD\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structICurveProxy.PermitData[]\",\"name\":\"permit\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address[3]\",\"name\":\"token\",\"type\":\"address[3]\"},{\"internalType\":\"uint256[3]\",\"name\":\"amount\",\"type\":\"uint256[3]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"executionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8[2]\",\"name\":\"v\",\"type\":\"uint8[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"r\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"s\",\"type\":\"bytes32[2]\"}],\"internalType\":\"structRouter.DelegatedCallReceipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"delegatedMintEusdRequestVia3pool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"removeAtCrosschainPool\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinAmountC\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"removeAtHubPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountH\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"y\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinAmountH\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"internalType\":\"structICurveProxy.MetaRedeemEUSD\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structICurveProxy.PermitData\",\"name\":\"permit\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"executionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8[2]\",\"name\":\"v\",\"type\":\"uint8[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"r\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"s\",\"type\":\"bytes32[2]\"}],\"internalType\":\"structRouter.DelegatedCallReceipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"delegatedRedeemEusdRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tokenReal\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenSynth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structISynthesis.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"executionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8[2]\",\"name\":\"v\",\"type\":\"uint8[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"r\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"s\",\"type\":\"bytes32[2]\"}],\"internalType\":\"structRouter.DelegatedCallReceipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"delegatedSynthTransferRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structIPortal.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"executionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8[2]\",\"name\":\"v\",\"type\":\"uint8[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"r\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"s\",\"type\":\"bytes32[2]\"}],\"internalType\":\"structRouter.DelegatedCallReceipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"delegatedTokenSynthesizeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"pubkeys\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes1\",\"name\":\"txStateBump\",\"type\":\"bytes1\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"executionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8[2]\",\"name\":\"v\",\"type\":\"uint8[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"r\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"s\",\"type\":\"bytes32[2]\"}],\"internalType\":\"structRouter.DelegatedCallReceipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"delegatedTokenSynthesizeRequestToSolana\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structIPortal.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structIPortal.PermitData\",\"name\":\"permitData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"executionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8[2]\",\"name\":\"v\",\"type\":\"uint8[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"r\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"s\",\"type\":\"bytes32[2]\"}],\"internalType\":\"structRouter.DelegatedCallReceipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"delegatedTokenSynthesizeRequestWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stoken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"executionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8[2]\",\"name\":\"v\",\"type\":\"uint8[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"r\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"s\",\"type\":\"bytes32[2]\"}],\"internalType\":\"structRouter.DelegatedCallReceipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"delegatedUnsynthesizeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stoken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"pubkeys\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"executionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8[2]\",\"name\":\"v\",\"type\":\"uint8[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"r\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"s\",\"type\":\"bytes32[2]\"}],\"internalType\":\"structRouter.DelegatedCallReceipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"delegatedUnsynthesizeRequestToSolana\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"emergencyUnburnRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"emergencyUnsyntesizeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remove\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinMintAmount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinDy\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"chain2address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"initialBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialChainID\",\"type\":\"uint256\"}],\"internalType\":\"structICurveProxy.MetaExchangeParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structICurveProxy.PermitData[]\",\"name\":\"permit\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address[3]\",\"name\":\"token\",\"type\":\"address[3]\"},{\"internalType\":\"uint256[3]\",\"name\":\"amount\",\"type\":\"uint256[3]\"}],\"name\":\"metaExchangeRequestVia3pool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addAtCrosschainPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinMintAmountC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lpIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addAtHubPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinMintAmountH\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialChainID\",\"type\":\"uint256\"}],\"internalType\":\"structICurveProxy.MetaMintEUSD\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structICurveProxy.PermitData[]\",\"name\":\"permit\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address[3]\",\"name\":\"token\",\"type\":\"address[3]\"},{\"internalType\":\"uint256[3]\",\"name\":\"amount\",\"type\":\"uint256[3]\"}],\"name\":\"mintEusdRequestVia3pool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"removeAtCrosschainPool\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinAmountC\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"removeAtHubPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountH\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"y\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinAmountH\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"internalType\":\"structICurveProxy.MetaRedeemEUSD\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structICurveProxy.PermitData\",\"name\":\"permit\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"redeemEusdRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"}],\"name\":\"removeTrustedWorker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"}],\"name\":\"setTrustedWorker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tokenReal\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structISynthesis.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"}],\"name\":\"synthTransferRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structIPortal.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"}],\"name\":\"tokenSynthesizeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"pubkeys\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes1\",\"name\":\"txStateBump\",\"type\":\"bytes1\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"tokenSynthesizeRequestToSolana\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structIPortal.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structIPortal.PermitData\",\"name\":\"permitData\",\"type\":\"tuple\"}],\"name\":\"tokenSynthesizeRequestWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stoken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"unsynthesizeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stoken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"pubkeys\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"unsynthesizeRequestToSolana\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162003ace38038062003ace8339810160408190526200003491620000e6565b600080546001600160a01b031916339081178255604051909182917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350600380546001600160a01b039586166001600160a01b031991821617909155600480549486169482169490941790935560028054928516928416929092179091556001805491909316911617905562000143565b80516001600160a01b0381168114620000e157600080fd5b919050565b60008060008060808587031215620000fd57600080fd5b6200010885620000c9565b93506200011860208601620000c9565b92506200012860408601620000c9565b91506200013860608601620000c9565b905092959194509250565b61397b80620001536000396000f3fe608060405234801561001057600080fd5b50600436106101c45760003560e01c80638da5cb5b116100f9578063c7026a4711610097578063e840f9d311610071578063e840f9d314610395578063f2fde38b146103c8578063f3a08186146103db578063f9bee062146103ee57600080fd5b8063c7026a471461035c578063d7c7d4161461036f578063db1c62e51461038257600080fd5b8063ad1dbe71116100d3578063ad1dbe7114610310578063b981b48514610323578063bd7d4b9614610336578063c067de171461034957600080fd5b80638da5cb5b146102ca578063905e4be0146102ea5780639500125b146102fd57600080fd5b8063605ae54411610166578063715018a611610140578063715018a61461028957806372b90c021461029157806379da42b9146102a45780637fe04792146102b757600080fd5b8063605ae5441461025057806361bf99ac1461026357806362e41ded1461027657600080fd5b80633d5ac6cb116101a25780633d5ac6cb1461020457806343961668146102175780634b7568fb1461022a578063604477a61461023d57600080fd5b80630b414a0d146101c9578063259ab8a4146101de5780633cce3045146101f1575b600080fd5b6101dc6101d7366004611f65565b610401565b005b6101dc6101ec366004612009565b610492565b6101dc6101ff3660046120b9565b6105ba565b6101dc610212366004612141565b6106b3565b6101dc6102253660046121d4565b610746565b6101dc610238366004612242565b610782565b6101dc61024b366004612331565b61080a565b6101dc61025e366004612445565b610922565b6101dc6102713660046124d2565b6109ee565b6101dc610284366004612573565b610b1d565b6101dc610bbd565b6101dc61029f36600461268e565b610c3a565b6101dc6102b23660046126be565b610c85565b6101dc6102c536600461268e565b610d10565b6000546040516001600160a01b0390911681526020015b60405180910390f35b6101dc6102f836600461273f565b610d5e565b6101dc61030b3660046121d4565b610db9565b6101dc61031e3660046127dc565b610df4565b6101dc6103313660046128d8565b610e50565b6101dc61034436600461296b565b610f31565b6101dc610357366004612a76565b610fd4565b6101dc61036a366004612ae0565b6110c9565b6101dc61037d366004612b39565b61111f565b6101dc610390366004612bbd565b6111e0565b6103b86103a336600461268e565b60056020526000908152604090205460ff1681565b60405190151581526020016102e1565b6101dc6103d636600461268e565b611236565b6101dc6103e9366004612c49565b611320565b6101dc6103fc366004612d67565b611642565b60025461041f90869086906001600160a01b031660808b013561169b565b600254604051635fcc4a3f60e01b81526001600160a01b0390911690635fcc4a3f90610457908a908a90889088908890600401612e3b565b600060405180830381600087803b15801561047157600080fd5b505af1158015610485573d6000803e3d6000fd5b5050505050505050505050565b6001600160a01b03851663d505accf84306104b360a0860160808701612f1c565b6104bd57876104c1565b6000195b60608601356104d36020880188612f39565b876020013588604001356040518863ffffffff1660e01b81526004016104ff9796959493929190612f54565b600060405180830381600087803b15801561051957600080fd5b505af115801561052d573d6000803e3d6000fd5b505060035461054b925087915085906001600160a01b03168761169b565b600354604051630529fc9f60e31b81526001600160a01b039091169063294fe4f890610581908890889088908890600401612fe4565b600060405180830381600087803b15801561059b57600080fd5b505af11580156105af573d6000803e3d6000fd5b505050505050505050565b6105c68686308561169b565b6004805460405163095ea7b360e01b81526001600160a01b03918216928101929092526024820184905287169063095ea7b3906044016020604051808303816000875af115801561061b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061063f9190613010565b506004805460405163966e0fe960e01b81526001600160a01b039091169163966e0fe991610679918a918a918a918a918a918a9101613063565b600060405180830381600087803b15801561069357600080fd5b505af11580156106a7573d6000803e3d6000fd5b50505050505050505050565b6004546000906106d2908890889088906001600160a01b0316866116fb565b6004805460405163b04640b760e01b81529293506001600160a01b03169163b04640b79161070a918c9186918b918b918b91016130a5565b600060405180830381600087803b15801561072457600080fd5b505af1158015610738573d6000803e3d6000fd5b505050505050505050505050565b600354604051630872c2cd60e31b81526001600160a01b0390911690634396166890610457908a908a908a908a908a908a908a90600401613107565b60035461079c90859084906001600160a01b03168661169b565b600354604051630529fc9f60e31b81526001600160a01b039091169063294fe4f8906107d2908790879087908790600401612fe4565b600060405180830381600087803b1580156107ec57600080fd5b505af1158015610800573d6000803e3d6000fd5b5050505050505050565b60005b8a81101561089e5760008a8a8381811061082957610829613146565b90506020020135111561088c5761088c8c8c8381811061084b5761084b613146565b9050602002016020810190610860919061268e565b6003548a906001600160a01b03168d8d8681811061088057610880613146565b9050602002013561169b565b8061089681613172565b91505061080d565b50600354604051630b505d3760e41b81526001600160a01b039091169063b505d370906108e3908e908e908e908e908e908b908f908f908f908e908e9060040161323d565b600060405180830381600087803b1580156108fd57600080fd5b505af1158015610911573d6000803e3d6000fd5b505050505050505050505050505050565b60005b60038110156109b557600082826003811061094257610942613146565b602002013511156109a3576109a383826003811061096257610962613146565b602002016020810190610975919061268e565b6003805487916001600160a01b039091169086908690811061099957610999613146565b602002013561169b565b806109ad81613172565b915050610925565b506002546040516301c2bc7d60e51b81526001600160a01b03909116906338578fa0906106799089908990899088908890600401613316565b6004805460405163bca7382360e01b81529182018790526000916001600160a01b039091169063bca7382390602401602060405180830381865afa158015610a3a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a5e91906133e3565b9050610a6c8185308861169b565b6004805460405163095ea7b360e01b81526001600160a01b03918216928101929092526024820187905282169063095ea7b3906044016020604051808303816000875af1158015610ac1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ae59190613010565b506004805460405163b04640b760e01b81526001600160a01b039091169163b04640b791610679918a918a918a918a918a91016130a5565b6080880151600254600091610b409189919089906001600160a01b0316866117eb565b60808a01819052600254604051635fcc4a3f60e01b81529192506001600160a01b031690635fcc4a3f90610b80908c908c908a908a908a90600401613400565b600060405180830381600087803b158015610b9a57600080fd5b505af1158015610bae573d6000803e3d6000fd5b50505050505050505050505050565b6000546001600160a01b03163314610bf05760405162461bcd60e51b8152600401610be79061347d565b60405180910390fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b03163314610c645760405162461bcd60e51b8152600401610be79061347d565b6001600160a01b03166000908152600560205260409020805460ff19169055565b60005b6003811015610cd7576000828260038110610ca557610ca5613146565b60200201351115610cc557610cc583826003811061096257610962613146565b80610ccf81613172565b915050610c88565b50600254604051634638ef3f60e11b81526001600160a01b0390911690638c71de7e906106799089908990899088908890600401613611565b6000546001600160a01b03163314610d3a5760405162461bcd60e51b8152600401610be79061347d565b6001600160a01b03166000908152600560205260409020805460ff19166001179055565b600454600090610d7d908a908a908a906001600160a01b0316866116fb565b600480546040516317ddce6f60e31b81529293506001600160a01b03169163beee737891610b80918d9186918d918d918d918d918d910161365d565b60048054604051639500125b60e01b81526001600160a01b0390911691639500125b91610457918b918b918b918b918b918b918b9101613107565b600354600090610e13908a908a908a906001600160a01b0316866117eb565b600354604051632b17790f60e21b81529192506001600160a01b03169063ac5de43c90610b80908c9085908c908c908c908c908c9060040161369f565b60005b6003811015610ef8576000838260038110610e7057610e70613146565b60200201511115610ee657610ece848260038110610e9057610e90613146565b602002016020810190610ea3919061268e565b848360038110610eb557610eb5613146565b602002015160025488906001600160a01b0316866117eb565b838260038110610ee057610ee0613146565b60200201525b80610ef081613172565b915050610e53565b50600254604051634638ef3f60e11b81526001600160a01b0390911690638c71de7e90610457908a908a908a9089908990600401613716565b60005b6003811015610f9b576000838260038110610f5157610f51613146565b60200201511115610f8957610f71848260038110610e9057610e90613146565b838260038110610f8357610f83613146565b60200201525b80610f9381613172565b915050610f34565b506002546040516301c2bc7d60e51b81526001600160a01b03909116906338578fa090610457908a908a908a9089908990600401613758565b6001600160a01b03861663d505accf8530610ff560a0870160808801612f1c565b610fff5788611003565b6000195b60608701356110156020890189612f39565b886020013589604001356040518863ffffffff1660e01b81526004016110419796959493929190612f54565b600060405180830381600087803b15801561105b57600080fd5b505af115801561106f573d6000803e3d6000fd5b50506003546000925061109291508890889088906001600160a01b0316866117eb565b600354604051630529fc9f60e31b81529192506001600160a01b03169063294fe4f890610457908a9085908a908a90600401612fe4565b6003546000906110e8908790879087906001600160a01b0316866117eb565b600354604051630529fc9f60e31b81529192506001600160a01b03169063294fe4f890610679908990859089908990600401612fe4565b61112b8786308961169b565b6004805460405163095ea7b360e01b81526001600160a01b03918216928101929092526024820188905288169063095ea7b3906044016020604051808303816000875af1158015611180573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111a49190613010565b50600480546040516317ddce6f60e31b81526001600160a01b039091169163beee737891610457918b918b918b918b918b918b918b910161365d565b6003546111fa90889087906001600160a01b03168961169b565b600354604051632b17790f60e21b81526001600160a01b039091169063ac5de43c90610457908a908a908a908a908a908a908a9060040161369f565b6000546001600160a01b031633146112605760405162461bcd60e51b8152600401610be79061347d565b6001600160a01b0381166112c55760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610be7565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6000896001600160401b0381111561133a5761133a612522565b604051908082528060200260200182016040528015611363578160200160208202803683370190505b50905060005b8c8110156115ad5760008c8c8381811061138557611385613146565b90506020020135111561159b578484828181106113a4576113a4613146565b6113ba92602060a0909202019081019150612f39565b60ff16156115205760008585838181106113d6576113d6613146565b905060a0020160800160208101906113ee9190612f1c565b611410578c8c8381811061140457611404613146565b90506020020135611414565b6000195b90508e8e8381811061142857611428613146565b905060200201602081019061143d919061268e565b6001600160a01b031663d505accf8c30848a8a8881811061146057611460613146565b905060a00201606001358b8b8981811061147c5761147c613146565b61149292602060a0909202019081019150612f39565b8c8c8a8181106114a4576114a4613146565b905060a00201602001358d8d8b8181106114c0576114c0613146565b905060a00201604001356040518863ffffffff1660e01b81526004016114ec9796959493929190612f54565b600060405180830381600087803b15801561150657600080fd5b505af115801561151a573d6000803e3d6000fd5b50505050505b61157c8e8e8381811061153557611535613146565b905060200201602081019061154a919061268e565b8d8d8481811061155c5761155c613146565b60035460209091029290920135918e91506001600160a01b0316876117eb565b82828151811061158e5761158e613146565b6020026020010181815250505b806115a581613172565b915050611369565b50600360009054906101000a90046001600160a01b03166001600160a01b031663b505d3708e8e848d8a8e8e8e8d8d6040518b63ffffffff1660e01b81526004016116019a999897969594939291906137f9565b600060405180830381600087803b15801561161b57600080fd5b505af115801561162f573d6000803e3d6000fd5b5050505050505050505050505050505050565b60045460009061166190899086908a906001600160a01b0316866116fb565b6004805460405163966e0fe960e01b81529293506001600160a01b03169163966e0fe99161070a918c918c918c918c9189918c9101613063565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b1790526116f5908590611864565b50505050565b600061170886858461193b565b611715868533853561169b565b6117208235866138b3565b60405163095ea7b360e01b81526001600160a01b038581166004830152602482018390529192509087169063095ea7b3906044016020604051808303816000875af1158015611773573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117979190613010565b50604080516001600160a01b038881168252843560208301523392908716917f8adfc333732362694bf9044b570e1ae5d25a5b143b8aa3c267aee92adae48bbe91015b60405180910390a395945050505050565b60006117f886858461193b565b611805868533853561169b565b6118108235866138b3565b905061181e8685858461169b565b604080516001600160a01b038881168252843560208301523392908716917f8adfc333732362694bf9044b570e1ae5d25a5b143b8aa3c267aee92adae48bbe91016117da565b60006118b9826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316611b9c9092919063ffffffff16565b80519091501561193657808060200190518101906118d79190613010565b6119365760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610be7565b505050565b60408051466020808301919091526bffffffffffffffffffffffff19606087811b8216848601528535605485015286811b8216607485015233901b16608883015283810135609c808401919091528351808403909101815260bc90920183528151910120906000908301816020020160208101906119b99190612f39565b60405160f89190911b6001600160f81b03191660208201526080840135602182015260c084013560418201526061016040516020818303038152906040528051906020012090506000611a34611a0e84611bb5565b611a1e6060870160408801612f39565b608087013560c0880160005b6020020135611c08565b90506000611a64611a4484611bb5565b611a546080880160608901612f39565b60a088013560c089016001611a2a565b6001600160a01b03831660009081526005602052604090205490915060ff16611add5760405162461bcd60e51b815260206004820152602560248201527f526f757465723a20696e76616c6964207369676e61747572652066726f6d207760448201526437b935b2b960d91b6064820152608401610be7565b856001600160a01b0316816001600160a01b031614611b4c5760405162461bcd60e51b815260206004820152602560248201527f526f757465723a20696e76616c6964207369676e61747572652066726f6d207360448201526432b73232b960d91b6064820152608401610be7565b8460200135421115611b935760405162461bcd60e51b815260206004820152601060248201526f526f757465723a20646561646c696e6560801b6044820152606401610be7565b50505050505050565b6060611bab8484600085611db1565b90505b9392505050565b6040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0821115611c855760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610be7565b8360ff16601b1480611c9a57508360ff16601c145b611cf15760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610be7565b6040805160008082526020820180845288905260ff871692820192909252606081018590526080810184905260019060a0016020604051602081039080840390855afa158015611d45573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116611da85760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610be7565b95945050505050565b606082471015611e125760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610be7565b843b611e605760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610be7565b600080866001600160a01b03168587604051611e7c91906138f6565b60006040518083038185875af1925050503d8060008114611eb9576040519150601f19603f3d011682016040523d82523d6000602084013e611ebe565b606091505b5091509150611ece828286611ed9565b979650505050505050565b60608315611ee8575081611bae565b825115611ef85782518084602001fd5b8160405162461bcd60e51b8152600401610be79190613912565b60006101008284031215611f2557600080fd5b50919050565b600060a08284031215611f2557600080fd5b6001600160a01b0381168114611f5257600080fd5b50565b8035611f6081611f3d565b919050565b6000806000806000806000610240888a031215611f8157600080fd5b611f8b8989611f12565b9650611f9b896101008a01611f2b565b95506101a0880135611fac81611f3d565b94506101c0880135611fbd81611f3d565b93506101e0880135611fce81611f3d565b9250610200880135611fdf81611f3d565b80925050610220880135905092959891949750929550565b600060808284031215611f2557600080fd5b6000806000806000610180868803121561202257600080fd5b853561202d81611f3d565b945060208601359350604086013561204481611f3d565b92506120538760608801611ff7565b91506120628760e08801611f2b565b90509295509295909350565b60008083601f84011261208057600080fd5b5081356001600160401b0381111561209757600080fd5b6020830191508360208260051b85010111156120b257600080fd5b9250929050565b60008060008060008060a087890312156120d257600080fd5b86356120dd81611f3d565b955060208701356120ed81611f3d565b945060408701356001600160401b0381111561210857600080fd5b61211489828a0161206e565b979a9699509760608101359660809091013595509350505050565b600060608284031215611f2557600080fd5b6000806000806000806000610200888a03121561215d57600080fd5b87359650602088013561216f81611f3d565b955060408801359450606088013561218681611f3d565b9350608088013561219681611f3d565b92506121a58960a08a0161212f565b91506121b5896101008a01611f12565b905092959891949750929550565b803560ff81168114611f6057600080fd5b600080600080600080600060e0888a0312156121ef57600080fd5b87359650602088013561220181611f3d565b9550604088013561221181611f3d565b945060608801359350612226608089016121c3565b925060a0880135915060c0880135905092959891949750929550565b60008060008060e0858703121561225857600080fd5b843561226381611f3d565b935060208501359250604085013561227a81611f3d565b91506122898660608701611ff7565b905092959194509250565b80356001600160e01b031981168114611f6057600080fd5b60008083601f8401126122be57600080fd5b5081356001600160401b038111156122d557600080fd5b6020830191508360208285010111156120b257600080fd5b60008083601f8401126122ff57600080fd5b5081356001600160401b0381111561231657600080fd5b60208301915083602060a0830285010111156120b257600080fd5b60008060008060008060008060008060006101408c8e03121561235357600080fd5b6001600160401b03808d35111561236957600080fd5b6123768e8e358f0161206e565b909c509a5060208d013581101561238c57600080fd5b61239c8e60208f01358f0161206e565b909a5098506123ad60408e01611f55565b97506123bb60608e01612294565b96508060808e013511156123ce57600080fd5b6123de8e60808f01358f016122ac565b90965094506123f08e60a08f01611ff7565b9350806101208e0135111561240457600080fd5b506124168d6101208e01358e016122ed565b81935080925050509295989b509295989b9093969950565b806060810183101561243f57600080fd5b92915050565b600080600080600080610200878903121561245f57600080fd5b6124698888611f12565b95506101008701356001600160401b0381111561248557600080fd5b61249189828a016122ed565b9096509450506101208701356124a681611f3d565b92506124b688610140890161242e565b91506124c6886101a0890161242e565b90509295509295509295565b600080600080600060e086880312156124ea57600080fd5b8535945060208601359350604086013561250381611f3d565b9250606086013561251381611f3d565b9150612062876080880161212f565b634e487b7160e01b600052604160045260246000fd5b60405161010081016001600160401b038111828210171561255b5761255b612522565b60405290565b8035600f81900b8114611f6057600080fd5b600080600080600080600080888a0361034081121561259157600080fd5b610100808212156125a157600080fd5b6125a9612538565b91508a356125b681611f3d565b82526125c460208c01612561565b602083015260408b013560408301526125df60608c01611f55565b606083015260808b013560808301526125fa60a08c01612561565b60a083015260c08b013560c083015261261560e08c01611f55565b60e08301528199506126298c828d01611f2b565b9850505061263a6101a08a01611f55565b95506126496101c08a01611f55565b94506126586101e08a01611f55565b93506126676102008a01611f55565b9250610220890135915061267f8a6102408b01611f12565b90509295985092959890939650565b6000602082840312156126a057600080fd5b8135611bae81611f3d565b60006102008284031215611f2557600080fd5b60008060008060008061030087890312156126d857600080fd5b6126e288886126ab565b95506102008701356001600160401b038111156126fe57600080fd5b61270a89828a016122ed565b90965094505061022087013561271f81611f3d565b925061272f88610240890161242e565b91506124c6886102a0890161242e565b6000806000806000806000806101e0898b03121561275c57600080fd5b883561276781611f3d565b975060208901359650604089013561277e81611f3d565b9550606089013561278e81611f3d565b9450608089013561279e81611f3d565b935060a08901356127ae81611f3d565b925060c0890135915061267f8a60e08b01611f12565b80356001600160f81b031981168114611f6057600080fd5b6000806000806000806000806101c0898b0312156127f957600080fd5b883561280481611f3d565b975060208901359650604089013561281b81611f3d565b955060608901356001600160401b0381111561283657600080fd5b6128428b828c0161206e565b9096509450612855905060808a016127c4565b925060a0890135915061267f8a60c08b01611f12565b600082601f83011261287c57600080fd5b604051606081018181106001600160401b038211171561289e5761289e612522565b6040528060608401858111156128b357600080fd5b845b818110156128cd5780358352602092830192016128b5565b509195945050505050565b6000806000806000806000610400888a0312156128f457600080fd5b6128fe89896126ab565b96506102008801356001600160401b0381111561291a57600080fd5b6129268a828b016122ed565b90975095505061022088013561293b81611f3d565b935061294b896102408a0161242e565b925061295b896102a08a0161286b565b91506121b5896103008a01611f12565b600080600080600080600087890361030081121561298857600080fd5b6101008082121561299857600080fd5b6129a0612538565b915089356129ad81611f3d565b8083525060208a0135602083015260408a013560408301526129d160608b01611f55565b606083015260808a013560808301526129ec60a08b01611f55565b60a08301526129fd60c08b01611f55565b60c083015260e08a810135908301529097508801356001600160401b03811115612a2657600080fd5b612a328a828b016122ed565b9097509550612a4690506101208901611f55565b9350612a56896101408a0161242e565b9250612a66896101a08a0161286b565b91506121b5896102008a01611f12565b6000806000806000806102808789031215612a9057600080fd5b8635612a9b81611f3d565b9550602087013594506040870135612ab281611f3d565b9350612ac18860608901611ff7565b9250612ad08860e08901611f2b565b91506124c6886101808901611f12565b60008060008060006101e08688031215612af957600080fd5b8535612b0481611f3d565b9450602086013593506040860135612b1b81611f3d565b9250612b2a8760608801611ff7565b91506120628760e08801611f12565b600080600080600080600060e0888a031215612b5457600080fd5b8735612b5f81611f3d565b9650602088013595506040880135612b7681611f3d565b94506060880135612b8681611f3d565b93506080880135612b9681611f3d565b925060a0880135612ba681611f3d565b8092505060c0880135905092959891949750929550565b600080600080600080600060c0888a031215612bd857600080fd5b8735612be381611f3d565b9650602088013595506040880135612bfa81611f3d565b945060608801356001600160401b03811115612c1557600080fd5b612c218a828b0161206e565b9095509350612c349050608089016127c4565b915060a0880135905092959891949750929550565b6000806000806000806000806000806000806102408d8f031215612c6c57600080fd5b6001600160401b038d351115612c8157600080fd5b612c8e8e8e358f0161206e565b909c509a506001600160401b0360208e01351115612cab57600080fd5b612cbb8e60208f01358f0161206e565b909a509850612ccc60408e01611f55565b9750612cda60608e01612294565b96506001600160401b0360808e01351115612cf457600080fd5b612d048e60808f01358f016122ac565b9096509450612d168e60a08f01611ff7565b93506001600160401b036101208e01351115612d3157600080fd5b612d428e6101208f01358f016122ed565b9093509150612d558e6101408f01611f12565b90509295989b509295989b509295989b565b60008060008060008060006101a0888a031215612d8357600080fd5b8735612d8e81611f3d565b96506020880135612d9e81611f3d565b955060408801356001600160401b03811115612db957600080fd5b612dc58a828b0161206e565b90965094505060608801359250608088013591506121b58960a08a01611f12565b8015158114611f5257600080fd5b60ff612dff826121c3565b1682526020810135602083015260408101356040830152606081013560608301526080810135612e2e81612de6565b8015156080840152505050565b61020081018635612e4b81611f3d565b6001600160a01b03168252612e6260208801612561565b600f0b6020830152604087013560408301526060870135612e8281611f3d565b6001600160a01b0316606083015260808781013590830152612ea660a08801612561565b612eb560a0840182600f0b9052565b5060c087013560c0830152612ecc60e08801611f55565b6001600160a01b03811660e08401525b50612eeb610100830187612df4565b6001600160a01b0385166101a08301526001600160a01b0384166101c0830152826101e08301529695505050505050565b600060208284031215612f2e57600080fd5b8135611bae81612de6565b600060208284031215612f4b57600080fd5b611bae826121c3565b6001600160a01b0397881681529590961660208601526040850193909352606084019190915260ff16608083015260a082015260c081019190915260e00190565b8035612fa081611f3d565b6001600160a01b039081168352602082013590612fbc82611f3d565b9081166020840152604082013590612fd382611f3d565b166040830152606090810135910152565b6001600160a01b038581168252602082018590528316604082015260e08101611da86060830184612f95565b60006020828403121561302257600080fd5b8151611bae81612de6565b81835260006001600160fb1b0383111561304657600080fd5b8260051b8083602087013760009401602001938452509192915050565b6001600160a01b0387811682528616602082015260a060408201819052600090613090908301868861302d565b60608301949094525060800152949350505050565b858152602081018590526001600160a01b038481166040830152838116606083015260e082019083356130d781611f3d565b8116608084015260208401356130ec81611f3d565b1660a08301526040929092013560c090910152949350505050565b9687526001600160a01b039586166020880152939094166040860152606085019190915260ff16608084015260a083019190915260c082015260e00190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60006000198214156131865761318661315c565b5060010190565b8183526000602080850194508260005b858110156131cb5781356131b081611f3d565b6001600160a01b03168752958201959082019060010161319d565b509495945050505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b81835260208301925060008160005b84811015613233576132208683612df4565b60a095860195919091019060010161320e565b5093949350505050565b6101408152600061325361014083018d8f61318d565b8281036020840152613266818c8e61302d565b6001600160a01b038b1660408501529050613284606084018a612f95565b6001600160e01b0319881660e08401528281036101008401526132a88187896131d6565b90508281036101208401526132be8185876131ff565b9e9d5050505050505050505050505050565b8060005b60038110156116f55781356132e881611f3d565b6001600160a01b0316845260209384019391909101906001016132d4565b6060818337506000606082015250565b60006101e0873561332681611f3d565b6001600160a01b039081168452602089810135908501526040808a01359085015260608901359061335682611f3d565b1660608401526080888101359084015261337260a08901611f55565b6001600160a01b031660a084015261338c60c08901611f55565b6001600160a01b03811660c08501525060e088013560e0840152806101008401526133ba81840187896131ff565b9150506133cb6101208301856132d0565b6133d9610180830184613306565b9695505050505050565b6000602082840312156133f557600080fd5b8151611bae81611f3d565b60006102008201905060018060a01b038088511683526020880151600f0b602084015260408801516040840152806060890151166060840152506080870151608083015260a087015161345860a0840182600f0b9052565b5060c087015160c083015260e0870151612edc60e08401826001600160a01b03169052565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6134cc826134bf83611f55565b6001600160a01b03169052565b6134d860208201611f55565b6001600160a01b031660208301526134f260408201611f55565b6001600160a01b031660408301526060818101359083015261351660808201612561565b6135256080840182600f0b9052565b5061353260a08201612561565b61354160a0840182600f0b9052565b5060c081013560c083015261355860e08201612561565b61356760e0840182600f0b9052565b506101008181013590830152610120613581818301611f55565b6001600160a01b03169083015261014061359c828201611f55565b6001600160a01b0316908301526101606135b7828201611f55565b6001600160a01b0316908301526101806135d2828201611f55565b6001600160a01b0316908301526101a081810135908301526101c06135f8818301611f55565b6001600160a01b0316908301526101e090810135910152565b60006102e0808301613623848a6134b2565b8161020085015261363581888a6131ff565b92506136456102208501876132d0565b60608561028086013760009052509695505050505050565b6001600160a01b0397881681526020810196909652938616604086015291851660608501528416608084015290921660a082015260c081019190915260e00190565b6001600160a01b038881168252602082018890528616604082015260c0606082018190526000906136d3908301868861302d565b6001600160f81b03199490941660808301525060a0015295945050505050565b8060005b60038110156116f55781518452602093840193909101906001016136f7565b60006102e061372583896134b2565b8061020084015261373981840187896131ff565b91505061374a6102208301856132d0565b6133d96102808301846136f3565b60006101e060018060a01b0380895116845260208901516020850152604089015160408501528060608a0151166060850152608089015160808501528060a08a01511660a08501525060c08801516137bb60c08501826001600160a01b03169052565b5060e088015160e0840152806101008401526137da81840187896131ff565b9150506137eb6101208301856132d0565b6133d96101808301846136f3565b600061014080835261380e8184018d8f61318d565b9050602083820381850152818c518084528284019150828e01935060005b818110156138485784518352938301939183019160010161382c565b50506001600160a01b038c166040860152613866606086018c612f95565b6001600160e01b03198a1660e086015284810361010086015261388a81898b6131d6565b925050508281036101208401526138a28185876131ff565b9d9c50505050505050505050505050565b6000828210156138c5576138c561315c565b500390565b60005b838110156138e55781810151838201526020016138cd565b838111156116f55750506000910152565b600082516139088184602087016138ca565b9190910192915050565b60208152600082518060208401526139318160408501602087016138ca565b601f01601f1916919091016040019291505056fea2646970667358221220d933d9560e79d3a20c15755ef6cb7fbfe06c48f9feccd5af8554823e3d3fb2eb64736f6c634300080a0033",
}

// RouterABI is the input ABI used to generate the binding from.
// Deprecated: Use RouterMetaData.ABI instead.
var RouterABI = RouterMetaData.ABI

// RouterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RouterMetaData.Bin instead.
var RouterBin = RouterMetaData.Bin

// DeployRouter deploys a new Ethereum contract, binding an instance of Router to it.
func DeployRouter(auth *bind.TransactOpts, backend bind.ContractBackend, portal common.Address, synthesis common.Address, curveProxy common.Address, localTreasury common.Address) (common.Address, *types.Transaction, *Router, error) {
	parsed, err := RouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RouterBin), backend, portal, synthesis, curveProxy, localTreasury)
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

// DelegatedBatchSynthesizeRequestWithDataTransit is a paid mutator transaction binding the contract method 0xf3a08186.
//
// Solidity: function delegatedBatchSynthesizeRequestWithDataTransit(address[] token, uint256[] amount, address from, bytes4 selector, bytes transitData, (address,address,address,uint256) synthParams, (uint8,bytes32,bytes32,uint256,bool)[] permitData, (uint256,uint256,uint8[2],bytes32[2],bytes32[2]) receipt) returns()
func (_Router *RouterTransactor) DelegatedBatchSynthesizeRequestWithDataTransit(opts *bind.TransactOpts, token []common.Address, amount []*big.Int, from common.Address, selector [4]byte, transitData []byte, synthParams IPortalSynthParams, permitData []IPortalPermitData, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "delegatedBatchSynthesizeRequestWithDataTransit", token, amount, from, selector, transitData, synthParams, permitData, receipt)
}
func (_Router *RouterTransactor) DelegatedBatchSynthesizeRequestWithDataTransitOverGsn(opts *bind.TransactOpts, token []common.Address, amount []*big.Int, from common.Address, selector [4]byte, transitData []byte, synthParams IPortalSynthParams, permitData []IPortalPermitData, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "delegatedBatchSynthesizeRequestWithDataTransit", token, amount, from, selector, transitData, synthParams, permitData, receipt)
}

// DelegatedBatchSynthesizeRequestWithDataTransit is a paid mutator transaction binding the contract method 0xf3a08186.
//
// Solidity: function delegatedBatchSynthesizeRequestWithDataTransit(address[] token, uint256[] amount, address from, bytes4 selector, bytes transitData, (address,address,address,uint256) synthParams, (uint8,bytes32,bytes32,uint256,bool)[] permitData, (uint256,uint256,uint8[2],bytes32[2],bytes32[2]) receipt) returns()
func (_Router *RouterSession) DelegatedBatchSynthesizeRequestWithDataTransit(token []common.Address, amount []*big.Int, from common.Address, selector [4]byte, transitData []byte, synthParams IPortalSynthParams, permitData []IPortalPermitData, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.DelegatedBatchSynthesizeRequestWithDataTransit(&_Router.TransactOpts, token, amount, from, selector, transitData, synthParams, permitData, receipt)
}
func (_Router *RouterSession) DelegatedBatchSynthesizeRequestWithDataTransitOverGsn(token []common.Address, amount []*big.Int, from common.Address, selector [4]byte, transitData []byte, synthParams IPortalSynthParams, permitData []IPortalPermitData, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.DelegatedBatchSynthesizeRequestWithDataTransitOverGsn(&_Router.TransactOpts, token, amount, from, selector, transitData, synthParams, permitData, receipt)
}

// DelegatedBatchSynthesizeRequestWithDataTransit is a paid mutator transaction binding the contract method 0xf3a08186.
//
// Solidity: function delegatedBatchSynthesizeRequestWithDataTransit(address[] token, uint256[] amount, address from, bytes4 selector, bytes transitData, (address,address,address,uint256) synthParams, (uint8,bytes32,bytes32,uint256,bool)[] permitData, (uint256,uint256,uint8[2],bytes32[2],bytes32[2]) receipt) returns()
func (_Router *RouterTransactorSession) DelegatedBatchSynthesizeRequestWithDataTransit(token []common.Address, amount []*big.Int, from common.Address, selector [4]byte, transitData []byte, synthParams IPortalSynthParams, permitData []IPortalPermitData, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.DelegatedBatchSynthesizeRequestWithDataTransit(&_Router.TransactOpts, token, amount, from, selector, transitData, synthParams, permitData, receipt)
}
func (_Router *RouterTransactorSession) DelegatedBatchSynthesizeRequestWithDataTransitOverGsn(token []common.Address, amount []*big.Int, from common.Address, selector [4]byte, transitData []byte, synthParams IPortalSynthParams, permitData []IPortalPermitData, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.DelegatedBatchSynthesizeRequestWithDataTransitOverGsn(&_Router.TransactOpts, token, amount, from, selector, transitData, synthParams, permitData, receipt)
}

// DelegatedMetaExchangeRequestVia3pool is a paid mutator transaction binding the contract method 0xb981b485.
//
// Solidity: function delegatedMetaExchangeRequestVia3pool((address,address,address,uint256,int128,int128,uint256,int128,uint256,address,address,address,address,uint256,address,uint256) params, (uint8,bytes32,bytes32,uint256,bool)[] permit, address from, address[3] token, uint256[3] amount, (uint256,uint256,uint8[2],bytes32[2],bytes32[2]) receipt) returns()
func (_Router *RouterTransactor) DelegatedMetaExchangeRequestVia3pool(opts *bind.TransactOpts, params ICurveProxyMetaExchangeParams, permit []ICurveProxyPermitData, from common.Address, token [3]common.Address, amount [3]*big.Int, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "delegatedMetaExchangeRequestVia3pool", params, permit, from, token, amount, receipt)
}
func (_Router *RouterTransactor) DelegatedMetaExchangeRequestVia3poolOverGsn(opts *bind.TransactOpts, params ICurveProxyMetaExchangeParams, permit []ICurveProxyPermitData, from common.Address, token [3]common.Address, amount [3]*big.Int, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "delegatedMetaExchangeRequestVia3pool", params, permit, from, token, amount, receipt)
}

// DelegatedMetaExchangeRequestVia3pool is a paid mutator transaction binding the contract method 0xb981b485.
//
// Solidity: function delegatedMetaExchangeRequestVia3pool((address,address,address,uint256,int128,int128,uint256,int128,uint256,address,address,address,address,uint256,address,uint256) params, (uint8,bytes32,bytes32,uint256,bool)[] permit, address from, address[3] token, uint256[3] amount, (uint256,uint256,uint8[2],bytes32[2],bytes32[2]) receipt) returns()
func (_Router *RouterSession) DelegatedMetaExchangeRequestVia3pool(params ICurveProxyMetaExchangeParams, permit []ICurveProxyPermitData, from common.Address, token [3]common.Address, amount [3]*big.Int, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.DelegatedMetaExchangeRequestVia3pool(&_Router.TransactOpts, params, permit, from, token, amount, receipt)
}
func (_Router *RouterSession) DelegatedMetaExchangeRequestVia3poolOverGsn(params ICurveProxyMetaExchangeParams, permit []ICurveProxyPermitData, from common.Address, token [3]common.Address, amount [3]*big.Int, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.DelegatedMetaExchangeRequestVia3poolOverGsn(&_Router.TransactOpts, params, permit, from, token, amount, receipt)
}

// DelegatedMetaExchangeRequestVia3pool is a paid mutator transaction binding the contract method 0xb981b485.
//
// Solidity: function delegatedMetaExchangeRequestVia3pool((address,address,address,uint256,int128,int128,uint256,int128,uint256,address,address,address,address,uint256,address,uint256) params, (uint8,bytes32,bytes32,uint256,bool)[] permit, address from, address[3] token, uint256[3] amount, (uint256,uint256,uint8[2],bytes32[2],bytes32[2]) receipt) returns()
func (_Router *RouterTransactorSession) DelegatedMetaExchangeRequestVia3pool(params ICurveProxyMetaExchangeParams, permit []ICurveProxyPermitData, from common.Address, token [3]common.Address, amount [3]*big.Int, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.DelegatedMetaExchangeRequestVia3pool(&_Router.TransactOpts, params, permit, from, token, amount, receipt)
}
func (_Router *RouterTransactorSession) DelegatedMetaExchangeRequestVia3poolOverGsn(params ICurveProxyMetaExchangeParams, permit []ICurveProxyPermitData, from common.Address, token [3]common.Address, amount [3]*big.Int, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.DelegatedMetaExchangeRequestVia3poolOverGsn(&_Router.TransactOpts, params, permit, from, token, amount, receipt)
}

// DelegatedMintEusdRequestVia3pool is a paid mutator transaction binding the contract method 0xbd7d4b96.
//
// Solidity: function delegatedMintEusdRequestVia3pool((address,uint256,uint256,address,uint256,address,address,uint256) params, (uint8,bytes32,bytes32,uint256,bool)[] permit, address from, address[3] token, uint256[3] amount, (uint256,uint256,uint8[2],bytes32[2],bytes32[2]) receipt) returns()
func (_Router *RouterTransactor) DelegatedMintEusdRequestVia3pool(opts *bind.TransactOpts, params ICurveProxyMetaMintEUSD, permit []ICurveProxyPermitData, from common.Address, token [3]common.Address, amount [3]*big.Int, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "delegatedMintEusdRequestVia3pool", params, permit, from, token, amount, receipt)
}
func (_Router *RouterTransactor) DelegatedMintEusdRequestVia3poolOverGsn(opts *bind.TransactOpts, params ICurveProxyMetaMintEUSD, permit []ICurveProxyPermitData, from common.Address, token [3]common.Address, amount [3]*big.Int, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "delegatedMintEusdRequestVia3pool", params, permit, from, token, amount, receipt)
}

// DelegatedMintEusdRequestVia3pool is a paid mutator transaction binding the contract method 0xbd7d4b96.
//
// Solidity: function delegatedMintEusdRequestVia3pool((address,uint256,uint256,address,uint256,address,address,uint256) params, (uint8,bytes32,bytes32,uint256,bool)[] permit, address from, address[3] token, uint256[3] amount, (uint256,uint256,uint8[2],bytes32[2],bytes32[2]) receipt) returns()
func (_Router *RouterSession) DelegatedMintEusdRequestVia3pool(params ICurveProxyMetaMintEUSD, permit []ICurveProxyPermitData, from common.Address, token [3]common.Address, amount [3]*big.Int, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.DelegatedMintEusdRequestVia3pool(&_Router.TransactOpts, params, permit, from, token, amount, receipt)
}
func (_Router *RouterSession) DelegatedMintEusdRequestVia3poolOverGsn(params ICurveProxyMetaMintEUSD, permit []ICurveProxyPermitData, from common.Address, token [3]common.Address, amount [3]*big.Int, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.DelegatedMintEusdRequestVia3poolOverGsn(&_Router.TransactOpts, params, permit, from, token, amount, receipt)
}

// DelegatedMintEusdRequestVia3pool is a paid mutator transaction binding the contract method 0xbd7d4b96.
//
// Solidity: function delegatedMintEusdRequestVia3pool((address,uint256,uint256,address,uint256,address,address,uint256) params, (uint8,bytes32,bytes32,uint256,bool)[] permit, address from, address[3] token, uint256[3] amount, (uint256,uint256,uint8[2],bytes32[2],bytes32[2]) receipt) returns()
func (_Router *RouterTransactorSession) DelegatedMintEusdRequestVia3pool(params ICurveProxyMetaMintEUSD, permit []ICurveProxyPermitData, from common.Address, token [3]common.Address, amount [3]*big.Int, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.DelegatedMintEusdRequestVia3pool(&_Router.TransactOpts, params, permit, from, token, amount, receipt)
}
func (_Router *RouterTransactorSession) DelegatedMintEusdRequestVia3poolOverGsn(params ICurveProxyMetaMintEUSD, permit []ICurveProxyPermitData, from common.Address, token [3]common.Address, amount [3]*big.Int, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.DelegatedMintEusdRequestVia3poolOverGsn(&_Router.TransactOpts, params, permit, from, token, amount, receipt)
}

// DelegatedRedeemEusdRequest is a paid mutator transaction binding the contract method 0x62e41ded.
//
// Solidity: function delegatedRedeemEusdRequest((address,int128,uint256,address,uint256,int128,uint256,address) params, (uint8,bytes32,bytes32,uint256,bool) permit, address payToken, address from, address receiveSide, address oppositeBridge, uint256 chainId, (uint256,uint256,uint8[2],bytes32[2],bytes32[2]) receipt) returns()
func (_Router *RouterTransactor) DelegatedRedeemEusdRequest(opts *bind.TransactOpts, params ICurveProxyMetaRedeemEUSD, permit ICurveProxyPermitData, payToken common.Address, from common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "delegatedRedeemEusdRequest", params, permit, payToken, from, receiveSide, oppositeBridge, chainId, receipt)
}
func (_Router *RouterTransactor) DelegatedRedeemEusdRequestOverGsn(opts *bind.TransactOpts, params ICurveProxyMetaRedeemEUSD, permit ICurveProxyPermitData, payToken common.Address, from common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "delegatedRedeemEusdRequest", params, permit, payToken, from, receiveSide, oppositeBridge, chainId, receipt)
}

// DelegatedRedeemEusdRequest is a paid mutator transaction binding the contract method 0x62e41ded.
//
// Solidity: function delegatedRedeemEusdRequest((address,int128,uint256,address,uint256,int128,uint256,address) params, (uint8,bytes32,bytes32,uint256,bool) permit, address payToken, address from, address receiveSide, address oppositeBridge, uint256 chainId, (uint256,uint256,uint8[2],bytes32[2],bytes32[2]) receipt) returns()
func (_Router *RouterSession) DelegatedRedeemEusdRequest(params ICurveProxyMetaRedeemEUSD, permit ICurveProxyPermitData, payToken common.Address, from common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.DelegatedRedeemEusdRequest(&_Router.TransactOpts, params, permit, payToken, from, receiveSide, oppositeBridge, chainId, receipt)
}
func (_Router *RouterSession) DelegatedRedeemEusdRequestOverGsn(params ICurveProxyMetaRedeemEUSD, permit ICurveProxyPermitData, payToken common.Address, from common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.DelegatedRedeemEusdRequestOverGsn(&_Router.TransactOpts, params, permit, payToken, from, receiveSide, oppositeBridge, chainId, receipt)
}

// DelegatedRedeemEusdRequest is a paid mutator transaction binding the contract method 0x62e41ded.
//
// Solidity: function delegatedRedeemEusdRequest((address,int128,uint256,address,uint256,int128,uint256,address) params, (uint8,bytes32,bytes32,uint256,bool) permit, address payToken, address from, address receiveSide, address oppositeBridge, uint256 chainId, (uint256,uint256,uint8[2],bytes32[2],bytes32[2]) receipt) returns()
func (_Router *RouterTransactorSession) DelegatedRedeemEusdRequest(params ICurveProxyMetaRedeemEUSD, permit ICurveProxyPermitData, payToken common.Address, from common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.DelegatedRedeemEusdRequest(&_Router.TransactOpts, params, permit, payToken, from, receiveSide, oppositeBridge, chainId, receipt)
}
func (_Router *RouterTransactorSession) DelegatedRedeemEusdRequestOverGsn(params ICurveProxyMetaRedeemEUSD, permit ICurveProxyPermitData, payToken common.Address, from common.Address, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.DelegatedRedeemEusdRequestOverGsn(&_Router.TransactOpts, params, permit, payToken, from, receiveSide, oppositeBridge, chainId, receipt)
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

// DelegatedTokenSynthesizeRequestToSolana is a paid mutator transaction binding the contract method 0xad1dbe71.
//
// Solidity: function delegatedTokenSynthesizeRequestToSolana(address token, uint256 amount, address from, bytes32[] pubkeys, bytes1 txStateBump, uint256 chainId, (uint256,uint256,uint8[2],bytes32[2],bytes32[2]) receipt) returns()
func (_Router *RouterTransactor) DelegatedTokenSynthesizeRequestToSolana(opts *bind.TransactOpts, token common.Address, amount *big.Int, from common.Address, pubkeys [][32]byte, txStateBump [1]byte, chainId *big.Int, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "delegatedTokenSynthesizeRequestToSolana", token, amount, from, pubkeys, txStateBump, chainId, receipt)
}
func (_Router *RouterTransactor) DelegatedTokenSynthesizeRequestToSolanaOverGsn(opts *bind.TransactOpts, token common.Address, amount *big.Int, from common.Address, pubkeys [][32]byte, txStateBump [1]byte, chainId *big.Int, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "delegatedTokenSynthesizeRequestToSolana", token, amount, from, pubkeys, txStateBump, chainId, receipt)
}

// DelegatedTokenSynthesizeRequestToSolana is a paid mutator transaction binding the contract method 0xad1dbe71.
//
// Solidity: function delegatedTokenSynthesizeRequestToSolana(address token, uint256 amount, address from, bytes32[] pubkeys, bytes1 txStateBump, uint256 chainId, (uint256,uint256,uint8[2],bytes32[2],bytes32[2]) receipt) returns()
func (_Router *RouterSession) DelegatedTokenSynthesizeRequestToSolana(token common.Address, amount *big.Int, from common.Address, pubkeys [][32]byte, txStateBump [1]byte, chainId *big.Int, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.DelegatedTokenSynthesizeRequestToSolana(&_Router.TransactOpts, token, amount, from, pubkeys, txStateBump, chainId, receipt)
}
func (_Router *RouterSession) DelegatedTokenSynthesizeRequestToSolanaOverGsn(token common.Address, amount *big.Int, from common.Address, pubkeys [][32]byte, txStateBump [1]byte, chainId *big.Int, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.DelegatedTokenSynthesizeRequestToSolanaOverGsn(&_Router.TransactOpts, token, amount, from, pubkeys, txStateBump, chainId, receipt)
}

// DelegatedTokenSynthesizeRequestToSolana is a paid mutator transaction binding the contract method 0xad1dbe71.
//
// Solidity: function delegatedTokenSynthesizeRequestToSolana(address token, uint256 amount, address from, bytes32[] pubkeys, bytes1 txStateBump, uint256 chainId, (uint256,uint256,uint8[2],bytes32[2],bytes32[2]) receipt) returns()
func (_Router *RouterTransactorSession) DelegatedTokenSynthesizeRequestToSolana(token common.Address, amount *big.Int, from common.Address, pubkeys [][32]byte, txStateBump [1]byte, chainId *big.Int, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.DelegatedTokenSynthesizeRequestToSolana(&_Router.TransactOpts, token, amount, from, pubkeys, txStateBump, chainId, receipt)
}
func (_Router *RouterTransactorSession) DelegatedTokenSynthesizeRequestToSolanaOverGsn(token common.Address, amount *big.Int, from common.Address, pubkeys [][32]byte, txStateBump [1]byte, chainId *big.Int, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.DelegatedTokenSynthesizeRequestToSolanaOverGsn(&_Router.TransactOpts, token, amount, from, pubkeys, txStateBump, chainId, receipt)
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

// DelegatedUnsynthesizeRequestToSolana is a paid mutator transaction binding the contract method 0xf9bee062.
//
// Solidity: function delegatedUnsynthesizeRequestToSolana(address stoken, address from, bytes32[] pubkeys, uint256 amount, uint256 chainId, (uint256,uint256,uint8[2],bytes32[2],bytes32[2]) receipt) returns()
func (_Router *RouterTransactor) DelegatedUnsynthesizeRequestToSolana(opts *bind.TransactOpts, stoken common.Address, from common.Address, pubkeys [][32]byte, amount *big.Int, chainId *big.Int, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "delegatedUnsynthesizeRequestToSolana", stoken, from, pubkeys, amount, chainId, receipt)
}
func (_Router *RouterTransactor) DelegatedUnsynthesizeRequestToSolanaOverGsn(opts *bind.TransactOpts, stoken common.Address, from common.Address, pubkeys [][32]byte, amount *big.Int, chainId *big.Int, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "delegatedUnsynthesizeRequestToSolana", stoken, from, pubkeys, amount, chainId, receipt)
}

// DelegatedUnsynthesizeRequestToSolana is a paid mutator transaction binding the contract method 0xf9bee062.
//
// Solidity: function delegatedUnsynthesizeRequestToSolana(address stoken, address from, bytes32[] pubkeys, uint256 amount, uint256 chainId, (uint256,uint256,uint8[2],bytes32[2],bytes32[2]) receipt) returns()
func (_Router *RouterSession) DelegatedUnsynthesizeRequestToSolana(stoken common.Address, from common.Address, pubkeys [][32]byte, amount *big.Int, chainId *big.Int, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.DelegatedUnsynthesizeRequestToSolana(&_Router.TransactOpts, stoken, from, pubkeys, amount, chainId, receipt)
}
func (_Router *RouterSession) DelegatedUnsynthesizeRequestToSolanaOverGsn(stoken common.Address, from common.Address, pubkeys [][32]byte, amount *big.Int, chainId *big.Int, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.DelegatedUnsynthesizeRequestToSolanaOverGsn(&_Router.TransactOpts, stoken, from, pubkeys, amount, chainId, receipt)
}

// DelegatedUnsynthesizeRequestToSolana is a paid mutator transaction binding the contract method 0xf9bee062.
//
// Solidity: function delegatedUnsynthesizeRequestToSolana(address stoken, address from, bytes32[] pubkeys, uint256 amount, uint256 chainId, (uint256,uint256,uint8[2],bytes32[2],bytes32[2]) receipt) returns()
func (_Router *RouterTransactorSession) DelegatedUnsynthesizeRequestToSolana(stoken common.Address, from common.Address, pubkeys [][32]byte, amount *big.Int, chainId *big.Int, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.DelegatedUnsynthesizeRequestToSolana(&_Router.TransactOpts, stoken, from, pubkeys, amount, chainId, receipt)
}
func (_Router *RouterTransactorSession) DelegatedUnsynthesizeRequestToSolanaOverGsn(stoken common.Address, from common.Address, pubkeys [][32]byte, amount *big.Int, chainId *big.Int, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.DelegatedUnsynthesizeRequestToSolanaOverGsn(&_Router.TransactOpts, stoken, from, pubkeys, amount, chainId, receipt)
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
