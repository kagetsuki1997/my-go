package main

import (
	"context"
	"errors"
	"flag"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"my-go/internal/doge/app"
	"my-go/internal/doge/config"
	"my-go/pkg/utils"

	"github.com/chelpis/member-center/pkg/telemetry"
	"github.com/chelpis/member-center/pkg/telemetry/log"
	"go.uber.org/zap"
)

var defaultConfigPath = filepath.Join(utils.GetCurrentDir(), "config", "config.toml")
var configPath = flag.String("config", defaultConfigPath, "path to config file")

func main() {
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	ctx, err := telemetry.Init(ctx, "doge")
	if err != nil {
		panic(err)
	}
	_, logger := log.WithNamedLoggerContext(ctx, "Main")
	config, err := config.Load(*configPath)
	if err != nil {
		logger.Fatal("Error loading config: ", zap.Error(err))
	}
	logger.Info(config.Database.DBName)

	logger.Info("App started")
	if err = app.Run(ctx, config); err != nil {
		logger.Fatal("Error running app: ", zap.Error(err))
	}

	// sync logger
	err = logger.Sync()
	if err != nil && !errors.Is(err, syscall.ENOTTY) {
		logger.Error("Error syncing logger: ", zap.Error(err))
	}
	logger.Info("App stopped")
}
