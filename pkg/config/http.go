package config

import (
	"github.com/aeekayy/go-api-base/pkg/database"
)

// JwtConfig configuration object for JSON Web Toens
type JwtConfig struct {
	SecretKey       string `json:"secret_key,omitempty" yaml:"secret_key,omitempty"`
	Issuer          string `json:"issuer,omitempty" yaml:"issuer,omitempty"`
	ExpirationHours int64  `json:"expiration_hours" yaml:"expiration_hours"`
}

// HTTPConfig configuration object for API server
type HTTPConfig struct {
	DB            database.DBConfig `json:"db,omitempty" yaml:"db,omitempty"`
	SetupDB       bool              `json:"setup_db,omitempty" yaml:"setup_db,omitempty"`
	EnableMetrics bool              `json:"enable_metrics,omitempty" yaml:"enable_metrics,omitempty"`
	Port          string            `json:"port,omitempty" yaml:"port,omitempty"`
	EnableCORS    bool              `json:"enable_cors,omitempty" yaml:"enable_cors,omitempty"`
	Jwt           JwtConfig         `json:"jwt,omitempty" yaml:"jwt,omitempty"`
}

// NewHTTPConfig returns a new, blank configuration for API server
func NewHTTPConfig() *HTTPConfig {
	db := database.NewConfig()

	return &HTTPConfig{
		DB:            *db,
		SetupDB:       false,
		EnableMetrics: false,
		Port:          ":8080",
		EnableCORS:    true,
		Jwt:           JwtConfig{},
	}
}
