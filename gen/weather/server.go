package weather

import (
	"context"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/khusainnov/grpc-weather/gen/pb/weather"
	"github.com/khusainnov/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	logger = logging.GetLogger()
)

type Server struct {
	pb.UnimplementedWeatherServiceServer
	grpcServer *grpc.Server
}

func (s *Server) RunGRPCServer(port string) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Errorf("Cannot create listener: %s", err.Error())
		return err
	}
	s.grpcServer = grpc.NewServer()

	pb.RegisterWeatherServiceServer(s.grpcServer, &Server{})
	reflection.Register(s.grpcServer)

	return s.grpcServer.Serve(lis)
}

func (s *Server) RunGatewayServer(port string) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Errorf("Cannot create listener: %s", err.Error())
		return
	}

	grpcMux := runtime.NewServeMux()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err = pb.RegisterWeatherServiceHandlerServer(ctx, grpcMux, &Server{}); err != nil {
		logger.Errorf("Error due registering weather handler server: %s", err.Error())
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	err = http.Serve(lis, mux)
	if err != nil {
		logger.Errorf("Error due start the gateway server: %s", err.Error())
		return
	}
}
