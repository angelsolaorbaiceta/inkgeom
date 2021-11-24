package g3d

import "testing"

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
