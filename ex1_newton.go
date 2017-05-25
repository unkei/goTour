package main

import (
	"fmt"
	"math"
)

const (
	limit = 0.001
)

func Sqrt(x float64) float64 {
	//	z := 1.0
	//	for n := 0; n < 10; n++ {
	//		z = z - (z*z-x)/(2*z)
	//	}
	//	return z

	z := 1.0
	z2 := float64(0)
	for (z-z2) < -limit || limit < (z-z2) {
		z2 = z
		z = z - (z*z-x)/(2*z)
		fmt.Println(z, z2)
	}
	return z
}

func main() {
	s := Sqrt(2)
	ms := math.Sqrt(2)
	fmt.Println(s)
	fmt.Println("math.Sqrt(", ms, ") - ", s, " = ", ms - s)
}

