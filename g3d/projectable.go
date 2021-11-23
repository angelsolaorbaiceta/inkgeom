package g3d

import (
	"math"

	"github.com/angelsolaorbaiceta/inkgeom"
)

/*
Projectable is an entity with projections in the X, Y and Z axes.
Used to represent both points and vectors in two dimensions.
*/
type Projectable struct {
	x, y, z float64
}

/* <-- Construction --> */

func MakePoint(x, y, z float64) Projectable {
	return Projectable{x, y, z}
}

func MakeVector(x, y, z float64) Projectable {
	return Projectable{x, y, z}
}

func MakeVersor(x, y, z float64) Projectable {
	length := computeLength(x, y, z)
	return Projectable{x / length, y / length, z / length}
}

/* <-- Properties --> */

func (p Projectable) X() float64 {
	return p.x
}

func (p Projectable) Y() float64 {
	return p.y
}

func (p Projectable) Z() float64 {
	return p.z
}

// Length is the magnitude of the vector
func (p Projectable) Length() float64 {
	return computeLength(p.x, p.y, p.z)
}

// IsVersor returns true if the vector has a length of 1.
func (p Projectable) IsVersor() bool {
	return inkgeom.IsCloseToOne(p.Length())
}

/* <-- Utils --> */

func computeLength(x, y, z float64) float64 {
	return math.Sqrt(x*x + y*y + z*z)
}
