package g2d

import (
	"testing"

	"github.com/angelsolaorbaiceta/inkgeom/nums"
	"github.com/stretchr/testify/assert"
)

func TestCreateVector(t *testing.T) {
	var (
		x, y = 2.0, 3.0
		v    = MakeVector(x, y)
	)

	t.Run("has a X projection", func(t *testing.T) {
		assert.Equal(t, x, v.X())
	})

	t.Run("has a Y projection", func(t *testing.T) {
		assert.Equal(t, y, v.Y())
	})
}

func TestVectorLength(t *testing.T) {
	v := MakeVector(3, 4)

	assert.True(t, nums.FloatsEqual(v.Length(), 5.0))
}

func TestIsVersor(t *testing.T) {
	v := MakeVector(1, 0)

	assert.True(t, v.IsVersor())
}

func TestIsNotVersor(t *testing.T) {
	v := MakeVector(1, 2)

	assert.False(t, v.IsVersor())
}

func TestMakeVersor(t *testing.T) {
	v := MakeVersor(1, 2)

	assert.True(t, v.IsVersor())
}

func TestPerpendicular(t *testing.T) {
	v := MakeVersor(1, 2)

	assert.True(t, v.Perpendicular().Equals(MakeVersor(-2, 1)))
}

/* Operations */
func TestSumVectors(t *testing.T) {
	v := MakeVector(1, 2).Plus(MakeVector(3, 4))

	assert.True(t, v.Equals(MakeVector(4, 6)))
}

func TestSubtractVectors(t *testing.T) {
	v := MakeVector(1, 2).Minus(MakeVector(3, 5))

	assert.True(t, v.Equals(MakeVector(-2, -3)))
}

func TestVectorsDotProduct(t *testing.T) {
	p := MakeVector(1, 2).DotTimes(MakeVector(3, 5))

	assert.True(t, nums.FloatsEqual(p, 13.0))
}

func TestVectorsCrossProduct(t *testing.T) {
	p := MakeVector(1, 2).CrossTimes(MakeVector(3, 5))

	assert.True(t, nums.FloatsEqual(p, -1.0))
}
