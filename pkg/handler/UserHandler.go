package handler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"strconv"
	"user-service/pkg/common/headers"
	"user-service/pkg/common/requests"
	"user-service/pkg/service"
)

type UserHandler struct {
	E       *echo.Echo
	Service *service.UserService
}

func (handler *UserHandler) InitHandler(e *echo.Echo) {
	e.GET(requests.GetUserByIdPath, handler.getUser)
}

func (handler *UserHandler) getUser(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Request().Header.Get(headers.UserId), 10, 64)
	err, user := handler.Service.SearchUserById(id)
	if err != nil {
		return err
	}
	jsonUser, err := json.Marshal(user)
	return json.NewEncoder(c.Response().Writer).Encode(jsonUser)
}
