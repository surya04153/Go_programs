package main

import "fmt"

// Function definition
func greet() {
	fmt.Println("Hello, World!")
}

// Function with parameters
func greetUser(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// Function that returns a value
func add(a int, b int) int {
	return a + b
}

// Function returning multiple values
func swap(a, b int) (int, int) {
	return b, a
}

// Function with named return values
func divide(dividend, divisor int) (quotient int, remainder int) {
	quotient = dividend / divisor
	remainder = dividend % divisor
	return
}

// Variadic function
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Recursive function to calculate factorial
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	// Function call
	greet()

	// Passing an argument to the function
	greetUser("Surya")

	// Storing and printing the result of the function
	result := add(5, 3)
	fmt.Println("Sum:", result)

	x, y := 1, 2
	fmt.Println("Before swap:", x, y)

	// Calling the function and assigning the returned values
	x, y = swap(x, y)
	fmt.Println("After swap:", x, y)

	q, r := divide(10, 3)
	fmt.Printf("Quotient: %d, Remainder: %d\n", q, r)

	addresult := sum(1, 2, 3, 4, 5)
	fmt.Println("Sum:", addresult)

	factorial := factorial(5)
	fmt.Println("Factorial:", factorial)

}
