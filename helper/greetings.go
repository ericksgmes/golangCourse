package helper

import "fmt"

func InputLabel (inputType string) string {
	if len(inputType) == 0 {
		return "Invalid input"
	}
	return "Type in your " + inputType + " address: "
}

func GreetUser (confName string, confTickets uint, remainingTickets uint) {
	fmt.Printf("Get your tickets for %v here!\n", confName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", confTickets, remainingTickets)
}