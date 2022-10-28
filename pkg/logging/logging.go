package logging

import (
	"eva/internal/config"
	"os"

	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var e *zap.SugaredLogger

type Logger struct {
	*zap.SugaredLogger
}

func GetLogger() Logger {
	return Logger{e}
}

func init() {
	globalConfig := config.GetConfig()
	loggerConfig := zap.NewProductionEncoderConfig()
	loggerConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	loggerConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	fileEncoder := zapcore.NewJSONEncoder(loggerConfig)
	consoleEncoder := zapcore.NewConsoleEncoder(loggerConfig)
	logFile, _ := os.OpenFile(globalConfig.Server.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer := zapcore.AddSync(logFile)
	//defaultLogLevel := zapcore.DebugLevel
	defaultLogLevel := getLoglevelFromConfig(*globalConfig)
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(colorable.NewColorableStdout()), defaultLogLevel),
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	sugar := logger.Sugar()
	e = sugar
}

// debug -> info -> warn -> error -> dpanic -> panic -> fatal
func getLoglevelFromConfig(globalConfig config.Config) zapcore.Level {
	switch level := globalConfig.Server.LogLevel; level {
	case "DEBUG":
		return zapcore.DebugLevel
	case "INFO":
		return zapcore.DebugLevel
	case "WARNING":
		return zapcore.WarnLevel
	case "ERROR":
		return zapcore.ErrorLevel
	case "DPANIC":
		return zapcore.DPanicLevel
	case "PANIC":
		return zapcore.PanicLevel
	case "FATAL":
		return zapcore.FatalLevel
	default:
		return zapcore.ErrorLevel
	}
}
