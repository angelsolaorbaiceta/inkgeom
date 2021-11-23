package g3d

import (
	"math"
	"testing"

	"github.com/angelsolaorbaiceta/inkgeom/nums"
)

func TestCreateProjectable(t *testing.T) {
	var (
		x = 10.0
		y = 20.0
		z = 30.0
		p = Projectable{x, y, z}
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

	t.Run("has a length", func(t *testing.T) {
		want := math.Sqrt(x*x + y*y + z*z)

		if got := p.Length(); got != want {
			t.Errorf("Want length = %f, got %f", want, got)
		}
	})
}

func TestVersor(t *testing.T) {
	var (
		x = 10.0
		y = 20.0
		z = 30.0
		v = MakeVersor(x, y, z)
	)

	t.Run("has unit length", func(t *testing.T) {
		if !nums.IsCloseToOne(v.Length()) {
			t.Errorf("Expected %f to be 1.0", v.Length())
		}
	})

	t.Run("is versor", func(t *testing.T) {
		if !v.IsVersor() {
			t.Error("Expected the vector to be a versor")
		}
	})

	t.Run("As versor", func(t *testing.T) {
		vector := MakeVector(x, y, z)

		if vector.IsVersor() {
			t.Error("Expected the vector to not have unitary length")
		}

		if !vector.ToVersor().IsVersor() {
			t.Error("A versor should have unitary length")
		}
	})
}
