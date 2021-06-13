package database

// DBConfig object for database configuration
type DBConfig struct {
	Host        string `json:"host,omitempty" yaml:"host,omitempty"`
	Port        int    `json:"port,omitempty" yaml:"port,omitempty"`
	Username    string `json:"username,omitempty" yaml:"username,omitempty"`
	Password    string `json:"password,omitempty" yaml:"password,omitempty"`
	DBName      string `json:"dbname,omitempty" yaml:"dbname,omitempty"`
	SSLMode     bool   `json:"ssl_mode,omitempty" yaml:"ssl_mode,omitempty"`
	AutoMigrate bool   `json:"auto_migrate,omitempty" yaml:"auto_migrate,omitempty"`
}

// NewConfig returns a new, blank database configuration
func NewConfig() *DBConfig {
	return &DBConfig{
		Host:        "localhost",
		Port:        5432,
		Username:    "postgres",
		Password:    "postgres",
		DBName:      "postgres",
		SSLMode:     false,
		AutoMigrate: false,
	}
}
