package main

type Error struct {
	Code int
	Status int
	Message string
	Base error
}