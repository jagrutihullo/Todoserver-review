package main

import "time"

//task entity
type Task struct {
	ID          uint
	Description string
	Status      string
	Deadline    time.Time `json:"timestamp,omitempty"`
}
