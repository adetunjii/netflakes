package swapi

import (
	"context"
	"testing"

	"github.com/adetunjii/netflakes/pkg/logging"
	"github.com/adetunjii/netflakes/utils/testutils"
	"github.com/stretchr/testify/require"
)

func createSwapiInstance(t *testing.T) *Swapi {
	baseURL := "https://swapi.dev/api"
	zapSugarLogger := logging.NewZapSugarLogger()
	logger := logging.NewLogger(zapSugarLogger)

	return New(baseURL, logger)
}

func TestSwapi_Get(t *testing.T) {

	testSwapi := createSwapiInstance(t)

	rootPath := "/"

	res, err := testSwapi.Get(rootPath)
	require.NoError(t, err)
	require.NotNil(t, res)

	// test for 404 errors
	errorPath := "/not-exist"
	res, err = testSwapi.Get(errorPath)
	require.Error(t, err)
	require.Nil(t, res)

}

func TestSwapi_FetchMovie(t *testing.T) {

	testSwapi := createSwapiInstance(t)

	testMovieID := int64(1)
	movie, err := testSwapi.FetchMovie(context.Background(), testMovieID)
	require.NoError(t, err)
	require.NotEmpty(t, movie)
	require.Equal(t, testutils.TestMovieSample.Title, movie.Title)
	require.Equal(t, testutils.TestMovieSample.EpisodeID, movie.EpisodeID)
	require.Equal(t, testutils.TestMovieSample.Created, movie.Created)
	require.Equal(t, len(testutils.TestMovieSample.Characters), len(movie.Characters))

}

func TestSwapi_FetchMovies(t *testing.T) {

	testSwapi := createSwapiInstance(t)

	movies, err := testSwapi.FetchMovies(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, movies)
	require.GreaterOrEqual(t, len(movies), 5)

	// assumes the movie at the top of the list has an id of 1
	movie := movies[0]
	require.Equal(t, testutils.TestMovieSample.EpisodeID, movie.EpisodeID)
	require.Equal(t, testutils.TestMovieSample.OpeningCrawl, movie.OpeningCrawl)
	require.Equal(t, testutils.TestMovieSample.Created, movie.Created)
	require.Equal(t, testutils.TestMovieSample.Characters, movie.Characters)
	require.Equal(t, len(testutils.TestMovieSample.Characters), len(movie.Characters))
}
