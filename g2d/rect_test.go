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

	if got := rect.Left(); got != origin.x {
		t.Errorf("Want left %f, got %f", origin.x, got)
	}
	if got := rect.Right(); got != origin.x+width {
		t.Errorf("Want right %f, got %f", origin.x+width, got)
	}
	if got := rect.Bottom(); got != origin.y {
		t.Errorf("Want bottom %f, got %f", origin.y, got)
	}
	if got := rect.Top(); got != origin.y+height {
		t.Errorf("Want top %f, got %f", origin.y+height, got)
	}
}

func TestRectContainsPoint(t *testing.T) {
	var (
		rect          = MakeRect(*MakePoint(10, 20), 100, 200)
		pointInside   = MakePoint(50, 70)
		pointsOutside = []*Point{
			MakePoint(5, 70),
			MakePoint(50, 300),
			MakePoint(150, 70),
			MakePoint(50, 5),
		}
	)

	if !rect.ContainsPoint(pointInside) {
		t.Errorf("Expected %v to cointain point %v", rect, pointInside)
	}

	for _, point := range pointsOutside {
		if rect.ContainsPoint(point) {
			t.Errorf("Expected %v not to cointain point %v", rect, point)
		}
	}
}
