package repository

import (
	"github.com/zxcv9203/golang-be-sample/internal/model"
	"github.com/zxcv9203/golang-be-sample/testdata"
	"reflect"
	"testing"
)

func NewRepository() Repository {
	return &MemoryRepository{
		store: make(map[int64]model.Post),
	}
}

func InitializeData(repo Repository) {
	posts := testdata.Posts()
	for _, post := range posts {
		_, err := repo.Save(post)
		if err != nil {
			return
		}
	}
}

func cleanData(repo Repository) {
	repo = &MemoryRepository{
		store: make(map[int64]model.Post),
	}
}

func TestSave(t *testing.T) {
	repo := NewRepository()

	teardown := func() {
		cleanData(repo)
	}

	t.Run("게시글 저장", func(t *testing.T) {
		want := testdata.Post()

		got, err := repo.Save(want)
		if err != nil {
			t.Errorf("데이터 저장에 실패했습니다. 실패 사유 : %q", err)
		}

		assertEqual(t, got, want)
		// 저장한 데이터 삭제
		teardown()
	})
}

func TestUpdate(t *testing.T) {
	repo := NewRepository()

	beforeEach := func() {
		InitializeData(repo)
	}

	teardown := func() {
		cleanData(repo)
	}

	t.Run("[성공] 게시글 업데이트", func(t *testing.T) {
		beforeEach()

		want := testdata.Post()
		id := want.Id

		got, _ := repo.Update(id, want)

		assertEqual(t, got, want)

		teardown()
	})

	t.Run("[실패] 게시글 ID가 존재하지 않는 경우 에러 반환", func(t *testing.T) {
		beforeEach()

		want := testdata.Post()
		id := int64(0)

		_, err := repo.Update(id, want)

		assertError(t, err)

		teardown()
	})
}

func TestFindById(t *testing.T) {
	repo := NewRepository()

	beforeEach := func() {
		InitializeData(repo)
	}

	teardown := func() {
		cleanData(repo)
	}

	t.Run("[성공] 게시글 단일 조회", func(t *testing.T) {
		beforeEach()

		want := testdata.Post()
		id := want.Id

		got, err := repo.FindById(id)
		if err != nil {
			t.Errorf("에러가 발생했습니다.\n에러 내용 : %q", err)
		}

		assertEqual(t, got, want)

		teardown()
	})

	t.Run("[실패] 존재하지 않는 게시글 ID를 전달하는 경우 에러 반환", func(t *testing.T) {
		beforeEach()

		wrongId := int64(0)

		_, err := repo.FindById(wrongId)

		assertError(t, err)

		teardown()
	})
}

func TestDeleteById(t *testing.T) {
	repo := NewRepository()

	beforeEach := func() {
		InitializeData(repo)
	}

	teardown := func() {
		cleanData(repo)
	}

	t.Run("[성공] 게시글 단일 삭제", func(t *testing.T) {
		beforeEach()

		want := testdata.Post()
		id := want.Id

		err := repo.DeleteById(id)
		if err != nil {
			t.Errorf("에러가 발생했습니다.\n에러 내용 : %q", err)
		}

		teardown()
	})

	t.Run("[실패] 존재하지 않는 게시글 ID를 전달하는 경우 에러 반환", func(t *testing.T) {
		beforeEach()

		wrongId := int64(0)

		err := repo.DeleteById(wrongId)

		assertError(t, err)

		teardown()
	})
}

func TestFind(t *testing.T) {
	repo := NewRepository()

	beforeEach := func() {
		InitializeData(repo)
	}

	teardown := func() {
		cleanData(repo)
	}

	t.Run("[성공] 게시글 조회", func(t *testing.T) {
		beforeEach()

		request := testdata.PageRequest()
		want := testdata.Posts()

		got := repo.FindAll(request)

		assertSliceEqual(t, got, want)

		teardown()
	})

}

func assertSliceEqual(t *testing.T, got any, want any) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("저장된 데이터가 요청한 데이터와 다릅니다. got : %q want : %q ", got, want)
	}
}

func assertEqual(t *testing.T, got any, want any) {
	t.Helper()

	if got != want {
		t.Errorf("저장된 데이터가 요청한 데이터와 다릅니다. got : %q want : %q ", got, want)
	}
}

func assertError(t *testing.T, err error) {
	t.Helper()

	if err == nil {
		t.Errorf("에러가 발생하지 않았습니다.")
	}
}
