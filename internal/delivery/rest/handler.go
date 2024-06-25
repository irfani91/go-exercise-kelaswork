package rest

import "main.go/internal/usecase/resto"

type handler struct {
	restoUsecase resto.UseCase
}

func NewHandler(restoUsecase resto.UseCase) *handler {
	return &handler{
		restoUsecase: restoUsecase,
	}
}
