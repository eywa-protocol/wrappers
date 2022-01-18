// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gsn

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
	"github.com/sirupsen/logrus"
)

func GsnBridgeAddContractBind(
	__gsnCaller GsnCaller,
	__chainId *big.Int,
	__signer *ecdsa.PrivateKey,
	__contractAddress common.Address, from [32]byte, oppositeBridge [32]byte, to [32]byte) (txHash common.Hash, err error) {

	__contractABI, err := abi.JSON(strings.NewReader(wrappers.BridgeABI))
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

	logrus.Infof("forwarderAddress: %s", __forwarderAddress.String())
	logrus.Infof("nodeRegistryAddress: %s", __contractAddress.String())

	__signerAddress := crypto.PubkeyToAddress(__signer.PublicKey)

	logrus.Infof("ownerAddress: %s", __signerAddress.String())

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

	__typedData, err := NewForwardRequestTypedData(
		__req,
		__forwarderAddress.String(),
		wrappers.BridgeABI,
		"addContractBind", from, oppositeBridge, to)
	if err != nil {
		return
	}

	__typedDataSignature, _, err := NewSignature(__typedData, __signer)
	if err != nil {

		return
	}

	__domainSeparatorHash, err := NewDomainSeparatorHash(__typedData)
	if err != nil {

		return
	}

	__genericParams, err := __forwarder.GENERICPARAMS(&bind.CallOpts{})
	if err != nil {

		return
	}

	__reqTypeHash, err := NewRequestTypeHash(__genericParams)
	if err != nil {

		return
	}

	return __gsnCaller.Execute(__chainId, *__req, __domainSeparatorHash, __reqTypeHash, nil, __typedDataSignature)
}

func GsnBridgeDaoTransferOwnership(
	__gsnCaller GsnCaller,
	__chainId *big.Int,
	__signer *ecdsa.PrivateKey,
	__contractAddress common.Address, newDao common.Address) (txHash common.Hash, err error) {

	__contractABI, err := abi.JSON(strings.NewReader(wrappers.BridgeABI))
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

	logrus.Infof("forwarderAddress: %s", __forwarderAddress.String())
	logrus.Infof("nodeRegistryAddress: %s", __contractAddress.String())

	__signerAddress := crypto.PubkeyToAddress(__signer.PublicKey)

	logrus.Infof("ownerAddress: %s", __signerAddress.String())

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

	__typedData, err := NewForwardRequestTypedData(
		__req,
		__forwarderAddress.String(),
		wrappers.BridgeABI,
		"daoTransferOwnership", newDao)
	if err != nil {
		return
	}

	__typedDataSignature, _, err := NewSignature(__typedData, __signer)
	if err != nil {

		return
	}

	__domainSeparatorHash, err := NewDomainSeparatorHash(__typedData)
	if err != nil {

		return
	}

	__genericParams, err := __forwarder.GENERICPARAMS(&bind.CallOpts{})
	if err != nil {

		return
	}

	__reqTypeHash, err := NewRequestTypeHash(__genericParams)
	if err != nil {

		return
	}

	return __gsnCaller.Execute(__chainId, *__req, __domainSeparatorHash, __reqTypeHash, nil, __typedDataSignature)
}

func GsnBridgeDaoUpdateEpochRequest(
	__gsnCaller GsnCaller,
	__chainId *big.Int,
	__signer *ecdsa.PrivateKey,
	__contractAddress common.Address, resetEpoch bool) (txHash common.Hash, err error) {

	__contractABI, err := abi.JSON(strings.NewReader(wrappers.BridgeABI))
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

	logrus.Infof("forwarderAddress: %s", __forwarderAddress.String())
	logrus.Infof("nodeRegistryAddress: %s", __contractAddress.String())

	__signerAddress := crypto.PubkeyToAddress(__signer.PublicKey)

	logrus.Infof("ownerAddress: %s", __signerAddress.String())

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

	__typedData, err := NewForwardRequestTypedData(
		__req,
		__forwarderAddress.String(),
		wrappers.BridgeABI,
		"daoUpdateEpochRequest", resetEpoch)
	if err != nil {
		return
	}

	__typedDataSignature, _, err := NewSignature(__typedData, __signer)
	if err != nil {

		return
	}

	__domainSeparatorHash, err := NewDomainSeparatorHash(__typedData)
	if err != nil {

		return
	}

	__genericParams, err := __forwarder.GENERICPARAMS(&bind.CallOpts{})
	if err != nil {

		return
	}

	__reqTypeHash, err := NewRequestTypeHash(__genericParams)
	if err != nil {

		return
	}

	return __gsnCaller.Execute(__chainId, *__req, __domainSeparatorHash, __reqTypeHash, nil, __typedDataSignature)
}

func GsnBridgeInitialize(
	__gsnCaller GsnCaller,
	__chainId *big.Int,
	__signer *ecdsa.PrivateKey,
	__contractAddress common.Address, forwarder common.Address) (txHash common.Hash, err error) {

	__contractABI, err := abi.JSON(strings.NewReader(wrappers.BridgeABI))
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

	logrus.Infof("forwarderAddress: %s", __forwarderAddress.String())
	logrus.Infof("nodeRegistryAddress: %s", __contractAddress.String())

	__signerAddress := crypto.PubkeyToAddress(__signer.PublicKey)

	logrus.Infof("ownerAddress: %s", __signerAddress.String())

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

	__typedData, err := NewForwardRequestTypedData(
		__req,
		__forwarderAddress.String(),
		wrappers.BridgeABI,
		"initialize", forwarder)
	if err != nil {
		return
	}

	__typedDataSignature, _, err := NewSignature(__typedData, __signer)
	if err != nil {

		return
	}

	__domainSeparatorHash, err := NewDomainSeparatorHash(__typedData)
	if err != nil {

		return
	}

	__genericParams, err := __forwarder.GENERICPARAMS(&bind.CallOpts{})
	if err != nil {

		return
	}

	__reqTypeHash, err := NewRequestTypeHash(__genericParams)
	if err != nil {

		return
	}

	return __gsnCaller.Execute(__chainId, *__req, __domainSeparatorHash, __reqTypeHash, nil, __typedDataSignature)
}

func GsnBridgeReceiveRequestV2(
	__gsnCaller GsnCaller,
	__chainId *big.Int,
	__signer *ecdsa.PrivateKey,
	__contractAddress common.Address, _reqId [32]byte, _sel []byte, _receiveSide common.Address, _bridgeFrom [32]byte, _votersPubKey []byte, _votersSignature []byte, _votersMask *big.Int) (txHash common.Hash, err error) {

	__contractABI, err := abi.JSON(strings.NewReader(wrappers.BridgeABI))
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

	logrus.Infof("forwarderAddress: %s", __forwarderAddress.String())
	logrus.Infof("nodeRegistryAddress: %s", __contractAddress.String())

	__signerAddress := crypto.PubkeyToAddress(__signer.PublicKey)

	logrus.Infof("ownerAddress: %s", __signerAddress.String())

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

	__typedData, err := NewForwardRequestTypedData(
		__req,
		__forwarderAddress.String(),
		wrappers.BridgeABI,
		"receiveRequestV2", _reqId, _sel, _receiveSide, _bridgeFrom, _votersPubKey, _votersSignature, _votersMask)
	if err != nil {
		return
	}

	__typedDataSignature, _, err := NewSignature(__typedData, __signer)
	if err != nil {

		return
	}

	__domainSeparatorHash, err := NewDomainSeparatorHash(__typedData)
	if err != nil {

		return
	}

	__genericParams, err := __forwarder.GENERICPARAMS(&bind.CallOpts{})
	if err != nil {

		return
	}

	__reqTypeHash, err := NewRequestTypeHash(__genericParams)
	if err != nil {

		return
	}

	return __gsnCaller.Execute(__chainId, *__req, __domainSeparatorHash, __reqTypeHash, nil, __typedDataSignature)
}

func GsnBridgeTransmitRequestV2(
	__gsnCaller GsnCaller,
	__chainId *big.Int,
	__signer *ecdsa.PrivateKey,
	__contractAddress common.Address, _selector []byte, receiveSide common.Address, oppositeBridge common.Address, chainId *big.Int, requestId [32]byte, sender common.Address, nonce *big.Int) (txHash common.Hash, err error) {

	__contractABI, err := abi.JSON(strings.NewReader(wrappers.BridgeABI))
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

	logrus.Infof("forwarderAddress: %s", __forwarderAddress.String())
	logrus.Infof("nodeRegistryAddress: %s", __contractAddress.String())

	__signerAddress := crypto.PubkeyToAddress(__signer.PublicKey)

	logrus.Infof("ownerAddress: %s", __signerAddress.String())

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

	__typedData, err := NewForwardRequestTypedData(
		__req,
		__forwarderAddress.String(),
		wrappers.BridgeABI,
		"transmitRequestV2", _selector, receiveSide, oppositeBridge, chainId, requestId, sender, nonce)
	if err != nil {
		return
	}

	__typedDataSignature, _, err := NewSignature(__typedData, __signer)
	if err != nil {

		return
	}

	__domainSeparatorHash, err := NewDomainSeparatorHash(__typedData)
	if err != nil {

		return
	}

	__genericParams, err := __forwarder.GENERICPARAMS(&bind.CallOpts{})
	if err != nil {

		return
	}

	__reqTypeHash, err := NewRequestTypeHash(__genericParams)
	if err != nil {

		return
	}

	return __gsnCaller.Execute(__chainId, *__req, __domainSeparatorHash, __reqTypeHash, nil, __typedDataSignature)
}

func GsnBridgeTransmitRequestV2Solana(
	__gsnCaller GsnCaller,
	__chainId *big.Int,
	__signer *ecdsa.PrivateKey,
	__contractAddress common.Address, _selector []byte, receiveSide [32]byte, oppositeBridge [32]byte, chainId *big.Int, requestId [32]byte, sender common.Address, nonce *big.Int) (txHash common.Hash, err error) {

	__contractABI, err := abi.JSON(strings.NewReader(wrappers.BridgeABI))
	if err != nil {
		return common.Hash{}, fmt.Errorf("could not parse ABI: %w", err)
	}

	__fRequest, err := __contractABI.Pack("transmitRequestV2_solana", _selector, receiveSide, oppositeBridge, chainId, requestId, sender, nonce)
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

	logrus.Infof("forwarderAddress: %s", __forwarderAddress.String())
	logrus.Infof("nodeRegistryAddress: %s", __contractAddress.String())

	__signerAddress := crypto.PubkeyToAddress(__signer.PublicKey)

	logrus.Infof("ownerAddress: %s", __signerAddress.String())

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

	__typedData, err := NewForwardRequestTypedData(
		__req,
		__forwarderAddress.String(),
		wrappers.BridgeABI,
		"transmitRequestV2_solana", _selector, receiveSide, oppositeBridge, chainId, requestId, sender, nonce)
	if err != nil {
		return
	}

	__typedDataSignature, _, err := NewSignature(__typedData, __signer)
	if err != nil {

		return
	}

	__domainSeparatorHash, err := NewDomainSeparatorHash(__typedData)
	if err != nil {

		return
	}

	__genericParams, err := __forwarder.GENERICPARAMS(&bind.CallOpts{})
	if err != nil {

		return
	}

	__reqTypeHash, err := NewRequestTypeHash(__genericParams)
	if err != nil {

		return
	}

	return __gsnCaller.Execute(__chainId, *__req, __domainSeparatorHash, __reqTypeHash, nil, __typedDataSignature)
}

func GsnBridgeUpdateEpoch(
	__gsnCaller GsnCaller,
	__chainId *big.Int,
	__signer *ecdsa.PrivateKey,
	__contractAddress common.Address, _newKey []byte, _votersPubKey []byte, _votersSignature []byte, _votersMask *big.Int, _newEpochParticipantsNum uint8, _newEpochNum uint32) (txHash common.Hash, err error) {

	__contractABI, err := abi.JSON(strings.NewReader(wrappers.BridgeABI))
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

	logrus.Infof("forwarderAddress: %s", __forwarderAddress.String())
	logrus.Infof("nodeRegistryAddress: %s", __contractAddress.String())

	__signerAddress := crypto.PubkeyToAddress(__signer.PublicKey)

	logrus.Infof("ownerAddress: %s", __signerAddress.String())

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

	__typedData, err := NewForwardRequestTypedData(
		__req,
		__forwarderAddress.String(),
		wrappers.BridgeABI,
		"updateEpoch", _newKey, _votersPubKey, _votersSignature, _votersMask, _newEpochParticipantsNum, _newEpochNum)
	if err != nil {
		return
	}

	__typedDataSignature, _, err := NewSignature(__typedData, __signer)
	if err != nil {

		return
	}

	__domainSeparatorHash, err := NewDomainSeparatorHash(__typedData)
	if err != nil {

		return
	}

	__genericParams, err := __forwarder.GENERICPARAMS(&bind.CallOpts{})
	if err != nil {

		return
	}

	__reqTypeHash, err := NewRequestTypeHash(__genericParams)
	if err != nil {

		return
	}

	return __gsnCaller.Execute(__chainId, *__req, __domainSeparatorHash, __reqTypeHash, nil, __typedDataSignature)
}
