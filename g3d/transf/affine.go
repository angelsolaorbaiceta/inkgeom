package transf

import "github.com/angelsolaorbaiceta/inkgeom/g3d"

// An Affine transformation is a linear transformation plus a translation.
type Affine struct {
	linear      *Linear
	translation *g3d.Vector
}

// MakeRotationAround creates a rotation around the given center point, in the direciton
// of the axis by the given angle.
func MakeRotationAround(radians float64, axis *g3d.Vector, center *g3d.Point) *Affine {
	var (
		rotation    = MakeRotation(radians, axis)
		translation = rotation.values.Minus(IdentityMatrix).TimesProj(center)
	)

	return &Affine{rotation, translation}
}

func (transf *Affine) Apply(vec *g3d.Vector) *g3d.Vector {
	return transf.linear.Apply(vec).Plus(transf.translation)
}
