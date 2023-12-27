package router

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	redis2 "github.com/go-redis/redis"

	"github.com/ptflp/godecoder"
	"go.uber.org/zap"
	"main/internal/infrastructure/cache"
	controller3 "main/internal/infrastructure/clients/geo/grpc"

	"main/internal/infrastructure/middleware"
	"main/internal/infrastructure/responder"
	"main/internal/modules/auth/controller"
	"main/internal/modules/auth/service"
	controller2 "main/internal/modules/geo/controller"
	service2 "main/internal/modules/geo/service"
	"main/internal/modules/geo/storage"
	controller4 "main/internal/modules/user/controller"
	service3 "main/internal/modules/user/service"
	storage2 "main/internal/modules/user/storage"
	"net/http"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	logs, _ := zap.NewProduction()
	resp := responder.NewResponder(godecoder.NewDecoder(), logs)
	client := redis2.NewClient(&redis2.Options{
		Addr: "redis:6379",
	})

	ping := client.Ping()
	fmt.Println(ping.String())
	redis := cache.NewRedis(client)

	geoProviderController := controller3.NewGeoProviderController(rpcClient)

	geoStorage := storage.NewGeoStorage(sqlAdapter, redis)
	userStorage := storage2.NewUserStorage(sqlAdapter)
	authService := service.NewAuth()

	geoService := service2.NewGeoService(geoStorage, geoProviderController)
	userService := service3.NewUserService(userStorage)

	reverseProxy := middleware.NewReverseProxy("hugo", "1313")

	r.Use(reverseProxy.ReverseProxy)

	r.Route("/api/address", func(r chi.Router) {
		geo := controller2.NewGeoController(geoService, resp)
		r.Post("/geocode", geo.Geo)
		r.Post("/search", geo.Search)
	})

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
	})

	r.Get("/swagger", middleware.SwaggerUI)
	r.Get("/static/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/static/", http.FileServer(http.Dir("/static"))).ServeHTTP(w, r)
	})
	return r
}
