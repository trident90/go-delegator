package merkletree

import (
	"fmt"
	"testing"

	proxyCrypto "bitbucket.org/coinplugin/proxy/crypto"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// //TestContent implements the Content interface provided by merkletree and represents the content stored in the tree.
type TestContent struct {
	x string
}

//CalculateHash hashes the values of a TestContent
func (t TestContent) CalculateHash() []byte {
	// h := sha256.New()
	// h.Write([]byte(t.x))
	// return h.Sum(nil)
	return crypto.Keccak256([]byte(t.x))
}

//Equals tests for equality of two Contents
func (t TestContent) Equals(other Content) bool {
	return t.x == other.(TestContent).x
}

func TestMerkleTree(t *testing.T) {

	//Build list of Content to build tree
	var list []Content

	test := hexutil.MustDecode("0x3ac225168df54212a25c1c01fd35bebfea408fdac2e31ddd6f80a4bbf9a5f1cb")
	fmt.Println("MustDecode: ", test)
	hexTest := hexutil.Bytes(test)
	fmt.Println("MustDecode: ", hexTest)
	list = append(list, HashContent{h: hexutil.MustDecode("0x3ac225168df54212a25c1c01fd35bebfea408fdac2e31ddd6f80a4bbf9a5f1cb")})
	list = append(list, HashContent{h: hexutil.MustDecode("0xb5553de315e0edf504d9150af82dafa5c4667fa618ed0a6f19c69b41166c5510")})
	list = append(list, HashContent{h: hexutil.MustDecode("0x0b42b6393c1f53060fe3ddbfcd7aadcca894465a5a438f69c87d790b2299b9b2")})
	list = append(list, HashContent{h: hexutil.MustDecode("0xf1918e8562236eb17adc8502332f4c9c82bc14e19bfc0aa10ab674ff75b3d2f3")})
	list = append(list, HashContent{h: hexutil.MustDecode("0xa8982c89d80987fb9a510e25981ee9170206be21af3c8e0eb312ef1d3382e761")})

	// list = append(list, TestContent{x: "a"})
	// list = append(list, TestContent{x: "b"})
	// list = append(list, TestContent{x: "c"})
	// list = append(list, TestContent{x: "d"})
	// list = append(list, TestContent{x: "e"})

	//Create a new Merkle Tree from the list of Content
	tree, _ := NewTree(list)

	//Get the Merkle Root of the tree
	mr := tree.MerkleRoot()
	//common.BytesToHash(mr)
	roothash := hexutil.Bytes(mr)
	fmt.Printf("root hash: %x \n", mr)
	fmt.Println(roothash)

	//Verify the entire tree (hashes for each node) is valid
	vt := tree.VerifyTree()
	fmt.Println("Verify Tree: ", vt)

	//Verify a specific content in in the tree
	vc := tree.VerifyContent(tree.MerkleRoot(), list[0])
	fmt.Println("Verify Content: ", vc)

	//String representation
	//fmt.Println(tree)

}

func TestSign(t *testing.T) {
	//privKey := "348a1e57efc055f3b8243f6e78fb18d0508f68d6f2e7cadf92b87e84a0b51024"
	privKey := "01b149603ca8f537bbb4e45d22e77df9054e50d826bb5f0a34e9ce460432b596"
	key, _ := crypto.HexToECDSA(privKey)
	pubKey := key.PublicKey
	fmt.Println("pubkey : ", pubKey)
	sAddress := crypto.PubkeyToAddress(pubKey)
	fmt.Printf("sAddress : %x \n", sAddress)

	bMsg := crypto.Keccak256([]byte("test"))
	msg := hexutil.Encode(bMsg)
	//msg := "test"
	fmt.Println("msg : ", msg)
	s, err := proxyCrypto.Sign("test", privKey)
	if err != nil {
		t.Error("Err", err)
	}
	s[64] += 27
	fmt.Printf("sig : %x\n", s)

	//sig := "0x1d42443d80dfba7a06f32b8496bb0689b2c6963c48ba1e7dfbb8d39976e1629f7ed75dc66c57c278ea44a49e13871592d57eb8f19606e7df1c34032ebf3c86ef1b"
	sig := hexutil.Encode(s)
	s2, err := proxyCrypto.EcRecover(msg, sig)
	if err != nil {
		t.Error("Err", err)
	}

	fmt.Println("sign Address : ", sAddress)

	fmt.Println("ecrecover Address : ", s2)

}
