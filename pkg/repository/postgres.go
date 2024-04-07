package repository

import (
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBName   string `json:"DBName"`
	SSLMode  string `json:"SSLMode"`
}

func NewPostgresConnection(config *Config) (*sqlx.DB, error) {
	return sqlx.Connect(
		"pgx", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			config.Host, config.Port, config.Username, config.Password, config.DBName, config.SSLMode))
}
