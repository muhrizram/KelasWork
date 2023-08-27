package main

import "fmt"

func printHello() {
	fmt.Printf("Hello guys!\n")
}

func main() {
	// Defer
	defer printHello()

	age := 23

	// If Condition
	if age < 18 {
		fmt.Printf("Not old enough!\n")
	} else if age > 60 {
		fmt.Printf("Too old!\n")
	} else {
		fmt.Printf("Welcome young man!\n")
	}

	// Switch Case
	grade := "D"

	switch grade {
	case "A":
		fmt.Printf("Good job!\n")
	case "B", "C":
		fmt.Printf("You can improve!\n")
	case "D":
		fmt.Printf("You're in trouble. Find your teacher!\n")
		fallthrough
	default:
		fmt.Printf("You've to try again\n")
	}

	// Looping for
	for i := 1; i <= 5; i++ {
		fmt.Printf("Looping %d\n", i)
	}

	// Looping (while in Go)
	j := 1
	for j < 10 {
		fmt.Printf("Looping %d\n", j)
		j += 2
	}

	// Infinite loop with condition
	for {
		j += 1
		if j == 13 {
			continue
		}
		fmt.Printf("Looping %d\n", j)
		if j == 15 {
			break
		}
	}

	// For loop Array (Range Iteratur)
	numbers := []int{1, 2, 3, 4}
	for d, z := range numbers {
		fmt.Printf("I got %d in index %d\n", z, d)
	}
}
