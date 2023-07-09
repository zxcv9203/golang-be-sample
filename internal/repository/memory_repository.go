package repository

import (
	"fmt"
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
	if m.notExistsPost(id) {
		return model.Post{}, fmt.Errorf("게시글이 존재하지 않습니다. 게시글 ID : %d", id)
	}

	post.Id = id
	m.store[id] = post
	return post, nil
}

func (m *MemoryRepository) FindById(id int64) (model.Post, error) {
	if m.notExistsPost(id) {
		return model.Post{}, fmt.Errorf("게시글이 존재하지 않습니다. 게시글 ID : %d", id)
	}
	return m.store[id], nil
}

func (m *MemoryRepository) FindAll(request page.Request) ([]model.Post, error) {

	return []model.Post{}, nil
}

func (m *MemoryRepository) DeleteById(id int64) error {
	if m.notExistsPost(id) {
		return fmt.Errorf("게시글이 존재하지 않습니다. 게시글 ID : %d", id)
	}
	delete(m.store, id)
	return nil
}

func (m *MemoryRepository) notExistsPost(id int64) bool {
	_, ok := m.store[id]
	return !ok
}
