package jarvis

import "math/rand"

type Vector []float64

func NewRandomVector(size int) Vector {
	res := make(Vector, size)
	for i := 0; i < len(res); i++ {
		res[i] = rand.Float64()*(1+1) - 1
	}
	return res
}

func VecAdd(v1 Vector, v2 Vector) Vector {
	if len(v1) != len(v2) {
		panic("vector lenghts do not match")
	}
	res := make(Vector, len(v1))
	for i := 0; i < len(v1); i++ {
		res[i] = v1[i] + v2[i]
	}
	return res
}

func VecEqual(v1 Vector, v2 Vector) bool {
	if len(v1) != len(v2) {
		return false
	}

	for i := 0; i < len(v1); i++ {
		if v1[i] != v2[i] {
			return false
		}
	}

	return true

}

func VecSum(v1 Vector) float64 {
	sum := 0.0
	for i := 0; i < len(v1); i++ {
		sum += v1[i]
	}
	return sum
}

func VecScale(v1 Vector, scale float64) Vector {
	res := make(Vector, len(v1))
	for i := 0; i < len(v1); i++ {
		res[i] = v1[i] * scale
	}
	return res
}

func VecDot(v1 Vector, v2 Vector) float64 {
	if len(v1) != len(v2) {
		panic("vector lenghts do not match")
	}
	sum := 0.0
	for i := 0; i < len(v1); i++ {
		sum += v1[i] * v2[i]
	}
	return sum
}

func VecMap(v1 Vector, f func(float64) float64) Vector {
	res := make(Vector, len(v1))
	for i := 0; i < len(v1); i++ {
		res[i] = f(v1[i])
	}
	return res
}
