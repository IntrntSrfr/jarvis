package main

import (
	"github.com/intrntsrfr/jarvis"
)

func main() {
	n := jarvis.NewNetwork(2, 2, 2, 0.1)
	//fmt.Println("NETWORK:", n)

	//targets := jarvis.Matrix{{3}, {5}}

	inp := jarvis.Matrix{{0.5}, {0.10}, {1}}
	target := jarvis.Matrix{{0}, {1}}
	n.Train(inp, target)
	for i := 0; i < 1000; i++ {
		n.Train(inp, target)
	}

	/*
		dataset := []struct{
			inp jarvis.Matrix
			target jarvis.Matrix
		}{
			{
				inp:    jarvis.Matrix{{0}, {0}},
				target: jarvis.Matrix{{0}},
			},
			{
				inp:    jarvis.Matrix{{0}, {1}},
				target: jarvis.Matrix{{0}},
			},
			{
				inp:    jarvis.Matrix{{1}, {0}},
				target: jarvis.Matrix{{0}},
			},
			{
				inp:    jarvis.Matrix{{1}, {1}},
				target: jarvis.Matrix{{1}},
			},
		}
		for i:=0;i<100000;i++{
			for _, data := range dataset{
				n.Train(data.inp, data.target)
			}
		}

		fmt.Println(n.Guess(dataset[3].inp))
	*/
	//fmt.Println(n.Guess(inp))

}
