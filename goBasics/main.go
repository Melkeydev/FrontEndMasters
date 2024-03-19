package main

import (
	"femBasics/imports"
	"fmt"
)

func main() {
	newTicket := imports.Ticket{
		Number: 123,
		Event:  "theater",
	}

	newTicket.GetTicketNumber()

	fmt.Println(newTicket)
}
