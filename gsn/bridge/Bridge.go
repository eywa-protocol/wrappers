// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridge

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/eywa-protocol/wrappers"

	"github.com/eywa-protocol/wrappers/gsn"
)

var (
	_ = big.NewInt(0) // Fake usage
)

var __contractBridgeSourceABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"oldEpochKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newEpochKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"requested\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epochNum\",\"type\":\"uint32\"}],\"name\":\"NewEpoch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"requestType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"selector\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainid\",\"type\":\"uint256\"}],\"name\":\"OracleRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"requestType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"bridge\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"selector\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"oppositeBridge\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainid\",\"type\":\"uint256\"}],\"name\":\"OracleRequestSolana\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"reqId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"bridgeFrom\",\"type\":\"bytes32\"}],\"name\":\"ReceiveRequest\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_listNode\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"from\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"oppositeBridge\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"}],\"name\":\"addContractBind\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dao\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDao\",\"type\":\"address\"}],\"name\":\"daoTransferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"resetEpoch\",\"type\":\"bool\"}],\"name\":\"daoUpdateEpochRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochNum\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochParticipantsNum\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEpoch\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"oppositeBridge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"receiveSide\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"from\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"prepareRqId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_reqId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_sel\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_receiveSide\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_bridgeFrom\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_votersPubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_votersSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_votersMask\",\"type\":\"uint256\"}],\"name\":\"receiveRequestV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_forwarder\",\"type\":\"address\"}],\"name\":\"setTrustedForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_selector\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"transmitRequestV2\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_selector\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"receiveSide\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"oppositeBridge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"transmitRequestV2ToSolana\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedForwarder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_newKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_votersPubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_votersSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_votersMask\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_newEpochParticipantsNum\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"_newEpochNum\",\"type\":\"uint32\"}],\"name\":\"updateEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionRecipient\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

func AddContractBind(gsnParams *gsn.CallOpts, args ...interface{}) (txHash common.Hash, err error) {
	gsnCallParams := &wrappers.GsnCallOpts{
		GsnCaller:       gsnParams.GsnCaller,
		ChainId:         gsnParams.ChainId,
		Signer:          gsnParams.Signer,
		ContractAddress: gsnParams.ContractAddress,
	}
	return wrappers.GsnExecutor(gsnCallParams, __contractBridgeSourceABI, "addContractBind", args...)
}

func DaoTransferOwnership(gsnParams *gsn.CallOpts, args ...interface{}) (txHash common.Hash, err error) {
	gsnCallParams := &wrappers.GsnCallOpts{
		GsnCaller:       gsnParams.GsnCaller,
		ChainId:         gsnParams.ChainId,
		Signer:          gsnParams.Signer,
		ContractAddress: gsnParams.ContractAddress,
	}
	return wrappers.GsnExecutor(gsnCallParams, __contractBridgeSourceABI, "daoTransferOwnership", args...)
}

func DaoUpdateEpochRequest(gsnParams *gsn.CallOpts, args ...interface{}) (txHash common.Hash, err error) {
	gsnCallParams := &wrappers.GsnCallOpts{
		GsnCaller:       gsnParams.GsnCaller,
		ChainId:         gsnParams.ChainId,
		Signer:          gsnParams.Signer,
		ContractAddress: gsnParams.ContractAddress,
	}
	return wrappers.GsnExecutor(gsnCallParams, __contractBridgeSourceABI, "daoUpdateEpochRequest", args...)
}

func Initialize(gsnParams *gsn.CallOpts, args ...interface{}) (txHash common.Hash, err error) {
	gsnCallParams := &wrappers.GsnCallOpts{
		GsnCaller:       gsnParams.GsnCaller,
		ChainId:         gsnParams.ChainId,
		Signer:          gsnParams.Signer,
		ContractAddress: gsnParams.ContractAddress,
	}
	return wrappers.GsnExecutor(gsnCallParams, __contractBridgeSourceABI, "initialize", args...)
}

func ReceiveRequestV2(gsnParams *gsn.CallOpts, args ...interface{}) (txHash common.Hash, err error) {
	gsnCallParams := &wrappers.GsnCallOpts{
		GsnCaller:       gsnParams.GsnCaller,
		ChainId:         gsnParams.ChainId,
		Signer:          gsnParams.Signer,
		ContractAddress: gsnParams.ContractAddress,
	}
	return wrappers.GsnExecutor(gsnCallParams, __contractBridgeSourceABI, "receiveRequestV2", args...)
}

func RenounceOwnership(gsnParams *gsn.CallOpts, args ...interface{}) (txHash common.Hash, err error) {
	gsnCallParams := &wrappers.GsnCallOpts{
		GsnCaller:       gsnParams.GsnCaller,
		ChainId:         gsnParams.ChainId,
		Signer:          gsnParams.Signer,
		ContractAddress: gsnParams.ContractAddress,
	}
	return wrappers.GsnExecutor(gsnCallParams, __contractBridgeSourceABI, "renounceOwnership", args...)
}

func SetTrustedForwarder(gsnParams *gsn.CallOpts, args ...interface{}) (txHash common.Hash, err error) {
	gsnCallParams := &wrappers.GsnCallOpts{
		GsnCaller:       gsnParams.GsnCaller,
		ChainId:         gsnParams.ChainId,
		Signer:          gsnParams.Signer,
		ContractAddress: gsnParams.ContractAddress,
	}
	return wrappers.GsnExecutor(gsnCallParams, __contractBridgeSourceABI, "setTrustedForwarder", args...)
}

func TransferOwnership(gsnParams *gsn.CallOpts, args ...interface{}) (txHash common.Hash, err error) {
	gsnCallParams := &wrappers.GsnCallOpts{
		GsnCaller:       gsnParams.GsnCaller,
		ChainId:         gsnParams.ChainId,
		Signer:          gsnParams.Signer,
		ContractAddress: gsnParams.ContractAddress,
	}
	return wrappers.GsnExecutor(gsnCallParams, __contractBridgeSourceABI, "transferOwnership", args...)
}

func TransmitRequestV2(gsnParams *gsn.CallOpts, args ...interface{}) (txHash common.Hash, err error) {
	gsnCallParams := &wrappers.GsnCallOpts{
		GsnCaller:       gsnParams.GsnCaller,
		ChainId:         gsnParams.ChainId,
		Signer:          gsnParams.Signer,
		ContractAddress: gsnParams.ContractAddress,
	}
	return wrappers.GsnExecutor(gsnCallParams, __contractBridgeSourceABI, "transmitRequestV2", args...)
}

func TransmitRequestV2ToSolana(gsnParams *gsn.CallOpts, args ...interface{}) (txHash common.Hash, err error) {
	gsnCallParams := &wrappers.GsnCallOpts{
		GsnCaller:       gsnParams.GsnCaller,
		ChainId:         gsnParams.ChainId,
		Signer:          gsnParams.Signer,
		ContractAddress: gsnParams.ContractAddress,
	}
	return wrappers.GsnExecutor(gsnCallParams, __contractBridgeSourceABI, "transmitRequestV2ToSolana", args...)
}

func UpdateEpoch(gsnParams *gsn.CallOpts, args ...interface{}) (txHash common.Hash, err error) {
	gsnCallParams := &wrappers.GsnCallOpts{
		GsnCaller:       gsnParams.GsnCaller,
		ChainId:         gsnParams.ChainId,
		Signer:          gsnParams.Signer,
		ContractAddress: gsnParams.ContractAddress,
	}
	return wrappers.GsnExecutor(gsnCallParams, __contractBridgeSourceABI, "updateEpoch", args...)
}
