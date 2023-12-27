package client_adapter

import "main/internal/models"

type ClientAdapter interface {
	Search(input []byte) (models.AddressSearchReworked, error)
	GeoCode(input []byte) (models.AddressSearchReworked, error)
}
