package testdata

import "github.com/zxcv9203/fiber-post/internal/model"

func Post() model.Post {
	return model.Post{
		Id:      1,
		Title:   "test",
		Content: "test-content",
	}
}
