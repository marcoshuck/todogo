package app

import def "github.com/marcoshuck/todogo/definitions"

var appService def.ServiceMethods = &ApplicationService{
	ServiceName: "ApplicationService",
}

var appController def.ControllerMethods = &ApplicationController{
	ControllerName: "ApplicationController",
	Service:        &appService,
}

var ApplicationModule = def.Module{
	Route:      &def.Route{
		Name:        "Application",
		Description: "Applications endpoint",
		Endpoint:    "/app",
	},
	Entity:     nil,
	Controller: &appController,
	Service:    &appService,
}