package main

import "fmt"

func main() {
	Level1()
	fmt.Println("Program continues after panic recovery.")
}

func Level1() {
	recover()

	defer func() {
		fmt.Println("Deferred in Level 1")

		if r := recover(); r != nil {
			fmt.Println("Panic caught at Level 3!")
		}
	}()

	Level2()
	fmt.Println("This will never print")

}

func Level2() {
	// No defer here!

	Level2Branch()
	Level3()

	fmt.Println("This will never print")

	// panic("BOOM from Level 2")
}

func Level2Branch() {
	defer fmt.Println("Deferred in Level 2 Branch")

	// panic("BOOM from Level 2 Branch")
}

func Level3() {
	defer fmt.Println("Deferred in Level 3")

	panic("BOOM")
}
