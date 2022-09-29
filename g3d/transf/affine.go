package transf

import "github.com/angelsolaorbaiceta/inkgeom/g3d"

// An Affine transformation is a linear transformation plus a translation.
type Affine struct {
	linear      *Linear
	translation *g3d.Vector
}

// MakeRotationAround creates a rotation affine transformation around the given center
// point, in the direciton of the axis by the given angle.
func MakeRotationAround(radians float64, axis *g3d.Vector, center *g3d.Point) *Affine {
	var (
		rotation    = MakeRotation(radians, axis)
		translation = IdentityMatrix.Minus(rotation.values).TimesProj(center)
	)

	return &Affine{rotation, translation}
}

// Apply applies the affine transformation to the given point or vector.
func (transf *Affine) Apply(proj g3d.Projectable) *g3d.Vector {
	return transf.linear.Apply(proj).Plus(transf.translation)
}
