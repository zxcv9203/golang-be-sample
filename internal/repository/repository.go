package repository

import (
	"github.com/zxcv9203/golang-be-sample/internal/model"
	"github.com/zxcv9203/golang-be-sample/internal/transport/rest/request"
)

type Repository interface {
	FindById(id int64) (model.Post, error)
	Save(post model.Post) (model.Post, error)
	Update(id int64, post model.Post) (model.Post, error)
	FindAll(request request.Page) []model.Post
	DeleteById(id int64) error
}
