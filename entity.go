package main

import "time"

type Entity struct {
	ID int64
	Uuid string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
