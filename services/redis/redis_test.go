package redis

import (
	"context"
	"testing"
	"time"

	"github.com/adetunjii/netflakes/pkg/logging"
	"github.com/adetunjii/netflakes/utils/testutils"
	"github.com/stretchr/testify/require"
)

func createRedisInstance(t *testing.T) *Redis {

	zapSugarLogger := logging.NewZapSugarLogger()
	logger := logging.NewLogger(zapSugarLogger)
	require.NotEmpty(t, logger)

	redisConfig := &Config{
		Host:     "localhost",
		Port:     "6379",
		Password: "password",
		Db:       0,                // default db
		Expiry:   10 * time.Minute, // default expiry time
	}

	return New(redisConfig, logger)
}

func TestSetMovies(t *testing.T) {
	testRedisInstance := createRedisInstance(t)
	err := testRedisInstance.SetMovies(context.Background(), testutils.TestMovieArray)
	require.NoError(t, err)
}

func TestGetMovies(t *testing.T) {
	testRedisInstance := createRedisInstance(t)

	movies, err := testRedisInstance.GetMovies(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, movies)
	require.Len(t, movies, 1)
}

func TestCloseConnection(t *testing.T) {
	testRedisInstance := createRedisInstance(t)

	err := testRedisInstance.CloseConnection()
	require.NoError(t, err)
}
