package inkgeom

var (
	globalIVersor = MakeVersor(1, 0)
	globalJVersor = MakeVersor(0, 1)
)

/*
RefFrame represents an orthonormal reference frame.
*/
type RefFrame struct {
	iVersor, jVersor Projectable
}

/* ::::::::::::::: Construction ::::::::::::::: */

/*
MakeRefFrameWithIVersor returns a reference frame with unitary and normal versors from
a given vector which gives the direction of the i versor.
*/
func MakeRefFrameWithIVersor(iVersor Projectable) RefFrame {
	i := iVersor.ToVersor()
	return RefFrame{iVersor: i, jVersor: i.Perpendicular()}
}

/* ::::::::::::::: Methods ::::::::::::::: */

// ProjectVector returns the projection of a vector (given in global coordinates) in this reference frame.
func (r RefFrame) ProjectVector(p Projectable) Projectable {
	return MakeVector(p.DotTimes(r.iVersor), p.DotTimes(r.jVersor))
}

/*
ProjectVectorToGlobal returns the projection of a local vector (vector projected in this reference frame),
in the global reference frame.
*/
func (r RefFrame) ProjectVectorToGlobal(p Projectable) Projectable {
	return r.ProjectionsToGlobal(p.X, p.Y)
}

/*
ProjectionsToGlobal returns the projection of a local vector (vector projected in this reference frame),
in the global reference frame.
*/
func (r RefFrame) ProjectionsToGlobal(xProj, yProj float64) Projectable {
	var (
		x = r.iVersor.Scaled(xProj)
		y = r.jVersor.Scaled(yProj)
	)

	return x.Plus(y)
}

/*
Cos returns the cosine of the angle between:
	- this frame's i versor and
	- global frame i's versor.
*/
func (r RefFrame) Cos() float64 {
	return r.iVersor.DotTimes(globalIVersor)
}

/*
Sin returns the sine of the angle between:
	- this frame's i versor and
	- global frame's i versor
*/
func (r RefFrame) Sin() float64 {
	return r.iVersor.DotTimes(globalJVersor)
}
