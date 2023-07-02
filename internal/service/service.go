package service

import (
	"github.com/zxcv9203/golang-be-sample/internal/repository"
	"github.com/zxcv9203/golang-be-sample/internal/transport/rest/request"
	"log"
)

type Service struct {
	repo repository.Repository
}

func NewService(r repository.Repository) Service {
	return Service{
		repo: r,
	}
}

func (s *Service) Save(request *request.Post) int64 {
	post := request.ToEntity()
	save, err := s.repo.Save(post)
	if err != nil {
		log.Print(err)
	}

	return save.Id
}
