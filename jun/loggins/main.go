package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/rs/zerolog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func benchmarkZap() {
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:     "ts",
		LevelKey:    "level",
		MessageKey:  "msg",
		EncodeTime:  zapcore.EpochTimeEncoder,
		EncodeLevel: zapcore.LowercaseLevelEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.AddSync(os.Stdout),
		zapcore.InfoLevel,
	)
	logger := zap.New(core)
	defer logger.Sync()

	logger.Info("request processed",
		zap.String("method", "GET"),
		zap.String("path", "/api/v1/users"),
		zap.Int("status", 200),
		zap.Duration("latency", 12*time.Millisecond),
	)
}

func benchmarkZeroLog() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	logger.Info().
		Str("method", "GET").
		Str("path", "/api/v1/users").
		Int("status", 200).
		Dur("latency", 12*time.Millisecond).
		Msg("request processed")
}

func bencharkSlog() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("request processed",
		"method", "GET",
		"path", "/api/v1/users",
		"status", 200,
		"latency", 12*time.Millisecond,
	)
}
