package postgres

import (
	"auth/pkg/env"
	"auth/pkg/logger"
	"fmt"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type DBConnection interface {
	GetDB() *gorm.DB
	CloseDB() error
	Reconnect() error
	PingDB() error
}

type dbConnection struct {
	db *gorm.DB
}

var (
	instance *dbConnection
	once     sync.Once
	log      = logger.NewLogger()
)

func New() DBConnection {
	once.Do(func() {
		instance = &dbConnection{}
	})
	return instance
}

func (conn *dbConnection) GetDB() *gorm.DB {
	if conn.db == nil {
		if err := conn.connect(); err != nil {
			log.Fatal("failed to connect database: %v", err)
			panic("failed to connect database")
		}
	}

	if err := conn.PingDB(); err != nil {
		log.Warn("Database connection lost, attempting to reconnect")
		if err := conn.Reconnect(); err != nil {
			log.Fatal("Failed to reconnect to the database: %v", err)
			panic("Failed to reconnect to the database")
		}
	}
	return conn.db
}

func (conn *dbConnection) connect() error {
	envVars := env.LoadEnv()
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		envVars.DBHost, envVars.DBPort, envVars.DBUser, envVars.DBName, envVars.DBPassword)

	var logLevel gormLogger.LogLevel
	switch envVars.DBLogMode {
	case "silent":
		logLevel = gormLogger.Silent
	case "error":
		logLevel = gormLogger.Error
	case "warn":
		logLevel = gormLogger.Warn
	case "info":
		logLevel = gormLogger.Info
	default:
		logLevel = gormLogger.Silent
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(logLevel),
	})
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	conn.db = db
	log.Success("Database connection established")
	return nil
}

func (conn *dbConnection) CloseDB() error {
	sqlDB, err := conn.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}
	return sqlDB.Close()
}

func (conn *dbConnection) Reconnect() error {
	if err := conn.connect(); err != nil {
		return fmt.Errorf("failed to reconnect database: %w", err)
	}
	log.Info("Database reconnected successfully")
	return nil
}

func (conn *dbConnection) PingDB() error {
	sqlDB, err := conn.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}
	return sqlDB.Ping()
}
