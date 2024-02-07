// Package simpleinterest calculates the simple interest for a given principal, rate, and time.
package simpleinterest

import "fmt"

// init function is called automatically when the package is initialized.
func init() {
    fmt.Println("Simple Interest package initialized")
}

// Calculate computes and returns the simple interest for a principal amount p,
// an annual interest rate r, and time duration t (in years).
// The function is exported because it starts with an uppercase letter.
func Calculate(p, r, t float64) float64 {
    interest := p * (r / 100) * t
    return interest
}
