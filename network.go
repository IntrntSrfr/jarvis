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
	Layers    []int
	Weights   []Matrix
}

func NewNetwork(in, hid, out int) *Network {
	return &Network{
		Inputs:    in,
		Hidden1:   hid,
		Hidden2:   hid,
		Output:    out,
		Layers:    []int{},
		Weights:   []Matrix{},
		WeightsIH: NewRandomMatrix(hid+1, in+1),
		WeightsHH: NewRandomMatrix(hid+1, hid+1),
		WeightsHO: NewRandomMatrix(out, hid+1),
	}
}
func NewNetwork2() *Network {
	return &Network{
		Layers:  []int{},
		Weights: []Matrix{},
	}
}

func (n *Network) String() string {
	return fmt.Sprintf("NETWORK\nLayers: %v\nWeights: %v", n.Layers, n.Weights)
}

func (n *Network) AddLayer(nodes int) {
	n.Layers = append(n.Layers, nodes)
}

func (n *Network) Compile() {
	for i := 0; i < len(n.Layers)-1; i++ {
		if i == len(n.Layers)-2 {
			n.Weights = append(n.Weights, NewRandomMatrix(n.Layers[i+1], n.Layers[i]+1))
		} else {
			n.Weights = append(n.Weights, NewRandomMatrix(n.Layers[i+1]+1, n.Layers[i]+1))
		}
	}
}

func (n *Network) Guess2(inp []float64) []float64 {
	inpM := make(Matrix, len(inp))
	for i := range inpM {
		inpM[i] = []float64{inp[i]}
	}
	inpM = append(inpM, []float64{1})

	var outs Matrix
	for i := range n.Weights {
		res := MatrixDot(n.Weights[i], inpM)
		outs = MatrixMap(res, Sigmoid)
		inpM = outs
	}

	return MatrixTranspose(outs)[0]

}

func (n *Network) Train2(inp, target []float64, lr float64) float64 {

	// this can probably be changed so that we train a whole dataset, but for now only 1 input
	// or we can use this method in another method to deal with a whole dataset, in that case we just need to return
	// the mseErrors this input gives so we can inform about the average error of a dataset

	// bias input
	inpM := NewMatrixFromArray(inp)
	inpM = append(inpM, []float64{1})

	//fmt.Println(inpM)

	targetM := NewMatrixFromArray(target)

	outputs := []Matrix{}
	outputs = append(outputs, inpM)

	var activations Matrix
	for i := range n.Weights {
		res := MatrixDot(n.Weights[i], inpM)
		activations = MatrixMap(res, Sigmoid)
		inpM = activations
		outputs = append(outputs, activations)
	}

	errors := MatrixScale(MatrixSub(targetM, outputs[len(outputs)-1]), -1)
	for i := range n.Weights {
		subtracted := subtractMatrix(outputs[len(outputs)-(1+i)])
		updated := MatrixScale(MatrixDot(MatrixMultiply(errors, MatrixMultiply(outputs[len(outputs)-(1+i)], subtracted)), MatrixTranspose(outputs[len(outputs)-(2+i)])), lr)
		n.Weights[len(n.Weights)-(1+i)] = MatrixSub(n.Weights[len(n.Weights)-(1+i)], updated)
		errors = MatrixDot(MatrixTranspose(n.Weights[len(n.Weights)-(1+i)]), errors)
		/*
			fmt.Println(1 + i)
			fmt.Println(n.Weights[len(n.Weights)-(1+i)])
			fmt.Println("")
		*/
	}

	/*
		outputSubtracted := subtractMatrix(outputActivations)
		errors := MatrixScale(MatrixSub(targetM, outputActivations), -1)
		updatedHO := MatrixScale(MatrixDot(MatrixMultiply(errors, MatrixMultiply(outputActivations, outputSubtracted)), MatrixTranspose(hidden2Activations)), lr)
		n.WeightsHO = MatrixSub(n.WeightsHO, updatedHO)

		hidden2Subtracted := subtractMatrix(hidden2Activations)
		hidden2Errors := MatrixDot(MatrixTranspose(n.WeightsHO), errors)
		updatedHH := MatrixScale(MatrixDot(MatrixMultiply(hidden2Errors, MatrixMultiply(hidden2Activations, hidden2Subtracted)), MatrixTranspose(hidden1Activations)), lr)
		n.WeightsHH = MatrixSub(n.WeightsHH, updatedHH)

		hidden1Subtracted := subtractMatrix(hidden1Activations)
		hidden1Errors := MatrixDot(MatrixTranspose(n.WeightsHH), hidden2Errors)
		updatedIH := MatrixScale(MatrixDot(MatrixMultiply(hidden1Errors, MatrixMultiply(hidden1Activations, hidden1Subtracted)), MatrixTranspose(inpM)), lr)
		n.WeightsIH = MatrixSub(n.WeightsIH, updatedIH)
	*/

	totalError := 0.0
	finals := outputs[len(outputs)-1]
	for i := range finals {
		totalError += 0.5 * math.Pow(targetM[i][0]-finals[i][0], 2)
	}
	return totalError
}

// Guess takes an input and feeds it through the network
func (n *Network) Guess(inp Matrix) []float64 {

	// bias input
	inp = append(inp, []float64{1})

	hidden1Result := MatrixDot(n.WeightsIH, inp)
	hidden1Activations := MatrixMap(hidden1Result, Sigmoid)

	hidden2Result := MatrixDot(n.WeightsHH, hidden1Activations)
	hidden2Activations := MatrixMap(hidden2Result, Sigmoid)

	outputResult := MatrixDot(n.WeightsHO, hidden2Activations)
	outputActivations := MatrixMap(outputResult, Sigmoid)

	return MatrixTranspose(outputActivations)[0]
}

func (n *Network) Train(inp, target Matrix, lr float64) float64 {
	// this can probably be changed so that we train a whole dataset, but for now only 1 input
	// or we can use this method in another method to deal with a whole dataset, in that case we just need to return
	// the mseErrors this input gives so we can inform about the average error of a dataset

	// bias input
	inp = append(inp, []float64{1})

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

	updatedHO := MatrixScale(MatrixDot(MatrixMultiply(errors, MatrixMultiply(outputActivations, outputSubtracted)), MatrixTranspose(hidden2Activations)), lr)
	updatedHH := MatrixScale(MatrixDot(MatrixMultiply(hidden2Errors, MatrixMultiply(hidden2Activations, hidden2Subtracted)), MatrixTranspose(hidden1Activations)), lr)
	updatedIH := MatrixScale(MatrixDot(MatrixMultiply(hidden1Errors, MatrixMultiply(hidden1Activations, hidden1Subtracted)), MatrixTranspose(inp)), lr)

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
