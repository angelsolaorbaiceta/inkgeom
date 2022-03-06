package g2d

import "testing"

func TestCreateRect(t *testing.T) {
	var (
		origin = MakePoint(2, 3)
		width  = 20.0
		height = 30.0
		rect   = MakeRect(*origin, width, height)
	)

	if got := rect.Origin(); !got.Equals(origin) {
		t.Errorf("Want origin %v, got %v", origin, got)
	}
	if got := rect.Width(); got != width {
		t.Errorf("Want width %f, got %f", got, width)
	}
	if got := rect.Height(); got != height {
		t.Errorf("Want height %f, got %f", got, height)
	}
}
