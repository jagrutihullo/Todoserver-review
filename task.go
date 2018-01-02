package main

import "time"

//Task is an entity under List, used throughout the system.
type Task struct {
	ID          uint
	Description string
	Status      string
	Deadline    time.Time
}
