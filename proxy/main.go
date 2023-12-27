package main

import (
	_ "github.com/lib/pq"
	"main/internal/infrastructure/router"
	"main/internal/infrastructure/server"
	"net/http"
)

func main() {

	s := &http.Server{
		Addr:    ":8080",
		Handler: router.NewRouter(),
	}

	serv := server.NewServer(s)

	serv.Serve()
}
