package main

import (
	"io"
	"log/slog"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func BenchmarkZap(b *testing.B) {
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:     "ts",
		LevelKey:    "level",
		MessageKey:  "msg",
		EncodeTime:  zapcore.EpochTimeEncoder,
		EncodeLevel: zapcore.LowercaseLevelEncoder,
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.AddSync(io.Discard),
		zapcore.InfoLevel,
	)
	logger := zap.New(core)
	defer logger.Sync()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("request processed",
			zap.String("method", "GET"),
			zap.String("path", "/api/v1/users"),
			zap.Int("status", 200),
			zap.Duration("latency", 12*time.Millisecond),
		)
	}
}

func BenchmarkZeroLog(b *testing.B) {
	logger := zerolog.New(io.Discard).With().Timestamp().Logger()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info().
			Str("method", "GET").
			Str("path", "/api/v1/users").
			Int("status", 200).
			Dur("latency", 12*time.Millisecond).
			Msg("request processed")
	}
}

func BenchmarkSlog(b *testing.B) {
	logger := slog.New(slog.NewJSONHandler(io.Discard, nil))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("request processed",
			"method", "GET",
			"path", "/api/v1/users",
			"status", 200,
			"latency", 12*time.Millisecond,
		)
	}
}

func BenchmarkSlogAttrs(b *testing.B) {
	logger := slog.New(slog.NewJSONHandler(io.Discard, nil))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.LogAttrs(nil, slog.LevelInfo, "request processed",
			slog.String("method", "GET"),
			slog.String("path", "/api/v1/users"),
			slog.Int("status", 200),
			slog.Duration("latency", 12*time.Millisecond),
		)
	}
}
