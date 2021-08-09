package inkgeom

import "testing"

func TestFloatsEqual(t *testing.T) {
	t.Run("floats equal with epsilon", func(t *testing.T) {
		if !FloatsEqualEps(1.0, 1.0001, 0.01) {
			t.Error("Floats expected to equal")
		}
	})

	t.Run("floats don't equal with epsilon", func(t *testing.T) {
		if FloatsEqualEps(1.0, 1.01, 0.001) {
			t.Error("Floats expected to not equal")
		}
	})
}

func TestLinearInterpolation(t *testing.T) {
	var (
		want = 10.0
		got  = LinInterpol(0.0, 0.0, 20.0, 20.0, 10.0)
	)

	if !FloatsEqual(want, got) {
		t.Errorf("Wanted %f, but got %f", want, got)
	}
}
