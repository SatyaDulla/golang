package main

import "fmt"

// https://go.dev/play/p/fzS67oeL0zK

/* If statement syntax
if condition {
}
*/

/* If else statement
if condition {
} else {
}
*/

/* If … else if … else statement
if condition1 {
...
} else if condition2 {
...
} else {
...
}
*/

/*
if assignment-statement; condition {
}
*/
// calculateEven prints whether a number is even. It demonstrates a basic if statement.
func calculateEven(num int) {
	if num%2 == 0 {
		fmt.Println("The number is even:", num)
		// Using return here allows us to avoid the else block, making the code cleaner
		return
	}
	// No need for an else block due to the return statement above
	fmt.Println("The number is not even:", num)
}

// demonstrateIfElseIfElse shows a chained if … else if … else statement.
func demonstrateIfElseIfElse(num int) string {
	if num < 0 {
		return "The number is negative"
	} else if num == 0 {
		return "The number is zero"
	} else {
		return "The number is positive"
	}
}

// demonstrateShortStatement demonstrates the if statement with a short statement.
// The short statement (declaration) precedes the condition and is scoped to the if and else blocks.
func demonstrateShortStatement(num int) string {
	if mod := num % 2; mod == 0 {
		return fmt.Sprintf("The number %d is even", num)
	} else {
		// mod is also available here due to the short statement syntax
		return fmt.Sprintf("The number %d is odd, with a remainder of %d when divided by 2", num, mod)
	}
}
