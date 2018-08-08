package merkletree

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

// HashContent Content For metaID
type HashContent struct {
	H hexutil.Bytes
}

// CalculateHash function to calculate hash
func (t HashContent) CalculateHash() []byte {
	return t.H
}

// Equals function to compare other Content
func (t HashContent) Equals(other Content) bool {
	//	return t.H.String() == other.(HashContent).H.String()
	return bytes.Equal(t.H, other.(HashContent).H)
}
