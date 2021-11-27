package g3d

import "testing"

func TestPlaneCreationFromParams(t *testing.T) {
	t.Run("from the a, b, c and d parameters", func(t *testing.T) {
		var (
			a        = 1.0
			b        = 2.0
			c        = 3.0
			d        = 4.0
			plane, _ = MakePlane(a, b, c, d)
			want     = MakeVector(a, b, c)
		)

		t.Run("computes the normal vector", func(t *testing.T) {
			if got := plane.NormalVector(); !want.Equals(got) {
				t.Errorf("Want normal vector %v, but got %v", want, got)
			}
		})

		t.Run("can't be created from a zero normal vector", func(t *testing.T) {
			plane, err := MakePlane(0, 0, 0, 1)

			if plane != nil {
				t.Errorf("Want nil plane, but got %v", plane)
			}
			if err != ErrZeroVector {
				t.Error("Expected error")
			}
		})
	})
}
