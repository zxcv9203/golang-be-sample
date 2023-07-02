package repository

import (
	"github.com/zxcv9203/fiber-post/internal/model"
	"github.com/zxcv9203/fiber-post/pkg/page"
)

type Repository interface {
	findById(id int64) (model.Post, error)
	save(post model.Post) (model.Post, error)
	update(id int64, post model.Post) (model.Post, error)
	findAll(request page.Request) ([]model.Post, error)
	deleteById(id int64) error
}
