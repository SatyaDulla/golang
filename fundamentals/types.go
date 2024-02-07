package main

import "fmt"

// https://go.dev/play/p/W00T5PB2Wsu

func fundTypes() {
	fmt.Println("\n\n################################################")
	fmt.Println("####      types		####")
	fmt.Println("################################################")
	// int type
	// Declaring an integer variable 'x' with explicit type and 'y' with type inference.
	var x int = 1
	y := 2 // type inferred as int
	fmt.Println("value of x is", x, "and y is", y)
	fmt.Printf("type of x is %T and type of y is %T\n", x, y) // Printing types of x and y

	// Type Conversion: Demonstrating how to handle operations between different types.
	i := 55   // int
	j := 67.8 // float64
	// sum := i + j // This operation is not allowed directly due to type mismatch.
	// Converting float64 to int for addition.
	sum := i + int(j) // Explicit type conversion required
	fmt.Println("Corrected sum of i and j with type conversion is", sum)

	// string type
	// Concatenating strings to form a full name.
	first := "Satya"
	last := "D"
	name := first + " " + last // Concatenating strings with a space
	fmt.Println("My name is", name)

	// bool type
	// Demonstrating boolean operations.
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	c := a && b // Logical AND operation
	fmt.Println("c:", c)
	d := a || b // Logical OR operation
	fmt.Println("d:", d)

	// Type Conversion
	// Demonstrating type conversion with integer and float64 types.
	i = 10           // int
	j = 20.5         // float64
	sum = i + int(j) // Explicit type conversion from float64 to int
	fmt.Println("sum of i & j with type conversion is", sum)
	fmt.Println("\n\n----------------------------------------------")
}
