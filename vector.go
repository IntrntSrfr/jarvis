package jarvis

import "math/rand"

// Vector represents a euclidean vector
type Vector []float64

// NewRandomVector creates a vector of a specified size with random values for every entry in the vector
func NewRandomVector(size int) Vector {
	res := make(Vector, size)
	for i := 0; i < len(res); i++ {
		res[i] = rand.Float64()*(1+1) - 1
	}
	return res
}

// VecAdd adds two Vector objects together
func VecAdd(v1 Vector, v2 Vector) Vector {
	if len(v1) != len(v2) {
		panic("vector lengths do not match")
	}
	res := make(Vector, len(v1))
	for i := 0; i < len(v1); i++ {
		res[i] = v1[i] + v2[i]
	}
	return res
}

// VecEqual checks if two Vector objects contain the same values
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

// VecSum computes the sum of all the values in a Vector
func VecSum(v1 Vector) float64 {
	sum := 0.0
	for i := 0; i < len(v1); i++ {
		sum += v1[i]
	}
	return sum
}

// VecScale scales a Vector by a scalar
func VecScale(v1 Vector, scale float64) Vector {
	res := make(Vector, len(v1))
	for i := 0; i < len(v1); i++ {
		res[i] = v1[i] * scale
	}
	return res
}

// VecDot takes the dot product of two Vector objects
func VecDot(v1 Vector, v2 Vector) float64 {
	if len(v1) != len(v2) {
		panic("vector lengths do not match")
	}
	sum := 0.0
	for i := 0; i < len(v1); i++ {
		sum += v1[i] * v2[i]
	}
	return sum
}

// VecMap changes every entry of a Vector by a given function
func VecMap(v1 Vector, f func(float64) float64) Vector {
	res := make(Vector, len(v1))
	for i := 0; i < len(v1); i++ {
		res[i] = f(v1[i])
	}
	return res
}
