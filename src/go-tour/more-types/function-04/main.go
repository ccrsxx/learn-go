package main

import "fmt"

// 1. Define the shape (Struct of Functions)
// This acts like an Interface or a Class definition in JS
type MathKit struct {
	Add func(float64, float64) float64
	Sub func(float64, float64) float64
	Mul func(float64, float64) float64
}

// 2. The Factory Function
// Returns the struct filled with specific logic (closures)
func NewMathFactory(startValue float64) MathKit {
	// You can even have local state here that the functions share!
	offset := startValue

	return MathKit{
		Add: func(x, y float64) float64 {
			return x + y + offset
		},
		Sub: func(x, y float64) float64 {
			return x - y - offset
		},
		Mul: func(x, y float64) float64 {
			return x * y
		},
	}
}

func main() {
	// Create the kit (with an offset of 0 for normal math)
	math := NewMathFactory(0)

	fmt.Println("Add:", math.Add(10, 5)) // 15
	fmt.Println("Sub:", math.Sub(10, 5)) // 5
	fmt.Println("Mul:", math.Mul(10, 5)) // 50

	// Create a "biased" kit (adds 100 to everything)
	biasedMath := NewMathFactory(100)
	fmt.Println("Biased Add:", biasedMath.Add(10, 5)) // 115 (10 + 5 + 100)
}
