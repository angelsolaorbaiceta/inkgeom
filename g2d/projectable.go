package g2d

import (
	"fmt"
	"math"

	"github.com/angelsolaorbaiceta/inkgeom/nums"
)

/*
Projectable is an entity with projections both in the X and Y axis.
Used to represent both points and vectors in two dimensions.
*/
type Projectable struct {
	x, y float64
}

// MakePoint creates a new point.
func MakePoint(x, y float64) Projectable {
	return Projectable{x, y}
}

// MakeVector creates a new vector.
func MakeVector(x, y float64) Projectable {
	return Projectable{x, y}
}

// MakeVectorFromTo creates a new vector which goes from one point to another.
func MakeVectorFromTo(from, to Projectable) Projectable {
	return Projectable{to.x - from.x, to.y - from.y}
}

// MakeVersor creates a vector with unitary norm following the direction of the given projections.
func MakeVersor(x, y float64) Projectable {
	length := computeLength(x, y)
	return Projectable{x / length, y / length}
}

/* <-- Properties --> */

func (p Projectable) X() float64 {
	return p.x
}

func (p Projectable) Y() float64 {
	return p.y
}

// Length returns the magnitude of the vector.
func (p Projectable) Length() float64 {
	return computeLength(p.x, p.y)
}

// IsVersor returns true if the vector has a length of 1.
func (p Projectable) IsVersor() bool {
	return nums.IsCloseToOne(p.Length())
}

/* <-- Methods --> */

// Equals returns true if the projections of this and other projectable are equal.
func (p Projectable) Equals(other Projectable) bool {
	return nums.FloatsEqual(p.x, other.x) &&
		nums.FloatsEqual(p.y, other.y)
}

// DistanceTo computes the distance between two points.
func (p Projectable) DistanceTo(other Projectable) float64 {
	var (
		deltaX = p.x - other.x
		deltaY = p.y - other.y
	)

	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

// ToVersor returns a versor with the same direction as the vector.
func (p Projectable) ToVersor() Projectable {
	return MakeVersor(p.x, p.y)
}

// Perpendicular returns the vector result of rotating PI/2 radians this one.
func (p Projectable) Perpendicular() Projectable {
	return MakeVector(-p.y, p.x)
}

// Scaled creates a new vector with the projections scaled the given factor.
func (p Projectable) Scaled(factor float64) Projectable {
	return MakeVector(p.x*factor, p.y*factor)
}

/* <-- Operations --> */

// Plus creates a new projectable adding this with other.
func (p Projectable) Plus(other Projectable) Projectable {
	return MakeVector(p.x+other.x, p.y+other.y)
}

// Minus creates a new projectable subtracting this with other.
func (p Projectable) Minus(other Projectable) Projectable {
	return MakeVector(p.x-other.x, p.y-other.y)
}

// DotTimes computes the dot product of this vector with other.
func (p Projectable) DotTimes(other Projectable) float64 {
	return (p.x * other.x) + (p.y * other.y)
}

// CrossTimes computes the cross product of this vector with other.
func (p Projectable) CrossTimes(other Projectable) float64 {
	return (p.x * other.y) - (p.y * other.x)
}

/* <-- Utils --> */

func computeLength(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

/* <-- Comparing --> */

/*
Compare returns -1 if this node goes before the other, 0 if they are equal and 1 if
this node goes after the other.
*/
func (p Projectable) Compare(other Projectable) int {
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

/* <-- Stringer --> */

func (p Projectable) String() string {
	return fmt.Sprintf("{%f, %f}", p.x, p.y)
}
