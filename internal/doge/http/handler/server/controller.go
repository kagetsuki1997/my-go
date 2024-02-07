package httpserver

import (
	"context"
	"my-go/internal/doge/api"
	"my-go/internal/doge/config"

	"github.com/chelpis/member-center/pkg/telemetry/log"
)

const logLabel = "HTTPServer"

func Run(ctx context.Context, config *config.Config, api api.Api) <-chan error {
	ctx, _ = log.WithNamedLoggerContext(ctx, logLabel)

	errChan := make(chan error, 1)
	server := New(ctx, config, api)

	// start server and handle server error
	go func() {
		err := server.Start(ctx)
		if err != nil {
			errChan <- err
		}
	}()

	// handle stop server
	go func() {
		<-ctx.Done()
		server.Stop(ctx)
	}()

	return errChan
}
