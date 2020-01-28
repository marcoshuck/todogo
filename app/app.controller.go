package app

import (
	def "github.com/marcoshuck/todogo/definitions"
	"net/http"
)

type ApplicationController def.Controller

func (app *ApplicationController) Create(w http.ResponseWriter, r *http.Request) {
	app.Service.Count()
	w.WriteHeader(http.StatusOK)
}

func (app *ApplicationController) Read(w http.ResponseWriter, r *http.Request) {
	app.Service.Count()
	w.WriteHeader(http.StatusOK)
}

func (app *ApplicationController) ReadAll(w http.ResponseWriter, r *http.Request) {
	app.Service.Count()
	w.WriteHeader(http.StatusOK)
}

func (app *ApplicationController) Count(w http.ResponseWriter, r *http.Request) {
	app.Service.Count()
	w.WriteHeader(http.StatusOK)
}

func (app *ApplicationController) Update(w http.ResponseWriter, r *http.Request) {
	app.Service.Count()
	w.WriteHeader(http.StatusOK)
}

func (app *ApplicationController) Delete(w http.ResponseWriter, r *http.Request) {
	app.Service.Count()
	w.WriteHeader(http.StatusOK)
}