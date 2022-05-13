package logging

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var SUGARED *zap.SugaredLogger

func InitZapLog(debugMode bool) *zap.Logger {
	var config zap.Config
	if debugMode {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
	}
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	logger, _ := config.Build(zap.AddCallerSkip(1))
	return logger
}

func LoadLogger(debugMode bool) func() {
	if SUGARED != nil {
		return nil
	}
	loggerMgr := InitZapLog(debugMode)
	zap.ReplaceGlobals(loggerMgr)
	SUGARED = loggerMgr.Sugar()

	return func() {
		_ = loggerMgr.Sync()
	}
}

func Warn(args ...interface{}) {
	if SUGARED != nil {
		SUGARED.Warn(args...)
	} else {
		log.Println(args...)
	}
}

func Error(args ...interface{}) {
	if SUGARED != nil {
		SUGARED.Error(args...)
	} else {
		log.Println(args...)
	}
}

// UnexpectedError works like Error function and logs the error details to a
// specified log file (a new log file is used each time).
func UnexpectedError(err error) {
	if SUGARED != nil {
		SUGARED.Error("Unexpected Error: ", err)
	} else {
		log.Println("Unexpected Error: ", err)
	}
	_ = ioutil.WriteFile(GetLogErrorPath(), []byte(err.Error()), fs.ModePerm)
}

func Info(args ...interface{}) {
	if SUGARED != nil {
		SUGARED.Info(args...)
	} else {
		log.Println(args...)
	}
}

func Infof(template string, args ...interface{}) {
	if SUGARED != nil {
		SUGARED.Infof(template, args...)
	} else {
		log.Printf(template, args...)
	}
}

func Debug(args ...interface{}) {
	if SUGARED != nil {
		SUGARED.Debug(args...)
	} else {
		log.Println(args...)
	}
}

func Debugf(template string, args ...interface{}) {
	if SUGARED != nil {
		SUGARED.Debugf(template, args...)
	} else {
		log.Printf(template, args...)
	}
}

func Fatal(args ...interface{}) {
	if SUGARED != nil {
		SUGARED.Fatal(args...)
	} else {
		// skipcq: RVV-A0003
		log.Fatal(args...)
	}
}

func LogPanic(details []byte) {
	p := string(os.PathSeparator)
	path := "logs" + p + "panics" + p +
		"panic_" + time.Now().Format("2006-01-02_15-04-05") + ".log"
	err := ioutil.WriteFile(path, details, fs.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}

func GetLogErrorPath() string {
	p := string(os.PathSeparator)
	return "logs" + p + "errors/" +
		"error_" + time.Now().Format("2006-01-02_15-04-05") + ".log"
}
