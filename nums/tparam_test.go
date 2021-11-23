package nums

import (
	"sort"
	"testing"
)

func TestTParamIsMin(t *testing.T) {
	if !MinT.IsMin() {
		t.Error("Expected T parameter to be min")
	}
}

func TestTParamIsMax(t *testing.T) {
	if !MaxT.IsMax() {
		t.Error("Expected T parameter to be max")
	}
}

func TestTParamIsExtreme(t *testing.T) {
	t.Run("is extreme if min", func(t *testing.T) {
		if !MinT.IsExtreme() {
			t.Errorf("Expected %v to be extreme", MinT)
		}
	})

	t.Run("is extreme if max", func(t *testing.T) {
		if !MaxT.IsExtreme() {
			t.Errorf("Expected %v to be extreme", MaxT)
		}
	})

	t.Run("is not extreme if not min or max", func(t *testing.T) {
		if HalfT.IsExtreme() {
			t.Errorf("Expected %v to not be extreme", HalfT)
		}
	})
}

func TestTParamsEqual(t *testing.T) {
	var (
		t1 = MakeTParam(0.34)
		t2 = MakeTParam(0.34)
		t3 = MakeTParam(0.35)
	)

	t.Run("a t param equals itself", func(t *testing.T) {
		if !t1.Equals(t1) {
			t.Error("Expected t value to equal itself")
		}
	})

	t.Run("two t params are equal if they have equal values", func(t *testing.T) {
		if !t1.Equals(t2) {
			t.Error("Expected t parameters to be equal")
		}
	})

	t.Run("two t params are't equal if they have different values", func(t *testing.T) {
		if t1.Equals(t3) {
			t.Error("Expected t parameters to be equal")
		}
	})
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

func TestGreaterOrLessThan(t *testing.T) {
	t1, t2 := MakeTParam(0.5), MakeTParam(0.7)

	t.Run("is less than", func(t *testing.T) {
		if !t1.IsLessThan(t2) {
			t.Error("Expected t1 < t2")
		}
	})

	t.Run("is greater than", func(t *testing.T) {
		if !t2.IsGreaterThan(t1) {
			t.Error("Expected t2 > t1")
		}
	})

	t.Run("is equal", func(t *testing.T) {
		if t1.IsLessThan(t1) {
			t.Error("Expected t1 == t1")
		}
		if t1.IsGreaterThan(t1) {
			t.Error("Expected t1 == t1")
		}
	})
}

func TestAverageT(t *testing.T) {
	var (
		t1   = MakeTParam(0.25)
		t2   = MakeTParam(0.75)
		want = MakeTParam(0.5)
	)

	if got := AverageT(t1, t2); !got.Equals(want) {
		t.Errorf("Want average to be %v, got %v", want, got)
	}
}

func TestTParamSorting(t *testing.T) {
	var (
		a       = MakeTParam(0.2)
		b       = MakeTParam(0.5)
		c       = MakeTParam(0.7)
		tparams = []TParam{c, a, b}
	)

	sort.Sort(ByTParamValue(tparams))

	if !tparams[0].Equals(a) {
		t.Error("Wrong ordering of t params")
	}
	if !tparams[1].Equals(b) {
		t.Error("Wrong ordering of t params")
	}
	if !tparams[2].Equals(c) {
		t.Error("Wrong ordering of t params")
	}
}
