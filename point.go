package jarvis

import "math/rand"

type Point struct {
	X     float64
	Y     float64
	Label int
}

func NewPoint() Point {
	p := Point{
		X: rand.Float64() * 10,
		Y: rand.Float64() * 10,
	}
	if 0.431*p.X+3 >= p.Y {
		//if p.X >= p.Y {
		p.Label = 1
	} else {
		p.Label = 0
	}
	return p
}
