package jarvis

type Matrix []Vector

/*
func MatrixDot(m1 Matrix, m2 Matrix) Matrix{

}
*/

func MatrixAdd(m1 Matrix, m2 Matrix) Matrix {
	res := make(Matrix, len(m1))
	for i := range m1 {
		res = append(res, VecAdd(m1[i], m2[i]))
	}
	return res
}

func MatrixScale(m1 Matrix, scale float64) Matrix {
	res := make(Matrix, len(m1))
	for i := range m1 {
		res = append(res, VecScale(m1[i], scale))
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
	for i := range m1 {
		row := make(Vector, len(m1))
		for j := range m1[i] {
			row[j] = m1[j][i]
		}
		res = append(res, row)
	}
	return res
}

func MatrixMap(m1 Matrix, f func(float64) float64) Matrix {
	res := make(Matrix, len(m1))
	for i := range m1 {
		res = append(res, VecMap(m1[i], f))
	}
	return res
}

func MatrixSum(m1 Matrix) float64 {
	sum := 0.0
	for i := range m1 {
		sum += VecSum(m1[i])
	}
	return sum
}
