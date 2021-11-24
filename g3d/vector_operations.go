package g3d

// Plus creates a new projectable adding this with other.
func (addend *Vector) Plus(augend *Vector) *Vector {
	return MakeVector(
		addend.x+augend.x,
		addend.y+augend.y,
		addend.z+augend.z,
	)
}

// Minus creates a new projectable subtracting this with other.
func (minuend *Vector) Minus(subtrahend *Vector) *Vector {
	return MakeVector(
		minuend.x-subtrahend.x,
		minuend.y-subtrahend.y,
		minuend.z-subtrahend.z,
	)
}

// DotTimes computes the dot product of this vector with other.
func (multiplicand *Vector) DotTimes(multiplier *Vector) float64 {
	return (multiplicand.x * multiplier.x) +
		(multiplicand.y * multiplier.y) +
		(multiplicand.z * multiplier.z)
}

// CrossTimes computes the cross product of this vector with other.
func (multiplicand *Vector) CrossTimes(multiplier *Vector) *Vector {
	return MakeVector(
		(multiplicand.y*multiplier.z)-(multiplicand.z*multiplier.y),
		(multiplicand.z*multiplier.x)-(multiplicand.x*multiplier.z),
		(multiplicand.x*multiplier.y)-(multiplicand.y*multiplier.x),
	)
}
