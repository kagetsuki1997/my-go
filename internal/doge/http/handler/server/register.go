package httpserver

import (
	httphandler "my-go/internal/doge/http/handler"

	"github.com/labstack/echo/v4"
)

func register(e *echo.Group, h httphandler.Handler) {
	doge := e.Group("/doge")
	doge.POST("", h.CreateDoge)
	doge.GET("/:id", h.GetDoge)
	doge.GET("", h.ListDoge)
}
