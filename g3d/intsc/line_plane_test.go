package intsc

import (
	"testing"

	"github.com/angelsolaorbaiceta/inkgeom/g3d"
)

func TestLinePlaneIntersection(t *testing.T) {
	t.Run("a line parallel to a plane has no intersection", func(t *testing.T) {
		var (
			line, _      = g3d.MakeLine(g3d.Origin, g3d.IVersor)
			plane, _     = g3d.MakePlaneFromPointAndNormal(g3d.MakePoint(1, 1, 1), g3d.KVersor)
			intersection = ComputeLinePlane(line, plane)
		)

		if intersection.HasIntersection {
			t.Error("Expected no intersection")
		}
		if intersection.Point != nil {
			t.Error("Expected no intersection point")
		}
	})

	t.Run("intersection point", func(t *testing.T) {
		var (
			line, _      = g3d.MakeLine(g3d.MakePoint(5, 7, 9), g3d.KVersor)
			xyPlane, _   = g3d.MakePlaneFromPointAndNormal(g3d.Origin, g3d.KVersor)
			intersection = ComputeLinePlane(line, xyPlane)
			want         = g3d.MakePoint(5, 7, 0)
		)

		if !intersection.HasIntersection {
			t.Error("Expected intersection")
		}

		if !intersection.Point.Equals(want) {
			t.Errorf("Want intersection point %v, got %v", want, intersection.Point)
		}
	})
}
