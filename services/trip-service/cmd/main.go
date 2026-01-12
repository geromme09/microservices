package main

import (
	"fmt"
	"log"
	"net/http"
	"ride-sharing/services/trip-service/internal/infrastructure/handler"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
)

func main() {

	fmt.Println("Trip Service is running!!!!!!")
	inMemo := repository.NewInMemTripRepository()
	srv := service.NewTripService(inMemo)
	handler := handler.NewTripPreviewHandler(srv)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /trip/preview", handler.HandleTripPreview)

	server := &http.Server{
		Addr:    ":8082",
		Handler: mux,
	}

	fmt.Println("Trip service is running on port 8082")

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
