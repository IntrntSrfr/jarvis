package jarvis

import (
	"fmt"
	"math"
)

type Network struct {
	Inputs    int
	Hidden    int
	Output    int
	WeightsIH Matrix
	WeightsHO Matrix
	LR        float64
}

func NewNetwork(in, hid, out int, LR float64) *Network {
	return &Network{
		Inputs: in,
		Hidden: hid,
		Output: out,
		LR:     LR,
		//WeightsIH: NewMatrix(hid, in),
		//WeightsHO: NewMatrix(out, hid),
		WeightsIH: NewRandomMatrix(hid, in),
		WeightsHO: NewRandomMatrix(out, hid),
	}
}

func (n *Network) String() string {
	return fmt.Sprintf("NETWORK - inputs: %v, hidden: %v, outputs: %v, lr: %v\nIH Weights: \n%v\nHO Weights: \n%v", n.Inputs, n.Hidden, n.Output, n.LR, n.WeightsIH, n.WeightsHO)
}

// Guess takes an input and feeds it through the network
func (n *Network) Guess(inp Matrix) Matrix {
	// Hidden result: result from IH x inp
	hiddenResult := MatrixDot(n.WeightsIH, inp)
	hiddenActivations := MatrixMap(hiddenResult, Sigmoid)
	//fmt.Println("HIDDEN RESULT:\n", hr)

	// Output result: result from HO x hr
	outputResult := MatrixDot(n.WeightsHO, hiddenActivations)
	outputActivations := MatrixMap(outputResult, Sigmoid)
	//fmt.Println("OUTPUT RESULT:\n", or)

	return outputActivations

}

func (n *Network) Train(inp, target Matrix) {
	// this can probably be changed so that we train a whole dataset, but for now only 1 input
	// or we can use this method in another method to deal with a whole dataset, in that case we just need to return
	// the mseErrors this input gives so we can inform about the average error of a dataset

	hiddenResult := MatrixDot(n.WeightsIH, inp)
	hiddenActivations := MatrixMap(hiddenResult, Sigmoid)
	//fmt.Println("HIDDEN RESULT:\n", hr)

	// Output result: result from HO x hr
	outputResult := MatrixDot(n.WeightsHO, hiddenActivations)
	outputActivations := MatrixMap(outputResult, Sigmoid)

	mseErrors := make(Matrix, n.Output)

	totalError := 0.0
	for i := range outputActivations {
		outputError := 0.5 * math.Pow(target[i][0]-outputActivations[i][0], 2)
		totalError += outputError
		mseErrors[i] = Vector{outputError}
		//fmt.Println(fmt.Sprintf("error for output %v: %v", i, outputError))
	}
	fmt.Println("total error", totalError)
	//fmt.Println(fmt.Sprintf("HIDDEN OUTPUTS: \n\t%v", hiddenActivations))
	//fmt.Println(fmt.Sprintf("FINAL OUTPUTS: \n\t%v", outputActivations))
	//fmt.Println(fmt.Sprintf("FINAL MSE ERRORS: \n\t%v", mseErrors))

	errors := MatrixSub(target, outputActivations)
	errors = MatrixScale(errors, -1)
	//fmt.Println(errors)

	finalSubtracted := subtractMatrix(outputActivations)
	hiddenSubtracted := subtractMatrix(hiddenActivations)

	//fmt.Println(fmt.Sprintf("1 - FINAL OUTPUTS: \n\t%v", finalSubtracted))

	hiddenErrors := MatrixDot(MatrixTranspose(n.WeightsHO), errors)

	updatedHO := MatrixScale(MatrixDot(MatrixMultiply(errors, MatrixMultiply(outputActivations, finalSubtracted)), MatrixTranspose(hiddenActivations)), n.LR)
	updatedIH := MatrixScale(MatrixDot(MatrixMultiply(hiddenErrors, MatrixMultiply(hiddenActivations, hiddenSubtracted)), MatrixTranspose(inp)), n.LR)

	//fmt.Println(fmt.Sprintf("IH WEIGHTS:\n\t%v", n.WeightsIH))
	//fmt.Println(fmt.Sprintf("HO WEIGHTS:\n\t%v", n.WeightsHO))

	//fmt.Println(fmt.Sprintf("CHANGE IH WEIGHTS:\n\t%v", updatedIH))
	//fmt.Println(fmt.Sprintf("CHANGE HO WEIGHTS:\n\t%v", updatedHO))

	n.WeightsIH = MatrixSub(n.WeightsIH, updatedIH)
	n.WeightsHO = MatrixSub(n.WeightsHO, updatedHO)
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
