package mongo

import (
	"context"
	"testing"
	"time"

	"github.com/custard-technology/abakcus/backend/internal/config"
)

func TestNewClient_InvalidURI(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	cfg := config.MongoConfig{URI: "http://not-a-mongo-uri", Database: "test"}
	_, err := NewClient(ctx, cfg)
	if err == nil {
		t.Fatal("expected error for invalid URI")
	}
}
