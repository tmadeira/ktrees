package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/tmadeira/tcc/codec"
	"github.com/tmadeira/tcc/ktree"
)

func main() {
	log.Println("Welcome!")
	log.Println("This program expects input in the following format:")
	log.Println("")
	log.Println("===")
	log.Println("n")
	log.Println("k")
	log.Println("m")
	log.Println("x_1 y_1")
	log.Println("...")
	log.Println("x_m y_m")
	log.Println("===")
	log.Println("")
	log.Println("Where:")
	log.Println("- (x_i, y_i) correspond to an edge in the k-tree.")
	log.Println("- Nodes must be 0-indexed.")
	log.Println("")

	var Tk ktree.Ktree
	var n, m int

	fmt.Scanf("%d", &n)
	Tk.Adj = make([][]int, n)

	fmt.Scanf("%d", &Tk.K)

	fmt.Scanf("%d", &m)

	for i := 0; i < m; i++ {
		var x, y int
		fmt.Scanf("%d %d", &x, &y)
		Tk.Adj[x] = append(Tk.Adj[x], y)
		Tk.Adj[y] = append(Tk.Adj[y], x)
	}

	for i := 0; i < n; i++ {
		sort.Ints(Tk.Adj[i])
	}

	C, err := codec.CodingAlgorithm(&Tk)
	if err != nil {
		log.Printf("An error occurred: %v\n", err)
		return
	}

	fmt.Printf("Q = %v\n", C.Q)
	fmt.Printf("S = [")
	for i := 0; i < len(C.S.P); i++ {
		if i != 0 {
			fmt.Printf(", ")
		}
		fmt.Printf("(%d, %d)", C.S.P[i], C.S.L[i])
	}
	fmt.Println("]")
}
