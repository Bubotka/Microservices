package main

import (
	"github.com/Bubotka/Microservices/proxy/internal/infrastructure/clients/geo/grpc/client_adapter"
	"github.com/Bubotka/Microservices/proxy/internal/infrastructure/router"
	"github.com/Bubotka/Microservices/proxy/internal/infrastructure/server"
	_ "github.com/lib/pq"

	"net/http"
)

func main() {
	client := client_adapter.NewClientGRpcAdapter(":8081")

	s := &http.Server{
		Addr:    ":8080",
		Handler: router.NewRouter(client),
	}

	serv := server.NewServer(s)

	serv.Serve()
}
