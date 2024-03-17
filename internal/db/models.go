// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"encoding/json"
	"time"
)

type Station struct {
	ID                 string
	Name               string
	Lat                float64
	Lon                float64
	BikesAvailable     int32
	EbikesAvailable    int32
	Ebikes             json.RawMessage
	BikeDocksAvailable int32
	LastUpdatedMs      int64
	IsOffline          bool
	CreatedAt          time.Time
}

type StationsTimeseries struct {
	ID                 string
	Name               string
	Lat                float64
	Lon                float64
	BikesAvailable     int32
	EbikesAvailable    int32
	Ebikes             json.RawMessage
	BikeDocksAvailable int32
	LastUpdatedMs      int64
	IsOffline          bool
	CreatedAt          time.Time
}
