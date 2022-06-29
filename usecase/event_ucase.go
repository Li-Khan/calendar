package usecase

import "github.com/Li-Khan/calendar/domain"

type eventUcase struct {
	eventRepo domain.EventRepository
}

func NewEventUsecase(e domain.EventRepository) domain.EventUsecase {
	return &eventUcase{
		eventRepo: e,
	}
}

func (e *eventUcase) Add(event *domain.Event) error {
	return nil
}

func (e *eventUcase) Delete(id int64) error {
	return nil
}

func (e *eventUcase) Update(event *domain.Event) error {
	return nil
}

func (e *eventUcase) List() *[]domain.Event {
	return nil
}
