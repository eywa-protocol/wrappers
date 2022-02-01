package wrappers

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core"
)

const (
	domainType     = "EIP712Domain"
	forwardRequest = "ForwardRequest"
)

var (
	ErrStringTooLong = errors.New("string to long")
	ErrHexTooShort   = errors.New("hex string is shorter than bytes32")
	UseGsnFlag       bool
)

func init() {
	v := os.Getenv("GSN_USE_BRIDGE")
	UseGsnFlag = v != ""
}

type GsnCallOpts struct {
	GsnCaller       GsnCaller
	ChainId         *big.Int
	Signer          *ecdsa.PrivateKey
	ContractAddress common.Address
}

type GsnCaller interface {
	GetForwarder(chainId *big.Int) (*Forwarder, error)
	GetForwarderAddress(chainId *big.Int) (common.Address, error)
	Execute(chainId *big.Int, req IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) (common.Hash, error)
}

func NewForwardRequestTypedData(req *IForwarderForwardRequest, forwarderAddress, abiJson, methodName string, args ...interface{}) (*core.TypedData, error) {

	contractAbi, err := abi.JSON(strings.NewReader(abiJson))
	if err != nil {
		return nil, fmt.Errorf("could not parse ABI: %w", err)
	}

	contractAbiPacked, err := contractAbi.Pack(methodName, args...)
	if err != nil {
		return nil, fmt.Errorf("could not pack abi method: %s on error:  %w", methodName, err)
	}

	req.Data = contractAbiPacked

	return &core.TypedData{
		Types: core.Types{
			domainType: []core.Type{
				{Name: "verifyingContract", Type: "address"},
			},
			forwardRequest: []core.Type{
				{
					Name: "from", Type: "address"},
				{
					Name: "to", Type: "address"},
				{
					Name: "value", Type: "uint256"},
				{
					Name: "gas", Type: "uint256"},
				{
					Name: "nonce", Type: "uint256"},
				{
					Name: "data", Type: "bytes"},
			},
		},
		Domain: core.TypedDataDomain{
			VerifyingContract: forwarderAddress,
		},
		PrimaryType: "ForwardRequest",
		Message: core.TypedDataMessage{
			"from":  req.From.String(),
			"to":    req.To.String(),
			"value": req.Value.String(),
			"gas":   req.Gas.String(),
			"nonce": req.Nonce.String(),
			"data":  contractAbiPacked,
		},
	}, nil
}

func NewSignature(typedData *core.TypedData, signer *ecdsa.PrivateKey) (signature hexutil.Bytes, SignDataRequestHash hexutil.Bytes, err error) {

	if addr, err := common.NewMixedcaseAddressFromString(crypto.PubkeyToAddress(signer.PublicKey).String()); err != nil {

		return nil, nil, err
	} else {

		return SignTypedData(*signer, *addr, *typedData)
	}
}

func NewDomainSeparatorHash(typedData *core.TypedData) ([32]byte, error) {

	if domainSeparator, err := typedData.HashStruct(domainType, typedData.Domain.Map()); err != nil {

		return [32]byte{}, err
	} else {

		return FromHex(domainSeparator.String())
	}

}

func NewRequestTypeHash(genericParams string) ([32]byte, error) {

	forwardRequestType := fmt.Sprintf("%s(%s)", forwardRequest, genericParams)
	reqTypeHash := crypto.Keccak256([]byte(forwardRequestType))

	if reqTypeHashBytes32, err := BytesToBytes32(reqTypeHash); err != nil {

		return [32]byte{}, err
	} else {

		return reqTypeHashBytes32, nil
	}

}

func BytesToBytes32(b []byte) (r [32]byte, err error) {
	if len(b) > 32 {
		err = ErrStringTooLong
	} else {
		copy(r[:], b)
	}
	return
}

func FromHex(s string) ([32]byte, error) {
	var b [32]byte
	n, err := hex.Decode(b[:], []byte(strings.TrimPrefix(s, "0x")))
	if err == nil && n != 32 {
		return b, ErrHexTooShort
	}
	return b, err
}

func SignTypedData(key ecdsa.PrivateKey, addr common.MixedcaseAddress, typedData core.TypedData) (signature hexutil.Bytes, SignDataRequestHash hexutil.Bytes, err error) {

	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return nil, nil, err
	}

	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return nil, nil, err
	}
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	sigHash := crypto.Keccak256(rawData)
	messages, err := typedData.Format()
	if err != nil {
		return nil, nil, err
	}
	req := &core.SignDataRequest{
		ContentType: core.DataTyped.Mime,
		Rawdata:     rawData,
		Messages:    messages,
		Hash:        sigHash,
		Address:     addr}
	signature, err = crypto.Sign(req.Hash, &key)
	if err != nil {
		return nil, nil, err
	}
	signature[64] += 27
	return signature, req.Hash, nil
}

func GsnExecutor(gsnParams *GsnCallOpts, abiSrc, methodName string, args ...interface{}) (txHash common.Hash, err error) {
	__contractABI, err := abi.JSON(strings.NewReader(abiSrc))
	if err != nil {
		return common.Hash{}, fmt.Errorf("could not parse ABI: %w", err)
	}

	__fRequest, err := __contractABI.Pack(methodName, args...)
	if err != nil {
		return
	}

	__forwarder, err := gsnParams.GsnCaller.GetForwarder(gsnParams.ChainId)
	if err != nil {
		return
	}

	__forwarderAddress, err := gsnParams.GsnCaller.GetForwarderAddress(gsnParams.ChainId)
	if err != nil {
		return
	}

	__signerAddress := crypto.PubkeyToAddress(gsnParams.Signer.PublicKey)

	__nonce, err := __forwarder.GetNonce(&bind.CallOpts{}, __signerAddress)
	if err != nil {
		return
	}

	__req := &IForwarderForwardRequest{
		From:  __signerAddress,
		To:    gsnParams.ContractAddress,
		Value: big.NewInt(0),
		Gas:   big.NewInt(1e6),
		Nonce: __nonce,
		Data:  __fRequest,
	}

	__typedData, err := NewForwardRequestTypedData(
		__req,
		__forwarderAddress.String(),
		abiSrc,
		methodName, args...)
	if err != nil {
		return
	}

	__typedDataSignature, _, err := NewSignature(__typedData, gsnParams.Signer)
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

	return gsnParams.GsnCaller.Execute(gsnParams.ChainId, *__req, __domainSeparatorHash, __reqTypeHash, nil, __typedDataSignature)
}

func GsnWrap(directCall func() (common.Hash, error), gsnCall func() (common.Hash, error)) (common.Hash, error) {
	if UseGsnFlag && gsnCall != nil {
		return gsnCall()
	}
	if directCall == nil {
		return common.Hash{}, errors.New("directCall method absent")
	}

	return directCall()
}

// Contract structs

// IForwarderForwardRequest is an auto generated low-level Go binding around an user-defined struct.
type IForwarderForwardRequest struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Gas   *big.Int
	Nonce *big.Int
	Data  []byte
}

// NodeRegistryNode is an auto generated low-level Go binding around an user-defined struct.
type NodeRegistryNode struct {
	Owner     common.Address
	Pool      common.Address
	HostId    string
	BlsPubKey []byte
	NodeId    *big.Int
}

// PortalPermitData is an auto generated low-level Go binding around an user-defined struct.
type PortalPermitData struct {
	V          uint8
	R          [32]byte
	S          [32]byte
	Deadline   *big.Int
	ApproveMax bool
}

// PortalSynthParams is an auto generated low-level Go binding around an user-defined struct.
type PortalSynthParams struct {
	Chain2address  common.Address
	ReceiveSide    common.Address
	OppositeBridge common.Address
	ChainID        *big.Int
}

// SolanaSerializeSolanaAccountMeta is an auto generated low-level Go binding around an user-defined struct.
type SolanaSerializeSolanaAccountMeta struct {
	Pubkey     [32]byte
	IsSigner   bool
	IsWritable bool
}

// SolanaSerializeSolanaStandaloneInstruction is an auto generated low-level Go binding around an user-defined struct.
type SolanaSerializeSolanaStandaloneInstruction struct {
	ProgramId [32]byte
	Accounts  []SolanaSerializeSolanaAccountMeta
	Data      []byte
}

// TestForwardForwardRequest is an auto generated low-level Go binding around an user-defined struct.
type TestForwardForwardRequest struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Gas   *big.Int
	Nonce *big.Int
	Data  []byte
}
