package jarvis

type Perceptron struct {
	Weights Vector
	Lr      float64
	Bias    float64
}

// NewPerceptron creates a new perceptron with weights according to a specific size input and a learning rate
func NewPerceptron(size int, lr float64) Perceptron {
	return Perceptron{
		Weights: make(Vector, size),
		Lr:      lr,
		Bias:    1,
	}
}

// Train trains the perceptron using an input and a target value for said input
func (p *Perceptron) Train(inp Vector, target int) {
	guess := p.Guess(inp)
	miss := target - guess
	p.Weights = VecAdd(p.Weights, VecScale(inp, float64(miss)*p.Lr))
}

// Guess simply asks the perceptron to make a prediction
func (p *Perceptron) Guess(inp Vector) int {
	if VecDot(p.Weights, inp)+p.Bias >= 0 {
		return 1
	} else {
		return -1
	}
}
