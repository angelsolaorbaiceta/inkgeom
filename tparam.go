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
	"math"
)

const (
	minTValue = 0.0
	halfTVaue = 0.5
	maxTValue = 1.0
)

var (
	// MinT is the smallest T parameter
	MinT = TParam{minTValue}
	// HalfT is the average between min and max T values
	HalfT = TParam{halfTVaue}
	// MaxT is the biggest T Parameter
	MaxT = TParam{maxTValue}
)

/*
A TParam is a parameter which takes values between in the range [0, 1].
*/
type TParam struct {
	value float64
}

/* <-- Construction --> */

/*
MakeTParam returns a new T parameter with the given value.

If the value is out of range, it is approximated to the closest end.
*/
func MakeTParam(value float64) TParam {
	switch {
	case value < minTValue:
		return MinT
	case value > maxTValue:
		return MaxT
	default:
		return TParam{value}
	}
}

/*
AverageT creates a new T parameter which value is the average of the given two.
*/
func AverageT(a, b TParam) TParam {
	return MakeTParam(0.5 * (a.value + b.value))
}

/* <-- Methods --> */

/*
Equals compares the given t parameters and returns true if their values are equal.
*/
func (t TParam) Equals(other TParam) bool {
	return FloatsEqual(t.value, other.value)
}

/*
DistanceTo computes the distance between the values of two T parameters.
*/
func (t TParam) DistanceTo(other TParam) float64 {
	return math.Abs(t.value - other.value)
}

/*
IsGreaterThan returns true if this t parameter's value is greater than the other's.
*/
func (t TParam) IsGreaterThan(other TParam) bool {
	return t.value > other.value
}

/*
IsLessThan returns true if this t parameter's value is smaller than the other's.
*/
func (t TParam) IsLessThan(other TParam) bool {
	return t.value < other.value
}

/* <-- Properties --> */

/*
IsMin returns true if this T parameter's value is the minimum value allowed.
*/
func (t TParam) IsMin() bool {
	return FloatsEqual(t.value, minTValue)
}

/*
IsMax returns true if this T parameter's value is the maximum value allowed.
*/
func (t TParam) IsMax() bool {
	return FloatsEqual(t.value, maxTValue)
}

/*
IsExtreme returns true if this T parameter's value is either minimum or maximum.
*/
func (t TParam) IsExtreme() bool {
	return t.IsMax() || t.IsMin()
}

/*
Value returns the value of the parameter.
*/
func (t TParam) Value() float64 {
	return t.value
}

/* <-- sort.Interface --> */

/*
ByTParamValue implements sort.Interface for []TParam based on the value field.
*/
type ByTParamValue []TParam

func (a ByTParamValue) Len() int {
	return len(a)
}

func (a ByTParamValue) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByTParamValue) Less(i, j int) bool {
	return a[i].value < a[j].value
}

/* <-- Functions --> */

/*
SubTParamRangeTimes subdivides a given range of t parameters a given number of times,
resulting in a times + 1 size slice.
*/
func SubTParamRangeTimes(startT, endT TParam, times int) []TParam {
	tParams := make([]TParam, times+1)
	step := startT.DistanceTo(endT) / float64(times)

	if startT.IsLessThan(endT) {
		tParams[0] = startT
		tParams[times] = endT
	} else {
		tParams[0] = endT
		tParams[times] = startT
	}

	for i := 1; i < times; i++ {
		tParams[i] = TParam{tParams[i-1].Value() + step}
	}

	return tParams
}

/*
SubTParamCompleteRangeTimes subdivides the entire range of [t_min, t_max] a
given number of times.
*/
func SubTParamCompleteRangeTimes(times int) []TParam {
	return SubTParamRangeTimes(MinT, MaxT, times)
}
