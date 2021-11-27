package g3d

import (
	"math"
	"testing"

	"github.com/angelsolaorbaiceta/inkgeom/nums"
)

func TestCreateVector(t *testing.T) {
	var (
		x = 10.0
		y = 20.0
		z = 30.0
		p = Vector{x, y, z}
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
		x      = 10.0
		y      = 20.0
		z      = 30.0
		v, err = MakeVersor(x, y, z)
	)

	t.Run("is created without error", func(t *testing.T) {
		if err != nil {
			t.Error("Expected versor creation without error")
		}
	})

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

		versor, err := vector.ToVersor()

		if !versor.IsVersor() {
			t.Error("A versor should have unitary length")
		}
		if err != nil {
			t.Error("Expected versor without error")
		}
	})

	t.Run("Can't be created from three zero components", func(t *testing.T) {
		v, err := MakeVersor(0, 0, 0)

		if v != nil {
			t.Errorf("Want nil vector, but got %v", v)
		}
		if err != ErrZeroVersor {
			t.Error("Expected error")
		}
	})

	t.Run("Can't create a versor from a zero vector", func(t *testing.T) {
		v, err := MakeVector(0, 0, 0).ToVersor()

		if v != nil {
			t.Errorf("Want nil vector, but got %v", v)
		}
		if err != ErrZeroVersor {
			t.Error("Expected error")
		}
	})
}

func TestVectorOpposite(t *testing.T) {
	var (
		v    = MakeVector(1, 2, -3)
		want = MakeVector(-1, -2, 3)
	)

	if got := v.Opposite(); !want.Equals(got) {
		t.Errorf("Want vector %v, but got %v", want, got)
	}
}

func TestVectorScale(t *testing.T) {
	var (
		scale  = 3.0
		vector = MakeVector(1, 2, 3)
		want   = MakeVector(3, 6, 9)
	)

	if got := vector.Scaled(scale); !want.Equals(got) {
		t.Errorf("Want vector %v, but got %v", want, got)
	}
}

func TestVectorOperations(t *testing.T) {
	var (
		u = MakeVector(1, 2, 3)
		v = MakeVector(4, 6, 8)
	)

	t.Run("Vector addition", func(t *testing.T) {
		want := MakeVector(5, 8, 11)

		if got := u.Plus(v); !want.Equals(got) {
			t.Errorf("Want vector %v, but got %v", want, got)
		}
	})

	t.Run("Vector subtraction", func(t *testing.T) {
		want := MakeVector(-3, -4, -5)

		if got := u.Minus(v); !want.Equals(got) {
			t.Errorf("Want vector %v, but got %v", want, got)
		}
	})

	t.Run("Vector dot product", func(t *testing.T) {
		want := 40.0

		if got := u.DotTimes(v); !nums.FloatsEqual(want, got) {
			t.Errorf("Want vector %f, but got %f", want, got)
		}
	})

	t.Run("Vector cross product", func(t *testing.T) {
		want := MakeVector(-2, 4, -2)

		if got := u.CrossTimes(v); !want.Equals(got) {
			t.Errorf("Want vector %v, but got %v", want, got)
		}

		if got := v.CrossTimes(u); !want.Opposite().Equals(got) {
			t.Errorf("Want vector %v, but got %v", want.Opposite(), got)
		}
	})
}
