package inkgeom

import (
    "testing"
)

func TestDistanceBetweenTParams(t *testing.T)  {
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
