package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Handlers *EventsHandlers
}

func NewServer(handlers *EventsHandlers) *Server {
	return &Server{
		Handlers: handlers,
	}
}

func (s *Server) StartServer() error {
	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("./static"))
	router.Handle("/", fs)

	router.Path("/event").Methods("POST").HandlerFunc(s.Handlers.EventHappenedHandler)
	router.Path("/eventlist").Methods("GET").HandlerFunc(s.Handlers.EventListHandler)
	router.Path("/user/{user_id}/history").Methods("GET").HandlerFunc(s.Handlers.UserEventsHistoryHandler)

	fmt.Println("Server is listening on localhost:8081")
	
	if err := http.ListenAndServe(":8081", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}

	return nil

}
