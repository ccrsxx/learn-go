package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	slice := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		slice[y] = make([]uint8, dx)

		for x := 0; x < dx; x++ {
			slice[y][x] = uint8((x + y) / 2)
		}
	}

	fmt.Println(slice)

	return slice
}

func main() {
	pic.Show(Pic)
}
