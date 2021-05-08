package jarvis

import "testing"

func TestVecDot(t *testing.T) {
	v1 := Vector{2, 3, 4}
	v2 := Vector{3, 4, 5}

	got :=  VecDot(v1, v2)

	if got != 38{
		t.Errorf("VecDot(v1, v2) = %d; want 38", got)
	}
}
func TestVecDotDifferentLengths(t *testing.T) {
	v1 := Vector{1, 2, 3}
	v2 := Vector{1, 2}

	got := VecDot(v1, v2)

	if got != -1{
		t.Errorf("VecDot(v1, v2) = %d; want -1", got)
	}
}