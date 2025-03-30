package domain

import "errors"

type TicketType string

var (
	ErrInvalidTicketKind = errors.New("invalid ticket kind")
)

const (
	TicketTypeHalf TicketType = "half"
	TicketTypeFull TicketType = "full"
)

type Ticket struct {
	ID         string
	EventId    string
	Spot       *Spot
	TicketType TicketType
	Price      float64
}

func IsValidTicketKind(ticketKind TicketType) bool {
	return ticketKind == TicketTypeHalf || ticketKind == TicketTypeFull
}

func (t *Ticket) CalculatePrice() {
	if t.TicketType == TicketTypeHalf {
		t.Price /= 2
	}
}

func (t *Ticket) Validate() error {
	if t.Price <= 0 {
		return errors.New("ticket price must be greater than zero")
	}
	return nil
}
