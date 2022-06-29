package usecase

import (
	"github.com/Li-Khan/calendar/domain"
	"time"
)

type eventUcase struct {
	eventRepo domain.EventRepository
}

func NewEventUsecase(e domain.EventRepository) domain.EventUsecase {
	return &eventUcase{
		eventRepo: e,
	}
}

func (e *eventUcase) Add(event *domain.Event) error {
	err := e.eventRepo.Add(event)
	if err != nil {
		return err
	}
	return nil
}

func (e *eventUcase) Delete(name string) {
	e.eventRepo.Delete(name)
}

func (e *eventUcase) UpdateName(old string, new string) error {
	err := e.eventRepo.UpdateName(old, new)
	if err != nil {
		return err
	}
	return nil
}

func (e *eventUcase) UpdateDate(name string, date time.Time) error {
	err := e.eventRepo.UpdateDate(name, date)
	if err != nil {
		return err
	}
	return nil
}

func (e *eventUcase) List() *[]domain.Event {
	return e.eventRepo.List()
}
