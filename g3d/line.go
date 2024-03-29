package g3d

// A Line is an infinite set of aligned points in space.
type Line struct {
	origin    *Point
	direction *Vector
}

// MakeLine creates a line passing throught the given point and with the given direction, which is
// normalized (scaled to have unitary length).
// Returns an error if the direction vector has zero length.
func MakeLine(origin *Point, direction *Vector) (*Line, error) {
	dirVersor, err := direction.ToVersor()

	if err != nil {
		return nil, err
	}

	return &Line{
		origin:    origin,
		direction: dirVersor,
	}, nil
}

// The Origin is a base point the line goes through.
func (l *Line) Origin() *Point {
	return l.origin
}

// The Direction of a line is the vector defining the alignment of the line's points.
func (l *Line) Direction() *Vector {
	return l.direction
}

// PointAt yields a point on the line for a given value of the t parameter, which can go from minus
// infinity to infinity.
func (l *Line) PointAt(t float64) *Point {
	return l.origin.Displaced(l.direction, t)
}
