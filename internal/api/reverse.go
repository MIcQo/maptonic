package api

import (
	"context"
	"github.com/MIcQo/maptonic/internal/db"
)

type ReverseGeocodingRequest struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

type ReverseGeocodingResponse struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

func ReverseGeocodingHandler(_ context.Context, _ *ReverseGeocodingRequest) (*ReverseGeocodingResponse, error) {
	// @TODO: reverse geocoding
	return &ReverseGeocodingResponse{}, nil
}

func ReverseGeocode(lat, lon float64, zoom uint) (any, error) {
	return db.ReverseGecode(context.Background(), lat, lon)
}
