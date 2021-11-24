package g3d

import "math"

// Point is a position in space, with X, Y and Z coordinates.
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

func (p *Point) DistanceTo(other *Point) float64 {
	var (
		dx = other.x - p.x
		dy = other.y - p.y
		dz = other.z - p.z
	)

	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func (p *Point) Equals(other *Point) bool {
	return projectablesEqual(p, other)
}
