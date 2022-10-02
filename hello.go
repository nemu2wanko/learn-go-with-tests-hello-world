package main

import "fmt"

const englishHelloPrefix = "Hello, "

// Hello returns a personalised greeting, defaulting to Hello, World if an empty name is passed.
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
