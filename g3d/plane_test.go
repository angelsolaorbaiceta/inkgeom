package g3d

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlaneCreationFromParams(t *testing.T) {
	assert := assert.New(t)

	t.Run("from the a, b, c and d parameters", func(t *testing.T) {
		var (
			a                = 1.0
			b                = 2.0
			c                = 3.0
			d                = 4.0
			plane, err       = MakePlane(a, b, c, d)
			wantNormalVec    = MakeVector(a, b, c)
			wantNormalVer, _ = wantNormalVec.ToVersor()
		)

		t.Run("plane created without error", func(t *testing.T) {
			assert.Nil(err)
		})

		t.Run("computes the normal vector", func(t *testing.T) {
			assert.True(wantNormalVec.Equals(plane.NormalVector()))
		})

		t.Run("computes the normal versor", func(t *testing.T) {
			assert.True(wantNormalVer.Equals(plane.NormalVersor()))
		})

		t.Run("can't be created from a zero normal vector", func(t *testing.T) {
			plane, err := MakePlane(0, 0, 0, 1)

			assert.Nil(plane)
			assert.Equal(ErrZeroVector, err)
		})
	})
}

func TestPlaneCreationFromPointAndNormal(t *testing.T) {
	assert := assert.New(t)

	var (
		p                = MakePoint(10, 20, 30)
		normalVec        = MakeVector(1, 1, 1)
		wantNormalVer, _ = normalVec.ToVersor()
		plane, err       = MakePlaneFromPointAndNormal(p, normalVec)
	)

	t.Run("plane created without error", func(t *testing.T) {
		assert.Nil(err)
	})

	t.Run("uses the normal vector", func(t *testing.T) {
		assert.True(normalVec.Equals(plane.NormalVector()))
	})

	t.Run("computes the normal versor", func(t *testing.T) {
		assert.True(wantNormalVer.Equals(plane.NormalVersor()))
	})

	t.Run("contains the point", func(t *testing.T) {
		assert.True(plane.ContainsPoint(p))
	})

	t.Run("can't create with a zero length normal vector", func(t *testing.T) {
		plane, err := MakePlaneFromPointAndNormal(p, Zero)

		assert.Nil(plane)
		assert.Equal(ErrZeroVector, err)
	})
}
