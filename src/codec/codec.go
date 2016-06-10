// Package codec implements coding and decoding algorithms.
package codec

import (
	"errors"
	"fmt"

	"characteristic"
	"dandelion"
	"ktree"
)

type Code struct {
	Q []int
	S *dandelion.DandelionCode
}

// CodingAlgorithm receives a k-tree Tk and returns a code (Q, S).
// See Section 5 from Caminiti et al.
func CodingAlgorithm(Tk *ktree.Ktree) (*Code, error) {
	fmt.Println("Coding Algorithm received input %v\n", Tk)

	// Step 1: Identify Q. Transform Tk into Rk.
	fmt.Println("Step 1...")
	Rk, err := ktree.RkFrom(Tk)
	if err != nil {
		return nil, err
	}
	fmt.Println("Rk = %v\n", Rk)

	// Step 2: Generate the characteristic tree T for Rk.
	fmt.Println("Step 2...")
	T := characteristic.TreeFrom(Rk)
	fmt.Println("T = %v\n", T)

	// Identify q = min(v not in Q).
	q := Rk.Q[len(Rk.Q)-1] + 1
	if Rk.Q[0] != 1 {
		q = 1
	} else {
		for i := 0; i+1 < len(Rk.Q); i++ {
			if Rk.Q[i+1] > Rk.Q[i]+1 {
				q = Rk.Q[i] + 1
				break
			}
		}
	}

	// Make x = phi[q].
	phi := ktree.ComputePhi(len(Tk.Adj), Tk.K, Rk.Q)
	x := phi[q]

	// We increased the indices in Step 2. Increase x accordingly.
	x++

	// Step 3: Compute the Generalized Dandelion Code for T. TODO.

	// Step 4: Return the code (Q, S). TODO.
	return nil, errors.New("Not implemented.")
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
