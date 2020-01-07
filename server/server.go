package server

import (
	"fmt"
	"github.com/marcoshuck/todogo/router"
	"net/http"
	"time"
)

func New(router router.Router, address string, port string) *http.Server {
	server := &http.Server{
		Addr:              fmt.Sprintf("%s:%s", address, port),
		Handler:           router,
		// TLSConfig:         nil,
		ReadTimeout:       time.Second * 15,
		// ReadHeaderTimeout: 0,
		WriteTimeout:      time.Second * 15,
		IdleTimeout:       time.Second * 60,
		// MaxHeaderBytes:    0,
		// TLSNextProto:      nil,
		// ConnState:         nil,
		// ErrorLog:          nil,
		// BaseContext:       nil,
		// ConnContext:       nil,
	}
	return server
}