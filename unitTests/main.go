package main

import "fmt"

// Add function to add two numbers
func Add(a, b int) int {
	return a + b
}

func main() {
	result := Add(2, 3)
	fmt.Println("Result:", result)
}
