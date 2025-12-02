package main

import (
	"fmt"
	"math"
)

func main() {
	const x, y = 3, 4

	f := math.Sqrt(float64(x*x + y*y))

	z := uint(f)

	zz := int(z)

	const emilia = "Emilia"

	rem := "Rem"

	i := 42           // int
	zzz := 3.142      // float64
	g := 0.867 + 0.5i // complex128

	fmt.Println(x, y, z, zz, emilia, rem, i, zzz, g)
}
