package main

import (
	"github.com/marcoshuck/todogo/app"
	"github.com/marcoshuck/todogo/database"
	"github.com/marcoshuck/todogo/errors"
	log "github.com/marcoshuck/todogo/logger"
	"github.com/marcoshuck/todogo/router"
	"github.com/marcoshuck/todogo/server"
)

func main() {
	var err *errors.Error
	var logger *log.Logger
	logger = log.GetInstance()

	logger.Info("TO-DO API Rest")

	logger.Info("Setting database up.")
	db, err := database.New()
	if err != nil {
		msg := logger.FormatErrorMessage(*err)
		logger.Error(msg)
	}
	defer database.Close(db)

	logger.Info("Setting router up.")
	r := router.New()

	r.Register(&app.ApplicationModule)

	logger.Info("Setting http server up.")

	srv := server.New(r, "127.0.0.1", "8080")

	logger.Info("Listening TO-DO API Rest on 127.0.0.1:8080")
	err = server.Listen(srv)
	if err != nil {
		msg := logger.FormatErrorMessage(*err)
		logger.Error(msg)
	}
	return
}