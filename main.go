package main

import (
	"github.com/adetunjii/netflakes/api"
	"github.com/adetunjii/netflakes/config"
	"github.com/adetunjii/netflakes/db"
	"github.com/adetunjii/netflakes/pkg/logging"
	"github.com/adetunjii/netflakes/port"
	"github.com/adetunjii/netflakes/services/redis"
	"github.com/adetunjii/netflakes/services/swapi"
	"github.com/adetunjii/netflakes/store/sqlstore"
)

func main() {
	baseURL := "https://swapi.dev/api"
	zapSugarLogger := logging.NewZapSugarLogger()
	logger := logging.NewLogger(zapSugarLogger)

	dbConfig, redisConfig := loadConfig(logger)

	sw := swapi.New(baseURL, logger)
	redis := redis.New(redisConfig, logger)
	dbInstance := db.New(dbConfig, logger)
	sqlStore := sqlstore.New(dbInstance, logger)

	server := api.NewServer(redis, sqlStore, sw, logger)

	server.Start()
}

func loadConfig(logger port.Logger) (*db.Config, *redis.Config) {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		logger.Fatal("failed to load config: %w", err)
	}

	dbConfig := &db.Config{
		Host:        cfg.DbHost,
		Port:        cfg.DbPort,
		User:        cfg.DbUser,
		Password:    cfg.DbPassword,
		Name:        cfg.DbName,
		DatabaseUrl: cfg.DbUrl,
	}

	redisConfig := &redis.Config{
		Host:     cfg.RedisHost,
		Port:     cfg.RedisPort,
		Password: cfg.RedisPassword,
		Db:       0,
		User:     cfg.RedisUser,
	}

	return dbConfig, redisConfig
}
