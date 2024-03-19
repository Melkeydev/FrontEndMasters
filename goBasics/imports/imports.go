package imports

import "fmt"

type Ticket struct {
	Number int
	Event  string
}

func (t Ticket) GetTicketNumber() {
	fmt.Println(t.Number)
}
