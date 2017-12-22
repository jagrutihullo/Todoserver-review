package main

import "time"

type TodoListEntity struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TaskEntity struct {
	ID          uint
	Description string
	Status      string
	Deadline    time.Time
}
