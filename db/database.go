package db

import (
	"fmt"
	config "go-gorm-hierarchy/configs"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var (
	// Database global variable database
	Database = &gorm.DB{}
)

// New open initialize a new db connection.
func New(config config.SQLConfig) (err error) {
	postgreSQLCredentials := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		config.Host,
		config.Port,
		config.Username,
		config.Password,
		config.DatabaseName,
	)

	Database, err = gorm.Open(postgres.Open(postgreSQLCredentials), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return err
	}

	sqlDB, err := Database.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(config.ConnMaxLifetime)
	err = sqlDB.Ping()
	if err != nil {
		return err
	}

	return nil
}

// Debug set debug
func Debug() {
	Database = Database.Debug()
}
