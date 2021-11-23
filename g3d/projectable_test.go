package g3d

import "testing"

func TestCreateProjectable(t *testing.T) {
	var (
		x = 10.0
		y = 20.0
		z = 30.0
		p = Projectable{x, y, z}
	)

	t.Run("has a X projection", func(t *testing.T) {
		if got := p.X(); got != x {
			t.Errorf("Want X = %f, got %f", x, got)
		}
	})

	t.Run("has a Y projection", func(t *testing.T) {
		if got := p.Y(); got != y {
			t.Errorf("Want Y = %f, got %f", y, got)
		}
	})

	t.Run("has a Z projection", func(t *testing.T) {
		if got := p.Z(); got != z {
			t.Errorf("Want Z = %f, got %f", z, got)
		}
	})
}
