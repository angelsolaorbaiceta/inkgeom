package inkgeom

import (
	"testing"

	"github.com/angelsolaorbaiceta/inkgeom"
	"github.com/angelsolaorbaiceta/inkmath/nums"
)

func TestMakeSegmentFromCoords(t *testing.T) {
	sx, sy, ex, ey := 1.0, 2.0, 3.0, 4.0
	seg := MakeSegmentFromCoords(sx, sy, ex, ey)

	if seg.Start.X != sx {
		t.Error("Start X coord not as expected")
	}
	if seg.Start.Y != sy {
		t.Error("Start Y coord not as expected")
	}
	if seg.End.X != ex {
		t.Error("End X coord not as expected")
	}
	if seg.End.Y != ey {
		t.Error("End Y coord not as expected")
	}
}

func TestSegmentLength(t *testing.T) {
	p, q := MakePoint(1, 2), MakePoint(4, 6)
	seg := MakeSegment(p, q)

	if !nums.FuzzyEqual(seg.Length(), 5.0) {
		t.Error("Wrong segment length")
	}
}

func TestSegmentLengthBetweenPositions(t *testing.T) {
	p, q := MakePoint(1, 2), MakePoint(4, 6)
	st, et := inkgeom.MakeTParam(0.25), inkgeom.MakeTParam(0.75)
	seg := MakeSegment(p, q)

	if !nums.FuzzyEqual(seg.LengthBetween(st, et), 2.5) {
		t.Error("Wrong segment length")
	}
}

func TestPointAtPosition(t *testing.T) {
	seg := MakeSegment(MakePoint(1, 1), MakePoint(3, 5))
	pos := inkgeom.MakeTParam(0.5)
	expectedPoint := MakePoint(2, 3)

	if !seg.PointAt(pos).Equals(expectedPoint) {
		t.Error("Wrong point")
	}
}

func TestDirectionVersor(t *testing.T) {
	seg := MakeSegment(MakePoint(1, 1), MakePoint(3, 5))
	expectedDirVersor := MakeVersor(1, 2)

	if !seg.DirectionVersor().Equals(expectedDirVersor) {
		t.Error("Wrong direction versor for segment")
	}
}
