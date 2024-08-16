package g2d

import (
	"math"
	"testing"

	"github.com/angelsolaorbaiceta/inkgeom/nums"
	"github.com/stretchr/testify/assert"
)

func TestMakeSegmentFromCoords(t *testing.T) {
	sx, sy, ex, ey := 1.0, 2.0, 3.0, 4.0
	seg := MakeSegmentFromCoords(sx, sy, ex, ey)

	assert.Equal(t, sx, seg.Start().X())
	assert.Equal(t, sy, seg.Start().Y())
	assert.Equal(t, ex, seg.End().X())
	assert.Equal(t, ey, seg.End().Y())
}

func TestSegmentLength(t *testing.T) {
	p, q := MakePoint(1, 2), MakePoint(4, 6)
	seg := MakeSegment(p, q)

	assert.True(
		t,
		nums.FloatsEqual(seg.Length(), 5.0),
		"Expected segment length: 5.0, Actual segment length: %f", seg.Length(),
	)
}

func TestSegmentLengthBetweenPositions(t *testing.T) {
	p, q := MakePoint(1, 2), MakePoint(4, 6)
	st, et := nums.MakeTParam(0.25), nums.MakeTParam(0.75)
	seg := MakeSegment(p, q)
	got := seg.LengthBetween(st, et)

	assert.True(
		t,
		nums.FloatsEqual(got, 2.5),
		"Expected segment length: 2.5, Actual segment length: %f", got,
	)
}

func TestPointAtPosition(t *testing.T) {
	seg := MakeSegment(MakePoint(1, 1), MakePoint(3, 5))
	pos := nums.MakeTParam(0.5)
	expectedPoint := MakePoint(2, 3)

	assert.True(
		t,
		seg.PointAt(pos).Equals(expectedPoint),
	)
}

func TestDirectionVersor(t *testing.T) {
	seg := MakeSegment(MakePoint(1, 1), MakePoint(3, 5))
	expectedDirVersor := MakeVersor(1, 2)

	assert.True(t, seg.DirectionVersor().Equals(expectedDirVersor))
}

func TestAngleInRadsFromX(t *testing.T) {
	t.Run("Angle when i = (1, 0) is 0", func(t *testing.T) {
		var (
			ref   = MakeRefFrameWithIVersor(MakeVector(1, 0))
			angle = ref.AngleInRadsFromX()
		)

		assert.True(t, nums.FloatsEqual(angle, 0.0), "Expected 0, got %f", angle)
	})

	t.Run("Angle when i = (1, 1) is PI/4", func(t *testing.T) {
		var (
			ref   = MakeRefFrameWithIVersor(MakeVector(1, 1))
			angle = ref.AngleInRadsFromX()
		)
		assert.True(t, nums.FloatsEqual(angle, math.Pi/4), "Expected PI/4, got %f", angle)
	})

	t.Run("Angle when i = (0, 1) is PI/2", func(t *testing.T) {
		var (
			ref   = MakeRefFrameWithIVersor(MakeVector(0, 1))
			angle = ref.AngleInRadsFromX()
		)
		assert.True(t, nums.FloatsEqual(angle, math.Pi/2), "Expected PI/2, got %f", angle)
	})

	t.Run("Angle when i = (-1, 1) is 3*PI/4", func(t *testing.T) {
		var (
			ref   = MakeRefFrameWithIVersor(MakeVector(-1, 1))
			angle = ref.AngleInRadsFromX()
		)
		assert.True(t, nums.FloatsEqual(angle, 3*math.Pi/4), "Expected 3*PI/4, got %f", angle)
	})

	t.Run("Angle when i = (-1, 0) is PI", func(t *testing.T) {
		var (
			ref   = MakeRefFrameWithIVersor(MakeVector(-1, 0))
			angle = ref.AngleInRadsFromX()
		)
		assert.True(t, nums.FloatsEqual(angle, math.Pi), "Expected PI, got %f", angle)
	})

	t.Run("Angle when i = (-1, -1) is -3*PI/4", func(t *testing.T) {
		var (
			ref   = MakeRefFrameWithIVersor(MakeVector(-1, -1))
			angle = ref.AngleInRadsFromX()
		)
		assert.True(t, nums.FloatsEqual(angle, -3*math.Pi/4), "Expected -3*PI/4, got %f", angle)
	})

	t.Run("Angle when i = (0, -1) is -PI/2", func(t *testing.T) {
		var (
			ref   = MakeRefFrameWithIVersor(MakeVector(0, -1))
			angle = ref.AngleInRadsFromX()
		)
		assert.True(t, nums.FloatsEqual(angle, -math.Pi/2), "Expected -PI/2, got %f", angle)
	})

	t.Run("Angle when i = (1, -1) is -PI/4", func(t *testing.T) {
		var (
			ref   = MakeRefFrameWithIVersor(MakeVector(1, -1))
			angle = ref.AngleInRadsFromX()
		)
		assert.True(t, nums.FloatsEqual(angle, -math.Pi/4), "Expected -PI/4, got %f", angle)
	})
}
