package g3d

import "math"

var (
	Origin = MakePoint(0, 0, 0)
)

// Point is a position in space, defined by its X, Y and Z coordinates.
type Point struct {
	x, y, z float64
}

// MakePoint creates a new point given its X, Y and Z coordinates.
func MakePoint(x, y, z float64) *Point {
	return &Point{x, y, z}
}

// X is the x coordinate.
func (p *Point) X() float64 {
	return p.x
}

// Y is the y coordinate.
func (p *Point) Y() float64 {
	return p.y
}

// Z is the z coordinate.
func (p *Point) Z() float64 {
	return p.z
}

// DistanceTo computes the distance between this point and another point.
func (p *Point) DistanceTo(other *Point) float64 {
	var (
		dx = other.x - p.x
		dy = other.y - p.y
		dz = other.z - p.z
	)

	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

// Displaced cretes a new point, result of displacing this one a given vector a given number of times.
func (p *Point) Displaced(vector *Vector, times float64) *Point {
	return MakePoint(
		p.x+vector.x*times,
		p.y+vector.y*times,
		p.z+vector.z*times,
	)
}

// VectorTo creates a vector goint from this point to another given one.
func (from *Point) VectorTo(to *Point) *Vector {
	return MakeVector(
		to.x-from.x,
		to.y-from.y,
		to.z-from.z,
	)
}

func (p *Point) Equals(other *Point) bool {
	return projectablesEqual(p, other)
}
