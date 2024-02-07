package httpserver

import (
	"context"
	"my-go/internal/doge/api"
	"my-go/internal/doge/config"
	httphandler "my-go/internal/doge/http/handler"
	"my-go/internal/doge/validator"

	httpserver "github.com/chelpis/member-center/pkg/http/server"
	"github.com/labstack/echo/v4"
)

const version = "v1"

type Server = httpserver.Server

func New(ctx context.Context, config *config.Config, api api.Api) *Server {
	server := httpserver.New(ctx, config.HttpServer.Addr)
	server.Handler.HTTPErrorHandler = errorHandler
	server.Handler.Validator = validator.New()

	router := server.Handler.Group("/api/" + version)
	h := httphandler.New(api)
	register(router, h)

	return &server
}

func errorHandler(err error, c echo.Context) {
	httpserver.ErrorHandler(err, c)
}
