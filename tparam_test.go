package inkgeom

import (
	"testing"
)

func TestTParamIsMin(t *testing.T) {
	if !MIN_T.IsMin() {
		t.Error("Expected T parameter to be min")
	}
}

func TestTParamIsMax(t *testing.T) {
	if !MAX_T.IsMax() {
		t.Error("Expected T parameter to be max")
	}
}

func TestTParamsEqual(t *testing.T) {
	t1, t2 := MakeTParam(0.34), MakeTParam(0.34)
	if !t1.Equals(t2) {
		t.Error("Expected t parameters to be equal")
	}
}

func TestTParamsNotEqual(t *testing.T) {
	t1, t2 := MakeTParam(0.34), MakeTParam(0.35)
	if t1.Equals(t2) {
		t.Error("Expected t parameters to be equal")
	}
}

func TestDistanceBetweenTParams(t *testing.T) {
	t1, t2 := MakeTParam(0.5), MakeTParam(0.75)

	if t2.DistanceTo(t1) != 0.25 {
		t.Error("Wrong distance between t parameters")
	}
}

func TestSubdivideRange(t *testing.T) {
	t1, t2 := MakeTParam(0.5), MakeTParam(0.7)
	vals := SubTParamRangeTimes(t1, t2, 2)

	if len(vals) != 3 {
		t.Error("Range subdivided wrong number of times")
	}
	if vals[0].Value() != 0.5 || vals[1].Value() != 0.6 || vals[2].Value() != 0.7 {
		t.Error("Wrong values for t parameter")
	}
}
