package main

import (
	"fmt"
	"github.com/intrntsrfr/jarvis"
)

func main() {
	n := jarvis.NewNetwork(3, 4, 2, 0.1)
	fmt.Println("NETWORK:", n)

	//targets := jarvis.Matrix{{3}, {5}}
	inp := jarvis.Matrix{{3}, {2}, {4}}
	fmt.Println(n.Feed(inp))
}
