package api

import "context"

type ReverseGeocodingRequest struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

type ReverseGeocodingResponse struct {
}

func ReverseGeocodingHandler(_ context.Context, _ *ReverseGeocodingRequest) (*ReverseGeocodingResponse, error) {
	// @TODO: reverse geocoding
	return &ReverseGeocodingResponse{}, nil
}

func ReverseGeocode(lat, lon float64, zoom uint) (any, error) {

	return nil, nil
}
