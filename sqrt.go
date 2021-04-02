package main

import (
	"fmt"
)

const (
	epsilon = 0.00000001
)

func Abs(x float64) float64 {
	if x > 0 {
		return x
	}
	return -x
}

func Sqrt(x float64) float64 {

	if x < 0 {
		return Sqrt(-x) + "i"
	}

	z := x / 2

	t := (z*z - x) / (2 * z)

	for Abs(t) > epsilon {
		t = (z*z - x) / (2 * z)
		z -= t
	}
	return z
}

func main() {
	fmt.Println(Sqrt(-16))
}
