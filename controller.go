package main

type Controller interface {
	Create()
	Read()
	ReadAll()
	Update()
	Delete()
}