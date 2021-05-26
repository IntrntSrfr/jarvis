package main

import (
	"fmt"
	"github.com/intrntsrfr/jarvis"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	dataset := []data{
		{
			labels: jarvis.Matrix{{0}},
			data:   jarvis.Matrix{{0}, {0}},
		},
		{
			labels: jarvis.Matrix{{1}},
			data:   jarvis.Matrix{{0}, {1}},
		},
		{
			labels: jarvis.Matrix{{1}},
			data:   jarvis.Matrix{{1}, {0}},
		},
		{
			labels: jarvis.Matrix{{0}},
			data:   jarvis.Matrix{{1}, {1}},
		},
	}

	fmt.Println(len(dataset))

	n := jarvis.NewNetwork(2, 4, 1, .1)

	for i := 0; ; i++ {
		start := time.Now()
		t := 0.0
		for _, d := range dataset {
			t += n.Train(d.data, d.labels)
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

	fmt.Println(fmt.Sprintf("(0, 0) GUESS: %v; WANT 0", n.Guess(jarvis.Matrix{{0}, {0}})[0][0]))
	fmt.Println(fmt.Sprintf("(0, 1) GUESS: %v; WANT 1", n.Guess(jarvis.Matrix{{0}, {1}})[0][0]))
	fmt.Println(fmt.Sprintf("(1, 0) GUESS: %v; WANT 1", n.Guess(jarvis.Matrix{{1}, {0}})[0][0]))
	fmt.Println(fmt.Sprintf("(1, 1) GUESS: %v; WANT 0", n.Guess(jarvis.Matrix{{1}, {1}})[0][0]))
}

type data struct {
	labels jarvis.Matrix
	data   jarvis.Matrix
}
