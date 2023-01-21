package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Settings struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SslMode  string
}

func NewPostgresDb(settings *Settings) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s "+
		"dbname=%s port=%s sslmode=%s",
		settings.Host,
		settings.User,
		settings.Password,
		settings.DbName,
		settings.Port,
		settings.SslMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
