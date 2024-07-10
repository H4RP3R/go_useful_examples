// a * x^2 + b * x + c = 0; a != 0
// D = b^2 - 4 * a * c
// x1 = (-b + √D) / 2 * a
// x2 = (-b - √D) / 2 * a

package main

import (
	"fmt"
	"math"
)

var (
	ErrZeroA       = fmt.Errorf("coefficient a is zero")
	ErrNoRealRoots = fmt.Errorf("equation has no real roots")
)

// SolveQuadraticEquation finds real roots of equation defined with 3 real coefficients
// It returns 2 roots if no error encountered or default float64 values and error otherwise
func SolveQuadraticEquation(a, b, c float64) (x1, x2 float64, err error) {
	if a == 0 {
		err = ErrZeroA
		return
	}

	D := b*b - 4*a*c
	if D < 0 {
		err = ErrNoRealRoots
		return
	}

	if D == 0 { // the equation has two identical roots
		x1 = -b / (2 * a)
		x2 = x1
		return // err == nil by default
	}

	dRoot := math.Sqrt(D)
	x1 = (-b + dRoot) / (2 * a)
	x2 = (-b - dRoot) / (2 * a)
	return
}
