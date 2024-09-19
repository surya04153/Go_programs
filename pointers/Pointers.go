package main

import (
	"fmt"
)

// Function to change the ice cream name using a pointer
func changeIceCreamName(name *string, newName string) {
	*name = newName
}

func main() {
	// Declare a variable
	x := 10
	// Create a pointer to the variable 'x'
	ptr := &x

	// Print the value of x, the pointer, and the value the pointer points to
	fmt.Printf("x: %d\n", x)       // x: 10
	fmt.Printf("ptr: %v\n", ptr)   // ptr: (memory address of x)
	fmt.Printf("*ptr: %d\n", *ptr) // *ptr: 10 (dereferencing the pointer)

	// Modify the value of x through the pointer
	*ptr = 20
	fmt.Printf("x after modification: %d\n", x) // x: 20

	// Original ice cream name
	originalName := "Vanilla Delight"
	fmt.Println("Original Ice Cream Name:", originalName)

	// Change the ice cream name using a pointer
	changeIceCreamName(&originalName, "Choco Bliss")
	fmt.Println("New Ice Cream Name:", originalName)

}
