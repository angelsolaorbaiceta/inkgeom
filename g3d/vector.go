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
	Zero       = MakeVector(0, 0, 0)
)

var (
	// ErrZeroVersor happens when a versor is created either from all zero projections or by
	// normalizing a zero vector.
	ErrZeroVersor = errors.New("can't create a versor if all components are zero")
	// ErrZeroVector results from an operation that requires a vector with a non-zero length.
	ErrZeroVector = errors.New("can't use a vector with zero length")
)

// A Vector is a direction with length in space.
// Vectors have three projections: Z, Y and Z.
type Vector struct {
	x, y, z float64
}

// MakeVector creates a new vector given the X, Y and Z projections.
func MakeVector(x, y, z float64) *Vector {
	return &Vector{x, y, z}
}

// MakeNonZeroVector creates a new vector given the X, Y and Z projections, or returns an ErrZeroVector
// error if the three of them are zero, as this would result in a vector with zero length.
func MakeNonZeroVector(x, y, z float64) (*Vector, error) {
	length := computeLength(x, y, z)

	if nums.IsCloseToZero(length) {
		return nil, ErrZeroVector
	}

	return MakeVector(x, y, z), nil
}

// MakeVersor creates a versor (a vector of unit length) given the vector components X, Y and Z.
// Returns an error if all three components are zero, as the zero vector can't be normalized.
func MakeVersor(x, y, z float64) (*Vector, error) {
	length := computeLength(x, y, z)

	if nums.IsCloseToZero(length) {
		return nil, ErrZeroVersor
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

// ToPoint returns a point with the same X, Y and Z projections as this vector.
func (v *Vector) ToPoint() *Point {
	return MakePoint(v.x, v.y, v.z)
}

// Opposite returns a new vector in the opposite direction as this one.
func (v *Vector) Opposite() *Vector {
	return MakeVector(-v.x, -v.y, -v.z)
}

// Scaled creates a new vector with the projections scaled the given factor.
func (v *Vector) Scaled(factor float64) *Vector {
	return MakeVector(v.x*factor, v.y*factor, v.z*factor)
}

// Squared returns a matrix result of multiplying this vector by its transpose.
func (v *Vector) Squared() *Matrix3x3 {
	var (
		x = v.x
		y = v.y
		z = v.z
	)

	return Make3x3Matrix(
		x*x, x*y, x*z,
		y*x, y*y, y*z,
		z*x, z*y, z*z,
	)
}

// IsParalleTo checks whether this and other vectors have the same direction (are parallel).
func (v *Vector) IsParallelTo(other *Vector) bool {
	return v.CrossTimes(other).IsZero()
}

// IsPerpendicularTo checks whether this and other vectors have perpendicular directions.
func (v *Vector) IsPerpendicularTo(other *Vector) bool {
	return nums.IsCloseToZero(v.DotTimes(other))
}

// Equals checks whether this and other vector have equal X, Y and Z projections.
func (v *Vector) Equals(other *Vector) bool {
	return projectablesEqual(v, other)
}

func computeLength(x, y, z float64) float64 {
	return math.Sqrt(x*x + y*y + z*z)
}
