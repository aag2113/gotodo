package main

import (
	"net/http"

	"gotodo/db"
	"gotodo/server"
)

func main() {
	db.Load()
	defer db.Kill()

	http.HandleFunc("/tasks", server.TasksRouter)
	http.HandleFunc("/tasks/", server.TaskRouter)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
