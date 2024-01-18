package metaresolver

import (
	"bytes"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"go-delegator/log"
)

func verifySignature(reqID uint64, hash []byte, v uint8, r hexutil.Bytes, s hexutil.Bytes, address common.Address) Error {
	isValid, err := isValidSignature(reqID, hash, v, r, s, address)
	log.Debugfd(reqID, "isValid : %v\n", isValid)
	if err != nil {
		return err
	}

	isValidWithPrefix, err := isValidSignatureWithPrefix(reqID, hash, v, r, s, address)
	log.Debugfd(reqID, "isValidWithPrefix : %v\n", isValidWithPrefix)
	if err != nil {
		return err
	}
	if isValid || isValidWithPrefix {
		return nil
	}
	return &invalidSignatureError{"Failed to verify signature"}

}
func isValidSignatureWithPrefix(reqID uint64, hash []byte, v uint8, r hexutil.Bytes, s hexutil.Bytes, address common.Address) (bool, Error) {
	phash := signHash(hash)

	return isValidSignature(reqID, phash, v, r, s, address)
}
func isValidSignature(reqID uint64, hash []byte, v uint8, r hexutil.Bytes, s hexutil.Bytes, address common.Address) (bool, Error) {
	var sig hexutil.Bytes

	sig = append(sig, r...)
	sig = append(sig, s...)
	sig = append(sig, v-27)
	log.Debugfd(reqID, "signature : '%v' \n", sig)
	var signedAddress common.Address

	pKeyBytes, err := ethCrypto.Ecrecover(hash, sig)
	if err != nil {
		log.Errorfd(reqID, "EcRecover Error: %v", err)
		return false, &internalError{"Failed to EcRecover"}
	}
	pKey, err := ethCrypto.UnmarshalPubkey(pKeyBytes)
	if err != nil {
		log.Errorfd(reqID, "EcRecover Error: %v", err)
		return false, &internalError{"Failed to EcRecover"}
	}
	signedAddress = ethCrypto.PubkeyToAddress(*pKey)

	log.Debugfd(reqID, "signedAddress : %v\n", signedAddress.String())
	if !bytes.Equal(signedAddress.Bytes(), address.Bytes()) {
		log.Debugfd(reqID, "not same address %v / %v ", address.String(), signedAddress.String())
		return false, nil
	}
	return true, nil
}

func addressArrayToBytes(addresses []common.Address) []byte {
	var ret []byte
	if len(addresses) == 0 {
		return ret
	}
	for _, addr := range addresses {
		addrBytes := addressToByte32(addr)
		ret = append(ret, addrBytes[:]...)
	}
	return ret
}
func addressToByte32(address common.Address) [32]byte {
	var result [32]byte
	copy(result[12:], address.Bytes())
	return result
}
func bigIntToByte32(_val *big.Int) [32]byte {
	result := [32]byte{}
	if _val != nil {
		_bytes := _val.Bytes()
		copy(result[32-len(_bytes):], _bytes)
	}
	return result
}
func signBytes(bmsg []byte, privKey *ecdsa.PrivateKey) ([]byte, error) {
	bMsg := ethCrypto.Keccak256(bmsg)
	return ethCrypto.Sign(signHash(bMsg), privKey)
}
func signHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return ethCrypto.Keccak256([]byte(msg))
}
func intToByte32(_val *big.Int) [32]byte {
	result := [32]byte{}
	if _val != nil {
		_bytes := _val.Bytes()
		copy(result[32-len(_bytes):], _bytes)
	}
	return result
}
