package main

type Controller interface {
	Create() (interface{}, Error)
	Read() (interface{}, Error)
	ReadAll() ([]interface{}, Error)
	Update() (interface{}, Error)
	Delete() (interface{}, Error)
}