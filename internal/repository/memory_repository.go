package repository

import (
	"github.com/zxcv9203/fiber-post/internal/model"
	"github.com/zxcv9203/fiber-post/pkg/page"
)

type MemoryRepository struct {
	store map[int64]model.Post
}

func (m *MemoryRepository) save(post model.Post) (model.Post, error) {
	size := int64(len(m.store) + 1)
	post.Id = size
	m.store[size] = post
	return post, nil
}

func (m *MemoryRepository) update(id int64, post model.Post) (model.Post, error) {
	return post, nil
}

func (m *MemoryRepository) findById(id int64) (model.Post, error) {
	return model.Post{}, nil
}

func (m *MemoryRepository) findAll(request page.Request) ([]model.Post, error) {
	return []model.Post{}, nil
}

func (m *MemoryRepository) deleteById(id int64) error {
	return nil

}
