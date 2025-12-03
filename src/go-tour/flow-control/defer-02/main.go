package main

import "fmt"

func testDeferAndPanic(shouldPanic bool) {
	defer fmt.Println("defer in testDeferAndPanic")
	fmt.Println("before panic")

	if shouldPanic {
		panic("panic occurred")
	}

	fmt.Println("after panic") // This line will not be executed
}

func main() {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")

	testDeferAndPanic(false)

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
