package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/khusainnov/grpc-weather/internal/config"
	"github.com/khusainnov/grpc-weather/internal/model"
	wapi "github.com/khusainnov/grpc-weather/pkg/weatherapi"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) GetWeather(cfg *config.Config, req *wapi.WeatherRequest) (*wapi.WeatherResponse, error) {
	API := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", cfg.WeatherToken, req.GetCity())

	resp, err := http.Get(API)
	if err != nil {
		cfg.L.Error("cannot get weather data from API", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "cannot get weather data from API: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		cfg.L.Error("error due reading response body from API", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "error due reading response body from API %v", err)
	}

	locData := &model.Weather{}
	err = json.Unmarshal(body, &locData)
	if err != nil {
		cfg.L.Error("error due unmarshal data", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "error due unmarshalling body %v", err)
	}

	rspLoc := &wapi.LocationBody{
		Region:    locData.Location.Region,
		Name:      locData.Location.Name,
		Country:   locData.Location.Country,
		Localtime: locData.Location.Localtime,
	}

	rspCur := &wapi.Current{
		TempC: locData.Current.TempC,
		TempF: locData.Current.TempF,
	}

	rsp := &wapi.WeatherResponse{
		Location: rspLoc,
		Current:  rspCur,
	}

	if err = s.repo.UploadWeather(rsp.Location.Region); err != nil {
		return nil, fmt.Errorf("cannot write data into db, %w", err)
	}

	return rsp, nil
}
