package helper

import "strings"

// export function dalam golang hanya memberi huruf besar pada function

func ValidateUserInput(firstName string, lastName string, userTickets uint, remainingTickets uint, email string) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	isValidEmail := strings.Contains(email, "@")
	return isValidEmail, isValidName, isValidTicketNumber
}