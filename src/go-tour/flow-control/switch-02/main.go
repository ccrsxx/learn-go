package main

import (
	"fmt"
)

func main() {
	const emilia = 1

	switch emilia {
	case 1:
		fmt.Println("Emilia watches 1 episode")
		fallthrough
	case 2:
		fmt.Println("Emilia watches 2 episodes")
	case 3, 4:
		fmt.Println("Emilia watches 3 or 4 episodes")
	default:
		fmt.Println("Emilia binge-watches the entire season!")
	}
}
