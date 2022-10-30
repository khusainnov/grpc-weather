package weather

import (
	"context"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	userAgentHeader            = "user-agent"
	grpcForwardedForHeader     = "x-forwarded-for"
)

type Metadata struct {
	UserAgent string
	ClientIP  string
}

func (s *Server) extractMetadata(ctx context.Context) *Metadata {
	mda := &Metadata{}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		logrus.Infof("md: %+v", md)
		if userAgent := md.Get(grpcGatewayUserAgentHeader); len(userAgent) > 0 {
			mda.UserAgent = userAgent[0]
		}

		if userAgent := md.Get(userAgentHeader); len(userAgent) > 0 {
			mda.UserAgent = userAgent[0]
		}

		if clientIPs := md.Get(grpcForwardedForHeader); len(clientIPs) > 0 {
			mda.ClientIP = clientIPs[0]
		}
	}

	if p, ok := peer.FromContext(ctx); ok {
		mda.ClientIP = p.Addr.String()
	}

	return mda
}
