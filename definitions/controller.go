package definitions

import "github.com/marcoshuck/todogo/errors"

type Controller struct {
	controllerName string
}

// ControllerMethods represents a set of generic REST controller methods
type ControllerMethods interface {
	Create(entity interface{}) (interface{}, *errors.Error)
	Read(uuid string) (interface{}, *errors.Error)
	ReadAll() ([]interface{}, *errors.Error)
	Count() int
	Update(uuid string, entity interface{}) (interface{}, *errors.Error)
	Delete(uuid string) (interface{}, *errors.Error)
}