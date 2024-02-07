package httphandler

import (
	"my-go/internal/doge/api"

	"github.com/labstack/echo/v4"
)

var (
	ErrBadRequest = echo.ErrBadRequest
	ErrNotFound   = echo.ErrNotFound
	ErrInternal   = echo.ErrInternalServerError
)

type Handler struct {
	api api.Api
}

func New(api api.Api) Handler {
	return Handler{api}
}

func bindAndValidate[T any](c echo.Context) (*T, error) {
	t := new(T)
	if err := c.Bind(t); err != nil {
		return nil, err
	}
	if err := c.Validate(t); err != nil {
		return nil, err
	}

	return t, nil
}
