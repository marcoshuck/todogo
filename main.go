package main

import (
	"fmt"
	"github.com/marcoshuck/todogo/database"
	"github.com/marcoshuck/todogo/errors"
	"github.com/marcoshuck/todogo/logger"
	"github.com/marcoshuck/todogo/router"
	"github.com/marcoshuck/todogo/server"
)

func main() {
	var err *errors.Error
	var log *logger.Logger
	log = logger.GetInstance()

	log.Info("TO-DO API Rest")

	log.Info("Setting database up.")
	db, err := database.New()
	if err != nil {
		log.Error(*err)
	}
	defer database.Close(db)

	log.Info("Setting router up.")
	r := router.New()

	log.Info("Setting http server up.")

	srv := server.New(r, "127.0.0.1", "8080")

	err = server.Listen(srv)
	if err != nil {
		log.Error(*err)
	}

	fmt.Println("Listening and serving http server")
	return
}