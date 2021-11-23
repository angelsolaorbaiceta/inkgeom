package g3d

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
