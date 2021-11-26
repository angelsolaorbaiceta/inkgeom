package g2d

import (
	"github.com/angelsolaorbaiceta/inkgeom/nums"
)

// A Segment is a straignt line defined between two points.
type Segment struct {
	start, end *Point
}

// MakeSegment creates a new segment defined between the given start and end points.
func MakeSegment(start, end *Point) *Segment {
	return &Segment{start, end}
}

// MakeSegmentFromCoords creates a new segment from the projections of the start and end points.
func MakeSegmentFromCoords(startX, startY, endX, endY float64) *Segment {
	return MakeSegment(
		MakePoint(startX, startY),
		MakePoint(endX, endY),
	)
}

func (s *Segment) Start() *Point {
	return s.start
}

func (s *Segment) End() *Point {
	return s.end
}

// Length computes the total length of the segment.
func (s *Segment) Length() float64 {
	return s.start.DistanceTo(s.end)
}

/* <-- Methods--> */

// LengthBetween computes the length of a portion of the segment between two given t values.
func (s *Segment) LengthBetween(startT, endT nums.TParam) float64 {
	return s.Length() * startT.DistanceTo(endT)
}

// PointAt computes an intermediate point in the segment.
func (s *Segment) PointAt(t nums.TParam) *Point {
	var (
		minTVal = nums.MinT.Value()
		maxTVal = nums.MaxT.Value()
	)

	return MakePoint(
		nums.LinInterpol(minTVal, s.start.x, maxTVal, s.end.x, t.Value()),
		nums.LinInterpol(minTVal, s.start.y, maxTVal, s.end.y, t.Value()),
	)
}

/*
DirectionVersor computes the versor which points in the advancing direction of the
segment's [start -> end].
*/
func (s *Segment) DirectionVersor() *Vector {
	return s.start.VectorTo(s.end).ToVersor()
}

// NormalVersor computes the versor perpendicular to the direction versor of the segment.
func (s *Segment) NormalVersor() *Vector {
	return s.DirectionVersor().Perpendicular()
}

/*
RefFrame returns the reference frame of the segment.

The Reference Frame i's versor points in the direction of the direction versor.
*/
func (s *Segment) RefFrame() *RefFrame {
	return MakeRefFrameWithIVersor(s.DirectionVersor())
}
