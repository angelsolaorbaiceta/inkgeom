/*
Copyright 2020 Angel Sola Orbaiceta

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package g2d

import (
	"testing"

	"github.com/angelsolaorbaiceta/inkgeom"
)

func TestHasProjections(t *testing.T) {
	x, y := 2.0, 3.0
	proj := MakePoint(x, y)

	if proj.X != x {
		t.Error("Projectable has wrong X projection")
	}
	if proj.Y != y {
		t.Error("Projectable has wrong Y projection")
	}
}

func TestDistance(t *testing.T) {
	p, q := MakePoint(1, 2), MakePoint(4, 6)
	dist := p.DistanceTo(q)

	if !inkgeom.FloatsEqual(dist, 5.0) {
		t.Error("Wrong distance between points")
	}
}

func TestVectorNorm(t *testing.T) {
	v := MakeVector(3, 4)
	if !inkgeom.FloatsEqual(v.Norm(), 5.0) {
		t.Error("Wrong vector norm")
	}
}

func TestIsVersor(t *testing.T) {
	v := MakeVector(1, 0)
	if !v.IsVersor() {
		t.Error("Vector expected to be versor")
	}
}

func TestIsNotVersor(t *testing.T) {
	v := MakeVector(1, 2)
	if v.IsVersor() {
		t.Error("Vector expected NOT to be versor")
	}
}

func TestMakeVersor(t *testing.T) {
	v := MakeVersor(1, 2)
	if !v.IsVersor() {
		t.Error("Vector expected to be versor")
	}
}

func TestPerpendicular(t *testing.T) {
	v := MakeVersor(1, 2)
	if !v.Perpendicular().Equals(MakeVersor(-2, 1)) {
		t.Error("Wrong perpendicular vector")
	}
}

/* Operations */
func TestSumVectors(t *testing.T) {
	v := MakeVector(1, 2).Plus(MakeVector(3, 4))
	if !v.Equals(MakeVector(4, 6)) {
		t.Error("Addition yielded wrong result")
	}
}

func TestSubtractVectors(t *testing.T) {
	v := MakeVector(1, 2).Minus(MakeVector(3, 5))
	if !v.Equals(MakeVector(-2, -3)) {
		t.Error("Subtraction yielded wrong result")
	}
}

func TestVectorsDotProduct(t *testing.T) {
	p := MakeVector(1, 2).DotTimes(MakeVector(3, 5))
	if !inkgeom.FloatsEqual(p, 13.0) {
		t.Error("Dot product yielded wrong result")
	}
}

func TestVectorsCrossProduct(t *testing.T) {
	p := MakeVector(1, 2).CrossTimes(MakeVector(3, 5))
	if !inkgeom.FloatsEqual(p, -1.0) {
		t.Error("Cross product yielded wrong result")
	}
}

/* Ordering */
func TestEqualPointsOrder(t *testing.T) {
	p := MakePoint(2, 5)
	if p.Compare(p) != 0 {
		t.Error("Expected points to have equal order")
	}
}

func TestPointWithSmallerXGoesFirts(t *testing.T) {
	p, q := MakePoint(1, 5), MakePoint(2, 3)
	if p.Compare(q) != -1 {
		t.Error("Expected point p to go begore q")
	}
	if q.Compare(p) != 1 {
		t.Error("Expected point p to go begore q")
	}
}

func TestPointWithSmallerYGoesFirts(t *testing.T) {
	p, q := MakePoint(2, 5), MakePoint(2, 6)
	if p.Compare(q) != -1 {
		t.Error("Expected point p to go begore q")
	}
	if q.Compare(p) != 1 {
		t.Error("Expected point p to go begore q")
	}
}
