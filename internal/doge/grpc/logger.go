package grpc

import (
	"context"
	"time"

	"github.com/chelpis/member-center/pkg/telemetry/log"
	"github.com/chelpis/member-center/pkg/telemetry/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GrpcLogger(serverCtx context.Context) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		logger := log.FromContext(serverCtx)
		tracer := trace.FromContext(serverCtx)
		ctx = log.WithLoggerContext(ctx, logger)
		ctx = trace.WithTracerContext(ctx, tracer)

		ctx, span := tracer.Start(ctx, info.FullMethod)
		defer span.End()

		startTime := time.Now()
		result, err := handler(ctx, req)
		duration := time.Since(startTime)

		statusCode := codes.Unknown
		if st, ok := status.FromError(err); ok {
			statusCode = st.Code()
		}

		if err != nil {
			log.FromContext(ctx).Error(statusCode.String(), zap.Error(err))
		}

		log.FromContext(ctx).Info("Request",
			zap.String("method", info.FullMethod),
			zap.Any("request", req),
			zap.Duration("duration", duration),
			zap.Int("status_code", int(statusCode)),
			zap.String("status_text", statusCode.String()))

		return result, err
	}
}
