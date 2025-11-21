package export // <--- THIS LINE IS MANDATORY

import "fmt"

// Notice it is Capitalized so it is "Exported" (visible)
func BestGirl() {
	fmt.Println("Emilia is the best!")
}

func PrivateBestGirl() {
	privateBestGirl()
}

func privateBestGirl() {
	fmt.Println("This function is private to the package")
}
