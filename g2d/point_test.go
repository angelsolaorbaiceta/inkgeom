package g2d

import (
	"testing"

	"github.com/angelsolaorbaiceta/inkgeom/nums"
)

func TestCreatePoint(t *testing.T) {
	var (
		x, y = 2.0, 3.0
		p    = MakePoint(x, y)
	)

	t.Run("has a X projection", func(t *testing.T) {
		if got := p.X(); got != x {
			t.Errorf("Want X = %f, got %f", x, got)
		}
	})

	t.Run("has a Y projection", func(t *testing.T) {
		if got := p.Y(); got != y {
			t.Errorf("Want Y = %f, got %f", y, got)
		}
	})
}

func TestDistanceBetweenPoints(t *testing.T) {
	var (
		p    = MakePoint(1, 2)
		q    = MakePoint(4, 6)
		want = 5.0
	)

	if got := p.DistanceTo(q); !nums.FloatsEqual(want, got) {
		t.Errorf("Want distance %f, got %f", want, got)
	}
}

func TestPointOrdering(t *testing.T) {
	t.Run("equal points have the same position", func(t *testing.T) {
		p := MakePoint(2, 5)

		if p.Compare(p) != 0 {
			t.Error("Expected points to have equal order")
		}
	})

	t.Run("point with smaller X goes first", func(t *testing.T) {
		var (
			p = MakePoint(1, 5)
			q = MakePoint(2, 3)
		)

		if p.Compare(q) != -1 {
			t.Error("Expected point p to go begore q")
		}
		if q.Compare(p) != 1 {
			t.Error("Expected point p to go begore q")
		}
	})

	t.Run("point with same X but smaller Y goes first", func(t *testing.T) {
		var (
			p = MakePoint(2, 5)
			q = MakePoint(2, 6)
		)

		if p.Compare(q) != -1 {
			t.Error("Expected point p to go begore q")
		}
		if q.Compare(p) != 1 {
			t.Error("Expected point p to go begore q")
		}
	})

}
