// Package renyi implements the Renyi k-tree.
package renyi

import (
	"ktree"
)

// TODO: Implement Step 4 from Decoding Algorithm (R_k -> k-tree).

type RenyiKtree struct {
	Ktree *ktree.Ktree
	Q     []int
}

// FromTk returns a Renyi k-Tree from a given k-Tree.
// See Step 1 from Coding Algorithm in Section 5 of Caminiti et al.
func FromTk(t *ktree.Ktree) (*RenyiKtree, error) {
	Q, err := ktree.GetQ(t)
	if err != nil {
		return nil, err
	}

	phi := ktree.ComputePhi(len(t.Adj), t.K, Q)

	return &RenyiKtree{ktree.Relabel(t, phi), Q}, nil
}
