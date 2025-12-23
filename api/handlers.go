package api

import (
	"net/http"

	"github.com/freddiemo/events_api/internal/models"
	"github.com/freddiemo/events_api/internal/service"
)

type EventHandler interface {
	CreateEvent(event *models.Event) error
	GetEvents() ([]*models.Event, error)
	GetEventByID(id int) (*models.Event, error)
}

type eventHandler struct {
	eventService service.EventService
	httpHandler  http.Handler
}

func NewEventHandler(eventService service.EventService, httpHandler HTTPHandler) EventHandler {
	return &eventHandler{eventService: eventService, httpHandler: httpHandler}
}

func (eventHandler *eventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) error {
	// Implementation for handling event creation requests
	return nil
}

func (eventHandler *eventHandler) GetEvents(w http.ResponseWriter, r *http.Request) ([]*models.Event, error) {
	// Implementation for handling requests to get events
	return nil, nil
}

func (eventHandler *eventHandler) GetEventByID(w http.ResponseWriter, r *http.Request, id int) (*models.Event, error) {
	// Implementation for handling requests to get a single event by ID
	return nil, nil
}
