package main

import (
	"fmt"

	emt "github.com/ccrsxx/learn-go/src/go-tour/more-types/struct-02"
)

func main() {
	emilia := emt.LoveEmilia()

	fmt.Println("Before changeX:", emilia)

	changeX(emilia)

	fmt.Println("After changeX:", emilia)

	fmt.Println("Before changePointer:", emilia)

	changePointer(&emilia)

	fmt.Println("After changePointer:", emilia)
}

func changeX(v emt.Vertex) {
	v.X = 100

	fmt.Println("Inside changeX:", v)
}

func changePointer(v *emt.Vertex) {
	v.X = 100

	fmt.Println("Inside changePointer:", v)
}
