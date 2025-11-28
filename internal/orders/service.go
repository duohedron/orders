package orders

import (
	"context"

	"github.com/google/uuid"
)

type EventType string

const EventOrderCreated EventType = "order.created"

type Event struct {
	Type    EventType
	OrderID uuid.UUID
}

type Service struct {
	store  Store
	events chan<- Event
}

func NewService(store Store, events chan<- Event) *Service {
	return &Service{store: store, events: events}
}

func (s *Service) Create(ctx context.Context, o *Order) error {
	if err := s.store.Create(ctx, o); err != nil {
		return err
	}
	s.events <- Event{Type: EventOrderCreated, OrderID: o.ID}
	return nil
}

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*Order, error) {
	return s.store.GetByID(ctx, id)
}
