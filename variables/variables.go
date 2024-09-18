package main

import (
	"fmt"
)

func main() {
	// Declaring variables with explicit types
	var x int = 10
	var y float64 = 3.14
	var name string = "Alice"
	var isActive bool = true

	// Declaring variables without specifying the type (type inferred)
	var age = 25
	city := "New York" // shorthand for declaration and initialization

	// Multiple variable declarations
	var a, b int = 5, 7
	var c, d = "hello", 4.5

	// Printing variables
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("name:", name)
	fmt.Println("isActive:", isActive)
	fmt.Println("age:", age)
	fmt.Println("city:", city)
	fmt.Println("a:", a, "b:", b)
	fmt.Println("c:", c, "d:", d)

	// Modifying a variable
	x = 20
	fmt.Println("Updated x:", x)

	// Constant declaration
	const pi = 3.14159
	fmt.Println("pi:", pi)

	// Uninitialized variables (they will take default values)
	var i int
	var f float64
	var s string
	var bo bool

	// Printing the default values of the variables
	fmt.Println("Default int value:", i)     // 0
	fmt.Println("Default float64 value:", f) // 0
	fmt.Println("Default string value:", s)  // ""
	fmt.Println("Default bool value:", bo)   // false

	// Variable declaration block
	var (
		ig     int     = 42
		fl     float64 = 3.14
		str    string  = "Go Language"
		active bool    = true
	)

	// Printing the values
	fmt.Println("i:", ig)
	fmt.Println("f:", fl)
	fmt.Println("s:", str)
	fmt.Println("active:", active)

	// Modifying variables
	i = 100
	s = "Golang"
	fmt.Println("Updated i:", i)
	fmt.Println("Updated s:", s)
}
