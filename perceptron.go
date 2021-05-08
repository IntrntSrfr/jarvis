package jarvis

type Perceptron struct {
	Weights Vector
	Lr      float64
}

func NewPerceptron(size int, lr float64) Perceptron {
	return Perceptron{
		Weights: make(Vector, size),
		Lr:      lr,
	}
}

func (p *Perceptron) Train(inp Vector, target int) {
	guess := p.Guess(inp)
	miss := target - guess
	p.Weights = VecAdd(p.Weights, VecScale(inp, float64(miss)*p.Lr))
}

func (p *Perceptron) Guess(inp Vector) int {
	if VecDot(p.Weights, inp) >= 0 {
		return 1
	} else {
		return -1
	}
}
