package main

import "fmt"

func hello() string {
	return "Hello, CI/CD!"
}
func main() {
	message := hello()
	fmt.Print(message)
}
