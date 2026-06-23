package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New retorna um logger Zap pronto para produção.
//
// zap.NewProductionConfig() entrega:
//   - saída JSON (legível por ferramentas como Datadog, Loki, CloudWatch)
//   - nível Info e acima por padrão
//   - stacktrace automático em nível Error
//   - caller (arquivo:linha) em cada entrada
func New() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	return cfg.Build()
}
