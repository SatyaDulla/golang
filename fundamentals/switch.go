package main

//import "fmt"

// parseCard returns the integer value of a card following blackjack ruleset.
func parseCard(card string) int {
    // Using switch statement to match card strings to their values
    switch card {
    case "ace":
        return 11
    case "two":
        return 2
    case "three":
        return 3
    // Grouping cases with the same return value for brevity
    case "four", "five", "six", "seven", "eight", "nine":
        return int(card[0]-'0') + 1 // Convert character to int and adjust for zero-indexing
    case "ten", "jack", "queen", "king":
        return 10
    default:
        return 0
    }
}

// firstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func firstTurn(card1, card2, dealerCard string) string {
    c1, c2, d := parseCard(card1), parseCard(card2), parseCard(dealerCard)

    // Using a switch statement without an expression to create a "tagless" switch,
    // which functions like a series of if-else statements but is cleaner.
    switch {
    case c1 == 11 && c2 == 11:
        return "P" // Pair of aces - Split
    case c1+c2 == 21 && d < 10:
        return "W" // Blackjack and dealer shows a card less than 10 - Win
    case c1+c2 == 21 && d >= 10:
        return "S" // Blackjack and dealer shows 10 or a face card - Stand
    case c1+c2 <= 20 && c1+c2 >= 17:
        return "S" // Hand is between 17 and 20 - Stand
    case c1+c2 <= 16 && c1+c2 >= 12 && d < 7:
        return "S" // Hand is between 12 and 16 and dealer shows less than 7 - Stand
    case c1+c2 <= 16 && c1+c2 >= 12 && d >= 7:
        return "H" // Hand is between 12 and 16 and dealer shows 7 or higher - Hit
    case c1+c2 <= 11:
        return "H" // Hand is 11 or less - Hit
    default:
        return "N" // Unexpected case
    }
}
