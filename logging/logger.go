package logging

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger(appName string) (*zap.Logger, error) {
	currentDate := time.Now().UTC().Format("2006-01-02")
	debugURL := fmt.Sprintf("logs/%s_debug_%s.log", currentDate, appName)
	errorURL := fmt.Sprintf("logs/%s_error_%s.log", currentDate, appName)

	debugFile, err := os.OpenFile(debugURL, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return DefaultLogger(appName), fmt.Errorf("initLogger: error opening debug log file")
	}

	errorFile, err := os.OpenFile(errorURL, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return DefaultLogger(appName), fmt.Errorf("initLogger: error opening error log file")
	}

	localErrorsFile := zapcore.Lock(errorFile)
	localDebugFile := zapcore.Lock(debugFile)
	stdDebug := zapcore.Lock(os.Stdout)
	stdError := zapcore.Lock(os.Stderr)

	config := zap.NewDevelopmentConfig()
	fileEncoder := zapcore.NewJSONEncoder(config.EncoderConfig)
	consoleEncoder := zapcore.NewConsoleEncoder(config.EncoderConfig)
	highPriority, lowPriority := priorityLevels()

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, localErrorsFile, highPriority),
		zapcore.NewCore(fileEncoder, localDebugFile, lowPriority),
		zapcore.NewCore(consoleEncoder, stdDebug, highPriority),
		zapcore.NewCore(consoleEncoder, stdError, lowPriority),
	)

	return zap.New(core).With(
		zap.String("app", appName),
	), nil
}

func DefaultLogger(appName string) *zap.Logger {
	config := zap.NewDevelopmentConfig()
	consoleEncoder := zapcore.NewConsoleEncoder(config.EncoderConfig)

	stdDebug := zapcore.Lock(os.Stdout)
	stdError := zapcore.Lock(os.Stderr)

	highPriority, lowPriority := priorityLevels()

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, stdDebug, highPriority),
		zapcore.NewCore(consoleEncoder, stdError, lowPriority),
	)
	return zap.New(core).With(
		zap.String("app", appName),
	)
}

func priorityLevels() (zapcore.LevelEnabler, zapcore.LevelEnabler) {
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})
	return highPriority, lowPriority
}
