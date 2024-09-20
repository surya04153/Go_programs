package main

import (
	"fmt"
	"sync"
)

func printNumbers(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal the WaitGroup that this goroutine is done
	for i := 1; i <= 5; i++ {
		fmt.Printf("Goroutine %d: %d\n", id, i)
	}
}

func main() {
	var wg sync.WaitGroup

	// Launch multiple goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go printNumbers(i, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("All goroutines finished!")
}
