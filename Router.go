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

	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"portal\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"synthesis\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"curveProxy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userFrom\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionPrice\",\"type\":\"uint256\"}],\"name\":\"CrosschainPaymentEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_SYNTHESIZE_REQUEST_SIGNATURE_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_SYNTH_TRANSFER_REQUEST_SIGNATURE_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_UNSYNTHESIZE_REQUEST_SIGNATURE_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_trustedWorker\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"emergencyUnburnRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"emergencyUnsyntesizeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCachedChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remove\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinMintAmount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinDy\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"chain2address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"initialBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialChainID\",\"type\":\"uint256\"}],\"internalType\":\"structICurveProxy.MetaExchangeParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structICurveProxy.PermitData[]\",\"name\":\"permit\",\"type\":\"tuple[]\"},{\"internalType\":\"address[3]\",\"name\":\"token\",\"type\":\"address[3]\"},{\"internalType\":\"uint256[3]\",\"name\":\"amount\",\"type\":\"uint256[3]\"}],\"name\":\"metaExchangeRequestVia3pool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addAtCrosschainPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinMintAmountC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lpIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addAtHubPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinMintAmountH\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialChainID\",\"type\":\"uint256\"}],\"internalType\":\"structICurveProxy.MetaMintEUSD\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structICurveProxy.PermitData[]\",\"name\":\"permit\",\"type\":\"tuple[]\"},{\"internalType\":\"address[3]\",\"name\":\"token\",\"type\":\"address[3]\"},{\"internalType\":\"uint256[3]\",\"name\":\"amount\",\"type\":\"uint256[3]\"}],\"name\":\"mintEusdRequestVia3pool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"removeAtCrosschainPool\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinAmountC\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"removeAtHubPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountH\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"y\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinAmountH\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"internalType\":\"structICurveProxy.MetaRedeemEUSD\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structICurveProxy.PermitData\",\"name\":\"permit\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"redeemEusdRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"}],\"name\":\"removeTrustedWorker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"}],\"name\":\"setTrustedWorker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stoken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structISynthesis.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"}],\"name\":\"synthTransferRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenSynth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structISynthesis.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"executionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structRouter.DelegatedCallReceipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"synthTransferRequestPayNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stoken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structISynthesis.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structISynthesis.PermitData\",\"name\":\"permitData\",\"type\":\"tuple\"}],\"name\":\"synthTransferRequestWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenSynth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structISynthesis.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structISynthesis.PermitData\",\"name\":\"permitData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"executionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structRouter.DelegatedCallReceipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"synthTransferRequestWithPermitPayNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structIPortal.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"executionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structRouter.DelegatedCallReceipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"synthesizeRequestPayNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structIPortal.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structIPortal.PermitData\",\"name\":\"permitData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"executionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structRouter.DelegatedCallReceipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"synthesizeRequestWithPermitPayNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structIPortal.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"}],\"name\":\"tokenSynthesizeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"pubkeys\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes1\",\"name\":\"txStateBump\",\"type\":\"bytes1\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"tokenSynthesizeRequestToSolana\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structIPortal.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structIPortal.PermitData\",\"name\":\"permitData\",\"type\":\"tuple\"}],\"name\":\"tokenSynthesizeRequestWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stoken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structISynthesis.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"}],\"name\":\"unsynthesizeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenSynth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structISynthesis.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"executionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structRouter.DelegatedCallReceipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"unsynthesizeRequestPayNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stoken\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"pubkeys\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"unsynthesizeRequestToSolana\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stoken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structISynthesis.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structISynthesis.PermitData\",\"name\":\"permitData\",\"type\":\"tuple\"}],\"name\":\"unsynthesizeRequestWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenSynth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structISynthesis.SynthParams\",\"name\":\"synthParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"}],\"internalType\":\"structISynthesis.PermitData\",\"name\":\"permitData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"executionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structRouter.DelegatedCallReceipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"unsynthesizeRequestWithPermitPayNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x6101406040523480156200001257600080fd5b5060405162002edd38038062002edd8339810160408190526200003591620002b4565b60408051808201825260048152634559574160e01b60208083019182528351808501855260018152603160f81b908201529151902060e08190527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc661010081905260a085815284517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f8186018190528187019490945260608101929092526080808301879052308383018190528651808503909301835260c0938401968790528251929095019190912090529190915261012052600080546001600160a01b031916339081178255918291907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3506001600160a01b038416620001a35760405162461bcd60e51b815260206004820152601b60248201527f526f757465723a20706f7274616c207a65726f2061646472657373000000000060448201526064015b60405180910390fd5b6001600160a01b038316620001fb5760405162461bcd60e51b815260206004820152601e60248201527f526f757465723a2073796e746865736973207a65726f2061646472657373000060448201526064016200019a565b6001600160a01b038216620002535760405162461bcd60e51b815260206004820152601f60248201527f526f757465723a20637572766550726f7879207a65726f20616464726573730060448201526064016200019a565b50600280546001600160a01b039485166001600160a01b03199182161790915560038054938516938216939093179092556001805491909316911617905562000306565b80516001600160a01b0381168114620002af57600080fd5b919050565b60008060008060808587031215620002cb57600080fd5b620002d68562000297565b9350620002e66020860162000297565b9250620002f66040860162000297565b6060959095015193969295505050565b60805160a05160c05160e0516101005161012051612b806200035d60003960006118ba015260006118fc015260006118db0152600061186a0152600081816104d10152611c30015260006118930152612b806000f3fe6080604052600436106101d85760003560e01c80637ecebe0011610102578063d550f9e111610095578063e840f9d311610064578063e840f9d3146104f5578063f2fde38b14610535578063f804882114610555578063f8900f781461056857600080fd5b8063d550f9e11461046f578063d63f79b21461048f578063d9ab9764146104a2578063dd0e02c7146104c257600080fd5b8063a82681f7116100d1578063a82681f714610414578063c51e547c14610427578063c89efd6b1461043a578063d3c522841461045a57600080fd5b80637ecebe001461038c5780637fe04792146103ac5780638da5cb5b146103cc5780639500125b146103f457600080fd5b806344c82d231161017a57806371056f5e1161014957806371056f5e14610322578063715018a61461033757806372b90c021461034c578063781ab3821461036c57600080fd5b806344c82d23146102af57806348f1fce5146102cf5780634ec658e2146102ef57806369ab5d461461030257600080fd5b80632e88ce08116101b65780632e88ce081461023a57806333f4b9de1461025a5780633644e5151461027a578063439616681461028f57600080fd5b80630d0e5b6d146101dd57806321d7a41c146101ff5780632acbaacf14610212575b600080fd5b3480156101e957600080fd5b506101fd6101f8366004611e60565b610588565b005b6101fd61020d366004611ee9565b610607565b34801561021e57600080fd5b50610227610767565b6040519081526020015b60405180910390f35b34801561024657600080fd5b506101fd610255366004611f5b565b61078f565b34801561026657600080fd5b506101fd610275366004611fe2565b61081f565b34801561028657600080fd5b50610227610902565b34801561029b57600080fd5b506101fd6102aa366004612054565b610911565b3480156102bb57600080fd5b506101fd6102ca3660046120d1565b61094d565b3480156102db57600080fd5b506101fd6102ea366004612146565b6109a3565b6101fd6102fd366004611ee9565b610a1f565b34801561030e57600080fd5b506101fd61031d3660046121ea565b610b39565b34801561032e57600080fd5b50610227610c01565b34801561034357600080fd5b506101fd610c10565b34801561035857600080fd5b506101fd61036736600461225b565b610c8d565b34801561037857600080fd5b506101fd610387366004611fe2565b610cd8565b34801561039857600080fd5b506102276103a736600461225b565b610e06565b3480156103b857600080fd5b506101fd6103c736600461225b565b610e26565b3480156103d857600080fd5b506000546040516001600160a01b039091168152602001610231565b34801561040057600080fd5b506101fd61040f366004612054565b610e74565b6101fd610422366004611ee9565b610eb0565b6101fd610435366004611fe2565b610ffc565b34801561044657600080fd5b506101fd610455366004612146565b6110b2565b34801561046657600080fd5b50610227611104565b34801561047b57600080fd5b506101fd61048a366004612146565b611113565b6101fd61049d366004611fe2565b6111a1565b3480156104ae57600080fd5b506101fd6104bd366004611fe2565b611219565b3480156104ce57600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610227565b34801561050157600080fd5b5061052561051036600461225b565b60046020526000908152604090205460ff1681565b6040519015158152602001610231565b34801561054157600080fd5b506101fd61055036600461225b565b61130a565b6101fd610563366004611fe2565b6113f4565b34801561057457600080fd5b506101fd610583366004612276565b61145e565b610594853330856114e9565b60035460405163966e0fe960e01b81526001600160a01b039091169063966e0fe9906105ce90889033908990899089908990600401612324565b600060405180830381600087803b1580156105e857600080fd5b505af11580156105fc573d6000803e3d6000fd5b505050505050505050565b600061063c846040013560405160200161062090612366565b6040516020818303038152906040528051906020012084611549565b905061064982358261172c565b6001600160a01b03871663d505accf333061066a60a08801608089016123de565b6106745789610678565b6000195b606088013561068a60208a018a6123fb565b89602001358a604001356040518863ffffffff1660e01b81526004016106b69796959493929190612416565b600060405180830381600087803b1580156106d057600080fd5b505af11580156106e4573d6000803e3d6000fd5b505050506106f4873330896114e9565b600354604051633bb2a45160e11b81526001600160a01b039091169063776548a29061072c908a908a9033908b908b9060040161248b565b600060405180830381600087803b15801561074657600080fd5b505af115801561075a573d6000803e3d6000fd5b5050505050505050505050565b604051602001610776906124c9565b6040516020818303038152906040528051906020012081565b6002546107a990879033906001600160a01b0316886114e9565b600254604051632b17790f60e21b81526001600160a01b039091169063ac5de43c906107e5908990899033908a908a908a908a9060040161252f565b600060405180830381600087803b1580156107ff57600080fd5b505af1158015610813573d6000803e3d6000fd5b50505050505050505050565b6001600160a01b03851663d505accf333061084060a08601608087016123de565b61084a578761084e565b6000195b606086013561086060208801886123fb565b876020013588604001356040518863ffffffff1660e01b815260040161088c9796959493929190612416565b600060405180830381600087803b1580156108a657600080fd5b505af11580156108ba573d6000803e3d6000fd5b505050506108ca853330876114e9565b600354604051633bb2a45160e11b81526001600160a01b039091169063776548a2906105ce908890889033908990899060040161248b565b600061090c61185d565b905090565b600254604051630872c2cd60e31b81526001600160a01b039091169063439616689061072c908a908a908a908a908a908a908a90600401612583565b60015461096b90859033906001600160a01b031660808a01356114e9565b600154604051635fcc4a3f60e01b81526001600160a01b0390911690635fcc4a3f906107e5908990899088908890889060040161261b565b6109af843330866114e9565b600354604051633bb2a45160e11b81526001600160a01b039091169063776548a2906109e7908790879033908890889060040161248b565b600060405180830381600087803b158015610a0157600080fd5b505af1158015610a15573d6000803e3d6000fd5b5050505050505050565b6000610a388460400135604051602001610620906126f0565b90506001600160a01b03871663d505accf3330610a5b60a08801608089016123de565b610a655789610a69565b6000195b6060880135610a7b60208a018a6123fb565b89602001358a604001356040518863ffffffff1660e01b8152600401610aa79796959493929190612416565b600060405180830381600087803b158015610ac157600080fd5b505af1158015610ad5573d6000803e3d6000fd5b50505050610ae782600001358261172c565b600254610b0190889033906001600160a01b0316896114e9565b600254604051637aa8486d60e11b81526001600160a01b039091169063f55090da9061072c908a908a9033908b908b9060040161248b565b60005b6003811015610bc8576000828260038110610b5957610b59612754565b60200201351115610bb657610bb6838260038110610b7957610b79612754565b602002016020810190610b8c919061225b565b60025433906001600160a01b0316858560038110610bac57610bac612754565b60200201356114e9565b80610bc08161276a565b915050610b3c565b506001546040516301c2bc7d60e51b81526001600160a01b03909116906338578fa0906105ce9088908890889088908890600401612815565b604051602001610776906126f0565b6000546001600160a01b03163314610c435760405162461bcd60e51b8152600401610c3a906128cd565b60405180910390fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b03163314610cb75760405162461bcd60e51b8152600401610c3a906128cd565b6001600160a01b03166000908152600460205260409020805460ff19169055565b6001600160a01b03851663d505accf3330610cf960a08601608087016123de565b610d035787610d07565b6000195b6060860135610d1960208801886123fb565b876020013588604001356040518863ffffffff1660e01b8152600401610d459796959493929190612416565b600060405180830381600087803b158015610d5f57600080fd5b505af1158015610d73573d6000803e3d6000fd5b50505050610d83853330876114e9565b60035460405163451c763560e01b81526001600160a01b039091169063451c763590610dbb908890889033908990899060040161248b565b6020604051808303816000875af1158015610dda573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dfe9190612902565b505050505050565b6001600160a01b0381166000908152600560205260408120545b92915050565b6000546001600160a01b03163314610e505760405162461bcd60e51b8152600401610c3a906128cd565b6001600160a01b03166000908152600460205260409020805460ff19166001179055565b600354604051639500125b60e01b81526001600160a01b0390911690639500125b9061072c908a908a908a908a908a908a908a90600401612583565b6000610ec98460400135604051602001610620906124c9565b9050610ed682358261172c565b6001600160a01b03871663d505accf3330610ef760a08801608089016123de565b610f015789610f05565b6000195b6060880135610f1760208a018a6123fb565b89602001358a604001356040518863ffffffff1660e01b8152600401610f439796959493929190612416565b600060405180830381600087803b158015610f5d57600080fd5b505af1158015610f71573d6000803e3d6000fd5b50505050610f81873330896114e9565b60035460405163451c763560e01b81526001600160a01b039091169063451c763590610fb9908a908a9033908b908b9060040161248b565b6020604051808303816000875af1158015610fd8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a159190612902565b60006110158360400135604051602001610620906124c9565b905061102282358261172c565b61102e863330886114e9565b60035460405163451c763560e01b81526001600160a01b039091169063451c763590611066908990899033908a908a9060040161248b565b6020604051808303816000875af1158015611085573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110a99190612902565b50505050505050565b6002546110cc90859033906001600160a01b0316866114e9565b600254604051637aa8486d60e11b81526001600160a01b039091169063f55090da906109e7908790879033908890889060040161248b565b60405160200161077690612366565b61111f843330866114e9565b60035460405163451c763560e01b81526001600160a01b039091169063451c763590611157908790879033908890889060040161248b565b6020604051808303816000875af1158015611176573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061119a9190612902565b5050505050565b60006111ba8360400135604051602001610620906126f0565b90506111c782358261172c565b6002546111e190879033906001600160a01b0316886114e9565b600254604051637aa8486d60e11b81526001600160a01b039091169063f55090da906107e5908990899033908a908a9060040161248b565b6001600160a01b03851663d505accf333061123a60a08601608087016123de565b6112445787611248565b6000195b606086013561125a60208801886123fb565b876020013588604001356040518863ffffffff1660e01b81526004016112869796959493929190612416565b600060405180830381600087803b1580156112a057600080fd5b505af11580156112b4573d6000803e3d6000fd5b50506002546112d2925087915033906001600160a01b0316876114e9565b600254604051637aa8486d60e11b81526001600160a01b039091169063f55090da906105ce908890889033908990899060040161248b565b6000546001600160a01b031633146113345760405162461bcd60e51b8152600401610c3a906128cd565b6001600160a01b0381166113995760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610c3a565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b600061140d836040013560405160200161062090612366565b905061141a82358261172c565b611426863330886114e9565b600354604051633bb2a45160e11b81526001600160a01b039091169063776548a2906107e5908990899033908a908a9060040161248b565b60005b60038110156114b057600082826003811061147e5761147e612754565b6020020135111561149e5761149e838260038110610b7957610b79612754565b806114a88161276a565b915050611461565b50600154604051634638ef3f60e11b81526001600160a01b0390911690638c71de7e906105ce908890889088908890889060040161291b565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b179052611543908590611920565b50505050565b600080611555336119f7565b604080517fb2cd5b2d9baed0ca8e8805aae5275d69b6a747b2fdb5c114eb96a962fdc1a6996020808301919091523360601b6bffffffffffffffffffffffff19168284015260548201899052863560748301526094820188905260b482018490528681013560d4808401919091528351808403909101815260f49092019092528051910120909150600061163e6115eb83611a1f565b6040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b90506116638161165460608801604089016123fb565b87606001358860800135611a6d565b935084602001354211156116ac5760405162461bcd60e51b815260206004820152601060248201526f526f757465723a20646561646c696e6560801b6044820152606401610c3a565b6001600160a01b03841660009081526004602052604090205460ff166117225760405162461bcd60e51b815260206004820152602560248201527f526f757465723a20696e76616c6964207369676e61747572652066726f6d207760448201526437b935b2b960d91b6064820152608401610c3a565b5050509392505050565b813410156117755760405162461bcd60e51b8152602060048201526016602482015275149bdd5d195c8e881a5b9d985b1a5908185b5bdd5b9d60521b6044820152606401610c3a565b6000816001600160a01b03163460405160006040518083038185875af1925050503d80600081146117c2576040519150601f19603f3d011682016040523d82523d6000602084013e6117c7565b606091505b50509050806118185760405162461bcd60e51b815260206004820152601c60248201527f526f757465723a206661696c656420746f2073656e64204574686572000000006044820152606401610c3a565b6040518381526001600160a01b0383169033907f141cdd7ab02a0ea7399fa91a81781d764708704497d78344275c700a6caa0be69060200160405180910390a3505050565b6000306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614156118b557507f000000000000000000000000000000000000000000000000000000000000000090565b61090c7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000611c16565b6000611975826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316611c809092919063ffffffff16565b8051909150156119f257808060200190518101906119939190612ab2565b6119f25760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610c3a565b505050565b6001600160a01b03811660009081526005602052604090208054600181018255905b50919050565b6000610e20611a2c61185d565b8360405161190160f01b6020820152602281018390526042810182905260009060620160405160208183030381529060405280519060200120905092915050565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0821115611aea5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610c3a565b8360ff16601b1480611aff57508360ff16601c145b611b565760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610c3a565b6040805160008082526020820180845288905260ff871692820192909252606081018590526080810184905260019060a0016020604051602081039080840390855afa158015611baa573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116611c0d5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610c3a565b95945050505050565b6040805160208101859052908101839052606081018290527f000000000000000000000000000000000000000000000000000000000000000060808201523060a082015260009060c0016040516020818303038152906040528051906020012090505b9392505050565b6060611c8f8484600085611c97565b949350505050565b606082471015611cf85760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610c3a565b843b611d465760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610c3a565b600080866001600160a01b03168587604051611d629190612afb565b60006040518083038185875af1925050503d8060008114611d9f576040519150601f19603f3d011682016040523d82523d6000602084013e611da4565b606091505b5091509150611db4828286611dbf565b979650505050505050565b60608315611dce575081611c79565b825115611dde5782518084602001fd5b8160405162461bcd60e51b8152600401610c3a9190612b17565b80356001600160a01b0381168114611e0f57600080fd5b919050565b60008083601f840112611e2657600080fd5b50813567ffffffffffffffff811115611e3e57600080fd5b6020830191508360208260051b8501011115611e5957600080fd5b9250929050565b600080600080600060808688031215611e7857600080fd5b611e8186611df8565b9450602086013567ffffffffffffffff811115611e9d57600080fd5b611ea988828901611e14565b9699909850959660408101359660609091013595509350505050565b600060608284031215611a1957600080fd5b600060a08284031215611a1957600080fd5b6000806000806000806102008789031215611f0357600080fd5b611f0c87611df8565b955060208701359450611f2160408801611df8565b9350611f308860608901611ec5565b9250611f3f8860c08901611ed7565b9150611f4f886101608901611ed7565b90509295509295509295565b60008060008060008060a08789031215611f7457600080fd5b611f7d87611df8565b955060208701359450604087013567ffffffffffffffff811115611fa057600080fd5b611fac89828a01611e14565b90955093505060608701356001600160f81b031981168114611fcd57600080fd5b80925050608087013590509295509295509295565b60008060008060006101608688031215611ffb57600080fd5b61200486611df8565b94506020860135935061201960408701611df8565b92506120288760608801611ec5565b91506120378760c08801611ed7565b90509295509295909350565b803560ff81168114611e0f57600080fd5b600080600080600080600060e0888a03121561206f57600080fd5b8735965061207f60208901611df8565b955061208d60408901611df8565b9450606088013593506120a260808901612043565b925060a0880135915060c0880135905092959891949750929550565b60006101008284031215611a1957600080fd5b60008060008060008061022087890312156120eb57600080fd5b6120f588886120be565b9550612105886101008901611ed7565b94506121146101a08801611df8565b93506121236101c08801611df8565b92506121326101e08801611df8565b915061020087013590509295509295509295565b60008060008060c0858703121561215c57600080fd5b61216585611df8565b93506020850135925061217a60408601611df8565b91506121898660608701611ec5565b905092959194509250565b60008083601f8401126121a657600080fd5b50813567ffffffffffffffff8111156121be57600080fd5b60208301915083602060a083028501011115611e5957600080fd5b8060608101831015610e2057600080fd5b60008060008060006101e0868803121561220357600080fd5b61220d87876120be565b945061010086013567ffffffffffffffff81111561222a57600080fd5b61223688828901612194565b909550935061224b90508761012088016121d9565b91506120378761018088016121d9565b60006020828403121561226d57600080fd5b611c7982611df8565b60008060008060008587036102e081121561229057600080fd5b610200808212156122a057600080fd5b879650860135905067ffffffffffffffff8111156122bd57600080fd5b6122c988828901612194565b90955093506122de90508761022088016121d9565b91506120378761028088016121d9565b81835260006001600160fb1b0383111561230757600080fd5b8260051b8083602087013760009401602001938452509192915050565b6001600160a01b0387811682528616602082015260a06040820181905260009061235190830186886122ee565b60608301949094525060800152949350505050565b7f73796e74685472616e736665725265717565737428616464726573732c75696e81527f743235362c616464726573732c616464726573732c5b616464726573732c616460208201526e64726573732c75696e743235365d2960881b6040820152604f0190565b80151581146123db57600080fd5b50565b6000602082840312156123f057600080fd5b8135611c79816123cd565b60006020828403121561240d57600080fd5b611c7982612043565b6001600160a01b0397881681529590961660208601526040850193909352606084019190915260ff16608083015260a082015260c081019190915260e00190565b6001600160a01b038061246983611df8565b1683528061247960208401611df8565b16602084015250604090810135910152565b6001600160a01b0386811682526020820186905284811660408301528316606082015260e081016124bf6080830184612457565b9695505050505050565b7f756e73796e74686573697a655265717565737428616464726573732c75696e7481527f3235362c616464726573732c616464726573732c5b616464726573732c61646460208201526d726573732c75696e743235365d2960901b6040820152604e0190565b6001600160a01b038881168252602082018890528616604082015260c06060820181905260009061256390830186886122ee565b6001600160f81b03199490941660808301525060a0015295945050505050565b9687526001600160a01b039586166020880152939094166040860152606085019190915260ff16608084015260a083019190915260c082015260e00190565b8035600f81900b8114611e0f57600080fd5b60ff6125df82612043565b168252602081013560208301526040810135604083015260608101356060830152608081013561260e816123cd565b8015156080840152505050565b61020081016001600160a01b038061263289611df8565b168352612641602089016125c2565b600f0b6020840152604088013560408401528061266060608a01611df8565b166060840152506080870135608083015261267d60a088016125c2565b61268c60a0840182600f0b9052565b5060c087013560c08301526126a360e08801611df8565b6001600160a01b031660e08301526126bf6101008301876125d4565b6001600160a01b0385166101a08301526001600160a01b0384166101c0830152826101e08301529695505050505050565b7f73796e74686573697a655265717565737428616464726573732c75696e74323581527f362c616464726573732c616464726573732c5b616464726573732c616464726560208201526b73732c75696e743235365d2960a01b6040820152604c0190565b634e487b7160e01b600052603260045260246000fd5b600060001982141561278c57634e487b7160e01b600052601160045260246000fd5b5060010190565b81835260208301925060008160005b848110156127c7576127b486836125d4565b60a09586019591909101906001016127a2565b5093949350505050565b8060005b6003811015611543576001600160a01b036127ef83611df8565b16845260209384019391909101906001016127d5565b6060818337506000606082015250565b60006101e06001600160a01b038061282c8a611df8565b16845260208901356020850152604089013560408501528061285060608b01611df8565b166060850152608089013560808501528061286d60a08b01611df8565b1660a08501525061288060c08901611df8565b6001600160a01b03811660c08501525060e088013560e0840152806101008401526128ae8184018789612793565b9150506128bf6101208301856127d1565b6124bf610180830184612805565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b60006020828403121561291457600080fd5b5051919050565b60006102e061293a8361292d8a611df8565b6001600160a01b03169052565b61294660208901611df8565b6001600160a01b0316602084015261296060408901611df8565b6001600160a01b0316604084015260608881013590840152612984608089016125c2565b6129936080850182600f0b9052565b506129a060a089016125c2565b6129af60a0850182600f0b9052565b5060c088013560c08401526129c660e089016125c2565b6129d560e0850182600f0b9052565b5061010088810135908401526101206129ef818a01611df8565b6001600160a01b031690840152610140612a0a898201611df8565b6001600160a01b031690840152610160612a25898201611df8565b6001600160a01b031690840152610180612a40898201611df8565b6001600160a01b0316908401526101a088810135908401526101c0612a66818a01611df8565b6001600160a01b0316908401526101e088810135908401526102008301819052612a938184018789612793565b915050612aa46102208301856127d1565b6124bf610280830184612805565b600060208284031215612ac457600080fd5b8151611c79816123cd565b60005b83811015612aea578181015183820152602001612ad2565b838111156115435750506000910152565b60008251612b0d818460208701612acf565b9190910192915050565b6020815260008251806020840152612b36816040850160208701612acf565b601f01601f1916919091016040019291505056fea2646970667358221220fbc0e23eee95ac0afa746c7fdfc255c80cdf62bb1e89c3b8bef7893c92e5c2f064736f6c634300080a0033",
}

// RouterABI is the input ABI used to generate the binding from.
// Deprecated: Use RouterMetaData.ABI instead.
var RouterABI = RouterMetaData.ABI

// RouterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RouterMetaData.Bin instead.
var RouterBin = RouterMetaData.Bin

// DeployRouter deploys a new Ethereum contract, binding an instance of Router to it.
func DeployRouter(auth *bind.TransactOpts, backend bind.ContractBackend, portal common.Address, synthesis common.Address, curveProxy common.Address, chainID *big.Int) (common.Address, *types.Transaction, *Router, error) {
	parsed, err := RouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RouterBin), backend, portal, synthesis, curveProxy, chainID)
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

// GetCachedChainId is a free data retrieval call binding the contract method 0xdd0e02c7.
//
// Solidity: function getCachedChainId() view returns(uint256)
func (_Router *RouterCaller) GetCachedChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getCachedChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCachedChainId is a free data retrieval call binding the contract method 0xdd0e02c7.
//
// Solidity: function getCachedChainId() view returns(uint256)
func (_Router *RouterSession) GetCachedChainId() (*big.Int, error) {
	return _Router.Contract.GetCachedChainId(&_Router.CallOpts)
}

// GetCachedChainId is a free data retrieval call binding the contract method 0xdd0e02c7.
//
// Solidity: function getCachedChainId() view returns(uint256)
func (_Router *RouterCallerSession) GetCachedChainId() (*big.Int, error) {
	return _Router.Contract.GetCachedChainId(&_Router.CallOpts)
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

// SynthTransferRequest is a paid mutator transaction binding the contract method 0x48f1fce5.
//
// Solidity: function synthTransferRequest(address stoken, uint256 amount, address to, (address,address,uint256) synthParams) returns()
func (_Router *RouterTransactor) SynthTransferRequest(opts *bind.TransactOpts, stoken common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "synthTransferRequest", stoken, amount, to, synthParams)
}
func (_Router *RouterTransactor) SynthTransferRequestOverGsn(opts *bind.TransactOpts, stoken common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "synthTransferRequest", stoken, amount, to, synthParams)
}

// SynthTransferRequest is a paid mutator transaction binding the contract method 0x48f1fce5.
//
// Solidity: function synthTransferRequest(address stoken, uint256 amount, address to, (address,address,uint256) synthParams) returns()
func (_Router *RouterSession) SynthTransferRequest(stoken common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams) (*types.Transaction, error) {
	return _Router.Contract.SynthTransferRequest(&_Router.TransactOpts, stoken, amount, to, synthParams)
}
func (_Router *RouterSession) SynthTransferRequestOverGsn(stoken common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams) (common.Hash, error) {
	return _Router.Contract.SynthTransferRequestOverGsn(&_Router.TransactOpts, stoken, amount, to, synthParams)
}

// SynthTransferRequest is a paid mutator transaction binding the contract method 0x48f1fce5.
//
// Solidity: function synthTransferRequest(address stoken, uint256 amount, address to, (address,address,uint256) synthParams) returns()
func (_Router *RouterTransactorSession) SynthTransferRequest(stoken common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams) (*types.Transaction, error) {
	return _Router.Contract.SynthTransferRequest(&_Router.TransactOpts, stoken, amount, to, synthParams)
}
func (_Router *RouterTransactorSession) SynthTransferRequestOverGsn(stoken common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams) (common.Hash, error) {
	return _Router.Contract.SynthTransferRequestOverGsn(&_Router.TransactOpts, stoken, amount, to, synthParams)
}

// SynthTransferRequestPayNative is a paid mutator transaction binding the contract method 0xf8048821.
//
// Solidity: function synthTransferRequestPayNative(address tokenSynth, uint256 amount, address to, (address,address,uint256) synthParams, (uint256,uint256,uint8,bytes32,bytes32) receipt) payable returns()
func (_Router *RouterTransactor) SynthTransferRequestPayNative(opts *bind.TransactOpts, tokenSynth common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "synthTransferRequestPayNative", tokenSynth, amount, to, synthParams, receipt)
}
func (_Router *RouterTransactor) SynthTransferRequestPayNativeOverGsn(opts *bind.TransactOpts, tokenSynth common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "synthTransferRequestPayNative", tokenSynth, amount, to, synthParams, receipt)
}

// SynthTransferRequestPayNative is a paid mutator transaction binding the contract method 0xf8048821.
//
// Solidity: function synthTransferRequestPayNative(address tokenSynth, uint256 amount, address to, (address,address,uint256) synthParams, (uint256,uint256,uint8,bytes32,bytes32) receipt) payable returns()
func (_Router *RouterSession) SynthTransferRequestPayNative(tokenSynth common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.SynthTransferRequestPayNative(&_Router.TransactOpts, tokenSynth, amount, to, synthParams, receipt)
}
func (_Router *RouterSession) SynthTransferRequestPayNativeOverGsn(tokenSynth common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.SynthTransferRequestPayNativeOverGsn(&_Router.TransactOpts, tokenSynth, amount, to, synthParams, receipt)
}

// SynthTransferRequestPayNative is a paid mutator transaction binding the contract method 0xf8048821.
//
// Solidity: function synthTransferRequestPayNative(address tokenSynth, uint256 amount, address to, (address,address,uint256) synthParams, (uint256,uint256,uint8,bytes32,bytes32) receipt) payable returns()
func (_Router *RouterTransactorSession) SynthTransferRequestPayNative(tokenSynth common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.SynthTransferRequestPayNative(&_Router.TransactOpts, tokenSynth, amount, to, synthParams, receipt)
}
func (_Router *RouterTransactorSession) SynthTransferRequestPayNativeOverGsn(tokenSynth common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.SynthTransferRequestPayNativeOverGsn(&_Router.TransactOpts, tokenSynth, amount, to, synthParams, receipt)
}

// SynthTransferRequestWithPermit is a paid mutator transaction binding the contract method 0x33f4b9de.
//
// Solidity: function synthTransferRequestWithPermit(address stoken, uint256 amount, address to, (address,address,uint256) synthParams, (uint8,bytes32,bytes32,uint256,bool) permitData) returns()
func (_Router *RouterTransactor) SynthTransferRequestWithPermit(opts *bind.TransactOpts, stoken common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, permitData ISynthesisPermitData) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "synthTransferRequestWithPermit", stoken, amount, to, synthParams, permitData)
}
func (_Router *RouterTransactor) SynthTransferRequestWithPermitOverGsn(opts *bind.TransactOpts, stoken common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, permitData ISynthesisPermitData) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "synthTransferRequestWithPermit", stoken, amount, to, synthParams, permitData)
}

// SynthTransferRequestWithPermit is a paid mutator transaction binding the contract method 0x33f4b9de.
//
// Solidity: function synthTransferRequestWithPermit(address stoken, uint256 amount, address to, (address,address,uint256) synthParams, (uint8,bytes32,bytes32,uint256,bool) permitData) returns()
func (_Router *RouterSession) SynthTransferRequestWithPermit(stoken common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, permitData ISynthesisPermitData) (*types.Transaction, error) {
	return _Router.Contract.SynthTransferRequestWithPermit(&_Router.TransactOpts, stoken, amount, to, synthParams, permitData)
}
func (_Router *RouterSession) SynthTransferRequestWithPermitOverGsn(stoken common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, permitData ISynthesisPermitData) (common.Hash, error) {
	return _Router.Contract.SynthTransferRequestWithPermitOverGsn(&_Router.TransactOpts, stoken, amount, to, synthParams, permitData)
}

// SynthTransferRequestWithPermit is a paid mutator transaction binding the contract method 0x33f4b9de.
//
// Solidity: function synthTransferRequestWithPermit(address stoken, uint256 amount, address to, (address,address,uint256) synthParams, (uint8,bytes32,bytes32,uint256,bool) permitData) returns()
func (_Router *RouterTransactorSession) SynthTransferRequestWithPermit(stoken common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, permitData ISynthesisPermitData) (*types.Transaction, error) {
	return _Router.Contract.SynthTransferRequestWithPermit(&_Router.TransactOpts, stoken, amount, to, synthParams, permitData)
}
func (_Router *RouterTransactorSession) SynthTransferRequestWithPermitOverGsn(stoken common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, permitData ISynthesisPermitData) (common.Hash, error) {
	return _Router.Contract.SynthTransferRequestWithPermitOverGsn(&_Router.TransactOpts, stoken, amount, to, synthParams, permitData)
}

// SynthTransferRequestWithPermitPayNative is a paid mutator transaction binding the contract method 0x21d7a41c.
//
// Solidity: function synthTransferRequestWithPermitPayNative(address tokenSynth, uint256 amount, address to, (address,address,uint256) synthParams, (uint8,bytes32,bytes32,uint256,bool) permitData, (uint256,uint256,uint8,bytes32,bytes32) receipt) payable returns()
func (_Router *RouterTransactor) SynthTransferRequestWithPermitPayNative(opts *bind.TransactOpts, tokenSynth common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, permitData ISynthesisPermitData, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "synthTransferRequestWithPermitPayNative", tokenSynth, amount, to, synthParams, permitData, receipt)
}
func (_Router *RouterTransactor) SynthTransferRequestWithPermitPayNativeOverGsn(opts *bind.TransactOpts, tokenSynth common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, permitData ISynthesisPermitData, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "synthTransferRequestWithPermitPayNative", tokenSynth, amount, to, synthParams, permitData, receipt)
}

// SynthTransferRequestWithPermitPayNative is a paid mutator transaction binding the contract method 0x21d7a41c.
//
// Solidity: function synthTransferRequestWithPermitPayNative(address tokenSynth, uint256 amount, address to, (address,address,uint256) synthParams, (uint8,bytes32,bytes32,uint256,bool) permitData, (uint256,uint256,uint8,bytes32,bytes32) receipt) payable returns()
func (_Router *RouterSession) SynthTransferRequestWithPermitPayNative(tokenSynth common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, permitData ISynthesisPermitData, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.SynthTransferRequestWithPermitPayNative(&_Router.TransactOpts, tokenSynth, amount, to, synthParams, permitData, receipt)
}
func (_Router *RouterSession) SynthTransferRequestWithPermitPayNativeOverGsn(tokenSynth common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, permitData ISynthesisPermitData, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.SynthTransferRequestWithPermitPayNativeOverGsn(&_Router.TransactOpts, tokenSynth, amount, to, synthParams, permitData, receipt)
}

// SynthTransferRequestWithPermitPayNative is a paid mutator transaction binding the contract method 0x21d7a41c.
//
// Solidity: function synthTransferRequestWithPermitPayNative(address tokenSynth, uint256 amount, address to, (address,address,uint256) synthParams, (uint8,bytes32,bytes32,uint256,bool) permitData, (uint256,uint256,uint8,bytes32,bytes32) receipt) payable returns()
func (_Router *RouterTransactorSession) SynthTransferRequestWithPermitPayNative(tokenSynth common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, permitData ISynthesisPermitData, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.SynthTransferRequestWithPermitPayNative(&_Router.TransactOpts, tokenSynth, amount, to, synthParams, permitData, receipt)
}
func (_Router *RouterTransactorSession) SynthTransferRequestWithPermitPayNativeOverGsn(tokenSynth common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, permitData ISynthesisPermitData, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.SynthTransferRequestWithPermitPayNativeOverGsn(&_Router.TransactOpts, tokenSynth, amount, to, synthParams, permitData, receipt)
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

// UnsynthesizeRequestWithPermit is a paid mutator transaction binding the contract method 0x781ab382.
//
// Solidity: function unsynthesizeRequestWithPermit(address stoken, uint256 amount, address to, (address,address,uint256) synthParams, (uint8,bytes32,bytes32,uint256,bool) permitData) returns()
func (_Router *RouterTransactor) UnsynthesizeRequestWithPermit(opts *bind.TransactOpts, stoken common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, permitData ISynthesisPermitData) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "unsynthesizeRequestWithPermit", stoken, amount, to, synthParams, permitData)
}
func (_Router *RouterTransactor) UnsynthesizeRequestWithPermitOverGsn(opts *bind.TransactOpts, stoken common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, permitData ISynthesisPermitData) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "unsynthesizeRequestWithPermit", stoken, amount, to, synthParams, permitData)
}

// UnsynthesizeRequestWithPermit is a paid mutator transaction binding the contract method 0x781ab382.
//
// Solidity: function unsynthesizeRequestWithPermit(address stoken, uint256 amount, address to, (address,address,uint256) synthParams, (uint8,bytes32,bytes32,uint256,bool) permitData) returns()
func (_Router *RouterSession) UnsynthesizeRequestWithPermit(stoken common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, permitData ISynthesisPermitData) (*types.Transaction, error) {
	return _Router.Contract.UnsynthesizeRequestWithPermit(&_Router.TransactOpts, stoken, amount, to, synthParams, permitData)
}
func (_Router *RouterSession) UnsynthesizeRequestWithPermitOverGsn(stoken common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, permitData ISynthesisPermitData) (common.Hash, error) {
	return _Router.Contract.UnsynthesizeRequestWithPermitOverGsn(&_Router.TransactOpts, stoken, amount, to, synthParams, permitData)
}

// UnsynthesizeRequestWithPermit is a paid mutator transaction binding the contract method 0x781ab382.
//
// Solidity: function unsynthesizeRequestWithPermit(address stoken, uint256 amount, address to, (address,address,uint256) synthParams, (uint8,bytes32,bytes32,uint256,bool) permitData) returns()
func (_Router *RouterTransactorSession) UnsynthesizeRequestWithPermit(stoken common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, permitData ISynthesisPermitData) (*types.Transaction, error) {
	return _Router.Contract.UnsynthesizeRequestWithPermit(&_Router.TransactOpts, stoken, amount, to, synthParams, permitData)
}
func (_Router *RouterTransactorSession) UnsynthesizeRequestWithPermitOverGsn(stoken common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, permitData ISynthesisPermitData) (common.Hash, error) {
	return _Router.Contract.UnsynthesizeRequestWithPermitOverGsn(&_Router.TransactOpts, stoken, amount, to, synthParams, permitData)
}

// UnsynthesizeRequestWithPermitPayNative is a paid mutator transaction binding the contract method 0xa82681f7.
//
// Solidity: function unsynthesizeRequestWithPermitPayNative(address tokenSynth, uint256 amount, address to, (address,address,uint256) synthParams, (uint8,bytes32,bytes32,uint256,bool) permitData, (uint256,uint256,uint8,bytes32,bytes32) receipt) payable returns()
func (_Router *RouterTransactor) UnsynthesizeRequestWithPermitPayNative(opts *bind.TransactOpts, tokenSynth common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, permitData ISynthesisPermitData, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "unsynthesizeRequestWithPermitPayNative", tokenSynth, amount, to, synthParams, permitData, receipt)
}
func (_Router *RouterTransactor) UnsynthesizeRequestWithPermitPayNativeOverGsn(opts *bind.TransactOpts, tokenSynth common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, permitData ISynthesisPermitData, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return GsnExecutor(_Router.gsn, RouterMetaData.ABI, "unsynthesizeRequestWithPermitPayNative", tokenSynth, amount, to, synthParams, permitData, receipt)
}

// UnsynthesizeRequestWithPermitPayNative is a paid mutator transaction binding the contract method 0xa82681f7.
//
// Solidity: function unsynthesizeRequestWithPermitPayNative(address tokenSynth, uint256 amount, address to, (address,address,uint256) synthParams, (uint8,bytes32,bytes32,uint256,bool) permitData, (uint256,uint256,uint8,bytes32,bytes32) receipt) payable returns()
func (_Router *RouterSession) UnsynthesizeRequestWithPermitPayNative(tokenSynth common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, permitData ISynthesisPermitData, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.UnsynthesizeRequestWithPermitPayNative(&_Router.TransactOpts, tokenSynth, amount, to, synthParams, permitData, receipt)
}
func (_Router *RouterSession) UnsynthesizeRequestWithPermitPayNativeOverGsn(tokenSynth common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, permitData ISynthesisPermitData, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.UnsynthesizeRequestWithPermitPayNativeOverGsn(&_Router.TransactOpts, tokenSynth, amount, to, synthParams, permitData, receipt)
}

// UnsynthesizeRequestWithPermitPayNative is a paid mutator transaction binding the contract method 0xa82681f7.
//
// Solidity: function unsynthesizeRequestWithPermitPayNative(address tokenSynth, uint256 amount, address to, (address,address,uint256) synthParams, (uint8,bytes32,bytes32,uint256,bool) permitData, (uint256,uint256,uint8,bytes32,bytes32) receipt) payable returns()
func (_Router *RouterTransactorSession) UnsynthesizeRequestWithPermitPayNative(tokenSynth common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, permitData ISynthesisPermitData, receipt RouterDelegatedCallReceipt) (*types.Transaction, error) {
	return _Router.Contract.UnsynthesizeRequestWithPermitPayNative(&_Router.TransactOpts, tokenSynth, amount, to, synthParams, permitData, receipt)
}
func (_Router *RouterTransactorSession) UnsynthesizeRequestWithPermitPayNativeOverGsn(tokenSynth common.Address, amount *big.Int, to common.Address, synthParams ISynthesisSynthParams, permitData ISynthesisPermitData, receipt RouterDelegatedCallReceipt) (common.Hash, error) {
	return _Router.Contract.UnsynthesizeRequestWithPermitPayNativeOverGsn(&_Router.TransactOpts, tokenSynth, amount, to, synthParams, permitData, receipt)
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
