package g3d

import (
	"math"
	"testing"

	"github.com/angelsolaorbaiceta/inkgeom/nums"
	"github.com/stretchr/testify/assert"
)

func TestCreateVector(t *testing.T) {
	assert := assert.New(t)

	var (
		x = 10.0
		y = 20.0
		z = 30.0
		p = Vector{x, y, z}
	)

	t.Run("has a X projection", func(t *testing.T) {
		assert.Equal(x, p.X())
	})

	t.Run("has a Y projection", func(t *testing.T) {
		assert.Equal(y, p.Y())
	})

	t.Run("has a Z projection", func(t *testing.T) {
		assert.Equal(z, p.Z())
	})

	t.Run("has a length", func(t *testing.T) {
		want := math.Sqrt(x*x + y*y + z*z)
		assert.Equal(want, p.Length())
	})
}

func TestVersor(t *testing.T) {
	assert := assert.New(t)

	var (
		x      = 10.0
		y      = 20.0
		z      = 30.0
		v, err = MakeVersor(x, y, z)
	)

	t.Run("is created without error", func(t *testing.T) {
		assert.Nil(err)
	})

	t.Run("has unit length", func(t *testing.T) {
		assert.True(nums.IsCloseToOne(v.Length()))
	})

	t.Run("is versor", func(t *testing.T) {
		assert.True(v.IsVersor())
	})

	t.Run("As versor", func(t *testing.T) {
		vector := MakeVector(x, y, z)

		assert.False(vector.IsVersor())

		versor, err := vector.ToVersor()

		assert.True(versor.IsVersor())
		assert.Nil(err)
	})

	t.Run("Can't be created from three zero components", func(t *testing.T) {
		v, err := MakeVersor(0, 0, 0)

		assert.Nil(v)
		assert.NotNil(err)
	})

	t.Run("Can't create a versor from a zero vector", func(t *testing.T) {
		v, err := MakeVector(0, 0, 0).ToVersor()

		assert.Nil(v)
		assert.NotNil(err)
	})
}

func TestVectorOpposite(t *testing.T) {
	var (
		v    = MakeVector(1, 2, -3)
		want = MakeVector(-1, -2, 3)
	)

	assert.True(t, want.Equals(v.Opposite()))
}

func TestVectorScale(t *testing.T) {
	var (
		scale  = 3.0
		vector = MakeVector(1, 2, 3)
		want   = MakeVector(3, 6, 9)
	)

	assert.True(t, want.Equals(vector.Scaled(scale)))
}

func TestVectorParallelism(t *testing.T) {
	assert := assert.New(t)

	t.Run("parallel vectors", func(t *testing.T) {
		var (
			u = MakeVector(1, 2, 3)
			v = MakeVector(2, 4, 6)
		)

		assert.True(u.IsParallelTo(v))
	})

	t.Run("non parallel vectors", func(t *testing.T) {
		var (
			u = MakeVector(1, 2, 3)
			v = MakeVector(5, 1, -9)
		)

		assert.False(u.IsParallelTo(v))
	})
}

func TestVectorPerpendicularity(t *testing.T) {
	assert := assert.New(t)

	t.Run("perpendicular vectors", func(t *testing.T) {
		var (
			u = MakeVector(1, 0, 0)
			v = MakeVector(0, 1, 0)
		)

		assert.True(u.IsPerpendicularTo(v))
	})

	t.Run("non perpendicular vectors", func(t *testing.T) {
		var (
			u = MakeVector(1, 0, 0)
			v = MakeVector(2, 1, -5)
		)

		assert.False(u.IsPerpendicularTo(v))
	})
}

func TestVectorOperations(t *testing.T) {
	assert := assert.New(t)

	var (
		u = MakeVector(1, 2, 3)
		v = MakeVector(4, 6, 8)
	)

	t.Run("Vector addition", func(t *testing.T) {
		want := MakeVector(5, 8, 11)
		assert.True(want.Equals(u.Plus(v)))
	})

	t.Run("Vector subtraction", func(t *testing.T) {
		want := MakeVector(-3, -4, -5)
		assert.True(want.Equals(u.Minus(v)))
	})

	t.Run("Vector dot product", func(t *testing.T) {
		want := 40.0
		assert.True(nums.FloatsEqual(want, u.DotTimes(v)))
	})

	t.Run("Vector cross product", func(t *testing.T) {
		want := MakeVector(-2, 4, -2)

		assert.True(want.Equals(u.CrossTimes(v)))
		assert.True(want.Opposite().Equals(v.CrossTimes(u)))
	})
}

func TestVectorSquared(t *testing.T) {
	var (
		v    = MakeVector(3, 5, 7)
		want = Make3x3Matrix(9, 15, 21, 15, 25, 35, 21, 35, 49)
	)

	assert.True(t, want.Equals(v.Squared()))
}
