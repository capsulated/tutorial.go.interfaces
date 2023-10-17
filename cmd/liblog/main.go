package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ServiceLogger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(fields ...interface{})
}

func NewLogger() *zap.SugaredLogger {
	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Development:      true,
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, _ := config.Build()
	defer logger.Sync()

	sugar := logger.Sugar()

	return sugar
}

func main() {
	var log ServiceLogger
	log = NewLogger()

	// Example logs
	log.Debug("This is a debug message.")
	log.Info("This is an info message.")
	log.Warn("This is a warning message.")
	log.Error("This is an error message.")
}
