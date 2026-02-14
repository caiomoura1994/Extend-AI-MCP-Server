package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds the application configuration
type Config struct {
	APIKey           string
	APIVersion       string
	BaseURL          string
	EnableProcessors bool
	EnableWorkflows  bool
	EnableParse      bool
	EnableExtractors bool
	EnableFiles      bool
}

// parseBoolEnv parses a boolean environment variable with a default value
func parseBoolEnv(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	parsed, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}

	return parsed
}

// Load reads configuration from environment variables
func Load() (*Config, error) {
	apiKey := os.Getenv("EXTEND_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("EXTEND_API_KEY environment variable is required")
	}

	apiVersion := os.Getenv("EXTEND_API_VERSION")
	if apiVersion == "" {
		apiVersion = "2026-02-09" // Default version
	}

	baseURL := os.Getenv("EXTEND_BASE_URL")
	if baseURL == "" {
		baseURL = "https://api.extend.ai"
	}

	return &Config{
		APIKey:           apiKey,
		APIVersion:       apiVersion,
		BaseURL:          baseURL,
		EnableProcessors: parseBoolEnv("EXTEND_ENABLE_PROCESSORS", true),
		EnableWorkflows:  parseBoolEnv("EXTEND_ENABLE_WORKFLOWS", true),
		EnableParse:      parseBoolEnv("EXTEND_ENABLE_PARSE", true),
		EnableExtractors: parseBoolEnv("EXTEND_ENABLE_EXTRACTORS", true),
		EnableFiles:      parseBoolEnv("EXTEND_ENABLE_FILES", true),
	}, nil
}
