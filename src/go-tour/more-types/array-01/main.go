package main

import "fmt"

func main() {
	var arr [2]string

	arr[0] = "Hello"
	arr[1] = "World"

	fmt.Println(arr[0], arr[1])
	fmt.Println(arr)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
