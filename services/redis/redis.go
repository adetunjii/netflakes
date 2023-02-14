package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/adetunjii/netflakes/model"
	"github.com/adetunjii/netflakes/port"
	"github.com/redis/go-redis/v9"
)

const MOVIE_KEY = "movies"
const DEFAULT_EXPIRY = 12 * time.Hour

type Config struct {
	Host     string
	Port     string
	User     string
	Db       int
	Password string
	Expiry   time.Duration
}

type Redis struct {
	client *redis.Client
	expiry time.Duration
	logger port.Logger
}

var _ port.KVStore = (*Redis)(nil)

func New(config *Config, logger port.Logger) *Redis {
	redis := &Redis{
		client: nil,
		expiry: config.Expiry,
		logger: logger,
	}

	if err := redis.GetClient(config); err != nil {
		logger.Error("connection to redis failed", err)
	}
	return redis
}

func (r *Redis) GetClient(config *Config) error {
	redisAddr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	redisURL := &redis.Options{
		Addr:     redisAddr,
		Password: config.Password,
		DB:       config.Db,
		Username: config.User,
	}

	r.client = redis.NewClient(redisURL)

	_, err := r.client.Ping(context.Background()).Result()
	if err != nil {
		r.logger.Fatal("error connecting to redis %v", err)
	}

	r.logger.Info(fmt.Sprintf("Redis connected successfully on %s...", redisAddr))
	return nil
}

func (r *Redis) CloseConnection() error {
	if r.client != nil {
		return r.client.Close()
	}
	return nil
}

func (r *Redis) RestartConnection(config *Config) error {
	if r.client != nil {
		r.CloseConnection()
	}

	return r.GetClient(config)
}

func (r *Redis) GetMovies(ctx context.Context) ([]model.Movie, error) {
	movies := []model.Movie{}
	key := fmt.Sprintf("%s:movie_list", MOVIE_KEY)

	result, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if redisErr, ok := err.(redis.Error); ok {
			if redisErr == redis.Nil {
				return movies, nil
			}
			return nil, redisErr
		}

		r.logger.Error("failed to get query result: %w", err)
		return nil, err
	}

	if err = json.Unmarshal([]byte(result), &movies); err != nil {
		r.logger.Error("failed to unmarshal json data: %w", err)
		return nil, err
	}
	return movies, nil

}

func (r *Redis) SetMovies(ctx context.Context, movies []model.Movie) error {
	key := fmt.Sprintf("%s:movie_list", MOVIE_KEY)

	j, err := json.Marshal(movies)
	if err != nil {
		r.logger.Error("failed to marshal json data: %w", err)
		return err
	}

	return r.client.Set(ctx, key, string(j), DEFAULT_EXPIRY).Err()
}

func (r *Redis) GetMovie(ctx context.Context, url string) (*model.Movie, error) {
	key := fmt.Sprintf("%s:%s", MOVIE_KEY, url)

	movie := &model.Movie{}

	result, err := r.client.Get(ctx, key).Result()
	if err != nil {
		r.logger.Error("failed to get query result %w", err)
		return nil, err
	}

	if err = json.Unmarshal([]byte(result), &movie); err != nil {
		r.logger.Error("failed to unmarshal json data: %w", err)
		return nil, err
	}
	return movie, nil
}

func (r *Redis) SetMovie(ctx context.Context, url string, movie *model.Movie) error {
	key := fmt.Sprintf("%s:%s", MOVIE_KEY, url)

	j, err := json.Marshal(movie)
	if err != nil {
		r.logger.Error("failed to marshal json data: %w", err)
		return err
	}
	return r.client.Set(ctx, key, string(j), DEFAULT_EXPIRY).Err()
}

func (r *Redis) SetMovieCharacters(ctx context.Context, movieID int64, characters []model.Character) error {
	key := fmt.Sprintf("%s:%s:%d", MOVIE_KEY, "characters", movieID)

	j, err := json.Marshal(characters)
	if err != nil {
		r.logger.Error("failed to marshal json data: %w", err)
		return err
	}

	return r.client.Set(ctx, key, string(j), DEFAULT_EXPIRY).Err()
}

func (r *Redis) GetMovieCharacters(ctx context.Context, movieID int64) ([]model.Character, error) {
	key := fmt.Sprintf("%s:%s:%d", MOVIE_KEY, "characters", movieID)

	characters := []model.Character{}

	result, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if redisErr, ok := err.(redis.Error); ok {
			if redisErr == redis.Nil {
				return characters, nil
			}
			return nil, redisErr
		}
	}

	if err := json.Unmarshal([]byte(result), &characters); err != nil {
		r.logger.Error("failed to unmarshal json data: %w", err)
		return nil, err
	}
	return characters, nil
}
