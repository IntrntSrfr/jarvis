package jarvis

type Perceptron struct {
	Weights Vector
	Bias    int
}

func (p Perceptron) Guess(inp Vector) int {
	if VecSum(inp) >= 0 {
		return 1
	} else {
		return -1
	}
}
