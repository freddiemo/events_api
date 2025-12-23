package database

import (
	"github.com/freddiemo/events_api/internal/models"

	"github.com/freddiemo/events_api/internal/database/database"
)

type EventRepository interface {
	CreateEvent(event *models.Event) error
	GetEvents() ([]*models.Event, error)
	GetEventByID(id int) (*models.Event, error)
}

type eventRepository struct {
	db *database.Database
}

func NewEventRepository(db *database.Database) EventRepository {
	return &eventRepository{db: db}
}

func (eventRepo *eventRepository) CreateEvent(event *models.Event) error {
	// Implementation for creating an event in the database
	return nil
}

func (eventRepo *eventRepository) GetEvents() ([]*models.Event, error) {
	// Implementation for retrieving events from the database
	return nil, nil
}

func (eventRepo *eventRepository) GetEventByID(id int) (*models.Event, error) {
	// Implementation for retrieving a single event by ID from the database
	return nil, nil
}



