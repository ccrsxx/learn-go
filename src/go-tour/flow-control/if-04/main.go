package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0

	// Loop almost forever, we will break manually
	for i := 0; i < 1000; i++ {
		prevZ := z // 1. Store previous value

		// 2. Calculate new value
		z -= (z*z - x) / (2 * z)

		fmt.Printf("Iter %d: %.10f\n", i, z)

		// 3. Compare: Is the change smaller than 0.0001?
		// math.Abs handles negative differences
		if changes := math.Abs(z - prevZ); math.Abs(changes) < 0.0001 {
			fmt.Println("Converged!")
			break
		}
	}

	return z
}

func main() {
	fmt.Println("Final:", Sqrt(2))
}
