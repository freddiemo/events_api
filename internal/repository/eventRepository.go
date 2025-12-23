package repository

import (
	"fmt"
	"log"

	"github.com/freddiemo/events_api/internal/database"
	"github.com/freddiemo/events_api/models"
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
	eventsDb, err := eventRepo.db.DB.Query("SELECT * FROM events")
	if err != nil {
		log.Println("error getting events")
		return nil, err
	}
	events := []*models.Event{}
	for eventsDb.Next() {
		event := &models.Event{}
		err := eventsDb.Scan(&event.ID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	fmt.Println(events)

	return events, nil
}

func (eventRepo *eventRepository) GetEventByID(id int) (*models.Event, error) {
	// Implementation for retrieving a single event by ID from the database
	return nil, nil
}
