package techpalace

import (
	"fmt"
	"strings"
)

/*
For most customers who scan their loyalty cards, the store owner wants to see `Welcome to the Tech Palace, ` followed by the name of the customer in capital letters on the display.

Implement the function `WelcomeMessage` that accepts the name of the customer as a `string` argument and returns the desired message as a `string`.
*/

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

/*

Write a function `AddBorder` that accepts a welcome message (a `string`) and the number of stars per line (type `int`) as arguments.
It should return a `string` that consists of 3 lines, a line with the desired number of stars, then the welcome message as it was passed in, then another line of stars.
*/
// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	return fmt.Sprintf("%s\n%s\n%s", strings.Repeat("*", numStarsPerLine), welcomeMsg, strings.Repeat("*", numStarsPerLine))
}

/*
Implement a function `CleanUpMessage` that accepts the old marketing message as a string.
The function should first remove all stars from the text and afterwards remove the leading and trailing whitespaces from the remaining text.
The function should then return the cleaned up message.
*/
// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	cleanup := strings.ReplaceAll(oldMsg, "*", "")
	cleanup = strings.TrimSpace(cleanup)
	return cleanup
}
