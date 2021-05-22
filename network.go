package jarvis

import "fmt"

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

// Feed takes an input and feeds it through the network
func (n *Network) Feed(inp Matrix) Matrix {
	// Hidden result: result from IH x inp
	hr := MatrixMap(MatrixDot(n.WeightsIH, inp), Sigmoid)
	fmt.Println("HIDDEN RESULT:\n", hr)

	// Output result: result from HO x hr
	or := MatrixMap(MatrixDot(n.WeightsHO, hr), Sigmoid)
	fmt.Println("OUTPUT RESULT:\n", or)

	return or

	/*
		errors := MatrixSub(target, or)
		fmt.Println("ERRORS:\n", errors)

		fmt.Println(MatrixTranspose(n.WeightsHO))
	*/
}

func (n *Network) Train() {

}
