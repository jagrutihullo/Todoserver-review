package main

import "time"

//list entity
type List struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Tasks     []Task
}
