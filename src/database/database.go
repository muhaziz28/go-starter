package database

import (
	"app/src/config"
	"app/src/utils"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(dbHost, dbName string) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		dbHost,
		config.DBUser,
		config.DBPassword,
		dbName,
		config.DBPort,
	)

	utils.Log.Infof("Connecting to database: %s@%s:%d/%s", config.DBUser, dbHost, config.DBPort, dbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		TranslateError:         true,
	})
	if err != nil {
		utils.Log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
		utils.Log.Warnf("Failed to create uuid-ossp extension: %v", err)
	}

	sqlDB, errDB := db.DB()
	if errDB != nil {
		utils.Log.Fatalf("Failed to get database instance: %v", errDB)
	}

	if err := sqlDB.Ping(); err != nil {
		utils.Log.Fatalf("Failed to ping database: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)

	utils.Log.Info("✅ Database connected successfully")
	return db
}
