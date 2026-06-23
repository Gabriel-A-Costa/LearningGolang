package main

import (
	logger "github.com/gabriel-a-costa/zap-logging/log"
	"go.uber.org/zap"
)

func main() {
	log := logger.Amostragem()

	defer log.Sync()

	log.Info("login", zap.Any("user", logger.User{ID: "1", Name: "João", Email: "joao@email.com", Password: "2312314"}))
}
