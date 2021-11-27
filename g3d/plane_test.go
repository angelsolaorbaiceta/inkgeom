package g3d

import "testing"

func TestPlaneCreationFromParams(t *testing.T) {
	t.Run("from the a, b, c and d parameters", func(t *testing.T) {
		var (
			a          = 1.0
			b          = 2.0
			c          = 3.0
			d          = 4.0
			plane, err = MakePlane(a, b, c, d)
			want       = MakeVector(a, b, c)
		)

		t.Run("plane created without error", func(t *testing.T) {
			if err != nil {
				t.Errorf("Want no error, but got %v", err)
			}
		})

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

func TestPlaneCreationFromPointAndNormal(t *testing.T) {
	var (
		p          = MakePoint(10, 20, 30)
		n          = MakeVector(1, 1, 1)
		plane, err = MakePlaneFromPointAndNormal(p, n)
	)

	t.Run("plane created without error", func(t *testing.T) {
		if err != nil {
			t.Errorf("Want no error, but got %v", err)
		}
	})

	t.Run("uses the normal vector", func(t *testing.T) {
		if got := plane.NormalVector(); !n.Equals(got) {
			t.Errorf("Want normal vector %v, but got %v", n, got)
		}
	})

	t.Run("contains the point", func(t *testing.T) {
		if !plane.ContainsPoint(p) {
			t.Error("Expected plane to contain point")
		}
	})

	t.Run("can't create with a zero length normal vector", func(t *testing.T) {
		plane, err := MakePlaneFromPointAndNormal(p, Zero)

		if plane != nil {
			t.Errorf("Want nil plane, got %v", plane)
		}
		if err != ErrZeroVector {
			t.Error("Expected error")
		}
	})
}
