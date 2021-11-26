package g2d

import (
	"testing"

	"github.com/angelsolaorbaiceta/inkgeom/nums"
)

func TestCreateVector(t *testing.T) {
	var (
		x, y = 2.0, 3.0
		v    = MakeVector(x, y)
	)

	t.Run("has a X projection", func(t *testing.T) {
		if got := v.X(); got != x {
			t.Errorf("Want X = %f, got %f", x, got)
		}
	})

	t.Run("has a Y projection", func(t *testing.T) {
		if got := v.Y(); got != y {
			t.Errorf("Want Y = %f, got %f", y, got)
		}
	})
}

func TestVectorLength(t *testing.T) {
	v := MakeVector(3, 4)

	if !nums.FloatsEqual(v.Length(), 5.0) {
		t.Error("Wrong vector norm")
	}
}

func TestIsVersor(t *testing.T) {
	v := MakeVector(1, 0)

	if !v.IsVersor() {
		t.Error("Vector expected to be versor")
	}
}

func TestIsNotVersor(t *testing.T) {
	v := MakeVector(1, 2)

	if v.IsVersor() {
		t.Error("Vector expected NOT to be versor")
	}
}

func TestMakeVersor(t *testing.T) {
	v := MakeVersor(1, 2)

	if !v.IsVersor() {
		t.Error("Vector expected to be versor")
	}
}

func TestPerpendicular(t *testing.T) {
	v := MakeVersor(1, 2)

	if !v.Perpendicular().Equals(MakeVersor(-2, 1)) {
		t.Error("Wrong perpendicular vector")
	}
}

/* Operations */
func TestSumVectors(t *testing.T) {
	v := MakeVector(1, 2).Plus(MakeVector(3, 4))

	if !v.Equals(MakeVector(4, 6)) {
		t.Error("Addition yielded wrong result")
	}
}

func TestSubtractVectors(t *testing.T) {
	v := MakeVector(1, 2).Minus(MakeVector(3, 5))

	if !v.Equals(MakeVector(-2, -3)) {
		t.Error("Subtraction yielded wrong result")
	}
}

func TestVectorsDotProduct(t *testing.T) {
	p := MakeVector(1, 2).DotTimes(MakeVector(3, 5))

	if !nums.FloatsEqual(p, 13.0) {
		t.Error("Dot product yielded wrong result")
	}
}

func TestVectorsCrossProduct(t *testing.T) {
	p := MakeVector(1, 2).CrossTimes(MakeVector(3, 5))

	if !nums.FloatsEqual(p, -1.0) {
		t.Error("Cross product yielded wrong result")
	}
}
