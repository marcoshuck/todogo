package router

import "github.com/gorilla/mux"

type Router struct {
	*mux.Router
}

func New() *Router {
	router := mux.NewRouter()
	return &Router{
		Router: router,
	}
}