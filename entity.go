package main

import "time"

// Entity represents a generic entity to be used by Controllers and Services
type Entity struct {
	ID int64
	Uuid string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
