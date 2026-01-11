package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		// r := v + 10
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T %v!\n", v, v)
	}
}

func test() {
	// crashes if emilia is not a string
	// var emilia any = 1
	var emilia any = "Emilia"

	emt, ok := emilia.(string)

	// rem := emt + " is cute"

	if ok {
		fmt.Println(emt)
	} else {
		rem := emt
		fmt.Println("emilia is not a string", rem)
	}
}

func main() {
	test()

	do(21)
	do("hello")
	do(true)
}
