// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

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
	EbikesAvailable    int32
	BikeDocksAvailable int32
	Ebikes             json.RawMessage
	CreatedAt          time.Time
}

type StationsTimeseries struct {
	ID                 string
	Name               string
	Lat                float64
	Lon                float64
	EbikesAvailable    int32
	BikeDocksAvailable int32
	Ebikes             json.RawMessage
	LastUpdatedMs      int64
	CreatedAt          time.Time
}
