package app

import (
	"context"
	"my-go/internal/doge/api"
	"my-go/internal/doge/config"
	"my-go/internal/doge/grpc"
	httpserver "my-go/internal/doge/http/handler/server"
	"my-go/internal/doge/store"

	"github.com/chelpis/member-center/pkg/db"
	"github.com/chelpis/member-center/pkg/telemetry/log"
	"go.uber.org/zap"
)

const logLabel = "App"

func Run(ctx context.Context, config *config.Config) (err error) {
	ctx, logger := log.WithNamedLoggerContext(ctx, logLabel)

	logger.Info("Init database connection")
	database, err := db.New(ctx, config.Database)
	if err != nil {
		logger.Error("Error conneting to database", zap.Error(err))
		return
	}
	defer db.Close(database)

	if err = db.AutoMigrate(database, config.Database); err != nil {
		logger.Error("Error migrating database", zap.Error(err))
		return
	}
	logger.Info("Database migrated")

	store := store.New(database)
	api := api.New(store)

	logger.Info("Start HTTP server")
	httpErrChan := httpserver.Run(ctx, config, api)

	logger.Info("Start gRPC server")
	grpcErrChan := grpc.Run(ctx, config, api)

	// handle shutdown
	select {
	case err = <-httpErrChan:
	case err = <-grpcErrChan:
	case <-ctx.Done():
	}

	return
}
