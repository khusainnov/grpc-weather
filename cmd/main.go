package main

import (
	"os"

	"github.com/joho/godotenv"
	gw "github.com/khusainnov/grpc-weather/gen/weather"
	"github.com/khusainnov/logging"
)

func main() {
	logger := logging.GetLogger()

	if err := godotenv.Load("./config/envs/.env"); err != nil {
		logger.Infoln("Error due load config: %s", err.Error())
		return
	}

	s := new(gw.Server)

	logger.Infof("Starting Gateway on port:%s", os.Getenv("HTTP_WEATHER_PORT"))
	s.RunGatewayServer(os.Getenv("HTTP_WEATHER_PORT"))

	logger.Infof("Starting grpc server on port:%s", os.Getenv("GRPC_WEATHER_PORT"))
	if err := s.RunGRPCServer(os.Getenv("GRPC_WEATHER_PORT")); err != nil {
		logger.Errorf("%s", err.Error())
		return
	}
}
