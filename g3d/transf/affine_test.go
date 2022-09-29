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

	t.Run("Rotation in the YZ plane", func(t *testing.T) {
		var (
			angle    = math.Pi / 2
			origin   = g3d.MakePoint(0, 4, 5)
			rotation = MakeRotationAround(angle, g3d.IVersor, origin)
			original = g3d.MakeVector(0, 7, 5)
			want     = g3d.MakeVector(0, 4, 8)
		)

		got := rotation.Apply(original)
		fmt.Println(got)

		assert.True(got.Equals(want))
	})
}
