package main

import (
	"fmt"

	"github.com/tmadeira/tcc/codec"
	"github.com/tmadeira/tcc/ktree"
)

func main() {
	fmt.Println("Welcome!")
	fmt.Println("This program expects input in the following format:")
	fmt.Println("")
	fmt.Println("===")
	fmt.Println("n")
	fmt.Println("k")
	fmt.Println("m")
	fmt.Println("x_1 y_1")
	fmt.Println("...")
	fmt.Println("x_m y_m")
	fmt.Println("===")
	fmt.Println("")
	fmt.Println("Where:")
	fmt.Println("- (x_i, y_i) correspond to an edge in the k-tree.")
	fmt.Println("- Nodes must be 0-indexed.")
	fmt.Println("")

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

	C, err := codec.CodingAlgorithm(&Tk)
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		return
	}

	fmt.Println("")
	fmt.Println("OUTPUT")
	fmt.Println("======")
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
