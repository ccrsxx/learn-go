package main

import methods "github.com/ccrsxx/learn-go/src/go-tour/methods/methods-01"

func main() {
	v := methods.Vertex{
		X: 3,
		Y: 4,
	}

	println(v.Abs(), v.Emilia())
}
