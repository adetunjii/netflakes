package db

import (
	"fmt"

	"github.com/adetunjii/netflakes/port"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host        string `json:"host"`
	Port        string `json:"port"`
	User        string `json:"user"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	DatabaseUrl string `json:"url"`
}

type PostgresDB struct {
	instance *gorm.DB
	logger   port.Logger
}

var _ port.DB = (*PostgresDB)(nil)

type PreloadOption func(*gorm.DB)

func New(dbConfig *Config, logger port.Logger) *PostgresDB {
	db := &PostgresDB{
		instance: nil,
		logger:   logger,
	}

	if err := db.Connect(dbConfig); err != nil {
		logger.Fatal("connection to db failed: %v", err)
	}
	return db
}

func (p *PostgresDB) Connect(config *Config) error {

	var dsn string
	databaseUrl := config.DatabaseUrl

	if databaseUrl == "" {
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			config.Host, config.Port, config.User, config.Password, config.Name)
	} else {
		dsn = databaseUrl
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	p.instance = db

	p.logger.Info(fmt.Sprintf("Database Connected Successfully %v...", dsn))
	return nil
}

func (p *PostgresDB) CloseConnection() error {
	if p.instance != nil {
		connection, err := p.instance.DB()
		if err != nil {
			return err
		}

		connection.Close()
	}
	return nil
}

func (p *PostgresDB) RestartConnection(config *Config) error {
	if p.instance != nil {
		p.CloseConnection()
	}

	return p.Connect(config)
}

func (p *PostgresDB) Save(arg interface{}) error {
	return p.instance.Create(arg).Error
}

func (p *PostgresDB) List(dest interface{}, conditions map[string]interface{}, limit int, offset int) error {
	return p.instance.Limit(limit).Offset(offset).Where(conditions).Find(dest).Error

}

func (p *PostgresDB) FindOne(dest interface{}, conditions map[string]interface{}) error {
	return p.instance.Where(conditions).First(dest).Error

}

func (p *PostgresDB) FindById(dest interface{}, id int64) error {
	return p.instance.First(dest, id).Error
}

func (p *PostgresDB) DeleteOne(model interface{}, condition map[string]interface{}) error {
	return p.instance.Where(condition).Delete(model).Error
}

func (p *PostgresDB) Raw(dest interface{}, query string, values ...interface{}) error {
	return p.instance.Raw(query, values).Scan(dest).Error
}

func (p *PostgresDB) Count(model interface{}, conditions map[string]interface{}) (int64, error) {
	var count int64
	if err := p.instance.Model(model).Where(conditions).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
