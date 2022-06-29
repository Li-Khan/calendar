package repository

import (
	"github.com/Li-Khan/calendar/domain"
	"sort"
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

func (e *eventRepo) Delete(name string) {
	e.mutex.RLock()
	delete(e.events, name)
	e.mutex.RUnlock()
}

func (e *eventRepo) UpdateName(old string, new string) error {
	e.mutex.RLock()
	event := e.events[old]
	e.mutex.RUnlock()

	event.Name = new

	err := e.Add(event)
	if err != nil {
		return err
	}

	e.Delete(old)

	return nil
}

func (e *eventRepo) UpdateDate(name string, date time.Time) error {
	e.mutex.RLock()
	e.events[name].Date = date
	e.mutex.RUnlock()
	return nil
}

func (e *eventRepo) List() *[]domain.Event {
	var events []domain.Event

	e.mutex.RLock()
	for _, event := range e.events {
		events = append(events, *event)
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].Date.Before(events[j].Date)
	})
	e.mutex.RUnlock()

	return &events
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
