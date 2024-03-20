package imports

type Ticket struct {
	ID    int
	Event string
}

func (t *Ticket) ChangeEvent(newEvent string) {
	t.Event = newEvent
}
