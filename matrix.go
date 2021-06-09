package jarvis

import "math/rand"

type Matrix [][]float64

func NewMatrix(m, n int) Matrix {
	rows := make(Matrix, m)
	for i := range rows {
		rows[i] = make([]float64, n)
	}
	return rows
}

func NewRandomMatrix(m, n int) Matrix {
	rows := make(Matrix, m)
	for i := range rows {
		lol := make([]float64, n)
		for j := range lol {
			lol[j] = rand.Float64()*(1+1) - 1
		}
		rows[i] = lol
	}
	return rows
}

func MatrixDot(m1 Matrix, m2 Matrix) Matrix {

	res := make(Matrix, len(m1))
	for i := range res {
		row := make([]float64, len(m2[0]))
		for j := range row {
			sum := 0.0
			for k := 0; k < len(m1[i]); k++ {
				sum += m1[i][k] * m2[k][j]
			}
			row[j] = sum
		}
		res[i] = row
	}
	return res
}

func MatrixAdd(m1 Matrix, m2 Matrix) Matrix {
	res := make(Matrix, len(m1))
	for i := range res {
		row := make([]float64, len(m1[i]))
		for j := range row {
			row[j] = m1[i][j] + m2[i][j]
		}
		res[i] = row
	}
	return res
}

func MatrixScale(m1 Matrix, scale float64) Matrix {
	res := make(Matrix, len(m1))
	for i := range res {
		row := make([]float64, len(m1[i]))
		for j := range row {
			row[j] = m1[i][j] * scale
		}
		res[i] = row
	}
	return res
}

func MatrixSub(m1 Matrix, m2 Matrix) Matrix {
	return MatrixAdd(m1, MatrixScale(m2, -1))
}

func MatrixMultiply(m1 Matrix, m2 Matrix) Matrix {
	res := make(Matrix, len(m1))
	for i := range res {
		row := make([]float64, len(m1[i]))
		for j := range row {
			row[j] = m1[i][j] * m2[i][j]
		}
		res[i] = row
	}
	return res
}

func MatrixEqual(m1 Matrix, m2 Matrix) bool {
	for i := range m1 {
		for j := range m1[i] {
			if m1[i][j] != m2[i][j] {
				return false
			}
		}
	}
	return true
}

func MatrixTranspose(m1 Matrix) Matrix {
	res := make(Matrix, len(m1[0]))
	for i := range res {
		row := make([]float64, len(m1))
		for j := range row {
			row[j] = m1[j][i]
		}
		res[i] = row
	}
	return res
}

func MatrixMap(m1 Matrix, f func(float64) float64) Matrix {
	res := make(Matrix, len(m1))
	for i := range m1 {
		row := make([]float64, len(m1[i]))
		for j := range m1[i] {
			row[j] = f(m1[i][j])
		}
		res[i] = row
	}
	return res
}

func MatrixSum(m1 Matrix) float64 {
	sum := 0.0
	for i := range m1 {
		for j := range m1[i] {
			sum += m1[i][j]
		}
	}
	return sum
}
