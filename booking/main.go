package main

import (
	"fmt"
	"strings"
)

func main() {
	const conferenceName = "Developer Conference"
	const conferenceTickets = 50

	var remainingTickets uint = conferenceTickets
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	for remainingTickets > 0 {
		var firstName string
		var lastName string
		var emailAddress string
		var userTickets uint
		var privateNames []string

		fmt.Println()
		fmt.Print("Please enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Print("Please enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Print("Please enter your email address: ")
		fmt.Scan(&emailAddress)
		fmt.Print("Please enter the number of tickets you would like to purchase: ")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			fmt.Println()
			fmt.Printf("Thank you %v %v for booking %v tickets.\n", firstName, lastName, userTickets)
			fmt.Printf("You will receive a confirmation at your email address %v.\n", emailAddress)

			fmt.Println()
			fmt.Printf("There are %v tickets still available for %v.\n", remainingTickets, conferenceName)

			bookings = append(bookings, firstName+" "+lastName)
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				privateNames = append(privateNames, names[0]+" "+names[1][0:1])
			}

			fmt.Printf("The bookings are: %v\n", privateNames)
		} else {
			fmt.Printf("We only have %v tickets available, you wanted %v.\n", remainingTickets, userTickets)
			continue
		}

		if remainingTickets == 0 {
			fmt.Printf("Tickets for %v are sold out!\n", conferenceName)
			break
		}
	}
}
