package calendar

import (
	"github.com/Li-Khan/calendar/domain"
	"sort"
	"sync"
	"time"
)

type Event struct {
	Name string
	Date time.Time
	data eventData
}

type eventData struct {
	mutex  sync.Mutex
	events map[string]*Event
}

func NewEvent() *Event {
	return &Event{
		data: eventData{
			events: make(map[string]*Event),
		},
	}
}

func (e *Event) Add(name string, date time.Time) error {
	err := e.checkExist()
	if err != nil {
		return err
	}

	e.data.mutex.Lock()
	e.data.events[e.Name] = e
	e.data.mutex.Unlock()
	return nil
}

func (e *Event) UpdateName(old string, new string) error {
	e.data.mutex.Lock()
	event := e.data.events[old]
	e.data.mutex.Unlock()

	event.Name = new

	err := e.Add(event.Name, event.Date)
	if err != nil {
		return err
	}

	e.Delete(old)

	return nil
}
func (e *Event) UpdateDate(name string, date time.Time) error {
	e.data.mutex.Lock()
	e.data.events[name].Date = date
	e.data.mutex.Unlock()
	return nil
}

func (e *Event) List() *[]Event {
	var events []Event

	e.data.mutex.Lock()
	for _, event := range e.data.events {
		events = append(events, *event)
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].Date.Before(events[j].Date)
	})
	e.data.mutex.Unlock()

	return &events
}

func (e *Event) Delete(name string) {
	e.data.mutex.Lock()
	delete(e.data.events, name)
	e.data.mutex.Unlock()
}

func (e *Event) checkExist() error {
	if e.isNameAlreadyExist(e.Name) {
		return domain.ErrAlreadyExist
	}

	if e.isDateAlreadyExist(e.Date) {
		return domain.ErrDateAlreadyExist
	}

	return nil
}

func (e *Event) isNameAlreadyExist(name string) bool {
	e.data.mutex.Lock()
	_, ok := e.data.events[name]
	e.data.mutex.Unlock()
	return ok
}

func (e *Event) isDateAlreadyExist(date time.Time) bool {
	e.data.mutex.Lock()
	for _, val := range e.data.events {
		if val.Date == date {
			return true
		}
	}
	e.data.mutex.Unlock()

	return false
}
