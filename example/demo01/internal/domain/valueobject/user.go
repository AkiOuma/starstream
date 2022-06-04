package valueobject

import (
	"time"
)

// Value Object
type User struct {
	Name      string
	Height    float64
	Group     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Querier
type UserQuerier struct {
	Id             []int
	IdLower        int
	IdUpper        int
	Name           []string
	SearchName     string
	HeightLower    float64
	HeightUpper    float64
	Group          []int
	GroupLower     int
	GroupUpper     int
	CreatedAtLower time.Time
	CreatedAtUpper time.Time
	UpdatedAtLower time.Time
	UpdatedAtUpper time.Time
}
