package router

import (
	"github.com/gorilla/mux"
)

// Router represents an application router.
type Router struct {
	*mux.Router
}

// New returns a new Router instance.
func New() *Router {
	router := mux.NewRouter()
	return &Router{
		Router: router,
	}
}