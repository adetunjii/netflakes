package port

import (
	"context"

	"github.com/adetunjii/netflakes/model"
)

type MovieApi interface {
	FetchMovies(ctx context.Context) ([]model.Movie, error)
	FetchMovie(ctx context.Context, movieID int64) (*model.Movie, error)
	FetchMovieCharacters(ctx context.Context, id int64) ([]model.Character, error)
	GetCharacter(ctx context.Context, movieUrl string) (*model.Character, error)
}
