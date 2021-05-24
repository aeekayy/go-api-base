package api

import (
	"github.com/aeekayy/go-api-base/pkg/database"
)

type Config struct {
	DB				database.DBConfig	`json:"db,omitempty" yaml:"db,omitempty"`
	SetupDB			bool				`json:"setup_db,omitempty" yaml:"setup_db,omitempty"`
	EnableMetrics	bool				`json:"enable_metrics,omitempty" yaml:"enable_metrics,omitempty"`
}

func NewConfig() *Config {
	db := database.NewConfig()

	return &Config{
		DB:	*db,
		SetupDB: false,
		EnableMetrics: false,
	}
}
