package main

import (
	"fmt"
)

func main() {
	// Integer formatting
	i := 42
	fmt.Printf("Decimal: %d\n", i)     // Decimal format
	fmt.Printf("Binary: %b\n", i)      // Binary format
	fmt.Printf("Octal: %o\n", i)       // Octal format
	fmt.Printf("Hexadecimal: %x\n", i) // Hexadecimal format

	// String formatting
	s := "Hello, Go!"
	fmt.Printf("String: %s\n", s)         // Regular string
	fmt.Printf("Quoted string: %q\n", s)  // Quoted string
	fmt.Printf("Default format: %v\n", s) // Default format (same as %s)

	// Float formatting
	f := 3.14159
	fmt.Printf("Float: %f\n", f)             // Float in decimal format
	fmt.Printf("Scientific: %e\n", f)        // Float in scientific notation
	fmt.Printf("Formatted float: %.2f\n", f) // Float with 2 decimal places
}
