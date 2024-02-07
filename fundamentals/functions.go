package main

import "fmt"

// func fundFunction(parameterName type) returnType {
// function body
// }
func fundFunction(name string) string {
	fmt.Println("\n\n################################################")
	fmt.Println("####      functions		####")
	fmt.Println("################################################")
	fmt.Println(name)
	return name
}

func fundFunctionMultipleReturns(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	fmt.Printf("Area %.2f Perimeter %.2f", area, perimeter)
	return area, perimeter
}

func fundFunctionMultipleNamedReturns(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	fmt.Printf("\nArea %.2f Perimeter %.2f", area, perimeter)
	return //no explicit return value
}

func fundFunctionMultipleNamedReturnsBlank(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	fmt.Printf("\nArea %.2f Perimeter %.2f\n", area, perimeter)
	fmt.Println("\n\n------")
	return
}
