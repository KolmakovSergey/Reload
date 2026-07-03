package main

import (
	"fmt"
	"reload/internal/repository"
	"reload/server"
)

func main() {

	// storage := repository.NewEventSliceRepo([]models.Event{})

	dsn := "postgres://postgres:root@127.0.0.1:5432/Reload?sslmode=disable"

	var (
		storage *repository.PostgresRepo
		err     error
	)

	if storage, err = repository.NewPostgresRepo(dsn); err != nil {
		fmt.Println("connect to db failed: ", err.Error())
		return
	}

	eventHandlers := server.NewEventHandlers(storage)
	newServer := server.NewServer(eventHandlers)

	if err := newServer.StartServer(); err != nil {
		fmt.Println("failed to run server: " + err.Error())
	}
}
