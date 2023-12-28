package router

import (
	"fmt"
	"github.com/Bubotka/Microservices/proxy/internal/infrastructure/clients/geo/grpc/client_adapter"
	"github.com/go-chi/chi/v5"
	redis2 "github.com/go-redis/redis"

	controller3 "github.com/Bubotka/Microservices/proxy/internal/infrastructure/clients/geo/grpc"
	"github.com/ptflp/godecoder"
	"go.uber.org/zap"

	"github.com/Bubotka/Microservices/proxy/internal/infrastructure/middleware"
	"github.com/Bubotka/Microservices/proxy/internal/infrastructure/responder"
	controller2 "github.com/Bubotka/Microservices/proxy/internal/modules/geo/controller"
	service2 "github.com/Bubotka/Microservices/proxy/internal/modules/geo/service"
	"net/http"
)

func NewRouter(geoClientAdapter client_adapter.ClientAdapter) http.Handler {
	r := chi.NewRouter()
	logs, _ := zap.NewProduction()
	resp := responder.NewResponder(godecoder.NewDecoder(), logs)
	client := redis2.NewClient(&redis2.Options{
		Addr: "redis:6379",
	})

	ping := client.Ping()
	fmt.Println(ping.String())

	geoProviderController := controller3.NewGeoProviderController(geoClientAdapter)

	geoService := service2.NewGeoService(geoProviderController)

	reverseProxy := middleware.NewReverseProxy("hugo", "1313")

	r.Use(reverseProxy.ReverseProxy)

	r.Route("/api/address", func(r chi.Router) {
		geo := controller2.NewGeoController(geoService, resp)
		r.Post("/geocode", geo.Geo)
		r.Post("/search", geo.Search)
	})
	/*
		r.Route("/api", func(r chi.Router) {
			auth := controller.NewAuth(authService, resp)
			user := controller4.NewUser(userService, resp)
			r.Post("/login", auth.Login)
			r.Post("/register", auth.Register)
			r.Post("/create", user.Create)
			r.Post("/update", user.Update)
			r.Post("/delete", user.Delete)

			r.Get("/user/{id}", user.GetByUsername)

			r.Get("/list", user.List)
		})*/

	r.Get("/swagger", middleware.SwaggerUI)
	r.Get("/static/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/static/", http.FileServer(http.Dir("/static"))).ServeHTTP(w, r)
	})
	return r
}
