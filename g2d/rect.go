package g2d

import (
	"errors"
	"fmt"
	"math"

	"github.com/angelsolaorbaiceta/inkgeom/nums"
)

// A Rect is a two dimensional rectangle whose edges are horizontal and vertical.
// A Rect is defined by an origin point and two numbers: the width and height.
type Rect struct {
	origin Point
	width  float64
	height float64
}

// MakeRect creates a new rectangle given an origin point, the width and height.
//
// A non-nil error is returned if either the width or height are smaller than zero.
func MakeRect(origin Point, width, height float64) (*Rect, error) {
	if width < 0.0 {
		return nil, errors.New("the width must be greater or equal to zero")
	}

	if height < 0.0 {
		return nil, errors.New("the height must be greater or equal to zero")
	}

	return &Rect{
		origin: origin,
		width:  width,
		height: height,
	}, nil
}

// MakeRectContaining returns the smallest rectangle containing all the given points.
// Some of the input points end up in one of the edges, thus, those points won't pass the
// ContainsPoint test.
//
// A non-nil error is returned if the list of points is empty.
func MakeRectContaining(points []*Point) (*Rect, error) {
	if len(points) < 1 {
		return nil, errors.New("at least one point is required to construct the rectangle")
	}

	var (
		firstPoint = points[0]
		point      *Point
		minX       = firstPoint.x
		maxX       = firstPoint.x
		minY       = firstPoint.y
		maxY       = firstPoint.y
	)

	for i := 1; i < len(points); i++ {
		point = points[i]

		minX = math.Min(minX, point.x)
		maxX = math.Max(maxX, point.x)
		minY = math.Min(minY, point.y)
		maxY = math.Max(maxY, point.y)
	}

	return MakeRect(
		*MakePoint(minX, minY),
		maxX-minX,
		maxY-minY,
	)
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

// WithScaledSize returns a new rectangle resulting from scaling the width and height.
func (r *Rect) WithScaledSize(scale float64) (*Rect, error) {
	return MakeRect(r.origin, r.width*scale, r.height*scale)
}

// Equals checks if this and other rectangle are equal.
func (r *Rect) Equals(other *Rect) bool {
	return r.origin.Equals(&other.origin) &&
		nums.FloatsEqual(r.width, other.width) &&
		nums.FloatsEqual(r.height, other.height)
}

func (r *Rect) String() string {
	return fmt.Sprintf("Rect{origin: %v, size: (%f, %f)}", &r.origin, r.width, r.height)
}
