package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/custard-technology/abakcus/backend/internal/config"
	"github.com/custard-technology/abakcus/backend/internal/handler"
	mongopkg "github.com/custard-technology/abakcus/backend/internal/repository/mongo"
	"github.com/custard-technology/abakcus/backend/internal/service"
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

	menuRepo := mongopkg.NewMenuRepository(client, cfg.Database)
	menuSvc := service.NewMenuService(menuRepo)
	menuHandler := handler.NewMenuHandler(menuSvc)

	mux := http.NewServeMux()
	mux.HandleFunc("/menus", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/menus" {
			if r.Method == http.MethodPost {
				menuHandler.CreateMenu(w, r)
			} else if r.Method == http.MethodGet {
				menuHandler.ListMenus(w, r)
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
		} else {
			if r.Method == http.MethodGet {
				menuHandler.GetMenu(w, r)
			} else if r.Method == http.MethodPut {
				menuHandler.UpdateMenu(w, r)
			} else if r.Method == http.MethodDelete {
				menuHandler.DeleteMenu(w, r)
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
		}
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("starting server on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	log.Printf("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("server shutdown error: %v", err)
	}

	log.Printf("server stopped")
}
