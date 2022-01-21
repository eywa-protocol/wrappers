// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package test

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

// TestForwardForwardRequest is an auto generated low-level Go binding around an user-defined struct.
type TestForwardForwardRequest struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Gas   *big.Int
	Nonce *big.Int
	Data  []byte
}

var __contractTestForwardSourceABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_forwarder\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"FooCalled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_val\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_str\",\"type\":\"string\"}],\"name\":\"foo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"str\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structTestForward.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"requestTypeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"suffixData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"testExecute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"ret\",\"type\":\"string\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedForwarder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"val\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionRecipient\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

func GsnBridgeFoo(
	__gsnCaller gsn.GsnCaller,
	__chainId *big.Int,
	__signer *ecdsa.PrivateKey,
	__contractAddress common.Address, _val *big.Int, _str string) (txHash common.Hash, err error) {

	__contractABI, err := abi.JSON(strings.NewReader(__contractTestForwardSourceABI))
	if err != nil {
		return common.Hash{}, fmt.Errorf("could not parse ABI: %w", err)
	}

	__fRequest, err := __contractABI.Pack("foo", _val, _str)
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
		__contractTestForwardSourceABI,
		"foo", _val, _str)
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

func GsnBridgeTestExecute(
	__gsnCaller gsn.GsnCaller,
	__chainId *big.Int,
	__signer *ecdsa.PrivateKey,
	__contractAddress common.Address, req TestForwardForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) (txHash common.Hash, err error) {

	__contractABI, err := abi.JSON(strings.NewReader(__contractTestForwardSourceABI))
	if err != nil {
		return common.Hash{}, fmt.Errorf("could not parse ABI: %w", err)
	}

	__fRequest, err := __contractABI.Pack("testExecute", req, domainSeparator, requestTypeHash, suffixData, sig)
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
		__contractTestForwardSourceABI,
		"testExecute", req, domainSeparator, requestTypeHash, suffixData, sig)
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
