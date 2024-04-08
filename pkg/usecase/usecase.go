package usecase

import "api-architecture/pkg/service"

type Usecase struct {
	services *service.Service
}

func NewUsecase(services *service.Service) *Usecase {
	return &Usecase{services: services}
}
