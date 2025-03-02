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

	var gormLevel gormLogger.LogLevel
	switch envVars.DBLogMode {
	case "silent":
		gormLevel = gormLogger.Silent
	case "error":
		gormLevel = gormLogger.Error
	case "warn":
		gormLevel = gormLogger.Warn
	case "info":
		gormLevel = gormLogger.Info
	case "debug":
		gormLevel = gormLogger.Info
	default:
		gormLevel = gormLogger.Silent
	}

	// Se utiliza el logger personalizado que integra Zerolog.
	customGormLogger := logger.NewGormLogger(log, gormLevel)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: customGormLogger,
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
	log.Info("Database connection established")
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
