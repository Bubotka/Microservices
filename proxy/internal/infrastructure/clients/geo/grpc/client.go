package grpc

import (
	"main/internal/infrastructure/clients/geo/grpc/client_adapter"
	"main/internal/models"
)

type GeoProvider interface {
	AddressSearch(input []byte) (models.AddressSearchReworked, error)
	GeoCode(input []byte) (models.AddressSearchReworked, error)
}

type GeoProviderController struct {
	client client_adapter.ClientAdapter
}

func NewGeoProviderController(client client_adapter.ClientAdapter) *GeoProviderController {
	return &GeoProviderController{client: client}
}

func (g *GeoProviderController) AddressSearch(input []byte) (models.AddressSearchReworked, error) {
	addressSearchReworked, err := g.client.Search(input)
	return addressSearchReworked, err
}

func (g *GeoProviderController) GeoCode(input []byte) (models.AddressSearchReworked, error) {
	addressSearchReworked, err := g.client.GeoCode(input)
	return addressSearchReworked, err
}
