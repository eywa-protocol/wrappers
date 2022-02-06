// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registry

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/eywa-protocol/wrappers"
	"github.com/eywa-protocol/wrappers/gsn"
)

var (
	_ = big.NewInt(0) // Fake usage
)

// NodeRegistryNode is an auto generated low-level Go binding around an user-defined struct.
type NodeRegistryNode struct {
	Owner     common.Address
	Pool      common.Address
	HostId    string
	BlsPubKey []byte
	NodeId    *big.Int
}

var __contractNodeRegistrySourceABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayerPool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hostId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"blsPubKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nodeId\",\"type\":\"uint256\"}],\"name\":\"CreatedRelayer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"oldEpochKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newEpochKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"requested\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epochNum\",\"type\":\"uint32\"}],\"name\":\"NewEpoch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"snapNum\",\"type\":\"uint256\"}],\"name\":\"NewSnapshot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"requestType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"selector\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainid\",\"type\":\"uint256\"}],\"name\":\"OracleRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"requestType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"bridge\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"selector\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"oppositeBridge\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainid\",\"type\":\"uint256\"}],\"name\":\"OracleRequestSolana\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"reqId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"bridgeFrom\",\"type\":\"bytes32\"}],\"name\":\"ReceiveRequest\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EYWA\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_COLLATERAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_listNode\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"from\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"oppositeBridge\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"}],\"name\":\"addContractBind\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"checkPermissionTrustList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"hostId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"blsPubKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"nodeId\",\"type\":\"uint256\"}],\"internalType\":\"structNodeRegistry.Node\",\"name\":\"_node\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"createRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dao\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDao\",\"type\":\"address\"}],\"name\":\"daoTransferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"resetEpoch\",\"type\":\"bool\"}],\"name\":\"daoUpdateEpochRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochNum\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochParticipantsNum\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBLSPubKeys\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEpoch\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getNode\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"hostId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"blsPubKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"nodeId\",\"type\":\"uint256\"}],\"internalType\":\"structNodeRegistry.Node\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodes\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"hostId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"blsPubKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"nodeId\",\"type\":\"uint256\"}],\"internalType\":\"structNodeRegistry.Node[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSnapshot\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"hostIds\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_EYWA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_forwarder\",\"type\":\"address\"}],\"name\":\"initialize2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nodeExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ownedNodes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"hostId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"blsPubKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"nodeId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"oppositeBridge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"receiveSide\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"from\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"prepareRqId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_reqId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_sel\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_receiveSide\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_bridgeFrom\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_votersPubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_votersSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_votersMask\",\"type\":\"uint256\"}],\"name\":\"receiveRequestV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_forwarder\",\"type\":\"address\"}],\"name\":\"setTrustedForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"snapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"snapNum\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_selector\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiveSide\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oppositeBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"transmitRequestV2\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_selector\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"receiveSide\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"oppositeBridge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"transmitRequestV2ToSolana\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedForwarder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_newKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_votersPubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_votersSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_votersMask\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_newEpochParticipantsNum\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"_newEpochNum\",\"type\":\"uint32\"}],\"name\":\"updateEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionRecipient\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

func AddContractBind(gsnParams *gsn.CallOpts, args ...interface{}) (txHash common.Hash, err error) {
	gsnCallParams := &wrappers.GsnCallOpts{
		GsnCaller:       gsnParams.GsnCaller,
		ChainId:         gsnParams.ChainId,
		Signer:          gsnParams.Signer,
		ContractAddress: gsnParams.ContractAddress,
	}
	return wrappers.GsnExecutor(gsnCallParams, __contractNodeRegistrySourceABI, "addContractBind", args...)
}

func CreateRelayer(gsnParams *gsn.CallOpts, args ...interface{}) (txHash common.Hash, err error) {
	gsnCallParams := &wrappers.GsnCallOpts{
		GsnCaller:       gsnParams.GsnCaller,
		ChainId:         gsnParams.ChainId,
		Signer:          gsnParams.Signer,
		ContractAddress: gsnParams.ContractAddress,
	}
	return wrappers.GsnExecutor(gsnCallParams, __contractNodeRegistrySourceABI, "createRelayer", args...)
}

func DaoTransferOwnership(gsnParams *gsn.CallOpts, args ...interface{}) (txHash common.Hash, err error) {
	gsnCallParams := &wrappers.GsnCallOpts{
		GsnCaller:       gsnParams.GsnCaller,
		ChainId:         gsnParams.ChainId,
		Signer:          gsnParams.Signer,
		ContractAddress: gsnParams.ContractAddress,
	}
	return wrappers.GsnExecutor(gsnCallParams, __contractNodeRegistrySourceABI, "daoTransferOwnership", args...)
}

func DaoUpdateEpochRequest(gsnParams *gsn.CallOpts, args ...interface{}) (txHash common.Hash, err error) {
	gsnCallParams := &wrappers.GsnCallOpts{
		GsnCaller:       gsnParams.GsnCaller,
		ChainId:         gsnParams.ChainId,
		Signer:          gsnParams.Signer,
		ContractAddress: gsnParams.ContractAddress,
	}
	return wrappers.GsnExecutor(gsnCallParams, __contractNodeRegistrySourceABI, "daoUpdateEpochRequest", args...)
}

func Initialize(gsnParams *gsn.CallOpts, args ...interface{}) (txHash common.Hash, err error) {
	gsnCallParams := &wrappers.GsnCallOpts{
		GsnCaller:       gsnParams.GsnCaller,
		ChainId:         gsnParams.ChainId,
		Signer:          gsnParams.Signer,
		ContractAddress: gsnParams.ContractAddress,
	}
	return wrappers.GsnExecutor(gsnCallParams, __contractNodeRegistrySourceABI, "initialize", args...)
}

func Initialize2(gsnParams *gsn.CallOpts, args ...interface{}) (txHash common.Hash, err error) {
	gsnCallParams := &wrappers.GsnCallOpts{
		GsnCaller:       gsnParams.GsnCaller,
		ChainId:         gsnParams.ChainId,
		Signer:          gsnParams.Signer,
		ContractAddress: gsnParams.ContractAddress,
	}
	return wrappers.GsnExecutor(gsnCallParams, __contractNodeRegistrySourceABI, "initialize2", args...)
}

func ReceiveRequestV2(gsnParams *gsn.CallOpts, args ...interface{}) (txHash common.Hash, err error) {
	gsnCallParams := &wrappers.GsnCallOpts{
		GsnCaller:       gsnParams.GsnCaller,
		ChainId:         gsnParams.ChainId,
		Signer:          gsnParams.Signer,
		ContractAddress: gsnParams.ContractAddress,
	}
	return wrappers.GsnExecutor(gsnCallParams, __contractNodeRegistrySourceABI, "receiveRequestV2", args...)
}

func RenounceOwnership(gsnParams *gsn.CallOpts, args ...interface{}) (txHash common.Hash, err error) {
	gsnCallParams := &wrappers.GsnCallOpts{
		GsnCaller:       gsnParams.GsnCaller,
		ChainId:         gsnParams.ChainId,
		Signer:          gsnParams.Signer,
		ContractAddress: gsnParams.ContractAddress,
	}
	return wrappers.GsnExecutor(gsnCallParams, __contractNodeRegistrySourceABI, "renounceOwnership", args...)
}

func SetTrustedForwarder(gsnParams *gsn.CallOpts, args ...interface{}) (txHash common.Hash, err error) {
	gsnCallParams := &wrappers.GsnCallOpts{
		GsnCaller:       gsnParams.GsnCaller,
		ChainId:         gsnParams.ChainId,
		Signer:          gsnParams.Signer,
		ContractAddress: gsnParams.ContractAddress,
	}
	return wrappers.GsnExecutor(gsnCallParams, __contractNodeRegistrySourceABI, "setTrustedForwarder", args...)
}

func TransferOwnership(gsnParams *gsn.CallOpts, args ...interface{}) (txHash common.Hash, err error) {
	gsnCallParams := &wrappers.GsnCallOpts{
		GsnCaller:       gsnParams.GsnCaller,
		ChainId:         gsnParams.ChainId,
		Signer:          gsnParams.Signer,
		ContractAddress: gsnParams.ContractAddress,
	}
	return wrappers.GsnExecutor(gsnCallParams, __contractNodeRegistrySourceABI, "transferOwnership", args...)
}

func TransmitRequestV2(gsnParams *gsn.CallOpts, args ...interface{}) (txHash common.Hash, err error) {
	gsnCallParams := &wrappers.GsnCallOpts{
		GsnCaller:       gsnParams.GsnCaller,
		ChainId:         gsnParams.ChainId,
		Signer:          gsnParams.Signer,
		ContractAddress: gsnParams.ContractAddress,
	}
	return wrappers.GsnExecutor(gsnCallParams, __contractNodeRegistrySourceABI, "transmitRequestV2", args...)
}

func TransmitRequestV2ToSolana(gsnParams *gsn.CallOpts, args ...interface{}) (txHash common.Hash, err error) {
	gsnCallParams := &wrappers.GsnCallOpts{
		GsnCaller:       gsnParams.GsnCaller,
		ChainId:         gsnParams.ChainId,
		Signer:          gsnParams.Signer,
		ContractAddress: gsnParams.ContractAddress,
	}
	return wrappers.GsnExecutor(gsnCallParams, __contractNodeRegistrySourceABI, "transmitRequestV2ToSolana", args...)
}

func UpdateEpoch(gsnParams *gsn.CallOpts, args ...interface{}) (txHash common.Hash, err error) {
	gsnCallParams := &wrappers.GsnCallOpts{
		GsnCaller:       gsnParams.GsnCaller,
		ChainId:         gsnParams.ChainId,
		Signer:          gsnParams.Signer,
		ContractAddress: gsnParams.ContractAddress,
	}
	return wrappers.GsnExecutor(gsnCallParams, __contractNodeRegistrySourceABI, "updateEpoch", args...)
}