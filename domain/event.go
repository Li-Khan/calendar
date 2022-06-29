package domain

import "time"

type Event struct {
	Name string
	Date time.Time
}

type EventUsecase interface {
	Add(event *Event) error
	Delete(name string) error
	Update(event *Event) error
	List() []Event
}

type EventRepository interface {
	Add(event *Event) error
	Delete(name string) error
	Update(event *Event) error
	List() *[]Event
}
