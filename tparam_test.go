/*
Copyright 2020 Angel Sola Orbaiceta

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package inkgeom

import (
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
