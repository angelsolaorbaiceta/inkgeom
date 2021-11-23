package g2d

import (
	"github.com/angelsolaorbaiceta/inkgeom/nums"
)

// A Segment is a line between two points.
type Segment struct {
	Start, End Projectable
}

/* <-- Construction --> */

// MakeSegment creates a new segment defined between the given start and end points.
func MakeSegment(start, end Projectable) Segment {
	return Segment{start, end}
}

// MakeSegmentFromCoords creates a new segment from the projections of the start and end points.
func MakeSegmentFromCoords(startX, startY, endX, endY float64) Segment {
	return Segment{MakePoint(startX, startY), MakePoint(endX, endY)}
}

/* <-- Properties --> */

// Length computes the total length of the segment.
func (s Segment) Length() float64 {
	return s.Start.DistanceTo(s.End)
}

/* <-- Methods--> */

// LengthBetween computes the length of a portion of the segment between two given t values.
func (s Segment) LengthBetween(startT, endT nums.TParam) float64 {
	return s.Length() * startT.DistanceTo(endT)
}

// PointAt computes an intermediate point in the segment.
func (s Segment) PointAt(t nums.TParam) Projectable {
	var (
		minTVal = nums.MinT.Value()
		maxTVal = nums.MaxT.Value()
	)

	return MakePoint(
		nums.LinInterpol(minTVal, s.Start.X(), maxTVal, s.End.X(), t.Value()),
		nums.LinInterpol(minTVal, s.Start.Y(), maxTVal, s.End.Y(), t.Value()),
	)
}

/*
DirectionVersor computes the versor which points in the advancing direction of the
segment's [start -> end].
*/
func (s Segment) DirectionVersor() Projectable {
	return MakeVectorFromTo(s.Start, s.End).ToVersor()
}

// NormalVersor computes the versor perpendicular to the direction versor of the segment.
func (s Segment) NormalVersor() Projectable {
	return s.DirectionVersor().Perpendicular()
}

/*
RefFrame returns the reference frame of the segment.

The Reference Frame i's versor points in the direction of the direction versor.
*/
func (s Segment) RefFrame() RefFrame {
	return MakeRefFrameWithIVersor(s.DirectionVersor())
}
