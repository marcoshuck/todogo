package main

import (
	"fmt"
	"github.com/marcoshuck/todogo/database"
	"github.com/marcoshuck/todogo/router"
	"github.com/marcoshuck/todogo/server"
)

func main() {
	fmt.Println("TO-DO API Rest")

	fmt.Println("Setting database up.")
	db, err := database.New()
	defer database.Close(db)

	fmt.Println("Setting router up.")
	r := router.New()

	fmt.Println("Setting http server up.")
	srv := server.New(r, "127.0.0.1", "8080")
	srv.ListenAndServe()

	fmt.Println("Listening and serving http server")
	return
}