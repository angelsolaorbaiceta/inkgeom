package g3d

import (
	"math"
	"testing"
)

func TestCreateLine(t *testing.T) {
	var (
		origin        = MakePoint(1, 2, 4)
		direction     = MakeVector(1, 1, 1)
		wantDirection = MakeVersor(1, 1, 1)
		line          = MakeLine(origin, direction)
	)

	if got := line.Origin(); !origin.Equals(got) {
		t.Errorf("Want line origin %v, but got %v", origin, got)
	}

	if got := line.Direction(); !wantDirection.Equals(got) {
		t.Errorf("Want line direction %v, but got %v", wantDirection, got)
	}
}

func TestPointAt(t *testing.T) {
	var (
		origin = MakePoint(1, 2, 3)
		dir    = MakeVersor(1, 1, 1)
		line   = MakeLine(origin, dir)
	)

	t.Run("At t = 0, the point is the line's origin", func(t *testing.T) {
		if got := line.PointAt(0.0); !origin.Equals(got) {
			t.Errorf("Want point %v, but got %v", origin, got)
		}
	})

	t.Run("Get point at t > 0", func(t *testing.T) {
		want := MakePoint(2, 3, 4)

		if got := line.PointAt(math.Sqrt(3)); !want.Equals(got) {
			t.Errorf("Want point %v, but got %v", want, got)
		}
	})
}
