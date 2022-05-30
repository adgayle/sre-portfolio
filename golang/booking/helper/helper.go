package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, emailAddress string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmailAddress := strings.Contains(emailAddress, "@") && len(emailAddress) > 5
	isValidUserTickets := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmailAddress, isValidUserTickets
}
