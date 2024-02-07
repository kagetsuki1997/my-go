package httphandler

import (
	"my-go/internal/doge/store"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (handler *Handler) CreateDoge(c echo.Context) error {
	req, err := bindAndValidate[CreateDogeRequest](c)
	if err != nil {
		return ErrBadRequest.SetInternal(err)
	}

	resp, err := handler.api.CreateDoge(store.CreateDogeParams{Name: req.Name, MagicNumber: req.MagicNumber, Type: req.Type})
	if err != nil {
		return ErrInternal.SetInternal(err)
	}

	return c.JSON(http.StatusCreated, resp)
}

func (handler *Handler) GetDoge(c echo.Context) error {
	req, err := bindAndValidate[GetDogeRequest](c)
	if err != nil {
		return ErrBadRequest.SetInternal(err)
	}

	resp, err := handler.api.GetDoge(req.Id)
	if err != nil {
		if err.Error() == "record not found" {
			return ErrNotFound.SetInternal(err)
		}

		return ErrInternal.SetInternal(err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (handler *Handler) ListDoge(c echo.Context) error {
	resp, err := handler.api.ListDoge()
	if err != nil {
		return ErrInternal.SetInternal(err)
	}

	return c.JSON(http.StatusOK, resp)
}
