package main

import (
	"fmt"
	"github.com/marcoshuck/todogo/router"
	"github.com/marcoshuck/todogo/server"
)

func main() {
	fmt.Println("TO-DO API Rest")
	r := router.New()
	srv := server.New(r, "127.0.0.1", "8080")
	srv.ListenAndServe()
}