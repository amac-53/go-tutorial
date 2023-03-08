package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {	
	if x < 0 {
		err := ErrNegativeSqrt(x)
		return x, err
	}
	
	z := 1.0
	var prev float64
	for {
		prev = z
		z -= (z*z - x) / (2*z)
		if math.Abs(prev - z) < 1e-5 {
		  return z, nil
		}
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
