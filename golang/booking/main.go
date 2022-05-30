package main

import (
	"fmt"
	"sync"
	"time"

	"booking/helper"
)

const conferenceName = "Developer Conference"
const conferenceTickets = 50

var remainingTickets uint = conferenceTickets
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	firstName, lastName, emailAddress, userTickets := getUserInput()

	isValidName, isValidEmailAddress, isValidUserTickets := helper.ValidateUserInput(firstName, lastName, emailAddress, userTickets, remainingTickets)

	if isValidUserTickets && isValidName && isValidEmailAddress {
		bookTicket(userTickets, firstName, lastName, emailAddress)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, emailAddress)

		printFirstNameLastInitial()

		if remainingTickets == 0 {
			fmt.Printf("Tickets for %v are sold out!\n", conferenceName)
		}
	} else {
		if !isValidName {
			fmt.Println("The first or last name you entered must be 2 or more characters.")
		}
		if !isValidEmailAddress {
			fmt.Println("The email address must contain the @ character and have at least 5 characters.")
		}
		if !isValidUserTickets {
			fmt.Printf("The number of tickets must be more than 0 but less than the number we have available %v.\n", remainingTickets)
		}
		fmt.Println("Please try again.")
	}
	wg.Wait()
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
	firstNameLastInitial := []string{}

	for _, booking := range bookings {
		firstNameLastInitial = append(firstNameLastInitial, booking.firstName+" "+booking.lastName[0:1])
	}
	fmt.Printf("The bookings are: %v\n", firstNameLastInitial)
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
	var userData UserData

	userData.firstName = firstName
	userData.lastName = lastName
	userData.email = emailAddress
	userData.numberOfTickets = userTickets

	remainingTickets = remainingTickets - userTickets

	bookings = append(bookings, userData)

	fmt.Println()
	fmt.Printf("Thank you %v %v for booking %v tickets.\n", firstName, lastName, userTickets)
	fmt.Printf("You will receive a confirmation at your email address %v.\n", emailAddress)

	fmt.Println()
	fmt.Printf("There are %v tickets still available for %v.\n", remainingTickets, conferenceName)
	fmt.Println()
	fmt.Printf("Map of bookings %v\n", bookings)
}

func sendTicket(userTickets uint, firstName string, lastName string, emailAddress string) {
	time.Sleep(10 * time.Second)

	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println(("#############################################################"))
	fmt.Printf("Sending ticket: \n%v to email address %v\n", ticket, emailAddress)
	fmt.Println(("#############################################################"))
	wg.Done()
}
