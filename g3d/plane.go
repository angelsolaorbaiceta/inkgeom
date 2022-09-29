package g3d

import "github.com/angelsolaorbaiceta/inkgeom/nums"

// A Plane is a three dimensional straight surface defined by a point and a normal versor.
// The plane is defined in its ax + by + cz + d = 0 form.
type Plane struct {
	d            float64
	point        *Point
	normalVector *Vector
	normalVersor *Vector
}

// MakePlane creates a new plane give its a, b, c, and d components of the equation:
// 	ax + by + cz + d = 0
//
// Returns a ErrZeroVector if the resulting normal vector has zero length.
func MakePlane(a, b, c, d float64) (*Plane, error) {
	normalVec, err := MakeNonZeroVector(a, b, c)

	if err != nil {
		return nil, err
	}

	// We know the vector isn't zero, so we can safely normalize it.
	versor, _ := normalVec.ToVersor()

	return &Plane{
		d:            d,
		point:        findPointInPlane(a, b, c, d),
		normalVector: normalVec,
		normalVersor: versor,
	}, nil
}

// MakePlaneFromPointAndNormal returns a new plane passing through a given point and whose normal
// direction is given by the passed in vector.
// Returns a ErrZeroVector if the resulting normal vector has zero length.
func MakePlaneFromPointAndNormal(p *Point, normal *Vector) (*Plane, error) {
	if nums.IsCloseToZero(normal.Length()) {
		return nil, ErrZeroVector
	}

	// We know the vector isn't zero, so we can safely normalize it.
	normalVersor, _ := normal.ToVersor()

	return &Plane{
		d:            -(p.x*normal.x + p.y*normal.y + p.z*normal.z),
		point:        p,
		normalVector: normal,
		normalVersor: normalVersor,
	}, nil
}

func (p *Plane) a() float64 {
	return p.normalVector.x
}

func (p *Plane) b() float64 {
	return p.normalVector.y
}

func (p *Plane) c() float64 {
	return p.normalVector.z
}

// The NormalVector is a vector normal to the surface of the plane.
// Its projections are given by the a, b and c parameters in the plane's equation.
func (p *Plane) NormalVector() *Vector {
	return p.normalVector
}

// The NormalVersor is a versor normal to the surface of the plane.
func (p *Plane) NormalVersor() *Vector {
	return p.normalVersor
}

// The Point contained in the plane that's used as base point.
func (p *Plane) Point() *Point {
	return p.point
}

// ContainsPoint checks whether the given point is on the plane.
func (p *Plane) ContainsPoint(pt *Point) bool {
	return nums.IsCloseToZero(p.EvaluatePoint(pt))
}

// EvaluatePoint returns the result of evaluating a point in the plane's ax + by + cz + d equation.
func (p *Plane) EvaluatePoint(pt *Point) float64 {
	return p.a()*pt.x + p.b()*pt.y + p.c()*pt.z + p.d
}

func findPointInPlane(a, b, c, d float64) *Point {
	if !nums.IsCloseToZero(a) {
		return MakePoint(-d/a, 0, 0)
	}

	if !nums.IsCloseToZero(b) {
		return MakePoint(0, -d/b, 0)
	}

	if !nums.IsCloseToZero(c) {
		return MakePoint(0, 0, -d/c)
	}

	panic("Plane has zero normal vector")
}
