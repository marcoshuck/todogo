package definitions

import "github.com/marcoshuck/todogo/errors"

// Service represents a generic REST service
type Service interface {
	Add(entity interface{}) (interface{}, *errors.Error)
	FindOne(uuid string) (interface{}, *errors.Error)
	FindAll() ([]interface{}, *errors.Error)
	Count() int
	Update(uuid string, entity interface{}) (interface{}, *errors.Error)
	Delete(uuid string) (interface{}, *errors.Error)
}