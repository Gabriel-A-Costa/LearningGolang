package main

import (
	logger "github.com/gabriel-a-costa/zap-logging/log"
	"go.uber.org/zap"
)

func main() {
	log := logger.Amostragem()

	defer log.Sync()

	// Amostragem
	u := logger.User{ID: "1", Name: "João", Email: "joao@email.com", Password: "2312314"}

	// log.Info("login", zap.Any("user", u))

	child := log.With(zap.String("name", "main"))
	child.Info("an info log", zap.Any("user", u))
}
