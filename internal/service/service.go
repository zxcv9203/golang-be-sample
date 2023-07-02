package service

import (
	"github.com/zxcv9203/fiber-post/internal/repository"
	"github.com/zxcv9203/fiber-post/internal/transport/rest/request"
	"log"
)

type Service struct {
	repo repository.Repository
}

func (s *Service) Save(request request.Post) int64 {
	post := request.ToEntity()
	save, err := s.repo.Save(post)
	if err != nil {
		log.Print(err)
	}

	return save.Id
}
