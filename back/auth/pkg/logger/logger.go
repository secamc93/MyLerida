package logger

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"time"

	"github.com/rs/zerolog"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

type ILogger interface {
	Debug(msg string, params ...interface{})
	Info(msg string, params ...interface{})
	Warn(msg string, params ...interface{})
	Error(msg string, params ...interface{})
	Fatal(msg string, params ...interface{})
	SetOutput(w io.Writer)
	SetLogLevel(level LogLevel)
	Writer() io.Writer
}

// zerologLogger es una implementación de ILogger usando Zerolog.
type zerologLogger struct {
	logger zerolog.Logger
}

// NewLogger retorna una instancia de ILogger basada en Zerolog.
func NewLogger() ILogger {
	// Utilizamos ConsoleWriter para una salida legible en consola.
	writer := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}
	// Se quita la llamada a Caller para que no aparezca en todos los logs.
	zl := zerolog.New(writer).With().Timestamp().Logger()
	// Ajuste: configurar el nivel global según DB_LOG_MODE
	if os.Getenv("DB_LOG_MODE") == "debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	return &zerologLogger{logger: zl}
}

func (z *zerologLogger) Debug(msg string, params ...interface{}) {
	z.logger.Debug().Msgf(msg, params...)
}

func (z *zerologLogger) Info(msg string, params ...interface{}) {
	z.logger.Info().Msgf(msg, params...)
}

func (z *zerologLogger) Warn(msg string, params ...interface{}) {
	z.logger.Warn().Msgf(msg, params...)
}

func (z *zerologLogger) Error(msg string, params ...interface{}) {
	// Obtener la ubicación real de la llamada
	_, file, line, ok := runtime.Caller(2)
	caller := ""
	if ok {
		caller = fmt.Sprintf("%s:%d", file, line)
	}
	// Se añade el campo "caller" al log para indicar la ubicación
	z.logger.Error().Str("caller", caller).Msgf(msg, params...)
}

func (z *zerologLogger) Fatal(msg string, params ...interface{}) {
	z.logger.Fatal().Msgf(msg, params...)
}

func (z *zerologLogger) SetOutput(w io.Writer) {
	z.logger = z.logger.Output(w)
}

func (z *zerologLogger) SetLogLevel(level LogLevel) {
	// Mapear nuestros niveles personalizados a los de Zerolog.
	var zLevel zerolog.Level
	switch level {
	case DEBUG:
		zLevel = zerolog.DebugLevel
	case INFO:
		zLevel = zerolog.InfoLevel
	case WARN:
		zLevel = zerolog.WarnLevel
	case ERROR:
		zLevel = zerolog.ErrorLevel
	case FATAL:
		zLevel = zerolog.FatalLevel
	default:
		zLevel = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(zLevel)
}

func (z *zerologLogger) Writer() io.Writer {
	// Zerolog no expone un Writer directamente, por lo que devolvemos os.Stdout (o podrías almacenar el writer configurado).
	return os.Stdout
}
