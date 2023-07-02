package service

import (
	"github.com/zxcv9203/fiber-post/internal/repository"
	"github.com/zxcv9203/fiber-post/testdata"
	"testing"
)

func NewService() Service {
	return Service{
		repo: &repository.MockRepository{},
	}
}

func TestSave(t *testing.T) {
	service := NewService()

	t.Run("[성공] 게시글 저장", func(t *testing.T) {
		request := testdata.PostCreateRequest()
		want := testdata.Post().Id

		got := service.Save(request)

		assertEquals(t, got, want)
	})
}

func assertEquals(t *testing.T, got int64, want int64) {
	t.Helper()

	if got != want {
		t.Errorf("비교하는 데이터가 다릅니다. got : %q want %q", got, want)
	}
}
