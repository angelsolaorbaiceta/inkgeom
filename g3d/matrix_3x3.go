package g3d

import "github.com/angelsolaorbaiceta/inkgeom/nums"

var (
	Identity3x3Matrix = Make3x3Matrix(1, 0, 0, 0, 1, 0, 0, 0, 1)
)

type Matrix3x3 struct {
	a, d, g float64
	b, e, h float64
	c, f, i float64
}

func Make3x3Matrix(a, d, g, b, e, h, c, f, i float64) *Matrix3x3 {
	return &Matrix3x3{a, d, g, b, e, h, c, f, i}
}

func (m *Matrix3x3) TimesProj(proj Projectable) *Vector {
	var (
		x = proj.X()
		y = proj.Y()
		z = proj.Z()
	)

	return MakeVector(
		m.a*x+m.d*y+m.g*z,
		m.b*x+m.e*y+m.h*z,
		m.c*x+m.f*y+m.i*z,
	)
}

func (m *Matrix3x3) Plus(other *Matrix3x3) *Matrix3x3 {
	return Make3x3Matrix(
		m.a+other.a, m.d+other.d, m.g+other.g,
		m.b+other.b, m.e+other.e, m.h+other.h,
		m.c+other.c, m.f+other.f, m.i+other.i,
	)
}

func (m *Matrix3x3) Minus(other *Matrix3x3) *Matrix3x3 {
	return Make3x3Matrix(
		m.a-other.a, m.d-other.d, m.g-other.g,
		m.b-other.b, m.e-other.e, m.h-other.h,
		m.c-other.c, m.f-other.f, m.i-other.i,
	)
}

func (m *Matrix3x3) Equals(other *Matrix3x3) bool {
	return nums.FloatsEqual(m.a, other.a) &&
		nums.FloatsEqual(m.d, other.d) &&
		nums.FloatsEqual(m.g, other.g) &&
		nums.FloatsEqual(m.b, other.b) &&
		nums.FloatsEqual(m.e, other.e) &&
		nums.FloatsEqual(m.h, other.h) &&
		nums.FloatsEqual(m.c, other.c) &&
		nums.FloatsEqual(m.f, other.f) &&
		nums.FloatsEqual(m.i, other.i)
}
