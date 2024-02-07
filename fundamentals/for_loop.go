package main

import "fmt"

// https://go.dev/play/p/cnut6Lx33eB
/*
for initialisation; condition; post {
}
*/
// allNumbers demonstrates a basic for loop, printing numbers from 0 to num.
func allNumbers(num int) {
    for i := 0; i <= num; i++ {
        fmt.Println("Number:", i)
    }
}

// allNumbersBreak demonstrates using 'break' to exit a loop early.
// This loop exits when i equals 3.
func allNumbersBreak(num int) {
    for i := 0; i <= num; i++ {
        if i == 3 {
            fmt.Println("Breaking at", i)
            break
        }
        fmt.Println("Number:", i)
    }
}

// allNumbersContinue demonstrates skipping the rest of the loop body for the current iteration.
// This loop skips even numbers.
func allNumbersContinue(num int) {
    for i := 0; i <= num; i++ {
        if i%2 == 0 {
            continue // Skips printing even numbers
        }
        fmt.Println("Odd Number:", i)
    }
}

// allNumbersNestedLoop demonstrates nested for loops, printing a right-angled triangle of asterisks.
func allNumbersNestedLoop(num int) {
    for i := 0; i < num; i++ {
        for j := 0; j <= i; j++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
}

// demonstrateRangeLoop showcases iterating over a slice using the range clause.
func demonstrateRangeLoop() {
    fmt.Println("Iterating over slice with range loop:")
    numbers := []int{2, 3, 4, 5, 6}
    for index, value := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
}
