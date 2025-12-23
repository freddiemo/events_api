package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/freddiemo/events_api/api"
	"github.com/freddiemo/events_api/internal/database"
	"github.com/freddiemo/events_api/internal/repository"
	"github.com/freddiemo/events_api/internal/service"
)

func main() {
	db, err := database.NewDatabase()
	fmt.Println(db.DB)
	if err != nil {
		log.Println("Failed to connect to database:", err)
		panic(err)
	}
	defer db.DB.Close()

	eventRepo := repository.NewEventRepository(db)
	eventService := service.NewEventService(eventRepo)
	eventHandler := api.NewEventHandler(eventService, nil)

	http.Handle("/events", http.HandlerFunc(eventHandler.GetEvents))
	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Println("Failed to start server:", err)
	}
}

/*
func eventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		eventHandler.GetEvents(w, r)
	} else if r.Method == http.MethodPost {
		eventHandler.CreateEvent(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}*/
