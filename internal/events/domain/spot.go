package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidSpotNumber     = errors.New("invalid spot number")
	ErrSpotNotFound          = errors.New("spot not found")
	ErrSpotAlreadyReserved   = errors.New("spot already reserved")
	ErrSpotNameTwoCaracters  = errors.New("spot name must be at least 2 characters long")
	ErrSpotNameRequired      = errors.New("spot name is required")
	ErrSpotNameInvalidFormat = errors.New("spot name must start with a letter and end with a number")
)

type SpotStatus string

const (
	SpotStatusAvailable SpotStatus = "available"
	SpotStatusSold      SpotStatus = "sold"
)

type Spot struct {
	ID       string
	EventID  string
	Name     string
	Status   SpotStatus
	TicketID string
}

func NewSpot(event *Event, name string) (*Spot, error) {
	spot := &Spot{
		ID:      uuid.New().String(),
		EventID: event.ID,
		Name:    name,
		Status:  SpotStatusAvailable,
	}

	if err := spot.Validate(); err != nil {
		return nil, err
	}
	return spot, nil
}

func (s *Spot) Validate() error {
	if len(s.Name) == 0 {
		return ErrSpotNameRequired
	}
	if len(s.Name) < 2 {
		return ErrSpotNameTwoCaracters
	}
	if s.Name[0] < 'A' || s.Name[0] > 'Z' {
		return ErrSpotNameInvalidFormat
	}
	if s.Name[1] < '0' || s.Name[1] > '9' {
		return ErrSpotNameInvalidFormat
	}
	return nil
}

func (s *Spot) Reserve(ticketID string) error {
	if s.Status == SpotStatusSold {
		return ErrSpotAlreadyReserved
	}
	s.Status = SpotStatusSold
	s.TicketID = ticketID
	return nil
}
