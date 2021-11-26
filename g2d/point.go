package g2d

import (
	"fmt"
	"math"

	"github.com/angelsolaorbaiceta/inkgeom/nums"
)

type Point struct {
	x, y float64
}

// MakePoint creates a new point with the given X and Y coordinates.
func MakePoint(x, y float64) *Point {
	return &Point{x, y}
}

func (p *Point) X() float64 {
	return p.x
}

func (p *Point) Y() float64 {
	return p.y
}

// DistanceTo computes the distance between two points.
func (p *Point) DistanceTo(other *Point) float64 {
	var (
		dX = p.x - other.x
		dY = p.y - other.y
	)

	return math.Sqrt(dX*dX + dY*dY)
}

func (from *Point) VectorTo(to *Point) *Vector {
	return MakeVector(to.x-from.x, to.y-from.y)
}

// Equals returns true if the projections of this and other projectable are equal.
func (p *Point) Equals(other *Point) bool {
	return nums.FloatsEqual(p.x, other.x) &&
		nums.FloatsEqual(p.y, other.y)
}

/*
Compare returns -1 if this node goes before the other, 0 if they are equal and 1 if
this node goes after the other.
*/
func (p *Point) Compare(other *Point) int {
	if p.Equals(other) {
		return 0
	}

	if nums.FloatsEqual(p.x, other.x) {
		if p.y < other.y {
			return -1
		}
		return 1
	}

	if p.x < other.x {
		return -1
	}
	return 1
}

func (p *Point) String() string {
	return fmt.Sprintf("(%f, %f)", p.x, p.y)
}
