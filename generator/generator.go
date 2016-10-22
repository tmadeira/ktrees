// Package generator implements a random k-tree generator.
package generator

import (
	"errors"
	"log"
	"math/rand"
	"sort"
	"time"

	"github.com/tmadeira/tcc/characteristic"
	"github.com/tmadeira/tcc/codec"
	"github.com/tmadeira/tcc/dandelion"
	"github.com/tmadeira/tcc/ktree"
)

var seed = func() int64 {
	return time.Now().UnixNano()
}

// RandomKtree returns a random k-tree with n nodes.
func RandomKtree(n, k int) (*ktree.Ktree, error) {
	C, err := randomCode(n, k)
	if err != nil {
		return nil, err
	}

	log.Printf("C = {%v, %v}", C.Q, C.S)

	Tk, err := codec.DecodingAlgorithm(C)
	if err != nil {
		return nil, err
	}

	return Tk, nil
}

// randomCode returns a random code for a k-tree with n nodes.
func randomCode(n, k int) (*codec.Code, error) {
	if n-2 < k {
		// For more information, see Remark 1 in Caminiti et al.
		return nil, errors.New("This code requires n >= k+2.")
	}

	if k < 1 {
		return nil, errors.New("This code requires k > 0.")
	}

	rand.Seed(seed())

	C := &codec.Code{
		rand.Perm(n)[:k],
		&dandelion.DandelionCode{
			make([]int, n-k-2),
			make([]int, n-k-2),
		},
	}

	sort.Ints(C.Q)

	for i := 0; i < n-k-2; i++ {
		r := rand.Intn((n-k)*k + 1)
		if r == 0 {
			C.S.P[i] = 0
			C.S.L[i] = characteristic.E
		} else {
			r--
			C.S.P[i] = 1 + r/k
			C.S.L[i] = r % k
		}
	}

	return C, nil
}
