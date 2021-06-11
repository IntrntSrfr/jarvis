package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/intrntsrfr/jarvis"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	dataset := []data{
		{
			labels: []float64{0},
			data:   []float64{0, 0},
		},
		{
			labels: []float64{1},
			data:   []float64{0, 1},
		},
		{
			labels: []float64{1},
			data:   []float64{1, 0},
		},
		{
			labels: []float64{0},
			data:   []float64{1, 1},
		},
	}

	fmt.Println(len(dataset))

	n := jarvis.NewNetwork2()
	n.AddLayer(2)
	n.AddLayer(4)
	n.AddLayer(1)
	n.Compile()

	for i := 0; ; i++ {
		start := time.Now()
		t := 0.0
		for _, d := range dataset {
			t += n.Train2(d.data, d.labels, 0.01)
		}
		if i%10 == 0 {
			fmt.Println(fmt.Sprintf("> EPOCH %v - TOTAL ERROR: %v - TOTAL AVERAGE ERROR: %v - TIME TAKEN: %v", i, t, t/float64(len(dataset)), time.Since(start)))
		}
		if t/float64(len(dataset)) < 0.0001 {
			fmt.Println(fmt.Sprintf("> EPOCH %v - TOTAL ERROR: %v - TOTAL AVERAGE ERROR: %v - TIME TAKEN: %v", i, t, t/float64(len(dataset)), time.Since(start)))
			break
		}
	}
	/*
		dt, _ := json.MarshalIndent(n, "", "\t")
		fmt.Println(string(dt))*/

	fmt.Println(fmt.Sprintf("(0, 0) GUESS: %v; WANT 0", n.Guess2([]float64{0, 0})[0]))
	fmt.Println(fmt.Sprintf("(0, 1) GUESS: %v; WANT 1", n.Guess2([]float64{0, 1})[0]))
	fmt.Println(fmt.Sprintf("(1, 0) GUESS: %v; WANT 1", n.Guess2([]float64{1, 0})[0]))
	fmt.Println(fmt.Sprintf("(1, 1) GUESS: %v; WANT 0", n.Guess2([]float64{1, 1})[0]))
}

type data struct {
	labels []float64
	data   []float64
}
