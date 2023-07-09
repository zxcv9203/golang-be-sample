package repository

import (
	"fmt"
	"github.com/zxcv9203/golang-be-sample/internal/model"
	"github.com/zxcv9203/golang-be-sample/pkg/page"
	"github.com/zxcv9203/golang-be-sample/testdata"
)

type MockRepository struct {
}

func (m *MockRepository) Save(post model.Post) (model.Post, error) {
	return testdata.Post(), nil
}

func (m *MockRepository) Update(id int64, post model.Post) (model.Post, error) {
	if m.validatePostId(id) {
		return model.Post{}, fmt.Errorf("게시글이 존재하지 않습니다. 게시글 ID : %d", id)
	}
	post.Id = id
	return post, nil
}

func (m *MockRepository) FindById(id int64) (model.Post, error) {
	if m.validatePostId(id) {
		return model.Post{}, fmt.Errorf("게시글이 존재하지 않습니다. 게시글 ID : %d", id)
	}
	return testdata.Post(), nil
}

func (m *MockRepository) FindAll(request page.Request) ([]model.Post, error) {
	return []model.Post{testdata.Post()}, nil
}

func (m *MockRepository) DeleteById(id int64) error {
	if m.validatePostId(id) {
		return fmt.Errorf("게시글이 존재하지 않습니다. 게시글 ID : %d", id)
	}
	return nil
}

// id 0인 게시글이 들어올 경우 예외
func (m *MockRepository) validatePostId(id int64) bool {
	wrongId := int64(0)
	if id == wrongId {
		return true
	}
	return false
}
