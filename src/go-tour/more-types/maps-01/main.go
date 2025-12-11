package main

import "fmt"

type Vertex struct {
	str    string         // Value Type -> Defaults to ""
	Lat    float64        // Value Type -> Defaults to 0.0
	Long   float64        // Value Type -> Defaults to 0.0
	m      map[string]int // Reference Tyspe -> Defaults to nil
	arr    [5]int         // Value Type -> Defaults to [0 0 0 0 0]
	dArray []int          // Reference Type -> Defaults to nil
}

func main() {
	m := make(map[string]Vertex)

	// 1. Assign an empty Vertex (Go fills it with Zero Values)
	m["Bell Labs"] = Vertex{}

	v := m["Bell Labs"]

	// 2. The "Nice" Print (Hides the truth)
	fmt.Println("--- Nice Print (%v) ---")
	fmt.Printf("%+v\n\n", v)

	// 3. The "Real" Print (Shows nil vs empty)
	fmt.Println("--- Real Truth (%#v) ---")
	fmt.Printf("%#v\n", v)
}
