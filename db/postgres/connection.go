package postgres

import (
	"fmt"
	"time"
)

type Connection struct {
	Host                  string
	Port                  int64
	Database              string
	User                  string
	Password              string
	TimeZone              string
	SSLCert               string
	SSLKey                string
	SSLRootCert           string
	SSLMode               SSLMode
	MaxOpenConnections    int
	MaxIdleConnections    int
	ConnectionMaxIdleTime time.Duration
	ConnectionMaxLifeTime time.Duration
	ConnectionTimeout     time.Duration
}

// ToPostgresConnectionString returns the Postgres connection string based on the Connection struct.
//
// The method constructs a connection string with the following features:
//   - Basic connection parameters (user, password, host, database)
//   - TimeZone setting (defaults to "Asia/Ho_Chi_Minh" if not specified)
//   - SSL configuration based on SSLMode and related certificates
//
// Returns:
//   - string: A formatted PostgreSQL connection string suitable for database connection
//
// Example:
//
//	conn := Connection{
//	    Host:     "localhost",
//	    User:     "postgres",
//	    Password: "secret",
//	    Database: "mydb",
//	    SSLMode:  Disable,
//	}
//	connString := conn.ToPostgresConnectionString()
//	// Results in: "postgresql://postgres:secret@localhost/mydb?TimeZone=Asia/Ho_Chi_Minh&sslmode=disable"
func (conn Connection) ToPostgresConnectionString() string {
	// Use a default timezone if none is provided.
	if conn.TimeZone == "" {
		conn.TimeZone = "Asia/Ho_Chi_Minh" // Default timezone
	}

	// Start with the base connection string.
	dsn := fmt.Sprintf("postgresql://%s:%s@%s/%s?TimeZone=%s&sslmode=%s",
		conn.User,
		conn.Password,
		conn.Host,
		conn.Database,
		conn.TimeZone,
		conn.SSLMode,
	)

	// Append SSL parameters if SSL mode is not "disable".
	if conn.SSLMode != Disable {
		if conn.SSLCert != "" {
			dsn += fmt.Sprintf("&sslcert=%s", conn.SSLCert)
		}
		if conn.SSLKey != "" {
			dsn += fmt.Sprintf("&sslkey=%s", conn.SSLKey)
		}
		if conn.SSLRootCert != "" {
			dsn += fmt.Sprintf("&sslrootcert=%s", conn.SSLRootCert)
		}
	}

	return dsn
}
