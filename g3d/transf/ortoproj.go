package transf

import "github.com/angelsolaorbaiceta/inkgeom/g3d"

// An OrtographicProjection is an affine transformation that projects points into a plane.
//
// Ortographic projections aren't invertible: given a point in the plane, there are an
// infinite number of points in the 3D space that project into it (those in a line
// perpendicular to the plane).
type OrtographicProjection struct {
	affine *Affine
}

// MakeOrtographic creates a new ortographic projection transformation that projects points
// into the given plane.
func MakeOrtographic(plane *g3d.Plane) *OrtographicProjection {
	var (
		n   = plane.NormalVersor()
		nSq = n.Squared()
	)

	return &OrtographicProjection{
		affine: &Affine{
			linear:      &Linear{values: g3d.Identity3x3Matrix.Minus(nSq)},
			translation: nSq.TimesProj(plane.Point()),
		},
	}
}

// Apply projects the given point into the plane.
func (transf *OrtographicProjection) Apply(proj g3d.Projectable) *g3d.Point {
	return transf.affine.Apply(proj).ToPoint()
}
