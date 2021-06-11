package main

import (
	"fmt"

	"github.com/intrntsrfr/jarvis"
)

func main() {

	n := jarvis.NewNetwork2()

	n.AddLayer(2)
	n.AddLayer(4)
	n.AddLayer(4)
	n.AddLayer(1)

	n.Compile()

	for i := 0; i < 500; i++ {
		lol := n.Train2([]float64{0, 1}, []float64{1}, 0.1)
		fmt.Println(lol)
	}

}
