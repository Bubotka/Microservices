package docks

import "main/internal/modules/user/controller"

//swagger:route Post /api/create создание_пользователя CreateRequest
// Responses:
//   200: CreateResponse

//swagger:parameters CreateRequest
type CreateRequest struct {
	// User - пользователя для создания
	// in: body
	// required: true
	// example: {"username": "Kolia","password": "2332"}
	User string
}

// swagger:response CreateResponse
type CreateResponse struct {
	// in: body
	// Message ответ в виде сообщения.
	Message string
}

// swagger:route POST /api/update обновление_пользователя UpdateRequest
// Responses:
//   200: UpdateResponse

// swagger:parameters UpdateRequest
type UpdateRequest struct {
	// User - пользователя для создания
	// in: body
	// required: true
	// example: {"id":1,"username": "Kolia","password": "2332"}
	User string
}

// swagger:response UpdateResponse
type UpdateResponse struct {
	// in: body
	// Message ответ в виде сообщения.
	Message string
}

// swagger:route POST /api/delete удаление_пользователя DeleteRequest
// Responses:
//   200: DeleteResponse

// swagger:parameters DeleteRequest
type DeleteRequest struct {
	// ID -индекс пользователя для удаления
	// in: body
	// required: true
	// example: {"id":1}
	ID string
}

// swagger:response DeleteResponse
type DeleteResponse struct {
	// in: body
	// Message ответ в виде сообщения.
	Message string
}

// swagger:route GET /api/list список_пользователя ListRequest
// Responses:
//   200: ListResponseDock

// swagger:parameters ListRequest
type ListRequest struct {
	// Limit - количество элементов для отображения
	// in: query
	// required: true
	// example: 10
	Limit int
	// Offset - индекс, с которого начинать выборку
	// in: query
	// required: true
	// example: 0
	Offset int
}

// swagger:response ListResponseDock
type ListResponseDock struct {
	// in: body
	// Response ответ содержащий пользователей и их количество в бд.
	Response controller.ListResponse
}

// swagger:route GET /api/user/{id} пользователь_по_id GetByIdRequest
// Responses:
//   200: GetByIdResponse

// swagger:parameters GetByIdRequest
type GetByIdRequest struct {
	// ID - идентификатор пользователя
	// in: path
	// required: true
	// type: integer
	// format: int64
	// example: 1
	ID int `json:"id"`
}

// swagger:response GetByIdResponse
type GetByIdResponse struct {
	// in: body
	// Response ответ содержащий пользователя.
	Response controller.User
}
