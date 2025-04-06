package helper

import (
	"fmt"
	"strings"
)

func InputLabel (inputType string, flags... bool) string {
	tickets := false
	if len(flags) > 0 {
		tickets = flags[0]
	}

	if tickets {
		return "Enter number of tickets you want to buy: "
	}
	if len(inputType) == 0 {
		return "Invalid input"
	}
	return "Type in your " + inputType + ": "
}

func GreetUser (confName string, confTickets uint, remainingTickets uint) {
	fmt.Printf("Get your tickets for %v here!\n", confName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", confTickets, remainingTickets)
}

func ThankUser (fullname string, userTickets uint, email string) {
	fmt.Printf("Thank you %v for buying %v tickets with us!!\n" + 
	"You will recieve a confirmation email @ %v\n",
	fullname, userTickets, email)
}

func ShowBooked (bookings []string, remainingTickets uint, confName string) {
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		fmt.Println("\n" + names[0] + "\n")
	}
	fmt.Printf("%v remaining tickets in %v\n", remainingTickets, confName)
}