package logging

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/austinrgray/signalcodex/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger(svcName string) (*zap.Logger, error) {
	currentDate := time.Now().UTC().Format("2006-01-02")
	debugURL := fmt.Sprintf("logs/%s/%s_debug.log", svcName, currentDate)
	errorURL := fmt.Sprintf("logs/%s/%s_error.log", svcName, currentDate)

	debugFile, err := os.OpenFile(debugURL, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return DefaultLogger(svcName), fmt.Errorf(string(LogErrFmtDebugLoggerInit), err)
	}

	errorFile, err := os.OpenFile(errorURL, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return DefaultLogger(svcName), fmt.Errorf(string(LogErrFmtErrorLoggerInit), err)
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
		zap.String(string(LogKeySvcName), svcName),
	), nil
}

func DefaultLogger(svcName string) *zap.Logger {
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
		zap.String(string(LogKeySvcName), svcName),
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

func NewLoggerContext(parent context.Context, fields ...zap.Field) context.Context {
	parentLogger := parent.Value(utils.ContextLoggerKey).(*zap.Logger)
	logger := parentLogger.With(fields...)
	return context.WithValue(parent, utils.ContextLoggerKey, logger)
}

func LoggerFromContext(ctx context.Context) *zap.Logger {
	logger, ok := ctx.Value(utils.ContextLoggerKey).(*zap.Logger)
	if !ok {
		return zap.NewNop()
	}
	return logger
}
