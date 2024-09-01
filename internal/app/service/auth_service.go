package service

import "martpedia-backend/internal/app/model/web"

type AuthService interface {
	Register(request web.UserRegisterRequest) string
}
