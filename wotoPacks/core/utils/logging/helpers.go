package logging

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var SUGARED *zap.SugaredLogger

func InitZapLog() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	logger, _ := config.Build()
	return logger
}

func Debug(args ...interface{}) {
	if SUGARED != nil {
		SUGARED.Debug(args...)
	} else {
		log.Println(args...)
	}
}

func Warn(args ...interface{}) {
	if SUGARED != nil {
		SUGARED.Warn(args...)
	} else {
		log.Println(args...)
	}
}

func Info(args ...interface{}) {
	if SUGARED != nil {
		SUGARED.Info(args...)
	} else {
		log.Println(args...)
	}
}

func Error(args ...interface{}) {
	if SUGARED != nil {
		SUGARED.Error(args...)
		SUGARED.Error(args...)
	} else {
		log.Println(args...)
	}
}
