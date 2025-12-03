package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	v := 100
	z := "rem"

	fmt.Println("out", v, z)

	if v, z, zz := math.Pow(x, n), "Emilia", "Rin"; v < lim {
		fmt.Println("inside", v, z, zz)

		return v
	}

	fmt.Println("after", v, z, x)

	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
