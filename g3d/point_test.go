package g3d

import (
	"math"
	"testing"

	"github.com/angelsolaorbaiceta/inkgeom/nums"
)

func TestCreatePoint(t *testing.T) {
	var (
		x = 10.0
		y = 20.0
		z = 30.0
		p = Point{x, y, z}
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

	t.Run("has a Z projection", func(t *testing.T) {
		if got := p.Z(); got != z {
			t.Errorf("Want Z = %f, got %f", z, got)
		}
	})
}

func TestPointsDistance(t *testing.T) {
	var (
		p    = MakePoint(1, 2, 3)
		q    = MakePoint(4, 7, 9)
		want = math.Sqrt(70)
	)

	if got := p.DistanceTo(q); !nums.FloatsEqual(want, got) {
		t.Errorf("Want distance %f, got %f", want, got)
	}
}
