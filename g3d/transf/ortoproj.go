package transf

import "github.com/angelsolaorbaiceta/inkgeom/g3d"

// An OrtographicProjection is an affine transformation that projects points into a plane.
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
