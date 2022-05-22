package main

import (
	"fmt"
	"strings"
)

const conferenceName = "Developer Conference"
const conferenceTickets = 50

var remainingTickets uint = conferenceTickets
var bookings = []string{}

func main() {

	greetUser()

	for remainingTickets > 0 {
		firstName, lastName, emailAddress, userTickets := getUserInput()

		isValidName, isValidEmailAddress, isValidUserTickets := validateUserInput(firstName, lastName, emailAddress, userTickets)

		if isValidUserTickets && isValidName && isValidEmailAddress {
			bookTicket(userTickets, firstName, lastName, emailAddress)
			printFirstNameLastInitial()

			if remainingTickets == 0 {
				fmt.Printf("Tickets for %v are sold out!\n", conferenceName)
				break
			}
		} else {
			if !isValidName {
				fmt.Println("The first or last name you entered must be 2 or more characters.")
			}
			if !isValidEmailAddress {
				fmt.Println("The email address must contain the @ character.")
			}
			if !isValidUserTickets {
				fmt.Printf("The number of tickets must be more than 0 but less than the number we have available %v.\n", remainingTickets)
			}
			fmt.Println("Please try again.")
			continue
		}
	}
}

func greetUser() {

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf(
		"We have a total of %v tickets and %v are still available.\n",
		conferenceTickets,
		remainingTickets,
	)
	fmt.Println("Get your tickets here to attend.")
}

func printFirstNameLastInitial() {
	privateNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		privateNames = append(privateNames, names[0]+" "+names[1][0:1])
	}
	fmt.Printf("The bookings are: %v\n", privateNames)
}

func validateUserInput(firstName string, lastName string, emailAddress string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmailAddress := strings.Contains(emailAddress, "@") && len(emailAddress) > 5
	isValidUserTickets := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmailAddress, isValidUserTickets
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var emailAddress string
	var userTickets uint

	fmt.Println()
	fmt.Print("Please enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Please enter your email address: ")
	fmt.Scan(&emailAddress)

	fmt.Print("Please enter the number of tickets you would like to purchase: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, emailAddress, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, emailAddress string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Println()
	fmt.Printf("Thank you %v %v for booking %v tickets.\n", firstName, lastName, userTickets)
	fmt.Printf("You will receive a confirmation at your email address %v.\n", emailAddress)

	fmt.Println()
	fmt.Printf("There are %v tickets still available for %v.\n", remainingTickets, conferenceName)
}
