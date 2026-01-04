package main

import "fmt"

// Define a custom int type
type UserStatus int

const (
	Pending UserStatus = 0
	Active  UserStatus = 1
	Banned  UserStatus = 2
)

// Attach logic to the integer
func (s UserStatus) String() string {
	switch s {
	case Pending:
		return "Pending ⏳"
	case Active:
		return "Active ✅"
	case Banned:
		return "Banned 🚫"
	default:
		return "Unknown ❓"
	}
}

func main() {
	status := Active

	// It acts like an int, but behaves like an object!
	fmt.Println(status.String()) // Output: Active ✅
}
