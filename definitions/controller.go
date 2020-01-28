package definitions

import (
	"net/http"
)

type Controller struct {
	ControllerName string
	Service *ServiceMethods
}

// ControllerMethods represents a set of generic REST controller methods
type ControllerMethods interface {
	Create(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
	ReadAll(w http.ResponseWriter, r *http.Request)
	Count(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}