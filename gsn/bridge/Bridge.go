// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridge

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

var __contractBridgeSourceABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"oldEpochKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newEpochKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"requested\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epochNum\",\"type\":\"uint32\"}],\"name\":\"NewEpoch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"requestType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"selector\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainid\",\"type\":\"uint256\"}],\"name\":\"OracleRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"requestType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"bridge\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"selector\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"oppositeBridge\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainid\",\"type\":\"uint256\"}],\"name\":\"OracleRequestSolana\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousDao\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newDao\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"reqId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"bridgeFrom\",\"type\":\"bytes32\"}],\"name\":\"ReceiveRequest\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_listNode\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"from\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"oppositeBridge\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"}],\"name\":\"addContractBind\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dao\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDao\",\"type\":\"address\"}],\"name\":\"daoTransferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"resetEpoch\",\"type\":\"bool\"}],\"name\":\"daoUpdateEpochRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochNum\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochParticipantsNum\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEpoch\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"oppositeBridge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"receiveSide\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"from\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"prepareRqId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_reqId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_sel\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_receiveSide\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_bridgeFrom\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_votersPubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_votersSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_votersMask\",\"type\":\"uint256\"}],\"name\":\"receiveRequestV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_selector\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"transmitRequestV2\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_selector\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"receiveSide\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"oppositeBridge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"transmitRequestV2ToSolana\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedForwarder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_newKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_votersPubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_votersSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_votersMask\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_newEpochParticipantsNum\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"_newEpochNum\",\"type\":\"uint32\"}],\"name\":\"updateEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionRecipient\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

func GsnBridgeAddContractBind(
	__gsnCaller gsn.GsnCaller,
	__chainId *big.Int,
	__signer *ecdsa.PrivateKey,
	__contractAddress common.Address, from [32]byte, oppositeBridge [32]byte, to [32]byte) (txHash common.Hash, err error) {

	__contractABI, err := abi.JSON(strings.NewReader(__contractBridgeSourceABI))
	if err != nil {
		return common.Hash{}, fmt.Errorf("could not parse ABI: %w", err)
	}

	__fRequest, err := __contractABI.Pack("addContractBind", from, oppositeBridge, to)
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
		__contractBridgeSourceABI,
		"addContractBind", from, oppositeBridge, to)
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

func GsnBridgeDaoTransferOwnership(
	__gsnCaller gsn.GsnCaller,
	__chainId *big.Int,
	__signer *ecdsa.PrivateKey,
	__contractAddress common.Address, newDao common.Address) (txHash common.Hash, err error) {

	__contractABI, err := abi.JSON(strings.NewReader(__contractBridgeSourceABI))
	if err != nil {
		return common.Hash{}, fmt.Errorf("could not parse ABI: %w", err)
	}

	__fRequest, err := __contractABI.Pack("daoTransferOwnership", newDao)
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
		__contractBridgeSourceABI,
		"daoTransferOwnership", newDao)
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

func GsnBridgeDaoUpdateEpochRequest(
	__gsnCaller gsn.GsnCaller,
	__chainId *big.Int,
	__signer *ecdsa.PrivateKey,
	__contractAddress common.Address, resetEpoch bool) (txHash common.Hash, err error) {

	__contractABI, err := abi.JSON(strings.NewReader(__contractBridgeSourceABI))
	if err != nil {
		return common.Hash{}, fmt.Errorf("could not parse ABI: %w", err)
	}

	__fRequest, err := __contractABI.Pack("daoUpdateEpochRequest", resetEpoch)
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
		__contractBridgeSourceABI,
		"daoUpdateEpochRequest", resetEpoch)
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

func GsnBridgeInitialize(
	__gsnCaller gsn.GsnCaller,
	__chainId *big.Int,
	__signer *ecdsa.PrivateKey,
	__contractAddress common.Address, forwarder common.Address) (txHash common.Hash, err error) {

	__contractABI, err := abi.JSON(strings.NewReader(__contractBridgeSourceABI))
	if err != nil {
		return common.Hash{}, fmt.Errorf("could not parse ABI: %w", err)
	}

	__fRequest, err := __contractABI.Pack("initialize", forwarder)
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
		__contractBridgeSourceABI,
		"initialize", forwarder)
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

func GsnBridgeReceiveRequestV2(
	__gsnCaller gsn.GsnCaller,
	__chainId *big.Int,
	__signer *ecdsa.PrivateKey,
	__contractAddress common.Address, _reqId [32]byte, _sel []byte, _receiveSide common.Address, _bridgeFrom [32]byte, _votersPubKey []byte, _votersSignature []byte, _votersMask *big.Int) (txHash common.Hash, err error) {

	__contractABI, err := abi.JSON(strings.NewReader(__contractBridgeSourceABI))
	if err != nil {
		return common.Hash{}, fmt.Errorf("could not parse ABI: %w", err)
	}

	__fRequest, err := __contractABI.Pack("receiveRequestV2", _reqId, _sel, _receiveSide, _bridgeFrom, _votersPubKey, _votersSignature, _votersMask)
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
		__contractBridgeSourceABI,
		"receiveRequestV2", _reqId, _sel, _receiveSide, _bridgeFrom, _votersPubKey, _votersSignature, _votersMask)
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

func GsnBridgeTransmitRequestV2(
	__gsnCaller gsn.GsnCaller,
	__chainId *big.Int,
	__signer *ecdsa.PrivateKey,
	__contractAddress common.Address, _selector []byte, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, requestId [32]byte, sender common.Address, nonce *big.Int) (txHash common.Hash, err error) {

	__contractABI, err := abi.JSON(strings.NewReader(__contractBridgeSourceABI))
	if err != nil {
		return common.Hash{}, fmt.Errorf("could not parse ABI: %w", err)
	}

	__fRequest, err := __contractABI.Pack("transmitRequestV2", _selector, receiveSide, oppositeBridge, chainId, requestId, sender, nonce)
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
		__contractBridgeSourceABI,
		"transmitRequestV2", _selector, receiveSide, oppositeBridge, chainId, requestId, sender, nonce)
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

func GsnBridgeTransmitRequestV2ToSolana(
	__gsnCaller gsn.GsnCaller,
	__chainId *big.Int,
	__signer *ecdsa.PrivateKey,
	__contractAddress common.Address, _selector []byte, receiveSide [32]byte, oppositeBridge [32]byte, chainId *big.Int, requestId [32]byte, sender common.Address, nonce *big.Int) (txHash common.Hash, err error) {

	__contractABI, err := abi.JSON(strings.NewReader(__contractBridgeSourceABI))
	if err != nil {
		return common.Hash{}, fmt.Errorf("could not parse ABI: %w", err)
	}

	__fRequest, err := __contractABI.Pack("transmitRequestV2ToSolana", _selector, receiveSide, oppositeBridge, chainId, requestId, sender, nonce)
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
		__contractBridgeSourceABI,
		"transmitRequestV2ToSolana", _selector, receiveSide, oppositeBridge, chainId, requestId, sender, nonce)
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

func GsnBridgeUpdateEpoch(
	__gsnCaller gsn.GsnCaller,
	__chainId *big.Int,
	__signer *ecdsa.PrivateKey,
	__contractAddress common.Address, _newKey []byte, _votersPubKey []byte, _votersSignature []byte, _votersMask *big.Int, _newEpochParticipantsNum uint8, _newEpochNum uint32) (txHash common.Hash, err error) {

	__contractABI, err := abi.JSON(strings.NewReader(__contractBridgeSourceABI))
	if err != nil {
		return common.Hash{}, fmt.Errorf("could not parse ABI: %w", err)
	}

	__fRequest, err := __contractABI.Pack("updateEpoch", _newKey, _votersPubKey, _votersSignature, _votersMask, _newEpochParticipantsNum, _newEpochNum)
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
		__contractBridgeSourceABI,
		"updateEpoch", _newKey, _votersPubKey, _votersSignature, _votersMask, _newEpochParticipantsNum, _newEpochNum)
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
