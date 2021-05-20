package jarvis

import (
	"fmt"
	"math/rand"
)

type Point struct {
	X     float64
	Y     float64
	Label float64
}

func NewPoint() Point {
	p := Point{
		X: rand.Float64() * 10,
		Y: rand.Float64() * 10,
	}
	if -0.75*p.X+6 >= p.Y {
		//if p.X >= p.Y {
		p.Label = 1
	} else {
		p.Label = 0
	}
	return p
}

func (p Point) String() string {
	return fmt.Sprintf("(%v, %v)", p.X, p.Y)
}
