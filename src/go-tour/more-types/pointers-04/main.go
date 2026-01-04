package main

func main() {
	x := 10

	y := &x

	*y = 20

	println(x, *y, &x, y)
}
