package app

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"path/filepath"

	"github.com/khusainnov/grpc-weather/internal/app/weatherservice/endpoint"
	"github.com/khusainnov/grpc-weather/internal/app/weatherservice/service"
	"github.com/khusainnov/grpc-weather/internal/config"
	"github.com/khusainnov/grpc-weather/internal/db"
	"github.com/khusainnov/grpc-weather/internal/http"
	"github.com/khusainnov/grpc-weather/internal/repository"
	wapi "github.com/khusainnov/grpc-weather/pkg/weatherapi"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func New(cfg *config.Config) error {
	var s *grpc.Server

	if cfg.AppMode == config.ProdAppMode {
		tlsCert, err := loadTLSCertificate(cfg)
		if err != nil {
			return fmt.Errorf("cannot load tls certificate, %w", err)
		}

		s = grpc.NewServer(
			grpc.ChainUnaryInterceptor(
				logUnaryInterceptor(cfg.L),
			),
			grpc.ChainStreamInterceptor(
				logStreamInterceptor(cfg.L),
			),
			grpc.Creds(tlsCert),
		)
	}

	s = grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logUnaryInterceptor(cfg.L),
		),
		grpc.ChainStreamInterceptor(
			logStreamInterceptor(cfg.L),
		),
	)
	reflection.Register(s)

	dbClient, err := db.NewClient(cfg)
	if err != nil {
		cfg.L.Fatal("not connected to db", zap.Error(err))
	}

	lis, err := net.Listen("tcp", cfg.GRPCAddr)
	if err != nil {
		return fmt.Errorf("cannot create listener, %w", err)
	}

	httpServer := http.New(cfg)
	httpServer.Start()

	repo := repository.NewRepository(dbClient.GetDB())
	srv := service.NewService(repo)
	endpoints := endpoint.NewEndpoint(srv, cfg)

	wapi.RegisterWeatherServiceServer(s, endpoints)

	cfg.L.Info("starting listening grpc server", zap.Any("PORT", cfg.GRPCAddr))
	if err := s.Serve(lis); err != nil {
		cfg.L.Error("error serve grpc server", zap.Error(err))
		return fmt.Errorf("error serve grpc server, %w", err)
	}

	return nil
}

func logUnaryInterceptor(log *zap.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (_ interface{}, err error) {
		log.Info("unary interceptor", zap.Any("method", info.FullMethod))

		return handler(ctx, req)
	}
}

func logStreamInterceptor(log *zap.Logger) grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		log.Info("stream interceptor", zap.Any("method", info.FullMethod))

		return handler(srv, ss)
	}
}

func loadTLSCertificate(cfg *config.Config) (credentials.TransportCredentials, error) {
	serveCert, err := tls.LoadX509KeyPair(filepath.Join(cfg.CertPath, "server-cert.pem"), filepath.Join(cfg.CertPath, "server-key.pem"))
	if err != nil {
		cfg.L.Error("cannot load the tls certificate", zap.Error(err))
		return nil, fmt.Errorf("error due load the tls certificate, %w", err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{serveCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(tlsConfig), nil
}
