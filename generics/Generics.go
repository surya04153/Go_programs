package main

import (
	"fmt"
)

// Generic function that works with numeric types
func Add[T int | float64](a, b T) T {
	return a + b
}

func main() {
	// Call the generic function with different numeric types
	fmt.Println(Add(3, 4))     // int: 7
	fmt.Println(Add(3.5, 2.5)) // float64: 6
	// fmt.Println(Add(3, "4"))  // This would cause a compilation error because string is not allowed
}
