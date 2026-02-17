package config

import (
	"errors"
	"os"
)

// MongoConfig holds the values necessary to connect to MongoDB.
// It is intentionally small; additional configuration can be added later
// if needed (e.g. connection timeouts, max pool size, etc.).
//
// Fields are read from environment variables so the application can be
// configured via Renderer secrets or an .env file in development.
//
// Example usage:
//    cfg, err := config.LoadMongoConfig()
//    if err != nil {
//        log.Fatalf("failed to load config: %v", err)
//    }
//
// This follows the "Validate and Sanitize Input" rule from rules.md.

type MongoConfig struct {
	URI      string
	Database string
}

func LoadMongoConfig() (MongoConfig, error) {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		return MongoConfig{}, errors.New("MONGO_URI is required")
	}

	db := os.Getenv("MONGO_DB")
	if db == "" {
		return MongoConfig{}, errors.New("MONGO_DB is required")
	}

	return MongoConfig{URI: uri, Database: db}, nil
}
