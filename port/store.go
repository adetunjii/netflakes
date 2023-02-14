package port

import (
	"context"

	"github.com/adetunjii/netflakes/model"
)

type SqlStore interface {
	Movie() MovieStore
}

type MovieStore interface {
	SaveComment(ctx context.Context, comment *model.Comment) error
	FetchComments(ctx context.Context, movieID int64, page int64, size int64) ([]model.Comment, error)
	FetchCommentCounts(ctx context.Context, movieID int64) (int64, error)
	GetComment(ctx context.Context, commentID int64) (*model.Comment, error)
	DeleteComment(ctx context.Context, commentID int64) error
}

type KVStore interface {
	CloseConnection() error
	GetMovies(ctx context.Context) ([]model.Movie, error)
	GetMovie(ctx context.Context, url string) (*model.Movie, error)
	SetMovies(ctx context.Context, movies []model.Movie) error
	SetMovie(ctx context.Context, url string, movie *model.Movie) error
	GetMovieCharacters(ctx context.Context, movieID int64) ([]model.Character, error)
	SetMovieCharacters(ctx context.Context, movieID int64, characters []model.Character) error
}
