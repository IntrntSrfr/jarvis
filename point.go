package jarvis

import "math/rand"

type Point struct {
	X     float64
	Y     float64
	Label int
}

func NewPoint() Point {
	p := Point{
		X: rand.Float64() * 100,
		Y: rand.Float64() * 100,
	}
	if p.X >= p.Y {
		p.Label = 1
	} else {
		p.Label = -1
	}
	return p
}
