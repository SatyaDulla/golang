package main

import "fmt"

// https://go.dev/play/p/27TL4OcZ0hY
func fundVariables() {
	fmt.Println("\n\n################################################")
	fmt.Println("####      variables		####")
	fmt.Println("################################################")
	// Example of multiple variable declarations using var block
	var (
		name   = "Satya" // Type inferred as string
		age    = 30      // Type inferred as int
		height int       // Declared but not initialized; defaults to 0
	)

	fmt.Println("Hello World!")

	// Demonstrating variable assignment and re-assignment
	fmt.Println("My age is", age)
	age = 32 // Re-assignment
	fmt.Println("Now, my age is", age)

	// Short variable declaration for a new variable (type inference)
	wifeAge := 24 // Type inferred as int
	fmt.Println("Wife's age is", wifeAge)

	// Multiple variable declarations with type inference
	heights, width := 100.00, 200.10 // Both variables are inferred as float64
	fmt.Println("Dimensions are Height:", heights, "Width:", width)

	// Printing multiple variables in one statement
	fmt.Println("Details:", name, age, height) // Note: height will print 0 as it's uninitialized

	// Short variable declaration with initialization
	headcount := 10 // Type inferred as int
	fmt.Println("Count is", headcount)

	// Short declaration for multiple variables
	myHomeHeadCount, wifeHomeHeadCount := 3, 5 // Both inferred as int
	fmt.Println("Next counts are", myHomeHeadCount, "and", wifeHomeHeadCount)
	fmt.Println("\n\n----------------------------------------------")
}
