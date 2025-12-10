package main

import "fmt"

type User struct {
	Name string
	Age  *int // Pointer! Can be nil (missing) or point to an int.
}

func main() {
	// Case 1: Age is missing (nil)
	u1 := User{Name: "Emilia", Age: nil}

	fmt.Printf("%+v\n", u1)

	// Case 2: Age is present
	// myAge := 100

	// this shit yells
	// u := User{Name: "Rem", Age: myAge}
	u := User{Name: "Rem", Age: nil}

	test := *u.Age + 10

	fmt.Println(test)

	fmt.Printf("%+v\n", *u.Age)
}
