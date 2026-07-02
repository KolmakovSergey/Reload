package main

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

var storage repository.EventStorage

func SendError(rw http.ResponseWriter, code int, err error) {

	errDTO := models.NewErrorDTO(err.Error(), code)
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

func eventHandler(rw http.ResponseWriter, r *http.Request) {

	var newEvent models.EventDTO

	if err := json.NewDecoder(r.Body).Decode(&newEvent); err != nil {
		SendError(rw, http.StatusBadRequest, errors.New("Incorrect request content: "+err.Error()))
		return
	}

	customEvent, err := models.NewEvent(newEvent.UserID, newEvent.Action, newEvent.ProductID, newEvent.Timestamp)
	if err != nil {
		SendError(rw, http.StatusBadRequest, errors.New("Can`t create new event: "+err.Error()))
		return
	}

	if err := storage.SaveEvent(customEvent); err != nil {
		SendError(rw, http.StatusInternalServerError, errors.New("Can`t save event: "+err.Error()))
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

func eventListHandler(rw http.ResponseWriter, r *http.Request) {

	tempStorage, err := storage.GetAllEvents()
	if err != nil {
		SendError(rw, http.StatusInternalServerError, errors.New("Can`t get data: "+err.Error()))
		return
	}

	tempStorageBytes, err := json.Marshal(tempStorage)

	if err != nil {
		SendError(rw, http.StatusInternalServerError, errors.New("Can`t create json with data: "+err.Error()))
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	if _, err := rw.Write(tempStorageBytes); err != nil {
		fmt.Println("Can`t write http answer:", err.Error())
	}
}

func userEventsHistoryHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["user_id"]

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		SendError(rw, http.StatusBadRequest, errors.New("Incorrect user id: "+err.Error()))
		return
	}

	userEvents, err := storage.GetEventsByUserId(userIdInt)
	if err != nil {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusNoContent)
		return
	}

	userEventsBytes, err := json.Marshal(userEvents)

	if err != nil {
		SendError(rw, http.StatusInternalServerError, errors.New("Can`t marshal data: "+err.Error()))
		return

	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	if _, err := rw.Write(userEventsBytes); err != nil {
		fmt.Println("Can`t write http answer:", err.Error())
	}
}

func main() {

	storage = repository.NewEventSliceRepo([]models.Event{})

	fs := http.FileServer(http.Dir("C:/GO/Reload/static"))

	router := mux.NewRouter()
	router.Handle("/", fs)
	router.HandleFunc("/event", eventHandler).Methods("POST")
	router.HandleFunc("/eventlist", eventListHandler).Methods("GET")
	router.HandleFunc("/user/{user_id}/history", userEventsHistoryHandler).Methods("GET")

	fmt.Println("Server is listening on localhost:8080")

	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		fmt.Println(err.Error())
		return
	}

}
