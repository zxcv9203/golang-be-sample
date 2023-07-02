package config

import (
	"github.com/zxcv9203/golang-be-sample/internal/repository"
	"github.com/zxcv9203/golang-be-sample/internal/service"
	"github.com/zxcv9203/golang-be-sample/internal/transport/rest"
)

func InitDependencies() *rest.Handler {
	r := repository.NewMemoryRepository()
	s := service.NewService(r)
	h := rest.NewHandler(s)

	return h
}
