package tests

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/ptflp/godecoder"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"main/internal/infrastructure/clients/geo/grpc"
	"main/internal/infrastructure/responder"
	"main/internal/models"
	"main/internal/modules/geo/service"
	controller2 "main/internal/modules/user/controller"
	mocks2 "main/internal/modules/user/storage/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSearchHandler(t *testing.T) {
	type args struct {
		method string
		url    string
		body   string
		want   string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "base_test",
			args: args{
				method: "POST",
				url:    "/api/address/search",
				body:   "{\"query\":\"москва сухонская 11\"}\n",
				want:   "{\"addresses\":[{\"result\":\"г Москва, ул Сухонская, д 11\",\"lat\":\"55.878\",\"lon\":\"37.654\"}]}\n",
			},
		},
	}
	r := chi.NewRouter()
	logs, _ := zap.NewProduction()
	resp := responder.NewResponder(godecoder.NewDecoder(), logs)
	geoRepository := mocks2.NewGeoRepository(t)
	provider := factory.GetProvider("grpc")

	go provider.Listen(":8081")
	client, _ := provider.Connect("localhost:8081")
	geoService := service.NewGeoService(geoRepository, grpc.NewGeoProviderController(client))
	r.Route("/api/address", func(r chi.Router) {
		geo := controller2.NewGeoController(geoService, resp)
		r.Post("/search", geo.Search)
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(tt.args.method, tt.args.url, strings.NewReader(tt.args.body))

			geoRepository.On("ListLevenshtein", context.Background(), "search_request", "москва сухонская 11").
				Return(models.SearchHistoryAddress{}, fmt.Errorf("no cache"))

			geoRepository.On("Create", context.Background(),
				models.SearchHistoryAddress{
					Id:              0,
					SearchRequest:   "москва сухонская 11",
					AddressResponse: "{\"addresses\":[{\"result\":\"г Москва, ул Сухонская, д 11\",\"lat\":\"55.878\",\"lon\":\"37.654\"}]}"}).
				Return(nil)

			r.ServeHTTP(rec, req)
			assert.Equal(t, tt.args.want, rec.Body.String())
		})
	}
}

func TestGeoHandler(t *testing.T) {
	type args struct {
		method string
		url    string
		body   string
		want   string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "base_test",
			args: args{
				method: "POST",
				url:    "/api/address/geocode",
				body:   "{\"lat\":\"55.746953903751084\",\"lng\":\"37.61718750000001\"}\n",
				want:   "{\"addresses\":[{\"result\":\"г Москва, Софийская наб, д 14 стр 1\",\"lat\":\"55.747\",\"lon\":\"37.617\"}]}\n",
			},
		},
	}
	r := chi.NewRouter()
	logs, _ := zap.NewProduction()
	resp := responder.NewResponder(godecoder.NewDecoder(), logs)
	geoRepository := mocks2.NewGeoRepository(t)
	provider := factory.GetProvider("grpc")

	go provider.Listen(":8082")
	client, _ := provider.Connect("localhost:8082")
	geoService := service.NewGeoService(geoRepository, grpc.NewGeoProviderController(client))
	r.Route("/api/address", func(r chi.Router) {
		geo := controller2.NewGeoController(geoService, resp)
		r.Post("/geocode", geo.Geo)
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(tt.args.method, tt.args.url, strings.NewReader(tt.args.body))
			r.ServeHTTP(rec, req)
			assert.Equal(t, tt.args.want, rec.Body.String())
		})
	}
}
