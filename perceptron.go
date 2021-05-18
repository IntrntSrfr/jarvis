package jarvis

import (
	"fmt"
	"math"
)

type Perceptron struct {
	Weights Vector
	Lr      float64
	Bias    float64
}

// NewPerceptron creates a new perceptron with weights according to a specific size input and a learning rate
func NewPerceptron(size int, lr float64) *Perceptron {

	// do size + 1 do account for the bias weight
	return &Perceptron{
		Weights: make(Vector, size+1),
		Lr:      lr,
		Bias:    1,
	}
}

// String returns a string representation of the perceptron
func (p *Perceptron) String() string {
	return fmt.Sprintf("Bias: %v - Learning rate: %v - Weights: %v", p.Bias, p.Lr, p.Weights)
}

// Train trains the perceptron using an input and a target value for said input
func (p *Perceptron) Train(inp Vector, target int) int {
	guess := p.Guess(inp, nil)
	miss := target - guess
	p.Weights = VecAdd(p.Weights, VecScale(inp, float64(miss)*p.Lr))
	return guess
}

type ActivationFunction func(x float64) float64

// The HeavisideStep function is a step function that returns x if x >= 0, and 0 otherwise
var HeavisideStep = func(x float64) float64 {
	if x >= 0 {
		return 1
	}
	return 0
}

// The Sigmoid function returns an output clamped between 0 and 1, 0 if the number is very negative and
// 1 if the number is very positive.
var Sigmoid = func(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

// ReLU is the Rectified Linear Unit formula, which returns x when x >= 0 and 0 if its less than 0
var ReLU = func(x float64) float64 {
	if x >= 0 {
		return x
	}
	return 0
}

// Guess simply asks the perceptron to make a prediction
func (p *Perceptron) Guess(inp Vector, f ActivationFunction) int {
	if VecDot(p.Weights, inp) >= 0 {
		return 1
	} else {
		return 0
	}
}
