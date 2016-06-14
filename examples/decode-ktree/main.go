package main

import (
	"fmt"
	"sort"

	"github.com/tmadeira/tcc/codec"
	"github.com/tmadeira/tcc/dandelion"
)

func main() {
	fmt.Println("Welcome!")
	fmt.Println("This program expects input in the following format:")
	fmt.Println("")
	fmt.Println("===")
	fmt.Println("k")
	fmt.Println("Q_1")
	fmt.Println("...")
	fmt.Println("Q_k")
	fmt.Println("s")
	fmt.Println("p_1 l_1")
	fmt.Println("...")
	fmt.Println("p_s l_s")
	fmt.Println("===")
	fmt.Println("")
	fmt.Println("Where:")
	fmt.Println("- Q_i correspond to the values in Q.")
	fmt.Println("- (p_i, l_i) correspond to the pairs in the code S.")
	fmt.Println("")

	var C codec.Code

	var k int
	fmt.Scanf("%d", &k)
	C.Q = make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Scanf("%d", &C.Q[i])
	}
	sort.Ints(C.Q)

	var s int
	fmt.Scanf("%d", &s)
	C.S = &dandelion.DandelionCode{make([]int, s), make([]int, s)}
	for i := 0; i < s; i++ {
		fmt.Scanf("%d %d", &C.S.P[i], &C.S.L[i])
	}

	Tk, err := codec.DecodingAlgorithm(&C)
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		return
	}

	fmt.Println("")
	fmt.Println("OUTPUT")
	fmt.Println("======")
	fmt.Printf("n = %d\n", len(Tk.Adj))
	fmt.Printf("k = %d\n", Tk.K)
	fmt.Printf("edges = [")
	for u := 0; u < len(Tk.Adj); u++ {
		for i := 0; i < len(Tk.Adj[u]); i++ {
			v := Tk.Adj[u][i]
			if v > u {
				if u != 0 || i != 0 {
					fmt.Printf(", ")
				}
				fmt.Printf("(%d, %d)", u, v)
			}
		}
	}
	fmt.Println("]")
}
