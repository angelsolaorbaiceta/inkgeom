package inkgeom

import "math"

const defaultEpsilon = 1e-10

/*
FloatsEqualEps compares two float64 values and returns true if the difference between
the two is smaller than a given epsilon.
*/
func FloatsEqualEps(a, b, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}

/*
FloatsEqual compares two float64 values and returns true if the difference between the
two is smaller than a default epsilon value (1E-10).
*/
func FloatsEqual(a, b float64) bool {
	return FloatsEqualEps(a, b, defaultEpsilon)
}

/*
LinInterpol computes the linear interpolation for a given position given two points on
the desired line: (startPos, startVal) and (endPos, endVal).
*/
func LinInterpol(startPos, startVal, endPos, endVal, posToInterpolate float64) float64 {
	return startVal + (posToInterpolate-startPos)*(endVal-startVal)/(endPos-startPos)
}
