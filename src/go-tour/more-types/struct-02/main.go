package emt

import "fmt"

type Vertex struct {
	X      int
	Y      int
	inside int
}

func Random() {
	v := Vertex{
		X:      1,
		Y:      2,
		inside: 3,
	}

	fmt.Println(v)
	fmt.Printf("%+v\n", v)

	v.X = 4

	fmt.Println(v)
	fmt.Printf("%+v\n", v)
}

func LoveEmilia() Vertex {
	data := Vertex{X: 1, Y: 2, inside: 3}

	return data
}
