package mongo

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/custard-technology/abakcus/backend/internal/config"
)

// NewClient creates a MongoDB client using the supplied configuration.
// It attempts to connect and then pings the server to verify the connection.
//
// The caller is responsible for calling Disconnect when the client is no longer
// needed (usually via defer in main).
//
// Errors are returned verbatim so that callers can include them in logs or decide
// whether to retry/exit. Logging should happen upstream to keep this package
// focused on database concerns.
func NewClient(ctx context.Context, cfg config.MongoConfig) (*mongo.Client, error) {
	if cfg.URI == "" {
		return nil, errors.New("mongo: empty URI")
	}

	// create a context with timeout to avoid hanging indefinitely
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(cfg.URI)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}

	// ping the primary to verify connectivity
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		// attempt clean disconnect before returning
		_ = client.Disconnect(context.Background())
		return nil, err
	}

	return client, nil
}
