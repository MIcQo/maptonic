package api

import "context"

type SearchGeocodingRequest struct {
	Term string `json:"term" required:"true"`
}

type SearchGeocodingResponse struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

func SearchGeocodingHandler(_ context.Context, _ *SearchGeocodingRequest) (*SearchGeocodingResponse, error) {
	// @TODO: reverse geocoding
	return &SearchGeocodingResponse{}, nil
}

func SearchGeocode(lat, lon float64, zoom uint) (any, error) {

	return nil, nil
}
