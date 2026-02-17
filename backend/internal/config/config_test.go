package config

import (
	"os"
	"testing"
)

func TestLoadMongoConfig(t *testing.T) {
	origURI := os.Getenv("MONGO_URI")
	origDB := os.Getenv("MONGO_DB")
	defer func() {
		os.Setenv("MONGO_URI", origURI)
		os.Setenv("MONGO_DB", origDB)
	}()

	// both missing
	os.Unsetenv("MONGO_URI")
	os.Unsetenv("MONGO_DB")
	if _, err := LoadMongoConfig(); err == nil {
		t.Fatal("expected error when both env vars are missing")
	}

	// missing db
	os.Setenv("MONGO_URI", "mongodb://localhost:27017")
	os.Unsetenv("MONGO_DB")
	if _, err := LoadMongoConfig(); err == nil {
		t.Fatal("expected error when MONGO_DB is missing")
	}

	// missing uri
	os.Unsetenv("MONGO_URI")
	os.Setenv("MONGO_DB", "testdb")
	if _, err := LoadMongoConfig(); err == nil {
		t.Fatal("expected error when MONGO_URI is missing")
	}

	// valid
	os.Setenv("MONGO_URI", "mongodb://localhost:27017")
	os.Setenv("MONGO_DB", "testdb")
	cfg, err := LoadMongoConfig()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.URI != "mongodb://localhost:27017" {
		t.Errorf("uri mismatch: got %s", cfg.URI)
	}
	if cfg.Database != "testdb" {
		t.Errorf("db mismatch: got %s", cfg.Database)
	}
}
