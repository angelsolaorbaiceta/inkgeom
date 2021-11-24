package g3d

import "github.com/angelsolaorbaiceta/inkgeom/nums"

// A Projectable is an entity with immutable projections in the X, Y and Z axes.
type Projectable interface {
	X() float64
	Y() float64
	Z() float64
}

func projectablesEqual(a, b Projectable) bool {
	return nums.FloatsEqual(a.X(), b.X()) &&
		nums.FloatsEqual(a.Y(), b.Y()) &&
		nums.FloatsEqual(a.Z(), b.Z())
}
