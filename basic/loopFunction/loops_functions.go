package loopfunction

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0 // initial guess
	fmt.Printf("x=%v, z=%v\n", x, z)
	for i := 1; i < 11; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(i, "\t", z)
	}
	return z
}

const delta = 1e-10

func QuickSqrt(x float64) float64 {
	z := 1.0
	t := 0.0
	count := 1
	for math.Abs(t-z) > delta {
		t, z = z, z-(z*z-x)/(2*z)
		count++
		fmt.Println(count, "\t", z)
	}
	return z
}

func LoopFunction() {
	fmt.Println("Result: ", Sqrt(2))
	fmt.Println("Result: ", QuickSqrt(2))
	fmt.Println("Accurate:", math.Sqrt(2))
}
