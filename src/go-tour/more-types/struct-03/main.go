package main

import (
	"fmt"

	emt "github.com/ccrsxx/learn-go/src/go-tour/more-types/struct-02"
)

func main() {
	emilia := emt.LoveEmilia()

	data := emt.Vertex{X: 10, Y: 20}

	dataX := data.X
	dataY := data.Y

	data.Y = 40

	fmt.Println(emilia, dataX, dataY)
	fmt.Printf("%+v", data)
}
