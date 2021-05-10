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

	// make a list of points
	points := make([]jarvis.Point, 100)
	for i := 0; i < len(points); i++ {
		points[i] = jarvis.NewPoint()
	}

	// make a perceptron of size 2 and learning rate of 0.1
	p := jarvis.NewPerceptron(2, 0.1)
	fmt.Println(p)

	epochs := 0
	for {
		failed := false
		for _, point := range points {
			input := jarvis.Vector{1, point.X, point.Y}
			lol := p.Train(input, point.Label)
			if lol != point.Label {
				failed = true
			}
		}
		if !failed {
			break
		}
		epochs++
	}

	fmt.Println(epochs)

	for _, point := range points {
		inp := jarvis.Vector{1, point.X, point.Y}
		if p.Guess(inp, nil) != point.Label {
			fmt.Println("FAILED GUESS FOR", point)
		}
	}

	fmt.Println(p)

}
