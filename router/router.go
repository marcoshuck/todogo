package router

import (
	"github.com/gorilla/mux"
	def "github.com/marcoshuck/todogo/definitions"
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

func (r *Router) Register(mod *def.Module) {
	r.HandleFunc(mod.Route.Endpoint, (*mod.Controller).Create).Methods("POST").Name(mod.Route.Name + " - Create")
	r.HandleFunc(mod.Route.Endpoint, (*mod.Controller).ReadAll).Methods("GET").Name(mod.Route.Name + " - GetAll")
	r.HandleFunc(mod.Route.Endpoint + "{slug}", (*mod.Controller).Read).Methods("POST").Name(mod.Route.Name + " - Get")
	r.HandleFunc(mod.Route.Endpoint + "{slug}", (*mod.Controller).Update).Methods("PATCH").Name(mod.Route.Name + " - Update")
	r.HandleFunc(mod.Route.Endpoint + "{slug}", (*mod.Controller).Delete).Methods("DELETE").Name(mod.Route.Name + " - Delete")
}