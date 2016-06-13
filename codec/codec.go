// Package codec implements coding and decoding algorithms.
package codec

import (
	"errors"
	"fmt"

	"github.com/tmadeira/tcc/characteristic"
	"github.com/tmadeira/tcc/dandelion"
	"github.com/tmadeira/tcc/ktree"
)

type Code struct {
	Q []int
	S *dandelion.DandelionCode
}

// CodingAlgorithm receives a k-tree Tk and returns a code (Q, S).
// See Section 5 from Caminiti et al.
func CodingAlgorithm(Tk *ktree.Ktree) (*Code, error) {
	fmt.Printf("Coding Algorithm received input %v\n", Tk)

	// Step 1: Identify Q. Transform Tk into Rk.
	fmt.Println("Step 1...")
	Rk, err := ktree.RkFrom(Tk)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Rk = %v\nQ = %v\n", Rk.Ktree, Rk.Q)

	// Step 2: Generate the characteristic tree T for Rk.
	fmt.Println("Step 2...")
	T := characteristic.TreeFrom(Rk)
	fmt.Printf("T = %v\n", T)

	// Identify q = min(v not in Q).
	q := Rk.Q[len(Rk.Q)-1] + 1
	if Rk.Q[0] != 0 {
		q = 0
	} else {
		for i := 0; i+1 < len(Rk.Q); i++ {
			if Rk.Q[i+1] > Rk.Q[i]+1 {
				q = Rk.Q[i] + 1
				break
			}
		}
	}
	fmt.Printf("q = %v\n", q)

	// Make x = phi[q].
	phi := ktree.ComputePhi(len(Tk.Adj), Tk.K, Rk.Q)
	fmt.Printf("phi = %v\n", phi)
	x := phi[q]
	fmt.Printf("x = %v\n", x)

	// We increased the indices in Step 2. Increase x and Q accordingly.
	x++
	Q := Rk.Q
	for i := 0; i < len(Q); i++ {
		Q[i]++
	}

	// Step 3: Compute the Generalized Dandelion Code for T.
	S := dandelion.Code(T, x)
	fmt.Printf("S (with lm) = %v\n", S)

	// Remove the pair corresponding to phi[lm].
	lm, err := ktree.FindLm(Tk)
	if err != nil {
		return nil, err
	}
	// cor is the index of the pair corresponding to phi[lm].
	cor := phi[lm]
	if x < cor {
		cor--
	}
	fmt.Printf("lm = %v; phi[lm] = %v; cor = %v\n", lm, phi[lm], cor)
	S.P = append(S.P[:cor], S.P[cor+1:]...)
	S.L = append(S.L[:cor], S.L[cor+1:]...)

	fmt.Printf("Final S = %v\n", S)

	// Step 4: Return the code (Q, S).
	return &Code{Q, S}, nil
}

// DecodingAlgorithm receives a code (Q, S) and returns a k-tree Tk.
// See Section 6 in Caminiti et al.
func DecodingAlgorithm(code *Code) (*ktree.Ktree, error) {
	fmt.Printf("Decoding Algorithm received input %v\n", code)

	// Step 1: Compute phi starting from Q; find lm and q.
	fmt.Println("Step 1...")
	k := len(code.Q)
	n := len(code.S.P) + k + 1
	phi := ktree.ComputePhi(n, k, code.Q)
	fmt.Printf("phi = %v\n", phi)
	// TODO: Find lm and q.

	// Step 2: Insert the pair (0, e) and decode S to obtain T. TODO.

	// Step 3: Rebuild Rk by visiting T. TODO.

	// Step 4: Apply phi^(-1) to Rk to obtain Tk. TODO.

	// Step 5: Return Tk. TODO.
	return nil, errors.New("Not implemented.")
}
