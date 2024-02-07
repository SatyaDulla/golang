package main

import (
	"fmt"
	"fundamentals/simpleinterest"
	"log"
)

/*
The initialization sequence proceeds as follows:

1. Imported packages are initialized first.
2. The `simpleinterest` package is initialized at the outset, including any init functions it contains.
3. Package-level variables (p, r, t) are initialized subsequently.
4. The init function within the main package is executed following the initialization of package-level variables.
5. Lastly, the main function is executed.
*/
var p, r, t = 100.0, 10.0, 2.0

// init function is used to perform initial checks and setup.
func init() {
	fmt.Println("Main package initialized")
	if p < 0 {
		log.Fatal("Principal is less than zero")
	}
	if r < 0 {
		log.Fatal("Rate of interest is less than zero")
	}
	if t < 0 {
		log.Fatal("Duration is less than zero")
	}
}

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

	fmt.Println("\nSimple Interest Calculation")
	si := simpleinterest.Calculate(p, r, t)
	fmt.Printf("Simple interest is: %.2f\n", si)

	fmt.Println("\n----------------------------------------------")

	num := 4
	calculateEven(num)
	fmt.Println(demonstrateIfElseIfElse(num))
	fmt.Println(demonstrateShortStatement(5))
}
