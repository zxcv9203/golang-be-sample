package repository

import (
	"fmt"
	"github.com/zxcv9203/golang-be-sample/internal/model"
	"github.com/zxcv9203/golang-be-sample/internal/transport/rest/request"
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

func (m *MemoryRepository) FindAll(request request.Page) []model.Post {
	storeSize := int64(len(m.store))
	requestSize := request.Size()
	startIdx := request.Page()

	if requestSize > storeSize {
		requestSize = storeSize
	}

	posts := make([]model.Post, 0, requestSize)

	for i := int64(0); i < requestSize; i++ {
		idx := i + startIdx
		if m.notExistsPost(idx) {
			continue
		}
		posts = append(posts, m.store[idx])
	}

	return posts
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
