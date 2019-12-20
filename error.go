package main

// Error is a wrapper that contains a Base error and additional fields
type Error struct {
	Code int
	Status int
	Message string
	Base error
}