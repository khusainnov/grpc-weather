package endpoint

import (
	"context"

	wapi "github.com/khusainnov/grpc-weather/pkg/weatherapi"
)

func (e *Endpoint) WeatherRequest(ctx context.Context, req *wapi.WeatherRequest) (*wapi.WeatherResponse, error) {

	return nil, nil
}
