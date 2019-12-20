package main

type Service interface {
	Add()
	FindOne()
	FindAll()
	Count()
	Update()
	Delete()
}