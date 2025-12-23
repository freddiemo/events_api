package api

import (
	"encoding/json"
	"net/http"

	"github.com/freddiemo/events_api/internal/service"
	"github.com/freddiemo/events_api/models"
)

type EventHandler interface {
	CreateEvent(w http.ResponseWriter, r *http.Request) error
	GetEvents(w http.ResponseWriter, r *http.Request)
	GetEventByID(w http.ResponseWriter, r *http.Request, id int) (*models.Event, error)
}

type eventHandler struct {
	eventService service.EventService
	httpHandler  http.Handler
}

func NewEventHandler(eventService service.EventService, httpHandler http.Handler) EventHandler {
	return &eventHandler{eventService: eventService, httpHandler: httpHandler}
}

func (eventHandler *eventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) error {
	// Implementation for handling event creation requests
	return nil
}

func (eventHandler *eventHandler) GetEvents(w http.ResponseWriter, r *http.Request) {
	events, err := eventHandler.eventService.GetEvents()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	eventsJSON, _ := json.Marshal(events)
	w.Write(eventsJSON)
	w.WriteHeader(http.StatusOK)
}

func (eventHandler *eventHandler) GetEventByID(w http.ResponseWriter, r *http.Request, id int) (*models.Event, error) {
	// Implementation for handling requests to get a single event by ID
	return nil, nil
}
