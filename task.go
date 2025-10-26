package main

import (
	"fmt"
)

type Todolist struct {
	Title string   `json:"title"`
	Tasks []string `json:"tasks"`
}

func createTask() *Todolist {
	return &Todolist{Title: "", Tasks: []string{}}
}

func (l *Todolist) addTask(task string) {
	l.Tasks = append(l.Tasks, task)

	err := fileAppend(l.Title, len(l.Tasks), task)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (l *Todolist) removeTask(position int) {
	l.Tasks = append(l.Tasks[:position], l.Tasks[position+1:]...)

	err := removeFileLine(l.Title, position)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (l *Todolist) updateTask(position int, newTask string) {
	l.Tasks[position] = newTask

	err := updateFileLine(l.Title, position, newTask)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (l *Todolist) clearTask() {
	l.Title = ""
	l.Tasks = []string{}
}
