package repository

import (
	"github.com/zxcv9203/fiber-post/internal/model"
	"github.com/zxcv9203/fiber-post/pkg/page"
)

type Repository interface {
	FindById(id int64) (model.Post, error)
	Save(post model.Post) (model.Post, error)
	Update(id int64, post model.Post) (model.Post, error)
	FindAll(request page.Request) ([]model.Post, error)
	DeleteById(id int64) error
}
