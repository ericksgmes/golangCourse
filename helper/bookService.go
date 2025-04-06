package helper

type User struct {
	UserTickets uint
	FirstName string
	LastName string
	Email string
}

func Book(remainingTickets uint, bookings []User, user User) (uint, []User) {
	remainingTickets = remainingTickets - user.UserTickets
	bookings = append(bookings, user)
	return remainingTickets, bookings
}