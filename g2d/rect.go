package g2d

import "fmt"

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

// Left is the left-most X coordinate of the rectangle.
func (r *Rect) Left() float64 {
	return r.origin.x
}

// Right is the right-most X coordinate of the rectangle.
func (r *Rect) Right() float64 {
	return r.origin.x + r.width
}

// Bottom is the low-most Y coordinate of the rectangle.
func (r *Rect) Bottom() float64 {
	return r.origin.y
}

// Top is the top-most Y coordinate of the rectangle.
func (r *Rect) Top() float64 {
	return r.origin.y + r.height
}

// ContainsPoint checks whether this rectangle contains the given point.
// Edges are not considered part of the rectangle, so a point on an edge is considered to be
// outside of the rectangle.
func (r *Rect) ContainsPoint(point *Point) bool {
	return point.x > r.Left() &&
		point.x < r.Right() &&
		point.y > r.Bottom() &&
		point.y < r.Top()
}

func (r *Rect) String() string {
	return fmt.Sprintf("Rect{origin: %v, size: (%f, %f)}", &r.origin, r.width, r.height)
}
