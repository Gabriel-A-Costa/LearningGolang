package logger

import (
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func LoggerWithContext() *zap.Logger {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	logger := zap.Must(zap.NewProduction())

	if os.Getenv("APP_ENV") == "development" {
		logger = zap.Must(zap.NewDevelopment())
	}

	// childLogger := logger.With(
	// 	zap.String("service", "userService"),
	// 	zap.String("requestID", "abc123"),
	// )

	return logger
}
