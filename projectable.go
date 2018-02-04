package inkgeom

import (
	"math"

	"github.com/angelsolaorbaiceta/inkmath"
)

// Projectable is an entity with projections both in X and Y axis.
type Projectable struct {
	X, Y float64
}

/* ::::::::::::::: Construction ::::::::::::::: */

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
	return Projectable{to.X - from.X, to.Y - from.Y}
}

// MakeVersor creates a vector with unitary norm following the direction of the given projections.
func MakeVersor(x, y float64) Projectable {
	norm := normForProjs(x, y)
	return Projectable{x / norm, y / norm}
}

/* ::::::::::::::: Properties ::::::::::::::: */

// Norm returns the magnitude of the vector.
func (p Projectable) Norm() float64 {
	return normForProjs(p.X, p.Y)
}

// IsVersor returns true if the vector has a norm of 1.
func (p Projectable) IsVersor() bool {
	return inkmath.FuzzyEqual(p.Norm(), 1.0)
}

/* ::::::::::::::: Methods ::::::::::::::: */

// Equals returns true if the projections of this and other projectable are equal.
func (p Projectable) Equals(other Projectable) bool {
	return inkmath.FuzzyEqual(p.X, other.X) && inkmath.FuzzyEqual(p.Y, other.Y)
}

// DistanceTo computes the distance between two points.
func (p Projectable) DistanceTo(other Projectable) float64 {
	deltaX, deltaY := p.X-other.X, p.Y-other.Y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

// ToVersor returns a versor with the same direction as the vector.
func (p Projectable) ToVersor() Projectable {
	return MakeVersor(p.X, p.Y)
}

// Perpendicular returns the vector result of rotating PI/2 radians this one.
func (p Projectable) Perpendicular() Projectable {
	return MakeVector(-p.Y, p.X)
}

// Scaled creates a new vector with the projections scaled the given factor.
func (p Projectable) Scaled(factor float64) Projectable {
	return MakeVector(p.X*factor, p.Y*factor)
}

/* ::::::::::::::: Operations ::::::::::::::: */

// Plus creates a new projectable adding this with other.
func (p Projectable) Plus(other Projectable) Projectable {
	return MakeVector(p.X+other.X, p.Y+other.Y)
}

// Minus creates a new projectable subtracting this with other.
func (p Projectable) Minus(other Projectable) Projectable {
	return MakeVector(p.X-other.X, p.Y-other.Y)
}

// DotTimes computes the dot product of this vector with other.
func (p Projectable) DotTimes(other Projectable) float64 {
	return (p.X * other.X) + (p.Y * other.Y)
}

// CrossTimes computes the cross product of this vector with other.
func (p Projectable) CrossTimes(other Projectable) float64 {
	return (p.X * other.Y) - (p.Y * other.X)
}

/* ::::::::::::::: Utils ::::::::::::::: */
func normForProjs(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

/* ::::::::::::::: Comparing ::::::::::::::: */

/*
Compare returns -1 if this node goes before the other, 0 if they are equal and 1 if
this node goes after the other.
*/
func (p Projectable) Compare(other Projectable) int {
	if p.Equals(other) {
		return 0
	}

	if inkmath.FuzzyEqual(p.X, other.X) {
		if p.Y < other.Y {
			return -1
		}
		return 1
	}

	if p.X < other.X {
		return -1
	}
	return 1
}
