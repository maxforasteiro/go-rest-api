package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// DBConfig represents the db configuration
type DBConfig struct {
	Host string
	Port int
	Username string
	Password string
	DBName string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
			Host: "localhost",
			Port: 5432,
			Username: "postgres",
			Password: "password",
			DBName: "todo_development",
		}

	return &dbConfig
}

func DbURL(c *DBConfig) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host,
		c.Port,
		c.Username,
		c.Password,
		c.DBName,
	)
}
