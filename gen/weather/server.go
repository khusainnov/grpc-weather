package weather

import (
	"net"
	"net/http"
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
	httpServer *http.Server
	grpcServer *grpc.Server
}

type WeatherServer struct {
	pb.UnimplementedWeatherServiceServer
}

func (s *Server) RunHTTPServer(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:              ":" + port,
		Handler:           handler,
		ReadTimeout:       timeTTL,
		ReadHeaderTimeout: timeTTL,
		WriteTimeout:      timeTTL,
		IdleTimeout:       timeTTL,
		MaxHeaderBytes:    1 << 20,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) RunGRPCServer(port string) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Errorf("Cannot create listener: %s", err.Error())
		return err
	}
	s.grpcServer = grpc.NewServer()

	pb.RegisterWeatherServiceServer(s.grpcServer, &WeatherServer{})
	reflection.Register(s.grpcServer)

	return s.grpcServer.Serve(lis)
}
