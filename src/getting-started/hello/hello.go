package main

import (
	"fmt"

	"github.com/ccrsxx/learn-go/src/getting-started/greetings"
)

func getHelloMessage() string {
	return greetings.Hello("Emilia")
}

func sayHello() {
	message := getHelloMessage()

	fmt.Println(message)
}

func main() {
	sayHello()
}
