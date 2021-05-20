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
	/*
		points := make([]jarvis.Point, 100)
		for i := 0; i < len(points); i++ {
			points[i] = jarvis.NewPoint()
		}
	*/

	nandPoints := []jarvis.Point{
		{X: 0.0, Y: 0.0, Label: 1},
		{X: 0.0, Y: 1.0, Label: 1},
		{X: 1.0, Y: 0.0, Label: 1},
		{X: 1.0, Y: 1.0, Label: 0},
	}

	// make a perceptron of size 2 and learning rate of 0.1
	nandPerceptron := jarvis.NewPerceptron(2, 0.1)
	fmt.Println("NAND:", nandPerceptron)

	nandPerceptron.TrainEpochs(nandPoints, 5000, jarvis.Sigmoid)
	//fmt.Println("NANX EPOCHS TRAINED:", e)

	fmt.Println("NAND:", nandPerceptron)

	fmt.Println("NAND(0, 0):", nandPerceptron.Guess(jarvis.Vector{1, 0, 0}, jarvis.Sigmoid))
	fmt.Println("NAND(0, 1):", nandPerceptron.Guess(jarvis.Vector{1, 0, 1}, jarvis.Sigmoid))
	fmt.Println("NAND(1, 0):", nandPerceptron.Guess(jarvis.Vector{1, 1, 0}, jarvis.Sigmoid))
	fmt.Println("NAND(1, 1):", nandPerceptron.Guess(jarvis.Vector{1, 1, 1}, jarvis.Sigmoid))
	fmt.Println()

	orPoints := []jarvis.Point{
		{X: 0.0, Y: 0.0, Label: 0},
		{X: 0.0, Y: 1.0, Label: 1},
		{X: 1.0, Y: 0.0, Label: 1},
		{X: 1.0, Y: 1.0, Label: 1},
	}

	// make a perceptron of size 2 and learning rate of 0.1
	orPerceptron := jarvis.NewPerceptron(2, 0.1)
	fmt.Println("OR:", orPerceptron)

	orPerceptron.TrainEpochs(orPoints, 5000, jarvis.Sigmoid)

	fmt.Println("OR:", orPerceptron)

	fmt.Println("OR(0, 0):", orPerceptron.Guess(jarvis.Vector{1, 0, 0}, jarvis.Sigmoid))
	fmt.Println("OR(0, 1):", orPerceptron.Guess(jarvis.Vector{1, 0, 1}, jarvis.Sigmoid))
	fmt.Println("OR(1, 0):", orPerceptron.Guess(jarvis.Vector{1, 1, 0}, jarvis.Sigmoid))
	fmt.Println("OR(1, 1):", orPerceptron.Guess(jarvis.Vector{1, 1, 1}, jarvis.Sigmoid))
	fmt.Println()

	/*
		xorPoints := []jarvis.Point{
			{X: nandPerceptron.Guess(jarvis.Vector{1, 0, 0}, nil), Y: orPerceptron.Guess(jarvis.Vector{1, 0, 0}, nil), Label: 0},
			{X: nandPerceptron.Guess(jarvis.Vector{1, 0, 1}, nil), Y: orPerceptron.Guess(jarvis.Vector{1, 0, 1}, nil), Label: 1},
			{X: nandPerceptron.Guess(jarvis.Vector{1, 1, 0}, nil), Y: orPerceptron.Guess(jarvis.Vector{1, 1, 0}, nil), Label: 1},
			{X: nandPerceptron.Guess(jarvis.Vector{1, 1, 1}, nil), Y: orPerceptron.Guess(jarvis.Vector{1, 1, 1}, nil), Label: 0},
	*/
	xorPoints := []jarvis.Point{
		{X: 1, Y: 0, Label: 0},
		{X: 1, Y: 1, Label: 1},
		{X: 1, Y: 1, Label: 1},
		{X: 0, Y: 1, Label: 0},
	}

	// make a perceptron of size 2 and learning rate of 0.1
	xorPerceptron := jarvis.NewPerceptron(2, 0.1)
	fmt.Println("XOR:", xorPerceptron)

	xorPerceptron.TrainEpochs(xorPoints, 5000, jarvis.HeavisideStep)

	fmt.Println("XOR:", xorPerceptron)

	fmt.Println("XOR(0, 0):", xorPerceptron.Guess(jarvis.Vector{1, 0, 0}, jarvis.HeavisideStep))
	fmt.Println("XOR(0, 1):", xorPerceptron.Guess(jarvis.Vector{1, 0, 1}, jarvis.HeavisideStep))
	fmt.Println("XOR(1, 0):", xorPerceptron.Guess(jarvis.Vector{1, 1, 0}, jarvis.HeavisideStep))
	fmt.Println("XOR(1, 1):", xorPerceptron.Guess(jarvis.Vector{1, 1, 1}, jarvis.HeavisideStep))

}
