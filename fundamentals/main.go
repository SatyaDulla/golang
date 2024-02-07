package main

import "fmt"

func main() {
	fundVariables()
	fundTypes()
	fundConstants()
	fundFunction("Satya")
	fundFunctionMultipleReturns(10, 20)
	fundFunctionMultipleNamedReturns(10, 20)
	fundFunctionMultipleNamedReturnsBlank(15, 25)

	//length, width := 10.0, 20.0
	area, _ := fundFunctionMultipleNamedReturnsBlank(12, 12) // Using _ to ignore the perimeter
	fmt.Println("Area with Perimeter using blank identifier:", area)

	// Variadic function call
	sum(1, 2, 3, 4, 5)

    // Higher-order function call
    // Here we define an anonymous function inline and pass it to applyFunc
    applyFunc(func(x int) int {
        return x * x
    }, 5)
	fmt.Println("\n\n----------------------------------------------")
}
