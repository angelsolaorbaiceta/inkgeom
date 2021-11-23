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

/*
Projectable is an entity with projections in the X, Y and Z axes.
Used to represent both points and vectors in two dimensions.
*/
type Projectable struct {
	x, y, z float64
}

/* <-- Construction --> */

func MakePoint(x, y, z float64) *Projectable {
	return &Projectable{x, y, z}
}

func MakeVector(x, y, z float64) *Projectable {
	return &Projectable{x, y, z}
}

func MakeVersor(x, y, z float64) *Projectable {
	length := computeLength(x, y, z)
	return &Projectable{x / length, y / length, z / length}
}

/* <-- Properties --> */

func (p *Projectable) X() float64 {
	return p.x
}

func (p *Projectable) Y() float64 {
	return p.y
}

func (p *Projectable) Z() float64 {
	return p.z
}

// Length is the magnitude of the vector
func (p *Projectable) Length() float64 {
	return computeLength(p.x, p.y, p.z)
}

// IsVersor returns true if the vector has a length of 1.
func (p *Projectable) IsVersor() bool {
	return nums.IsCloseToOne(p.Length())
}

/* <-- Methods --> */

// Opposite returns a new vector in the opposite direction as this one.
func (p *Projectable) Opposite() *Projectable {
	return MakeVector(-p.x, -p.y, -p.z)
}

// ToVersor returns a versor with the same direction as this vector.
func (p *Projectable) ToVersor() *Projectable {
	if p.IsVersor() {
		return p
	}

	return MakeVersor(p.x, p.y, p.y)
}

// Equals evaluates to true if the two projectables have identical X, Y and Z projections.
func (p *Projectable) Equals(other *Projectable) bool {
	return nums.FloatsEqual(p.x, other.x) &&
		nums.FloatsEqual(p.y, other.y) &&
		nums.FloatsEqual(p.z, other.z)
}

/* <-- Operations --> */

// Plus creates a new projectable adding this with other.
func (addend *Projectable) Plus(augend *Projectable) *Projectable {
	return MakeVector(
		addend.x+augend.x,
		addend.y+augend.y,
		addend.z+augend.z,
	)
}

// Minus creates a new projectable subtracting this with other.
func (minuend *Projectable) Minus(subtrahend *Projectable) *Projectable {
	return MakeVector(
		minuend.x-subtrahend.x,
		minuend.y-subtrahend.y,
		minuend.z-subtrahend.z,
	)
}

// DotTimes computes the dot product of this vector with other.
func (multiplicand *Projectable) DotTimes(multiplier *Projectable) float64 {
	return (multiplicand.x * multiplier.x) +
		(multiplicand.y * multiplier.y) +
		(multiplicand.z * multiplier.z)
}

// CrossTimes computes the cross product of this vector with other.
func (u *Projectable) CrossTimes(v *Projectable) *Projectable {
	return MakeVector(
		(u.y*v.z)-(u.z*v.y),
		(u.z*v.x)-(u.x*v.z),
		(u.x*v.y)-(u.y*v.x),
	)
}

/* <-- Utils --> */

func computeLength(x, y, z float64) float64 {
	return math.Sqrt(x*x + y*y + z*z)
}
