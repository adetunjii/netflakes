package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DbHost        string `mapstructure:"DB_HOST"`
	DbUser        string `mapstructure:"DB_USER"`
	DbPassword    string `mapstructure:"DB_PASSWORD"`
	DbPort        string `mapstructure:"DB_PORT"`
	DbUrl         string `mapstructure:"DB_URL"`
	DbName        string `mapstructure:"DB_NAME"`
	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     string `mapstructure:"REDIS_PORT"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	RedisUser     string `mapstructure:"REDIS_USER"`
}

func LoadConfig(path string) (config Config, err error) {
	// viper.AddConfigPath(path)
	// viper.SetConfigName("app")
	// viper.SetConfigType("env")

	// enables system environment variables take precedence over the ones
	// from environment variable files.
	viper.AutomaticEnv()

	// Access environment variables
	dbUser := viper.GetString("DB_USER")
	dbPassword := viper.GetString("DB_PASSWORD")
	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetString("DB_PORT")
	dbUrl := viper.GetString("DB_URL")
	dbName := viper.GetString("DB_NAME")
	redisHost := viper.GetString("REDIS_HOST")
	redisPort := viper.GetString("REDIS_PORT")
	redisPassword := viper.GetString("REDIS_PASSWORD")
	redisUser := viper.GetString("REDIS_USER")

	config = Config{
		DbHost:        dbHost,
		DbPort:        dbPort,
		DbUser:        dbUser,
		DbPassword:    dbPassword,
		DbUrl:         dbUrl,
		DbName:        dbName,
		RedisHost:     redisHost,
		RedisPassword: redisPassword,
		RedisUser:     redisUser,
		RedisPort:     redisPort,
	}

	return config, nil
}
