package utils

import (
	"os"

	"go.uber.org/zap"
)

func ProvideLogger() *zap.Logger {
	var logger *zap.Logger

	if os.Getenv("ENVIRONMENT") == "production" {
		logger, _ = zap.NewProduction()
	} else {
		logger, _ = zap.NewDevelopment()
	}
	return logger
}
