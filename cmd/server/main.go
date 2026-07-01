package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reload/internal/models"

	"github.com/gorilla/mux"
)


var mainSlice []models.Event

func eventHandler(rw http.ResponseWriter, r *http.Request) {

	fmt.Println("Event exists!")

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
		fmt.Println("can`t create new event Purchase: ", err)
		rw.WriteHeader(http.StatusBadRequest)
	}

	mainSlice = append(mainSlice, customEvent)
	fmt.Println("Event save!")

	// Устанавливаем заголовок Content-Type
    rw.Header().Set("Content-Type", "application/json")
    // Отправляем JSON-объект
    rw.WriteHeader(http.StatusOK)
    rw.Write([]byte(`{"status":"ok","event":"saved"}`))
}

func main() {

	fs := http.FileServer(http.Dir("C:/GO/Reload/static"))

	router := mux.NewRouter()
	router.Handle("/", fs)
	router.HandleFunc("/event", eventHandler).Methods("POST")

	fmt.Println("Server is listening on localhost:8080")

	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		fmt.Println(err.Error())
		return
	}

}
