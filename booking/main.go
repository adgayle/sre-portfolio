package main

import "fmt"

func main() {
	const conferenceName = "Developer Conference"
	const conferenceTickets = 50

	var remainingTickets uint = conferenceTickets

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	var firstName string
	var lastName string
	var emailAddress string
	var userTickets uint
	var bookings []string

	fmt.Println()
	fmt.Print("Please enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Please enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Please enter your email address: ")
	fmt.Scan(&emailAddress)
	fmt.Print("Please enter the number of tickets you would like to purchase: ")
	fmt.Scan(&userTickets)

	fmt.Println()
	fmt.Printf("Thank you %v %v for booking %v tickets.\n", firstName, lastName, userTickets)
	fmt.Printf("You will receive a confirmation at your email address %v.\n", emailAddress)

	remainingTickets = remainingTickets - userTickets
	fmt.Println()
	fmt.Printf("There are %v tickets still available for %v.\n", remainingTickets, conferenceName)

	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("The bookings are: %v\n", bookings)
}
