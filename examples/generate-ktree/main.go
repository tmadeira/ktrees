package main

import (
	"fmt"
	"log"

	"github.com/tmadeira/tcc/generator"
)

func main() {
	log.Println("Welcome!")
	log.Println("This program expects input in the following format:")
	log.Println("")
	log.Println("===")
	log.Println("n k")
	log.Println("===")
	log.Println("")

	var n, k int

	fmt.Scanf("%d %d", &n, &k)

	Tk, err := generator.RandomKtree(n, k)
	if err != nil {
		log.Printf("An error occurred: %v\n", err)
		return
	}

	fmt.Printf("%d\n", len(Tk.Adj))
	fmt.Printf("%d\n", Tk.K)
	m := 0
	for u := 0; u < len(Tk.Adj); u++ {
		for i := 0; i < len(Tk.Adj[u]); i++ {
			v := Tk.Adj[u][i]
			if v > u {
				m++
			}
		}
	}
	fmt.Printf("%d\n", m)
	for u := 0; u < len(Tk.Adj); u++ {
		for i := 0; i < len(Tk.Adj[u]); i++ {
			v := Tk.Adj[u][i]
			if v > u {
				fmt.Printf("%d %d\n", u, v)
			}
		}
	}
}
