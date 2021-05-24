package cron 

import (
	"github.com/aeekayy/go-api-base/pkg/database"
)

type Config struct {
	DB		database.DBConfig	`json:"db,omitempty" yaml:"db,omitempty"`
	SetupDB	bool				`json:"setup_db,omitempty" yaml:"setup_db,omitempty"`
}

func NewConfig() *Config {
	db := database.NewConfig()

	return &Config{
		DB:	*db,
		SetupDB: false,
	}
}
