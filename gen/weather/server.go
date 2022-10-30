package weather

import (
	"context"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/khusainnov/grpc-weather/doc/statik"
	pb "github.com/khusainnov/grpc-weather/gen/pb/weather"
	"github.com/khusainnov/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
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

	jsonOptions := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(jsonOptions)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err = pb.RegisterWeatherServiceHandlerServer(ctx, grpcMux, &Server{}); err != nil {
		logger.Errorf("Error due registering weather handler server: %s", err.Error())
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	fs := http.FileServer(http.Dir("./doc/swagger"))
	/*statikFS, err := fs.New()
	if err != nil {
		logger.Fatalf("Cannot create statik fs: %s", err.Error())
	}*/

	//sgHandler := http.StripPrefix("/swagger/", http.FileServer(statikFS))
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

	err = http.Serve(lis, mux)
	if err != nil {
		logger.Errorf("Error due start the gateway server: %s", err.Error())
		return
	}
}
