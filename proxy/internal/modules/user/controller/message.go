package controller

import "main/internal/models"

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateRequest struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type DeleteRequest struct {
	Id int `json:"id"`
}

type ListResponse struct {
	Users []models.User `json:"users"`
	Count int           `json:"count"`
}
