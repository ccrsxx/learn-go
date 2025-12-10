package main

import "fmt"

type Vertex struct {
	X, Y int
	z    string
}

var (
	v1 = Vertex{1, 2, "ss"}  // has type Vertex
	v2 = Vertex{X: 1}        // Y:0 is implicit
	v3 = Vertex{}            // X:0 and Y:0
	p  = &Vertex{1, 2, "ss"} // has type *Vertex

)

func main() {
	fmt.Println(v1, p, *p, v2, v3)
}
