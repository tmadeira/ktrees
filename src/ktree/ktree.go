// Package ktree implements k-trees.
package ktree

import (
	"errors"
	"sort"
)

type Ktree struct {
	// Keep the adjacency lists sorted.
	Adj [][]int
	K   int
}

type RenyiKtree struct {
	Ktree *Ktree
	Q     []int
}

// GetQ returns Q, the k-clique adjacent to the maximum labeled leaf of T_k.
func GetQ(t *Ktree) ([]int, error) {
	// Compute degree of each node v and find lm: maximum v such that d(v) = k.
	lm := -1
	for v := 0; v < len(t.Adj); v++ {
		if len(t.Adj[v]) == t.K {
			lm = v
		}
	}

	if lm == -1 {
		return nil, errors.New("Can't find v with d(v) = k.")
	}

	Q := t.Adj[lm]
	sort.Ints(Q)
	return Q, nil
}

// ComputePhi returns a vector phi from Q (see Program 4 in Caminiti et al).
func ComputePhi(n, k int, Q []int) []int {
	const unassigned = -1

	phi := make([]int, n)
	for i := 0; i < n; i++ {
		phi[i] = unassigned
	}

	for i := 0; i < len(Q); i++ {
		phi[Q[i]] = n - k + i
	}

	for i := 0; i < n-k; i++ {
		j := i
		for phi[j] != unassigned {
			j = phi[j]
		}
		phi[j] = i
	}

	return phi
}

// Relabel receives a k-Tree and relabels it by phi.
func Relabel(old *Ktree, phi []int) *Ktree {
	new := &Ktree{make([][]int, len(old.Adj)), old.K}

	for u := 0; u < len(old.Adj); u++ {
		for i := 0; i < len(old.Adj[u]); i++ {
			v := old.Adj[u][i]
			new.Adj[phi[u]] = append(new.Adj[phi[u]], phi[v])
		}
		sort.Ints(new.Adj[phi[u]])
	}

	return new
}

// RkFrom returns a Renyi k-Tree from a given k-Tree.
// See Step 1 from Coding Algorithm in Section 5 of Caminiti et al.
func RkFrom(t *Ktree) (*RenyiKtree, error) {
	Q, err := GetQ(t)
	if err != nil {
		return nil, err
	}

	phi := ComputePhi(len(t.Adj), t.K, Q)

	return &RenyiKtree{Relabel(t, phi), Q}, nil
}

// TODO: Implement Step 4 from Decoding Algorithm (R_k -> k-tree).
