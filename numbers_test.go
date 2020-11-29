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

package inkgeom

import "testing"

func TestFloatsEqual(t *testing.T) {
	t.Run("floats equal with epsilon", func(t *testing.T) {
		if !FloatsEqualEps(1.0, 1.0001, 0.01) {
			t.Error("Floats expected to equal")
		}
	})

	t.Run("floats don't equal with epsilon", func(t *testing.T) {
		if FloatsEqualEps(1.0, 1.01, 0.001) {
			t.Error("Floats expected to not equal")
		}
	})
}

func TestLinearInterpolation(t *testing.T) {
	var (
		want = 10.0
		got  = LinInterpol(0.0, 0.0, 20.0, 20.0, 10.0)
	)

	if !FloatsEqual(want, got) {
		t.Errorf("Wanted %f, but got %f", want, got)
	}
}
