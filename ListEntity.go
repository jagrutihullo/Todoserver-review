package main

import "time"

//list entity
type List struct {
	ID        uint
	Name      string
	CreatedAt time.Time `json:"timestamp,omitempty"`
	UpdatedAt time.Time `json:"utimestamp,omitempty"`
	Tasks     []Task
}
