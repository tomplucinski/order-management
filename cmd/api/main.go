package main

import (
	"fmt"
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
	"order-management/internal/handler"
)

func main() {
	fmt.Println("Start")

	mux := http.NewServeMux()
	mux.HandleFunc("/orders", handler.OrdersHandler)
	mux.HandleFunc("/orders/", handler.OrderByIDHandler)

	srv := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	go func() {
		log.Println("Starting server on :8080")

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Forced shutdown: %v", err)
	}

	log.Println("Server gracefully stopped.")
}