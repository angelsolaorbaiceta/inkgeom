package inkgeom

import (
    "math"
    "github.com/angelsolaorbaiceta/inkmath"
)

const (
    MIN_T_VALUE = 0.0
    MAX_T_VALUE = 1.0
)

var (
    MIN_T = TParam{MIN_T_VALUE}
    MAX_T = TParam{MAX_T_VALUE}
)

type TParam struct {
    value float64
}

/* Construction */
func MakeTParam(value float64) TParam {
    switch {
    case value < MIN_T_VALUE:
        return MIN_T
    case value > MAX_T_VALUE:
        return MAX_T
    default:
        return TParam{value}
    }
}

/* Methods */
func (t TParam) Equals(other TParam) bool {
    return inkmath.FuzzyEqual(t.value, other.value)
}

func (t TParam) DistanceTo(other TParam) float64 {
    return math.Abs(t.value - other.value)
}

func (t TParam) IsGreaterThan(other TParam) bool {
    return t.value > other.value
}

func (t TParam) IsLessThan(other TParam) bool {
    return !t.IsGreaterThan(other)
}

/* Properties */
func (t TParam) IsMin() bool {
    return inkmath.FuzzyEqual(t.value, MIN_T_VALUE)
}

func (t TParam) IsMax() bool {
    return inkmath.FuzzyEqual(t.value, MAX_T_VALUE)
}

func (t TParam) Value() float64 {
    return t.value
}

/* Functions */

// Subdivides a given range of t parameters a given number of times, resulting
// in a times + 1 size slice.
func SubTParamRangeTimes(startT, endT TParam, times int) []TParam {
    tParams := make([]TParam, times + 1)
    step := startT.DistanceTo(endT) / float64(times)

    if startT.IsLessThan(endT) {
        tParams[0] = startT
        tParams[times] = endT
    } else {
        tParams[0] = endT
        tParams[times] = startT
    }

    for i := 1; i < times; i++ {
        tParams[i] = TParam{tParams[i - 1].Value() + step}
    }

    return tParams
}

func SubTParamCompleteRangeTimes(times int) []TParam {
    return SubTParamRangeTimes(MIN_T, MAX_T, times)
}
