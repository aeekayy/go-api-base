package cron

import (
	"github.com/aeekayy/go-api-base/pkg/database"
)

// Config represents cron configuration
type Config struct {
	DB      database.DBConfig `json:"db,omitempty" yaml:"db,omitempty"`
	SetupDB bool              `json:"setup_db,omitempty" yaml:"setup_db,omitempty"`
}

// NewConfig returns a new, blank configuration for Cron
func NewConfig() *Config {
	db := database.NewConfig()

	return &Config{
		DB:      *db,
		SetupDB: false,
	}
}
