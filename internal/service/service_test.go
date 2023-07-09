package service

import (
	"github.com/zxcv9203/golang-be-sample/internal/repository"
	"github.com/zxcv9203/golang-be-sample/testdata"
	"testing"
)

func NewMockService() Service {
	return Service{
		repo: &repository.MockRepository{},
	}
}

func TestSave(t *testing.T) {
	service := NewMockService()

	t.Run("[성공] 게시글 저장", func(t *testing.T) {
		request := testdata.PostCreateRequest()
		want := testdata.Post().Id

		got := service.Save(request)

		assertEquals(t, got, want)
	})
}

func TestUpdate(t *testing.T) {
	service := NewMockService()

	t.Run("[성공] 게시글 업데이트", func(t *testing.T) {
		request := testdata.PostCreateRequest()
		want := testdata.Post().Id

		got, _ := service.Update(want, request)

		assertEquals(t, got, want)
	})

	t.Run("[실패] 존재하지 않는 ID 입력시 예외처리", func(t *testing.T) {
		request := testdata.PostCreateRequest()
		wrongId := int64(0)

		_, err := service.Update(wrongId, request)

		assertError(t, err)
	})
}

func TestFindById(t *testing.T) {
	service := NewMockService()

	t.Run("[성공] 게시글 조회", func(t *testing.T) {
		want := testdata.Post()
		id := want.Id

		got, err := service.FindById(id)
		if err != nil {
			t.Errorf("에러 발생 : %q", err)
		}

		assertEquals(t, got, want)
	})

	t.Run("[실패] 존재하지 않는 게시글 ID를 전달하는 경우 에러 반환", func(t *testing.T) {
		wrongId := int64(0)

		_, err := service.FindById(wrongId)

		assertError(t, err)
	})
}

func TestDeleteById(t *testing.T) {
	service := NewMockService()

	t.Run("[성공] 게시글 단일 삭제", func(t *testing.T) {
		want := testdata.Post().Id

		err := service.DeleteById(want)
		if err != nil {
			t.Errorf("에러 발생 : %q", err)
		}
	})

	t.Run("[실패] 존재하지 않는 게시글 ID를 전달하는 경우 에러 반환", func(t *testing.T) {
		wrongId := int64(0)

		err := service.DeleteById(wrongId)

		assertError(t, err)
	})
}

func assertEquals(t *testing.T, got any, want any) {
	t.Helper()

	if got != want {
		t.Errorf("비교하는 데이터가 다릅니다. got : %q want %q", got, want)
	}
}

func assertError(t *testing.T, err error) {
	t.Helper()

	if err == nil {
		t.Errorf("에러가 발생하지 않았습니다.")
	}
}
