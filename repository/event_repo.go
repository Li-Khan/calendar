package repository

import (
	"github.com/Li-Khan/calendar/domain"
	"sync"
	"time"
)

type eventRepo struct {
	events map[string]*domain.Event
	mutex  sync.RWMutex
}

func NewEventRepo(events map[string]*domain.Event) domain.EventRepository {
	return &eventRepo{
		events: events,
	}
}

func (e *eventRepo) Add(event *domain.Event) error {
	err := e.checkExist(event)
	if err != nil {
		return err
	}

	e.mutex.RLock()
	e.events[event.Name] = event
	e.mutex.RUnlock()

	return nil
}

func (e *eventRepo) Delete(name string) error {
	return nil
}

func (e *eventRepo) Update(event *domain.Event) error {
	return nil
}

func (e *eventRepo) List() *[]domain.Event {
	return nil
}

func (e *eventRepo) checkExist(event *domain.Event) error {
	if e.isNameAlreadyExist(event.Name) {
		return domain.ErrAlreadyExist
	}

	if e.isDateAlreadyExist(event.Date) {
		return domain.ErrDateAlreadyExist
	}

	return nil
}

func (e *eventRepo) isNameAlreadyExist(name string) bool {
	e.mutex.RLock()
	_, ok := e.events[name]
	e.mutex.RUnlock()
	return ok
}

func (e *eventRepo) isDateAlreadyExist(date time.Time) bool {
	e.mutex.RLock()
	for _, val := range e.events {
		if val.Date == date {
			return true
		}
	}
	e.mutex.RUnlock()

	return false
}
