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
	fmt.Println("\n\n----------------------------------------------")
}
