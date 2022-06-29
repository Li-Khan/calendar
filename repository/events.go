package repository

import "github.com/Li-Khan/calendar/domain"

func NewEventsRepository() map[string]*domain.Event {
	events := make(map[string]*domain.Event)
	return events
}
