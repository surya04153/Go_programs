package main

import "fmt"

func main() {
	// Basic for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	j := 0
	// for loop behaving like a while loop
	for j < 5 {
		fmt.Println(j)
		j++
	}

	numbers := []int{10, 20, 30, 40, 50}

	// for loop with range
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	capitals := map[string]string{"France": "Paris", "Japan": "Tokyo", "India": "New Delhi"}

	// for loop with range on a map
	for country, capital := range capitals {
		fmt.Printf("The capital of %s is %s\n", country, capital)
	}

	word := "Hello"

	// for loop with range on a string
	for index, char := range word {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}
}
