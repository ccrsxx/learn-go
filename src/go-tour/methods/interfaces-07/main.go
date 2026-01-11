package main

import "fmt"

func emilia() (int, int) {
	return 42, 7
}

func main() {
	x, y := emilia()

	fmt.Println(x, y)

	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	fmt.Println()
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
