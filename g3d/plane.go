package g3d

import "github.com/angelsolaorbaiceta/inkgeom/nums"

// A Plane is a three dimensional straight surface defined by a point and a normal versor.
// The plane is defined in its ax + by + cz + d = 0 form.
type Plane struct {
	d            float64
	normalVector *Vector
}

// MakePlane creates a new plane give its a, b, c, and d components of the equation
// ax + by + cz + d = 0.
// Returns a ErrZeroVector if the resulting normal vector has zero length.
func MakePlane(a, b, c, d float64) (*Plane, error) {
	normalVec, err := MakeNonZeroVector(a, b, c)

	if err != nil {
		return nil, err
	}

	return &Plane{d, normalVec}, nil
}

// MakePlaneFromPointAndNormal returns a new plane passing through a given point and whose normal
// direction is given by the passed in vector.
// Returns a ErrZeroVector if the resulting normal vector has zero length.
func MakePlaneFromPointAndNormal(p *Point, normal *Vector) (*Plane, error) {
	if nums.IsCloseToZero(normal.Length()) {
		return nil, ErrZeroVector
	}

	return &Plane{
		d:            -(p.x*normal.x + p.y*normal.y + p.z*normal.z),
		normalVector: normal,
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

// ContainsPoint checks whether the given point is on the plane.
func (p *Plane) ContainsPoint(pt *Point) bool {
	return nums.IsCloseToZero(p.evaluatePoint(pt) + p.d)
}

func (p *Plane) evaluatePoint(pt *Point) float64 {
	return p.a()*pt.x + p.b()*pt.y + p.c()*pt.z
}
