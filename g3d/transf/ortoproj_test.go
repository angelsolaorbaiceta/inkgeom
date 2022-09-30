package transf

import (
	"testing"

	"github.com/angelsolaorbaiceta/inkgeom/g3d"
	"github.com/stretchr/testify/assert"
)

func TestOrtographicProjection(t *testing.T) {
	assert := assert.New(t)

	t.Run("Projection in the XY plane", func(t *testing.T) {
		var (
			plane = g3d.XYPlane
			proj  = MakeOrtographic(plane)
			orig  = g3d.MakePoint(1, 2, 3)
			want  = g3d.MakePoint(1, 2, 0)
		)

		got := proj.Apply(orig)

		assert.True(got.Equals(want))
	})

	t.Run("Projection in the YZ plane", func(t *testing.T) {
		var (
			plane = g3d.YZPlane
			proj  = MakeOrtographic(plane)
			orig  = g3d.MakePoint(1, 2, 3)
			want  = g3d.MakePoint(0, 2, 3)
		)

		got := proj.Apply(orig)

		assert.True(got.Equals(want))
	})

	t.Run("Projection in the XZ plane", func(t *testing.T) {
		var (
			plane = g3d.XZPlane
			proj  = MakeOrtographic(plane)
			orig  = g3d.MakePoint(1, 2, 3)
			want  = g3d.MakePoint(1, 0, 3)
		)

		got := proj.Apply(orig)

		assert.True(got.Equals(want))
	})
}
