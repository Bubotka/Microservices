package client_adapter

import (
	"context"
	gp "geo/pkg/go/geo"
	"main/internal/models"
)

type ClientGRpcAdapter struct {
	client gp.GeoProviderClient
}

func NewClientGRpcAdapter(client gp.GeoProviderClient) *ClientGRpcAdapter {
	return &ClientGRpcAdapter{client: client}
}

func (c *ClientGRpcAdapter) Search(input []byte) (models.AddressSearchReworked, error) {
	req := &gp.SearchRequest{Place: string(input)}
	response, err := c.client.Search(context.Background(), req)
	if err != nil {
		return nil, err
	}
	var elements []models.AddressElementSearch
	for _, element := range response.Elements {
		elementSearch := models.AddressElementSearch{
			Result: element.GetResult(),
			GeoLat: element.GetGeoLat(),
			GeoLon: element.GetGeoLon(),
		}
		elements = append(elements, elementSearch)
	}
	return elements, nil
}

func (c *ClientGRpcAdapter) GeoCode(input []byte) (models.AddressSearchReworked, error) {
	req := &gp.GeoRequest{Coordinates: string(input)}
	response, err := c.client.GeoCode(context.Background(), req)
	if err != nil {
		return nil, err
	}
	var elements []models.AddressElementSearch
	for _, element := range response.Elements {
		elementSearch := models.AddressElementSearch{
			Result: element.GetResult(),
			GeoLat: element.GetGeoLat(),
			GeoLon: element.GetGeoLon(),
		}
		elements = append(elements, elementSearch)
	}
	return elements, nil
}
