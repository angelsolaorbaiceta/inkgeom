package inkgeom

import (
    "math"
    "github.com/angelsolaorbaiceta/inkmath"
)

const (
    MIN_T = 0.0
    MAX_T = 1.0
)

type TParam struct {
    value float64
}

/* Construction */
func MakeTParam(value float64) TParam {
    switch {
    case value < MIN_T:
        return TParam{MIN_T}
    case value > MAX_T:
        return TParam{MAX_T}
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

/* Properties */
func (t TParam) IsMin() bool {
    return inkmath.FuzzyEqual(t.value, MIN_T)
}

func (t TParam) IsMax() bool {
    return inkmath.FuzzyEqual(t.value, MAX_T)
}

func (t TParam) Value() float64 {
    return t.value
}

/* Functions */

// Subdivides a given range of t parameters a given number of times, resulting
// in a times + 1 size slice.
func SubdivideTimes(startT, endT TParam, times int) []TParam {
    tParams := make([]TParam, times + 1)
    step := startT.DistanceTo(endT) / float64(times)

    if startT.Value() < endT.Value() {
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
