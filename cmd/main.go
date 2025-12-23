package main

import (
	"net/http"

	"github.com/freddiemo/events_api/internal/database"
	"github.com/freddiemo/events_api/internal/service"
	"github.com/freddiemo/events_api/api"
)

func main() {
	db := database.NewDatabase()
	defer db.DB.Close()

	eventRepo := database.NewEventRepository(db)
	eventService := service.NewEventService(eventRepo)
	eventHandler := api.NewEventHandler(eventService, nil)

	http.Handle("/events", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			eventHandler.GetEvents(w, r)
		} else if r.Method == http.MethodPost {
			eventHandler.CreateEvent(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
	http.ListenAndServe(":8080", nil)
}

