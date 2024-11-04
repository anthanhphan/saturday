package logger

type Config struct {
	// DisableCaller specifies whether to omit the caller's file and line number in log messages.
	DisableCaller bool
	// DisableStacktrace specifies whether to disable stacktrace logging on errors.
	DisableStacktrace bool
	// EnableDevMode indicates whether the logger should operate in development mode with more verbose output.
	EnableDevMode bool
	// Level defines the logging level (e.g., debug, info, warn, error).
	Level LevelType
	// Encoding specifies the format for log output (e.g., json, console).
	Encoding EncodingType
}
