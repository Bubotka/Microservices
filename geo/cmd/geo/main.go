package main

import (
	"fmt"
	"github.com/Bubotka/Microservices/geo/domain/models"
	"github.com/Bubotka/Microservices/geo/internal/grpc"
	"github.com/Bubotka/Microservices/geo/pkg/db"
	"github.com/Bubotka/Microservices/geo/pkg/db/tools/Initializer"
	"github.com/Bubotka/Microservices/geo/pkg/db/tools/migrator"

	_ "github.com/lib/pq"
	"os"
)

func main() {
	postgresDB, err := db.NewPostgresDB(db.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	})

	if err != nil {
		fmt.Println("Не получилось подключиться к бд")
	}

	initializer := Initializer.NewInitializer(postgresDB)
	initializer.Init()

	var generator migrator.SQLiteGenerator
	m := migrator.NewMigrator(postgresDB, &generator)
	err = m.Migrate(&models.SearchHistoryAddress{})
	if err != nil {
		fmt.Println("Не удалось мигрировать")
	}

	grpcServer := grpc.NewGrpcServer()
	go grpcServer.Listen(":8082")
}