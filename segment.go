package inkgeom

import (
    "github.com/angelsolaorbaiceta/inkmath"
)

type Segment struct {
    Start, End Projectable
}

/* Construction */
func MakeSegment(start, end Projectable) Segment {
    return Segment{start, end}
}

func MakeSegmentFromCoords(startX, startY, endX, endY float64) Segment {
    return Segment{MakePoint(startX, startY), MakePoint(endX, endY)}
}

/* Properties */
func (s Segment) Length() float64 {
    return s.Start.DistanceTo(s.End)
}

/* Methods */
func (s Segment) LengthBetween(startT, endT TParam) float64 {
    return s.Length() * startT.DistanceTo(endT)
}

func (s Segment) PointAt(t TParam) Projectable {
    return MakePoint(
        inkmath.LinInterpol(MIN_T, s.Start.X, MAX_T, s.End.X, t.Value()),
        inkmath.LinInterpol(MIN_T, s.Start.Y, MAX_T, s.End.Y, t.Value()))
}

func (s Segment) DirectionVersor() Projectable {
    return MakeVectorFromTo(s.Start, s.End).ToVersor()
}

func (s Segment) NormalVersor() Projectable {
    return s.DirectionVersor().Perpendicular()
}
