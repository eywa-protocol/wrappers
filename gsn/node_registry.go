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

func NodeRegistryCreateNode(gsnCaller GsnCaller, chainId *big.Int, signer *ecdsa.PrivateKey, contractAddress common.Address, node wrappers.NodeRegistryNode, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (txHash common.Hash, err error) {

	contractABI, err := abi.JSON(strings.NewReader(wrappers.NodeRegistryABI))
	if err != nil {
		return common.Hash{}, fmt.Errorf("could not parse ABI: %w", err)
	}

	fRequest, err := contractABI.Pack("createRelayer", node, deadline, v, r, s)
	if err != nil {

		return
	}

	forwarder, err := gsnCaller.GetForwarder(chainId)
	if err != nil {

		return
	}

	forwarderAddress, err := gsnCaller.GetForwarderAddress(chainId)
	if err != nil {

		return
	}

	logrus.Infof("forwarderAddress: %s", forwarderAddress.String())
	logrus.Infof("nodeRegistryAddress: %s", contractAddress.String())

	signerAddress := crypto.PubkeyToAddress(signer.PublicKey)

	logrus.Infof("ownerAddress: %s", signerAddress.String())

	nonce, err := forwarder.GetNonce(&bind.CallOpts{}, signerAddress)
	if err != nil {

		return
	}

	req := &wrappers.IForwarderForwardRequest{
		From:  signerAddress,
		To:    contractAddress,
		Value: big.NewInt(0),
		Gas:   big.NewInt(1e6),
		Nonce: nonce,
		Data:  fRequest,
	}

	typedData, err := NewForwardRequestTypedData(
		req,
		forwarderAddress.String(),
		wrappers.NodeRegistryABI,
		"createRelayer", node, deadline, v, r, s)
	if err != nil {

		return
	}

	typedDataSignature, _, err := NewSignature(typedData, signer)
	if err != nil {

		return
	}

	domainSeparatorHash, err := NewDomainSeparatorHash(typedData)
	if err != nil {

		return
	}

	genericParams, err := forwarder.GENERICPARAMS(&bind.CallOpts{})
	if err != nil {

		return
	}

	reqTypeHash, err := NewRequestTypeHash(genericParams)
	if err != nil {

		return
	}

	return gsnCaller.Execute(chainId, *req, domainSeparatorHash, reqTypeHash, nil, typedDataSignature)
}
