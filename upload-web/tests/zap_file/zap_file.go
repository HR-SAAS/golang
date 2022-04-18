package main

import (
	"go.uber.org/zap"
	"time"
)

func NewLogger() (logger *zap.Logger, err error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.OutputPaths = []string{
		"./log.log",
		"stderr",
	}
	return cfg.Build()
}
func main() {
	logger, _ := NewLogger()
	defer logger.Sync()
	url := "www.hr-saas.com"
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	// sugar方便但是 没有logger快速
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}
