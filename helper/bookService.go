package helper


func Book(remainingTickets uint, bookings []string, userTickets uint, fullname string) (uint, []string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, fullname)
	return remainingTickets, bookings
}