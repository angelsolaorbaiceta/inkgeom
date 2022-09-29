package transf

import (
	"github.com/angelsolaorbaiceta/inkgeom/g3d"
)

var (
	AffineIdentity = &Affine{LinearIdentity, g3d.Zero}
)

// An Affine transformation is a linear transformation plus a translation.
type Affine struct {
	linear      *Linear
	translation *g3d.Vector
}

// MakeTranslation creates a translation affine transformation.
func MakeTranslation(x, y, z float64) *Affine {
	return &Affine{LinearIdentity, g3d.MakeVector(x, y, z)}
}

// MakeRotationAround creates a rotation affine transformation around the given center
// point, in the direciton of the axis by the given angle.
func MakeRotationAround(radians float64, axis *g3d.Vector, center *g3d.Point) *Affine {
	var (
		rotation    = MakeRotation(radians, axis)
		translation = g3d.Identity3x3Matrix.Minus(rotation.values).TimesProj(center)
	)

	return &Affine{rotation, translation}
}

// Apply applies the affine transformation to the given point or vector.
func (transf *Affine) Apply(proj g3d.Projectable) *g3d.Vector {
	return transf.linear.Apply(proj).Plus(transf.translation)
}
