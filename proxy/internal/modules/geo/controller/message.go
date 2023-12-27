package controller

import "main/internal/models"

type SearchResponse struct {
	Addresses models.AddressSearchReworked `json:"addresses"`
}

type GeocodeResponse struct {
	Addresses models.AddressSearchReworked `json:"addresses"`
}

type SearchRequest struct {
	Place string `json:"query"`
}
