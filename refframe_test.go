package inkgeom

import (
	"math"
	"testing"

	"github.com/angelsolaorbaiceta/inkmath"
)

func TestJVersor(t *testing.T) {
	i := MakeVersor(1, 1)
	frame := MakeRefFrameWithIVersor(i)

	if !frame.jVersor.Equals(i.Perpendicular()) {
		t.Error("j versor not as expected")
	}
}

func TestProjectVector(t *testing.T) {
	v := MakeVector(2, 3)
	proj := MakeVector(5.0/math.Sqrt2, 1.0/math.Sqrt2)
	frame := MakeRefFrameWithIVersor(MakeVersor(1, 1))

	if !frame.ProjectVector(v).Equals(proj) {
		t.Error("Projected vector not as expected")
	}
}

func TestPositiveCosine(t *testing.T) {
	i := MakeVersor(1, 1)
	frame := MakeRefFrameWithIVersor(i)
	expectedCos := math.Sqrt2 / 2.0

	if cos := frame.Cos(); !inkmath.FuzzyEqual(cos, expectedCos) {
		t.Errorf("Wrong cosine, expected %f, but got %f", expectedCos, cos)
	}
}

func TestNegativeCosine(t *testing.T) {
	i := MakeVersor(-1, 1)
	frame := MakeRefFrameWithIVersor(i)
	expectedCos := -math.Sqrt2 / 2.0

	if cos := frame.Cos(); !inkmath.FuzzyEqual(cos, expectedCos) {
		t.Errorf("Wrong cosine, expected %f, but got %f", expectedCos, cos)
	}
}

func TestPositiveSine(t *testing.T) {
	i := MakeVersor(1, 1)
	frame := MakeRefFrameWithIVersor(i)
	expectedSin := math.Sqrt2 / 2.0

	if sin := frame.Sin(); !inkmath.FuzzyEqual(sin, expectedSin) {
		t.Errorf("Wrong sine, expected %f, but got %f", expectedSin, sin)
	}
}

func TestNegativeSine(t *testing.T) {
	i := MakeVersor(1, -1)
	frame := MakeRefFrameWithIVersor(i)
	expectedSin := -math.Sqrt2 / 2.0

	if sin := frame.Sin(); !inkmath.FuzzyEqual(sin, expectedSin) {
		t.Errorf("Wrong sine, expected %f, but got %f", expectedSin, sin)
	}
}
