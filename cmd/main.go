package main

import (
	"fmt"
	"github.com/Li-Khan/calendar/repository"
)

func main() {
	events := repository.NewEventsRepository()
	eventsRepo := repository.NewEventRepo(events)
	fmt.Println(eventsRepo)
}
