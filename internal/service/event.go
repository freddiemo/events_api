package service

import (
	"github.com/freddiemo/events_api/internal/repository"
	"github.com/freddiemo/events_api/models"
)

type EventService interface {
	CreateEvent(event *models.Event) error
	GetEvents() ([]*models.Event, error)
	GetEventByID(id int) (*models.Event, error)
}

type eventService struct {
	eventRepo repository.EventRepository
}

func NewEventService(eventRepo repository.EventRepository) EventService {
	return &eventService{eventRepo: eventRepo}
}

func (eventService *eventService) CreateEvent(event *models.Event) error {
	// Implementation for creating an event
	return nil
}

func (eventService *eventService) GetEvents() ([]*models.Event, error) {
	// Implementation for retrieving events
	return eventService.eventRepo.GetEvents()
}

func (eventService *eventService) GetEventByID(id int) (*models.Event, error) {
	// Implementation for retrieving a single event by ID
	return nil, nil
}
