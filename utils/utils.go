package utils

import (
	"fmt"
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger = getLogger()

func ReadFile(filePath string) []byte {
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func GetEnv(key string, dflt ...string) string {
	result, success := os.LookupEnv(key)
	if success {
		return result
	}
	if len(dflt) > 0 {
		return dflt[0]
	}
	return ""
}

func getLogger() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.DisableStacktrace = true
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	if err != nil {
		panic(interface{}(fmt.Errorf("failed to create a logger: %s", err)))
	}
	return logger.Named("dev-platform-gateway")
}
