package g2d

// A Rect is a two dimensional rectangle whose edges are horizontal and vertical.
// A Rect is defined by an origin point and two numbers: the width and height.
type Rect struct {
	origin Point
	width  float64
	height float64
}

// MakeRect creates a new rectangle given an origin point, the width and height.
func MakeRect(origin Point, width, height float64) *Rect {
	return &Rect{
		origin: origin,
		width:  width,
		height: height,
	}
}

// The Origin of the rectangle
func (r *Rect) Origin() *Point {
	return &r.origin
}

// The Width of the rectangle: the horizontal dimension.
func (r *Rect) Width() float64 {
	return r.width
}

// The Height of the rectangle: the vertical dimension.
func (r *Rect) Height() float64 {
	return r.height
}
