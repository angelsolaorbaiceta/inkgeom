package g2d

import "testing"

func TestCreateRect(t *testing.T) {
	var (
		origin  = MakePoint(2, 3)
		width   = 20.0
		height  = 30.0
		rect, _ = MakeRect(*origin, width, height)
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

	t.Run("Returns an error if the width is less than zero", func(t *testing.T) {
		if rect, err := MakeRect(*origin, -width, height); err == nil || rect != nil {
			t.Error("Want error, got nil")
		}
	})

	t.Run("Returns an error if the height is less than zero", func(t *testing.T) {
		if rect, err := MakeRect(*origin, width, -height); err == nil || rect != nil {
			t.Error("Want error, got nil")
		}
	})
}

func TestMakeRectContainingPoints(t *testing.T) {
	var (
		points = []*Point{
			MakePoint(0, 5),
			MakePoint(10, 0),
			MakePoint(5, 7),
		}
		wantRect, _ = MakeRect(*MakePoint(0, 0), 10, 7)
	)

	if got, _ := MakeRectContaining(points); !got.Equals(wantRect) {
		t.Errorf("Want %v, got %v", wantRect, got)
	}
}

func TestRectContainsPoint(t *testing.T) {
	var (
		rect, _       = MakeRect(*MakePoint(10, 20), 100, 200)
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

func TestScaleRect(t *testing.T) {
	var (
		origin      = MakePoint(2, 3)
		width       = 20.0
		height      = 30.0
		scale       = 5.0
		rect, _     = MakeRect(*origin, width, height)
		wantRect, _ = MakeRect(*origin, width*scale, height*scale)
	)

	if got, _ := rect.WithScaledSize(scale); !got.Equals(wantRect) {
		t.Errorf("Want %v, got %v", wantRect, got)
	}
}

func TestRectWithMargin(t *testing.T) {
	rect, _ := MakeRect(*MakePoint(0, 0), 100, 200)

	t.Run("compute a larger rect", func(t *testing.T) {
		var (
			lateralMargin  = 10.0
			verticalMargin = 20.0
			wantRect, _    = MakeRect(*MakePoint(-10, -20), 120, 240)
		)

		if got, _ := rect.WithMargins(lateralMargin, verticalMargin); !got.Equals(wantRect) {
			t.Errorf("Want %v, got %v", wantRect, got)
		}
	})

	t.Run("compute a smaller rect", func(t *testing.T) {
		var (
			lateralMargin  = -10.0
			verticalMargin = -20.0
			wantRect, _    = MakeRect(*MakePoint(10, 20), 80, 160)
		)

		if got, _ := rect.WithMargins(lateralMargin, verticalMargin); !got.Equals(wantRect) {
			t.Errorf("Want %v, got %v", wantRect, got)
		}
	})

	t.Run("can't create a rect with negative width", func(t *testing.T) {
		if got, err := rect.WithMargins(-51, 0); got != nil || err == nil {
			t.Error("Want error")
		}
	})

	t.Run("can't create a rect with negative height", func(t *testing.T) {
		if got, err := rect.WithMargins(0, -101); got != nil || err == nil {
			t.Error("Want error")
		}
	})
}
