package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	pb "github.com/khusainnov/grpc-weather/gen/pb/weather"
	"github.com/khusainnov/grpc-weather/internal/entity"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) WeatherRequest(ctx context.Context, req *pb.RequestData) (*pb.ResponseBody, error) {
	API := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", os.Getenv("WEATHER_API_TOKEN"), req.GetCity())

	logger.Infof("Sending request on API: %s", API)
	resp, err := http.Get(API)
	if err != nil {
		logger.Infof("cannot get weather data from API: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "cannot get weather data from API: %s", err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Infof("Error due reading response body from API: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "Error due reading response body from API: %s", err.Error())
	}

	locData := &entity.Weather{}
	err = json.Unmarshal(body, &locData)
	if err != nil {
		logger.Infof("%s", err.Error())
		return nil, status.Errorf(codes.Internal, "error due unmarshalling body: %s", err.Error())
	}

	rspLoc := &pb.ResponseBody_LocationBody{
		Region:    locData.Location.Region,
		Name:      locData.Location.Name,
		Country:   locData.Location.Country,
		Localtime: locData.Location.Localtime,
	}
	rspCur := &pb.ResponseBody_Current{
		TempC: locData.Current.TempC,
		TempF: locData.Current.TempF,
	}

	rsp := &pb.ResponseBody{
		Location: rspLoc,
		Current:  rspCur,
	}

	return rsp, nil
}
