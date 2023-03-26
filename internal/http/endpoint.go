package http

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/khusainnov/grpc-weather/internal/app/weatherservice/endpoint"
	wapi "github.com/khusainnov/grpc-weather/pkg/weatherapi"
	"go.uber.org/zap"
)

func (s *Server) setupRoutes(e *endpoint.Endpoint) *http.ServeMux {
	grpcMux := runtime.NewServeMux()

	if err := wapi.RegisterWeatherServiceHandlerServer(context.Background(), grpcMux, e); err != nil {
		s.cfg.L.Error("cannot register weather gateway", zap.Error(err))
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
