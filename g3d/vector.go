package g3d

import (
	"errors"
	"math"

	"github.com/angelsolaorbaiceta/inkgeom/nums"
)

var (
	IVersor, _ = MakeVersor(1, 0, 0)
	JVersor, _ = MakeVersor(0, 1, 0)
	KVersor, _ = MakeVersor(0, 0, 1)
)

var errZeroVersor = errors.New("can't create a versor if all components are zero")

// A Vector is a direction with length in space.
type Vector struct {
	x, y, z float64
}

func MakeVector(x, y, z float64) *Vector {
	return &Vector{x, y, z}
}

// MakeVersor creates a versor (a vector of unit length) given the vector components X, Y and Z.
// Returns an error if all three components are zero, as the zero vector can't be normalized.
func MakeVersor(x, y, z float64) (*Vector, error) {
	length := computeLength(x, y, z)

	if nums.IsCloseToZero(length) {
		return nil, errZeroVersor
	}

	return &Vector{x / length, y / length, z / length}, nil
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

// IsZero returns true if all X, Y and Z componets of this vector are zero.
func (v *Vector) IsZero() bool {
	return nums.IsCloseToZero(v.x) && nums.IsCloseToZero(v.y) && nums.IsCloseToZero(v.z)
}

// ToVersor returns a versor with the same direction as this vector.
// Returns an error if all three components are zero, as the zero vector can't be normalized.
func (v *Vector) ToVersor() (*Vector, error) {
	if v.IsVersor() {
		return v, nil
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
