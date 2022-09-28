package transf

import "github.com/angelsolaorbaiceta/inkgeom/g3d"

type Matrix3x3 struct {
	a, d, g float64
	b, e, h float64
	c, f, i float64
}

func MakeMatrix(a, d, g, b, e, h, c, f, i float64) *Matrix3x3 {
	return &Matrix3x3{a, d, g, b, e, h, c, f, i}
}

func (m *Matrix3x3) TimesVector(vector *g3d.Vector) *g3d.Vector {
	var (
		x = vector.X()
		y = vector.Y()
		z = vector.Z()
	)

	return g3d.MakeVector(
		m.a*x+m.d*y+m.g*z,
		m.b*x+m.e*y+m.h*z,
		m.c*x+m.f*y+m.i*z,
	)
}

func (m *Matrix3x3) Plus(other *Matrix3x3) *Matrix3x3 {
	return MakeMatrix(
		m.a+other.a, m.d+other.d, m.g+other.g,
		m.b+other.b, m.e+other.e, m.h+other.h,
		m.c+other.c, m.f+other.f, m.i+other.i,
	)
}

func (m *Matrix3x3) Minus(other *Matrix3x3) *Matrix3x3 {
	return MakeMatrix(
		m.a-other.a, m.d-other.d, m.g-other.g,
		m.b-other.b, m.e-other.e, m.h-other.h,
		m.c-other.c, m.f-other.f, m.i-other.i,
	)
}
