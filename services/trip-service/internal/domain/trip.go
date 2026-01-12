package domain

import (
	"context"
	"ride-sharing/shared/types"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TripModel struct {
	ID       primitive.ObjectID
	UserId   string //this is a rider
	Status   string
	RideFair RideFairModel
}

type TripRepository interface {
	CreateTrip(ctx context.Context, trip TripModel) (*TripModel, error)
}

type TripService interface {
	CreateTrip(ctx context.Context, trip RideFairModel) (*TripModel, error)
	GetRoute(ctx context.Context, pickup, destination types.Coordinate) (*types.OsmrApiResponse, error)
}
