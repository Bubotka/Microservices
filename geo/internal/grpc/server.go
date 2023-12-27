package grpc

import (
	"github.com/Bubotka/Microservices/geo/internal/services/geo"
	gp "github.com/Bubotka/Microservices/geo/pkg/go/geo"
	"google.golang.org/grpc"
	"log"

	"net"
)

type GrpcServer struct {
}

func NewGrpcServer() *GrpcServer {
	return &GrpcServer{}
}

func (g *GrpcServer) Listen(address string) {
	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Ошибка при прослушивании порта: %v", err)
	}

	server := grpc.NewServer()
	gp.RegisterGeoProviderServer(server, geo.NewGeoService())

	log.Println("Запуск gRPC сервера...")
	if err := server.Serve(l); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
