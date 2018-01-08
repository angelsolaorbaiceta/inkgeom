package inkgeom

/*
RefFrame represents an orthonormal reference frame.
*/
type RefFrame struct {
	iVersor, jVersor Projectable
}

/* Construction */

// MakeRefFrameWithIVersor returns a reference frame with unitary and normal versors from
// a given vector which gives the direction of the i versor.
func MakeRefFrameWithIVersor(iVersor Projectable) RefFrame {
	i := iVersor.ToVersor()
	return RefFrame{iVersor: i, jVersor: i.Perpendicular()}
}

/* Methods */

// ProjectVector returns the projection of a vector (given in global coordinates) in this reference frame.
func (r RefFrame) ProjectVector(p Projectable) Projectable {
	return MakePoint(p.DotTimes(r.iVersor), p.DotTimes(r.jVersor))
}
