package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log level
const (
	LevelInfo  LevelType = "info"
	LevelWarn  LevelType = "warn"
	LevelDebug LevelType = "debug"
	LevelError LevelType = "error"
)

// Log encoding
const (
	EncodingJSON    EncodingType = "json"
	EncodingConsole EncodingType = "console"
)

// Log level map
var logLevelMap = map[LevelType]zapcore.Level{
	LevelDebug: zapcore.DebugLevel,
	LevelInfo:  zap.InfoLevel,
	LevelWarn:  zapcore.WarnLevel,
	LevelError: zapcore.ErrorLevel,
}

// Log key
const (
	LogMessageKey = "message"
	LogTimeKey    = "time"
	LogLevelKey   = "level"
	LogNameKey    = "log"
)
