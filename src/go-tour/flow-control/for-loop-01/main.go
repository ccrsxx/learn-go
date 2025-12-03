package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		sum += i
	}
	fmt.Println(sum)
}
