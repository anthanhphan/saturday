package main

import (
	"github.com/anthanhphan/saturday/logger"
	"go.uber.org/zap"
)

func main() {
	logInstance, undo := logger.InitLogger(&logger.Config{
		// Disables caller info in logs (default: false, accepts: bool)
		DisableCaller: false,
		// Disables stack trace in logs (default: false, accepts: bool)
		DisableStacktrace: true,
		// Enables development mode (default: false, accepts: bool)
		EnableDevMode: true,
		// Sets log level (default: Info; defaults to Debug if EnableDevMode is true;
		// accepts: logger.LevelInfo, logger.LevelWarn, logger.LevelError, logger.LevelDebug)
		Level: logger.LevelInfo,
		// Sets log output format (default: JSON; defaults to CONSOLE if EnableDevMode is true;
		// accepts: logger.EncodingConsole, logger.EncodingJSON)
		Encoding: logger.EncodingConsole,
	})
	defer func() {
		_ = logInstance.Sync()
	}()
	defer undo()

	log := zap.L().With(zap.String("app-name", "saturday"), zap.String("version", "1.0")).Sugar()
	log.Debug("Debug message...") // Logged only in development mode or when log level is set to Debug
	log.Info("Info message...")
	log.Warn("Warn message...")
	log.Error("Error message...")
}
