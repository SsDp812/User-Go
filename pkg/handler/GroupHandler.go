package handler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"strconv"
	"user-service/pkg/common/headers"
	"user-service/pkg/common/requests"
	"user-service/pkg/service"
)

type GroupHandler struct {
	E       *echo.Echo
	Service *service.GroupService
}

func (handler *GroupHandler) InitHandler(e *echo.Echo) {
	e.GET(requests.GetGroupByIdPath, handler.getGroup)
}

func (handler *GroupHandler) getGroup(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Request().Header.Get(headers.GroupId), 10, 64)
	err, group := handler.Service.SearchGroupById(id)
	if err != nil {
		return err
	}
	jsonGroup, err := json.Marshal(group)
	return json.NewEncoder(c.Response().Writer).Encode(jsonGroup)
}
