package main

import (
	"fmt"
	"github.com/marcoshuck/todogo/database"
	"github.com/marcoshuck/todogo/errors"
	"github.com/marcoshuck/todogo/router"
	"github.com/marcoshuck/todogo/server"
)

func main() {
	fmt.Println("TO-DO API Rest")
	var err *errors.Error

	fmt.Println("Setting database up.")
	db, err := database.New()
	if err != nil {
		fmt.Println(fmt.Sprintf("[Error] Status: %d | Code: %d | Message: %s", err.Status, err.Code, err.Message))
		fmt.Println(fmt.Sprintf("Stack trace: %v", err.Base))
	}
	defer database.Close(db)

	fmt.Println("Setting router up.")
	r := router.New()

	fmt.Println("Setting http server up.")
	srv := server.New(r, "127.0.0.1", "8080")
	err = server.Listen(srv)
	if err != nil {
		fmt.Println(fmt.Sprintf("[Error] Status: %d | Code: %d | Message: %s", err.Status, err.Code, err.Message))
		fmt.Println(fmt.Sprintf("Stack trace: %v", err.Base))
	}

	fmt.Println("Listening and serving http server")
	return
}