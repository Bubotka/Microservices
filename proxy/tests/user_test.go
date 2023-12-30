package tests

import (
	"context"
	"geo/internal/db/adapter"
	"geo/internal/db/adapter/mocks"
	"github.com/go-chi/chi/v5"
	"github.com/ptflp/godecoder"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"main/internal/infrastructure/responder"
	"main/internal/models"
	"main/internal/modules/user/controller"
	service3 "main/internal/modules/user/service"
	"main/internal/modules/user/storage"
	mocks2 "main/internal/modules/user/storage/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateUser(t *testing.T) {
	type args struct {
		method string
		url    string
		body   string
		want   string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "base_test",
			args: args{
				method: "POST",
				url:    "/api/create",
				body:   "{\"username\": \"Kolia\",\"password\": \"2332\"}",
				want:   "\"Пользователь успешно создан\"\n",
			},
		},
	}
	r := chi.NewRouter()
	logs, _ := zap.NewProduction()
	resp := responder.NewResponder(godecoder.NewDecoder(), logs)
	mockAdapter := mocks.NewSqlAdapterer(t)
	userStorage := storage.NewUserStorage(mockAdapter)
	userService := service3.NewUserService(userStorage)
	r.Route("/api", func(r chi.Router) {
		user := controller.NewUser(userService, resp)
		r.Post("/create", user.Create)
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			mockAdapter.On("Create", context.Background(), models.User{Username: "Kolia", Password: "2332"}).Return(nil)
			req, _ := http.NewRequest(tt.args.method, tt.args.url, strings.NewReader(tt.args.body))
			r.ServeHTTP(rec, req)

			assert.Equal(t, tt.args.want, rec.Body.String())
		})
	}
}

func TestUpdateUser(t *testing.T) {
	type args struct {
		method string
		url    string
		body   string
		want   string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "base_test",
			args: args{
				method: "POST",
				url:    "/api/update",
				body:   "{\"id\":1,\"username\": \"Kolia\",\"password\": \"2332\"}",
				want:   "\"Пользователь успешно обновлён\"\n",
			},
		},
	}
	r := chi.NewRouter()
	logs, _ := zap.NewProduction()
	resp := responder.NewResponder(godecoder.NewDecoder(), logs)
	mockAdapter := mocks.NewSqlAdapterer(t)
	userStorage := storage.NewUserStorage(mockAdapter)
	userService := service3.NewUserService(userStorage)
	r.Route("/api", func(r chi.Router) {
		user := controller.NewUser(userService, resp)
		r.Post("/update", user.Update)
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			mockAdapter.On("Update", context.Background(),
				models.User{ID: 1, Username: "Kolia", Password: "2332"},
				adapter.Condition{Equal: map[string]interface{}{"id": 1}}, map[string]interface{}{"id": 1},
				map[string]interface{}{"username": "Kolia"},
				map[string]interface{}{"password": "2332"}).
				Return(nil)

			req, _ := http.NewRequest(tt.args.method, tt.args.url, strings.NewReader(tt.args.body))
			r.ServeHTTP(rec, req)

			assert.Equal(t, tt.args.want, rec.Body.String())
		})
	}
}

func TestDeleteUser(t *testing.T) {
	type args struct {
		method string
		url    string
		body   string
		want   string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "base_test",
			args: args{
				method: "POST",
				url:    "/api/delete",
				body:   "{\"id\":1}",
				want:   "\"Пользователь успешно удалён\"\n",
			},
		},
	}
	r := chi.NewRouter()
	logs, _ := zap.NewProduction()
	resp := responder.NewResponder(godecoder.NewDecoder(), logs)
	mockAdapter := mocks.NewSqlAdapterer(t)
	userStorage := storage.NewUserStorage(mockAdapter)
	userService := service3.NewUserService(userStorage)
	r.Route("/api", func(r chi.Router) {
		user := controller.NewUser(userService, resp)
		r.Post("/delete", user.Delete)
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			mockAdapter.On("Update", context.Background(),
				models.User{Username: "", Password: ""},
				adapter.Condition{Equal: map[string]interface{}{"id": 1}},
				map[string]interface{}{"is_delete": true}).
				Return(nil)
			req, _ := http.NewRequest(tt.args.method, tt.args.url, strings.NewReader(tt.args.body))
			r.ServeHTTP(rec, req)

			assert.Equal(t, tt.args.want, rec.Body.String())
		})
	}
}

func TestListUser(t *testing.T) {
	type args struct {
		method string
		url    string
		want   string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "base_test",
			args: args{
				method: "GET",
				url:    "/api/list",
				want:   "{\"users\":[{\"ID\":1,\"Username\":\"fgs\",\"Password\":\"123\",\"IsDelete\":false}],\"count\":1}\n",
			},
		},
	}
	r := chi.NewRouter()
	logs, _ := zap.NewProduction()
	resp := responder.NewResponder(godecoder.NewDecoder(), logs)
	mockUserRep := mocks2.NewUserRepository(t)
	userService := service3.NewUserService(mockUserRep)
	r.Route("/api", func(r chi.Router) {
		user := controller.NewUser(userService, resp)
		r.Get("/list", user.List)
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			mockUserRep.On("List", context.Background(),
				adapter.Condition{LimitOffset: &adapter.LimitOffset{Limit: 100, Offset: 0}}).
				Return([]models.User{{ID: 1, Username: "fgs", Password: "123", IsDelete: false}}, nil)

			mockUserRep.On("Count", context.Background(), adapter.Condition{}).
				Return(1, nil)

			req, _ := http.NewRequest(tt.args.method, tt.args.url, nil)
			r.ServeHTTP(rec, req)
			assert.Equal(t, tt.args.want, rec.Body.String())
		})
	}
}

/*
func TestGetByIDUser(t *testing.T) {
	type args struct {
		method string
		url    string
		want   string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "base_test",
			args: args{
				method: "GET",
				url:    "/api/user/1",
				want:   "{\"ID\":1,\"Username\":\"fgs\",\"Password\":\"123\",\"IsDelete\":false}\n",
			},
		},
	}
	r := chi.NewRouter()
	logs, _ := zap.NewProduction()
	resp := responder.NewResponder(godecoder.NewDecoder(), logs)
	mockUserRep := mocks2.NewUserRepository(t)
	userService := service3.NewUserService(mockUserRep)
	r.Route("/api", func(r chi.Router) {

		r.Get("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
			id := chi.URLParam(r, "id")
			out := userService.GetByUsername(service3.GetIn{Id: id})
			if out.Error != nil {
				resp.ErrorBadRequest(w, out.Error)
				return
			}

			w.WriteHeader(http.StatusOK)
			resp.OutputJSON(w, out.User)
		})
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			mockUserRep.On("GetByUsername", context.Background(), "1").
				Return(models.User{ID: 1, Username: "fgs", Password: "123", IsDelete: false}, nil)

			req, _ := http.NewRequest(tt.args.method, tt.args.url, nil)
			r.ServeHTTP(rec, req)
			assert.Equal(t, tt.args.want, rec.Body.String())
		})
	}
}
*/
