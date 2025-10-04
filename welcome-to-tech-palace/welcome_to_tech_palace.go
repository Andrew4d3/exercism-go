package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	lineOfStars := strings.Repeat("*", numStarsPerLine)

	return fmt.Sprintf("%s\n%s\n%s", lineOfStars, welcomeMsg, lineOfStars)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.Trim(oldMsg, "\n*\x20") // x20 is the hex code for space
}
