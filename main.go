package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to our booking app")
	const confTickets = 50
	var remainingTickets uint = 50
	var confName = "Go Conference"
	bookings := []string{}

	var userTickets uint
	var firstName string
	var lastName string
	var email string

	for {
		fmt.Printf("Get your tickets for %v here!\n", confName)
		fmt.Printf("We have a total of %v tickets and %v are still available.\n", confTickets, remainingTickets)
		
		fmt.Print("Type in your email address: ")
		fmt.Scan(&email)

		fmt.Print("Type in your first name: ")
		fmt.Scan(&firstName)

		fmt.Print("Type in your last name: ")
		fmt.Scan(&lastName)

		fmt.Print("Enter number of tickets you want to buy: ")
		fmt.Scanln(&userTickets)
		
		if userTickets > remainingTickets {continue}

		fullname := firstName + " " + lastName
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, fullname)

		fmt.Printf("Thank you %v for buying %v tickets with us!!\n" + 
		"You will recieve a confirmation email @ %v\n",
		firstName + " " + lastName, userTickets, email)

		for _, booking := range bookings {
			var names = strings.Fields(booking)
			fmt.Println("\n" + names[0] + "\n")
		}

		fmt.Printf("%v remaining tickets in %v\n", remainingTickets, confName)

		if remainingTickets == 0 {return}
	}
}
