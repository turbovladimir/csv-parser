package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var con *gorm.DB

func NewConnection() *gorm.DB {

	if con != nil {
		return con
	}

	config := ResolveDBConfig()

	con, _ = gorm.Open(postgres.Open(config.ToString()))

	return con
}

type DsnConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func ResolveDBConfig() *DsnConfig {
	return &DsnConfig{
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	}
}

func (c DsnConfig) ToString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		c.Host, c.User, c.Password, c.Database, c.Port,
	)
}
