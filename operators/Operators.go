package main

import (
	"fmt"
)

func main() {
	// 1. Arithmetic Operators
	a, b := 10, 3
	fmt.Println("Arithmetic Operators:")
	fmt.Println("a + b =", a+b) // Addition
	fmt.Println("a - b =", a-b) // Subtraction
	fmt.Println("a * b =", a*b) // Multiplication
	fmt.Println("a / b =", a/b) // Division (integer division)
	fmt.Println("a % b =", a%b) // Modulus (remainder)

	// 2. Comparison Operators
	fmt.Println("\nComparison Operators:")
	fmt.Println("a == b:", a == b) // Equal to
	fmt.Println("a != b:", a != b) // Not equal to
	fmt.Println("a > b:", a > b)   // Greater than
	fmt.Println("a < b:", a < b)   // Less than
	fmt.Println("a >= b:", a >= b) // Greater than or equal to
	fmt.Println("a <= b:", a <= b) // Less than or equal to

	// 3. Logical Operators
	x, y := true, false
	fmt.Println("\nLogical Operators:")
	fmt.Println("x && y:", x && y) // Logical AND
	fmt.Println("x || y:", x || y) // Logical OR
	fmt.Println("!x:", !x)         // Logical NOT

	// 4. Bitwise Operators
	fmt.Println("\nBitwise Operators:")
	fmt.Println("a & b:", a&b)   // AND
	fmt.Println("a | b:", a|b)   // OR
	fmt.Println("a ^ b:", a^b)   // XOR
	fmt.Println("a << 1:", a<<1) // Left shift
	fmt.Println("a >> 1:", a>>1) // Right shift

	// 5. Assignment Operators
	fmt.Println("\nAssignment Operators:")
	c := 5
	fmt.Println("c =", c)
	c += 3 // Equivalent to c = c + 3
	fmt.Println("c += 3:", c)
	c *= 2 // Equivalent to c = c * 2
	fmt.Println("c *= 2:", c)

	// 6. Unary Operators
	fmt.Println("\nUnary Operators:")
	d := 10
	fmt.Println("d:", d)
	d++ // Increment (equivalent to d = d + 1)
	fmt.Println("d++:", d)
	d-- // Decrement (equivalent to d = d - 1)
	fmt.Println("d--:", d)
}
