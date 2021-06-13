package database

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres" // using postgres sql
	"gorm.io/gorm"

	"github.com/aeekayy/go-api-base/pkg/models"
)

func SetupDatabase(config *DBConfig) (*gorm.DB, error) {
	// To get the value from the config file using key
	// viper package read .env
	viper_user := config.Username
	viper_password := config.Password
	viper_db := config.DBName
	viper_host := config.Host
	viper_port := config.Port

	prosgret_conname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", viper_host, viper_port, viper_user, viper_db, viper_password)

	log.Infof("conname is\t\t%s", prosgret_conname)

	db, err := gorm.Open(postgres.Open(prosgret_conname), &gorm.Config{})
	if err != nil {
		log.Error("Failed to connect to database!")
		return nil, err
	}

	log.Info("Connected to the database")

	log.Infof("Auto-migrate is set to %+v", config.AutoMigrate)

	if config.AutoMigrate {
		log.Info("Moving Migration")
	}

	return db, nil
}

func MigrateDatabase(sqlDB *gorm.DB, config *DBConfig) error {
	log.Info("Migrating the database")
	var err error

	if sqlDB == nil {
		// To get the value from the config file using key
		// viper package read .env
		viper_user := config.Username
		viper_password := config.Password
		viper_db := config.DBName
		viper_host := config.Host
		viper_port := config.Port

		prosgret_conname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", viper_host, viper_port, viper_user, viper_db, viper_password)

		log.Infof("conname is\t\t%s", prosgret_conname)

		sqlDB, err = gorm.Open(postgres.Open(prosgret_conname), &gorm.Config{})
		if err != nil {
			log.Error("Error detected with the migration")
			return err
		}
	}

	log.Info("Migrating users")
	sqlDB.AutoMigrate(&models.User{})
	log.Info("Migrating events")
	sqlDB.AutoMigrate(&models.Event{})
	log.Info("Database migration is complete")

	return nil
}
