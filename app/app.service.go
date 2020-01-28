package app

import (
	def "github.com/marcoshuck/todogo/definitions"
	"github.com/marcoshuck/todogo/errors"
)

type ApplicationService def.Service

func (app *ApplicationService) Add(entity interface{}) (interface{}, *errors.Error) {
	return nil, nil
}

func (app *ApplicationService) FindOne(uuid string) (interface{}, *errors.Error) {
	return nil, nil
}

func (app *ApplicationService) FindAll() ([]interface{}, *errors.Error) {
	return nil, nil
}

func (app *ApplicationService) Count() int {
	return -1
}

func (app *ApplicationService) Update(uuid string, entity interface{}) (interface{}, *errors.Error) {
	return nil, nil

}

func (app *ApplicationService) Delete(uuid string) (interface{}, *errors.Error) {
	return nil, nil
}