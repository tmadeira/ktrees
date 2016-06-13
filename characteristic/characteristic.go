// Package characteristic implements the Characteristic Tree.
package characteristic

import (
	"sort"

	"github.com/tmadeira/tcc/ktree"
)

type Tree struct {
	P, L []int
}

const E = -1

// TreeFrom returns a characteristic tree from a given Renyi k-Tree.
// See Step 2 from Coding Algorithm in Section 5 of Caminiti et al.
func TreeFrom(Rk *ktree.RenyiKtree) *Tree {
	K, order := pruneRk(Rk)
	p := addEdges(order, K)
	l := buildLabels(K, p)
	return &Tree{p, l}
}

// pruneRk is implemented as seen in Program 6 in Section 5 of Caminiti et al.
// It returns the constructed K and the pruning order.
func pruneRk(Rk *ktree.RenyiKtree) ([][]int, []int) {
	n, k := len(Rk.Ktree.Adj), Rk.Ktree.K
	m := make([]bool, n)
	d := make([]int, n)
	K := make([][]int, n-k)
	order := make([]int, 0)

	// Fill degrees vector.
	for v := 0; v < n; v++ {
		d[v] = len(Rk.Ktree.Adj[v])
	}

	for v := 0; v < n-k; v++ {
		w := v
		if d[w] == k {
			remove(w, Rk, K, m, d, &order)

			// While exists an unmarked u in adj(w) such that u < v and d(u) = k...
			for i := 0; i < len(K[w]); i++ {
				u := K[w][i]
				if !m[u] && u < v && d[u] == k {
					w = u
					remove(w, Rk, K, m, d, &order)
					i = -1
				}
			}
		}
	}

	return K, order
}

// remove is implemented as seen in Program 6 in Section 5 of Caminiti et al.
func remove(x int, Rk *ktree.RenyiKtree, K [][]int, m []bool, d []int, o *[]int) {
	// Let K[x] be adj[x] without all marked elements.
	K[x] = make([]int, 0)
	for i := 0; i < len(Rk.Ktree.Adj[x]); i++ {
		v := Rk.Ktree.Adj[x][i]
		if !m[v] {
			K[x] = append(K[x], v)
			d[v]--
		}
	}

	// Keep track of the pruning order.
	*o = append(*o, x)

	// Mark x as removed.
	m[x] = true
}

// addEdges is implemented as seen in Program 7 in Section 5 of Caminiti et al.
// It returns a vector p.
func addEdges(order []int, K [][]int) []int {
	k := len(K[0])
	n := len(K) + k

	// p and level will be 1-indexed (relative to K). 0 is the new root.
	p := make([]int, n-k+1)
	level := make([]int, n+1)
	p[0] = E

	for i := len(order) - 1; i >= 0; i-- {
		v := order[i]
		if K[v][0] == n-k {
			// This means K[v] == root.
			p[v+1] = 0
			level[v+1] = 1
		} else {
			// Choose w in K[v] such that level[w] is maximum.
			w := K[v][0]
			for j := 1; j < len(K[v]); j++ {
				if level[K[v][j]+1] > level[w+1] {
					w = K[v][j]
				}
			}

			p[v+1] = w + 1
			level[v+1] = level[w+1] + 1
		}
	}
	return p
}

// buildLabels constructs a label vector from K and p.
// It is used in the end of the Step 2 of the Coding Algorithm.
// Note that indices in p are 1-indexed relative to K. That's why we use i-1
// to translate an index from p to K.
func buildLabels(K [][]int, p []int) []int {
	n := len(p)
	l := make([]int, n)
	for v := 0; v < n; v++ {
		if p[v] <= 0 {
			l[v] = E
		} else {
			kvi := 0
			kpvi := 0
			for kvi < len(K[v-1]) && kpvi < len(K[p[v]-1]) {
				x := K[v-1][kvi]
				y := K[p[v]-1][kpvi]
				if x == y {
					kvi++
					kpvi++
				} else if x < y {
					kvi++
				} else { // if y < x.
					kvi = len(K[v-1])
				}
			}
			l[v] = kpvi
		}
	}
	return l
}

// RenyiKtreeFrom returns a Renyi k-Tree from a given characteristic tree.
// See Step 3 (Prog. 8) from Decoding Algorithm in Section 6 of Caminiti et al.
func RenyiKtreeFrom(n, k int, Q []int, T *Tree) *ktree.RenyiKtree {
	adj := make([][]int, n)

	// Initialize Rk as the k-clique R on [n-k, ..., n)
	for i := n - k; i < n; i++ {
		for j := n - k; j < n; j++ {
			if i != j {
				adj[i] = append(adj[i], j)
			}
		}
	}

	// Create children vector from T.P.
	children := make([][]int, n)
	for i := 0; i < len(T.P); i++ {
		if T.P[i] == -1 {
			continue
		}
		children[T.P[i]] = append(children[T.P[i]], i)
	}

	// Visit T in BFS order, starting with the children of R.
	K := make([][]int, n)
	queue := make([]int, n)
	m := make([]bool, n)
	start, end := 0, 0
	for i := 0; i < len(children[0]); i++ {
		m[children[0][i]] = true
		queue[end] = children[0][i]
		end++
	}
	for start != end {
		v := queue[start]
		start++
		if T.P[v] == 0 {
			for i := n - k + 1; i <= n; i++ {
				K[v] = append(K[v], i)
			}
		} else {
			for i := 0; i < len(K[T.P[v]]); i++ {
				if i != T.L[v] {
					K[v] = append(K[v], K[T.P[v]][i])
				}
			}
			K[v] = append(K[v], T.P[v])
			sort.Ints(K[v])
		}
		for i := 0; i < len(K[v]); i++ {
			u := K[v][i]
			adj[u-1] = append(adj[u-1], v-1)
			adj[v-1] = append(adj[v-1], u-1)
		}
		for i := 0; i < len(children[v]); i++ {
			if !m[children[v][i]] {
				m[children[v][i]] = true
				queue[end] = children[v][i]
				end++
			}
		}
	}

	Rk := &ktree.RenyiKtree{
		&ktree.Ktree{make([][]int, n), k},
		Q,
	}

	// Order adjacency lists in O(nk).
	for u := 0; u < n; u++ {
		for i := 0; i < len(adj[u]); i++ {
			v := adj[u][i]
			Rk.Ktree.Adj[v] = append(Rk.Ktree.Adj[v], u)
		}
	}

	return Rk
}
