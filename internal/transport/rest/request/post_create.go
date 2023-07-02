package request

import "github.com/zxcv9203/fiber-post/internal/model"

type Post struct {
	Title   string
	Content string
}

func (p *Post) ToEntity() model.Post {
	return model.Post{
		Title:   p.Title,
		Content: p.Content,
	}
}
