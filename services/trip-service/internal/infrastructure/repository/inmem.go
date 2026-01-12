package repository

import (
	"context"
	"ride-sharing/services/trip-service/internal/domain"
)

type inMemTripRepository struct {
	trips    map[string]domain.TripModel
	rideFair map[string]domain.RideFairModel
}

func NewInMemTripRepository() *inMemTripRepository {
	return &inMemTripRepository{
		trips:    make(map[string]domain.TripModel),
		rideFair: make(map[string]domain.RideFairModel),
	}
}

func (r *inMemTripRepository) CreateTrip(ctx context.Context, trip domain.TripModel) (*domain.TripModel, error) {
	r.trips[trip.ID.Hex()] = trip
	return &trip, nil
}
