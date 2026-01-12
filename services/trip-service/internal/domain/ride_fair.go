package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RideFairModel struct {
	ID                primitive.ObjectID
	UserId            string
	PackageSlug       string // ex: van,luxury,suv
	TotalPriceInCents float64
	ExpiredAt         time.Time
}
