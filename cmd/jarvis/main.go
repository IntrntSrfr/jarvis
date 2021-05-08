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

	// train for each point in list 10 times
	for i := 0; i < 1000; i++ {
		for _, point := range points {
			input := jarvis.Vector{point.X, point.Y}
			p.Train(input, point.Label)
		}
	}

	for _, point := range points {
		inp := jarvis.Vector{point.X, point.Y}
		if p.Guess(inp) != point.Label {
			fmt.Println("FAILED GUESS FOR", point)
		}
	}

	fmt.Println(p)

}
