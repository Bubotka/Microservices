package service

import (
	"context"
	"fmt"
	"main/internal/infrastructure/clients/geo/grpc"
	"main/internal/modules/geo/storage"
)

type GeoService struct {
	storage     storage.GeoRepository
	geoProvider grpc.GeoProvider
}

func NewGeoService(storage storage.GeoRepository, geoProvider grpc.GeoProvider) *GeoService {
	return &GeoService{storage: storage, geoProvider: geoProvider}
}

func (g *GeoService) ListLevenshtein(in ListlIn) ListlOut {
	out, err := g.storage.ListLevenshtein(context.Background(), in.Column, in.Text)
	return ListlOut{out, err}
}

func (g *GeoService) Create(in CreateIn) {
	g.storage.Create(context.Background(), in.SHA)
}

func (g *GeoService) Search(in SearchIn) SearchOut {
	fmt.Println("Зашли в Search")
	addressReworked, err := g.geoProvider.AddressSearch(in.Data)
	if err != nil {
		return SearchOut{}
	}
	return SearchOut{
		Addresses: addressReworked,
		Error:     nil,
	}
}

func (g *GeoService) Geo(in GeoIn) GeoOut {
	addressReworked, err := g.geoProvider.GeoCode(in.Data)
	if err != nil {
		return GeoOut{}
	}
	return GeoOut{
		Addresses: addressReworked,
		Error:     nil,
	}
}
