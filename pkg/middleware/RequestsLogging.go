package middleware_user_service

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func LogRequest(next echo.HandlerFunc) echo.HandlerFunc {
	log.Debug("Request caught!")
	return next
}
