package metaresolver

import (
	"bytes"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/metadium/go-delegator/json"
	"github.com/metadium/go-delegator/log"
	"github.com/metadium/go-delegator/metaresolver/sc/identityregistry"
	"github.com/metadium/go-delegator/metaresolver/sc/servicekeyresolver"
)

func createIdentity(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {

	log.Debugd(reqID, " Call createIdentity Function")

	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	// var providers [1]common.Address
	// delegatorAddr, _ := getDelegatorAddress()
	// providers[0] = *delegatorAddr
	reqParam := tmpParams.(identityCreateParams)

	log.Debugfd(reqID, "parameter[RecoveryAddress] : %x \n", reqParam.RecoveryAddress)
	log.Debugfd(reqID, "parameter[AssociatedAddress] : %x \n", reqParam.AssociatedAddress)
	log.Debugfd(reqID, "parameter[Providers] : %x \n", reqParam.Providers)
	log.Debugfd(reqID, "parameter[Resolvers] : %v \n", reqParam.Resolvers)
	log.Debugfd(reqID, "parameter[V] : %v", reqParam.V)
	log.Debugfd(reqID, "parameter[R] : %v", reqParam.R)
	log.Debugfd(reqID, "parameter[S] : %v", reqParam.S)
	log.Debugfd(reqID, "parameter[Timestamp] : %v", reqParam.Timestamp)

	log.Debugd(reqID, "PASS - 01. Check Parameter \n")
	// // 2. verify signature
	// identityRegistryAddr, _ := identityregistry.GetAddress()
	// keccakBytes, _ := reqParam.Keccak256(identityRegistryAddr)
	// errObj = verifySignature(reqID, keccakBytes, reqParam.V[0], reqParam.R, reqParam.S, reqParam.AssociatedAddress)
	// if errObj != nil {
	// 	resp.Error = makeErrorResponse(errObj)
	// 	return
	// }

	// 3. CallCreateIdentity
	var rBytes, sBytes [32]byte
	copy(rBytes[:], reqParam.R)
	copy(sBytes[:], reqParam.S)
	trx, err := identityregistry.CallCreateIdentity(reqID, reqParam.RecoveryAddress, reqParam.AssociatedAddress, reqParam.Providers, reqParam.Resolvers, reqParam.V[0], rBytes, sBytes, reqParam.Timestamp)
	if err != nil {
		log.Errorfd(reqID, "CallCreateIdentity Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugfd(reqID, "PASS - Call CreateIdentity : %v", trx.Hash().String())

	//  return txid
	resp.Result = trx.Hash().String()

	return
}

func addAssociatedAddressDelegated(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	log.Debugd(reqID, " Call addAssociatedAddressDelegated Function")

	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	reqParam := tmpParams.(addAssociatedAddressDelegatedParams)

	log.Debugfd(reqID, "parameter[ApprovingAddress] : %x \n", reqParam.ApprovingAddress)
	log.Debugfd(reqID, "parameter[AddressToAdd] : %x \n", reqParam.AddressToAdd)
	log.Debugfd(reqID, "parameter[V] : %v", reqParam.V)
	log.Debugfd(reqID, "parameter[R] : %v", reqParam.R)
	log.Debugfd(reqID, "parameter[S] : %v", reqParam.S)
	log.Debugfd(reqID, "parameter[Timestamp] : %v", reqParam.Timestamp)

	log.Debugd(reqID, "PASS - 01. Check Parameter \n")
	var vBytes [2]byte
	vBytes[0] = reqParam.V[0][0]
	vBytes[1] = reqParam.V[1][0]
	var rBytes, sBytes [2][32]byte
	copy(rBytes[0][:], reqParam.R[0])
	copy(rBytes[1][:], reqParam.R[1])
	copy(sBytes[0][:], reqParam.S[0])
	copy(sBytes[1][:], reqParam.S[1])

	trx, err := identityregistry.CallAddAssociatedAddressDelegated(reqID, reqParam.ApprovingAddress, reqParam.AddressToAdd, vBytes, rBytes, sBytes, reqParam.Timestamp)
	if err != nil {
		log.Errorfd(reqID, "CallAddAssociatedAddressDelegated Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugfd(reqID, "PASS - Call AddAssociatedAddressDelegated : %v", trx.Hash().String())

	//  return txid
	resp.Result = trx.Hash().String()

	return
}

func removeAssociatedAddressDelegated(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	log.Debugd(reqID, " Call removeAssociatedAddressDelegated Function")

	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	reqParam := tmpParams.(removeAssociatedAddressDelegatedParams)

	log.Debugfd(reqID, "parameter[AddressToRemove] : %x \n", reqParam.AddressToRemove)
	log.Debugfd(reqID, "parameter[V] : %v", reqParam.V)
	log.Debugfd(reqID, "parameter[R] : %v", reqParam.R)
	log.Debugfd(reqID, "parameter[S] : %v", reqParam.S)
	log.Debugfd(reqID, "parameter[Timestamp] : %v", reqParam.Timestamp)

	log.Debugd(reqID, "PASS - 01. Check Parameter \n")

	var rBytes, sBytes [32]byte
	copy(rBytes[:], reqParam.R)
	copy(sBytes[:], reqParam.S)
	trx, err := identityregistry.CallRemoveAssociatedAddressDelegated(reqID, reqParam.AddressToRemove, reqParam.V[0], rBytes, sBytes, reqParam.Timestamp)
	if err != nil {
		log.Errorfd(reqID, "CallRemoveAssociatedAddressDelegated Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugfd(reqID, "PASS - Call RemoveAssociatedAddressDelegated : %v", trx.Hash().String())

	//  return txid
	resp.Result = trx.Hash().String()

	return
}

func addKeyDelegated(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {

	log.Debugd(reqID, " Call addKeyDelegated Function")

	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc
	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	reqParam := tmpParams.(addKeyDelegatedParams)
	log.Debugfd(reqID, "parameter[ResolverAddress] : %x", reqParam.ResolverAddress)
	log.Debugfd(reqID, "parameter[AssociatedAddress] : %x", reqParam.AssociatedAddress)
	log.Debugfd(reqID, "parameter[Key] : %x", reqParam.Key)
	log.Debugfd(reqID, "parameter[Symbol] : %v", reqParam.Symbol)
	log.Debugfd(reqID, "parameter[V] : %v", reqParam.V)
	log.Debugfd(reqID, "parameter[R] : %v", reqParam.R)
	log.Debugfd(reqID, "parameter[S] : %v", reqParam.S)
	log.Debugfd(reqID, "parameter[Timestamp] : %v", reqParam.Timestamp)
	log.Debugd(reqID, "PASS - 01. Check Parameter")

	//2. get instance ServiceKeyResolver
	instance, err := servicekeyresolver.GetInstance(reqParam.ResolverAddress)
	if err != nil {
		errObj := &notExistsAddressError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	if instance == nil {
		err = fmt.Errorf("Cannot get ServiceKeyResolver Instance")
		errObj := &notExistsAddressError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(reqID, "PASS - 02. Get ServiceKeyResolver")

	// 3. CallAddKeyDelegated
	var rBytes, sBytes [32]byte
	copy(rBytes[:], reqParam.R)
	copy(sBytes[:], reqParam.S)
	trx, err := servicekeyresolver.CallAddKeyDelegated(reqID, instance, reqParam.AssociatedAddress, reqParam.Key, reqParam.Symbol, reqParam.V[0], rBytes, sBytes, reqParam.Timestamp)
	if err != nil {
		log.Errorfd(reqID, "CallAddKeyDelegated Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugfd(reqID, "PASS 05. Call addKeyDelegated : %v", trx.Hash().String())

	//  return txid
	resp.Result = trx.Hash().String()
	return
}

func removeKeyDelegated(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	log.Debugd(reqID, " Call removeKeyDelegated Function")

	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc
	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	reqParam := tmpParams.(removeKeyDelegatedParams)
	log.Debugfd(reqID, "parameter[ResolverAddress] : %x", reqParam.ResolverAddress)
	log.Debugfd(reqID, "parameter[AssociatedAddress] : %x", reqParam.AssociatedAddress)
	log.Debugfd(reqID, "parameter[Key] : %x", reqParam.Key)
	log.Debugfd(reqID, "parameter[V] : %v", reqParam.V)
	log.Debugfd(reqID, "parameter[R] : %v", reqParam.R)
	log.Debugfd(reqID, "parameter[S] : %v", reqParam.S)
	log.Debugfd(reqID, "parameter[Timestamp] : %v", reqParam.Timestamp)
	log.Debugd(reqID, "PASS - 01. Check Parameter")

	//2. get instance ServiceKeyResolver
	instance, err := servicekeyresolver.GetInstance(reqParam.ResolverAddress)
	if err != nil {
		errObj := &notExistsAddressError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	if instance == nil {
		err = fmt.Errorf("Cannot get ServiceKeyResolver Instance")
		errObj := &notExistsAddressError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(reqID, "PASS - 02. Get ServiceKeyResolver")

	// 3. CallDelegatedApprove
	var rBytes, sBytes [32]byte
	copy(rBytes[:], reqParam.R)
	copy(sBytes[:], reqParam.S)
	trx, err := servicekeyresolver.CallRemoveKeyDelegated(reqID, instance, reqParam.AssociatedAddress, reqParam.Key, reqParam.V[0], rBytes, sBytes, reqParam.Timestamp)
	if err != nil {
		log.Errorfd(reqID, "CallRemoveKeyDelegated Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugfd(reqID, "PASS 05. Call RemoveKeyDelegated : %v", trx.Hash().String())

	//return txid
	resp.Result = trx.Hash().String()

	return
}

func removeKeysDelegated(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	log.Debugd(reqID, " Call removeKeysDelegated Function")

	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc
	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	reqParam := tmpParams.(removeKeysDelegatedParams)
	log.Debugfd(reqID, "parameter[ResolverAddress] : %x", reqParam.ResolverAddress)
	log.Debugfd(reqID, "parameter[AssociatedAddress] : %x", reqParam.AssociatedAddress)
	log.Debugfd(reqID, "parameter[V] : %v", reqParam.V)
	log.Debugfd(reqID, "parameter[R] : %v", reqParam.R)
	log.Debugfd(reqID, "parameter[S] : %v", reqParam.S)
	log.Debugfd(reqID, "parameter[Timestamp] : %v", reqParam.Timestamp)
	log.Debugd(reqID, "PASS - 01. Check Parameter")

	//2. get instance ServiceKeyResolver
	instance, err := servicekeyresolver.GetInstance(reqParam.ResolverAddress)
	if err != nil {
		errObj := &notExistsAddressError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}

	if instance == nil {
		err = fmt.Errorf("Cannot get ServiceKeyResolver Instance")
		errObj := &notExistsAddressError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(reqID, "PASS - 02. Get ServiceKeyResolver")

	// 3. CallDelegatedApprove
	var rBytes, sBytes [32]byte
	copy(rBytes[:], reqParam.R)
	copy(sBytes[:], reqParam.S)
	trx, err := servicekeyresolver.CallRemoveKeysDelegated(reqID, instance, reqParam.AssociatedAddress, reqParam.V[0], rBytes, sBytes, reqParam.Timestamp)
	if err != nil {
		log.Errorfd(reqID, "CallRemoveKeysDelegated Error : %v", err)
		errObj := &internalError{err.Error()}
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugfd(reqID, "PASS 05. Call RemoveKeysDelegated : %v", trx.Hash().String())

	//return txid
	resp.Result = trx.Hash().String()

	return
}
func createIdentityTest(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {

	log.Debugd(reqID, " Call createIdentity Function")

	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	var providers [1]common.Address
	delegatorAddr, _ := getDelegatorAddress()
	providers[0] = *delegatorAddr
	reqParam := tmpParams.(identityCreateParams)

	identityRegistryAddr, _ := identityregistry.GetAddress()

	log.Debugfd(reqID, "parameter[RecoveryAddress] : %x \n", reqParam.RecoveryAddress)
	log.Debugfd(reqID, "parameter[AssociatedAddress] : %x \n", reqParam.AssociatedAddress)
	log.Debugfd(reqID, "parameter[Providers] : %x \n", providers)
	log.Debugfd(reqID, "parameter[Resolvers] : %v \n", reqParam.Resolvers)
	log.Debugfd(reqID, "parameter[V] : %v", reqParam.V)
	log.Debugfd(reqID, "parameter[R] : %v", reqParam.R)
	log.Debugfd(reqID, "parameter[S] : %v", reqParam.S)
	log.Debugfd(reqID, "parameter[Timestamp] : %v", reqParam.Timestamp)

	log.Debugd(reqID, "PASS - 01. Check Parameter \n")

	keccakBytes, signData := reqParam.Keccak256(identityRegistryAddr)
	/*
		irAddrBytes := identityRegistryAddr.Bytes()

		providersBytes := addressArrayToBytes(reqParam.Providers)
		resolversBytes := addressArrayToBytes(reqParam.Resolvers)
		timestampBytes := bigIntToByte32(reqParam.Timestamp)

		var signData []byte
		signData = append(signData, headerBytes...)
		signData = append(signData, irAddrBytes...)
		signData = append(signData, []byte("I authorize the creation of an Identity on my behalf.")...)
		signData = append(signData, reqParam.RecoveryAddress.Bytes()...)
		signData = append(signData, reqParam.AssociatedAddress.Bytes()...)
		signData = append(signData, providersBytes...)
		signData = append(signData, resolversBytes...)
		signData = append(signData, timestampBytes[:]...)
		keccakBytes := ethCrypto.Keccak256(signData)
	*/
	log.Debugfd(reqID, "SignData : %x \n", signData)

	log.Debugfd(reqID, "keccakBytes : %x \n", keccakBytes)

	errObj = verifySignature(reqID, keccakBytes, reqParam.V[0], reqParam.R, reqParam.S, reqParam.AssociatedAddress)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}

	//make sig
	// var sig hexutil.Bytes

	// sig = append(sig, reqParam.R...)
	// sig = append(sig, reqParam.S...)
	// sig = append(sig, reqParam.V[0]+27)
	// log.Debugfd(reqID,"signature : %v %v\n", sig, len(sig))
	/*
		prikey, err := ethCrypto.HexToECDSA("01b149603ca8f537bbb4e45d22e77df9054e50d826bb5f0a34e9ce460432b596")
		if err != nil {
			log.Debugfd(reqID,"Failed to get key %s", err)
		}
		var tsig hexutil.Bytes
		tsig, err = signBytes(signData, prikey)
		//tsig, err := ethCrypto.Sign(signHash(signData), prikey)
		if err != nil {
			log.Debugfd(reqID,"Failed to sign %s", err)
		}
		log.Debugfd(reqID,"signature : %v %v\n", tsig, len(tsig))
		rsig := tsig[:32]
		ssig := tsig[32:64]
		v := tsig[64]

		log.Debugfd(reqID,"V,R,S : %v,%v,%v \n", v, rsig, ssig)
	*/
	// var pubkeyBytes, shash hexutil.Bytes
	// shash = signHash(keccakBytes)
	// log.Debugfd(reqID,"shash : %v \n", shash)

	// pubkeyBytes, err := ethCrypto.Ecrecover(shash, sig)
	// if err != nil {
	// 	log.Errorfd(reqID, "-- EcRecover Error: %v", err)
	// 	errObj := &internalError{"Failed to EcRecover"}
	// 	resp.Error = makeErrorResponse(errObj)
	// 	return
	// }
	// log.Debugfd(reqID,"pubkeyBytes : %v \n", pubkeyBytes)
	// pubKey, _ := ethCrypto.UnmarshalPubkey(pubkeyBytes)
	// //var addrBytes common.Address
	// signedAddress := ethCrypto.PubkeyToAddress(*pubKey)
	// log.Debugfd(reqID,"signedByAddress : %x \n", signedAddress)

	// if !bytes.Equal(reqParam.AssociatedAddress.Bytes(), signedAddress.Bytes()) {
	// 	log.Errorfd(reqID, "Error  :Sign Error %v / %v ", reqParam.AssociatedAddress.String(), signedAddress.String())
	// 	errObj := &invalidSignatureError{"Failed to verify signature"}
	// 	resp.Error = makeErrorResponse(errObj)
	// 	return
	// }
	/* sign Test code
	var msg1, msg1Hash, msg1Sig hexutil.Bytes
	msg1 = []byte("test meg")
	msg1Hash = ethCrypto.Keccak256(msg1)
	log.Debugfd(reqID,"msg1Hash : %v \n", msg1Hash)
	msg1Sig, _ = ethCrypto.Sign(msg1Hash, prikey)
	log.Debugfd(reqID,"msg1sign : %v \n", msg1Sig)
	msg1R := msg1Sig[:32]
	msg1S := msg1Sig[32:64]
	msg1V := msg1Sig[64] + 27
	log.Debugfd(reqID,"V,R,S : %v,%v,%v \n", msg1V, msg1R, msg1S)

	msg1pKeyBytes, _ := ethCrypto.Ecrecover(msg1Hash, msg1Sig)
	msg1pKey, _ := ethCrypto.UnmarshalPubkey(msg1pKeyBytes)
	msg1Addr := ethCrypto.PubkeyToAddress(*msg1pKey)
	log.Debugfd(reqID,"msg1address : %x \n", msg1Addr)
	*/
	//ethCrypto.Keccak256(headerBytes,
	// 	identityRegistryAddr.Bytes(),
	// 	[]byte("I authorize the creation of an Identity on my behalf."),
	// 	reqParam.RecoveryAddress.Bytes(),
	// 	reqParam.AssociatedAddress.Bytes(),
	// 	providersBytes, resolversBytes,
	// 	timestampBytes[:])

	//TODO: 2. Verify signature
	// _, errObj = verifySignature(reqID, reqParam.Address.String(), reqParam.Signature.String(), &reqParam.Address)
	// if errObj != nil {
	// 	resp.Error = makeErrorResponse(errObj)
	// 	return
	// }
	/*
		// 3. CallCreateIdentity
		trx, err := identityregistry.CallCreateIdentity(reqParam.RecoveryAddress, reqParam.AssociatedAddress, providers, reqParam.Resolvers, reqParam.V, reqParam.R, reqParam.S, reqParam.Timestamp)
		if err != nil {
			log.Errorfd(reqID, "CallCreateMetaID Error : %v", err)
			errObj := &internalError{err.Error()}
			resp.Error = makeErrorResponse(errObj)
			return
		}
		log.Debugfd(reqID, "PASS - Call CreateMetaID : %v", trx.Hash().String())

		//  return txid
		resp.Result = trx.Hash().String()
	*/
	err := fmt.Errorf("Cannot Call Contract")
	errObj = &internalError{err.Error()}
	resp.Error = makeErrorResponse(errObj)
	return
}
func addAssociatedAddressDelegatedTest(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	log.Debugd(reqID, " Call addAssociatedAddressDelegated Function")

	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	reqParam := tmpParams.(addAssociatedAddressDelegatedParams)

	log.Debugfd(reqID, "parameter[ApprovingAddress] : %x \n", reqParam.ApprovingAddress)
	log.Debugfd(reqID, "parameter[AddressToAdd] : %x \n", reqParam.AddressToAdd)
	log.Debugfd(reqID, "parameter[V] : %v", reqParam.V)
	log.Debugfd(reqID, "parameter[R] : %v", reqParam.R)
	log.Debugfd(reqID, "parameter[S] : %v", reqParam.S)
	log.Debugfd(reqID, "parameter[Timestamp] : %v", reqParam.Timestamp)
	log.Debugd(reqID, "PASS - 01. Check Parameter \n")

	// 2. verify signature
	identityRegistryAddr, _ := identityregistry.GetAddress()
	ein := big.NewInt(4)
	approveHash, approveData, addHash, addData := reqParam.Keccak256(identityRegistryAddr, ein)
	log.Debugfd(reqID, "approveData : %x \n", approveData)
	errObj = verifySignature(reqID, approveHash, reqParam.V[0][0], reqParam.R[0], reqParam.S[0], reqParam.ApprovingAddress)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(reqID, "PASS - 02. approve \n")

	log.Debugfd(reqID, "addData : %x \n", addData)
	errObj = verifySignature(reqID, addHash, reqParam.V[1][0], reqParam.R[1], reqParam.S[1], reqParam.AddressToAdd)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(reqID, "PASS - 02. add \n")
	var vBytes [2]uint8
	vBytes[0] = reqParam.V[0][0]
	vBytes[1] = reqParam.V[1][0]
	var rBytes, sBytes [2][32]byte
	copy(rBytes[0][:], reqParam.R[0])
	copy(rBytes[1][:], reqParam.R[1])
	copy(sBytes[0][:], reqParam.S[0])
	copy(sBytes[1][:], reqParam.S[1])

	err := fmt.Errorf("Cannot Call Contract")
	errObj = &internalError{err.Error()}
	resp.Error = makeErrorResponse(errObj)
	return
}
func removeAssociatedAddressDelegatedTest(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {
	log.Debugd(reqID, " Call removeAssociatedAddressDelegatedTest Function")

	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc

	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	reqParam := tmpParams.(removeAssociatedAddressDelegatedParams)

	log.Debugfd(reqID, "parameter[AddressToAdd] : %x \n", reqParam.AddressToRemove)
	log.Debugfd(reqID, "parameter[V] : %v", reqParam.V)
	log.Debugfd(reqID, "parameter[R] : %v", reqParam.R)
	log.Debugfd(reqID, "parameter[S] : %v", reqParam.S)
	log.Debugfd(reqID, "parameter[Timestamp] : %v", reqParam.Timestamp)
	log.Debugd(reqID, "PASS - 01. Check Parameter \n")

	// 2. verify signature
	identityRegistryAddr, _ := identityregistry.GetAddress()
	ein := big.NewInt(4)
	hash, data := reqParam.Keccak256(identityRegistryAddr, ein)
	log.Debugfd(reqID, "Data : %x \n", data)
	errObj = verifySignature(reqID, hash, reqParam.V[0], reqParam.R, reqParam.S, reqParam.AddressToRemove)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	log.Debugd(reqID, "PASS - 02. verify signature \n")

	err := fmt.Errorf("Cannot Call Contract")
	errObj = &internalError{err.Error()}
	resp.Error = makeErrorResponse(errObj)
	return
}
func addKeyDelegatedTest(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {

	log.Debugd(reqID, " Call addKeyDelegatedTest Function")

	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc
	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	reqParam := tmpParams.(addKeyDelegatedParams)
	log.Debugfd(reqID, "parameter[ResolverAddress] : %x", reqParam.ResolverAddress)
	log.Debugfd(reqID, "parameter[AssociatedAddress] : %x", reqParam.AssociatedAddress)
	log.Debugfd(reqID, "parameter[Key] : %x", reqParam.Key)
	log.Debugfd(reqID, "parameter[Symbol] : %v", reqParam.Symbol)
	log.Debugfd(reqID, "parameter[V] : %v", reqParam.V)
	log.Debugfd(reqID, "parameter[R] : %v", reqParam.R)
	log.Debugfd(reqID, "parameter[S] : %v", reqParam.S)
	log.Debugfd(reqID, "parameter[Timestamp] : %v", reqParam.Timestamp)
	log.Debugd(reqID, "PASS - 01. Check Parameter")
	//2. verify signature
	hash, data := reqParam.Keccak256()
	log.Debugfd(reqID, "SignData : '%x' \n", data)
	log.Debugfd(reqID, "hash : '%x' \n", hash)
	errObj = verifySignature(reqID, hash, reqParam.V[0], reqParam.R, reqParam.S, reqParam.AssociatedAddress)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}

	err := fmt.Errorf("Cannot Call Contract")
	errObj = &internalError{err.Error()}
	resp.Error = makeErrorResponse(errObj)
	return
}
func removeKeyDelegatedTest(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {

	log.Debugd(reqID, " Call removeKeyDelegatedTest Function")

	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc
	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	reqParam := tmpParams.(removeKeyDelegatedParams)
	log.Debugfd(reqID, "parameter[ResolverAddress] : %x", reqParam.ResolverAddress)
	log.Debugfd(reqID, "parameter[AssociatedAddress] : %x", reqParam.AssociatedAddress)
	log.Debugfd(reqID, "parameter[Key] : %x", reqParam.Key)
	log.Debugfd(reqID, "parameter[V] : %v", reqParam.V)
	log.Debugfd(reqID, "parameter[R] : %v", reqParam.R)
	log.Debugfd(reqID, "parameter[S] : %v", reqParam.S)
	log.Debugfd(reqID, "parameter[Timestamp] : %v", reqParam.Timestamp)
	log.Debugd(reqID, "PASS - 01. Check Parameter")
	//2. verify signature
	keccakBytes, _ := reqParam.Keccak256()
	errObj = verifySignature(reqID, keccakBytes, reqParam.V[0], reqParam.R, reqParam.S, reqParam.AssociatedAddress)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}

	err := fmt.Errorf("Cannot Call Contract")
	errObj = &internalError{err.Error()}
	resp.Error = makeErrorResponse(errObj)
	return
}
func removeKeysDelegatedTest(reqID uint64, req json.RPCRequest) (resp json.RPCResponse, errRet error) {

	log.Debugd(reqID, " Call removesKeyDelegatedTest Function")

	resp.ID = req.ID
	resp.Jsonrpc = req.Jsonrpc
	tmpParams, errObj := getParameter(req.Method, req.Params)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}
	reqParam := tmpParams.(removeKeysDelegatedParams)
	log.Debugfd(reqID, "parameter[ResolverAddress] : %x", reqParam.ResolverAddress)
	log.Debugfd(reqID, "parameter[AssociatedAddress] : %x", reqParam.AssociatedAddress)
	log.Debugfd(reqID, "parameter[V] : %v", reqParam.V)
	log.Debugfd(reqID, "parameter[R] : %v", reqParam.R)
	log.Debugfd(reqID, "parameter[S] : %v", reqParam.S)
	log.Debugfd(reqID, "parameter[Timestamp] : %v", reqParam.Timestamp)
	log.Debugd(reqID, "PASS - 01. Check Parameter")
	//2. verify signature
	keccakBytes, _ := reqParam.Keccak256()
	errObj = verifySignature(reqID, keccakBytes, reqParam.V[0], reqParam.R, reqParam.S, reqParam.AssociatedAddress)
	if errObj != nil {
		resp.Error = makeErrorResponse(errObj)
		return
	}

	err := fmt.Errorf("Cannot Call Contract")
	errObj = &internalError{err.Error()}
	resp.Error = makeErrorResponse(errObj)
	return
}
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
