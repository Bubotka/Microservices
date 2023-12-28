package grpc

import (
	"github.com/Bubotka/Microservices/proxy/internal/infrastructure/clients/geo/grpc/client_adapter"
	"github.com/Bubotka/Microservices/proxy/internal/models"
)

type GeoProvider interface {
	ListLevenshtein(column, text string) (models.SearchHistoryAddress, error)
	Create(sha models.SearchHistoryAddress) error
	AddressSearch(input []byte) (models.AddressSearchReworked, error)
	GeoCode(input []byte) (models.AddressSearchReworked, error)
}

type GeoProviderController struct {
	client client_adapter.ClientAdapter
}

func (g *GeoProviderController) ListLevenshtein(column, text string) (models.SearchHistoryAddress, error) {
	response, err := g.client.ListLevenshtein(column, text)
	if err != nil {
		return models.SearchHistoryAddress{}, err
	}
	return response, nil
}

func (g *GeoProviderController) Create(sha models.SearchHistoryAddress) error {
	err := g.client.Create(sha)
	return err
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
