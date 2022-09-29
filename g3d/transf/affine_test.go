package transf

import (
	"fmt"
	"math"
	"testing"

	"github.com/angelsolaorbaiceta/inkgeom/g3d"
	"github.com/stretchr/testify/assert"
)

func TestAffineTransformation(t *testing.T) {
	assert := assert.New(t)

	t.Run("Translation transformation", func(t *testing.T) {
		var (
			x, y, z     = 2.0, 3.0, 4.0
			translation = MakeTranslation(x, y, z)
			original    = g3d.MakeVector(1, 2, 3)
			want        = g3d.MakeVector(3, 5, 7)
		)

		got := translation.Apply(original)

		assert.True(got.Equals(want))
	})

	t.Run("Rotation in the YZ plane", func(t *testing.T) {
		var (
			angle    = math.Pi / 2
			origin   = g3d.MakePoint(0, 4, 5)
			rotation = MakeRotationAround(angle, g3d.IVersor, origin)
			original = g3d.MakeVector(0, 7, 5)
			want     = g3d.MakeVector(0, 4, 8)
		)

		got := rotation.Apply(original)

		assert.True(got.Equals(want))
	})

	t.Run("Rotation in the XZ plane", func(t *testing.T) {
		var (
			angle    = math.Pi / 2
			origin   = g3d.MakePoint(4, 0, 5)
			rotation = MakeRotationAround(angle, g3d.JVersor, origin)
			original = g3d.MakeVector(7, 0, 5)
			want     = g3d.MakeVector(4, 0, 2)
		)

		got := rotation.Apply(original)
		fmt.Println(got)

		assert.True(got.Equals(want))
	})

	t.Run("Rotation in the XY plane", func(t *testing.T) {
		var (
			angle    = math.Pi / 2
			origin   = g3d.MakePoint(4, 5, 0)
			rotation = MakeRotationAround(angle, g3d.KVersor, origin)
			original = g3d.MakeVector(7, 5, 0)
			want     = g3d.MakeVector(4, 8, 0)
		)

		got := rotation.Apply(original)
		fmt.Println(got)

		assert.True(got.Equals(want))
	})
}
