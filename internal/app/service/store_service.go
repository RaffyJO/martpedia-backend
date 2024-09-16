package service

import "martpedia-backend/internal/app/model/web"

type StoreService interface {
	Create(request web.StoreCreateRequest) (web.StoreCreateResponse, error)
}
