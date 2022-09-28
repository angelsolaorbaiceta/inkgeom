package transf

import "github.com/angelsolaorbaiceta/inkgeom/g3d"

// Linear is a linear transformation applied to a vector, such that, given the vectors
// U and V, the linear transformation L:
//
//	L(cU + dV) = cL(U) + dL(V)
//
// where c and d are scalars.
//
// A linear transformation is represented by a matrix:
//
//	⌈ a  d  g ⌉
//	| b  e  h |
//	⌊ c  f  i ⌋
//
// A linear transformation is typically created to perform a specific transformation,
// such as a rotation, a translation or a scaling.
// Use the MakeRotation, MakeTranslation and MakeScaling functions to create
// these transformations.
//
// Alternatively, a generic linear transformation can be created using the `MakeLinear`
type Linear struct {
	a, d, g float64
	b, e, h float64
	c, f, i float64
}

// MakeLinear creates a new linear transformation from the given matrix values.
//
// The matrix is represented by the following values:
//
//	⌈ a  d  g ⌉
//	| b  e  h |
//	⌊ c  f  i ⌋
func MakeLinear(a, d, g, b, e, h, c, f, i float64) *Linear {
	return &Linear{a, d, g, b, e, h, c, f, i}
}

// MakeScaling creates a scaling linear transformation, where a vector is scaled by the
// given factors in the X, Y and Z axes.
func MakeScaling(x, y, z float64) *Linear {
	return MakeLinear(x, 0, 0, 0, y, 0, 0, 0, z)
}

// Apply applies the linear transformation to the given vector.
func (l *Linear) Apply(vec *g3d.Vector) *g3d.Vector {
	var (
		x = vec.X()
		y = vec.Y()
		z = vec.Z()
	)

	return g3d.MakeVector(
		l.a*x+l.d*y+l.g*z,
		l.b*x+l.e*y+l.h*z,
		l.c*x+l.f*y+l.i*z,
	)
}
