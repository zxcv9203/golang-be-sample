package repository

import (
	"github.com/zxcv9203/golang-be-sample/internal/model"
	"github.com/zxcv9203/golang-be-sample/pkg/page"
)

type MemoryRepository struct {
	store map[int64]model.Post
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		store: make(map[int64]model.Post),
	}
}

func (m *MemoryRepository) Save(post model.Post) (model.Post, error) {
	size := int64(len(m.store) + 1)
	post.Id = size
	m.store[size] = post
	return post, nil
}

func (m *MemoryRepository) Update(id int64, post model.Post) (model.Post, error) {
	return post, nil
}

func (m *MemoryRepository) FindById(id int64) (model.Post, error) {
	return model.Post{}, nil
}

func (m *MemoryRepository) FindAll(request page.Request) ([]model.Post, error) {
	return []model.Post{}, nil
}

func (m *MemoryRepository) DeleteById(id int64) error {
	return nil

}
