package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

const (
	limit = 0.001
)

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %.0f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	z := 1.0
	z2 := float64(0)
	for (z-z2) < -limit || limit < (z-z2) {
		z2 = z
		z = z - (z*z-x)/(2*z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

