package transf

import (
	"math"
	"testing"

	"github.com/angelsolaorbaiceta/inkgeom/g3d"
	"github.com/stretchr/testify/assert"
)

func TestLinearTransformation(t *testing.T) {
	assert := assert.New(t)

	t.Run("Scaling transformation", func(t *testing.T) {
		var (
			x, y, z  = 2.0, 3.0, 4.0
			scaling  = MakeScaling(x, y, z)
			original = g3d.MakeVector(1, 2, 3)
			want     = g3d.MakeVector(2, 6, 12)
		)

		got := scaling.Apply(original)

		assert.True(got.Equals(want))
	})

	t.Run("Uniform scaling transformation", func(t *testing.T) {
		var (
			factor   = 2.0
			scaling  = MakeUniformScaling(factor)
			original = g3d.MakeVector(1, 2, 3)
			want     = g3d.MakeVector(2, 4, 6)
		)

		got := scaling.Apply(original)

		assert.True(got.Equals(want))
	})

	t.Run("Rotation in the YZ plane", func(t *testing.T) {
		var (
			angle    = math.Pi / 2
			rotation = MakeRotation(angle, g3d.IVersor)
			original = g3d.MakeVector(0, 1, 2)
			want     = g3d.MakeVector(0, -2, 1)
		)

		got := rotation.Apply(original)

		assert.True(got.Equals(want))
	})

	t.Run("Rotation in the XZ plane", func(t *testing.T) {
		var (
			angle    = math.Pi / 2
			rotation = MakeRotation(angle, g3d.JVersor)
			original = g3d.MakeVector(1, 0, 2)
			want     = g3d.MakeVector(2, 0, -1)
		)

		got := rotation.Apply(original)

		assert.True(got.Equals(want))
	})

	t.Run("Rotation in the XY plane", func(t *testing.T) {
		var (
			angle    = math.Pi / 2
			rotation = MakeRotation(angle, g3d.KVersor)
			original = g3d.MakeVector(1, 2, 0)
			want     = g3d.MakeVector(-2, 1, 0)
		)

		got := rotation.Apply(original)

		assert.True(got.Equals(want))
	})
}
