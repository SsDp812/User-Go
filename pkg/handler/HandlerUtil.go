package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
)

type Handler interface {
	InitHandler(e *echo.Echo)
}

func InitializeAllHandlers(e *echo.Echo, handlers ...Handler) error {
	if e == nil {
		return errors.New("Error initializing all handlers")
	}
	for _, handler := range handlers {
		handler.InitHandler(e)
	}
	return nil
}
