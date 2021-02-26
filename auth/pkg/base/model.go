package base

import (
	"time"

	"github.com/google/uuid"
)

type (
	Identification struct {
		ID       uuid.UUID `bson:"id"`
		TenantID string    `bson:"tenant_id"`
		Slug     string    `bson:"slug"`
	}

	Audit struct {
		CreatedByID string    `bson:"created_by_id"`
		UpdatedByID string    `bson:"updated_by_id"`
		CreatedAt   time.Time `bson:"created_at"`
		UpdatedAt   time.Time `bson:"updated_at"`
	}
)

type (
	GeoJson struct {
		Type        string
		Coordinates []float64
	}
)
