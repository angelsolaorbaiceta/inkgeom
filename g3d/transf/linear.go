package transf

import (
	"math"

	"github.com/angelsolaorbaiceta/inkgeom/g3d"
)

var (
	IdentityLinear = MakeLinear(1, 0, 0, 0, 1, 0, 0, 0, 1)
)

// A Linear transformation that transforms vectors.
// Given the vectors U and V, the linear transformation L:
//
//	L(cU + dV) = cL(U) + dL(V)
//
// where c and d are scalars.
//
// The transformation of a linear combination is a linear combination of the
// transformations.
//
// A linear transformation is represented by a matrix:
//
//	⌈ a  d  g ⌉
//	| b  e  h |
//	⌊ c  f  i ⌋
//
// A linear transformation is typically created to perform a specific transformation,
// such as a rotation or a scaling.
// Use the MakeRotation and MakeScaling functions to create these transformations.
//
// Alternatively, a generic linear transformation can be created using the MakeLinear.
type Linear struct {
	m *Matrix3x3
}

// MakeLinear creates a new linear transformation from the given matrix values.
//
// The matrix is represented by the following values:
//
//	⌈ a  d  g ⌉
//	| b  e  h |
//	⌊ c  f  i ⌋
func MakeLinear(a, d, g, b, e, h, c, f, i float64) *Linear {
	return &Linear{MakeMatrix(a, d, g, b, e, h, c, f, i)}
}

// MakeScaling creates a scaling linear transformation, where a vector is scaled by the
// given factors in the X, Y and Z axes.
//
// The resulting matrix is:
//
//	⌈ x  0  0 ⌉
//	| 0  y  0 |
//	⌊ 0  0  z ⌋
func MakeScaling(x, y, z float64) *Linear {
	return MakeLinear(x, 0, 0, 0, y, 0, 0, 0, z)
}

// MakeUniformScaling creates a scaling linear transformation, where a vector is scaled by the
// given factor in all axes.
//
// The resulting matrix is:
//
//	⌈ s  0  0 ⌉
//	| 0  s  0 |
//	⌊ 0  0  s ⌋
func MakeUniformScaling(s float64) *Linear {
	return MakeScaling(s, s, s)
}

// MakeRotation creates a rotation linear transformation, where a vector is rotated by the
// given angle (in radians) around the origin in the direction of the given axis.
func MakeRotation(radians float64, axis *g3d.Vector) *Linear {
	var (
		x           = axis.X()
		y           = axis.Y()
		z           = axis.Z()
		cos         = math.Cos(radians)
		sin         = math.Sin(radians)
		oneMinusCos = 1 - cos
	)

	return MakeLinear(
		cos+x*x*oneMinusCos, x*y*oneMinusCos-z*sin, x*z*oneMinusCos+y*sin,
		y*x*oneMinusCos+z*sin, cos+y*y*oneMinusCos, y*z*oneMinusCos-x*sin,
		z*x*oneMinusCos-y*sin, z*y*oneMinusCos+x*sin, cos+z*z*oneMinusCos,
	)
}

// Apply applies the linear transformation to the given vector.
func (l *Linear) Apply(vec *g3d.Vector) *g3d.Vector {
	return l.m.TimesVector(vec)
}
