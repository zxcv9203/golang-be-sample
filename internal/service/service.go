package service

import (
	"github.com/zxcv9203/golang-be-sample/internal/model"
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
		log.Println(err)
	}

	return save.Id
}

func (s *Service) Update(id int64, request *request.Post) (int64, error) {
	post := request.ToEntity()
	update, err := s.repo.Update(id, post)
	if err != nil {
		return 0, err
	}

	return update.Id, nil
}

func (s *Service) FindById(id int64) (model.Post, error) {
	post, err := s.repo.FindById(id)
	if err != nil {
		return model.Post{}, err
	}
	return post, nil
}

func (s *Service) DeleteById(id int64) error {
	err := s.repo.DeleteById(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Find(request request.Page) []model.Post {
	return s.repo.FindAll(request)
}
