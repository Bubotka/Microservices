package router

import (
	geogrpc "github.com/Bubotka/Microservices/proxy/internal/infrastructure/clients/geo"
	clientadaptergeo "github.com/Bubotka/Microservices/proxy/internal/infrastructure/clients/geo/grpc/client_adapter"
	ucontroller "github.com/Bubotka/Microservices/proxy/internal/modules/user/controller"
	uservice "github.com/Bubotka/Microservices/proxy/internal/modules/user/service"
	usergrpc "github.com/Bubotka/Microservices/proxy/pkg/clients/user/grpc"
	clientadapteruser "github.com/Bubotka/Microservices/proxy/pkg/clients/user/grpc/client_adapter"

	"github.com/go-chi/chi/v5"
	"github.com/ptflp/godecoder"
	"go.uber.org/zap"

	"github.com/Bubotka/Microservices/proxy/internal/infrastructure/middleware"
	"github.com/Bubotka/Microservices/proxy/internal/infrastructure/responder"
	gcontroller "github.com/Bubotka/Microservices/proxy/internal/modules/geo/controller"
	gservice "github.com/Bubotka/Microservices/proxy/internal/modules/geo/service"
	"net/http"
)

func NewRouter(geoClientAdapter clientadaptergeo.GeoClientAdapter, userClientAdapter clientadapteruser.UserClientAdapter) http.Handler {
	r := chi.NewRouter()
	logs, _ := zap.NewProduction()
	resp := responder.NewResponder(godecoder.NewDecoder(), logs)

	geoProvider := geogrpc.NewGeoProvider(geoClientAdapter)
	userProvider := usergrpc.NewUserProvider(userClientAdapter)

	geoService := gservice.NewGeoService(geoProvider)
	userService := uservice.NewUserService(userProvider)

	geoController := gcontroller.NewGeoController(geoService, resp)
	userController := ucontroller.NewUserController(userService, resp)

	reverseProxy := middleware.NewReverseProxy("hugo", "1313")

	r.Use(reverseProxy.ReverseProxy)

	r.Route("/api/address", func(r chi.Router) {

		r.Post("/geocode", geoController.Geo)
		r.Post("/search", geoController.Search)
	})

	r.Route("/api/user", func(r chi.Router) {
		r.Get("/profile/{username}", userController.GetByUsername)
		r.Get("/list", userController.List)
	})

	r.Get("/swagger", middleware.SwaggerUI)
	r.Get("/static/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/static/", http.FileServer(http.Dir("/static"))).ServeHTTP(w, r)
	})
	return r
}
