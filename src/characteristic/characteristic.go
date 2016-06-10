// Package characteristic implements the Characteristic Tree.
package characteristic

import "ktree"

type Tree struct {
	P, L []int
}

const E = -1

// TreeFrom returns a characteristic tree from a given Renyi k-Tree.
// See Step 2 from Coding Algorithm in Section 5 of Caminiti et al.
func TreeFrom(Rk *ktree.RenyiKtree) *Tree {
	// TODO.
	return nil
}

// RenyiKtreeFrom returns a Renyi k-Tree from a given characteristic tree.
// See Step 3 from Decoding Algorithm in Section 6 of Caminiti et al.
func RenyiKtreeFrom(T *Tree) *ktree.RenyiKtree {
	// TODO.
	return nil
}
