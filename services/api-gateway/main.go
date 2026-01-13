package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/ws/riders", handleRiderWebSocket)
	mux.HandleFunc("/ws/drivers", handleDriverWebSocket)

	server := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	serveErrors := make(chan error, 1)

	go func() {
		log.Printf("Trip Service is listening on %s", server.Addr)
		serveErrors <- server.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serveErrors:
		log.Fatalf("server error: %v", err)
	case sig := <-shutdown:
		log.Printf("shutdowning down due to %v signal", sig)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("graceful shutdown did not complete in time: %v", err)
			if err := server.Close(); err != nil {
				log.Printf("could not stop server: %v", err)
			}
		}

	}

}
