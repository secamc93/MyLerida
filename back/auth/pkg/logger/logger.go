package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	SUCCESS
	WARN
	ERROR
	FATAL
)

type ILogger interface {
	Debug(msg string, params ...interface{})
	Info(msg string, params ...interface{})
	Success(msg string, params ...interface{})
	Warn(msg string, params ...interface{})
	Error(msg string, params ...interface{})
	Fatal(msg string, params ...interface{})
	SetOutput(w io.Writer)
	SetLogLevel(level LogLevel)
	Writer() io.Writer
}

type logger struct {
	colorReset  string
	colorCyan   string
	colorRed    string
	colorYellow string
	colorGreen  string
	logLevel    LogLevel
	output      io.Writer
}

func NewLogger() ILogger {
	return &logger{
		colorReset:  "\033[0m",
		colorCyan:   "\033[36m",
		colorRed:    "\033[31m",
		colorYellow: "\033[33m",
		colorGreen:  "\033[32m",
		logLevel:    INFO,
		output:      os.Stdout,
	}
}

func (l *logger) log(level LogLevel, color, msg string, params ...interface{}) {
	if level < l.logLevel {
		return
	}

	_, file, line, _ := runtime.Caller(2)
	filename := filepath.Base(file)

	var origin string
	if filename == "db.go" {
		origin = fmt.Sprintf("[config.ConnectDB:%d]", line)
	} else {
		origin = fmt.Sprintf("[%s:%d]", strings.TrimSuffix(filename, filepath.Ext(filename)), line)
	}

	logMsg := fmt.Sprintf("[%s] %s [%s]", level.String(), origin, fmt.Sprintf(msg, params...))
	log.SetOutput(l.output)
	log.Printf("%s%s%s", color, logMsg, l.colorReset)
}

func (l *logger) Debug(msg string, params ...interface{}) {
	l.log(DEBUG, l.colorCyan, msg, params...)
}

func (l *logger) Info(msg string, params ...interface{}) {
	l.log(INFO, l.colorCyan, msg, params...)
}

func (l *logger) Success(msg string, params ...interface{}) {
	l.log(SUCCESS, l.colorGreen, msg, params...)
}

func (l *logger) Warn(msg string, params ...interface{}) {
	l.log(WARN, l.colorYellow, msg, params...)
}

func (l *logger) Error(msg string, params ...interface{}) {
	l.log(ERROR, l.colorRed, msg, params...)
}

func (l *logger) Fatal(msg string, params ...interface{}) {
	l.log(FATAL, l.colorRed, msg, params...)
	log.Panicf("[%s]", fmt.Sprintf(msg, params...))
}

func (l *logger) SetOutput(w io.Writer) {
	l.output = w
}

func (l *logger) SetLogLevel(level LogLevel) {
	l.logLevel = level
}

func (l *logger) Writer() io.Writer {
	return l.output
}

func (level LogLevel) String() string {
	switch level {
	case DEBUG:
		return "DEBU"
	case INFO:
		return "INFO"
	case SUCCESS:
		return "SUCC"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERRO"
	case FATAL:
		return "FATA"
	default:
		return "UNKN"
	}
}
