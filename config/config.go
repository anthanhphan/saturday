package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// NewConfig reads and parses a configuration file into a provided struct type.
// It supports both JSON (.json) and YAML (.yml, .yaml) file formats, automatically
// detecting the format based on the file extension.
//
// The function uses generics to support any struct type that can be unmarshalled
// from JSON or YAML.
//
// Parameters:
//   - configPath: string - Path to the configuration file
//   - configModel: *T - Pointer to the struct that will hold the configuration
//
// Returns:
//   - *T - Pointer to the populated configuration struct
//   - error - Error if any occurred during reading or parsing
//
// Example:
//
//	type MyConfig struct {
//	    Host string `json:"host" yaml:"host"`
//	    Port int    `json:"port" yaml:"port"`
//	}
//
//	config, err := NewConfig("config.yaml", &MyConfig{})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Host: %s, Port: %d\n", config.Host, config.Port)
func NewConfig[T any](configPath string, configModel *T) (*T, error) {
	if configPath == "" {
		return nil, fmt.Errorf("config path is required")
	}

	// Read the file content.
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Determine the file format based on the file extension.
	switch ext := filepath.Ext(configPath); ext {
	case ".json":
		// Unmarshal JSON content
		if err := json.Unmarshal(data, configModel); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
		}
	case ".yml", ".yaml":
		// Unmarshal YAML content
		if err := yaml.Unmarshal(data, configModel); err != nil {
			return nil, fmt.Errorf("failed to unmarshal YAML: %w", err)
		}
	default:
		return nil, fmt.Errorf("unsupported file extension: %s", ext)
	}

	return configModel, nil
}

// GetConfigPath constructs the path to a configuration file based on the environment
// and desired file extension. It follows a standardized file naming convention:
// "./env/env.<environment>.<extension>".
//
// The function supports multiple environments and will fall back to a local configuration
// if an unknown environment is specified.
//
// Supported environments:
//   - "qc" -> Quality Control environment
//   - "staging" -> Staging environment
//   - "production" -> Production environment
//   - any other value -> Falls back to local environment
//
// Parameters:
//   - env: string - The environment name (e.g., "qc", "staging", "production")
//   - ext: ...string - Optional file extension (defaults to "json" if not provided)
//
// Returns:
//   - string - The constructed configuration file path
//
// Example:
//
//	path := GetConfigPath("staging", "json")  // Returns "./env/env.staging.json"
//	path := GetConfigPath("production")       // Returns "./env/env.production.yml"
//	path := GetConfigPath("unknown")          // Returns "./env/env.local.yml"
func GetConfigPath(env string, ext ...string) string {
	fileExt := "json"
	if len(ext) != 0 {
		fileExt = ext[0]
	}

	configPaths := map[string]string{
		"qc":         "./env/env.qc." + fileExt,
		"staging":    "./env/env.staging." + fileExt,
		"production": "./env/env.production." + fileExt,
	}

	// Return the config path for the environment or the default local file.
	if path, exists := configPaths[env]; exists {
		return path
	}

	// Default to a local configuration file based on the provided extension.
	return "./env/env.local." + fileExt
}
