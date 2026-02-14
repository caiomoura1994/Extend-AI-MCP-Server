package config

import (
	"fmt"
	"os"
)

// Config holds the application configuration
type Config struct {
	APIKey     string
	APIVersion string
	BaseURL    string
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
		APIKey:     apiKey,
		APIVersion: apiVersion,
		BaseURL:    baseURL,
	}, nil
}
