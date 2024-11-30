package gzlog

import (
	"context"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm/logger"
)

type ContextFn func(ctx context.Context) []zapcore.Field

type GormLogger struct {
	LogLevel      logger.LogLevel
	SlowThreshold time.Duration
}

// NewGormLogger creates a new GormLogger instance with customizable log level and slow query threshold.
//
// Parameters:
//   - level: The logging level to be used by the logger (Error = 1, Warn = 2, Info = 3, Silent = -1)
//   - slowThreshold: The threshold in milliseconds after which a query is considered slow
//
// Returns:
//   - GormLogger: A new instance of GormLogger configured with the specified settings
//
// Example:
//
//	// Create a logger with Info level and 200ms slow query threshold
//	logger := NewGormLogger(gormlogger.Info, 200)
//	db.Logger = logger
//
//	// Create a logger with Warning level and 500ms slow query threshold
//	logger := NewGormLogger(gormlogger.Warn, 500)
//	db.Logger = logger
func NewGormLogger(level logger.LogLevel, slowThreshold int64) GormLogger {
	return GormLogger{
		LogLevel:      level,
		SlowThreshold: time.Duration(slowThreshold) * time.Millisecond,
	}
}

// Trace logs SQL execution time and query details based on configured thresholds.
//
// Parameters:
//   - ctx: Context for the operation
//   - begin: Start time of the operation
//   - fc: Function that returns the SQL query and affected rows
//   - err: Any error that occurred during execution
//
// Example:
//
//	logger.Trace(ctx, time.Now(), func() (string, int64) {
//	    return "SELECT * FROM users", 10
//	}, nil)
func (gl GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if gl.LogLevel <= 0 {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && gl.LogLevel >= logger.Error:
		sql, rows := fc()
		zap.L().Sugar().Error("[TRACE] ", zap.Error(err), zap.Duration("duration", elapsed), zap.Int64("rows", rows), zap.String("sql query", sql))
	case gl.SlowThreshold != 0 && elapsed > gl.SlowThreshold && gl.LogLevel >= logger.Warn:
		sql, rows := fc()
		zap.L().Sugar().Warn("[TRACE] ", zap.Duration("duration", elapsed), zap.Int64("rows", rows), zap.String("sql query", sql))
	case gl.LogLevel >= logger.Info:
		sql, rows := fc()
		zap.L().Sugar().Info("[TRACE] ", zap.Duration("duration", elapsed), zap.Int64("rows", rows), zap.String("sql query", sql))
	}
}

func (gl GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return GormLogger{
		SlowThreshold: gl.SlowThreshold,
		LogLevel:      level,
	}
}

func (gl GormLogger) Info(ctx context.Context, str string, args ...interface{}) {
	if gl.LogLevel < logger.Info {
		return
	}

	zap.L().Sugar().Infof(str, args...)
}

func (gl GormLogger) Warn(ctx context.Context, str string, args ...interface{}) {
	if gl.LogLevel < logger.Warn {
		return
	}

	zap.L().Sugar().Warnf(str, args...)
}

func (gl GormLogger) Error(ctx context.Context, str string, args ...interface{}) {
	if gl.LogLevel < logger.Error {
		return
	}

	zap.L().Sugar().Errorf(str, args...)
}
