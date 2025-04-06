package main

import (
	"bookingApp/helper"
	"fmt"
)

const GlobalVariableJustForFunsies = 1
const confTickets = 50
var remainingTickets uint = 50
var confName = "Go Conference"
var bookings = make([]helper.User, 0)

func main() {
	fmt.Println("Welcome to our booking app")

	var user helper.User

	for {
		helper.GreetUser(confName, confTickets, remainingTickets)

		fmt.Println(helper.InputLabel("email"))
		fmt.Scan(&user.Email)

		fmt.Println(helper.InputLabel("first name"))
		fmt.Scan(&user.FirstName)

		fmt.Println(helper.InputLabel("last name"))
		fmt.Scan(&user.LastName)

		fmt.Println(helper.InputLabel("tickets", true))
		fmt.Scanln(&user.UserTickets)
		
		if user.UserTickets > remainingTickets {continue}

		remainingTickets, bookings = helper.Book(remainingTickets, bookings, user)

		helper.ShowBooked(bookings, remainingTickets, confName)
		//fmt.Println(GlobalVariableJustForFunsies)
		if remainingTickets == 0 {return}
	}
}
