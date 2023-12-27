package main

import (
	"github.com/Bubotka/Microservices/proxy/internal/infrastructure/router"
	"github.com/Bubotka/Microservices/proxy/internal/infrastructure/server"
	_ "github.com/lib/pq"

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
