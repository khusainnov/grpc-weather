package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/khusainnov/grpc-weather/internal/app/weatherservice/endpoint"
	wapi "github.com/khusainnov/grpc-weather/pkg/weatherapi"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s *Server) setupRoutes(e *endpoint.Endpoint) *http.ServeMux {
	grpcMux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err := wapi.RegisterWeatherServiceHandlerFromEndpoint(context.Background(), grpcMux, fmt.Sprintf("localhost%s", s.cfg.GRPCAddr), opts); err != nil {
		s.cfg.L.Error("cannot register gateway handler from endpoint", zap.Error(err))
		return nil
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	mux.HandleFunc("/alive", func(w http.ResponseWriter, r *http.Request) {
		s.cfg.L.Info("alive endpoint")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<h1>Server is working</h1>"))
		return
	})

	return mux
}
