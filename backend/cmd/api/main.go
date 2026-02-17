package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/custard-technology/abakcus/backend/internal/config"
	mongopkg "github.com/custard-technology/abakcus/backend/internal/repository/mongo"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	cfg, err := config.LoadMongoConfig()
	if err != nil {
		log.Fatalf("configuration error: %v", err)
	}

	log.Printf("connecting to MongoDB at %s", cfg.URI)
	ctx := context.Background()
	client, err := mongopkg.NewClient(ctx, cfg)
	if err != nil {
		log.Fatalf("MongoDB connection failed: %v", err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Printf("error disconnecting MongoDB client: %v", err)
		}
	}()
	log.Printf("MongoDB connection successful")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	log.Printf("shutting down")
}
