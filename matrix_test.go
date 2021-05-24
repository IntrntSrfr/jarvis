package jarvis

import (
	"testing"
)

func TestMatrixDot(t *testing.T) {
	m1 := Matrix{{2, 3}, {4, 7}, {6, 2}}
	m2 := Matrix{{2, 3, 2}, {1, 4, 2}}

	got := MatrixDot(m1, m2)
	expected := Matrix{{7, 18, 10}, {15, 40, 22}, {14, 26, 16}}

	if !MatrixEqual(got, expected) {
		t.Errorf("MatrixDot(m1, m2): %v; expected: %v", got, expected)
	}

}

func TestMatrixEqual(t *testing.T) {
	m1 := Matrix{{1, 2}, {3, 4}}
	m2 := Matrix{{1, 2}, {3, 4}}

	if !MatrixEqual(m1, m2) {
		t.Errorf("MatrixEqual(m1, m2): %v; expected: %v", false, true)
	}
}

func TestMatrixAdd(t *testing.T) {

	m1 := Matrix{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}}
	m2 := Matrix{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}}

	got := MatrixAdd(m1, m2)
	expected := Matrix{{0, 2, 4}, {6, 8, 10}, {12, 14, 16}}

	if !MatrixEqual(got, expected) {
		t.Errorf("MatrixAdd(m1, m2): %v; expected: %v", got, expected)
	}
}

func TestMatrixSub(t *testing.T) {

	m1 := Matrix{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}}
	m2 := Matrix{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}}

	got := MatrixSub(m1, m2)
	expected := Matrix{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}

	if !MatrixEqual(got, expected) {
		t.Errorf("MatrixSub(m1, m2): %v; expected: %v", got, expected)
	}
}

func TestMatrixMap(t *testing.T) {
	m1 := Matrix{{1, 2}, {2, 4}}
	f := func(i float64) float64 {
		return i + 4
	}

	got := MatrixMap(m1, f)
	expected := Matrix{{5, 6}, {6, 8}}

	if !MatrixEqual(got, expected) {
		t.Errorf("MatrixScale(m1, scale): %v; expected: %v", got, expected)
	}
}

func TestMatrixScale(t *testing.T) {
	m1 := Matrix{{1, 2}, {2, 4}}
	scale := 2.0

	got := MatrixScale(m1, scale)
	expected := Matrix{{2, 4}, {4, 8}}

	if !MatrixEqual(got, expected) {
		t.Errorf("MatrixScale(m1, scale): %v; expected: %v", got, expected)
	}
}

func TestMatrixSum(t *testing.T) {
	m1 := Matrix{{1, 2}, {3, 4}}

	got := MatrixSum(m1)
	if got != 10 {
		t.Errorf("MatrixSum(m1): %v; expected: %v", got, 10)
	}
}

func TestMatrixTranspose(t *testing.T) {

	m1 := Matrix{{1, 2}, {3, 4}, {5, 6}}

	got := MatrixTranspose(m1)
	expected := Matrix{{1, 3, 5}, {2, 4, 6}}

	if !MatrixEqual(got, expected) {
		t.Errorf("MatrixTranspose(m1): %v; expected: %v", got, expected)
	}

}
