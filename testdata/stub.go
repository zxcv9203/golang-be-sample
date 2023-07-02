package testdata

import (
	"github.com/zxcv9203/golang-be-sample/internal/model"
	"github.com/zxcv9203/golang-be-sample/internal/transport/rest/request"
)

func Post() model.Post {
	return model.Post{
		Id:      1,
		Title:   "test",
		Content: "test-content",
	}
}

func PostCreateRequest() *request.Post {
	return &request.Post{
		Title:   "test",
		Content: "test-content",
	}
}
