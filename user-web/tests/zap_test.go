package main

import (
	"go.uber.org/zap"
	"testing"
	"time"
)

func Test_Zap(t *testing.T) {
	// json格式
	//logger, _ := zap.NewProduction()
	// 日志格式
	logger, _ := zap.NewDevelopment()
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
	zap.S().Info("hello")
	sugar.Infof("Failed to fetch URL: %s", url)
}
