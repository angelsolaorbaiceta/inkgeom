package intsc

import (
	"github.com/angelsolaorbaiceta/inkgeom/g3d"
	"github.com/angelsolaorbaiceta/inkgeom/nums"
)

// LinePlane is the resulting intersection of a line with a plane.
type LinePlane struct {
	HasIntersection bool
	Point           *g3d.Point
}

// ComputeLinePlane computes the intersection between a line and a plane.
// If the line and the plane are parallel, no intersection is recorded. In any other case, the
// intersection yields the intersection point.
func ComputeLinePlane(line *g3d.Line, plane *g3d.Plane) *LinePlane {
	dirDotNorm := plane.NormalVector().DotTimes(line.Direction())

	if nums.IsCloseToZero(dirDotNorm) {
		return &LinePlane{false, nil}
	}

	lineT := -(plane.EvaluatePoint(line.Origin())) / dirDotNorm

	return &LinePlane{
		HasIntersection: true,
		Point:           line.PointAt(lineT),
	}
}
