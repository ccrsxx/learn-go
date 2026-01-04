package main

import "fmt"

type Settings map[string]string

type Vertex struct {
	settings Settings
}

func main() {
	v := Vertex{
		settings: Settings{},
	}

	fmt.Println(v)
}
