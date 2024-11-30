package postgres

type SSLMode string

const (
	Disable    SSLMode = "disable"
	VerifyFull SSLMode = "verify-full"
	VerifyCA   SSLMode = "verify-ca"
	Require    SSLMode = "require"
	Prefer     SSLMode = "prefer"
	Allow      SSLMode = "allow"
)
