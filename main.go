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

	http.HandleFunc("/load", todoList.loadTask)
	http.HandleFunc("/add", todoList.addTaskUI)
	http.HandleFunc("/remove", todoList.removeTaskUI)
	http.HandleFunc("/update", todoList.updateTaskUI)
	http.HandleFunc("/clear", todoList.clearSelection)

	log.Println("Server running on port http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
