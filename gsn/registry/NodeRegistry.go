// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registry

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/eywa-protocol/wrappers"
	"github.com/eywa-protocol/wrappers/gsn"
)

// NodeRegistryNode is an auto generated low-level Go binding around an user-defined struct.
type NodeRegistryNode struct {
	Owner     common.Address
	Pool      common.Address
	HostId    string
	BlsPubKey string
	NodeId    *big.Int
}

var __contractNodeRegistrySourceABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_EYWA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_forwarder\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayerPool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hostId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"blsPubKey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nodeId\",\"type\":\"uint256\"}],\"name\":\"CreatedRelayer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"snapNum\",\"type\":\"uint256\"}],\"name\":\"NewSnapshot\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EYWA\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_COLLATERAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"checkPermissionTrustList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"hostId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"blsPubKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nodeId\",\"type\":\"uint256\"}],\"internalType\":\"structNodeRegistry.Node\",\"name\":\"_node\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"createRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBLSPubKeys\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getNode\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"hostId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"blsPubKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nodeId\",\"type\":\"uint256\"}],\"internalType\":\"structNodeRegistry.Node\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodes\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"hostId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"blsPubKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nodeId\",\"type\":\"uint256\"}],\"internalType\":\"structNodeRegistry.Node[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSnapshot\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"hostIds\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nodeExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ownedNodes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"hostId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"blsPubKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nodeId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"snapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"snapNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastTouchTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"touchSnapshot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedForwarder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionRecipient\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

func GsnBridgeCreateRelayer(
	__gsnCaller gsn.GsnCaller,
	__chainId *big.Int,
	__signer *ecdsa.PrivateKey,
	__contractAddress common.Address, _node NodeRegistryNode, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (txHash common.Hash, err error) {

	__contractABI, err := abi.JSON(strings.NewReader(__contractNodeRegistrySourceABI))
	if err != nil {
		return common.Hash{}, fmt.Errorf("could not parse ABI: %w", err)
	}

	__fRequest, err := __contractABI.Pack("createRelayer", _node, _deadline, _v, _r, _s)
	if err != nil {
		return
	}

	__forwarder, err := __gsnCaller.GetForwarder(__chainId)
	if err != nil {
		return
	}

	__forwarderAddress, err := __gsnCaller.GetForwarderAddress(__chainId)
	if err != nil {
		return
	}

	__signerAddress := crypto.PubkeyToAddress(__signer.PublicKey)

	__nonce, err := __forwarder.GetNonce(&bind.CallOpts{}, __signerAddress)
	if err != nil {

		return
	}

	__req := &wrappers.IForwarderForwardRequest{
		From:  __signerAddress,
		To:    __contractAddress,
		Value: big.NewInt(0),
		Gas:   big.NewInt(1e6),
		Nonce: __nonce,
		Data:  __fRequest,
	}

	__typedData, err := gsn.NewForwardRequestTypedData(
		__req,
		__forwarderAddress.String(),
		__contractNodeRegistrySourceABI,
		"createRelayer", _node, _deadline, _v, _r, _s)
	if err != nil {
		return
	}

	__typedDataSignature, _, err := gsn.NewSignature(__typedData, __signer)
	if err != nil {
		return
	}

	__domainSeparatorHash, err := gsn.NewDomainSeparatorHash(__typedData)
	if err != nil {
		return
	}

	__genericParams, err := __forwarder.GENERICPARAMS(&bind.CallOpts{})
	if err != nil {
		return
	}

	__reqTypeHash, err := gsn.NewRequestTypeHash(__genericParams)
	if err != nil {
		return
	}

	return __gsnCaller.Execute(__chainId, *__req, __domainSeparatorHash, __reqTypeHash, nil, __typedDataSignature)
}

func GsnBridgeTouchSnapshot(
	__gsnCaller gsn.GsnCaller,
	__chainId *big.Int,
	__signer *ecdsa.PrivateKey,
	__contractAddress common.Address) (txHash common.Hash, err error) {

	__contractABI, err := abi.JSON(strings.NewReader(__contractNodeRegistrySourceABI))
	if err != nil {
		return common.Hash{}, fmt.Errorf("could not parse ABI: %w", err)
	}

	__fRequest, err := __contractABI.Pack("touchSnapshot")
	if err != nil {
		return
	}

	__forwarder, err := __gsnCaller.GetForwarder(__chainId)
	if err != nil {
		return
	}

	__forwarderAddress, err := __gsnCaller.GetForwarderAddress(__chainId)
	if err != nil {
		return
	}

	__signerAddress := crypto.PubkeyToAddress(__signer.PublicKey)

	__nonce, err := __forwarder.GetNonce(&bind.CallOpts{}, __signerAddress)
	if err != nil {

		return
	}

	__req := &wrappers.IForwarderForwardRequest{
		From:  __signerAddress,
		To:    __contractAddress,
		Value: big.NewInt(0),
		Gas:   big.NewInt(1e6),
		Nonce: __nonce,
		Data:  __fRequest,
	}

	__typedData, err := gsn.NewForwardRequestTypedData(
		__req,
		__forwarderAddress.String(),
		__contractNodeRegistrySourceABI,
		"touchSnapshot")
	if err != nil {
		return
	}

	__typedDataSignature, _, err := gsn.NewSignature(__typedData, __signer)
	if err != nil {
		return
	}

	__domainSeparatorHash, err := gsn.NewDomainSeparatorHash(__typedData)
	if err != nil {
		return
	}

	__genericParams, err := __forwarder.GENERICPARAMS(&bind.CallOpts{})
	if err != nil {
		return
	}

	__reqTypeHash, err := gsn.NewRequestTypeHash(__genericParams)
	if err != nil {
		return
	}

	return __gsnCaller.Execute(__chainId, *__req, __domainSeparatorHash, __reqTypeHash, nil, __typedDataSignature)
}
