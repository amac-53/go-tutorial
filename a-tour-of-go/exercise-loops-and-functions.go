package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	var prev float64
	for {
		prev = z
		z -= (z*z - x) / (2*z)
		fmt.Println("z:", z)
		if math.Abs(prev - z) < 1e-5 {
		  return z
		}
	}
}

func main() {
	fmt.Println(Sqrt(5))
}
