package env

import (
	"auth/pkg/logger"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	ServerPort string
	DBLogMode  string
	JwtSecret  string
	BaseUrl    string
}

func LoadEnv() *Env {
	log := logger.NewLogger()

	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env := &Env{
		DBHost:     getEnv("DB_HOST", log),
		DBPort:     getEnv("DB_PORT", log),
		DBName:     getEnv("DB_NAME", log),
		DBUser:     getEnv("DB_USER", log),
		DBPassword: getEnv("DB_PASSWORD", log),
		ServerPort: getEnv("SERVER_PORT", log),
		DBLogMode:  getEnv("DB_LOG_MODE", log),
		JwtSecret:  getEnv("JWT_SECRET", log),
		BaseUrl:    getEnv("BASE_URL", log),
	}

	return env
}

func getEnv(key string, log logger.ILogger) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatal("Environment variable %s not set", key)
	}
	return value
}
