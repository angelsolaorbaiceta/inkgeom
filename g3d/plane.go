package g3d

// A Plane is a three dimensional straight surface defined by a point and a normal versor.
// The plane is defined in its ax + by + cz + d = 0 form.
type Plane struct {
	a, b, c, d   float64
	normalVector *Vector
}

// MakePlane creates a new plane give its a, b, c, and d components of the equation
// ax + by + cz + d = 0.
// Returns an error if the resulting normal vector has zero length.
func MakePlane(a, b, c, d float64) (*Plane, error) {
	normalVec, err := MakeNonZeroVector(a, b, c)

	if err != nil {
		return nil, err
	}

	return &Plane{a, b, c, d, normalVec}, nil
}

// The NormalVector is a vector normal to the surface of the plane.
// Its projections are given by the a, b and c parameters in the plane's equation.
func (p *Plane) NormalVector() *Vector {
	return p.normalVector
}
