package logger

import (
	"time"

	"github.com/anthanhphan/saturday/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Package logger provides a structured logging implementation using uber-go/zap.
// It supports both development and production configurations with customizable
// encoding formats, log levels, and output options.

// createLogger initializes and returns a new zap.Logger instance based on the provided configuration.
// It handles both development and production environments with appropriate defaults.
//
// Parameters:
//   - config: Logger configuration options including environment mode, log level, and encoding preferences
//   - opts: Additional zap.Option parameters for further logger customization
//
// Returns:
//   - *zap.Logger: Configured logger instance
//   - error: Any error encountered during logger creation
func createLogger(config *Config, opts ...zap.Option) (*zap.Logger, error) {
	// Get the appropriate zap config (development or production)
	zapConfig := getZapConfigByMode(config.EnableDevMode)

	// Set log level
	if !config.EnableDevMode {
		zapConfig.Level = zap.NewAtomicLevelAt(logLevelMap[utils.DefaultIfEmpty(config.Level, LevelInfo)])
	}

	// Caller and stacktrace configuration
	zapConfig.DisableCaller = config.DisableCaller
	zapConfig.DisableStacktrace = config.DisableStacktrace

	// Encode configuration
	zapConfig.Encoding = string(getEncodeByMode(config.EnableDevMode, config.Encoding))
	zapConfig.EncoderConfig = zapcore.EncoderConfig{
		MessageKey:    LogMessageKey,
		TimeKey:       LogTimeKey,
		LevelKey:      LogLevelKey,
		NameKey:       LogNameKey,
		CallerKey:     optionalKey(config.DisableCaller, "caller"),
		StacktraceKey: optionalKey(config.DisableStacktrace, "stacktrace"),
		EncodeLevel:   getEncodeLevelByMode(config.EnableDevMode),
		EncodeTime:    getTimeEncoder,
		EncodeCaller:  zapcore.FullCallerEncoder,
	}

	return zapConfig.Build(opts...)
}

// getZapConfigByMode returns the appropriate zap.Config based on the development mode flag.
// Development mode enables more verbose logging suitable for debugging.
func getZapConfigByMode(devMode bool) zap.Config {
	if devMode {
		return zap.NewDevelopmentConfig()
	}
	return zap.NewProductionConfig()
}

// getEncodeByMode determines the encoding format based on the development mode
// and user-specified encoding type. Development mode defaults to console encoding,
// while production defaults to JSON encoding if not specified.
func getEncodeByMode(devMode bool, encode EncodingType) EncodingType {
	if devMode {
		return EncodingConsole
	}
	return utils.DefaultIfEmpty(encode, EncodingJSON)
}

// getEncodeLevelByMode returns the level encoder based on the development mode.
// Development mode includes colored output, while production uses capital letters.
func getEncodeLevelByMode(devMode bool) zapcore.LevelEncoder {
	if devMode {
		return zapcore.CapitalColorLevelEncoder
	}
	return zapcore.CapitalLevelEncoder
}

// getTimeEncoder implements zapcore.TimeEncoder to format timestamps
// using RFC3339 format (e.g., "2006-01-02T15:04:05Z07:00").
func getTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(time.RFC3339))
}

// optionalKey returns either the provided key or an empty string based on the disabled flag.
// This is used to conditionally include or exclude certain fields in the log output.
func optionalKey(disabled bool, key string) string {
	if disabled {
		return ""
	}
	return key
}
