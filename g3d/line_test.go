package g3d

import (
	"math"
	"testing"
)

func TestCreateLine(t *testing.T) {
	var (
		origin           = MakePoint(1, 2, 4)
		direction        = MakeVector(1, 1, 1)
		wantDirection, _ = MakeVersor(1, 1, 1)
		line, err        = MakeLine(origin, direction)
	)

	if err != nil {
		t.Error("Expected line to be created without error")
	}

	t.Run("has origin", func(t *testing.T) {
		if got := line.Origin(); !origin.Equals(got) {
			t.Errorf("Want line origin %v, but got %v", origin, got)
		}
	})

	t.Run("has direction", func(t *testing.T) {
		if got := line.Direction(); !wantDirection.Equals(got) {
			t.Errorf("Want line direction %v, but got %v", wantDirection, got)
		}
	})

	t.Run("can't be created from a zero vector as direction", func(t *testing.T) {
		line, err := MakeLine(origin, Zero)

		if line != nil {
			t.Error("Expected line to be nil")
		}
		if err != ErrZeroVersor {
			t.Error("Expected creation to yield error")
		}
	})
}

func TestPointAt(t *testing.T) {
	var (
		origin  = MakePoint(1, 2, 3)
		dir, _  = MakeVersor(1, 1, 1)
		line, _ = MakeLine(origin, dir)
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
