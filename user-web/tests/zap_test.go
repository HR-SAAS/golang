package main

import (
	"go.uber.org/zap"
	"testing"
	"time"
)

func Test_Zap(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	url := "www.hr-saas.com"
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}
