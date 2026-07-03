package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reload/internal/models"
	"reload/internal/repository"
	"strconv"

	"github.com/gorilla/mux"
)

type EventsHandlers struct {
	Repo repository.EventStorage
}

func NewEventHandlers(repo repository.EventStorage) *EventsHandlers {
	return &EventsHandlers{
		Repo: repo,
	}
}

func sendError(rw http.ResponseWriter, code int, err error) {

	errDTO := NewErrorDTO(err.Error(), code)
	errBytes, err := json.Marshal(errDTO)
	if err != nil {
		fmt.Println("Can`t Marshal ErrorDTO: ", err)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(code)
	if _, err := rw.Write(errBytes); err != nil {
		fmt.Println("Can`t write http answer:", err.Error())
	}
}

func (e *EventsHandlers) EventHappenedHandler(rw http.ResponseWriter, r *http.Request) {

	var newEvent models.EventDTO

	if err := json.NewDecoder(r.Body).Decode(&newEvent); err != nil {
		sendError(rw, http.StatusBadRequest, errors.New("Incorrect request content: "+err.Error()))
		return
	}

	if err := e.Repo.SaveEvent(newEvent); err != nil {
		sendError(rw, http.StatusInternalServerError, errors.New("Can`t save event: "+ err.Error()))
		return
	}

	fmt.Println("Event save!")

	// Устанавливаем заголовок Content-Type
	rw.Header().Set("Content-Type", "application/json")
	// Отправляем JSON-объект
	rw.WriteHeader(http.StatusCreated)
	if _, err := rw.Write([]byte(`{"status":"ok","event":"saved"}`)); err != nil {
		fmt.Println("Can`t write http answer:", err.Error())
	}
}

func (e *EventsHandlers) EventListHandler(rw http.ResponseWriter, r *http.Request) {

	tempStorage, err := e.Repo.GetAllEvents()
	if err != nil {
		sendError(rw, http.StatusInternalServerError, errors.New("Can`t get data: "+err.Error()))
		return
	}

	tempStorageBytes, err := json.Marshal(tempStorage)

	if err != nil {
		sendError(rw, http.StatusInternalServerError, errors.New("Can`t create json with data: "+err.Error()))
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	if _, err := rw.Write(tempStorageBytes); err != nil {
		fmt.Println("Can`t write http answer:", err.Error())
	}
}

func (e *EventsHandlers) UserEventsHistoryHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["user_id"]

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		sendError(rw, http.StatusBadRequest, errors.New("Incorrect user id: "+err.Error()))
		return
	}

	userEvents, err := e.Repo.GetEventsByUserId(userIdInt)
	if err != nil {
		sendError(rw, http.StatusInternalServerError, errors.New("Can`t execute query: "+err.Error()))
		return
	}

	userEventsBytes, err := json.Marshal(userEvents)

	if err != nil {
		sendError(rw, http.StatusInternalServerError, errors.New("Can`t marshal data: "+err.Error()))
		return

	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	if _, err := rw.Write(userEventsBytes); err != nil {
		fmt.Println("Can`t write http answer:", err.Error())
	}
}
