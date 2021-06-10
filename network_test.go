package jarvis

import "testing"

func TestNetwork_AddLayer(t *testing.T) {
	n := NewNetwork2()
	n.AddLayer(784)
	n.AddLayer(256)
	n.AddLayer(256)
	n.AddLayer(10)
	if len(n.Layers) != 4 {
		t.Errorf("Layer length is not right; want 4, got %v", len(n.Layers))
	}
}
