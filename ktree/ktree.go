// Package ktree implements k-trees.
package ktree

import (
	"errors"
)

type Ktree struct {
	// The adjacency lists must be always sorted.
	Adj [][]int
	K   int
}

type RenyiKtree struct {
	Ktree *Ktree
	Q     []int
}

// GetQ returns Q, the k-clique adjacent to the maximum labeled leaf of T_k.
func GetQ(t *Ktree) ([]int, error) {
	lm, err := FindLm(t)
	if err != nil {
		return nil, err
	}

	Q := t.Adj[lm]
	return Q, nil
}

// FindLm computes degree of each node to find lm: maximum v s.t. d(v) = k.
func FindLm(t *Ktree) (int, error) {
	lm := -1
	for v := 0; v < len(t.Adj); v++ {
		if len(t.Adj[v]) == t.K {
			lm = v
		}
	}

	if lm == -1 {
		return lm, errors.New("Can't find v with d(v) = k.")
	}

	return lm, nil
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
	n := len(old.Adj)
	adj := make([][]int, n)
	new := &Ktree{make([][]int, n), old.K}

	for u := 0; u < n; u++ {
		for i := 0; i < len(old.Adj[u]); i++ {
			v := old.Adj[u][i]
			adj[phi[u]] = append(adj[phi[u]], phi[v])
		}
	}

	// Order adjacency lists in O(nk).
	for u := 0; u < n; u++ {
		for i := 0; i < len(adj[u]); i++ {
			v := adj[u][i]
			new.Adj[v] = append(new.Adj[v], u)
		}
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

// TkFrom returns a k-Tree from a given Renyi k-Tree.
// See Step 4 from Decoding Algorithm in Section 6 of Caminiti et al.
func TkFrom(r *RenyiKtree) *Ktree {
	n, k := len(r.Ktree.Adj), r.Ktree.K
	phi := ComputePhi(n, k, r.Q)

	// Construct phi^(-1).
	inv := GetInverse(phi)

	return Relabel(r.Ktree, inv)
}

// GetInverse receives a permutation vector phi and returns its inverse.
func GetInverse(phi []int) []int {
	n := len(phi)
	inv := make([]int, n)
	for i := 0; i < n; i++ {
		inv[phi[i]] = i
	}
	return inv
}
