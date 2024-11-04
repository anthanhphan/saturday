package logger

import (
	"log"

	"go.uber.org/zap"
)

// InitLogger creates and initializes a new zap logger instance based on the provided configuration.
// It also sets up the logger as the global zap logger.
//
// Parameters:
//   - cfg: Configuration settings for the logger
//   - opts: Optional zap options to customize logger behavior
//
// Returns:
//   - *zap.Logger: The configured logger instance
//   - func(): A cleanup function that restores the previous global logger when called
func InitLogger(cfg *Config, opts ...zap.Option) (*zap.Logger, func()) {
	logger, err := createLogger(cfg, opts...)
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	undo := zap.ReplaceGlobals(logger)

	return logger, undo
}
