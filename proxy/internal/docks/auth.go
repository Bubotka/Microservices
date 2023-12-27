package docks

//swagger:route Post /api/register регистрация registerRequest
// Регистрация пользователя
// Responses:
//   200: registerResponse

//swagger:parameters registerRequest
type registerRequest struct {
	// Userdata - данные пользователя
	// in: body
	// required: true
	// example: {"username":"Kolia","password":"123"}
	Userdata string
}

// swagger:response registerResponse
type registerResponse struct {
	// in: body
	// Result резултат операции.
	// example: {"Register is succeed"}
	Result string
}

//swagger:route Post /api/login авторизация loginRequest
// Авторизация пользователя
// Responses:
//   200: loginResponse

//swagger:parameters loginRequest
type loginRequest struct {
	// Userdata - данные пользователя
	// in: body
	// required: true
	// example: {"username":"Kolia","password":"123"}
	Userdata string
}

// swagger:response loginResponse
type loginResponse struct {
	// in: body
	// Token токен пользователя.
	// example: {"23owiudfghjklasdfgsdw"}
	Token string
}
