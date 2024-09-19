package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Printf("Hello, %s!\n", name)

	var firstName, lastName string
	fmt.Print("Enter your first and last name: ")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hello, %s %s!\n", firstName, lastName)

	var age int
	var fname string
	fmt.Print("Enter your name and age (format: name age): ")
	fmt.Scanf("%s %d", &fname, &age)
	fmt.Printf("Hello, %s! You are %d years old.\n", fname, age)

	var number int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)
	fmt.Printf("You entered: %d\n", number)
}
