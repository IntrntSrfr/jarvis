package jarvis

type Perceptron struct {
	Weights Vector
	Lr      float64
}

func NewPerceptron(size int, lr float64) Perceptron {
	return Perceptron{
		//Weights: NewRandomVector(size),
		Weights: Vector{0, 0},
		Lr:      lr,
	}
}

func (p Perceptron) Train(inp Vector, target int) {
	guess := p.Guess(inp)
	miss := target - guess

	for i := range p.Weights {
		p.Weights[i] += inp[i] * float64(miss) * p.Lr
	}

	/*
		lol := VecScale(inp, float64(miss))
		lol = VecScale(lol, p.Lr)
		fmt.Println("DELTA W", lol)
		lol = VecAdd(p.Weights, lol)
		fmt.Println("NEW W", lol)
		p.Weights = VecAdd(p.Weights, VecScale(VecScale(inp, float64(miss)), p.Lr))
	*/
}

func (p Perceptron) Guess(inp Vector) int {

	if VecDot(p.Weights, inp) >= 0 {
		return 1
	} else {
		return -1
	}
}
