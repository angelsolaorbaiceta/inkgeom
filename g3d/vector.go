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

/* <-- Construction --> */

func MakeVector(x, y, z float64) *Vector {
	return &Vector{x, y, z}
}

func MakeVersor(x, y, z float64) *Vector {
	length := computeLength(x, y, z)
	return &Vector{x / length, y / length, z / length}
}

/* <-- Properties --> */

func (v *Vector) X() float64 {
	return v.x
}

func (v *Vector) Y() float64 {
	return v.y
}

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

/* <-- Methods --> */

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

/* <-- Utils --> */

func computeLength(x, y, z float64) float64 {
	return math.Sqrt(x*x + y*y + z*z)
}
