package main

import "time"

//List is an entity used throughout the system.
type List struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Tasks     []Task
}
