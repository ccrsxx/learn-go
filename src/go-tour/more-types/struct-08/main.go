package main

import "fmt"

type Config struct {
	Host string
	Port int
}

func createConfig() *Config {
	return &Config{
		Host: "localhost",
		Port: 8080,
	}
}

func main() {
	config := createConfig()

	fmt.Printf("Server running at %s:%d\n", config.Host, config.Port)
}
