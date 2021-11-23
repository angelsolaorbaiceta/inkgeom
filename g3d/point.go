package g3d

/*
Point is an entity with projections in the X, Y and Z axes.
Used to represent both points and vectors in two dimensions.
*/
type Point struct {
	x, y, z float64
}

/* <-- Construction --> */

func MakePoint(x, y, z float64) *Point {
	return &Point{x, y, z}
}

/* <-- Properties --> */

func (p *Point) X() float64 {
	return p.x
}

func (p *Point) Y() float64 {
	return p.y
}

func (p *Point) Z() float64 {
	return p.z
}

/* <-- Methods --> */

func (p *Point) Equals(other *Point) bool {
	return ProjectablesEqual(p, other)
}
