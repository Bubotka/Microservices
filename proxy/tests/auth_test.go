package tests

import (
	"github.com/go-chi/chi/v5"
	"github.com/ptflp/godecoder"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"main/internal/infrastructure/responder"
	"main/internal/modules/auth/controller"
	"main/internal/modules/auth/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_registerHandler(t *testing.T) {
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
				url:    "/api/register",
				body:   "{\"username\":\"Kolia\",\"password\":\"123\"}",
				want:   "\"Пользователь успешно зарегестрирован\"\n",
			},
		},
	}
	r := chi.NewRouter()
	logs, _ := zap.NewProduction()
	resp := responder.NewResponder(godecoder.NewDecoder(), logs)
	authService := service.NewAuth()
	auth := controller.NewAuth(authService, resp)
	r.Route("/api", func(r chi.Router) {
		r.Post("/register", auth.Register)
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(tt.args.method, tt.args.url, strings.NewReader(tt.args.body))
			r.ServeHTTP(rec, req)
			assert.Equal(t, tt.args.want, rec.Body.String())
			assert.Equal(t, http.StatusOK, rec.Code)
		})
	}
}

func Test_loginHandler(t *testing.T) {
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
				url:    "/api/login",
				body:   "{\"username\":\"Kolia\",\"password\":\"123\"}",
				want:   "\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IktvbGlhIn0.UXrAYgIUuwyq8iKZa7t_QB6j9oGd_UQJjX44yCW80AU\"\n",
			},
		},
	}
	r := chi.NewRouter()
	logs, _ := zap.NewProduction()
	resp := responder.NewResponder(godecoder.NewDecoder(), logs)
	authService := service.NewAuth()
	auth := controller.NewAuth(authService, resp)
	r.Route("/api", func(r chi.Router) {
		r.Post("/login", auth.Login)
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(tt.args.method, tt.args.url, strings.NewReader(tt.args.body))
			r.ServeHTTP(rec, req)
			assert.Equal(t, tt.args.want, rec.Body.String())
			assert.Equal(t, http.StatusOK, rec.Code)
		})
	}
}
