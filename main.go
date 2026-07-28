package main

import (
	"log"
	"net/http"
)

func main() {
	if err := setApi(); err != nil {
		log.Fatal(err)
	}
	todoList := createTask()

	http.HandleFunc("POST /load", todoList.loadTask)
	http.HandleFunc("POST /add", todoList.addTaskUI)
	http.HandleFunc("DELETE /remove", todoList.removeTaskUI)
	http.HandleFunc("PUT /update", todoList.updateTaskUI)
	http.HandleFunc("POST /clear", todoList.clearSelection)

	log.Println("Server running on port http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
