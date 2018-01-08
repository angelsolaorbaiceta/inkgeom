package inkgeom

import (
    "testing"
    "math"
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
    proj := MakeVector(5.0 / math.Sqrt2, 1.0 / math.Sqrt2)
    frame := MakeRefFrameWithIVersor(MakeVersor(1, 1))

    if !frame.ProjectVector(v).Equals(proj) {
        t.Error("Projected vector not as expected")
    }
}
