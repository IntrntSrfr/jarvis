package jarvis

import (
	"fmt"
	"math"
)

type Network struct {
	Inputs    int
	Hidden1   int
	Hidden2   int
	Output    int
	WeightsIH Matrix
	WeightsHH Matrix
	WeightsHO Matrix
	LR        float64
}

func NewNetwork(in, hid, out int, LR float64) *Network {
	return &Network{
		Inputs:    in,
		Hidden1:   hid,
		Hidden2:   hid,
		Output:    out,
		LR:        LR,
		WeightsIH: NewRandomMatrix(hid+1, in+1),
		WeightsHH: NewRandomMatrix(hid+1, hid+1),
		WeightsHO: NewRandomMatrix(out, hid+1),
	}
}

func (n *Network) String() string {
	return fmt.Sprintf("NETWORK - inputs: %v, hidden: %v, outputs: %v, lr: %v\nIH Weights: \n%v\nHO Weights: \n%v", n.Inputs, n.Hidden1, n.Output, n.LR, n.WeightsIH, n.WeightsHO)
}

// Guess takes an input and feeds it through the network
func (n *Network) Guess(inp Matrix) []float64 {
	inp = append(inp, Vector{1})

	hidden1Result := MatrixDot(n.WeightsIH, inp)
	hidden1Activations := MatrixMap(hidden1Result, Sigmoid)

	hidden2Result := MatrixDot(n.WeightsHH, hidden1Activations)
	hidden2Activations := MatrixMap(hidden2Result, Sigmoid)

	outputResult := MatrixDot(n.WeightsHO, hidden2Activations)
	outputActivations := MatrixMap(outputResult, Sigmoid)

	return MatrixTranspose(outputActivations)[0]
}

func (n *Network) Train(inp, target Matrix) float64 {
	// this can probably be changed so that we train a whole dataset, but for now only 1 input
	// or we can use this method in another method to deal with a whole dataset, in that case we just need to return
	// the mseErrors this input gives so we can inform about the average error of a dataset

	inp = append(inp, Vector{1})

	hidden1Result := MatrixDot(n.WeightsIH, inp)
	hidden1Activations := MatrixMap(hidden1Result, Sigmoid)

	hidden2Result := MatrixDot(n.WeightsHH, hidden1Activations)
	hidden2Activations := MatrixMap(hidden2Result, Sigmoid)

	outputResult := MatrixDot(n.WeightsHO, hidden2Activations)
	outputActivations := MatrixMap(outputResult, Sigmoid)

	totalError := 0.0
	for i := range outputActivations {
		outputError := 0.5 * math.Pow(target[i][0]-outputActivations[i][0], 2)
		totalError += outputError
	}

	errors := MatrixSub(target, outputActivations)
	errors = MatrixScale(errors, -1)

	outputSubtracted := subtractMatrix(outputActivations)
	hidden2Subtracted := subtractMatrix(hidden2Activations)
	hidden1Subtracted := subtractMatrix(hidden1Activations)

	hidden2Errors := MatrixDot(MatrixTranspose(n.WeightsHO), errors)
	hidden1Errors := MatrixDot(MatrixTranspose(n.WeightsHH), hidden2Errors)

	updatedHO := MatrixScale(MatrixDot(MatrixMultiply(errors, MatrixMultiply(outputActivations, outputSubtracted)), MatrixTranspose(hidden2Activations)), n.LR)
	updatedHH := MatrixScale(MatrixDot(MatrixMultiply(hidden2Errors, MatrixMultiply(hidden2Activations, hidden2Subtracted)), MatrixTranspose(hidden1Activations)), n.LR)
	updatedIH := MatrixScale(MatrixDot(MatrixMultiply(hidden1Errors, MatrixMultiply(hidden1Activations, hidden1Subtracted)), MatrixTranspose(inp)), n.LR)

	n.WeightsIH = MatrixSub(n.WeightsIH, updatedIH)
	n.WeightsHH = MatrixSub(n.WeightsHH, updatedHH)
	n.WeightsHO = MatrixSub(n.WeightsHO, updatedHO)
	return totalError
}

func subtractMatrix(m Matrix) Matrix {
	res := Matrix{}

	for i := range m {
		var v []float64
		for j := range m[0] {
			v = append(v, 1-m[i][j])
		}
		res = append(res, v)
	}

	return res
}
