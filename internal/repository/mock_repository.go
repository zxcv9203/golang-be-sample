package repository

import (
	"github.com/zxcv9203/fiber-post/internal/model"
	"github.com/zxcv9203/fiber-post/pkg/page"
	"github.com/zxcv9203/fiber-post/testdata"
)

type MockRepository struct {
}

func (m *MockRepository) Save(post model.Post) (model.Post, error) {
	return testdata.Post(), nil
}

func (m *MockRepository) Update(id int64, post model.Post) (model.Post, error) {
	return post, nil
}

func (m *MockRepository) FindById(id int64) (model.Post, error) {
	return testdata.Post(), nil
}

func (m *MockRepository) FindAll(request page.Request) ([]model.Post, error) {
	return []model.Post{testdata.Post()}, nil
}

func (m *MockRepository) DeleteById(id int64) error {
	return nil
}
