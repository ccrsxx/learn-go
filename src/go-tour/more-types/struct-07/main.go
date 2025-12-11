package main

import "fmt"

type Vertex struct {
	X int
	Y int
	z *int
}

func main() {
	var a []int
	var s *string
	var v Vertex

	fmt.Println(a == nil)
	fmt.Printf("%#v %#v %#v\n", a, s, v)

	// s = new(string)
	// v = &Vertex{
	// 	X: 10,
	// 	Y: 20,
	// }

	fmt.Printf("%#v %#v %#v\n", a, s, v)
}
