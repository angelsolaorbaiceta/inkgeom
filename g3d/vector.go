package g3d

import (
	"math"

	"github.com/angelsolaorbaiceta/inkgeom/nums"
)

var (
	IVersor = MakeVersor(1, 0, 0)
	JVersor = MakeVersor(0, 1, 0)
	KVersor = MakeVersor(0, 0, 1)
)

// A Vector is a direction with length in space.
type Vector struct {
	x, y, z float64
}

func MakeVector(x, y, z float64) *Vector {
	return &Vector{x, y, z}
}

func MakeVersor(x, y, z float64) *Vector {
	length := computeLength(x, y, z)
	return &Vector{x / length, y / length, z / length}
}

// X is the vector's projection in the X axis.
func (v *Vector) X() float64 {
	return v.x
}

// Y is the vector's projection in the Y axis.
func (v *Vector) Y() float64 {
	return v.y
}

// Z is the vector's projection in the Z axis.
func (v *Vector) Z() float64 {
	return v.z
}

// Length is the magnitude of the vector
func (v *Vector) Length() float64 {
	return computeLength(v.x, v.y, v.z)
}

// IsVersor evaluates to true if the vector has a length of 1.
func (v *Vector) IsVersor() bool {
	return nums.IsCloseToOne(v.Length())
}

// ToVersor returns a versor with the same direction as this vector.
func (v *Vector) ToVersor() *Vector {
	if v.IsVersor() {
		return v
	}

	return MakeVersor(v.x, v.y, v.z)
}

// Opposite returns a new vector in the opposite direction as this one.
func (v *Vector) Opposite() *Vector {
	return MakeVector(-v.x, -v.y, -v.z)
}

// Scaled creates a new vector with the projections scaled the given factor.
func (v *Vector) Scaled(factor float64) *Vector {
	return MakeVector(v.x*factor, v.y*factor, v.z*factor)
}

func (v *Vector) Equals(other *Vector) bool {
	return projectablesEqual(v, other)
}

func computeLength(x, y, z float64) float64 {
	return math.Sqrt(x*x + y*y + z*z)
}
