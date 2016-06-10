// Package dandelion implements Generalized Dandelion Coding/Decoding.
package dandelion

import (
	"sort"

	"characteristic"
)

// DandelionCode represents a coded tree (see Section 4 of Caminiti et al).
type DandelionCode struct {
	// P is the parent vector.
	P []int

	// L is the label vector.
	L []int

	// TODO: Use Compact Representation from Section 7 of Caminiti et al.
}

const r = 0

const notProcessed = 0
const inProgress = 1
const processed = 2

// Code receives a characteristic tree and returns its Dandelion code.
// See Program 1: Generalized Dandelion Coding (in Caminiti et al).
func Code(t *characteristic.Tree, x int) *DandelionCode {
	p, l := t.P, t.L

	// Make swaps while p[x] != r.
	for p[x] != r {
		w := r
		for v := p[x]; v != r; v = p[v] {
			if v > w {
				w = v
			}
		}
		l[x], l[w] = l[w], l[x]
		p[x], p[w] = p[w], p[x]
	}

	// Remove r and x from vectors p and l.
	s := &DandelionCode{make([]int, 0), make([]int, 0)}
	for i := 0; i < len(p); i++ {
		if i != r && i != x {
			s.P = append(s.P, p[i])
			s.L = append(s.L, l[i])
		}
	}

	return s
}

// Decode receives a Dandelion code and returns its characteristic tree.
// See Program 2: Generalized Dandelion Decoding (in Caminiti et al).
func Decode(s *DandelionCode, x int) *characteristic.Tree {
	// Construct graph from code.
	n := len(s.P) + 2
	p, l := make([]int, n), make([]int, n)
	p[r] = characteristic.E
	l[r] = characteristic.E
	p[x] = r
	l[x] = characteristic.E
	jump := 0
	for v := 0; v < n; v++ {
		if v == r || v == x {
			jump++
		} else {
			p[v] = s.P[v-jump]
			l[v] = s.L[v-jump]
		}
	}

	// Identify all cycles.
	m := make([]int, 0) // m is the vector of maximal nodes.
	status := make([]int, n)
	for v := 0; v < n; v++ {
		if v != r && v != x {
			analyze(v, p, &status, &m)
		}
	}

	// Make swaps.
	sort.Ints(m)
	for i := 0; i < len(m); i++ {
		l[x], l[m[i]] = l[m[i]], l[x]
		p[x], p[m[i]] = p[m[i]], p[x]
	}

	return &characteristic.Tree{p, l}
}

// analyze is implemented as seen in Program 3 of Caminiti et al.
func analyze(v int, p []int, status, m *[]int) {
	if (*status)[v] == processed || p[v] < 0 {
		return
	}

	(*status)[v] = inProgress
	if (*status)[p[v]] == inProgress {
		// A cycle has been identified.
		maximal_node := v
		for u := p[v]; u != v; u = p[u] {
			if u > maximal_node {
				maximal_node = u
			}
		}
		*m = append(*m, maximal_node)
	} else {
		analyze(p[v], p, status, m)
	}
	(*status)[v] = processed
}
