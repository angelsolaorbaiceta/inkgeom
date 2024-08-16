package g2d

import (
	"fmt"
	"math"

	"github.com/angelsolaorbaiceta/inkgeom/nums"
)

var (
	IVersor = MakeVersor(1, 0)
	JVersor = MakeVersor(0, 1)
)

// Vector is a direction in the plane represented by two projections: in the X and Y axes.
type Vector struct {
	x, y float64
}

// MakeVector creates a new vector.
func MakeVector(x, y float64) *Vector {
	return &Vector{x, y}
}

// MakeVersor creates a vector with unitary norm following the direction of the given projections.
func MakeVersor(x, y float64) *Vector {
	length := computeLength(x, y)
	return &Vector{x / length, y / length}
}

func (v *Vector) X() float64 {
	return v.x
}

func (v *Vector) Y() float64 {
	return v.y
}

// Length returns the magnitude of the vector.
func (v *Vector) Length() float64 {
	return computeLength(v.x, v.y)
}

// IsVersor returns true if the vector has a length of 1.
func (v *Vector) IsVersor() bool {
	return nums.IsCloseToOne(v.Length())
}

// Equals returns true if the projections of this and other vector are equal.
func (v *Vector) Equals(other *Vector) bool {
	return nums.FloatsEqual(v.x, other.x) &&
		nums.FloatsEqual(v.y, other.y)
}

// ToVersor returns a versor with the same direction as this vector.
func (v *Vector) ToVersor() *Vector {
	if v.IsVersor() {
		return v
	}

	return MakeVersor(v.x, v.y)
}

// Perpendicular returns the vector result of rotating PI/2 radians this one.
func (v *Vector) Perpendicular() *Vector {
	return MakeVector(-v.y, v.x)
}

// Scaled creates a new vector with the projections scaled the given factor.
func (v *Vector) Scaled(factor float64) *Vector {
	return MakeVector(v.x*factor, v.y*factor)
}

// Plus creates a new vector adding this with other.
func (v *Vector) Plus(other *Vector) *Vector {
	return MakeVector(v.x+other.x, v.y+other.y)
}

// Minus creates a new vector subtracting this with other.
func (v *Vector) Minus(other *Vector) *Vector {
	return MakeVector(v.x-other.x, v.y-other.y)
}

// DotTimes computes the dot product of this vector with other.
func (v *Vector) DotTimes(other *Vector) float64 {
	return (v.x * other.x) + (v.y * other.y)
}

// CrossTimes computes the cross product of this vector with other.
func (v *Vector) CrossTimes(other *Vector) float64 {
	return (v.x * other.y) - (v.y * other.x)
}

// AngleInRadsFromX returns the angle (in radians) between this vector and the X axis.
// The returned angle is in the range [-π, π].
func (v *Vector) AngleInRadsFromX() float64 {
	var (
		versor    = v.ToVersor()
		value     = math.Acos(versor.DotTimes(IVersor))
		crossProd = IVersor.CrossTimes(versor)
	)

	return math.Copysign(value, crossProd)
}

func computeLength(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func (v Vector) String() string {
	return fmt.Sprintf("{%f, %f}", v.x, v.y)
}
