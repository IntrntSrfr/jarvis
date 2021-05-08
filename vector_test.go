package jarvis

import "testing"

func TestVecDot(t *testing.T) {
	v1 := Vector{2, 3, 4}
	v2 := Vector{3, 4, 5}

	got := VecDot(v1, v2)

	if got != 38 {
		t.Errorf("VecDot(v1, v2) = %d; want 38", got)
	}
}
func TestVecDotDifferentLengths(t *testing.T) {
	v1 := Vector{1, 2, 3}
	v2 := Vector{1, 2}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("VecDot(v1, v2) should have had panicked")
		}
	}()
	VecDot(v1, v2)
}

func TestVecEqual(t *testing.T) {
	v1 := Vector{1, 2, 3}
	v2 := Vector{1, 2, 3}

	if !VecEqual(v1, v2) {
		t.Errorf("VecEqual(v1, v2) = %v; want true", false)
	}
}

func TestVecAdd(t *testing.T) {
	v1 := Vector{1, 2, 3}
	v2 := Vector{1, 2, 3}

	got := VecAdd(v1, v2)

	expected := Vector{2, 4, 6}

	if !VecEqual(got, expected) {
		t.Errorf("VecEqual(v1, v2) = %v; want true", false)
	}
}
