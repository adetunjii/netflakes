package sqlstore

import (
	"context"

	"github.com/adetunjii/netflakes/model"
)

const DEFAULT_PAGE_SIZE = 20

type SqlMovieStore struct {
	*SqlStore
}

func newMovieStore(sqlstore *SqlStore) *SqlMovieStore {
	return &SqlMovieStore{sqlstore}
}

func (m *SqlMovieStore) SaveComment(ctx context.Context, comment *model.Comment) error {
	return m.db.Save(comment)
}

func (m *SqlMovieStore) FetchComments(ctx context.Context, movieID int64, page int64, size int64) ([]model.Comment, error) {

	if page == 0 {
		page = 1
	}

	if size == 0 {
		size = DEFAULT_PAGE_SIZE
	}

	limit := size
	offset := (page - 1) * size
	dest := []model.Comment{}
	conditions := map[string]interface{}{
		"movie_id": movieID,
	}

	err := m.db.List(&dest, conditions, int(limit), int(offset))
	return dest, err
}

func (m *SqlMovieStore) FetchCommentCounts(ctx context.Context, movieID int64) (int64, error) {
	model := &model.Comment{}
	condition := map[string]interface{}{
		"movie_id": movieID,
	}
	return m.db.Count(model, condition)
}

func (m *SqlMovieStore) GetComment(ctx context.Context, commentID int64) (*model.Comment, error) {
	dest := &model.Comment{}
	if err := m.db.FindById(dest, commentID); err != nil {
		return nil, err
	}
	return dest, nil

}

func (m *SqlMovieStore) DeleteComment(ctx context.Context, commentID int64) error {
	c := &model.Comment{}
	condition := map[string]interface{}{
		"id": commentID,
	}
	return m.db.DeleteOne(c, condition)
}
