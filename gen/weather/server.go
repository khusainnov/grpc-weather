package weather

import (
	"net"
	"time"

	pb "github.com/khusainnov/grpc-weather/gen/pb/proto"
	"github.com/khusainnov/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	logger = logging.GetLogger()
)

const (
	timeTTL = time.Second * 15
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
