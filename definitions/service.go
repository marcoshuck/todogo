package definitions

import "github.com/marcoshuck/todogo/errors"

type Service struct {
	serviceName string
}

// ServiceMethods represents a set of generic REST service methods
type ServiceMethods interface {
	Add(entity interface{}) (interface{}, *errors.Error)
	FindOne(uuid string) (interface{}, *errors.Error)
	FindAll() ([]interface{}, *errors.Error)
	Count() int
	Update(uuid string, entity interface{}) (interface{}, *errors.Error)
	Delete(uuid string) (interface{}, *errors.Error)
}

