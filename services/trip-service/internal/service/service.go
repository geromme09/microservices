package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/shared/types"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type tripService struct {
	repo domain.TripRepository
}

func NewTripService(repo domain.TripRepository) domain.TripService {
	return &tripService{
		repo: repo,
	}
}

func (s *tripService) CreateTrip(ctx context.Context, fare domain.RideFairModel) (*domain.TripModel, error) {
	t := domain.TripModel{
		ID:       primitive.NewObjectID(),
		UserId:   fare.UserId,
		Status:   "pending",
		RideFair: fare,
	}
	return s.repo.CreateTrip(ctx, t)
}

func (s *tripService) GetRoute(ctx context.Context, pickup, destination types.Coordinate) (*types.OsmrApiResponse, error) {

	url := fmt.Sprintf("http://router.project-osrm.org/route/v1/driving/%f,%f;%f,%f?overview=full&geometries=geojson", pickup.Longitude, pickup.Latitude, destination.Longitude, destination.Latitude)

	client := http.Client{}

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get route: %w", err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var respData types.OsmrApiResponse
	if err := json.Unmarshal(b, &respData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal route data: %w", err)
	}

	return &respData, nil
}
