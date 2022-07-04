package calendar

import (
	"errors"
	"sort"
	"sync"
	"time"
)

type Calendar struct {
	mutex  sync.Mutex
	events map[string]*Event
}

type Event struct {
	Name string    `json:"name,omitempty"`
	Date time.Time `json:"date,omitempty"`
}

var (
	ErrNameAlreadyExist error = errors.New("name already exist")
	ErrDateAlreadyExist error = errors.New("date already exist")
	ErrEventNotExist    error = errors.New("the event does not exist")
)

func NewCalendar() *Calendar {
	return &Calendar{
		events: make(map[string]*Event),
	}
}

func (c *Calendar) Add(name string, date time.Time) error {
	event := Event{
		Name: name,
		Date: date,
	}
	err := c.checkExist(event)
	if err != nil {
		return err
	}

	c.mutex.Lock()
	c.events[name] = &event
	c.mutex.Unlock()

	return nil
}

func (c *Calendar) UpdateName(old string, new string) error {
	c.mutex.Lock()
	event := c.events[old]
	if event == nil {
		c.mutex.Unlock()
		return ErrEventNotExist
	}
	c.mutex.Unlock()

	event.Name = new
	c.Delete(old)

	err := c.Add(event.Name, event.Date)
	if err != nil {
		return err
	}

	return nil
}
func (c *Calendar) UpdateDate(name string, date time.Time) error {
	c.mutex.Lock()
	var err error
	_, ok := c.events[name]
	if !ok {
		err = ErrEventNotExist
	} else {
		c.events[name].Date = date
	}
	c.mutex.Unlock()
	return err
}

func (c *Calendar) List() *[]Event {
	var events []Event

	c.mutex.Lock()
	for _, event := range c.events {
		events = append(events, *event)
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].Date.Before(events[j].Date)
	})
	c.mutex.Unlock()

	return &events
}

func (c *Calendar) ListEventsForDay() (*[]Event, error) {
	var events []Event

	c.mutex.Lock()
	now, err := timeNow()
	if err != nil {
		return nil, err
	}

	for _, event := range c.events {
		if inTimeSpan(now, now.Add(24*time.Hour), event.Date) {
			events = append(events, *event)
		}
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].Date.Before(events[j].Date)
	})
	c.mutex.Unlock()

	return &events, nil
}

func (c *Calendar) ListEventsForWeek() (*[]Event, error) {
	var events []Event

	c.mutex.Lock()
	now, err := timeNow()
	if err != nil {
		return nil, err
	}

	for _, event := range c.events {
		if inTimeSpan(now, now.Add((24*time.Hour)*7), event.Date) {
			events = append(events, *event)
		}
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].Date.Before(events[j].Date)
	})
	c.mutex.Unlock()

	return &events, nil
}

func (c *Calendar) ListEventsForMonth() (*[]Event, error) {
	var events []Event

	c.mutex.Lock()
	now, err := timeNow()
	if err != nil {
		return nil, err
	}

	for _, event := range c.events {
		if inTimeSpan(now, now.Add((24*time.Hour)*30), event.Date) {
			events = append(events, *event)
		}
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].Date.Before(events[j].Date)
	})
	c.mutex.Unlock()

	return &events, nil
}

func (c *Calendar) Delete(name string) error {
	c.mutex.Lock()
	if !c.isNameAlreadyExist(name) {
		c.mutex.Unlock()
		return ErrEventNotExist
	}
	delete(c.events, name)
	c.mutex.Unlock()
	return nil
}

func (c *Calendar) checkExist(event Event) error {
	if c.isNameAlreadyExist(event.Name) {
		return ErrNameAlreadyExist
	}

	if c.isDateAlreadyExist(event.Date) {
		return ErrDateAlreadyExist
	}

	return nil
}

func (c *Calendar) isNameAlreadyExist(name string) bool {
	c.mutex.Lock()
	_, ok := c.events[name]
	c.mutex.Unlock()
	return ok
}

func (c *Calendar) isDateAlreadyExist(date time.Time) bool {
	c.mutex.Lock()
	for _, val := range c.events {
		if val.Date == date {
			c.mutex.Unlock()
			return true
		}
	}
	c.mutex.Unlock()

	return false
}

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func timeNow() (time.Time, error) {
	now, err := time.Parse(time.Now().Format("2006-01-02"), "2006-01-02")
	if err != nil {
		return time.Time{}, err
	}
	return now, err
}
