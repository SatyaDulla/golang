package main

import "fmt"

func fundConstants() {
	fmt.Println("\n\n################################################")
	fmt.Println("####      Constants Overview      ####")
	fmt.Println("################################################")

	// Basic constant declaration
	const a = 50 // Declaring a constant
	fmt.Println("Printing a const value:", a)

	// Grouped constant declaration
	const (
		name = "Satya"
		age  = 30
	)
	fmt.Println("Printing const name and age:", name, age)

	fmt.Println("\n\n---------")

	// Demonstrating untyped and typed constants
	const city = "New York" // Untyped constant
	var cityCopy = city     // The type of cityCopy is inferred from city (string)
	fmt.Printf("Type %T Value %v\n", city, cityCopy)

	// Constants with compile-time arithmetic
	const x = 3
	const y = x * 2 // Compile-time arithmetic
	fmt.Println("Value of y:", y)

	fmt.Println("\n\n---------")

	// Enumeration using constants and iota
	const (
		jan = iota + 1 // 1
		feb            // 2
		mar            // 3
	)
	fmt.Println("January:", jan, "February:", feb, "March:", mar)
	fmt.Println("\n\n----------------------------------------------")
}
