package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		fmt.Printf("fibonacci: x=%d, y=%d\n", x, y)

		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}
