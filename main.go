package main

import (
	"fmt"
)

const conferenceTicket int = 50
var conferenceName = "Go Conference"
var remainingTicket uint = 50
var bookings = make([] UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	userOfTickets uint
}

func main() {
	greetUser()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)


			bookTickets(firstName, lastName, email, userTickets)


		if isValidEmail && isValidName && isValidTicketNumber {
			remainingTicket = remainingTicket - userTickets

				fmt.Printf("\nHi %v %v, your email is %v and you want %v tickets\n", firstName, lastName, email, userTickets)

					// print first names
			firstNames :=	getFirstNames()

			fmt.Printf("The first names of bookings is %v", firstNames)

				if remainingTicket == 0 {
					// end program
					fmt.Println("Sorry, we are sold out!")
					break
				}
		} else {
			fmt.Printf("Sorry, we have %v tickets remaining", remainingTicket)

		}
	}
}

func greetUser()  {
	fmt.Printf("Welcome to %v booking app!\n", conferenceName)
	fmt.Printf("There are total of %v and you have %v tickets remaining.\n",conferenceTicket, remainingTicket)
	fmt.Println("Get your tickets here to begin")
}

func getFirstNames () []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)

	}

	return firstNames
}



func getUserInput () (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Print("Enter your name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email: ")
	fmt.Scan(&email)
	fmt.Print("Enter the number of tickets you want: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(firstName string, lastName string, email string, userTickets uint) {
	remainingTicket = remainingTicket - userTickets

		// create a map for a user
		var userData = UserData {
			firstName: firstName,
			lastName: lastName,
			email: email,
			userOfTickets: userTickets,
		}




		bookings = append(bookings, userData)

		fmt.Printf("List of bookings: %v\n", bookings)

		fmt.Printf("The whole slice %v\n", bookings)
		fmt.Printf("The first element %v\n", bookings[0])
		fmt.Printf("The last element %T\n", bookings)
		fmt.Printf("The length of the slice %v\n", len(bookings))
}

