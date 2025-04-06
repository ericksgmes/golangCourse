package main

import (
	"fmt"
	"bookingApp/helper"
)

const GlobalVariableJustForFunsies = 1
const confTickets = 50
var remainingTickets uint = 50
var confName = "Go Conference"
var bookings = []string{}

func main() {
	fmt.Println("Welcome to our booking app")

	var userTickets uint
	var firstName string
	var lastName string
	var email string

	for {
		helper.GreetUser(confName, confTickets, remainingTickets)

		fmt.Println(helper.InputLabel("email"))
		fmt.Scan(&email)

		fmt.Println(helper.InputLabel("first name"))
		fmt.Scan(&firstName)

		fmt.Println(helper.InputLabel("last name"))
		fmt.Scan(&lastName)

		fmt.Println(helper.InputLabel("tickets", true))
		fmt.Scanln(&userTickets)
		
		if userTickets > remainingTickets {continue}

		fullname := firstName + " " + lastName
		remainingTickets, bookings = helper.Book(remainingTickets, bookings, userTickets, fullname)

		helper.ShowBooked(bookings, remainingTickets, confName)
		//fmt.Println(GlobalVariableJustForFunsies)
		if remainingTickets == 0 {return}
	}
}
