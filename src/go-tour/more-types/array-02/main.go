package main

import "fmt"

func main() {
	var arr []string

	arr = append(arr, "Hello")
	arr = append(arr, "World")

	fmt.Println(arr)

	primes := [6]int{2, 3, 5, 7, 11, 13}

	s := primes[1:4]

	fmt.Println(s)
}
