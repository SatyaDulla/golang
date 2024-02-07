package main

import "fmt"

// https://go.dev/play/p/_1wEKsrxfmj

// Simple function demonstration with a single string parameter and return value.
// It prints and returns the passed string.
func fundFunction(name string) string {
    fmt.Println("\n\n################################################")
    fmt.Println("####              Functions                ####")
    fmt.Println("################################################")
    fmt.Println("Hello,", name)
    return name
}

// Demonstrates a function with multiple return values.
// Calculates and returns the area and the perimeter of a rectangle.
func fundFunctionMultipleReturns(length, width float64) (float64, float64) {
    var area = length * width
    var perimeter = (length + width) * 2
    fmt.Printf("Area: %.2f, Perimeter: %.2f", area, perimeter)
    return area, perimeter
}

// Demonstrates named return values for a function.
// Allows the function to return without specifying return values explicitly.
func fundFunctionMultipleNamedReturns(length, width float64) (area, perimeter float64) {
    area = length * width
    perimeter = (length + width) * 2
    fmt.Printf("\nNamed Returns -> Area: %.2f, Perimeter: %.2f", area, perimeter)
    return // No explicit return value needed because of named return values
}

func fundFunctionMultipleNamedReturnsBlank(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	fmt.Printf("\nArea %.2f Perimeter %.2f\n", area, perimeter)
	fmt.Println("\n\n------")
	return
}

// Demonstrates a variadic function that can accept a variable number of int arguments.
// Calculates and returns the sum of all arguments.
func sum(numbers ...int) int {
    total := 0
    for _, number := range numbers {
        total += number
    }
    fmt.Printf("\nSum of numbers: %d", total)
    return total
}

// Demonstrates a higher-order function by taking a function and an integer as arguments.
// Applies the passed function to the integer and returns the result.
func applyFunc(f func(int) int, val int) int {
    result := f(val)
    fmt.Printf("\nResult of applying function: %d", result)
    return result
}
