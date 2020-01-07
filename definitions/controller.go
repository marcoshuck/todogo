package definitions

import "github.com/marcoshuck/todogo/errors"

// Controller represents a generic REST controller
type Controller interface {
	Create(entity interface{}) (interface{}, *errors.Error)
	Read(uuid string) (interface{}, *errors.Error)
	ReadAll() ([]interface{}, *errors.Error)
	Count() int
	Update(uuid string, entity interface{}) (interface{}, *errors.Error)
	Delete(uuid string) (interface{}, *errors.Error)
}