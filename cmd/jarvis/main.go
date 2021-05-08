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

	points := make([]jarvis.Point, 100)
	for i := 0; i < len(points); i++ {
		points[i] = jarvis.NewPoint()
	}
	fmt.Println(points)

	p := jarvis.NewPerceptron(2, 0.1)
	fmt.Println(p)

	for i := 0; i < 1000; i++ {
		for _, point := range points {
			input := jarvis.Vector{point.X, point.Y}

			p.Train(input, point.Label)

			//fmt.Println()

			//fmt.Println(jarvis.VecDot(v1, p.Weights))

			//guess := p.Guess(input)
			//fmt.Println("GUESS", guess)

			//miss := point.Label - guess
			//fmt.Println("ERROR:", miss)
			//fmt.Println(jarvis.VecAdd(p.Weights, jarvis.VecScale(v1, float64(miss))))

			//p.Weights = jarvis.VecAdd(p.Weights, jarvis.VecScale(input, float64(miss)))
		}
	}
	fmt.Println(p)
}
