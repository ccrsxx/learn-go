package main

import "fmt"

type Config struct {
	z int
	// is it necessary to have a pointer to a reference type that already defaults to nil?
	// testMap *map[string]int
	testMap map[string]int
}

func main() {
	var x *Config

	x = &Config{
		z: 10,
	}

	// x.testMap = nil
	// x.testMap = map[string]int{}

	if x.testMap == nil {
		fmt.Println("testMap is nil")
	}

	fmt.Printf("%#v\n", x)
}
