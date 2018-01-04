package inkgeom

import (
	"math"
	"github.com/angelsolaorbaiceta/inkmath"
)

type Projectable struct {
	X, Y float64
}

/* Construction */
func MakePoint(x, y float64) Projectable {
	return Projectable{x, y}
}

func MakeVector(x, y float64) Projectable {
	return Projectable{x, y}
}

func MakeVectorFromTo(from, to Projectable) Projectable {
	return Projectable{to.X - from.X, to.Y - from.Y}
}

func MakeVersor(x, y float64) Projectable {
	norm := normForProjs(x, y)
	return Projectable{x / norm, y / norm}
}

/* Properties */
func (p Projectable) Norm() float64 {
	return normForProjs(p.X, p.Y)
}

// Returns true if the vector has a norm of 1.
func (p Projectable) IsVersor() bool {
	return inkmath.FuzzyEqual(p.Norm(), 1.0)
}

/* Methods */
func (p Projectable) Equals(other Projectable) bool {
	return inkmath.FuzzyEqual(p.X, other.X) && inkmath.FuzzyEqual(p.Y, other.Y)
}

// Computes the distance between two points.
func (p Projectable) DistanceTo(other Projectable) float64 {
	deltaX, deltaY := p.X - other.X, p.Y - other.Y
	return math.Sqrt(deltaX * deltaX + deltaY * deltaY)
}

// Returns a versor with the same direction as the vector.
func (p Projectable) ToVersor() Projectable {
	return MakeVersor(p.X, p.Y)
}

// Returns a vector which is the result of rotating PI/2 radians this one.
func (p Projectable) Perpendicular() Projectable {
	return MakeVector(-p.Y, p.X)
}

/* Operations */
func (p Projectable) Plus(other Projectable) Projectable {
	return MakeVector(p.X + other.X, p.Y + other.Y)
}

func (p Projectable) Minus(other Projectable) Projectable {
	return MakeVector(p.X - other.X, p.Y - other.Y)
}

func (p Projectable) DotTimes(other Projectable) float64 {
	return (p.X * other.X) + (p.Y * other.Y)
}

func (p Projectable) CrossTimes(other Projectable) float64 {
	return (p.X * other.Y) - (p.Y * other.X)
}

/* Utils */
func normForProjs(x, y float64) float64 {
	return math.Sqrt(x * x + y * y)
}
