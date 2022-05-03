package helpers

import (
	"fmt"
	"strings"
)

const Line string = "----------------------------------------------------------------------------------------------------"

func ValidateUserInput(firstName, lastName, email string, userTickets, remainingTickets uint16) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func ShowValidationErrors(isValidName, isValidEmail, isValidTicketNumber bool) {
	fmt.Println()

	if !isValidName {
		fmt.Println("First name or last name you entered is too short. 😕")
	}
	if !isValidEmail {
		fmt.Println("Email you entered is invalid. 😕")
	}
	if !isValidTicketNumber {
		fmt.Println("Number of tickets you entered is invalid. 😕")
	}

	fmt.Println(Line)
	fmt.Println("\nTry again 🙂")
}
