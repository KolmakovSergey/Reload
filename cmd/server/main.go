package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reload/internal/models"
	"reload/internal/repository"
	"strconv"

	"github.com/gorilla/mux"
)

var storage repository.EventRepo

func eventHandler(rw http.ResponseWriter, r *http.Request) {

	event, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("can`t read event: ", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	var newEvent models.EventDTO

	if err := json.Unmarshal(event, &newEvent); err != nil {
		fmt.Println("can`t unmarshal event dto: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	customEvent, err := models.NewEvent(newEvent.UserID, newEvent.Action, newEvent.ProductID, newEvent.Timestamp)

	if err != nil {
		fmt.Println("can`t create new event: ", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	storage.SaveEvent(customEvent)
	fmt.Println("Event save!")

	// Устанавливаем заголовок Content-Type
	rw.Header().Set("Content-Type", "application/json")
	// Отправляем JSON-объект
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`{"status":"ok","event":"saved"}`))
}

func eventListHandler(rw http.ResponseWriter, r *http.Request) {

	StorageBytes, err := json.Marshal(storage.Storage)

	if err != nil {
		fmt.Println("can`t create json with event data: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(StorageBytes)
}

func userEventsHistoryHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["user_id"]
	userIdInt, err := strconv.Atoi(userId)

	if err != nil {
		fmt.Println("can`t convert user id: ", err)
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	userEvents := storage.GetEventsByUserId(userIdInt)

	userEventsBytes, err := json.Marshal(userEvents)

	if err != nil {
		fmt.Println("can`t marshal events: ", err)
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(userEventsBytes)
}

func main() {

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
