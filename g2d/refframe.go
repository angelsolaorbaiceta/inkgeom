package g2d

// RefFrame represents an orthonormal reference frame in two dimensions.
type RefFrame struct {
	iVersor, jVersor *Vector
}

// MakeRefFrameWithIVersor returns a reference frame with unitary and normal versors from
// a given vector which gives the direction of the i versor.
func MakeRefFrameWithIVersor(iVersor *Vector) *RefFrame {
	i := iVersor.ToVersor()
	return &RefFrame{iVersor: i, jVersor: i.Perpendicular()}
}

// ProjectVector returns the projection of a vector (given in global coordinates) in this
// reference frame.
func (r *RefFrame) ProjectVector(p *Vector) *Vector {
	return MakeVector(p.DotTimes(r.iVersor), p.DotTimes(r.jVersor))
}

// ProjectProjections returns the projectio of a vector given by its projections in global
// coordinates in this reference frame.
func (r *RefFrame) ProjectProjections(xProj, yProj float64) *Vector {
	return r.ProjectVector(MakeVector(xProj, yProj))
}

// ProjectVectorToGlobal returns the projection of a local vector (vector projected in this
// reference frame), in the global reference frame.
func (r *RefFrame) ProjectVectorToGlobal(p *Vector) *Vector {
	return r.ProjectionsToGlobal(p.X(), p.Y())
}

// ProjectionsToGlobal returns the projection of a local vector (vector projected in this
// reference frame), in the global reference frame.
func (r *RefFrame) ProjectionsToGlobal(xProj, yProj float64) *Vector {
	var (
		x = r.iVersor.Scaled(xProj)
		y = r.jVersor.Scaled(yProj)
	)

	return x.Plus(y)
}

// Cos returns the cosine of the angle between this frame's i versor and te global
// frame i's versor.
func (r *RefFrame) Cos() float64 {
	return r.iVersor.DotTimes(IVersor)
}

// Sin returns the sine of the angle between this frame's i versor and the global
// frame's i versor
func (r *RefFrame) Sin() float64 {
	return r.iVersor.DotTimes(JVersor)
}

// AngleInRadsFromX returns the angle in radians from the global frame's i versor.
// The angle returned is in the range [-π, π].
func (r *RefFrame) AngleInRadsFromX() float64 {
	return r.iVersor.AngleInRadsFromX()
}
