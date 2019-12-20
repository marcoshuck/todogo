package main

type Controller interface {
	Create(entity interface{}) (interface{}, Error)
	Read(uuid string) (interface{}, Error)
	ReadAll() ([]interface{}, Error)
	Count() int
	Update(uuid string, entity interface{}) (interface{}, Error)
	Delete(uuid string) (interface{}, Error)
}