package logger

import (
	"context"
	"time"

	gormLogger "gorm.io/gorm/logger"
)

// GormLogger implementa la interfaz gormLogger.Interface usando nuestro logger.
type GormLogger struct {
	logger   ILogger
	logLevel gormLogger.LogLevel
}

// NewGormLogger crea una instancia de GormLogger.
func NewGormLogger(logger ILogger, level gormLogger.LogLevel) *GormLogger {
	return &GormLogger{
		logger:   logger,
		logLevel: level,
	}
}

func (g *GormLogger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	g.logLevel = level
	return g
}

func (g *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if g.logLevel >= gormLogger.Info {
		g.logger.Info(msg, data...)
	}
}

func (g *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if g.logLevel >= gormLogger.Warn {
		g.logger.Warn(msg, data...)
	}
}

func (g *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if g.logLevel >= gormLogger.Error {
		g.logger.Error(msg, data...)
	}
}

func (g *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	if err != nil {
		if g.logLevel >= gormLogger.Error {
			g.logger.Error("SQL Error: %v [rows:%d] [elapsed:%s]", err, rows, elapsed)
		}
	} else {
		if g.logLevel >= gormLogger.Info {
			g.logger.Info("SQL: %s [rows:%d] [elapsed:%s]", sql, rows, elapsed)
		}
	}
}
