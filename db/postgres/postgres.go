package postgres

import (
	"database/sql"
	"fmt"

	"github.com/anthanhphan/saturday/gzlog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	Executor *gorm.DB
}

// NewDatabase creates a new Database instance with customizable connection and logging settings.
//
// Parameters:
//   - conn: Connection configuration containing database connection parameters
//   - logLevel: The logging level for database operations ("error", "warn", "info")
//   - slowQueryThreshold: The threshold in milliseconds after which a query is considered slow
//
// Returns:
//   - *Database: A new Database instance with configured GORM executor
//   - error: Any error encountered during database initialization
//
// Example:
//
//	conn := postgres.Connection{
//	    Host:     "localhost",
//	    Port:     5432,
//	    Database: "mydb",
//	}
//	db, err := NewDatabase(conn, "info", 200)
//	if err != nil {
//	    log.Fatal(err)
//	}
func NewDatabase(conn Connection, logLevel string, slowQueryThreshold int64) (*Database, error) {
	// Create the Postgres connection string
	dsn := conn.ToPostgresConnectionString()

	// Open a database connection with GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 createLogger(logLevel, slowQueryThreshold),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Retrieve the underlying *sql.DB instance
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get SQL DB instance: %w", err)
	}

	// Check database connection
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("database ping failed: %w", err)
	}

	// Set connection pool settings
	configureConnectionPool(sqlDB, conn)

	return &Database{Executor: db}, nil
}

// configureConnectionPool sets the connection pool settings for the database.
func configureConnectionPool(sqlDB *sql.DB, conn Connection) {
	sqlDB.SetMaxOpenConns(conn.MaxOpenConnections)
	sqlDB.SetMaxIdleConns(conn.MaxIdleConnections)
	sqlDB.SetConnMaxIdleTime(conn.ConnectionMaxIdleTime)
	sqlDB.SetConnMaxLifetime(conn.ConnectionMaxLifeTime)
}

// createLogger creates a GORM logger based on the specified log level.
func createLogger(logLevel string, slowQueryThreshold int64) logger.Interface {
	logLevels := map[string]logger.LogLevel{
		"error": logger.Error,
		"warn":  logger.Warn,
		"info":  logger.Info,
	}

	// Default to "info" if the log level is not recognized
	level, exists := logLevels[logLevel]
	if !exists {
		level = logger.Info
	}

	return gzlog.NewGormLogger(level, slowQueryThreshold)
}
