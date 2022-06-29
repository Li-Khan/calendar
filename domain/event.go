package domain

import "time"

type Event struct {
	Name string
	Date time.Time
}

type EventUsecase interface {
	Add(event *Event) error
	Delete(name string)
	UpdateName(old string, new string) error
	UpdateDate(name string, date time.Time) error
	List() *[]Event
}

type EventRepository interface {
	Add(event *Event) error
	Delete(name string)
	UpdateName(old string, new string) error
	UpdateDate(name string, date time.Time) error
	List() *[]Event
}
