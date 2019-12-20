package main

// Service represents a generic REST service
type Service interface {
	Add(entity interface{}) (interface{}, Error)
	FindOne(uuid string) (interface{}, Error)
	FindAll() ([]interface{}, Error)
	Count() int
	Update(uuid string, entity interface{}) (interface{}, Error)
	Delete(uuid string) (interface{}, Error)
}