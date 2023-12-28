package client_adapter

import (
	"context"
	gp "github.com/Bubotka/Microservices/geo/pkg/go/geo"
	"github.com/Bubotka/Microservices/proxy/internal/models"
	"google.golang.org/grpc"
	"log"
	"time"
)

type ClientGRpcAdapter struct {
	client gp.GeoProviderClient
}

func NewClientGRpcAdapter(address string) *ClientGRpcAdapter {
	for i := 0; i < 5; i++ {
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Println("Ошибка при подключении к серверу:", err)
			time.Sleep(1 * time.Second)
			continue
		}
		client := gp.NewGeoProviderClient(conn)
		log.Println("Клиент подключился по адресу: ", address)

		return &ClientGRpcAdapter{client: client}
	}
	return nil
}

func (c *ClientGRpcAdapter) ListLevenshtein(column, text string) (models.SearchHistoryAddress, error) {
	req := &gp.ListLevenshteinRequest{
		Column: column,
		Text:   text,
	}

	response, err := c.client.ListLevenshtein(context.Background(), req)
	if err != nil {
		return models.SearchHistoryAddress{}, err
	}

	sha := models.SearchHistoryAddress{
		Id:              int(response.Sha.Id),
		SearchRequest:   response.Sha.SearchRequest,
		AddressResponse: response.Sha.AddressResponse,
	}

	return sha, nil
}

func (c *ClientGRpcAdapter) Create(sha models.SearchHistoryAddress) error {
	req := &gp.CreateRequest{Sha: &gp.SearchHistoryAddress{
		Id:              int32(sha.Id),
		SearchRequest:   sha.SearchRequest,
		AddressResponse: sha.AddressResponse,
	}}

	_, err := c.client.Create(context.Background(), req)
	return err
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
