package main

import (
	"fmt"
	"go-booking-app/helpers"
	"net/smtp"
	"strconv"
	"sync"
)

const conferenceTickets uint16 = 50

var conferenceName string = "\"Go Conference\""
var remainingTickets uint16 = conferenceTickets
var bookings = make([]UserData, 0)

var firstName string
var lastName string
var email string
var userTickets uint16

type UserData struct {
	firstName string
	lastName  string
	email     string
	tickets   uint16
}

var wg = sync.WaitGroup{}

func main() {
	greetUser()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helpers.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket()
			informUser()

			wg.Add(1)
			go sendTicket()

			printUserNames()

			if remainingTickets == 0 {
				break
			}
		} else {
			helpers.ShowValidationErrors(isValidName, isValidEmail, isValidTicketNumber)
		}
	}

	wg.Wait()
}

func greetUser() {
	fmt.Println(helpers.Line)
	fmt.Printf("Welcome to our %v booking application 🙋\n", conferenceName)
	fmt.Println(helpers.Line)
	fmt.Println("Get your tickets(🎫) here to attend")
	fmt.Println("We have total of", conferenceTickets, "tickets(🎫)! and", remainingTickets, "are still available.")
}

func getUserInput() (string, string, string, uint16) {
	fmt.Print("\nEnter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email: ")
	fmt.Scan(&email)
	fmt.Printf("How many tickets(🎫) do you want? (Remaining tickets(🎫): %v): ", remainingTickets)
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket() {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		tickets:   userTickets,
	}

	bookings = append(bookings, userData)
}

func printUserNames() {
	userNames := []string{}

	for _, booking := range bookings {
		userNames = append(userNames, booking.firstName+" "+booking.lastName)
	}

	fmt.Println("These are our all bookings: ", userNames)
	fmt.Println(helpers.Line)
}

func informUser() {
	fmt.Println(helpers.Line)
	fmt.Printf("Thank you \"%v\" for booking %v tickets(🎫). 😄\n", firstName+" "+lastName, userTickets)

	if remainingTickets != 0 {
		fmt.Printf("%v tickets reamining for %v\n", remainingTickets, conferenceName)
	}

	fmt.Println(helpers.Line)
}

func sendTicket() {
	from := "toghrulnasirli111@gmail.com"
	password := "pjpsyurwsgryqmmy"
	to := []string{email}
	host := "smtp.gmail.com"
	port := "587"

	body := "Hello " + firstName + " " + lastName + ". Here are " + strconv.FormatUint(uint64(userTickets), 10) + " ticktes for " + conferenceName
	message := []byte(body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(host+":"+port, auth, from, to, message)

	if err != nil {
		panic(err)
	}

	wg.Done()
}
